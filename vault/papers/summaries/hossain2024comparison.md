# Hossain & Yilmaz Acar (2024) — Comparison of New and Old Optimization Algorithms for Traveling Salesman Problem on Small, Medium, and Large-scale Benchmark Instances

**DOI:** 10.17798/bitlisfen.1380086  
**Journal:** Bitlis Eren University Journal of Science, 13(1), 216–231

---

## 1. Problem and Motivation

The Traveling Salesman Problem (TSP) is a canonical NP-hard combinatorial optimization problem with broad real-world applications. While many metaheuristics have been proposed over decades, there is no clear consensus on how recently developed algorithms (ABC, SSA, GWO) compare against classical ones (GA, ACO, SA) across different problem scales. The authors aim to fill this gap through a controlled empirical comparison on small, medium, and large TSP benchmark instances.

## 2. Core Method or Approach

- **Algorithms compared:** Six metaheuristics grouped as _old_ (Genetic Algorithm — GA, Ant Colony Optimization — ACO, Simulated Annealing — SA) and _new_ (Artificial Bee Colony — ABC, Salp Swarm Algorithm — SSA, Grey Wolf Optimization — GWO).
- **Benchmark instances:** Seven TSPLIB instances — burma14, berlin52, kroA100 (small, 14–100 nodes); ts225, att532 (medium, 225–532 nodes); rat783, dsj1000 (large, 783–1000 nodes).
- **Experimental setup:** All algorithms run with population size = 100, iterations = 1000, 10 independent runs. Implementation in Python (Anaconda) using Matplotlib, Seaborn, SciPy, and Pandas.
- **Evaluation metrics:** Best optimal solution, average solution, standard deviation, standard deviation (%), and computation time.
- **Statistical analysis:** Two-sample t-tests at α = 0.05 comparing convergence and efficiency across algorithm pairs and between old vs. new groups.

## 3. Main Results

- **Medium-scale instances:** The clearest finding — new algorithms (ABC, SSA, GWO) significantly outperform old algorithms in solution quality, convergence, and computation time on ts225 and att532. Among old algorithms, ACO achieves the best optimal solutions for medium-scale instances (133,285 for ts225; 99,268 for att532).
- **Small-scale instances:** No statistically significant differences between any algorithm pairs. All algorithms converge to near-optimal solutions reliably.
- **Large-scale instances:** No statistically significant differences between old and new groups. GA achieves the best optimal solutions (170,626 for rat783; 534,398,427 for dsj1000), but GWO also shows competitive performance with lower computation time.
- **GWO characteristics:** Fastest algorithm among the new group; demonstrates low standard deviation percentages (0.94%–7.53%), indicating stable convergence across runs; scales efficiently to large instances.
- **ACO characteristics:** Excels at medium-scale instances due to pheromone-based exploration/exploitation balance; degrades on large instances with very high computation times (up to 43,420s for dsj1000).
- **SA characteristics:** Fastest computation times overall (sub-second for all instances) but solution quality degrades sharply as instance size grows.

## 4. Strengths and Limitations

**Strengths:**
- Side-by-side comparison of six algorithms across three distinct scale categories, controlled with standardized parameters.
- Statistical significance testing (t-tests) avoids over-interpreting noisy differences.
- Reports multiple metrics beyond final objective value (standard deviation %, computation time), enabling nuanced algorithmic assessment.

**Limitations:**
- Only seven TSPLIB instances used; findings may not generalize to other instance families (symmetric, asymmetric, real-world).
- Single set of algorithm parameters (no hyperparameter tuning per instance/algorithm), which may favor some algorithms over others.
- No hybrid or ensemble methods evaluated, despite hybrids being common in practice.
- Implementation in Python rather than compiled languages — computation time comparisons may be affected by language-level overhead.

## 5. Practical Takeaway for Researchers

- **Scale-dependent algorithm selection matters.** For medium-scale TSP instances (~200–500 nodes), prefer newer metaheuristics (ABC, SSA, GWO) or ACO over GA and SA. For large-scale instances (700+ nodes), GA and GWO are more competitive. For small instances (<100 nodes), choice is indifferent.
- **GWO is a strong general-purpose choice** with good solution quality, low variance, and competitive runtime across all scales.
- Future work should investigate hybrid approaches and algorithm-specific hyperparameter tuning to close the gap on small and large instances where no method clearly dominates.
