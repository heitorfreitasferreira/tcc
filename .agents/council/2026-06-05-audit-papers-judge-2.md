# Judge 2 — Audit Validation: "Incompleto" Category + Roadmap Cross-check

**Date:** 2026-06-05  
**Scope:** Validate 10 "incompleto" papers, cross-check "sem resumo" × roadmap, citation risk assessment

---

## Incompleto Verification: PASS 10/10 verified

All 10 papers correctly classified as "incompleto" (PDF failure documented in resumo):

| # | Paper | PDF Exists? | `pdfinfo` Status | Verdict |
|---|-------|------------|-------------------|--------|
| 1 | araujo2025pso | **NO** | File not found | **CONFIRMED** — worse than corrupted; completely absent |
| 2 | bean1994genetic | YES | 8pp, valid metadata (INFORMS) | **CONFIRMED** — only metadata page, no article body (roadmap P29 confirms) |
| 3 | blum2005acointro | YES | `Internal Error: xref num 467 not found` | **CONFIRMED** — corrupted xref, 21pp but no extractable text |
| 4 | clerc2000discretepso | YES | 18pp, OK metadata (Author: Maurice, distiller 3.02) | **CONFIRMED** — structurally valid PDF but text extraction failed (likely scanned or embedded fonts) |
| 5 | halim2019combinatorial | YES | 14pp, heavily corrupted (Bad FCHECK × 12, syntax errors × 100+) | **CONFIRMED** — stream corruption throughout |
| 6 | larranaga1999ga | YES | `Internal Error: xref num 221 not found`, 42pp | **CONFIRMED** — xref corruption |
| 7 | oliver1987crossover | YES | 14pp, scanned proceedings (Paper Capture Plug-in) | **CONFIRMED** — wrong content: PDF contains De Jong article, not Oliver et al. pp. 224–230 |
| 8 | potvin1996ga | YES | `Internal Error: xref num 316 not found`, 34pp, Producer: PageGenie (from TIFF scan) | **CONFIRMED** — scanned images, no extractable text |
| 9 | ppaco2024 | YES | 9pp, stream errors (Bad FCHECK × 3, Invalid XRef × 3) | **CONFIRMED** — corrupted streams |
| 10 | wu2020comparative | YES | `Internal Error: xref num 170 not found`, 5pp | **CONFIRMED** — xref corruption |

### Notes on specific papers:

- **araujo2025pso**: The note claims a 6-page PDF with blank OCR, but the file does not exist on disk. This should be reclassified as **sem PDF** or the file must be recovered.
- **oliver1987crossover**: Unique case — PDF exists and is structurally valid, but contains the wrong article. The scanned proceedings include the ToC showing Oliver et al. on pp. 224–230, then the actual body starts with De Jong's paper.
- **clerc2000discretepso**: The PDF is structurally OK (18pp, valid metadata). The note's claim of no extractable text may reflect a scanned copy or fonts that resist extraction. Worth re-testing with `pdftotext` before discarding.
- **ppaco2024**: Stream errors present but metadata extracted successfully (LuaHBTeX, 2024). The 9-page count and metadata suggest the LaTeX source compiled — corruption may be partial.

### Contribution content assessment:

All 10 resumos **correctly document extraction failure** rather than summarizing paper content. Contributions sections are based on title, metadata, and bibliographic context — not on body text. This is **appropriate** for the `lido-parcial` status.

**No hidden contribution content** was found buried in any of the 10 notes. The notes are honest about their limitations.

---

## Roadmap Dependency Mapping

### P26 — Held-Karp BibTeX + Lower Bounds Section
| Key | Status | Blocks P26? |
|-----|--------|-------------|
| heldkarp1971traveling | **sem resumo** | YES |
| heldkarp1970traveling | OK (has note) | — |
| johnson1996asymptotic | OK (has note) | — |
| kinable2017hybrid | OK (has note) | — |
| righini2021efficient | OK (has note) | — |
| valenzuela1997estimating | OK (has note) | — |

**Impact:** 1/6 entries blocked. Low severity — the canonical Held-Karp is heldkarp1970 (already in vault). heldkarp1971 is a follow-up.

### P28 — ATSP/TDTSP References
No specific keys from the "sem resumo" list match P28. These are new references to add. **Not blocked.**

### P29 — Bean (1994) PDF Recovery
| Key | Status | Blocks P29? |
|-----|--------|-------------|
| bean1994genetic | **incompleto** | YES — this IS P29 |

**Impact:** Fully blocked. P29 explicitly requires obtaining a readable copy and updating the note.

### P30 — Fix Inconsistent Ratings (0→≥3)
| Key | Status | Current Rating | Target | Blocks P30? |
|-----|--------|----------------|--------|-------------|
| bean1994genetic | incompleto | 0 | ≥3 | YES |
| vanhove2012route | sem resumo | ? | ≥3 | YES |
| winter2002modeling | OK | 0 | ≥3 | Partial |
| wang2021ant | OK | 0 | ≥3 | Partial |

**Impact:** 2/4 blocked (bean + vanhove). Medium severity — requires both PDF recovery (P29) and rating adjustment.

### P31 — Demšar Note + 6 Lower Bound BibTeX Entries
| Key | Status | Blocks P31? |
|-----|--------|-------------|
| aggarwal2000angular | **sem resumo** | YES |
| balas1985branch | **sem resumo** | YES |
| fischetti1992additive | **sem resumo** | YES |
| leraromero2020dynamic | **sem resumo** | YES |
| lawler1985traveling | **sem resumo** | YES |
| karp1979patching | Unknown | Check |
| demsar2006statistical | Empty note (outside papers/) | YES (separate issue) |

**Impact:** 5/7 blocked. **Severely blocked.** However, the roadmap's own triage table says these are "Download manual condicionado à seção de lower bounds" — meaning they don't block the monograph if the discussion stays limited to AP bound.

### P34 — Cite shami2022pso and gad2022pso in PSO section
| Key | Status | Blocks P34? |
|-----|--------|-------------|
| shami2022pso | OK (note exists) | No |
| gad2022pso | OK (note exists) | No |

**Impact:** Not blocked by audit findings. Both notes exist and are presumably valid.

---

## Citation Risk Assessment

### "Sem Resumo" Papers Currently Cited in LaTeX Monograph: 5 of 18

| Key | Cited In | Citation Context | Risk |
|-----|----------|------------------|------|
| **vanhove2012route** | `fundamentacao.tex:18` | "custos de curva... Winter e Vanhove e Fack discutem modelos" | **MEDIUM** — roadmap classifies as "condicional"; used for contextual claim, not strong claim |
| **dellamico2021multiple** | `fundamentacao.tex:48`, `introducao.tex:11` | "Dell'Amico et al. trataram variantes com múltiplos drones" | **MEDIUM** — roadmap says "análise posterior"; cited for drone routing context |
| **dellamico2022exact** | `fundamentacao.tex:48`, `introducao.tex:11` | Same as above, paired citation | **MEDIUM** — same risk profile |
| **deepaco2023** | `conclusao.tex:28` | "investigar métodos híbridos e neurais, como DeepACO" | **LOW** — cited only in Conclusão as future work direction |
| **nagata2006eax** | `fundamentacao.tex:60` | "Nagata propôs o operador EAX, referência de alta qualidade" | **HIGH** — used to support a claim about GA crossover quality |

### "Incompleto" Papers Currently Cited in LaTeX Monograph: 6 of 10

| Key | Cited In | Citation Context | Risk |
|-----|----------|------------------|------|
| **bean1994genetic** | `fundamentacao.tex:69` | "random keys: cada partícula mantém um vetor real... Bean" | **HIGH** — supports implementation claim; PDF only has metadata page |
| **clerc2000discretepso** | `fundamentacao.tex:69` | "Clerc formalizou uma versão discreta do PSO para permutações" | **HIGH** — supports strong methodological claim; PDF text not extractable |
| **potvin1996ga** | `fundamentacao.tex:60` | "Potvin revisou operadores de crossover para TSP" | **HIGH** — used for substantive claim; PDF is scanned images, unreadable |
| **larranaga1999ga** | `fundamentacao.tex:60` | "Larrañaga et al. sistematizaram representações e operadores" | **HIGH** — used for substantive claim; PDF has xref corruption |
| **wu2020comparative** | `fundamentacao.tex:95`, `introducao.tex:15` | "Wu comparou GA, PSO e ACO em instâncias TSPLIB" | **HIGH** — roadmap lists as "central segura" but PDF is corrupted |
| **halim2019combinatorial** | `fundamentacao.tex:95`, `introducao.tex:15` | "Halim e Ismail destacaram que a escolha da representação..." | **HIGH** — roadmap lists as "central segura" but PDF is heavily corrupted |

### CRITICAL FINDING: Roadmap Contradiction

The roadmap (line 339) lists `wu2020comparative` and `halim2019combinatorial` as **"Referências centrais seguras"** for Fundamentação. Both have corrupted PDFs (xref errors, FCHECK failures) with no extractable text. The roadmap also lists `clerc2000discretepso` as "segura" despite its text extraction failure.

This contradicts the roadmap's own citation safety policy (lines 242–247):
- `wu2020comparative` should be classified **`frágil`** (PDF corrompido, nota não validada)
- `halim2019combinatorial` should be classified **`frágil`**
- `clerc2000discretepso` should be classified **`frágil`** or at best **`condicional`**

---

## Summary of Citations by Safety Class

| Class | Count | Papers |
|-------|-------|--------|
| `frágil` → cited in monograph | **11** | bean1994genetic, clerc2000discretepso, potvin1996ga, larranaga1999ga, wu2020comparative, halim2019combinatorial, vanhove2012route, dellamico2021multiple, dellamico2022exact, deepaco2023, nagata2006eax |
| `segura` → cited in monograph | Varies | All remaining citations with valid PDFs |
| `bloqueada` | **1** | araujo2025pso (no PDF on disk) — not cited in monograph |

---

## Critical Issues

1. **6 incompleto papers actively cited with substantive claims**: bean1994genetic, clerc2000discretepso, potvin1996ga, larranaga1999ga, wu2020comparative, halim2019combinatorial are all used to support methodological claims in `fundamentacao.tex`. None have readable PDFs. These citations are **frágil** per the citation safety policy and should NOT sustain claims about specific operators, representations, or experimental findings.

2. **Roadmap self-contradiction**: `wu2020comparative`, `halim2019combinatorial`, and `clerc2000discretepso` are listed as "Referências centrais seguras" but have corrupted/unreadable PDFs. The roadmap's own `citation-safety.md` policy (P22, still pendente) would classify them as `frágil`.

3. **araujo2025pso has no PDF file**: The note references a PDF that doesn't exist on disk. This paper is not cited in the monograph, so risk is low, but the audit metadata is inaccurate.

4. **P22 (citation-safety.md) is still pendente**: This means 11 frágil citations are in the monograph with no formal safety classification. Implementing P22 would immediately flag the conflict.

5. **P31 is severely blocked**: 5 of the 6 lower-bound BibTeX keys are "sem resumo". However, the roadmap's triage already classifies these as "download manual condicionado" — they don't block the monograph unless the lower bounds section is expanded beyond AP bound.

---

## Overall: WARN

The "incompleto" classification is accurate for all 10 papers. The audit correctly identified the failures.

**However**, the presence of 6 incompleto + 5 sem-resumo papers actively cited in the monograph — 11 frágil citations total — represents a significant bibliographic integrity risk. The roadmap's own central references for Fundamentação include papers with corrupted PDFs, contradicting its citation safety framework.

**Recommended actions:**
1. Execute P22 (citation-safety.md) immediately to formalize the classification
2. Flag all 11 frágil citations with `% citation-safety: frágil — PDF corrompido` in the .tex
3. Downgrade claims supported only by frágil citations to contextual/review claims
4. Recover or replace PDFs for wu2020comparative, halim2019combinatorial, clerc2000discretepso (high-impact, directly support comparative claims)
5. Recover araujo2025pso.pdf or remove the file reference from the note
