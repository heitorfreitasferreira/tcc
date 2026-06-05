# Efficient optimization of the Held–Karp lower bound

**Author:** Giovanni Righini (University of Milan)  
**Venue:** Open Journal of Mathematical Optimization, Volume 2, Article no. 9, 2021  
**DOI:** [10.5802/ojmo.11](https://doi.org/10.5802/ojmo.11)

---

## 1. Problem and motivation

The Held–Karp lower bound for the TSP is computed by deleting an arbitrary vertex *p*, building a minimum spanning tree (MST) on the remaining *n−1* vertices, and adding the two cheapest edges incident to *p*. Different choices of *p* yield different lower bounds. Righini addresses the problem of **optimally** selecting *p* to maximize this bound — a task whose naive solution costs *O(mn + n² log n)* — and shows it can be done in *O(m + n log n)*, the same complexity as computing a single MST.

## 2. Core method or approach

- Pre-computes all *n* alternative MSTs (one per deleted vertex) in a single pass over the edge list sorted by non-decreasing cost, mimicking Kruskal’s algorithm.
- Uses an oriented spanning tree with depth-first-search numbering (Dn/Up indices) to efficiently determine, for any non-tree edge, which vertices lie on the induced cycle — and therefore which vertices can use that edge as a replacement.
- Maintains a **local subgraph** *G(p)* for each vertex *p*, represented by a Union–Find structure tracking p-components; each inserted link records which alternative edge reconnects which pair of components.
- Introduces a **Tree–Union–Find** structure to merge overlapping oriented paths in *O(log n)* amortized, skipping already-explored segments of cycles in *O(1)*.
- After all alternative edges are found, selects the vertex *p* that maximizes (`AltTreeCost(p) − StarCost(p) + c₁ + c₂`) to obtain the optimal Held–Karp lower bound.

## 3. Main results

- **Complexity:** worst-case *O(m + n log n)* — same as one execution of Kruskal or Prim (with Fibonacci heaps).
- **Experimental validation:** two datasets of 100 instances each (100 vertices); dataset A with 15% random graphs, dataset B with Euclidean complete graphs.
- **Key numbers:** For dataset A, avg. LB_HK (random *p*) = 8990, optimal LB*_HK = 11759 (~23% improvement); optimal LB*_HK was strictly tighter than Helsgaun's LB*_H in 55/100 instances. For dataset B, avg. LB_HK = 530, LB*_HK = 597 (~11% improvement); tighter in 63/100 instances.

## 4. Strengths and limitations

**Strengths:**
- Same asymptotic cost as a single MST computation, enabling practical integration into branch-and-bound solvers.
- Guaranteed to be at least as tight as the Helsgaun leaf-optimized bound.
- Algorithm is more general: it pre-computes alternative MSTs for *any* vertex deletion, useful beyond TSP (e.g., network resilience).

**Limitations:**
- The improvement over Helsgaun's method is instance-dependent — substantial for some, negligible for others.
- The paper provides only limited computational experiments (100-vertex instances, synthetic data); no results on standard TSPLIB benchmarks or larger instances.
- The lower bound is not a TSP solution itself; its practical impact depends on the parent solver (branch-and-bound, etc.), which is not evaluated here.

## 5. Practical takeaway for researchers

For any TSP solver that uses 1-tree lower bounds, replacing an arbitrary vertex choice with this *O(m + n log n)* optimal selection provides a strictly better (and often 10–20% higher) lower bound at no asymptotic extra cost. The algorithm also serves as a general tool for maintaining MST-based structures under vertex deletions.
