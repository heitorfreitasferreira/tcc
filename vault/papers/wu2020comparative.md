---
title: "A Comparative Study of Solving Traveling Salesman Problem with Genetic Algorithm, Ant Colony Algorithm, and Particle Swarm Optimization"
authors: [Wu, Zefeng]
year: 2020
doi: "10.1145/3450292.3450308"
bibtex-key: wu2020comparative
pdf: "papers/pdfs/wu2020.pdf"
tags: [tsp, ga, aco, pso, comparison, metaheuristic]
status: lido-parcial
rating: 4
---

## PDF

![[wu2020.pdf]]

## Resumo

Compara GA, ACO e PSO na resolução do TSP. Simulações mostram que PSO produz soluções mais estáveis e de melhor qualidade, ACO tem o menor tempo de execução, e GA é menos influenciado pela escala do problema.

## Contribuições Principais

- Comparação head-to-head dos três métodos mais populares no mesmo framework experimental
- Análise de estabilidade das soluções vs. escala do problema

## Relevância para o TCC

Comparação direta dos três métodos implementados no projeto (GA, PSO, ACO) com as mesmas métricas.

## Métodos e Abordagens

- GA com crossover e mutação
- ACO com depósito/evaporação de feromônio
- PSO com atualização de velocidade/posição
- Benchmarks TSPLIB

## Conexões

- [[tsp]] — problema-alvo
- [[genetic-algorithms]] — método comparado
- [[particle-swarm]] — método comparado
- [[ant-colony]] — método comparado
- [[comparative-studies]] — área temática

## Citações-chave

> PSO provided a better and more stable solution, ACO took the shortest running time and GA is less influenced by the problem scale.
