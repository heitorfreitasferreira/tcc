---
tags: [area, metaheuristica, bio-inspirado]
created: 2026-06-02
updated: 2026-06-02
---

# Algoritmos Genéticos (GA)

## Definição
Método de otimização inspirado na seleção natural ([[holland1975adaptation]]). População de soluções evolui via cruzamento, mutação e seleção.

## Aplicação em TSP
- Representação: permutação de cidades
- Operadores de cruzamento: PMX, OX, CX ([[oliver1987crossover]])
- Mutação: swap, inversão, deslocamento
- Seleção: torneio, roleta, elitismo

## Conexões
- [[particle-swarm]] — outra populacional
- [[ant-colony]] — outra bio-inspirada
- [[tsp]]
- [[ga]] — implementação no projeto (Go)

## Papers Relacionados
- [[holland1975adaptation]] — GA fundacional
- [[goldberg1989genetic]] — textbook de referência
- [[oliver1987crossover]] — operadores de cruzamento para TSP (PMX, OX, CX)
- [[bean1994genetic]] — random keys para sequenciamento
- [[potvin1996ga]] — survey de operadores de crossover para TSP
- [[larranaga1999ga]] — survey exaustivo de GA para TSP (representações e operadores)
- [[nagata2006eax]] — EAX, crossover estado-da-arte para TSP
- [[hga2024hybrid]] — GA-ACO híbrido para TSP
