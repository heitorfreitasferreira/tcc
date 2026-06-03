---
title: "Ant System: Optimization by a Colony of Cooperating Agents"
authors: [Dorigo, Marco]
year: 1996
doi: "10.1109/3477.484436"
bibtex-key: dorigo1996ant
tags: [aco metaheuristic]
status: lido
rating: 5
pdf: "papers/pdfs/dorigo1996ant.pdf"
---

## PDF

![[dorigo1996ant.pdf]]

## Resumo

O artigo fundacional do Ant System (AS) propõe um novo paradigma computacional inspirado no comportamento de colônias de formigas reais para resolver problemas de otimização combinatória. O AS combina três mecanismos: feedback positivo (depósito de feromônio), computação distribuída (múltiplos agentes) e heurística construtiva gulosa (visibilidade baseada na distância). Três variantes são apresentadas — ant-cycle, ant-density e ant-quantum — sendo a primeira superior por usar informação global (tour completo) em vez de local. O método é aplicado e validado no TSP, ATSP, QAP e Job-Shop Scheduling, demonstrando versatilidade e robustez.

## Contribuições Principais

- Definição do Ant System (AS), primeiro algoritmo de otimização por colônia de formigas
- Formulação matemática completa com trilha de feromônio, lista tabu e probabilidade de transição
- Três modelos de atualização de feromônio: ant-cycle (global), ant-density e ant-quantity (locais)
- Demonstração de versatilidade: aplicação a TSP, ATSP, QAP e JSP com modificações mínimas
- Comparação experimental com Simulated Annealing e Tabu Search no problema Oliver30
- Estratégia elitista para reforço do melhor tour encontrado

## Relevância para o TCC

O ACO é um dos três métodos implementados no TCC para comparação sobre TSP/rTSP. Este artigo estabelece a base conceitual do Ant System, que evoluiu para o Ant Colony System ([[dorigo1997ant]]) efetivamente usado no código de otimização. A capacidade do AS de lidar com ATSP é particularmente relevante para cenários de patrulha com drones, onde assimetrias podem surgir de vento, terreno ou limites operacionais.

## Métodos e Abordagens

- Agentes artificiais (formigas) constroem tours incrementalmente com probabilidade baseada em feromônio τ_{ij}(t) e visibilidade η_{ij}=1/d_{ij}
- Parâmetros α e β controlam peso relativo de feromônio vs. distância
- Evaporação de feromônio (fator ρ) evita convergência prematura
- Lista tabu para garantir tours factíveis (cada cidade visitada uma vez)
- Três variantes de atualização de feromônio: ant-cycle (Δτ = Q/L_k no fim do tour), ant-density (Δτ = Q a cada passo), ant-quantity (Δτ = Q/d_{ij} a cada passo)
- Experimentos em Oliver30 (30 cidades), Eilon50, Eilon75, grades T×T

## Conexões

- [[dorigo1997ant]] — Ant Colony System (ACS), evolução direta do AS
- [[wang2021ant]] — SOS-ACO, otimização de parâmetros do ACO
- [[kappagantula2025dpso]] — DPSO-Q usa Ant-Q (derivado do AS)
- [[kennedy1995particle]] — PSO (outra metaheurística bio-inspirada contemporânea)
- [[ant-colony]]
- [[rajwar2023exhaustive]] — classifica ACO como método canônico

## Notas e Insights

- O AS é computacionalmente caro (O(NC·n³)) comparado a heurísticas especializadas como Lin-Kernighan, mas é um framework geral
- A estratégia elitista melhora convergência, mas excesso de elitismo causa estagnação prematura
- O número ótimo de formigas é aproximadamente igual ao número de cidades (m ≈ n)
- A principal fraqueza do ant-cycle é o tempo de processamento, não a qualidade da solução
- O conceito de "estagnação" (todas as formigas seguem o mesmo tour) antecipa problemas de diversidade em metaheurísticas populacionais
- O artigo é denso e estabelece a nomenclatura que perdura na área (feromônio, evaporação, lista tabu)

## Citações-chave

> "Ant System is a viable new approach to stochastic combinatorial optimization. The main characteristics of this model are positive feedback, distributed computation, and the use of a constructive greedy heuristic."

> "The most interesting aspect of this autocatalytic process is that finding the shortest path around the obstacle seems to be an emergent property of the interaction between the obstacle shape and ants distributed behavior."
