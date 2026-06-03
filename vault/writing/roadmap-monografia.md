---
title: Roadmap de Escrita da Monografia
tags:
  - writing
  - monografia
  - roadmap
  - agentes
status: pronto-para-execucao
created: 2026-06-02
---

# Roadmap de Escrita da Monografia

Esta nota é o ponto de partida operacional para agentes que irão escrever a monografia. A escrita deve usar o `vault/` como camada organizada de conhecimento e a estrutura de diretórios da `monografia/` como destino, sem assumir conteúdo já escrito nos arquivos `.tex`.

## Hierarquia de Informação

A fonte determinística de informação do projeto é sempre o código em `src/` e os dados brutos gerados por ele em `src/data/`. O `vault/` não é fonte primária: ele é uma camada de organização mantida por agentes a comando do autor. Portanto, antes de escrever qualquer afirmação técnica, metodológica ou experimental, o agente deve conferir se ela ainda é refletida pelo código e pelos dados.

Fluxo obrigatório de informação:

```text
src/                 → define implementação, algoritmos, flags, schemas e renderizações
src/data/            → contém instâncias, grafos e resultados brutos gerados pelo código
leitura código+dados → valida o que foi realmente implementado e medido
literatura           → contextualiza, fundamenta e compara com trabalhos externos
vault/               → organiza síntese, claims, referências e plano de escrita
monografia/          → recebe o texto final, figuras, tabelas e citações
```

Ordem de autoridade quando houver conflito:

| Ordem | Fonte | Papel |
|---:|---|---|
| 1 | `src/` | Verdade sobre implementação, parâmetros disponíveis, algoritmos, schemas e endpoints |
| 2 | `src/data/` | Verdade sobre instâncias, resultados brutos, tempos, evolução e artefatos gerados |
| 3 | Literatura | Verdade externa para conceitos, métodos, lacunas e trabalhos relacionados |
| 4 | `vault/` | Organização e síntese produzida por agentes; deve ser revalidada contra código/dados |
| 5 | `monografia/` | Documento final; deve herdar informação validada das camadas anteriores |

> [!warning] Regra de validação
> Se uma nota do `vault/` contradizer o código ou os dados brutos, o código e os dados vencem. O agente deve atualizar a nota ou marcar a inconsistência antes de usar essa informação na monografia.

## Regra Central

Escrever em português acadêmico, com foco no cenário de patrulha com drones e no problema TSP-SD-ATP. Toda afirmação técnica ou experimental forte deve estar ancorada primeiro em `src/` e `src/data/`; o `vault/` serve para localizar e organizar essa evidência.

| Tipo de afirmação | Fonte obrigatória |
|---|---|
| Conceito clássico de TSP, NP-dificuldade, variantes | Literatura organizada em `vault/areas/` e `vault/papers/` |
| GA, PSO, ACO, busca exaustiva e lower bound AP como implementados | Código em `src/optimization/`, `src/graph/`, `src/shared/`; síntese em `vault/projeto/` |
| Formulação TSP-SD-ATP | Código em `src/graph/` e `src/points/`; síntese em [[problem-formulation]] |
| Arquitetura e pipeline | Código em `src/cmd/`, scripts `src/run_*.sh`, schemas em `src/shared/`; síntese em [[architecture]], [[experiment-pipeline]] |
| Resultados e conclusões empíricas | Dados brutos em `src/data/results/`; síntese em [[resultados]], [[analysis-methodology]] |
| Figuras e tabelas | Dados em `src/data/`, scripts em `scripts/`, renderização em `src/web/`, saídas em `monografia/figs/` |

> [!warning] Claims fortes
> Não escrever como fato absoluto que este é o "primeiro" estudo publicado sem formulação defensável. Preferir: "não foram identificados, na revisão realizada, estudos que comparem sistematicamente GA, PSO e ACO na variante TSP-SD-ATP".

## Ordem Recomendada de Escrita

| Ordem | Capítulo | Por que escrever nesta posição | Nota-base |
|---:|---|---|---|
| 1 | Fundamentação Teórica | Define vocabulário, citações e limites conceituais | [[fundamentacao]] |
| 2 | Proposta | Descreve a formulação e a implementação, usando vocabulário já estabilizado | [[proposta]] |
| 3 | Experimentos | Depende da proposta e precisa de números/figuras finais | [[experimentos]] |
| 4 | Introdução | Deve refletir exatamente o que a monografia entrega | [[introducao]] |
| 5 | Conclusão | Só deve fechar depois dos resultados e limitações estarem estabilizados | [[conclusao]] |

## Entregáveis Globais Antes da Escrita Final

| Entregável | Status esperado | Fonte |
|---|---|---|
| Lista de claims centrais | Criada antes de Introdução e Conclusão | [[resultados]], [[analysis-methodology]] |
| Tabela de referências por capítulo | Criada antes de Fundamentação | `vault/papers/index.md` |
| Padronização terminológica | Criada antes de qualquer capítulo | [[problem-formulation]], [[tsp-variants]] |
| Figuras/tabelas mínimas de resultados | Criadas antes de Experimentos | `scripts/`, `monografia/figs/` |
| Protocolo estatístico final | Fechado antes de Experimentos | [[analysis-methodology]] |

## Protocolo de Validação Antes de Escrever

Todo agente deve executar esta sequência antes de produzir texto para a monografia:

1. Identificar quais afirmações o capítulo precisa fazer.
2. Conferir no `src/` como o problema, método, pipeline ou renderização está implementado.
3. Conferir em `src/data/` quais instâncias, resultados e tempos existem de fato.
4. Buscar e entender a literatura necessária para contextualizar os achados.
5. Organizar a síntese no `vault/`, corrigindo notas que estejam desatualizadas.
6. Usar o `vault/` validado como base de escrita para `monografia/`.

O agente não deve tratar uma nota do `vault/` como evidência suficiente quando a afirmação depender de implementação ou resultado experimental. Nesses casos, a nota deve apontar para o código ou para o dado bruto que sustenta a afirmação.

## Terminologia Oficial

Usar os termos abaixo de forma consistente:

| Conceito | Termo recomendado | Observação |
|---|---|---|
| Problema principal | TSP-SD-ATP | Definir na primeira ocorrência |
| Tradução explicativa | TSP com penalidades angulares dependentes da sequência | Usar em texto corrido |
| Apelido interno | rTSP | Evitar na monografia, exceto se explicar que é nomenclatura interna |
| Métrica objetivo | makespan | Definir como tempo/custo total da rota |
| Pontos visitados | pontos de interesse ou POIs | Manter uma forma predominante |
| Veículo | drone ou VANT | Escolher uma forma predominante; "drone" é mais direto |

## Claims Centrais Permitidos

Estes claims podem orientar a escrita, mas devem ser verificados contra os dados finais antes de entrar no texto definitivo.

| Claim | Força recomendada | Evidência esperada |
|---|---|---|
| O problema implementado difere do TSP clássico porque o custo depende do nó anterior | Forte | [[problem-formulation]] |
| O tensor 3D permite avaliar rotas em O(n), após pré-computação O(n³) | Forte | [[problem-formulation]] |
| GA, PSO e ACO são metaheurísticas canônicas para TSP e variantes | Forte | [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]] |
| ACO obteve melhor qualidade de solução nas instâncias avaliadas | Forte se os dados finais confirmarem | [[resultados]] |
| GA foi muito mais rápido que ACO, com pior qualidade em instâncias grandes | Forte se os dados finais confirmarem | [[resultados]], timing |
| PSO teve desempenho inferior nesta variante específica | Moderado; explicar pela codificação random keys com cautela | [[resultados]], [[pso]] |
| O lower bound via relaxação AP sobre matriz reduzida é um limitante inferior válido para o TSP-SD-ATP | Forte (por construção matemática) | [[lower-bounds]], código em `src/optimization/lowerbound/` |
| O bound AP é válido mas frouxo para o TSP-SD-ATP: gap ~50% (40–65%) vs brute-force nas instâncias pequenas | Forte (dados confirmam) | Auditoria em `src/data/results/` vs BF; [[auditoria-codigo-dados-vault]] |
| A variante TSP-SD-ATP não aparece como benchmark consolidado na literatura revisada | Moderado; escrever como lacuna da revisão realizada | [[comparative-studies]] |

## Roteiro por Capítulo

### Capítulo 1 — Introdução

**Objetivo:** apresentar o problema, justificar o cenário de patrulha com drones, delimitar a lacuna e declarar contribuições.

**Entradas obrigatórias:** [[introducao]], [[problem-formulation]], [[drone-routing]], [[comparative-studies]], [[resultados]].

**Estrutura sugerida:** contexto de patrulha com drones; roteamento como problema central; TSP como abstração inicial; limitação do TSP clássico para curvas; definição sucinta do TSP-SD-ATP; lacuna da revisão; objetivo geral; objetivos específicos; contribuições; organização do texto.

**Critérios de aceite:** a introdução não promete mais do que os experimentos entregam; a lacuna está escrita de forma defensável; os objetivos são mensuráveis; as contribuições correspondem a capítulos e artefatos reais.

**Riscos:** exagerar novidade; transformar a aplicação em drone delivery em vez de patrulha; antecipar resultados com números ainda não validados.

### Capítulo 2 — Fundamentação Teórica

**Objetivo:** dar suporte conceitual para o problema, os métodos e a comparação experimental.

**Entradas obrigatórias:** [[fundamentacao]], [[TSP]], [[tsp-variants]], [[bio-inspired-optimization]], [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]], [[drone-routing]], [[comparative-studies]].

**Estrutura sugerida:** TSP clássico e complexidade; variantes relevantes do TSP; custos de curva e dependência de sequência; roteamento de drones; metaheurísticas bio-inspiradas; GA para TSP; PSO para TSP; ACO para TSP; estudos comparativos relacionados.

**Referências centrais:** [[garey1979computers]], [[lawler1985traveling]], [[applegate2006traveling]], [[winter2002modeling]], [[vanhove2012route]], [[holland1975adaptation]], [[goldberg1989genetic]], [[kennedy1995particle]], [[clerc2000discretepso]], [[dorigo1996ant]], [[dorigo1997ant]], [[stutzle2000mmas]], [[murray2015flying]], [[agatz2018optimization]], [[chandra2022comparative]], [[wu2020comparative]], [[halim2019combinatorial]], Demšar (2006) para protocolo estatístico.

**Critérios de aceite:** cada método tem origem, mecanismo básico e relação com TSP; a variante TSP-SD-ATP é posicionada sem inventar uma taxonomia não suportada; trabalhos relacionados são sintetizados, não listados; o capítulo prepara diretamente a proposta.

**Riscos:** excesso de revisão genérica; citar metaheurísticas modernas sem conexão com o que foi implementado; confundir TSP-D/FSTSP com o problema deste trabalho; esquecer de incluir lower bound AP na descrição dos métodos implementados.

### Capítulo 3 — Proposta

**Objetivo:** descrever o problema implementado, a representação computacional e os algoritmos usados no benchmark.

**Entradas obrigatórias:** código em `src/graph/`, `src/points/`, `src/optimization/`, `src/cmd/`, `src/shared/`; depois [[proposta]], [[problem-formulation]], [[architecture]], [[ga]], [[pso]], [[aco]], [[bruteforce]], [[lower-bounds]], [[experiment-pipeline]].

**Estrutura sugerida:** formulação do TSP-SD-ATP; geração de instâncias; tensor 3D de custos; função objetivo; arquitetura da implementação; GA; PSO; ACO; busca exaustiva; lower bound via relaxação AP; pipeline experimental; política de reprodutibilidade.

**Critérios de aceite:** a formulação matemática bate com o código real em `src/`; parâmetros são explicitados conforme flags e defaults implementados; decisões de design são apresentadas como escolhas pragmáticas, não como ótimos universais; busca exaustiva aparece como baseline para instâncias pequenas; lower bound AP (redução 3D→2D + Hungarian) é justificado teoricamente e sua implementação descrita; o capítulo permite reproduzir o experimento em alto nível.

**Riscos:** misturar descrição da proposta com resultados; justificar parâmetros sem evidência; omitir que não houve tuning sistemático; não explicar a distinção entre lower bound AP (bound numérico) e os métodos que produzem rotas completas.

### Capítulo 4 — Experimentos e Resultados

**Objetivo:** apresentar configuração experimental, dados coletados, análise de qualidade, análise de tempo e discussão dos resultados.

**Entradas obrigatórias:** arquivos em `src/data/`, `src/data/results/`, código/scripts que geram resultados e figuras; depois [[experimentos]], [[resultados]], [[analysis-methodology]], [[experiment-pipeline]], figuras em `monografia/figs/`.

**Estrutura sugerida:** configuração experimental; instâncias e sementes; métodos e parâmetros; baseline por brute-force; lower bound para instâncias grandes (relaxação AP via Hungarian, redução 3D→2D); qualidade das soluções com gap vs AP bound; tempo computacional; trade-off qualidade-tempo; análise estatística; discussão por método; ameaças à validade.

**Figuras e tabelas mínimas:** tabela de ótimos brute-force; tabela comparativa AP bound vs ótimo (instâncias pequenas); tabela de gap médio (AP bound como referência); gráfico de qualidade por método e tamanho; gráfico de tempo por método e tamanho (incluindo lower bound, ∼0ms); gráfico qualidade versus tempo; curvas de convergência; visualização de rotas selecionadas.

**Critérios de aceite:** todos os números têm fonte rastreável em `src/data/results/`; gráficos têm escala, unidade e legenda; testes estatísticos são reportados apenas se executados; conclusões são proporcionais aos dados brutos; limitações experimentais aparecem antes da conclusão final.

**Riscos:** declarar significância sem teste; comparar médias sem variabilidade; usar apenas melhor caso para métodos estocásticos; esquecer que brute-force só cobre instâncias pequenas; usar AP bound como se fosse ótimo verdadeiro em vez de limitante inferior.

### Capítulo 5 — Conclusão

**Objetivo:** sintetizar o que foi realizado, responder aos objetivos, delimitar limitações e propor trabalhos futuros.

**Entradas obrigatórias:** [[conclusao]], [[resultados]], [[analysis-methodology]], [[comparative-studies]], referências futuras como [[deepaco2023]], [[neufaco2025]], [[gpaco2025]], [[dellamico2021multiple]].

**Estrutura sugerida:** retomada do objetivo; resposta direta aos objetivos específicos; síntese dos principais achados; contribuições; limitações; trabalhos futuros.

**Critérios de aceite:** não apresenta resultado novo; não usa linguagem mais forte que o capítulo de Experimentos; limitações são explícitas; trabalhos futuros derivam das limitações; encerra voltando ao cenário de patrulha com drones.

**Riscos:** concluir que ACO é melhor em geral, em vez de melhor no desenho experimental realizado; minimizar limitações; propor trabalhos futuros desconectados do projeto.

> As perguntas e respostas que originaram estas tarefas estão documentadas em [[entrevista-sessao]].

## Tarefas Preparatórias Para Agentes

| ID | Tarefa | Saída esperada | Prioridade |
|---|---|---|---|
| P1 | Criar matriz claim-evidência | Nota `vault/writing/claim-evidence-matrix.md` | Concluída |
| P2 | Atualizar índice de papers | `vault/papers/index.md` com contagem e categorias atuais | Concluída |
| P3 | Criar glossário terminológico | Nota `vault/writing/glossario-monografia.md` | Concluída |
| P4 | Fechar protocolo estatístico | Validado em [[analysis-methodology]]; `scripts/analise-estatistica.py` corrigido; `cd-diagram.{svg,png}` regenerado | Concluída |
| P5 | Listar figuras/tabelas finais | Nota `vault/writing/figuras-tabelas-monografia.md` | Concluída |
| P6 | Mapear cada capítulo para arquivos `.tex` | Nota `vault/writing/mapa-capitulos-tex.md` | Concluída |
| P7 | Auditar divergências entre `src/`, `src/data/` e `vault/` | Nota `vault/writing/auditoria-codigo-dados-vault.md` | Concluída |
| P8 | Implementar e validar lower bound para instâncias grandes (n ≥ 15) | Implementação completa: algoritmo Hungarian O(n³) em Go puro, redução 3D→2D, comando `tcc optimize lowerbound`, 30/30 instâncias executadas, figuras/overlays geradas, integração nos scripts de análise. Bound validado contra brute-force (10a: AP=4.86 ≤ BF=8.25). Documentado em [[lower-bounds]], [[justificativa-lowerbound]] | Concluída |
| P9 | Estudar efeito da codificação na comparação justa entre métodos | Concluído em [[auditoria-codificacao-metodos]]. GA usa permutação direta; PSO usa random keys; ACO usa transições 3D. Explicitar como ameaça à validade, sem nova implementação nesta monografia. | Concluída |
| P10 | Análise de sensibilidade a hiperparâmetros | Concluída em [[auditoria-hiperparametros]] como limitação/trabalho futuro. Não bloqueia a escrita; declarar que não houve tuning sistemático. | Concluída como limitação |

## Protocolo Para Cada Agente Escritor

Antes de escrever:

1. Ler este roadmap.
2. Ler o código e/ou dados brutos relacionados ao capítulo em `src/` e `src/data/`.
3. Ler a nota-base do capítulo em `vault/writing/`.
4. Ler as notas de projeto e área exigidas para o capítulo.
5. Separar claims em três classes: conceituais, metodológicos e experimentais.
6. Verificar se cada claim metodológico tem fonte no código.
7. Verificar se cada claim experimental tem fonte em `src/data/results/`.
8. Atualizar ou sinalizar qualquer nota do `vault/` que não reflita código ou dados atuais.

Durante a escrita:

1. Escrever em português acadêmico direto.
2. Evitar parágrafos genéricos sobre "importância" sem dado, referência ou função argumentativa.
3. Inserir citações apenas quando a referência sustenta a frase específica.
4. Preservar a narrativa de patrulha com drones.
5. Manter o problema como TSP-SD-ATP, não como benchmark genérico de otimizadores.

Após a escrita:

1. Conferir se objetivos, métodos, dados e conclusões estão alinhados.
2. Marcar claims que ainda precisam de validação.
3. Listar figuras/tabelas citadas e verificar se existem.
4. Verificar se não há promessa sem evidência.
5. Registrar pendências em nota separada ou no topo do capítulo.

## Protocolo de Análise Estatística (P4 — Concluído)

Com 3 métodos estocásticos × 30 instâncias × 51 sementes = 4590 execuções estocásticas, comparar apenas médias/melhores casos é insuficiente. O protocolo segue Demšar (2006), detalhado em [[analysis-methodology#Protocolo-de-Análise-Estatística]]. O script estatístico foi corrigido e validado em [[auditoria-script-analise-estatistica]].

### Resumo

| Etapa | Teste | O que responde |
|-------|-------|---------------|
| 1 | **Friedman** | Rejeita H₀: "todos os métodos são equivalentes" |
| 2 | **Nemenyi post-hoc** | Se Friedman rejeitar, mostra quais pares (GA×PSO, GA×ACO, PSO×ACO) diferem com significância |
| 3 | **Diagrama CD** | Visualização da diferença crítica entre métodos |
| 4 | **Wilcoxon signed-rank** | Alternativa pareada por instância para reforçar/confrontar Friedman+Nemenyi |

### Implementação

Script `scripts/analise-estatistica.py`:

```python
# Friedman/Iman-Davenport: 3 métodos, 30 instâncias, mediana por instância
# Nemenyi: CD com q_alpha compatível com Demšar
# Wilcoxon/Holm: pareado por instância entre cada par de métodos
```

### Saídas validadas

- Friedman/Iman-Davenport: F(2,58)=293.2222, p=4.710129e-31
- Ranks médios: ACO=1.1000, GA=1.9000, PSO=3.0000
- Nemenyi: CD=0.6050; todos os pares significativos
- Wilcoxon/Holm: todos os pares significativos
- Diagrama CD exportado para `monografia/figs/cd-diagram.{png,svg}`

## Critério de Pronto Para Escrita Definitiva

A monografia estará pronta para escrita definitiva quando estas condições forem atendidas:

| Condição | Aceite |
|---|---|
| Claims principais mapeados | Cada claim tem fonte e força definida |
| Código e dados auditados | Claims metodológicos conferidos em `src/`; claims experimentais conferidos em `src/data/` |
| Estatística resolvida | Testes executados ou análise descritiva assumida explicitamente |
| Figuras mínimas disponíveis | Cada figura tem escala, unidade e fonte |
| Referências centrais selecionadas | Cada capítulo tem bibliografia mínima definida |
| Terminologia estabilizada | TSP-SD-ATP, makespan, drone e POI usados de forma consistente |
| Limitações declaradas | Parâmetros fixos, instâncias sintéticas e baseline limitado aparecem no texto |

## Próxima Ação Recomendada

Com P1–P10 concluídos, a próxima ação é escrever/revisar os capítulos `.tex` usando as notas atualizadas e os artefatos regenerados. Começar por Experimentos ou Proposta, pois são os capítulos que mais dependem dos dados e do código.
