---
title: "A Comparative Study of Metaheuristics Methods for Solving Traveling Salesman Problem"
authors: [Chandra, Agung, Naro, Aulia]
year: 2022
doi: "10.57675/IMIST.PRSM/ijist-v6i2.211"
bibtex-key: chandra2022comparative
pdf: "papers/pdfs/chandra2022comparative.pdf"
tags: [tsp, ga, sa, ts, aco, pso, abc, comparison, statistics, metaheuristic]
status: lido
rating: 4
---

## PDF

[[papers/pdfs/chandra2022comparative.pdf]]

## Resumo

Compara 8 metaheurísticas — GA, SA, TS, ACO, PSO, ABC, EFOA e A3 (Artificial Atom Algorithm) — na resolução do TSP simétrico para 70 cidades reais na ilha de Java, Indonésia. O ABC (Artificial Bee Colony) obteve a melhor distância (2.447 km) e a melhor média (2.458 km). Os resultados foram submetidos a ANOVA e teste post-hoc de Tukey, que indicaram que 20 dos 28 pares de métodos apresentam diferenças estatisticamente significativas. PSO apresentou o pior desempenho, com distância média de 10.932 km.

## Contribuições Principais

- Aplicação de rigor estatístico (ANOVA + Tukey) para comparar 8 metaheurísticas no mesmo framework
- Utilização de caso real (70 cidades em Java) complementado por validação com TSPLIB (tsplib58)
- Ranking completo: ABC > SA > {A3, EFOA, GA} > ACO > TS > PSO
- Documentação detalhada de parâmetros e operadores de cada método

## Relevância para o TCC

Demonstra como ANOVA e Tukey podem validar diferenças entre metaheurísticas no TSP — metodologia estatística diretamente aplicável à análise dos experimentos do projeto com GA, PSO e ACO no rTSP. O ranking obtido (ABC superior a GA, PSO e ACO) fornece referência para expectativas de desempenho.

## Métodos e Abordagens

- 8 metaheurísticas implementadas em MATLAB 2015a no mesmo framework
- ANOVA de um fator + teste post-hoc Tukey-Kramer (α = 0,05)
- 70 cidades reais em Java (coordenadas geográficas) + instância TSPLIB (tsplib58)
- Métrica: distância euclidiana total (km)
- Hardware: Intel Core i5 7200U, 2.5 GHz, 32 bits
- Iterações: 10.000 para ABC e A3; parâmetros específicos para cada método

## Conexões

- [[tsp]] — problema-alvo
- [[comparative-studies]] — área temática
- [[genetic-algorithms]] — GA incluído
- [[particle-swarm]] — PSO incluído
- [[ant-colony]] — ACO incluído
- [[simulated-annealing]] — SA incluído
- [[tabu-search]] — TS incluído
- [[artificial-bee-colony]] — ABC (melhor método)
- [[alexander2020comparison]] — comparação GA vs. ACO
- [[halim2019combinatorial]] — survey comparativo de heurísticas TSP
- [[toaza2023review]] — revisão bibliométrica de metaheurísticas

## Notas e Insights

- ABC superou todos os métodos, incluindo GA, PSO e ACO — relevante para justificar a inclusão de ABC em estudos comparativos futuros
- PSO apresentou desempenho muito inferior (média 10.932 km vs. 2.447 km do ABC); possivelmente devido à dificuldade de adaptação do PSO discreto ao TSP
- GA e ACO tiveram desempenho intermediário comparável (GA: 2.727 km, ACO: 2.851 km)
- Apenas 8 dos 28 pares não apresentaram diferença significativa — a maioria dos métodos difere estatisticamente
- Os métodos no mesmo framework facilitam comparação justa, mas a sintonia de parâmetros pode favorecer alguns algoritmos

## Citações-chave

> "ABC optimization method has the shortest distance: 2,447 and the best average value: 2,458 in distance."

> "The ANOVA test indicates that the data of distances are not the same for every method of metaheuristics."
