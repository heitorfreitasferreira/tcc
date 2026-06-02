---
title: "Automated Design of State Transition Rules in Ant Colony Optimization by Genetic Programming"
authors: [Liu, Xiao, Chen, Yu, Zhang, Meng]
year: 2025
doi: "10.1007/s12293-025-00435-9"
bibtex-key: gpaco2025
pdf: "papers/pdfs/gpaco2025.pdf"
tags: [aco metaheuristic gp]
status: lido
rating: 4
---

## PDF

[[papers/pdfs/gpaco2025.pdf]]

## Resumo

Este estudo usa Programação Genética (GP) para projetar automaticamente regras de transição de estado em ACO. GP-ACO evolui regras de transição eficazes sem conhecimento de especialistas. Cinco questões investigadas: generalidade do GP-ACO, impacto de variantes ACO, efeito de busca local 2-opt, impacto de informação global adicional, e interpretabilidade. xGP-ACO (com informação global) supera até heurísticas projetadas por LLMs (ReEvo).

## Contribuições Principais

- Demonstração de que GP pode projetar regras de transição ACO competitivas com heurísticas manuais
- xGP-ACO com informação global supera LLM-designed heuristics (ReEvo)
- Análise de interpretabilidade: GP produz regras mais compreensíveis que redes neurais
- Generalidade robusta: regras evoluídas em uma instância funcionam bem em outras

## Relevância para o TCC

Abordagem alternativa de hibridização (GP + ACO) interessante para discussão de trabalhos relacionados. Diferente de neural-enhanced ACO, GP-ACO produz regras interpretáveis. Útil para contrastar com DeepACO/NeuFACO na seção de estado-da-arte.

## Métodos e Abordagens

- Programação Genética (árvores) para representar regras de transição
- Conjunto de terminais: distância, feromônio, informação global (nº de arestas, etc.)
- Variantes ACO testadas: AS, MMAS, ACS como base
- xGP-ACO: adiciona estatísticas globais como terminais
- Busca local 2-opt como pós-processamento
- Comparação com ReEvo (LLM-designed heuristics) e ACO clássico

## Conexões

- [[dorigo1996ant]] — Ant System (AS)
- [[dorigo1997ant]] — ACO para TSP
- [[blum2005acointro]] — introdução ao ACO (contexto)
- [[deepaco2023]] — DeepACO (neural-enhanced, contrasta com GP)
- [[lin1973effective]] — 2-opt usado como busca local
- [[ant-colony]]
