---
tags:
- tipo/area
- tipo/index
created: 2026-06-02
updated: 2026-06-02
type: area
---

# Catálogo de Áreas — Base de Conhecimento

10 notas de área organizando o conhecimento da literatura por tema. Cada área agrega artigos em [[papers/index]] e se conecta à implementação em [[projeto/index]].

## Áreas

| Nota | Conteúdo | Tags |
|------|----------|------|
| [[tsp]] | TSP clássico: definição, NP-dificuldade, métodos de solução | `area`, `problema-classico` |
| [[tsp-variants]] | Classificação de variantes do TSP na literatura | `area`, `variante-tsp` |
| [[routing]] | Problemas de roteamento em grafos | `area`, `problema-classico` |
| [[drone-routing]] | Roteamento com drones: FSTSP, TSP-D, entregas | `area`, `aplicacao` |
| [[bio-inspired-optimization]] | Otimização bio-inspirada: definição e panorama | `area`, `metodologia` |
| [[genetic-algorithms]] | Algoritmos Genéticos: operadores, representações | `area`, `metaheuristica` |
| [[particle-swarm]] | Particle Swarm Optimization: discreto para TSP | `area`, `metaheuristica` |
| [[ant-colony]] | Ant Colony Optimization: feromônio, construção de rotas | `area`, `metaheuristica` |
| [[comparative-studies]] | Estudos comparativos GA vs PSO vs ACO em TSP | `area`, `benchmark` |
| [[lower-bounds]] | Métodos de lower bound para TSP e aplicabilidade ao TSP-SD-ATP | `area`, `lower-bound` |

## Conexões

- [[projeto/index]] — implementação dos métodos no código Go
- [[papers/index]] — artigos fichados que fundamentam cada área
- [[writing/index]] — scaffolds de escrita que usam estas áreas

## Convenções

- Tags: prefixo `area` + tema específico
- Cada área conecta-se aos artigos em [[papers/index]] que a fundamentam e às notas de [[projeto/index]] que a implementam
