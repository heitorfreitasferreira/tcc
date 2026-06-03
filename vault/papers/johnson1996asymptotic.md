---
title: "Asymptotic Experimental Analysis for the Held-Karp Traveling Salesman Bound"
authors: [Johnson, David S., McGeoch, Lyle A., Rothberg, Edward E., Schreiber, Robert]
year: 1996
doi: "10.1137/1.9781611971486"
bibtex-key: johnson1996asymptotic
pdf: "papers/pdfs/johnson1996asymptotic.pdf"
tags: [tsp, lower-bound, held-karp, experimental]
status: lido
rating: 5
---

## Resumo

Estudo empírico em larga escala do Held-Karp lower bound para TSP, computando-o exatamente (via simplex + separação de subtour) para instâncias de até 30.000 cidades, e aproximadamente (via subgradiente) para até 1 milhão de cidades. Demonstra que o ótimo do TSP está em média < 0.8% acima do HK bound para instâncias Euclideanas aleatórias, e < 2% para instâncias TSPLIB. Estima a constante assintótica C_OPT ≈ 0.7124 ± 0.0002 para o TSP Euclideano 2D, corrigindo estimativas anteriores.

## Contribuições Principais

- Validação empírica definitiva: HK bound ≈ ótimo do TSP (gap < 0.8%)
- Estimativa do limiting ratio C_HK = lim_{N→∞} E[HK]/√N para várias classes de instâncias
- Correção da constante assintótica do TSP Euclideano: C_OPT ≈ 0.7124
- Demonstração de que o HK bound é adequado como "stand-in" para o ótimo em estudos experimentais

## Relevância para o TCC

Estabelece que, para qualquer método de lower bound implementado, o gap em relação ao ótimo precisa ser avaliado. Os resultados do paper servem como referência: se o gap da relaxação escolhida (AP, HK reduzido, etc.) for consistentemente maior que 2-5%, o bound precisa ser melhorado.

## Métodos e Abordagens

- Solução exata do LP de subtour (simplex + separação por fluxo máximo)
- Aproximação via subgradiente (Lagrangiano) com MSTs em subgrafo esparso
- Geração de instâncias: uniforme 2D/3D/4D, distâncias Euclideanas, retilíneas, supremo, matrizes aleatórias
- TSPLIB como conjunto de instâncias reais
- Regressão estatística para estimativas assintóticas

## Conexões

- [[heldkarp1970traveling]] — bound HK original
- [[heldkarp1971traveling]] — subgradiente para bound HK
- [[valenzuela1997estimating]] — implementação prática alternativa de subgradiente
- [[righini2021efficient]] — otimização da seleção de vértice
- [[applegate2006traveling]] — Concorde e computação exata do subtour LP
- [[tsp]]
- [[lower-bounds]]

## Notas e Insights

- Na prática, usar o HK bound como substituto do ótimo é válido para avaliar heurísticas
- O gap depende do tipo de instância: Euclideana < 0.8%, TSPLIB < 2%, matriz aleatória ≈ 2.042
- A computação exata do HK bound requer código especializado (Concorde), mas a aproximação subgradiente funciona bem
- Para instâncias com n < 1000, o subgradiente obtém tipicamente 99.5%+ do bound exato

## Citações-chave

> "We provide empirical evidence in support of using the HK bound as a stand-in for the optimal tour length when evaluating the quality of near-optimal tours."

> "For a wide variety of randomly generated instance types the optimal tour length averages less than 0.8% over the HK bound, and even for the real-world instances in TSPLIB the gap is almost always less than 2%."

![[johnson1996asymptotic.pdf]]
