# Cluster Based Branching for the Asymmetric Traveling Salesman Problem

**Reference**: Lysgaard, J. (1999). Cluster based branching for the asymmetric traveling salesman problem. *European Journal of Operational Research*, 119(3), 314–325.

---

## 1. Problem and Motivation

Existing branch-and-bound algorithms for the ATSP used the assignment problem (AP) relaxation for lower bounds and subtour elimination schemes for branching (Bellmore & Malone, 1971; Garfinkel, 1973). These branching schemes were efficient to compute but ignored additional problem-structure information gathered during the bounding procedures. The paper addresses this gap by exploiting dual information from the additive bounding phase to guide branching decisions.

## 2. Core Method or Approach

- Defines a **cluster** as a node set $K \subset N$ ($2 \leq |K| \leq n-2$) for which there exists an optimal ATSP solution satisfying $x(K, N\setminus K) = 1$ — i.e., the nodes of $K$ are visited consecutively.
- **Cluster identification** proceeds in two stages:
  - **Arborescence-based**: cutsets from the additive bounding procedure (Fischetti & Toth, 1992) with dual variable $U_k \geq \text{GUB} - \text{LB}$ are promoted directly to clusters.
  - **Artificial nodes-based** (SEAN procedure): when arborescence identification fails, candidate cutsets satisfying $U_k \geq F \cdot (\text{GUB} - \text{LB})$ are tested by introducing two artificial nodes; if the lower bound on the expanded instance reaches GUB, the node set is confirmed as a cluster.
- **Cluster branching**: once a cluster $S$ is identified, dominance tests reduce the set of candidate entry/exit nodes, then branching is performed by fixing exit nodes (exit branching) or entry nodes (entry branching), whichever yields fewer subproblems. The cluster is contracted into a single node in each child subproblem.
- **Subtour reduction tests** (fallback when no cluster is found): inversion-based reduction (branch 2) and dynamic programming-based reduction (branches $k \ge 3$) prune the search tree by detecting dominated continuation paths in the chosen subtour.
- The full algorithm is a best-bound branch-and-bound that embeds the Fischetti & Toth (1992) additive bounding steps (AP → shortest paths → SAAP → SADP → SAADP → cost tightening) and replaces standard subtour branching with the cluster-based scheme in steps 8–10.

## 3. Main Results

- **Dataset**: 22 ATSP instances from TSPLIB (Reinelt, 1991), ranging from 33 to 161 nodes (FT53, FT70, RY48P, KRO124P, P43X2, and the FTV series).
- **Key performance comparison** (cluster branching vs. subtour branching on a DEC Alpha 2100 5/250):
  - Cluster branching was consistently faster on 21 of 22 instances (only FTV33 was slightly slower, time ratio 1.12).
  - Dramatic outlier: **P43X2** — unsolved after 1 CPU hour with subtour branching; solved in **7 seconds** with cluster branching (time ratio < 0.002).
  - Time ratios (cluster/subtour) ranged from **0.08 to 0.96** on solved instances; BB node count was consistently lower with cluster branching.
  - Cluster identification was successful at **10%–30%** of branch-and-bound nodes on most instances.
- Four largest instances (RBG323, RBG358, RBG403, RBG443) were solved at the root node (AP bound = heuristic upper bound). FTV170 could not be solved due to memory limits; FTV160 timed out (>3330 s) even with cluster branching.

## 4. Strengths and Limitations

**Strengths**:
- Exploits dual information from bounding procedures to derive branching decisions, unlike purely subtour-based schemes.
- Cluster contraction reduces problem size in child nodes; dominance tests aggressively prune the set of candidate branches.
- Integrates cleanly with the widely adopted Fischetti–Toth additive bounding framework.

**Limitations**:
- Cluster size is capped at ≈8–16 nodes due to the exponential time/space cost of the Held–Karp DP algorithm for computing shortest Hamiltonian paths within a cluster.
- The artificial-nodes cluster test adds non-trivial computational overhead when arborescence identification fails.
- Large instances (FTV160, FTV170) remain out of reach due to CPU time or memory.
- The approach is tied to the branching structure of Fig. 1, requiring that the subproblem $x(S, N\setminus S) \ge 2$ be immediately fathomable — which holds for network-based lower bounds but not necessarily for LP-based bounds.

## 5. Practical Takeaway for Researchers

Cluster branching demonstrates that incorporating structural information from dual variables into branching can yield order-of-magnitude improvements over standard subtour elimination, especially on ATSP instances with symmetric sub-structures (e.g., identical nodes, as in P43X2). The principle of identifying consecutively-visitable node sets and contracting them is transferable to other sequencing/routing problems where the lower-bounding infrastructure produces meaningful duals on cutsets.
