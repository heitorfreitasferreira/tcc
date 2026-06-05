# Particle Swarm Optimization — Kennedy & Eberhart (1995)

## 1. Problem and Motivation

Continuous nonlinear function optimization lacks algorithms that are simultaneously simple, computationally inexpensive, and effective across diverse problem classes. Kennedy and Eberhart propose that social interaction—specifically information sharing among conspecifics—offers an evolutionary advantage for locating resources in complex environments. They hypothesize that a computational model mimicking such social dynamics can serve as a general-purpose optimizer, bridging the gap between slow evolutionary search and fast neural processing.

## 2. Core Method or Approach

- **Particle Swarm Optimizer (PSO):** A population of candidate solutions ("particles") flies through the search space, each adjusting its velocity based on two attractors: its own best-known position (*pbest*, autobiographical memory) and the swarm's global best-known position (*gbest*, social knowledge).
- **Velocity update rule:** `v = v + 2·rand()·(pbest − present) + 2·rand()·(gbest − present)`, where the stochastic coefficient (mean = 1) promotes overshooting the target roughly half the time, balancing exploration and exploitation.
- **Incremental simplification:** The algorithm evolved from a bird-flocking simulation; nearest-neighbor velocity matching, "craziness" (random jitter), and separate p/g increment parameters were progressively stripped away, yielding the minimal three-term velocity update.
- **Computational minimalism:** Requires only primitive arithmetic; implementable in a few lines of code with low memory and speed overhead.
- **Design philosophy:** Aligns with Millonas's five principles of swarm intelligence (proximity, quality, diverse response, stability, adaptability); uses stochastic processes analogous to crossover in genetic algorithms but operates via acceleration toward better solutions rather than direct manipulation of candidates.

## 3. Main Results

- **XOR neural network training:** A 2-3-1 feedforward NN (13 parameters) was trained to `e < 0.05` criterion in an average of 30.7 iterations with 20 particles.
- **Fisher Iris classification:** Trained NN weights as effectively as backpropagation, averaging 284 epochs over 10 training sessions.
- **EEG spike classification:** Achieved 92% test accuracy versus 89% for backpropagation, suggesting better generalization from PSO-trained weights.
- **Schaffer f6 benchmark:** Found the global optimum on every run; approximated the evaluation-budget performance of elementary genetic algorithms reported in Davis (1991).

## 4. Strengths and Limitations

**Strengths:**
- Extreme simplicity with competitive performance on nonlinear, multimodal problems.
- No gradient information required; applicable to non-differentiable objective functions.
- Naturally balances exploration and exploitation via momentum (persistence of velocity) and stochastic attraction to personal/global bests.

**Limitations:**
- Early-stage algorithm; parameter sensitivity (e.g., the constant *2*) not yet systematically analyzed.
- Evaluated on small-dimensional problems (13-D NN weights, 2-D Schaffer f6); scalability to high-dimensional problems unknown at the time.
- The "explorers and settlers" variant and midpoint-only version both failed, indicating that the dual stochastic terms and velocity momentum are fragile if removed.

## 5. Practical Takeaway for Researchers

PSO offers a *default-first-try* optimizer for continuous nonlinear problems where gradient methods are impractical. Its minimal implementation cost and robust performance on multimodal benchmarks make it an attractive baseline. However, practitioners should verify sensitivity to the acceleration constant and population size for their specific problem dimensionality, as these hyperparameters were left under-explored in this foundational paper.
