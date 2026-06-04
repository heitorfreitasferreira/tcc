---
title: "Estimating the Held-Karp Lower Bound for the Geometric TSP"
authors: [Valenzuela, Christine L., Jones, Antonia J.]
year: 1997
doi: "10.1016/S0377-2217(96)00214-7"
bibtex_key: valenzuela1997estimating
bibtex-key: valenzuela1997estimating
pdf: "papers/pdfs/valenzuela1997estimating.pdf"
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: disponivel
rating: 4
role: "revisao"
areas:
  - tsp, "lower-bounds"]
methods: ["held-karp", "linear-programming", "estimation"]
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/lower-bound
  - metodo/exact
  - metodo/lower-bound
  - metodo/held-karp
  - metodo/lagrangean
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/4
---

## PDF

![[valenzuela1997estimating.pdf]]

## Tese Central

O limite inferior de Held-Karp (HK) é extraordinariamente justo para o TSP geométrico — frequentemente atinge mais de 99% do valor ótimo — mas seu cálculo exato via relaxação lagrangeana é caro. Valenzuela e Jones propõem métodos de estimação computacionalmente eficientes que aproximam o valor HK com erro controlado, permitindo seu uso prático como referência de qualidade em metaheurísticas e algoritmos aproximativos para TSP geométrico.

## Resumo

O artigo examina os fundamentos teóricos do limite inferior de Held-Karp para o TSP geométrico e investiga a relação entre o gap HK e características estruturais das instâncias (distribuição espacial, número de cidades, propriedades de agrupamento). A contribuição central é um conjunto de métodos de estimação que aproximam o limite HK sem resolver o problema de otimização lagrangeana completo, viabilizando seu uso como critério de parada (lower bound stopping rule) em algoritmos heurísticos. Os autores analisam o comportamento do limite em instâncias com diferentes distribuições espaciais, mostrando que o gap HK tende a ser maior em instâncias com clusters pronunciados.

## Contribuições Principais

- Análise teórica detalhada do limite inferior de Held-Karp para o TSP geométrico
- Métodos de estimação do limite HK com custo computacional reduzido
- Caracterização empírica do gap HK em função de propriedades das instâncias (n, distribuição, clusterização)
- Proposta de uso do limite HK estimado como critério de parada para algoritmos heurísticos
- Conexão entre estrutura espacial das instâncias e qualidade dos limites inferiores

## Relevância para o TCC

O limite de Held-Karp é a referência teórica de qualidade contra a qual os gaps de otimalidade dos métodos bio-inspirados (GA, PSO, ACO) são avaliados nos experimentos do TCC. Este artigo fornece a justificativa empírica para usar HK como baseline: ele é suficientemente justo para instâncias geométricas (erro típico < 1%) e sua estimação é computacionalmente viável. No contexto de roteamento de drones sobre pontos geográficos (TSP euclidiano/geométrico), os resultados do artigo são diretamente aplicáveis. O uso do limite HK como critério de parada é uma técnica que pode ser adaptada para os experimentos do TCC.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): justificativa do uso do limite de Held-Karp como referência de qualidade; viabilidade de estimar HK para instâncias geométricas do cenário de patrulha
- Como citar na monografia: \cite{valenzuela1997estimating}

## Métodos e Abordagens

- Relaxação lagrangeana do TSP via 1-tree (Held-Karp)
- Método do subgradiente para otimização do dual lagrangeano
- Estimação do limite HK por amostragem e técnicas de extrapolação
- Análise do gap HK como função de: número de cidades (n), distribuição espacial (uniforme, clusterizada), propriedades geométricas
- Geração de instâncias geométricas com características controladas para experimentação
- Validação contra valores ótimos conhecidos do TSPLIB

## Evidência / Resultado Relevante

- O limite HK atinge tipicamente 99%+ do valor ótimo para TSP geométrico com distribuição uniforme
- O gap HK é maior em instâncias clusterizadas (distribuição não uniforme dos pontos)
- A taxa de convergência do método do subgradiente depende da estrutura da instância
- Métodos de estimação alcançam erro inferior a 1% em relação ao valor HK exato com custo computacional significativamente menor

## Limitações de Uso

> [!warning] Limitação
> O foco em TSP geométrico (pontos em R² com distância euclidiana) é diretamente aplicável ao cenário do TCC, mas os resultados sobre clusterização espacial podem não se transferir para instâncias com restrições temporais (makespan) do rTSP. Os métodos de estimação são de 1997 e existem abordagens mais modernas (e.g., Concorde TSP solver). A implementação da relaxação lagrangeana requer tuning do passo do subgradiente, que é sensível à escala da instância. Para instâncias muito pequenas (n < 20), o gap HK pode ser menos informativo como critério de parada.

## Conexões

- Fundamenta: [[heldkarp1970traveling]], [[heldkarp1971traveling]], TSP, lower-bounds
- Relacionado a: [[johnson1996asymptotic]], [[fischetti1992additive]], [[saller2025approximability]]
- Contrasta com: [[gutekunst2020relaxations]] — este foca em estimação do HK; Gutekunst cobre relaxações mais amplas incluindo SDP
- Apoia claim: viabilidade de usar HK como limite inferior de referência para TSP geométrico
- Usado em capítulo: [[fundamentacao]]

## Notas e Insights

- O artigo revela uma propriedade notável: o limite HK é consistentemente justo para TSP geométrico, mas a justificativa teórica para esse fenômeno permanece parcial — é mais uma observação empírica robusta do que um teorema
- A descoberta de que instâncias clusterizadas produzem gaps HK maiores é relevante para o TCC: pontos de patrulha (POIs) podem exibir agrupamento natural em áreas urbanas
- O método do subgradiente para otimização lagrangeana é conceitualmente simples (atualiza multiplicadores proporcionalmente à violação das restrições), mas seu desempenho prático depende criticamente da escolha da sequência de passos
- A noção de "estimar" em vez de "calcular exatamente" o limite HK reflete uma filosofia pragmática: em otimização heurística, uma boa estimativa de 99% do ótimo é frequentemente mais útil do que o valor exato do ótimo

## Citações-chave

> "The Held-Karp lower bound has proved to be remarkably tight for the geometric traveling salesman problem, often achieving more than 99% of the optimal tour length."

> "We show that the characteristics of the underlying point set, particularly the degree of clustering, have a significant effect on the quality of the Held-Karp bound."

> "Estimating the Held-Karp bound rather than computing it exactly provides a practical stopping criterion for heuristic algorithms without incurring the full computational cost of Lagrangian optimization."
