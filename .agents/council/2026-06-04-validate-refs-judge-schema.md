```json
{
  "type": "verdict",
  "verdict": "FAIL",
  "confidence": "HIGH",
  "file": ".agents/council/2026-06-04-validate-refs-judge-schema.md"
}
```
---

# JUDGE 2 — Schema Auditor Report

## 1. Per-Note Compliance Table

| Key | Frontmatter OK | Tags OK | Wiki Links OK | Sections OK | Issues |
|-----|---------------|---------|---------------|-------------|--------|
| khoufi2019survey | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/5`, `capitulo/introducao`, `capitulo/fundamentacao` tags |
| saller2025approximability | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; method `approximation` non-canonical |
| yang2023review | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; method `machine-learning` non-canonical |
| ilavarasi2014variants | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/3`, `capitulo/fundamentacao` tags |
| alkhalifa2025comparative | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao`, `capitulo/proposta` tags; method `hybrid` non-canonical; arXiv preprint (`validation_status` should perhaps be `requer-validacao`) |
| alhijawi2024genetic | PASS | **FAIL** | PASS | PASS | Tag `area/bio-inspired` should be `area/bio-inspired-optimization`; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags |
| hassanat2019crossover | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao`, `capitulo/proposta` tags |
| umbarkar2015crossover | PASS | MINOR | PASS | PASS | Missing `papel/revisao`, `relevancia/3`, `capitulo/fundamentacao` tags |
| waysi2025optimization | PASS | **FAIL** | PASS | PASS | Tag `area/bio-inspired` should be `area/bio-inspired-optimization`; missing `papel/revisao`, `relevancia/3`, `capitulo/fundamentacao` tags |
| shami2022pso | **FAIL** | **FAIL** | **FAIL** | PASS | **MISSING `bibtex_key` field** (has `bibtex-key` only); missing all `area/*` tags; missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; body uses `[[pso]]` wiki link (file does not exist — should be `[[particle-swarm]]`); `authors` field uses non-YAML array (semicolons instead of list items) |
| gad2022pso | **FAIL** | **FAIL** | **FAIL** | PASS | **MISSING `bibtex_key` field** (has `bibtex-key` only); missing all `area/*` tags; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; `authors` uses semicolons instead of YAML list |
| zhu2025cumulative | **FAIL** | **FAIL** | **FAIL** | PASS | **MISSING `bibtex_key` field** (has `bibtex-key` only); missing all `area/*` tags; missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; body uses `[[pso]]` wiki link; `authors` format non-standard |
| zhang2015comprehensive | **FAIL** | **FAIL** | **FAIL** | PASS | **MISSING `bibtex_key` field** (has `bibtex-key` only); missing all `area/*` tags; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; body uses `[[pso]]` wiki link |
| blum2024acobibliometric | PASS | **FAIL** | **FAIL** | PASS | Missing all `area/*` tags; missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; Conexões `[[aco]]` does not resolve (no `aco.md` in vault/ — should be `[[ant-colony]]`) |
| dorigo2018acooverview | PASS | **FAIL** | **FAIL** | PASS | Missing all `area/*` tags; missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; Conexões: `[[aco]]` and `[[ant-system]]` do not resolve |
| abdulghani2024comprehensive | PASS | **FAIL** | **FAIL** | PASS | Missing all `area/*` tags; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; Conexões `[[aco]]` does not resolve |
| pathak2025acoprinciples | PASS | MINOR | **FAIL** | PASS | Missing `papel/revisao`, `relevancia/3`, `capitulo/fundamentacao` tags; tag `area/routing` present but `routing` not in `areas` frontmatter; Conexões `[[aco]]` does not resolve |
| misra2024acorecent | PASS | **FAIL** | **FAIL** | PASS | Missing all `area/*` tags; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; Conexões `[[aco]]` does not resolve |
| hoffman2013tspencyclopedia | **FAIL** | MINOR | PASS | PASS | `areas: ["TSP"]` — uppercase "TSP" should be lowercase `"tsp"`; `chapters: ["cap_referencial_teorico"]` should be `["fundamentacao"]`; Conexões uses `[[cap_referencial_teorico]]` which is not a standard area name; missing `papel/revisao`, `relevancia/5`, `capitulo/fundamentacao` tags; methods use non-canonical forms (`branch-and-bound`, `linear-programming`) |
| valenzuela1997estimating | **FAIL** | MINOR | PASS | PASS | `areas: ["TSP"]` — should be `"tsp"`; `chapters: ["cap_referencial_teorico"]` should be `["fundamentacao"]`; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; methods `held-karp`, `linear-programming`, `estimation` — `estimation` is non-standard |
| huang2021branchsurvey | **FAIL** | MINOR | PASS | PASS | `areas: ["TSP"]` — should be `"tsp"`; `chapters: ["cap_referencial_teorico"]` should be `["fundamentacao"]`; missing `papel/revisao`, `relevancia/3`, `capitulo/fundamentacao` tags; methods `branch-and-bound`, `mixed-integer-programming` — `mixed-integer-programming` non-canonical; arXiv preprint, `validation_status` should perhaps be `requer-validacao` |
| gutekunst2020relaxations | **FAIL** | MINOR | PASS | PASS | `areas: ["TSP"]` — should be `"tsp"`; `chapters: ["cap_referencial_teorico"]` should be `["fundamentacao"]`; `doi: ""` empty; missing `papel/revisao`, `relevancia/4`, `capitulo/fundamentacao` tags; methods `held-karp`, `semidefinite-programming` — non-canonical; PhD thesis (`type: paper` may be inaccurate — could be `type: thesis`) |

---

## 2. Summary Statistics

| Category | Count |
|----------|-------|
| **Total notes audited** | 22 |
| **Fully compliant (0 issues)** | 0 |
| **FAIL — blocking issues** | 10 |
| **Minor/Non-blocking issues only** | 12 |
| **Notes passing Frontmatter** | 12 |
| **Notes passing Tags** | 2 (khoufi2019survey, ilavarasi2014variants — no tag errors; minor tag omissions) |
| **Notes passing Wiki Links** | 13 |
| **Notes passing Sections** | 22 |

**Blocking failures breakdown:**
- Missing `bibtex_key`: 4 notes (shami2022pso, gad2022pso, zhu2025cumulative, zhang2015comprehensive)
- Wrong `areas` casing (`"TSP"` instead of `"tsp"`): 4 notes (hoffman2013tspencyclopedia, valenzuela1997estimating, huang2021branchsurvey, gutekunst2020relaxations)
- Wrong `chapters` format (`"cap_referencial_teorico"`): 4 notes (same as above)
- Broken wiki links: 7 notes
- Wrong area tags: 2 notes (alhijawi2024genetic, waysi2025optimization)

---

## 3. Concrete Fixes Required

### 3.1 Missing `bibtex_key` (BLOCKING — scripts need both fields)

**Files:** `shami2022pso.md`, `gad2022pso.md`, `zhu2025cumulative.md`, `zhang2015comprehensive.md`

Add after `doi:` line:
```yaml
bibtex_key: <same-as-bibtex-key>
```

### 3.2 Uppercase `"TSP"` in `areas` (BLOCKING — breaks query matching)

**Files:** `hoffman2013tspencyclopedia.md`, `valenzuela1997estimating.md`, `huang2021branchsurvey.md`, `gutekunst2020relaxations.md`

Change:
```yaml
areas: ["TSP", "lower-bounds"]
```
to:
```yaml
areas:
  - "tsp"
  - "lower-bounds"
```

### 3.3 Wrong `chapters` format (BLOCKING — uses long format)

**Files:** `hoffman2013tspencyclopedia.md`, `valenzuela1997estimating.md`, `huang2021branchsurvey.md`, `gutekunst2020relaxations.md`

Change:
```yaml
chapters: ["cap_referencial_teorico"]
```
to:
```yaml
chapters:
  - fundamentacao
```

### 3.4 Broken `[[aco]]` wiki links (BLOCKING — 5 notes)

**Files:** `blum2024acobibliometric.md`, `dorigo2018acooverview.md`, `abdulghani2024comprehensive.md`, `misra2024acorecent.md`, `pathak2025acoprinciples.md`

In Conexões section, change `[[aco]]` → `[[ant-colony]]`

### 3.5 Broken `[[ant-system]]` wiki link (BLOCKING)

**File:** `dorigo2018acooverview.md`

In Conexões section, change `[[ant-system]]` → plain text `Ant System` or link to `[[dorigo1996ant]]`/`[[dorigo1997ant]]`

### 3.6 Broken `[[pso]]` wiki links in body text (BLOCKING — 3 notes)

**Files:** `shami2022pso.md`, `zhu2025cumulative.md`, `zhang2015comprehensive.md`

In body text, change `[[pso]]` → `[[particle-swarm]]`

### 3.7 Wrong tag `area/bio-inspired` (BLOCKING — 2 notes)

**Files:** `alhijawi2024genetic.md`, `waysi2025optimization.md`

Change:
```yaml
  - area/bio-inspired
```
to:
```yaml
  - area/bio-inspired-optimization
```

### 3.8 Missing `area/*` tags in frontmatter (BLOCKING — 10 notes)

**Files:** shami2022pso, gad2022pso, zhu2025cumulative, zhang2015comprehensive, blum2024acobibliometric, dorigo2018acooverview, abdulghani2024comprehensive, misra2024acorecent, alhijawi2024genetic, waysi2025optimization

For PSO notes (first 4): add `area/particle-swarm` and `area/bio-inspired-optimization`
For ACO notes (blum2024 through misra2024): add `area/ant-colony` and `area/bio-inspired-optimization` plus any domain areas (e.g., `area/tsp`)
For GA notes (last 2): add `area/bio-inspired-optimization` (already have `area/genetic-algorithms`)

### 3.9 Empty `doi` (BLOCKING — 1 note)

**File:** `gutekunst2020relaxations.md`

The value should not be empty. If a DOI truly does not exist (e.g., unpublished thesis), set to a URL or remove the field.

### 3.10 Non-standard `authors` format (MINOR — 3 notes)

**Files:** `shami2022pso.md`, `gad2022pso.md`, `zhu2025cumulative.md`

Use YAML list format (each author as a separate `-` item or quoted comma-separated string), not semicolons.

---

## 4. Systemic Issues (Affect ALL 22 Notes)

1. **Missing `papel/revisao` tag** — All 22 notes have `role: "revisao"` but NONE has the corresponding `papel/revisao` tag. The schema requires hierarchical tags that mirror the frontmatter properties. Fix: add `  - papel/revisao` to every note's tags.

2. **Missing `relevancia/{1-5}` tag** — All 22 notes have a `rating` field (3, 4, or 5) but NONE has the corresponding `relevancia/{N}` tag. Fix: add `  - relevancia/{N}` matching the rating value.

3. **Missing `capitulo/{...}` tag** — All 22 notes have `chapters` set (most to `fundamentacao`) but NONE has the corresponding `capitulo/{chapter}` tag. Fix: add `  - capitulo/fundamentacao` (or appropriate chapter tag) to each note.

4. **Inconsistent `bibtex_key` vs `bibtex-key`** — 4 notes (all PSO surveys from the same batch: shami2022pso, gad2022pso, zhu2025cumulative, zhang2015comprehensive) are missing `bibtex_key`. This suggests an import script used a format that only writes `bibtex-key`. The other 18 notes have both fields correctly.

5. **Inconsistent formatting** — Notes split into two distinct formatting batches:
   - **Batch A** (18 notes): proper YAML list format for `areas`, `methods`, `chapters`; both `bibtex_key` and `bibtex-key` present; authors as YAML list items.
   - **Batch B** (4 PSO notes + 4 TSP notes): inline array format `areas: ["..."]`; missing `bibtex_key` (PSO notes) or wrong case (TSP notes); authors as semicolons (PSO notes).

6. **Conexões `Fundamenta` line style** — Mixed usage: some notes use plain text (e.g., `Fundamenta: TSP, tsp-variants`), others use wiki links (e.g., `Fundamenta: [[aco]], ant-colony`). The plain-text ones avoid broken links but are inconsistent with the wiki-link style used for `Relacionado a`.

---

## 5. Recommendations

1. **Write a sanitization script** to bulk-apply the systemic fixes (add `papel/revisao`, `relevancia/{N}`, `capitulo/{N}` tags) to all 22 notes at once.
2. **Fix the 4 PSO notes first** (`bibtex_key` is the most critical blocker for script compatibility).
3. **Fix the 4 TSP notes second** (wrong `areas` casing and `chapters` format).
4. **Fix the 6 broken wiki links** (`[[aco]]`, `[[ant-system]]`, `[[pso]]`).
5. After all fixes, re-run the `scripts/import-bib-to-vault.sh` validation to confirm `bibtex_key`/`bibtex-key` alignment.
6. Consider adding a CI-like YAML validation step: `yamllint` on frontmatter or a simple Ruby/Python script checking required fields exist.
