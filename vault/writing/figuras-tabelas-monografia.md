---
title: Figuras e Tabelas da Monografia
tags:
  - writing
  - monografia
  - figuras
  - tabelas
status: inicial
created: 2026-06-02
---

# Figuras e Tabelas da Monografia

Catálogo de todos os artefatos visuais disponíveis em `monografia/figs/` e tabelas geradas por `scripts/consolidate_results.py`. Organizado por capítulo.

Referência cruzada com as figuras/tabelas mínimas listadas em [[roadmap-monografia#Capítulo-4-—-Experimentos-e-Resultados]].

---

## Capítulo 1 — Introdução

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig1 | ilustração | Diagrama do cenário de patrulha com drone (a criar) | — | Pendente |

> Nenhuma figura obrigatória para Introdução. Opcional: figura conceitual do cenário.

---

## Capítulo 2 — Fundamentação Teórica

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig2 | diagrama | Penalidade angular: ângulo de virada entre dois segmentos | `diagram-angular-penalty.{png,svg}` | Pronto |
| Fig3 | diagrama | Tensor 3D de custo G[prev][curr][next] | `diagram-tensor-3d.{png,svg}` | Pronto |
| Fig4 | fluxograma | Algoritmo Genético (GA) | `flowchart-ga.{tex,pdf,png,svg}` | Pronto |
| Fig5 | fluxograma | PSO | `flowchart-pso.{tex,pdf,png,svg}` | Pronto |
| Fig6 | fluxograma | ACO | `flowchart-aco.{tex,pdf,png,svg}` | Pronto |
| Fig7 | fluxograma | Busca exaustiva (brute-force) | `flowchart-bruteforce.{tex,pdf,png,svg}` | Pronto |
| Fig8 | fluxograma | Lower bound AP (Hungarian) | `flowchart-lowerbound.{tex,pdf,png,svg}` | Pronto |

> Fig2–3 na seção de TSP-SD-ATP e custos de curva. Fig4–8 na seção de métodos.

---

## Capítulo 3 — Proposta

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| — | — | Nenhuma figura específica (a formulação usa equações, não figuras) | — | — |

> Opcional: incluir Fig2–3 da Fundamentação se não usadas lá, ou referenciar como "conforme visto na Seção 2".

---

## Capítulo 4 — Experimentos e Resultados

### Seção: Configuração e Baseline

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| T1 | tabela | Instâncias: nome, n, variante, nº POIs | gerada por `consolidate_results.py` | Gerável |
| T2 | tabela | Ótimos brute-force (10a..14c) com makespan | gerada por `consolidate_results.py` | Gerável |
| T3 | tabela | Parâmetros dos métodos (pop, iter, defaults) | manual | Rascunho |
| T4 | tabela | AP bound vs brute-force: bound, ótimo, gap % (10a..14c) | gerada por `consolidate_results.py` | Gerável |

### Seção: Qualidade das Soluções

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig9 | heatmap | Makespan mediano por instância×método | `heatmap-makespan-median.{png,svg}` | Pronto |
| Fig10 | heatmap | Gap % vs brute-force (instâncias 10a..14c) | `heatmap-gap-vs-bf.{png,svg}` | Pronto |
| Fig11 | heatmap | Taxa de acerto (% sementes que encontram ótimo) | `heatmap-success-rate.{png,svg}` | Pronto |
| Fig12 | boxplot | Distribuição de qualidade (makespan) por método e instância | `boxplot-estabilidade.{png,svg}` | Pronto |
| Fig13 | gráfico | Distribuição de qualidade agregada (método × n) | `fig-quality-distribution.{png,svg}` | Pronto |
| T5 | tabela | Gap médio, min, max por método (instâncias com BF) | gerada por `consolidate_results.py` | Gerável |
| T6 | tabela | Makespan médio por método e instância (grandes: 50a..100c) | gerada por `consolidate_results.py` | Gerável |
| T7 | tabela | Taxa de acerto por método e instância (10a..14c) | gerada por `consolidate_results.py` | Gerável |

### Seção: Tempo Computacional

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig14 | gráfico | Escalabilidade: tempo de otimização por n (escala log) | `scalability-runtime.{png,svg}` | Pronto |
| Fig15 | scatter | Trade-off qualidade×tempo (média por método×instância) | `scatter-quality-vs-time.{png,svg}` | Pronto |
| T8 | tabela | Tempo médio (ms) por método e instância (n=50, 100) | gerada por `consolidate_results.py` | Gerável |
| T9 | tabela | Razão ACO/GA e ACO/PSO em n=100 | gerada por `consolidate_results.py` | Gerável |

### Seção: Convergência

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig16 | curva | Convergência: fração de avaliações × gap mediano (3 métodos) | `fig-runtime-convergence.{png,svg}` | Pronto |
| Fig17 | curva | Última iteração com melhoria por método | `convergence-last-improvement.{png,svg}` | Pronto |
| Fig18 | overlay | Curvas de convergência sobrepostas (10a, 30a, 50a, 100a) | `convergence-overlay-{10a,30a,50a,100a}.{png,svg}` | Pronto |

### Seção: Rotas

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| Fig19 | rota | Rota ótima BF + rotas GA/PSO/ACO/LB para 10a | `overlay-10a-methods.png` | Pronto |
| Fig20 | rota | Rota ótima BF + rotas GA/PSO/ACO/LB para 10b | `overlay-10b-methods.png` | Pronto |
| Fig21 | rota | Rota ótima BF + rotas GA/PSO/ACO/LB para 10c | `overlay-10c-methods.png` | Pronto |
| Fig22 | rota | Rotas GA/PSO/ACO/LB para 14a | `overlay-14a-methods.png` | Pronto |
| Fig23 | rota | Rotas GA/PSO/ACO/LB para 15a | `overlay-15a-methods.png` | Pronto |
| Fig24 | rota | Rotas GA/PSO/ACO/LB para 20a | `overlay-20a-methods.png` | Pronto |
| Fig25 | rota | Rotas GA/PSO/ACO/LB para 30a | `overlay-30a-methods.png` | Pronto |
| Fig26 | rota | Rotas GA/PSO/ACO/LB para 50a | `overlay-50a-methods.png` | Pronto |
| Fig27 | rota | Rotas GA/PSO/ACO/LB para 100a | `overlay-100a-methods.png` | Pronto |
| Fig28 | painel | Painel multi-instância (10a, 30a, 100a) com 4 métodos cada | `route-panel-{10a,30a,100a}.png` | Pronto |
| Fig29 | tex | Small multiples: rotas de todas as instâncias | `fig-route-small-multiples.tex` | Pronto |

> Selecionar 3–5 overlays representativos para o corpo da monografia; os demais podem ir para apêndice.

### Seção: Estatística

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| T10 | tabela | Resultado do teste de Friedman (χ², p-valor) | gerada por `scripts/analise-estatistica.py` | Pendente |
| T11 | tabela | Matriz de p-valores Nemenyi ou Wilcoxon+Bonferroni | gerada por `scripts/analise-estatistica.py` | Pendente |
| Fig30 | diagrama | Diagrama de diferença crítica (CD) | `monografia/figs/diagrama-cd.{png,svg}` | Pendente |

> A seção de estatística e as tabelas T10–T11 e Fig30 dependem da execução de `scripts/analise-estatistica.py`.

---

## Capítulo 5 — Conclusão

| ID | Tipo | Descrição | Arquivo | Status |
|----|------|-----------|---------|--------|
| — | — | Nenhuma figura ou tabela nova | — | — |

> Conclusão não apresenta novos resultados. Pode referenciar figuras dos capítulos anteriores.

---

## Resumo por Status

| Status | Quantidade |
|--------|-----------:|
| Pronto (PNG+SVG) | 20+ figuras de publicação + 30+ rotas + 6 fluxogramas + 2 diagramas conceituais |
| Gerável (script existe) | Tabelas T1–T9 via `consolidate_results.py` |
| Pendente (a criar) | T10, T11, Fig30 (dependem de `scripts/analise-estatistica.py`); Fig1 (cenário de patrulha) |
| A selecionar | Escolher ~5 overlays de rota para corpo, restante para apêndice |

## Legenda de Nomenclatura

- **heatmap-***: mapas de calor 2D (instância × método)
- **scatter-***: gráfico de dispersão
- **boxplot-***: gráfico de caixa
- **convergence-***: curvas de evolução do makespan
- **scalability-***: tempo × tamanho da instância
- **overlay-***: rotas sobrepostas no mesmo mapa
- **route-panel-***: painéis comparativos lado a lado
- **diagram-***: conceitos abstratos (tensor, ângulo)
- **flowchart-***: diagrama de fluxo do algoritmo
- **graph-***-bare: grafo sem rota sobreposta
- **fig-***: figuras de alta qualidade para publicação
