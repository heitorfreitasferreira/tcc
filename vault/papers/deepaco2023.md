---
title: "DeepACO: Neural-enhanced Ant Colony Optimization for Combinatorial Optimization"
authors: [Ye, Haopeng, Wang, Jian, Liang, Hong, Cao, Zhiguang, Li, Yong, Li, Fanzhang]
year: 2023
doi: ""
bibtex-key: deepaco2023
tags: [aco metaheuristic tsp deep-learning]
status: lido
rating: 4
---

## Resumo

DeepACO é um framework não-autorregressivo que integra deep reinforcement learning com ACO. Utiliza Graph Neural Network (GNN) treinada com PPO para gerar heurísticas específicas por instância, guiando a decisão probabilística das formigas. DeepACO melhora significativamente o ACO clássico mantendo escalabilidade para 1000+ nós. Serve como base para NeuFACO (2025) e GFACS (2025).

## Contribuições Principais

- Primeiro framework que integra GNN + PPO com ACO de forma não-autorregressiva
- Heurísticas aprendidas por instância (instance-specific) substituem heurísticas fixas (e.g., 1/d_ij)
- Escalável para TSP de até 1000+ nós sem retreino
- Generalizável para outros COPs (CVRP, OP, KP, etc.)

## Relevância para o TCC

Representa a fronteira mais recente em hibridização neural + swarm. Útil para a seção de trabalhos relacionados/estado-da-arte, mostrando direções futuras além do escopo do TCC. DeepACO é precursor de NeuFACO, que é mais eficiente.

## Métodos e Abordagens

- GNN (Graph Neural Network) para embedding do grafo TSP
- PPO (Proximal Policy Optimization) como algoritmo de RL
- Heurística neural substitui η_ij = 1/d_ij na regra de transição ACO
- Framework não-autorregressivo (tour completo gerado em paralelo)

## Conexões

- [[dorigo1996ant]] — ACO original, base conceitual
- [[dorigo1997ant]] — ACO para TSP
- [[neufaco2025]] — evolução neural ACO mais eficiente
- [[ppaco2024]] — abordagem RL paralela (policy gradient + ACO)
- [[kappagantula2025dpso]] — DPSO-Q, swarm+RL paralela
- [[ant-colony]]
