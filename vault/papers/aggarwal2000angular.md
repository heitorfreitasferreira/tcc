---
title: "The Angular-Metric Traveling Salesman Problem"
authors: [Aggarwal, Alok, Coppersmith, Don, Khanna, Sanjeev, Motwani, Rajeev, Schieber, Baruch]
year: 2000
doi: "10.1137/S0097539796312719"
bibtex-key: aggarwal2000angular
tags: [tsp, angular-cost, turn-penalty, approximation]
status: lido
rating: 4
---

## Resumo

Formula o Angular-Metric TSP (AM-TSP): dados pontos no plano Euclideano, encontrar um tour que minimize a soma dos ângulos de mudança de direção em cada vértice. Estabelece NP-hardness do problema e de sua relaxação para cycle cover. Apresenta algoritmos de aproximação O(log n) para o AM-TSP e para a versão conjunta (ângulo + distância). Fornece bounds justos para o valor extremo da medida angular.

## Contribuições Principais

- Primeira formulação e análise do AM-TSP como problema de otimização combinatória
- NP-hardness do AM-TSP e do cycle cover com custo angular
- Algoritmo de aproximação O(log n) para AM-TSP
- Algoritmo de aproximação para a soma ponderada ângulo + distância
- Bounds justos (tight) para o valor extremo da medida angular

## Relevância para o TCC

O AM-TSP é a variante mais próxima do componente de ângulo no TSP-SD-ATP. A diferença é que o TSP-SD-ATP tem custo total = distância + penalidade angular (makespan), enquanto o AM-TSP minimiza apenas ângulo. A NP-hardness do AM-TSP confirma que o TSP-SD-ATP é NP-difícil. A aproximação O(log n) serve como referência superior para métodos heurísticos.

## Métodos e Abordagens

- Redução do AM-TSP ao problema do cycle cover em grafo direcionado
- Algoritmo de aproximação: amostragem de direções + ATSP no grafo resultante
- Análise combinatória do número mínimo de curvas em um tour
- Bounds assintóticos para ângulo total mínimo

## Conexões

- [[winter2002modeling]] — modelagem de custos de curva via pseudo-dual graph
- [[vanhove2012route]] — experimentos computacionais com turn restrictions
- [[fischetti1992additive]] — additive bounding para ATSP (aplicável à redução)
- [[kinable2017hybrid]] — MDD para sequenciamento (generaliza AM-TSP)
- [[tsp]]
- [[tsp-variants]]
- [[lower-bounds]]

## Notas e Insights

- O AM-TSP puro (só ângulo) não captura o makespan completo do TSP-SD-ATP
- NP-hardness do cycle cover angular implica que relaxações simples (assignment) não são triviais
- A abordagem de discretização de direções (amostragem) pode ser usada para construir lower bounds
- Fekete & Woeginger (1996) também estudaram ângulo-restricted tour problems

## Citações-chave

> "We establish the NP-hardness of both this problem and its relaxation to the cycle cover problem."

> "We show that both problems can be approximated to within a ratio of O(log n) in polynomial time."
