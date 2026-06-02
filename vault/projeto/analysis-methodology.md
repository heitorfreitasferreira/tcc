---
tags: [projeto, implementacao, analise, estatistica, metodologia]
---

# Metodologia de Análise — Experimentos

## Pipeline de Dados

```
src/data/results/summary/*.json    (4635 runs)
src/data/results/timing/*.json     (4635 runs)
src/data/results/evolution/*.jsonl (4635 runs)
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
| Brute-force | 15 | 1 (s0) | 15 (10a..14c) |

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
- **Nemenyi**: p < 0.05. Calcular CD = q_α · √(k(k+1)/6N) onde k = 3 métodos, N = 30 instâncias, q_α (3, ∞) = 3.314.
- **Wilcoxon**: pareado por instância, com correção Bonferroni-Holm para 3 comparações (GA×PSO, GA×ACO, PSO×ACO).
- **Exclusão do lower bound**: LB é determinístico (1 execução, sem variabilidade) e não compete — é referência. Não entra nos testes.
- **Exclusão do brute-force**: apenas 15 instâncias, 1 seed. Subconjunto separado para validação do LB e gap analysis.

### Script de Análise

Implementado em `scripts/analise-estatistica.py` (a criar) usando `scipy.stats`:

```python
from scipy.stats import friedmanchisquare, wilcoxon
import numpy as np

# friedmanchisquare: recebe 3 arrays (GA, PSO, ACO) com 30 valores (mediana por instância)
# wilcoxon: pareado entre dois métodos sobre as mesmas 30 instâncias
# CD: calculado a partir da tabela de postos de Friedman
```

### Saídas esperadas

- Tabela com p-valor do Friedman
- Matriz de p-valores Nemenyi (ou ajuste Bonferroni-Holm para Wilcoxon)
- Diagrama CD exportado para `monografia/figs/diagrama-cd.png` e `.svg`
- Conclusão textual: "GA/PSO/ACO diferem significativamente (p < 0.05)" ou equivalente, apenas se os testes confirmarem

### Limitações

- Apenas 30 instâncias sintéticas — poder estatístico limitado
- 3 métodos apenas — CD de Nemenyi é conservativo para poucos tratamentos
- Parâmetros fixos (pop=100, iter=100) — a ordenação pode mudar sob tuning

## Mapeamento Claims → Dados

| Claim na monografia | Dado de suporte | Localização |
|--------------------|----------------|-------------|
| "ACO produz makespan médio menor que GA e PSO em todas as 30 instâncias" | Summary: mediana por instância | `src/data/results/summary/` |
| "ACO é significativamente mais lento que GA e PSO (p < 0.05)" | Timing + Friedman/Wilcoxon | `src/data/results/timing/` |
| "GA oferece melhor equilíbrio qualidade-tempo em instâncias grandes" | Scatter quality×time | `scripts/visualizacoes.ipynb` |
| "O lower bound AP é válido (≤ ótimo) mas frouxo (gap ~50%)" | Summary LB vs BF | `src/data/results/summary/` |
| "Diferenças entre métodos são estatisticamente significativas" | Friedman + Nemenyi | `scripts/analise-estatistica.py` |

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
