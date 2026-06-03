---
title: "Hybrid Particle Swarm Optimization Algorithm for Traveling Salesman Problem Based on Ternary Optical Computer"
authors: [Sun, Shaojiang]
year: 2024
doi: "10.1016/j.procs.2024.09.151"
bibtex-key: sun2024hybrid
tags: [pso tsp metaheuristic]
status: lido
rating: 3
---

## Resumo

Propõe um algoritmo híbrido HPSO (Hybrid Particle Swarm Optimization) para TSP, combinando três estratégias: (1) geração greedy da população inicial para qualidade superior de partida, (2) critério de Metropolis para aceitação probabilística de soluções inferiores (aumentando diversidade), e (3) 2-opt como operador de mutação para refino local. O algoritmo é implementado sobre Ternary Optical Computer (TOC), explorando mega-paralelismo óptico para acelerar computação. Resultados mostram convergência rápida e soluções de alta qualidade em instâncias TSP.

## Contribuições Principais

- HPSO integrando greedy strategy + Metropolis criterion + 2-opt em framework unificado
- Primeira implementação de PSO híbrido para TSP em Ternary Optical Computer
- Demonstração de que técnicas clássicas (greedy, Metropolis, 2-opt) combinadas superam limitações individuais do PSO
- Convergência acelerada com qualidade de solução competitiva

## Relevância para o TCC

Ilustra abordagem híbrida relevante para PSO aplicado a TSP, com técnicas transferíveis ao código do repositório. A geração greedy de população e o 2-opt como busca local são particularmente aplicáveis ao GA e PSO já implementados. O critério de Metropolis (inspirado em Simulated Annealing) é uma técnica de escape de ótimos locais que pode beneficiar qualquer metaheurística no projeto.

## Métodos e Abordagens

- PSO discreto com representação de permutação de cidades
- Geração greedy da população inicial (abordagem construtiva)
- Critério de Metropolis para aceitação probabilística de soluções piores
- 2-opt como operador de mutação para eliminação de cruzamentos
- Ternary Optical Computer (TOC) para paralelização massiva
- Validação em instâncias TSP (não especificadas no abstract)

## Conexões

- [[kennedy1995particle]] — PSO original, base do método
- [[araujo2025pso]] — PSO discreto com 2-opt/3-opt para TSP
- [[kappagantula2025dpso]] — DPSO com RL para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[lin1973effective]] — 2-opt usado como busca local
- [[particle-swarm]]
- [[tsp]]
- [[bio-inspired-optimization]]

## Notas e Insights

- Open Access (CC BY-NC-ND) na ScienceDirect: https://doi.org/10.1016/j.procs.2024.09.151 — PDF disponível na página do DOI
- Geração greedy de população melhora significativamente a qualidade inicial, reduzindo iterações necessárias
- Metropolis criterion é importado do Simulated Annealing — crossover interessante entre metaheurísticas
- 2-opt é eficaz para eliminar cruzamentos em rotas TSP; técnica já usada no GA do repositório
- A implementação em TOC é o diferencial, mas limita a reprodutibilidade (hardware especializado não disponível)
- Sem TOC, as técnicas de hibridização (greedy + Metropolis + 2-opt) são aplicáveis em hardware convencional
- Limitação: abstract não detalha datasets específicos nem métricas de comparação

## Citações-chave

> By introducing greedy strategy, the algorithm generates excellent populations. And the improved Metropolis criterion is added to individual learning and population learning to accept inferior solutions with certain probability, which can effectively increase the population diversity.

> The results show that HPSO solves the TSP with fast convergence speed and high quality of the optimal solution.
