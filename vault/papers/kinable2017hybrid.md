---
title: "Hybrid Optimization Methods for Time-Dependent Sequencing Problems"
authors: [Kinable, Joris, Cire, Andre A., van Hoeve, Willem-Jan]
year: 2017
doi: "10.1016/j.ejor.2016.11.035"
bibtex-key: kinable2017hybrid
pdf: "papers/pdfs/kinable2017hybrid.pdf"
tags: [tdtsp, lower-bound, decision-diagram, constraint-programming]
status: lido
rating: 5
---

## PDF

![[kinable2017hybrid.pdf]]

## Resumo

Kinable, Cire e van Hoeve propõem métodos híbridos exatos para problemas de sequenciamento em que o tempo de setup entre duas tarefas depende da posição relativa das tarefas na sequência. O estudo modela essa classe por variantes do time-dependent TSP: TD-TSP, TD-TSP com janelas de tempo e TD-SOP com precedências. A abordagem combina programação por restrições com relaxações discretas por diagramas de decisão multivalorados e relaxações contínuas de programação linear. Os MDDs capturam estrutura combinatória e fornecem limites durante a busca; a relaxação LP fornece custos reduzidos e informação dual, integrados por additive bounding.

## Contribuições Principais

- Propõe uma abordagem híbrida CP + MDD + LP para sequenciamento dependente de posição.
- Aplica o método a TD-TSP, TD-TSPTW e TD-SOP.
- Mostra como MDDs fortalecem propagação de domínios e poda na árvore de busca.
- Integra informação dual de relaxações LP em MDDs via additive bounding.
- Demonstra desempenho superior a modelos genéricos de MILP e CP em várias instâncias.

## Relevância para o TCC

O artigo é relevante para extensões do TCC em que o custo de visitar um ponto dependa da posição na rota, do tempo acumulado, de janelas de operação ou de precedências. Em patrulha por drones, isso pode representar degradação de bateria, prioridade temporal de áreas, restrições de visita ou mudança de custo conforme a missão avança. O trabalho também oferece referência metodológica para comparar metaheurísticas com abordagens exatas híbridas.

## Métodos e Abordagens

- Modelagem CP com variáveis de posição e restrição global `alldifferent`.
- Extensão para janelas de tempo por variáveis de conclusão e restrições de canalização.
- Extensão para precedências por variáveis de posição e relações de ordem.
- Construção de MDDs derivados de programação dinâmica para representar o espaço de sequências.
- Uso de relaxações LP baseadas em rede tempo-espaço para TD-TSP e variantes.

## Conexões

- [[garey1979computers]]
- [[heldkarp1970traveling]]
- [[fischetti1992additive]]
- [[bock2025survey]]
- [[winter2002modeling]]
- [[tsp-variants]]

## Notas e Insights

- O artigo diferencia TD-TSP dependente da posição de versões em que o tempo de viagem depende do instante de saída.
- CP puro com `alldifferent` é considerado fraco porque a restrição de unicidade fica desconectada da função objetivo.
- MDDs são usados não apenas para representar soluções, mas para gerar limites que ajudam a provar otimalidade.
- Em TD-TSPTW, MILP não encontrou solução viável para as instâncias geradas, enquanto as abordagens com MDD encontraram soluções para todas.

## Citações-chave

> “Our proposed methods rely on a hybrid approach where a constraint programming model is enhanced with two distinct relaxations: One discrete relaxation based on multivalued decision diagrams, and one continuous relaxation based on linear programming.”

> “Computational experiments clearly show that the hybrid approach outperforms traditional MILP and CP formulations for both sequencing problems in terms of time and solution quality.”
