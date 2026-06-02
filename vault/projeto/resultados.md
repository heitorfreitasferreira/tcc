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
| **ACO** | **0.18%** | 0.00% | 3.78% | **76.86%** |
| **GA** | 4.36% | 0.00% | 31.57% | 29.41% |
| **PSO** | 19.02% | 0.00% | 57.89% | 5.10% |

Detalhamento: ACO encontra o ótimo em ~77% das sementes (instâncias 10-14). GA acerta ~29%. PSO raramente encontra o ótimo mesmo em n=10.

## Qualidade da Solução — Instâncias Grandes (50-100)

Comparação do melhor makespan entre os métodos (sem brute-force). 51 sementes:

| Instância | ACO (best) | ACO (média±dp) | GA (best) | GA (média±dp) | PSO (best) | PSO (média±dp) |
|-----------|:----------:|:--------------:|:---------:|:-------------:|:----------:|:--------------:|
| 50a | **24.67** | 25.65 ± 0.49 | 41.01 | 45.39 ± 2.76 | 53.77 | 59.73 ± 3.46 |
| 50b | **26.35** | 27.48 ± 0.66 | 42.65 | 48.63 ± 3.84 | 57.03 | 64.32 ± 4.52 |
| 50c | **26.24** | 27.01 ± 0.72 | 42.01 | 46.94 ± 3.10 | 54.41 | 62.35 ± 4.16 |
| 100a | **42.67** | 45.35 ± 0.98 | 97.59 | 110.66 ± 5.97 | 127.40 | 136.51 ± 5.07 |
| 100b | **45.47** | 48.06 ± 1.11 | 97.53 | 110.50 ± 6.80 | 126.40 | 136.00 ± 5.55 |
| 100c | **45.49** | 46.95 ± 1.03 | 95.16 | 111.32 ± 7.65 | 126.18 | 138.19 ± 6.62 |

ACO domina com folga: makespan **~2.4× menor que GA**, **~2.9× menor que PSO** em n=100 (médias).

## Tempo de Execução

Média sobre 51 sementes por instância:

| Instância | GA | PSO | ACO |
|-----------|:--:|:---:|:---:|
| 10a | 0.001s | 0.002s | 0.025s |
| 50a | 0.023s | 0.042s | 1.064s |
| 100a | 0.039s | 0.076s | 4.262s |

GA é **~110× mais rápido que ACO** em n=100, mas com makespan 2.4× pior. PSO é ~56× mais rápido que ACO.

## Trade-off Qualidade × Tempo

| Método | n=100 (makespan médio) | n=100 (tempo médio) | Gap vs ACO |
|--------|:---------------------:|:------------------:|:----------:|
| ACO | **46.79** | 4.268s | — |
| GA | 110.83 | **0.039s** | +137% |
| PSO | 136.90 | 0.076s | +193% |

Nenhum método domina: ACO para qualidade, GA para velocidade, PSO intermediário em ambos.

## Observações

- ACO: qualidade superior mas ~4s em n=100. Feromônio 3D e roleta por candidato são caros (evaporação O(N³)).
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
