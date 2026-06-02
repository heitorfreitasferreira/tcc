---
tags: [projeto, implementacao, arquitetura, go, pacotes]
---

# Arquitetura do Código

## Estrutura de Pacotes

```
src/
├── main.go                        ← entrypoint
├── cmd/                           ← CLI (Cobra)
│   ├── root.go                    ←   flags globais (--seed, --folder)
│   ├── create.go                  ←   geração de instâncias
│   ├── map.go                     ←   geração de pontos
│   ├── graph.go                   ←   geração de grafos (tensor 3D)
│   ├── optimize.go                ←   otimização (flags comuns)
│   ├── ga.go                      ←   GA CLI
│   ├── pso.go                     ←   PSO CLI
│   ├── aco.go                     ←   ACO CLI
│   ├── bruteforce.go              ←   brute force CLI
│   ├── reporting.go               ←   persistência de resultados
│   └── serve.go                   ←   servidor web
├── points/                        ← pontos 2D
│   ├── types.go                   ←   Coordinate2D, Points2D
│   ├── creater.go                 ←   geração aleatória [-1,+1)
│   ├── math.go                    ←   distância euclidiana, vetores
│   └── io.go                      ←   load/save JSON
├── graph/                         ← tensor de custo 3D
│   ├── types.go                   ←   Graph = [][][]float64
│   ├── creater.go                 ←   New() — constrói tensor
│   ├── math.go                    ←   angle(), turnCost()
│   ├── makespan.go                ←   Makespan(), IsValidSolution()
│   └── io.go                      ←   load/save JSON
├── optimization/                  ← metaheurísticas
│   ├── ga/                        ← GA
│   │   └── main.go                ←   Optimize(), crossover, mutação
│   ├── pso/                       ← PSO
│   │   ├── main.go                ←   Swarm, Optimize()
│   │   └── particle.go            ←   partícula, random keys
│   ├── aco/                       ← ACO
│   │   ├── main.go                ←   ACO struct, Optimize()
│   │   └── ant.go                 ←   formiga, walk, feromônio
│   └── brute/                     ← exaustiva
│       └── main.go                ←   Optimize(), Heap permutations
├── shared/                        ← tipos compartilhados
│   ├── types.go                   ←   HyperParams
│   ├── optimization.go            ←   OptimizationResult, Improvement
│   ├── slices.go                  ←   Shuffle, RandomizeSlice
│   ├── stats.go                   ←   utilitários
│   └── reporting/                 ←   schemas JSON, escritores
├── data/                          ← instâncias embutidas
│   ├── assets.go                  ←   //go:embed
│   ├── {n}{letra}.points          ←   pontos
│   └── {n}{letra}.graph           ←   tensor 3D
├── run_all.sh                     ← executor batch
├── run_experiments_multi_seed.sh  ← multi-semente
├── run_experiments_bruteforce_missing.sh
├── go.mod
└── Dockerfile
```

## Fluxo de Dados

```
create (CLI)
  │
  ├── points.CreateInstances()    →  Points2D (nós em [-1,+1)²)
  │
  └── graph.New(pts)              →  Graph (tensor 3D G[prev][curr][next])
       │                                Cada célula = distance + turnPenalty
       │
       ▼
optimize ga|pso|aco|bruteforce
  │
  ├── g.Makespan(order)           →  float64  (avalia rota)
  │
  └── persistOptimizeRun()        →  JSON (summary, evolution, timing)
```

## Padrões de Projeto

- **Strategy**: CLI otimiza chama diferentes pacotes (`ga.Optimize`, `pso.Optimize`, etc.)
- **Callback**: `onImprovement` para logging/persistência durante otimização
- **Embedded Turn Cost**: penalidade angular pré-computada no tensor (não durante otimização)
- **Generator**: `generatePermutations` usa canal Go para streaming de permutações
- **Random Keys**: PSO decodifica posição contínua em permutação via ordenação

## Dependências Externas

- `github.com/spf13/cobra` — CLI
- Nenhuma para otimização (matemática pura stdlib)

## Padrões de Projeto (Comentários)

- **Strategy**: CLI escolhe pacote otimizador em runtime → análogo ao padrão Strategy de [[gamma1994designpatterns]]. Cada método (GA, PSO, ACO) implementa a mesma interface implícita.
- **Embedded Turn Cost**: pré-computação da penalidade angular no tensor → análogo ao pseudo-dual graph de [[winter2002modeling]], mas adaptado para TSP (shortest-path vs permutação).
- **Generator (permutações)**: `generatePermutations` usa canal Go para streaming → evita alocar O(n!) na memória.

## Conexões

- [[experiment-pipeline]] — uso da CLI e scripts
- [[problem-formulation]] — o tensor que o pipeline processa
- [[ga]], [[pso]], [[aco]], [[bruteforce]] — otimizadores
- [[bio-inspired-optimization]] — contexto geral dos métodos
- [[winter2002modeling]] — abordagem alternativa para turn costs (pseudo-dual graph)

## Código-fonte

- `src/main.go`
- `src/go.mod`
- `src/Dockerfile`
