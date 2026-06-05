---
title: "Exact models for the flying sidekick traveling salesman problem"
authors: "Mauro Dell'Amico, Roberto Montemanni, Stefano Novellani"
year: 2022
journal: "International Transactions in Operational Research"
doi: "10.1111/itor.13030"
tags:
  - fstsp
  - exact-methods
  - milp
  - drone-routing
  - tsp-d
---

## 1. Problem and motivation

The Flying Sidekick Traveling Salesman Problem (FSTSP) models a truck--drone delivery system where both vehicles cooperate to serve customers. The drone can depart from the truck, serve exactly one customer per flight (sortie), and return to the truck at a different node, all within a finite battery endurance. Synchronization between truck and drone must be maintained, and the objective is to minimize the overall completion time. Prior MILP formulations relied heavily on "big-M" constraints and could not solve all 10-customer benchmark instances to proven optimality.

## 2. Core method or approach

- Three novel MILP formulations are proposed: **3IF** (3-indexed sortie variables), **2IF** (2-indexed sortie variables with a drone-position binary variable z_i), and **2IF-BC** (2IF without z, using branch-and-cut separation of exponentially many crossing-sortie elimination constraints).
- A single set of time variables t_i replaces the two separate sets (truck + drone) used in prior formulations, halving the number of timing variables and reducing "big-M" constraint counts from O(n³) to O(n²).
- The binary variable z_i tracks whether the drone is on the truck at node i, which eliminates crossing sorties without path-based enumeration for 3IF and 2IF; 2IF-BC separates tournament constraints dynamically.
- The objective function is decomposed into truck travel time + launch/rendezvous times + waiting time, which improves the root-node lower bound gap from ~81.5% (Murray and Chu, 2015) to ~22.5%.
- Extensions are provided for variants: loops (same-node launch/rendezvous), unlimited endurance, and drone waiting on the ground.

## 3. Main results

- **10-customer benchmarks (Murray & Chu, 2015):** All 72 instances (E=20 and E=40) solved to optimality. 14 of these had no previously known optimal solution. 2IF is the fastest overall (avg. ~38 s for E=20, ~233 s for E=40).
- **20-customer instances:** With E=20, 2IF-BC solved 58/120 optimally (avg. gap 5.6%); 2IF solved 52/120 (avg. gap 6.4%). With E=40, only 2/120 reached optimality (gap ~17--19%).
- **Loop variant (vs. Poikonen et al., 2019):** Both 2IF and 2IF-BC solved all 9-customer instances in ~4 s and all 14-customer instances in ~10--15 min, while the prior B&B solved none of the 14-customer instances.
- **Comparison with CF-RR (Roberti & Ruthmair, 2020):** 2IF/2IF-BC are faster on the FSTSP with loops and on the endurance-constrained FSTSP; CF-RR is superior for the unlimited-endurance variant.
- Depot location strongly affects difficulty: instances with the depot near the center of gravity are hardest due to more feasible sorties.

## 4. Strengths and limitations

**Strengths:**

- Dramatically more compact than prior formulations (fewer big-M constraints, fewer variables).
- First exact method to close all 10-customer FSTSP benchmarks.
- Competitive or superior to other compact formulations on loop variants and endurance-constrained problems.
- Modular: same formulation family accommodates loops, unlimited endurance, and wait-on-ground variants.

**Limitations:**

- Scalability plateaus at ~20 customers, especially with larger endurance (E=40: only 2/120 solved optimally).
- Branch-and-price methods outperform on unlimited-endurance and larger instances.
- Experiments run on modest hardware (Intel Core i3-2100, 3.1 GHz, 8 GB RAM); absolute times vs. competitors on faster machines must be interpreted cautiously.
- Single truck/single drone only; no multiple-vehicle extensions.

## 5. Practical takeaway for researchers

The 2IF and 2IF-BC formulations provide the most effective compact MILP approaches for the endurance-constrained FSTSP and its loop variant at small-to-medium scale (≤20 customers). For problems where endurance is binding and the drone must hover while waiting, these formulations should be preferred over CF-RR and similar compact models. For unlimited endurance or larger instances (>20 customers), branch-and-price or decomposition-based methods remain the better choice. The objective decomposition trick (separating waiting time from travel time) is transferable to other synchronized vehicle routing models and yields significantly tighter lower bounds.
