---
title: "An exhaustive review of the metaheuristic algorithms for search and optimization: taxonomy, applications, and open challenges"
authors: "Kanchan Rajwar, Kusum Deep, Swagatam Das"
year: 2023
journal: "Artificial Intelligence Review"
doi: "10.1007/s10462-023-10470-y"
tags: [metaheuristics, survey, taxonomy, parameter-classification, nature-inspired, optimization]
---

## 1. Problem and motivation

With over 500 metaheuristic algorithms (MAs) developed to date and roughly 38 new ones appearing each year in the last decade, the field of metaheuristics is in a state of uncontrolled proliferation. Many recently proposed algorithms show substantial similarities to existing methods under different names, raising the question of what constitutes genuine novelty. This survey aims to provide a comprehensive, critical landscape of the field, tracking approximately 540 MAs, offering a novel parameter-based taxonomy, and identifying open challenges and future directions.

## 2. Core method or approach

- Compiled a catalog of approximately 540 MAs from the literature (up to 2022), accompanied by statistical data from Scopus on publication trends, document types, and citations of top algorithms.
- Provided constructive criticism documenting known cases of "duplicate" or PSO-like algorithms: BHO, GWO, FA, and BA are all shown to be variants of PSO; IWD is a variant of ACO; HS is equivalent to ES. Approximately 57 algorithms identified as PSO-like, 24 as GA-like, and 24 as DE-like (Molina et al. 2020).
- Proposed a novel taxonomy based on the number of primary control parameters, classifying MAs into seven categories: free-parameter, mono-parameter, bi-parameter, tri-parameter, tetra-parameter, penta-parameter, and miscellaneous (6+ parameters). This provides mathematical insight into algorithmic complexity and tuning burden.
- Reviewed four existing taxonomies (source of inspiration, population size, population movement/behavior, and the new parameter-based classification) with their relative advantages.
- Surveyed real-world applications across NP-hard problems (TSP, VRP, JSS), medical science, semantic web, industry (5G deployment, self-driving), swarm drones/robotics, differential equations, and image processing.

## 3. Main results

- **Scale of the field**: Approximately 540 MAs cataloged; Scopus data shows ~99,877 published documents (articles + conference papers) in the domain as of Dec 2022, with a strong upward trend (R² = 0.995).
- **Top-cited MAs (2012–2022)**: PSO, GA, and DE dominate as the most cited algorithms.
- **Parameter taxonomy findings**: The majority of algorithms fall into the 1–5 parameter range. Algorithms with fewer primary parameters (e.g., free-parameter and mono-parameter) are easier to tune and adapt, making them preferable for industrial black-box problems. PSO has five primary parameters; DE and GWO each have two; ABC has one; TLBO has zero.
- **Novelty crisis**: Many recent MAs lack originality and are effectively re-branded variants of PSO, GA, DE, ACO, or ABC. The paper calls for constructive debate and suggests new algorithms should only be introduced when existing ones fail on real-world problems or when a genuinely more intelligent mechanism is discovered.
- **Application evidence**: GA-based methods achieved ~86% sensitivity and ~98% accuracy for lung nodule detection; MAs approximate ODE/PDE solutions where traditional methods fail; metaheuristics are widely used for TSP, UAV path planning, 5G deployment, and medical image processing.

## 4. Strengths and limitations

**Strengths**: Extraordinarily comprehensive catalog; the parameter-based taxonomy is a genuine novelty; strong critical stance on the field's methodological issues; bridges theoretical frameworks (NFL theorem, structural bias, exploration-exploitation) with practical applications.

**Limitations**: Survey-style paper — no new experimental benchmarks or empirical comparisons of algorithms; the parameter count taxonomy is descriptive rather than predictive; does not deeply analyze the runtime complexity or convergence properties of specific algorithms; the large table of 540 algorithms lists names and references but does not evaluate individual algorithm quality.

## 5. Practical takeaway for researchers

Resist the urge to invent yet another metaphor-based MA without rigorous proof of novelty. Before proposing a new algorithm, benchmark against modern, well-tuned versions of PSO, GA, DE, ACO, and ABC — not their original 1990s implementations. Prefer algorithms with fewer primary parameters when tackling real-world black-box problems, as tuning burden increases with parameter count. Key open problems remain: scaling to large-scale global optimization (dimensions > 1000), developing a unified mathematical framework for convergence/stability analysis, structural bias evaluation, and bridging the gap between benchmark function performance and real-world problem performance.
