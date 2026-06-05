# Routing Problem for Unmanned Aerial Vehicle Patrolling Missions — A Progressive Hedging Algorithm

**Authors:** Sudarshan Rajan, Kaarthik Sundar, Natarajan Gautam  
**Venue:** Preprint (arXiv:2106.08379v2), submitted to Elsevier, July 2021  
**Keywords:** two-stage stochastic program, progressive hedging, integer programming, routing, UAV

---

## 1. Problem and motivation

The paper addresses the Single Vehicle Data Gathering Problem (SVDGP): a UAV must visit a set of targets and, depending on the (uncertain) fidelity of information collected at each target, may need to detour through supplemental locations before proceeding to the next target. The goal is to find a first-stage tour that minimizes the sum of the deterministic tour cost and the expected additional traversal cost incurred by recourse detours. This models real-world patrolling missions — e.g., precision agriculture or surveillance — where data quality at a point of interest determines whether further inspection is required. Despite practical relevance, the stochastic version of this UAV routing problem had not been formulated prior to this work.

## 2. Core method or approach

- **Two-stage stochastic integer program:** First stage selects a TSP tour over targets; second stage decides supplemental-target visits per scenario (binary variables in both stages). Uncertainty is modeled via independent Bernoulli random variables per target.
- **Scenario-based uncertainty:** Each scenario ω is a binary vector indicating which targets require supplemental visits; probabilities can be independent or extended to Markov-chain dependence. The objective minimizes first-stage cost plus expected recourse cost over all scenarios.
- **Progressive Hedging (PH) algorithm:** A scenario-decomposition method that solves one MILP per scenario per iteration, enforces non-anticipativity via quadratic penalty terms, and aggregates solutions through weighted averaging.
- **Cost-proportional ρ heuristic:** The penalty parameter ρ is set by scaling all edge costs to (0,1) and taking their average, following Watson & Woodruff (2011).
- **Dubins-path travel costs:** Edge costs are computed as shortest Dubins paths for a fixed-wing UAV with minimum turn-radius (5 units), yielding asymmetric travel costs.

## 3. Main results

- **Optimality and speed:** On 720 test instances with 10 targets, the PH algorithm solved all instances to optimality, whereas solving the extensive form via CPLEX timed out (7200 s limit) on 21 of 240 instances with 50 scenarios and 42 of 240 with 200 scenarios. PH was on average **2.5× faster** than the extensive form.
- **Value of the stochastic solution (VSS):** Relative improvement over the expected-value problem reached 5–17.5%, increasing with higher probability of insufficient fidelity (p ≥ 0.5).
- **Scalability:** PH solved instances up to 40 targets and 200 scenarios within ~4.5 hours (longest: 40 targets/200 scenarios at ~16,315 s). Computation time grows steeply beyond 25–30 targets.
- **Supplemental locations and radius:** No clear monotonic trend between m (3–9) or R (5–10 units) and computation time; algorithm converged in 1–37 iterations across all experiments.
- **Residual convergence:** PH residuals decreased to < 10⁻⁵, typically within 1–25 iterations depending on instance size and scenario count.

## 4. Strengths and limitations

**Strengths:**
- First two-stage stochastic formulation for UAV data-gathering with binary recourse, filling a gap in the literature.
- PH algorithm provides an effective decomposition heuristic even with binary variables in both stages; empirically converged to optimal solutions on all tested instances.
- General framework: supports asymmetric costs, arbitrary probability distributions (via sampling), and is extensible to multi-stage settings.
- Rigorous experimental design with 56 distinct random instances tested across multiple parameter dimensions.

**Limitations:**
- Single-vehicle only; multi-UAV extension requires additional target-partitioning logic.
- No theoretical convergence guarantee for PH with discrete variables (heuristic only, though empirically effective).
- Computation time becomes prohibitive for >40 targets; the authors suggest clustering heuristics for practical missions with ~100 targets.
- Only two-stage formulation evaluated; multi-stage extension described but not implemented or tested.
- All experiments use synthetic random instances on a 100×100 grid; no real-world UAV mission data.

## 5. Practical takeaway for researchers

The PH algorithm is a robust, parallelizable heuristic for two-stage stochastic routing problems with binary recourse — it outperforms solving the extensive form directly and yields empirically optimal solutions. For UAV patrolling applications with up to ~30–40 targets, it provides tractable offline route plans that account for data-fidelity uncertainty. Researchers extending this work should focus on multi-vehicle decomposition, multi-stage rolling-horizon implementations, and real-world validation with actual UAV flight data.
