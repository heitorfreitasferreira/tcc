---
title: 'Route Planning with Turn Restrictions: A Computational Experiment'
authors:
- Vanhove
- Stéphanie
- Fack
- Veerle
year: 2012
doi: 10.1016/j.orl.2012.06.001
bibtex_key: vanhove2012route
bibtex-key: vanhove2012route
pdf: papers/pdfs/vanhove2012route.pdf
tags:
- area/routing
- area/tsp-variants
- evidencia/referencia
- metodo/heuristic
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 3
type: paper
areas:
- routing
- tsp-variants
methods:
- heuristic
role: revisao
reading_status: resumo-lido
validation_status: nao-validado
---

## Resumo

Publicado na Operations Research Letters (2012). Compara experimentalmente três métodos para lidar com restrições de conversão (turn costs e proibições de curva) em algoritmos de caminho mais curto: (1) node splitting — divide nós com restrições em múltiplos nós, um por arco incidente; (2) line graph — transforma o grafo inteiro, onde cada nó do line graph representa um arco do grafo original e cada arco representa uma curva legal; (3) direct method — modifica o algoritmo de Dijkstra para rotular arcos em vez de nós, permitindo ciclos na solução. O estudo mostra que a densidade de restrições de curva (t_r e t_l) é o fator determinante de desempenho, e fornece diretrizes para escolha do algoritmo conforme as características do grafo.

## Contribuições Principais

- Primeiro estudo comparativo experimental de algoritmos de caminho mais curto com turn restrictions
- Análise teórica de complexidade de tempo e memória para os três métodos
- Diretriz prática: line graph é preferível para grafos com muitas restrições; direct method para poucas restrições
- Modelagem de turn costs como custos adicionais de transição entre arcos consecutivos

## Relevância para o TCC

Citado na Seção 2.2 como referência sobre custos dependentes de curva em roteamento. O conceito de que "o custo de passar por um nó não pode ser descrito apenas por uma aresta" é análogo à motivação do TSP-SD-ATP. Não trata de TSP — é sobre caminho mais curto com restrições de conversão, mas fornece o contexto conceitual para a dependência de sequência.

## Conexões

- [[tsp-variants]]
- [[winter2002modeling]] — modelagem de turn costs via pseudo-dual graph
