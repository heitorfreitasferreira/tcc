---
title: "Discrete Particle Swarm Optimization, Illustrated by the Traveling Salesman Problem"
author: "Maurice Clerc"
year: 2000
aliases: [clerc2000discretepso]
tags: [pso, discrete-optimization, tsp, combinatorial-optimization, swarm-intelligence]
---

## Problem and Motivation

Particle Swarm Optimization (PSO) was originally designed for continuous optimization problems. Many real-world problems, including the Traveling Salesman Problem (TSP), are inherently discrete. Clerc aimed to extend PSO to combinatorial domains by redefining the core operators — velocity, position update, and particle movement — in a way that respects discrete search spaces while preserving the swarm's collaborative dynamics.

## Core Method or Approach

- Redefines the particle state as a permutation (tour) and velocity as a list of transpositions (pairwise swaps) rather than continuous vectors.
- Introduces arithmetic operations on the search space: subtraction between two positions yields a velocity (set of swaps), addition of velocity to a position applies those swaps to produce a new position, and scalar multiplication on velocity retains a coefficient-weighted subset of swaps.
- The velocity update equation combines weighted contributions from: current velocity (inertia), the difference between personal best and current position, and the difference between global best and current position — each expressed as a list of transpositions.
- Applies a stochastic proportional coefficient to each transposition source, controlling how many swaps are inherited from each component, analogous to the acceleration coefficients in canonical PSO.
- The resulting algorithm is demonstrated on TSP instances, showing that the discrete PSO converges to near-optimal tours using only the redefined algebraic operators, without problem-specific crossover or mutation.

## Main Results

The method was tested on several classical TSP instances (e.g., eil51, berlin52). The discrete PSO found solutions within a few percent of the known optimum in a modest number of iterations. Convergence behavior showed that the swarm collectively homed in on good permutations, with the best particle's tour length decreasing steadily. The approach demonstrated that PSO could be generalized to permutation problems without losing its essential search dynamics.

## Strengths and Limitations

**Strengths:** Formally clean extension of PSO to discrete permutation spaces via algebraic redefinition; preserves the intuitive velocity-position metaphor; avoids problem-specific genetic operators; applicable to any permutation-based combinatorial problem beyond TSP.

**Limitations:** The transposition-based velocity representation can scale poorly with instance size (velocity lists grow with problem dimension); convergence guarantees are weaker than for continuous PSO; no systematic comparison against state-of-the-art TSP solvers (e.g., Lin-Kernighan heuristic, ant colony optimization) is provided; parameter sensitivity (coefficients, swap rates) remains empirically tuned rather than theoretically characterized.

## Practical Takeaway for Researchers

The paper's main contribution is conceptual: it shows that PSO can be ported to discrete problems by redefining the underlying algebra of positions and velocities. For researchers implementing PSO on combinatorial tasks (scheduling, routing, assignment), the transposition model provides a direct template — but modern discrete PSO variants (e.g., binary PSO, set-based PSO, permutation PSO with priority encoding) generally offer better scalability and should be considered for larger instances.
