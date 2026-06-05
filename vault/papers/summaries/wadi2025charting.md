# wadi2025charting — Summary

- **Title:** Charting New Routes: Comparing Swarm-Based Approaches to the Traveling Salesman Problem
- **Authors:** Ali Hassan Ahmed Wadi, Shahla Uthman Umar (University of Kirkuk, Iraq)
- **Venue:** International Journal of Computational Methods and Experimental Measurements, Vol. 13, No. 2, pp. 371–379, 2025
- **DOI:** https://doi.org/10.18280/ijcmem.130214

## 1. Problem and motivation

The Traveling Salesman Problem (TSP) is NP-hard; exact solvers such as Branch and Bound (BB) and Dynamic Programming (DP) explode in time and memory beyond small instances (≈10–15 cities). The authors argue that swarm-based metaheuristics — PSO, ACO, and especially the less-studied Elephant Herding Optimization (EHO) — need a systematic head-to-head comparison on execution time, solution quality, and scalability to guide practitioners selecting algorithms for logistics, transport, and smart manufacturing.

## 2. Core method or approach

- Benchmarked three swarm algorithms (ACO, PSO, EHO) plus two exact baselines (BB, DP) on TSP instances with 5, 10, 14, 19, 100, and 150 cities.
- Measured four metrics: best tour cost, execution time, CPU time used, and time complexity.
- Described each algorithm with pseudocode, flowcharts, and canonical parameter roles (α, β, ρ for ACO; ω, c₁, c₂ for PSO; clan size, learning rate α, migration operator for EHO).
- Used manual parameter tuning (no adaptive or automated strategy) on a single laptop (MATLAB, Intel i5-8350U, 16 GB RAM).
- No mention of repetitions, variance, or statistical tests — results appear to be single-run values.

## 3. Main results

- **EHO dominated solution quality across all city sizes.** For 5 cities, EHO reported best cost 24.93 (vs. BB 160.64, DP 152.98, ACO 265.32, PSO 259.80). For 150 cities, EHO cost 6352.34 vs. ACO 1333.12 and PSO 6886.67.
- **BB and DP became infeasible beyond 14 cities.** BB failed to terminate at 19 cities; DP reached memory/computation limits at 100 cities.
- **ACO offered the best cost–time tradeoff** on large instances (100–150 cities), achieving lower costs than PSO with execution times of 2–3 s (vs. PSO <0.14 s but far worse costs).
- **PSO was fastest** but produced poor solution quality on large instances (e.g., 6886.67 at 150 cities vs. EHO 6352.34).
- The paper claims EHO is "quicker than the others" — but the raw timing data in Table 2 does not consistently support this (EHO and PSO have similar low execution times on large instances).

## 4. Strengths and limitations

**Strengths:**
- Direct comparison of an underexplored algorithm (EHO) against mature baselines on a classic benchmark.
- Covers a wide span of problem sizes (5–150 cities), from trivial to computationally demanding.
- Provides pseudocode and flowcharts for reproducibility.

**Limitations:**
- No statistical rigor: single-run results without confidence intervals, standard deviations, or repeated trials; stochastic algorithms cannot be compared meaningfully with one sample.
- The EHO result for 5 cities (24.93) is implausibly low given that BB returned 160.64 and DP 152.98; likely an implementation or reporting error that undermines the core claim.
- No use of standard TSP benchmark libraries (TSPLIB), making results unverifiable and incomparable with the wider literature.
- Parameter tuning is acknowledged as important but handled manually without systematic grid search or sensitivity analysis.
- The literature review is dated and contains unusual entries (e.g., references to atmospheric TSP concentration, which is unrelated).
- Prose quality has noticeable errors throughout, suggesting limited revision.

## 5. Practical takeaway for researchers

The paper highlights EHO as a promising metaheuristic for large TSP instances, but the evidence is weakened by single-run reporting and an anomalous small-instance result. A fairer comparison would repeat each algorithm 30+ times on standard TSPLIB instances, report distributions (not single points), and apply non-parametric statistical tests. Researchers building on this work should treat the EHO-vs-ACO ranking as a hypothesis to verify rather than a settled conclusion.
