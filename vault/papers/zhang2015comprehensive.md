---
title: "A Comprehensive Survey on Particle Swarm Optimization Algorithm and Its Applications"
authors: [Zhang, Yudong, Wang, Shuihua, Ji, Genlin]
year: 2015
doi: "10.1155/2015/931256"
bibtex_key: zhang2015comprehensive
bibtex-key: zhang2015comprehensive
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
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
  - relevancia/4
---

## PDF

<!-- PDF não disponível -->

## Tese Central

O PSO transcendeu sua formulação original como otimizador contínuo para se tornar um framework generalista de otimização com aplicações em domínios radicalmente diversos — processamento de sinais, bioinformática, operações de manufatura e otimização combinatória —, e a chave para essa versatilidade está nas estratégias de adaptação do espaço de busca (discretização, hibridização e representação específica do domínio) mais do que em modificações no núcleo do algoritmo. Zhang, Wang e Ji (2015) oferecem a primeira survey verdadeiramente interdisciplinar do PSO, organizada por domínio de aplicação e não por mecanismo algorítmico.

## Resumo

Zhang, Wang e Ji publicam em Mathematical Problems in Engineering (Hindawi, 2015) uma das surveys mais citadas de PSO, distinguindo-se das contemporâneas pelo foco em aplicações interdisciplinares em vez de taxonomia algorítmica. O artigo cobre: (1) **fundamentos do PSO** — equações canônicas, inércia (Shi & Eberhart), fator de constrição (Clerc & Kennedy) e topologias de vizinhança; (2) **variantes paramétricas** — PSO com parâmetros adaptativos (APSO), PSO com aprendizagem abrangente (CLPSO), PSO com distância (DPSO); (3) **PSO para otimização multiobjetivo** — MOPSO com dominância de Pareto, NSGA-II + PSO; (4) **aplicações em processamento de sinais** — design de filtros digitais (FIR/IIR), beamforming, estimação de direção de chegada (DOA); (5) **aplicações em pesquisa operacional** — escalonamento de job-shop, roteamento de veículos (VRP), problema do caixeiro viajante (TSP); (6) **aplicações em bioinformática** — alinhamento de sequências, predição de estrutura de proteínas, seleção de genes; e (7) **aplicações em manufatura** — otimização de parâmetros de usinagem, cadeia de suprimentos. O artigo conclui que a eficácia do PSO em cada domínio depende criticamente da estratégia de representação (codificação da solução no espaço da partícula), mais do que da variante algorítmica escolhida.

## Contribuições Principais

- Primeira survey de PSO organizada por domínio de aplicação (interdisciplinar), complementando surveys organizadas por mecanismo.
- Cobertura de aplicações em processamento de sinais e bioinformática — domínios raramente cobertos em surveys de otimização.
- Discussão detalhada sobre estratégias de representação (encoding) como fator crítico de sucesso do PSO em cada domínio.
- Levantamento de PSO aplicado a TSP e VRP — diretamente relevante para o TCC.
- Identificação de que PSO para problemas combinatórios requer adaptação de representação (ordinal, path, adjacency) e não apenas ajuste de parâmetros.
- Catálogo de funções benchmark usadas na literatura até 2015, com análise de adequação por classe de problema.

## Relevância para o TCC

Zhang et al. (2015) é a survey de PSO que mais diretamente aborda TSP e roteamento entre as quatro surveys enriquecidas. A seção sobre PSO para TSP cobre a representação de tours como permutações, o uso de operadores de troca (swap operators) e a hibridização com busca local (2-opt, 3-opt) — exatamente os mecanismos implementados no [[particle-swarm]] do TCC. A discussão sobre VRP estende a relevância para o cenário de patrulha multi-drone (que pode ser modelado como variante de VRP). Além disso, o artigo fornece uma referência histórica importante: mostra o estado do PSO aplicado a problemas combinatórios antes da era deep learning, servindo como baseline para comparar com os avanços documentados em [[zhu2025cumulative]].

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Fundamentação de PSO para TSP; estratégias de representação discreta; adaptação de operadores de swap para roteamento.
- Como citar na monografia: \cite{zhang2015comprehensive} — para a conexão histórica entre PSO e TSP/VRP e para justificar a estratégia de representação de tours adotada no TCC.

## Métodos e Abordagens

- PSO canônico: v(t+1) = w·v(t) + c₁r₁(pbest − x) + c₂r₂(gbest − x); x(t+1) = x(t) + v(t+1).
- Variantes cobertas: PSO com inércia linear decrescente, PSO com fator de constrição (Clerc & Kennedy), APSO (parâmetros adaptativos), CLPSO (comprehensive learning), DMS-PSO (dynamic multi-swarm).
- Representação para TSP: (a) codificação ordinal — tour representado por sequência de índices relativos; (b) codificação de caminho (path representation) — permutação direta de cidades; (c) operadores de swap — troca de arestas como análogo discreto da atualização de velocidade.
- Hibridização para TSP: PSO + 2-opt, PSO + 3-opt, PSO + Lin-Kernighan como pós-otimização local.
- MOPSO: NSGA-II adaptado com partículas, SPEA2 + PSO, crowding distance para diversidade do arquivo externo.
- Aplicações cobertas: processamento de sinais (FIR/IIR, DOA, beamforming), pesquisa operacional (TSP, VRP, job-shop scheduling), bioinformática (alinhamento de sequências, folding de proteínas), manufatura (usinagem, supply chain), redes de sensores (localização, cobertura).

## Evidência / Resultado Relevante

- A survey documenta que PSO híbrido com 2-opt/3-opt atinge soluções a ~3–5% do ótimo para instâncias TSP de até 100 cidades, competitivo com ACO básico.
- A representação por permutação com operadores de swap é identificada como a estratégia mais eficaz para TSP — consistente com a implementação do [[particle-swarm]] no TCC.
- O artigo é altamente citado (Google Scholar), indicando aceitação como referência canônica para PSO aplicado.

## Limitações de Uso

> [!warning] Limitação
> A survey é de 2015 e não cobre os avanços pós-deep learning (RL-PSO, surrogate PSO, CSO) documentados em [[zhu2025cumulative]]. A cobertura de TSP é qualitativa — o artigo não reporta resultados experimentais próprios em benchmarks TSPLIB, apenas compila achados da literatura. Para benchmarks quantitativos de PSO em TSP, é necessário complementar com artigos primários como [[clerc2000discretepso]] ou surveys mais recentes. O periódico (Mathematical Problems in Engineering, Hindawi) tem fator de impacto inferior ao Archives of Computational Methods in Engineering — a survey é citada por cobertura, não por prestígio do veículo.

## Conexões

- Fundamenta: [[kennedy1995particle]] — PSO original, ponto de partida de todas as variantes cobertas.
- Relacionado a: [[shami2022pso]] — cobre o mesmo terreno com taxonomia por mecanismo (complementar, mais atualizada).
- Relacionado a: [[gad2022pso]] — revisão sistemática posterior com metodologia PRISMA.
- Relacionado a: [[zhu2025cumulative]] — cobre os avanços posteriores a 2015 que esta survey não alcança.
- Relacionado a: [[clerc2000discretepso]] — formalização do PSO discreto para TSP, referência aplicada que a survey cita.
- Relacionado a: [[lin1973effective]] — heurística Lin-Kernighan usada como pós-otimização em PSO híbrido para TSP.
- Contrasta com: [[larranaga1999ga]] — survey de GA para TSP, mesmo gênero (survey de método bio-inspirado para problema combinatório) para método diferente.
- Apoia claim: PSO é aplicável a TSP com adaptações de representação adequadas; a representação é mais crítica que a variante algorítmica.
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- O insight principal de Zhang et al. — "a representação importa mais que a variante" — é metodologicamente valioso para o TCC: justifica dedicar atenção ao design da codificação de tours no espaço de partículas, não apenas à escolha de parâmetros.
- A cobertura de VRP conecta o TCC com a literatura de roteamento de veículos, que é o corpo teórico mais próximo do problema de patrulha com drones.
- A seção de processamento de sinais, embora fora do escopo direto do TCC, demonstra a versatilidade do PSO e pode ser citada para argumentar que o método não é "toy algorithm" — tem aplicações industriais sérias.
- Por ser de 2015, esta survey serve como "fotografia do passado" que, combinada com [[zhu2025cumulative]] (fotografia do presente), permite à monografia traçar a trajetória de evolução do PSO ao longo de uma década.
- As aplicações em bioinformática (alinhamento de sequências) são conceitualmente análogas a TSP (problemas de permutação), reforçando a adequação do PSO para esta classe de problemas.

## Citações-chave

> "The performance of PSO in solving a specific problem depends critically on the solution representation and the mapping between the particle space and the problem space."

> "For combinatorial optimization problems such as TSP and VRP, hybrid PSO algorithms that incorporate local search operators (2-opt, 3-opt) consistently outperform pure PSO variants."

> "PSO has evolved from a simple continuous optimizer into a versatile optimization framework with applications spanning signal processing, bioinformatics, operations research, and manufacturing."
