---
tags:
- area/bio-inspired-optimization
- metodo/metaheuristic
- tipo/area
created: 2026-06-02
updated: 2026-06-05
type: area
---

# Particle Swarm Optimization (PSO)

## Definição
Método de otimização inspirado no comportamento social de pássaros/peixes ([[kennedy1995particle]]). Partículas movem-se no espaço de busca combinando memória individual e informação social.

## Aplicação em TSP
- Representação: permutação via vetores de posição
- Estratégias: troca de arestas, atualização de velocidade com operadores discretos
- Topologias: global, local, anel

## Conexões
- [[genetic-algorithms]] — outra populacional
- [[tsp]]
- [[pso]] — implementação no projeto (Go)

## Papers Relacionados
- [[kennedy1995particle]] — PSO original
- [[shami2022pso]] — survey mais citada (1249 cites), taxonomia de variantes (inércia, hibridização, multiobjetivo, discreto), convergência prematura
- [[gad2022pso]] — revisão sistemática PRISMA, análise bibliométrica, subexploração em otimização combinatória discreta
- [[araujo2025pso]] — PSO discreto para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[sun2024hybrid]] — PSO híbrido para TSP
- [[clerc2000discretepso]] — PSO discreto formalizado (NoHope/ReHope)
- [[kappagantula2025dpso]] — DPSO com RL para TSP

## Citações na Monografia
- Seção 2.6 (fundamentacao.tex): `shami2022pso` e `gad2022pso` citados para reforçar revisão de PSO e balancear cobertura ACO vs PSO (P34, 2026-06-05)
