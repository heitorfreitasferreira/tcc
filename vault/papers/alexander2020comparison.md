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

[[papers/pdfs/alexander2020comparison.pdf]]

## Resumo

Compara Genetic Algorithm (GA) e Ant Colony Optimization (ACO) na resolução do TSP simétrico com tamanhos de 10 a 100 cidades (aleatórias). O ACO encontra rotas mais curtas que o GA em todas as configurações testadas, enquanto o GA apresenta tempo de execução significativamente menor. O GA utiliza ordered crossover e mutação recíproca como operadores de permutação. O estudo quantifica o trade-off clássico entre qualidade da solução e velocidade computacional.

## Contribuições Principais

- Avaliação sistemática do trade-off qualidade vs. tempo em 10 cenários (10–100 cidades)
- GA com operadores de permutação adaptados para TSP (ordered crossover + reciprocal mutation)
- Confirmação quantitativa da superioridade do ACO em qualidade de rota e do GA em velocidade

## Relevância para o TCC

Corrobora diretamente o trade-off qualidade vs. velocidade entre GA e ACO, que é central para a escolha do método de otimização no cenário de patrulha com drones. Em aplicações com restrições de tempo real, GA pode ser preferível; em planejamento off-line, ACO pode oferecer rotas mais curtas.

## Métodos e Abordagens

- GA: ordered crossover (OX), reciprocal mutation, seleção por torneio
- ACO: modelo clássico com parâmetros α, β, ρ, Q
- Cidades geradas aleatoriamente em 10 tamanhos de problema (10 a 100)
- Métricas: distância total do tour e tempo de execução
- Implementação em MATLAB

## Conexões

- [[TSP]] — problema-alvo
- [[comparative-studies]] — área temática
- [[genetic-algorithms]] — GA comparado
- [[ant-colony]] — ACO comparado
- [[chandra2022comparative]] — estudo com 8 metaheurísticas (inclui GA, ACO, PSO)
- [[halim2019combinatorial]] — survey comparativo com 6 heurísticas
- [[toaza2023review]] — evidência bibliométrica: ACO mais citado, GA mais aplicado

## Notas e Insights

- Estudo simples e limitado (apenas 2 métodos, cidades aleatórias sem benchmark padronizado)
- Resultado consistente com a literatura: ACO geralmente produz melhores rotas que GA em TSP
- GA é 2–5x mais rápido que ACO dependendo do tamanho do problema
- O trade-off identificado é relevante para decidir entre otimização on-line (GA) e off-line (ACO) em sistemas de drones
- Falta análise estatística (não usa ANOVA ou testes de significância como [[chandra2022comparative]])

## Citações-chave

> "Ant colony algorithm is able to find a shorter distance than the genetic algorithm, but the genetic algorithm shows a better speed of completion."

> "GA can complete the TSP faster than ACO, but ACO can find a shorter route."
