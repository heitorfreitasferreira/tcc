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

[[papers/pdfs/murray2015flying.pdf]]

## Resumo

Artigo seminal que define dois novos problemas de otimização de entregas com drones: o FSTSP (Flying Sidekick Traveling Salesman Problem) e o PDSTSP (Parallel Drone Scheduling TSP). No FSTSP, um único drone opera em sincronia com um caminhão de entregas, sendo lançado e recuperado pelo caminhão em pontos de encontro (clientes). No PDSTSP, múltiplos drones partem diretamente do depósito para atender clientes dentro do raio de voo, enquanto o caminhão atende os demais. Ambos os problemas são formalizados como MILP, e heurísticas eficientes (route-and-reassign) são propostas e validadas em instâncias de até 75 clientes. A análise numérica mostra que a velocidade do drone é mais crítica que a autonomia de voo para redução do tempo de entrega.

## Contribuições Principais

- Definição formal do FSTSP (drone lançado de caminhão) e PDSTSP (drones a partir do depósito)
- Formulações MILP para ambos os problemas
- Heurística route-and-reassign para FSTSP: constrói rota TSP, reassinala clientes ao drone iterativamente
- Heurística para PDSTSP: particiona clientes entre TSP (caminhão) e PMS (parallel machine scheduling, drones)
- Análise de tradeoff velocidade vs. autonomia: velocidade é fator mais determinante

## Relevância para o TCC

Fundação de toda a literatura de otimização de rotas com drones. O conceito de sincronização entre veículos e a modelagem de restrições de autonomia informam diretamente o rTSP do TCC. A estrutura de heurística route-and-reassign (resolver TSP base + ajustar atribuições) é análoga ao fluxo de trabalho do TCC (gerar rota TSP, aplicar metaheurísticas). A análise de tradeoff velocidade-autonomia pode ser transportada para o cenário de patrulha com drones.

## Métodos e Abordagens

- Programação Inteira Mista (MILP) resolvida com Gurobi
- Heurística FSTSP: TSP → cálculo de savings → reassinalmento iterativo de clientes ao drone
- Heurística PDSTSP: TSP (caminhão) + algoritmo LPT (parallel machine scheduling) para drones
- TSPs resolvidos com: solver IP, Clarke-Wright savings, nearest neighbor, sweep
- Instâncias: 10, 20, 25, 50, 75 clientes; 1-3 drones; velocidades 15-40 mph; autonomia 20-40 min

## Conexões

- [[agatz2018optimization]] — TSP-D, variante do problema definido aqui
- [[freitas2020vns]] — VNS aplicado ao FSTSP
- [[dellamico2021multiple]] — extensão para múltiplos drones
- [[dellamico2022exact]] — modelos exatos para o FSTSP
- [[rajan2022routing]] — roteamento de drones em patrulha (contexto diferente)
- [[lin1973effective]] — heurísticas LK usadas como base nos métodos
- [[garey1979computers]] — FSTSP é NP-difícil (generalização do TSP)
- [[drone-routing]]
- [[TSP]]
- [[bio-inspired-optimization]]

## Notas e Insights

- A heurística de savings para TSP (Clarke-Wright) mostrou o melhor equilíbrio qualidade/tempo no FSTSP
- A formulação MILP do FSTSP não conseguiu resolver nenhuma instância à otimalidade no limite de 30 min (72 instâncias de 10 clientes)
- PDSTSP com heurística Savings+LPT teve gap médio < 4% com tempo de frações de segundo
- Velocidade do drone importa mais que autonomia: drone mais rápido com menos bateria supera drone lento com mais bateria
- Limitações importantes: drone visita 1 cliente por sortie, drone não pode relançar do depósito, pontos de encontro apenas em clientes
- Futuras direções sugeridas: múltiplos drones, relaxar restrições de ponto de encontro, combinar FSTSP + PDSTSP

## Citações-chave

> "The primary contribution of this paper is to introduce a new variant of the traditional traveling salesman problem (TSP) that addresses the challenge of determining optimal customer assignments for a UAV working in tandem with a delivery truck."

> "Speed, even at the expense of endurance, is a critical factor in leveraging UAVs in last-mile delivery operations."
