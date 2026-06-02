---
title: "A Study of Permutation Crossover Operators on the Traveling Salesman Problem"
authors: [Oliver, Ian M., Smith, D. J., Holland, John R. C.]
year: 1987
doi: ""
bibtex-key: oliver1987crossover
tags: [ga tsp]
status: lido
rating: 4
---

## Resumo

Artigo seminal que analisa três operadores de cruzamento por permutação — Parcialmente Mapeado (PMX), Order Crossover (OX) e Cycle Crossover (CX) — para aplicação de Algoritmos Genéticos ao Problema do Caixeiro Viajante (TSP). Os autores caracterizam como cada operador amostra o espaço de o-schemas (esquemas de ordem) e, portanto, a que tipos de problema cada um é mais adequado. Experimentos em instâncias TSP corroboram a análise teórica, estabelecendo a base para o uso de GAs em problemas de permutação.

## Contribuições Principais

- Definição e análise formal dos operadores PMX, OX e CX para representação por permutação
- Introdução do conceito de o-schema para análise teórica de operadores de ordem
- Demonstração experimental de que OX é superior para TSP entre os três operadores
- Estabelecimento da agenda de pesquisa para operadores de cruzamento em GAs para problemas combinatórios
- Mais de 1.000 citações, tornando-se referência obrigatória em GA para TSP

## Relevância para o TCC

O operador OX (Order Crossover) é utilizado na implementação de GA do repositório (`src/pkg/ga/`). A análise comparativa dos operadores de cruzamento justifica a escolha de OX como operador padrão. Os conceitos de preservação de ordem e posição são relevantes para o rTSP, onde a sequência de visitação dos pontos de interesse é a variável de decisão.

## Métodos e Abordagens

- Representação por permutação (path representation) para tours TSP
- Operador PMX (Partially Mapped Crossover) — preserva posições absolutas
- Operador OX (Order Crossover) — preserva ordem relativa
- Operador CX (Cycle Crossover) — preserva posição absoluta de subconjuntos
- Análise de o-schemas (esquemas de ordem)
- Experimentos em instâncias TSP

## Conexões

- [[holland1975adaptation]] — GA, fundação teórica (esquemas e crossovers)
- [[goldberg1989genetic]] — textbook que sistematizou GAs e expandiu análise de crossover
- [[bean1994genetic]] — random keys (outra representação para TSP em GA)
- [[lin1973effective]] — TSP, benchmark para operadores de cruzamento
- [[genetic-algorithms]]

## Notas e Insights

- OX é o operador mais utilizado em GAs para TSP até hoje
- PMX tende a preservar posições absolutas, OX preserva ordem relativa, CX preserva posições de ciclo
- A escolha do operador de crossover impacta significativamente a convergência e qualidade da solução
- O artigo não aborda mutação — foco exclusivo em crossover

## Citações-chave

> "Three permutation crossovers are analyzed to characterize how they sample the o-schema space, and hence what type of problems they may be applicable to."

> "Experiments performed on the Traveling Salesman Problem go some way to support the theoretical analysis."
