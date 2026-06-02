---
title: "Edge Assembly Crossover for the Traveling Salesman Problem"
authors: [Nagata, Yuichi]
year: 2006
doi: ""
bibtex-key: nagata2006eax
tags: [ga tsp metaheuristic]
status: lido
rating: 5
---

## Resumo

Edge Assembly Crossover (EAX) é um operador de crossover para GAs aplicados ao TSP. Diferente de operadores tradicionais (PMX, OX, CX), o EAX constrói ciclos (AB-cycles) a partir das arestas de duas soluções pai e os rearranja para gerar filhos de alta qualidade. O GA+EAX é considerado estado-da-arte para TSP, competindo com LKH. O algoritmo alterna entre otimização local (estágio I: EAX sem mutação) e global (estágio II: EAX + mutação/restart). Variantes recentes (2025) melhoram a eficiência eliminando necessidade de reparo.

## Contribuições Principais

- Operador EAX que explora a estrutura de arestas do TSP de forma mais eficaz que PMX/OX/ERX
- GA+EAX atinge qualidade estado-da-arte, superando LKH em diversas instâncias
- Estratégia de dois estágios: exploração local (estágio I) alternada com global (estágio II)
- Código aberto e amplamente utilizado como baseline em competições TSP

## Relevância para o TCC

EAX representa o estado-da-arte em crossover para TSP. Embora complexo para implementar do zero, serve como benchmark superior para comparar com operadores mais simples (PMX, OX). A discussão de EAX na fundamentação teórica mostra conhecimento do estado-da-arte.

## Métodos e Abordagens

- AB-cycles: ciclos alternando arestas de cada pai (similar a operação XOR em grafo)
- Estratégia de dois estágios (estágio I: local, estágio II: global com restart)
- Busca local: 2-opt, LK opcional
- Hibridizado com restarts quando a população converge

## Conexões

- [[larranaga1999ga]] — survey de operadores de crossover (contexto histórico)
- [[potvin1996ga]] — categorização de crossover (edge-preserving)
- [[oliver1987crossover]] — PMX, OX, CX (operadores anteriores)
- [[lin1973effective]] — LKH, heurística concorrente estado-da-arte
- [[goldberg1989genetic]] — GA textbook
- [[genetic-algorithms]]
