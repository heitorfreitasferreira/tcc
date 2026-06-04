---
tags: [writing, capitulo, proposta]
created: 2026-06-02
updated: 2026-06-02
---

# Proposta — Scaffold

## Estrutura do Capítulo

### 3.1 Formulação do Problema
- Pontos em [-1,1]², nó 0 = base
- Tensor 3D G[prev][curr][next] ([[problem-formulation]])
- Custo = distância euclidiana + penalidade angular
- Makespan como função objetivo

### 3.2 Arquitetura da Implementação
- Estrutura de pacotes Go ([[architecture]])
- Geração de instâncias (points → graph)
- Interface comum de otimização (`shared.OptimizationResult`)

### 3.3 Algoritmo Genético
- Representação: permutação
- Crossover OX ([[ga]])
- Mutação swap, seleção por torneio, elitismo
- Design rationale: por que OX? Torneio tamanho 2? Mutação 5%? ([[ga]])

### 3.4 Particle Swarm Optimization
- Random keys para codificação contínua→discreta ([[pso]])
- Atualização de velocidade com inércia W
- Design rationale: por que random keys? C1=C2=2.0? W=0.7?

### 3.5 Ant Colony Optimization
- Feromônio 3D (τ[i][j][k]) ([[aco]])
- Heurística η = 1/G[i][j][k]
- Construção de rota + atualização global
- Design rationale: por que roleta e não pseudo-aleatória proporcional?

### 3.6 Busca Exaustiva
- Heap's algorithm para permutações ([[bruteforce]])
- Uso como baseline ótima nas instâncias executadas `10a..15c`

### 3.7 Pipeline Experimental
- CLI em Cobra ([[experiment-pipeline]])
- Resultados estruturados (summary, evolution, timing)
- 51 sementes × 30 instâncias × 3 métodos

## Design Rationale (para discussão na proposta)

| Decisão | Alternativas | Por que esta? |
|---------|-------------|---------------|
| OX crossover | PMX, CX, EAX | Simplicidade + bom desempenho empírico; EAX seria caro demais |
| Torneio tamanho 2 | Roleta, ranking | Baixa pressão seletiva → mantém diversidade |
| Mutação swap 5% | Inversão, deslocamento | Swap é a mais simples e suficiente |
| Random keys (PSO) | Swap-operator | RK permite usar PSO contínuo padrão sem modificar equação |
| C1=C2=2.0, W=0.7 | Diversos da literatura | Valores canônicos; sem tuning específico (limitação) |
| Ant System com roleta | Pseudo-aleatória proporcional | Mais exploratória; sem q₀ para calibragem |
| Tensor 3D | Pseudo-dual graph (Winter) | Pré-computação O(n³) → avaliação O(n); simples de implementar |

## Material de Apoio

- [[ga]], [[pso]], [[aco]], [[bruteforce]] — implementações
- [[architecture]] — estrutura do código
- [[experiment-pipeline]] — pipeline
- [[problem-formulation]] — formulação matemática
