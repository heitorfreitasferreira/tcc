---
title: "A Comparative Study of Metaheuristics Methods for Solving Traveling Salesman"
authors: "A. Chandra, A. Naro"
year: 2022
journal: "International Journal of Information Science & Technology (iJIST), Vol. 6, No. 2"
doi: null
tags: [metaheuristics, TSP, comparative-study, ABC, ANOVA, Tukey]
---

## 1. Problem and Motivation

Transportation logistics efficiency relies on effective route optimization algorithms. The symmetric Traveling Salesman Problem (TSP) — an NP-hard combinatorial problem with \((n-1)!/2\) possible tours — cannot be solved by exact methods for real-world instances within practical time limits. This paper compares eight metaheuristic methods on a TSP instance of 70 major cities across Java island, Indonesia, motivated by the island's high population density and rapid growth in motor vehicle traffic, which demand greener routing solutions.

## 2. Core Method or Approach

- Eight metaheuristic algorithms are benchmarked: **Genetic Algorithm (GA, 1975)**, **Simulated Annealing (SA, 1983)**, **Tabu Search (TS, 1986)**, **Ant Colony Optimization (ACO, 1992)**, **Particle Swarm Optimization (PSO, 1995)**, **Artificial Bee Colony (ABC, 2005)**, **Improved Fruit Fly Optimization Algorithm (EFOA, 2016)**, and **Artificial Atom Algorithm (A³, 2018)**.
- All algorithms implemented in MATLAB 2015a; each runs 20 independent trials on the 70-city Java dataset and on the TSPLIB58 benchmark for external validation.
- Distance matrix computed via Euclidean distances between geographic coordinates; output measured in kilometers.
- Statistical validation via **one-way ANOVA** to test equality of means across methods, followed by **Tukey-Kramer post-hoc test** to identify which specific method pairs differ significantly.

## 3. Main Results

- **Best performer**: ABC achieved the shortest route (2,447 km) and best average (2,458 km). Ranking by distance: ABC > SA (2,683) > A³ ≈ EFOA ≈ GA (all 2,727) > ACO (2,851) > TS (3,190) > PSO (9,623).
- **PSO severely underperformed**: its best distance (9,623 km) is ~4× worse than ABC, suggesting poor suitability of the discrete PSO encoding or inadequate parameter tuning.
- On **TSPLIB58**: ABC again led (25,400), with GA second (25,902); PSO was again the worst (99,887).
- **ANOVA**: \(F(7, 152) = 2519.06\), \(p \approx 1.03 \times 10^{-153}\), \(\eta^2 = 99.21\%\) — overwhelmingly significant differences among methods.
- **Tukey test**: 20 out of 28 pairwise comparisons are statistically significant. Non-significant pairs include GA–ACO, GA–A³, GA–EFOA, SA–TS, SA–ACO, ACO–A³, ACO–EFOA, and A³–EFOA.

## 4. Strengths and Limitations

**Strengths**: Broad scope covering eight methods spanning four decades of metaheuristic development; includes both classical (GA, SA) and recent (EFOA, A³) algorithms; dual evaluation on real-world and benchmark instances; rigorous statistical analysis with ANOVA and Tukey-Kramer post-hoc comparison; clear parameter reporting for several methods.

**Limitations**: No ground-truth optimal or best-known tour length is reported for the Java dataset, making it unclear how close any method is to optimality; PSO performance is anomalously poor and not adequately discussed; the paper does not report convergence behavior, runtime costs, or parameter sensitivity; only 20 runs per method, which limits power for non-parametric comparisons; implementations rely on third-party MATLAB code, introducing possible variability in algorithmic fidelity; the TSPLIB58 results are mentioned but not subjected to the same statistical treatment.

## 5. Practical Takeaway for Researchers

ABC is a strong, simple choice for symmetric TSP instances of moderate size (~70 cities), outperforming both classical methods (GA, SA, TS) and newer algorithms (EFOA, A³) at comparable levels of parameter tuning effort. PSO in its discrete form should be applied with caution to TSP — or at minimum benchmarked against standard permutation-based encodings before being included in comparisons. Always report an optimal or best-known baseline and include convergence curves to contextualize final solution quality.
