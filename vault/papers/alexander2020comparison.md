---
title: "The Comparison of Genetic Algorithm and Ant Colony Optimization in Completing Travelling Salesman Problem"
authors: [Alexander, Alexander, Sriwindono, Haris]
year: 2020
doi: "10.4108/eai.20-9-2019.2292121"
bibtex-key: alexander2020comparison
pdf: "papers/pdfs/alexander2020comparison.pdf"
tags: [tsp, ga, aco, comparison, tradeoff]
status: lido
rating: 2
---

## PDF

![[alexander2020comparison.pdf]]

## Resumo

Alexander e Sriwindono comparam [[genetic-algorithms]] e [[ant-colony]] para resolver o [[tsp]]. O estudo usa dados com 10, 20, ..., 100 cidades, executa cada configuração 10 vezes e compara distância percorrida e tempo de processamento. O GA usa seleção por roleta, Order Crossover e mutação reciprocal-exchange; o ACO é descrito como ACS, com parâmetros fixos. O resultado central é um trade-off: ACS encontra rotas mais curtas, mas GA tende a ser mais rápido em instâncias acima de 20 cidades.

## Contribuições Principais

- Comparação empírica direta entre GA e ACO/ACS no TSP.
- Mostra que qualidade de solução e tempo computacional podem favorecer algoritmos diferentes.
- Aplica um desenho experimental simples com tamanhos crescentes de instância.
- Explicita operadores de GA usados: roleta, Ordered Crossover e mutação reciprocal-exchange.
- Usa distância total e tempo médio como métricas de comparação.

## Relevância para o TCC

O artigo reforça uma premissa importante para os resultados do TCC: não basta comparar apenas makespan ou qualidade da solução; o custo computacional também muda o ranking prático dos métodos. Para patrulha com drones, uma rota ligeiramente melhor pode não compensar se o tempo de otimização for muito maior. A conclusão “ACO melhor em distância, GA melhor em tempo” dialoga diretamente com a análise de trade-off entre [[ga]], [[aco]], [[pso]] e [[bruteforce]].

## Métodos e Abordagens

- Instâncias com 10 a 100 cidades.
- Cada conjunto de dados foi testado 10 vezes.
- GA com população inicial de 10 cromossomos.
- Seleção por roulette wheel.
- Crossover: Order Crossover.
- Mutação: reciprocal-exchange.
- ACS com `alpha = 1`, `beta = 0.5` e `rho = 0.9`.

## Conexões

- [[haroun2015performance]]
- [[chandra2022comparative]]
- [[wu2020comparative]]
- [[comparative-studies]]
- [[genetic-algorithms]]
- [[ant-colony]]
- [[tsp]]

## Notas e Insights

- O artigo é útil para discutir trade-offs, mas a metodologia tem poucos detalhes sobre geração das cidades, controle de sementes e significância estatística.
- A população pequena do GA sugere cautela ao interpretar a diferença de qualidade entre GA e ACS.
- A conclusão final é equilibrada: não declara ACS universalmente melhor, pois GA vence em tempo em parte das instâncias.

## Citações-chave

> “the ant colony algorithm is able to find a shorter distance than the genetic algorithm, but the genetic algorithm shows a better speed of completion than the ant colony algorithm”

> “Generally, it cannot be ascertained that the ACS is better than GA and vice versa.”
