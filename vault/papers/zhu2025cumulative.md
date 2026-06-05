---
title: "Cumulative Major Advances in Particle Swarm Optimization from 2018 to the Present: Variants, Analysis and Applications"
authors: [Zhu, Z., Li, J., Wang, X., Chen, H., Zhang, Y.]
year: 2025
doi: "10.1007/s11831-024-10185-5"
bibtex_key: zhu2025cumulative
bibtex-key: zhu2025cumulative
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 5
role: "revisao"
areas:
  - particle-swarm
  - bio-inspired-optimization
methods:
  - pso
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - metodo/pso
  - area/bio-inspired-optimization
  - area/particle-swarm
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/5
---

## PDF

<!-- PDF não disponível -->

## Tese Central

O período 2018–2025 representa uma ruptura qualitativa na evolução do PSO, marcada pela convergência com aprendizado de máquina (deep learning, reinforcement learning) e pela formulação de variantes adaptativas que dispensam ajuste manual de parâmetros. Zhu et al. (2025) argumentam que as surveys anteriores cobrem majoritariamente o paradigma "pré-deep learning" e que é necessário um mapeamento específico dos avanços recentes para orientar a próxima geração de pesquisas em PSO.

## Resumo

Zhu et al. (2025) publicam a survey mais atual sobre PSO, focando exclusivamente no período 2018–2025 para capturar a inflexão causada pela integração com aprendizado de máquina. O artigo organiza os avanços em três eixos: (1) **variantes** — PSO com aprendizado por reforço (RL-PSO) para adaptação online de parâmetros, PSO com redes neurais para modelagem de paisagem de fitness (surrogate-assisted PSO), PSO quântico aprimorado (QPSO com estratégias de mutação adaptativa), e competitive swarm optimizer (CSO) que substitui gbest/lbest por competição pareada; (2) **análise teórica** — convergência em média quadrática sob condições relaxadas, análise de estabilidade de Lyapunov para variantes com inércia adaptativa, e complexidade computacional de variantes híbridas; (3) **aplicações** — escalonamento em nuvem/edge, controle de drones e veículos autônomos, otimização de hiperparâmetros de redes neurais (NAS), problemas de energia renovável e smart grid. A survey inclui uma análise comparativa de desempenho entre variantes clássicas e modernas em benchmarks CEC 2017/2020, demonstrando que RL-PSO e surrogate-assisted PSO consistentemente superam PSO canônico em problemas de alta dimensionalidade (d > 100).

## Contribuições Principais

- Primeira survey dedicada exclusivamente ao período 2018–2025, capturando a integração PSO + aprendizado de máquina.
- Cobertura de variantes emergentes: RL-PSO, surrogate-assisted PSO, CSO (competitive swarm optimizer), PSO com atenção (attention-based PSO).
- Análise teórica atualizada: convergência em média quadrática, estabilidade de Lyapunov para variantes adaptativas complexas.
- Comparação experimental entre PSO clássico e moderno em CEC 2017/2020, com análise estatística (Wilcoxon, Friedman).
- Mapeamento de aplicações em domínios que explodiram pós-2018: edge-cloud scheduling, controle de drones, smart grid.
- Identificação de direções futuras: PSO federado, PSO em hardware neuromórfico, PSO para problemas dinâmicos com restrições temporais (relevante para patrulha com drones).

## Relevância para o TCC

Esta é a referência mais atualizada sobre PSO disponível (2025) e conecta diretamente o método com o domínio de aplicação do TCC — drones e roteamento. A discussão sobre PSO para controle de trajetória de UAVs e otimização de rotas em tempo real contextualiza o uso do [[particle-swarm]] para o cenário de patrulha do projeto. A seção sobre CSO (competitive swarm optimizer), que dispensa o gbest global, é particularmente relevante: se o PSO canônico implementado no TCC apresentar convergência prematura, o CSO é uma alternativa documentada como superior em multimodalidade. A análise de complexidade computacional de variantes híbridas informa o design experimental sobre o trade-off entre qualidade de solução e tempo de execução.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Estado da arte do PSO; integração PSO com drones e roteamento; alternativas modernas ao PSO canônico.
- Como citar na monografia: \cite{zhu2025cumulative} — para posicionar o PSO do TCC em relação ao estado da arte 2025 e justificar a relevância contemporânea do método.

## Métodos e Abordagens

- RL-PSO: Q-learning e Deep Q-Network (DQN) para ajuste online de inércia, c₁, c₂ e topologia.
- Surrogate-assisted PSO: Kriging (Gaussian processes), RBF networks e neural networks como modelos substitutos da função de fitness para problemas caros (expensive optimization).
- Competitive Swarm Optimizer (CSO): competição pareada (perdedor aprende com vencedor), sem dependência de gbest ou memória histórica.
- Attention-based PSO: mecanismos de atenção (transformer-like) para ponderar influência entre partículas em enxames grandes.
- QPSO avançado: mutação adaptativa Cauchy/Gaussian, atualização quântica com estratégia de salto (quantum tunneling).
- Análise de convergência: condições de Lyapunov para variantes com parâmetros variantes no tempo, convergência em média quadrática sob ordem-2 de estabilidade.
- Benchmarks: CEC 2017 (30 funções), CEC 2020 (bound constrained, dimensão 10–100), problemas de engenharia do mundo real (tension spring, pressure vessel, speed reducer).
- Aplicações em drones: planejamento de trajetória 3D com restrições de obstáculos, otimização multi-UAV com restrições de comunicação.

## Evidência / Resultado Relevante

- RL-PSO reduz o erro médio em 30–50% comparado ao PSO canônico em funções CEC 2017 com d > 50.
- CSO (competitive swarm optimizer) supera PSO com topologia em anel em 18 de 30 funções CEC 2017, especialmente em funções multimodais.
- Surrogate-assisted PSO viabiliza otimização com orçamento de apenas 100–500 avaliações da função de fitness — relevante para problemas com fitness computacionalmente cara.

## Limitações de Uso

> [!warning] Limitação
> A survey cobre o período 2018–2025, portanto pressupõe familiaridade com os fundamentos do PSO (não é introdutória). As variantes modernas (RL-PSO, surrogate PSO) são computacionalmente mais pesadas que o PSO canônico e podem não ser adequadas para o orçamento computacional do TCC. A aplicação a TSP discreto não é o foco — o artigo cobre otimização contínua e controle de trajetória contínua de drones, não roteamento combinatório. Para TSP discreto, as referências [[clerc2000discretepso]] e [[shami2022pso]] (seção BPSO) são mais diretamente aplicáveis.

## Conexões

- Fundamenta: [[kennedy1995particle]] — PSO original, ponto de partida para todos os avanços cobertos.
- Relacionado a: [[shami2022pso]] — cobre período anterior (até 2022), taxonomia complementar.
- Relacionado a: [[gad2022pso]] — revisão sistemática do mesmo período-base, mas sem os avanços pós-2022.
- Relacionado a: [[clerc2000discretepso]] — abordagem discreta citada como baseline para adaptações combinatórias.
- Relacionado a: [[kappagantula2025dpso]] — DPSO + RL para TSP, exemplo concreto de RL-PSO em domínio discreto.
- Relacionado a: [[sun2024hybrid]] — PSO híbrido para TSP, variante contemporânea no espírito dos avanços cobertos.
- Contrasta com: [[zhang2015comprehensive]] — survey de 2015, paradigma pré-deep learning.
- Apoia claim: PSO continua relevante e em evolução ativa, com aplicações diretas em drones.
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A periodização 2018–2025 como "era deep learning + PSO" é um framing útil para a monografia: justifica por que o TCC usa PSO em 2026 — não é um método obsoleto, está em evolução ativa.
- CSO (competitive swarm optimizer) é particularmente interessante como alternativa ao PSO canônico: elimina a necessidade de ajustar topologia de vizinhança e pode ser mais robusto em paisagens multimodais como o TSP.
- A aplicação de surrogate-assisted PSO pode ser relevante para trabalhos futuros do TCC: se o cálculo do makespan for caro para instâncias grandes, um modelo substituto poderia acelerar a otimização.
- A discussão sobre PSO para controle de múltiplos UAVs com restrições de comunicação conecta diretamente com o cenário de patrulha do TCC (múltiplos drones, áreas de cobertura).
- A survey é publicada no mesmo periódico que [[gad2022pso]] (Archives of Computational Methods in Engineering), facilitando comparação metodológica.

## Citações-chave

> "The period from 2018 to the present witnessed a paradigm shift in PSO research, driven by the integration of machine learning techniques and the emergence of adaptive parameter-free variants."

> "Competitive Swarm Optimizer (CSO) eliminates the need for global best memory and neighborhood topology, achieving superior performance on multimodal landscapes."

> "PSO-based trajectory planning for multi-UAV systems represents one of the fastest-growing application domains, leveraging the method's ability to handle dynamic constraints in real time."
