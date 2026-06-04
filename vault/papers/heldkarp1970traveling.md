---
title: "The Traveling-Salesman Problem and Minimum Spanning Trees"
authors: [Held, Michael, Karp, Richard M.]
year: 1970
doi: "10.1287/opre.18.6.1138"
bibtex_key: heldkarp1970traveling
bibtex-key: heldkarp1970traveling
pdf: "papers/pdfs/heldkarp1970traveling.pdf"
tags: [tsp, lower-bound, lagrangean]
status: lido
rating: 5
---

## Resumo

Propõe o método de relaxação Lagrangiana para o TSP baseado em 1-trees. A ideia central é relaxar as restrições de grau de cada vértice (cada vértice deve ter grau 2 num tour) usando multiplicadores de Lagrange, resultando num subproblema de 1-tree de custo mínimo que pode ser resolvido eficientemente (MST + duas arestas). O máximo da função Lagrangiana sobre os multiplicadores produz o **Held-Karp lower bound**, que é idêntico à solução da relaxação LP do TSP (subtour LP). Apresenta três métodos de solução: geração de colunas, ascent method, e branch-and-bound.

## Contribuições Principais

- Formulação do TSP como maximização de uma função côncava sobre 1-trees
- Equivalência entre o bound Lagrangiano e a relaxação LP do TSP
- Três métodos para encontrar o bound: column generation, ascent method (subgradiente), e branch-and-bound
- Demonstração de que o método fornece limites apertados na prática

## Relevância para o TCC

O Held-Karp bound é o lower bound mais utilizado para TSP simétrico, com gap empírico < 0.8% do ótimo. Para o TSP-SD-ATP (tensor 3D), o método não se aplica diretamente, mas a redução `c'[j][k] = min_i cost[i][j][k]` permite aplicar o HK bound sobre uma matriz 2D reduzida, fornecendo um lower bound válido.

## Métodos e Abordagens

- Relaxação Lagrangiana de restrições de grau
- Subproblema de 1-tree: MST sobre V\{1} + 2 arestas mais baratas incidentes a 1
- Geração de colunas no espaço dual
- Ascent method para otimização subgradiente
- Estrutura de branch-and-bound baseada em graus

## Conexões

- [[heldkarp1971traveling]] — Part II, com formulação de programação dinâmica O(n²2ⁿ)
- [[johnson1996asymptotic]] — validação empírica do HK bound em larga escala
- [[valenzuela1997estimating]] — implementação prática do subgradiente para HK
- [[righini2021efficient]] — otimização da seleção de vértice p no bound HK
- [[lawler1985traveling]] — survey que consolida o método
- [[applegate2006traveling]] — Concorde usa subtour LP (=HK bound) como relaxação
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- A transformação de custos c'_{ij} = c_{ij} + π_i + π_j não altera o custo de tours, mas altera o custo de 1-trees
- A função w(π) é côncava por partes, maximizada por subgradiente
- O bound independe do vértice 1 escolhido, mas Righini (2021) mostra que a escolha ótima do vértice pode ser computada em O(m + n log n)
- Para o tensor 3D do TSP-SD-ATP, a redução por min sobre a dimensão anterior permite usar este bound

## Citações-chave

> "We describe a method for computing lower bounds for the traveling-salesman problem which is based on the concept of a minimum spanning 1-tree."

> "The method proceeds by maximizing a concave function defined on a certain convex polytope; the value of this function is a lower bound on the cost of an optimal tour."

![[heldkarp1970traveling.pdf]]
