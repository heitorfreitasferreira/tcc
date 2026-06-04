---
title: "NeuFACO: Neural Focused Ant Colony Optimization for Traveling Salesman Problem"
authors: [Tran, Dat Thanh, Tran, Khai Quang, Pham, Khoi Anh, Vu, Van Khu, Do, Dong Duc]
year: 2025
doi: "10.48550/arXiv.2503.08812"
bibtex-key: neufaco2025
pdf: "papers/pdfs/neufaco2025.pdf"
tags: [aco metaheuristic tsp deep-learning]
status: lido
rating: 5
---

## PDF

![[neufaco2025.pdf]]

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

## Notas e Insights

- NeuFACO resolve uma limitação importante do DeepACO: em vez de reconstruir tours completos a cada iteração, o Focused ACO (FACO) modifica seletivamente nós ao redor de uma solução de referência, preservando subestruturas fortes
- A redução de 60× no tempo de execução em relação ao DeepACO e GFACS é impressionante, mantendo qualidade comparável (gaps de 1.16-2.98% em TSPLIB)
- A regularização de entropia no PPO evita colapso de política, um problema comum em RL para otimização combinatória
- Para o TCC, NeuFACO representa o estado-da-arte em neural-enhanced ACO e é essencial citar na seção de trabalhos relacionados
- O gap pequeno (~1-3%) mesmo em instâncias grandes sugere que métodos clássicos (GA, PSO, ACO puros) chegam perto do ótimo, validando a abordagem do TCC
- A escalabilidade para 1500 nós TSPLIB mostra que a abordagem é prática para problemas reais de patrulha com drones

## Citações-chave

> "NeuFACO integrates deep reinforcement learning with a refined ACO framework to address limitations of non-autoregressive models for the TSP."

> "By combining PPO-based policy learning with targeted refinement around high-quality solutions, it achieves a strong balance between global guidance and local exploitation."

> "Experiments show that NeuFACO consistently achieves superior or highly competitive performance compared to a wide range of neural baselines on both randomized and benchmark datasets."
