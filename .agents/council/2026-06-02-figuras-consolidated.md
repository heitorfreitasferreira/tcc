# Council Report: Figures/Tables Plan Review

**Date:** 2026-06-02
**Target:** `vault/writing/figuras-tabelas-monografia.md`
**Mode:** validate (2 judges)

## Verdict: FAIL (HIGH confidence)

| Judge | Verdict | Confidence |
|-------|---------|------------|
| **Figura** (Scientific Figures Expert) | FAIL | HIGH |
| **Bench** (Literature Benchmark Analyst) | WARN | HIGH |

**Consensus rule applied:** Any FAIL → **FAIL**

---

## Cross-Judge Agreement Matrix

| Finding | Figura | Bench | Consolidated |
|---------|--------|-------|--------------|
| Script `analise-estatistica.py` inexistente (T10, T11, Fig30 bloqueados) | CRITICAL | flagged | **CRITICAL** |
| AP bound faltando para instâncias grandes (50a–100c) | — | CRITICAL | **CRITICAL** |
| Penalidade angular invisível nas rotas | — | CRITICAL | **CRITICAL** |
| Fig1 (cenário de patrulha) ausente | CRITICAL | — | **CRITICAL** |
| Bandas de incerteza ausentes nas curvas de convergência | MAJOR | — | **MAJOR** |
| Tabela best/mean/worst unificada ausente | — | MAJOR | **MAJOR** |
| Tabela de ranking de métodos ausente | — | MAJOR | **MAJOR** |
| Fig13 tipo não especificado | MAJOR | — | **MAJOR** |
| Eixos/unidades/escala não verificados no plano | MAJOR | — | **MAJOR** |
| Rotas usam semente única (s0) sem disclaimer | MAJOR | — | **MAJOR** |
| Performance profile (Dolan-Moré) ausente | — | MAJOR | **MAJOR** |
| Análise de sensibilidade a parâmetros ausente | — | MAJOR | **MAJOR** |
| Redundância de figuras de qualidade | — | MINOR | **MINOR** |
| Heatmap vs bar chart | MINOR | MINOR | **MINOR** |
| Boxplot sem pontos sobrepostos | MINOR | — | **MINOR** |
| Fig17 (last improvement) não padrão | — | MINOR | **MINOR** |
| T9 tendencioso (ACO como referência) | — | MINOR | **MINOR** |
| Pairwise dominance scatter ausente | — | MINOR | **MINOR** |
| Visualização bound tightness ausente | — | MINOR | **MINOR** |
| Fig16 filename pode não corresponder à descrição | MINOR | — | **MINOR** |

---

## Consolidated Findings

### Critical (must fix before monograph)

| ID | Finding | Fix | Judge(s) |
|----|---------|-----|----------|
| C01 | `scripts/analise-estatistica.py` não existe. T10 (Friedman), T11 (Nemenyi), Fig30 (CD diagram) dependem dele. Três artefatos bloqueados por pipeline faltante. | Criar `scripts/analise-estatistica.py` com Friedman, Nemenyi e CD diagram, OU remover T10/T11/Fig30 do plano e declarar análise descritiva. | Ambos |
| C02 | AP bound não é comparado com metaheurísticas para instâncias grandes (50a–100c). O roadmap lista "tabela de gap médio (AP bound como referência)" como requisito mínimo, mas o plano só tem T4 (bound vs BF para instâncias pequenas). | Adicionar tabela/figura do gap% (best-makespan − LB) / LB para todas as 30 instâncias, especialmente 50a–100c onde BF não existe. | Bench |
| C03 | Penalidade angular — característica definidora do TSP-SD-ATP — é invisível em todas as 9+ figuras de rota (Fig19–29). Um leitor familiar com TSP clássico não vê nada novo ou diferente. | Adicionar codificação visual (cor, espessura, ou anotação) destacando curvas fechadas em pelo menos uma rota representativa. Ou adicionar painel mostrando distribuição da penalidade angular por método. | Bench |
| C04 | Fig1 (diagrama do cenário de patrulha com drone) é "Pendente". Introdução não tem âncora visual para o problema motivador. | Criar diagrama TikZ: drone, POIs, rota de patrulha, anotação da penalidade angular. Salvar como `monografia/figs/fig-scenario-diagram.{tex,png,svg}`. | Figura |

### Major (strongly recommended)

| ID | Finding | Fix | Judge(s) |
|----|---------|-----|----------|
| M01 | Curvas de convergência (Fig16–18) não mencionam bandas de incerteza. Métodos estocásticos produzem trajetórias diferentes por semente — linha única sem variabilidade é anti-padrão. | Adicionar bandas sombreadas (mediana ± IQR ou ± 1σ). Especificar significado no texto. | Figura |
| M02 | Nenhuma tabela best/mean/worst unificada cobre todas as instâncias. T5 (gap min/max/médio) só para instâncias com BF; T6 (média) só para grandes. Literatura (Chandra, Wu, Halim) apresenta tabela única por instância. | Unificar T5–T6 em tabela best/mean/worst (ou gap%) por método×instância para 10a–100c. | Bench |
| M03 | Nenhuma tabela de ranking médio por método para complementar Friedman+CD. Demšar (2006) espera tabela de postos. | Adicionar tabela: mean rank por método, posicionada antes de T10/Fig30. | Bench |
| M04 | Fig13 tem tipo "gráfico" — ambíguo. "Distribuição de qualidade agregada (método × n)" não permite reprodução ou avaliação. | Especificar encoding exato: violin + pontos, boxplot + jitter, ou ECDF facetado por n. | Figura |
| M05 | Nenhuma descrição de figura no plano menciona eixos, unidades, ou escala. Critério de aceite exige "escala, unidade e legenda" mas não há verificação. | Adicionar coluna ao plano: métrica por eixo, unidades, escala log, e para rotas: escala de coordenadas ou scale bar. | Figura |
| M06 | Figuras de rota (Fig19–29) usam semente 0 (s0). Sem declaração se s0 é melhor/mediana/ilustrativa, o leitor não sabe se é representativa. | Declarar estratégia de seleção: melhor semente, mediana, ou "ilustrativo (s0)". Se ilustrativo, rotular explicitamente. | Figura |
| M07 | Performance profile (Dolan-Moré 2002) ausente. É a figura mais informativa para comparação multi-método em OR benchmarking e é padrão desde 2002. | Adicionar `fig-performance-profile` mostrando fração de instâncias resolvidas dentro de fator τ do best, τ ∈ [1, 1.5]. | Bench |
| M08 | Nenhuma análise de sensibilidade a parâmetros. Literatura recente (pós-2015) espera justificativa ou análise do impacto de parâmetros. | Adicionar seção (pode ser small multiples ou tabela única) mostrando impacto de pop/iter/parâmetros específicos em instância representativa (ex.: 30a). | Bench |

### Minor (nice to have)

| ID | Finding | Fix | Judge(s) |
|----|---------|-----|----------|
| m01 | Três figuras de qualidade (Fig9 heatmap + Fig12 boxplot + Fig13 distribution) sobrepostas. Literatura apresenta 1-2 figuras de qualidade. | Consolidar: manter Fig12 como principal, mover Fig9 para painel de Fig13, ou remover uma. | Bench |
| m02 | Heatmaps (Fig9–11) usam saturação de cor como encoding principal. Posição em escala comum é encoding mais forte (dot plot, slopegraph). | Justificar escolha de heatmap OU suplementar com dot plot. Garantir colormap perceptualmente uniforme (viridis) e colorblind-safe. | Ambos |
| m03 | Fig12 (boxplot) não especifica se pontos individuais por semente são sobrepostos. | Adicionar jittered points overlay. | Figura |
| m04 | Fig17 (última iteração com melhoria) não é padrão na literatura e tem interpretabilidade questionável para métodos estocásticos. | Substituir por bandas de confiança nas curvas de convergência OU tabela de iteração para atingir 90%/95%/99% do ótimo. | Bench |
| m05 | T9 (ACO/GA e ACO/PSO ratio em n=100) usa ACO como referência, presumindo que ACO é baseline. | Normalizar pelo mais rápido por instância ou apresentar tempos absolutos. | Bench |
| m06 | Sem scatter de dominância pareada (makespan método X vs Y por instância). Halim 2019 inclui similar. | Adicionar scatter matrix ou dominance chart para cada par de métodos. | Bench |
| m07 | Bound tightness do AP mostrado apenas como tabela (T4), sem visualização por tamanho de instância. | Adicionar line/scatter plot: gap% do AP bound em função de n (10–14) + bound vs best-known para n≥15. | Bench |
| m08 | Fig16 (fig-runtime-convergence): filename sugere "runtime" mas descrição diz "fração de avaliações × gap mediano". Verificar se o arquivo corresponde à descrição. | Verificar conteúdo real do arquivo. Se mostra iterações, renomear. | Figura |

---

## Summary for Monograph Writer

### O que está forte
- Infraestrutura de figuras robusta: maioria dos arquivos existe em PNG+SVG, script `consolidate_results.py` funcional, renderizador Go funcional
- Fluxogramas com fonte TikZ (`.tex`) — padrão acadêmico correto
- Figuras de rota abrangentes (30+ overlays + small multiples + painéis)
- Pipeline experimental e critérios de aceite bem definidos

### O que precisa ser criado (bloqueante)
1. **`scripts/analise-estatistica.py`** — sem ele, T10, T11, Fig30 não existem. Toda a seção de estatística fica inviável.
2. **Fig1** — diagrama do cenário de patrulha para a Introdução

### O que precisa ser adicionado ao plano
3. **Gap vs AP bound para instâncias grandes** (C02) — referência absoluta para n≥15
4. **Visualização da penalidade angular nas rotas** (C03) — o diferencial do TSP-SD-ATP
5. **Performance profile (Dolan-Moré)** (M07)
6. **Tabela best/mean/worst unificada** (M02) + **ranking table** (M03)

### O que precisa ser corrigido nos artefatos existentes
7. Adicionar bandas de incerteza nas curvas de convergência (M01)
8. Especificar tipo da Fig13 (M04)
9. Adicionar verificação de escala/unidades no plano (M05)
10. Rotular estratégia de seleção de semente nas rotas (M06)
11. Sobrepor pontos individuais no boxplot Fig12 (m03)

---

## Files Reviewed

- `vault/writing/figuras-tabelas-monografia.md` (target)
- `vault/writing/roadmap-monografia.md` (acceptance criteria)
- `monografia/figs/*` (artifact inventory)
- `scripts/consolidate_results.py` (pipeline)
- `scripts/analise-estatistica.py` (missing)

## Judges

- **Figura** — Scientific Figures Expert (loaded with `scientific-figures` skill)
- **Bench** — Literature Benchmark Analyst (familiar with Chandra 2022, Wu 2020, Halim 2019, Dorigo & Stützle 2004)
