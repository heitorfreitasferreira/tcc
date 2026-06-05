# Estimating the Held-Karp lower bound for the geometric TSP

**Authors:** Christine L. Valenzuela, Antonia J. Jones  
**Journal:** European Journal of Operational Research 102 (1997), 157–175  
**DOI:** S0377-2217(96)00214-7

## 1. Problem and motivation

Avaliar a qualidade de soluções aproximadas para o TSP em instâncias grandes exige um limitante inferior confiável. O limitante de Held-Karp (HK) pode ser calculado exatamente via programação linear, mas implementações eficientes de PL não estão disponíveis para problemas acima de algumas centenas de cidades e não escalam bem. O artigo fornece diretrizes práticas para estimar HK via otimização de subgradiente baseada em 1-árvores, tornando-o acessível para instâncias com milhares de cidades.

## 2. Core method or approach

- **Relaxação lagrangiana via 1-árvore mínima:** uma 1-árvore é uma MST sobre os vértices 2..n mais as duas arestas de menor custo incidentes ao vértice 1. Um tour é uma 1-árvore onde todo vértice tem grau 2.
- **Otimização de subgradiente:** pesos π_i são atribuídos aos vértices e iterativamente atualizados com base no desvio do grau da 1-árvore em relação a 2 (aumenta se grau > 2, diminui se grau < 2). O processo força a 1-árvore a aproximar-se de um tour, maximizando w(π) = min_T {C_T + Σ π_i(d_i^T − 2)}.
- **Duas fórmulas de iteração comparadas:** (i) Held-Wolfe-Crowder (HWC) — utiliza um limitante superior U e λ_m decrescente; (ii) Volgenant-Jonker (VJ) — sequência decrescente de passos sem necessidade de limitante superior, com termo de amortecimento baseado no grau da iteração anterior.
- **Técnica de aceleração por subgrafo:** em vez de grafo completo, usa-se apenas os 20 vizinhos mais próximos por cidade para instâncias aleatórias (40 para TSPLIB, com representantes de cada quadrante), reduzindo a complexidade de O(n²) para O(n log n) por iteração.
- **Passo final:** uma única avaliação O(n²) sobre o grafo completo usando os melhores pesos π obtidos, produzindo HK(completegraph).

## 3. Main results

- **Qualidade do limitante:** em média, HK fica dentro de 0,8% do ótimo para instâncias euclidianas aleatórias com milhares de cidades. Para instâncias da TSPLIB, o desvio do HK exato é geralmente muito abaixo de 1%, com piores casos (~2%) em distribuições patológicas (ex.: dsj1000, fl3795).
- **Fórmulas equivalentes:** HWC e VJ produzem resultados praticamente idênticos para instâncias aleatórias e TSPLIB; VJ é preferível por eliminar a necessidade de estimar um limitante superior.
- **Vizinhos suficientes:** 20 vizinhos mais próximos bastam para pontos aleatórios no quadrado unitário; instâncias TSPLIB com agrupamentos requerem ~40 vizinhos e representação garantida dos quatro quadrantes.
- **Comprimento de sequência:** M ≈ 28n^0.62 como guia; para n entre 100 e 10.000, sequências de 100–300 iterações são frequentemente suficientes. Sequências muito curtas (100) já produzem excelentes aproximações.
- **Constante de Goemans-Bertsimas:** c₂ ≈ 0,70787 ± 0,0008, confirmando independentemente a estimativa de Johnson et al. (1996) de 0,70805 ± 0,00007.
- **Desempenho computacional:** ~5 s para n=100, ~1,5 h para n=5000 em SPARC 10. Em todos os experimentos com pontos aleatórios, HK(subgraph) = HK(completegraph).

## 4. Strengths and limitations

**Strengths:**
- Extremo detalhamento de implementação (pseudocódigo completo, derivação das fórmulas de passo).
- Validação em duas famílias complementares de instâncias (aleatórias uniformes e TSPLIB).
- Robustez demonstrada pela equivalência prática de duas fórmulas de iteração independentes.
- Confirmação independente da constante c₂, reforçando a credibilidade dos resultados.

**Limitations:**
- O gap entre HK e o ótimo pode chegar a ~2% em instâncias com distribuições patológicas de cidades.
- O cálculo final O(n²) sobre grafo completo impõe restrições de memória para n > 10.000.
- A estimativa da constante c₁ (comprimento esperado do tour ótimo) permanece computacionalmente inviável para n grande.
- A fórmula-guia M ≈ 28n^0.62 é empírica e válida apenas no intervalo 100–5.000 cidades.

## 5. Practical takeaway for researchers

A fórmula de Volgenant-Jonker é a escolha recomendada por não exigir limitante superior. Para instâncias aleatórias, 20 vizinhos mais próximos são suficientes; para instâncias com estrutura irregular, usar ~40 vizinhos com representação forçada dos quatro quadrantes. Sequências curtas (100–300 iterações) produzem estimativas de HK com erro desprezível para a maioria das aplicações práticas. O passo final de grafo completo pode ser omitido para pontos aleatórios, pois HK(subgraph) converge para o mesmo valor.
