---
name: scientific-figures
description: Create scientific figures, statistical plots, visual analyses, SVG/PNG artifacts, and publication-ready images for papers, monographs, theses, and LaTeX documents. Use when generating graphs, charts, plots, route/map images, visual analysis panels, or figure scripts for scientific documents.
---

# Scientific Figures

Use this skill to produce rigorous, publication-ready visual analyses as images for scientific documents. The visual standard follows Edward Tufte's principles from *The Visual Display of Quantitative Information*: maximize information density, remove non-data ink, make comparisons easy, and never let decoration compete with evidence.

## Core Rule

Every figure must answer a research question. If the image does not support a claim, reveal a pattern, compare alternatives, or document a method, do not generate it.

## Repository Workflow

- Prefer reproducible scripts under `scripts/` for generated figures.
- Prefer editable vector output (`.svg`, `.tex`, or `.pdf`) as the primary artifact.
- Also save `.png` previews when useful for quick inspection or web display.
- Save final monograph figures under `monografia/figs/`.
- Prefer text, captions, panel labels, explanatory notes, and multi-panel composition in LaTeX (`.tex`) rather than embedding them inside image files.
- Prefer rendering multiple small, focused image artifacts and arranging them afterward in `.tex`; the image-generation step should focus only on the visual content itself.
- For route/map visualizations already supported by the Go web server, use `tcc serve` and `/api/render` instead of duplicating rendering logic.
- If the server lacks the required route/map visualization, extend `src/web/render/` and wire it through `src/web/handlers/render.go` before creating duplicate plotting code.
- Python is appropriate for statistical charts and data analysis figures when it is simpler than Go rendering.
- Keep code, data, and text traceable: a figure used in `monografia/` must be reproducible from `src/data/results/`, `scripts/`, or the Go render endpoint.

## Tufte-Inspired Design Principles

- Maximize data-ink ratio: keep marks that encode data; remove borders, shadows, gradients, heavy grids, redundant legends, and decorative backgrounds.
- Avoid chartjunk: no 3D bars, textures, icons, ornamental colors, exploded pies, or effects that do not encode information.
- Prefer direct labeling over legends when labels fit cleanly near the data.
- Use small multiples when comparing the same structure across methods, instances, seeds, or iterations.
- Use consistent scales across panels when the comparison depends on magnitude.
- Use local rescaling only when it answers a different question, and label that choice explicitly.
- Reveal variation, not only central tendency: include confidence intervals, standard deviation, standard error, boxplots, violin plots, or individual points when appropriate.
- Show the data whenever feasible: for small or moderate samples, overlay points on summaries.
- Use muted, purposeful colors. Color must encode method, group, state, or uncertainty; it must not decorate.
- Use redundant encodings when the figure may be printed in grayscale: line style, marker shape, labels, or facet panels.
- Preserve quantitative honesty: axes, baselines, transformations, smoothing, exclusions, and normalization must be explicit.

## Figure Requirements

- Every graph must expose its scale.
- Every axis must have a label and units where applicable.
- Every multi-series figure must identify each method or group through direct labels or a legend.
- Every statistical chart must show sample size or make it recoverable from the caption/table.
- Every map/route image must include axes, coordinate ticks, or a scale indicator.
- Every figure must be readable at the size it will appear in the document.
- Every panel must have a clear title or panel label when used in a composite figure.
- Every file name must describe the content, not just the script step.
- If a composite image becomes dense or hard to read, split it into separate artifacts and use LaTeX to arrange them vertically, across pages, or as subfigures.

## Recommended Scientific Figure Types

| Question | Prefer | Avoid |
|---|---|---|
| Compare methods across instances | Dot plot, slopegraph, small multiples, ranked bar/dot chart | Pie chart, 3D bars |
| Show convergence over iterations | Line plot with uncertainty band or small multiples | One noisy line per seed without summary |
| Show distribution across seeds | Boxplot plus points, violin plus points, ECDF | Mean-only bar chart |
| Show runtime scaling | Log-scale line/dot plot with units and instance size | Unlabeled bars |
| Show route/map solution | Coordinate map with route, nodes, scale, method, instance, metric | Decorative drone/map background |
| Show relationships | Scatter plot with fit only if justified, residuals when needed | Smoothed curve without method description |
| Show process or algorithm | TikZ flowchart or concise schematic | Screenshot-style boxes with decorative arrows |

## Visual Analysis Workflow

### 1. Define the Claim

Before generating a figure, write one sentence:

`This figure supports the claim that ...`

Then choose the minimum visual form that tests or communicates that claim.

### 2. Inspect the Data

- Verify source files exist.
- Check units, metric definitions, missing values, repeated seeds, and run IDs.
- Confirm whether lower or higher values are better.
- Confirm whether comparisons should be absolute, relative, normalized, ranked, or statistical.

### 3. Choose Encodings

- Position on a common scale is the strongest encoding; prefer it for primary metrics.
- Length is acceptable for simple magnitudes.
- Color should encode categories, not ordered magnitude unless using a perceptually ordered colormap.
- Area, volume, and angle are weak encodings; avoid them for precise comparison.

### 4. Generate Reproducibly

- Put reusable constants at the top of the script: paths, figure size, colors, method order, labels, units.
- Write one function per panel when creating composite figures.
- For charts that are self-contained visual objects, save SVG first and PNG second.
- For route/map panels or dense multi-panel layouts, save individual image panels first and generate a `.tex` fragment that handles titles, panel labels, captions, notes, and layout.
- Avoid manual edits to generated images unless the source artifact also records the change.

## LaTeX-First Composition Standard

- Treat LaTeX as the preferred composition layer for scientific documents.
- Do not bake figure titles, long explanatory text, captions, ABNT source notes, or page-level typography into raster images.
- Generate the smallest useful image artifacts: one map, one route panel, one chart panel, or one schematic component per file unless a single combined image is demonstrably more readable.
- Compose panels in `.tex` using `figure`, `subfigure`/`subcaption` when available, `minipage`, or repeated figure environments when vertical layout is clearer.
- Let LaTeX control final width with commands such as `\includegraphics[width=0.9\textwidth]{...}` rather than hard-coding dense page typography into the image.
- Keep generated `.tex` fragments under `monografia/figs/` when they belong to a figure stack, and include a comment showing how to `\input{...}` them from the monograph.
- Generate PNG previews for quick inspection, but do not treat the preview as the typographic source of truth.
- If an image contains text that LaTeX could reasonably typeset, move that text into `.tex` unless it is an axis tick, direct data label, map scale, or visual encoding that must stay attached to the graphic.
- For route/map figures in this TCC, render each route/map panel through the Go endpoint as a separate PNG, then organize panel labels and captions in `.tex`.

### 5. Audit the Figure

Use the checklist below before finishing.

## Matplotlib Style Baseline

Use this as a starting point for Python figures, then adapt to the document.

```python
import matplotlib.pyplot as plt

METHOD_COLORS = {
    "bruteforce": "#4c4c4c",
    "ga": "#4c78a8",
    "pso": "#f58518",
    "aco": "#54a24b",
}

plt.rcParams.update({
    "figure.dpi": 120,
    "savefig.dpi": 300,
    "font.size": 9,
    "axes.titlesize": 10,
    "axes.labelsize": 9,
    "xtick.labelsize": 8,
    "ytick.labelsize": 8,
    "legend.fontsize": 8,
    "axes.spines.top": False,
    "axes.spines.right": False,
    "axes.grid": True,
    "grid.color": "#d9d9d9",
    "grid.linewidth": 0.5,
    "grid.alpha": 0.7,
    "legend.frameon": False,
})
```

## Composite Figure Standard

- Prefer one logical composite figure with panels `A`, `B`, `C` when panels support one argument, but compose it in `.tex` when readability or page layout benefits from separate panel images.
- Keep method order, colors, markers, and labels consistent across panels.
- Use shared axes where valid.
- Keep panel titles short and move detail to axis labels or caption.
- Export at the expected final size, commonly near full text width for monograph figures.
- Avoid wide, dense single-image composites when each panel will become too small at document size; split into individual panels and let LaTeX stack or paginate them.

## Captions And Self-Containment

A figure should be understandable without reading surrounding paragraphs. Ensure the figure or its caption identifies:

- Dataset or instance IDs.
- Methods compared.
- Metric and unit.
- Whether values are mean, median, best, final iteration, or all observations.
- Number of seeds/runs.
- Error band meaning, if present.
- Any filtering, normalization, or log scale.

## Statistical Honesty

- Do not use bar charts for distributions unless the data are naturally counts or proportions.
- Do not truncate axes to amplify small effects unless the truncation is explicit and justified.
- Do not smooth convergence or time-series curves unless the smoothing method and window are stated.
- Do not hide failed runs; encode or report them.
- Do not mix metrics with different units on one axis.
- Do not compare stochastic methods from single seeds unless the figure explicitly says it is illustrative.

## TSP And Drone Patrol Figures

For this repository's TCC context:

- Frame figures around patrol route quality, makespan, convergence, computational cost, and robustness across seeds.
- For route images, include instance ID, method, seed/run ID, iteration or final state, and objective value.
- For optimization comparisons, use method names consistently: `bruteforce`, `ga`, `pso`, `aco`.
- When comparing against brute force, show optimality gap as `%` when useful: `(method - optimum) / optimum * 100` for minimization metrics.
- For runtime, label units as `ms`, `s`, or `min` and use log scale only when orders of magnitude differ.

## Quality Checklist

- The figure supports a specific claim.
- The data source is real and documented.
- Axes, ticks, units, and legends/direct labels are present.
- Scale is visible for every plot or map.
- Colors are consistent, muted, and meaningful.
- Uncertainty or distribution is shown when reporting stochastic results.
- Visual clutter is removed without removing evidence.
- The image remains readable at document size.
- The output is reproducible from a checked-in script or render endpoint.
- SVG/PDF/TEX source and PNG preview are saved when appropriate.
- Text and composition are in `.tex` whenever that improves readability, typographic consistency, or maintainability.

## Anti-Patterns

- Decorative dashboards for academic figures.
- Mean-only bar charts for stochastic optimizer results.
- Figures without units or sample size.
- Legends that repeat information already visible in labels.
- Palettes that rely only on red/green contrast.
- Screenshots of plots instead of generated artifacts.
- Manual spreadsheet graphics without source code.
- Route maps without coordinate scale.
- Dense multi-panel SVG/PNG figures with titles, captions, and notes baked into the image when LaTeX could compose them more readably.
- Claims in the text that are stronger than what the figure shows.

## Related Skills

- Use `data-analysis` or `statistical-analysis` for statistical test selection and effect-size reporting.
- Use `academic-writing` when writing figure captions or discussing results in prose.
- Use `latex-document-skill` when integrating figures into LaTeX or compiling the monograph.
- Use `go` and `golang-cli` when extending the Go render endpoint.
