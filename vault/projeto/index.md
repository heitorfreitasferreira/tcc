---
tags: [projeto, implementacao, visao-geral]
---

# Projeto — Otimização Bio-Inspirada para Patrulha com Drones

## Sobre

Este diretório documenta a **implementação em Go** do TCC. Cada nota reflete diretamente o código em `src/`, diferentemente das notas em `areas/` (conhecimento da literatura) e `papers/` (fichamento de artigos).

## Notas do Projeto

| Nota | Conteúdo | Código-fonte |
|------|----------|-------------|
| [[problem-formulation]] | Formulação TSP-SD-ATP, tensor 3D, função objetivo | `src/graph/` |
| [[ga]] | Algoritmo Genético — seleção, crossover, mutação | `src/optimization/ga/` |
| [[pso]] | Particle Swarm — posição, velocidade, random keys | `src/optimization/pso/` |
| [[aco]] | Ant Colony — feromônio 3D, heurística, construção de rota | `src/optimization/aco/` |
| [[bruteforce]] | Busca exaustiva (Heap's algorithm) | `src/optimization/brute/` |
| [[experiment-pipeline]] | CLI, scripts, execução, resultados | `src/cmd/`, `src/run_*.sh` |
| [[architecture]] | Estrutura de pacotes, fluxo de dados | `src/` |

## Convenções

- Tags: `projeto`, `implementacao`, `go`, `tcc`
- [[Links]] para `areas/` quando um conceito da literatura é usado
- [[Links]] para `papers/` quando um artigo específico fundamenta uma decisão
- Caminhos de código relativos a `src/`

## Conexões

- [[TSP]] — o problema clássico que esta implementação estende
- [[tsp-variants]] — classificação da variante TSP-SD-ATP na literatura
- [[drone-routing]] — contexto de aplicação
- [[bio-inspired-optimization]] — metaheurísticas implementadas
- [[comparative-studies]] — contexto experimental
