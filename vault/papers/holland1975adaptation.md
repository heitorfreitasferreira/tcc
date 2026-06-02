---
title: "Adaptation in Natural and Artificial Systems"
authors: [Holland, John H.]
year: 1975
doi: ""
bibtex-key: holland1975adaptation
pdf: "papers/pdfs/holland1975adaptation.pdf"
tags: [ga foundational]
status: lido
rating: 4
---

## PDF

![[holland1975adaptation.pdf]]

## Resumo

Livro fundacional dos Algoritmos Genéticos, publicado originalmente em 1975. Holland propõe um framework matemático formal para sistemas adaptativos inspirado na evolução natural: seleção, crossover, mutação. Apresenta o Teorema dos Esquemas (Schema Theorem), que explica porque os GAs funcionam — os esquemas de baixa ordem e curta distância definidora (building blocks) crescem exponencialmente na população. O livro estabelece as bases para toda a área de computação evolucionária, influenciando Goldberg, De Jong, Koza e gerações seguintes.

## Contribuições Principais

- Formalização matemática dos algoritmos genéticos como sistemas adaptativos
- Teorema dos Esquemas (Schema Theorem) — base teórica que explica o poder dos GAs
- Hipótese dos Building Blocks — construção de soluções complexas pela recombinação de sub-soluções simples

## Relevância para o TCC

Referência histórica e teórica fundamental. Citado para justificar a escolha de Algoritmos Genéticos como método de otimização, apoiando-se na base teórica do Teorema dos Esquemas e na Hipótese dos Building Blocks.

## Métodos e Abordagens

- Algoritmo Genético Canônico: seleção proporcional (roleta), crossover de um ponto, mutação bit-flip
- Representação binária (cromossomos como strings binárias)
- Teorema dos Esquemas como ferramenta de análise

## Conexões

- [[goldberg1989genetic]] — textbook que consolidou e popularizou os GAs
- [[oliver1987crossover]] — primeira aplicação de GA ao TSP
- [[bean1994genetic]] — random keys para GA em problemas de sequenciamento
- [[larranaga1999ga]] — survey de GA para TSP, herdeiro direto
- [[genetic-algorithms]]

## Notas e Insights

Holland não aplicou GAs ao TSP diretamente. A ligação com TSP veio com Oliver (1987) e foi consolidada por Larrañaga (1999). O livro é denso e matemático; para implementação prática, Goldberg (1989) é mais acessível.

## Citações-chave

> The schema theorem shows that the GA allocates exponentially increasing trials to above-average schemas.
