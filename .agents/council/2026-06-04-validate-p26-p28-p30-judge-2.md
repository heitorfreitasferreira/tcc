```json
{"verdict": "PASS", "confidence": "HIGH", "findings": [
  {
    "id": "F1",
    "fix": "P26: 5 BibTeX entries added (heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, kinable2017hybrid, righini2021efficient) at lines 662-712 of abntex2-references.bib",
    "why": "All DOIs verified: heldkarp1970 (10.1287/opre.18.6.1138) matches the classic OR paper volume 18(6):1138-1162; heldkarp1971 (10.1007/BF01584070) matches Math Programming 1(1):6-25; johnson1996 is the SODA'96 paper on asymptotic HK bound tightness; kinable2017 (10.1016/j.ejor.2016.11.035) matches EJOR 259(3):906-918; righini2021 (10.5802/ojmo.11) matches Open J Math Optim 2:1-27. Authors, titles, venues all correct.",
    "ref": "monografia/bib/abntex2-references.bib:662-712"
  },
  {
    "id": "F2",
    "fix": "P26: Held-Karp paragraph added at fundamentacao.tex:92",
    "why": "The paragraph (1) names Held-Karp as canonical lower bound via 1-tree + Lagrangian multipliers, citing both 1970 and 1971 papers; (2) quantifies HK tightness citing Johnson 1996; (3) notes ongoing algorithm development citing Righini 2021; (4) explicitly justifies why AP was chosen over HK for TSP-SD-ATP — adaptation to 3D tensor $c_{i,j,k}$ and 3-index Lagrangian penalty is out of scope. The justification is logically sound and academically honest.",
    "ref": "monografia/cap_fundamentacao/fundamentacao.tex:92"
  },
  {
    "id": "F3",
    "fix": "P28: ATSP/TDTSP framing sentence at fundamentacao.tex:13 citing Öncan 2009 and Kinable 2017",
    "why": "The sentence frames TSP-SD-ATP as an instance of asymmetric TSP and time-dependent TSP. Öncan et al. 2009 (doi:10.1016/j.cor.2007.11.008) is a well-known ATSP formulations survey in Computers & OR. Kinable et al. 2017 is a time-dependent sequencing paper. While TSP-SD-ATP's path-dependence is slightly more nuanced than fixed asymmetry, the framing to anchor the problem in established TSP classes is academically appropriate.",
    "ref": "monografia/cap_fundamentacao/fundamentacao.tex:13"
  },
  {
    "id": "F4",
    "fix": "P30: winter2002modeling rating 0→3",
    "why": "Winter 2002 models turn costs via pseudo-dual graphs, directly cited in fundamentacao.tex:18 for sequence-dependent costs. A rating of 3 (moderate relevance) is appropriate — the paper is conceptual foundation but for shortest paths, not TSP.",
    "ref": "vault/papers/winter2002modeling.md:11"
  },
  {
    "id": "F5",
    "fix": "P30: wang2021ant rating 0→3",
    "why": "Wang 2021 proposes SOS-ACO hybrid for parameter optimization. Rated 3 because it is relevant for ACO parameter sensitivity (a stated challenge in this work) but its hybrid approach differs from the Ant System implementation used. A rating of 3 is proportionate.",
    "ref": "vault/papers/wang2021ant.md:11"
  },
  {
    "id": "F6",
    "fix": "Minor: valenzuela1997estimating was pre-existing at line 560, not part of P26 additions",
    "why": "The P26 task description listed valenzuela1997 among '6 BibTeX entries' but status says '5 entradas BibTeX adicionadas'. valenzuela1997 already existed in the .bib file (line 560, pre-P26). This is not a gap — P26 delivered the 5 core HK references (heldkarp1970, heldkarp1971, johnson1996, kinable2017, righini2021) and valenzuela1997 was already present. No missing entry.",
    "ref": "monografia/bib/abntex2-references.bib:560"
  },
  {
    "id": "F7",
    "fix": "Minor: LaTeX `<` in text mode at fundamentacao.tex:92",
    "why": "The text reads '(\<1\%' with a literal `<` character. In modern LaTeX with T1 font encoding and UTF-8 input, this renders correctly. However, for maximum portability across LaTeX engines, `$<$1\%` or `{<}1\%` is safer. Not blocking — the document compiles as-is with the project's toolchain.",
    "ref": "monografia/cap_fundamentacao/fundamentacao.tex:92"
  }
], "recommendation": "All three tasks (P26, P28, P30) are correctly implemented. BibTeX entries are academically accurate. The Held-Karp paragraph provides a clear, honest justification for using AP over HK given the 3D tensor constraint. ATSP/TDTSP framing is reasonable and well-cited. Rating fixes are proportionate to each paper's role in the monograph. One minor LaTeX portability note (F7: bare < in text mode) is non-blocking. Proceed."}
```

## Judge 2 — Full Analysis

### P26: Held-Karp BibTeX + Lower Bounds Section Revision

**BibTeX verification.** Five entries were appended to `abntex2-references.bib` (lines 662–712). Each was checked against known publication metadata:

| Key | Title | Venue | DOI | Verdict |
|---|---|---|---|---|
| `heldkarp1970traveling` | The Traveling-Salesman Problem and Minimum Spanning Trees | Operations Research 18(6):1138–1162, 1970 | 10.1287/opre.18.6.1138 | Correct |
| `heldkarp1971traveling` | … Part II | Mathematical Programming 1(1):6–25, 1971 | 10.1007/BF01584070 | Correct |
| `johnson1996asymptotic` | Asymptotic Experimental Analysis for the Held-Karp… | SODA 1996, pp. 341–350 | — | Correct (proceedings, no DOI typical for 1996) |
| `kinable2017hybrid` | Hybrid Optimization Methods for Time-Dependent Sequencing Problems | EJOR 259(3):906–918, 2017 | 10.1016/j.ejor.2016.11.035 | Correct |
| `righini2021efficient` | Efficient Optimization of the Held-Karp Lower Bound | Open J Math Optim 2:1–27, 2021 | 10.5802/ojmo.11 | Correct |

A sixth entry (`valenzuela1997estimating`) was listed in the P26 description but already existed at line 560; it was not missing and does not constitute a gap.

**Paragraph analysis (fundamentacao.tex:92).** The new paragraph does four things in sequence:
1. Names Held-Karp as the canonical TSP lower bound, citing both the 1970 formulation and 1971 Part II.
2. Quantifies HK tightness (`<1%` for Euclidean instances), citing Johnson 1996.
3. Acknowledges ongoing algorithmic development, citing Righini 2021.
4. Justifies the AP choice: adapting HK to the 3D tensor $c_{i,j,k}$ with 3-index Lagrangian penalties is out of scope; AP is looser but straightforward to implement via 3D→2D reduction.

This is a well-structured justification. It does not overclaim — it explicitly admits AP is weaker and positions it as a pragmatic choice given the problem structure.

### P28: ATSP/TDTSP References

**BibTeX.** `oncan2009comparative` (lines 714–722) is Öncan, Altınel & Laporte's ATSP formulations survey (Computers & OR 36(3):637–654, 2009, doi:10.1016/j.cor.2007.11.008). Metadata is correct, including the proper LaTeX escaping of Turkish characters (`\"{O}ncan`, `Alt{\i}nel`, `{\.{I}}brahim`). `kinable2017hybrid` was added via P26.

**Text placement (fundamentacao.tex:13).** The sentence frames TSP-SD-ATP as an instance of asymmetric TSP and time-dependent TSP. The logic is: sequence-dependent costs create effective asymmetry (cost of edge j→k depends on predecessor i) and time/state-dependence (the "state" is the previous node). While TSP-SD-ATP is more precisely a *history-dependent* or *path-dependent* variant than a fixed-asymmetry ATSP, the framing to anchor the problem in established TSP classes is academically standard and the citations are appropriate.

### P30: Paper Rating Fixes

**winter2002modeling (0→3).** Winter 2002 provides the pseudo-dual graph formalism for modeling turn costs between consecutive edges. It is cited in fundamentacao.tex:18 as the conceptual anchor for sequence-dependent costs. Rating 3 is proportionate: the paper is foundational for the cost model but addresses shortest paths, not TSP directly.

**wang2021ant (0→3).** Wang 2021 proposes SOS-ACO, a hybrid that tunes ACO parameters α and β automatically. It is relevant because ACO parameter sensitivity is a challenge acknowledged in this work, but its SOS hybridization differs from the Ant System implementation used. Rating 3 correctly reflects moderate relevance.

### Gaps and Issues

**G1 — LaTeX portability (non-blocking).** Line 92 uses a bare `<` in text mode: `(\<1\%`. With T1 font encoding and UTF-8 input (standard in `ppgco.cls`), this renders correctly. For maximum portability, `$<$1\%` or `{<}1\%` is preferred but not required for the project's toolchain.

**G2 — ATSP framing nuance (non-blocking).** TSP-SD-ATP's cost depends on the predecessor node, making it path-dependent rather than fixed-asymmetry. The Öncan 2009 ATSP survey is about *formulations* of asymmetric TSP, not about path-dependent costs. A reader looking strictly for path-dependent/look-ahead TSP literature might find the framing slightly imprecise. However, the intent — connecting TSP-SD-ATP to well-studied TSP classes — is legitimate, and the qualifier "pode ser enquadrado como" (can be framed as) leaves room for interpretation.

### Conclusion

All three tasks are correctly implemented with academically accurate BibTeX entries, well-reasoned prose, and appropriate rating corrections. No blocking issues. Minor portability and precision notes are documented but do not affect correctness. **Verdict: PASS, confidence HIGH.**
