---
tags: [projeto, implementacao, formulacao, tsp, tensor, siglas]
sigla: "TSP-SD-ATP"
definicao: "Traveling Salesman Problem with Sequence-Dependent Angular Turn Penalties"
incluir: sim
ocorrencias_ac: 0
ocorrencias_texto: 0
arquivos_ac: ""
---

# Formulação do Problema — TSP-SD-ATP

> **TSP com Custos de Curva Angulares Dependentes da Sequência** (*Sequence-Dependent Angular Turn Penalties*)

Esta nota descreve exclusivamente a formulação implementada em `src/`. A classificação na literatura está em [[tsp-variants]].

## Instâncias

Geradas por `src/cmd/map.go` → `src/points/creater.go`:

````go
x := rng.Float64()*2 - 1   // [-1, +1)
y := rng.Float64()*2 - 1   // [-1, +1)
````

- $n$ pontos 2D em $[-1,1)^2$
- Nó **0** = base (origem e retorno obrigatório)
- Nós $1 \dots n-1$ = POIs (*Points of Interest*)
- Instâncias nomeadas `{n}{letra}.points` (ex: `10a.points`, `10b.points`)

## Tensor 3D de Custo

Pré-computado por `src/graph/creater.go` → `New()`:

```go
penalty[k][i][j] = turnCost(currVector, nextVector) + time
```

Onde:

- $time = \text{EuclideanDistance}(i, j) / \text{droneSpeed}$ (`droneSpeed = 1`)
- $\text{turnCost} = \text{maxPenalty} \times (\theta / \pi)$ (`maxPenalty = 1`)
- $\theta = \arccos\left(\frac{\vec{v}_{ki} \cdot \vec{v}_{ij}}{\|\vec{v}_{ki}\|\|\vec{v}_{ij}\|}\right)$

Resultado: tensor $G[i][j][k]$ em `src/graph/types.go`:

```go
type Graph [][][]float64 // [anterior][atual][proximo]
```

`Graph` é um alias para `[][][]float64` — 3D slice. Cada célula contém **tempo de percurso + penalidade de curva**.

## Função Objetivo — Makespan

`src/graph/makespan.go`:

```go
func (g Graph) Makespan(order []int) float64 {
    lastNode := 0
    currNode := 0
    for _, nextNode := range order {
        makespan += g[lastNode][currNode][nextNode]
        lastNode = currNode
        currNode = nextNode
    }
    makespan += g[lastNode][currNode][0] // retorno à base
    return makespan
}
```

- Parte do nó 0, percorre a sequência `order`, retorna a 0
- Cada passo indexa $G[anterior][atual][próximo]$ — $O(n)$ lookups
- A penalidade angular já está embutida no tensor

## Por que Tensor 3D?

Diferente do TSP clássico (matriz 2D $C[i][j]$), aqui o custo de ir de $i$ para $j$ **depende de onde viemos** (nó anterior $k$). Isso porque:

- A curva entre $\vec{v}_{ki}$ e $\vec{v}_{ij}$ depende do ângulo de entrada e saída
- Drone não pode virar instantaneamente — curvas fechadas custam tempo

O tensor pré-computado $O(n^3)$ é uma **técnica de embedded turn cost**: as metaheurísticas (GA, PSO, ACO) não precisam calcular ângulos durante a otimização.

> **Consequência para ACO:** o tensor exige feromônio 3D ($\tau_{ijk}$) com N³ entradas — ~98× mais que ACO 2D clássico para n=100. Esta **diluição estrutural** é inerente ao TSP-SD-ATP, não um bug de implementação. Ver [[aco#Limitações Inerentes ao TSP-SD-ATP]].

## Contraste com TSP Clássico

| Aspecto | TSP Clássico | Este Projeto |
|---------|-------------|--------------|
| Espaço de busca | $(n-1)!$ permutações | $(n-1)!$ permutações (mesmo) |
| Custo da aresta | $C[i][j]$ fixo (matriz 2D) | $G[k][i][j]$ depende do anterior (tensor 3D) |
| Avaliação | $O(n)$ lookups na matriz | $O(n)$ lookups no tensor (mesmo) |
| Penalidade angular | Não | Sim, normalizada $[0,1]$ |
| Simetria | $C[i][j]=C[j][i]$ (simétrico) | Assimétrico devido à penalidade |

## Nome Interno

No código e scripts, ocasionalmente chamado de **"rTSP"** (*routing TSP*). Não é nomenclatura formal na literatura — apenas apelido interno.

## Código-fonte

- Geração de pontos: `src/points/creater.go`
- Cálculo de ângulo: `src/graph/math.go`
- Construção do tensor: `src/graph/creater.go`
- Avaliação: `src/graph/makespan.go`
- Tipos: `src/graph/types.go`, `src/points/types.go`
