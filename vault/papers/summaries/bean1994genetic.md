# Bean (1994) — Genetic Algorithms and Random Keys for Sequencing and Optimization

**Reference:** James C. Bean. Genetic Algorithms and Random Keys for Sequencing and Optimization. *ORSA Journal on Computing*, 6(2):154–160, 1994. DOI: [10.1287/ijoc.6.2.154](https://doi.org/10.1287/ijoc.6.2.154)

---

## 1. Problem and Motivation

Traditional genetic algorithms struggle with sequencing and permutation problems because standard crossover of two feasible parent permutations almost always produces infeasible offspring (e.g., duplicate or missing elements in a TSP tour). Prior work addressed this with problem-specific repair operators (PMX crossover, edge recombination, subsequence-swap, etc.), but no general, robust encoding existed that guarantees offspring feasibility across diverse problem classes without per-problem customization.

## 2. Core Method: Random Keys

- Solutions are encoded as vectors of random numbers in [0,1]^n (one gene per decision variable).
- Decoding: sort the random keys in ascending order; the resulting index permutation defines the solution (e.g., job sequence, machine assignment, variable priority).
- For multi-resource problems, the integer part of a key encodes resource assignment (e.g., machine ID) and the fractional part encodes sequencing priority.
- All genetic operators (crossover, mutation) operate on the real-valued random-key vectors — never on the discrete permutation — so *any* vector decodes to a feasible solution.
- The GA implementation uses: **elitist reproduction** (copy best individuals), **parametrized uniform crossover** (0.7 bias toward first parent), and **immigration** (1% of population replaced by new random individuals each generation) instead of traditional bit-flip mutation.

## 3. Main Results

| Problem Class | Instances | Target | Performance |
|---|---|---|---|
| Multiple-machine scheduling (min. total tardiness) | fshx4 (200 jobs, 4 m/c), fshx8 (400 jobs, 8 m/c) | 2.5% of best known | ~5 min (fshx4), ~25 min (fshx8) on IBM RS/6000-320H |
| Resource allocation (0-1 knapsack) | ra1, ra2 (100 vars, 5 constraints); ra1d, ra2d (200 vars, 5 constraints, dual-degenerate) | 2.5% of optimum | ~1.5–10 min; on dual-degenerate problems, GA solved in ~500 s vs OSL branch-and-bound taking >14 h (ra2d) |
| Quadratic assignment (Nugent) | n15, n20, n30 | 5% of best known | n15: <1 min; n20: ~2 min; n30: 2/10 runs failed to reach target in 30,000 generations |

- Low variance across 10 random seeds for scheduling and resource allocation.
- QAP results are weaker than specialized methods like simulated annealing or tabu search.

## 4. Strengths and Limitations

**Strengths:**
- Single encoding scheme applies to scheduling, resource allocation, QAP, vehicle routing, and generalized TSP without problem-specific repair.
- Guarantees offspring feasibility by construction — feasibility logic moves into the decoder/evaluator.
- Empirically scalable and robust across random seeds on scheduling and knapsack problems.

**Limitations:**
- Performance on QAP is only moderate; specialized heuristics (SA, tabu search) outperform it.
- No theoretical analysis of convergence rate; justification is purely empirical.
- Population size and immigration rate were tuned by pilot studies, not systematically parameterized.
- The decoder mapping is still problem-specific (sorting logic varies by problem), so the approach is not a "black box."

## 5. Practical Takeaway

Random keys provide a simple, drop-in encoding strategy that eliminates the infeasible-offspring problem for any sequencing or assignment problem where a valid solution can be derived by sorting priorities. It trades specialization for generality: researchers who need a quick, robust GA for a new permutation problem can adopt random keys with the described elitist+immigration GA and expect decent results, especially for scheduling and resource allocation problems, though it may not beat highly tuned problem-specific heuristics on harder combinatorial problems like QAP.
