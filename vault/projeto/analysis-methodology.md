---
tags: [projeto, implementacao, analise, estatistica, metodologia]
status: estatistica-validada
updated: 2026-06-02
---

# Metodologia de Análise — Experimentos

## Pipeline de Dados

```
src/data/results/summary/*.json    (4638 runs)
src/data/results/timing/*.json     (4638 runs)
src/data/results/evolution/*.jsonl (4638 runs)
       │
       ▼
scripts/consolidate_results.py     →  LaTeX tables + summary stats
scripts/analise-estatistica.py     →  Friedman + Nemenyi + Wilcoxon
scripts/visualizacoes.ipynb        →  gráficos e análises
scripts/gerar-graficos-estatisticos.py → figuras de publicação
       │
       ▼
vault/projeto/resultados.md        →  síntese para monografia
```

### Cobertura

| Método | Runs | Sementes | Instâncias |
|--------|-----:|----------|------------|
| GA | 1530 | 51 (s0..s50) | 30 (10a..100c) |
| PSO | 1530 | 51 (s0..s50) | 30 (10a..100c) |
| ACO | 1530 | 51 (s0..s50) | 30 (10a..100c) |
| Lower bound | 30 | 1 (s0) | 30 (10a..100c) |
| Brute-force | 18 | 1 (s0) | 18 (10a..15c) |

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
- **Gap vs ótimo/BF**: `(best − optimal) / optimal × 100%` (apenas instâncias com brute-force)
- **Gap vs LB**: `(best − LB) / LB × 100%` (todas as instâncias; LB por construção ≤ ótimo)
- **Tempo de otimização**: `durations_ms.optimize` (milissegundos)
- **Iterações até convergência**: última iteração com melhoria significativa

### Por método (agregado sobre 51 sementes)
- **Média**, **mediana**, **desvio padrão** do best makespan
- **Taxa de sucesso** (instâncias pequenas): fração de sementes que encontram o ótimo
- **Tempo médio de execução**

## Protocolo de Análise Estatística

Referência metodológica: Demšar (2006) — *Statistical Comparisons of Classifiers over Multiple Data Sets*.

### Justificativa

Com 3 métodos estocásticos × 30 instâncias × 51 sementes, comparar apenas medianas é insuficiente. O protocolo segue a recomendação de Demšar (2006) para comparação de múltiplos algoritmos sobre múltiplos conjuntos de dados, usando testes não-paramétricos (não assume normalidade).

### Testes

| Etapa | Teste | O que responde |
|-------|-------|---------------|
| 1 | **Friedman** | Rejeita H₀: "todos os métodos são equivalentes" |
| 2 | **Nemenyi post-hoc** | Se Friedman rejeitar, mostra quais pares diferem com significância (usando diferença crítica CD) |
| 3 | **Diagrama CD** | Visualização da diferença crítica entre postos médios |
| 4 | **Wilcoxon signed-rank** | Confirmação pareada por instância para cada par de métodos |

### Decisões de Implementação

- **Friedman**: usar mediana do makespan por instância como valor-resumo (1 valor por método×instância). Isso dá 30 blocos (instâncias), 3 tratamentos (métodos).
- **Nemenyi**: p < 0.05. Calcular CD = q_α · √(k(k+1)/6N) onde k = 3 métodos, N = 30 instâncias. Usar q_α da Tabela 5a de Demšar (2006) para a fórmula CD: q_{3,0.05} = 2.343 (já dividido por √2). Não usar o valor bruto da Studentized range (3.314), que infla o CD por √2.
- **Wilcoxon**: pareado por instância, com correção Bonferroni-Holm para 3 comparações (GA×PSO, GA×ACO, PSO×ACO).
- **Exclusão do lower bound**: LB é determinístico (1 execução, sem variabilidade) e não compete — é referência. Não entra nos testes.
- **Exclusão do brute-force**: apenas 18 instâncias, 1 seed. Subconjunto separado para validação do LB e análise de gap vs ótimo.

### Script de Análise

Implementado em `scripts/analise-estatistica.py` (stdlib puro, sem dependências externas):

```python
# Friedman: implementação manual com ranks médios por instância (mediana como valor-resumo)
# Iman-Davenport: F_F = (N-1)·χ² / (N·(k-1) - χ²)
# Nemenyi: CD = q_α · √(k(k+1)/6N) com q_α da Tabela 5a de Demšar
# Wilcoxon: implementação manual com distribuição exata (DP) para n ≤ 30
# Holm: correção sequencial de Bonferroni para m comparações
```

### Saídas Validadas em 2026-06-02

| Resultado | Valor |
|---|---:|
| N | 30 instâncias |
| k | 3 métodos |
| Rank médio ACO | 1.1000 |
| Rank médio GA | 1.9000 |
| Rank médio PSO | 3.0000 |
| χ²_F de Friedman | 54.6000 |
| F_F de Iman-Davenport | F(2,58) = 293.2222 |
| p-valor Iman-Davenport | 4.710129e-31 |
| CD Nemenyi (α=0.05, q=2.343) | 0.6050 |

Todas as comparações Nemenyi são significativas: ACO×GA (Δrank=0.8000), ACO×PSO (Δrank=1.9000) e GA×PSO (Δrank=1.1000). O Wilcoxon signed-rank pareado com correção Bonferroni-Holm também rejeita H₀ nos três pares: ACO×GA (W=1.0, n=26, p=5.960464e-08), ACO×PSO (W=0.0, n=30, p=1.862645e-09) e GA×PSO (W=0.0, n=30, p=1.862645e-09).

O diagrama CD foi regenerado em `monografia/figs/cd-diagram.svg` e `monografia/figs/cd-diagram.png`.

### Limitações

- Apenas 30 instâncias sintéticas — poder estatístico limitado
- 3 métodos apenas — CD de Nemenyi é conservativo para poucos tratamentos
- Parâmetros fixos (pop=100, iter=100) — a ordenação pode mudar sob tuning

## Mapeamento Claims → Dados

| Claim na monografia | Dado de suporte | Localização |
|--------------------|----------------|-------------|
| "ACO produz makespan médio menor que GA e PSO em todas as 30 instâncias" | Summary: mediana por instância | `src/data/results/summary/` |
| "ACO é significativamente mais lento que GA e PSO (p < 0.05)" | Timing + Friedman/Wilcoxon | `src/data/results/timing/` |
| "GA oferece melhor tempo computacional, mas ACO oferece melhor qualidade de solução" | Scatter quality×time + timing | `src/data/results/summary/`, `src/data/results/timing/` |
| "O lower bound AP é válido (≤ ótimo) mas frouxo (gap ~50%)" | Summary LB vs BF | `src/data/results/summary/` |
| "Diferenças entre métodos são estatisticamente significativas" | Friedman/Iman-Davenport + Nemenyi + Wilcoxon/Holm | `scripts/analise-estatistica.py`, `monografia/figs/cd-diagram.svg` |

## Conexões

- [[resultados]] — síntese dos números
- [[experiment-pipeline]] — como os dados foram gerados
- [[comparative-studies]] — metodologia de comparação na literatura
- [[demsar2006statistical]] — referência do protocolo estatístico
- [[chandra2022comparative]] — exemplo de ANOVA+Tukey para TSP (nota: este TCC usa Demšar, não ANOVA)

## Código-fonte

- `scripts/consolidate_results.py` — consolidação e tabelas LaTeX
- `scripts/analise-estatistica.py` — testes estatísticos (Friedman, Nemenyi, Wilcoxon, CD)
- `scripts/visualizacoes.ipynb` — análise visual
- `scripts/gerar-graficos-estatisticos.py` — figuras de publicação
