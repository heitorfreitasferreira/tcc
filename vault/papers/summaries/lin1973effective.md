# Lin & Kernighan (1973) — Summary

**Title:** An Effective Heuristic Algorithm for the Traveling-Salesman Problem
**Authors:** S. Lin, B. W. Kernighan
**Source:** Operations Research, 21(2), pp. 498–516, INFORMS
**BibTeX key:** `lin1973effective`

---

## 1. Problem and Motivation

The paper addresses the symmetric traveling-salesman problem (TSP): given an \(n \times n\) symmetric distance matrix, find a minimum-length tour visiting each city exactly once. Exact algorithms at the time required prohibitive running times and existing heuristics degraded in effectiveness beyond roughly 60 cities. The authors sought a heuristic that produces optimal or near-optimal solutions with high frequency and whose running time scales modestly with \(n\).

## 2. Core Method or Approach

- **Variable-depth \(k\)-opt local search**: generalizes fixed-\(k\) interchange (e.g., 2-opt, 3-opt) by building an exchange sequence element by element rather than fixing \(k\) in advance.
- **Sequential edge exchange**: at each step, one tour edge \(x_i\) is broken and replaced by a non-tour edge \(y_i\). Gains are additive: \(G = \sum (|x_i| - |y_i|)\).
- **Gain criterion as stopping rule**: only sequences whose every partial sum is positive are explored. When the cumulative gain \(G_k\) drops below the best improvement \(G^*\) found so far, search along that path stops.
- **Limited backtracking**: full backtracking is restricted to levels 1 and 2 (first two choices of \(y_1\) and \(y_2\)); only a few alternatives (up to 5) are tried at each level.
- **Key refinements**:
  - **Lookahead**: \(y_i\) is chosen to maximize \(|x_{i+1}| - |y_i|\) (not just nearest neighbor), using the five shortest available candidates.
  - **Reduction**: after several local optima are found, links common to all of them are frozen (cannot be broken at depth \(\geq 4\)), dramatically cutting search space and run time.
  - **Checkout avoidance**: previously seen local optima are recognized and skipped.
  - **Nonsequential exchange**: a post-optimization test for nonsequential 4-opt improvements is applied as insurance.

## 3. Main Results

- **Optimality frequency**: optimum found in a single random-start trial with probability ~1.0 for \(n \leq 42\), dropping to 0.2–0.3 for 100-city problems.
- **Running time**: grows approximately as \(n^{2.2}\). A typical 100-city problem requires ~25 seconds pre-reduction and ~8 seconds post-reduction on a GE635.
- **Classical benchmark problems** (20–100 cities): optimum or best-known tour obtained in every case. Improved on 3 of 5 Krolak et al. (1971) 100-city problems: \(\#25\) (22148 vs 22193), \(\#26\) (20749 vs 20852), \(\#28\) (22068 vs 22115).
- **Random Euclidean points** (unit square, 30–110 cities): optimum frequency 0.85 (n=30) to 0.21 (n=110); post-reduction frequencies are higher.
- **Random distance matrix** (non-metric, 30–110 cities): harder class — optimum frequency 0.46 (n=30) to 0.09 (n=110).
- **Real-world test**: 318-point laser drilling problem solved by optimizing ~100-point sub-tours and splicing them; result was 0.85 inches shorter (~2% improvement) than the hand-crafted solution.
- **Small set of distinct local optima**: e.g., 48-city problem produced only 4 distinct solutions over 200+ trials, with the minimum occurring 30% of the time — enabling strong statistical confidence in optimality.

## 4. Strengths and Limitations

**Strengths:**
- First variable-depth local search for TSP; foundation for the most successful TSP heuristics for decades.
- Empirically strong: finds optimum with high probability even for 100 cities.
- Well-engineered with multiple refinements (lookahead, reduction, checkout avoidance) that substantially improve run time without sacrificing solution quality.
- Statistical framework for assessing confidence in optimality via distribution of local optima.
- Broad applicability demonstrated beyond TSP (graph partitioning in Kernighan & Lin 1970).

**Limitations:**
- No theoretical optimality guarantee; relies on statistical confidence.
- Does not handle all non-sequential exchange configurations; some tour improvements are missed.
- Storage requirements restrict experiments to \(\leq 110\) cities in the original implementation.
- Effectiveness degrades with problem size and problem class (non-metric problems harder than Euclidean).
- Reduction may bias search if early local optima are suboptimal.

## 5. Practical Takeaway for Researchers

The Lin-Kernighan heuristic established the dominant paradigm for TSP local search: variable-depth sequential exchanges governed by a greedy gain criterion with limited backtracking. The key insight — that the sequence of gains must maintain positive cumulative sums — dramatically prunes the search space and remains the conceptual core of modern TSP solvers (including LKH). The reduction technique (freezing edges common to multiple local optima) is an early and effective instance of what later became *candidate set* pruning in advanced implementations. For any stochastic local search on combinatorial problems, the paper's philosophy of iterating random starts through a powerful local-optimum mapper remains relevant.
