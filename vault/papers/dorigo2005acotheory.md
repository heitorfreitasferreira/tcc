---
title: 'Ant Colony Optimization Theory: A Survey'
authors:
- Dorigo
- Marco
- Blum
- Christian
year: 2005
doi: 10.1016/j.tcs.2005.05.020
bibtex_key: dorigo2005acotheory
bibtex-key: dorigo2005acotheory
tags:
- evidencia/referencia
- metodo/aco
- papel/revisao
- papel/teorico
- status/lido
- tipo/paper
status: lido
rating: 4
pdf: papers/pdfs/dorigo2005acotheory.pdf
type: paper
methods:
- aco
role: revisao
reading_status: lido
validation_status: nao-validado
---

## PDF

![[dorigo2005acotheory.pdf]]

## Resumo

Dorigo e Blum fazem um survey teórico de [[ant-colony]], deslocando o foco de aplicações de prova de conceito para perguntas formais sobre como e por que ACO funciona. O artigo revisa resultados de convergência, aproxima ACO de métodos de busca baseada em modelos, discute relações com stochastic gradient ascent e cross-entropy, e analisa vieses de busca. O TSP aparece como aplicação inicial e como exemplo natural para definir modelos de problema, componentes de solução e trilhas de feromônio.

## Contribuições Principais

- Organiza resultados teóricos de convergência para ACO.
- Distingue convergência em valor e convergência em solução.
- Formaliza ACO como procedimento estocástico baseado em modelo probabilístico parametrizado por feromônios.
- Relaciona ACO a model-based search, stochastic gradient ascent e cross-entropy.
- Discute vieses negativos de busca, incluindo competição injusta entre componentes e selection fix-points.

## Relevância para o TCC

É uma referência central para fundamentar [[ant-colony]] além da metáfora das formigas. O TCC implementa ACO como metaheurística para TSP/rTSP; este artigo fornece linguagem formal para descrever construção probabilística de soluções, atualização de feromônio, informação heurística, evaporação e exploração/explotação. Também ajuda a evitar afirmações fortes demais: convergência assintótica não implica desempenho prático em tempo finito.

## Métodos e Abordagens

- Define problema de otimização combinatória como espaço finito de soluções, restrições e função objetivo.
- Define componentes de solução e parâmetros de feromônio associados a esses componentes.
- Apresenta o framework básico de ACO: inicialização, construção de soluções, busca local opcional, atualização de feromônio e retorno da melhor solução.
- Discute transição probabilística combinando feromônio e informação heurística.
- Analisa ACO com limites inferiores de feromônio para garantir probabilidade positiva de construir qualquer solução.

## Conexões

- [[ant-colony]]
- [[ant-colony]]
- [[dorigo1996ant]]
- [[dorigo1997ant]]
- [[dorigo2004book]]
- [[blum2005acointro]]
- [[stutzle2000mmas]]
- [[tsp]]

## Notas e Insights

- A distinção entre garantia assintótica e desempenho experimental é essencial para o TCC.
- A formalização por componentes de solução combina bem com TSP, pois arestas ou decisões de próximo nó podem receber feromônio.
- A discussão sobre viés ajuda a interpretar resultados ruins de ACO como efeito de representação e construção, não apenas de parâmetros.

## Citações-chave

> “candidate solutions are constructed using a pheromone model, that is, a parametrized probability distribution over the solution space”

> “The proofs that were presented in this section do not say anything about the time required to find an optimal solution, which can be astronomically large.”
