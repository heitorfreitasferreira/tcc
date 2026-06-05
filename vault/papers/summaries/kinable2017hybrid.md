# kinable2017hybrid — Summary

**Reference:** J. Kinable, A. A. Cire, W.-J. van Hoeve, "Hybrid optimization methods for time-dependent sequencing problems," *European Journal of Operational Research*, Vol. 259, no. 3, pp. 887–897, 2017. DOI: [10.1016/j.ejor.2016.11.035](https://doi.org/10.1016/j.ejor.2016.11.035)

---

## 1. Problem and Motivation

Sequencing problems in manufacturing, transport, and distribution often involve setup times that depend on the *position* of a task in the sequence (e.g., aging/learning effects in machines, or the time-dependent TSP). Position-dependent setup times make the problem substantially harder than the classical position-independent case — exact solutions for the position-dependent TSP were limited to roughly 100 nodes as of 2013, versus 100,000+ for the classical TSP. Most existing approaches are either dedicated to a single problem variant or cannot easily incorporate practical side constraints (time windows, precedence relations).

## 2. Core Method or Approach

- **CP as backbone.** Problem is formulated as a constraint programming model with `alldifferent` and `element` constraints. Traditional CP alone is ineffective because setup-time objectives are decoupled from the feasibility constraints, yielding weak bounds.
- **MDD discrete relaxation.** A multivalued decision diagram (relaxed, width-limited) is embedded as a global constraint (`MDDConstraint`). At each CP search node, the MDD is refined and filtered, producing lower bounds and removing infeasible transitions. The MDD encodes the solution space compactly via layered acyclic graphs; bound computation uses a dynamic-programming recurrence over arcs.
- **LP-based additive bounding.** The linear programming relaxation of a MILP time-space network formulation (Picard–Queyranne / Vander Wiel–Sahinidis) is solved. Residual reduced costs are projected onto MDD arc costs via the additive bounding framework of Fischetti and Toth (1989). This produces bounds at least as strong as the tighter of the MDD or LP bounds individually, and enables additional arc filtering.
- **Problem variants handled.** The same framework accommodates TD-TSP, TD-TSPTW (time windows), and TD-SOP (precedence constraints) with minimal changes — additional constraints (time-window, precedence) are incorporated into the LP and, for TD-SOP, summed into the arc cost recurrence.

## 3. Main Results

- **TD-TSP (43 TSPLib-derived instances, ≤76 nodes):** Both MDD approaches substantially outperformed standalone MIP. The CP+MDD+AdditiveBounding variant (CP^ab_MDD) outperformed CP+MDD for instances up to ~52 vertices; beyond that, LP-weakness or LP cost eroded the benefit. The specialized column-generation solver of Abeledo et al. (2013) generally outperformed the hybrid approach on pure TD-TSP.
- **TD-TSPTW (270 randomly generated instances, 30–40 nodes):** MIP failed to find any feasible solution. CP+MDD+AB found feasible solutions for *all* instances and significantly outperformed CP+MDD alone, especially for larger instances and smaller MDD widths. The AB benefit decreased as MDD width increased (since a wider MDD already captures more structure).
- **TD-SOP (29 TSPLib-derived SOP instances, 7–100 nodes):** Both MDD methods significantly outperformed MIP (MIP found no feasible solution for most instances). The AB improvement over pure CP+MDD was more modest due to weak LP relaxations for the precedence-constrained case; solving the full LP with separation often exceeded 24 hours for 100-vertex instances.
- **Bound analysis:** Combining a weak LP with MDD yielded proportionally larger gains than combining a strong (cut-strengthened) LP with MDD — the MDD compensates structurally when the LP provides limited information.

## 4. Strengths and Limitations

**Strengths:**
- Generic framework: easily incorporates time windows and precedence constraints without problem-specific re-engineering.
- Clear, orders-of-magnitude improvement over pure CP and generic MIP on all three problem classes.
- Synergistic bound improvement: MDD + additive LP bounds are consistently stronger than either alone.
- Novel integration of discrete (MDD) and continuous (LP) relaxations within CP search.

**Limitations:**
- Outperformed by specialized TD-TSP solvers (Abeledo et al. 2013) on pure TD-TSP.
- Performance degrades when the LP relaxation is weak or too expensive (TD-SOP with large instances).
- MDD refinement and LP separation have non-trivial overhead; for small/easy instances, the overhead can outweigh benefits.
- MDD width trades bound strength against runtime; AB benefit diminishes at large widths.

## 5. Practical Takeaway for Researchers

For position-dependent sequencing problems that also involve side constraints (time windows, precedence), a CP+MDD framework with additive LP bounding provides a strong generic baseline that substantially outperforms off-the-shelf MIP/CP while remaining adaptable. The key design insight is that projecting continuous dual information (reduced costs) onto a discrete state-space relaxation (MDD) yields complementary tightening — especially when each individual relaxation is moderately weak. Future work could substitute stronger LP formulations (e.g., Held-Karp, Lagrangian) into the same additive bounding pipeline.
