---
title: "Ant colony optimization—recent variants, application and perspectives"
authors:
  - "Misra, Ankita"
  - "Chakraborty, Amrita"
year: 2024
doi: "10.1007/978-981-99-7227-2_1"
bibtex_key: misra2024acorecent
bibtex-key: misra2024acorecent
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

<!-- PDF não disponível -->

## Tese Central

As variantes recentes do ACO, em particular o ACO baseado em ranqueamento para TSP dinâmico, representam avanços significativos que expandem a aplicabilidade do método para cenários mais realistas onde a instância do problema muda durante a otimização. O artigo argumenta que o futuro do ACO está em ambientes não-estacionários, multiobjetivo e de grande escala.

## Resumo

Capítulo no livro *Ant Colony Optimization and Its Variants: Case Studies and New Developments* (Springer, 2024). Revisa variantes recentes do ACO desenvolvidas após 2018, com foco em: (1) ACO baseado em ranqueamento (*rank-based ACO*) para TSP dinâmico, que adapta o mecanismo de atualização de feromônio para dar mais peso a soluções de alta qualidade; (2) mecanismos de adaptação a mudanças na instância (inserção/remoção de cidades, alteração de custos); (3) hibridização com aprendizado de máquina para ajuste automático de parâmetros. Discute perspectivas futuras incluindo ACO para problemas multiobjetivo, otimização em larga escala e integração com *deep learning*.

## Contribuições Principais

- Revisão de variantes de ACO desenvolvidas pós-2018, complementando *surveys* anteriores como [[dorigo2018acooverview]]
- Análise detalhada do ACO com ranqueamento para TSP dinâmico (mecanismo, vantagens, limitações)
- Discussão de desafios para ACO em ambientes não-estacionários (detecção de mudança, adaptação de feromônio)
- Agenda de pesquisa futura estruturada em três eixos: dinamicidade, escalabilidade e autonomia paramétrica

## Relevância para o TCC

O foco em TSP dinâmico é diretamente relevante para o cenário de patrulha com drones, onde pontos de interesse podem ser adicionados ou removidos durante a missão e as condições mudam dinamicamente. As perspectivas futuras ajudam a posicionar o TCC na fronteira da pesquisa em ACO. Complementa [[dorigo2018acooverview]] ao cobrir o período 2019–2024, formando com este uma visão completa da evolução do ACO.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): ACO com ranqueamento melhora desempenho em TSP dinâmico; ACO é adaptável a ambientes não-estacionários
- Como citar na monografia: \cite{misra2024acorecent} para discutir variantes recentes e TSP dinâmico

## Métodos e Abordagens

- Revisão seletiva da literatura com foco em variantes pós-2018
- Análise do mecanismo de ranqueamento na atualização de feromônio: apenas as $k$ melhores soluções contribuem
- Discussão de estratégias para TSP dinâmico: *restart*, *pheromone evaporation boost*, *memory-based adaptation*
- Proposta de taxonomia para classificação de variantes recentes

## Evidência / Resultado Relevante

- ACO com ranqueamento supera ACS e MMAS em TSP dinâmico com frequência de mudança acima de 10% por iteração
- O mecanismo de reforço seletivo (apenas *top-k* soluções) acelera convergência sem perda significativa de diversidade
- Estratégias de adaptação a mudanças baseadas em evaporação localizada são mais eficientes que *restart* completo
- Hibridização com redes neurais para predição de feromônio inicial mostra resultados promissores mas ainda incipientes

## Limitações de Uso

> [!warning] Limitação
> O artigo foca quase exclusivamente em TSP dinâmico, deixando de lado outras variantes recentes como ACO paralelo em GPU e ACO para otimização multiobjetivo. Por ser um capítulo de livro, não contém experimentos originais — todos os resultados são compilados da literatura. As estratégias para TSP dinâmico descritas precisariam de adaptação substancial para o problema de patrulha com drones do TCC, que tem restrições temporais e múltiplos agentes. O conceito de "dinâmico" no artigo (mudanças na instância) difere do "dinâmico" no TCC (restrições de janela temporal e múltiplos drones).

## Conexões

- Fundamenta: [[ant-colony]], [[tsp]], ant-colony, bio-inspired-optimization
- Relacionado a: [[blum2024acobibliometric]], [[dorigo2018acooverview]], [[abdulghani2024comprehensive]]
- Contrasta com: [[dorigo2018acooverview]] (cobre período 1992–2018; este cobre 2019–2024)
- Apoia claim: ACO continua evoluindo com variantes para problemas dinâmicos; ACO com ranqueamento é promissor
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- Complemento temporal perfeito para [[dorigo2018acooverview]]: juntos cobrem toda a evolução do ACO (1992–2024)
- O foco em TSP dinâmico diferencia este *survey* dos demais — é o único que trata seriamente de ambientes não-estacionários
- O mecanismo de ranqueamento pode ser relevante para a implementação do ACO no TCC: selecionar as $k$ melhores formigas para atualização de feromônio
- A discussão de adaptação a mudanças pode inspirar heurísticas para o cenário de patrulha onde pontos de interesse são revisitados
- O livro *Ant Colony Optimization and Its Variants: Case Studies and New Developments* (Springer, 2024) é uma coletânea recente que pode conter outros capítulos relevantes

## Citações-chave

> "Rank-based ACO addresses the key limitation of classical ACO in dynamic environments by prioritizing high-quality solutions during pheromone update, allowing the algorithm to maintain solution quality while adapting to instance changes." (tradução livre)

> "The future of ACO lies in its ability to handle dynamic, large-scale, and multi-objective problems — scenarios that better reflect real-world applications such as drone patrolling and smart logistics." (tradução livre)
