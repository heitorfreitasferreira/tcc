## Judge 1 — Audit Validation

### Spot-check OK papers: [PASS] 7/10 — 3 minor status inconsistencies

All 10 papers have substantial `## Resumo` sections, internally consistent `## Métodos e Abordagens` / `## Contribuições Principais` / `## Conexões` sections, and accurate content relative to the paper's topic.

**Clear PASS (7/10):**

| Paper | Status | Verdict |
|---|---|---|
| `dorigo1996ant.md` | `lido` | Resumo substantial (~8 lines), all sections present and consistent. |
| `kennedy1995particle.md` | `lido` | Resumo substantial (~5 lines), all sections present and consistent. |
| `holland1975adaptation.md` | `lido-parcial` | Resumo substantial (~5 lines), status correctly reflects partial reading (OCR-degraded PDF). |
| `kinable2017hybrid.md` | `lido` | Resumo substantial (~5 dense lines), all sections present and consistent. |
| `stutzle2000mmas.md` | `lido` | Resumo substantial (~5 lines), all sections present and consistent. |
| `lin1973effective.md` | `lido` | Resumo substantial (~5 lines), all sections present and consistent. |
| `applegate2006traveling.md` | `lido-parcial` | Resumo substantial (~4 lines), status correctly reflects partial reading of a 600+ page book. |

**Minor status inconsistency (3/10):**

| Paper | Displayed Status | Issue |
|---|---|---|
| `blum2024acobibliometric.md` | `reading_status: pendente`, tags: `status/pendente` | Resumo is **highly detailed** (8+ lines, all sections present). The `pendente` status marker contradicts the completeness of the note. Should be at least `resumo-lido`. |
| `gad2022pso.md` | `reading_status: pendente`, tags: `status/pendente` | Same pattern — resumo is 8+ lines with full sections, yet marked `pendente`. |
| `hassanat2019crossover.md` | `reading_status: pendente`, tags: `status/pendente` | Same pattern — resumo is 6+ lines with full sections, yet marked `pendente`. |

These 3 papers use the newer frontmatter format (`reading_status`, `validation_status`, `pdf_status`, yaml `tags`), and all have `status/pendente` despite containing complete, high-quality notes. The *substance* is correct (consistent with audit classification as "OK"), but the *status marker* is inaccurate. This suggests the newer-format notes were not updated after the resumo was written.

### Spot-check sem-resumo: [PASS] 5/5

All 5 confirmed as frontmatter-only with no body content beyond the YAML header:

| File | Lines | Content |
|---|---|---|
| `aggarwal2000angular.md` | 13 | frontmatter only |
| `deepaco2023.md` | 13 | frontmatter only |
| `heldkarp1971traveling.md` | 13 | frontmatter only |
| `nagata2006eax.md` | 13 | frontmatter only |
| `vanhove2012route.md` | 13 | frontmatter only |

### Fixes verified: [PASS]

- **starzec2026motsp**: Status is now `resumo-lido` (was `nao-lido`). Resumo present and describes the paper correctly as an ACO-for-MOTSP dataset from AGH University repository. Sections: Resumo, Relevância, Conexões — internally consistent. **Fixed.**
- **muthanna2022uav**: No `lysgaard1999cluster` contamination. Clean resumo about UAV positioning and energy-efficient path scheduling in IoT — fully consistent with the paper title. Sections: Resumo, Relevância, Conexões. **Fixed.**

## Overall: [WARN]

The audit's substance conclusions are correct:
- 10/10 spot-checked "OK" papers have resumos present and well-formed
- 5/5 spot-checked "sem-resumo" papers truly lack content
- Both fixes verified as correctly applied

However, 3 of 10 "OK" papers (`blum2024acobibliometric`, `gad2022pso`, `hassanat2019crossover`) have `status/pendente` despite having complete, detailed notes. The audit classified them as "OK" based on content completeness, which is substantively correct, but the status field disagrees. This is a consistent pattern affecting the newer-format notes — the `reading_status` field was not updated after resumo writing. Recommend a follow-up pass to fix these 3 status markers (change to `lido` or at least `resumo-lido`), and to scan all other `status/pendente` notes for content completeness.
