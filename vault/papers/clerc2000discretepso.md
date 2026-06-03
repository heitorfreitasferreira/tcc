---
title: "Discrete Particle Swarm Optimization, Illustrated by the Traveling Salesman Problem"
authors: [Clerc, Maurice]
year: 2000
doi: ""
bibtex-key: clerc2000discretepso
tags: [pso tsp survey]
status: lido
rating: 5
---

## Resumo

O artigo seminal que estende o PSO clássico (contínuo) para problemas discretos/combinatórios, usando o TSP como caso de estudo. Define formalmente os elementos do PSO discreto: posições como permutações (tours), velocidade como lista de transposições, e operações de subtração/adição de permutações. Introduz os conceitos de NoHope/ReHope para evitar estagnação: quando o swarm não melhora, ele é reexpandido usando métodos como Lazy Descent Method (LDM), Deep Descent Method (DDM) e Local Iterative Levelling (LIL). A abordagem é apresentada como genérica: "se você não tem um algoritmo específico para seu problema discreto, use PSO — funciona".

## Contribuições Principais

- Primeira formalização matemática de PSO discreto para problemas de permutação
- Definição das operações de velocidade (transposições) e posição (permutações) para TSP
- Mecanismo NoHope/ReHope para escape de ótimos locais
- Demonstração prática em instâncias TSP de até 100 cidades

## Relevância para o TCC

Referência fundamental para justificar a escolha de PSO discreto para TSP/rTSP. A definição de velocidade como lista de transposições e o mecanismo NoHope/ReHope são implementados diretamente.

## Métodos e Abordagens

- Posição: permutação (tour) — vetor de inteiros
- Velocidade: lista ordenada de transposições (swap operators) com prioridades
- Operações: subtração (diferença entre duas permutações), adição (aplicar transposições)
- NoHope: detecta estagnação e reexpande o swarm
- ReHope: reexpansão via LDM, DDM, ou LIL
- Coeficientes: c1 (cognitive), c2 (social), w (inércia) adaptados para o discreto

## Conexões

- [[kennedy1995particle]] — PSO original, base conceitual
- [[araujo2025pso]] — PSO discreto moderno para TSP (DPSO+2-opt/3-opt)
- [[sun2024hybrid]] — PSO híbrido com busca local greedy+Metropolis
- [[kappagantula2025dpso]] — DPSO com RL
- [[particle-swarm]]
- [[tsp]]
