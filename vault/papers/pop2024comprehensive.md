---
title: "A Comprehensive Survey on the Generalized Traveling Salesman Problem"
authors: [Pop, Petrică C.]
year: 2024
doi: "10.1016/j.ejor.2023.07.022"
bibtex-key: pop2024comprehensive
pdf: "papers/pdfs/pop2024.pdf"
tags: [tsp gtsp survey combinatorial-optimization drone-routing]
status: lido
rating: 5
---

## PDF

![[pop2024.pdf]]

## Resumo

Survey abrangente (invited review) publicado no *European Journal of Operational Research* sobre o Generalized Traveling Salesman Problem (GTSP). O GTSP estende o TSP clássico particionando vértices em clusters e exigindo que exatamente um vértice de cada cluster seja visitado no tour ótimo. O artigo cobre definição formal, variantes (CTSP, GTSP-TW, SGTSP, PCGTSP, CGTSP, FTSP), aplicações reais (roteirização postal, manufatura, image retrieval, entrega com drones), formulações matemáticas (programação inteira, fluxo em rede, abordagem local-global), algoritmos exatos (branch-and-cut, programação dinâmica), de aproximação, heurísticos e metaheurísticos. Inclui análise comparativa dos algoritmos state-of-the-art nos datasets GTSP_LIB, BAF_LIB, MOM_LIB e LARGE_LIB.

## Contribuições Principais

- Primeiro survey dedicado exclusivamente ao GTSP na literatura
- Cobertura completa: definição, 10+ variantes, 8 aplicações reais, 4 formulações matemáticas, taxonomia de solução
- Análise comparativa quantitativa de 5 algoritmos state-of-the-art (MA, LKH, LNS, Basic ILS, Refined ILS) em 4 bibliotecas de benchmark
- Identificação de problemas em aberto: algoritmos híbridos para GTSP em larga escala, GTSP dinâmico e estocástico
- Discussão de aplicações emergentes em logística drone e sistemas automatizados de armazenagem

## Relevância para o TCC

Altamente relevante: o GTSP modela diretamente o cenário de patrulha com drones, onde pontos de interesse são agrupados em clusters (regiões) e apenas um ponto por cluster precisa ser visitado. A seção de aplicações menciona explicitamente *drone-assisted parcel delivery service*. As formulações matemáticas e algoritmos comparados fornecem base teórica e prática para extensões do TCC.

## Métodos e Abordagens

- Survey com taxonomia de métodos: exatos, transformação, redução, aproximação, heurísticos, metaheurísticos
- Cobertura de programação inteira (subtour elimination, cutset, fluxo em rede, local-global)
- Algoritmos comparados: LNS (Smith & Imeson, 2017), LKH (Helsgaun, 2015), MA (Gutin & Karapetyan, 2010), Basic ILS e Refined ILS (Schmidt & Irnich, 2022)
- Datasets: GTSP_LIB (88 instâncias), BAF_LIB (56), MOM_LIB (45), LARGE_LIB (44)
- Métricas: taxa de sucesso (BKS), erro percentual médio, tempo de execução

## Conexões

- [[tsp]]
- [[gtsp]]
- [[drone-routing]]
- [[bock2025survey]] — survey de variantes TSP em warehousing (contemporâneo)
- [[lawler1985traveling]] — survey clássico do TSP
- [[applegate2006traveling]] — estudo computacional do TSP
- [[chandra2022comparative]] — estudo comparativo de metaheurísticas TSP
- [[genetic-algorithms]] — GA é mais usado para GTSP
- [[ant-colony]] — ACO é o segundo mais usado para GTSP

## Notas e Insights

- LNS (Large Neighborhood Search) apresentou o melhor desempenho global nos 4 datasets
- GA é o metaheurístico mais aplicado ao GTSP, seguido por ACO — consistente com os achados de [[toaza2023review]]
- Transformação GTSP→TSP (Noon & Bean, 1993) permite usar solvers TSP maduros (LKH) sem aumentar número de vértices
- Aplicações com drones mencionadas: clustered GTSP para *drone-assisted parcel delivery* (Baniasadi et al., 2020)
- GTSP é NP-difícil (contém TSP como caso particular com clusters unitários)
- Lacuna identificada: falta de framework unificado comparando as relaxações LP das diferentes formulações

## Citações-chave

> "The GTSP offers a promising way to model various real-world applications. Its hierarchical structure and versatility offers accurate models for several practical applications."

> "The LNS algorithm developed by Smith & Imeson (2017) showed the best overall performance."

> "The genetic algorithm (GA) approach is most used for the GTSP, while the second most used approach is the ant colony optimization (ACO) metaheuristic."
