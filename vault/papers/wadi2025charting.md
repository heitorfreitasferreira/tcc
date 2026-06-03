---
title: "Charting New Routes: Comparing Swarm-Based Approaches to the Traveling Salesman Problem"
authors: [Wadi, Ali Hassan Ahmed, Umar, Shahla Uthman]
year: 2025
doi: "10.18280/ijcmem.130214"
bibtex-key: wadi2025charting
pdf: "papers/pdfs/wadi2025charting.pdf"
tags: [tsp, pso, aco, eho, swarm, comparison]
status: lido
rating: 3
---

## PDF

[[papers/pdfs/wadi2025charting.pdf]]

## Resumo

Compara Particle Swarm Optimization (PSO), Ant Colony Optimization (ACO) e Elephant Herding Optimization (EHO) na resolução do TSP com 5 a 150 cidades, usando Branch and Bound (BB) e Dynamic Programming (DP) como baselines exatos. EHO supera significativamente os demais em custo ótimo na maioria dos cenários, especialmente em instâncias maiores (100–150 cidades). ACO apresenta o melhor equilíbrio entre custo e tempo de execução. PSO mostra desempenho competitivo em velocidade mas qualidade de solução inferior. O estudo inclui análise detalhada de complexidade temporal, uso de CPU e memória.

## Contribuições Principais

- Inclusão do EHO como metaheurística menos convencional frente a PSO e ACO
- Análise abrangente de escalabilidade de 5 a 150 cidades com 5 algoritmos
- Métricas múltiplas: custo ótimo, tempo de execução, uso de CPU, complexidade temporal e espacial
- EHO identificado como melhor custo-benefício para instâncias grandes
- Baselines exatos (BB, DP) confirmam intratabilidade para >20 cidades

## Relevância para o TCC

Demonstra que metaheurísticas menos convencionais (EHO) podem superar as clássicas (PSO, ACO) — relevante para discussão do teorema "no free lunch" no TCC. A metodologia de comparação com múltiplas métricas (custo, tempo, CPU, complexidade) serve como template para os experimentos do repositório. Os baselines exatos (BB, DP) confirmam a necessidade de metaheurísticas para instâncias >20 cidades.

## Métodos e Abordagens

- Algoritmos swarm: PSO, ACO, EHO
- Baselines exatos: Branch and Bound (BB), Dynamic Programming (DP)
- Instâncias: 5, 10, 14, 19, 100, 150 cidades (coordenadas geradas randomicamente)
- Métricas: best cost, execution time, CPU time used, time complexity, memory consumption
- Implementação em MATLAB (HP EliteBook x360, i5-8350U, 16GB RAM)
- 1 run por configuração (sem repetições estatísticas)

## Conexões

- [[tsp]] — problema-alvo
- [[comparative-studies]] — área temática
- [[particle-swarm]] — PSO incluído e analisado
- [[ant-colony]] — ACO incluído e analisado
- [[hossain2024comparison]] — estudo comparativo similar com GA, ACO, SA vs. ABC, GWO, SSA
- [[almufti2025comparative]] — comparação de 9 metaheurísticas em TSP
- [[bio-inspired-optimization]]
- [[drone-routing]] — swarm-based routing para patrulha

## Notas e Insights

- EHO consistentemente supera PSO e ACO em qualidade de solução, especialmente em 100–150 cidades
- ACO oferece melhor compromisso custo-tempo; PSO é rápido mas de baixa qualidade
- BB e DP tornam-se inviáveis acima de 14–19 cidades (tempo exponencial)
- Limitação séria: apenas 1 run por configuração — sem análise estatística
- Coordenadas de cidades geradas randomicamente (não TSPLIB) — difícil comparar com literatura
- PSO teve desempenho fraco em TSP, consistente com a literatura (PSO é naturalmente contínuo)
- EHO como alternativa promissora para considerar em extensões futuras do repositório
- Estudo não reporta desvio padrão nem intervalo de confiança

## Citações-chave

> EHO surpasses the others in achieving lower optimal costs, particularly as the number of cities increases.

> Traditional algorithms such as BB and DP executed perfectly with small problem sizes (e.g., 10 or 5 cities), yielding accurate and efficient outcomes. However, their lack of computational efficiency was felt as problems turned increasingly complex.
