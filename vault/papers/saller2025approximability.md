---
title: "A Survey on Approximability of Traveling Salesman Problems Using the TSP-T3CO Definition Scheme"
authors:
  - "Saller, Sebastian"
  - "Koehler, Jana"
  - "Karrenbauer, Andreas"
year: 2025
doi: "10.1007/s10479-025-06641-5"
bibtex_key: saller2025approximability
bibtex-key: saller2025approximability
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 5
role: "revisao"
areas:
  - tsp
  - "tsp-variants"
  - lower-bounds
methods:
  - exact
  - approximation
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
  - area/lower-bounds
  - metodo/exact
  - metodo/approximation
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/5
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/saller2025approximability.pdf]] -->

## Tese Central

A aproximabilidade das variantes do TSP pode ser sistematicamente organizada e comparada por meio do esquema de definição TSP-T3CO (*Type, Time, Cost, Capacity, Order*), que unifica a descrição formal dos problemas e permite mapear, para cada combinação de restrições, os melhores limites inferiores e superiores de aproximação conhecidos na literatura.

## Resumo

Saller, Koehler e Karrenbauer (2025) apresentam o primeiro levantamento sistemático sobre a aproximabilidade das variantes mais conhecidas do Problema do Caixeiro Viajante, utilizando o esquema TSP-T3CO como framework unificado de definição. O artigo cataloga os melhores resultados de aproximação (limites inferiores de inaproximabilidade e limites superiores de algoritmos de aproximação) para dezenas de variantes, incluindo TSP com janelas de tempo, TSP com *deadlines*, TSP assimétrico, TSP com coleta e entrega, entre outras. Publicado no *Annals of Operations Research*, o trabalho estabelece uma referência canônica para pesquisadores que precisam situar a dificuldade teórica de uma nova variante do TSP. O esquema TSP-T3CO permite descrever qualquer variante por meio de cinco dimensões: tipo de grafo, restrições temporais, custos, capacidades e ordenação.

## Contribuições Principais

- Introduz e aplica o esquema TSP-T3CO como linguagem comum para definição precisa de variantes do TSP, permitindo comparação formal entre problemas.
- Realiza o primeiro *survey* abrangente de resultados de aproximabilidade para variantes do TSP, consolidando décadas de resultados dispersos em teoria da computação e otimização combinatória.
- Para cada variante catalogada, fornece os melhores limites inferiores (dureza de aproximação) e superiores (algoritmos de aproximação conhecidos), identificando *gaps* onde os limites ainda não coincidem.
- Destaca que várias variantes permanecem sem resultados de aproximabilidade publicados, revelando direções de pesquisa em aberto.

## Relevância para o TCC

- Oferece uma fundamentação teórica rigorosa para posicionar o TSP-SD-ATP no espectro de dificuldade computacional das variantes do TSP, utilizando o esquema TSP-T3CO.
- Os resultados de inaproximabilidade justificam teoricamente a escolha de métodos heurísticos e meta-heurísticos no TCC: se uma variante é NP-difícil de aproximar além de certo fator, métodos exatos tornam-se inviáveis.
- Fornece limites inferiores teóricos que podem ser usados como referência para avaliar a qualidade das soluções obtidas pelos algoritmos bio-inspirados implementados no TCC.
- O esquema TSP-T3CO pode ser usado para descrever formalmente o TSP-SD-ATP no capítulo de metodologia, conferindo precisão à definição do problema.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico (complexidade e aproximabilidade do TSP e suas variantes; justificativa teórica para métodos aproximativos)
- Claim(s) apoiado(s): O TSP e suas variantes são NP-difíceis e, portanto, métodos exatos não escalam para instâncias realistas; meta-heurísticas são justificadas como abordagem prática
- Como citar na monografia: Usar o esquema TSP-T3CO para descrever o TSP-SD-ATP no referencial teórico; citar os limites de aproximabilidade para contextualizar a dificuldade do problema e a necessidade de métodos heurísticos.

## Métodos e Abordagens

- Esquema TSP-T3CO (*Type, Time, Cost, Capacity, Order*): framework de definição formal que decompõe variantes do TSP em cinco dimensões ortogonais.
- Algoritmos de aproximação com garantia de fator constante: *Christofides-Serdyukov* (fator 3/2 para TSP métrico), Held-Karp (limite inferior via *minimum spanning tree*).
- Reduções de inaproximabilidade: demonstrações de que certas variantes não admitem aproximação de fator constante a menos que P = NP.
- Revisão sistemática da literatura de aproximabilidade, cobrindo resultados desde os anos 1970 até 2024.

## Evidência / Resultado Relevante

- O TSP métrico simétrico admite aproximação de fator 3/2 (Christofides-Serdyukov), e este é o melhor limite superior conhecido há quase 50 anos — um resultado clássico que ilustra a dificuldade de melhorar garantias mesmo para a variante mais básica.
- Várias variantes com restrições temporais (janelas de tempo, *deadlines*) são APX-difíceis, ou seja, admitem aproximação de fator constante mas não admitem PTAS.
- Para muitas variantes combinando duas ou mais restrições (e.g., TSP com janelas de tempo e capacidades), não há resultados de aproximabilidade publicados — o que inclui potencialmente o TSP-SD-ATP.

## Limitações de Uso

> [!warning] Limitação
> O artigo foca exclusivamente em resultados de aproximabilidade teórica (limites de pior caso), não em desempenho empírico de algoritmos. Os fatores de aproximação garantidos são frequentemente conservadores e não refletem o comportamento típico dos algoritmos em instâncias práticas. O esquema TSP-T3CO cobre variantes clássicas do TSP, mas pode não capturar completamente restrições específicas de aplicações com drones (como autonomia energética dependente do trajeto e condições de vento).

## Conexões

- Fundamenta: TSP, tsp-variants, complexidade computacional, limites inferiores teóricos
- Relacionado a: [[lawler1985traveling]], [[applegate2006traveling]], [[garey1979computers]], [[heldkarp1970traveling]]
- Contrasta com: [[khoufi2019survey]] (foco em aplicações práticas de drones, não em teoria da aproximação) e [[yang2023review]] (foco em métodos de solução práticos, não em garantias teóricas)
- Apoia claim: A necessidade de métodos heurísticos e meta-heurísticos para variantes realistas do TSP, dada a dureza de aproximação de muitas variantes
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- Este artigo é mais valioso para o TCC como fundamentação teórica do que como guia prático: ele fornece os argumentos formais para justificar por que meta-heurísticas são necessárias.
- O esquema TSP-T3CO é uma contribuição importante para a padronização terminológica — pode ser usado no TCC para definir precisamente o TSP-SD-ATP e evitar ambiguidades.
- O fato de o artigo ser de 2025 e publicado em periódico de alto impacto (*Annals of Operations Research*) confere atualidade e credibilidade às citações.
- A identificação de *gaps* na aproximabilidade de variantes combinadas sugere que o TSP-SD-ATP (que combina restrições de autonomia, múltiplos depósitos e patrulha) provavelmente está em uma classe sem resultados teóricos estabelecidos — o que reforça a relevância da abordagem experimental do TCC.

## Citações-chave

> "We provide the first systematic survey on best approximability results for well-known TSP variants using the TSP-T3CO definition scheme." (resumo) — declaração da contribuição central do artigo.

> "The TSP-T3CO scheme decomposes TSP variants along five independent dimensions: Type, Time, Cost, Capacity, and Order." — citação útil para descrever formalmente o TSP-SD-ATP.

> "For many variants combining two or more constraints from different dimensions, no approximability results have been published." — relevante para justificar a abordagem experimental do TCC diante da ausência de garantias teóricas.
