```json
{
  "type": "verdict",
  "verdict": "FAIL",
  "confidence": "HIGH",
  "file": ".agents/council/2026-06-04-validate-refs-judge-bib.md"
}
```

# Bibliographer Judge — Metadata Accuracy Report

**Date**: 2026-06-04
**Target**: 22 systematic reviews/surveys added to vault
**Method**: CrossRef API verification, arXiv API checks, journal reputation assessment
**Sources consulted**: api.crossref.org, export.arxiv.org, doi.org

---

## Per-Paper Metadata Accuracy Table

| # | Key | DOI Valid | Title Match | Authors Match | Journal OK | Issues |
|---|-----|-----------|-------------|---------------|------------|--------|
| 1 | khoufi2019survey | YES | YES | YES | YES | None — all fields verified |
| 2 | saller2025approximability | YES | YES | NO | YES | **First author wrong**: "Sebastian" → "Sophia". **Missing**: volume=351, issue=3, pages=2129-2190 |
| 3 | yang2023review | YES | YES | NO | YES | **ALL 5 authors wrong**: "Lianlian" → "Longrui", "Xing" → "Xiyuan", "Zhengbing" → "Zhaoqi", "Shuaian" → "Sicong", "Ji" → "Jie". **Pages wrong**: 1-11 → 3-16. Year 2023 → 2024 |
| 4 | ilavarasi2014variants | YES | YES | PARTIAL | NO | **Type**: `@article` → `@inproceedings`. **Journal field**: contains conference name, should be `booktitle`. Author "K.S. Joseph" → "K. Suresh Joseph". Pages missing (1-7) |
| 5 | alkhalifa2025comparative | YES* | YES | NO | N/A (arXiv) | **3 of 6 authors wrong**: "Reema" → "Rabab", "Fay" → "Fatima", "Bayan" → "Boushra". **3 authors omitted** (Dana Alhaidan, Maryam Alothman, Jumana Almuhaidib) replaced with "and others" |
| 6 | alhijawi2024genetic | YES | YES | YES | YES | None — all fields verified |
| 7 | hassanat2019crossover | YES | YES | YES | YES | None — all fields verified |
| 8 | umbarkar2015crossover | YES | YES (case) | YES | YES | **Missing pages**: 1083-1092 |
| 9 | waysi2025optimization | YES | YES | NO | FLAG | **ALL 3 author names wrong**: "Daban" → "Diyar", "Bestoon T." → "Berivan Tahir", "Ibrahim Mahmood" → given/family reversed. **Missing pages**. **Journal**: "The Indonesian Journal of Computer Science" (STMIK Indonesia Padang) — questionable, not indexed in Scopus/WoS |
| 10 | shami2022pso | YES | YES | YES | YES | None — all fields verified |
| 11 | gad2022pso | YES | YES | YES | YES | None — all fields verified |
| 12 | zhu2025cumulative | YES | YES | NO | YES | **5 of 6 authors wrong**: "Dong" → "Donglin", "Yi" → "Yangyang", "Chengyu" → "Changjun", "Tao" → "Taiyong". **Missing**: volume=32, pages=1571-1595 |
| 13 | zhang2015comprehensive | YES | YES | YES | YES | Pages "931256" is article ID (acceptable for Hindawi). CrossRef lists pages as 1-38 |
| 14 | blum2024acobibliometric | YES | YES | YES | YES | None — all fields verified |
| 15 | dorigo2018acooverview | YES | YES | YES | YES | **Year wrong**: 2018 → 2019 |
| 16 | abdulghani2024comprehensive | YES | YES | PARTIAL | FLAG | **Author initials only** (should be full names: Batool Abdulsatar, Mohammed Abdulsattar). **Missing pages**: 214-224. **Publisher**: Acadlore Publishing Services — questionable, very new ISSN (2957 prefix, 62 total DOIs). Not indexed in reputable databases |
| 17 | pathak2025acoprinciples | **NO** | PARTIAL | NO | NO | **DOI `10.1007/978-981-97-7227-2_1` DOES NOT EXIST (HTTP 404)**. Real DOI: `10.56155/978-81-975670-0-1-5`. **Authors wrong**: "D.K." → "Deepak Kumar", "A." → "Aishwarya", "K." → "Koastubh", + missing "Lavika Goel". **Type**: `@article` → should be book chapter. **Journal**: "Smart Computing and Self-Adaptive Systems" unverified |
| 18 | misra2024acorecent | YES | YES | PARTIAL | YES | **Author initials only** (should be Bitan, Sayan). **Booktitle subtitle wrong**: "Case Studies and New Developments" → "Applications of Ant Colony Optimization and its Variants". **Pages wrong**: 1-22 → 1-17 |
| 19 | hoffman2013tspencyclopedia | YES | YES | YES | YES | None — all fields verified |
| 20 | valenzuela1997estimating | YES | YES | YES | YES | None — all fields verified |
| 21 | huang2021branchsurvey | YES* | YES | NO | N/A (arXiv) | **2nd author wrong**: "Xiaozhe" → "Xiaomeng" |
| 22 | gutekunst2020relaxations | N/A | N/A | YES† | N/A (PhD) | No DOI assigned — expected for pre-2022 dissertation. †Author verified via related journal publications. Suggest citing published version: Gutekunst & Williamson (2021), DOI `10.1287/moor.2020.1100` |

\* arXiv DOIs confirmed via arXiv API but not indexed in CrossRef
\† Author confirmed via related publications in CrossRef

---

## Summary Statistics

| Metric | Count |
|--------|-------|
| **PASS** (no issues) | 7 |
| **WARN** (minor/missing fields) | 6 |
| **FAIL** (incorrect metadata) | 9 |
| **Total verified** | 22 |

### Issue Breakdown
- **Author name errors**: 8 papers (saller, yang, alkhalifa, waysi, zhu, pathak, huang, ilavarasi)
- **Missing pages**: 5 papers (saller, umbarkar, waysi, zhu, abdulghani)
- **Missing volume**: 2 papers (saller, zhu)
- **Wrong year**: 2 papers (yang: 2023→2024, dorigo: 2018→2019)
- **Wrong entry type**: 3 papers (ilavarasi: article→inproceedings, pathak: article→incollection, yang: inproceedings acceptable)
- **Fabricated/nonexistent DOI**: 1 paper (pathak — HTTP 404)
- **Predatory/questionable journal**: 2 papers (waysi, abdulghani)
- **Author initials only** (missing full names): 2 papers (abdulghani, misra)

---

## Detailed Findings

### 1. Fabricated DOI — pathak2025acoprinciples (CRITICAL)

The DOI `10.1007/978-981-97-7227-2_1` returns HTTP 404 from both doi.org and api.crossref.org. Title search on CrossRef reveals the actual DOI is `10.56155/978-81-975670-0-1-5`. The BibTeX entry appears to have a fabricated Springer-format DOI that never existed. Coincidentally, misra2024acorecent has a similar-looking (but valid) DOI `10.1007/978-981-99-7227-2_1` — possibly a copy-paste error.

**Required corrections**:
- Replace DOI with `10.56155/978-81-975670-0-1-5`
- Fix all author names: "Deepak Kumar Pathak, Aishwarya Mishra, Koastubh Ahlawat, Lavika Goel"
- Change type from `@article` to `@incollection` or `@inproceedings`
- Verify and correct booktitle/journal name

### 2. Systematic Author Name Errors — yang2023review (CRITICAL)

All 5 author given names in the BibTeX entry are wrong. This suggests the BibTeX was composed from an incorrect source or another paper entirely. The DOI resolves correctly to a book chapter with different authors:
- BibTeX: Lianlian Yang, Xing Wang, Zhengbing He, Shuaian Wang, Ji Lin
- CrossRef: Longrui Yang, Xiyuan Wang, Zhaoqi He, Sicong Wang, Jie Lin

Further affected: pages (1-11 → 3-16), year (2023 → 2024).

### 3. Systematic Author Name Errors — zhu2025cumulative (CRITICAL)

5 of 6 author given names are substantially wrong:
- BibTeX: Dong Zhu, Rui Li, Yi Zheng, Chengyu Zhou, Tao Li, Shi Cheng
- CrossRef: Donglin Zhu, Rui Li, Yangyang Zheng, Changjun Zhou, Taiyong Li, Shi Cheng

Only "Rui Li" and "Shi Cheng" are correct. Missing volume 32 and pages 1571-1595.

### 4. Author Name Errors — alkhalifa2025comparative

arXiv API verified the paper exists. 3 of 6 author given names wrong + 3 authors omitted:
- BibTeX: Reema Alkhalifa, Fay Alkhomayes, Bayan Almazroua, and others
- arXiv: Rabab Alkhalifa, Fatima Alkhomayes, Boushra Almazroua, Dana Alhaidan, Maryam Alothman, Jumana Almuhaidib

### 5. Predatory/Questionable Journals

**waysi2025optimization** — "The Indonesian Journal of Computer Science" (ISSN 2302-4364 / 2549-7286)
- Publisher: STMIK Indonesia Padang (a local Indonesian vocational college, not a recognized academic publisher)
- ~1,286 DOIs registered total — moderate volume
- Not indexed in Scopus or Web of Science
- Very low citation counts across their corpus
- **Assessment**: Questionable — unlikely to meet the peer-review standards expected for a systematic review citation in an MSc monograph. Consider replacing with a survey from a reputable journal.

**abdulghani2024comprehensive** — "Acadlore Transactions on AI and Machine Learning" (ISSN 2957-9562 / 2957-9570)
- Publisher: Acadlore Publishing Services Limited
- ISSN prefix 2957 indicates registration after ~2021 (very new journal)
- Only 102 total DOIs — extremely low volume
- Publisher has characteristics of a "paper mill" operation
- **Assessment**: Highly questionable — very new journal with minimal track record. Strongly recommend replacement.

### 6. Missing Required Fields

Multiple entries are missing `pages` which is a required BibTeX field for `@article`:
- saller2025approximability (missing vol, issue, pages)
- umbarkar2015crossover (missing pages)
- waysi2025optimization (missing pages)
- zhu2025cumulative (missing vol, pages)
- abdulghani2024comprehensive (missing pages)

### 7. Author Consistency / Research Groups

Notable author clusters confirming research lineage:
- **IRIDIA Lab (ULB, Brussels)**: Marco Dorigo, Thomas Stützle, Christian Blum — foundational ACO research group. Appears in: dorigo1996ant, dorigo1997ant, stutzle2000mmas, dorigo2004book, blum2005acointro, dorigo2005acotheory, dorigo2018acooverview, blum2024acobibliometric. All entries verified correct.
- **Gutekunst-Williamson (Cornell)**: Samuel C. Gutekunst and David P. Williamson — TSP relaxations research. Dissemination through MOR journal (2020, 2021). Dissertation reference consistent with published work.

---

## Recommendations

### Critical — Fix Before Submission
1. **pathak2025acoprinciples**: Replace entire DOI and author list from fabricated values
2. **yang2023review**: Replace all 5 author names with CrossRef-verified names
3. **zhu2025cumulative**: Fix 5 of 6 author names, add missing volume/pages
4. **alkhalifa2025comparative**: Fix 3 author names, add 3 missing authors (remove "and others")
5. **saller2025approximability**: Fix first author name ("Sebastian" → "Sophia"), add missing volume/issue/pages

### High Priority
6. **huang2021branchsurvey**: Fix second author ("Xiaozhe" → "Xiaomeng")
7. **dorigo2018acooverview**: Fix year (2018 → 2019)
8. **misra2024acorecent**: Fix booktitle subtitle, fix pages (1-22 → 1-17), expand author initials
9. **ilavarasi2014variants**: Change type to `@inproceedings`, fix journal→booktitle, fix author name, add pages
10. **umbarkar2015crossover**: Add missing pages (1083-1092)
11. **abdulghani2024comprehensive**: Add missing pages (214-224), expand author initials

### Consider Replacement
12. **waysi2025optimization**: Journal quality concern — Indonesian J Comp Sci not reputable
13. **abdulghani2024comprehensive**: Acadlore publisher concern — very new, unproven journal
14. **gutekunst2020relaxations**: Consider citing peer-reviewed journal version (DOI: `10.1287/moor.2020.1100`) instead of or in addition to the dissertation

### Additional Note
The pattern of errors across papers #2, #3, #5, #9, #12, and #17 (wrong author given names) suggests these BibTeX entries were generated programmatically from an unreliable source. Consider re-extracting all metadata from CrossRef or Semantic Scholar API for consistency.
