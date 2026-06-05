# pop2024comprehensive

**Pop, P. C., Cosma, O., Sabo, C., & Pop Sitar, C.** (2024). *A comprehensive survey on the generalized traveling salesman problem.* European Journal of Operational Research, 314, 819–835.

---

## 1. Problem and motivation

The generalized traveling salesman problem (GTSP) extends the classical TSP by partitioning vertices into clusters and requiring that exactly one vertex from each cluster is visited on a minimum-cost Hamiltonian cycle. Despite being among the most researched combinatorial optimization problems — with applications in location-routing, material flow design, medical supply distribution, waste collection, airport routing, and image retrieval — no prior dedicated survey existed for the GTSP. This paper fills that gap.

## 2. Core method or approach

- **Survey structure.** The paper is organized around six axes: formal problem definition and variants, real-life applications, mathematical formulations (both exponential-size and compact), solution approaches (exact, transformation, reduction, approximation, heuristic/metaheuristic), benchmark datasets, and comparative computational analysis.
- **Formulation taxonomy.** Integer and mixed-integer programming formulations are reviewed, including the generalized subtour elimination formulation, the generalized cutset formulation, single-commodity flow formulations, and the local-global decomposition approach that separates macro-level (inter-cluster) and micro-level (intra-cluster) decisions.
- **Solution approach classification.** Methods are divided into five classes: exact algorithms (branch-and-bound, branch-and-cut, dynamic programming), transformation methods (reducing GTSP to TSP), reduction algorithms (eliminating non-optimal vertices/edges), approximation algorithms (with quality guarantees), and heuristic/metaheuristic algorithms (GA, ACO, PSO, ILS, LNS, beam search).
- **Comparative benchmark.** Five state-of-the-art algorithms — memetic algorithm (MA), Lin-Kernighan-Helsgaun (LKH), large neighborhood search (LNS), basic iterated local search (Basic ILS), and refined ILS — are compared across four benchmark libraries (GTSP_LIB, BAF_LIB, MOM_LIB, LARGE_LIB) using success rate and average percentage error.
- **Future directions.** The paper identifies six open research avenues: theoretical analysis of formulation strength, decomposition and hybrid methods, dynamic GTSP variants, uncertainty modeling (stochastic/fuzzy/rough), multigraph-based formulations, and updating benchmark instances.

## 3. Main results

- **Datasets.** Four standard benchmark libraries totaling 233 instances: GTSP_LIB (88 instances, 31–217 clusters, up to 1084 vertices), BAF_LIB (56 instances, up to 1084 vertices), MOM_LIB (45 instances, 2–200 clusters, up to 3000 vertices), and LARGE_LIB (44 instances, 10–17,180 clusters, up to 85,900 vertices).
- **Best overall solver.** The Large Neighborhood Search (LNS) algorithm by Smith & Imeson (2017) achieved the best overall performance: most best-known solutions (BKS) found across all libraries, 100% median success rate on MOM_LIB, and lowest average percentage errors on BAF_LIB (0.07%) and MOM_LIB (0.02%).
- **Efficiency.** On directly comparable GTSP_LIB instances, LKH is the most efficient (lower average percentage error in shorter time), followed by MA and then LNS. LKH holds the best success rates on GTSP_LIB (0.0065% average error) and LARGE_LIB (0.52% average error).
- **Algorithm landscape.** Genetic algorithms and ant colony optimization are the most widely used metaheuristics historically; recent trends favor hybrid algorithms and neighborhood search techniques (ILS, LNS, beam search).

## 4. Strengths and limitations

- **Strengths.** The first comprehensive GTSP survey; systematic taxonomy of formulations and solution approaches; covers 50+ years of literature; provides the only cross-library comparative analysis of the five leading algorithms with processed, normalized results; identifies concrete open problems.
- **Limitations.** Efficiency comparison is incomplete due to lack of standardized execution-time reporting across papers, different hardware, and different programming languages; no unifying theoretical comparison of the LP relaxation strength of competing formulations; the benchmark instances have not been revised since their introduction and may not reflect the scale of modern applications.

## 5. Practical takeaway for researchers

- For new GTSP work, benchmark against all four libraries (not just GTSP_LIB) and use LNS (Smith & Imeson, 2017) as the primary state-of-the-art baseline.
- When computational efficiency matters, LKH (Helsgaun, 2015) is the best choice for symmetric instances.
- Decomposition methods (local-global approach) and hybrid exact/metaheuristic combinations are the most promising directions for scaling to larger, more realistic instances.
- Dynamic and uncertain variants of the GTSP remain largely unexplored and represent a significant research opportunity.
