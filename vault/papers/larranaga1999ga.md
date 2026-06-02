---
title: "Genetic Algorithms for the Travelling Salesman Problem: A Review of Representations and Operators"
authors: [Larrañaga, Pedro, Kuijpers, C. M. H., Murga, R. H., Inza, I., Dizdarevic, S.]
year: 1999
doi: "10.1023/A:1006529012972"
bibtex-key: larranaga1999ga
pdf: "papers/pdfs/larranaga1999ga.pdf"
tags: [ga tsp survey]
status: lido
rating: 5
---

## PDF

[[papers/pdfs/larranaga1999ga.pdf]]

## Resumo

This paper is the result of a literature study carried out by the authors. It is a review of the different attempts made to solve the Travelling Salesman Problem with Genetic Algorithms. We present crossover and mutation operators, developed to tackle the TSP with Genetic Algorithms with different representations such as: binary representation, path representation, adjacency representation, ordinal representation and matrix representation. Likewise, we show the experimental results obtained with different standard examples using combinations of crossover and mutation operators in relation with path representation.

**Operadores de crossover cobertos:** PMX (Partially Mapped Crossover), OX (Order Crossover), CX (Cycle Crossover), ERX (Edge Recombination Crossover), Heuristic Crossover, SCX (Sequential Constructive Crossover).

**Operadores de mutação cobertos:** Swap (exchange), Insertion, Inversion (reversal), Displacement, Scramble.

## Contribuições Principais

- Revisão exaustiva e sistemática de representações e operadores de GA para TSP
- Experimentos comparativos: PMX + Swap/Inversion se destacam como combinação mais robusta
- Categorização clara de operadores em grupos (preservação de posição, ordem, arestas)

## Relevância para o TCC

Artigo de referência para justificar a escolha de operadores de crossover (e.g., PMX, OX, ERX) e mutação (e.g., swap, inversão) em GAs para TSP/rTSP. Os experimentos comparativos fornecem evidência empírica para decisões de implementação.

## Métodos e Abordagens

- Crossover: PMX, OX, CX, ERX, SCX, Heuristic Crossover
- Mutação: Swap, Insertion, Inversion, Displacement, Scramble
- Representações: binária, path (permutação), adjacência, ordinal, matricial
- Experimentos em benchmarks TSP (Oliver30, KroA100, etc.)

## Conexões

- [[holland1975adaptation]] — GA fundacional
- [[goldberg1989genetic]] — GA textbook
- [[oliver1987crossover]] — crossover operators (PMX, OX, CX)
- [[bean1994genetic]] — random keys
- [[nagata2006eax]] — EAX (crossover estado-da-arte, evolução posterior)
- [[potvin1996ga]] — survey anterior focado em categorização de crossover
- [[genetic-algorithms]]
- [[TSP]]
