---
title: "Variants of Travelling Salesman Problem: A Survey"
authors:
  - "Ilavarasi, K."
  - "Joseph, K. S."
year: 2014
doi: "10.1109/ICICES.2014.7033850"
bibtex_key: ilavarasi2014variants
bibtex-key: ilavarasi2014variants
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 3
role: "revisao"
areas:
  - tsp
  - "tsp-variants"
methods:
  - heuristic
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/tsp-variants
  - metodo/heuristic
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/3
---

## PDF

<!-- PDF não disponível -->

## Tese Central

O TSP possui inúmeras variantes que estendem o problema clássico com restrições adicionais (lucro, janelas de tempo, cinemática, múltiplos caixeiros), e compreender essa diversidade é essencial para selecionar a modelagem e o método de solução adequados a cada aplicação prática.

## Resumo

Ilavarasi e Joseph (2014) apresentam um levantamento das principais variantes do Problema do Caixeiro Viajante, categorizando-as por tipo de restrição adicional em relação ao TSP clássico. O artigo cobre quatro famílias de variantes: (i) variantes baseadas em lucro (*prize-collecting TSP*, *orienteering problem*, *maximum TSP*), nas quais nem todos os vértices precisam ser visitados mas há recompensa associada a cada visita; (ii) variantes com janelas de tempo (*TSP with time windows*), que impõem intervalos de visita para cada vértice; (iii) variantes com restrições cinemáticas (*kinetic TSP*), nas quais os vértices se movem ao longo do tempo; e (iv) variantes com múltiplos caixeiros (*multiple TSP*). Publicado na conferência IEEE ICICES 2014, o artigo tem caráter introdutório e é adequado como porta de entrada para o estudo de variantes do TSP.

## Contribuições Principais

- Organiza as variantes do TSP em quatro famílias conceituais (lucro, temporal, cinemática, múltiplos agentes), oferecendo uma taxonomia didática.
- Para cada variante, descreve a formulação do problema, as principais aplicações práticas e os métodos heurísticos de solução propostos na literatura.
- Discute a relação entre as variantes e suas aplicações no mundo real, conectando problemas teóricos a contextos como logística, roteamento de veículos e sequenciamento de produção.
- Identifica o *multiple TSP* e o TSP com janelas de tempo como as variantes mais estudadas e com maior relevância prática.

## Relevância para o TCC

- Fornece uma introdução acessível às variantes do TSP, útil para contextualizar o TSP-SD-ATP como uma variante que combina restrições de múltiplos agentes (múltiplos drones) e autonomia (restrição temporal/energética).
- A discussão sobre *multiple TSP* é diretamente relevante para o cenário de patrulha com múltiplos drones do TCC.
- A cobertura de aplicações práticas ajuda a justificar a relevância do problema de patrulha com drones como instância concreta das variantes do TSP.
- O artigo serve como referência complementar de baixa complexidade para leitores que precisam de uma visão geral das variantes antes de consultar surveys mais densos como [[khoufi2019survey]] ou [[saller2025approximability]].

## Uso no TCC

- Capítulo(s): cap_referencial_teorico (introdução às variantes do TSP; caracterização do TSP-SD-ATP como variante com múltiplos agentes e restrições de autonomia)
- Claim(s) apoiado(s): O TSP clássico possui extensões relevantes para problemas reais de roteamento; o TSP com múltiplos caixeiros modela cenários com múltiplos drones
- Como citar na monografia: Usar a classificação de variantes para introduzir o tema no referencial teórico; citar a definição de *multiple TSP* para justificar a modelagem com múltiplos drones.

## Métodos e Abordagens

- Revisão descritiva de variantes com foco em formulação do problema e aplicações, sem avaliação experimental.
- Para cada variante, enumera heurísticas propostas na literatura (e.g., inserção mais próxima, 2-opt, algoritmos genéticos), mas sem análise comparativa de desempenho.
- Abordagem taxonômica baseada no tipo de restrição adicional (lucro, tempo, movimento, multiplicidade de agentes).

## Evidência / Resultado Relevante

- O *multiple TSP* (mTSP) é identificado como uma das variantes mais estudadas, sendo a base para problemas de roteamento de veículos e patrulha multi-agente.
- As variantes com janelas de tempo são as que apresentam maior número de aplicações práticas reportadas, especialmente em logística urbana.
- O *kinetic TSP* (com alvos móveis) é apontado como a variante menos explorada, representando uma lacuna de pesquisa que persiste até hoje.

## Limitações de Uso

> [!warning] Limitação
> Artigo de conferência com escopo limitado e profundidade modesta — não cobre meta-heurísticas em detalhe nem oferece análise experimental. Publicado em 2014, não inclui a literatura recente sobre variantes do TSP para drones (que explodiu a partir de 2016). A classificação em quatro famílias é simplificada e não captura variantes que combinam múltiplas restrições (como o TSP-SD-ATP, que combina múltiplos agentes com restrição de autonomia). Deve ser usado apenas como referência introdutória, complementado por surveys mais abrangentes como [[khoufi2019survey]] e [[saller2025approximability]].

## Conexões

- Fundamenta: TSP, tsp-variants, classificação de variantes do TSP
- Relacionado a: [[khoufi2019survey]] (variantes para UAVs, extensão natural das variantes clássicas), [[lawler1985traveling]] (referência clássica sobre TSP)
- Contrasta com: [[yang2023review]] (foco em métodos de solução, não em variantes do problema) e [[saller2025approximability]] (foco em aproximabilidade teórica, não em aplicações)
- Apoia claim: O TSP possui variantes que modelam cenários reais de roteamento com múltiplos agentes
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- Artigo adequado como "porta de entrada" para o tema de variantes do TSP: linguagem acessível, escopo delimitado, boas conexões com aplicações práticas.
- A cobertura do *multiple TSP* é particularmente útil para o TCC, pois estabelece a ponte entre o TSP clássico (um caixeiro) e o cenário de patrulha com múltiplos drones.
- Por ser um artigo de conferência de 2014 com foco em heurísticas simples, seu valor para o TCC é principalmente didático e de contextualização — não deve ser a única referência sobre variantes do TSP.
- A classificação proposta (lucro, tempo, movimento, multiplicidade) pode ser usada como estrutura inicial no referencial teórico, complementada por taxonomias mais modernas como o TSP-T3CO de [[saller2025approximability]].

## Citações-chave

> "The travelling salesman problem has many variants such as Profitable tour problem, Prize Collecting TSP, Orienteering Problem, Time dependent TSP, Max TSP, and many more." — abertura do artigo que estabelece a diversidade de variantes.

> "Multiple Travelling Salesman Problem (mTSP) is a generalization of the TSP where more than one salesman is used in the solution." — definição relevante para a modelagem com múltiplos drones no TCC.

> "The time window constraint is the most widely studied variant due to its practical importance in logistics and transportation." — justificativa para a relevância prática das variantes com restrições temporais.
