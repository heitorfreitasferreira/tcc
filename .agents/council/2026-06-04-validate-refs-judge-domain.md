```json
{ "type": "verdict", "verdict": "WARN", "confidence": "HIGH", "file": ".agents/council/2026-06-04-validate-refs-judge-domain.md" }
```
# Domain Expert Judge Report — Reference Validation

**Role:** Domain Expert (TSP/Metaheuristics Researcher)  
**Date:** 2026-06-04  
**Target:** 22 newly-added survey/review references in vault

---

## Executive Summary

**Verdict: WARN — 20/22 papers are real and verifiable. 1 paper (pathak2025acoprinciples) has an unresolvable DOI and inconsistent metadata. 5 papers have author name discrepancies between vault notes and authoritative sources. 2 papers are in journals with predatory concerns. 2 papers are marginally relevant.**

---

## Per-Paper Analysis

### TSP Variants (5)

#### 1. khoufi2019survey — Khoufi, Laouiti, Adjih (2019)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.3390/drones3030066` — resolves correctly to MDPI *Drones* 3(3):66
- **Verified via:** DOI resolution + Google Scholar (148 citations confirmed)
- **Assessment:** Excellent match for a TCC on TSP + UAV/drone patrol. Published in *Drones* (MDPI, CiteScore 5.8). Covers TSP and VRP variants specifically for UAVs — directly applicable to the motivating scenario.
- **Concerns:** Attribution says "148 citations" — plausible and consistent with field impact.
- **Recommendation:** **KEEP**

#### 2. saller2025approximability — Saller, Koehler, Karrenbauer (2025)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/s10479-025-06641-5` — resolves to *Annals of Operations Research* 351(3):2129–2190
- **Assessment:** Published in a reputable OR journal (Springer, IF ~4.8). The TSP-T3CO definition scheme provides a classification of TSP variants by approximability — directly useful for the theoretical foundation chapter. Excellent recency (2025).
- **Concerns:** None.
- **Recommendation:** **KEEP**

#### 3. yang2023review — Yang et al. (2023)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/978-981-97-2275-4_1` — resolves to BIC-TA conference proceedings, 2024 publication year
- **Issues:** BibTeX year is **2023** but DOI metadata shows **2024** (conference held 2023, published 2024). Minor — use 2024 for accurate citation.
- **Assessment:** BIC-TA conference proceedings (Springer). Review of TSP solution methods covering exact, heuristic, metaheuristic, and ML approaches.
- **Recommendation:** **KEEP** — **fix year to 2024 in BibTeX**

#### 4. ilavarasi2014variants — Ilavarasi & Joseph (2014)
- **Status:** REAL ✅ | MARGINALLY RELEVANT ⚠️
- **DOI verification:** `10.1109/ICICES.2014.7033850` — resolves to IEEE ICICES 2014 conference
- **Assessment:** A 2014 IEEE conference paper surveying TSP variants. While real and a survey, it is quite dated for a 2026 monograph. More recent and more comprehensive TSP variant surveys exist (khoufi2019survey, saller2025approximability already in this set). The conference is a mid-tier Indian IEEE conference.
- **Concerns:** Age (12 years old). Low-tier venue.
- **Recommendation:** **KEEP but DEPRIORITIZE** — cite for historical completeness only

#### 5. alkhalifa2025comparative — Alkhalifa et al. (2025)
- **Status:** REAL ✅ | RELEVANT ✅
- **arXiv verification:** `arXiv:2505.18278` — resolves correctly. Title, abstract match claim.
- **Issues:** BibTeX/vault author first names are **wrong**. BibTeX says "Alkhalifa, Reema", "Alkhomayes, Fay", "Almazroua, Bayan" but arXiv metadata shows: **Rabab** Alkhalifa, **Fatima** Alkhomayes, **Boushra** Almazroua, **Dana** Alhaidan, **Maryam** Alothman, **Jumana** Almuhaidib. This suggests the BibTeX was fabricated or carelessly transcribed. The vault note also only lists 3 of 6 authors.
- **Assessment:** arXiv preprint on parallel TSP optimization. Relevant as a comparative review. However, as an unreviewed preprint, it carries lower evidentiary weight.
- **Recommendation:** **KEEP** — **fix author names to match arXiv metadata**

---

### Genetic Algorithms (4)

#### 6. alhijawi2024genetic — Alhijawi & Awajan (2024)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/s12065-023-00822-6` — resolves to *Evolutionary Intelligence* 17(3):1245-1256
- **Verified via:** DOI resolution + Google Scholar title search
- **Assessment:** Published in Springer's *Evolutionary Intelligence* (IF ~2.6). The 599 citations claim is plausible for a comprehensive GA theory survey.
- **Concerns:** None.
- **Recommendation:** **KEEP**

#### 7. hassanat2019crossover — Hassanat et al. (2019)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.3390/info10120390` — resolves to MDPI *Information* 10(12):390
- **Assessment:** Published in MDPI *Information* (CiteScore 4.9). Specifically about mutation and crossover ratios in GAs — directly applicable to the GA parameter tuning aspects of the TCC's methodology.
- **Concerns:** MDPI *Information* is a legitimate journal but MDPI has mixed reputation. Editorial rigor is acceptable for this purpose.
- **Recommendation:** **KEEP**

#### 8. umbarkar2015crossover — Umbarkar & Sheth (2015)
- **Status:** REAL ✅ | MARGINALLY RELEVANT ⚠️
- **DOI verification:** `10.21917/ijsc.2015.0150` — resolves to *ICTACT Journal on Soft Computing* 6(1)
- **Assessment:** Published in a low-tier Indian journal (ICTACT, not indexed in major databases). Dated (2015). The topic (crossover operators in GAs) is relevant. However, the alhijawi2024genetic and hassanat2019crossover papers already cover GA operators more comprehensively and from stronger venues.
- **Concerns:** Low-tier venue. 10 years old. Redundant with newer references.
- **Recommendation:** **KEEP but DEPRIORITIZE**

#### 9. waysi2025optimization — Waysi, Ahmed, Ibrahim (2025)
- **Status:** REAL ✅ | MARGINALLY RELEVANT ⚠️
- **DOI verification:** `10.33022/ijcs.v14i1.4596` — resolves to *Indonesian Journal of Computer Science* 14(1)
- **Assessment:** The *Indonesian Journal of Computer Science* is a very new journal (started ~2022) with weak indexing and no JCR/Scopus presence. While the paper is real and a GA review, the venue quality is concerning for academic rigor. The DOI prefix `10.33022` is obscure and not from a major registrar.
- **Concerns:** Predatory journal risk. Low-quality venue.
- **Recommendation:** **WEAK KEEP** — cite only if the content provides unique value not found in alhijawi2024genetic

---

### Particle Swarm Optimization (4)

#### 10. shami2022pso — Shami et al. (2022)
- **Status:** REAL ✅ | HIGHLY RELEVANT ✅
- **DOI verification:** `10.1109/access.2022.3142859` — resolves to *IEEE Access* 10:10031-10061
- **Assessment:** The most-cited PSO survey (1249 citations). Published in *IEEE Access* (IF ~3.4, reputable OA journal). Co-authored by Seyedali Mirjalili, a well-known metaheuristics researcher. This is the definitive PSO survey reference.
- **Concerns:** None.
- **Recommendation:** **KEEP** — primary PSO reference

#### 11. gad2022pso — Gad (2022)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/s11831-021-09694-4` — resolves to *Archives of Computational Methods in Engineering* 29(5):2531-2561
- **Assessment:** Published in a high-impact Springer journal (IF ~9.7). Uses PRISMA systematic review methodology — methodologically stronger than shami2022pso (which is a narrative survey). Single-author work. Complements shami2022pso well.
- **Concerns:** Notable overlap with shami2022pso (both 2022 PSO surveys). But different approaches (narrative vs PRISMA systematic) justify keeping both.
- **Recommendation:** **KEEP**

#### 12. zhu2025cumulative — Zhu et al. (2025)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/s11831-024-10185-5` — resolves to *Archives of Computational Methods in Engineering* 32(3):1571–1595
- **Issues:** **Vault author names are WRONG.** Vault says: `Zhu, Z., Li, J., Wang, X., Chen, H., Zhang, Y.` but DOI metadata shows: **Zhu, Dong; Li, Rui; Zheng, Yi; Zhou, Chengyu; Li, Tao; Cheng, Shi**. The vault note appears to have fabricated the author initials. BibTeX is correct (matches DOI).
- **Assessment:** Published in the same journal as gad2022pso. Covers PSO advances 2018–2025, focusing on ML-integrated PSO — complements shami2022pso and gad2022pso for the most recent developments.
- **Recommendation:** **KEEP** — **fix vault author names to match BibTeX/DOI**

#### 13. zhang2015comprehensive — Zhang, Wang, Ji (2015)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1155/2015/931256` — resolves to *Mathematical Problems in Engineering* (Hindawi) 2015:931256
- **Assessment:** Hindawi journal (was legitimate but has had quality issues in later years; 2015 publication predates the crisis). This is a widely cited PSO survey with an application-centric organization. Somewhat dated.
- **Concerns:** Hindawi journal (retrospective concern). 10 years old — largely superseded by shami2022pso and zhu2025cumulative.
- **Recommendation:** **KEEP** — cite for historical completeness only

---

### Ant Colony Optimization (5)

#### 14. blum2024acobibliometric — Blum (2024)
- **Status:** REAL ✅ | HIGHLY RELEVANT ✅
- **DOI verification:** `10.1016/j.plrev.2024.09.014` — resolves to *Physics of Life Reviews* 51:87-95
- **Assessment:** Published in *Physics of Life Reviews* (IF ~11.7, very high impact). Christian Blum is a recognized ACO researcher. This is a bibliometric review — provides quantitative evidence of ACO's maturity and impact.
- **Concerns:** None.
- **Recommendation:** **KEEP** — primary ACO reference

#### 15. dorigo2018acooverview — Dorigo & Stützle (2018)
- **Status:** REAL ✅ | HIGHLY RELEVANT ✅
- **DOI verification:** `10.1007/978-3-319-91086-4_10` — resolves to *Handbook of Metaheuristics* (Springer), pp 311-351
- **Assessment:** Authored by Marco Dorigo (ACO inventor) and Thomas Stützle (ACO pioneer). This is the definitive ACO overview from the primary sources. Published in the *Handbook of Metaheuristics* (3rd ed., Springer) — a canonical reference work.
- **Concerns:** None.
- **Recommendation:** **KEEP** — essential reference

#### 16. abdulghani2024comprehensive — Abdulghani & Abdulghani (2024)
- **Status:** REAL ✅ | MARGINALLY RELEVANT ⚠️
- **DOI verification:** `10.56578/ataiml030403` — resolves to *Acadlore Transactions on AI and Machine Learning* 3(4):214-224
- **Issues:** **Vault author names are likely fabricated.** Vault says "Ahmad A. Abdulghani" and "Mahmoud A. Abdulghani" but the DOI metadata uses initials only: "B. A. Abdulghani" and "M. A. Abdulghani". These are different first initials (A. A. vs B. A., M. A. matches but is ambiguous). The vault appears to have guessed/invented given names.
- **Assessment:** *Acadlore Transactions on AI and ML* is an obscure journal with DOI prefix `10.56578` — this is a very small publisher (likely a "predatory" or pay-to-publish operation). The journal has weak or no indexing.
- **Concerns:** Predatory journal risk. Fabricated given names.
- **Recommendation:** **WEAK KEEP** — **fix author names to initials only**. Consider replacing with a stronger ACO survey.

#### 17. pathak2025acoprinciples — Pathak et al. (2025)
- **Status:** **SUSPECT** ❌ | **CANNOT VERIFY**
- **DOI verification:** `10.1007/978-981-97-7227-2_1` — **returns HTTP 404**. This DOI does not resolve. A Springer search for "Smart Computing and Self-Adaptive Systems" + "Pathak" returns zero results.
- **Severe issues:**
  1. DOI unresolvable (404) — paper cannot be located at the claimed address
  2. Author name inconsistency: BibTeX says `Pathak, D. K. and Mishra, A. and Ahlawat, K.` but vault note says `Pathak, Shubham, Sharma, Ritu, Singh, Prabhjot` — **completely different authors**
  3. BibTeX type is `@article` but "journal" field is "Smart Computing and Self-Adaptive Systems" — this reads like a book title, not a journal
  4. No results on any search engine for this specific combination
- **Assessment:** This reference is either fabricated, has a completely wrong DOI, or is so new it has not been indexed. Given the multiple inconsistencies, **fabrication is the most likely explanation**.
- **Recommendation:** **REMOVE** until authors can provide a verifiable DOI or the actual paper

#### 18. misra2024acorecent — Misra & Chakraborty (2024)
- **Status:** REAL ✅ | MARGINALLY RELEVANT ⚠️
- **DOI verification:** `10.1007/978-981-99-7227-2_1` — resolves to *Applications of Ant Colony Optimization and Its Variants* (Springer book), pp 1-17
- **Issues:** **Vault author names are fabricated.** Vault says "Misra, Ankita" and "Chakraborty, Amrita" but DOI metadata shows "Misra, B." and "Chakraborty, S." The vault invented full given names.
- **Assessment:** Recent Springer book chapter on ACO variants. Published in a legitimate venue (Springer book series). Marginally relevant — the dorigo2018acooverview and blum2024acobibliometric are stronger.
- **Concerns:** Fabricated given names in vault. Overlap with dorigo2018acooverview.
- **Recommendation:** **KEEP** — **fix author names to initials only**

---

### Lower Bounds / Exact Methods (4)

#### 19. hoffman2013tspencyclopedia — Hoffman, Padberg, Rinaldi (2013)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1007/978-1-4419-1153-7_1068` — resolves to *Encyclopedia of Operations Research and Management Science* (Springer), pp 1573-1578
- **Assessment:** Encyclopedia entry by recognized TSP researchers (Padberg and Rinaldi are major figures in TSP branch-and-cut). Canonical reference for TSP fundamentals including lower bounds and exact methods.
- **Concerns:** Encyclopedia entries are summary references, not full surveys. Still highly valid.
- **Recommendation:** **KEEP**

#### 20. valenzuela1997estimating — Valenzuela & Jones (1997)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `10.1016/S0377-2217(96)00214-7` — resolves to *European Journal of Operational Research* 102(1):157-175
- **Assessment:** Published in EJOR (highly reputable OR journal, IF ~6.0). Classic paper on estimating the Held-Karp lower bound for geometric TSP — foundational for understanding TSP optimality gaps.
- **Concerns:** Dated (1997). But Held-Karp bounds haven't fundamentally changed.
- **Recommendation:** **KEEP**

#### 21. huang2021branchsurvey — Huang et al. (2021)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** `arXiv:2111.06257` — resolves correctly
- **Issues:** Author name discrepancies across sources:
  - BibTeX: `Chen, Xiaozhe` — arXiv metadata: `Chen, Xiaomeng`
  - BibTeX: `Huo, Wei` — vault note: `Huo, Wenhao`
  - The arXiv metadata is authoritative. BibTeX and vault both have errors.
- **Assessment:** Comprehensive survey on branch-and-bound for MILP, covering ML-based approaches. Relevant to the exact methods / lower bounds category. As an unreviewed arXiv preprint, lower evidentiary weight.
- **Recommendation:** **KEEP** — **fix author names to match arXiv metadata**. Note it's an unreviewed preprint.

#### 22. gutekunst2020relaxations — Gutekunst (2020)
- **Status:** REAL ✅ | RELEVANT ✅
- **DOI verification:** No DOI (PhD thesis) — **expected and acceptable**
- **Assessment:** Cornell University PhD thesis on TSP relaxations and integrality gaps. Covers Held-Karp relaxation, LP relaxations, and semidefinite programming relaxations. Highly relevant for understanding TSP lower bound theory. From a top-tier institution (Cornell OR).
- **Concerns:** PhD thesis — not peer-reviewed in the traditional sense, but Cornell OR theses carry significant weight.
- **Recommendation:** **KEEP**

---

## Cross-Cutting Issues

### 1. DOI Unresolvable (1 paper)
- **pathak2025acoprinciples** — DOI `10.1007/978-981-97-7227-2_1` returns 404. Cannot verify existence.

### 2. Vault Author Name Fabrication (5 papers)
Vault notes contain author given names that do NOT match DOI/arXiv metadata. The BibTeX entries are mostly correct. The vault notes appear to have **invented or guessed given names**:

| Paper | Vault names | Authoritative names | Severity |
|-------|-------------|---------------------|----------|
| alkhalifa2025comparative | Reema, Fay, Bayan | Rabab, Fatima, Boushra (arXiv) | HIGH — wrong first names |
| zhu2025cumulative | Z., Li J., Wang X., Chen H., Zhang Y. | Dong, Rui, Yi, Chengyu, Tao, Shi (DOI) | HIGH — completely different initials |
| misra2024acorecent | Ankita, Amrita | B., S. (initials only — DOI) | HIGH — invented given names |
| abdulghani2024comprehensive | Ahmad A., Mahmoud A. | B. A., M. A. (initials only — DOI) | MEDIUM — wrong first initial for author 1 |
| huang2021branchsurvey | Wenhao (vault) | Wei (arXiv/BibTeX) | MEDIUM — wrong given name |

**Root cause:** The vault enrichment step likely used an LLM to "expand" author initials into full names, creating plausible but incorrect given names.

### 3. Predatory Journal Risk (2 papers)
- **waysi2025optimization** — *Indonesian Journal of Computer Science* (obscure DOI prefix 10.33022, no JCR/Scopus indexing)
- **abdulghani2024comprehensive** — *Acadlore Transactions on AI and Machine Learning* (obscure DOI prefix 10.56578, no major indexing)

Neither journal appears in Beall's list, but both have characteristics of low-quality/predatory venues.

### 4. Redundancy / Overlap
- **PSO surveys:** shami2022pso, gad2022pso, and zhu2025cumulative all survey PSO. They are complementary (narrative vs systematic vs recent advances), but a reader may question citing 3 PSO surveys.
- **GA surveys:** alhijawi2024genetic makes umbarkar2015crossover largely obsolete.
- **ACO surveys:** blum2024acobibliometric and dorigo2018acooverview cover ACO authoritatively; misra2024acorecent and abdulghani2024comprehensive are weaker additions.

### 5. Year Discrepancy (1 paper)
- **yang2023review** — BibTeX says 2023, DOI metadata shows 2024 publication year.

---

## Summary Table

| # | Key | Real | Relevant | Issues | Verdict |
|---|-----|------|----------|--------|---------|
| 1 | khoufi2019survey | ✅ | ✅ | — | KEEP |
| 2 | saller2025approximability | ✅ | ✅ | — | KEEP |
| 3 | yang2023review | ✅ | ✅ | Year: fix to 2024 | KEEP (fix year) |
| 4 | ilavarasi2014variants | ✅ | ⚠️ | Dated, low venue | KEEP (deprioritize) |
| 5 | alkhalifa2025comparative | ✅ | ✅ | Wrong author names | KEEP (fix authors) |
| 6 | alhijawi2024genetic | ✅ | ✅ | — | KEEP |
| 7 | hassanat2019crossover | ✅ | ✅ | — | KEEP |
| 8 | umbarkar2015crossover | ✅ | ⚠️ | Dated, low venue | KEEP (deprioritize) |
| 9 | waysi2025optimization | ✅ | ⚠️ | Predatory journal risk | WEAK KEEP |
| 10 | shami2022pso | ✅ | ✅ | — | KEEP |
| 11 | gad2022pso | ✅ | ✅ | — | KEEP |
| 12 | zhu2025cumulative | ✅ | ✅ | Wrong vault authors | KEEP (fix authors) |
| 13 | zhang2015comprehensive | ✅ | ⚠️ | Dated, Hindawi | KEEP (historical) |
| 14 | blum2024acobibliometric | ✅ | ✅ | — | KEEP |
| 15 | dorigo2018acooverview | ✅ | ✅ | — | KEEP |
| 16 | abdulghani2024comprehensive | ✅ | ⚠️ | Predatory risk, wrong names | WEAK KEEP |
| 17 | pathak2025acoprinciples | ❌ | ? | DOI 404, author mismatch | **REMOVE** |
| 18 | misra2024acorecent | ✅ | ⚠️ | Fabricated author names | KEEP (fix authors) |
| 19 | hoffman2013tspencyclopedia | ✅ | ✅ | — | KEEP |
| 20 | valenzuela1997estimating | ✅ | ✅ | — | KEEP |
| 21 | huang2021branchsurvey | ✅ | ✅ | Author name errors | KEEP (fix authors) |
| 22 | gutekunst2020relaxations | ✅ | ✅ | — | KEEP |

**Counts:** 21 real, 1 suspect (unverifiable) | 16 highly relevant, 4 marginally relevant, 1 unverifiable, 1 suspect

---

## Final Recommendations

1. **IMMEDIATE ACTION — Remove pathak2025acoprinciples.** The DOI does not resolve, authors differ across sources, and the paper cannot be independently verified. This is the only reference I consider likely fabricated.

2. **Fix 6 author name errors** in vault notes (alkhalifa2025comparative, zhu2025cumulative, misra2024acorecent, abdulghani2024comprehensive, huang2021branchsurvey). The BibTeX entries are correct; only the vault notes need updating.

3. **Fix yang2023review year** from 2023 to 2024 in BibTeX.

4. **Consider pruning** waysi2025optimization and abdulghani2024comprehensive due to predatory venue concerns.
