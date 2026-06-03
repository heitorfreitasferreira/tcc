---
title: "Ant Colony Optimization: Introduction and Recent Trends"
authors: [Blum, Christian]
year: 2005
doi: "10.1016/j.plrev.2005.10.001"
bibtex-key: blum2005acointro
tags: [aco survey]
status: lido
rating: 5
---

## Resumo

Uma introdução completa ao ACO e suas variantes mais bem-sucedidas. Apresenta: (1) inspiração biológica (forrageamento de formigas), (2) Ant System (AS) aplicado ao TSP, (3) a metaheurística ACO formalizada, (4) variantes: Elitist AS, Rank-based AS, MAX-MIN Ant System (MMAS) e Ant Colony System (ACS), (5) resultados teóricos de convergência, (6) extensão para otimização contínua, (7) hibridização com AI/OR.

## Contribuições Principais

- Comparação sistemática das variantes ACO (AS, EAS, RAS, MMAS, ACS)
- Framework unificado da metaheurística ACO com componentes identificáveis
- Discussão de convergência, parâmetros e quando usar cada variante
- Cobertura de aplicações além do TSP (routing em redes, scheduling, etc.)

## Relevância para o TCC

Artigo ideal para a seção de fundamentação teórica do ACO. Explica de forma clara as diferenças entre variantes (AS, ACS, MMAS) e fornece diretrizes para escolha — MMAS é robusto para instâncias grandes, ACS é mais rápido para convergência.

## Métodos e Abordagens

- Ant System (AS): todas as formigas depositam feromônio
- Elitist AS (EAS): reforço extra para a melhor solução
- Rank-based AS (RAS): top-k formigas depositam
- MAX-MIN Ant System (MMAS): limites [τ_min, τ_max]
- Ant Colony System (ACS): regra pseudo-aleatória + atualização local
- Busca local 2-opt/3-opt como componente opcional

## Conexões

- [[dorigo1996ant]] — Ant System (AS)
- [[dorigo1997ant]] — ACS (Ant Colony System)
- [[dorigo2005acotheory]] — teoria ACO
- [[stutzle2000mmas]] — MMAS
- [[dorigo2004book]] — livro referência
- [[ant-colony]]
- [[tsp]]
