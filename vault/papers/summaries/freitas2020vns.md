# A Variable Neighborhood Search for Flying Sidekick Traveling Salesman Problem

**Authors:** Júlia Cária de Freitas, Puca Huachi Vaz Penna (UFOP, 2018)
**arXiv:** 1804.03954v2 [math.OC]

## 1. Problem and Motivation

The Flying Sidekick Traveling Salesman Problem (FSTSP) models parcel delivery where a single truck and a single drone collaborate: the drone launches from and returns to the truck at possibly different customer locations, while the truck services other customers in parallel. The objective is to minimize total delivery time, subject to the drone's limited flight endurance and unitary payload capacity. Rising industry interest (Amazon, DHL, JD.com) motivates efficient algorithms for this NP-hard problem.

## 2. Core Method

- **HGVNS algorithm** — Hybrid heuristic with three stages: (i) Concorde MIP solver finds the optimal TSP tour via the truck cost matrix; (ii) `CreateInitialSolution` converts selected truck customers into drone customers via a savings heuristic; (iii) General Variable Neighborhood Search (GVNS) refines the joint truck-drone route.
- **RVND local search** — Randomized Variable Neighborhood Descent explores 7 neighborhood structures with Best Improvement: Reinsertion, Or-opt2, Exchange, Exchange(2,1), Exchange(2,2), 2-opt, and Relocate Customer. Neighborhoods are shuffled and reinitialized upon improvement.
- **Two problem variants** — Handles both FSTSP (endurance and service time constraints) and TSP-D (unlimited endurance, instant launch/recover) by relaxing constraints.
- **New TSPLIB-based benchmark** — Introduces 25 instances (51–200 nodes) with Manhattan distance for truck and Euclidean distance for drone, 85–90% drone-eligible customers.

## 3. Main Results

| Benchmark | Best Improvement | Avg Improvement (TSP baseline) |
|-----------|-----------------|-------------------------------|
| Ponza (2016) instances (10 runs) | 24.84% over BKS (instance 150.2) | 21.70% over TSP optimum |
| Agatz et al. (2016) TSP-D (uniform/single-center/double-center, α=1,2,3) | 62.24% (single-center, n=75, α=2) | 43–53% (α=1 vs α=2/3) |
| New TSPLIB instances | 45.48% (pr107) | 13.49% |

- New best-known solutions for **all** Ponza (2016) literature instances.
- Drone speed α shows diminishing returns: α=2 and α=3 produce similar improvements (approx. 40–54%) — doubling speed matters more than tripling.
- Overall maximum improvement over TSP-only delivery: **67.79%**.

## 4. Strengths and Limitations

**Strengths:**

- Sets new BKS on every tested instance from the literature.
- Handles two distinct problem formulations (FSTSP and TSP-D).
- Seven diverse, well-diagrammed neighborhood operators tailored for mixed vehicle routes.
- Provides a new reproducible benchmark set grounded in TSPLIB.
- Moderate runtime: avg 23.5s on TSPLIB instances.

**Limitations:**

- Restricted to one truck / one drone; does not scale to heterogeneous fleets.
- No MILP formulation for the FSTSP is provided (left as future work).
- TSPLIB experiments use Manhattan distance for truck (approximates grid roads); real road networks are not considered.
- No direct comparison against Agatz et al. (2016) results due to incomplete experimental reporting in the original paper.
- 10-run average used for stochastic evaluation — modest sample size.

## 5. Practical Takeaway for Researchers

A GVNS initialized from an optimal TSP solution is a strong, reproducible baseline for drone-assisted routing problems. The 7 neighborhood structures (especially Relocate Customer and the Exchange family) provide a reusable toolkit for future metaheuristics on mixed-vehicle routing. When designing drone speed configurations, results suggest that making the drone twice as fast as the truck already captures most of the achievable time savings.
