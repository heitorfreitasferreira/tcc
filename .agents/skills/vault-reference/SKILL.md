---
name: vault-reference
description: Project-specific Obsidian vault structure, installed skills, basic MCPs, and the standard workflow for importing and enriching new academic references.
tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - bash
---

<role>
You are an agent that works with the TCC Obsidian vault. You know the vault directory layout, the available skills and MCPs, and the step-by-step workflow for importing new papers, enriching notes, and maintaining the knowledge graph.
</role>

## Project skill

`knowledge-base` (`.agents/skills/knowledge-base/SKILL.md`) — vault expansion: paper search, DOI validation, BibTeX import, note enrichment, canvas updates. Full 8-step workflow.

## Installed skills

- `obsidian-markdown` — Obsidian-compatible markdown (callouts, [[links]])
- `obsidian-vault` — vault operations (create, read, search notes)
- `obsidian-bases` — structured queries over the vault
- `json-canvas` — create/update `.canvas` knowledge graphs
- `academic-researcher` — find and understand academic papers
- `academic-search` — academic search strategies

## Available basic MCPs

- `scihub` — download PDFs and metadata
- `google-scholar` — search academic articles

## Vault structure

```
vault/
├── papers/           ← One markdown note per paper
├── areas/            ← Research area notes
├── templates/        ← Note templates
├── canvas/           ← Knowledge graph (.canvas)
└── opencode-vault.md ← Full instructions for agents
```

## New reference workflow

1. Use `google-scholar` + `academic-search` skill to find papers
2. Use `scihub` to fetch metadata/PDF (or `bash scripts/download-pdfs.sh` for batch)
3. Add entry to `monografia/bib/abntex2-references.bib`
4. Run `bash scripts/import-bib-to-vault.sh <bibtex-key>` to create note
5. Use `obsidian-markdown` skill to write summary/contributions
6. Link to existing areas with `[[wiki links]]`
7. Update canvas at `vault/canvas/tcc-knowledge-graph.canvas`

## Import existing BibTeX

```bash
./scripts/import-bib-to-vault.sh                          # all entries
./scripts/import-bib-to-vault.sh kennedy1995particle      # single entry
BIB=path/to/file.bib ./scripts/import-bib-to-vault.sh     # custom .bib
```
