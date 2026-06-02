---
title: "Ant Colonies for the Travelling Salesman Problem"
authors: [Dorigo, Marco]
year: 1997
doi: "10.1016/S0303-2647(97)01708-5"
bibtex-key: dorigo1997ant
tags: [tsp]
status: lido
rating: 5
pdf: "papers/pdfs/dorigo1997ant.pdf"
---

## PDF

[[papers/pdfs/dorigo1997ant.pdf]]

## Resumo

O artigo apresenta o Ant Colony System (ACS), uma evolução significativa do Ant System para o Problema do Caixeiro Viajante (TSP). O ACS introduz três inovações principais: (1) uma regra de transição pseudo-aleatória proporcional que equilibra exploração e explotação, (2) atualização global de feromônio que recompensa apenas o melhor tour, e (3) atualização local de feromônio que desencoraja convergência prematura ao reduzir feromônio nas arestas visitadas. O ACS demonstra desempenho superior a GA, SA, EP e redes neurais elásticas em benchmarks TSP, encontrando soluções ótimas ou próximas do ótimo para problemas de até 1577 cidades usando listas candidatas.

## Contribuições Principais

- Proposição do Ant Colony System (ACS) com regra de transição pseudo-aleatória proporcional (q₀)
- Atualização global de feromônio restrita à melhor solução (diferente do AS que usa todas as formigas)
- Atualização local de feromônio para promover diversidade (diminui feromônio nas arestas percorridas)
- Uso de listas candidatas (candidate list) para escalabilidade em problemas grandes
- Demonstração de eficácia em ATSP (problemas assimétricos), superando códigos exatos estado-da-arte
- Comparação sistemática com GA, EP, SA, AG e redes neurais em benchmarks TSPLIB

## Relevância para o TCC

O ACS é a versão do ACO implementada no código de otimização do TCC (referenciada como "aco" nos experimentos). O artigo mostra que o ACS encontra soluções de alta qualidade para TSP com complexidade O(n²·t), diretamente comparável aos outros métodos (GA, PSO) do estudo. A capacidade de lidar com TSP assimétrico é relevante para cenários de patrulha onde custos de deslocamento podem variar com direção (vento, elevação). As listas candidatas são uma técnica de otimização prática para instâncias maiores.

## Métodos e Abordagens

- Regra de transição: pseudo-random-proportional (q ≤ q₀ → explota; senão → explora)
- Atualização global: τ(r,s) ← (1-α)·τ(r,s) + α·Δτ(r,s), apenas para o melhor tour
- Atualização local: τ(r,s) ← (1-α)·τ(r,s) + α·τ₀ (reduz feromônio em arestas visitadas)
- Lista candidata: cl cidades mais próximas (cl=20 tipicamente) para reduzir branching factor
- Parâmetros típicos: m=10, β=2, α=0.1, q₀=0.9, τ₀=(n·L_nn)^(-1)
- Benchmarks: Oliver30, Eil51, Eil76, KroA100, d198, pcb442, att532, rat778, fl1577 (TSPLIB)
- Problemas ATSP: ry48p, 43×2 (encontra ótimo em 220s vs. 32h de método exato)

## Conexões

- [[dorigo1996ant]] — Ant System, base conceitual deste trabalho
- [[wang2021ant]] — SOS-ACO (evolução do ACO para TSP)
- [[lin1973effective]] — TSP, heurísticas clássicas de referência; ACS usa 3-opt como pós-otimização
- [[ant-colony]]
- [[rajwar2023exhaustive]] — classifica ACO como método canônico

## Notas e Insights

- O ACS resolve Oliver30 em apenas 830 tours (vs. 1830 do GA, 40.000 do SA, 325.000 do EP)
- Em problemas grandes (fl1577), erro médio de ~3.5% — notável para método geral sem ajuste fino
- A atualização LOCAL de feromônio (diminuir trilha) é contraintuitiva mas crucial para diversidade
- O ACS é comparável ao Lin-Kernighan em qualidade, mas mais lento; sua vantagem é a adaptabilidade a variações do problema
- A regra pseudo-aleatória proporcional é uma contribuição elegante que influenciou aprendizado por reforço
- Limitação: sensibilidade a parâmetros (q₀, α, β, tamanho da lista candidata)
- A abordagem de comunicação mediada por feromônio é diretamente aplicável a enxames de drones

## Citações-chave

> "ACS is capable of generating good solutions to both symmetric and asymmetric instances of the TSP."

> "Although when applied to the symmetric TSP ACS is not competitive with specialized heuristic methods like Lin-Kernighan, its performance can become very interesting when applied to a slightly different problem."
