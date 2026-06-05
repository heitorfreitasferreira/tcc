# Nagata & Kobayashi (2013) — A Powerful Genetic Algorithm Using Edge Assembly Crossover for the TSP

**Reference:** Nagata, Y., & Kobayashi, S. (2013). A Powerful Genetic Algorithm Using Edge Assembly Crossover for the Traveling Salesman Problem. *INFORMS Journal on Computing*, 25(2), 346–363. DOI: [10.1287/ijoc.1120.0506](https://doi.org/10.1287/ijoc.1120.0506)

## 1. Problem and Motivation

For nearly 40 years, the best heuristic algorithms for the TSP have been based on the Lin–Kernighan (LK) local search. Previous genetic algorithms (GAs) could not compete with LK-based methods due to (i) high computational cost of crossover operators and (ii) lack of local optimization capability. The authors aim to construct a GA that outperforms state-of-the-art LK-based heuristics **without using any LK-based algorithm**, thereby demonstrating that GAs can be a viable direction for TSP.

## 2. Core Method or Approach

- **Edge Assembly Crossover (EAX):** Generates offspring by combining edges from two parent tours. The two parent edge sets are overlaid into a multigraph, partitioned into alternating AB-cycles, and a subset (E-set) is selected to replace edges from parent A with edges from parent B, plus a few new short edges. The resulting intermediate solution (set of subtours) is merged into a valid tour via a greedy minimum-spanning-tree-like procedure.

- **Localization of EAX (Stage I):** Uses the *single strategy* (selecting one AB-cycle at a time) so offspring differ from parent A by only a few edges. Coupled with efficient O(1) update techniques, this reduces generation cost far below O(N), making the GA computationally competitive with local search.

- **Block2 strategy (Stage II):** A global version of EAX that increases E-set size while keeping the number of subtours low. It redefines AB-cycles to merge adjacent ineffective cycles, then uses **tabu search** to select AB-cycles that minimize the number of C-vertices (#C) in the intermediate solution. Lower #C means fewer subtours and better offspring quality when parents are already high-quality tours.

- **Entropy-preserving selection:** Replaces the traditional thermodynamical GA (TDGA) framework. Offspring are evaluated by the ratio of tour-length improvement to loss of population edge-entropy (ΔL/ΔH), preserving diversity at negligible computational cost with no tunable temperature parameter. Only offspring that improve the parent or increase diversity are accepted.

- **Two-stage framework:** Stage I uses localized EAX until stagnation; Stage II switches to block2 strategy for the final refinement phase. Population size = 300, offspring per parent pair = 30.

## 3. Main Results

- **Benchmarks:** 57 well-studied TSP instances from TSPLIB, National TSPs, and VLSI TSPs (up to 85,900 cities), plus 6 Art TSP instances (100,000–200,000 cities).
- **Solution quality:** GA-EAX found optimal or best-known solutions for nearly all instances up to 39,000 vertices; improved **11 best-known solutions** (5 on standard benchmarks, 6 Art TSPs).
- **Comparison to LKH-2:** GA-EAX significantly outperformed LKH-2 (full patching, K=6) on 50 out of 57 instances (Wilcoxon rank sum test, p < 0.05), with generally shorter computation times.
- **Ablation evidence:** Localization drastically reduces computation time vs. random strategy; block2 strategy finds optima where five-multiple strategy cannot; entropy-preserving selection dominates both greedy and distance-preserving selection.

## 4. Strengths and Limitations

**Strengths:**
- First GA to surpass LK-based heuristics on TSP without using LK, challenging the 40-year consensus.
- Entropy-preserving selection is parameter-free, computationally cheap, and empirically superior to greedy/distance-based alternatives.
- Source code is publicly available; methodology is transparent and reproducible.
- Scales to 200,000-city instances with memory-saving implementation.

**Limitations:**
- Underperforms on VLSI TSP instances where vertices are arranged in lattice patterns (e.g., pla7397, pla33810, pla85900) — the authors attribute this to many distinct near-optimal solutions hindering population convergence.
- Not tested on very large instances (millions of cities) where LK-based algorithms have been applied; the authors note that parallelization would be required.
- The two-stage switch criterion (G-based stagnation) is somewhat heuristic.

## 5. Practical Takeaway for Researchers

EAX with localization and block2 strategy provides a state-of-the-art GA for TSP that does not depend on LK-based local search. The entropy-preserving selection model — evaluating offspring by ΔL/ΔH — is a lightweight, parameter-free diversity mechanism applicable to other combinatorial optimization GAs. For TSP instances without lattice structure, GA-EAX is a strong alternative to LKH-2, particularly for instances with 10,000–200,000 cities where its advantage is most pronounced.
