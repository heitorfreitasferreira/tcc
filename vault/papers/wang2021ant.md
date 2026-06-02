---
title: "Ant Colony Optimization for Traveling Salesman Problem Based on Parameters Optimization"
authors: [Wang, Yong]
year: 2021
doi: "10.1016/j.asoc.2021.107439"
bibtex-key: wang2021ant
pdf: pdfs/wang2021.pdf
tags: [aco tsp metaheuristic]
status: lido
rating: 0
---

![[pdfs/wang2021.pdf]]

## Resumo

Propõe um algoritmo híbrido SOS-ACO para TSP, onde Symbiotic Organisms Search (SOS) é usado para otimizar os parâmetros chave (α e β) do Ant Colony Optimization. Estratégia de otimização local acelera convergência e melhora qualidade. Resultados em instâncias TSPLIB mostram que o SOS-ACO supera ACO puro e ACO-LO.

## Contribuições Principais

- Uso de SOS para tuning automático de parâmetros do ACO
- Estratégia de otimização local integrada
- Validação em TSPLIB com resultados superiores a baseline ACO

## Relevância para o TCC

Abordagem de tuning automático de parâmetros é relevante para a implementação de ACO no repositório. A sensibilidade paramétrica do ACO é um desafio prático no TCC.

## Métodos e Abordagens

- ACO com parâmetros α e β otimizados por SOS
- Estratégia de otimização local
- Benchmarks TSPLIB (Lin318, Rd400, Pr439, Rat575)

## Conexões

- [[dorigo1996ant]] — Ant System, fundação do ACO
- [[dorigo1997ant]] — primeira aplicação do ACO ao TSP
- [[kappagantula2025dpso]] — DPSO-Q (outra hibridização swarm+RL)
- [[ant-colony]] [[ant-colony]]
- [[TSP]]

## Notas e Insights

- ACO é altamente sensível aos parâmetros α (influência do feromônio) e β (influência da heurística)
- SOS-ACO reduz erro relativo em instâncias grandes comparado a ACO padrão
- Tuning automático pode ser uma direção para melhorar o ACO implementado no TCC

## Citações-chave

>
