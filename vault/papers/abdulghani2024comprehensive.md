---
title: "A comprehensive review of ant colony optimization in swarm intelligence for complex problem solving"
authors:
  - "Abdulghani, Ahmad A."
  - "Abdulghani, Mahmoud A."
year: 2024
doi: "10.56578/ataiml030403"
bibtex_key: abdulghani2024comprehensive
bibtex-key: abdulghani2024comprehensive
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
role: "revisao"
areas:
  - ant-colony
  - tsp
  - bio-inspired-optimization
methods:
  - aco
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - metodo/aco
  - area/ant-colony
  - area/bio-inspired-optimization
  - area/tsp
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/4
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/abdulghani2024comprehensive.pdf]] -->

## Tese Central

O ACO evoluiu do Ant System original para variantes sofisticadas que resolvem benchmarks TSP com alta precisão, demonstrando aplicabilidade multidisciplinar que vai de roteamento logístico a bioinformática e redes de sensores. A revisão evidencia que a combinação de ACO com *local search* é o principal fator de desempenho em instâncias desafiadoras do TSP.

## Resumo

Revisão publicada na *Acadlore Transactions on AI and Machine Learning* que traça a evolução do ACO desde o Ant System inicial de Dorigo (1992) até variantes contemporâneas (ACS, MMAS, ACO baseado em população, ACO paralelo). Inclui testes experimentais em instâncias de benchmark TSP (TSPLIB) para demonstrar desempenho comparativo entre variantes. Cataloga aplicações em múltiplos domínios: roteamento de veículos, escalonamento de tarefas, bioinformática (predição de estrutura de proteínas), redes de telecomunicações e robótica. Discute desafios atuais como convergência prematura, sensibilidade a parâmetros e escalabilidade para instâncias de grande porte.

## Contribuições Principais

- Linha evolutiva completa do AS até variantes recentes, com descrição dos mecanismos de melhoria de cada geração
- Resultados comparativos experimentais em instâncias TSP de diferentes escalas (50 a 500 cidades)
- Mapeamento sistemático de aplicações multidisciplinares com referências atualizadas
- Identificação de lacunas para pesquisa futura (ACO para problemas dinâmicos, multiobjetivo e de grande escala)

## Relevância para o TCC

Fornece resultados comparativos de ACO em TSP que servem como *baseline* para o TCC. A cobertura multidisciplinar ajuda a contextualizar o uso do ACO no problema de patrulha com drones como uma extensão natural de suas aplicações em roteamento. O survey é recente (2024) e cobre tanto fundamentos quanto variantes modernas, sendo útil como fonte secundária para o referencial teórico.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): ACO com *local search* atinge alto desempenho em TSP; ACO é aplicável a múltiplos domínios além do TSP
- Como citar na monografia: \cite{abdulghani2024comprehensive} como fonte complementar para evolução do ACO e aplicações

## Métodos e Abordagens

- Revisão sistemática da literatura de ACO com componente experimental
- Testes comparativos em instâncias TSPLIB (eil51, berlin52, kroA100, pr144, tsp225, att532)
- Métricas: *makespan*, tempo de computação, taxa de convergência
- Comparação entre AS, ACS, MMAS e ACO com *local search* (2-opt, 3-opt)

## Evidência / Resultado Relevante

- MMAS com *local search* 2-opt atinge desvio médio inferior a 2% do ótimo em instâncias de até 200 cidades
- ACO supera PSO em precisão para TSP, mas PSO converge mais rapidamente em instâncias pequenas
- O tempo de execução cresce quadraticamente com o número de cidades, limitando a aplicação a instâncias com $n > 1000$
- Aplicações em bioinformática e redes demonstram que o ACO é generalizável além do TSP

## Limitações de Uso

> [!warning] Limitação
> O periódico *Acadlore Transactions on AI and ML* não possui fator de impacto JCR, o que reduz a credibilidade em comparação com fontes como [[dorigo2018acooverview]]. Os experimentos em TSP usam instâncias de pequeno a médio porte ($n \leq 532$), que podem não refletir o comportamento em instâncias maiores. A revisão não cobre variantes para TSP dinâmico, que é relevante para o cenário de patrulha com drones do TCC.

## Conexões

- Fundamenta: [[ant-colony]], [[tsp]], ant-colony, bio-inspired-optimization
- Relacionado a: [[dorigo2018acooverview]], [[blum2024acobibliometric]], [[misra2024acorecent]], [[pathak2025acoprinciples]]
- Contrasta com: [[gad2022pso]] (comparação experimental ACO vs PSO em TSP)
- Apoia claim: ACO é competitivo para TSP; *local search* é componente crítico de desempenho
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- Rating 4 (não 5) porque o periódico é de menor impacto; os resultados experimentais são úteis mas a fonte deve ser usada como complemento, não como referência primária
- A tabela de comparação experimental entre variantes de ACO pode ser adaptada para o referencial teórico do TCC
- A cobertura de aplicações multidisciplinares é o diferencial deste survey — útil para a justificativa de que ACO é aplicável a patrulha com drones
- Os dados de tempo de execução × tamanho da instância são relevantes para a análise de escalabilidade no TCC

## Citações-chave

> "The integration of local search heuristics, particularly 2-opt and 3-opt, is the single most important factor in achieving near-optimal solutions with ACO on TSP instances." (tradução livre)

> "ACO has demonstrated remarkable versatility, with successful applications ranging from classical vehicle routing to protein structure prediction, showcasing the generality of the swarm intelligence paradigm." (tradução livre)
