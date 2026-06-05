# AGENTS.md

## Project Context

- This is a TCC/monografia repository comparing bio-inspired optimization methods over a TSP/rTSP base problem; the motivating scenario is patrol drones visiting points of interest with minimal route/time.
- Keep code, generated experiment data, analysis notebooks, and monograph text aligned: claims in `monografia/` should be supported by outputs under `src/data/results/` or analysis in `scripts/visualizacoes.ipynb`.

## Repository Shape

- `src/` is the Go module (`module tcc`, Go 1.23.7) and Cobra CLI entrypoint; run Go commands from `src/` or use `make -C src ...`.
- `src/data/` contains checked-in `.points`, `.graph`, and experiment results; the web server embeds `*.graph`, `*.points`, `results/summary`, `results/evolution`, and `results/timing`, but not `results/logs`.
- `monografia/` is the LaTeX monograph using `ppgco.cls`; chapter files are included by `main_ppgco_ufu.tex` from `cap_*` directories and references live in `monografia/bib/abntex2-references.bib`.
- `scripts/visualizacoes.ipynb` analyzes `src/data/results` and expects run file names like `<instance>__<method>__s...__h....json`.

## Image & Artifact Generation

- **Web server is the canonical source for visual analysis images.** The `tcc serve` command exposes `/api/render?map=X&run=Y&iteration=N` which returns `image/png` — pure Go rendering (no browser dependency). All analysis figures for the monograph should be generated through this endpoint when possible.
- Use `bash scripts/gerar-figuras.sh` to batch-generate the standard set of figures into `monografia/figs/`. The script starts the server, calls `/api/render` for each desired figure, then stops the server.
- **If the render endpoint lacks a visualization you need, extend it rather than duplicating logic.** Add new rendering functions in `src/web/render/` and wire them through `src/web/handlers/render.go`. Keep the server as the single source of truth for route/graph visualization.
- **Python is also acceptable** for generating images, plots, tables, or analysis artifacts when it is genuinely easier (e.g., statistical plots with matplotlib/seaborn, complex data transformations with pandas). Prefer adding such code to `scripts/` or the Jupyter notebook `scripts/visualizacoes.ipynb`. The key rule: **do not duplicate visualization logic** — if the web server already renders what you need, use it; if not, choose the tool (Go render endpoint or Python script) that best fits the task.
- When a new analysis/visualization is needed, first check if the web server already exposes the required data (evolution, summary, timing, graph structure) and consider adding a new render function or query parameter before reaching for a separate tool.
- **Every graph/plot must expose its scale.** Route/map images must include coordinate axes or a scale indicator; statistical charts must include labeled axes, units where applicable, tick labels, and a legend when multiple series/methods are shown. A figure without scale is not acceptable for the monograph.
- **Prefer text and composition in LaTeX.** When a figure needs titles, panel labels, explanatory text, captions, or multi-panel organization, render the smallest useful image artifacts separately and arrange them in `.tex` afterward. The image-generation step should focus on the image content itself, not on final page typography or dense composite layout.
- Academic figures should be specific and self-contained: title/caption-ready content, method names, instance IDs, units (`makespan`, `%`, `ms`, `n`), and visual encodings documented by labels or legends. Avoid decorative graphics that do not support a claim.
- Figures intended for the monograph must be written in Portuguese (pt-BR) and follow ABNT-style academic wording: use `Figura`, `Fonte`, concise descriptive titles, metric names in Portuguese when natural (`tempo de otimização`, `taxa de acerto`, `iteração`) and keep technical terms such as `makespan` only when they are the defined objective metric.
- **Flowcharts are preferably TikZ `.tex` artifacts**, even if not yet included by the monograph. Keep standalone flowcharts under `monografia/figs/` as `.tex` plus generated `.svg`/`.png` previews when useful; this makes final LaTeX integration and typography consistent.
- For generated charts, prefer keeping both source and output: script under `scripts/`, vector output (`.svg` or `.tex`) in `monografia/figs/`, and PNG preview for quick inspection.
- Publication-style scientific figures should use a shared visual standard rather than ad-hoc styling: fixed page-aware width (roughly full text width), consistent font sizes, muted method colors, no decorative backgrounds, no unnecessary top/right spines, and panel labels (`A`, `B`, `C`, ...). Prefer composite multi-panel figures over many isolated plots when the panels support one claim.
- Figure-generating code should be organized around reusable constants/helpers and per-panel drawing functions. Save SVG as the primary editable artifact and PNG as a preview; the final figure should be generated at the size it is expected to appear in the document.
- Use the local `scientific-figures` skill for any new figure stack or substantial figure revision. Its baseline is Tufte-style: each figure must support a specific research claim, avoid chartjunk, show distributions for stochastic methods, expose scale/units/sample sizes, and keep outputs reproducible from scripts or the Go render endpoint.
- `scripts/gerar-analises.sh` is the canonical batch command for final visual artifacts. It removes old generated figure outputs (preserving logos), builds/starts the Go renderer, generates TikZ method flowcharts, and creates the publication-ready composite figures.

## Go CLI Commands

- Build: `make -C src build` creates `src/tcc`; batch scripts require this executable before running.
- Focused tests: from `src/`, run `go test ./path/to/pkg -run TestName`; full tests: `go test ./...`.
- `make -C src test` runs `go mod tidy`, `go mod vendor`, then writes `coverage.out` and `report.json`; use it only when those side effects are acceptable.
- Lint target is `make -C src lint`, but it requires `golangci-lint` and runs `golangci-lint run --enable-all`.
- CLI examples: `./src/tcc create -s 42 -f ./src/data`, `./src/tcc optimize ga --instance ./src/data/10a.graph --results-dir ./src/data/results`, `./src/tcc serve --addr :8080`.

## Experiment Workflow

- Optimization methods are `bruteforce`, `ga`, `pso`, and `aco`; `ga`, `pso`, and `aco` use `--population` and `--iterations` plus method-specific flags.
- Structured outputs are written below `--results-dir` as `summary/<run_id>.json`, `evolution/<run_id>.jsonl`, and `timing/<run_id>.json`; `--if-exists` accepts `skip|overwrite|error` and defaults to `skip`.
- `src/run_all.sh` does not create graphs; it expects files matching the requested `--frequency` such as `10a.graph`, `10b.graph`, etc., and runs jobs through GNU `parallel`.
- `src/run_experiments_multi_seed.sh` intentionally rejects `bruteforce`; use `src/run_experiments_bruteforce_missing.sh` for missing brute-force baselines.
- Docker compose builds from `src/Dockerfile`; `docker compose run --rm cli --help` runs the CLI, and `docker compose run --rm run-all --method=ga ...` runs the batch entrypoint with `./src/data` mounted.

## Experiment Data Structure

See [AGENTS-experiments.md](./AGENTS-experiments.md) for the full experiment workflow, artifact schemas, coverage, analysis pipeline, and caveats for agents.

## Knowledge Base — Obsidian Vault (`vault/`)

### Project Skill (`.agents/skills/knowledge-base/SKILL.md`)
- `knowledge-base` — agent skill for vault expansion: search papers, validate DOIs, add BibTeX, import, enrich notes, update canvas. Full 8-step workflow with quality standards, anti-patterns, and example entry.

### Skills Installed

- `obsidian-markdown` — formatação markdown compatível com Obsidian (callouts, [[links]], etc.)
- `obsidian-vault` — operações do agente no vault (criar, ler, buscar notas)
- `obsidian-bases` — queries estruturadas sobre o vault
- `json-canvas` — criar/atualizar grafos de conhecimento `.canvas`
- `academic-researcher` — buscar e entender papers acadêmicos
- `academic-search` — estratégias de busca acadêmica

### MCPs Available

- `scihub` — baixar PDFs e metadados de papers
- `google-scholar` — buscar artigos acadêmicos

### Vault Structure

```
vault/
├── papers/           ← Uma nota markdown por artigo (19 existentes)
├── areas/            ← Notas sobre áreas de pesquisa
├── templates/        ← Template para criar novas notas
├── canvas/           ← Grafo de conhecimento (.canvas)
└── opencode-vault.md ← Instruções completas para agentes
```

### Agent Workflow for New References

1. Use `google-scholar` + `academic-search` skill to find papers
2. Use `scihub` to fetch metadata/PDF (ou `bash scripts/download-pdfs.sh` para lote com Unpaywall/Semantic Scholar/OpenAlex e validação `pdfinfo`)
3. Add entry to `monografia/bib/abntex2-references.bib`
4. Run `bash scripts/import-bib-to-vault.sh <bibtex-key>` to create note
5. Use `obsidian-markdown` skill to write resumo/contribuições
6. Link to existing areas with `[[wiki links]]`
7. Update canvas at `vault/canvas/tcc-knowledge-graph.canvas`

### Import Existing BibTeX

```bash
./scripts/import-bib-to-vault.sh                          # all entries
./scripts/import-bib-to-vault.sh kennedy1995particle      # single entry
BIB=path/to/file.bib ./scripts/import-bib-to-vault.sh     # custom .bib
```

## Academic MCP Servers (PDF Download & Literature Search)

New MCPs installed; each adds specific tools for academic workflows:

| MCP | Tools | Instalação |
|-----|-------|------------|
| `crossref` | `crossref_get_work`, `crossref_search_works`, `crossref_get_references`, `crossref_search_journals`, `crossref_search_funders` | `npm i -g @cyanheads/crossref-mcp-server` |
| `scholar-sidekick` | `resolveIdentifier`, `formatCitation`, `checkOpenAccess`, `checkRetraction`, `verifyCitation` | `npm i -g scholar-sidekick-mcp` |
| `semantic-scholar` | 16 tools: search, citations, authors, recommendations, OA PDF URLs | `uv tool install semantic-scholar-fastmcp` |
| `doiget` | `doiget_fetch_paper`, `doiget_resolve_paper`, `doiget_metadata_only`, `doiget_bibtex_export`, `doiget_batch_from_bibliography`, +9 more | `cargo install --path crates/doiget-cli` (from gh:sotashimozono/doiget) |
| `academic-search` | `search_papers`, `fetch_paper_details`, `search_by_topic` (Crossref + S2) | `scripts/academic-search-mcp-server.py` |

### OpenCode Commands

- `/baixar-pdf <doi|bibtex-key>` — download PDF via MCP cascade + extract summary
- `/consultar-academico <query>` — search multiple databases in parallel

### Helper Scripts

- `scripts/baixar-pdf-chave.sh <bibtex-key>` — CLI fallback for PDF download + summary extraction
- `scripts/download-pdfs.sh [--keys K] [--use-scihub]` — batch PDF download pipeline (Unpaywall → S2 → OpenAlex → Sci-Hub)

### Env Variables (optional, for higher rate limits)

```bash
export CROSSREF_MAILTO="heitor@ufu.br"           # polite Crossref pool
export SCHOLAR_API_KEY="ssk_..."                 # free key from scholar-sidekick.com
export SEMANTIC_SCHOLAR_API_KEY="..."             # higher S2 rate limits
```

## Monograph Notes

- For academic prose, write in Portuguese unless the target section explicitly requires English, and keep the problem framing as drone patrol/TSP comparison rather than a generic optimizer benchmark.
- Compile from `monografia/` with the BibTeX cycle when references change: `pdflatex main_ppgco_ufu.tex`, `bibtex main_ppgco_ufu`, then `pdflatex main_ppgco_ufu.tex` twice.
- The current chapter files still contain template text in places; verify claims against actual experiment artifacts before strengthening conclusions.
