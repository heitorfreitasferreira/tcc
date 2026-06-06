# AGENTS — Servidores MCP

Novos MCPs instalados; cada um adiciona ferramentas específicas para workflow acadêmico:

| MCP | Ferramentas | Instalação |
|-----|-------------|------------|
| `crossref` | `crossref_get_work`, `crossref_search_works`, `crossref_get_references`, `crossref_search_journals`, `crossref_search_funders` | `npm i -g @cyanheads/crossref-mcp-server` |
| `scholar-sidekick` | `resolveIdentifier`, `formatCitation`, `checkOpenAccess`, `checkRetraction`, `verifyCitation` | `npm i -g scholar-sidekick-mcp` |
| `semantic-scholar` | 16 ferramentas: search, citations, authors, recommendations, OA PDF URLs | `uv tool install semantic-scholar-fastmcp` |
| `doiget` | `doiget_fetch_paper`, `doiget_resolve_paper`, `doiget_metadata_only`, `doiget_bibtex_export`, `doiget_batch_from_bibliography`, +9 | `cargo install --path crates/doiget-cli` |
| `academic-search` | `search_papers`, `fetch_paper_details`, `search_by_topic` (Crossref + S2) | `scripts/academic-search-mcp-server.py` |

## Scripts Auxiliares

- `scripts/baixar-pdf-chave.sh <bibtex-key>` — fallback CLI para download de PDF + extração de resumo
- `scripts/download-pdfs.sh [--keys K] [--use-scihub]` — pipeline batch de download (Unpaywall → S2 → OpenAlex → Sci-Hub)

## Env Variables (opcionais, para rate limits mais altos)

```bash
export CROSSREF_MAILTO="heitor@ufu.br"           # polite Crossref pool
export SCHOLAR_API_KEY="ssk_..."                 # free key de scholar-sidekick.com
export SEMANTIC_SCHOLAR_API_KEY="..."             # rate limits mais altos no S2
```
