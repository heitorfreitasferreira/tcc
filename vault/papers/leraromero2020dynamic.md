---
title: "A Dynamic Programming Algorithm for the Time-Dependent Traveling Salesman Problem with Time Windows"
authors: [Lera-Romero, Gonzalo, Miranda-Bront, Juan José, Soulignac, Francisco J.]
year: 2020
doi: "10.1007/s10288-020-00450-x"
bibtex-key: leraromero2020dynamic
tags: [tdtsp, lower-bound, labeling, ng-path]
status: lido
rating: 4
---

## Resumo

Desenvolve um algoritmo de labeling (programação dinâmica) para o TDTSP com janelas de tempo. A principal contribuição é uma nova relaxação de estado-espaço específica para custos time-dependent, chamada **ti-relaxation**, que reduz significativamente o tempo de execução sem comprometer a qualidade dos bounds. Combina técnicas de ng-path relaxation (Baldacci et al.), relaxação Lagrangiana para fortalecer bounds, e algoritmos bidirecionais de labeling. O algoritmo supera métodos anteriores para TDTSPTW.

## Contribuições Principais

- Relaxação ti para TDTSP: versão time-dependent do ng-path relaxation
- Algoritmo de labeling bidirecional para TDTSP sem upper bounds apertados
- Regras de factibilidade para ngL-tour relaxation adaptadas ao caso time-dependent
- Solução de instâncias TDTSP com até 120 vértices

## Relevância para o TCC

A relaxação ti (time-dependent) e o ng-path relaxation são adaptáveis ao tensor 3D do TSP-SD-ATP. O labeling bidirecional permite computar lower bounds sem precisar de upper bounds de alta qualidade prévios. A abordagem de relaxação de estado-espaço (ng-path) é uma alternativa aos MDDs para obter bounds em problemas de sequenciamento.

## Métodos e Abordagens

- Programação dinâmica com labeling (forward + backward)
- ng-path relaxation: cada estado mantém apenas um subconjunto NG(i) de nós visitados
- ti-relaxation: simplificação dos tempos de chegada para TDTSP
- Relaxação Lagrangiana para fortalecer bounds (penalização de estados)
- Dynamic Neighbor Augmentation (DNA) para acelerar convergência

## Conexões

- [[kinable2017hybrid]] — abordagem alternativa com MDD + LP para TDTSP
- [[fischetti1992additive]] — additive bounding (combinável com labeling)
- [[heldkarp1970traveling]] — bound HK como referência de qualidade
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- ng-path relaxation é padrão em VRP para obter bounds apertados sem enumerar todos os estados
- A ti-relaxation é específica para custos que dependem do tempo — no TSP-SD-ATP sem tempo explícito, a versão padrão ng-path se aplica
- Labeling bidirecional reduz o espaço de busca sem exigir upper bounds apertados
- Relaxação Lagrangiana sobre o ng-path pode fortalecer o bound adicionalmente

## Citações-chave

> "We propose a new state-space relaxation specifically designed for the time-dependent context."

> "Extensive computational experiments show the effectiveness of the overall approach and the impact of the new relaxation."
