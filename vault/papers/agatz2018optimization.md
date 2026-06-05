---
title: Optimization Approaches for the Traveling Salesman Problem with Drone
authors:
- Agatz
- Niels
year: 2018
doi: 10.1287/trsc.2017.0791
bibtex_key: agatz2018optimization
bibtex-key: agatz2018optimization
tags:
- area/drone-routing
- area/tsp
- evidencia/referencia
- metodo/metaheuristic
- status/lido
- tipo/paper
status: lido
rating: 5
pdf: papers/pdfs/agatz2018optimization.pdf
type: paper
areas:
- drone-routing
methods:
- metaheuristic
role: revisao
reading_status: lido
validation_status: nao-validado
---

## PDF

![[agatz2018optimization.pdf]]

## Resumo

Agatz, Bouman e Schmidt estudam o Traveling Salesman Problem with Drone (TSP-D), formulação em que um caminhão e um drone colaboram para atender clientes. O caminhão tem maior capacidade e alcance; o drone é mais rápido, mas atende um cliente por operação e precisa sincronizar com o caminhão. O objetivo é minimizar o tempo total da rota colaborativa. O artigo propõe um modelo de programação inteira e heurísticas route first-cluster second: primeiro se constrói uma rota de caminhão que visita todos os nós; depois a rota é particionada em operações de caminhão e drone.

## Contribuições Principais

- Formula um modelo IP para o TSP-D.
- Propõe heurísticas rápidas route first-cluster second.
- Apresenta algoritmo de programação dinâmica para encontrar a melhor atribuição caminhão-drone para uma sequência fixa.
- Prova garantias de aproximação baseadas em TSP e MST.
- Compara heurísticas contra soluções ótimas em instâncias pequenas.

## Relevância para o TCC

Este artigo conecta diretamente o TSP clássico ao cenário de drones preservando uma estrutura algorítmica analisável. Para o TCC, ele fornece ponte entre o modelo simplificado TSP/rTSP e aplicações mais realistas de drones de patrulha ou entrega. A ideia route first-cluster second também é útil como interpretação: primeiro encontra-se uma boa sequência de visitação, depois adapta-se a rota para restrições operacionais.

## Métodos e Abordagens

- Modelo em grafo com depósito e clientes.
- Tempos de viagem distintos para caminhão e drone.
- Objetivo de minimizar o tempo total, considerando esperas por sincronização.
- Construção inicial por TSP ótimo via Concorde ou por MST.
- Particionamento guloso com operações MakeFly, PushLeft e PushRight.
- Particionamento exato por programação dinâmica em `O(n^3)` para uma rota fixa.

## Conexões

- [[murray2015flying]]
- [[dellamico2021multiple]]
- [[dellamico2022exact]]
- [[freitas2020vns]]
- [[rajan2022routing]]
- [[lin1973effective]]
- [[applegate2006traveling]]
- [[drone-routing]]

## Notas e Insights

- Uma boa rota TSP não é necessariamente uma boa rota TSP-D, porque o caminhão pode precisar facilitar bons voos do drone.
- O particionamento exato supera o guloso, mas cresce em custo computacional.
- A busca local é essencial para explorar bem o benefício do drone; sem adaptar a sequência, há mais espera e menor paralelismo.
- O artigo é boa base para explicar por que o TCC começa com TSP/rTSP: mesmo uma extensão com um drone acoplado já exige modelos e heurísticas mais complexos.

## Citações-chave

> “An innovative last-mile delivery concept in which a truck collaborates with a drone to make deliveries gives rise to a new variant of the traveling salesman problem (TSP) that we call the TSP with drone.”

> “The exact partitioning algorithm has minimal time duration among all solutions (R, D) to the TSP-D with R and D both subsequences of R1. This solution can be found in time O(n3).”
