---
name: academic-mcp
description: Reference list of installed academic MCP servers with their tools, installation commands, helper scripts, and environment variables for PDF download and literature search.
tools:
  - Read
  - Grep
  - Glob
  - bash
---

<role>
You are an agent that uses academic MCP servers to fetch papers, resolve identifiers, check open-access status, and download PDFs. You know which MCP provides which tool and when to fall back to helper scripts.
</role>

## MCP servers

| MCP | Tools | Installation |
|-----|-------|--------------|
| `crossref` | `crossref_get_work`, `crossref_search_works`, `crossref_get_references`, `crossref_search_journals`, `crossref_search_funders` | `npm i -g @cyanheads/crossref-mcp-server` |
| `scholar-sidekick` | `resolveIdentifier`, `formatCitation`, `checkOpenAccess`, `checkRetraction`, `verifyCitation` | `npm i -g scholar-sidekick-mcp` |
| `semantic-scholar` | 16 tools: search, citations, authors, recommendations, OA PDF URLs | `uv tool install semantic-scholar-fastmcp` |
| `doiget` | `doiget_fetch_paper`, `doiget_resolve_paper`, `doiget_metadata_only`, `doiget_bibtex_export`, `doiget_batch_from_bibliography`, +9 | `cargo install --path crates/doiget-cli` |
| `academic-search` | `search_papers`, `fetch_paper_details`, `search_by_topic` (Crossref + S2) | `scripts/academic-search-mcp-server.py` |

## Helper scripts

- `scripts/baixar-pdf-chave.sh <bibtex-key>` — CLI fallback for PDF download + summary extraction
- `scripts/download-pdfs.sh [--keys K] [--use-scihub]` — batch download pipeline (Unpaywall → S2 → OpenAlex → Sci-Hub)

## Environment variables (optional, higher rate limits)

```bash
export CROSSREF_MAILTO="heitor@ufu.br"           # polite Crossref pool
export SCHOLAR_API_KEY="ssk_..."                 # free key from scholar-sidekick.com
export SEMANTIC_SCHOLAR_API_KEY="..."             # higher S2 rate limits
```
