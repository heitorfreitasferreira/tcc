---
title: "Genetic Algorithms for the Traveling Salesman Problem"
authors: [Potvin, Jean-Yves]
year: 1996
doi: "10.1007/BF02125403"
bibtex-key: potvin1996ga
pdf: "papers/pdfs/potvin1996ga.pdf"
tags: [ga tsp survey]
status: lido
rating: 4
---

## PDF

[[papers/pdfs/potvin1996ga.pdf]]

## Resumo

This paper surveys the genetic algorithm (GA) approach for solving the traveling salesman problem (TSP). It categorizes crossover operators into three groups based on their emphasis on relative order, absolute position, or edge preservation. Key findings include: edge-preserving crossover outperforms other crossover operators; local hill-climbing is crucial for good performance; separation of the population into subpopulations helps prevent premature convergence; larger populations correspond to better solutions. The paper also prominently features matrix-based encoding and mentions that parallel GAs will greatly improve solutions in the future.

**Categorias de crossover:** (1) Ordem relativa (e.g., OX1), (2) Posição absoluta (e.g., PMX), (3) Preservação de arestas (e.g., ERX).

## Contribuições Principais

- Taxonomia dos operadores de crossover para TSP em 3 categorias
- Evidência empírica de que edge-preserving crossover (ERX) supera as demais categorias
- Demonstração da importância de hibridização com busca local (hill-climbing)
- Discussão sobre mitigação de convergência prematura via subpopulações

## Relevância para o TCC

Útil para justificar a escolha de edge-preserving crossover e a importância de incluir busca local (e.g., 2-opt) no GA. A taxonomia dos operadores é útil para a seção de fundamentação teórica.

## Métodos e Abordagens

- Crossover: PMX, OX, CX, ERX (categorizados por ênfase)
- Busca local: hill-climbing (2-opt)
- Estratégias populacionais: subpopulações para diversidade
- Encoding: path representation (permutations) e matrix-based

## Conexões

- [[larranaga1999ga]] — review mais abrangente (1999)
- [[oliver1987crossover]] — crossover operators
- [[lin1973effective]] — busca local LKH
- [[nagata2006eax]] — EAX (edge-preserving crossover, estado-da-arte)
- [[genetic-algorithms]]
- [[tsp]]
