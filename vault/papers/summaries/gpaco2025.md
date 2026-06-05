# GP-ACO 2025 — Automated design of state transition rules in ACO by genetic programming

**Reference:** Lin, B.-C., Mei, Y., & Zhang, M. (2025). Automated design of state transition rules in ant colony optimization by genetic programming: a comprehensive investigation. *Memetic Computing*, 17:2. DOI: [10.1007/s12293-025-00435-9](https://doi.org/10.1007/s12293-025-00435-9)

---

## 1. Problem and motivation

Manual crafting of ACO state transition rules demands deep domain expertise and becomes intractable for complex combinatorial problems where subtle variable interactions escape human designers. Existing automated alternatives — neural-network predictors requiring large pre-solved datasets and LLM-based design with unknown autonomy — suffer from data dependency, poor interpretability, or inadequate systematic evaluation. The paper addresses these gaps by conducting the first comprehensive experimental study of GP-ACO, investigating generality, sensitivity to ACO variants, local-search interactions, global-information augmentation, and interpretability.

## 2. Core method or approach

- **GP-ACO framework:** Genetic Programming evolves tree-structured state transition rules that replace handcrafted formulas (e.g., $\phi^\alpha \eta^\beta$) within the ACO solution construction. Each GP individual is compiled into a state transition rule, integrated into an ACO variant (AS, ACS, or MMAS), and evaluated by running a full ACO solve on a training TSP instance; the tour length determines fitness.
- **Matrix-based acceleration:** A PyTorch-accelerated matrix formulation restructures scalar edge-by-edge priority calculations into batched matrix operations, leveraging multi-core CPUs to reduce training time by approximately 80% versus the scalar baseline.
- **xGP-ACO extension:** Expands the terminal set from the original two ($\phi$ = pheromone, $d$ = distance) to six by adding four global-information terminals: average distance $\bar{d}$, average pheromone $\bar{\phi}$, total nodes $n$, and number of candidate nodes $n_u$. The function set comprises $+$, $-$, $\times$, protected division, and negation.
- **Experimental protocol:** Four progressive experiments on TSP rue instances (TSP20, TSP20-100, TSP100) and 15 TSPLIB instances, with 30 independent GP training runs per configuration, each tested 30 times. Comparisons include ACO baselines and LLM-designed ACO (ReEvo).
- **Hyper-heuristic character:** GP-ACO is a simulation-based generative hyper-heuristic: it does not require pre-solved data and evaluates rules through full ACO execution, offering superior generality over supervised-learning alternatives.

## 3. Main results

- **Generality (Q1):** GP-ACO consistently outperforms ACO baselines and LLM-ACO across all three TSP scales; the performance advantage widens with problem size. On TSP100: GP-AS mean normalized cost = 1.0105 vs. AS baseline 1.0784.
- **ACO variant independence (Q2):** The choice of ACO framework (AS, ACS, MMAS) does not degrade GP-ACO learning capability. GP-AS and GP-ACS notably surpass LLM-ACO; GP-MMAS is competitive but slightly behind LLM-MMAS.
- **Local search effect (Q3):** 2-opt local search weakens GP-ACO learning in MMAS (which depends less on state transition rules) but not in AS or ACS. GP-ACS+2-opt and GP-AS+2-opt still significantly outperform ACO+2-opt baselines.
- **Global information boost (Q4):** xGP-ACO (six terminals) without local search significantly outperforms GP-ACO, ACO baselines, and LLM-ACO; xGP-MMAS finds solutions below existing lower bounds. With 2-opt, xGP-AS+2-opt improves further, but xGP-MMAS+2-opt weakens due to reduced dependency on learned rules.
- **Cross-distribution TSPLIB (Q5):** Rules trained on TSP100 transfer effectively to diverse TSPLIB instances. xGP-AS-best reduces the AS baseline gap from 13.51% to 5.91%; with 2-opt, xGP-AS+2-opt-best achieves 0.70% gap to optimum and xGP-MMAS+2-opt-best achieves 0.25%.
- **Interpretability (Q6):** Evolved GP trees reveal structured patterns — e.g., high powers of $d$ in denominators, negative $\phi d$ terms in numerators, and selective use of global terminals $\bar{d}$ and $\bar{\phi}$ — supporting human-readable analysis of learned rules.

## 4. Strengths and limitations

**Strengths:**
- Does not require pre-solved training data or expert domain knowledge.
- Robust generality across problem scales and distributions (rue → TSPLIB).
- Excellent interpretability via tree-structured rule expressions.
- Systematic ablation design covering ACO variants, local search, and terminal sets.
- Matrix acceleration makes nested GP-ACO training practically feasible.

**Limitations:**
- Training remains expensive: ACO is nested inside GP, even with ~80% acceleration.
- xGP-MMAS+2-opt exhibits overfitting on larger out-of-distribution instances (gaps > 5% for problems larger than training scale).
- Experiments restricted to TSP; no evaluation on VRP, scheduling, or other CO problems.
- Terminal $n_u$ was unused in all best-evolved rules; the benefit of six terminals may saturate.
- Only 50 GP generations were used; deeper evolution or larger trees may yield further gains but increase cost.

## 5. Practical takeaway for researchers

GP-ACO, especially the xGP variant with global-information terminals, offers an effective, interpretable, and data-free method for automatically designing ACO state transition rules. Researchers working on metaheuristic-based optimization should consider GP-ACO as a hyper-heuristic alternative whenever manual rule design is costly or problem characteristics shift between instances. Augmenting the terminal set with average-distance and average-pheromone features provides the largest performance jump at minimal implementation cost. When local search is available, prefer pairing xGP with AS or ACS rather than MMAS to preserve learning effectiveness.
