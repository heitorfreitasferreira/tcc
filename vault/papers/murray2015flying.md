---
title: "The Flying Sidekick Traveling Salesman Problem: Optimization of Drone-Assisted Parcel Delivery"
authors: [Murray, Chase C.]
year: 2015
doi: "10.1016/j.trc.2015.03.005"
bibtex-key: murray2015flying
tags: [tsp drone metaheuristic]
status: lido
rating: 5
pdf: "papers/pdfs/murray2015flying.pdf"
---

## PDF

![[murray2015flying.pdf]]

## Resumo

Murray e Chu introduzem o Flying Sidekick Traveling Salesman Problem (FSTSP), variante do TSP em que um caminhão de entrega opera em conjunto com um drone. O caminhão pode lançar o drone durante sua rota; o drone entrega a um cliente elegível e depois retorna ao caminhão ou ao depósito. O objetivo é minimizar o tempo total até que todos os clientes sejam atendidos e ambos os veículos retornem. O artigo também define o Parallel Drone Scheduling TSP (PDSTSP), cenário em que uma frota de drones sai diretamente do centro de distribuição enquanto um caminhão atende os demais clientes.

## Contribuições Principais

- Introduz formalmente o FSTSP como variante do TSP com caminhão e drone sincronizados.
- Introduz o PDSTSP para casos em que drones atendem clientes diretamente a partir do depósito.
- Formula modelos MILP para os dois problemas.
- Propõe heurística route and re-assign para FSTSP.
- Analisa velocidade versus endurance e mostra que nem sempre é ótimo atribuir todos os clientes elegíveis aos drones.

## Relevância para o TCC

Este é um dos artigos mais importantes para conectar [[tsp]] e drones. Embora o TCC trate patrulha de drones, e não entrega de pacotes, o artigo fornece motivação logística para variantes de TSP com veículos aéreos, restrições operacionais e objetivo de minimizar tempo total. Ele ajuda a posicionar o modelo atual como uma base controlada antes de avançar para variantes com sincronização, endurance, elegibilidade de pontos e paralelismo.

## Métodos e Abordagens

- Formulação MILP para FSTSP.
- Formulação MILP para PDSTSP.
- Heurística FSTSP baseada em rota inicial de TSP e realocação iterativa de clientes para drone quando há economia positiva.
- Heurísticas de TSP testadas: IP, Clarke-Wright savings, nearest neighbor e sweep.
- Para PDSTSP, combinação de rota TSP para caminhão com escalonamento em máquinas paralelas idênticas para drones.

## Conexões

- [[agatz2018optimization]]
- [[dellamico2021multiple]]
- [[dellamico2022exact]]
- [[freitas2020vns]]
- [[rajan2022routing]]
- [[tsp]]
- [[drone-routing]]

## Notas e Insights

- A formulação FSTSP é mais próxima de sincronização caminhão-drone; o PDSTSP é mais próximo de paralelização a partir do depósito.
- O artigo mostra que velocidade pode ser mais importante que endurance, especialmente quando múltiplos drones estão disponíveis.
- A heurística FSTSP depende bastante da qualidade da rota TSP inicial; isso conecta diretamente com o TCC.
- Para patrulha, a analogia mais forte não é “entregar pacote”, mas “usar autonomia aérea para reduzir makespan sob restrições operacionais”.

## Citações-chave

> “The primary contribution of this paper is to introduce a new variant of the traditional traveling salesman problem (TSP) that addresses the challenge of determining optimal customer assignments for a UAV working in tandem with a delivery truck.”

> “As an extension of the TSP, it is clear that the FSTSP is NP-hard.”
