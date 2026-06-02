---
title: "{PSO} and the Traveling Salesman Problem: An Intelligent Optimization Approach"
authors: [Araújo, Kael Silva, Barboza, Francisco Márcio]
year: 2025
doi: ""
bibtex-key: araujo2025pso
tags: [pso tsp metaheuristic]
status: lido
rating: 3
---

## Resumo

Explora a aplicação de Particle Swarm Optimization (PSO) ao TSP, adaptando o algoritmo originalmente contínuo para o domínio discreto via representação por permutação. Incorpora busca local (2-opt e 3-opt) para melhoria das soluções. Resultados mostram bom desempenho para instâncias pequenas, mas degradação em problemas maiores.

## Contribuições Principais

- Adaptação de PSO para TSP com representação por permutação
- Uso de 2-opt e 3-opt como busca local
- Comparação com GA e Simulated Annealing em benchmarks

## Relevância para o TCC

Referência direta para a implementação de PSO para TSP no repositório. A análise de limitações (desempenho reduzido em instâncias grandes) é relevante para discussão dos resultados experimentais.

## Métodos e Abordagens

- PSO discreto com permutações
- Busca local: 2-opt e 3-opt
- Benchmarks TSPLIB

## Conexões

- [[kennedy1995particle]] — PSO original, base do método
- [[sun2024hybrid]] — PSO híbrido para TSP (mesma linha)
- [[kappagantula2025dpso]] — DPSO com RL para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[lin1973effective]] — busca local 2-opt e 3-opt usada aqui
- [[particle-swarm]]
- [[TSP]]

## Notas e Insights

- PSO funciona bem para TSPs pequenos/médios mas sofre em grandes instâncias
- Hibridização com busca local (2-opt/3-opt) é essencial para desempenho competitivo
- Limitação em escala motivou abordagens híbridas como DPSO-Q e matrix-PSO
- O artigo confirma a importância de mecanismos de diversidade para evitar ótimos locais em TSP
- Preprint disponível no arXiv (2501.15319) — publicado como preprint no Qeios

## Citações-chave

> "PSO performs well for small to medium-sized problems, though its performance diminishes for larger instances due to difficulties in escaping local optima."

> "PSO is a promising approach for solving TSP, with potential for further improvement through hybridization with other optimization techniques."
