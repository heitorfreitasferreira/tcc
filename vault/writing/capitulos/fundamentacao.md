---
tags:
- capitulo/fundamentacao
- tipo/writing
- topico/monografia
created: 2026-06-02
updated: 2026-06-02
type: writing
---

# Fundamentação Teórica — Scaffold

## Estrutura do Capítulo

### 2.1 Problema do Caixeiro Viajante
- Definição, NP-dificuldade ([[garey1979computers]])
- TSP clássico vs ATSP vs variantes ([[tsp-variants]])
- Formulação matemática

### 2.2 TSP com Custos de Curva
- Motivação física: drones e ângulos de virada
- Trabalhos relacionados: Winter ([[winter2002modeling]]), Vanhove ([[vanhove2012route]])
- AM-TSP (Fekete & Krupke) vs TSP-SD-ATP
- Tensor 3D como estratégia de embedded turn cost ([[problem-formulation]])

### 2.3 Algoritmos Genéticos
- Inspiração biológica ([[holland1975adaptation]], [[goldberg1989genetic]])
- Representação, seleção, crossover, mutação
- Aplicação ao TSP: PMX, OX, CX ([[oliver1987crossover]], [[potvin1996ga]], [[larranaga1999ga]])
- EAX ([[nagata2006eax]])

### 2.4 Particle Swarm Optimization
- Inspiração social ([[kennedy1995particle]])
- PSO contínuo: velocidade, inércia, componentes cognitiva/social
- PSO discreto para TSP: random keys, swap operators ([[clerc2000discretepso]])

### 2.5 Ant Colony Optimization
- Inspiração em formigas ([[dorigo1996ant]])
- Ant System → Ant Colony System ([[dorigo1997ant]])
- MAX-MIN Ant System ([[stutzle2000mmas]])
- ACO para TSP: feromônio em arestas, heurística de visibilidade

### 2.6 Trabalhos Relacionados
- Estudos comparativos ([[chandra2022comparative]], [[wu2020comparative]], [[halim2019combinatorial]])
- Survey de metaheurísticas para TSP ([[toaza2023review]])
- GTSP ([[pop2024comprehensive]])
- Drone routing ([[murray2015flying]], [[agatz2018optimization]], [[dellamico2021multiple]])

## Material de Apoio

- [[tsp]] — visão geral
- [[tsp-variants]] — classificação completa
- [[genetic-algorithms]] — GA na literatura
- [[particle-swarm]] — PSO na literatura
- [[ant-colony]] — ACO na literatura
- [[drone-routing]] — contexto drone
- [[problem-formulation]] — formulação implementada
- [[comparative-studies]] — estudos comparativos
