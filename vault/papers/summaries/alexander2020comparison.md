---
title: "The Comparison of Genetic Algorithm and Ant Colony Optimization in Completing Travelling Salesman Problem"
authors: "Alexander, Haris Sriwindono"
year: 2020
venue: "ICSTI 2019 (EAI)"
doi: "10.4108/eai.20-9-2019.2292121"
tags:
  - comparison
  - ga
  - aco
  - tsp
  - metaheuristic
---

## 1. Problem and Motivation

The Traveling Salesman Problem (TSP) is an NP-hard combinatorial optimization problem whose exact solution scales exponentially with the number of cities. Heuristic methods such as Genetic Algorithms (GA) and Ant Colony Optimization (ACO) offer approximate solutions by trading optimality for speed. This paper compares GA and ACO on the TSP to determine which yields shorter tour distances and which is computationally more efficient.

## 2. Core Method or Approach

- **GA configuration:** Roulette-wheel parent selection, Order Crossover (OX), Reciprocal-exchange mutation (1:1000 mutation-to-crossover ratio), population size = 10 chromosomes.
- **ACO configuration:** Ant Colony System (ACS) variant with pheromone weight α = 1, visibility weight β = 0.5, evaporation rate ρ = 0.9.
- **Instances:** 10 datasets with 10, 20, 30, …, 100 randomly generated city points; each tested 10 times with results averaged.
- **Metrics:** Total tour distance and computational time (seconds).
- All parameters held constant across runs; no parameter tuning reported.

## 3. Main Results

- **Solution quality:** ACS consistently produces shorter distances than GA across all instance sizes. The quality gap widens with problem size — GA/ACS distance ratio grows from approximately 2× at 10 cities to approximately 4× at 90–100 cities.
- **Computational time:** GA is slower than ACS at 10–20 cities but becomes faster than ACS for all instances above 30 cities. Time results exhibit fluctuations across instance sizes.
- Both algorithms produce distances that scale directly with the number of cities, confirming the NP-hard nature of TSP.

## 4. Strengths and Limitations

**Strengths:**
- Clear side-by-side comparison of two widely used metaheuristics on the same benchmark instances.
- Covers a range of instance sizes (10–100 cities), providing a scaling perspective.
- Reports both solution quality and runtime — a practical two-dimensional comparison.

**Limitations:**
- Small and fixed GA population (10 chromosomes) likely handicaps GA; no attempt to tune or vary parameters for either algorithm.
- No statistical significance testing across the 10 runs — only averages are reported.
- No comparison against known optimal solutions or optimality gaps (makes it hard to assess absolute quality).
- Limited to two metrics; convergence behavior, robustness (variance), and scalability beyond 100 cities are unexplored.
- GA operators (OX, reciprocal mutation) and ACO variant (ACS) may not represent the state of the art for either family.

## 5. Practical Takeaway for Researchers

ACO (specifically ACS) finds shorter TSP tours at the cost of higher runtime on instances above ~20 cities, while GA converges faster but with substantially worse solution quality under the tested configuration. When both solution quality and speed matter, the trade-off depends on the application. Researchers should not interpret these results as a definitive ranking of GA vs. ACO — parameter tuning, larger populations, and modern operator variants can significantly alter the outcome.
