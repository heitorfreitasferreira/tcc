---
title: Matriz Claim-Evidência da Monografia
tags:
  - writing
  - claims
  - evidencias
  - monografia
status: atualizada-pos-p7
created: 2026-06-02
updated: 2026-06-02
---

# Matriz Claim-Evidência da Monografia

Esta matriz lista claims que podem orientar a escrita da monografia. A evidência primária é sempre `src/` para claims metodológicos e `src/data/` para claims experimentais. O `vault/` aparece como camada de organização, não como fonte primária.

## Legenda

| Campo | Significado |
|---|---|
| Forte | Pode ser usado se citado com a evidência indicada |
| Moderado | Pode ser usado com escopo delimitado |
| Fraco | Usar apenas como hipótese, interpretação ou discussão |
| Bloqueado | Não usar até resolver divergência |

## Claims Conceituais e de Literatura

| ID | Claim | Força | Evidência primária | Vault de apoio | Uso recomendado |
|---|---|---|---|---|---|
| C01 | O TSP clássico modela a busca por uma rota que visita todos os pontos e retorna à origem com custo mínimo. | Forte | Literatura | [[TSP]], [[garey1979computers]], [[lawler1985traveling]] | Fundamentação e Introdução |
| C02 | O TSP é computacionalmente difícil e motiva o uso de heurísticas/metaheurísticas em instâncias maiores. | Forte | Literatura | [[garey1979computers]], [[applegate2006traveling]] | Fundamentação |
| C03 | Problemas de roteamento com drones aparecem em patrulha, inspeção, monitoramento e entrega, mas FSTSP/TSP-D não são o mesmo problema deste trabalho. | Forte | Literatura | [[drone-routing]], [[murray2015flying]], [[agatz2018optimization]], [[rajan2022routing]] | Introdução e Fundamentação |
| C04 | GA, PSO e ACO são metaheurísticas canônicas aplicadas ao TSP e variantes. | Forte | Literatura | [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]], [[comparative-studies]] | Fundamentação |
| C05 | Não foram identificados, na revisão organizada no vault, estudos que comparem sistematicamente GA, PSO e ACO especificamente no TSP-SD-ATP. | Moderado | Revisão bibliográfica auditável | [[comparative-studies]] | Introdução; escrever como lacuna da revisão, não como prova absoluta de inexistência |

## Claims de Formulação e Implementação

| ID | Claim | Força | Evidência primária | Vault de apoio | Uso recomendado |
|---|---|---|---|---|---|
| M01 | A variante implementada usa um tensor de custo 3D `G[anterior][atual][proximo]`. | Forte | `src/graph/types.go:8`, `src/graph/makespan.go:13` | [[problem-formulation]] | Proposta |
| M02 | O custo de transição soma distância euclidiana e penalidade angular. | Forte | `src/graph/creater.go:24-33`, `src/graph/math.go` | [[problem-formulation]] | Proposta/Fundamentação |
| M03 | A avaliação de uma rota percorre a sequência de POIs e adiciona o retorno ao nó 0. | Forte | `src/graph/makespan.go:7-19` | [[problem-formulation]] | Proposta |
| M04 | O nó 0 é origem/base implícita e não aparece na permutação dos métodos. | Forte | `src/graph/makespan.go:8-18`, `src/optimization/ga/main.go:57-61`, `src/optimization/brute/main.go:10-12` | [[problem-formulation]], [[ga]], [[bruteforce]] | Proposta |
| M05 | As instâncias possuem arquivos `.points` e `.graph` em `src/data/`. | Forte | `src/data/*.points`, `src/data/*.graph` | [[experiment-pipeline]] | Proposta/Experimentos |
| M06 | O pipeline de resultados grava summaries, evolution e timing em `src/data/results/`. | Forte | `src/shared/reporting/reporting.go:148-156`, `src/shared/reporting/reporting.go:37-84` | [[experiment-pipeline]] | Proposta/Experimentos |

## Claims Sobre Métodos Implementados

| ID | Claim | Força | Evidência primária | Vault de apoio | Uso recomendado |
|---|---|---|---|---|---|
| A01 | O GA usa representação por permutação dos nós `1..n-1`. | Forte | `src/optimization/ga/main.go:57-67` | [[ga]] | Proposta |
| A02 | O GA usa seleção por torneio, OX, mutação swap e elitismo. | Forte | `src/optimization/ga/main.go:75-95`, `src/optimization/ga/main.go:130-201` | [[ga]] | Proposta |
| A03 | Os defaults do GA são população 100, iterações 100, elitismo 1, mutação 0.05 e torneio 2. | Forte | `src/cmd/optimize.go:36-37`, `src/cmd/ga.go:64-66` | [[ga]] | Proposta/Experimentos |
| A04 | O PSO usa random keys: ordena posições contínuas para produzir permutação. | Forte | `src/optimization/pso/particle.go:23-40` | [[pso]] | Proposta |
| A05 | O PSO usa atualização contínua com inércia, componente cognitivo e componente social. | Forte | `src/optimization/pso/particle.go:42-54` | [[pso]] | Proposta |
| A06 | Os defaults atuais do PSO são `c1=2.0`, `c2=2.0`, `w=0.7`. | Forte | `src/cmd/pso.go:68-70` | [[pso]] | Proposta/Experimentos |
| A07 | O ACO usa feromônio 3D com a mesma dependência de sequência do tensor de custo. | Forte | `src/optimization/aco/main.go:21`, `src/optimization/aco/main.go:84-104` | [[aco]] | Proposta |
| A08 | O ACO constrói rotas por roleta proporcional usando `tau^alpha * eta^beta`, com `eta=1/G`. | Forte | `src/optimization/aco/ant.go:42-84` | [[aco]] | Proposta |
| A09 | O ACO deposita feromônio a partir de todas as formigas, não apenas da melhor global. | Forte | `src/optimization/aco/ant.go:27-39` | [[aco]] | Proposta |
| A10 | Os defaults atuais do ACO são `alpha=1.0`, `beta=2.0`, `rho=0.2`, `q=100`. | Forte | `src/cmd/aco.go:66-69` | [[aco]] | Proposta/Experimentos |
| A11 | A busca exaustiva enumera permutações por algoritmo de Heap e serve como baseline exata nas instâncias executadas. | Forte | `src/optimization/brute/main.go:20-79` | [[bruteforce]] | Proposta/Experimentos |
| A12 | O lower bound reduz o tensor 3D para matriz 2D via `c'[j][k] = min_i cost[i][j][k]` e aplica Hungarian O(n³). | Forte | `src/optimization/lowerbound/main.go:42-60`, `src/optimization/lowerbound/hungarian.go` | [[lower-bounds]] | Proposta |
| A13 | O lower bound produz um limitante inferior válido (≤ makespan ótimo) mas não gera rota factível (pode conter subtours). | Forte | `src/optimization/lowerbound/main.go:62-89` | [[lower-bounds]] | Proposta (explicar que é bound, não rota) |

## Claims Experimentais Auditados

Estes claims usam a cobertura auditada em [[auditoria-codigo-dados-vault]]: GA 1530, PSO 1530, ACO 1530, lowerbound 30 e brute-force 18. Para agrupamentos por instância, normalizar o campo `instance` pelo nome-base (`10a`, `30b`, etc.), pois os summaries misturam caminhos absolutos e relativos.

| ID | Claim | Força | Evidência primária | Vault de apoio | Uso recomendado |
|---|---|---|---|---|---|
| E01 | Existem 30 instâncias em `src/data/`, de 10 a 100 nós, com três variantes por tamanho usado. | Forte | `src/data/*.points`, `src/data/*.graph` | [[experiment-pipeline]] | Experimentos |
| E02 | O diretório contém 4638 summaries: 51 sementes para GA/PSO/ACO em 30 instâncias, 30 lowerbound e 18 brute-force. | Forte | `src/data/results/summary/*.json` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E03 | Brute-force fornece ótimos para `10a` a `15c`. | Forte | Summaries `*__bruteforce__*.json` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E05 | No subconjunto atual, ACO tem menor gap médio que GA e PSO nas instâncias com brute-force. | Forte descritivo | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Experimentos |
| E06 | Nas 18 instâncias com brute-force (`10a..15c`), ACO tem gap médio 0.4296%, GA 5.2814% e PSO 22.2915%. | Forte descritivo | `src/data/results/summary/*.json` filtrado e normalizado por nome-base | [[auditoria-codigo-dados-vault]] | Experimentos |
| E07 | No subconjunto atual, ACO não encontra o ótimo em 100% das execuções até `13c`. | Forte como correção | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Usar para evitar claim errado |
| E08 | Em `100a..100c`, ACO apresenta makespan médio substancialmente menor que GA e PSO. | Forte descritivo | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Experimentos/Conclusão |
| E09 | Em `100a`, médias de makespan: ACO 45.3528, GA 110.6608, PSO 136.5131. | Forte descritivo | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Experimentos |
| E10 | Em `100b`, médias de makespan: ACO 48.0596, GA 110.4959, PSO 135.9987. | Forte descritivo | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Experimentos |
| E11 | Em `100c`, médias de makespan: ACO 46.9487, GA 111.3243, PSO 138.1908. | Forte descritivo | `src/data/results/summary/*.json` filtrado | [[auditoria-codigo-dados-vault]] | Experimentos |
| E12 | Em n=100, ACO é muito mais lento que GA e PSO no tempo de otimização. | Forte descritivo | `timing_file` ligado aos summaries filtrados | [[auditoria-codigo-dados-vault]] | Experimentos/Conclusão |
| E13 | Em `100a`, tempos médios: ACO 4262.06 ms, GA 38.82 ms, PSO 76.24 ms. | Forte descritivo | `src/data/results/timing/*.json` via `timing_file` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E14 | Em `100b`, tempos médios: ACO 4276.55 ms, GA 36.27 ms, PSO 75.25 ms. | Forte descritivo | `src/data/results/timing/*.json` via `timing_file` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E15 | Em `100c`, tempos médios: ACO 4266.55 ms, GA 39.37 ms, PSO 73.84 ms. | Forte descritivo | `src/data/results/timing/*.json` via `timing_file` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E16 | O AP bound é um limitante inferior válido para todas as 30 instâncias (determinístico) e ≤ BF ótimo nas 18 com brute-force. | Forte descritivo | `src/data/results/summary/*__lowerbound__*.json` | [[auditoria-codigo-dados-vault]] | Experimentos |
| E17 | O AP bound é frouxo para o TSP-SD-ATP: gap médio 51.36% (min 40.79%, max 65.35%) vs BF ótimo. | Forte descritivo | `src/data/results/summary/*__lowerbound__*.json` vs BF | [[auditoria-codigo-dados-vault]] | Experimentos (bound válido mas não tight) |
| E18 | O lower bound é computacionalmente trivial (~0ms n≤15, ~5ms n=100). | Forte descritivo | `src/data/results/timing/*__lowerbound__*.json` | [[auditoria-codigo-dados-vault]] | Experimentos (contraste com metaheurísticas) |
| E19 | A análise estatística confirma ordenação por makespan mediano: ACO, depois GA, depois PSO. | Forte | `scripts/analise-estatistica.py`, `src/data/results/summary/*.json` | [[analysis-methodology]], [[auditoria-script-analise-estatistica]] | Experimentos |
| E20 | Com mediana por instância, os ranks médios são ACO 1.1000, GA 1.9000 e PSO 3.0000. | Forte | `scripts/analise-estatistica.py` | [[analysis-methodology]] | Experimentos |
| E21 | Friedman/Iman-Davenport rejeita equivalência entre métodos: F(2,58)=293.2222, p=4.710129e-31. | Forte | `scripts/analise-estatistica.py` | [[analysis-methodology]] | Experimentos |
| E22 | Nemenyi e Wilcoxon/Holm indicam diferença significativa nos três pares ACO×GA, ACO×PSO e GA×PSO. | Forte | `scripts/analise-estatistica.py`, `monografia/figs/cd-diagram.svg` | [[analysis-methodology]] | Experimentos |

## Claims Bloqueados ou Que Exigem Correção

| ID | Claim bloqueado | Motivo | Ação necessária |
|---|---|---|---|
| B01 | ACO encontra o ótimo em 100% das sementes para instâncias pequenas. | Contradiz dados atuais com 18 instâncias BF | Remover ou recalcular com subconjunto explicitado |
| B02 | ACO tem gap médio 0.00% nas instâncias pequenas. | Contradiz dados atuais com 18 instâncias BF | Usar 0.4296% para `10a..15c` ou explicitar outro subconjunto |
| B03 | GA tem gap médio 0.04% nas instâncias pequenas. | Contradiz dados atuais com 18 instâncias BF | Usar 5.2814% para `10a..15c` ou explicitar outro subconjunto |
| B04 | PSO tem gap médio 1.57% nas instâncias pequenas. | Contradiz dados atuais com 18 instâncias BF | Usar 22.2915% para `10a..15c` ou explicitar outro subconjunto |
| B05 | GA é ~75x mais rápido que ACO em n=100. | Subconjunto atual indica ~108x a ~118x em `100a..100c` | Atualizar para razão calculada ou reportar por instância |
| B06 | ACO em `100a` leva cerca de 2.859s. | Subconjunto atual indica média 4.262s | Atualizar com dados filtrados |
| ~~B07~~ | ~~Diferenças entre métodos são estatisticamente significativas.~~ | Resolvido: script corrigido e validado | Usar E21/E22 com os valores atuais |
| B08 | "AP bound tem gap < 30% nas instâncias pequenas." | Dados mostram gap real 40.79–65.35% nas 18 instâncias BF | Substituir por "AP bound é válido mas frouxo (gap 51.36% médio)" |
| ~~B09~~ | ~~O diagrama CD atual é evidência final.~~ | Resolvido: `cd-diagram.{svg,png}` regenerado após correção | Usar com E22 |

## Claims Interpretativos Permitidos com Cautela

| ID | Claim | Força | Condição de uso |
|---|---|---|---|
| I01 | ACO parece mais adequado à dependência de sequência do TSP-SD-ATP porque modela feromônio em triplas. | Moderado | Apresentar como interpretação apoiada por desempenho e estrutura, não como prova causal |
| I02 | GA oferece melhor custo computacional, mas sacrifica qualidade em instâncias grandes. | Moderado | Usar dados de makespan e tempo lado a lado |
| I03 | PSO com random keys teve desempenho inferior neste desenho experimental. | Moderado | Delimitar à implementação atual e aos parâmetros usados |
| I04 | A ausência de tuning sistemático limita generalizações. | Forte | Consta do desenho experimental; mencionar em limitações com apoio em [[auditoria-hiperparametros]] |
| I05 | O lower bound AP via redução 3D→2D é válido mas pouco informativo para o TSP-SD-ATP (gap ~50%). Métodos 3D-nativos (MDD, ng-path) poderiam produzir bounds mais justos. | Moderado | Dados mostram gap 40–65%; literatura indica MDD como alternativa | Apresentar como limitação do método e direção futura |
| I06 | As diferenças de representação entre GA, PSO e ACO são ameaça à validade externa da comparação. | Forte | Apoiar em [[auditoria-codificacao-metodos]]; usar como limitação, não como explicação causal definitiva |

## Próximas Tarefas

1. Usar [[auditoria-codificacao-metodos]] para redigir limitações sobre comparabilidade entre representações.
2. Usar [[auditoria-hiperparametros]] para redigir limitações sobre parâmetros fixos.
3. Escrever os capítulos `.tex` evitando reaproveitar parágrafos com cobertura antiga.
