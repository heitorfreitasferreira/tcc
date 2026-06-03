---
title: "MAX-MIN Ant System"
authors: [Stützle, Thomas, Hoos, Holger H.]
year: 2000
doi: "10.1016/S0167-739X(00)00043-1"
bibtex-key: stutzle2000mmas
tags: [aco tsp]
status: lido
rating: 4
pdf: "papers/pdfs/stutzle2000mmas.pdf"
---

## PDF

![[stutzle2000mmas.pdf]]

## Resumo

Apresenta o MAX-MIN Ant System (MMAS), uma das variantes ACO mais bem-sucedidas. MMAS difere do AS original em três aspectos: (1) apenas a melhor formiga deposita feromônio (iteration-best ou global-best), (2) valores limitados a [τ_min, τ_max] para evitar estagnação, (3) feromônios inicializados em τ_max para promover exploração. MMAS supera AS e compete com ACS, especialmente em problemas grandes onde o risco de convergência prematura é maior. O artigo também introduz o uso de reinicializações when estagnação é detectada.

## Contribuições Principais

- Introdução do mecanismo de limites de feromônio [τ_min, τ_max] como forma de evitar estagnação
- Demonstração de que apenas a melhor formiga depositar feromônio acelera convergência sem sacrificar qualidade
- Inicialização em τ_max como estratégia de exploração inicial
- Competitivo ou superior ao ACS em instâncias TSP padrão

## Relevância para o TCC

Justifica a escolha do MMAS como variante ACO, especialmente relevante para instâncias grandes ou com muitos ótimos locais (rTSP). Os limites de feromônio são a principal inovação para balanceamento exploração/explotation.

## Métodos e Abordagens

- Atualização: apenas iteration-best ou global-best deposita feromônio
- Limites: τ_max = 1/(ρ·L_best), τ_min = τ_max / (2n)
- Inicialização: τ₀ = τ_max (exploração máxima no início)
- Reinicialização: quando τ_ij ≈ τ_min ≈ τ_max, reinicia feromônios
- Combinado com busca local 2-opt/3-opt para melhores resultados

## Conexões

- [[dorigo1996ant]] — Ant System (base)
- [[dorigo1997ant]] — ACS (variante contemporânea)
- [[blum2005acointro]] — introduction covering MMAS
- [[dorigo2004book]] — livro referência
- [[ant-colony]]
- [[tsp]]

## Notas e Insights

- MMAS introduz três inovações simples mas eficazes: (1) apenas a melhor formiga deposita feromônio, (2) limites [τ_min, τ_max] para evitar estagnação, (3) inicialização em τ_max para exploração máxima
- O mecanismo de limites de feromônio é a principal contribuição e é amplamente adotado em implementações ACO modernas
- A análise de fitness-distance correlation (FDC) para TSP e QAP mostra que a correlação entre qualidade e distância do ótimo justifica a exploração das melhores soluções
- Para o TCC, MMAS é a variante ACO mais adequada para o rTSP, pois instâncias grandes e complexas se beneficiam dos limites de feromônio para evitar convergência prematura
- A fórmula τ_max = 1/(ρ·L_best) e τ_min = τ_max / (2n) é implementada diretamente no código ACO do TCC
- O artigo mostra que MMAS combinado com busca local 2-opt atinge resultados competitivos com ACS, validando a escolha de MMAS+2-opt para o TCC

## Citações-chave

> "The use of a rather simple mechanism for limiting the strengths of the pheromone trails effectively avoids premature convergence of the search."

> "MMAS achieves a strongly improved performance compared to AS and is among the best available algorithms for the QAP."

> "The exploitation of the best solutions found during the search strongly improves performance, but must be combined with effective mechanisms for search space exploration to avoid premature convergence."
