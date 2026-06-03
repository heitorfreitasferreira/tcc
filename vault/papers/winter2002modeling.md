---
title: "Modeling Costs of Turns in Route Planning"
authors: [Winter, Stephan]
year: 2002
doi: "10.1023/A:1020853410145"
bibtex-key: winter2002modeling
pdf: "papers/pdfs/winter2002modeling.pdf"
tags: [routing]
status: lido
rating: 0
---

## PDF

![[winter2002modeling.pdf]]

## Resumo

Winter propõe um modelo para incorporar custos de conversão entre arestas consecutivas em planejamento de rotas. O argumento central é que certos custos não pertencem naturalmente nem a nós nem a arestas isoladas: restrições de conversão, ângulo de giro, espera em conexões, raio de curva e simplicidade cognitiva da rota dependem do par “aresta anterior, aresta seguinte”. Para resolver isso sem modificar algoritmos de caminho mínimo, o artigo constrói um grafo pseudo-dual no qual arestas do grafo original viram nós e relações de continuidade entre arestas viram arestas ponderáveis. Assim, custos de giro passam a ser tratados como pesos normais no grafo transformado.

## Contribuições Principais

- Define formalmente o grafo pseudo-dual completo e restrito para representar relações entre arestas consecutivas.
- Mostra que custos de giro não são atributos adequados de nós ou arestas simples no grafo primal.
- Compara o grafo pseudo-dual com expansão de nós, argumentando que o pseudo-dual é menor e mais limpo topologicamente.
- Demonstra que algoritmos de caminho mínimo podem rodar sem modificações após a transformação.
- Discute aplicações em navegação de pedestres, restrições de conversão, minimização de número de giros e minimização de ângulo total.

## Relevância para o TCC

O artigo é relevante para extensões do rTSP do TCC em que o custo da rota não dependa apenas da distância entre pontos, mas também da transição entre segmentos. Em patrulha com drones, isso pode modelar penalidades por curvas bruscas, consumo adicional em manobras, restrições de orientação ou preferência por rotas mais suaves. A contribuição conceitual é importante: se o custo depende do par de movimentos consecutivos, a matriz simples de distâncias do TSP não representa todo o problema.

## Métodos e Abordagens

- Representação de cada aresta do grafo primal como um nó no grafo pseudo-dual.
- Criação de arestas pseudo-duais para cada par de arestas consecutivas no grafo original.
- Uso de grafos pseudo-duais restritos para impedir retornos imediatos indesejados.
- Combinação de custo de deslocamento e custo de giro em uma função de custo no grafo transformado.
- Comparação de requisitos de armazenamento entre pseudo-dual e expansão de nós.

## Conexões

- [[tsp]]
- [[routing]]
- [[bock2025survey]]
- [[agatz2018optimization]]
- [[murray2015flying]]
- [[kinable2017hybrid]]

## Notas e Insights

- O artigo é sobre caminho mínimo, não TSP, mas a modelagem de custo dependente de transição é transferível para roteamento mais complexo.
- A ideia de “rota simples” como rota com menos instruções diferencia menor distância de menor complexidade operacional.
- Para drones, penalidades de giro podem tornar a rota geometricamente mais realista que uma métrica puramente euclidiana.
- A principal limitação apontada é que os protótipos não são otimizados para redes grandes e multiusuário.

## Citações-chave

> “Turn costs can be represented by a pseudo-dual graph in a way that shortest path algorithms run without modifications.”

> “The pseudo-dual graph represents all pairs of consecutive edges and allows individual weighting.”
