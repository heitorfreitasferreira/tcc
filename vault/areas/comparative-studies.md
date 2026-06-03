---
title: Estudos Comparativos de Metaheurísticas
tags: [metaheuristic, comparison, benchmark, experimental]
status: atualizado-pos-p7
updated: 2026-06-02
---

## Descrição

Área que agrupa estudos que comparam experimentalmente múltiplas metaheurísticas (GA, PSO, ACO, ABC, GWO, SA, etc.) aplicadas a problemas de otimização combinatória, especialmente o TSP. Estes estudos fornecem evidências empíricas sobre trade-offs entre qualidade de solução, tempo computacional, escalabilidade e robustez.

## Conexões

- [[TSP]] — problema-alvo da maioria dos benchmarks
- [[bio-inspired-optimization]] — metaheurísticas comparadas
- [[genetic-algorithms]] — método frequentemente incluído
- [[particle-swarm]] — método frequentemente incluído
- [[ant-colony]] — método frequentemente incluído
- [[experiment-pipeline]] — pipeline experimental do projeto
- [[ga]], [[pso]], [[aco]], [[bruteforce]] — implementações comparadas no projeto

## Gap na Literatura

Estudos comparativos existentes usam majoritariamente **TSP clássico (matriz de custo 2D simétrica)** em datasets TSPLIB. Na revisão organizada neste vault, não foram identificados estudos que comparem sistematicamente GA, PSO e ACO na variante **TSP-SD-ATP** (TSP com penalidades angulares dependentes de sequência), que é o objeto deste TCC.

Isto significa que:
- Não há baseline na literatura para a qualidade esperada de cada método no TSP-SD-ATP
- A influência da penalidade angular sobre o ranking dos métodos é desconhecida
- Este TCC aborda essa lacuna com 4638 summaries em 30 instâncias, incluindo 51 sementes por metaheurística, busca exaustiva em 18 instâncias e lower bound AP em todas as 30

Ver [[problem-formulation]] para a definição da variante e [[resultados]] para os resultados.

## Principais Insights

- Nenhum algoritmo domina universalmente ("No Free Lunch Theorem")
- ACO tende a produzir soluções de maior qualidade em TSPs grandes, com maior custo computacional
- GA é mais rápido e consistente em instâncias pequenas/médias
- PSO oferece bom equilíbrio entre qualidade e tempo (mas veja [[resultados]] para a variante TSP-SD-ATP)
- Estatística não-paramétrica pareada (Friedman/Nemenyi e Wilcoxon/Holm) é necessária para validar diferenças entre métodos neste TCC

## Referências na Base

- [[wu2020comparative]] — GA vs ACO vs PSO
- [[haroun2015performance]] — GA vs ACO
- [[almufti2025comparative]] — 9 metaheurísticas
- [[hossain2024comparison]] — old vs new algorithms
- [[chandra2022comparative]] — 8 metaheurísticas + ANOVA
- [[wadi2025charting]] — PSO vs ACO vs EHO
- [[alexander2020comparison]] — GA vs ACO
- [[halim2019combinatorial]] — 6 heurísticas
