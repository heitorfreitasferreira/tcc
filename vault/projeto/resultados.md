---
tags: [projeto, implementacao, resultados, experimentos]
---

# Resultados dos Experimentos

Dados consolidados de **4605 execuções** (51 sementes × 3 métodos × 30 instâncias + 15 brute-force). Gerados pelo pipeline em [[experiment-pipeline]].

## Cobertura

| Instâncias | Nós | Brute-force | GA | PSO | ACO |
|-----------|-----|:-----------:|:--:|:---:|:---:|
| 10a, 10b, 10c | 10 | 1 seed | 51 | 51 | 51 |
| 11a..14c | 11-14 | 1 seed | 51 | 51 | 51 |
| 15a..20c | 15-20 | — | 51 | 51 | 51 |
| 30a..50c | 30-50 | — | 51 | 51 | 51 |
| 100a, 100b, 100c | 100 | — | 51 | 51 | 51 |

## Baseline Ótima (Brute-force)

Instâncias pequenas resolvidas exatamente por [[bruteforce]]:

| Instância | Makespan ótimo |
|-----------|:--------------:|
| 10a | 8.251 |
| 10b | 7.906 |
| 10c | 9.545 |
| 11a | 9.493 |
| 11b | 8.823 |
| 11c | 10.260 |
| 12a | 9.501 |
| 12b | 9.769 |
| 12c | 10.480 |
| 13a | 10.950 |
| 13b | 11.474 |
| 13c | 10.042 |
| 14a | 11.557 |
| 14b | 11.467 |
| 14c | 11.486 |

## Qualidade da Solução — Gap vs Brute-force (instâncias 10-14)

Média dos gaps percentuais entre a melhor solução encontrada e o ótimo, sobre 51 sementes:

| Método | Gap médio | Gap min | Gap max | Taxa de acerto do ótimo |
|--------|:---------:|:-------:|:-------:|:-----------------------:|
| **ACO** | **0.00%** | 0.00% | 0.00% | **100%** (10a-13c), ~60% (14) |
| **GA** | 0.04% | 0.00% | 0.58% | ~50% (10a-13c), ~3% (14) |
| **PSO** | 1.57% | 0.00% | 6.55% | ~15% (10), ~5% (11-12), 0% (13-14) |

Detalhamento: ACO encontra o ótimo em **100% das sementes** para instâncias até 13c. GA encontra em ~50% das sementes. PSO raramente encontra o ótimo mesmo em n=10.

## Qualidade da Solução — Instâncias Grandes (50-100)

Comparação do melhor makespan entre os métodos (sem brute-force). 51 sementes:

| Instância | ACO (best) | GA (best) | PSO (best) | ACO (média±dp) |
|-----------|:----------:|:---------:|:----------:|:--------------:|
| 50a | **24.37** | 41.01 | 53.77 | 25.86 ± 0.49 |
| 50b | **26.78** | 42.65 | 57.03 | 27.69 ± 0.61 |
| 50c | **25.76** | 42.01 | 54.41 | 27.42 ± 0.75 |
| 100a | **45.64** | 97.59 | 127.40 | 47.53 ± 0.96 |
| 100b | **48.22** | 97.53 | 126.40 | 50.31 ± 0.89 |
| 100c | **44.94** | 95.16 | 126.18 | 48.48 ± 1.24 |

ACO domina com folga: makespan **~2× menor que GA**, **~2.7× menor que PSO** em n=100.

## Tempo de Execução

Média sobre 51 sementes por instância:

| Instância | GA | PSO | ACO |
|-----------|:--:|:---:|:---:|
| 10a | 0.002s | 0.003s | 0.024s |
| 50a | 0.026s | 0.037s | 0.971s |
| 100a | 0.038s | 0.075s | 2.859s |

GA é **~75× mais rápido que ACO** em n=100, mas com makespan 2× pior. PSO é ~38× mais rápido que ACO.

## Trade-off Qualidade × Tempo

| Método | n=100 (makespan) | n=100 (tempo) | Gap vs ACO |
|--------|:----------------:|:-------------:|:----------:|
| ACO | **47.53** | 2.859s | — |
| GA | 110.66 | **0.038s** | +133% |
| PSO | 136.51 | 0.075s | +187% |

Nenhum método domina: ACO para qualidade, GA para velocidade, PSO intermediário em ambos.

## Observações

- ACO: qualidade superior mas 2-3s mesmo em n=100. Feromônio 3D e roleta por candidato são caros.
- GA: ótimo custo-benefício. Encontra ótimo em instâncias pequenas em frações de segundo. Crossover OX + mutação swap são eficientes.
- PSO: random keys + atualização contínua não se adaptam bem ao TSP-SD-ATP. A decodificação por ordenação perde informação posicional.

## Conexões

- [[experiment-pipeline]] — como os experimentos foram executados
- [[analysis-methodology]] — protocolo estatístico e notebook de análise
- [[ga]], [[pso]], [[aco]], [[bruteforce]] — métodos comparados
- [[comparative-studies]] — contexto na literatura

## Dados Brutos

- `src/data/results/summary/` — 4605 arquivos JSON
- `src/data/results/evolution/` — 4605 arquivos JSONL (histórico de convergência)
- `src/data/results/timing/` — tempos de execução
- `scripts/consolidate_results.py` — script de consolidação
- `scripts/visualizacoes.ipynb` — notebook de análise visual
