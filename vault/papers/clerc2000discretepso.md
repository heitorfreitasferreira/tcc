---
title: "Discrete Particle Swarm Optimization, Illustrated by the Traveling Salesman Problem"
authors: [Clerc, Maurice]
year: 2000
doi: ""
bibtex_key: clerc2000discretepso
bibtex-key: clerc2000discretepso
pdf: "papers/pdfs/clerc2000discretepso.pdf"
tags: [pso tsp survey]
status: lido-parcial
rating: 5
---

## PDF

![[clerc2000discretepso.pdf]]

## Resumo

Clerc propõe, pelo título e enquadramento local, uma adaptação discreta do [[particle-swarm]] para o [[tsp]], problema em que a representação contínua original do PSO não se aplica diretamente. O PDF local tem 18 páginas, mas não forneceu texto extraível; portanto, a descrição detalhada dos operadores e experimentos não pôde ser validada diretamente pelo corpo do PDF.

## Contribuições Principais

- Trata explicitamente a adaptação do PSO para um domínio combinatório.
- Usa o TSP como exemplo de problema em permutações.
- Serve como ponte entre o PSO contínuo de [[kennedy1995particle]] e aplicações discretas recentes.
- Não foi possível confirmar operadores, parâmetros ou resultados pelo PDF local.

## Relevância para o TCC

Esta referência é importante para justificar por que o PSO do TCC precisa de escolhas específicas de representação. No TSP/rTSP, uma partícula não pode simplesmente somar velocidade a uma posição contínua; ela precisa codificar tours, trocas, prioridades ou permutações. Essa afirmação é compatível com o título e escopo do arquivo, mas os detalhes técnicos dependem de recuperar um PDF legível.

## Métodos e Abordagens

- PSO discreto aplicado ao TSP.
- Representação baseada em tours/permutação, inferida pelo título.
- Não foi possível extrair equações, pseudocódigo, parâmetros ou resultados experimentais.

## Conexões

- [[kennedy1995particle]]
- [[particle-swarm]]
- [[particle-swarm]]
- [[tsp]]
- [[araujo2025pso]]
- [[sun2024hybrid]]
- [[huang2025matrix]]
- [[kappagantula2025dpso]]

## Notas e Insights

- A principal utilidade para o TCC é conceitual: mostrar que PSO discreto não é aplicação trivial do algoritmo original.
- Deve ser citado ao discutir a necessidade de adaptação de representação para permutações.
- Limitação de extração: PDF sem texto recuperável; OCR local resultou vazio.

## Citações-chave

Não extraídas. O PDF local não forneceu texto legível.
