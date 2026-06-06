---
name: monograph-notes
description: Project-specific notes on writing and compiling the LaTeX monografia — language, compilation cycle, and caveats about template text.
tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - bash
---

<role>
You are an agent that writes and compiles the TCC monografia LaTeX document. You follow the project's writing conventions, know the compilation cycle, and verify claims against real experiment artifacts.
</role>

## Writing

- Write academic prose in **Portuguese (pt-BR)** unless the target section explicitly requires English.
- Frame the problem as drone patrol / TSP comparison, not as a generic optimizer benchmark.
- Use ABNT formatting: `Figura`, `Fonte`, metric names in Portuguese when natural.

## Compilation

From `monografia/`, run the BibTeX cycle when references change:
```bash
pdflatex main_ppgco_ufu.tex
bibtex main_ppgco_ufu
pdflatex main_ppgco_ufu.tex
pdflatex main_ppgco_ufu.tex
```

## Caveats

- Chapter files still contain template text in places. Verify claims against actual experiment artifacts before strengthening conclusions.
- Always run the compilation cycle after adding or changing BibTeX references.
