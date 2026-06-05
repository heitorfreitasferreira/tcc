# A Patching Algorithm for the Nonsymmetric Traveling-Salesman Problem

**Autor:** Richard M. Karp  
**Ano:** 1979 (Memorandum UCB/ERL M78/24, January 1978)  
**Publicação:** University of California, Berkeley — Electronics Research Laboratory  

---

## 1. Problema e Motivação

O problema do caixeiro-viajante não simétrico (nonsymmetric TSP) é NP-difícil, e mesmo a aproximação com erro relativo uniformemente limitado é NP-difícil (Sahni & Gonzales). O artigo propõe um algoritmo de aproximação polinomial que, embora não garanta erro limitado no pior caso, tende probabilisticamente a soluções quase ótimas quando as distâncias são i.i.d. uniformes em [0,1] e o número de cidades é grande. A motivação prática inclui problemas de sequenciamento de máquinas, onde d_ij representa o custo de setup da tarefa j após a tarefa i.

## 2. Método Principal

- **Etapa 1 — Problema de Atribuição:** resolve-se o problema de atribuição n×n para a matriz de distâncias D em tempo O(n³). A permutação ótima π da atribuição pode conter múltiplos ciclos disjuntos.
- **Etapa 2 — Patching:** seleciona-se o ciclo de maior comprimento C e resolve-se um segundo problema de atribuição (k−1)×|C| para determinar o modo ótimo de unir cada um dos demais k−1 ciclos a C. O custo de cada operação de patching entre i ∈ C e j ∉ C é A_ij = d_i,π(j) + d_j,π(i) − d_i,π(i) − d_j,π(j).
- **Etapa 3 — Conversão em Tour:** aplicam-se k−1 operações de patching (que comutam) para fundir todos os ciclos em um único ciclo hamiltoniano, produzindo a permutação cíclica τ.
- **Saída:** a permutação τ e um limitante superior do erro relativo ε(D) = (c(τ, D) − c(τ*, D)) / c(τ*, D).
- **Algoritmo Modificado (Seção 4):** define d_ii = ∞ para eliminar pontos fixos e, iterativamente, une o ciclo mais curto a outro usando a operação de patching de menor custo disponível, sem restringir as junções ao ciclo mais longo.

## 3. Principais Resultados

- **Teorema 1 (Original):** para distâncias i.i.d. uniformes em [0,1], com probabilidade tendendo a 1 quando n → ∞, o erro relativo satisfaz ε(D) < 9√(8+2√7) · (ln n)² · n^(−1/2). O limitante tende a zero, mas muito lentamente — só é aceitavelmente pequeno para n astronomicamente grande.
- **Teorema 3 (Modificado — heurístico):** sob independência dos custos de patching, lim sup n^(1/2) · E[Y_n] < 2, onde Y_n é o custo total do patching modificado.
- **Simulação de Monte Carlo (100 execuções por n):**

| n      | 100  | 1.000 | 10.000 |
|--------|------|-------|--------|
| Média  | 0,18 | 0,067 | 0,018  |
| Média × √n | 2,1 | 1,8 | 1,8 |
| Mediana × √n | 1,6 | 2,0 | 1,7 |
| Máximo × √n | 5,4 | 4,9 | 4,0 |

- O custo esperado do patching decresce com ≈ 2/√n, consistente com a análise heurística.

## 4. Pontos Fortes e Limitações

**Pontos fortes:**
- Complexidade O(n³), comparável à resolução do problema de atribuição.
- Primeiro algoritmo de aproximação para TSP assimétrico com garantia probabilística de quase-otimalidade.
- A análise probabilística rigorosa usando distribuição de ciclos de permutações aleatórias, cotas de Chernoff e dominância estocástica é tecnicamente elegante.

**Limitações:**
- O limitante de erro do Teorema 1 converge muito lentamente (termo em (ln n)² / √n); utilidade prática requer n extremamente grande.
- A análise probabilística supõe distribuição uniforme [0,1]; não há garantia de pior caso.
- O algoritmo original pode falhar se nenhum ciclo tiver comprimento ≥ k (número de ciclos), embora esse evento tenha probabilidade desprezível.
- O algoritmo modificado é analisado apenas heuristicamente; a conjectura de dominância estocástica Z ≺ Y permanece em aberto.

## 5. Implicações Práticas para Pesquisadores

O algoritmo de patching oferece um paradigma reutilizável: reduzir TSP a um problema de atribuição (polinomial) e depois corrigir os ciclos com operações locais de baixo custo esperado. Para instâncias aleatórias de grande porte, a abordagem tende a produzir tours quase ótimos. Em aplicações práticas com centenas de cidades, o algoritmo modificado com junção iterativa do ciclo mais curto é preferível ao original, pois evita a restrição de unir todos os ciclos ao ciclo mais longo e mostrou empiricamente custo médio de patching decrescente com √n.

**Palavras-chave:** traveling-salesman problem, combinatorial optimization, approximation algorithms, probabilistic analysis of algorithms.
