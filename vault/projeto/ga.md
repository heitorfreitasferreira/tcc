---
tags: [projeto, implementacao, ga, algoritmo-genetico]
---

# Algoritmo Genético — Implementação

Implementação em `src/optimization/ga/`. A fundamentação teórica está em [[genetic-algorithms]].

## Parâmetros

| Parâmetro | Flag CLI | Default | Descrição |
|-----------|----------|---------|-----------|
| População | `--population` | 100 | Nº de indivíduos por geração |
| Iterações | `--iterations` | 100 | Nº de gerações |
| Elitismo | `--elitism` | 1 | Nº de melhores mantidos |
| Mutação | `--mutation-rate` | 0.05 | Prob. de swap por indivíduo |
| Torneio | `--tournament-size` | 2 | K do torneio |

## Algoritmo

`src/optimization/ga/main.go` — `Optimize()`:

1. **População inicial**: $n-1$ nós (1..n-1) permutados aleatoriamente
2. **Avaliação**: `g.Makespan(gen)` — usa o tensor 3D
3. **Seleção**: torneio binário (`selectParentTournament`)
4. **Crossover**: Ordered Crossover (OX) — `orderedCrossover`
   - Seleciona dois pontos de corte
   - Copia segmento do pai 1 para filho 1
   - Preenche resto com genes do pai 2 em ordem
   - Simétrico para filho 2
5. **Mutação**: swap de dois genes com prob. `mutationRate`
6. **Elitismo**: `elitism` melhores copiados para próxima geração

## Representação

Cromossomo = slice `[]int` com permutação de `{1, 2, ..., n-1}`.
Nó 0 (base) não aparece — é implícito no início e fim da rota via `Makespan()`.

## Detalhes de Implementação

- Indivíduo: `struct{ gen []int; fen float64 }`
- Crossover OX: `src/optimization/ga/main.go:141` — $O(n)$
- Mutação swap: `src/optimization/ga/main.go:196` — troca dois índices aleatórios
- Relatório de melhoria: `reportImprovement()` — callback `onImprovement` para logging/persistência

## Design Rationale

| Decisão | Alternativas | Por que esta? | Evidência |
|---------|-------------|---------------|-----------|
| **Crossover OX** | PMX, CX, EAX | OX preserva ordem relativa dos genes, adequado para TSP onde a ordem importa. Mais simples que PMX (mapeamento) e muito mais barato que EAX (O(n³)). | [[oliver1987crossover]] mostra OX comparável a PMX em qualidade |
| **Seleção por torneio (k=2)** | Roleta, ranking, torneio k>2 | Torneio binário tem pressão seletiva moderada, não requer fitness escalado (ao contrário da roleta). k=2 é padrão na literatura. | [[goldberg1989genetic]]; torneio é padrão em GAs modernos |
| **Mutação swap (5%)** | Inversão, deslocamento, scramble | Swap é a mutação mais simples para permutações. 5% é valor padrão que evita convergência prematura sem destruir boas soluções. | Testes empíricos mostraram que taxas >10% degradam qualidade |
| **Elitismo (1)** | 0, 2+ | Preserva o melhor indivíduo garantindo monotonicidade. 1 é suficiente; mais reduz diversidade. | Goldberg (1989) prova convergência com elitismo |
| **População 100** | 50, 200, 500 | 100 oferece diversidade adequada para n≤100 com custo computacional aceitável. | Compromisso entre exploração e tempo |
| **Avaliação via Makespan()** | Recalcular ângulos online | Tensor 3D pré-computado → avaliação O(n) vs O(n²) se recalculasse ângulos. | [[problem-formulation]] |

## Conexões

- [[problem-formulation]] — tensor 3D usado na avaliação
- [[pso]] — outra metaheurística populacional no projeto
- [[aco]] — outra metaheurística no projeto
- [[genetic-algorithms]] — fundamentação teórica
- [[oliver1987crossover]] — OX (Ordered Crossover) definido neste paper

## Código-fonte

- `src/optimization/ga/main.go` — `Optimize()`, `selectParentTournament()`, `orderedCrossover()`, `mutateSwap()`
- `src/cmd/ga.go` — CLI flags e integração
