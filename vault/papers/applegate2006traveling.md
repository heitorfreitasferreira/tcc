---
title: "The Traveling Salesman Problem: A Computational Study"
authors: [Applegate, David L., Bixby, Robert E., Chvátal, Vašek, Cook, William J.]
year: 2006
doi: "10.1515/9781400841103"
bibtex-key: applegate2006traveling
pdf: "papers/pdfs/applegate2006traveling.pdf"
tags: [tsp]
status: lido-parcial
rating: 5
---

## PDF

![[applegate2006traveling.pdf]]

## Resumo

Livro-monografia que documenta duas décadas de pesquisa dos autores sobre o Problema do Caixeiro Viajante (TSP), culminando no desenvolvimento do solver *Concorde*. Apresenta a abordagem *branch-and-cut* combinada com heurísticas sofisticadas (Lin-Kernighan encadeada, LKH) para resolver instâncias de até 85.900 cidades. Vencedor do Prêmio Lanchester de 2007, a obra cobre desde a história do problema até detalhes de implementação do código, passando por planos de corte, gerenciamento de PL, estratégias de branching e heurísticas de busca de tours.

## Contribuições Principais

- Desenvolvimento do solver Concorde, estado-da-arte para solução exata de TSP
- Avanços significativos em planos de corte (subtour, blossom, comb, domino-parity)
- Heurística *Chained Lin-Kernighan* para obtenção de limites superiores de alta qualidade
- Estratégias de *branching* (strong branching, tentative branching) para reduzir a árvore de busca
- Demonstração de que TSPs com dezenas de milhares de cidades podem ser resolvidos exatamente

## Relevância para o TCC

O Concorde serve como referência ótima para validação de meta-heurísticas em instâncias TSP. No contexto do TCC, a execução de algoritmos bio-inspirados (GA, PSO, ACO) pode ser comparada contra soluções ótimas obtidas pelo Concorde para instâncias pequenas. As heurísticas LK/LKH descritas no livro também inspiram operadores de busca local nos métodos implementados.

## Métodos e Abordagens

- Branch-and-cut (programação linear inteira com planos de corte)
- Planos de corte: subtour, blossom, comb, domino-parity, hipergrafos
- Heurísticas: Lin-Kernighan, Chained Lin-Kernighan, LKH (Helsgaun)
- Gerenciamento de PL: core LP, cut storage, edge pricing
- Strong branching e tentative branching
- Safe shrinking para redução de instância

## Conexões

- [[lawler1985traveling]] — survey clássico que este livro atualiza e expande
- [[garey1979computers]] — NP-completude do TSP
- [[lin1973effective]] — heurística Lin-Kernighan, base das buscas locais do livro
- [[dellamico2022exact]] — modelos exatos para FSTSP (inspirados em branch-and-cut)
- [[goldberg1989genetic]] — GAs como alternativa heurística contrastada com métodos exatos
- [[tsp]]

## Notas e Insights

- A combinação de planos de corte + branch-and-bound (branch-and-cut) é a abordagem mais bem-sucedida para solução exata de TSP
- O Concorde resolveu a instância mundial com 85.900 cidades, um marco histórico
- Heurísticas de alta qualidade (LK encadeada) são essenciais para podar a árvore de branch-and-cut
- Livro essencial para entender o que constitui uma solução ótima de TSP — referência fundamental para qualquer trabalho que proponha novas meta-heurísticas

## Citações-chave

> "Our primary concern in this book is to describe a method and computer code that have succeeded in solving a wide range of large-scale instances of the TSP."

> "The TSP is also of theoretical interest in computer science because it is one of the important class of NP-Hard combinatorial optimization problems."
