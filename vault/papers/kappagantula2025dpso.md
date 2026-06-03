---
title: "{DPSO-Q}: A Reinforcement Learning--Enhanced Swarm Algorithm for Solving the Traveling Salesman Problem"
authors: [Kappagantula, Sivayazi]
year: 2025
doi: "10.1155/int/8918171"
bibtex-key: kappagantula2025dpso
tags: [pso tsp]
status: lido
rating: 4
---

## Resumo

Propõe o DPSO-Q, algoritmo que combina Discrete Particle Swarm Optimization (DPSO) com reinforcement learning inspirado no Ant-Q para resolver o TSP. O Q-learning é usado para ajustar dinamicamente os parâmetros de busca do DPSO, balanceando exploração versus explotação ao longo das iterações. Resultados experimentais mostram redução de até 7.5% no comprimento das rotas comparado ao DPSO puro e velocidade 90% superior ao ACO e Ant-Q em datasets TSPLIB (ch130, zi929). O artigo posiciona o DPSO-Q como ferramenta promissora para otimização logística em larga escala no contexto do e-commerce moderno.

## Contribuições Principais

- Integração de reinforcement learning (mecanismo Ant-Q) com DPSO para TSP
- Balanceamento adaptativo entre eficiência computacional e qualidade da solução via RL
- Validação em datasets benchmark TSPLIB demonstrando superioridade sobre DPSO puro, ACO e Ant-Q
- Redução de até 7.5% no tour length e 90% mais rápido que ACO/Ant-Q

## Relevância para o TCC

Abordagem state-of-the-art de hibridização PSO-RL para TSP, diretamente aplicável ao contexto de otimização de rotas de patrulha com drones. A técnica de usar RL para guiar a busca do PSO discreto pode inspirar extensões no código do repositório (ex.: PSO com Q-learning adaptativo). Demonstra que swarm intelligence combinada com aprendizado por reforço supera algoritmos clássicos como ACO.

## Métodos e Abordagens

- Discrete Particle Swarm Optimization (DPSO) com representação permutacional
- Reinforcement learning via Q-learning adaptado do algoritmo Ant-Q
- Datasets TSPLIB: ch130 (130 cidades), zi929 (929 cidades)
- Métricas: tour length, tempo de execução, desvio da solução ótima conhecida
- Implementação validada contra baselines: DPSO puro, ACO, Ant-Q

## Conexões

- [[kennedy1995particle]] — PSO original, base do DPSO
- [[dorigo1996ant]] — Ant System (inspiração para o Q-learning/Ant-Q)
- [[dorigo1997ant]] — Ant Colony System
- [[araujo2025pso]] — PSO discreto para TSP
- [[sun2024hybrid]] — PSO híbrido com 2-opt e Metropolis para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[particle-swarm]]
- [[ant-colony]]
- [[tsp]]
- [[bio-inspired-optimization]]

## Notas e Insights

- Open Access (Wiley/Hindawi) em: https://onlinelibrary.wiley.com/doi/10.1155/int/8918171 — PDF disponível na página do periódico
- Demonstra que RL pode melhorar significativamente a exploração do PSO discreto, evitando convergência prematura
- Tradeoff claro: DPSO-Q é mais lento que DPSO puro, mas compensa com melhor qualidade de solução
- Abordagem inspiradora para estender o PSO do repositório com heurísticas adaptativas
- O uso de Ant-Q como mecanismo RL é uma escolha interessante — conecta duas famílias de swarm intelligence
- Limitação: testado apenas em dois datasets; generalização para outras topologias não verificada
- Trabalho futuro promissor: aplicar DPSO-Q a variantes como rTSP e mTSP

## Citações-chave

> DPSO-Q reduces tour lengths by up to 7.5% compared to DPSO and achieves execution times over 90% faster than ACO and Ant-Q on standard datasets such as ch130 and zi929.

> By leveraging swarm intelligence and adaptive learning mechanisms, DPSO-Q achieves a balance between computational efficiency and high-quality solutions.
