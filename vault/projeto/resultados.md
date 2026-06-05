---
tags:
- capitulo/experimentos
- evidencia/codigo
- status/atualizado
- tipo/projeto
- topico/experimentos
- topico/implementacao
status: atualizado-pos-p7
updated: 2026-06-02
type: projeto
---

# Resultados dos Experimentos

Síntese dos resultados auditados em [[auditoria-codigo-dados-vault]]. Os números abaixo devem orientar a escrita, mas tabelas finais devem ser geradas novamente a partir de `src/data/results/` antes de entrar na monografia.

> [!warning] Estatística em correção
> Claims de significância estatística dependem da correção de `scripts/analise-estatistica.py`. Até lá, usar apenas linguagem descritiva: “ACO apresentou menor makespan”, “GA foi mais rápido”, “PSO teve desempenho inferior neste desenho experimental”.

## Cobertura

| Método | Runs | Sementes | Instâncias | Observação |
|---|---:|---:|---:|---|
| GA | 1530 | 51 | 30 | Método estocástico |
| PSO | 1530 | 51 | 30 | Método estocástico |
| ACO | 1530 | 51 | 30 | Método estocástico |
| Lower bound | 30 | 1 | 30 | Determinístico; limitante inferior |
| Brute-force | 18 | 1 | 18 | Ótimo exato para `10a..15c` |

Artefatos existentes: 4638 summaries, 4638 arquivos de evolução e 4638 arquivos de timing.

## Baseline Ótima (Brute-force)

Instâncias resolvidas exatamente por [[bruteforce]]:

| Instância | Makespan ótimo |
|---|---:|
| 10a | 8.251283 |
| 10b | 7.906130 |
| 10c | 9.545034 |
| 11a | 9.493458 |
| 11b | 8.823027 |
| 11c | 10.260222 |
| 12a | 9.500701 |
| 12b | 9.768963 |
| 12c | 10.479634 |
| 13a | 10.950258 |
| 13b | 11.473574 |
| 13c | 10.041955 |
| 14a | 11.556925 |
| 14b | 11.466929 |
| 14c | 11.486447 |
| 15a | 9.128680 |
| 15b | 12.184839 |
| 15c | 12.160210 |

## Qualidade da Solução — Gap vs Brute-force

Agregado sobre as 18 instâncias com ótimo (`10a..15c`) e 51 sementes por método:

| Método | Runs | Gap médio | Gap min | Gap max | Taxa global de ótimo |
|---|---:|---:|---:|---:|---:|
| **ACO** | 918 | **0.4296%** | 0.0000% | 9.4445% | **71.79%** |
| GA | 918 | 5.2814% | 0.0000% | 41.4267% | 26.47% |
| PSO | 918 | 22.2915% | 0.0000% | 76.5549% | 4.25% |

Interpretação descritiva: ACO apresentou melhor qualidade média nas instâncias com ótimo conhecido. GA encontrou ótimos em parte relevante das sementes, mas teve gap médio maior. PSO foi o método com pior gap médio neste desenho experimental.

## Qualidade da Solução — Instâncias Grandes

Comparação em `50a..100c`, com 51 sementes por método:

| Instância | ACO best | ACO média | GA best | GA média | PSO best | PSO média |
|---|---:|---:|---:|---:|---:|---:|
| 50a | 24.6688 | 25.6475 | 41.0120 | 45.3925 | 53.7651 | 59.7284 |
| 50b | 26.3528 | 27.4843 | 42.6487 | 48.6345 | 57.0330 | 64.3194 |
| 50c | 26.2372 | 27.0062 | 42.0122 | 46.9438 | 54.4114 | 62.3455 |
| 100a | 42.6718 | 45.3528 | 97.5921 | 110.6608 | 127.4034 | 136.5131 |
| 100b | 45.4689 | 48.0596 | 97.5343 | 110.4959 | 126.3967 | 135.9987 |
| 100c | 45.4858 | 46.9487 | 95.1632 | 111.3243 | 126.1838 | 138.1908 |

Interpretação descritiva: ACO apresentou makespan médio substancialmente menor que GA e PSO nas instâncias grandes auditadas. Em `100a..100c`, a média do ACO ficou em torno de 45–48, enquanto GA ficou em torno de 110–111 e PSO em torno de 136–138.

## Tempo de Otimização

Tempos médios de otimização em n=100, lidos dos arquivos `timing` associados aos summaries:

| Instância | ACO média ms | GA média ms | PSO média ms | Razão ACO/GA | Razão ACO/PSO |
|---|---:|---:|---:|---:|---:|
| 100a | 4262.06 | 38.82 | 76.24 | 109.78x | 55.91x |
| 100b | 4276.55 | 36.27 | 75.25 | 117.89x | 56.83x |
| 100c | 4266.55 | 39.37 | 73.84 | 108.37x | 57.78x |

Interpretação descritiva: GA foi muito mais rápido que ACO em n=100, mas com makespan médio muito pior. PSO também foi muito mais rápido que ACO, mas teve a pior qualidade média.

## Lower Bound AP

O lower bound AP é válido como limitante inferior: nas 18 instâncias com brute-force, `LB <= BF` em todos os casos. O bound é frouxo para o TSP-SD-ATP.

| Métrica | Valor |
|---|---:|
| Gap médio LB→BF | 51.36% |
| Gap mínimo | 40.79% |
| Gap máximo | 65.35% |

Interpretação: a redução 3D→2D via `min_i G[i][j][k]` remove contexto angular e torna o AP bound pouco informativo como aproximação do ótimo, embora ele permaneça matematicamente válido.

## Claims Liberados Para Escrita

| Claim | Força | Condição |
|---|---|---|
| ACO obteve melhor qualidade descritiva de solução nas instâncias avaliadas. | Forte descritivo | Não escrever “significativo” até script corrigido |
| GA foi muito mais rápido que ACO, especialmente em n=100. | Forte descritivo | Reportar trade-off qualidade-tempo |
| PSO teve desempenho inferior nesta variante específica. | Moderado | Delimitar à implementação random keys e parâmetros fixos |
| AP bound é válido mas frouxo. | Forte | Explicar que não é rota factível nem ótimo |

## Conexões

- [[experiment-pipeline]] — como os experimentos foram executados
- [[analysis-methodology]] — protocolo estatístico e pendências
- [[auditoria-codigo-dados-vault]] — fonte da cobertura e dos números auditados
- [[claims.base]] — claims permitidos e bloqueados
- [[ga]], [[pso]], [[aco]], [[bruteforce]], [[lower-bounds]] — métodos comparados

## Dados Brutos

- `src/data/results/summary/` — 4638 arquivos JSON
- `src/data/results/evolution/` — 4638 arquivos JSONL
- `src/data/results/timing/` — 4638 arquivos JSON
- `scripts/consolidate_results.py` — script de consolidação, ainda precisa auditoria antes de gerar tabelas finais
- `scripts/analise-estatistica.py` — em correção por outro agente
