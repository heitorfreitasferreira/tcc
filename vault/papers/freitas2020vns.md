---
title: "A Variable Neighborhood Search for Flying Sidekick Traveling Salesman Problem"
authors: [de Freitas, Júlia Cária, Penna, Puca Huachi Vaz]
year: 2020
doi: "10.1111/itor.12671"
bibtex-key: freitas2020vns
pdf: "papers/pdfs/freitas2020vns.pdf"
tags: [tsp drone]
status: lido
rating: 4
---

## PDF

![[freitas2020vns.pdf]]

## Resumo

Propõe uma heurística híbrida (HGVNS) para o *Flying Sidekick Traveling Salesman Problem* (FSTSP), variante do TSP onde um caminhão e um drone realizam entregas colaborativamente. A solução inicial é obtida resolvendo o TSP clássico de forma exata (com o solver Concorde) e, em seguida, um *General Variable Neighborhood Search* (GVNS) com *Randomized Variable Neighborhood Descent* (RVND) otimiza as rotas combinadas de caminhão e drone. Experimentos mostram redução de até 67,79% no tempo total de entrega, estabelecendo novos *best-known solutions* (BKS) para todas as instâncias da literatura.

## Contribuições Principais

- Heurística HGVNS que combina solução exata do TSP com GVNS para o FSTSP
- Novos BKS para todas as instâncias FSTSP da literatura (Ponza 2016, Agatz et al. 2016)
- Redução de até 30,38% sobre a rota ótima de caminhão (TSP tradicional)
- Novo conjunto de 25 instâncias FSTSP baseadas no TSPLIB (51 a 200 nós)
- Demonstração de que a meta-heurística VNS é eficaz para problemas de roteamento com drones

## Relevância para o TCC

Conexão direta: o FSTSP é um dos problemas de roteamento com drones que motivam o rTSP usado no TCC. A abordagem VNS ilustra como meta-heurísticas podem resolver variantes do TSP com restrições adicionais (bateria do drone, sincronização caminhão-drone). A técnica de RVND pode inspirar operadores de busca local para os métodos GA, PSO e ACO implementados.

## Métodos e Abordagens

- Mixed-Integer Programming (MIP) com Concorde para solução exata do TSP subjacente
- General Variable Neighborhood Search (GVNS) com Randomized Variable Neighborhood Descent (RVND)
- Múltiplas estruturas de vizinhança para remoção/inserção de rotas de drone
- Busca local Best Improvement (BI) em cada vizinhança
- Heurística de criação de solução inicial (adaptação da abordagem de Murray & Chu 2015)

## Conexões

- [[murray2015flying]] — define o FSTSP, problema resolvido aqui
- [[agatz2018optimization]] — mesma variante de problema
- [[dellamico2021multiple]] — extensão para múltiplos drones
- [[dellamico2022exact]] — modelos exatos para o mesmo problema
- [[applegate2006traveling]] — Concorde usado como solver exato na fase inicial
- [[lin1973effective]] — VNS usa busca local (inspirada em LK)
- [[drone-routing]]

## Notas e Insights

- Preprint arXiv:1804.03954 (versão do autor anterior à publicação final) — PDF incluso no vault
- A abordagem híbrida (exato + meta-heurística) é eficaz para problemas com estrutura TSP subjacente
- RVND escolhe aleatoriamente a próxima vizinhança, evitando viés de ordenação
- O drone é significativamente mais rápido que o caminhão, mas limitado por bateria (endurance)
- O ganho do drone é maior em instâncias com muitos clientes (paralelismo nas entregas)
- A abordagem HGVNS inspirou o design dos operadores de busca local nos métodos bio-inspirados do TCC

## Citações-chave

> "The efficiency and dynamism of unmanned aerial vehicles, or drones, have presented substantial application opportunities in several industries in the last years."

> "Computational experiments show the potential of the algorithm to improve the total delivery time up to 67.79%."
