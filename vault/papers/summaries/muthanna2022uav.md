# Muthanna et al. (2022) — Towards optimal positioning and energy-efficient UAV path scheduling in IoT applications

**Journal:** Computer Communications 191 (2022) 145–160  
**DOI:** 10.1016/j.comcom.2022.04.029

---

## 1. Problem and motivation

UAV base stations are a promising solution for emergency communication (disaster management, search and rescue), but existing UAV positioning and path planning approaches neglect the impact of adverse weather conditions — wind, temperature, rain, and extreme events — on QoS, reliability, and energy efficiency. The paper addresses the lack of a joint optimization framework that simultaneously considers weather-aware positioning and energy-efficient trajectory planning for multi-UAV networks.

## 2. Core method or approach

The proposed IWPOP-UAV framework operates in four stages:

- **C-LSTM weather prediction:** A Cerebral Long Short-Term Memory (C-LSTM) network predicts weather attributes (temperature, wind speed/direction, precipitation, humidity, pressure) from historical and current environmental data, achieving lower training loss than conventional LSTM.
- **Density-aware cell partitioning:** The emergency area is tiled with hexagonal cells, partitioned via graph entropy based on user density and inter-user distance to localize target UEs.
- **A3C-based multi-UAV positioning:** An Asynchronous Advantage Actor-Critic (A3C) deep reinforcement learning algorithm determines — per cell — the number of UAVs required and their optimal 3D positions, considering predicted weather, coverage area, interference, data rate, delay sensitivity, elevation angle, and LOS/NLOS characteristics.
- **Mayfly Optimization Algorithm (MOA) for path planning:** UAV trajectory is formulated as a multi-objective problem (minimize flight time and energy consumption) and solved via MOA, which combines features of GA, PSO, and firefly optimization. Eleven parameters (wind, temperature, energy, speed, distance, obstacles, LOS/NLOS, optimality, completeness, cost efficiency, collision avoidance) guide path selection.

## 3. Main results

- **Simulation environment:** NS 3.26, 2000 × 1000 × 880 m area, 10 UAVs, 100 UEs, 1 GBS; benchmarked against SWIM and Multi-rotor UAV approaches.
- **QoS:** coverage ratio 0.847 (vs. 0.771 SWIM, 0.646 Multi-rotor); cell coverage 0.934 (vs. 0.893, 0.882); delay 36.5 ms (vs. 73.5, 89 ms).
- **Reliability:** path gain −13.3 dB (vs. −18.8, −29.8 dB); collected packets ~71k–78k (vs. ~38k–59k).
- **Energy:** UAV transmit power 4.46 W (vs. 5.53, 6.3 W); energy consumption 7.86 kJ (vs. 8.68, 10.35 kJ).
- **Time complexity:** O(K) for C-LSTM, O(N|A|²) for A3C, O(U·V·W·X) for MOA.

## 4. Strengths and limitations

**Strengths:**
- Jointly optimizes UAV positioning *and* path planning, unlike prior work that tackles only one.
- Explicitly models diverse weather effects on both A2G and A2A channels.
- Clear numerical improvements across all QoS, reliability, and energy metrics against two baselines.
- Illustrated with a concrete fire/disaster use case.

**Limitations:**
- Evaluation is simulation-only (NS 3.26); no hardware-in-the-loop or field experiments.
- Fixed-scale scenario (10 UAVs, 100 UEs); scalability beyond these numbers is untested.
- No sensitivity analysis on the accuracy of C-LSTM weather predictions or their effect on downstream A3C/MOA decisions.
- Three-stage pipeline introduces cumulative computational overhead; real-time feasibility is not analyzed.
- MOA is compared only against MOA-less baselines, not against other swarm/RL path planners.

## 5. Practical takeaway for researchers

Weather-aware joint optimization of UAV deployment and trajectory yields measurable gains in coverage, delay, packet collection, and energy consumption. The modular pipeline (C-LSTM → A3C → MOA) offers a template for emergency UAV communication systems, but real-world validation and scalability analysis remain open challenges.
