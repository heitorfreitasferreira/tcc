---
title: Glossário Terminológico da Monografia
tags:
  - writing
  - monografia
  - glossario
  - terminologia
status: atualizado-pos-p7
created: 2026-06-02
updated: 2026-06-02
---

# Glossário Terminológico da Monografia

Este glossário padroniza os termos usados na monografia. Sempre que houver alternativa (ex.: drone vs VANT), usar o termo marcado como **preferencial**. Os campos seguem o esquema:

- **Termo preferencial**: o que deve aparecer no texto
- **Sinônimos/evitar**: variantes encontradas na literatura ou no código interno
- **Definição**: descrição precisa
- **Uso no código**: como aparece em `src/` quando aplicável
- **Notas**: observações para o escritor

> [!warning] Regra de consistência
> Em caso de conflito entre este glossário e uma nota antiga do `vault/`, conferir [[auditoria-codigo-dados-vault]] e o código em `src/`. Este glossário foi atualizado após a auditoria P7.

---

## Problema e Formulação

### TSP-SD-ATP
- **Termo preferencial**: TSP-SD-ATP
- **Sinônimos/evitar**: rTSP (apelido interno, evitar na monografia)
- **Definição**: *Traveling Salesman Problem with Sequence-Dependent Angular Turn Penalties* — variante do TSP onde o custo de transição entre dois nós depende do nó anterior, devido à penalidade angular da curva.
- **Uso no código**: `tcc` (nome do módulo Go), `rTSP` em nomes de arquivos e scripts
- **Notas**: Definir na primeira ocorrência; usar tradução explicativa "TSP com penalidades angulares dependentes da sequência" em texto corrido. Não apresentar como variante consolidada na literatura; escrever como formulação adotada neste trabalho.

### Makespan
- **Termo preferencial**: makespan
- **Sinônimos/evitar**: custo total, tempo total de rota, *fitness*
- **Definição**: tempo total para o drone percorrer a rota completa, partindo da base (nó 0), visitando todos os POIs na sequência determinada, e retornando à base. Soma dos valores do tensor 3D nos passos da rota.
- **Uso no código**: `g.Makespan(order)`, `BestMakespan` em `shared.OptimizationResult`
- **Notas**: Métrica objetivo única do problema. Definir na primeira ocorrência como tempo/custo total da rota; depois usar `makespan`. Evitar chamar de *fitness* na monografia, pois *fitness* é termo interno de metaheurística e pode inverter a intuição de minimização.

### Tensor 3D de Custos
- **Termo preferencial**: tensor 3D de custos
- **Sinônimos/evitar**: matriz de adjacência 3D, grafo 3D, *cost tensor*
- **Definição**: estrutura `G[anterior][atual][próximo]` que armazena o custo de transição (distância + penalidade angular) para cada tripleta de nós consecutivos. Pré-computado em O(n³).
- **Uso no código**: `type Graph [][][]float64` em `src/graph/types.go`
- **Notas**: Diferencia este trabalho do TSP clássico (matriz 2D)

### Penalidade Angular
- **Termo preferencial**: penalidade angular
- **Sinônimos/evitar**: custo de curva, *turn cost*, *angular penalty*
- **Definição**: tempo adicional proporcional ao ângulo entre o vetor de chegada e o vetor de partida. Calculada como `maxPenalty × (θ/π)`, com `maxPenalty = 1`, θ em radianos.
- **Uso no código**: `turnCost()` em `src/graph/math.go`, `maxPenalti` em `src/graph/types.go`
- **Notas**: Normalizada em [0,1]; adicionada à distância euclidiana

### Ponto de Interesse
- **Termo preferencial**: ponto de interesse ou POI
- **Sinônimos/evitar**: nó, vértice, cliente, cidade
- **Definição**: localização geográfica 2D que o drone deve visitar. Nós 1 a n-1 do grafo.
- **Uso no código**: `Points2D` em `src/points/types.go`; `pts[i]` no intervalo `[1, n-1]`
- **Notas**: Usar "ponto de interesse" no texto corrido e `POI` em tabelas, figuras ou quando a repetição ficar pesada. Usar "nó" apenas ao discutir grafo/índices matemáticos.

### Base
- **Termo preferencial**: base
- **Sinônimos/evitar**: origem, depósito, nó 0, *depot*
- **Definição**: ponto de partida e retorno obrigatório do drone. Nó 0 do grafo.
- **Uso no código**: nó 0 implícito em `g.Makespan(order)` — não aparece na permutação
- **Notas**: A permutação dos métodos contém apenas os POIs (1..n-1); o Makespan adiciona a partida e o retorno

### Drone
- **Termo preferencial**: drone
- **Sinônimos/evitar**: VANT, UAV, veículo aéreo não tripulado
- **Definição**: veículo aéreo não tripulado que realiza a missão de patrulha. Velocidade constante = 1 unidade de distância por unidade de tempo.
- **Uso no código**: `droneSpeed metersPerSecond = 1` em `src/graph/types.go`
- **Notas**: Usar "drone" de forma consistente. Se o orientador exigir formalidade, definir na primeira ocorrência como "drone (veículo aéreo não tripulado)"; evitar alternar com VANT/UAV ao longo do texto.

### Instância
- **Termo preferencial**: instância
- **Sinônimos/evitar**: mapa, cenário, problema
- **Definição**: conjunto de pontos 2D em [-1,1)² que define um problema específico. Nomeada como `{n}{letra}` (ex.: `10a`, `100c`).
- **Uso no código**: arquivos `.points` e `.graph` em `src/data/`
- **Notas**: 30 instâncias no total (10 tamanhos × 3 variantes)

---

## Métodos de Otimização

### Algoritmo Genético (GA)
- **Termo preferencial**: GA ou Algoritmo Genético
- **Sinônimos/evitar**: Genetic Algorithm
- **Definição**: metaheurística populacional inspirada na seleção natural. Implementação com permutação de POIs, crossover OX, mutação swap, seleção por torneio e elitismo. Parâmetros: pop=100, iter=100, elitismo=1, mutação=0.05, torneio=2.
- **Uso no código**: `src/optimization/ga/`
- **Notas**: A representação por permutação é a codificação canônica para TSP

### Particle Swarm Optimization (PSO)
- **Termo preferencial**: PSO ou Otimização por Enxame de Partículas
- **Sinônimos/evitar**: Particle Swarm Optimization
- **Definição**: metaheurística populacional inspirada no comportamento social de aves/peixes. Implementação com random keys (codificação contínua → permutação via ordenação), atualização de velocidade com inércia, componentes cognitivo e social. Parâmetros: pop=100, iter=100, c1=c2=2.0, w=0.7.
- **Uso no código**: `src/optimization/pso/`
- **Notas**: Random keys perdem informação de adjacência — limitação estrutural conhecida. Sem *velocity clamping*.

### Ant Colony Optimization (ACO)
- **Termo preferencial**: ACO ou Otimização por Colônia de Formigas
- **Sinônimos/evitar**: Ant Colony Optimization, Ant System
- **Definição**: metaheurística populacional inspirada no forrageamento de formigas. Implementação com feromônio 3D (τ[prev][curr][next]), heurística η = 1/G, transição por roleta proporcional, depósito de todas as formigas (Ant System adaptado). Parâmetros: pop=100, iter=100, α=1.0, β=2.0, ρ=0.2, Q=100.
- **Uso no código**: `src/optimization/aco/`
- **Notas**: Feromônio 3D é consequência estrutural do tensor 3D; diluição ~98× vs ACO 2D para n=100

### Busca Exaustiva
- **Termo preferencial**: busca exaustiva
- **Sinônimos/evitar**: brute-force, enumeração completa, *exhaustive search*
- **Definição**: algoritmo determinístico que gera todas as (n-1)! permutações de POIs via algoritmo de Heap e retorna a rota de menor makespan. Baseline exata para instâncias pequenas executadas (`10a..15c`).
- **Uso no código**: `src/optimization/brute/`
- **Notas**: Complexidade O(n!·n); inviável para instâncias maiores. Serve como ótimo de referência apenas nas 18 instâncias com summary brute-force disponível, conforme [[auditoria-codigo-dados-vault]].

### Lower Bound (AP)
- **Termo preferencial**: lower bound ou limitante inferior
- **Sinônimos/evitar**: cota inferior, *assignment problem bound*
- **Definição**: limitante inferior para o makespan ótimo, obtido via relaxação do problema de designação (*Assignment Problem*) sobre matriz 2D reduzida do tensor 3D: `c'[j][k] = min_i G[i][j][k]`. Resolvido pelo algoritmo Hungarian O(n³). O bound é válido (≤ makespan ótimo) mas não produz rota factível (pode conter subtours).
- **Uso no código**: `src/optimization/lowerbound/`
- **Notas**: Gap médio 51.36% (40.79–65.35%) vs brute-force nas 18 instâncias auditadas. Válido mas frouxo para o TSP-SD-ATP. Útil como referência adicional para instâncias grandes onde brute-force é inviável, mas não deve ser descrito como ótimo ou rota.

### Rota Factível
- **Termo preferencial**: rota factível
- **Sinônimos/evitar**: solução válida, tour factível
- **Definição**: sequência que parte da base, visita cada POI exatamente uma vez e retorna à base. GA, PSO, ACO e busca exaustiva retornam rotas factíveis; o lower bound AP retorna um valor de bound e pode conter subtours.
- **Uso no código**: `BestSequence` em `shared.OptimizationResult`; validação em `src/graph/makespan.go`
- **Notas**: Usar esta distinção ao comparar lower bound com métodos de otimização.

---

## Parâmetros e Componentes dos Métodos

### Elitismo (GA)
- **Definição**: estratégia que preserva os melhores indivíduos de cada geração sem alteração. Default: 1.
- **Uso no código**: `Elitism int` em `ga.Params`

### Crossover OX
- **Termo preferencial**: crossover OX ou *Ordered Crossover*
- **Definição**: operador de recombinação para permutações que preserva a ordem relativa dos elementos. Seleciona dois pontos de corte, copia o segmento de um progenitor e preenche as posições restantes na ordem do outro progenitor.
- **Uso no código**: implementado em `src/optimization/ga/main.go:141-194`

### Mutação Swap
- **Definição**: operador que troca dois genes (POIs) de posição na permutação com probabilidade `mutationRate`. Default: 0.05.
- **Uso no código**: implementado em `src/optimization/ga/main.go:196-201`

### Seleção por Torneio
- **Definição**: método de seleção onde k indivíduos são amostrados aleatoriamente e o melhor (menor makespan) é selecionado para reprodução. Default: k=2.
- **Uso no código**: `src/optimization/ga/main.go:130-139`

### Random Keys
- **Termo preferencial**: random keys
- **Definição**: codificação contínua [0,1]^d para problemas de permutação. A ordem dos valores, quando ordenados, define a permutação. Usada pelo PSO para mapear posição contínua → rota.
- **Uso no código**: `src/optimization/pso/particle.go:23-40`
- **Notas**: Decodificação O(n log n). Não preserva adjacências explicitamente; tratar como possível fator no desempenho inferior do PSO, não como causa demonstrada isoladamente.

### Feromônio 3D (ACO)
- **Definição**: matriz τ[prev][curr][next] que representa a atratividade de visitar `next` a partir de `curr` tendo vindo de `prev`. Dimensão n×n×n.
- **Uso no código**: `pheromones` em `src/optimization/aco/main.go:21`
- **Notas**: Estruturalmente necessário para o TSP-SD-ATP; diluição é inerente, não bug

### Heurística ACO (η)
- **Definição**: informação local de custo usada na transição probabilística. η = 1/G[prev][curr][next]. Combinada com feromônio via τ^α · η^β.
- **Uso no código**: `eta` em `src/optimization/aco/ant.go:52-55`

### Evaporação (ρ)
- **Definição**: taxa de decaimento do feromônio a cada iteração. τ ← τ·(1−ρ). Default: 0.2.
- **Uso no código**: `Rho` em `aco.Params` (flag `--rho`)

### Depósito (Q)
- **Definição**: constante escalar que controla a quantidade de feromônio depositada por cada formiga. Default: 100.
- **Uso no código**: `Q` em `aco.Params` (flag `--q`)

---

## Pipeline e Experimentos

### Run ID
- **Definição**: identificador único de uma execução experimental. Formato: `{instancia}__{metodo}__s{seed}__h{hash}`. Ex.: `10a__ga__s0__hd186dc51`.
- **Uso no código**: `reporting.BuildRunID()` em `src/shared/reporting/reporting.go`
- **Notas**: O hash SHA1 previne colisões. A run ID nomeia os três arquivos de saída.

### Nome-base da Instância
- **Termo preferencial**: nome-base da instância
- **Sinônimos/evitar**: instância normalizada, ID da instância
- **Definição**: identificador curto extraído do arquivo de instância, como `10a`, `30b` ou `100c`.
- **Uso no código/dados**: aparece no `run_id`; o campo `instance` dos summaries pode conter caminho absoluto ou relativo.
- **Notas**: Em análises e tabelas, agrupar sempre pelo nome-base para evitar separar artificialmente caminhos absolutos e relativos.

### Summary
- **Definição**: arquivo JSON com o resultado final de uma execução (schema `tcc.summary.v1`). Contém: método, instância, seed, parâmetros, makespan final, sequência, arquivos de evolução/timing associados.
- **Uso no código**: `shared.ReportRunSummary` em `src/shared/reporting/reporting.go`

### Evolution
- **Definição**: arquivo JSONL (uma linha por iteração) com o histórico de evolução (schema `tcc.evolution.v1`). Cada linha contém: iteração, makespan, delta, sequência.
- **Uso no código**: `shared.EvolutionRecord` em `src/shared/reporting/reporting.go`

### Timing
- **Definição**: arquivo JSON com a decomposição temporal da execução (schema `tcc.timing.v1`). Campos: load_instance, optimize, serialize, total (em ms).
- **Uso no código**: `shared.RunTiming` em `src/shared/reporting/reporting.go`

### Seed / Semente
- **Definição**: valor inteiro (int64) que inicializa o gerador de números aleatórios. Usado para reprodutibilidade. 51 sementes (s0..s50) para métodos estocásticos.
- **Uso no código**: flag `--seed` / `-s` no CLI
- **Notas**: Métodos determinísticos (brute-force, lower bound) ignoram a seed ou usam s0

### Gap
- **Definição**: diferença relativa entre o makespan de uma solução e uma referência (ótimo ou bound). Calculado como `(makespan − referência) / referência × 100%`.
- **Notas**: Pode ser gap vs brute-force nas 18 instâncias com ótimo (`10a..15c`) ou gap vs lower bound AP para todas as instâncias. Sempre declarar a referência usada.

### Taxa de Acerto
- **Termo preferencial**: taxa de acerto
- **Sinônimos/evitar**: *success rate*, proporção de ótimos encontrados
- **Definição**: proporção de execuções (entre as 51 sementes) em que o método encontra o makespan ótimo (para instâncias com brute-force disponível).

---

## Análise Estatística

### Teste de Friedman
- **Definição**: teste não-paramétrico equivalente à ANOVA de duas vias por postos. Usado para rejeitar H₀: "todos os métodos têm desempenho equivalente". Não assume normalidade.
- **Notas**: Seguir Demšar (2006). Aplicar sobre as medianas por instância. O script atual foi auditado em [[auditoria-script-analise-estatistica]] e está em correção por outro agente; não usar claims finais até essa correção estar concluída.

### Teste de Nemenyi
- **Definição**: teste post-hoc para comparações múltiplas após Friedman. Calcula diferença crítica (CD) entre postos médios. Dois métodos diferem com significância se a distância entre postos excede o CD.
- **Notas**: Visualizado por diagrama CD

### Teste de Wilcoxon
- **Definição**: teste não-paramétrico pareado por instância entre dois métodos. Alternativa ao Nemenyi para confirmação.
- **Notas**: Usar com correção de Bonferroni-Holm se múltiplos pares

### Significância Estatística
- **Definição**: conclusão de que a diferença observada entre métodos não é atribuível ao acaso (p < 0.05). Deve ser baseada nos testes acima, não apenas na diferença de médias.
- **Notas**: Não declarar sem executar o protocolo completo corrigido. Enquanto isso, usar "diferença descritiva" ou "ordenação observada".

---

## Convenções de Escrita

| Conceito | Regra |
|----------|-------|
| Nomes de métodos em inglês | GA, PSO, ACO (siglas maiúsculas sem pontos) |
| Nomes de métodos em português | Algoritmo Genético, Otimização por Enxame de Partículas, Otimização por Colônia de Formigas |
| Primeira ocorrência | Sempre escrever por extenso com sigla entre parênteses: "Otimização por Colônia de Formigas (ACO)" |
| Tabelas e legendas | Usar siglas (GA, PSO, ACO, BF, LB) |
| Números decimais | Ponto como separador decimal (padrão científico internacional) |
| Milissegundos | ms |
| Porcentagem | % com um espaço antes (exceto em tabelas) |
| Drones | "drone" |
| Brute-force | "busca exaustiva" no texto; `BF` em tabelas |
| Lower bound | "limitante inferior" na primeira ocorrência; `LB` em tabelas |
| Significância | Usar com os valores validados em [[analysis-methodology]]; delimitar às configurações avaliadas |

---

## Relação com o Roadmap

Este glossário serve como léxico de referência para todos os capítulos. Ver [[roadmap-monografia]] para a terminologia oficial e validação cruzada com claims.
