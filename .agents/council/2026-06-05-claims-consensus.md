# Council Report — Claims Linking Validation

**Date:** 2026-06-05
**Mode:** Default (2 judges)
**Target:** vault/claims/ — 48 claims linked to code, data, literature, vault notes

## Consensus Verdict: WARN

| Judge | Verdict | Confidence | Focus |
|-------|---------|------------|-------|
| Judge 1 | PASS | HIGH | Completeness & correctness |
| Judge 2 | WARN | HIGH | Traceability & path resolution |

**Rule:** Mixed PASS/WARN → WARN

## Summary

The linking is **structurally complete and correct** — all 48 claims have appropriate YAML fields, body sections, and evidence references. The design is sound.

However, **two systematic path prefix bugs** affect ~68 markdown body links: the relative paths have one extra `../` that resolves outside the vault.

## Findings

### FAIL: Systematic Path Prefix Bug (Body Links Only)

| Link type | Used prefix | Resolves to | Correct prefix | Should resolve to |
|-----------|------------|-------------|----------------|-------------------|
| Code links | `../../../src/` | `/home/heitor/src/` ❌ | `../../src/` | `/home/heitor/tcc/src/` ✓ |
| PDF links | `../../papers/pdfs/` | `tcc/papers/pdfs/` ❌ | `../papers/pdfs/` | `vault/papers/pdfs/` ✓ |

**Affected:** A01-A13, C01-C05, I01-I06, M01-M06 (~68 links total)
**Not affected:** YAML `source_refs[].file`, `data_refs[].pattern`, `literature_pdfs[].path` (these use repo-relative paths, not filesystem links)

### OK: Wiki-Links ✓

All 37 unique wiki-link targets resolve to existing files:
- `[[aco]]` → `vault/projeto/aco.md`
- `[[resultados]]` → `vault/projeto/resultados.md`
- `[[ga]]` → `vault/projeto/ga.md`
- `[[pso]]` → `vault/projeto/pso.md`
- `[[auditoria-hiperparametros]]` → `vault/writing/auditorias/...`
- All area and paper notes ✓

### OK: E-Series Data Patterns ✓

All glob patterns match real files:
- `*__aco__*.json` → 1530 files
- `*__ga__*.json` → 1530 files
- `*__pso__*.json` → 1530 files
- `*__bruteforce__*.json` → 18 files
- `*__lowerbound__*.json` → 30 files
- `scripts/analise-estatistica.py` → EXISTS
- `monografia/figs/cd-diagram.svg` → EXISTS

### OK: I-Series Evidence Chains ✓

All three evidence types (code + data + vault) connect coherently in sampled claims (I01, I03, I05).

### OK: YAML Schema Consistency ✓

All source_refs have `file`+`lines`+`describes`. All data_refs have `pattern`+`scope`+`metric`. All literature_pdfs have `path`+`note`.

## Recommendation

Fix the two systematic path prefix bugs in all markdown body links:
1. Replace `../../../src/` with `../../src/` in Evidência no Código / Cadeia de Evidência sections
2. Replace `../../papers/pdfs/` with `../papers/pdfs/` in Literatura de Suporte sections

**Status: FIXED** — both bugs corrected via sed replaceAll across all affected files (A01-A13, C01-C05, E01-E19, I01-I06, M01-M06). Path resolution verified with `realpath`.

## Judge Reports

- Judge 1: `.agents/council/2026-06-05-claims-judge-1.md`
- Judge 2: `.agents/council/2026-06-05-claims-judge-2.md`
