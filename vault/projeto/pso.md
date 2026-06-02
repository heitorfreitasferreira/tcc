---
tags: [projeto, implementacao, pso, particle-swarm]
---

# Particle Swarm Optimization — Implementação

Implementação em `src/optimization/pso/`. A fundamentação teórica está em [[particle-swarm]].

## Parâmetros

| Parâmetro | Flag CLI | Default | Descrição |
|-----------|----------|---------|-----------|
| População | `--population` | 100 | Nº de partículas |
| Iterações | `--iterations` | 100 | Nº de iterações |
| $C_1$ (cognitivo) | `--c1` | 1.5 | Peso da memória individual |
| $C_2$ (social) | `--c2` | 1.5 | Peso da informação global |
| $W$ (inércia) | `--w` | 0.7 | Peso da velocidade anterior |

## Algoritmo

`src/optimization/pso/main.go` — `Optimize()`:

1. **Inicialização**: cada partícula recebe posição $x_i \in [0,1]^d$ aleatória, velocidade $v_i = 0$
2. **Decodificação**: `setSequence()` ordena os $x_i$ por valor (random keys) → permutação
3. **Avaliação**: `sw.Makespan(sequence)` — usa o tensor 3D
4. **Atualização** (para cada partícula):
   - $v_i = W \cdot v_i + C_1 \cdot r_1 \cdot (pBest_i - x_i) + C_2 \cdot r_2 \cdot (gBest - x_i)$
   - $x_i = x_i + v_i$
5. **Decodificação** → sequência → avaliação
6. Repete iterações

## Representação — Random Keys

Decodificação em `src/optimization/pso/particle.go` — `setSequence()`:

```go
slices.SortFunc(indices, func(a, b int) int {
    return cmp.Compare(p.x[a], p.x[b])
})
```

Cada partícula mantém vetor $x \in \mathbb{R}^d$ (posição contínua). A sequência é obtida **ordenando os índices** por valor crescente de $x$.

Exemplo: $x = [0.3, 0.1, 0.8, 0.5]$ → índices ordenados $[1, 0, 3, 2]$ → sequência $[2, 1, 4, 3]$ (soma +1).

## Estrutura

`src/optimization/pso/particle.go` — `particle` struct:
```go
type particle struct {
    x            []float64   // posição (random keys)
    v            []float64   // velocidade
    bestX        []float64   // melhor posição individual
    sequence     []int       // rota decodificada (nós 1..n-1)
    makespan     float64     // custo atual
    bestMakespan float64     // melhor custo individual
}
```

Enxame (`Swarm` struct) mantém: partículas, $gBest$ (melhor global), referência ao grafo.

## Design Rationale

| Decisão | Alternativas | Por que esta? | Evidência |
|---------|-------------|---------------|-----------|
| **Random Keys** | Swap-operator, permutação direta | RK permite usar equação de velocidade do PSO contínuo padrão sem modificações. A decodificação por ordenação é O(n log n). | [[kennedy1995particle]] define PSO para ℝⁿ; [[bean1994genetic]] introduz random keys |
| **C1=C2=1.5** | C1=C2=2.0, assimétrico | Valores canônicos que equilibram exploração individual e social. 1.5 é mais conservador que 2.0, reduzindo overshooting. | Clerc & Kennedy (2002) mostram que C1+C2 > 4 causa divergência |
| **W=0.7 (inércia)** | W adaptativo, W decrescente | Inércia fixa 0.7 favorece exploração global. Valor abaixo de 1 desacelera partículas gradualmente. | Shi & Eberhart (1998) — valores entre 0.4-0.9 são eficazes |
| **População 100** | 30, 50, 200 | Consistente com GA/ACO para comparação justa. Dimensão do espaço = n-1, população 100 é suficiente. | — |
| **Limitação conhecida** | — | Random keys perdem informação de adjacência. Diferente de operadores de permutação, a ordenação não captura arestas. Isto explica o desempenho inferior do PSO no TSP-SD-ATP. | [[clerc2000discretepso]] propõe alternativas (NoHope/ReHope) não implementadas |

## Conexões

- [[problem-formulation]] — tensor 3D usado na avaliação
- [[ga]] — outra populacional (contraste: GA usa crossover, PSO usa atualização de velocidade)
- [[aco]] — outra metaheurística no projeto
- [[particle-swarm]] — fundamentação teórica
- [[clerc2000discretepso]] — base do PSO discreto (NoHope/ReHope)

## Código-fonte

- `src/optimization/pso/main.go` — `Optimize()`, `update()`, `evaluate()`
- `src/optimization/pso/particle.go` — `setSequence()`, `update()`
- `src/cmd/pso.go` — CLI flags e integração
