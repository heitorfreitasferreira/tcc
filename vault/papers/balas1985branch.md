---
title: "Branch and Bound Methods for the Traveling Salesman Problem"
authors: [Balas, Egon, Toth, Paolo]
year: 1985
bibtex-key: balas1985branch
tags: [tsp, atsp, lower-bound, branch-and-bound, survey]
status: lido
rating: 5
---

## Resumo

Capítulo do livro editado por Lawler et al. (1985) que apresenta e compara métodos de branch-and-bound para TSP baseados em três relaxações principais: (1) **Assignment Problem (AP)** — relaxação padrão para ATSP, (2) **1-tree com função Lagrangiana** — para TSP simétrico, (3) **Assignment Problem com função Lagrangiana** — alternativa para ATSP. Discute estratégias de branching (subtour elimination, arco forbid/include) e apresenta resultados computacionais comparativos.

## Contribuições Principais

- Survey abrangente de métodos B&B baseados em relaxação AP para ATSP
- Comparação experimental de AP, 1-tree e AP Lagrangiano
- Análise do branching por eliminação de subtours (Carpaneto & Toth)
- Identificação de que AP bound é extremamente bom para ATSP com custos aleatórios U[0,1]

## Relevância para o TCC

Estabelece a relaxação AP como o padrão-ouro para ATSP. Para o TSP-SD-ATP (assimétrico, 3D), a relaxação AP sobre a matriz reduzida `c'[j][k] = min_i cost[i][j][k]` fornece um lower bound válido. O branching por subtour elimination é a estratégia natural para refinar o bound.

## Métodos e Abordagens

- Relaxação Assignment Problem (AP): O(n³) via algoritmo Hungaro
- Relaxação 1-tree Lagrangiana (Held-Karp)
- Branching por subtour elimination: quebra subtours fixando/removendo arestas
- Redução de custos via reduced costs do AP

## Conexões

- [[fischetti1992additive]] — additive bounding (AP + arborescências)
- [[heldkarp1970traveling]] — 1-tree Lagrangiana
- [[karp1979patching]] — patching heuristic para converter AP em tour
- [[lawler1985traveling]] — survey volume onde este capítulo está inserido
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- Para ATSP com custos no intervalo [0, L], quanto maior L/n mais difícil o problema
- O AP bound tipicamente está a 1-3% do ótimo para ATSP com custos aleatórios
- Reduced costs do AP permitem fixar variáveis: se custo reduzido > gap, arco não pode estar em tour ótimo
- O branching por subtour é natural: escolhe-se um subtour para quebrar

## Citações-chave

> "Assignment Problem based approaches are extensively used for the Asymmetric TSP."

> "The quality of the AP bound is such that, for random problems, the optimal solution is often found at the root node."
