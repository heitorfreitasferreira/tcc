# NeuFACO: Neural Focused Ant Colony Optimization for Traveling Salesman Problem

**arXiv:2509.16938v2** (2025) — Dat Thanh Tran, Khai Quang Tran, Khoi Anh Pham, Van Khu Vu, Dong Duc Do (VinUniversity / HUST / VNU Hanoi)

## 1. Problem and Motivation

The Traveling Salesman Problem (TSP) is NP-hard; exact solvers scale exponentially and heuristics risk local optima without problem-specific knowledge. Prior non-autoregressive (NAR) hybrid methods that combine neural networks with Ant Colony Optimization suffer from high variance, low sample efficiency (REINFORCE-based), and weak local refinement. NeuFACO addresses these issues by replacing REINFORCE with stable PPO training and upgrading the ACO engine with focused refinement and bounded pheromone updates.

## 2. Core Method or Approach

- **PPO-based policy learning**: A graph neural network (GNN) is trained with clipped PPO + entropy regularization to produce instance-conditioned heuristic matrices `H ∈ R^(n×n)` and value estimates, without pheromone during training.
- **Min–Max Ant System (MMAS)**: Pheromone trails are clamped to `[τ_min, τ_max]` and updated only from the best-so-far or iteration-best solution, preventing the neural prior from prematurely dominating search.
- **Focused ACO (FACO) with node relocation**: Instead of reconstructing full tours, each ant copies a reference solution (global-best or iteration-best) and modifies it by relocating a small number of nodes (minimum new edges = 8), preserving strong substructures.
- **Hierarchical node selection**: Candidate lists (k-nearest neighbors, k=20) and precomputed backup lists (BKP=64) reduce node-selection complexity; a full O(n) scan is used only as a fallback.
- **Scalable 2-opt local search**: Applied at every iteration, restricted to candidate edges and edges that changed since the last optimization, making it computationally feasible without prohibitive overhead.

## 3. Main Results

**NAR baselines (M=100 ants, I=100 iterations)** — all models retrained under identical hyperparameters:

| Instance   | Method   | Gap (%) | Time    |
|------------|----------|---------|---------|
| TSP200     | DeepACO  | 1.89    | 16.86s  |
| TSP200     | GFACS    | 1.80    | 17.36s  |
| TSP200     | NeuFACO  | 1.66    | 0.28s   |
| TSP500     | DeepACO  | 3.03    | 119.58s |
| TSP500     | GFACS    | 1.54    | 136.14s |
| TSP500     | NeuFACO  | 1.50    | 0.74s   |
| TSP1000    | DeepACO  | 4.32    | 659.43s |
| TSP1000    | GFACS    | 2.48    | 730.56s |
| TSP1000    | NeuFACO  | 2.00    | 1.98s   |

**RL baselines (M=256 ants, I=1000 iterations)**:

| Instance   | Method    | Obj.  | Gap (%) | Time   |
|------------|-----------|-------|---------|--------|
| TSP500     | Concorde  | 16.55 | –       | 10.7s  |
| TSP500     | NeuFACO   | 16.43 | 1.33    | 0.91s  |
| TSP500     | SIT+Greedy| 16.70 | 1.08    | 14.95s |
| TSP1000    | Concorde  | 23.12 | –       | 108s   |
| TSP1000    | NeuFACO   | 23.34 | 1.16    | 0.98s  |
| TSP1000    | LEHD+Greedy|23.84 | 3.11    | 0.8s   |

Evaluation on TSPLib instances up to 1500 nodes and randomized TSP (128 instances per size). NeuFACO achieves up to 60× lower wall-clock time than prior NAR methods and consistently lower variance.

## 4. Strengths and Limitations

**Strengths**: PPO provides stable, sample-efficient training compared to REINFORCE; FACO + MMAS preserves useful substructures and prevents premature neural dominance; very fast amortized inference; competitive with or outperforms a broad range of neural TSP solvers; scales to 1500 nodes.

**Limitations**: Solution sampling remains CPU-bound, making total runtime longer than some recent methods (e.g., SIT, LEHD on TSP1000); evaluated only on Euclidean TSP — generalization to other COPs not shown; no ablation study quantifying contribution of individual ACO components (MMAS, FACO, 2-opt) in isolation.

## 5. Practical Takeaway for Researchers

Combining PPO-trained neural heuristics with FACO (targeted node relocation) and MMAS (bounded pheromone) produces a strong, stable, and fast hybrid solver for Euclidean TSP. The hierarchical node selection (candidate → backup → fallback) and restricted per-iteration 2-opt are practical efficiency tricks that could transfer to other metaheuristic + neural hybrids. The CPU-bound sampling bottleneck suggests future work should focus on accelerating solution construction, possibly via GPU-accelerated ACO or further search-space pruning.
