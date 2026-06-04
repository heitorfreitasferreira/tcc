---
title: "DeepACO: Neural-enhanced Ant Systems for Combinatorial Optimization"
authors: [Ye, Haoran, Wang, Jiarui, Cao, Zhiguang, Liang, Helan, Li, Yong]
year: 2023
doi: "10.48550/arXiv.2305.19416"
bibtex_key: deepaco2023
bibtex-key: deepaco2023
pdf: "papers/pdfs/deepaco2023.pdf"
tags: [aco, deep-learning, neural-combinatorial-optimization, neurips]
status: resumo-lido
rating: 4
---

## Resumo

Publicado no NeurIPS 2023. Propõe o DeepACO, um framework genérico que utiliza deep reinforcement learning para automatizar o design de medidas heurísticas em ACO. O DeepACO fortalece as heurísticas de algoritmos ACO existentes, dispensando o design manual trabalhoso. Como meta-heurística neuralmente aprimorada, o DeepACO supera consistentemente seus equivalentes ACO padrão em oito COPs usando um único modelo neural e um único conjunto de hiperparâmetros. Como método de Neural Combinatorial Optimization, o DeepACO tem desempenho comparável ou superior a métodos específicos de problema em problemas canônicos de roteamento.

## Contribuições Principais

- Uso de deep RL para aprender heurísticas de ACO automaticamente, substituindo design manual
- Único modelo neural + único conjunto de hiperparâmetros funcionam em 8 problemas diferentes
- Supera ACO padrão e compete com métodos NCO específicos de problema
- Código aberto disponível

## Relevância para o TCC

Citado na Conclusão (Seção 5) como direção de trabalho futuro — investigar métodos híbridos e neurais como DeepACO. Não é usado na implementação atual (que é Ant System puro). Fornece contexto sobre a fronteira de pesquisa em ACO.

## Conexões

- [[ant-colony]]
- [[neufaco2025]] — NeuFACO, estado-da-arte neural ACO
- [[gpaco2025]] — GP-ACO, projeto automático de regras de transição
- [[dorigo1996ant]] — Ant System original (base do ACO implementado)
