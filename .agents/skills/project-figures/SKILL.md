---
name: project-figures
description: Project-specific conventions for generating figures, charts, and visual artifacts for the TCC monografia — render server, batch scripts, style guidelines, and preferred tooling.
tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - bash
---

<role>
You are an agent that generates publication-ready figures for the TCC monografia. You follow project conventions: prefer the Go render server for route/graph images, accept Python for statistical plots, and always include scale, labels, and legend. You compose multi-panel figures in LaTeX rather than cramming everything into one image.
</role>

## Canonical source

The **web server** (`tcc serve`) is the canonical source for visual analysis images. It exposes `/api/render?map=X&run=Y&iteration=N` returning `image/png` — pure Go rendering, no browser dependency. All analysis figures should use this endpoint when possible.

- `bash scripts/gerar-figuras.sh` batch-generates standard figures into `monografia/figs/`. It starts the server, calls `/api/render` for each figure, then stops.
- **If the render endpoint lacks what you need, extend it** in `src/web/render/` and wire through `src/web/handlers/render.go`. Keep the server as single source of truth.
- **Python is acceptable** for statistical plots (matplotlib/seaborn) or complex transforms (pandas). Add code to `scripts/` or `scripts/visualizacoes.ipynb`. Do not duplicate visualization logic.
- First check if the web server already exposes the required data before building a separate tool.

## Style rules

- **Every plot must expose its scale.** Route images: coordinate axes or scale indicator. Charts: labeled axes, units, tick labels, legend for multiple series.
- **Prefer LaTeX composition.** Render minimal image artifacts and arrange panels in `.tex` afterward.
- Figures must be self-contained: method names, instance IDs, units (`makespan`, `%`, `ms`, `n`), documented visual encodings.
- Monograph figures in **Portuguese (pt-BR)**, ABNT style: `Figura`, `Fonte`, metric names in Portuguese when natural.
- **Flowcharts: TikZ `.tex`** under `monografia/figs/` with `.svg`/`.png` previews.
- Keep source scripts under `scripts/`, vector output (`.svg`/`.tex`) in `monografia/figs/`, PNG for quick inspection.
- Publication style: fixed width, consistent font sizes, muted method colors, no decorative backgrounds, minimal spines, panel labels (`A`, `B`, `C`...).
- Figure code organized around reusable constants/helpers and per-panel drawing functions.
- Use the generic `scientific-figures` skill for new figure stacks or substantial revisions.
- `scripts/gerar-analises.sh` is the canonical batch command for final visual artifacts.
