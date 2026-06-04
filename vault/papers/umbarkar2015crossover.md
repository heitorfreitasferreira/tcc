---
title: "Crossover Operators in Genetic Algorithms: A Review"
authors: ["Umbarkar, A. J.", "Sheth, P. D."]
year: 2015
doi: "10.21917/ijsc.2015.0150"
bibtex_key: umbarkar2015crossover
bibtex-key: umbarkar2015crossover
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 3
role: "revisao"
areas:
  - genetic-algorithms
methods:
  - ga
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - metodo/ga
  - area/genetic-algorithms
  - area/bio-inspired-optimization
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/3
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/umbarkar2015crossover.pdf]] -->

## Tese Central

O operador de crossover é o mecanismo central de busca dos algoritmos genéticos, e sua eficácia depende criticamente da adequação entre o tipo de operador e a representação do problema; uma taxonomia que classifica operadores por representação (binária, inteira/permutação, real) e por mecanismo (discreto, aritmético, baseado em ordem) fornece diretrizes para seleção informada de operadores em aplicações práticas de GA.

## Resumo

Artigo de revisão publicado no ICTACT Journal on Soft Computing (2015) que categoriza e analisa os principais operadores de crossover utilizados em algoritmos genéticos. Organiza os operadores por tipo de representação: crossover para codificação binária (one-point, two-point, k-point, uniform, half-uniform, shuffle), para codificação de permutação/inteira (PMX — Partially Mapped Crossover, OX — Order Crossover, CX — Cycle Crossover, POS — Position-based Crossover, OBX — Order-based Crossover) e para codificação real (arithmetic, heuristic, BLX-α, SBX — Simulated Binary Crossover, Laplace crossover). Discute vantagens, desvantagens e adequação de cada operador a diferentes classes de problemas. Embora menos abrangente que surveys mais recentes, é uma referência didática útil como ponto de partida para compreensão de operadores de crossover.

## Contribuições Principais

- Taxonomia de operadores de crossover organizada por tipo de representação (binária, permutação, real)
- Descrição detalhada dos operadores de crossover para permutação (PMX, OX, CX, POS, OBX) — diretamente relevantes para TSP
- Comparação qualitativa de vantagens/desvantagens de cada operador (preservação de ordem, preservação de adjacência, viabilidade dos filhos)
- Discussão do papel do crossover como operador de explotação, complementar à mutação como operador de exploração

## Relevância para o TCC

- Fornece base teórica para a seção de operadores de crossover no cap_referencial_teorico
- A taxonomia de operadores de permutação (PMX, OX, CX) cobre exatamente os operadores candidatos para o GA-TSP do TCC
- A discussão sobre preservação de ordem vs. adjacência ajuda a justificar a escolha do operador de crossover na implementação
- Complementa [[alhijawi2024genetic]] com detalhamento mais focado exclusivamente em crossover

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Descrição e justificativa dos operadores de crossover para permutação utilizados na implementação do GA para TSP
- Como citar na monografia: `\cite{umbarkar2015crossover}` como referência complementar na seção de operadores genéticos

## Métodos e Abordagens

- Revisão narrativa com categorização taxonômica
- Análise por tipo de representação: binária, permutação/inteira, valor real
- Para cada operador: descrição do mecanismo, pseudocódigo ou explicação passo a passo, vantagens e limitações
- Discussão qualitativa (sem experimentação empírica própria)

## Evidência / Resultado Relevante

- PMX e OX são os operadores mais utilizados para TSP por preservarem, respectivamente, mapeamento de posição e ordem relativa
- CX preserva posições absolutas dos pais, o que pode ser vantajoso ou limitante dependendo da estrutura da paisagem de fitness do TSP
- Para problemas com representação real, SBX (Simulated Binary Crossover) é amplamente adotado por simular o comportamento de crossover binário em domínio contínuo

## Limitações de Uso

> [!warning] Limitação
> Survey de 2015, anterior a desenvolvimentos recentes em GAs (aprendizado de máquina para seleção de operadores, GAs adaptativos). A revisão é qualitativa e não inclui benchmarks empíricos próprios — não fornece evidência quantitativa de superioridade de operadores. Para o TCC, deve ser complementado com [[potvin1996ga]] e [[larranaga1999ga]] que fornecem comparações empíricas de operadores de crossover em TSP. Rating 3 reflete escopo mais limitado e idade da publicação.

## Conexões

- Fundamenta: [[goldberg1989genetic]], [[holland1975adaptation]] — obras fundamentais que introduziram os operadores de crossover clássicos
- Relacionado a: [[alhijawi2024genetic]] — survey mais recente e abrangente que inclui crossover; [[hassanat2019crossover]] — foco complementar em taxas (não tipos) de crossover
- Contrasta com: [[larranaga1999ga]] — abordagem empírica com benchmarks de desempenho de operadores em TSP
- Apoia claim: Caracterização dos operadores de crossover para permutação (PMX, OX, CX) como base para seleção no GA do TCC
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A distinção entre operadores que preservam posição absoluta (PMX, CX) vs. ordem relativa (OX, OBX) vs. adjacência (Edge Recombination) é crucial para entender por que OX tende a performar melhor em TSP: a qualidade de um tour depende mais da ordem relativa das cidades do que de suas posições absolutas
- O artigo é particularmente útil como referência didática: as explicações passo a passo de cada operador facilitam a implementação
- A ausência de experimentação empírica própria limita seu uso como evidência, mas sua taxonomia é útil como estrutura organizadora no referencial teórico
- Complementar com [[potvin1996ga]] para validação empírica de que OX e edge recombination superam PMX e CX em TSP

## Citações-chave

> "Crossover operator is the backbone of the genetic algorithm, playing a vital role in the convergence of the algorithm by combining the genetic information of two parents to generate new offspring."

> "For permutation-based problems like TSP, specialized crossover operators such as PMX, OX, and CX are required to maintain the validity of solutions, as standard binary crossover operators would produce infeasible tours with duplicate or missing cities."

> "The choice of crossover operator should be based on the problem representation and the properties of the solution space that need to be preserved."
