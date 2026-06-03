---
title: "Ant Colony Optimization"
authors: [Dorigo, Marco, Stützle, Thomas]
year: 2004
doi: ""
bibtex-key: dorigo2004book
tags: [aco survey book]
status: lido
rating: 5
pdf: "papers/pdfs/dorigo2004book.pdf"
---

## PDF

[[papers/pdfs/dorigo2004book.pdf]]

## Resumo

Livro referência fundamental sobre ACO, publicado pelo MIT Press. Cobre: fundamentos biológicos, a metaheurística ACO, implementações para TSP e outros problemas NP-difíceis, variantes (AS, EAS, RAS, MMAS, ACS), aplicações em roteamento em redes (AntNet), aspectos teóricos (convergência, parâmetros), e diretrizes para projeto de algoritmos ACO. É a obra mais completa e citada sobre o tema.

## Contribuições Principais

- Framework completo e unificado da metaheurística ACO
- Guia prático de implementação com pseudocódigo detalhado
- Análise de parâmetros (α, β, ρ, τ₀, m) e seu impacto no desempenho
- Estudos de caso: TSP, QAP, VRP, scheduling, AntNet
- Revisão de aspectos teóricos (convergência, parâmetros de controle)

## Relevância para o TCC

Referência máxima para justificar e implementar ACO. Contém todas as diretrizes de parametrização (α=1, β=2-5, ρ=0.1-0.5, τ₀ = 1/(n·L_nn)) e análise de convergência. Essencial para a seção de metodologia.

## Métodos e Abordagens

- Ant System (AS): τ_ij ← (1-ρ)·τ_ij + ΣΔτ_ij^k
- Elitist AS, Rank-based AS (RAS)
- Ant Colony System (ACS): pseudo-random proportional rule + local pheromone update
- MAX-MIN Ant System (MMAS): bounds [τ_min, τ_max]
- AntNet: aplicação a roteamento dinâmico em redes
- Daemon actions: busca local (2-opt, LK) como pós-processamento

## Conexões

- [[dorigo1996ant]] — Ant System (artigo original)
- [[dorigo1997ant]] — ACS (artigo original)
- [[dorigo2005acotheory]] — teoria ACO
- [[blum2005acointro]] — introduction and trends
- [[stutzle2000mmas]] — MMAS (artigo original)
- [[ant-colony]]
- [[tsp]]
