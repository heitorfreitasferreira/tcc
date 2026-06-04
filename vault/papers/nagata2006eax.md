---
title: "A Powerful Genetic Algorithm Using Edge Assembly Crossover for the Traveling Salesman Problem"
authors: [Nagata, Yuichi, Kobayashi, Shigenobu]
year: 2013
doi: "10.1287/ijoc.1120.0506"
bibtex_key: nagata2006eax
bibtex-key: nagata2006eax
pdf: "papers/pdfs/nagata2006eax.pdf"
tags: [ga, tsp, crossover, eax, edge-assembly]
status: resumo-lido
rating: 4
---

## Resumo

Publicado no INFORMS Journal on Computing (2013). Propõe um GA com Edge Assembly Crossover (EAX) para o TSP, com três melhorias substanciais: (i) localização do EAX com implementação eficiente, (ii) busca local integrada ao EAX para determinar boas combinações de blocos das soluções pais, e (iii) modelo de seleção inovador para manter diversidade populacional com custo computacional desprezível. Resultados experimentais em benchmarks TSP mostram que o GA proposto supera heurísticas estado-da-arte, incluindo variantes do Lin--Kernighan, em instâncias de até 200.000 cidades — alcançando desempenho de ponta sem usar LK.

## Contribuições Principais

- EAX como operador de crossover baseado em arestas que preserva e combina edges dos pais
- Mecanismo de diversidade populacional que mantém exploração sem custo adicional significativo
- Supera LK-based heuristics em instâncias grandes (até 200k cidades)
- Referência canônica para crossover em GA para TSP

## Relevância para o TCC

Citado na Seção 2.5 (GA) como referência de alta qualidade para recombinação baseada em arestas. O EAX é mencionado como contraste ao OX (Order Crossover) implementado neste trabalho. Referência importante para fundamentar a discussão sobre operadores de crossover no TSP.

## Conexões

- [[genetic-algorithms]]
- [[tsp]]
- [[potvin1996ga]] — revisão de operadores de crossover
- [[larranaga1999ga]] — sistematização de representações GA para TSP
