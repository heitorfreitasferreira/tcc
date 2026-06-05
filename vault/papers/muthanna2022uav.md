---
title: Towards Optimal Positioning and Energy-Efficient UAV Path Scheduling in IoT Applications
authors:
- Muthanna
- Mohammed Saleh Ali
- Muthanna
- Ammar
- Nguyen
- Tu N.
- Alshahrani
- Abdullah
- Abd El-Latif
- Ahmed A.
year: 2022
doi: 10.1016/j.comcom.2022.04.028
bibtex_key: muthanna2022uav
bibtex-key: muthanna2022uav
pdf: papers/pdfs/muthanna2022uav.pdf
tags:
- area/drone-routing
- area/iot
- evidencia/referencia
- metodo/c-lstm
- metodo/a3c
- metodo/mayfly-optimization
- status/lido
- tipo/paper
status: lido
rating: 3
classificacao: recuperado
type: paper
areas:
- drone-routing
- iot
- emergency-communications
methods:
- c-lstm
- a3c
- mayfly-optimization
- iot
role: revisao
reading_status: lido
validation_status: nao-validado
pdf_status: lido
---

## Resumo

Publicado em Computer Communications 191 (2022) 145–160, Elsevier. Propõe o arcabouço **IWPOP-UAV** (*Impact of Weather-based Positioning and Path planning of UAVs*) para comunicações de UAV em IoT em situações de emergência (desastre, busca e resgate). O sistema tem três módulos encadeados:

1. **Predição climática com C-LSTM** (*Cerebral Long Short-Term Memory*): variante de LSTM com célula recorrente inovadora, com *forget gate* e *input gate* otimizados, que afirma ter menor *training loss* que LSTM convencional. A C-LSTM recebe parâmetros climáticos históricos (precipitação, vento, pressão, temperatura, etc.) e prediz o comportamento do tempo nas próximas 24 h.
2. **Posicionamento multi-UAV com A3C** (*Asynchronous Actor Critic*): o A3C é um algoritmo de *deep reinforcement learning* que, usando a previsão climática da C-LSTM, decide a posição ótima de cada UAV para maximizar cobertura e QoS.
3. **Path planning com Mayfly Optimization Algorithm (MOA)**: meta-heurística bioinspirada híbrida que combina GA + PSO + Firefly, resolvendo o planejamento de trajetória energeticamente eficiente para cada UAV já posicionado.

A avaliação experimental é feita em NS-3.26, comparando IWPOP-UAV com abordagens de referência em métricas de QoS, cobertura, latência, energia e PDR (*Packet Delivery Ratio*).

**Atenção**: este artigo **NÃO trata de TSP nem de meta-heurísticas clássicas de otimização combinatória** (ACO, GA, PSO puro). É de comunicações sem fio em IoT/5G com UAVs, e o "PSO" mencionado na literatura relacionada não é o algoritmo de Kennedy/Eberhart, mas sim um ingrediente da Mayfly. Para o TCC, a relevância é apenas contextual: modelagem de restrições energéticas em missões com UAV.

## Contribuições Principais

- C-LSTM: célula recorrente melhorada para predição de séries temporais climáticas
- A3C para posicionamento multi-UAV em tempo real, considerando condições meteorológicas
- Mayfly Optimization Algorithm (MOA) para path planning energeticamente eficiente
- Integração end-to-end clima→posicionamento→trajetória
- Avaliação em NS-3.26 com métricas de rede (PDR, delay, throughput) e de cobertura

## Relevância para o TCC

Relevância **baixa** para o problema central (TSP/rTSP patrulha de drones). O artigo modela restrições energéticas e climáticas, mas usa arcabouço de deep RL + meta-heurística híbrida (Mayfly) em vez de ACO/PSO/GA canônicos. Pode ser citado como exemplo de aplicação de UAV em IoT/5G com restrições meteorológicas, mas **não** como referência de otimização combinatória bioinspirada. Não trata de TSP nem do cenário de patrulha com múltiplos pontos de interesse.

## Métodos e Abordagens

- **C-LSTM**: rede recorrente com *forget gate* `f(t)`, *input gate* e célula de memória; predição de 24 h de variáveis climáticas (temperatura, vento, chuva, pressão, umidade)
- **A3C**: actor-critic assíncrono com múltiplos workers; estado = condições meteorológicas atuais + posições dos UAVs; ação = ajuste de posição (x, y, altitude)
- **MOA (Mayfly Optimization)**: híbrido que combina:
  - Movimento de macho inspirado em PSO (componente `v_ij`)
  - Movimento de fêmea inspirado em Firefly (atração pelo melhor)
  - Operadores de crossover/mutação de GA para acasalamento
- Simulação em NS-3.26 com modelo de canal sem fio para IoT/5G

## Conexões

- [[murray2015flying]] — FSTSP truck-drone (problema diferente)
- [[agatz2018optimization]] — TSP-D (problema diferente)
- [[uav]] — categoria ampla
- Deep RL — A3C é método de deep RL
- IoT/5G — comunicações

## Notas e Insights

- O claim "C-LSTM tem menor training loss que LSTM" não é comparado quantitativamente com LSTM padrão no abstract — verificação necessária no corpo do artigo
- A Mayfly Optimization é uma meta-heurística obscura sem ampla adoção na literatura de TSP
- A integração clima→posicionamento→trajetória é o diferencial conceitual
- Limitação: a avaliação é simulada (NS-3), sem experimentos de campo
- Tabela 7 do PDF compara IWPOP-UAV com baselines (GA, PSO, Firefly) e reporta superioridade em QoS

## Citações-chave

> "The Unmanned Aerial Vehicle (UAV) based communication has been emerged as a feasible solution for remote applications such as disaster management, and search and rescue due to its mobility and cost efficiency."

> "In this paper, the impact of weather-based positioning and path planning of UAVs (IWPOP-UAV) is carried out to achieve increased QoS, reliability and energy efficiency in UAV communications. Initially, the prediction of weather conditions is performed by utilizing the Cerebral Long Short-Term Memory (C-LSTM)..."

> "Each cell is provided by the A3C algorithm based on weather conditions and other significant factors thereby determining the optimal positioning of the UAV. The path planning problem is formulated as an optimization problem and executed by using Mayfly Optimization Algorithm (MOA)."
