---
title: "Routing Problem for Unmanned Aerial Vehicle Patrolling Missions -- A Progressive Hedging Algorithm"
authors: [Rajan, Sudarshan, Sundar, Kaarthik, Gautam, Natarajan]
year: 2022
doi: "10.1016/j.cor.2022.105702"
bibtex-key: rajan2022routing
pdf: "papers/pdfs/rajan2022routing.pdf"
tags: [drone routing]
status: lido
rating: 5
---

## Resumo

Apresenta um modelo de programação estocástica de dois estágios para roteirização de UAVs em missões de patrulha. Dado um conjunto de alvos e alvos suplementares, as decisões de primeiro estágio definem a sequência de visita; ao chegar em cada alvo, se a informação coletada for insuficiente, o UAV visita alvos suplementares. Resolvido via Progressive Hedging Algorithm.

## Contribuições Principais

- Modelo two-stage stochastic programming para patrulha com UAV
- Progressive Hedging Algorithm como método de solução
- Experimentos computacionais extensivos demonstrando eficácia

## Relevância para o TCC

Altamente relevante: aborda exatamente o cenário de patrulha com drones sobre pontos de interesse, com incerteza na coleta de informação. Alinha-se diretamente com a motivação do TCC.

## Métodos e Abordagens

- Two-stage stochastic programming
- Progressive Hedging Algorithm
- Simulação de cenários de patrulha

## Conexões

- [[murray2015flying]] — FSTSP (roteamento de drones, contexto de entrega)
- [[agatz2018optimization]] — TSP-D (entrega com drone)
- [[ahmed2024receding]] — planejamento de caminho para UAVs
- [[drone-routing]] [[drone-routing]]
- [[TSP]]

## Notas e Insights

- Preprint arXiv:2106.08379 (versão do autor) — PDF incluso no vault
- Modelo captura incerteza inerente a missões de patrulha/reconhecimento
- Abordagem estocástica é mais realista que TSP determinístico para cenários de drone
- Pode inspirar extensão estocástica para o rTSP do TCC (ex.: qualidade de informação variável)
- Publicado no *Computers & Operations Research* (vol. 142, 2022) — periódico de alto impacto em OR
- Pesquisa realizada no Los Alamos National Laboratory (LANL)
- PH algoritmo convergiu em todas as instâncias testadas sem problemas de convergência
- Extensões futuras sugeridas incluem cenário multi-UAV e multi-estágio

## Citações-chave

> "This paper presents a two-stage stochastic program to model a routing problem involving an Unmanned Aerial Vehicle (UAV) in the context of patrolling missions."

> "The progressive hedging algorithm has been tested on its effectiveness has been corroborated on a large set of test instances. On all the instances, it was observed that no convergence issues have been observed."
