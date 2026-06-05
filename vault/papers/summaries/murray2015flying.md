# The Flying Sidekick Traveling Salesman Problem: Optimization of Drone-assisted Parcel Delivery

**Murray, C. C. & Chu, A. G.** (2015). *Transportation Research Part C: Emerging Technologies*, 54, 86–109. DOI: `10.1016/j.trc.2015.03.005`

## 1. Problem and Motivation

While UAV technology for last-mile parcel delivery was gaining commercial traction (Amazon Prime Air, DHL Parcelcopter, Google Project Wing), no prior work addressed the *operational routing and scheduling* challenges of coordinating drones with delivery trucks. The paper introduces two new variants of the traveling salesman problem to model two distinct operating scenarios: (i) a truck–drone tandem where the drone launches from and returns to the truck mid-route, and (ii) a depot-centric scenario where drones and trucks operate in parallel from a distribution center.

## 2. Core Method or Approach

- **FSTSP (Flying Sidekick TSP):** A single UAV works in synchronization with a delivery truck, launching from and recovering to the truck at customer locations. Formulated as a mixed-integer linear program (MILP) minimizing the latest vehicle return time to the depot, with constraints for time coordination, flight endurance, and sortie feasibility.
- **PDSTSP (Parallel Drone Scheduling TSP):** A fleet of identical UAVs operates directly from the depot in parallel with a truck serving a TSP route. The problem decomposes into a TSP (truck customers) and a parallel identical machine scheduling (PMS) problem (drone customers), linked by the customer partition.
- **FSTSP Heuristic:** A route-and-reassign procedure that starts with a TSP tour for the truck, then iteratively removes UAV-eligible customers from the truck route and assigns them to the drone, guided by a savings calculation that accounts for synchronization constraints.
- **PDSTSP Heuristic:** Initializes with all UAV-eligible within-range customers assigned to drones, then iteratively moves customers between drone and truck partitions (and performs pairwise swaps) to balance makespans, calling TSP and PMS solvers at each step.
- **TSP and PMS subproblem solvers:** Evaluated four TSP approaches (IP, Clarke-Wright savings, nearest neighbor, sweep) and two PMS approaches (IP, longest processing time first — LPT).

## 3. Main Results

- **FSTSP heuristic performance (72 instances, 10 customers):** The IP-based TSP subproblem within the heuristic averaged −1.16% gap vs. Gurobi's 30-minute-limited MILP solves (negative gaps = heuristic beat Gurobi). The savings-based TSP heuristic achieved +0.33% average gap in ~0.004 s. Nearest neighbor and sweep were less competitive (+2.91%, +8.33%).
- **PDSTSP heuristic performance (720 instances, 10–20 customers):** IP/IP combination yielded 0.12% avg gap for 10-customer and 0.22% for 20-customer problems. Savings/LPT achieved 1.58% and 3.90% avg gaps, respectively, with sub-millisecond runtimes. LPT delivered near-optimal PMS solutions.
- **Speed vs. endurance (90 instances, 25–75 customers):** Speed dominates endurance — faster UAVs with shorter flight times consistently outperform slower UAVs with longer endurance, even at equivalent total flight distance. The effect is more pronounced with multiple UAVs (PDSTSP) than with the single-UAV FSTSP.
- **FSTSP vs. PDSTSP regime:** PDSTSP is superior when many customers lie within drone range of the depot; FSTSP becomes preferable as the depot becomes remote and fewer customers are in direct range. The crossover depends on depot location, customer distribution, UAV count, and road network constraints.

## 4. Strengths and Limitations

**Strengths:** First formal MILP formulations and heuristic frameworks for two fundamentally distinct drone–truck delivery architectures; heuristics are simple, computationally cheap, and competitive with commercial solvers even on small instances; clear empirical guidance on speed-vs-endurance trade-offs; well-structured problem definitions that spawned extensive follow-up literature.

**Limitations:** FSTSP restricts UAV to a single sortie per launch node, requires rendezvous only at customer locations, and considers only one UAV — all simplifications acknowledged as model tractability constraints; heuristics are greedy with no metaheuristic escape from local optima (simulated annealing or tabu search suggested but not implemented); TSP heuristic choice strongly influences overall heuristic quality; empirical evaluation limited to 10–75 customer instances and synthetic rectangular-grid data with Manhattan truck travel; no cost-based objective; no real-world field validation.

## 5. Practical Takeaway for Researchers

The FSTSP and PDSTSP formulations provide the foundational operational models for drone-assisted last-mile delivery. Researchers extending this work should: (i) relax the rendezvous-at-customer-nodes and single-sortie-per-launch constraints to improve achievable savings; (ii) incorporate metaheuristic escapes (tabu search, SA) into the greedy reassignment framework; (iii) consider multi-UAV, multi-truck extensions; and (iv) when designing drone systems, prioritize flight speed over endurance — the numerical evidence strongly favors faster drones even with reduced range.
