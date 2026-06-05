# Agatz, Bouman, Schmidt (2016) — Optimization Approaches for the Traveling Salesman Problem with Drone

## 1. Problem and Motivation
Last-mile parcel delivery is costly and slow. A truck-drone tandem — where a drone launches from a truck to serve nearby customers while the truck continues its route — can parallelize deliveries and reduce total service time. This gives rise to the **Traveling Salesman Problem with Drone (TSP-D)**, a new NP-hard variant combining assignment and routing decisions.

## 2. Core Method or Approach
- Integer Programming (IP) formulation that solves small instances to optimality.
- **Route-first, cluster-second** heuristics: first construct a TSP truck tour (via Concorde optimal or MST-based), then partition nodes into truck/drone via dynamic programming.
- Four heuristic variants: greedy partitioning (gp) vs exact partitioning (ep), each with optional local search improvement (swap/relocate), evaluated on uniform- and cluster-distributed Euclidean instances.
- Worst-case approximation bounds: TSP to TSP-D is a (2+α)-approximation using the MST heuristic; TSP-only is a (1+α)-approximation, where α is the max speed-up factor of drone over truck.

## 3. Main Results
- On small instances (≤10 nodes), exact partitioning heuristics (TSP-ep) reach ~8% gap to optimal for uniform instances and ~24% for clustered instances.
- For larger instances (100 nodes), TSP-ep-all achieves ~30% savings over truck-only TSP in uniform settings; greedy variants (TSP-gp) yield ~20% savings.
- Greedy heuristics solve 100-node instances in seconds; exact partitioning run time grows quickly with instance size.
- The ~30% savings are consistent across instance sizes for uniform customer distributions.

## 4. Strengths and Limitations
- **Strengths**: First paper to compare heuristic TSP-D solutions against exact TSP-D optima (not just TSP bounds); provides both theoretical approximation guarantees and empirical validation; IP formulation supports extensions (endurance, restricted nodes).
- **Limitations**: Single truck + single drone only; no partial recharging modeled; exact partitioning heuristic does not scale well beyond ~50 nodes.

## 5. Practical Takeaway
For uniform customer distributions, truck-drone collaboration can reduce total service time by ~30% over truck-only delivery. Route-first, cluster-second heuristics with dynamic programming partitioning offer a strong cost/quality trade-off for instances up to ~100 nodes.
