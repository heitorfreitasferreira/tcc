---
title: "Ant colony optimization: overview and recent advances"
authors:
  - "Dorigo, Marco"
  - "Stützle, Thomas"
year: 2018
doi: "10.1007/978-3-319-91086-4_10"
bibtex_key: dorigo2018acooverview
bibtex-key: dorigo2018acooverview
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 5
role: "revisao"
areas:
  - ant-colony
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
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/5
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/dorigo2018acooverview.pdf]] -->

## Tese Central

O ACO é uma meta-heurística consolidada para problemas de otimização combinatória, com três variantes canônicas — Ant System (AS), Ant Colony System (ACS) e MAX-MIN Ant System (MMAS) — que representam diferentes compromissos entre intensificação (*exploitation*) e diversificação (*exploration*). O capítulo demonstra que, além do domínio discreto clássico, o ACO foi estendido com sucesso para otimização contínua (ACO_R), problemas multiobjetivo e aplicações dinâmicas.

## Resumo

Capítulo do *Handbook of Metaheuristics* (Springer, 3ª edição) escrito pelos criadores do ACO, Marco Dorigo e Thomas Stützle. Apresenta a meta-heurística ACO de forma abrangente e autoritativa: fundamentos biológicos (comportamento de forrageamento de formigas), descrição algorítmica formal do *Ant System* original, evolução para as três variantes principais (AS, ACS, MMAS), aplicações clássicas (TSP, roteamento de veículos, escalonamento, roteamento em redes) e avanços recentes incluindo otimização contínua com ACO_R, otimização multiobjetivo e integração com outras meta-heurísticas. O capítulo também discute a relação entre ACO e outras técnicas como *stochastic gradient descent* e *probabilistic learning*.

## Contribuições Principais

- Descrição canônica e formal da meta-heurística ACO pelos seus criadores, incluindo pseudocódigo das variantes principais
- Comparação estruturada entre AS, ACS e MMAS quanto a mecanismos de atualização de feromônio, elitismo e *local search*
- Extensão do ACO para otimização contínua via ACO_R (*ACO for continuous domains*), com distribuições de probabilidade em vez de grafos de construção
- Discussão de direções de pesquisa e problemas abertos (otimização multiobjetivo, problemas estocásticos, ACO automático)

## Relevância para o TCC

Fonte primária e autoritativa para o capítulo de referencial teórico sobre ACO. Serve como definição canônica dos algoritmos AS, ACS e MMAS. O código do TCC implementa Ant System com feromônio 3D. Por ser escrito pelos criadores do método, é a referência mais confiável para descrever formalmente o funcionamento do ACO. As discussões sobre TSP como *benchmark* clássico e sobre o equilíbrio exploração-intensificação são diretamente aplicáveis à comparação experimental do TCC.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): ACO é meta-heurística estabelecida com variantes AS, ACS, MMAS; feromônio codifica memória coletiva do enxame
- Como citar na monografia: \cite{dorigo2018acooverview} como referência primária para fundamentação do ACO

## Métodos e Abordagens

- Revisão estruturada da literatura de ACO com descrição algorítmica formal
- Apresentação de pseudocódigo para AS, ACS e MMAS
- Análise de componentes algorítmicos: regra de transição, atualização de feromônio (global, local, evaporação), heurística de visibilidade
- Discussão de hibridização com *local search* e outras meta-heurísticas

## Evidência / Resultado Relevante

- ACS e MMAS consistentemente superam o AS original em instâncias TSP de médio e grande porte
- ACO_R demonstra desempenho competitivo com métodos especializados em otimização contínua
- O uso de *local search* é crítico para que o ACO atinja desempenho de estado da arte em TSP
- Parâmetros como taxa de evaporação ($\rho$) e número de formigas ($m$) têm impacto significativo no desempenho

## Limitações de Uso

> [!warning] Limitação
> O capítulo cobre avanços até 2018, portanto não inclui variantes desenvolvidas nos últimos anos (ex.: ACO com ranqueamento para TSP dinâmico, abordado em [[misra2024acorecent]]). Por ser um *handbook chapter*, não reporta experimentos originais — os resultados citados provêm da literatura anterior. A implementação de ACO para o TCC deve considerar os parâmetros sugeridos, mas ajustá-los ao problema específico de patrulha com drones, que difere do TSP clássico por incluir múltiplos drones e restrições temporais.

## Conexões

- Fundamenta: [[ant-colony]], ant-colony, [[dorigo1996ant]], bio-inspired-optimization
- Relacionado a: [[blum2024acobibliometric]], [[abdulghani2024comprehensive]], [[misra2024acorecent]], [[pathak2025acoprinciples]]
- Contrasta com: [[gad2022pso]] (PSO), [[alhijawi2024genetic]] (GA) — métodos alternativos no referencial
- Apoia claim: ACO com *local search* é competitivo para TSP; parâmetros de feromônio controlam *exploration vs exploitation*
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- Escrito por Dorigo e Stützle, os criadores do ACO — é a referência mais autoritativa disponível sobre o tema
- O *Handbook of Metaheuristics* é uma coletânea de referência na área, o que amplifica a credibilidade do capítulo
- A descrição do ACO_R (otimização contínua) pode ser relevante se o TCC evoluir para formulações com variáveis contínuas
- Essencial para definir formalmente AS, ACS e MMAS na Seção 2.X do referencial teórico
- Combinar com [[blum2024acobibliometric]] (justificativa de relevância) e [[misra2024acorecent]] (variantes pós-2018)

## Citações-chave

> "The main underlying idea of ACO is that of simulating the foraging behavior of ants as a stochastic construction procedure that builds solutions probabilistically, guided by artificial pheromone trails and heuristic information." (tradução livre)

> "ACS and MMAS are the two most successful ACO variants, both significantly improving over the original AS by introducing mechanisms that better balance exploration and exploitation." (tradução livre)

> "The use of local search is crucial for ACO to achieve state-of-the-art performance on many combinatorial optimization problems, including the TSP." (tradução livre)
