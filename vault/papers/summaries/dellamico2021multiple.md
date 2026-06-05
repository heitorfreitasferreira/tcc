---
aliases: ["Dell'Amico et al. 2021 - MFSTSP"]
tags: [papers, routing, drones, scheduling, TSP, MILP, branch-and-cut]
paper_doi: "10.1002/net.22022"
year: 2021
---

# Modeling the Flying Sidekick Traveling Salesman Problem with Multiple Drones

**Authors:** Mauro Dell'Amico, Roberto Montemanni, Stefano Novellani  
**Journal:** *Networks* (2021)  
**DOI:** `10.1002/net.22022`

---

## 1. Problem and Motivation

The paper addresses the **Multiple Flying Sidekick Traveling Salesman Problem (MFSTSP)**, where a truck equipped with a fleet of identical drones delivers parcels to minimize total completion time. Each drone performs sorties (launch → serve one customer → rendezvous with the truck) subject to battery endurance limits. When multiple drones launch or return at the same truck stop, the *order* of those operations matters: ignoring the scheduling component can produce solutions that exceed drone endurance in practice. The work is motivated by the rapid growth of e-commerce and same-day delivery demands.

## 2. Core Method or Approach

- **Four MILP formulations:** Two branch-and-cut formulations using four-index sortie variables (`4I-BC`, `3I-BC`) and two using auxiliary "crossing sortie" binary variables (`4I-z`, `3I-z`) that track whether a drone is available on the truck at each node.
- **Explicit scheduling constraints:** Binary variables encode the precedence of launches and rendezvous at each node (`λ`, `ρ`, `α`, `β`), treating the truck as a single-machine scheduling problem at each stop.
- **Branch-and-cut with tournament constraints:** Tournament crossing sortie (TCS) and tournament backward sortie (TBS) inequalities are separated dynamically to cut infeasible fractional solutions; subtour elimination constraints (SEC) are also separated via max-flow.
- **Objective decomposition:** The completion time objective is decomposed into the truck's travel time plus node-level waiting times (`w_i`), improving lower bounds over prior formulations.

## 3. Main Results

- **Benchmark:** Murray & Chu (2015) instances with 10 customers, 2–5 drones, endurance *E* = 20 or 40 min; also tested on Murray & Raj (2020) instances with 8 and 10 customers.
- **Best formulation:** `4I-z` dominates all others: with 2 drones it achieves 2.97% average gap and 49/72 optimal solutions in 1 h on a modest CPU; on an improved setting (8 threads, faster CPU), instances with 2–4 drones are solved to optimality and |D| = 5 reaches only 0.38% gap.
- **Multiple drones benefit:** Using multiple drones reduces cost vs. single-drone; the improvement is larger with higher endurance (E = 40) and faster drones (25–35 mph). Depot location closer to customers amplifies the gain.
- **Comparison with Murray & Raj:** `4I-z` solves 88% of 8-customer instances optimally vs. 66% reported by Murray & Raj; average solution time is ~665 s vs. ~1428 s in the reference.

### Variant Studies

| Variant | Key finding |
|---|---|
| No scheduling of drone operations | Solutions are ~5–20% cheaper but can be infeasible in reality — strongly motivates the scheduling component |
| Launching time applied at depot | Large cost penalty; depot-launched sorties without launch time account for much of the multi-drone advantage |
| Drones allowed to wait on the ground | Small additional improvement; increasing drone count can compensate for lack of ground-wait capability |

## 4. Strengths and Limitations

**Strengths:**
- First multi-drone FSTSP formulation that explicitly schedules launch/rendezvous operations at each node, capturing a previously overlooked source of infeasibility.
- Comprehensive comparison of four formulations with transparent cut statistics and gap analysis.
- Public benchmark reproducibility and comparison against prior state-of-the-art.

**Limitations:**
- Tested only up to 10 customers; problem remains NP-hard and scaling to realistic instance sizes requires heuristics.
- Drones are identical; heterogeneous fleets (different speeds, payloads) are not modeled.
- Drone service time is embedded in travel times; the comparison with Murray & Raj ignores their explicit service time component.
- 1-hour time limit leaves several instances unsolved to optimality, especially for |D| ≥ 4 and E = 40.

## 5. Practical Takeaway for Researchers

The scheduling of multiple drone operations at the truck is **non-negligible** — ignoring it can underestimate completion times significantly. The `4I-z` formulation with crossing sortie variables is the strongest exact approach for small instances. For scaling up, future work should focus on matheuristics or decomposition methods that preserve the scheduling component. When modeling drone delivery problems, the depot launch-time assumption has a first-order effect on solution quality and should be explicitly justified.

---

*Summary generated from full-text PDF reading.*
