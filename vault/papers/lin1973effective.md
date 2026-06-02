---
title: "An Effective Heuristic Algorithm for the Traveling-Salesman Problem"
authors: [Lin, Shen]
year: 1973
doi: "10.1287/opre.21.2.498"
bibtex-key: lin1973effective
tags: [tsp metaheuristic]
status: lido
rating: 5
pdf: "papers/pdfs/lin1973effective.pdf"
---

## PDF

[[papers/pdfs/lin1973effective.pdf]]

## Resumo

O artigo introduz a heurística Lin-Kernighan (LK) para o Problema do Caixeiro Viajante Simétrico, considerada por décadas o algoritmo heurístico mais eficaz para TSP. A ideia central é generalizar o k-opt tradicional: em vez de fixar k antecipadamente, o algoritmo determina sequencialmente o conjunto ótimo de arestas a serem trocadas usando um critério de ganho acumulado. A heurística produz soluções ótimas para todos os problemas clássicos testados (até 110 cidades) com tempo de execução O(n^2.2). O artigo também introduz refinamentos como lookahead, redução por interseção de soluções e backtracking limitado.

## Contribuições Principais

- Heurística Lin-Kernighan (LK) para TSP baseada em trocas sequenciais de k arestas com k variável
- Critério de ganho acumulado (soma parcial positiva) como mecanismo de parada elegante e eficaz
- Técnicas de redução: uso de arestas comuns a múltiplos ótimos locais para guiar busca
- Demonstração empírica: solução ótima para problemas clássicos com alta frequência (até 100 cidades)
- Melhoria de soluções ótimas conhecidas para 3 dos 5 problemas de 100 cidades de Krolak et al.
- Aplicação a problema real de roteamento de máquina de perfuração a laser (318 pontos)

## Relevância para o TCC

O LK é o padrão-ouro contra o qual metaheurísticas como GA, PSO e ACO são comparadas em qualidade de solução para TSP. No contexto do TCC, entender o LK é essencial para: (1) calibrar expectativas sobre a qualidade das soluções encontradas pelos métodos bio-inspirados, (2) justificar o uso de 2-opt e 3-opt como busca local nos operadores dos algoritmos populacionais, e (3) contextualizar os resultados — se GA/PSO/ACO se aproximam do LK, isso é evidência de eficácia. O problema real de 318 pontos (perfuração a laser) é análogo ao roteamento de drones: ambos exigem boas soluções em tempo viável.

## Métodos e Abordagens

- Troca sequencial de arestas: a cada passo i, escolhe xi (aresta a remover) e yi (aresta a adicionar) para maximizar ganho acumulado G_k = Σ g_i
- Critério de ganho: g_i = |x_i| - |y_i|; só aceita se G_k > 0 e todas as somas parciais são positivas
- Backtracking limitado nos níveis 1 e 2 (até 5 alternativas para y₁ e y₂)
- Lookahead: escolhe y_i maximizando |x_{i+1}| - |y_i| (não apenas menor |y_i|)
- Redução: fixa arestas comuns a 2-5 soluções ótimas locais para podar a busca
- Testes: problemas clássicos (20-100 cidades), pontos aleatórios no quadrado unitário, matrizes aleatórias não-métricas, problema real de 318 pontos

## Conexões

- [[lawler1985traveling]] — survey clássico do TSP
- [[applegate2006traveling]] — estudo computacional do TSP (Concorde)
- [[garey1979computers]] — NP-completude do TSP
- [[araujo2025pso]] — usa 2-opt e 3-opt como busca local
- [[sun2024hybrid]] — usa 2-opt como busca local
- [[freitas2020vns]] — VNS usa busca local inspirada em LK
- [[TSP]]
- [[dorigo1997ant]] — ACS compara resultados com LK como baseline

## Notas e Insights

- A heurística LK é tão eficaz que problemas de 100 cidades têm probabilidade ~50% de encontrar o ótimo em uma tentativa aleatória
- A redução por interseção de soluções é uma ideia poderosa: arestas que aparecem em múltiplos ótimos locais têm alta probabilidade de pertencer ao ótimo global
- O LK produz resultados tão bons que Held e Karp (1970) usaram-no como guia para seu algoritmo exato de branch-and-bound
- O artigo antecipa conceitos de VNS (Variable Neighborhood Search) e GRASP com sua alternância entre construção aleatória e busca local
- O problema real de 318 pontos (Figs. 6-7) mostra a aplicabilidade prática do TSP: a rota manual vs. a rota otimizada têm diferença de ~2% em 42 polegadas de deslocamento total
- Limitação: o LK é específico para TSP simétrico; sua generalização para problemas assimétricos ou variantes (rTSP) não é trivial

## Citações-chave

> "The procedure produces optimum solutions for all problems tested, 'classical' problems appearing in the literature, as well as randomly generated test problems, up to 110 cities."

> "The better the heuristic is, the smaller the set of local optima will be, and the higher will be the fraction of random starts that lead to the global optimum."
