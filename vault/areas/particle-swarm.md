---
tags: [area, metaheuristica, bio-inspirado]
created: 2026-06-02
updated: 2026-06-02
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
- [[araujo2025pso]] — PSO discreto para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[sun2024hybrid]] — PSO híbrido para TSP
- [[clerc2000discretepso]] — PSO discreto formalizado (NoHope/ReHope)
- [[kappagantula2025dpso]] — DPSO com RL para TSP
