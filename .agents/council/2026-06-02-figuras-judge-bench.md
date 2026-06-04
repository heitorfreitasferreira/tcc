# Bench Judge — Literature Benchmark Analyst

**Target:** `vault/writing/figuras-tabelas-monografia.md`
**Date:** 2026-06-02
**Role:** Literature Benchmark Analyst — evaluates the figure/table plan against what comparable TSP metaheuristic comparison papers publish.

---

## Verdict JSON

```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "findings": [
    {
      "id": "f-001",
      "severity": "major",
      "finding": "Missing performance profile (Dolan-Moré plot) — standard in optimization benchmarking literature for comparing multiple solvers across many instances.",
      "fix": "Add a 'fig-performance-profile' showing the fraction of instances solved within factor τ of the best solution, for τ ∈ [1, 1.5].",
      "why": "Every comparative solver study in operations research (including TSP metaheuristic papers post-2002) includes a performance profile. It is the single most informative figure for multi-method comparison and its absence is a noticeable gap against the literature standard."
    },
    {
      "id": "f-002",
      "severity": "major",
      "finding": "No best/mean/worst makespan table covering ALL instances (small and large) for all methods.",
      "fix": "Extend T5-T6 into a unified table showing best/mean/worst makespan (or gap%) per method per instance, covering the full range 10a–100c.",
      "why": "Chandra 2022, Wu 2020, and Halim 2019 all present tables with best/mean/worst values per method per instance. The current plan splits this into T5 (gap min/max for BF instances) and T6 (mean only for large instances), which obscures the variability information for large instances and the absolute makespan values."
    },
    {
      "id": "f-003",
      "severity": "major",
      "finding": "No method ranking table (mean rank across instances) to complement Friedman/CD diagram.",
      "fix": "Add a table showing the mean rank of each method across instances (both for quality gap and for runtime), positioned before T10-Fig30.",
      "why": "Demšar (2006) framework expects rank tables alongside the CD diagram. Halim 2019 includes explicit ranking. Without the rank table, Fig30 (CD diagram) is disconnected from the underlying data and the reader cannot verify the ranking that produced the critical differences."
    },
    {
      "id": "f-004",
      "severity": "minor",
      "finding": "Fig17 (last improvement iteration) is non-standard and has questionable interpretability for stochastic methods.",
      "fix": "Consider replacing Fig17 with a convergence robustness chart (e.g., confidence bands around median convergence curves, or a table showing iteration at which each method reaches 90%/95%/99% of final solution).",
      "why": "Wu 2020, Chandra 2022, Dorigo & Stützle 2004 do not present 'last improvement' plots. For stochastic methods with plateau behavior, 'last improvement' can be dominated by random chance rather than meaningful convergence behavior. This figure risks confusing readers without adding clear value."
    },
    {
      "id": "f-005",
      "severity": "minor",
      "finding": "Over-reliance on heatmaps (Fig9-Fig11) versus bar/column charts preferred in reference papers.",
      "fix": "Replace at least one of Fig9 or Fig10 with a grouped bar chart showing makespan per method per instance size (n=10, 14, 15, 20, 30, 50, 100). Keep heatmaps as supplementary.",
      "why": "Chandra 2022, Wu 2020, Halim 2019 all use bar charts for quality comparisons. Heatmaps show patterns well but obscure precise values. A bar chart makes exact comparisons across methods at a given instance size immediately visible. Heatmaps are acceptable but the literature standard is bar charts."
    },
    {
      "id": "f-006",
      "severity": "minor",
      "finding": "T9 (ACO/GA and ACO/PSO runtime ratio at n=100) uses ACO as the reference, which presumes ACO is the baseline method rather than treating all methods symmetrically.",
      "fix": "Replace T9 with a general pairwise runtime ratio table, or normalize to the fastest method per instance (ratio over best), preserving the data without methodological favoritism.",
      "why": "The plan should not imply ACO is the reference method before presenting results. Comparative literature (Chandra, Wu, Halim) normalizes to best per instance or presents absolute times."
    },
    {
      "id": "f-007",
      "severity": "minor",
      "finding": "No pairwise method dominance scatter plot (method X makespan vs method Y makespan per instance).",
      "fix": "Add a scatter matrix or dominance chart: for each pair of methods, plot the makespan (or gap%) on one axis vs the other, one point per instance. Show diagonal where methods tie.",
      "why": "This is a standard diagnostic in comparative studies (Halim 2019 includes similar). It immediately shows which instances favor which method and whether dominance is systematic or instance-dependent."
    },
    {
      "id": "f-008",
      "severity": "minor",
      "finding": "Lower bound tightness is shown only as a table (T4); no visualization across instance sizes.",
      "fix": "Add a simple line/scatter plot showing AP bound gap% (bound vs BF optimal) as a function of n (10–14), and extend with AP bound vs best-known for larger instances.",
      "why": "A figure showing bound tightness degradation with n is more informative than a static table. It would support the narrative of when the bound is useful vs when it becomes loose."
    },
    {
      "id": "f-009",
      "severity": "major",
      "finding": "No parameter sensitivity or ablation analysis — literature increasingly expects this for stochastic methods.",
      "fix": "Add a brief ablation section (can be small multiples or a single table): show impact of population size, iterations, or key method-specific parameters on makespan for at least one representative instance (e.g., 30a).",
      "why": "While T3 lists parameter values, comparative papers (notably more recent ones post-2015) are expected to justify or analyze parameter choices. Without this, the reader cannot assess whether results are sensitive to arbitrary parameter settings. This is a growing expectation, not yet universal, hence MAJOR rather than CRITICAL."
    },
    {
      "id": "f-010",
      "severity": "critical",
      "finding": "No visualization of angular/turning cost in route overlays — the defining feature of TSP-SD-ATP is invisible in all 9 route figures.",
      "fix": "Add callouts or color-coded segments in at least one route overlay highlighting sharp turns (high angular penalty). Alternatively, add a small panel showing the angular penalty distribution of each method's solution.",
      "why": "The problem is defined by angular penalty; route overlays that show only paths without any indicator of turning costs fail to communicate what makes TSP-SD-ATP different from standard TSP. This is a critical omission because the monograph's entire contribution rests on this problem variant, yet the main visual evidence (route maps) does not depict it."
    },
    {
      "id": "f-011",
      "severity": "critical",
      "finding": "No table or figure comparing the stochastic methods against the AP lower bound for large instances (50a–100c).",
      "fix": "Add a table or figure showing the gap% of each method against the AP lower bound for all instances. This is essential because brute-force optima are unavailable for large instances.",
      "why": "The roadmap states 'tabela de gap médio (AP bound como referência)' as a minimum requirement. The current plan has T4 (AP vs BF for small instances) and T5 (gap vs BF), but no AP-bound comparison for large instances. Without it, the reader has no absolute quality reference for n≥15 results, making claims of 'good quality' unsubstantiated."
    },
    {
      "id": "f-012",
      "severity": "minor",
      "finding": "Redundancy cluster: Fig9 (median heatmap) + Fig12 (boxplot) + Fig13 (aggregated distribution) all show makespan quality in different forms, potentially overwhelming the reader.",
      "fix": "Consolidate: keep Fig12 (boxplot) as the primary quality figure (it shows distribution per method×instance); move Fig9 to the aggregated distribution Fig13 as a panel; or drop one of them.",
      "why": "Three separate quality figures risk section-bloat. Most papers (Chandra, Wu, Halim) present 1-2 quality figures: a main comparison chart + a distribution/variability figure. The plan has 5 quality figures (Fig9-13) which is excessive."
    }
  ]
}
```

---

## Detailed Analysis

### 1. Overall Assessment

The figure/table plan for Chapter 4 is **substantial and well-structured**, covering 30 labeled artifacts (9 tables + 21 figures) plus composite route panels. Most figures are already marked "Pronto" and traceable to the Go render endpoint or Python scripts, which is excellent.

However, from the **literature benchmark perspective**, there are two **critical gaps**, three **major omissions**, and several **minor issues** that would prevent this plan from matching the standard published in comparable TSP metaheuristic comparison papers.

**Verdict: WARN** — the plan is solid but has identifiable gaps against the literature standard that need addressing before it is monograph-ready.

---

### 2. Literature Gap Analysis

**Covered adequately (literature-aligned):**
- ✅ Baseline tables (T1-T4): instances, BF optima, parameters, AP vs BF — standard
- ✅ Gap summary table (T5): min/mean/max gap — standard (Halim, Chandra)
- ✅ Runtime scalability (Fig14): log-scale — standard
- ✅ Quality×time scatter (Fig15): standard (less common but valuable)
- ✅ Convergence curves (Fig16, Fig18): essential for stochastic methods (Wu, Dorigo)
- ✅ Statistical testing planned (T10, T11, Fig30): follows Demšar 2006 framework
- ✅ Route visualizations (Fig19-29): comprehensive, including small multiples
- ✅ Boxplots (Fig12): increasingly expected for stochastic methods

**Partially covered (literature-aligned but incomplete):**
- ⚠️ Quality comparison (Fig9-11, Fig13): heatmaps are a valid choice, but the literature standard (Chandra Fig3, Wu Fig2, Halim Fig1-2) uses bar/column charts for easier cross-method comparison at a specific instance size. Heatmaps are better for pattern recognition across many cells, worse for precise quantitative comparison.
- ⚠️ Best/mean/worst tables: only T5 (gap min/mean/max for small instances) and T6 (mean only for large instances). No unified best/mean/worst table for all instances.
- ⚠️ Convergence: three separate figures (Fig16, Fig17, Fig18) — Fig17 (last improvement) is non-standard and questionable.

**Missing entirely (gap vs literature):**
- ❌ **Performance profile (Dolan-Moré)** — standard since 2002 for comparing multiple solvers across many instances. Every major TSP comparison study since uses this. Its absence is the single most noticeable gap.
- ❌ **Method ranking table** — Halim 2019 presents explicit ranking tables. The Friedman test + CD diagram (Fig30) is incomplete without an intermediate ranking table showing mean ranks.
- ❌ **AP bound comparison for large instances** — the roadmap lists this as a minimum requirement (gap médio vs AP bound), but the plan only has T4 (small instances). No figure/table compares methods against AP bound for n≥15.
- ❌ **Angular penalty visualization** — the defining feature of TSP-SD-ATP is invisible in all route overlays.

---

### 3. Comparison with Key Papers

#### vs Chandra et al. (2022)

Chandra presents:
- Instance details table ✅ (T1)
- Parameter settings table ✅ (T3)
- Makespan results table (best/mean/worst) ❌ partial (T5-T6 cover partially)
- Bar chart of makespan per method ❌ uses heatmap instead — acceptable but different
- Convergence curves ✅ (Fig16, Fig18)
- Route maps for selected instances ✅ (Fig19-29)
- ANOVA + Tukey post-hoc ✅ (Friedman + Nemenyi planned)

**Chandra advantage:** simpler, more readable quality figures (bar charts). The plan's heatmaps are richer but less accessible.

#### vs Wu et al. (2020)

Wu presents:
- Best/mean/worst makespan per method per instance ❌ partial (only mean for large)
- Convergence curves ✅
- Bar charts of quality ✅ heatmap substitute
- Route maps for selected instances ✅

**Wu advantage:** cleaner presentation of results. No extraneous figures.

#### vs Halim et al. (2019)

Halim presents:
- Makespan table (6 heuristics × instances) ⚠️ partial
- Gap% table (explicit percentage from optimal) ✅ (T5)
- Runtime comparison table ✅ (T8)
- Ranking table ❌ missing
- Bar charts of gap% ✅ heatmap substitute

**Halim advantage:** explicit ranking, simpler quality figures. Halim does not include convergence curves or route maps, so the plan is stronger in those dimensions.

---

### 4. Missing Figure Types (Literature Standard vs Plan)

| Figure Type | Chandra | Wu | Halim | Dorigo | This Plan |
|------------|---------|----|-------|--------|-----------|
| Performance profile | ❌ | ❌ | ❌ | ❌ | ❌ |
| Best/mean/worst table | ✅ | ✅ | ✅ | — | ⚠️ partial |
| Ranking table | ❌ | ❌ | ✅ | — | ❌ |
| Method dominance scatter | ❌ | ❌ | ✅ | — | ❌ |
| Quality bar chart | ✅ | ✅ | ✅ | — | ⚠️ heatmap |
| Convergence curves | ✅ | ✅ | ❌ | ✅ | ✅ |
| Route visualizations | ✅ | ✅ | ❌ | ✅ | ✅ |
| Statistical CD diagram | ❌ | ❌ | ❌ | — | ✅ planned |
| Boxplot/violin | ❌ | ❌ | ❌ | — | ✅ |
| Angular penalty viz | N/A | N/A | N/A | N/A | ❌ |
| Parameter sensitivity | ❌ | ❌ | ❌ | ❌ | ❌ |

**Key insight:** While many papers omit performance profiles, they have become the expected standard in operations research benchmarking (Dolan & Moré 2002, cited 3000+ times). For a monograph comparing 4+ methods across 15+ instances, a performance profile is the single best summary figure. Its omission is the most significant gap.

---

### 5. Acceptance Criteria Evaluation

**"Todos os números têm fonte rastreável em src/data/results/"**
- ✅ T1-T2 (instances, BF optima) — generated from results
- ✅ T4 (AP bound vs BF) — generated
- ✅ T5-T9 — generated by consolidate_results.py
- ✅ All figures — generated by render server
- ⚠️ T10-T11, Fig30 (pending) — depend on analise-estatistica.py, not yet created

**"Gráficos têm escala, unidade e legenda"**
- The plan states this as a requirement. Cannot verify directly from plan but it's mandated.

**"Testes estatísticos são reportados apenas se executados"**
- ✅ T10/T11/Fig30 are marked Pendente, conditional on script execution.

**"Conclusões são proporcionais aos dados brutos"**
- Planning concern, not figure-specific.

**Overall AC assessment:** The acceptance criteria are well-addressed in the plan design. The main gap is that analise-estatistica.py does not exist yet, making the entire statistics section (T10, T11, Fig30) at risk.

---

### 6. Specific Finding Justifications

#### Finding f-010 (Critical — Angular Penalty Invisible)

The TSP-SD-ATP problem differs from standard TSP precisely because of the angular penalty term in the cost function (Fig2 visualizes this). Yet all route overlays (Fig19-27) show only path traces — there is no visual encoding of turning angles. This means:

1. The reader cannot see which methods produce smoother routes (lower angular cost).
2. The defining contribution of the problem variant is invisible in the primary visual evidence.
3. A reader familiar with standard TSP would see these route figures and see nothing new or different.

**Fix:** At minimum, add color-coded turning points (e.g., red for sharp >90° turns, amber for moderate, green for smooth) in one representative overlay. Better: add a small panel to the route figure showing the angular penalty distribution per method's solution.

#### Finding f-011 (Critical — AP Bound Missing for Large Instances)

The roadmap explicitly requires "tabela de gap médio (AP bound como referência)" as a minimum figure/table. The plan has T4 (AP vs BF for small instances only) but nothing for 50a-100c where BF is unavailable. This means:

1. Quality claims for large instances have no absolute reference point.
2. The reader cannot assess whether a 5% gap from best-found is good (if the bound is tight) or meaningless (if the bound is loose).
3. T6 only shows makespan values, not gap vs any reference.

**Fix:** Add a column or separate table showing (best-found makespan / AP bound) for each large instance, or a figure plotting this ratio across all instance sizes.

---

### 7. Recommendations (Priority Order)

**Critical (must fix before monograph):**

1. **Add angular penalty visualization to at least one route overlay** (f-010)
2. **Add AP bound gap table/figure for large instances** (f-011)

**Major (strongly recommended):**

3. **Add performance profile (Dolan-Moré)** (f-001)
4. **Add unified best/mean/worst table** (f-002)
5. **Add method ranking table** (f-003)
6. **Add parameter sensitivity analysis** (f-009)

**Minor (nice to have or redesign):**

7. Add method dominance scatter plot (f-007)
8. Add bound tightness visualization (f-008)
9. Replace one heatmap with bar chart (f-005)
10. Fix T9 reference bias (f-006)
11. Consider dropping Fig17 or replacing with confidence bands (f-004)
12. Consolidate quality figures to reduce redundancy (f-012)

---

### 8. Conclusion

The figure/table plan is **ambitious and largely complete** — it covers more ground (30+ artifacts) than typical comparative TSP papers (typically 8-12 figures/tables). The render server approach, small multiples, and statistical framework are strengths.

However, **two critical blind spots** (invisible angular penalty in route figures, missing AP bound comparison for large instances) mean the plan fails to communicate what is novel about this specific study. Additionally, **three major omissions** (performance profile, best/mean/worst table, ranking table) leave gaps against the literature standard.

The WARN verdict reflects that the plan is structurally sound but requires targeted additions to meet the benchmark set by comparable published work. A PASS would require addressing at least f-010 and f-011.
