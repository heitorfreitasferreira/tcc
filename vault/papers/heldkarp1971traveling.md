---
title: "The Traveling-Salesman Problem and Minimum Spanning Trees: Part II"
authors: [Held, Michael, Karp, Richard M.]
year: 1971
doi: "10.1007/BF01584070"
bibtex-key: heldkarp1971traveling
tags: [tsp, lower-bound, lagrangean]
status: lido
rating: 5
---

## Resumo

Segunda parte do trabalho seminal de Held e Karp sobre lower bounds para TSP. Apresenta o algoritmo de programação dinâmica O(n²2ⁿ) para solução exata do TSP (conhecido como algoritmo de Held-Karp), e aprofunda a otimização subgradiente para maximizar a função Lagrangiana do bound de 1-tree. Estabelece formalmente a convergência do método de subgradiente e fornece diretrizes para tamanho de passo.

## Contribuições Principais

- Algoritmo de programação dinâmica O(n²2ⁿ) para TSP exato — estado-da-arte em pior caso até hoje
- Formalização do subgradiente para maximizar w(π) na relaxação 1-tree
- Prova de convergência do método de subgradiente para bound HK
- Estratégias de passo (step size) baseadas em limitantes superior e inferior

## Relevância para o TCC

A programação dinâmica O(n²2ⁿ) é exata para TSP clássico e serve como referência. O método de subgradiente para o HK bound pode ser adaptado para o TSP-SD-ATP via redução do tensor 3D para matriz 2D. A estrutura de DP também pode ser adaptada para solução exata em instâncias pequenas (n ≤ 15-20) do TSP-SD-ATP usando estados (S, j, i) com custo O(n³2ⁿ).

## Métodos e Abordagens

- Programação dinâmica: dp[S][j] = min custo de caminho de 1 a j visitando S
- Subgradiente para maximizar w(π) com ajuste iterativo de multiplicadores
- Tamanho de passo: t^(m) → 0, Σ t^(m) → ∞
- Critério de parada baseado em grau dos vértices na 1-tree

## Conexões

- [[heldkarp1970traveling]] — Part I com formulação Lagrangiana original
- [[johnson1996asymptotic]] — implementação prática e validação empírica
- [[righini2021efficient]] — otimização da seleção de vértice p
- [[balas1985branch]] — branch-and-bound baseado em relaxação assignment (ATSP)
- [[tsp]]

## Notas e Insights

- O algoritmo de DP O(n²2ⁿ) permanece como o melhor pior caso conhecido para TSP exato
- A atualização dos multiplicadores: π_i^{(m+1)} = π_i^{(m)} + t^(m)(deg_i - 2)
- A convergência não é monótona — a função w(π) pode oscilar

## Citações-chave

> "We describe a dynamic programming algorithm which solves the traveling-salesman problem in O(n²2ⁿ) steps."

> "The results of the previous paper are extended to include a subgradient optimization procedure which, in practice, converges rapidly to the maximum of the lower bound function."
