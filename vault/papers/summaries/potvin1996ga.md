---
title: "Genetic Algorithms for the Traveling Salesman Problem"
authors: "Jean-Yves Potvin"
year: 1996
journal: "Annals of Operations Research"
volume: 63
pages: "337–370"
doi: "10.1007/BF02125403"
tags: [survey, genetic-algorithms, tsp, crossover-operators, metaheuristics]
---

## 1. Problema e motivação

O TSP é um problema NP-difícil clássico que serve como *benchmark* para novas meta-heurísticas. Em meados dos anos 1990, algoritmos genéticos ganhavam popularidade como paradigma de busca estocástica, e esta *survey* organiza e avalia as diversas abordagens de AG propostas até então para o TSP, cobrindo representações, operadores e hibridizações.

## 2. Método ou abordagem central

- **Representações do cromossomo:** representação por caminho (lista de cidades na ordem de visita), por adjacência (lista de arestas) e ordinal (codificação posicional relativa), cada uma com implicações distintas para os operadores aplicáveis.
- **Operadores de cruzamento:** PMX (*Partially Mapped Crossover*), OX (*Order Crossover*), CX (*Cycle Crossover*), ER (*Edge Recombination*) e variantes. O ER preserva informações de adjacência entre pais e é apontado como o mais eficaz para TSP.
- **Operadores de mutação:** troca (*swap*), inversão de subsequência, inserção e deslocamento, sendo a inversão (equivalente ao 2-opt) a mais relevante por preservar a maior parte da estrutura da rota.
- **Mecanismos de seleção:** roleta proporcional ao *fitness*, torneio e seleção por *ranking*, com discussão sobre pressão seletiva e diversidade populacional.
- **Hibridização com busca local:** incorporação de heurísticas como 2-opt, 3-opt e Lin-Kernighan dentro do ciclo evolutivo (AG híbrido ou *memetic algorithm*), mostrando ganhos substanciais sobre AGs puros.

## 3. Principais resultados

- Operadores baseados em arestas (ER e variantes) consistentemente superam operadores baseados em posição (PMX, OX, CX) em qualidade de solução para instâncias TSP.
- AGs híbridos que aplicam busca local a cada indivíduo da população produzem soluções significativamente melhores que AGs puros, aproximando-se de métodos *state-of-the-art* da época (e.g., Lin-Kernighan iterado).
- A diversidade populacional é crítica; sem mecanismos de manutenção de diversidade, a convergência prematura degrada a qualidade das soluções.
- Resultados reportados sobre instâncias euclidianas aleatórias e problemas clássicos da literatura de *Operational Research* (referências ao TSPLIB).

## 4. Pontos fortes e limitações

**Pontos fortes:** taxonomia abrangente das abordagens de AG para TSP; comparação sistemática de representações e operadores genéticos; ponte entre teoria de AG e prática de *Operational Research*; referência seminal com 323+ citações.

**Limitações:** não apresenta resultados empíricos novos — é uma *survey* puramente bibliográfica; comparações entre estudos diferentes não são padronizadas (parâmetros, hardware, critérios de parada distintos); anterior à consolidação de meta-heurísticas populacionais modernas e a *frameworks* como *memetic algorithms*.

## 5. Implicação prática para pesquisadores

Operadores de cruzamento que preservam arestas (em especial *Edge Recombination*), combinados com busca local (2-opt/LK), produzem os melhores resultados de AG para problemas de roteamento. A escolha da representação cromossômica restringe fundamentalmente quais operadores são aplicáveis — para TSP, a representação por caminho com operadores baseados em adjacência oferece o melhor compromisso entre expressividade e eficácia dos operadores.
