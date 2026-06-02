---
title: "A Survey on the Traveling Salesman Problem and its Variants in a Warehousing Context"
authors: [Bock, Stefan]
year: 2025
doi: "10.1016/j.ejor.2024.04.014"
bibtex-key: bock2025survey
pdf: "papers/pdfs/bock2025survey.pdf"
tags: [tsp tsp-variants warehousing complexity]
status: lido
rating: 3
---

## Resumo

Survey publicado no *European Journal of Operational Research* que sistematiza variantes do TSP no contexto de armazéns e centros de distribuição, motivado por e-commerce e fast-delivery. Aborda o Clustered TSP, Generalized TSP e Prize-Collecting TSP, analisando a complexidade computacional sob a estrutura de corredores paralelos (*parallel-aisle warehouse*). O artigo apresenta novos resultados de complexidade, identificando casos polinomiais tratáveis e variantes que permanecem NP-difíceis, além de propor uma agenda de pesquisa para a área.

## Contribuições Principais

- Primeira sistematização unificada de variantes TSP em contexto logístico de armazéns
- Novos resultados de complexidade para a topologia *parallel-aisle*
- Distinção clara entre casos polinomiais e NP-difíceis
- Agenda de pesquisa com direções para extensões dinâmicas e estocásticas

## Relevância para o TCC

Embora o foco seja armazéns, a metodologia de classificação de complexidade para variantes do TSP (clustered, generalized, prize-collecting) é diretamente transferível para o contexto de roteirização de drones. As estruturas de clusterização do GTSP e CTSP têm paralelos diretos com zonas de patrulha no rTSP.

## Métodos e Abordagens

- Survey com análise formal de complexidade computacional
- Classificação de variantes TSP por aplicabilidade em *warehousing*
- Novas provas de complexidade para casos específicos com topologia *parallel-aisle*
- Análise de algoritmos exatos, heurísticos e de aproximação

## Conexões

- [[TSP]]
- [[tsp-variants]]
- [[pop2024comprehensive]] — survey GTSP (contemporâneo)
- [[lawler1985traveling]] — survey clássico do TSP
- [[garey1979computers]] — NP-completude
- [[chandra2022comparative]] — estudo comparativo de metaheurísticas TSP

## Notas e Insights

- PDF obtido via RWTH Aachen institutional repository (open access): https://publications.rwth-aachen.de/record/985849
- O TSP clássico é tratável em corredores paralelos, mas variantes como GTSP e CTSP permanecem NP-difíceis
- A metodologia de classificação pode inspirar abordagem similar para variantes de drone routing
- A ausência de aplicações diretas em drones limita a transferibilidade imediata, mas os conceitos de clusterização são úteis
- Survey de 2025 muito recente; não captura desenvolvimentos posteriores

## Citações-chave

> "The classical TSP is polynomially solvable in parallel-aisle warehouses, but its variants remain NP-hard in the same setting."

> "Understanding the complexity of TSP variants in structured environments is essential for designing efficient logistics systems."
