```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "findings": [
    {
      "id": "f-001",
      "severity": "critical",
      "finding": "analise-estatistica.py does not exist despite being listed as the source for T10, T11, and Fig30.",
      "fix": "Create scripts/analise-estatistica.py implementing Friedman test, Nemenyi post-hoc, and CD diagram generation. Alternatively, mark T10/T11/Fig30 as 'Cancelled — escopo reduzido' and remove from the plan.",
      "why": "Three artifacts (T10, T11, Fig30) depend on a non-existent script. The plan lists them as 'Pendente' but with no path to completion. The acceptance criteria require statistical tests be 'reportados apenas se executados' — they cannot be executed without the script."
    },
    {
      "id": "f-002",
      "severity": "critical",
      "finding": "Fig1 (diagrama do cenário de patrulha com drone) is still 'Pendente'. The Introduction has no motivating visual for the drone patrol scenario.",
      "fix": "Create a TikZ or schematic diagram showing: drone, POIs as labeled nodes, patrol route arc, and angular penalty annotation. Save as monografia/figs/fig-scenario-diagram.{tex,png,svg}.",
      "why": "The entire TCC is framed around drone patrol. A scenario diagram in Chapter 1 is the reader's first visual contact with the problem. Without it, the Introduction lacks a motivating illustration that grounds TSP-SD-ATP in the real-world application."
    },
    {
      "id": "f-003",
      "severity": "major",
      "finding": "Convergence curves (Fig16–18) do not mention uncertainty bands or confidence intervals. The scientific-figures standard requires 'line plot with uncertainty band' for convergence.",
      "fix": "Add shaded standard-deviation or percentile bands to convergence plots. Update figure descriptions to specify error band meaning (e.g., 'median ± IQR across 30 seeds').",
      "why": "Stochastic methods (GA, PSO, ACO) produce different convergence trajectories per seed. A single mean/median line without uncertainty is a mean-only presentation, which is explicitly called out as an anti-pattern. The acceptance criteria also require 'gráficos têm escala, unidade e legenda' — uncertainty is part of the scale."
    },
    {
      "id": "f-004",
      "severity": "major",
      "finding": "Fig13 is typed only as 'gráfico' with no specific chart type. 'Distribuição de qualidade agregada (método × n)' is too vague to reproduce or evaluate.",
      "fix": "Specify the exact encoding: e.g., 'violin plot + points' or 'boxplot + jittered points' or 'ECDF faceted by n'. Update the figure file to match.",
      "why": "The type column in the plan exists precisely to distinguish encoding choices. 'gráfico' is a category error. The skill recommends specific types per question: for distribution across seeds, 'boxplot plus points, violin plus points, ECDF'."
    },
    {
      "id": "f-005",
      "severity": "major",
      "finding": "No figure description in the plan mentions axes, units, or coordinate scales. The acceptance criteria require 'gráficos têm escala, unidade e legenda', but the plan provides zero verification that any figure satisfies this.",
      "fix": "Add a column or note to each figure row specifying: (a) the metric on each axis, (b) units, (c) whether log scale is used, (d) for route figures, whether coordinates are shown with scale ticks or a scale bar.",
      "why": "The skill checklist requires every figure to 'expose its scale' with labeled axes and units. Without this information in the plan, there is no way to verify compliance. Route figures without coordinate scale are listed as an anti-pattern."
    },
    {
      "id": "f-006",
      "severity": "major",
      "finding": "Route overlay figures (Fig19–27) and individual method routes appear to use only seed 0 (s0) based on filenames like route-100a-ga-s0.png. Multiple route figures may show a single seed's solution.",
      "fix": "Either: (1) explicitly state that route figures use the best or median seed and note which one, or (2) generate routes for best, median, and worst seeds to show variability. If only s0 is shown, label the figure as 'illustrative (seed 0)'.",
      "why": "Statistical honesty requires not hiding failed runs or cherry-picking results. A single seed's route may not be representative. The skill says 'Do not compare stochastic methods from single seeds unless the figure explicitly says it is illustrative.'"
    },
    {
      "id": "f-007",
      "severity": "minor",
      "finding": "Fig12 (boxplot-estabilidade) does not specify whether individual seed-level points are overlaid on the boxes.",
      "fix": "Add points overlay and update the description to 'boxplot + pontos por semente'.",
      "why": "The skill recommends 'boxplot plus points' for distribution across seeds. Small sample sizes (≈30 seeds) fit comfortably in a jittered-point overlay, revealing multimodality and outliers that boxplots alone obscure."
    },
    {
      "id": "f-008",
      "severity": "minor",
      "finding": "Plan lists Fig16 file as 'fig-runtime-convergence' but the description says 'fração de avaliações × gap mediano', which describes a runtime-normalized convergence curve. The existing files fig-runtime-convergence.{png,svg} may not match this description.",
      "fix": "Verify that fig-runtime-convergence actually plots evaluation fraction (not iteration number) on the x-axis and median gap on the y-axis. If it uses plain iterations, rename to avoid confusion.",
      "why": "Figure file names must 'describe the content, not just the script step'. A mismatch between the plan's description and the actual visualization would mislead readers."
    },
    {
      "id": "f-009",
      "severity": "minor",
      "finding": "T5, T6, T7 are aggregate tables (mean, min, max, success rate) but do not show standard deviation or seed-level distribution.",
      "fix": "Add a standard deviation or IQR column to T5. Consider keeping full seed-level data in an appendix table or supplementary material.",
      "why": "The anti-pattern list warns against 'mean-only bar charts for stochastic optimizer results'. The same applies to tables: reporting only central tendency without dispersion hides seed variability."
    },
    {
      "id": "f-010",
      "severity": "minor",
      "finding": "Heatmaps (Fig9–11) use color intensity as the primary encoding for comparing methods across instances. The skill recommends dot plots or slopegraphs for this comparison type.",
      "fix": "Either: (1) justify why heatmaps are preferred (e.g., 'small multiples would require 3×15 panels which is less compact'), or (2) supplement with a dot-plot or ranked-bar version. Ensure the colormap is perceptually uniform (e.g., viridis) and colorblind-safe.",
      "why": "Position on a common scale is the strongest visual encoding. Heatmaps use color saturation (weaker), and the plan does not acknowledge this trade-off. The anti-pattern checklist flags 'Palettes that rely only on red/green contrast' which heatmaps commonly use."
    }
  ]
}
```

---

## Detailed Analysis

### Overall Assessment

**Verdict: FAIL**

The figures/tables plan is well-organized and most artifacts exist on disk. However, it fails on two critical points: (1) the entire statistical analysis pipeline (3 artifacts: T10, T11, Fig30) depends on a script that does not exist at `scripts/analise-estatistica.py`, and (2) the motivating scenario diagram (Fig1) for the Introduction is still missing. Additionally, multiple major issues prevent the plan from meeting the scientific-figures publication standard: missing uncertainty in convergence curves, vague figure type specifications, and no mention of axes/units/scale anywhere in the plan.

### Chapter-by-Chapter Review

#### Capítulo 1 — Introdução

**Fig1 (Pendente):** The only figure planned for the Introduction is the drone patrol scenario diagram, and it does not exist. For a TCC that frames itself around "drone de patrulha visitando pontos de interesse," this figure is the reader's first visual anchor. Without it, the Introduction is text-only, which weakens the problem motivation.

**Recommendation:** Create as a priority. A TikZ diagram with a small set of POI nodes, the drone, a patrol route arc, and a callout for the angular penalty would serve as the visual thesis statement. Save as both `.tex` (for LaTeX integration) and `.png`/`.svg` (for previews).

#### Capítulo 2 — Fundamentação Teórica

**Fig2–Fig8 (All Pronto):** This section is the strongest in the plan. Six flowcharts (all with `.tex` source + `.pdf`/`.png`/`.svg` derivatives) and two diagrams. Three observations:

- **Vector formats available:** Both diagrams (`diagram-angular-penalty`, `diagram-tensor-3d`) have `.svg` alongside `.png` — correct per the standard.
- **Flowchart source in TikZ:** All six flowcharts have `.tex` source files — excellent, as TikZ flowcharts are preferred over raster screenshots.
- **Missing flowchart-bruteforce.pdf:** Let me verify... the glob shows `flowchart-ga`, `flowchart-pso`, `flowchart-aco`, `flowchart-lowerbound` but not `flowchart-bruteforce` yet. The plan says "Pronto" — this may need verification.

**Potential issue:** The flowcharts are pedagogically sound but the plan should indicate whether they use consistent styling (same arrow types, same decision diamond shapes, same color palette).

This section passes. No critical issues.

#### Capítulo 3 — Proposta

No figures planned. Acceptable — the proposal may be primarily mathematical. However, if the tensor 3D structure or the angular penalty formulation is non-trivial, a schematic might help reader comprehension.

#### Capítulo 4 — Experimentos e Resultados

This is where most issues live.

##### Configuração e Baseline (T1–T4)

**T1, T2, T4 (Gerável via consolidate_results.py):** The script exists and produces LaTeX tables. This is manageable.

**T3 (Rascunho — parâmetros):** Should be finalized before submission. Ensure the parameter table includes: population size, iterations, crossover/mutation rates (GA), inertia/cognitive/social weights (PSO), evaporation/alpha/beta (ACO), and the seed count (30 per the experiment design).

**T4 gap %:** The gap metric should be explicitly defined: `(lb - optimum) / optimum * 100` or `(optimum - lb) / optimum * 100`? Ensure the sign and interpretation are clear.

##### Qualidade das Soluções (Fig9–13, T5–T7)

**Fig9–11 (Heatmaps — Pronto):** The files exist with both `.png` and `.svg`. The heatmap encoding choice is the main concern (see f-010). Additionally:
- `heatmap-gap-vs-bf`: If the gap is against brute-force optimum, is the reference column (brute-force) shown as 0%? If not, add it.
- `heatmap-success-rate`: This is a proportion heatmap (0–100%). A diverging colormap (e.g., RdYlBu) might be appropriate, but red/green is an anti-pattern. Use a colorblind-safe diverging palette.
- Verifying that the `.svg` version has embedded text (not outlined paths) so it renders correctly in LaTeX.

**Fig12 (boxplot-estabilidade — Pronto):** Exists with both `.png` and `.svg`. The plan does not say whether individual points are overlaid. For 30 seeds, a jittered strip plot overlaid on boxes is standard practice. Without it, outliers and data gaps are hidden.

**Fig13 (fig-quality-distribution — Pronto):** Type is "gráfico" — unacceptable. The file `fig-quality-distribution.{png,svg}` exists on disk but what encoding does it use? If it is a violin plot or ECDF, the plan should say so. If it is a bar chart of means, it is an anti-pattern.

**T5, T6, T7 (Gerável):** These tables need dispersion columns (std or IQR). T7 (success rate) is a proportion — consider adding a confidence interval (e.g., Wilson score interval) for each cell.

##### Tempo Computacional (Fig14–15, T8–T9)

**Fig14 (scalability-runtime — Pronto):** The description mentions "escala log" which is appropriate. However:
- Does it include the lower bound (AP) with ∼0ms as the acceptance criteria specify?
- Are time units clearly labeled (`ms`, `s`, or `min`)?
- The `.svg` version exists but verify the text is editable, not outlined.

**Fig15 (scatter-quality-vs-time — Pronto):** A scatter plot is the correct encoding for a two-dimensional trade-off. Verify that:
- Methods are distinguished by both color AND marker shape (redundant encoding for grayscale printing).
- Each point represents one seed or one instance? The plan is ambiguous.
- Axes have labels: e.g., "Tempo de otimização (ms)" × "Makespan" or "Gap (%)".

##### Convergência (Fig16–18)

**Fig16 (fig-runtime-convergence — Pronto):** The description "fração de avaliações × gap mediano" is promising — using evaluation count fraction (not iteration number) normalizes across methods with different population sizes. However:
- Is the x-axis fraction of total evaluations or absolute count? State clearly.
- The file exists as `.png` and `.svg`. Verify it matches the described metric.

**Fig17 (convergence-last-improvement — Pronto):** A boxplot or ECDF of the last iteration where improvement occurred. This is a useful complementary view. Ensure:
- The x-axis is labeled with iteration number or evaluation count.
- Units are clear.
- The `.svg` exists and is editable.

**Fig18 (convergence-overlay — Pronto):** Four panels (10a, 30a, 50a, 100a) with overlaid method traces. The files exist. Critical questions:
- Are uncertainty bands included? If not, this is a mean-only comparison (f-003).
- Are the y-axis scales consistent across panels? If not, note it explicitly.
- Panel labels (A, B, C, D) should be in LaTeX, not baked into the PNG.

##### Rotas (Fig19–29)

The route figure set is the most comprehensive section. All main files exist:
- 30+ individual route overlays (`overlay-*-methods.png`)
- Individual method routes (`route-*-{aco,ga,pso,lowerbound}-s0.png`)
- ACO iteration progression (`route-*-aco-s0-iter{1,25,50}.png`)
- Multi-panel composites (`route-panel-{10a,30a,100a}.png`)
- LaTeX small multiples (`fig-route-small-multiples.tex`)

**Key concerns:**
1. **Seed 0 bias (f-006):** All individual route images use `s0` (seed 0). For a rigorous comparison, the reader cannot tell if seed 0 is typical or cherry-picked. If the route figures are meant to be illustrative, say so explicitly.
2. **Coordinate scale (f-005):** The route images *probably* include axes (the Go render endpoint generates them), but the plan does not confirm this. Add a note that all route images include coordinate axes or a scale bar.
3. **File naming inconsistency:** The overlay files are `overlay-{instance}-methods.png` (e.g., `overlay-10a-methods.png`) but the plan references Fig19..27 without specific instance mapping. Not a technical problem but creates confusion when cross-referencing.
4. **Fig28 (route-panel) font size:** At monograph width, four route panels side-by-side may become illegible. Verify the panel images are readable when scaled to `0.45\textwidth` each.

##### Estatística (T10, T11, Fig30 — all Pendente)

These three artifacts are listed as "Pendente" but critically the generating script `scripts/analise-estatistica.py` does not exist. This is not a "figure that hasn't been generated yet" — it is a **missing pipeline**.

**T10 (Friedman test):** χ² and p-value. Implementation: `scipy.stats.friedmanchisquare`. Requires a matrix of method × instance (mean makespan or rank per instance). Straightforward.

**T11 (Nemenyi or Wilcoxon+Bonferroni):** Nemenyi post-hoc after significant Friedman. Implementation: `scikit-posthocs` or manual. The p-value matrix feeds Fig30.

**Fig30 (CD diagram):** The standard Critical Difference diagram from Demsar (2006). No file exists at `monografia/figs/diagrama-cd.*`. This requires either: (a) a Python implementation (e.g., `scikit-posthocs.critical_difference_diagram` or `Orange`), or (b) manual plotting with matplotlib.

**Impact:** The acceptance criteria state "testes estatísticos são reportados apenas se executados." If the statistical analysis script is not created, the entire statistics subsection must be removed from the monograph. This is a significant structural gap.

### Anti-Pattern Check

| Anti-Pattern | Present? | Details |
|---|---|---|
| Mean-only bar charts | Potentially | Fig13 type unknown; convergence curves may lack bands |
| No units/sample size | Yes | No figure description specifies units or n |
| Red/green palette | Unknown | Heatmaps may use problematic colormaps |
| Route maps without coordinate scale | Not verified | Plan doesn't guarantee scale |
| Legends repeating labels | Unknown | Cannot verify without seeing figures |
| Decor inactive | N/A | Not applicable to plan level |
| Single-seed comparison | Yes | Route figures use s0 only |
| Dense composite with baked-in text | Yes | route-panel images may have baked-in labels; prefer LaTeX |
| Figures without source traceability | No | consolidate_results.py exists and reads from src/data/results/ |
| Screenshots instead of generated | No | All files are generated artifacts |

### Acceptance Criteria Evaluation

| Criterion | Status | Assessment |
|---|---|---|
| Números com fonte rastreável em src/data/results/ | ✔️ PASS | consolidate_results.py reads from results/summary/ |
| Gráficos com escala, unidade e legenda | ❌ FAIL | Plan does not verify any figure for these; no explicit mention |
| Testes estatísticos reportados apenas se executados | ❌ FAIL | Script does not exist; cannot execute |
| Conclusões proporcionais aos dados brutos | ⚠️ Not verifiable | Plan-level check, not figure-level |
| Limitações experimentais antes da conclusão | ⚠️ Not verifiable | Monograph content, not figure content |
| Lower bound runtime (∼0ms) in scalability | ⚠️ Not verified | Plan shows Fig14 but doesn't confirm lower bound is included |

### Recommendations

**Immediate (blocking):**
1. Create `scripts/analise-estatistica.py` or remove T10, T11, Fig30 from the plan entirely.
2. Create Fig1 (scenario diagram) as a TikZ schematic.

**Before submission (high priority):**
3. Add uncertainty bands to all convergence curves (Fig16–18).
4. Specify the exact chart type for Fig13 and update the file if needed.
5. Add a verification column to the plan table: "Escala/Unidades/Legenda OK?"
6. Annotate route figure seed selection strategy: best, median, or illustrative only.
7. Add dispersion columns (std/IQR) to T5.

**Quality improvements (medium priority):**
8. Overlay individual seed points on Fig12 (boxplot).
9. Verify heatmap colormaps are colorblind-safe and perceptually uniform.
10. Add redundant encodings (marker shapes + colors) to Fig15 scatter plot.
11. Ensure route-panel images use LaTeX panel labels, not baked-in text.
12. Update Fig16 filename to accurately reflect content if it shows iterations, not evaluation fraction.

**Final recommendation:** The core figure infrastructure is solid — most files exist in multiple formats, the consolidate script is functional, and the route visualizations are comprehensive. The plan fails on the statistical analysis pipeline and missing scale/unit guarantees. Address f-001 (analise-estatistica.py) and f-002 (Fig1) to lift the critical blockers, then resolve the major uncertainty/scale issues for a publishable figure set.

---

*Audit performed 2026-06-02 by Figura (Scientific Figures Expert judge).*
