---
title: "Automated Design of State Transition Rules in Ant Colony Optimization by Genetic Programming"
authors: [Lin, Bo-Cheng, Mei, Yi, Zhang, Mengjie]
year: 2025
doi: "10.1007/s12293-025-00435-9"
bibtex-key: gpaco2025
pdf: "papers/pdfs/gpaco2025.pdf"
tags: [aco metaheuristic gp]
status: lido
rating: 4
---

## PDF

![[gpaco2025.pdf]]

## Resumo

O artigo propõe GP-ACO, uma abordagem para desenhar automaticamente regras de transição de estado em [[ant-colony]] usando programação genética. O foco é reduzir a dependência de conhecimento especialista, dados pré-resolvidos e modelos pouco interpretáveis. A aplicação experimental é o [[tsp]], usando variantes clássicas de ACO como AS, ACS e MMAS, além de comparações com heurísticas projetadas por LLM. Uma extensão, xGP-ACO, adiciona informações globais como distância média, feromônio médio, número de nós e número de candidatos.

## Contribuições Principais

- Formula um processo para evoluir regras de transição de ACO via programação genética.
- Avalia três variantes de ACO: AS, ACS e MMAS.
- Testa instâncias TSP sintéticas e 15 instâncias TSPLIB.
- Propõe xGP-ACO com terminais globais para melhorar as regras aprendidas.
- Analisa interpretabilidade das árvores evoluídas.

## Relevância para o TCC

Este artigo é diretamente relevante para a discussão de [[ant-colony]] aplicada ao TSP. Embora o TCC implemente uma versão mais clássica de ACO, `gpaco2025` mostra uma direção moderna: automatizar a regra que equilibra feromônio e informação heurística. Para patrulha com drones modelada como TSP/rTSP, a contribuição útil é conceitual: a regra de escolha de próximo nó pode incorporar informação global da instância.

## Métodos e Abordagens

- ACO aplicado ao [[tsp]].
- Programação genética para gerar regras de transição.
- Terminais básicos: feromônio e distância.
- Terminais globais em xGP-ACO: distância média, feromônio médio, número total de nós e número de nós candidatos.
- Comparações em AS, ACS e MMAS, com e sem 2-opt.

## Conexões

- [[dorigo1996ant]]
- [[dorigo1997ant]]
- [[stutzle2000mmas]]
- [[blum2005acointro]]
- [[dorigo2005acotheory]]
- [[deepaco2023]]
- [[neufaco2025]]
- [[ppaco2024]]

## Notas e Insights

- O artigo reforça que a regra de transição é componente crítico da ACO, não apenas detalhe de parametrização.
- A comparação entre AS, ACS e MMAS ajuda a explicar que variantes dependem em graus diferentes da regra de transição e da atualização de feromônio.
- O 2-opt melhora a estabilidade, mas pode reduzir a pressão seletiva sobre a regra de transição.
- A inclusão de informação global melhora xGP-ACO, sugerindo que regras puramente locais podem ser insuficientes em instâncias complexas.

## Citações-chave

> “The automated design of Ant Colony Optimization (ACO) algorithms has become increasingly significant, particularly in addressing complex combinatorial optimization problems.”

> “GP-ACO and xGP-ACO exhibit excellent interpretability.”
