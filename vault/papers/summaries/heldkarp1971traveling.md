# The traveling-salesman problem and minimum spanning trees: Part II

**Authors:** Michael Held, Richard M. Karp  
**Venue:** Mathematical Programming, vol. 1, pp. 6–25 (1971)  
**DOI:** [to be verified]  

---

## 1. Problem and motivation

The paper addresses the exact solution of the symmetric traveling-salesman problem (TSP). The authors’ prior work (Held & Karp, 1970) showed that the minimum-weight 1-tree yields a sharp lower bound for TSP, but the methods to compute that bound were insufficient. This paper provides an efficient iterative ascent method to approximate the bound and integrates it into a branch-and-bound solver that produces proven-optimal tours for instances up to 64 cities with remarkably small search trees.

## 2. Core method or approach

- **1-tree lower bound.** A 1-tree generalizes a spanning tree on vertices {2,…,n} plus two edges incident to vertex 1. Minimizing its weight under edge-cost perturbations π gives a family of lower bounds `w(π)`, and the tightest bound is `max_π w(π)`.
- **Ascent method (relaxation-style iteration).** The vector π is updated iteratively via `π_{m+1} = π_m + t_m · v`, where `v` is the degree-deficit vector of the current minimum-weight 1-tree. The update follows the normal of the most-violated inequality in the dual linear program, reducing Euclidean distance to the optimum.
- **Constant-step heuristic.** In practice a fixed step-size `t` is used. Theorem 1 guarantees that the achieved bound converges to within `½t·lim sup‖v‖²` of the true maximum. Empirically the method produces 1-trees that closely resemble tours (many degree-2 vertices).
- **Branch-and-bound integration.** When the ascent stalls for `p` consecutive iterations, branching excludes/includes edges ranked by how much the bound would rise if each were forbidden. Derived subproblems inherit the same ascent machinery.
- **Heuristic upper bound.** An upper bound `C̄` (from any constructive heuristic) prunes branches whose lower bound already exceeds `C̄`.

## 3. Main results

The solver was tested on 20 problems (20–64 cities) including random, Euclidean, knight’s-tour, and structured instances, running on an IBM 360/91.

| Category | Example |
|---|---|
| **Solved without branching** (initial ascent bound = optimum) | Croes-20 (4 s), Random-30 (19 s), Slant-64 (182 s) |
| **Solved with branching** | Dantzig-42 (0.9 min), Held-Karp-48 (1.4 min), Karg-Thompson-57 (13 min), Random-Euclidean-64 (5.5 min), 8×8 Knight’s tour (6.9 min) |

Key evidence:
- The initial-ascent bound was extremely close to the optimum in all cases; most of the residual gap is structural (`C* − max w(π)`), not algorithmic.
- The search trees are minuscule — e.g., only 44 interior nodes for Random-1500 (64 cities), 31 for 8×8 knight’s tour.
- Several full search trees are published (the authors claim this is unprecedented for large combinatorial problems).

## 4. Strengths and limitations

**Strengths:**
- The 1-tree relaxation provides an extraordinarily tight lower bound for symmetric TSP.
- The combined ascent + branch-and-bound routine solved, to proven optimality, instances that were considered large in 1971 (up to 64 cities), with tiny search trees.
- Formal convergence analysis (lemmas, Theorem 1) links the iteration to the relaxation method for linear inequalities.

**Limitations:**
- Constant step-size `t` is crude; variable step-size strategies are only sketched as future work.
- Fixed parameter `p` for termination of the ascent wastes computation.
- Computing time remains high (up to 15 min) despite the small search trees; the O(n²) minimum 1-tree recomputation per iteration is the bottleneck.
- No theoretical characterization of the duality gap (`C* − max w(π)`) is given.

## 5. Practical takeaway for researchers

The Held–Karp bound is a foundational TSP lower bound still relevant today. For heuristics or metaheuristics that need a quality certificate, computing `max w(π)` via the ascent method provides a cheap, tight lower bound. In exact solvers, the bound serves as the backbone of branch-and-cut frameworks (e.g., Concorde). The paper also demonstrates that deep structural relaxations (1-tree) can dramatically outperform naive bounding in combinatorial optimization.
