---
tags: [area, lower-bound, otimizacao]
created: 2026-06-02
updated: 2026-06-02
---

# Lower Bounds para TSP — Abordagens e Aplicabilidade ao TSP-SD-ATP

## Definição

Um **lower bound** (limitante inferior) para um problema de minimização é um valor `LB` tal que `LB ≤ OPT` para toda instância, onde `OPT` é o valor da solução ótima. Lower bounds são essenciais para:
- **Avaliar qualidade** de soluções heurísticas quando o ótimo é desconhecido (gap = (heurística - LB) / LB)
- **Podar** árvores de branch-and-bound
- **Guiar** algoritmos de aproximação

## Métodos Clássicos (TSP Simétrico)

### Held-Karp (1-tree + Relaxação Lagrangiana)
O bound mais utilizado para TSP simétrico, proposto por [[heldkarp1970traveling]].

**Mecanismo:** Remove um vértice `p`, calcula MST sobre V\{p}, adiciona as duas arestas mais baratas incidentes a `p`. Multiplicadores de Lagrange ajustam os custos das arestas para forçar grau 2 em todos os vértices. O máximo da função Lagrangiana é o lower bound.

**Equivalência:** O HK bound é idêntico ao valor da relaxação LP do TSP (subtour LP) — [[heldkarp1970traveling]], [[applegate2006traveling]].

**Qualidade:** Gap < 0.8% para instâncias Euclideanas aleatórias, < 2% para TSPLIB — [[johnson1996asymptotic]].

**Implementação prática:** Subgradiente sobre 1-trees — [[valenzuela1997estimating]].

**Otimização:** A escolha do vértice `p` pode ser otimizada em O(m + n log n) — [[righini2021efficient]].

**Pré-requisito:** O HK bound assume custos de aresta **2D simétricos** (`c_ij = c_ji`).

### Minimum Spanning Tree (MST)
Um bound simples: o custo de uma MST é um lower bound para o TSP (pois remover uma aresta de um tour produz uma árvore geradora). Gap típico: 15-20% do ótimo.

## Métodos para ATSP (Assimétrico)

### Assignment Problem (AP)
O bound mais usado para ATSP, proposto por [[balas1985branch]].

**Mecanismo:** Relaxa as restrições de conectividade do ATSP, mantendo apenas as restrições de grau (cada vértice tem exatamente uma aresta de entrada e uma de saída). O resultado é um conjunto de subtours. Resolvido pelo algoritmo Hungaro em O(n³).

**Qualidade:** Para ATSP com custos U[0,1], o gap AP-ATSP é o(1) — [[karp1979patching]].

### Additive Bounding (Fischetti & Toth)
Combina AP + r-arborescência + r-antiarborescência sequencialmente, onde cada relaxação opera sobre custos residuais da anterior. O bound final é sempre ≥ o melhor bound individual — [[fischetti1992additive]].

## Métodos para TSP com Dependência de Sequência (3D)

### Diagramas de Decisão Multivalorados (MDD)
Proposto por [[kinable2017hybrid]] para TDTSP (Time-Dependent TSP). O MDD é construído a partir do DP natural: estados (S, prev, curr). Limitando a largura do MDD via merge de estados, obtém-se uma relaxação que fornece lower bounds diretamente sobre o tensor 3D, sem redução para 2D.

**Vantagem:** Funciona nativamente com o tensor `cost[i][j][k]` do TSP-SD-ATP — cada transição no MDD usa `cost[prev][curr][next]`.

**Complexidade:** O(W · n²), onde W é a largura máxima do MDD.

### ng-path Relaxation
Adaptado para TDTSP por [[leraromero2020dynamic]]. Cada estado mantém apenas um subconjunto limitado NG(i) de nós visitados, reduzindo o espaço de estados do DP. O labeling bidirecional permite computar lower bounds sem upper bounds apertados.

### Angular-Metric TSP
[[aggarwal2000angular]] define o AM-TSP (minimizar soma de ângulos de curva). Provam NP-hardness e dão aproximação O(log n). O TSP-SD-ATP difere porque o makespan = distância + penalidade angular.

## Aplicação ao TSP-SD-ATP (Tensor 3D)

O tensor `Graph [][][]float64` com índices `[prev][curr][next]` impede aplicação direta de HK ou AP.

### Abordagem por Redução 3D→2D
Definir `c'[j][k] = min_i cost[i][j][k]`. Esta matriz 2D subestima cada transição, portanto qualquer lower bound computado sobre `c'` é válido para o TSP-SD-ATP. Permite aplicar:
- **AP bound** (Hungarian) — O(n³)
- **Held-Karp** (1-tree + subgradiente) — O(K·n²)

### Abordagem MDD Nativa
Sem redução, usando o MDD com largura limitada W. Implementável em Go puro. O estado (S, prev, curr) captura a dependência completa. Merge de estados por projeção (ex: ignorar nós "distantes" no conjunto S).

## Resumo Comparativo

| Método | 2D | 3D Nativo | Complexidade | Força LB |
|--------|:--:|:----------:|:------------:|:--------:|
| MST | ✓ | via redução | O(n²) | Fraca |
| AP (Hungarian) | ✓ | via redução | O(n³) | Forte |
| Held-Karp (1-tree) | ✓ | via redução | O(K·n²) | Muito forte |
| Additive Bounding | ✓ | via redução | O(n³ + n²) | Muito forte |
| MDD (largura W) | ✓ | **✓ nativo** | O(W·n²) | Forte (W-dependente) |
| ng-path relaxation | ✓ | **✓ nativo** | O(n²·NG) | Forte |

## Referências no Vault

- [[heldkarp1970traveling]] — bound HK original
- [[heldkarp1971traveling]] — DP O(n²2ⁿ) + subgradiente
- [[johnson1996asymptotic]] — validação empírica HK
- [[valenzuela1997estimating]] — implementação subgradiente
- [[righini2021efficient]] — otimização vértice HK
- [[fischetti1992additive]] — additive bounding ATSP
- [[balas1985branch]] — B&B com relaxação AP
- [[karp1979patching]] — patching AP→ATSP
- [[kinable2017hybrid]] — MDD para TDTSP
- [[leraromero2020dynamic]] — ng-path para TDTSP
- [[aggarwal2000angular]] — AM-TSP NP-hard
