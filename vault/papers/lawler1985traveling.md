---
title: "The Traveling Salesman Problem: A Guided Tour of Combinatorial Optimization"
authors: [Lawler, Eugene L., Lenstra, Jan Karel, Rinnooy Kan, Alexander H. G., Shmoys, David B.]
year: 1985
doi: "10.1002/net.3230170210"
bibtex-key: lawler1985traveling
# pdf: "papers/pdfs/lawler1985traveling.pdf"  # unavailable — all known sources behind Cloudflare/login
tags: [tsp metaheuristic]
status: lido
rating: 5
---

## Resumo

Coletânea editada por Lawler, Lenstra, Rinnooy Kan e Shmoys que reúne capítulos escritos por especialistas renomados cobrindo todos os aspectos do Problema do Caixeiro Viajante (TSP). Cada capítulo aborda uma dimensão diferente do problema — desde a NP-completude e formulações de programação linear até métodos exatos, heurísticas, aspectos poliédricos e variações. A obra consolidou o TSP como o "problema-modelo" da otimização combinatória e serviu como referência padrão por duas décadas.

## Contribuições Principais

- Primeira obra abrangente a tratar o TSP como problema central da otimização combinatória
- Capítulos sobre o poliedro do TSP, planos de corte, branch-and-bound e heurísticas de busca local
- Discussão aprofundada de variações (TSP assimétrico, TSP com múltiplos caixeiros, TSP estocástico)
- Bibliografia extensa e exercícios que orientaram gerações de pesquisadores
- Estabeleceu o TSP como *benchmark* padrão para novos algoritmos de otimização

## Relevância para o TCC

Como obra de referência, fornece a base teórica completa para entender o TSP e suas variantes. O capítulo sobre heurísticas contextualiza as meta-heurísticas bio-inspiradas utilizadas no TCC. As variações discutidas (mTSP, TSP estocástico) têm conexão direta com o cenário de patrulha com drones (rTSP).

## Métodos e Abordagens

- Programação linear inteira e planos de corte
- Branch-and-bound exato
- Heurísticas de construção (vizinhança mais próxima, inserção)
- Busca local (2-opt, 3-opt, Lin-Kernighan)
- Programação dinâmica
- Aspectos poliédricos e desigualdades válidas

## Conexões

- [[applegate2006traveling]] — estudo computacional que atualizou e expandiu este survey
- [[garey1979computers]] — NP-completude do TSP
- [[lin1973effective]] — heurística LK para TSP
- [[oliver1987crossover]] — operadores de cruzamento para GA em TSP (constrói sobre este survey)
- [[goldberg1989genetic]] — GAs como alternativa às heurísticas clássicas aqui descritas
- [[rajan2022routing]] — extensão estocástica do TSP para patrulha com UAV
- [[TSP]]

## Notas e Insights

- O TSP é descrito como "problema-modelo" — cada nova técnica de otimização é testada primeiro no TSP
- A estrutura do livro (capítulos independentes por especialistas) permite consulta direcionada
- Apesar de datado em aspectos computacionais, o conteúdo teórico permanece atual
- A seção sobre variações do TSP inspirou diretamente problemas como FSTSP e rTSP

## Citações-chave

> "The Traveling Salesman Problem is central to the area of Combinatorial Optimization, and it is through this problem that many of the most important developments in the area have been made."

> "Each chapter deals with a different aspect of the problem, and has been written by an acknowledged expert in the field. Focuses on the essential ideas in a self-contained manner."
