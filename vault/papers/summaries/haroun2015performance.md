## Haroun et al. (2015) — A Performance Comparison of GA and ACO Applied to TSP

**Reference:** Sabry Ahmed Haroun, Benhra Jamal, El Hassani Hicham. *A Performance Comparison of GA and ACO Applied to TSP.* International Journal of Computer Applications, Vol. 117, No. 19, May 2015.

**DOI/Link:** Not available (IJCA, open access at www.ijcaonline.org).

---

### 1. Problem and Motivation

The Traveling Salesman Problem (TSP) is an NP-hard combinatorial optimization problem ubiquitous in engineering, logistics, and planning. This work directly compares the empirical performance of two nature-inspired metaheuristics — Genetic Algorithm (GA) and Ant Colony Optimization (ACO) — in solving both classical symmetric TSP benchmarks and a real-world asymmetric TSP from urban transportation in Casablanca, Morocco.

### 2. Core Method or Approach

- **GA:** Canonical genetic algorithm with crossover and mutation operators, implemented in C++ using GAlib, operating over a randomly initialized population of chromosomes (solutions).
- **ACO:** Ant System (AS) variant by Dorigo et al., implemented in native C#, where artificial ants construct tours stochastically using pheromone trails and visibility heuristics, with local and global pheromone update rules.
- **Benchmark instances:** Three symmetric Euclidean TSPs from TSPLIB — Berlin52 (52 cities, optimum 7542), Eil76 (76 cities, optimum 538), A280 (280 cities, optimum 2579).
- **Real-world instance:** Casablanca40 — an asymmetric TSP (ATSP) modeled as a directed weighted graph with 40 nodes and 1560 arcs, representing urban route markers in Casablanca; no mathematically proven optimum exists (search space ≈ 2.04 × 10^46 solutions).
- **Hardware/environment:** Windows 64-bit, Intel i5-3470 3.20 GHz, 4 GB RAM.

### 3. Main Results

| Instance | GA Tour | GA Error % | GA Time (s) | ACO Tour | ACO Error % | ACO Time (s) |
|----------|---------|------------|-------------|----------|-------------|---------------|
| Berlin52 | 7542 (opt.) | 0.00 | 3.67 | 7575 | 0.44 | 17.73 |
| Eil76 | 570 | 5.95 | 13.52 | 551 | 2.42 | 73.91 |
| A280 | 3218 | 24.78 | 531.09 | 3092 | 19.89 | 767.71 |
| Casablanca40 | 125.97 km | — | 24.76 | 125.22 km | — | 20.62 |

- GA found the optimal tour for Berlin52; ACO did not match the optimum on any symmetric instance but had consistently lower error rates than GA on Eil76 and A280.
- ACO achieved a better tour (125.216 km) than GA (125.974 km) on Casablanca40 — qualitatively 0.60% better and computationally 20.08% faster.
- Error rate increases with problem size for both algorithms; the gap favoring ACO widens as instances grow larger.
- GA shows steady improvement before stagnating; ACO exhibits two-phase convergence — rapid early descent driven by pheromone accumulation, followed by stagnation, with late minor improvements via pheromone evaporation enabling diversification.
- Both algorithms scaled to 280 cities without pathologically failing, demonstrating robustness over problem size variation.

### 4. Strengths and Limitations

**Strengths:**
- Side-by-side comparison on both synthetic benchmarks and a real asymmetric logistics case, strengthening ecological validity.
- Clear, tabulated quantitative results including error rates, absolute tour lengths, and wall-clock times.
- Discussion of convergence dynamics (stagnation, two-phase behavior in ACO) goes beyond final-metric reporting.

**Limitations:**
- Single-run reporting — no statistical measures (mean, std, confidence intervals) across multiple seeds; stochastic algorithm comparison is therefore anecdotal.
- No systematic parameter tuning study; parameters were set based on author experience alone, limiting reproducibility.
- Language discrepancy (GA in C++, ACO in C#) introduces an uncontrolled implementation-performance confound.
- Real-world instance lacks an optimal baseline, so relative performance is only pairwise, not absolute.
- No inclusion of local search hybrids or modern ACO variants (e.g., MAX–MIN Ant System, ACS) beyond basic Ant System.

### 5. Practical Takeaway for Researchers

For TSP instances up to ~50 cities, GA can find optimal solutions faster and with simpler implementation. For larger or asymmetric problems (≥ 75 cities), ACO yields higher-quality solutions at the cost of longer runtime and greater sensitivity to stagnation. Neither algorithm is universally superior; the choice should be guided by instance size, computational budget, and tolerance for suboptimality. Future work should include multi-seed statistical analysis, systematic parameter tuning, and hybridization with local search.
