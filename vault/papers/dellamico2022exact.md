---
title: "Exact Models for the Flying Sidekick Traveling Salesman Problem"
authors: [Dell'Amico, Mauro]
year: 2022
doi: "10.1111/itor.13030"
bibtex-key: dellamico2022exact
tags: [tsp drone]
status: lido
rating: 4
pdf: "papers/pdfs/dellamico2022exact.pdf"
---

## PDF

![[dellamico2022exact.pdf]]

## Resumo

Propõe modelos exatos (MILP) melhorados para o FSTSP (Flying Sidekick Traveling Salesman Problem), superando as limitações da formulação original de [[murray2015flying]]. Desenvolve três formulações alternativas: (i) formulação de fluxo em três índices baseada em sorties do drone, (ii) formulação de particionamento de rota com desigualdades válidas, e (iii) formulação com variáveis de tempo contínuo. As novas formulações permitem resolver instâncias de até 20 clientes de forma ótima, enquanto a formulação original resolvia apenas 10 clientes com gap. Inclui refinamentos de pré-processamento e planos de corte para fortalecimento dos limites inferiores.

## Contribuições Principais

- Três formulações MILP alternativas para o FSTSP com melhor desempenho computacional
- Desigualdades válidas e planos de corte para fortalecimento dos limites inferiores
- Capacidade de resolver instâncias de até 20 clientes de forma ótima (vs. 10 clientes na formulação original)
- Análise comparativa detalhada entre as formulações propostas
- Pré-processamento para redução do espaço de busca de sorties do drone

## Relevância para o TCC

Fornece a base exata para comparação de qualidade das soluções heurísticas/metaheurísticas implementadas no TCC. Embora o TCC use rTSP (não FSTSP), a metodologia de validação de gap (solução heurística vs. ótimo exato) é diretamente aplicável. As técnicas de pré-processamento e desigualdades válidas podem inspirar refinamentos nas implementações GA/PSO/ACO.

## Métodos e Abordagens

- Programação Inteira Mista (MILP) com CPLEX
- Três formulações: (a) fluxo em 3 índices, (b) particionamento de rota, (c) tempo contínuo
- Desigualdades válidas e planos de corte (cutting planes)
- Pré-processamento de sorties elegíveis do drone
- Instâncias de 10 a 20 clientes com variação de velocidade/autonomia do drone

## Conexões

- [[murray2015flying]] — define o FSTSP, problema modelado exatamente aqui
- [[dellamico2021multiple]] — formulações para múltiplos drones (mesmos autores)
- [[agatz2018optimization]] — heurísticas para TSP-D
- [[freitas2020vns]] — VNS para FSTSP (comparação de gap)
- [[applegate2006traveling]] — métodos exatos para TSP (inspiração)
- [[drone-routing]]

## Notas e Insights

- Instâncias com 20 clientes requerem até 2h para solução ótima mesmo com as formulações melhoradas
- A formulação de fluxo em três índices é a mais eficiente computacionalmente
- O gap das heurísticas existentes (e.g., [[freitas2020vns]]) pode ser avaliado contra estes ótimos
- Para o TCC, métodos exatos são inviáveis para instâncias grandes, mas servem como ground truth para instâncias pequenas

## Citações-chave

> "The proposed models allow to solve to optimality FSTSP instances with up to 20 customers, whereas the original formulation could only handle instances with up to 10 customers."
