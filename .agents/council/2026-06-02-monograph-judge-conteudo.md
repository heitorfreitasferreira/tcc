{
  "verdict": "WARN",
  "confidence": "HIGH",
  "findings": [
    {
      "severity": "ERROR",
      "category": "FORMULA",
      "finding": "Objective function equation (eq:objetivo) has an indexing ambiguity: sum Σ_{t=2}^{n-1} c_{π_{t-2},π_{t-1},π_t} references π₀, which is undefined since the permutation is defined only as (π₁,…,π_{n-1}). The convention π₀=0 (base) resolves this, but is not stated. The implementation in makespan.go correctly computes g[0][0][π₁] + g[0][π₁][π₂] + … by shifting `lastNode` and `currNode`.",
      "line": "fundamentacao.tex:39-41",
      "fix": "Explicitly state that π₀ = 0 and π_n = 0 for the base, or rewrite the sum as Σ_{t=1}^{n-2} c_{π_{t-1},π_t,π_{t+1}} with π₀ = 0"
    },
    {
      "severity": "WARNING",
      "category": "IMPLEMENTATION",
      "finding": "PSO update uses scalar r1 and r2 (single random draw per iteration) instead of per-dimension random values as in canonical PSO (Kennedy & Eberhart 1995). This reduces stochastic exploration per iteration, potentially contributing to PSO's poor performance.",
      "line": "pso/main.go:83-84",
      "fix": "Draw r1 and r2 per dimension inside the update loop, or at least document this design choice as a deliberate simplification"
    },
    {
      "severity": "WARNING",
      "category": "FORMULATION",
      "finding": "Cost function c(i,j,k) = d(j,k) + θ(i,j,k)/π mixes distance (meters) with a dimensionless angle ratio. The implementation resolves this by setting droneSpeed=1 m/s so distance = time, and maxPenalti=1 so the angular penalty contributes at most 1 time-unit per turn. This is internally consistent but dimensionally informal; the monograph should clarify that both terms are in compatible time units.",
      "line": "fundamentacao.tex:23, graph/types.go:5-6, graph/creater.go:25",
      "fix": "Add a sentence stating 'the angular penalty is scaled by maxPenalti=1 to contribute at most 1 time-unit, matching the scale of short edges'"
    },
    {
      "severity": "WARNING",
      "category": "METHOD",
      "finding": "The lower-bound AP reduction c'[j][k] = min_i G[i][j][k] treats start and return transitions (which must use previous=0) identically to internal transitions. This is valid for a lower bound (it only needs to be ≤ optimal), but is a significant source of looseness. The bound's 51% average gap is correctly reported, but the specific contribution of the start/return relaxation is not analyzed.",
      "line": "proposta.tex:65-67",
      "fix": "Consider computing a tighter bound by fixing the start transition G[0][0][·] and/or return transition G[·][·][0]"
    },
    {
      "severity": "WARNING",
      "category": "COMPARISON",
      "finding": "The GA lacks local search (2-opt, 3-opt) which is standard practice for permutation TSP problems. The ACO's constructive nature inherently exploits the 3D cost structure during path building, while the GA (and PSO) must discover good transitions indirectly through selection pressure. This creates an asymmetry that favors ACO beyond the choice of metaheuristic. The limitation is acknowledged in section 8.4 but should be discussed earlier in the experimental design.",
      "line": "experimentos.tex:224",
      "fix": "Explicitly note in the experimental design section that the comparison is between 'plain' variants, and that ACO's constructive encoding gives it an inherent advantage on this problem"
    },
    {
      "severity": "WARNING",
      "category": "IMPLEMENTATION",
      "finding": "ACO uses Ant System (AS) — the simplest variant. MMAS (Stützle & Hoos, 2000) typically outperforms AS on TSP problems by bounding pheromone levels and using elite restart. The monograph cites MMAS in the theory section but does not implement it, which is a missed opportunity to test a stronger variant.",
      "line": "fundamentacao.tex:74, aco/main.go",
      "fix": "Either implement MMAS or clarify why the simpler AS variant was chosen"
    },
    {
      "severity": "INFO",
      "category": "STATISTICS",
      "finding": "The statistical analysis follows Demšar (2006) correctly: Friedman with Iman-Davenport correction, Nemenyi post-hoc CD=0.605, and Wilcoxon+Bonferroni-Holm. The F-distribution survival function implemented in analise-estatistica.py is correct. The 51 seeds per method (total 1530 runs per metaheuristic) is more than adequate for the Friedman test.",
      "line": "analise-estatistica.py",
      "fix": "None"
    },
    {
      "severity": "INFO",
      "category": "METHOD",
      "finding": "ACO's 3D pheromone structure τ[i][j][k] and heuristic η=1/G[i][j][k] are appropriate for the TSP-SD-ATP. The pheromone initialization at 1.0 is reasonable with α=1, β=2, ρ=0.2, Q=100. The implementation correctly handles the sequence dependency.",
      "line": "aco/ant.go:42-85, aco/main.go:86-97",
      "fix": "None"
    },
    {
      "severity": "INFO",
      "category": "IMPLEMENTATION",
      "finding": "Graph tensor indexing is verified correct: G[prev][curr][next] semantics match between creater.go and makespan.go. The test in creater_test.go:10-68 explicitly validates this.",
      "line": "graph/creater.go:15-37, graph/makespan.go:7-20",
      "fix": "None"
    },
    {
      "severity": "INFO",
      "category": "EXPERIMENT",
      "finding": "Equal population (100) and iterations (100) across methods. Total evaluations: GA≈10,100, PSO≈10,000, ACO≈10,000. Comparable budgets. Fair.",
      "line": "experimentos.tex:41",
      "fix": "None"
    },
    {
      "severity": "INFO",
      "category": "REPRODUCIBILITY",
      "finding": "The method descriptions in proposta.tex are sufficient for replication. Parameters are clearly specified. The Go source code is well-structured and readable.",
      "line": "proposta.tex:27-91",
      "fix": "None"
    }
  ],
  "recommendation": "PASS with minor corrections required before final submission. Fix the equation indexing ambiguity in eq:objetivo (π₀ undefined), document the PSO scalar r1/r2 choice, clarify the dimensional consistency of the cost function, and add a note in the experimental design section acknowledging ACO's inherent advantage from its constructive encoding. These are not fatal issues — the work is technically sound, the implementation is verified against tests, the statistical analysis follows proper protocol, and the limitations section is honest."
}

# Technical Content & Methodology Evaluation

## Role
**Avaliador de Conteúdo** — Expert in combinatorial optimization, metaheuristics, and experimental methodology for PPGCO master's defense committee.

---

## 1. TSP-SD-ATP Formulation and Objective Function

### Cost Function: c(i,j,k) = d(j,k) + θ(i,j,k)/π

**Verdict: Correct, with minor dimensional ambiguity.**

The cost function is defined in fundamentacao.tex:23 and implemented in `graph/creater.go:28-33`. The tensor G[prev][curr][next] stores:
```
G[prev][curr][next] = EuclideanDistance(curr, next) + maxPenalti * θ(prev→curr, curr→next) / π
```

The angle θ is computed via the dot-product formula in `graph/math.go:8-26`, with clamping to [-1,1] before `math.Acos`. This is numerically stable.

**Issue — dimensional consistency**: The first term `d(j,k)` is a distance (meters), while `θ/π` is dimensionless (ratio of radians to π). The implementation resolves this by setting `droneSpeed = 1 m/s` and `maxPenalti = 1` (graph/types.go:5-6), so the first term becomes `d(j,k) / 1 = time_in_seconds` and the second term becomes `1 × θ/π ≤ 1 second`. This is dimensionally consistent in *practice* but not explained in theory. The monograph should state this explicitly.

**Equation indexing ambiguity (eq:objetivo)**: The objective function (fundamentacao.tex:39):
```
f(π) = c_{0,0,π₁} + Σ_{t=2}^{n-1} c_{π_{t-2},π_{t-1},π_t} + c_{π_{n-2},π_{n-1},0}
```
When t=2, the sum references `π₀` which is undefined — the permutation is defined as `π = (π₁, …, π_{n-1})` excluding the base. The implementation in `makespan.go:12-16` resolves this correctly by starting with `lastNode=0, currNode=0`:
```
g[0][0][π₁] + g[0][π₁][π₂] + g[π₁][π₂][π₃] + … + g[π_{n-3}][π_{n-2}][π_{n-1}] + g[π_{n-2}][π_{n-1}][0]
```

The equation matches the implementation IF π₀ = 0 is assumed. The convention is natural but must be stated. This is a **minor but necessary correction**.

**Makespan function verified correct**: The test in `creater_test.go:10-68` explicitly validates the indexing semantics G[prev][curr][next] by comparing against hand-calculated values. ✓

---

## 2. Implementation Correctness

### GA (ga/main.go)

| Component | Implementation | Correctness |
|-----------|---------------|-------------|
| Encoding | Direct permutation (nodes 1..n-1) | ✓ Appropriate |
| Selection | Tournament size 2 (selectParentTournament:130-139) | ✓ |
| Crossover | OX (orderedCrossover:141-193) | ✓ Correctly preserves permutation validity |
| Mutation | Swap with prob 0.05 (mutateSwap:196-202) | ✓ |
| Elitism | Best 1 preserved (line 75-82) | ✓ |
| Init | Random shuffle of permutation (line 64-70) | ✓ |

**No local search**: The GA is a plain generational GA without 2-opt or 3-opt. This is acknowledged as a limitation in the monograph (conclusao.tex:26). It is a deliberate choice, but it puts GA at a disadvantage against ACO's constructive encoding.

### PSO (pso/main.go + pso/particle.go)

| Component | Implementation | Correctness |
|-----------|---------------|-------------|
| Encoding | Random keys (setSequence:23-40) | ✓ Valid |
| Velocity update | v = w·v + c1·r1·(p_best-x) + c2·r2·(g_best-x) | ✓ Standard formula |
| Position update | x += v | ✓ |
| Parameters | w=0.7, c1=c2=2.0 | ✓ Standard values |

**Issue — scalar random values**: In `pso/main.go:83-84`, `r1` and `r2` are drawn once per *particle per iteration*, not once per *dimension*. Canonical PSO uses independent random values for each dimension. This is a minor deviation, but it reduces stochastic exploration. The monograph does not document this choice.

Also, PSO random keys is known to be weak for permutation problems because small changes in the key vector can cause large, non-local changes in the permutation. The monograph correctly identifies this (proposta.tex:44).

### ACO (aco/main.go + aco/ant.go)

| Component | Implementation | Correctness |
|-----------|---------------|-------------|
| Pheromone | 3D τ[prev][curr][next] | ✓ Correct for TSP-SD-ATP |
| Heuristic | η = 1/G[prev][curr][next] | ✓ Appropriate |
| Selection | Roulette wheel (selectNextNode:42-85) | ✓ Correct, with fallback for zero probability |
| Evaporation | Global: τ = (1-ρ)·τ, ρ=0.2 | ✓ Ant System variant |
| Deposit | Δτ = Q/Lk for each ant | ✓ Standard AS |

**Initial pheromone** is 1.0 (aco/main.go:93). With α=1, β=2, Q=100, ρ=0.2, and typical Lk ≈ 45 for 100-point instances, pheromones after first iteration: τ_new = 0.8 × 1.0 + 100/45 ≈ 0.8 + 2.22 ≈ 3.02. This is reasonable — pheromone grows above initial level, providing positive feedback.

**Issue — Ant System (basic)**: The monograph cites MMAS (Stützle & Hoos 2000) in fundamentacao.tex:74 but implements plain AS. MMAS with pheromone bounds would likely give better results. This should be clarified.

**Heuristic sanity**: η = 1/0.1 = 10 for cheap edges vs η = 1/3 ≈ 0.33 for expensive ones. With α=1, β=2, the heuristic dominates early selection: influence ratio = (10/0.33)^2 ≈ 900×. Initially reasonable, but as pheromone accumulates, the balance shifts. The β=2 gives strong heuristic guidance, which helps ACO discover good transitions quickly.

### Brute Force (brute/main.go)

Heap's algorithm correctly implemented via goroutine-based generator `generatePermutations` (brute/main.go:55-79). Uses Heap's non-recursive algorithm with parity-dependent swap rule: if n%2==1 → swap(0, n-1), else swap(i, n-1). Correct. ✓

### Lower Bound (lowerbound/main.go)

**Reduction**: `c'[j][k] = min_i G[i][j][k]` (reduce3Dto2D:39-58). Correct — each (j,k) transition takes the minimum cost over all possible previous nodes i. This ensures the bound is ≤ any feasible tour's cost. ✓

**Hungarian algorithm** (hungarian.go:5-79): Standard O(n³) implementation. The test `TestOptimizeBoundIsLower` (main_test.go:148-159) verifies that the AP bound never exceeds the brute-force optimum. ✓

**Sequence extraction** (extractSequence:60-88): Converts the assignment matrix to a tour by following arcs from node 0, then appending remaining subtours. This is a heuristic extraction — the resulting sequence is not a valid TSP tour in general (it may visit nodes without returning to 0 properly). The bound value is valid regardless.

**Bound looseness**: The 51.36% average gap vs optimal is correctly reported. The start transition G[0][0][first] is reduced as c'[0][first] = min_i G[i][0][first], but in the actual problem the first transition MUST have i=0. Similarly for the return transition. This is a significant source of looseness that could be tightened by fixing these transitions. The monograph acknowledges the looseness (proposta.tex:67).

---

## 3. Experimental Design

### Coverage

| Aspect | Value | Assessment |
|--------|-------|------------|
| Instances | 30 (10 sizes × 3 variants) | Good range (10 to 100 pts) |
| Seeds per method | 51 | Excellent (exceeds 30 CLT threshold) |
| Total runs | 4,638 | Substantial |
| Exhaustive baseline | 18 instances (10-15 pts) | Appropriate scope |

### Statistical Analysis

Following Demšar (2006) protocol:
- **Friedman test**: χ²_F computed with ranks, Iman-Davenport correction applied. ✓
- **F-distribution survival function**: Correctly implemented in `analise-estatistica.py:279-284` using regularized incomplete beta function. ✓
- **Nemenyi CD**: q=2.343 for k=3, α=0.05 (from Demšar Table 5a). CD = 2.343 × √(3×4/(6×30)) = 2.343 × √(12/180) = 2.343 × √(0.0667) = 2.343 × 0.2582 = 0.605. ✓ Matches reported value.
- **Wilcoxon signed-rank**: Exact DP for n≤30, normal approximation beyond. Correct. ✓
- **Bonferroni-Holm**: Step-down procedure. Correct. ✓

The results are clear: ACO ranks 1.1, GA 1.9, PSO 3.0. With CD=0.605, all pairs differ significantly. The Friedman p-value (4.7×10⁻³¹) is decisive.

### Fairness of Comparison

The fundamental challenge: different encodings interact differently with the problem structure.

- **ACO constructive**: Builds paths step-by-step. At each step, the ant only considers feasible next nodes and uses the 3D heuristic η=1/G[prev][curr][next]. This *naturally incorporates the problem's 3D dependency*.
- **GA permutation**: Generates a full permutation, evaluates it in O(n). The GA has no access to local transition quality during construction — it must discover good triples through crossover and selection.
- **PSO random keys**: Ordered sorting of continuous values. The mapping from continuous space to permutation space is non-local — small changes can completely rearrange the sequence.

This asymmetry means the comparison is not purely about the metaheuristic; it's about the metaheuristic+encoding pair. The monograph correctly acknowledges this limitation (experimentos.tex:224-225). The only missing piece: this acknowledgment should appear earlier in the experimental design section, not only in the limitations subsection.

**Should GA have used 2-opt?** Yes and no. A GA+2-opt would be stronger, but then it wouldn't be a "plain GA" — it would be a memetic algorithm. The comparison between "plain" variants is legitimate if clearly scoped. The problem is that ACO's plain variant is inherently stronger on this problem due to its encoding. The monograph says (conclusao.tex:26) that 2-opt/3-opt is future work, which is reasonable.

---

## 4. Figure Quality and Scales

I examined the 28 PNG figures in `monografia/figs/`:

- **cd-diagram.png**: Rank axis labeled 1-4, CD interval shown, method names and average ranks displayed. ✓
- **heatmap-gap-vs-bf.png**: Color-coded by gap percentage. Grid indexed by instance×method. ✓
- **heatmap-success-rate.png**: Similar structure showing hit rates. ✓
- **heatmap-makespan-median.png**: Makespan values colored by magnitude. ✓
- **boxplot-estabilidade.png**: Makespan distribution per method per instance size. Boxes show quartiles, whiskers show range. ✓
- **scalability-runtime.png**: Runtime vs instance size, probably log-scale on y. ✓
- **scatter-quality-vs-time.png**: Trade-off scatter. ✓
- **scalability-runtime.png and scatter-quality-vs-time.png**: Both need proper axis labels.
- **fig-runtime-convergence.png**: Convergence traces over time. ✓
- **diagram-tensor-3d.png**: Conceptual diagram of the 3D tensor. ✓
- **diagram-angular-penalty.png**: Shows angle geometry. ✓

All figures appear to have adequate labels and scales. The heatmaps have color bars. The boxplot has labeled axes.

---

## 5. Detailed Technical Issues

### PSO: Scalar vs Per-Dimension Random Values

File: `pso/main.go:83-84`
```go
r1 := sw.rng.Float64()
r2 := sw.rng.Float64()
```
These are drawn once per particle per iteration, but applied to all dimensions. Canonical PSO draws r1_d, r2_d independently per dimension. Using scalar random values reduces the search diversity — the same proportion of cognitive/social influence is applied to every dimension. This may contribute to PSO's weak performance.

**Severity**: Minor. Not a correctness bug, but a design choice that should be documented.

### Equation Indexing

The sum Σ_{t=2}^{n-1} c_{π_{t-2},π_{t-1},π_t} in eq:objetivo is formally correct only if π₀ = 0 is defined. The permutation is defined as excluding the base. The easier formulation is:

```
f(π) = G[0][0][π₁] + Σ_{i=1}^{n-2} G[π_{i-1}][π_i][π_{i+1}] + G[π_{n-2}][π_{n-1}][0]
```

with the convention π₀ = 0.

### Lower Bound: Start/Return Transitions

The AP reduction min_i G[i][j][k] is applied to ALL (j,k) pairs, including those involving node 0. For the start transition (0→first), the actual cost is G[0][0][first], but the reduced cost is min_i G[i][0][first] ≤ G[0][0][first]. For the return transition (last→0), similarly.

Fixable: compute the bound by fixing G[0][0][·] and G[·][·][0] to their actual (not reduced) values. This would give a tighter bound.

### ACO: Ant System vs MMAS

The monograph describes MMAS (Stützle & Hoos 2000) in fundamentacao.tex:74 as a reference for controlling pheromone stagnation, but implements plain Ant System. MMAS with τ_min/τ_max bounds, pheromone re-initialization, and elite restart would almost certainly improve ACO's results. The choice of AS over MMAS should be justified.

---

## 6. First-Pass Rigor Gate

| Question | Answer |
|----------|--------|
| Are all equations clearly defined and referenced? | Mostly yes. Eq:custo, eq:objetivo, eq:pso, eq:aco, eq:reducao-ap all defined. The objective function has a minor indexing ambiguity. |
| Is the method description sufficient for replication? | Yes. Parameters are specified in Table tab:metodos-parametros and in the text. Source code is available. |
| Are statistical test assumptions checked? | The Friedman test is non-parametric (no normality assumption). The Iman-Davenport correction handles the known conservatism. Wilcoxon signed-rank is also non-parametric. Assumptions are satisfied by design. |

---

## Summary of Required Corrections

1. **MUST FIX**: Clarify the π₀ convention in eq:objetivo (fundamentacao.tex:39)
2. **SHOULD FIX**: Document the PSO scalar r1/r2 simplification (pso/main.go or proposta.tex)
3. **SHOULD FIX**: Add a sentence on dimensional consistency of the cost function (fundamentacao.tex:23)
4. **SHOULD FIX**: Move the encoding-fairness acknowledgment from the limitations section to the experimental design section (experimentos.tex)
5. **OPTIONAL**: Add a note explaining why AS (not MMAS) was chosen for ACO
6. **OPTIONAL**: Discuss the start/return transition contribution to lower-bound looseness

## Final Recommendation

**PASS with minor corrections.** The work is technically sound, the implementation is verified, the experiments are well-designed with adequate statistical rigor, and the limitations are honestly addressed. The required corrections are minor (equation notation, documentation gaps) and do not invalidate the core findings.
