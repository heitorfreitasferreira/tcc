---
title: "Choosing Mutation and Crossover Ratios for Genetic Algorithms—A Review with a New Dynamic Approach"
authors: ["Hassanat, Ahmad", "Almohammadi, Khalid", "Alkafaween, Esra'a", "Abunawas, Eman", "Hammouri, Awni", "Prasath, V. B. Surya"]
year: 2019
doi: "10.3390/info10120390"
bibtex_key: hassanat2019crossover
bibtex-key: hassanat2019crossover
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
role: "revisao"
areas:
  - genetic-algorithms
  - tsp
methods:
  - ga
chapters:
  - fundamentacao
  - proposta
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - metodo/ga
  - area/genetic-algorithms
  - area/tsp
  - area/bio-inspired-optimization
  - capitulo/fundamentacao
  - capitulo/proposta
  - papel/revisao
  - relevancia/4
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/hassanat2019crossover.pdf]] -->

## Tese Central

As taxas de crossover e mutação são os parâmetros mais críticos para o desempenho de algoritmos genéticos, mas não há consenso sobre valores ótimos universais; uma abordagem dinâmica que ajusta essas taxas ao longo da evolução — reduzindo a mutação e aumentando o crossover conforme a convergência avança — supera configurações estáticas tradicionais, especialmente em instâncias de TSP.

## Resumo

Artigo publicado na Information (2019) que revisa as práticas de parametrização de taxas de crossover (Pc) e mutação (Pm) em GAs, identificando ausência de consenso na literatura. Propõe uma nova abordagem dinâmica que ajusta Pc e Pm com base no progresso geracional: começa com mutação alta e crossover baixo (favorecendo exploração) e gradualmente inverte (favorecendo explotação). Valida a proposta experimentalmente em instâncias de TSP (Travelling Salesman Problem), demonstrando convergência mais rápida e soluções de melhor qualidade em comparação com taxas fixas convencionais. Diretamente relevante por testar em TSP — o mesmo problema-base do TCC.

## Contribuições Principais

- Revisão das práticas de escolha de Pc e Pm na literatura de GA, evidenciando a diversidade e contradição de recomendações
- Proposta de um método dinâmico determinístico (não-adaptativo) de variação de Pc e Pm ao longo das gerações: Pm decresce de ~50% a ~0% e Pc cresce de ~50% a ~100%
- Validação experimental em instâncias de TSP com comparação contra taxas fixas (Pc=0.9, Pm=0.01; Pc=0.8, Pm=0.2; etc.)
- Demonstração de que a abordagem dinâmica reduz o número de gerações necessárias para convergência e melhora a qualidade da solução final

## Relevância para o TCC

- Altamente relevante: testa a proposta diretamente em TSP, fornecendo evidência empírica de que parâmetros dinâmicos de GA melhoram desempenho no mesmo domínio do TCC
- Fundamenta escolhas de parametrização do GA nos experimentos comparativos (cap_metodologia)
- A revisão das práticas de Pc/Pm serve como justificativa para a configuração experimental adotada
- A abordagem dinâmica pode ser incorporada como variante do GA nos experimentos ou discutida como trabalho futuro

## Uso no TCC

- Capítulo(s): cap_referencial_teorico, cap_metodologia
- Claim(s) apoiado(s): Justificativa empírica para a parametrização do GA nos experimentos; evidência de que o ajuste de parâmetros impacta significativamente o desempenho em TSP
- Como citar na monografia: `\cite{hassanat2019crossover}` para fundamentar a escolha de taxas de crossover e mutação e discutir abordagens dinâmicas

## Métodos e Abordagens

- Revisão de literatura sobre taxas de crossover e mutação em GAs
- Proposta de função linear determinística para Pc e Pm: Pc(g) = Pc_min + (Pc_max − Pc_min) × (g/G) e Pm(g) = Pm_max − (Pm_max − Pm_min) × (g/G), onde g é a geração atual e G o total
- Experimentos em instâncias de TSP (biblioteca TSPLIB) com diferentes tamanhos de instância
- Comparação contra configurações estáticas típicas (Pc=0.8-1.0, Pm=0.001-0.1)
- Métricas: qualidade da solução (tour length), convergência (gerações até estabilização)

## Evidência / Resultado Relevante

- A abordagem dinâmica proposta superou as taxas fixas em todas as instâncias de TSP testadas, tanto em qualidade da solução quanto em velocidade de convergência
- Taxas fixas convencionais (Pc=0.9, Pm=0.01) frequentemente levam à convergência prematura em TSP
- O método é simples (função linear) e não introduz custo computacional adicional significativo

## Limitações de Uso

> [!warning] Limitação
> A abordagem dinâmica proposta é determinística (função linear do progresso geracional) e não adaptativa (não responde ao estado da população). O TCC pode usar a revisão de práticas de Pc/Pm como fundamentação, mas a abordagem dinâmica específica exigiria implementação adicional. As instâncias de TSP testadas são relativamente pequenas; generalização para instâncias maiores (30+ pontos) requer verificação. Complementar com [[larranaga1999ga]] para benchmarks de GA-TSP em larga escala.

## Conexões

- Fundamenta: [[alhijawi2024genetic]] — taxonomia geral de operadores de GA; [[goldberg1989genetic]] — fundamentos de parametrização de GA
- Relacionado a: [[umbarkar2015crossover]] — revisão complementar focada em operadores de crossover; [[potvin1996ga]] — GA aplicado a TSP com análise de operadores
- Contrasta com: [[larranaga1999ga]] — abordagem estática com ampla validação empírica em TSP
- Apoia claim: Evidência de que parametrização de GA impacta significativamente desempenho em TSP
- Usado em capítulo: cap_referencial_teorico, cap_metodologia

## Notas e Insights

- A função linear de decaimento/aumento de Pm/Pc é notavelmente simples e de fácil implementação — pode ser adicionada ao GA do TCC com poucas linhas de código
- O artigo destaca que valores altos de Pm (>0.1) são benéficos nas fases iniciais para manter diversidade, contrariando a prática comum de manter Pm ≤ 0.05 fixo — insight importante para a configuração do GA do TCC
- A validação em TSP torna este artigo uma das poucas fontes que combinam análise de parâmetros de GA com experimentação no domínio exato do TCC
- O apêndice com tabela de práticas de Pc/Pm na literatura (Tabela 1 do artigo) é uma referência rápida valiosa para justificar escolhas de parâmetros

## Citações-chave

> "There is no consensus on choosing a specific ratio or rate for crossover and mutation; in fact, there is a significant difference in the ratios used by researchers for the same problem." — justificativa central para o estudo

> "The proposed dynamic approach changes the crossover and mutation ratios linearly based on the number of generations, starting with high mutation and low crossover to encourage exploration, and ending with low mutation and high crossover to encourage exploitation."

> "The dynamic approach outperformed the fixed rates in all the TSP datasets used, achieving better solutions in fewer generations."
