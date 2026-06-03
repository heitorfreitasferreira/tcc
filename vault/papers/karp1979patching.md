---
title: "A Patching Algorithm for the Nonsymmetric Traveling-Salesman Problem"
authors: [Karp, Richard M.]
year: 1979
doi: "10.1137/0208007"
bibtex-key: karp1979patching
tags: [atsp, assignment, patching]
status: lido
rating: 4
---

## Resumo

Propõe um algoritmo de patching para converter uma solução do Assignment Problem (AP) — que produz um conjunto de subtours — em um tour válido para o ATSP. O algoritmo opera em O(n³) e consiste em intercalar os subtours dois a dois, escolhendo a melhor maneira de concatená-los (break-and-repair). Para custos U[0,1], Karp prova que o gap ATSP - AP é o(1), explicando por que a relaxação AP é tão eficaz para ATSP.

## Contribuições Principais

- Algoritmo de patching O(n³) para converter subtours do AP em tour ATSP
- Prova probabilística: para ATSP com custos U[0,1], E[ATSP - AP] = o(1)
- Explicação teórica para a eficácia empírica da relaxação AP

## Conexões

- [[balas1985branch]] — B&B para ATSP com relaxação AP
- [[fischetti1992additive]] — additive bounding usa AP como primeiro componente
- [[heldkarp1970traveling]] — bound HK alternativo para TSP simétrico
- [[tsp]]
- [[lower-bounds]]
