---
title: "A Patching Algorithm for the Nonsymmetric Traveling-Salesman Problem"
authors: [Karp, Richard M.]
year: 1979
doi: "10.1137/0208007"
bibtex-key: karp1979patching
pdf: "papers/pdfs/karp1979patching.pdf"
tags: [atsp, assignment, patching]
status: lido-parcial
rating: 4
---

## PDF

![[karp1979patching.pdf]]

## Resumo

Karp propõe um algoritmo de aproximação para o TSP assimétrico baseado em resolver primeiro um problema de atribuição e depois remendar os ciclos da permutação ótima de atribuição para formar um único ciclo hamiltoniano. O algoritmo explora o fato de que o problema de atribuição pode ser resolvido em tempo polinomial, enquanto a solução resultante pode conter vários ciclos. A etapa de patching escolhe operações de troca que unem ciclos diferentes com menor acréscimo de custo. O artigo combina projeto algorítmico com análise probabilística: para matrizes de distância aleatórias uniformes, a razão entre o custo da rota obtida e o ótimo tende a `1 + e(n)`, com `e(n)` indo a zero quando `n` cresce.

> [!warning] Leitura parcial
> O OCR do PDF é legível, mas apresenta erros moderados em fórmulas e símbolos. As citações e expressões matemáticas devem ser conferidas visualmente antes de uso na monografia.

## Contribuições Principais

- Apresenta um algoritmo polinomial de aproximação para o TSP assimétrico.
- Usa o problema de atribuição como relaxação inicial do ATSP.
- Define operações de patching que unem ciclos de uma permutação em um único tour.
- Mostra que o tempo de execução é comparável ao de resolver uma atribuição `n x n`, com ordem `O(n^3)`.
- Fornece análise probabilística para distâncias independentes uniformes.

## Relevância para o TCC

O artigo oferece uma ponte entre métodos exatos, relaxações e heurísticas para TSP. Para o TCC, ele ajuda a explicar uma estratégia clássica: resolver um problema mais fácil que preserva parte da estrutura do TSP e depois reparar a solução para obter uma rota viável. Mesmo que o TCC implemente metaheurísticas bio-inspiradas, a lógica de usar limites, relaxações e reparos é útil para interpretar por que alguns métodos conseguem boas soluções sem explorar todas as permutações.

## Métodos e Abordagens

- Formulação do ATSP como busca por uma permutação cíclica de custo mínimo.
- Solução inicial pelo problema de atribuição sobre a matriz de distâncias.
- Identificação dos ciclos na permutação ótima de atribuição.
- Escolha de um ciclo de comprimento máximo como ciclo base para unir os demais.
- Construção de um problema de atribuição auxiliar para selecionar pontos de patching.

## Conexões

- [[garey1979computers]]
- [[heldkarp1970traveling]]
- [[heldkarp1971traveling]]
- [[fischetti1992additive]]
- [[lin1973effective]]
- [[applegate2006traveling]]

## Notas e Insights

- A relaxação por atribuição produz ciclos disjuntos; a dificuldade prática está em convertê-los em um tour único sem aumentar muito o custo.
- O resultado probabilístico não é garantia uniforme para qualquer matriz de distâncias; ele depende do modelo aleatório assumido.
- O artigo reconhece que não se deve esperar aproximação polinomial com erro relativo uniformemente limitado para TSP geral.
- O algoritmo é clássico para ATSP, mas menos diretamente aplicável ao TSP simétrico euclidiano usado em muitos benchmarks.

## Citações-chave

> “The algorithm first solves the assignment problem for the matrix D, and then patches the cycles of the optimum assignment together to form a tour.”

> “The execution time of the patching algorithm is O(n3).”
