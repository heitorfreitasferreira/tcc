---
title: "Policy Gradient and Experience Replay in Ant Colony Optimization"
authors: [Sheppard, John, et al.]
year: 2024
doi: ""
bibtex-key: ppaco2024
tags: [aco metaheuristic rl]
status: lido
rating: 4
---

## Resumo

Integra princípios de RL (Policy Gradient e PPO) ao ACO tradicional. PGACO e PPOACO substituem a atualização de feromônio por gradientes de política, usando experience replay como generalização de estratégias elitistas. Experimentos em 8 instâncias TSPLIB mostram que PPOACO supera consistentemente variantes clássicas (AS, MMAS) e baseadas em gradiente (ACOSGD, ADACO). A unificação de otimização populacional com RL melhora significativamente a capacidade de busca.

## Contribuições Principais

- Substituição da atualização de feromônio por gradientes de política (PGACO)
- PPOACO: PPO + ACO com experience replay
- Experience replay como generalização natural do elitismo ACO
- Vantagem estatisticamente significativa sobre AS, MMAS, ACOSGD, ADACO em TSPLIB

## Relevância para o TCC

Terceira abordagem neural+ACO (com DeepACO e NeuFACO) que demonstra a tendência de hibridização swarm + deep RL. Útil para discussão de estado-da-arte e trabalhos futuros. Abordagem mais simple que DeepACO (sem GNN), usando apenas policy gradient.

## Métodos e Abordagens

- Policy Gradient (REINFORCE) aplicado à seleção de arestas
- PPO (Proximal Policy Optimization) para treinamento estável
- Experience replay buffer armazena transições (estado, ação, recompensa)
- Grafo TSP completo como estado observável
- Função de recompensa: tour length negativo (minimização)

## Conexões

- [[dorigo1996ant]] — ACO original
- [[dorigo1997ant]] — ACO para TSP
- [[deepaco2023]] — DeepACO (GNN + PPO, abordagem neural alternativa)
- [[neufaco2025]] — NeuFACO (focused ACO neural, estado-da-arte)
- [[kappagantula2025dpso]] — DPSO-Q, swarm+RL paralela no PSO
- [[ant-colony]]
