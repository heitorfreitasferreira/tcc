---
tags: [projeto, implementacao, analise, estatistica, metodologia]
---

# Metodologia de Análise — Experimentos

## Pipeline de Dados

```
src/data/results/summary/*.json    (4605 runs)
src/data/results/timing/*.json     (4605 runs)
src/data/results/evolution/*.jsonl (4605 runs)
       │
       ▼
scripts/consolidate_results.py     →  LaTeX tables + summary stats
scripts/visualizacoes.ipynb        →  gráficos e análises
       │
       ▼
vault/projeto/resultados.md        →  síntese para monografia
```

## Notebook de Análise

`scripts/visualizacoes.ipynb` — Jupyter Notebook com:

### DataFrames principais
- `summary_df`: uma linha por execução (instância, método, seed, best_makespan, params)
- `evolution_df`: uma linha por ponto de melhoria (iteração, makespan, delta)
- `timing_df`: tempos de carregamento, otimização e serialização
- `runs_df`: mesclagem dos anteriores com métricas derivadas

### Visualizações implementadas

| Visualização | O que mostra |
|-------------|-------------|
| Heatmap de cobertura | Quais instância×método foram executados |
| Heatmap de makespan | Melhor makespan por instância×método + média por n |
| Tempo médio por n | Linha (escala log) + boxplot por método |
| Scatter qualidade×tempo | Trade-off por instância (ou mediana agregada) |
| Curva de convergência | Gap mediano ao final vs fração de avaliações |
| Gap vs bruteforce | Heatmap + boxplot do gap percentual |
| Mapas de rota | Visualização geográfica das melhores rotas |

## Métricas

### Por execução
- **Best makespan**: menor custo encontrado (objetivo principal)
- **Gap vs ótimo**: `(best - optimal) / optimal × 100%` (apenas instâncias com brute-force)
- **Tempo de otimização**: `durations_ms.optimize` (milissegundos)
- **Iterações até convergência**: última iteração com melhoria significativa

### Por método (agregado sobre 51 sementes)
- **Média**, **mediana**, **desvio padrão** do best makespan
- **Taxa de sucesso** (instâncias pequenas): fração de sementes que encontram o ótimo
- **Tempo médio de execução**

## Testes Estatísticos

Atualmente implementados em `scripts/consolidate_results.py` (geração de tabelas LaTeX). Previstos para o notebook:

- **Shapiro-Wilk**: normalidade dos resíduos
- **ANOVA one-way** (ou **Kruskal-Wallis** se não-normal): diferenças entre métodos
- **Post-hoc Tukey HSD**: pares com diferença significativa

Referência metodológica: [[chandra2022comparative]] — ANOVA + Tukey para comparação de metaheurísticas em TSP.

## Mapeamento Claims → Dados

| Claim na monografia | Dado de suporte | Localização |
|--------------------|----------------|-------------|
| "ACO encontra o ótimo em todas as 51 sementes para n≤12" | Summary: `10a..12c` × `aco` × `s0..s50` | `resultados.md` |
| "GA é 75× mais rápido que ACO em n=100" | Timing: `100*__ga__*.json` vs `100*__aco__*.json` | `resultados.md` |
| "GA encontra o ótimo em ~50% das sementes para n≤12" | Summary: `10a..12c` × `ga` × `s0..s50` | `resultados.md` |

## Conexões

- [[resultados]] — síntese dos números
- [[experiment-pipeline]] — como os dados foram gerados
- [[comparative-studies]] — metodologia de comparação na literatura
- [[chandra2022comparative]] — referência de ANOVA+Tukey para TSP

## Código-fonte

- `scripts/consolidate_results.py` — consolidação e tabelas LaTeX
- `scripts/visualizacoes.ipynb` — análise visual
