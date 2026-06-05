# Ahmed & Sheltami (2024) — Receding Horizon and Optimization-based Control for UAV Path Planning with Collision Avoidance

**Venue:** Procedia Computer Science 251 (2024), pp. 15–22  
**DOI:** 10.1016/j.procs.2024.11.079  
**Institution:** King Fahd University of Petroleum and Minerals, Saudi Arabia

---

## 1. Problem and Motivation

UAVs operating in obstacle-dense environments require path planning strategies that simultaneously minimize energy consumption and guarantee collision avoidance. Existing global planning approaches struggle with real-time constraints when the environment is only partially known or dynamically updated. The authors target an **online** energy-efficient path planner that re-plans within a limited detection range, making it suitable for reactive navigation in cluttered 3D spaces.

## 2. Core Method

- **Receding Horizon Control (RHC):** The full path is decomposed into overlapping spatial segments (finite horizons). Only the current segment is optimized; the horizon shifts forward as the UAV progresses.
- **MILP formulation:** Within each segment, the path is cast as a mixed-integer linear program. Decision variables encode waypoint visitation (`x_ij`) and intermediate location indicators (`y_i`). A collision penalty (`OC_ij`) is set to infinity for obstacle-crossing edges.
- **Energy-minimization objective:** The cost function minimizes total flight energy (`E_ij`) along the path. The energy model is adopted from Di Franco & Buttazzo (2016), which incorporates horizontal flight, vertical flight, hovering, and turning costs from real UAV experiments.
- **CPLEX solver:** Each segment's MILP is solved exactly using IBM CPLEX, generating the locally optimal collision-free waypoints.
- **Quadratic Bézier smoothing:** Linear waypoint-to-waypoint segments are smoothed to reduce energy wasted on sharp deceleration/acceleration turns.

## 3. Main Results

- **Simulation environment:** MATLAB, 100 m × 100 m area, UAV altitude 50 m, segment size 20 m, altitude bounds [30 m, 120 m]. Tested with 9 or 15 static obstacles and 15 dynamic obstacles.
- **Path quality:** The approach successfully generated collision-free paths in all tested scenarios, including narrow-passage configurations.
- **Energy metric:** Normalized energy consumption (total consumed energy divided by the energy of the shortest straight-line path at maximum speed).
- **Smoothing impact:** Across 0–8 obstacles, the Bézier-smoothed paths consistently consumed less energy than unsmoothed paths. The percentage improvement from smoothing is plotted and reported as notable (specific values are shown in figures but not enumerated in the text).

## 4. Strengths and Limitations

**Strengths:**
- Online-capable: the RHC decomposition bounds the MILP size, keeping solve times manageable.
- Exact solutions per horizon (CPLEX), avoiding approximation errors of heuristic planners.
- Energy model grounded in empirical UAV flight data.
- Smoothing step directly addresses a real-world source of UAV energy waste (sharp turns).

**Limitations:**
- Evaluation is simulation-only (MATLAB); no hardware-in-the-loop or field experiments.
- Environment size is modest (100 × 100 m) and the segment/horizon size (20 m) is fixed.
- Relies on a commercial solver (CPLEX), which may limit reproducibility in resource-constrained settings.
- Short conference paper; no comparison against meta-heuristic or sampling-based planners beyond the smoothed-vs-unsmoothed ablation.
- Dynamic obstacles are mentioned but no quantitative results on success rate under varying obstacle speeds or densities are provided.

## 5. Practical Takeaway for Researchers

The RHC + MILP pipeline offers a principled way to achieve **exact local optimality** under hard collision constraints, which is valuable when safety guarantees matter more than asymptotic scalability. For TSP/rTSP drone patrol contexts, the segment-decomposition idea could be adapted to **multi-goal visitation**: each horizon might contain a subset of target points, with the global tour constructed incrementally. The smoothing strategy is directly transferable to any waypoint-based planner and provides a low-cost energy improvement.
