---
title: "A Review of Metaheuristic Algorithms for Solving {TSP}-Based Scheduling Optimization Problems"
authors: [Toaza, Bladimir]
year: 2023
doi: "10.1016/j.asoc.2023.110908"
bibtex-key: toaza2023review
pdf: null
tags: [tsp metaheuristic review scheduling]
status: lido
rating: 4
---

## Resumo

Revisão bibliométrica sistemática de 120 metaheurísticas aplicadas a problemas de otimização de scheduling baseados em TSP, publicada na *Applied Soft Computing*. O estudo analisa tendências de publicação, impacto de citação e características dos algoritmos mais prevalentes. GA é o mais aplicado em número de publicações, mas ACO é o mais citado, sugerindo que o impacto científico do ACO supera sua adoção prática. A revisão identifica lacunas na comparação padronizada entre metaheurísticas e na aplicação a problemas dinâmicos de scheduling.

## Contribuições Principais

- Catálogo de 120 metaheurísticas com features descritivas e métricas de avaliação
- Análise de tendências temporais: GA domina frequência, ACO domina citações
- Tabulação comparativa de operadores, parâmetros e domínios de aplicação
- Identificação de lacunas: falta de benchmarks unificados e escassez de estudos em TSP dinâmico

## Relevância para o TCC

Fornece justificativa bibliométrica para a escolha de GA, PSO e ACO como metaheurísticas centrais na comparação do rTSP para patrulha com drones. Os dados de citação e adoção corroboram a seleção do tripé metodológico do projeto.

## Métodos e Abordagens

- Busca sistemática com automação via API em bases acadêmicas
- 120 metaheurísticas classificadas por tipo (populacional vs. trajetória), operadores e domínio
- Métricas: frequência de publicação, citações, features dos algoritmos
- Foco em problemas de scheduling com estrutura TSP

## Conexões

- [[TSP]]
- [[genetic-algorithms]]
- [[particle-swarm]]
- [[ant-colony]]
- [[comparative-studies]]
- [[rajwar2023exhaustive]] — survey complementar (taxonomia de metaheurísticas)
- [[holland1975adaptation]] — GA fundacional
- [[kennedy1995particle]] — PSO fundacional
- [[dorigo1996ant]] — ACO fundacional
- [[halim2019combinatorial]] — survey comparativo contemporâneo

## Notas e Insights

- Open Access (CC BY-NC-ND) na ScienceDirect: https://doi.org/10.1016/j.asoc.2023.110908 — PDF disponível na página do DOI mediante browser com JavaScript
- GA é o mais usado, ACO o mais citado — possível indicador de que ACO tem maior impacto científico mas menor adoção prática
- Revisão cobre literatura até 2023, ponto de partida útil para estado da arte
- Ênfase em scheduling; seria valioso estender a análise para TSP dinâmico e robótico
- A ausência de benchmarks padronizados é uma limitação metodológica relevante

## Citações-chave

> "GA is the most applied algorithm in publications, but ACO is the most cited one."

> "Metaheuristic algorithms can provide satisfactory solutions for complex combinatorial optimization problems at a reasonable computational cost."
