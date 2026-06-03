---
title: "Matrix-Based Particle Swarm Optimization with Hybrid Strategy for Multi-Traveling Salesman Problem"
authors: [Huang, Zhaoquan, Zhu, Donglin, Zhou, Changjun, Cheng, Shi]
year: 2025
doi: "10.1007/s12065-025-01082-2"
bibtex-key: huang2025matrix
tags: [pso tsp metaheuristic]
status: lido
rating: 3
---

## Resumo

Propõe um algoritmo PSO baseado em representação matricial com estratégias múltiplas para o Multiple Traveling Salesman Problem (mTSP). A representação matricial permite computação paralela eficiente. Inclui mecanismo de diversidade dinâmica para balancear busca local e global. Nova função de fitness também aborda diferenças de rota entre vendedores.

## Contribuições Principais

- Representação matricial para PSO em mTSP (paralelizável)
- Mecanismo de diversidade dinâmica
- Função de fitness multi-objetivo (distância total + balanceamento)
- Superior a MPSO, HJSPSO e PSO padrão em 5 benchmarks TSP

## Relevância para o TCC

Extensão natural do TSP para múltiplos agentes (múltiplos drones). A abordagem de balanceamento de rotas é relevante para cenários de patrulha com múltiplos UAVs.

## Métodos e Abordagens

- PSO matricial com paralelismo
- Mecanismo de diversidade dinâmica
- Função de fitness ponderada (caminho total + equilíbrio entre rotas)

## Conexões

- [[kennedy1995particle]] — PSO original, base do método
- [[araujo2025pso]] — PSO discreto para TSP (mono-agente)
- [[sun2024hybrid]] — PSO híbrido para TSP
- [[kappagantula2025dpso]] — DPSO com RL para TSP
- [[particle-swarm]]
- [[tsp]]

## Notas e Insights

- Springer paywall — sem versão OA encontrada; disponível em: https://doi.org/10.1007/s12065-025-01082-2
- Balanceamento entre agentes é tão importante quanto minimizar distância total
- Representação matricial facilita implementação em GPU
- Abordagem publicada em 2025 no periódico *Evolutionary Intelligence* (vol. 18)
- A função de fitness multi-objetivo (distância + balanceamento) é diretamente aplicável ao rTSC com múltiplos drones

## Citações-chave

> "Matrix-based representation enables efficient parallel computation for the multi-traveling salesman problem."

> "The dynamic diversity mechanism balances local and global search, preventing premature convergence in PSO."
