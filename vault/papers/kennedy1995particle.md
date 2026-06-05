---
title: Particle Swarm Optimization
authors:
- Kennedy
- James
year: 1995
doi: 10.1109/ICNN.1995.488968
bibtex_key: kennedy1995particle
bibtex-key: kennedy1995particle
tags:
- evidencia/referencia
- metodo/metaheuristic
- metodo/pso
- status/lido
- tipo/paper
status: lido
rating: 5
pdf: papers/pdfs/kennedy1995particle.pdf
type: paper
methods:
- metaheuristic
- pso
role: revisao
reading_status: lido
validation_status: nao-validado
---

## PDF

![[kennedy1995particle.pdf]]

## Resumo

O artigo fundacional do Particle Swarm Optimization (PSO) introduz um algoritmo de otimização inspirado no comportamento social de bandos de pássaros e cardumes de peixes. Cada partícula ajusta sua trajetória no espaço de busca combinando sua memória individual (pbest — "nostalgia simples") com o conhecimento coletivo do enxame (gbest — "norma social"). O algoritmo é extremamente simples, requerendo apenas operadores matemáticos primitivos, e demonstrou eficácia na otimização de funções não-lineares contínuas e no treinamento de redes neurais. Os autores também discutem as conexões entre PSO, vida artificial e algoritmos evolucionários.

## Contribuições Principais

- Proposição do PSO como método de otimização para funções contínuas não-lineares
- Demonstração de que comportamento social simulado (troca de informação entre agentes) pode resolver problemas de otimização
- Algoritmo extremamente simples (poucas linhas de código), computacionalmente barato em memória e velocidade
- Conexão conceitual entre evolução, swarm intelligence e aprendizado por reforço
- Validação em funções benchmark (Schaffer f6) e treinamento de redes neurais (XOR, Iris Data)

## Relevância para o TCC

O PSO é um dos três métodos metaheurísticos implementados no TCC para comparação sobre TSP/rTSP. Sua simplicidade e baixo custo computacional o tornam candidato natural para otimização de rotas de drones de patrulha. O artigo original estabelece os princípios que serão adaptados para o domínio discreto (TSP) nas variantes posteriores arquivadas na base. A metáfora social — agentes compartilhando informação sobre boas rotas — é diretamente análoga ao problema de coordenação entre múltiplos drones.

## Métodos e Abordagens

- Agentes (partículas) com posição e velocidade em espaço n-dimensional
- Duas forças de atração: pbest (melhor posição individual) e gbest (melhor posição global)
- Atualização de velocidade com componentes estocásticas (rand() × 2) para equilibrar exploração/explotação
- Testes em função Schaffer f6, treinamento de MLP para XOR, classificação Iris Data
- Comparação qualitativa com algoritmos genéticos e programação evolucionária

## Conexões

- [[araujo2025pso]] — PSO discreto para TSP
- [[huang2025matrix]] — PSO matricial para mTSP
- [[sun2024hybrid]] — PSO híbrido para TSP
- [[kappagantula2025dpso]] — DPSO com RL para TSP
- [[dorigo1996ant]] — ACO (outra abordagem bio-inspirada contemporânea)
- [[particle-swarm]]
- [[rajwar2023exhaustive]] — survey que classifica PSO como método canônico

## Notas e Insights

- O artigo evolui de uma simulação social para um otimizador funcional, ilustrando como metáforas biológicas podem gerar algoritmos práticos
- A versão "simplificada" final removeu parâmetros como craziness e nearest-neighbor velocity matching, usando apenas pbest/gbest com fator estocástico 2 — exemplo de elegância por poda
- Os autores (um psicólogo social e um engenheiro elétrico) vêm de áreas distintas, o que pode explicar a interdisciplinaridade da abordagem
- O algoritmo original é contínuo; para aplicação em TSP (domínio discreto) são necessárias adaptações como codificação por permutação ou operadores de posição discretos
- Limitação: não há garantia de convergência para o ótimo global; o overshooting (momentum) é feature, não bug

## Citações-chave

> "Particle swarm optimization is an extremely simple algorithm that seems to be effective for optimizing a wide range of functions."

> "Why is social behavior so ubiquitous in the animal kingdom? Because it optimizes. What is a good way to solve engineering optimization problems? Modeling social behavior."
