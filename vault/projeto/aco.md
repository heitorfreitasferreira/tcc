---
tags: [projeto, implementacao, aco, ant-colony]
---

# Ant Colony Optimization — Implementação

Implementação em `src/optimization/aco/`. A fundamentação teórica está em [[ant-colony]].

Baseada no **Ant System (AS)** de [[dorigo1996ant]], adaptado com feromônio tridimensional para suportar o tensor de custo.

## Parâmetros

| Parâmetro | Flag CLI | Default | Descrição |
|-----------|----------|---------|-----------|
| População | `--population` | 100 | Nº de formigas por iteração |
| Iterações | `--iterations` | 100 | Nº de iterações |
| $\alpha$ (feromônio) | `--alpha` | 1.0 | Peso do feromônio na decisão |
| $\beta$ (heurística) | `--beta` | 2.0 | Peso da visibilidade na decisão |
| $\rho$ (evaporação) | `--rho` | 0.2 | Taxa de evaporação do feromônio |
| $Q$ (depósito) | `--q` | 100 | Constante de depósito |

## Feromônio 3D

Diferente do ACO clássico (feromônio 2D $\tau_{ij}$), este ACO usa feromônio **tridimensional** $\tau_{ijk}$:

```go
pheromones[prev][curr][next] float64
```

Mesma dimensionalidade do tensor de custo $G[prev][curr][next]$. Inicializado com $\tau_{ijk} = 1.0$ em `src/optimization/aco/main.go:84-96`.

## Algoritmo

`src/optimization/aco/main.go` — `Optimize()`:

1. **Inicializar feromônio**: $\tau_{ijk} = 1.0$ para arestas viáveis ($i \neq j \neq k$)
2. **Para cada iteração**:
   - Cada formiga constrói rota via `walk()`
   - `updatePheromones()` — evaporação global + depósito
3. **Retornar melhor rota encontrada**

## Construção da Rota (Ant Walk)

`src/optimization/aco/ant.go` — `walk()`:

1. Parte do nó 0, escolhe primeiro destino aleatoriamente
2. Para cada passo:
   ```go
   func (aco *ACO) selectNextNode(prev, curr int, visited []bool) int
   ```
   - Para cada candidato $k$ não visitado:
     - Heurística: $\eta = 1 / G[prev][curr][k]$
     - Probabilidade: $p_k = \tau_{prev,curr,k}^\alpha \cdot \eta^\beta$
   - Seleção por roleta proporcional
3. Avalia rota completa via `aco.Makespan()`

## Atualização de Feromônio

`src/optimization/aco/ant.go` — `updatePheromones()`:

1. **Evaporação**: $\tau_{ijk} \leftarrow (1 - \rho) \cdot \tau_{ijk}$ (todas as arestas)
2. **Depósito**: para cada formiga, $\tau_{ijk} \leftarrow \tau_{ijk} + Q / L_k$ (arestas do tour)

## Design Rationale

| Decisão | Alternativas | Por que esta? | Evidência |
|---------|-------------|---------------|-----------|
| **Feromônio 3D** | Feromônio 2D (ACO clássico) | O tensor de custo é 3D (G[prev][curr][next]), então o feromônio precisa da mesma dimensionalidade para modelar dependência de sequência. | [[dorigo1997ant]] usa feromônio 2D para TSP clássico; aqui a variante exige 3D |
| **Seleção por roleta** | Pseudo-aleatória proporcional (ACS) | Roleta é mais exploratória, não requer calibragem do parâmetro q₀. Escolha conservadora para evitar convergência prematura. | [[dorigo1997ant]] mostra que q₀ alto (0.9) favorece explotação; sem tuning, roleta é mais robusta |
| **α=1.0, β=2.0** | α≠1, β≠2 | Valores canônicos. β>α prioriza heurística (distância) sobre feromônio, evitando estagnação. | [[dorigo1996ant]] usa α=1, β=2-5; [[dorigo1997ant]] usa β=2 |
| **ρ=0.2 (evaporação)** | 0.1, 0.5 | Compromisso entre reter memória (0.1) e evitar estagnação (0.5). Valor encontrado por calibragem empírica nas instâncias do estudo. | [[dorigo1997ant]] usa ρ=0.1; aqui o feromônio 3D diluído requer evaporação mais agressiva |
| **Depósito de todas as formigas** | Só a melhor (global-best) | Depósito de todas (estilo AS) mantém diversidade. Global-best acelera convergência mas arrisca estagnação. | [[dorigo1997ant]] usa global-best no ACS; [[stutzle2000mmas]] usa elitismo |
| **Heurística η=1/G[i][j][k]** | η=1/d(i,j) (só distância) | Usar o custo completo do tensor como heurística já embute a penalidade angular na decisão da formiga. | [[problem-formulation]] — tensor já contém distância+ângulo |
| **População 100** | n formigas (m=n) | 100 formigas é consistente com GA/PSO para comparação. Dorigo sugere m≈n, mas isso faria m variar por instância. | [[dorigo1996ant]] sugere m≈n; aqui fixamos m=100 por consistência |

## Detalhes

- Ant System adaptado: roleta proporcional + depósito de todas as formigas
- $\eta = 1/G[i][j][k]$ — já incorpora penalidade angular
- Feromônio é global-best? Não: deposita de todas as formigas (estilo Ant System)

## Limitações Inerentes ao TSP-SD-ATP

### Feromônio 3D e Diluição Estrutural

O ACO clássico usa $\tau_{ij}$ (N² entradas). Este ACO usa $\tau_{ijk}$ (N³ entradas) porque o tensor de custo $G[prev][curr][next]$ **exige** a mesma dimensionalidade — a penalidade angular depende da sequência completa de 3 nós. Não há representação 2D que preserve essa informação sem perda.

| Tamanho | Entradas τ_{ij} (2D) | Entradas τ_{ijk} (3D) | Diluição |
|---------|---------------------|----------------------|----------|
| n=10    | 90                  | 720                  | 8×       |
| n=100   | 9.900               | 970.200              | 98×      |

Consequência: com 100 formigas depositando, cada triple específico recebe ~98× menos depósitos que cada aresta em 2D. A relação sinal-ruído do feromônio é **estruturalmente pior** — não é evitável dentro desta representação.

### Escalabilidade O(N³)

- Memória: N³ floats → 8 MB para n=100, 8 GB para n=1000
- Tempo: evaporação O(N³), seleção O(N) por passo → O(pop × iters × N²) no total
- Viável para n ≤ 100; inviável para n ≥ 1000 sem poda ou lista candidata

Essas limitações não são bugs de implementação — são consequências diretas do TSP-SD-ATP. A seção [[#Problemas de Implementação (BUGs)]] abaixo documenta o que pode e deve ser corrigido.

## Problemas de Implementação (BUGs)

Ver `//BUG` no código-fonte. O próximo agente deve resolver:

1. **`src/graph/types.go:10`** — `maxPenalti` como `time.Duration` (int64) causa type confusion no cast para float64
2. **`src/graph/math.go:28-29`** — TODO antigo reportava valores patológicos ~5e+08; raiz no type confusion acima
3. ~~`src/cmd/aco.go:69` — `--rho` default 0.5 é muito agressivo; literatura usa 0.1-0.3~~ **RESOLVIDO (rho=0.2 desatualizado)**
4. **`src/optimization/aco/ant.go:16-21`** — evaporação varre N³ inteiro em vez de só triplas válidas
5. **`src/optimization/aco/ant.go:58`** — `total == 0` retorna -1 sem fallback; tour incompleto perde exploração
6. **`src/optimization/aco/ant.go:85-90`** — primeiro passo é sorteado uniformemente, sem heurística nem feromônio
7. **`src/optimization/aco/ant.go:96-98`** — tour incompleto (next==-1) contribui 0 feromônio

## Direções de Pesquisa

Ver `//NOTE: (pesquisa)` no código-fonte. Aprimoramentos opcionais:

- **MMAS (Stützle & Hoos, 2000)**: limites [τ_min, τ_max] mitigariam diluição 3D e evitariam estagnação
- **Lista candidata**: nearest-neighbor reduz branching factor de O(N) para O(k)
- **Global-best deposition (ACS)**: em vez de todas as formigas depositarem, só a melhor iteração + melhor global
- **ACO 2D alternativo**: τ_{ij} com heurística que fatora o turn cost separadamente (comparação justa)
- ~~**Parâmetro Gama** (`main.go:12`): declarado mas nunca usado — remover ou implementar~~ **RESOLVIDO: gama removido**

## Conexões

- [[problem-formulation]] — tensor 3D $G$ usado como heurística $\eta = 1/G$
- [[ga]] — outra metaheurística no projeto (contraste: ACO é construtivo, GA é populacional)
- [[pso]] — outra metaheurística no projeto
- [[ant-colony]] — fundamentação teórica
- [[dorigo1997ant]] — ACS, base conceitual
- [[dorigo1996ant]] — Ant System original

## Código-fonte

- `src/optimization/aco/main.go` — `Optimize()`, `new()`, `pathWithoutOrigin()`
- `src/optimization/aco/ant.go` — `ant` struct, `walk()`, `selectNextNode()`, `updatePheromones()`
- `src/cmd/aco.go` — CLI flags e integração
