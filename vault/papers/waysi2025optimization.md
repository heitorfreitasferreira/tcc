---
title: "Optimization by Nature: A Review of Genetic Algorithm Techniques"
authors: ["Waysi, Daban", "Ahmed, Bestoon T.", "Ibrahim, Ibrahim Mahmood"]
year: 2025
doi: "10.33022/ijcs.v14i1.4596"
bibtex_key: waysi2025optimization
bibtex-key: waysi2025optimization
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 3
role: "revisao"
areas:
  - genetic-algorithms
  - bio-inspired-optimization
methods:
  - ga
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - metodo/ga
  - area/genetic-algorithms
  - area/bio-inspired-optimization
  - area/bio-inspired-optimization
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/3
---

## PDF

<!-- PDF não disponível -->

## Tese Central

Algoritmos genéticos permanecem uma das meta-heurísticas bio-inspiradas mais relevantes e amplamente adotadas para problemas de otimização, e os avanços recentes em GAs adaptativos, híbridos e paralelos demonstram sua evolução contínua como ferramenta de otimização, superando limitações clássicas de convergência prematura e escalabilidade.

## Resumo

Survey recente (2025) publicado no Indonesian Journal of Computer Science que revisa as técnicas de algoritmos genéticos e o estado atual da pesquisa na área. Organiza o campo em fundamentos clássicos (representação, seleção, crossover, mutação, elitismo) e avanços recentes (GAs adaptativos com parâmetros auto-ajustáveis, GAs híbridos combinados com busca local ou outras meta-heurísticas, GAs paralelos para escalabilidade). Discute aplicações em engenharia, otimização combinatória e aprendizado de máquina. Por ser de 2025, captura o estado da arte mais recente, embora com escopo mais limitado e em periódico de menor impacto em comparação com [[alhijawi2024genetic]].

## Contribuições Principais

- Revisão atualizada (2025) do estado da arte em algoritmos genéticos
- Organização das variantes modernas: GAs adaptativos (ajuste dinâmico de Pc, Pm, tamanho de população), híbridos (GA + simulated annealing, GA + busca tabu, GA + redes neurais), e paralelos (master-slave, coarse-grained, fine-grained)
- Discussão de tendências contemporâneas: integração com aprendizado de máquina, otimização multiobjetivo, GAs para big data
- Identificação de lacunas e direções futuras de pesquisa em GAs

## Relevância para o TCC

- Complementa [[alhijawi2024genetic]] com perspectiva ainda mais recente (2025), cobrindo avanços que podem não estar no survey de 2024
- A discussão sobre GAs híbridos e adaptativos fornece contexto para posicionar a implementação do GA no TCC (que adota uma versão canônica, não híbrida) e apontar direções de trabalho futuro
- Reforça a relevância contínua de GAs como método de otimização bio-inspirada, justificando sua inclusão na comparação experimental do TCC
- Rating 3 (em vez de 4) devido ao periódico de menor impacto; usar como referência complementar, não principal

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Evidência da relevância atual dos GAs como método de otimização; contextualização de variantes modernas (adaptativas, híbridas) como direções de trabalho futuro
- Como citar na monografia: `\cite{waysi2025optimization}` como referência complementar de atualidade na seção de algoritmos genéticos

## Métodos e Abordagens

- Revisão narrativa da literatura com organização temática (clássicos → variantes modernas → aplicações)
- Categorização de avanços: adaptativos, híbridos, paralelos
- Síntese qualitativa de tendências e lacunas de pesquisa
- Sem experimentação empírica própria

## Evidência / Resultado Relevante

- GAs híbridos (memetic algorithms) consistentemente superam GAs canônicos em benchmarks de otimização combinatória
- GAs adaptativos eliminam a necessidade de tuning manual de parâmetros, reduzindo o esforço de configuração experimental
- Paralelização é apontada como estratégia principal para escalar GAs a problemas de grande porte

## Limitações de Uso

> [!warning] Limitação
> Publicado em periódico de menor fator de impacto (Indonesian Journal of Computer Science); a revisão é qualitativa e genérica, sem experimentação própria ou benchmarks quantitativos. A cobertura de aplicações em TSP é superficial. Para o TCC, deve ser usado como referência terciária de atualidade, não como fonte primária de fundamentação teórica ou empírica. O survey [[alhijawi2024genetic]] (Evolutionary Intelligence, 599 citações) é preferível como referência principal.

## Conexões

- Fundamenta: [[alhijawi2024genetic]] — survey mais robusto com taxonomia detalhada de operadores; [[goldberg1989genetic]] — fundamentos clássicos
- Relacionado a: [[hassanat2019crossover]] — abordagem dinâmica de parâmetros como exemplo de GA adaptativo; [[umbarkar2015crossover]] — taxonomia complementar de operadores
- Contrasta com: [[shami2022pso]] — estado da arte equivalente para PSO; [[dorigo2018acooverview]] — estado da arte para ACO
- Apoia claim: Relevância atual e evolução contínua de GAs como justificativa para sua inclusão no TCC
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A data de publicação (2025) é uma vantagem estratégica: demonstra que a literatura sobre GAs continua ativa e relevante no momento da escrita da monografia
- A seção sobre GAs híbridos pode ser usada para contextualizar a limitação do escopo do TCC (GA canônico, sem hibridização) e sugerir work futuro
- O survey discute GAs para big data — tangencial ao TCC, mas útil para posicionar o TSP como um problema de otimização combinatória clássico, distinto de aplicações de larga escala em dados
- Por ser genérico e sem experimentação, usar com moderação no texto; preferir [[alhijawi2024genetic]] como referência principal de GA no referencial teórico

## Citações-chave

> "Genetic algorithms continue to be one of the most popular and effective bio-inspired optimization techniques, with ongoing research addressing their classical limitations through adaptive, hybrid, and parallel approaches."

> "Hybrid genetic algorithms, also known as memetic algorithms, combine the global search capability of GAs with local search methods to achieve better convergence and solution quality."

> "Adaptive genetic algorithms dynamically adjust their parameters during the evolutionary process, eliminating the need for manual tuning and improving robustness across different problem instances."
