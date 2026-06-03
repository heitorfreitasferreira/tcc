---
title: "Ant Colony Optimization Theory: A Survey"
authors: [Dorigo, Marco, Blum, Christian]
year: 2005
doi: "10.1016/j.tcs.2005.05.020"
bibtex-key: dorigo2005acotheory
tags: [aco survey theory]
status: lido
rating: 4
pdf: "papers/pdfs/dorigo2005acotheory.pdf"
---

## PDF

[[papers/pdfs/dorigo2005acotheory.pdf]]

## Resumo

Survey dos resultados teóricos sobre Ant Colony Optimization. Aborda: (1) provas de convergência de algoritmos ACO (Gutjahr, 2000 — convergência em probabilidade para o ótimo global), (2) relações entre ACO e Model-Based Search (MBS), Cross-Entropy (CE), Estimation of Distribution Algorithms (EDAs), (3) análise de comportamento (parâmetros, feromônio, estagnação). O artigo mostra que ACO é um caso particular de Model-Based Search, onde a distribuição de probabilidade sobre soluções é iterativamente atualizada.

## Contribuições Principais

- Conexão formal entre ACO e Model-Based Search (MBS, CE, EDAs)
- Revisão das provas de convergência dos algoritmos ACO
- Análise do papel dos parâmetros (α, β, ρ) no comportamento do ACO
- Framework Hyper-Cube (HCF) como generalização

## Relevância para o TCC

Suporte teórico para justificar a escolha do ACO. Permite afirmar que (1) ACO converge em probabilidade para o ótimo, (2) tem fundamentos alinhados com busca baseada em modelo, (3) os parâmetros têm interpretação clara.

## Métodos e Abordagens

- Graph-based Ant System (GBAS) e GBAS*/GBAS† (algoritmos com prova de convergência)
- Hyper-Cube Framework (HCF): normalização do feromônio para [0,1]
- Relação com Cross-Entropy (CE) e Estimation of Distribution Algorithms (EDAs)
- Análise de estagnação: quando τ_max/τ_min → ∞

## Conexões

- [[dorigo1996ant]] — Ant System (ACO original)
- [[dorigo1997ant]] — ACO aplicado ao TSP (ACS)
- [[blum2005acointro]] — introdução e variantes práticas
- [[stutzle2000mmas]] — MMAS (convergência + limites de feromônio)
- [[ant-colony]]
- [[tsp]]
