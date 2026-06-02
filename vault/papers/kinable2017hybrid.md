---
title: "Hybrid Optimization Methods for Time-Dependent Sequencing Problems"
authors: [Kinable, Joris, Cire, Andre A., van Hoeve, Willem-Jan]
year: 2017
doi: "10.1016/j.ejor.2016.11.035"
bibtex-key: kinable2017hybrid
tags: [tdtsp, lower-bound, decision-diagram, constraint-programming]
status: lido
rating: 5
---

## Resumo

Propõe métodos híbridos para problemas de sequenciamento onde os custos de setup entre tarefas dependem da posição relativa na ordenação — o Time-Dependent Traveling Salesman Problem (TDTSP). Combina programação por restrições (CP) com duas relaxações: (1) **Diagramas de Decisão Multivalorados (MDDs)** como relaxação discreta, (2) **Programação Linear (LP)** como relaxação contínua. O MDD é construído a partir do DP do problema com largura limitada, fornecendo lower bounds. O additive bounding (Fischetti & Toth) incorpora informação das relaxações LP no MDD. Resultados superam MILP e CP puros.

## Contribuições Principais

- Framework híbrido CP + MDD + LP para sequenciamento posição-dependente
- MDD com largura limitada W: merges de estados fornecem relaxação controlável
- Additive bounding usando custos reduzidos do LP para fortalecer o MDD
- Filtragem baseada em bound no MDD (podagem de arcos inviáveis)
- Solução de instâncias TDTSP significativamente maiores que o estado-da-arte anterior

## Relevância para o TCC

O TSP-SD-ATP com tensor 3D é um caso particular de TDTSP/sequenciamento posição-dependente. O MDD com largura limitada é a abordagem mais promissora para obter lower bounds que **não dependem de redução para 2D** — opera diretamente sobre a estrutura 3D. O estado (S, i, j) do DP captura perfeitamente a dependência do nó anterior. A implementação em Go puro é viável (sem solver LP externo — pode-se usar só o MDD sem o componente LP).

## Métodos e Abordagens

- Constraint Programming como framework principal
- MDD exato (largura total) vs relaxado (largura limitada W)
- Merge de estados: agrupa nós com mesmos (prev, curr) mas diferentes sets visitados
- Additive bounding para incorporar custos reduzidos do LP
- Filtragem de arcos no MDD por dominância e bound

## Conexões

- [[fischetti1992additive]] — additive bounding usado para combinar relaxações
- [[heldkarp1970traveling]] — bound HK como alternativa de relaxação contínua
- [[leraromero2020dynamic]] — abordagem alternativa com labeling para TDTSP
- [[aggarwal2000angular]] — angular-metric TSP, problema similar
- [[TSP]]
- [[lower-bounds]]

## Notas e Insights

- O MDD captura naturalmente a estrutura 3D: estados são (S, curr, prev), transições usam cost[prev][curr][next]
- O parâmetro W controla o trade-off precisão vs. desempenho
- O additive bounding fica mais forte com custos residuais do LP, mas só o MDD já fornece um bound válido
- Para o TSP-SD-ATP sem time windows, o MDD puro (sem CP) já pode fornecer lower bounds úteis

## Citações-chave

> "We propose novel optimization methods for sequencing problems in which the setup times between a pair of tasks depend on the relative position of the tasks in the ordering."

> "Our techniques are based on exact hybrid methods that combine linear programming, constraint programming, and multivalued decision diagrams."
