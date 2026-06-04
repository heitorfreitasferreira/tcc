---
title: "Efficient Optimization of the Held–Karp Lower Bound"
authors: [Righini, Giovanni]
year: 2021
doi: "10.5802/ojmo.11"
bibtex_key: righini2021efficient
bibtex-key: righini2021efficient
pdf: "papers/pdfs/righini2021efficient.pdf"
tags: [tsp, lower-bound, held-karp, algorithm]
status: lido
rating: 4
---

## Resumo

Demonstra que a escolha do vértice p no bound HK (que Held e Karp selecionavam arbitrariamente) pode ser otimizada para fornecer o maior bound possível com a mesma complexidade O(m + n log n) de uma única MST. Apresenta um algoritmo que computa todas as arestas alternativas necessárias para reconstruir uma MST quando um vértice é removido do grafo, permitindo selecionar o vértice p* que maximiza LB*HK.

## Contribuições Principais

- Algoritmo para otimizar a escolha do vértice p no bound HK (LB*HK ≥ LB*H de Helsgaun)
- Mesma complexidade O(m + n log n) que uma MST
- Demonstração de que LB*HK domina o bound de Helsgaun (LB*H)
- Experimentos mostram melhoria de 10-20% sobre o bound HK com escolha arbitrária

## Relevância para o TCC

O bound HK otimizado é útil como referência para avaliar a qualidade das soluções dos métodos bio-inspirados. Para o TSP-SD-ATP, aplicável após redução `c'[j][k] = min_i cost[i][j][k]`. A melhoria de 10-20% sobre a versão com escolha arbitrária de vértice pode fazer diferença na avaliação de gap.

## Métodos e Abordagens

- Algoritmo para pré-computação de arestas alternativas em MSTs
- Identificação das duas arestas mais baratas incidentes a cada vértice
- Uso de estruturas Union-Find para processamento eficiente de arestas em ordem crescente
- Otimização sobre todos os vértices candidatos p ∈ V

## Conexões

- [[heldkarp1970traveling]] — bound HK original com escolha arbitrária de p
- [[johnson1996asymptotic]] — validação empírica
- [[valenzuela1997estimating]] — implementação de subgradiente
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- Três bound de 1-tree comparados: LB0 (mínimo global), LBH (Helsgaun — folha da MST), LBHK (Held-Karp — vértice removido)
- LB*HK é sempre ≥ LB*H ≥ LB0
- O algoritmo permite escolher o melhor vértice sem computar n MSTs separadas

## Citações-chave

> "The algorithm presented here shows that the vertex p that Held and Karp selected in an arbitrary way can be selected in an optimal way with the same worst-case time complexity required to compute a minimum spanning tree."

![[righini2021efficient.pdf]]
