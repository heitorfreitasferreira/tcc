---
title: Receding Horizon and Optimization-based Control for {UAV} Path Planning with Collision Avoidance
authors:
- Ahmed
- Gamil
- Sheltami
- Tarek
year: 2024
doi: 10.1016/j.procs.2024.11.079
bibtex_key: ahmed2024receding
bibtex-key: ahmed2024receding
pdf: papers/pdfs/ahmed2024receding.pdf
tags:
- area/drone-routing
- area/routing
- evidencia/referencia
- metodo/milp
- metodo/metaheuristic
- metodo/path-smoothing
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 3
classificacao: recuperado
type: paper
areas:
- drone-routing
- routing
methods:
- milp
- rhc
- path-smoothing
- cplex
role: revisao
reading_status: resumo-lido
validation_status: nao-validado
pdf_status: lido
---

## Resumo

Publicado em Procedia Computer Science 251 (2024) 15–22, Elsevier. Proposta de planejamento de trajetória *online* energeticamente eficiente para UAVs em ambientes complexos, com formulação de **Programação Linear Inteira Mista (MILP)** e técnica de **Receding Horizon Control (RHC)**. O caminho é dividido em segmentos (sub-paths) restritos ao campo de visão do drone, com janelas de horizonte finito atualizadas continuamente conforme o UAV avança; o solver **CPLEX** resolve o MILP de cada segmento para obter o caminho ótimo local em tempo real. Para reduzir o consumo de energia causado por curvas fechadas, é aplicada uma estratégia de *path smoothing* nos caminhos gerados. A função objetivo penaliza consumo energético e custo de colisão com obstáculos (`OC_ij`), com variáveis de decisão para a sequência de waypoints. A modelagem energética usada é a derivada de referência experimental do artigo [14] do próprio grupo.

**Contribuições declaradas (do PDF, Seção 1):**
1. Abordagem energeticamente eficiente com horizonte finito dinâmico, resolvendo o MILP localmente em janelas restritas que se atualizam com a posição do UAV.
2. Formulação MILP com restrições para minimização de energia e evasão de colisões, incluindo condição de retorno seguro para recarga quando a energia for insuficiente.
3. Estratégia de suavização de trajetória (*path smoothing*) para reduzir o efeito de curvas acentuadas, que consomem energia adicional nas fases de desaceleração e aceleração.

## Contribuições Principais

- Formulação MILP para o problema de planejamento de trajetória com função de custo energético e restrições de colisão
- Acoplamento de RHC com MILP resolvido em horizonte finito dinâmico (sub-regiões do espaço percorridas em sequência)
- Solução exata via CPLEX para cada sub-problema dentro do campo de visão do drone
- Estratégia de *path smoothing* pós-processamento para reduzir energia em curvas
- Simulação em MATLAB variando número de obstáculos (0–8) e medindo consumo normalizado

## Relevância para o TCC

O artigo referencia diretamente o problema de planejamento de rotas para drones com restrições de colisão, mas usa um arcabouço determinístico (MILP + RHC) e não meta-heurísticas populacionais — não trata de TSP, ACO, PSO ou GA. Serve como referência complementar para modelagem de restrições energéticas, e o método RHC aplicado a sub-paths é conceitualmente similar a esquemas de janela deslizante (rolling horizon) que podem ser combinados com meta-heurísticas para melhorar a escalabilidade.

## Métodos e Abordagens

- Formulação: MILP com variáveis de decisão inteiras para sequência de waypoints em `N_p` posições possíveis do segmento atual
- Função objetivo: minimizar `E_total` (consumo total de energia do segmento) com termo de penalidade por colisão
- Restrições principais:
  - (1.f) simetria do consumo de energia: `E_ij = E_ji`
  - (1.g) energia suficiente para o trecho
  - Restrições de avoidance com obstáculos
- Solver: CPLEX
- Loop RHC: a cada intervalo, monta MILP do segmento dentro do campo de visão, resolve, executa o trecho, move a janela
- Path smoothing: aplicado ao caminho gerado para suavizar curvas e reduzir consumo
- Simulação: MATLAB, com medições de consumo normalizado sobre o consumo ótimo (sem obstáculos) e variação de 0 a 8 obstáculos

## Conexões

- [[murray2015flying]] — FSTSP truck-drone, MILP para scheduling
- [[agatz2018optimization]] — TSP-D, formulações IP
- [[dellamico2021multiple]] — MFSTSP, scheduling MILP
- [[uav]] — categoria ampla

## Notas e Insights

- O método RHC com horizonte finito é o que torna viável resolver MILP *online* dentro do campo de visão do drone
- O ganho energético do *path smoothing* é documentado no artigo (Fig. 6b do PDF mostra percentual de melhoria vs. número de obstáculos)
- A função de custo combina energia de voo com penalidade por colisão `OC_ij`
- Limitação: dependência do solver exato (CPLEX) pode ser custosa para grandes instâncias em tempo real
- O modelo de energia é empírico (referência [14] do próprio grupo) e restrito a uma classe específica de drones

## Citações-chave

> "This paper proposes an online energy-efficient path planning approach for UAVs in complex environments."

> "The problem is formulated as a minimization optimization problem based on Mixed Integer Linear Programming (MILP), where a cost function is designed to minimize energy consumption while ensuring terrain obstacle avoidance within a limited detection range."

> "To achieve this, we apply a Receding Horizon Control (RHC) and optimization approach. The entire path is divided into segments or sub-paths, with constraints in place to prevent collisions with obstacles."
