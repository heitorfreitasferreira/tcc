---
id: council-2026-06-02-validate-monograph
type: council
date: 2026-06-02
---

## Council Consensus: WARN

**Target:** Monografia "Otimização de Rotas de Patrulhamento com Drones: Um Estudo Comparativo de Meta-heurísticas Bioinspiradas"
**Author:** Heitor Freitas | **Advisor:** Prof. Dr. Claudiney Ramos Tinoco
**Mode:** validate
**Judges:** 4 (Orientador, Conteúdo, Redação, Estatística)

---

### Verdicts

| Judge | Perspective | Verdict | Confidence |
|-------|------------|---------|------------|
| **Orientador** | Overall Quality & Contribution | WARN | HIGH |
| **Conteúdo** | Technical Content & Methodology | WARN | HIGH |
| **Redação** | Writing, Structure & ABNT | WARN | HIGH |
| **Estatística** | Experimental Design & Statistical Rigor | WARN | HIGH |

---

### Shared Findings (All Judges Agree)

| Finding | Severity | Fix | Ref |
|---------|----------|-----|-----|
| **Encoding confound** — GA (permutation), PSO (random keys), ACO (3D constructive) use fundamentally different representations. Performance hierarchy conflates algorithm vs. encoding. | CRITICAL | Add explicit construct-validity threat discussion; reframe conclusions as "nas condições avaliadas (codificações distintas)" | Orientador F04, Conteúdo §3, Estatística F4 |
| **Success rate tolerance undocumented** — ACO 71.79% hit rate depends on floating-point tolerance (~1e-6). Exact match yields 60.35%. | MAJOR | Document ε threshold in §4.3.1 | Orientador F01, Estatística F12 |
| **Equal iterations ≠ fair comparison** — ACO uses 112× more CPU than GA at 100 pts (4.26s vs 38ms), yet compared on equal iteration budget. | MAJOR | Add "equidade da comparação" subsection discussing wall-clock vs. iteration budget | Orientador F03, Conteúdo §3.10 |
| **No effect sizes reported** — p-values alone (Friedman p=4.7e-31) conflate statistical and practical significance. With 1530 obs/method, trivial differences become significant. | MAJOR | Add Kendall's W, Cliff's δ, or rank η² alongside all p-values | Estatística §5, Orientador F07 |
| **No confidence intervals** — all point estimates (gap, hit rate, runtime) lack uncertainty measures. | MODERATE | Add bootstrap 95% CI for gap, hit rate, and timing means | Estatística §11 |
| **Hyperparameter sensitivity absent** — single fixed configuration per method; ACO may benefit from tuning more than GA/PSO. | MODERATE | Add sensitivity analysis or justify choices against literature | Estatística §8, Conteúdo §5 |
| **Lower bound 51% gap too loose** — LB valid but uninformative for quality assessment on large instances. | MODERATE | De-emphasize LB as quality metric; keep only as correctness guarantee | Orientador F06, Estatística §4 |
| **Equation indexing ambiguity** — eq:objetivo references π₀ without defining π₀=0 convention. | MODERATE | Explicitly state π₀=0 in sum notation | Conteúdo §1, F1 |
| **Class option mismatch** — `\documentclass[monografia]` outputs "Bacharel em Ciência da Computação" instead of PPGCO master's. | HIGH | Change to `dissertmst` and adjust preâmbulo | Redação §9 |
| **Acronym package unused** — `acronym` loaded but `\ac{}` never used; FSTSP appears without expansion. | MODERATE | Either use `\ac{}` consistently or remove the package | Redação §6 |

---

### Disagreements

| Issue | Orientador | Conteúdo | Redação | Estatística |
|-------|-----------|----------|---------|-------------|
| **Severity of PSO scalar r1/r2** | Minor (not flagged) | Minor — should document | Not flagged | Not flagged |
| **Section 5.4 (publications)** | Remove or fill | Not flagged | Remove or integrate | Not flagged |
| **Gap generalization (n≤15→n≤100)** | Not flagged (accepts as limitation) | Not flagged (accepts as limitation) | Not flagged | **High severity** — explicit threat to conclusion validity |

---

### Judge-Specific Insights

**Orientador (Advisor) — Strengths:**
- Formalization rigor (TSP-SD-ATP equations precise, O(n) evaluation after O(n³) precomputation)
- Reproducibility (full deterministic seeding, 4638 verified structured outputs, open-source code)
- Statistical methodology (Demšar 2006 protocol correctly applied — nonparametric, paired-by-instance, multiple-testing correction)
- Limitation transparency (four explicit limitations in §4.7, restated in §5.2)
- Scale of experimentation exceeds typical master's expectations

**Conteúdo (Technical Content) — Strengths:**
- Implementation verified against source code: GA OX crossover, ACO 3D pheromone, brute-force Heap algorithm, Hungarian O(n³) all correct
- Graph tensor indexing G[prev][curr][next] semantics validated by unit tests
- Statistical analysis: Friedman F-distribution via regularized incomplete beta function ✓, Nemenyi CD=0.605 ✓, Bonferroni-Holm ✓
- ACO 3D pheromone τ[i][j][k] and heuristic η=1/G[i][j][k] appropriate for TSP-SD-ATP

**Redação (Writing) — Strengths:**
- Consistent academic register, logical chapter flow, proper paragraph transitions
- All figures have "Fonte:" attribution (ABNT requirement)
- Proper Portuguese decimal separator {,} throughout tables and equations
- 30+ citations verified against .bib — no phantom/orphan references
- DOIs present in 26/32 entries (missing are classic books without DOI — acceptable)
- Flowcharts exist as TikZ standalone for all 5 methods

**Estatística (Statistical Design) — Strengths:**
- 51 seeds per method is generous (literature uses 10-30), enables reliable rank-based tests
- Friedman with Iman-Davenport correction correctly applied (randomized block design, n=30 blocks ≫ minimum 5)
- Bonferroni-Holm sequential rejection correctly applied (α/3, α/2, α/1)
- RCBD design (instances=blocks, methods=treatments, seeds=replicates) appropriate for stochastic algorithm comparison

---

### Critical Issues Requiring Action Before Defense

1. **🔴 Class option mismatch** (Redação F1) — `ppgco.cls` with `monografia` option outputs "Bacharel" text on title page. If this is PPGCO (master's), change to `dissertmst`. If intentional, adjust preâmbulo accordingly.

2. **🔴 Encoding confound not discussed as validity threat** (Orientador F04, Estatística F4) — The 3 representations (direct permutation, random keys, 3D pheromone) are confounded with algorithm identity. Must be reframed as "metaheuristic+encoding pairs" and discussed as construct-validity threat per Shadish-Cook-Campbell.

3. **🔴 Success rate tolerance** (Orientador F01) — Document floating-point ε used. 71.79% drops to 60.35% on exact match.

4. **🔴 Equal budget confound** (Orientador F03) — Add "fair comparison" discussion: equal iterations ≠ equal computational effort. ACO runs 112× longer at 100 pts.

5. **🟡 Effect sizes absent** (Estatística F3, F5) — Add Kendall's W and Cliff's δ to complement p-values.

6. **🟡 Confidence intervals absent** (Estatística F14) — Add 95% CI for gap, hit rate, timing estimates.

7. **🟡 Equation indexing** (Conteúdo F1) — Clarify π₀=0 convention in eq:objetivo.

8. **🟡 Acronym package** (Redação F2) — Either use `\ac{}` consistently or remove the package.

---

### Strengths (Preserve and Highlight)

| Aspect | Detail |
|--------|--------|
| **Experimental scale** | 4638 runs across 30 instances × 51 seeds — exceeds typical master's scope |
| **Implementation quality** | Go/Cobra modular architecture, verified by tests, 4638 structured output files |
| **Statistical protocol** | Demšar (2006) correctly applied: Friedman + Nemenyi + Wilcoxon-Holm |
| **Limitation transparency** | 4 explicit limitations discussed honestly in dedicated sections |
| **Reproducibility** | Deterministic seeding, open-source code, structured JSON/JSONL output artifacts |
| **Formalization** | TSP-SD-ATP equations precise, 3D tensor representation clearly defined |

---

### Recommendation

**Conditional PASS — address the 8 issues above before submission/defense.**

The monograph is fundamentally sound: the TSP-SD-ATP formulation is well-motivated, the implementation is correct and verified, the experimental scale (4638 runs) is impressive, the statistical protocol follows established methodology (Demšar 2006), and the writing quality is above average for a TCC/monograph. None of the critical issues require re-running experiments — they are all textual clarifications, methodological framing adjustments, or documentation gaps.

The three major findings (encoding confound, equal budget, success rate tolerance) are framing/transparency issues that a defense committee will identify. Addressing them converts a WARN into a clear PASS.

---

*Council completed in ~120s. 4/4 judges responded.*
