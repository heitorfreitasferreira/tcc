---
title: Ant Colony Optimization for Traveling Salesman Problem Based on Parameters Optimization
authors:
- Wang
- Yong
year: 2021
doi: 10.1016/j.asoc.2021.107439
bibtex_key: wang2021ant
bibtex-key: wang2021ant
pdf: papers/pdfs/wang2021ant.pdf
tags:
- area/tsp
- evidencia/referencia
- metodo/aco
- metodo/metaheuristic
- status/lido
- tipo/paper
status: lido
rating: 3
type: paper
methods:
- aco
- metaheuristic
role: revisao
reading_status: lido
validation_status: nao-validado
pdf_status: integro
areas:
- tsp
- aco
chapters:
- fundamentacao
claim_support: []
aliases:
- SOS-ACO
- ACO parameter optimization
---

## PDF

![[wang2021ant.pdf]]

## Resumo

Propõe um algoritmo híbrido SOS-ACO para TSP, onde Symbiotic Organisms Search (SOS) é usado para otimizar os parâmetros chave (α e β) do Ant Colony Optimization. Estratégia de otimização local acelera convergência e melhora qualidade. Resultados em instâncias TSPLIB mostram que o SOS-ACO supera ACO puro e ACO-LO.

## Contribuições Principais

- Uso de SOS para tuning automático de parâmetros do ACO
- Estratégia de otimização local integrada
- Validação em TSPLIB com resultados superiores a baseline ACO

## Evidência / Resultado Relevante

- O SOS-ACO usa Symbiotic Organisms Search para otimizar `α` e `β`, dois parâmetros centrais da construção de soluções em ACO.
- A escolha de SOS é motivada por ser uma meta-heurística sem parâmetros próprios de controle, evitando transferir o problema de tuning para outro algoritmo parametrizado.
- O summary registra testes em 10 instâncias TSPLIB, de 51 a 575 cidades, com erro máximo de 2,33% em relação ao melhor conhecido; validar esse número no PDF antes de usar como claim quantitativo na monografia.

## Limitações de Uso

- O artigo otimiza apenas `α` e `β`; outros parâmetros de ACO continuam definidos pelo usuário.
- O método combina tuning e busca local, portanto não deve ser comparado diretamente ao Ant System puro implementado no repositório.
- O estudo não substitui uma análise de sensibilidade própria do TCC.

## Relevância para o TCC

Abordagem de tuning automático de parâmetros é relevante para a implementação de ACO no repositório. A sensibilidade paramétrica do ACO é um desafio prático no TCC.

## Métodos e Abordagens

- ACO com parâmetros α e β otimizados por SOS
- Estratégia de otimização local
- Benchmarks TSPLIB (Lin318, Rd400, Pr439, Rat575)

## Conexões

- [[dorigo1996ant]] — Ant System, fundação do ACO
- [[dorigo1997ant]] — primeira aplicação do ACO ao TSP
- [[kappagantula2025dpso]] — DPSO-Q (outra hibridização swarm+RL)
- [[ant-colony]] [[ant-colony]]
- [[tsp]]

## Notas e Insights

- ACO é altamente sensível aos parâmetros α (influência do feromônio) e β (influência da heurística)
- SOS-ACO reduz erro relativo em instâncias grandes comparado a ACO padrão
- Tuning automático pode ser uma direção para melhorar o ACO implementado no TCC
- Decisão P46: incorporar como referência útil para P33/tuning, sem transformar os resultados do paper em benchmark direto do projeto.

## Citações-chave

>
