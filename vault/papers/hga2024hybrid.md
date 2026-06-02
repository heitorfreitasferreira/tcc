---
title: "An Efficient Hybrid Algorithm with Novel Inver-over Operator and Ant Colony Optimization for Traveling Salesman Problem"
authors: [Singh, D. R., Singh, M. K., Chaurasia, S. N.]
year: 2024
doi: ""
bibtex-key: hga2024hybrid
tags: [ga aco tsp metaheuristic hybrid]
status: lido
rating: 4
---

## Resumo

Este artigo propõe um algoritmo híbrido GA-ACO em dois estágios para TSP euclidiano. Estágio 1: ACO gera população inicial + operador Inver-over refina. Estágio 2: GA com crossover personalizado e 2-opt refinam em direção à otimalidade global. O híbrido supera métodos recentes em qualidade, combinando exploração global do GA com explotação local do ACO + 2-opt.

## Contribuições Principais

- Framework híbrido GA-ACO em dois estágios que explora sinergia entre as metaheurísticas
- Operador Inver-over adaptado para refinar soluções ACO
- Crossover personalizado específico para TSP
- Resultados superiores a GA puro, ACO puro e híbridos anteriores em benchmarks TSPLIB

## Relevância para o TCC

Demonstra que GA e ACO podem ser combinados com benefício mútuo, o que é relevante para a discussão de resultados comparativos. Se o TCC testar GA e ACO separadamente, este artigo oferece uma perspectiva de hibridização como possível extensão.

## Métodos e Abordagens

- Estágio 1: ACO gera população inicial diversa
- Operador Inver-over (inversão adaptativa) refina soluções ACO
- Estágio 2: GA com crossover personalizado (similar a SCX)
- Busca local 2-opt como pós-processamento
- Benchmarks TSPLIB (eil51, berlin52, st70, pr76, kroA100, etc.)

## Conexões

- [[dorigo1996ant]] — ACO, base do primeiro estágio
- [[dorigo1997ant]] — ACO para TSP
- [[holland1975adaptation]] — GA, base do segundo estágio
- [[goldberg1989genetic]] — GA textbook
- [[larranaga1999ga]] — survey de operadores GA para TSP
- [[lin1973effective]] — 2-opt usado como busca local
- [[ant-colony]]
- [[genetic-algorithms]]
