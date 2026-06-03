---
title: "Combinatorial Optimization: Comparison of Heuristic Algorithms in Travelling Salesman Problem"
authors: [Halim, A. Hanif, Ismail, I.]
year: 2019
doi: "10.1007/s11831-017-9247-y"
bibtex-key: halim2019combinatorial
pdf: "papers/pdfs/halim2019combinatorial.pdf"
tags: [tsp, ga, sa, ts, aco, heuristic, comparison, survey, tpo]
status: lido
rating: 4
---

## PDF

[[papers/pdfs/halim2019combinatorial.pdf]]

## Resumo

Survey comparativo publicado na *Archives of Computational Methods in Engineering* (fator de impacto elevado) que analisa 6 heurísticas aplicadas ao TSP: Nearest Neighbor (construtivo), Genetic Algorithm (GA), Simulated Annealing (SA), Tabu Search (TS), Ant Colony Optimization (ACO) e Tree Physiology Optimization (TPO). O estudo avalia precisão das soluções, velocidade de convergência e tempo computacional utilizando instâncias da TSPLIB. O GA apresentou o melhor equilíbrio entre precisão e velocidade, enquanto o ACO demonstrou boa capacidade de encontrar soluções de qualidade mas com maior custo computacional. O TPO é introduzido como método bio-inspirado recente.

## Contribuições Principais

- Survey comparativo cobrindo 6 classes de algoritmos em um mesmo arcabouço experimental
- Inclusão do Tree Physiology Optimization (TPO), método bio-inspirado pouco explorado na literatura de TSP
- Métricas padronizadas de comparação: precisão (desvio da solução ótima), convergência (iterações) e tempo
- Experimentos em instâncias clássicas da TSPLIB permitindo reprodutibilidade

## Relevância para o TCC

Survey consolidado em periódico de alto impacto. Útil para fundamentar a escolha dos métodos — GA, ACO e SA — que também são utilizados no projeto. A inclusão de NN como baseline construtivo fornece referência inferior para comparação. A análise de convergência informa a escolha do número de iterações nos experimentos do TCC.

## Métodos e Abordagens

- Nearest Neighbor (NN): heurística construtiva gulosa
- GA: codificação de permutação, crossover OX, mutação por troca
- SA: cooling schedule geométrico, perturbação 2-opt
- TS: lista tabu, critério de aspiração
- ACO: modelo clássico (Dorigo), parâmetros α, β, ρ
- TPO: modelagem de competição por luz em árvores
- Benchmarks: instâncias selecionadas da TSPLIB
- Métricas: desvio percentual do ótimo, iterações até convergência, tempo de CPU

## Conexões

- [[tsp]] — problema-alvo
- [[comparative-studies]] — área temática
- [[genetic-algorithms]] — GA incluído
- [[ant-colony]] — ACO incluído
- [[simulated-annealing]] — SA incluído
- [[tabu-search]] — TS incluído
- [[alexander2020comparison]] — comparação específica GA vs. ACO
- [[chandra2022comparative]] — estudo comparativo com 8 metaheurísticas
- [[toaza2023review]] — revisão bibliométrica de metaheurísticas TSP

## Notas e Insights

- GA apresentou o melhor equilíbrio geral entre qualidade de solução e tempo — consistente com a adoção massiva de GA na literatura (confirmado por [[toaza2023review]])
- ACO tende a encontrar boas soluções mas com convergência mais lenta
- SA é competitivo em qualidade mas sensível ao resfriamento
- NN serve como baseline inferior rápido mas de baixa qualidade
- TPO mostrou potencial mas requer mais estudos para validação
- Publicado em 2019 — dados podem estar desatualizados para o estado da arte
- Survey útil como referência de linha de base, mas não inclui PSO (relevante para o TCC)

## Citações-chave

> "Meta-heuristic algorithms are an optimization algorithm that able to solve TSP problem towards a satisfactory solution."

> "Genetic algorithm achieves the best balance between solution quality and computational time among the tested methods."
