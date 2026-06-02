---
title: "Receding Horizon and Optimization-based Control for {UAV} Path Planning with Collision Avoidance"
authors: [Ahmed, Gamil]
year: 2024
doi: "10.1016/j.procs.2024.11.079"
bibtex-key: ahmed2024receding
tags: [drone metaheuristic routing]
status: lido
rating: 3
pdf: null
---

## Resumo

Propõe uma abordagem de controle preditivo baseado em horizonte recedente (Receding Horizon Control — RHC) integrado com Particle Swarm Optimization (PSO) para planejamento de caminho de UAVs com desvio de obstáculos. O RHC otimiza a trajetória local em uma janela deslizante, enquanto o PSO ajusta os parâmetros de controle em cada iteração. Simulações em ambientes 2D e 3D demonstram que o método RHC-PSO gera trajetórias suaves e sem colisões, superando abordagens RHC isoladas em termos de tempo de computação e qualidade da trajetória. O artigo foca em aspectos de navegação em tempo real, diferentemente dos problemas de roteirização combinatória como TSP.

## Contribuições Principais

- Integração de Receding Horizon Control com PSO para planejamento local de trajetória UAV
- Função objetivo multi-critério: distância percorrida, suavidade da trajetória, distância de obstáculos
- Simulações 2D e 3D demonstrando eficácia em desvio de obstáculos estáticos
- Comparação quantitativa entre RHC-PSO e RHC puro

## Relevância para o TCC

Oferece perspectiva complementar: planejamento de caminho local com desvio de obstáculos, enquanto o TCC foca em roteirização global (TSP/rTSP). A integração PSO + controle preditivo ilustra como metaheurísticas podem ser usadas em tempo real, relevante para cenários híbridos de patrulha com drones que exigem replanejamento dinâmico. Contudo, o problema tratado (path planning com obstáculos) é distinto do problema central do TCC (roteirização combinatória).

## Métodos e Abordagens

- Receding Horizon Control (RHC) / Model Predictive Control (MPC)
- Particle Swarm Optimization (PSO) para otimização local em cada horizonte
- Função objetivo: combinação linear de distância, suavidade (curvatura) e distância mínima de obstáculos
- Simulações em ambientes 2D e 3D com obstáculos estáticos
- Malha de pontos de controle (B-splines) para representação da trajetória

## Conexões

- [[rajan2022routing]] — roteamento estocástico para patrulha de UAVs (perspectiva global)
- [[murray2015flying]] — FSTSP (otimização de rota com drone)
- [[kennedy1995particle]] — PSO original
- [[particle-swarm]]
- [[drone-routing]]

## Notas e Insights

- Open Access (CC BY) na ScienceDirect: https://doi.org/10.1016/j.procs.2024.11.079 — PDF disponível na página do DOI
- O foco é planejamento local (tempo real), não roteirização global (TSP)
- O PSO resolve o subproblema de otimização dos parâmetros de controle a cada janela do RHC
- Abordagem computacionalmente leve: adequada para execução embarcada em UAVs
- Não aborda problemas de roteirização (sequenciamento de múltiplos pontos de interesse)
- Relevância limitada para o núcleo combinatório do TCC, mas útil como referência para extensões de replanejamento dinâmico

## Citações-chave

> "The proposed RHC-PSO approach generates smooth and collision-free trajectories in both 2D and 3D environments, outperforming the standard RHC approach in terms of computational time."
