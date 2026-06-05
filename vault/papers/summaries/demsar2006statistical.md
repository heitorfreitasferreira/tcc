# Demšar (2006) — Statistical Comparisons of Classifiers over Multiple Data Sets

**Reference:** Janez Demšar. *Statistical Comparisons of Classifiers over Multiple Data Sets.* Journal of Machine Learning Research 7 (2006), 1–30.

## 1. Problem and Motivation

Machine learning papers routinely compare multiple classifiers across multiple data sets, yet the community lacked an established, statistically sound procedure for doing so. A survey of ICML proceedings (1999–2003) revealed that 25–50% of relevant papers attempted some form of statistical comparison, but most relied on inappropriate practices — typically many uncorrected pairwise t-tests, ignoring the multiple-comparisons problem. The paper addresses this gap by reviewing existing tests, evaluating their theoretical suitability for typical ML data, and providing empirical evidence to guide recommendations.

## 2. Core Method or Approach

- **Survey of current practice:** Analyzed 174 ICML papers (1999–2003) that compared at least two classifiers, cataloguing sampling methods, evaluation measures, and statistical tests actually used.
- **Theoretical analysis of candidate tests:** Examined assumptions and suitability of the paired t-test, Wilcoxon signed-ranks test, sign test (two classifiers), and repeated-measures ANOVA, Friedman test, and multiple post-hoc procedures (Nemenyi, Bonferroni-Dunn, Holm, Hochberg, Hommel) for three or more classifiers.
- **Empirical comparison:** Evaluated test behavior on 7 real classifiers (C4.5 variants, naive Bayes, kNN) over 40 UCI data sets, using 1000 random draws of 10 data sets each. Introduced a bias parameter *k* to control the degree of difference between classifiers and measured rejection rates and replicability (both Bouckaert's *R(e)* and a variance-based *R(p)*).
- **CD (Critical Difference) diagrams:** Proposed a compact graphical representation for post-hoc test results, showing average ranks on an axis with critical difference intervals.
- **Recommendation framework:** Wilcoxon for two classifiers; Friedman + Nemenyi (all-pairs) or Bonferroni-Dunn/Holm (vs. control) for multiple classifiers.

## 3. Main Results

- **ICML survey:** ~50% of comparison papers used t-tests without multiple-comparison correction; only a few applied Bonferroni correction. Classification accuracy was the dominant metric (67–84%), with AUC rarely used (0–13%).
- **Two-classifier comparison:** The Wilcoxon signed-ranks test consistently yielded lower p-values and higher rejection rates than the paired t-test (both absolute and relative differences), suggesting real-world ML data violates parametric assumptions (normality, commensurability). The sign test was substantially weaker.
- **Multiple classifiers:** The Friedman test rejected the null hypothesis more often than ANOVA at smaller inter-algorithm differences; at larger differences the two converged. Replicability of the Friedman test was higher than ANOVA when measured by *R(p)*.
- **Post-hoc tests:** The Bonferroni-Dunn test (for comparisons against a control) was more powerful than Nemenyi (for all-pairs) since it adjusts for fewer comparisons. Holm and Hochberg stepwise procedures gave practically identical results and rejected slightly more hypotheses than Bonferroni-Dunn.
- **Replicability concern:** All tests showed limited replicability when differences were marginal, motivating the use of as many data sets as possible.

## 4. Strengths and Limitations

**Strengths:**
- First comprehensive treatment of statistical comparisons over multiple data sets in ML.
- Combines theoretical argumentation with empirical validation on real classifiers and data sets.
- Clear, actionable recommendations with worked numerical examples.
- CD diagrams provide an intuitive, space-efficient visualization standard.

**Limitations:**
- No formal Type I / Type II error analysis; the author argues the "ground truth" depends on the kind of difference one intends to measure, making error-rate experiments ill-posed without artificial data.
- Empirical study used only 10 data sets per random draw and 7 classifiers from a single family (decision trees, naive Bayes, kNN).
- Limited to classification accuracy and AUC; other metrics (precision/recall, F-measure) not tested.
- The 40 UCI data sets are dated relative to modern benchmarks.

## 5. Practical Takeaway for Researchers

- **Two classifiers → Wilcoxon signed-ranks test.** It is non-parametric, robust to outliers, does not assume normality, and empirically outperforms the paired t-test on typical ML data.
- **Three or more classifiers → Friedman test + post-hoc.** Use Nemenyi for all-pairs comparisons; use Bonferroni-Dunn or Holm when comparing against a single control method.
- **Visualize with CD diagrams.** Plot average ranks on a horizontal axis, mark critical difference intervals, and connect groups that are not significantly different.
- **Avoid:** averaging raw scores across data sets (incommensurable), uncorrected pairwise t-tests (inflated Type I error), and counting only "significant wins/losses" (arbitrary threshold).
- **Use as many data sets as possible** — replicability degrades when differences are marginal and sample sizes are small.
