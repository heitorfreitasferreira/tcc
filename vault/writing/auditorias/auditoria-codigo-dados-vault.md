---
title: Auditoria Código-Dados-Vault
tags:
- evidencia/auditoria
- evidencia/codigo
- evidencia/dados
- status/atualizado
- tipo/auditoria
- tipo/writing
- topico/monografia
status: atualizado-2026-06-02
created: 2026-06-02
updated: 2026-06-02
type: writing
---

# Auditoria Código-Dados-Vault

Esta auditoria compara as notas centrais do `vault/` com a fonte determinística do projeto: implementação em `src/`, dados brutos em `src/data/` e artefatos gerados em `monografia/figs/`.

> [!warning] Regra para escrita
> Claims metodológicos podem ser usados quando estiverem ancorados no código indicado. Claims experimentais devem ser recalculados a partir de `src/data/results/` antes de entrar em texto final, porque a cobertura atual mudou em relação a notas anteriores.

## Resultado Executivo

| Item | Status auditado | Decisão para escrita |
|---|---|---|
| Formulação TSP-SD-ATP | Confirmada pelo código | Pode ser usada |
| GA | Confirmado pelo código | Pode ser usado |
| PSO | Confirmado pelo código | Pode ser usado, com limitação da codificação random keys |
| ACO | Confirmado pelo código | Pode ser descrito como Ant System adaptado com feromônio 3D |
| Brute-force | Confirmado pelo código e dados atuais | Agora cobre 18 instâncias (`10a..15c`), não 15 |
| Lower bound AP | Confirmado pelo código | Pode ser usado como limitante inferior, não como rota factível |
| Cobertura experimental atual | 4638 summaries, 4638 timings, 4638 evolutions | Atualizar notas que ainda dizem 4605 ou 4635 |
| Resultados agregados em [[resultados]] | Atualizados no vault | Tabelas finais ainda devem ser geradas a partir dos dados brutos |
| Testes estatísticos | Script existe, mas está metodologicamente inconsistente | Não declarar significância até corrigir/auditar o script |

## Fonte de Verdade Auditada

| Camada | Caminhos | Papel |
|---|---|---|
| Código | `src/graph/`, `src/points/`, `src/optimization/`, `src/cmd/`, `src/shared/reporting/` | Formulação, algoritmos, flags, schemas e persistência |
| Dados | `src/data/*.points`, `src/data/*.graph`, `src/data/results/{summary,evolution,timing}/` | Instâncias e resultados brutos |
| Scripts | `scripts/analise-estatistica.py`, `scripts/gerar-graficos-estatisticos.py`, `scripts/gerar-analises.sh` | Análise estatística e geração de figuras |
| Figuras | `monografia/figs/` | Artefatos visuais disponíveis |
| Vault | `vault/projeto/*.md`, `vault/writing/**/*.md` | Síntese intermediária; não é evidência primária |

## Achados Metodológicos Confirmados

### Formulação e Grafo

| Claim | Status | Evidência primária |
|---|---|---|
| O grafo é um tensor 3D | Confirmado | `src/graph/types.go:8` define `type Graph [][][]float64` |
| A indexação é `[anterior][atual][proximo]` | Confirmado | `src/graph/types.go:8`, `src/graph/makespan.go:7-19` |
| O custo soma distância euclidiana e penalidade angular | Confirmado | `src/graph/creater.go:24-33`, `src/graph/math.go:33` |
| `droneSpeed = 1` e `maxPenalti = 1` | Confirmado | `src/graph/types.go:5-6` |
| A rota parte do nó 0 e retorna ao nó 0 | Confirmado | `src/graph/makespan.go:7-19` |
| Existem 30 instâncias `.points` e 30 `.graph` | Confirmado | `src/data/10a..100c.{points,graph}` |

### Métodos de Otimização

| Método | Claim confirmado | Evidência primária |
|---|---|---|
| GA | Representação por permutação dos nós `1..n-1` | `src/optimization/ga/main.go:57-67` |
| GA | Seleção por torneio, crossover OX, mutação swap e elitismo | `src/optimization/ga/main.go:75-95`, `src/optimization/ga/main.go:130-201` |
| GA | Defaults: população 100, iterações 100, elitismo 1, mutação 0.05, torneio 2 | `src/cmd/optimize.go:36-37`, `src/cmd/ga.go:64-66` |
| PSO | Representação por random keys e ordenação de posições contínuas | `src/optimization/pso/particle.go:23-40` |
| PSO | Atualização contínua com inércia, componente cognitivo e social | `src/optimization/pso/particle.go:42-54` |
| PSO | Defaults: `c1=2.0`, `c2=2.0`, `w=0.7` | `src/cmd/pso.go:68-70` |
| ACO | Feromônio 3D `pheromones[prev][curr][next]` | `src/optimization/aco/main.go:21`, `src/optimization/aco/main.go:84-104` |
| ACO | Seleção por roleta proporcional com `tau^alpha * eta^beta`, `eta=1/G` | `src/optimization/aco/ant.go:42-84` |
| ACO | Depósito por todas as formigas | `src/optimization/aco/ant.go:27-39` |
| ACO | Defaults: `alpha=1.0`, `beta=2.0`, `rho=0.2`, `q=100` | `src/cmd/aco.go:66-69` |
| Brute-force | Enumera permutações dos nós `1..n-1` por algoritmo de Heap | `src/optimization/brute/main.go:9-79` |
| Lower bound | Reduz tensor 3D para matriz 2D via `c'[j][k] = min_i cost[i][j][k]` | `src/optimization/lowerbound/main.go:42-60` |
| Lower bound | Aplica Hungarian O(n^3) e retorna bound possivelmente com subtours | `src/optimization/lowerbound/hungarian.go`, `src/optimization/lowerbound/main.go:62-89` |

## Cobertura Experimental Atual

Contagem recalculada em `src/data/results/` nesta auditoria:

| Artefato | Quantidade |
|---|---:|
| `summary/*.json` | 4638 |
| `evolution/*.jsonl` | 4638 |
| `timing/*.json` | 4638 |
| Links `timing_file` e `evolution_file` ausentes | 0 |

Distribuição por método nos summaries:

| Método | Runs | Instâncias | Sementes |
|---|---:|---:|---:|
| ACO | 1530 | 30 | 51 (`s0..s50`) |
| GA | 1530 | 30 | 51 (`s0..s50`) |
| PSO | 1530 | 30 | 51 (`s0..s50`) |
| Lower bound | 30 | 30 | 1 (`s0`) |
| Brute-force | 18 | 18 | 1 (`s0`) |

> [!success] Cobertura normalizada no vault
> [[resultados]] e [[analysis-methodology]] foram atualizadas após esta auditoria para refletir 4638 summaries e 18 brute-force, incluindo `15a`, `15b` e `15c`.

### Observação Sobre Caminhos de Instância

Os summaries usam uma mistura de caminhos absolutos (`/home/heitor/tcc/src/data/10a.graph`) e relativos (`src/data/30a.graph`) no campo `instance`. O `run_id` continua normalizado por nome de instância, e não há duplicatas lógicas para `(instância, método, seed)` quando a instância é normalizada por `Path(instance).stem`.

Decisão para análise: scripts e tabelas devem normalizar `instance` pelo nome-base (`10a`, `30b`, etc.) antes de agrupar resultados.

## Baseline Brute-force Confirmada

| Instância | Makespan ótimo |
|---|---:|
| 10a | 8.251282810997422 |
| 10b | 7.906130315558171 |
| 10c | 9.545033907523141 |
| 11a | 9.493457779404203 |
| 11b | 8.823027373557352 |
| 11c | 10.260221738390818 |
| 12a | 9.500700789073989 |
| 12b | 9.768963312506418 |
| 12c | 10.47963364504159 |
| 13a | 10.950258153646407 |
| 13b | 11.473574261140902 |
| 13c | 10.041955396200756 |
| 14a | 11.55692517246078 |
| 14b | 11.466929014715355 |
| 14c | 11.486447067382262 |
| 15a | 9.128679854598571 |
| 15b | 12.184839159096176 |
| 15c | 12.16021037136399 |

## Gap vs Brute-force nas 18 Instâncias com Ótimo

Agregado sobre 18 instâncias e 51 sementes por método, normalizando o campo `instance` pelo nome-base.

| Método | Runs | Gap médio | Gap mínimo | Gap máximo | Taxa global de ótimo |
|---|---:|---:|---:|---:|---:|
| ACO | 918 | 0.4296% | 0.0000% | 9.4445% | 71.79% |
| GA | 918 | 5.2814% | 0.0000% | 41.4267% | 26.47% |
| PSO | 918 | 22.2915% | 0.0000% | 76.5549% | 4.25% |

> [!success] Claims experimentais atualizados
> [[resultados]] e [[claim-evidence-matrix]] foram atualizadas para usar os gaps das 18 instâncias com brute-force (`10a..15c`).

## Instâncias Grandes

Os números abaixo continuam consistentes com [[resultados]] para `50a..100c`, considerando 51 sementes por método.

| Instância | ACO best | ACO média | GA best | GA média | PSO best | PSO média |
|---|---:|---:|---:|---:|---:|---:|
| 50a | 24.6688 | 25.6475 | 41.0120 | 45.3925 | 53.7651 | 59.7284 |
| 50b | 26.3528 | 27.4843 | 42.6487 | 48.6345 | 57.0330 | 64.3194 |
| 50c | 26.2372 | 27.0062 | 42.0122 | 46.9438 | 54.4114 | 62.3455 |
| 100a | 42.6718 | 45.3528 | 97.5921 | 110.6608 | 127.4034 | 136.5131 |
| 100b | 45.4689 | 48.0596 | 97.5343 | 110.4959 | 126.3967 | 135.9987 |
| 100c | 45.4858 | 46.9487 | 95.1632 | 111.3243 | 126.1838 | 138.1908 |

## Tempo de Otimização em n=100

Tempos médios em milissegundos, lidos dos `timing_file` ligados aos summaries.

| Instância | ACO média ms | GA média ms | PSO média ms | Razão ACO/GA | Razão ACO/PSO |
|---|---:|---:|---:|---:|---:|
| 100a | 4262.06 | 38.82 | 76.24 | 109.78x | 55.91x |
| 100b | 4276.55 | 36.27 | 75.25 | 117.89x | 56.83x |
| 100c | 4266.55 | 39.37 | 73.84 | 108.37x | 57.78x |

## Lower Bound AP vs Brute-force

O AP bound foi encontrado para todas as 30 instâncias e é menor ou igual ao ótimo brute-force em todas as 18 instâncias com BF disponível. O gap abaixo é `(BF - LB) / BF`.

| Instância | BF ótimo | LB bound | Gap LB->BF |
|---|---:|---:|---:|
| 10a | 8.251283 | 4.859238 | 41.11% |
| 10b | 7.906130 | 2.979065 | 62.32% |
| 10c | 9.545034 | 4.956950 | 48.07% |
| 11a | 9.493458 | 4.529660 | 52.29% |
| 11b | 8.823027 | 4.158859 | 52.86% |
| 11c | 10.260222 | 6.075174 | 40.79% |
| 12a | 9.500701 | 4.549388 | 52.12% |
| 12b | 9.768963 | 4.091362 | 58.12% |
| 12c | 10.479634 | 4.088736 | 60.98% |
| 13a | 10.950258 | 5.799255 | 47.04% |
| 13b | 11.473574 | 6.078440 | 47.02% |
| 13c | 10.041955 | 5.347435 | 46.75% |
| 14a | 11.556925 | 4.004625 | 65.35% |
| 14b | 11.466929 | 5.675375 | 50.51% |
| 14c | 11.486447 | 6.420654 | 44.10% |
| 15a | 9.128680 | 4.716761 | 48.33% |
| 15b | 12.184839 | 6.102700 | 49.92% |
| 15c | 12.160210 | 5.259865 | 56.75% |

Resumo: gap médio 51.36%, mínimo 40.79%, máximo 65.35%. O bound é válido, mas frouxo para o TSP-SD-ATP.

## Auditoria dos Scripts de Análise

| Script | Status | Achado |
|---|---|---|
| `scripts/analise-estatistica.py` | Existe e executa | Gera `monografia/figs/cd-diagram.svg`; não usa SciPy |
| `scripts/analise-estatistica.py` | Inconsistente com [[analysis-methodology]] | Usa média por instância, não mediana |
| `scripts/analise-estatistica.py` | Incompleto para o protocolo declarado | Não implementa Wilcoxon nem ajuste Holm |
| `scripts/analise-estatistica.py` | Problema numérico/metodológico | A fórmula da estatística Friedman aparenta usar somas de ranks de forma incorreta, gerando `χ²_F = 377640.0000` para 30 blocos e 3 métodos |
| `scripts/analise-estatistica.py` | Mensagem incorreta | Imprime `Execuções por instância: 30`, mas os métodos estocásticos têm 51 sementes por instância |
| `scripts/gerar-graficos-estatisticos.py` | Existe | Gera heatmaps, dispersão qualidade-tempo, escalabilidade e estabilidade sem dependências externas |
| `scripts/gerar-analises.sh` | Existe | Pipeline canônico atual para figuras: build do servidor, diagramas TikZ e figuras compostas |

Decisão para escrita: não usar claims de significância estatística ainda. Antes, corrigir o script para seguir o protocolo de [[analysis-methodology]] ou revisar a metodologia para refletir exatamente o script executado.

## Artefatos Visuais Disponíveis

Há artefatos em `monografia/figs/`, incluindo:

| Tipo | Exemplos |
|---|---|
| Diagrama CD | `cd-diagram.svg`, `cd-diagram.png` |
| Gráficos estatísticos | `heatmap-success-rate.png`, `heatmap-gap-vs-bf.png`, `scatter-quality-vs-time.png`, `scalability-runtime.png`, `boxplot-estabilidade.png`, `performance-profile.svg` |
| Rotas renderizadas | `route-10a-aco-s0-iter1.png`, `route-100a-aco-s0.png`, `overlay-100a-methods.png` |
| Diagramas conceituais | `diagram-tensor-3d.svg`, `diagram-angular-penalty.svg`, `flowchart-ga.tex`, `flowchart-pso.tex`, `flowchart-aco.tex`, `flowchart-lowerbound.tex` |

Esses artefatos existem, mas a seleção final de figuras ainda deve passar por [[figuras-tabelas-monografia]] e pela regra de escala/unidade do roadmap.

## Divergências Prioritárias no Vault

| Nota | Divergência | Ação recomendada |
|---|---|---|
| ~~[[resultados]]~~ | ~~Declara 4605 execuções, 15 brute-force e gaps do subconjunto `10a..14c`~~ | Resolvido: atualizado para cobertura P7 |
| ~~[[analysis-methodology]]~~ | ~~Declara 4635 runs e 15 brute-force~~ | Resolvido quanto à cobertura e estatística |
| ~~[[claim-evidence-matrix]]~~ | ~~Claims E02, E03, E06, E17 usam cobertura antiga~~ | Resolvido |
| ~~[[experiment-pipeline]]~~ | ~~Omite `lowerbound`~~ | Resolvido |
| ~~[[experiment-pipeline]]~~ | ~~Formato de run ID contém `p{pop}__i{iter}`~~ | Resolvido |
| ~~[[roadmap-monografia]]~~ | ~~P4 aparece resolvido, mas o script estatístico não cumpre integralmente o protocolo~~ | Resolvido após correção e validação de P4 |

## Claims Permitidos Após Esta Auditoria

| Claim | Força | Condição |
|---|---|---|
| O TSP-SD-ATP implementado usa custo dependente da sequência com tensor 3D | Forte | Citar `src/graph/` |
| GA, PSO, ACO, brute-force e lowerbound estão implementados como comandos do `tcc optimize` | Forte | Citar `src/cmd/` |
| ACO obteve menor makespan médio que GA e PSO nas instâncias avaliadas | Forte descritivo | Usar dados recalculados, sem significância estatística ainda |
| GA foi muito mais rápido que ACO em n=100 | Forte descritivo | Usar tempos médios por instância |
| PSO teve desempenho inferior neste desenho experimental | Moderado | Delimitar à implementação random keys e parâmetros fixos |
| Brute-force fornece ótimo exato para `10a..15c` | Forte | Citar 18 summaries BF |
| AP bound é válido mas frouxo | Forte | Citar LB ≤ BF e gap médio 51.36% nas 18 instâncias |

## Pendências Antes da Escrita Definitiva

1. Aguardar/corroborar a correção de `scripts/analise-estatistica.py` antes de liberar claims de significância.
2. Regerar ou validar tabelas finais com `scripts/consolidate_results.py` contra a cobertura atual.
3. Normalizar agrupamentos por nome-base de instância em qualquer script novo.
4. Reescrever os `.tex` conforme [[mapa-capitulos-tex]], evitando reaproveitar resultados antigos.
