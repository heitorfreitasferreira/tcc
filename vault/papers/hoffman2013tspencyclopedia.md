---
title: Traveling Salesman Problem
authors:
- Hoffman
- Karla L.
- Padberg
- Manfred
- Rinaldi
- Giovanni
year: 2013
doi: 10.1007/978-1-4419-1153-7_1068
bibtex_key: hoffman2013tspencyclopedia
bibtex-key: hoffman2013tspencyclopedia
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 5
role: revisao
areas:
- lower-bound
- tsp, "lower-bounds"]
methods:
- branch-and-bound
- exact
- linear-programming
- lower-bound
chapters:
- fundamentacao
claim_support: []
aliases: []
tags:
- area/lower-bound
- area/tsp
- capitulo/fundamentacao
- evidencia/referencia
- metodo/exact
- metodo/lower-bound
- papel/revisao
- relevancia/5
- status/pendente
- tipo/paper
---

## PDF

<!-- PDF não disponível -->

## Tese Central

O Problema do Caixeiro Viajante (TSP) é um dos problemas combinatórios mais estudados da pesquisa operacional, e sua dificuldade teórica (NP-difícil) contrasta com a existência de métodos exatos capazes de resolver instâncias com milhares de cidades. A combinação de limites inferiores fortes (Held-Karp, relaxações LP) com técnicas de enumeração implícita (branch-and-bound, branch-and-cut) forma a espinha dorsal dos resolvedores exatos modernos, enquanto heurísticas fornecem soluções de boa qualidade para aplicações práticas.

## Resumo

Enciclopédia de referência que cobre de forma abrangente o TSP: formulações matemáticas (formulação de Dantzig-Fulkerson-Johnson, formulação de Miller-Tucker-Zemlin), classes de instâncias (simétrico, assimétrico, euclidiano), limites inferiores clássicos (1-tree de Held-Karp, relaxação de atribuição, relaxações LP), métodos exatos (branch-and-bound, branch-and-cut, planos de corte), e heurísticas (construção, melhoria local, Lin-Kernighan). O verbete contextualiza historicamente o problema — desde os trabalhos seminais de Dantzig, Fulkerson e Johnson (1954) até os avanços computacionais que permitiram resolver instâncias do TSPLIB com milhares de cidades — e discute variantes como o TSP com janelas de tempo e o TSP probabilístico.

## Contribuições Principais

- Síntese autoritativa das formulações matemáticas do TSP (DFJ, MTZ, multifluxo)
- Exposição sistemática dos limites inferiores: 1-tree, Held-Karp, relaxação de atribuição, LP
- Descrição da arquitetura branch-and-cut que viabilizou a solução exata de instâncias grandes
- Cobertura de heurísticas clássicas (vizinho mais próximo, inserção, Christofides, 2-opt, 3-opt, Lin-Kernighan)
- Contextualização do TSP como benchmark para métodos de otimização e teoria da complexidade

## Relevância para o TCC

Esta é a referência canônica para o capítulo de referencial teórico (cap_referencial_teorico). Fornece a fundamentação matemática sobre limites inferiores que permite ao TCC contextualizar por que métodos exatos são inviáveis para rTSP com muitas requisições e justificar o uso de metaheurísticas (GA, PSO, ACO). A discussão sobre o limite de Held-Karp como referência de qualidade é diretamente usada para interpretar os gaps de otimalidade nos experimentos. A taxonomia de heurísticas clássicas também serve de baseline conceitual contra a qual as metaheurísticas bio-inspiradas são comparadas.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): caracterização do TSP como NP-difícil; definição de limites inferiores como referência para avaliação de soluções heurísticas
- Como citar na monografia: \cite{hoffman2013tspencyclopedia}

## Métodos e Abordagens

- Formulações matemáticas: Dantzig-Fulkerson-Johnson (DFJ), Miller-Tucker-Zemlin (MTZ), multifluxo
- Limites inferiores: 1-tree de Held-Karp, relaxação lagrangeana, relaxação de atribuição, relaxações de programação linear
- Métodos exatos: branch-and-bound, branch-and-cut com planos de corte (comb inequalities, subtour elimination)
- Heurísticas construtivas: vizinho mais próximo, inserção (mais próximo, mais distante, mais barata), Christofides (3/2-aproximação para TSP métrico)
- Heurísticas de melhoria: 2-opt, 3-opt, Or-opt, Lin-Kernighan (LK)
- Complexidade: NP-difícil, não aproximável por fator constante no caso geral (a menos que P=NP)

## Evidência / Resultado Relevante

- O limite de Held-Karp frequentemente atinge mais de 99% do valor ótimo para instâncias euclidianas típicas
- Branch-and-cut resolveu instâncias com mais de 10.000 cidades (TSPLIB) de forma exata
- A heurística de Christofides garante solução no máximo 50% pior que o ótimo para TSP métrico

## Limitações de Uso

> [!warning] Limitação
> Por ser um verbete enciclopédico, não contém resultados experimentais originais nem cobre metaheurísticas bio-inspiradas (GA, PSO, ACO) com profundidade. O foco é TSP clássico, não rTSP ou variantes com restrições temporais (makespan). Os métodos exatos descritos são inviáveis para as instâncias com 10-200 pontos e restrições de tempo do cenário de patrulha do TCC. Usar como referência conceitual e histórica, não como fonte direta de algoritmos implementados.

## Conexões

- Fundamenta: TSP, lower-bounds
- Relacionado a: [[heldkarp1970traveling]], [[heldkarp1971traveling]], [[lawler1985traveling]], [[applegate2006traveling]]
- Contrasta com: [[rajwar2023exhaustive]] — este é exato/clássico; rajwar cobre métodos bio-inspirados modernos
- Apoia claim: definição formal do TSP e seus limites inferiores teóricos
- Usado em capítulo: [[fundamentacao]]

## Notas e Insights

- O TSP ocupa posição única na pesquisa operacional: é suficientemente simples para enunciar em uma frase e suficientemente difícil para motivar décadas de pesquisa em otimização combinatória
- A evolução dos resolvedores exatos — de 49 cidades em 1954 para 85.900 cidades em 2006 (Applegate et al.) — é um dos feitos mais impressionantes da computação científica
- O limite de Held-Karp (1970/1971) permanece surpreendentemente justo (tight) para instâncias euclidianas, tipicamente dentro de 1% do ótimo
- A conexão entre TSP e spanning trees (1-tree) é um insight profundo que conecta teoria dos grafos, otimização combinatória e relaxação lagrangeana
- O verbete ilustra como um problema "de brinquedo" motivou avanços fundamentais em otimização inteira mista, teoria poliédrica e algoritmos de separação

## Citações-chave

> "The traveling salesman problem (TSP) is one of the most widely studied combinatorial optimization problems."

> "The TSP is easy to state, but it is very difficult to solve, and it belongs to the class of NP-hard problems."

> "The Held-Karp lower bound is one of the most celebrated results in the TSP literature, providing a bound that is typically within 1% of the optimal value."
