---
title: "Genetic Algorithms in Search, Optimization, and Machine Learning"
authors: [Goldberg, David E.]
year: 1989
doi: ""
bibtex_key: goldberg1989genetic
bibtex-key: goldberg1989genetic
pdf: "papers/pdfs/goldberg1989genetic.pdf"
tags: [ga metaheuristic]
status: lido-parcial
rating: 5
---

## PDF

![[goldberg1989genetic.pdf]]

## Resumo

Goldberg apresenta GAs como procedimentos de busca baseados nos mecanismos da seleção natural e da genética natural. O livro organiza a metodologia para busca, otimização e aprendizado de máquina, explicando o algoritmo genético simples, o teorema dos esquemas, operadores de reprodução, crossover e mutação, além de aplicações. O PDF local é escaneado; OCR em páginas selecionadas funcionou, embora com erros típicos de reconhecimento.

> [!warning] Leitura parcial
> O OCR é utilizável em trechos centrais, mas há caracteres trocados e fórmulas degradadas. Citações literais devem ser conferidas visualmente.

## Contribuições Principais

- Consolida GAs como método de busca e otimização.
- Apresenta o Simple Genetic Algorithm.
- Explica schemata, ordem, comprimento definidor e teorema fundamental dos GAs.
- Discute o efeito combinado de reprodução, crossover e mutação.
- Liga GAs, otimização e sistemas classificadores.

## Relevância para o TCC

Goldberg é a referência prática para descrever o GA implementado no TCC: população, aptidão, seleção probabilística, crossover, mutação e iterações. Para o problema de patrulha com drones, ele ajuda a explicar por que soluções parciais de boa qualidade podem se propagar: esquemas curtos, de baixa ordem e acima da média tendem a receber mais amostras em gerações posteriores.

## Métodos e Abordagens

- Formulação de GAs como busca em populações de strings.
- Seleção proporcional à aptidão, incluindo exemplo de roulette wheel.
- Definição de schema com símbolo wildcard.
- Análise de ordem do schema e comprimento definidor.
- Crossover como troca estruturada e aleatória de informação entre strings.
- Mutação como alteração aleatória de alelos.

## Conexões

- [[holland1975adaptation]]
- [[genetic-algorithms]]
- [[tsp]]
- [[oliver1987crossover]]
- [[bean1994genetic]]
- [[potvin1996ga]]
- [[larranaga1999ga]]

## Notas e Insights

- Goldberg sacrifica parte do rigor formal para construir intuição; isso torna a obra útil para fundamentação metodológica do TCC.
- A análise por esquemas fornece justificativa mais forte para crossover do que a simples analogia com reprodução biológica.
- O livro deve ser usado para explicar o GA canônico, não operadores específicos de TSP.

## Citações-chave

> “This book is about genetic algorithms (GAs)—search procedures based on the mechanics of natural selection and natural genetics.”

> “Crossover is a structured yet randomized information exchange between strings.”
