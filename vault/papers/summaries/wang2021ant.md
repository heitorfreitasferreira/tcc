# Wang & Han (2021) — Ant Colony Optimization for TSP Based on Parameters Optimization

**Reference:** Wang, Y., & Han, Z. (2021). Ant colony optimization for traveling salesman problem based on parameters optimization. *Applied Soft Computing, 107*, 107439. DOI: [10.1016/j.asoc.2021.107439](https://doi.org/10.1016/j.asoc.2021.107439)

---

## 1. Problem and Motivation

The performance of Ant Colony Optimization (ACO) for the Traveling Salesman Problem (TSP) critically depends on its parameter values — particularly the pheromone importance factor α and the heuristic importance factor β. However, there is no established theory for assigning these parameters; they are problem- and instance-dependent, and finding suitable values typically requires extensive trial-and-error experimentation. Existing approaches that use PSO or GA to tune ACO parameters merely shift the burden, because those metaheuristics themselves have parameters that must be set.

## 2. Core Method: SOS–ACO

- **SOS as parameter optimizer.** Symbiotic Organisms Search (SOS) is used to adaptively optimize ACO's key parameters α and β. SOS is chosen because it is *parameter-free* — it has no tunable parameters of its own, unlike PSO or GA.
- **Three-phase evolution.** SOS evolves a population of organisms (each encoding an [α, β] pair) through:
  - *Mutualism* — two organisms are updated toward the current best, with a mutual benefit factor.
  - *Commensalism* — one organism benefits from the best without affecting the other.
  - *Parasitism* — a mutated organism replaces a host if it has higher fitness.
- **Fitness function.** Fitness(Y) = 1 / (best tour length found by ACO using parameters Y), measured after all m ants construct routes.
- **Local optimization.** A nearest-neighbor + path-reversal local search is applied to the best tour after each ACO run, improving exploitation without significant computational overhead.

## 3. Main Results

- **Dataset:** 10 symmetric TSP instances from TSPLIB, ranging from 51 to 575 cities (Eil51, Eil76, KroA100, Ch150, KroA200, Pr264, Lin318, Rd400, Pr439, Rat575).
- **Solution quality:** Best solutions within 2.33% of BKS across all instances. Exact BKS found for Eil51, Eil76, KroA100, and Pr264. Average error ≤ 2.33% in all cases.
- **Robustness:** PE(%) (average-vs-best relative error) ≤ 1% for all instances, indicating low variance across 10 runs.
- **Comparison to baselines:**
  - Against standard ACO (α=1, β=5): SOS–ACO's Error(%) ≤ 2.33% vs. ACO's 6–20%.
  - Against ACO with local search (ACO–LO): SOS–ACO's Error(%) ≤ 2.33% vs. ACO–LO's 2.4–10.8%.
  - Against published hybrids (PACO-3Opt, HGA, PSO–ACO-3Opt, ACSFA): SOS–ACO finds 8 of the 10 best solutions.
  - For large instances (Lin318, Rd400, Pr439, Rat575), SOS–ACO consistently outperforms all competitors.
- **Parameter sensitivity:** SOS–ACO maintains stable performance across varying ant counts (m=10,20,30), evaporation rates (ρ=0.1–0.7), and pheromone enhancement values (Q=1000–5000). Only excessively high ρ (0.7) causes slight degradation.

## 4. Strengths and Limitations

**Strengths:**
- Parameter-free SOS eliminates the meta-parameter tuning problem that plagues other hybrid ACO approaches.
- Strong empirical performance with consistent convergence across diverse instance scales.
- Demonstrated robustness to variations in both the optimized parameters (α, β) and the fixed ACO parameters (m, ρ, Q).

**Limitations:**
- Only two ACO parameters (α, β) are optimized; others are fixed by the user.
- Tested only up to 575 cities; scalability beyond TSPLIB medium instances not evaluated.
- No theoretical convergence analysis or runtime complexity bounds.
- Local search strategy is relatively simple (nearest-neighbor reversal); more sophisticated operators (e.g., k-Opt) are not incorporated.

## 5. Practical Takeaway for Researchers

Using a parameter-free metaheuristic (SOS) to tune ACO parameters is an effective strategy that reduces the experimental burden of manual parameter tuning. The hybrid approach generalizes well across instance sizes and remains stable even when other fixed ACO parameters vary, making it a practical alternative to manually calibrated ACO for TSP and potentially related combinatorial optimization problems. For future work, combining SOS–ACO with stronger local search operators (e.g., 2-Opt, 3-Opt) is a natural extension.
