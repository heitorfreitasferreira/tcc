```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "key_insight": "Pipeline conceptually complete but has 7+ critical gaps in MCP availability, batch error handling, ID resolution chains, and bidirectional data integrity that block reliable end-to-end execution.",
  "findings": [
    {
      "severity": "critical",
      "category": "integration",
      "description": "semantic-scholar MCP listed as dependency in Fase 2 but not present in available MCP tools (only academic-search_search_papers wraps S2 internally — different invocation pattern)",
      "fix": "Replace semantic-scholar references with academic-search_search_papers or arxiv_search_papers in Fase 2 fallback path",
      "why": "The agent will try to call a non-existent tool and fail silently or error",
      "ref": ".opencode/command/incorporar.md:41"
    },
    {
      "severity": "critical",
      "category": "integration",
      "description": "doiget_bibtex_export returns null if entry not in doiget store, which happens when PDF was obtained via scihub or download-pdfs.sh (not doiget_fetch_paper)",
      "fix": "Always call doiget_metadata_only before bibtex_export to ensure store entry; or fall back to crossref_get_work → formatCitation for BibTeX generation",
      "why": "Fase 1 cascade may succeed via scihub, leaving doiget store empty; Fase 2 then silently produces null BibTeX with no fallback",
      "ref": ".opencode/command/incorporar.md:28,43"
    },
    {
      "severity": "critical",
      "category": "completeness",
      "description": "No abort/error path defined when ALL PDF sources fail (paywalled paper, scihub blocked, download-pdfs.sh unavailable)",
      "fix": "Add explicit Fase 1 abort condition: if no PDF obtained, report 'PDF indisponível — criando nota apenas com metadados' and skip to Fase 2 metadata-only path",
      "why": "Pipeline has no terminal state for unobtainable PDFs; agent will loop or hallucinate next steps",
      "ref": ".opencode/command/incorporar.md:26-30"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "BibTeX key → DOI resolution missing when $ARGUMENTS is a BibTeX key like 'kennedy1995particle'",
      "fix": "Add step: lookup key in monografia/bib/abntex2-references.bib → extract doi field → proceed with DOI path. If key not found, search crossref by author+title.",
      "why": "Fase 1 says 'Se DOI/arXiv/key recebido' but never explains how to turn a BibTeX key into a downloadable identifier",
      "ref": ".opencode/command/incorporar.md:13,28"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "'fila' mode reads .opencode/log/pdf-events.log which does not exist (mkdir -p .opencode/log is only called in pdf-watcher on file event, never pre-created)",
      "fix": "Ensure .opencode/log/ directory exists before first use; add 'log vazio ou ausente → nenhum PDF pendente' handling",
      "why": "pdf-watcher only fires on live file.watcher.updated events; PDFs added before plugin start are never logged",
      "ref": ".opencode/plugins/pdf-watcher.ts:49"
    },
    {
      "severity": "significant",
      "category": "architecture",
      "description": "Fase 5 'atualiza campos vazios' has no defined emptiness heuristic — sections with <!-- Pendente --> comments vs '-' placeholder vs truly empty need different treatment",
      "fix": "Define explicit emptiness rules: frontmatter null/0/[]=vazio, section with only <!--...--> comment=vazio, section with only '-' or empty bullet=vazio",
      "why": "Agent will inconsistently decide what to update, potentially overwriting intentional placeholders or skipping genuinely empty sections",
      "ref": ".opencode/command/incorporar.md:76"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "Fase 7 claim verification against ~60 claims is unbounded — reading all claim files for every paper is token-prohibitive",
      "fix": "Add filtering: first check paper's inferred areas/methods → match only claims tagged with overlapping areas; limit to top-10 most relevant claims by keyword similarity",
      "why": "Without filtering, a single paper incorporation would consume >30K tokens just reading claims before any analysis begins",
      "ref": ".opencode/command/incorporar.md:88-96"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "Fase 8 crossref_get_references returns DOIs but vault matching uses BibTeX-key filenames — no DOI→BibTeX key reverse lookup mechanism exists",
      "fix": "For each reference DOI: try doiget_bibtex_export → extract key; or search vault/papers/*.md frontmatter for matching doi field via grep",
      "why": "Direct DOI-to-filename matching will miss most papers since vault filenames use author-year-key convention, not DOIs",
      "ref": ".opencode/command/incorporar.md:100-106"
    },
    {
      "severity": "minor",
      "category": "architecture",
      "description": "No partial-failure rollback — if pipeline fails at Fase 7, BibTeX was already added (Fase 4), vault note created (Fase 5), and canvas modified (Fase 6) with no undo path",
      "fix": "Add phase-gating: only commit BibTeX and canvas changes after all phases succeed; or implement a '--dry-run' mode for pre-flight validation",
      "why": "Agent-created artifacts from a failed run pollute the vault and bib with half-validated data",
      "ref": ".opencode/command/incorporar.md:24-114"
    },
    {
      "severity": "minor",
      "category": "correctness",
      "description": "extract-pdf-doi.sh DOI regex matches citations/references on first 3 pages, not just the paper's own DOI — first match wins, could be wrong paper",
      "fix": "Add heuristic: prefer DOI found near page 1 header/metadata area (pdfinfo/exiftool) over regex body matches; for regex matches, confirm with user if ambiguous",
      "why": "Many PDFs have DOIs in footnotes and references on pages 1-3; first-match-wins yields false positives",
      "ref": "scripts/extract-pdf-doi.sh:30-31"
    }
  ],
  "recommendation": "Prioritize the 3 critical fixes (semantic-scholar→academic-search, doiget_bibtex_export store dependency, PDF-unobtainable abort path) before first use. Then address the BibTeX-key→DOI resolution and fila-mode log bootstrap. The remaining significant issues can be resolved iteratively as the pipeline is exercised on real papers.",
  "schema_version": 3
}
```

## Judge 2 — Practical Execution Analysis

I evaluated the `/incorporar` pipeline by tracing every phase against the actual tools, scripts, and data structures available. My lens: can an agent executing this spec complete all 9 phases for a real paper without getting stuck?

### Phase-by-Phase Trace

**Fase 1 — Obter PDF**: The cascade logic is sound in principle, but the BibTeX-key input path has a missing link: the spec says `Se DOI/arXiv/key recebido` as if these are equivalent, but a BibTeX key like `kennedy1995particle` is not a downloadable identifier. The agent must first resolve it to a DOI by searching `abntex2-references.bib` — this lookup step is absent. Additionally, `download-pdfs.sh` delegates to `download-pdfs.py` (a Python script whose dependencies are unchecked), and the spec has no terminal state when all three sources fail (paywalled + scihub down + script fails).

**Fase 2 — Extrair metadados**: The `semantic-scholar` MCP is listed as a tool to use but is **not present** in the agent's available MCPs. The closest equivalent is `academic-search_search_papers` which internally queries Semantic Scholar but has a different invocation signature. This would fail. More critically, `doiget_bibtex_export` only works if the entry is already in the doiget store — but if Fase 1 succeeded via scihub (not doiget), the store is empty and bibtex_export returns `null` with no error. The spec has no fallback for this case.

**Fase 3 — Gerar resumo**: No structural issues, but automatic tag/role/method inference via regex (mirroring `import-bib-to-vault.sh`'s `infer_tags()`) is fragile — a paper mentioning "genetic algorithm" in a limitations section gets tagged as `metodo/ga`. No validation step exists.

**Fase 4 — Integrar BibTeX**: The spec says "ABNT-style" but `doiget_bibtex_export` returns standard BibTeX. The conversion is left implicit.

**Fase 5 — Nota vault**: The "update empty fields" heuristic is dangerously vague. What is "empty"? A section with `<!-- Pendente -->`? With just `-`? The `bibtex_key` vs `bibtex-key` (underscore vs hyphen) duplication in both template and `import-bib-to-vault.sh` is a latent inconsistency.

**Fase 6 — Canvas**: Position-based placement without collision detection means new nodes will overlap existing ones in densely populated areas (e.g., the TSP region at x:400 already has >15 nodes stacked vertically).

**Fase 7 — Claims**: Reading all ~60 claim files per paper is token-prohibitive. No filtering or relevance pre-screening exists. Additionally, the linking is one-directional (paper→claim) unlike Fase 8's bidirectional approach.

**Fase 8 — Conexões**: The crossref reference list returns DOIs, but vault matching uses BibTeX-key filenames. No DOI→key reverse lookup exists. Most references won't match.

**Fase 9 — Sugerir próximos passos**: The `rating` field has no scoring rubric.

### Cross-Cutting Concerns

1. **No rollback**: A failure at Fase 7 leaves BibTeX (Fase 4), vault note (Fase 5), and canvas (Fase 6) already mutated.
2. **MCP availability**: 4 of 6 listed MCP dependencies (crossref, doiget, scihub, pdf-reader) are network-dependent. The spec has no offline/graceful-degradation mode.
3. **pdf-watcher plugin**: Logs only on live file events. If a PDF is copied to `vault/papers/pdfs/` while OpenCode is not running, it's invisible to both the watcher and the `fila` command until the next watcher event triggers.
4. **Template drift**: The template at `vault/templates/paper-note.md` uses Portuguese placeholder text but the command spec describes different section names (e.g., "Citações-chave" matches, but "Contribuições Principais" vs "Contribuições Principais" — actually they match). Template frontmatter has `role: revisao` as default but the spec says to infer `role` dynamically. The `import-bib-to-vault.sh` script generates notes with sections like "Uso no TCC" and "Evidência / Resultado Relevante" that don't appear in the command spec's list of sections to fill.

### Verdict: WARN

The pipeline architecture is well-structured and covers the right conceptual steps. However, 3 critical blockers (missing MCP, doiget store dependency, no PDF-unobtainable abort) and 4 significant gaps (BibTeX key→DOI resolution, fila log bootstrap, claim verification scale, DOI→vault filename matching) mean the pipeline cannot complete reliably for a real-world paper without agent improvisation that may produce incorrect results.
