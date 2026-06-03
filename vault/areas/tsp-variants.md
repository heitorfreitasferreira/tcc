---
tags: [area, variante-tsp, classificacao]
created: 2026-06-02
updated: 2026-06-02
---

# Variantes do Problema do Caixeiro Viajante (TSP)

## Classificação Geral

O TSP clássico (**Problema do Caixeiro Viajante**, do inglês *Traveling Salesman Problem*) serve como base para dezenas de variantes que relaxam ou adicionam restrições. A taxonomia abaixo organiza as principais:

### Núcleo: Matriz de Custo 2D

**TSP Simétrico**: $C[i][j] = C[j][i]$, distância euclidiana. NP-difícil [[garey1979computers]].

**TSP Assimétrico (ATSP)**: $C[i][j] \neq C[j][i]$, arestas direcionais. Mais geral que o TSP simétrico; problemas de roteamento urbano com mão única ou tempos variáveis por direção.

**TSP com Lucros** (*TSP with Profits*):
  - **PCTSP** (*Prize-Collecting TSP*): pode pular nós pagando penalidade [[bock2025survey]]
  - **Orienteering Problem**: maximizar pontuação coletada dentro de orçamento de distância
  - **Quota TSP**: coletar cota mínima de prêmio minimizando distância

**TSP com Janelas de Tempo (TSPTW)**: cada nó deve ser visitado dentro de intervalo [$a_i, b_i$].

**TSP dependente do tempo (TDTSP)**: custo da aresta $(i,j)$ varia com o instante de partida.

### Clusterização

**GTSP** (*Generalized TSP*): nós particionados em clusters; visita-se exatamente um nó por cluster [[pop2024comprehensive]].

**CTSP** (*Clustered TSP*): nós em clusters; visitam-se todos os nós, mas nós do mesmo cluster devem ser consecutivos na rota.

**TSP com Agrupamento Hierárquico**: combina GTSP com precedência entre clusters.

### Dependência de Sequência e Curva

**SDTSP** (*Sequence-Dependent TSP*): custo de $(j \to k)$ depende do nó visitado imediatamente antes de $j$. Comum em problemas de *setup* em manufatura e scheduling de máquinas.

**TSP com Custos de Curva (*TSP with Turn Costs*)**: o custo de transição incorpora penalidade angular pela mudança de direção ao passar por um nó.

  - **AM-TSP** (*Angular-Metric TSP*): função objetivo é **apenas** a soma dos ângulos de curva; distância não é contabilizada. Proposto por Fekete & Krupke (2017) no contexto de cobertura com robôs aéreos. Existe PTAS para o TSP euclidiano, mas apenas $O(\log n)$-aproximação para o AM-TSP.

### Drones e Veículos Híbridos

**FSTSP** (*Flying Sidekick TSP*): caminhão e drone cooperam em entregas [[murray2015flying], [agatz2018optimization], [dellamico2022exact]].

**TSP-D** (*TSP with Drone*): generalização do FSTSP, com variações de múltiplos drones [[dellamico2021multiple]].

**Patrulha com VANT**: roteamento de drone para vigilância/perímetro, frequentemente com horizonte finito e incerteza [[rajan2022routing]].

### Outras

**Bottleneck TSP**: minimizar a maior aresta do tour.

**Maximum Scatter TSP**: maximizar a menor aresta (roteamento seguro/distribuído).

**Covering TSP**: visitar pontos que cubram uma região (ex.: varredura com sensor).

**TSP com Múltiplos Caixeiros (mTSP)**: $m$ caixeiros partindo de depósito comum.

---

## Variante do TCC — TSP-SD-ATP

A **variante implementada neste projeto** combina distância euclidiana + penalidade angular via tensor 3D pré-computado. Difere do AM-TSP (que minimiza **só ângulo**) ao somar ambos os componentes.

> Documentação completa da implementação em [[problem-formulation]].

---

## Papers no Vault Relacionados a Cada Variante

| Variante | Papers |
|----------|--------|
| **TSP clássico** | [[lawler1985traveling]], [[applegate2006traveling]], [[garey1979computers]] |
| **GTSP** | [[pop2024comprehensive]] |
| **TSP com GA** | [[potvin1996ga]], [[larranaga1999ga]], [[nagata2006eax]], [[hga2024hybrid]] |
| **TSP com PSO** | [[clerc2000discretepso]], [[araujo2025pso]], [[kappagantula2025dpso]], [[sun2024hybrid]], [[huang2025matrix]] |
| **TSP com ACO** | [[dorigo1997ant]], [[stutzle2000mmas]], [[dorigo2004book]], [[dorigo2005acotheory]], [[blum2005acointro]], [[deepaco2023]], [[neufaco2025]], [[ppaco2024]], [[gpaco2025]], [[wang2021ant]] |
| **Turn costs / Angular** | [[winter2002modeling]], [[vanhove2012route]]; ver também [[problem-formulation]] (implementação) |
| **Drone routing** | [[murray2015flying]], [[agatz2018optimization]], [[dellamico2022exact]], [[dellamico2021multiple]], [[freitas2020vns]], [[rajan2022routing]], [[ahmed2024receding]] |
| **Survey TSP** | [[bock2025survey]], [[pop2024comprehensive]], [[halim2019combinatorial]] |
| **Comparativos** | [[wu2020comparative]], [[haroun2015performance]], [[chandra2022comparative]], [[almufti2025comparative]], [[wadi2025charting]], [[alexander2020comparison]], [[hossain2024comparison]], [[toaza2023review]] |

---

## Referências Externas Citadas

- Caldwell, T. (1961). *On Finding Minimum Routes in a Network With Turn Penalties*. CACM 4(2). — Trabalho pioneiro em penalidades de curva.
- Fekete, S.P. & Krupke, D. (2017). *Covering Tours with Turn Cost: Variants, Approximation and Practical Solution*. EuroCG 2017. — Define o AM-TSP (Angular-Metric TSP).
- Geisberger, R. & Vetter, C. (2011). *Efficient Routing in Road Networks with Turn Costs*. SEA'11. — Contraction hierarchies com turn costs.

---

## Conexões

- [[tsp]] — nota principal sobre o TSP clássico
- [[problem-formulation]] — implementação da variante TSP-SD-ATP no projeto
- [[routing]] — classe geral de problemas de roteamento
- [[drone-routing]] — roteamento de drones (contexto de aplicação)
- [[bio-inspired-optimization]] — metaheurísticas usadas para resolver a variante
- [[ant-colony]], [[genetic-algorithms]], [[particle-swarm]] — métodos da literatura
- [[comparative-studies]] — comparação entre métodos na variante
