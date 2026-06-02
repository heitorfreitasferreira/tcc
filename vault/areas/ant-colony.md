---
tags: [area, metaheuristica, bio-inspirado]
---

# Ant Colony Optimization (ACO)

## Definição
Método de otimização inspirado no comportamento de formigas ([[dorigo1996ant]]). Formigas artificiais depositam feromônio nas arestas do grafo, guiando futuras soluções.

## Aplicação em TSP
- Naturalmente adequado: TSP é um problema em grafo
- Atualização de feromônio: deposição e evaporação
- Heurística: informação de distância (visibilidade)
- Ant System (AS), Ant Colony System (ACS), MAX-MIN

## Conexões
- [[genetic-algorithms]]
- [[particle-swarm]]
- [[TSP]]
- [[aco]] — implementação no projeto (Go)

## Papers Relacionados
- [[dorigo1996ant]] — Ant System (ACO original)
- [[dorigo1997ant]] — ACO aplicado ao TSP
- [[dorigo2004book]] — livro referência (MIT Press)
- [[dorigo2005acotheory]] — survey teórico (convergência, Model-Based Search)
- [[blum2005acointro]] — introdução e variantes (AS, ACS, MMAS)
- [[stutzle2000mmas]] — MAX-MIN Ant System
- [[wang2021ant]] — SOS-ACO com otimização de parâmetros
- [[deepaco2023]] — DeepACO, neural-enhanced ACO
- [[neufaco2025]] — NeuFACO, estado-da-arte neural ACO
- [[gpaco2025]] — GP-ACO, projeto automático de regras de transição
- [[ppaco2024]] — PGACO/PPOACO, policy gradient ACO
- [[hga2024hybrid]] — GA-ACO híbrido para TSP
