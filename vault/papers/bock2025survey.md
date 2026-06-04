---
title: "A Survey on the Traveling Salesman Problem and its Variants in a Warehousing Context"
authors: [Bock, Stefan]
year: 2025
doi: "10.1016/j.ejor.2024.04.014"
bibtex_key: bock2025survey
bibtex-key: bock2025survey
pdf: "papers/pdfs/bock2025survey.pdf"
tags: [tsp tsp-variants warehousing complexity]
status: lido
rating: 3
---

## PDF

![[bock2025survey.pdf]]

## Resumo

Bock et al. fazem um survey específico sobre o TSP e variantes em armazéns com estrutura de corredores paralelos. O artigo parte do problema clássico de roteamento de separadores, modelável como TSP, mas mostra que processos modernos de e-commerce, robôs móveis, picking assistido por AMR, armazenamento disperso, múltiplos depósitos, prazos e inventário automatizado exigem variantes como TSP com precedência, clustered TSP, generalized TSP, prize-collecting TSP, orienteering, traveling repairman, TSPTW e covering salesman. A principal contribuição é cruzar três dimensões: uso operacional em armazém, literatura existente e status de complexidade quando o grafo tem estrutura de blocos.

## Contribuições Principais

- Identifica dez variantes de TSP relevantes para problemas de roteamento em armazéns.
- Diferencia a complexidade do TSP em grafos gerais da complexidade em layouts de corredores paralelos.
- Mostra que o TSP clássico em armazéns de um bloco, dois blocos e múltiplos blocos com número limitado de corredores transversais é polinomial.
- Apresenta resultados de complexidade para variantes como prize-collecting TSP, orienteering e TSP com janelas de tempo em armazéns.
- Organiza casos de uso reais: picking básico, AMR-assisted picking, scattered storage, stowing, perecíveis, coordenação de equipe e stock-taking robótico.

## Relevância para o TCC

O artigo é útil para enquadrar o TCC como problema de roteamento em ambiente físico estruturado. Embora o cenário do TCC seja patrulha de drones, a lógica de visitar posições de interesse com custo de deslocamento mínimo é análoga ao roteamento de pickers e robôs em armazém. O survey também ajuda a defender a escolha do TSP/rTSP como base experimental: o TSP clássico pode ter casos estruturados tratáveis, mas variantes com restrições operacionais recuperam dificuldade computacional.

## Métodos e Abordagens

- Survey de variantes de TSP filtradas por aplicabilidade a armazéns.
- Formulações matemáticas para ATSP e variantes, com variáveis de arco, posição, seleção e tempo.
- Análise de complexidade específica para layouts `1B`, `2B` e `MB`.
- Discussão de programas dinâmicos para TSP em corredores paralelos.
- Classificação de variantes por caso de uso e status: polinomial, NP-difícil binário, NP-difícil forte ou aberto.

## Conexões

- [[garey1979computers]]
- [[lawler1985traveling]]
- [[applegate2006traveling]]
- [[winter2002modeling]]
- [[murray2015flying]]
- [[agatz2018optimization]]
- [[tsp-variants]]

## Notas e Insights

- O artigo corrige uma generalização comum: TSP geral é NP-difícil, mas o TSP clássico em certos layouts de armazém é polinomial.
- A tabela final resume variantes, usos e complexidade de forma útil para a monografia.
- O texto defende que complexidade ainda importa na prática, pois algoritmos eficientes viram sub-rotinas em decomposições maiores.
- A distinção entre problemas estáticos/determinísticos e extensões dinâmicas/estocásticas é relevante para posicionar o rTSP do TCC.

## Citações-chave

> “Traditional picker routing, in which a single picker has to visit a given set of picking positions in a picker-to-parts process, can be modeled as the classical Traveling Salesman Problem (TSP).”

> “It is amazing to see that such an old-established field like routing in warehouses still offers so many unexplored use cases and unresolved methodological research challenges.”
