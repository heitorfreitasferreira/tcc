---
tags:
- evidencia/metodologia
- tipo/area
created: 2026-06-02
updated: 2026-06-02
type: area
---

# Otimização Bio-Inspirada

## Definição
Classe de algoritmos de otimização que se inspiram em fenômenos biológicos/naturais para resolver problemas computacionais complexos.

## Métodos Abordados no TCC

| Método | Inspiração | Criador(es) | Ano |
|--------|-----------|-------------|-----|
| [[genetic-algorithms]] | Seleção natural | Holland | 1975 |
| [[particle-swarm]] | Comportamento de pássaros | Kennedy, Eberhart | 1995 |
| [[ant-colony]] | Comportamento de formigas | Dorigo et al. | 1996 |

## Problema-alvo
[[tsp]] e [[drone-routing]] — patrulha com drones (ver [[problem-formulation]])

## Implementação no Projeto
Os três métodos estão implementados em Go: [[ga]], [[pso]], [[aco]]. Ver também [[architecture]] e [[experiment-pipeline]].

## Papers Relacionados
- [[holland1975adaptation]] — GA fundacional (1975)
- [[kennedy1995particle]] — PSO original (1995)
- [[dorigo1996ant]] — ACO original (1996)
- [[dorigo1997ant]] — ACO para TSP (1997)
- [[bean1994genetic]] — random keys GA (1994)
- [[oliver1987crossover]] — crossover operators GA (1987)
- [[potvin1996ga]] — survey de crossover GA para TSP (1996)
- [[larranaga1999ga]] — survey exaustivo GA+TSP (1999)
- [[clerc2000discretepso]] — PSO discreto para TSP (2000)
- [[stutzle2000mmas]] — MAX-MIN Ant System (2000)
- [[dorigo2004book]] — livro ACO, MIT Press (2004)
- [[blum2005acointro]] — introdução e variantes ACO (2005)
- [[dorigo2005acotheory]] — teoria ACO (2005)
- [[nagata2006eax]] — EAX, crossover GA estado-da-arte para TSP (2006)
- [[wang2021ant]] — SOS-ACO (2021)
- [[deepaco2023]] — DeepACO, neural-enhanced ACO (2023)
- [[ppaco2024]] — PGACO/PPOACO, policy gradient + ACO (2024)
- [[sun2024hybrid]] — PSO híbrido para TSP (2024)
- [[hga2024hybrid]] — GA-ACO híbrido para TSP (2024)
- [[araujo2025pso]] — PSO discreto para TSP (2025)
- [[huang2025matrix]] — PSO matricial para mTSP (2025)
- [[kappagantula2025dpso]] — DPSO+RL para TSP (2025)
- [[neufaco2025]] — NeuFACO, estado-da-arte neural ACO (2025)
- [[gpaco2025]] — GP-ACO, projeto automático de regras ACO (2025)
