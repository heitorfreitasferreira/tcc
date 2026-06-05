---
title: 'The Traveling-Salesman Problem and Minimum Spanning Trees: Part II'
authors:
- Held
- Michael
- Karp
- Richard M.
year: 1971
doi: 10.1007/BF01584070
bibtex_key: heldkarp1971traveling
bibtex-key: heldkarp1971traveling
pdf: papers/pdfs/heldkarp1971traveling.pdf
tags:
- area/tsp
- evidencia/referencia
- metodo/lagrangean
- metodo/lower-bound
- metodo/branch-and-bound
- metodo/relaxation
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 5
classificacao: recuperado
type: paper
areas:
- tsp
- lower-bound
methods:
- lagrangean
- lower-bound
- branch-and-bound
- relaxation-method
- 1-tree
role: revisao
chapters:
- fundamentacao
claim_support: []
aliases:
- Held-Karp Part II
- Held-Karp 1971
reading_status: resumo-lido
validation_status: nao-validado
pdf_status: lido
---

## Resumo

Publicado em Mathematical Programming 1 (1971) 6–25, North-Holland. **Parte II** da série Held-Karp, dá sequência ao artigo de 1970 ([7] na bibliografia do PDF) que estabelecera a relação TSP ↔ minimum 1-tree e a fórmula do *lower bound* `C* ≥ w(π) = min_k [c^k + π·v^k]`. Este paper apresenta:

1. Um **método de ascensão** (*ascent method*) iterativo para aproximar `max_π w(π)` de baixo, relacionado ao *relaxation method* de Agmon (1954) e Motzkin-Schoenberg (1954) para sistemas de desigualdades lineares. Iteração: `π^{m+1} = π^m + t_m v^{k(π^m)}`, onde `k(π^m)` é o índice do 1-tree de peso mínimo nos pesos modificados `c_ij + π_i + π_j`. Os autores demonstram convergência (Lemma 3) e propriedades de Fejér-monotonicidade.
2. Um **procedimento branch-and-bound** que combina o ascent method com particionamento por inclusão/exclusão de arestas (segue Bellmore-Nemhauser 1968). Em cada nó: aplica o ascent para obter `w_{X,Y}(π)`; se `w_{X,Y} ≥ C` (upper bound), descarta; senão, ramifica incluindo/excluindo arestas.
3. Resultados computacionais: o algoritmo **resolve problemas de até 64 cidades** (Dantzig 42, Random Euclidean 64, Held-Karp 48, Karg-Thompson 57, 8×8 Knight's tour 64, Join 64) provando o ótimo. As árvores de busca são minúsculas ("minuscule compared to those normally encountered in combinatorial problems"), atestando o *gap* pequeno entre o bound de Held-Karp e o ótimo.

**Atenção**: a formulação `O(n²2^n)` por *dynamic programming* mencionada em algumas referências (Held-Karp 1962) é de um **artigo separado** — referência [8] deste paper: *"Held, M. and Karp, R.M., 1962. A dynamic programming approach to sequencing problems. J. SIAM 10, 196–210"*. **Este paper de 1971 trata do ascent method + branch-and-bound para o *lower bound* Lagrangiano**, não da DP.

## Contribuições Principais

- Ascent method iterativo para `max w(π)`, com prova de convergência (Lemma 3) e Teorema 1
- Procedimento branch-and-bound que produz soluções ótimas certificadas para TSP simétrico
- Melhoria do bound HK: os autores observam que o ascent produz `w(π)` muito próximo de `max w(π)`, e que a maioria do gap `C* - max w(π)` é intrínseco (não do ascent)
- Soluções ótimas certificadas em todos os problemas testados, até 64 cidades
- Possíveis melhorias discutidas: starting point via assignment problem dual, passo `t_m` variável, ordenação dos custos para Kruskal

## Relevância para o TCC

Referência canônica de lower bound para TSP. O `C* - max w(π)` é o *gap* que define o quão bom é o bound HK; empiricamente (Johnson et al. 1996) é < 0.8% em instâncias euclidianas aleatórias. O método de ascensão deste paper é a base dos algoritmos modernos (Applegate/Concorde, Righini 2021) que computam o LP de subtour para TSP. Para o TCC, é a justificativa teórica do uso de `max w(π)` como aproximação do ótimo.

## Métodos e Abordagens

- 1-tree: árvore com `{2, 3, ..., n}` mais duas arestas no vértice 1; `c_ij + π_i + π_j` muda o peso do 1-tree de mínimo sem alterar o peso de tours
- Bound: `C* ≥ w(π) = min_k [c^k + π·v^k]` onde `v^k_i = d^k_i - 2` (desvio de grau 2)
- Ascent method: `π^{m+1} = π^m + t_m v^{k(π^m)}` com `t_m` constante (= 1 tipicamente)
- Relaxation method: equivalente a resolver `max w(π)` por violação iterativa das restrições lineares
- Branch-and-bound: particionamento `(X, Y, π, w_{X,Y}(π))`, com seleção de nó de menor bound e regra de branching que exclui arestas
- Heurística: qualquer procedimento (Lin, Croes, etc.) fornece o upper bound `C` inicial
- Complexidade do ascent: o método depende de sucessivos 1-trees mínimos — `O(n²)` cada (Dijkstra) ou `O(n log n)` com pré-ordenação (Kruskal)

## Conexões

- [[heldkarp1970traveling]] — paper original (referência [7] deste PDF), formula o bound HK
- [[johnson1996asymptotic]] — análise empírica em larga escala do gap HK (usa o ascent method)
- [[valenzuela1997estimating]] — implementação prática alternativa do ascent
- [[righini2021efficient]] — otimização da seleção de vértice p para HK bound
- [[applegate2006traveling]] — Concorde, branch-and-cut com planos de corte
- [[tsp]]
- [[lower-bounds]]
- [[karp1979patching]] — patching analysis, outro lower bound para TSP

## Notas e Insights

- O bound é tão apertado que o branch-and-bound "produziu soluções ótimas certificadas em todos os TSPs apresentados, variando até 64 cidades" (PDF, Abstract)
- As árvores de busca são publicadas no paper (Figs. 2–8) — "primeira vez que se apresenta uma árvore de busca inteira para um problema combinatório de grande porte"
- O ascent usa passo constante `t_m = 1`, que os autores reconhecem como "naive" mas que funciona bem na prática
- O paper nota que "the bounds computed in the initial ascent are extremely close to `max w(π)`" — a maior parte do gap `C* - w(π)` é intrínseco à relaxação 1-tree
- Trabalhos posteriores (Held-Karp 1970, depois [Johnson et al. 1996]) generalizaram para até 1 milhão de cidades
- Trabalhos posteriores: column generation (equivalente a simplex) pode ser mais rápido para instâncias grandes, mas o ascent method é surpreendentemente competitivo
- Decisão P46: preservar a nota canônica e o DOI `10.1007/BF01584070`; o summary registra DOI como `[to be verified]`, mas BibTeX e nota canônica já estão alinhados com a versão corrigida em P43.

## Citações-chave

> "An efficient iterative method for approximating this bound closely from below is presented. A branch-and-bound procedure based upon these considerations has easily produced proven optimum solutions to all traveling-salesman problems presented to it, ranging in size up to sixty-four cities."

> "The bounds used are so sharp that the resulting search trees are minuscule compared to those normally encountered in combinatorial problems of this type."

> "In fact, the experience with the traveling-salesman problem indicates that some form of the relaxation method may be superior to the simplex method for linear programs involving a very large number of inequalities."
