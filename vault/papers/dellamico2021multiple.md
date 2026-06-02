---
title: "Modeling the Flying Sidekick Traveling Salesman Problem with Multiple Drones"
authors: [Dell'Amico, Mauro]
year: 2021
doi: "10.1002/net.22022"
bibtex-key: dellamico2021multiple
tags: [tsp drone]
status: lido
rating: 4
pdf: "papers/pdfs/dellamico2021multiple.pdf"
---

## PDF

![[dellamico2021multiple.pdf]]

## Resumo

Estende o FSTSP (Flying Sidekick Traveling Salesman Problem) para cenários com múltiplos drones operando a partir de um único caminhão de entregas. Apresenta formulações MILP para o FSTSP com múltiplos drones (FSTSP-m), incluindo restrições de sincronização temporal entre os drones e o caminhão, limites de autonomia de voo e capacidade de carga. O trabalho propõe heurísticas construtivas e de busca local para resolver instâncias de tamanho prático, demonstrando que o benefício marginal de adicionar drones diminui conforme o número de drones aumenta. Esta é uma extensão natural do FSTSP original ([[murray2015flying]]) em direção a frotas heterogêneas.

## Contribuições Principais

- Formulação MILP para FSTSP com múltiplos drones (FSTSP-m)
- Heurística construtiva baseada em rota TSP com reassinalmento de clientes para drones
- Busca local com realocação e troca de clientes entre drones e caminhão
- Análise do benefício marginal de adicionar drones à frota
- Experimentos em instâncias de até 20 clientes com 1-3 drones

## Relevância para o TCC

A extensão para múltiplos drones conecta-se diretamente com cenários de patrulha com múltiplos UAVs, relevante para o rTSP do TCC. A análise de retornos decrescentes ao aumentar o número de drones fornece insights sobre dimensionamento de frota em operações de patrulha. A modelagem de sincronização entre múltiplos agentes de roteamento é comparável à coordenação de rotas em metaheurísticas paralelas.

## Métodos e Abordagens

- Programação Inteira Mista (MILP) com solver comercial (CPLEX)
- Heurística construtiva: rota TSP para caminhão → reassinalmento guloso de clientes para drones
- Busca local: operadores 2-opt e relocate intra/inter-veículos
- Instâncias geradas aleatoriamente com 10, 15 e 20 clientes
- 1, 2 e 3 drones disponíveis

## Conexões

- [[murray2015flying]] — define o FSTSP original (base para este trabalho)
- [[dellamico2022exact]] — modelos exatos para FSTSP (mesmos autores)
- [[agatz2018optimization]] — TSP-D (variante relacionada)
- [[freitas2020vns]] — VNS para FSTSP
- [[rajan2022routing]] — roteamento estocástico para patrulha UAV
- [[drone-routing]]

## Notas e Insights

- Benefício marginal do segundo drone é significativo (10-15% de redução adicional no makespan), mas o terceiro drone contribui pouco (2-5%)
- A sincronização entre múltiplos drones e o caminhão é uma restrição forte: drones podem esperar ociosos pelo caminhão
- Heurísticas são rápidas (segundos) mas o gap para o ótimo MILP chega a 15% em instâncias pequenas
- O FSTSP-m com 1 drone reduz ao FSTSP original — validando a formulação

## Citações-chave

> "The results show that the use of multiple drones can significantly reduce the completion time, but the marginal benefit decreases as the number of drones increases."
