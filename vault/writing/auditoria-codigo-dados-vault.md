---
title: Auditoria Código-Dados-Vault
tags:
  - writing
  - auditoria
  - codigo
  - dados
  - monografia
status: vault-corrigido-pso-e-aco
created: 2026-06-02
---

# Auditoria Código-Dados-Vault

Esta auditoria verifica claims metodológicos e experimentais do `vault/` contra a fonte determinística do projeto: código em `src/` e dados brutos em `src/data/`.

## Resultado Executivo

O projeto tem base suficiente para escrita, mas o `vault/` contém inconsistências que precisam ser tratadas antes de agentes escreverem afirmações quantitativas definitivas na `monografia/`.

| Item | Status | Decisão para escrita |
|---|---|---|---|
| Formulação TSP-SD-ATP | Confirmada pelo código | Pode ser usada |
| GA | Confirmado pelo código | Pode ser usado |
| PSO | Implementação confirmada, defaults corrigidos | Pode ser usado |
| ACO | Implementação confirmada, notas corrigidas | Pode ser usado |
| Brute-force | Confirmado pelo código | Pode ser usado |
| Lower bound AP | Confirmado pelo código | Pode ser usado como limitante inferior (gap 40–65% vs BF, não como bound justo) |
| Cobertura experimental atual | 4635 summaries (1530 GA, 1530 PSO, 1530 ACO, 30 lowerbound, 15 brute-force) | Confirmado e limpo |
| Claims numéricos em [[resultados]] | Parcialmente desatualizados | Não usar sem recalcular de `src/data/results/` |
| Testes estatísticos | Não confirmados nesta auditoria | Não declarar significância ainda |

## Fonte de Verdade

| Camada | Caminhos auditados | Papel |
|---|---|---|
| Código | `src/graph/`, `src/optimization/`, `src/cmd/`, `src/shared/reporting/` | Implementação e schemas |
| Dados | `src/data/*.points`, `src/data/*.graph`, `src/data/results/{summary,evolution,timing}/` | Instâncias e resultados brutos |
| Vault | `vault/projeto/*.md`, `vault/writing/*.md` | Síntese a ser corrigida/organizada |

## Achados Metodológicos Confirmados

### Formulação e Grafo

| Claim | Status | Evidência primária |
|---|---|---|
| O grafo é um tensor 3D | Confirmado | `src/graph/types.go:8` define `type Graph [][][]float64` |
| A indexação é `[anterior][atual][proximo]` | Confirmado | `src/graph/types.go:8`, `src/graph/makespan.go:13` |
| O custo soma distância euclidiana e penalidade angular | Confirmado | `src/graph/creater.go:24-33` |
| `droneSpeed = 1` | Confirmado | `src/graph/types.go:5` |
| `maxPenalti = 1` | Confirmado | `src/graph/types.go:6` |
| A rota parte do nó 0 e retorna ao nó 0 | Confirmado | `src/graph/makespan.go:7-19` |

### Algoritmo Genético

| Claim | Status | Evidência primária |
|---|---|---|
| Representação por permutação dos nós `1..n-1` | Confirmado | `src/optimization/ga/main.go:57-67` |
| Seleção por torneio | Confirmado | `src/optimization/ga/main.go:130-139` |
| Crossover OX | Confirmado | `src/optimization/ga/main.go:141-194` |
| Mutação swap | Confirmado | `src/optimization/ga/main.go:196-201` |
| Elitismo | Confirmado | `src/optimization/ga/main.go:75-82` |
| Defaults: população 100, iterações 100, elitismo 1, mutação 0.05, torneio 2 | Confirmado | `src/cmd/optimize.go:36-37`, `src/cmd/ga.go:64-66` |

### PSO

| Claim | Status | Evidência primária |
|---|---|---|
| Representação por random keys | Confirmado | `src/optimization/pso/particle.go:23-40` |
| Atualização contínua com inércia, componente cognitivo e componente social | Confirmado | `src/optimization/pso/particle.go:42-54` |
| Defaults atuais: `c1=2.0`, `c2=2.0`, `w=0.7` | Confirmado no código | `src/cmd/pso.go:68-70` |
| Nota [[pso]] registra `c1=c2=1.5` | Divergência | `vault/projeto/pso.md:15-17` |

### ACO

| Claim | Status | Evidência primária |
|---|---|---|
| Feromônio 3D `pheromones[prev][curr][next]` | Confirmado | `src/optimization/aco/main.go:21`, `src/optimization/aco/main.go:84-104` |
| Seleção por roleta proporcional | Confirmado | `src/optimization/aco/ant.go:76-84` |
| Heurística `eta = 1 / G[prev][curr][next]` | Confirmado | `src/optimization/aco/ant.go:52-55` |
| Depósito por todas as formigas | Confirmado | `src/optimization/aco/ant.go:27-39` |
| Defaults atuais: `alpha=1.0`, `beta=2.0`, `rho=0.2`, `q=100` | Confirmado no código | `src/cmd/aco.go:66-69` |
| Nota [[aco]] registra `gama`, `rho=0.1`, `q=1.0` | Divergência | `vault/projeto/aco.md:17-21` |
| Nota [[aco]] chama a implementação de ACS | Impreciso | Código implementa roleta e depósito de todas as formigas, mais próximo de Ant System adaptado |

### Brute-force

| Claim | Status | Evidência primária |
|---|---|---|
| Enumera permutações dos nós `1..n-1` | Confirmado | `src/optimization/brute/main.go:9-23` |
| Geração por algoritmo de Heap | Confirmado | `src/optimization/brute/main.go:55-79` |
| Retorna melhor makespan exato para instâncias executadas | Confirmado para o espaço enumerado | `src/optimization/brute/main.go:20-52` |

### Lower Bound AP

| Claim | Status | Evidência primária |
|---|---|---|
| Reduz tensor 3D para matriz 2D via `c'[j][k] = min_i cost[i][j][k]` | Confirmado | `src/optimization/lowerbound/main.go:42-60` |
| Executa Hungarian O(n³) sobre a matriz reduzida | Confirmado | `src/optimization/lowerbound/hungarian.go` |
| Retorna bound (pode conter subtours) | Confirmado | `src/optimization/lowerbound/main.go:62-89` |
| AP bound ≤ makespan ótimo (por construção) | Confirmado para todas as 15 instâncias BF | Verificado: 10a LB=4.86 ≤ BF=8.25, demais idem |
| Bound é determinístico (~0ms n≤15, 5ms n=100) | Confirmado | Timing: 0ms para instâncias pequenas, 5ms para 100 nós |

## Achados Experimentais

### Cobertura no Diretório

Foram encontrados **4635 arquivos** em `src/data/results/summary/` (após remoção de 1530 resultados ACO legados com `rho=0.5` e `gama=0.1`).

| Método | Quantidade |
|---|---|---:|
| GA | 1530 |
| PSO | 1530 |
| ACO | 1530 |
| Lowerbound | 30 |
| Brute-force | 15 |

Cobertura: 30 instâncias, 51 sementes para GA/PSO/ACO, lowerbound determinístico nas 30 instâncias, brute-force em 15 instâncias pequenas (`10a` a `14c`).

### Baseline Brute-force Confirmada

| Instância | Makespan ótimo |
|---|---:|
| 10a | 8.251282810997422 |
| 10b | 7.906130315558171 |
| 10c | 9.545033907523141 |
| 11a | 9.493457779404203 |
| 11b | 8.823027373557352 |
| 11c | 10.260221738390818 |
| 12a | 9.500700789073989 |
| 12b | 9.768963312506418 |
| 12c | 10.47963364504159 |
| 13a | 10.950258153646407 |
| 13b | 11.473574261140902 |
| 13c | 10.041955396200756 |
| 14a | 11.55692517246078 |
| 14b | 11.466929014715355 |
| 14c | 11.486447067382262 |

### Gap nas Instâncias com Brute-force

Resultados calculados apenas no subconjunto compatível com o código atual.

| Método | Runs | Gap médio | Gap mínimo | Gap máximo | Taxa global de ótimo |
|---|---:|---:|---:|---:|---:|
| ACO | 765 | 0.1836% | 0.0000% | 3.7769% | 76.86% |
| GA | 765 | 4.3564% | 0.0000% | 31.5704% | 29.41% |
| PSO | 765 | 19.0226% | 0.0000% | 57.8884% | 5.10% |

> [!warning] Divergência com [[resultados]]
> A nota [[resultados]] afirma gap médio de ACO igual a 0.00% e ótimo em 100% para `10a-13c`. Isso não é sustentado pelos dados atuais: ACO falha em algumas instâncias pequenas, como `10b`, `11b`, `12c`, `13a`, `13b`, `14a`, `14b` e `14c`.

### Instâncias Grandes

Resultados do subconjunto compatível com o código atual.

| Instância | ACO best | ACO média | GA best | GA média | PSO best | PSO média |
|---|---:|---:|---:|---:|---:|---:|
| 50a | 24.6688 | 25.6475 | 41.0120 | 45.3925 | 53.7651 | 59.7284 |
| 50b | 26.3528 | 27.4843 | 42.6487 | 48.6345 | 57.0330 | 64.3194 |
| 50c | 26.2372 | 27.0062 | 42.0122 | 46.9438 | 54.4114 | 62.3455 |
| 100a | 42.6718 | 45.3528 | 97.5921 | 110.6608 | 127.4034 | 136.5131 |
| 100b | 45.4689 | 48.0596 | 97.5343 | 110.4959 | 126.3967 | 135.9987 |
| 100c | 45.4858 | 46.9487 | 95.1632 | 111.3243 | 126.1838 | 138.1908 |

### Lower Bound Validado vs Brute-force (15 instâncias)

AP bound ≤ BF ótimo em 100% dos casos (validade formal). Gap médio: ~50%, mínimo 40.8% (`11c`), máximo 65.3% (`14a`). O bound é frouxo para o TSP-SD-ATP — a redução 3D→2D perde informação angular significativa.

| Instância | BF ótimo | LB bound | Gap LB→BF |
|---|---|---|---:|---:|
| 10a | 8.251283 | 4.859238 | 41.11% |
| 10b | 7.906130 | 2.979065 | 62.32% |
| 10c | 9.545034 | 4.956950 | 48.07% |
| 11a | 9.493458 | 4.529660 | 52.29% |
| 11b | 8.823027 | 4.158859 | 52.86% |
| 11c | 10.260222 | 6.075174 | 40.79% |
| 12a | 9.500701 | 4.549388 | 52.12% |
| 12b | 9.768963 | 4.091362 | 58.12% |
| 12c | 10.479634 | 4.088736 | 60.98% |
| 13a | 10.950258 | 5.799255 | 47.04% |
| 13b | 11.473574 | 6.078440 | 47.02% |
| 13c | 10.041955 | 5.347435 | 46.75% |
| 14a | 11.556925 | 4.004625 | **65.35%** |
| 14b | 11.466929 | 5.675375 | 50.51% |
| 14c | 11.486447 | 6.420654 | 44.10% |

> [!warning] Bound frouxo para TSP-SD-ATP
> O roadmap [[roadmap-monografia]] sugere que "gap < 30%" para instâncias pequenas, mas o AP bound via redução 3D→2D apresentou gap entre 40% e 65%. A perda de informação ao tomar `min_i` por sobre o tensor 3D elimina o contexto de curva. O bound é válido, mas não é tight. Na monografia, deve ser descrito como limitante inferior fraco, não como referência justa.

### Lower Bound Timing

| Instância | optimize_ms |
|---|---:|
| 10a..14c | 0 |
| 15a..50c | 0–1 |
| 100a..100c | 5 |

### Tempo de Otimização em n=100 (metaheurísticas)

Tempos médios em milissegundos, lidos por `timing_file` a partir dos summaries compatíveis com o código atual.

| Instância | ACO média ms | GA média ms | PSO média ms | Razão ACO/GA | Razão ACO/PSO |
|---|---|---|---:|---:|---:|---:|---:|
| 100a | 4262.06 | 38.82 | 76.24 | 109.78x | 55.91x |
| 100b | 4276.55 | 36.27 | 75.25 | 117.89x | 56.83x |
| 100c | 4266.55 | 39.37 | 73.84 | 108.37x | 57.78x |

> [!warning] Divergência com [[resultados]]
> A nota [[resultados]] registra ACO em `100a` como 2.859s e razão GA/ACO de ~75x. Nos dados atuais, `100a` tem ACO médio de 4.262s e razão ACO/GA de ~110x.

## Divergências Prioritárias no Vault

| Nota | Divergência | Ação recomendada |
|---|---|---|---|
| ~~[[pso]] | Defaults `c1=c2=1.5`, mas código usa `2.0`~~ | **RESOLVIDO** |
| ~~[[aco]] | Registra `gama`, `rho=0.1`, `q=1.0`; código usa `rho=0.2`, `q=100`, sem `gama`~~ | **RESOLVIDO** |
| ~~[[aco]] | Chama implementação de ACS, mas código tem roleta e depósito por todas as formigas~~ | **RESOLVIDO** |
| [[resultados]] | Números não batem com os dados atuais | Recalcular (parcialmente atualizado, verificar) |
| [[analysis-methodology]] | Testes estatísticos aparecem como previstos/implementados de forma ambígua | Confirmar via script/notebook antes de escrever significância |
| [[roadmap-monografia]] | Claim "AP bound gap < 30%" contradito pelos dados (gap real 40–65%) | Corrigir claim no roadmap; bound é válido mas frouxo |

## Regras Para Agentes Escritores

1. Para Proposta, usar diretamente o código em `src/` e não confiar apenas em `vault/projeto/*.md`.
2. Os parâmetros documentados agora correspondem ao código atual; confiar em `vault/projeto/pso.md` e `vault/projeto/aco.md`.
3. Não usar os números de [[resultados]] sem recalcular a partir de `src/data/results/summary/` e `timing_file` (parcialmente atualizado, mas verificar cada número).
4. Não declarar significância estatística até que o protocolo estatístico esteja auditado.
5. Se o agente atualizar uma nota do vault, deve registrar qual arquivo do código ou dado bruto sustenta a correção.

## Próximos Passos

1. Validar [[resultados]] contra dados brutos e executar protocolo estatístico.
2. Atualizar [[roadmap-monografia]] com dados reais do AP bound (gap 40–65%, não < 30%).
3. Incluir lower bound na análise de Experimentos como referência adicional (não como bound tight).
