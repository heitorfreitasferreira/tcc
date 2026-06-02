---
title: "NeuFACO: Neural Focused Ant Colony Optimization for Traveling Salesman Problem"
authors: [Ye, Haopeng, Wang, Jian, Liang, Hong, Cao, Zhiguang, Li, Yong, Li, Fanzhang]
year: 2025
doi: ""
bibtex-key: neufaco2025
tags: [aco metaheuristic tsp deep-learning]
status: lido
rating: 5
---

## Resumo

NeuFACO combina deep RL (PPO + entropy regularization) com Focused ACO (FACO). Diferente do ACO clássico que reconstrói tours completos, o FACO modifica seletivamente nós ao redor de uma solução de referência, preservando subestruturas fortes. Alcança gaps de 1.16-2.98% em TSPLIB (até 1500 nós) com tempos até 60× menores que DeepACO e GFACS. É o estado-da-arte em aprendizado neural para ACO.

## Contribuições Principais

- Focused ACO (FACO): reconstrói seletivamente apenas subconjuntos de nós, não o tour inteiro
- 60× mais rápido que DeepACO e GFACS com qualidade comparável
- Regularização de entropia no PPO evita colapso de política
- Escalável para instâncias de até 1500 nós TSPLIB
- Estado-da-arte em neural-enhanced ACO para TSP

## Relevância para o TCC

O resultado mais recente e relevante na linha neural + swarm. Embora fora do escopo de implementação, é essencial citar na seção de trabalhos relacionados para demonstrar conhecimento das fronteiras da área. O gap pequeno (~1-3%) mostra que mesmo métodos clássicos chegam perto do ótimo.

## Métodos e Abordagens

- Focused ACO: modifica nós ao redor de uma solução de referência
- GNN para embedding + PPO para treinamento
- Regularização de entropia para balancear exploração/explotation
- Reconexão gulosa com busca local 2-opt

## Conexões

- [[dorigo1996ant]] — ACO original
- [[dorigo1997ant]] — ACO para TSP
- [[deepaco2023]] — DeepACO, precursor direto
- [[ppaco2024]] — PGACO/PPOACO, abordagem RL alternativa
- [[kappagantula2025dpso]] — DPSO-Q, hibridização swarm+RL paralela
- [[ant-colony]]
