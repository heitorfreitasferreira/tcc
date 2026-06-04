# Análise do Visualizador Web de Experimentos

## 1. Visão Geral

O módulo `src/web/` implementa um **visualizador interativo** que permite navegar e
inspecionar os resultados dos experimentos de otimização (TSP/rTSP) sem depender de
ferramentas externas como Jupyter ou planilhas. Ele é servido por um binário Go
auto-contido (`./src/tcc serve --addr :8080`) que embarca todos os dados e assets
via `embed`.

### O que o visualizador oferece

| Funcionalidade | Descrição |
|---|---|
| Navegação hierárquica | Árvore colapsável: **Mapa → Método → Run** (sidebar) |
| Renderização do grafo | Pontos 2D em Canvas com grid, eixos, arestas direcionadas |
| Animação da evolução | Playback iteracão-a-iteracão da melhor rota ao longo do tempo |
| Animação da melhor rota | Percurso suave sobre a melhor sequência encontrada |
| Controle de velocidade | 0.5×, 1×, 2×, 4× |
| Deep linking | URL reflete o estado: `/?map=10a&method=aco&run=...` |

---

## 2. Análise da Arquitetura

### 2.1 Camadas do Backend

```
handlers/page.go         ← HTTP + HTMX endpoints
    ↓
service/page.go          ← Lógica de seleção, montagem do PageData
    ↓
repository/              ← Acesso a dados embarcados (fs.FS)
  ├── types.go            ← Interfaces: MapsReader, SummaryReader, EvolutionReader
  ├── data.go             ← EmbeddedDataRepository (wrapper sobre fs.FS)
  ├── maps.go             ← ListMaps, LoadPoints, LoadGraph
  ├── summary.go          ← ListRuns (lê results/summary/*.json)
  └── evolution.go        ← LoadEvolution (lê results/evolution/*.jsonl)
```

**Separação de responsabilidades clara.** Cada camada é testável isoladamente
(9 testes nos subpacotes). O repository usa interfaces (`MapsReader`, `SummaryReader`,
`EvolutionReader`) que permitem stub fácil nos testes.

**Observação crítica:** O `PageService` recebe *três* interfaces separadas em vez
de uma única interface de repositório. Isso é um design acertado — permite mockar
apenas o necessário em cada cenário de teste. Contudo, a injeção via construtor
com três parâmetros do mesmo tipo concreto (`*EmbeddedDataRepository`) é ruidosa.
Uma fábrica ou um struct de configuração reduziria o acoplamento se houver
necessidade de fontes de dados diferentes (ex: rede vs. embed).

### 2.2 Frontend (ES Modules sem Bundler)

```
static/js/
├── index.js              ← Entry point
├── constants.js          ← Configurações compartilhadas
├── app/
│   └── main-content.js   ← Orquestrador: parse → controller → bind HTMX
├── data/
│   └── parsers.js        ← Parsing seguro de JSON (points + evolution)
├── domain/
│   ├── evolution.js      ← Lógica de melhor frame, labels, formatação
│   └── route.js          ← Construção de rota, interpolação entre pontos
├── player/
│   ├── evolution-player.js       ← Timed iteration playback (setTimeout)
│   ├── best-sequence-player.js   ← Smooth traversal (requestAnimationFrame)
│   └── run-controller.js         ← Orquestrador dos dois modos
├── render/
│   ├── geometry.js       ← Sistema de coordenadas: mundo [-1,1] → canvas
│   └── map-renderer.js   ← Grid, eixos, setas, pontos, marcador
└── ui/
    ├── controls.js       ← Conexão botões ↔ controller
    └── metadata.js       ← Atualização dos elementos de métrica
```

**Análise:** A organização por domínio (domain/), rendering (render/), playback
(player/), e UI (ui/) segue boas práticas de arquitetura frontend. Cada módulo
tem responsabilidade única e é pequeno (maior: `map-renderer.js` com 232 linhas).

**Observações:**

- `domain/route.js` contém lógica pura (sem DOM, sem Canvas) — ideal para
  testes unitários que hoje não existem. Adicionar testes para
  `buildRouteIndices` e `interpolateRoutePoint` aumentaria a confiabilidade.

- `run-controller.js` atua como *state machine* entre modos "evolution" e
  "best-sequence". A troca de modo destrói e recria o player ativo. Isso
  simplifica o código mas significa que o estado de pause é mantido apenas
  no controller, não nos players individuais — uma escolha correta de design.

### 2.3 Navegação com HTMX

O visualizador usa HTMX 2.0.7 para navegação parcial sem JavaScript de
roteamento. Três endpoints REST-leves:

| Endpoint | Retorno | Comportamento |
|---|---|---|
| `/ui/select-map` | `main_with_sidebar_oob` | Troca mapa, expande/colapsa na sidebar (OOB swap) |
| `/ui/select-method` | `main_with_sidebar_oob` | Troca método, toggle expansão |
| `/ui/select-run` | `main_with_sidebar_oob` | Seleciona run, carrega evolução |

O template `main_with_sidebar_oob` renderiza tanto o conteúdo principal quanto
a sidebar com `hx-swap-oob="innerHTML"`, permitindo que uma única resposta
atualize dois elementos da página.

**Análise:** A escolha de HTMX sobre um SPA (React, Vue) é adequada para o
escopo — o visualizador tem três ações de usuário (selecionar mapa, método,
run) e não requer estado client-side complexo. Isso eliminou a necessidade de
bundler JS, roteamento client-side, e gerenciamento de estado.

---

## 3. Análise das Visualizações

### 3.1 Renderização do Mapa

O `geometry.js` mapeia coordenadas do mundo (normalizadas em `[-1, 1]`) para
pixels do canvas com padding uniforme. A escala é automática (usa o menor
entre largura e altura útil).

**O que é desenhado:**
- Grid quadriculado (passo 0.25) em cinza claro
- Eixos X e Y em cinza escuro, destacados
- Moldura da área de plotagem
- Pontos: vértice 0 em vermelho (origem/depósito), demais em azul, com rótulo
  numérico
- Arestas: setas laranja com pontas, jitter lateral para arestas paralelas

**Diferenciais técnicos:**
- **Jitter lateral em arestas paralelas:** Se o otimizador produzir uma rota
  que percorre a mesma aresta ida-e-volta, as setas são deslocadas lateralmente
  para evitar sobreposição. O deslocamento acumula: `base + sinal * lane * step`,
  onde `lane` é `ceil(ocorrência / 2)`. Isso produz um feixe de setas
  paralelas visíveis, em vez de uma única linha borrada.
- **Setas direcionais:** A cabeça da seta é desenhada como um triângulo
  preenchido, com tamanho proporcional ao comprimento da aresta (mín. 8px,
  máx. 12px). O shaft recua `pontoRaio + gap` da ponta para não sobrepor
  o vértice.

### 3.2 Animação da Evolução (Evolution Player)

O `evolution-player.js` implementa playback iterativo usando `setTimeout`:

```
delay = max(16ms, 100ms / speedMultiplier)
```

Para cada iteração do mundo (1, 2, 3, ..., N), ele localiza o frame de
evolução mais recente ≤ iteração atual via busca sequencial (`findFrameIndexForIteration`). Como o número de frames é pequeno (uma fração das iterações
totais), a complexidade O(n) por tick não é problema.

**Estado do ciclo:** Quando a iteração atinge `cycleEndIteration` (max entre
último frame e total de iterações configuradas), o cursor volta para
`cycleStartIteration` (iteração do primeiro frame). Isso cria um loop
contínuo.

**Observação:** O player lida com o caso de `frames.length === 0` (brute-force
não gera evolution) — simplesmente desenha o mapa sem rota e exibe métricas
vazias. Isso é importante para não quebrar a UI quando dados parciais existem.

### 3.3 Animação da Best Sequence

O `best-sequence-player.js` usa `requestAnimationFrame` para animação suave
do percurso. A lógica:

```
segmentProgress += (elapsed * speed) / BEST_SEQUENCE_EDGE_DURATION_MS
```

Quando `segmentProgress >= 1`, avança para o próximo segmento da rota (com
wrap-around). O marcador de percurso é interpolado linearmente entre os
pontos inicial e final do segmento atual via `interpolateRoutePoint()`.

**Construção da rota:** `buildRouteIndices` em `domain/route.js`:
1. Normaliza os índices da sequência (filtra valores inválidos).
2. Garante que o vértice 0 (depósito) está presente (insere no início se
   ausente).
3. Garante que a rota termina no vértice inicial (fecha o ciclo).

**Observação:** Se a melhor sequência encontrada pelo otimizador já termina
no vértice 0 (como esperado), o passo 3 adiciona uma aresta extra idêntica
à última. Isso resulta em uma duplicata visual — o marcador percorre a
aresta final duas vezes. Isso poderia ser evitado com uma verificação
`route[last] !== route[0]` antes de adicionar o fechamento.

### 3.4 Navegação com Estado nos Templates

A sidebar (`sidebar.html`) implementa uma árvore de três níveis usando
condicionais Go template:

```
Nível 1: Mapas
  └── se expandido → Nível 2: Métodos
        └── se expandido → Nível 3: Runs
```

Cada link modifica o estado via `hx-get` com parâmetros `current_map`,
`expanded_map`, `current_method`, `expanded_method`. A lógica de toggle
(resolve colapsar se já expandido, expandir se fechado, trocar se outro)
está no `PageService.ResolveMapSelection` / `ResolveMethodSelection`.

**Observação interessante:** O handler `SelectRun` não usa toggle — selecionar
uma run é uma ação definitiva que também expande o método pai. Isso é
consistente com a expectativa do usuário: ao clicar em uma run, quer vê-la,
não togglear.

---

## 4. Cobertura e Lacunas

### 4.1 O que o visualizador cobre bem

- **Inspeção qualitativa de rotas:** Permite ver como cada método converge
  para a solução, visualmente.
- **Comparação entre métodos:** A navegação rápida entre runs de ACO, PSO,
  GA e Bruteforce permite inspeção comparativa.
- **Reprodutibilidade:** Qualquer run pode ser acessada por URL direta
  (`/?map=30a&method=aco&run=30a__aco__s42__h<hash>`).

### 4.2 Lacunas identificadas

1. **Sem sobreposição de métodos:** Não é possível sobrepor duas rotas de
   métodos diferentes no mesmo canvas para comparação lado a lado.

2. **Sem métricas agregadas:** O visualizador mostra dados de uma run por vez.
   Não há visão agregada (média/mediana por método, box-plots, etc.) — isso
   fica para o notebook `scripts/visualizacoes.ipynb`.

3. **Sem dados de timing:** O diretório `results/timing/` é embarcado mas não
   exposto na UI. Tempo de execução por método não é visível.

4. **Sem destaque de arestas da solução ótima:** Quando disponível (instâncias
   com brute-force), não há comparação visual vs. solução ótima.

5. **Sem informação de instância:** O número de nós, densidade do grafo, e
   parâmetros do método não são exibidos.

---

## 5. Relação com a Monografia

### 5.1 Como o visualizador suporta a análise

| Tipo de análise | Como o visualizador contribui |
|---|---|
| Qualitativa | Screenshots da rota em diferentes iterações para figuras na monografia |
| Convergência | Animação mostra como a rota melhora ao longo das iterações |
| Melhor rota | Modo best-sequence permite inspecionar a qualidade topológica da rota |
| Reprodutibilidade | URLs diretas para cada run permitem verificação independente |

### 5.2 Possíveis figuras extraídas

1. **Exemplo de evolução:** Captura de 3-4 frames (iteração 1, 10, 50, final)
   para uma run de ACO em `20a`, mostrando melhoria progressiva da rota.

2. **Comparação de rotas finais:** Captura da melhor rota de ACO vs. GA para
   a mesma instância, lado a lado.

3. **Problema de arestas paralelas:** Captura mostrando o jitter lateral em
   uma rota que cruza a mesma aresta múltiplas vezes (evidencia comportamento
   do otimizador).

4. **Mapa com grid e eixos:** Captura do grafo renderizado com grid, mostrando
   a distribuição espacial dos POIs.

### 5.3 Inclusão na metodologia

Sugestão de redação para a seção de metodologia/implementação:

> Para permitir a inspeção visual qualitativa dos resultados, implementou-se
> um servidor web interativo embarcado no binário de linha de comando. A
> interface organiza os experimentos em uma árvore hierárquica de três níveis
> (instância, método de otimização, execução), permitindo ao usuário navegar
> entre as mais de 4.500 execuções e visualizar a evolução da rota ao longo
> das iterações. Um controle de velocidade (0,5× a 4×) e a alternância entre
> os modos "evolução" e "melhor rota" facilitam a inspeção tanto do processo
> de convergência quanto da solução final. O frontend utiliza Canvas API para
> renderização do grafo com arestas direcionadas, HTMX para navegação parcial
> sem recarregamento, e módulos ES nativos sem bundler.

---

## 6. Conclusões da Análise

O visualizador web cumpre seu objetivo de fornecer uma ferramenta de inspeção
visual para os experimentos. A separação em camadas no backend e a organização
modular no frontend facilitam manutenção e extensão. As principais virtudes
são a portabilidade (binário único), a navegação fluida (HTMX), e a qualidade
da renderização (jitter lateral em arestas paralelas, setas direcionais,
interpolação suave).

Para a monografia, o visualizador é mais útil como ferramenta de *inspeção
qualitativa* e geração de figuras, enquanto a *análise quantitativa* (tabelas
comparativas, testes estatísticos) permanece no notebook Jupyter. As lacunas
identificadas (sobreposição de métodos, métricas agregadas, timing) não
comprometem o uso atual mas seriam extensões naturais para trabalho futuro.
