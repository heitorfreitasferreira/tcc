---
tags: [area, problema-classico]
created: 2026-06-02
updated: 2026-06-02
---

# Traveling Salesman Problem (TSP)

## Definição
Dado um conjunto de cidades e distâncias entre cada par, encontrar o menor caminho que visita cada cidade uma única vez e retorna à origem. Problema NP-difícil ([[garey1979computers]]).

## Variantes

Classificação completa em [[tsp-variants]].

As principais variantes na literatura:

- **TSP Simétrico**: $C[i][j] = C[j][i]$, distância euclidiana. NP-difícil.
- **TSP Assimétrico (ATSP)**: $C[i][j] \neq C[j][i]$, arestas direcionais.
- **GTSP** (*Generalized TSP*): clusters, visita um nó por cluster [[pop2024comprehensive]]
- **CTSP** (*Clustered TSP*): clusters, visita todos com contiguidade
- **PCTSP** (*Prize-Collecting TSP*): permite pular nós com penalidade
- **TSP com Drone (FSTSP/TSP-D)**: [[drone-routing]]
- **TSP com Custos de Curva**: penalidades angulares na transição entre arestas

A **variante implementada neste TCC** (TSP-SD-ATP) está documentada em [[problem-formulation]].

## Métodos de Solução
- Exatos: Programação Linear Inteira, Branch-and-Bound
- Aproximativos: Christofides, heurísticas LKH
- Lower Bounds: [[heldkarp1970traveling]], [[fischetti1992additive]], [[lower-bounds]]
- Metaheurísticas: [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]]

## Aplicação no TCC
Problema base para comparação de métodos bio-inspirados (GA, PSO, ACO, busca exaustiva) em cenário de patrulha com drones. A formulação exata implementada está em [[problem-formulation]] e os métodos em [[ga]], [[pso]], [[aco]], [[bruteforce]].

## Conexões

- [[lower-bounds]] — métodos de lower bound e aplicabilidade ao TSP-SD-ATP

## Papers Relacionados

### TSP Clássico
- [[garey1979computers]] — NP-completude
- [[lawler1985traveling]] — survey clássico
- [[applegate2006traveling]] — estudo computacional (Concorde)
- [[lin1973effective]] — heurística LK
- [[pop2024comprehensive]] — survey GTSP
- [[bock2025survey]] — survey TSP variants em warehousing

### Turn Costs e Penalidade Angular
- [[winter2002modeling]] — modelagem de custos de curva (pseudo-dual graph)
- [[vanhove2012route]] — experimentos computacionais com turn restrictions

### GA para TSP
- [[oliver1987crossover]] — operadores de cruzamento para TSP
- [[potvin1996ga]] — survey de crossover GA para TSP
- [[larranaga1999ga]] — survey GA+TSP (representações e operadores)
- [[nagata2006eax]] — EAX, crossover estado-da-arte para TSP
- [[hga2024hybrid]] — GA-ACO híbrido para TSP

### PSO para TSP
- [[clerc2000discretepso]] — PSO discreto para TSP
- [[sun2024hybrid]] — PSO híbrido para TSP
- [[araujo2025pso]] — PSO discreto para TSP
- [[huang2025matrix]] — PSO para mTSP
- [[kappagantula2025dpso]] — DPSO+RL para TSP

### ACO para TSP
- [[dorigo1997ant]] — ACO para TSP
- [[stutzle2000mmas]] — MMAS para TSP
- [[dorigo2004book]] — ACO book com capítulos dedicados ao TSP
- [[dorigo2005acotheory]] — fundamentos teóricos ACO
- [[blum2005acointro]] — introdução ACO com exemplos TSP
- [[wang2021ant]] — SOS-ACO para TSP
- [[deepaco2023]] — DeepACO, neural-enhanced ACO
- [[ppaco2024]] — PGACO/PPOACO, policy gradient ACO
- [[neufaco2025]] — NeuFACO, estado-da-arte neural ACO
- [[gpaco2025]] — GP-ACO, projeto automático de regras ACO

### Estudos Comparativos
- [[wu2020comparative]] — GA vs PSO vs ACO em TSPLIB
- [[haroun2015performance]] — GA vs ACO em TSP
- [[chandra2022comparative]] — ANOVA/Tukey de 8 métodos
- [[almufti2025comparative]] — 9 metaheurísticas em TSP
- [[wadi2025charting]] — swarm-based para TSP
- [[hossain2024comparison]] — clássicos vs modernos em TSP
- [[toaza2023review]] — 120 metaheurísticas em scheduling TSP
