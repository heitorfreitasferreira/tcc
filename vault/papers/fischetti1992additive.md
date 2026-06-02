---
title: "An Additive Bounding Procedure for the Asymmetric Traveling Salesman Problem"
authors: [Fischetti, Matteo, Toth, Paolo]
year: 1992
doi: "10.1007/BF01585642"
bibtex-key: fischetti1992additive
tags: [atsp, lower-bound, additive-bounding]
status: lido
rating: 5
---

## Resumo

Propõe um procedimento de bounding aditivo (additive bounding) que combina múltiplas relaxações para o ATSP (Asymmetric TSP) para produzir um lower bound mais forte que qualquer relaxação isolada. A ideia é aplicar sequencialmente várias relaxações (assignment problem, r-arborescência, r-antiarborescência), onde cada relaxação opera sobre custos residuais da anterior. O bound resultante não é inferior ao mais forte dos bounds componentes e, em muitos casos, é significativamente superior.

## Contribuições Principais

- Framework de additive bounding aplicável a qualquer problema de otimização combinatória
- Combinação de assignment problem (AP) + r-arborescência (r-SAP) + r-antiarborescência (r-SAAP) + versões direcionadas (r-SADP, r-SAADP)
- Demonstração de que o bound aditivo domina cada bound individual
- Algoritmo B&B baseado neste bound resolve instâncias ATSP de até 2.000 vértices

## Relevância para o TCC

O additive bounding é a abordagem mais robusta para ATSP na literatura. Para o TSP-SD-ATP, que é essencialmente um ATSP com custo dependente do nó anterior, a redução `c'[j][k] = min_i cost[i][j][k]` permite aplicar o additive bounding (AP + arborescências) sobre a matriz 2D reduzida. Alternativamente, pode-se adaptar o AP diretamente sobre o tensor 3D.

## Métodos e Abordagens

- Assignment Problem (AP) — O(n³) via algoritmo húngaro
- r-arborescência (SAP) — MST dirigida O(n²) (Tarjan 1977)
- r-antiarborescência (SAAP) — análogo invertido
- Combinação aditiva: custos residuais c^(k+1) = c^(k) - π^(k)
- Algoritmo Karp (1979) para converter solução AP em tour via patching

## Conexões

- [[balas1985branch]] — B&B para ATSP com relaxação AP
- [[heldkarp1970traveling]] — arborescência como generalização de 1-tree para ATSP
- [[karp1979patching]] — patching heuristic para converter AP em TSP
- [[lawler1985traveling]] — survey que descreve relaxações do ATSP
- [[TSP]]
- [[lower-bounds]]

## Notas e Insights

- O additive bounding permite que cada relaxação "especialista" ataque uma parcela diferente do gap
- A ordem das relaxações importa: aplicar AP primeiro, depois as arborescências
- O bound final é tipicamente 1-3% acima do bound AP isolado
- O código CDT (Carpaneto, Dell'Amico, Toth) implementa esta abordagem e resolve ATSPs grandes rapidamente

## Citações-chave

> "The additive approach aims at computing a lower bound by combining different bounding procedures, each exploiting a different substructure of the problem."

> "The final bound is guaranteed to be not worse than the best bounding procedure among those used."
