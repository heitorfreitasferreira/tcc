---
title: "Estimating the Held-Karp Lower Bound for the Geometric TSP"
authors: [Valenzuela, Christine L., Jones, Antonia J.]
year: 1997
doi: "10.1016/S0377-2217(97)00360-1"
bibtex-key: valenzuela1997estimating
tags: [tsp, lower-bound, held-karp, lagrangean]
status: lido
rating: 4
---

## Resumo

Apresenta uma implementação prática e detalhada do método de subgradiente para estimar o Held-Karp lower bound para o TSP geométrico. Compara duas abordagens para calcular o 1-tree mínimo com custos modificados: uma baseada no algoritmo de Prim (O(n²)) e outra baseada na construção de uma árvore geradora mínima (MST) com atualização de pesos. Relata experimentos em instâncias TSPLIB e aleatórias, com foco na qualidade do bound obtido e no número de iterações necessárias.

## Contribuições Principais

- Comparação detalhada de duas implementações de subgradiente para o HK bound
- Análise do impacto de diferentes fórmulas de step size na convergência
- Demonstração de que o método subgradiente atinge consistentemente ≥ 99% do bound exato
- Discussão sobre critérios de parada e inicialização dos multiplicadores

## Relevância para o TCC

Fornece um blueprint de implementação do subgradiente HK bound em código. A abordagem pode ser adaptada para o TSP-SD-ATP via redução do tensor 3D para matriz 2D `c'[j][k] = min_i cost[i][j][k]`. A discussão sobre tamanho de passo e critérios de parada é diretamente aplicável.

## Métodos e Abordagens

- Algoritmo de Prim para MST com custos Lagrangianos modificados
- Subgradiente: π_i^{(m+1)} = π_i^{(m)} + t^(m)(deg_i - 2)
- Step size: t^(m) = α^(m)(UB - w(π^(m)))/Σ(deg_i - 2)²
- Estratégia de redução de α por fator 0.5 quando w(π) não melhora

## Conexões

- [[heldkarp1970traveling]] — formulação original
- [[heldkarp1971traveling]] — subgradiente
- [[johnson1996asymptotic]] — validação empírica em larga escala
- [[righini2021efficient]] — otimização da escolha do vértice p
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- O step size proposto por Held & Karp original é: t = α(UB - LB)/Σ(deg_i - 2)²
- α começa em 2.0 e é reduzido quando não há melhoria por K iterações (típico: α = α/2 a cada 2-5 iterações sem melhoria)
- O UB (upper bound) pode ser obtido de uma heurística construtiva (e.g., nearest neighbor, double tree)
- O bound converge rapidamente nas primeiras iterações, depois estabiliza

## Citações-chave

> "The Held-Karp lower bound, which is relatively quick and easy to compute using the above techniques, has enormous practical value when evaluating the quality of near optimal solutions for large problems where the true optima are not known."
