---
title: "Comparison of New and Old Optimization Algorithms for Traveling Salesman Problem on Small, Medium, and Large-Scale Benchmark Instances"
authors: [Hossain, Md Al Amin, Yılmaz Acar, Züleyha]
year: 2024
doi: "10.17798/bitlisfen.1380086"
bibtex-key: hossain2024comparison
pdf: "papers/pdfs/hossain2024.pdf"
tags: [tsp, ga, aco, sa, abc, gwo, ssa, comparison, benchmark]
status: lido
rating: 4
---

## PDF

![[hossain2024.pdf]]

## Resumo

Compara algoritmos clássicos (GA, ACO, SA) com algoritmos recentes (ABC, GWO, SSA) em 7 instâncias TSPLIB de 14 a 1000 nós, classificadas em small (14–100), medium (100–500) e large (500–1000). Utiliza teste t para significância estatística entre grupos. Resultados indicam que GA apresenta desempenho superior em instâncias médias, enquanto em instâncias pequenas e grandes não há diferença estatisticamente significativa entre grupos. O estudo conclui que algoritmos "novos" não superam necessariamente os clássicos em todas as escalas.

## Contribuições Principais

- Teste t de Student para significância estatística entre grupos old vs. new
- Classificação por porte: small (14–100), medium (100–500), large (500–1000)
- Evidência contrária à suposição de que algoritmos mais novos são sempre melhores
- GA identificado como significativamente superior em instâncias médias

## Relevância para o TCC

Fornece evidência estatística fundamental para o TCC: GA e ACO (implementados no projeto) permanecem competitivos frente a métodos mais recentes (GWO, SSA). A classificação por porte de instância é diretamente aplicável ao design experimental do repositório. O uso de teste t valida a abordagem de análise estatística que pode ser replicada nos experimentos do TCC.

## Métodos e Abordagens

- Algoritmos "old": GA, ACO, SA
- Algoritmos "new": ABC, GWO, SSA
- 7 instâncias TSPLIB: small (14–100 nós), medium (100–500), large (500–1000)
- Teste t de Student para comparação entre grupos
- Métricas: tour length médio, desvio padrão, significância estatística
- 30 runs independentes por algoritmo-instância

## Conexões

- [[TSP]] — problema-alvo
- [[comparative-studies]] — área temática
- [[genetic-algorithms]] — GA incluído e melhor em médias
- [[ant-colony]] — ACO incluído
- [[almufti2025comparative]] — estudo similar com 9 metaheurísticas
- [[wadi2025charting]] — comparação swarm-based com EHO, ACO, PSO
- [[bio-inspired-optimization]]

## Notas e Insights

- GA surpreendentemente competitivo em instâncias médias — relevante para justificar inclusão de GA no projeto
- Ausência de diferença significativa em small/large sugere que escolha do algoritmo importa mais para tamanhos intermediários
- Estudo bem desenhado com separação por porte e teste estatístico
- Limitação: apenas 6 algoritmos; PSO não foi incluído (relevante para o TCC)
- A classificação small/medium/large é adotável diretamente nos experimentos do repositório
- Sugere que investir em hibridização (ex.: GA+ACO) pode ser mais frutífero que buscar algoritmos mais novos

## Citações-chave

> GA algorithm showed significantly better performance compared to the ACO, SA, ABC, SSA, and GWO algorithms in medium instances.

> New algorithms do not necessarily outperform classical ones in all problem scales — statistical significance depends heavily on instance size category.
