---
title: "Particle Swarm Optimization: A Comprehensive Survey"
authors: [Shami, Tareq M., El-Saleh, Ayman A., Alswaitti, Mohammed, Al-Tashi, Qasem, Summakieh, Mhd Amen, Mirjalili, Seyedali]
year: 2022
doi: "10.1109/access.2022.3142859"
bibtex_key: shami2022pso
bibtex-key: shami2022pso
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

O PSO amadureceu ao longo de 27 anos em um ecossistema de variantes que atacam sistematicamente seus três problemas centrais — convergência prematura, dependência paramétrica e equilíbrio exploração-explotacão — por meio de quatro grandes famílias de aperfeiçoamento: modificação de parâmetros, hibridização com outros métodos, extensão multiobjetivo e adaptação a domínios discretos/binários. A survey fornece a taxonomia mais abrangente publicada até 2022 (324 referências), organizando as variantes por estratégia de melhoria e não apenas por aplicação.

## Resumo

Shami et al. (2022) publicam na IEEE Access a survey mais citada sobre PSO (1249 citações), cobrindo 324 trabalhos. O artigo organiza as variantes de PSO em quatro eixos: (1) modificação de parâmetros — inércia adaptativa, coeficientes de aceleração variáveis no tempo e topologias de vizinhança (global, anel, Von Neumann); (2) hibridização — PSO com algoritmos genéticos (GA-PSO), simulated annealing, busca local, lógica fuzzy e redes neurais; (3) PSO multiobjetivo (MOPSO) — dominância de Pareto, indicadores de desempenho e arquivos externos; e (4) PSO binário e discreto — funções de transferência (S-shaped, V-shaped) e operadores de mapeamento para espaços combinatórios. O texto dedica atenção especial ao problema da convergência prematura e cataloga soluções: perturbação estocástica (Turbulent PSO), reinicialização adaptativa, aprendizagem por oposição (OBL), e mecanismos de diversidade populacional. A survey também compila benchmarks padrão (funções CEC, problemas de engenharia) e oferece recomendações práticas para seleção de variantes conforme a classe de problema.

## Contribuições Principais

- Taxonomia em 4 eixos (parâmetros, hibridização, multiobjetivo, discreto/binário) cobrindo 324 artigos.
- Catálogo sistemático das causas e soluções para convergência prematura (Turbulent PSO, OBL, reinicialização).
- Análise detalhada de topologias de vizinhança e seu impacto no equilíbrio exploração-explotacão.
- Levantamento exaustivo de PSO binário (BPSO) e suas funções de transferência (S-shape, V-shape), relevante para adaptação a TSP.
- Recomendações práticas de qual variante usar por classe de problema (unimodal, multimodal, restrito, multiobjetivo).
- Compilação de benchmarks e métricas padronizadas para comparação justa entre variantes.

## Relevância para o TCC

Esta é a survey de PSO mais citada da literatura e cobre exatamente os aspectos que fundamentam a implementação do [[particle-swarm]] no TCC: o mecanismo canônico de Kennedy-Eberhart, as estratégias de adaptação de parâmetros (inércia, constantes de aceleração) e, crucialmente, as abordagens de PSO discreto (BPSO com funções de transferência) que são a base para aplicar PSO ao [[tsp]]. A discussão sobre convergência prematura é diretamente relevante para interpretar os resultados experimentais — se o PSO estagnar em ótimos locais nas instâncias TSP do projeto, a survey fornece o referencial teórico para explicar o fenômeno e apontar melhorias. A seção de hibridização (GA-PSO) também contextualiza comparações entre métodos populacionais.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Fundamentação do funcionamento do PSO, justificativa da escolha de parâmetros, explicação de convergência prematura.
- Como citar na monografia: \cite{shami2022pso} — para definição canônica de PSO, taxonomia de variantes, e discussão de convergência prematura.

## Métodos e Abordagens

- PSO canônico: atualização de velocidade v_i(t+1) = w·v_i(t) + c₁r₁(pbest_i − x_i) + c₂r₂(gbest − x_i).
- Modificação de parâmetros: inércia linear decrescente (Shi & Eberhart 1998), inércia adaptativa (Zhan et al. 2009), coeficientes de aceleração variáveis (Ratnaweera et al. 2004).
- Topologias de vizinhança: global (gbest PSO), anel (lbest PSO com k=2 vizinhos), Von Neumann, dinâmica hierárquica.
- Hibridização: GA-PSO (crossover + mutação integrados), PSO-SA (simulated annealing para escape de ótimos locais), PSO com busca local (Lamarckiana).
- MOPSO: dominância de Pareto, crowding distance, arquivo externo, líderes não-dominados.
- PSO binário (BPSO): função sigmoide S-shaped (Kennedy & Eberhart 1997), V-shaped (Rashedi et al. 2010), mapeamento de velocidade → probabilidade de bit.
- PSO discreto: operadores de swap para permutações, representação por chaves aleatórias, estratégias de reparo de factibilidade.
- Benchmarks: funções unimodais (Sphere, Rosenbrock), multimodais (Rastrigin, Griewank, Ackley), CEC 2005/2013/2017, problemas de engenharia (tension/compression spring, welded beam).

## Evidência / Resultado Relevante

- A survey é a referência canônica para PSO com 1249 citações (Google Scholar), indicando consenso da comunidade sobre sua taxonomia.
- Cobertura de 324 artigos fornece base estatística sólida para generalizações sobre eficácia relativa das variantes.
- Demonstra que PSO com inércia adaptativa e topologia em anel consistentemente supera PSO canônico em funções multimodais — insight relevante para TSP (problema multimodal).

## Limitações de Uso

> [!warning] Limitação
> A survey cobre PSO em domínio contínuo predominantemente; a adaptação a TSP requer PSO discreto (BPSO ou operadores de swap), que é tratado em uma subseção apenas. O artigo não avalia variantes de PSO em instâncias TSPLIB — a evidência empírica é majoritariamente sobre funções benchmark contínuas. Para o TCC, esta referência fundamenta a teoria do PSO, mas não fornece resultados experimentais diretamente transferíveis ao TSP-SD-ATP.

## Conexões

- Fundamenta: [[kennedy1995particle]] — PSO canônico que a survey referencia como ponto de partida.
- Relacionado a: [[gad2022pso]], [[zhu2025cumulative]], [[zhang2015comprehensive]] — outras surveys de PSO com escopos complementares.
- Relacionado a: [[clerc2000discretepso]] — PSO discreto, abordagem necessária para aplicação ao TSP.
- Contrasta com: [[pop2024comprehensive]] — survey de ACO vs. survey de PSO; métodos bio-inspirados distintos.
- Apoia claim: Eficácia de métodos populacionais bio-inspirados para otimização combinatória.
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A survey organiza PSO por *mecanismo de melhoria*, não por aplicação — isso facilita selecionar a variante certa para TSP com base nas características do problema (discreto, NP-difícil, multimodal).
- A seção de BPSO com funções de transferência S-shape e V-shape é essencial para entender como mapear posições contínuas de partículas para arestas de um tour TSP.
- O artigo é particularmente útil para justificar a escolha de parâmetros do PSO no TCC (tamanho do enxame, coeficientes de aceleração, estratégia de inércia).
- As topologias de vizinhança (global vs. anel) têm impacto documentado em diversidade populacional — o TCC pode explorar essa variante se o PSO canônico estagnar.
- Referência às CEC competitions fornece benchmarks padronizados que podem ser adaptados para comparações futuras.

## Citações-chave

> "PSO has gained increasing popularity among researchers due to its simplicity, fast convergence, and few control parameters."

> "Premature convergence is one of the main drawbacks of PSO that occurs when particles lose diversity and converge to a local optimum."

> "Binary PSO employs a transfer function to map the continuous search space to a binary search space, which is essential for solving combinatorial optimization problems such as the TSP."
