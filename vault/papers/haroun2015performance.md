---
title: "A Performance Comparison of GA and ACO Applied to TSP"
authors: [Haroun, Sabry Ahmed, Jamal, Benhra, Hicham, El Hassani]
year: 2015
doi: "10.5120/20674-3466"
bibtex_key: haroun2015performance
bibtex-key: haroun2015performance
pdf: "papers/pdfs/haroun2015performance.pdf"
tags: [tsp, ga, aco, comparison, metaheuristic]
status: lido-parcial
rating: 4
---

## PDF

![[haroun2015performance.pdf]]

## Resumo

Compara GA e ACO em 3 benchmarks TSPLIB (Berlin52, Eil76, A280) e 1 caso real (Casablanca40). GA é mais rápido e leve; ACO encontra rotas mais curtas em problemas grandes.

## Contribuições Principais

- Inclui caso real de TSP assimétrico (Casablanca40)
- Análise de erro e tempo por instância

## Relevância para o TCC

Comparação GA vs ACO com dados concretos de erro e runtime. Resultados corroboram trade-offs observados no projeto.

## Métodos e Abordagens

- GA com operadores clássicos
- ACO com evaporação de feromônio
- Benchmarks: Berlin52, Eil76, A280 + Casablanca40

## Conexões

- [[tsp]] — problema-alvo
- [[genetic-algorithms]] — método comparado
- [[ant-colony]] — método comparado
- [[comparative-studies]] — área temática

## Notas e Insights

- O estudo compara GA e ACO em 4 instâncias, mas a análise estatística é limitada — apenas médias e desvios, sem testes de significância (como Wilcoxon ou t-pareado)
- A inclusão de um caso real assimétrico (Casablanca40) é um diferencial em relação a benchmarks puramente sintéticos
- A conclusão de que ACO é melhor para problemas grandes e GA para problemas pequenos/médios é consistente com a literatura, mas o estudo não explora PSO, que é o terceiro método implementado no TCC
- Para o TCC, os resultados corroboram a necessidade de comparar GA e ACO no contexto do rTSP, onde o tamanho da instância pode influenciar a escolha do método
- A ausência de busca local (2-opt) em ambos os métodos limita a generalização dos resultados para implementações híbridas como as do TCC
- O artigo não detalha a parametrização usada (taxa de crossover, mutação, evaporação), dificultando a reprodução dos experimentos

## Citações-chave

> GA is excellent with small to medium size instances... ACO is greedier but provides better results, particularly with large problems.
