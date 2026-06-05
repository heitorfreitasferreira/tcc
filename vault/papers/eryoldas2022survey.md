---
title: 'A Literature Survey on Offline Automatic Algorithm Configuration'
authors:
  - Eryoldaş
  - Yasemin
year: 2022
doi: '10.3390/app12136316'
bibtex_key: eryoldas2022survey
bibtex-key: eryoldas2022survey
pdf: papers/pdfs/eryoldas2022survey.pdf
type: paper
reading_status: lido
validation_status: nao-validado
pdf_status: integro
rating: 4
role: revisao
areas:
  - tuning
  - metaheuristic
methods:
  - f-race
  - paramils
  - smac
  - irace
  - racing
  - meta-ea
chapters:
  - experimentos
claim_support: []
aliases:
  - Eryoldas survey 2022
  - offline algorithm configuration survey
tags:
  - area/tuning
  - evidencia/referencia
  - metodo/metaheuristic
  - papel/revisao
  - status/lido
  - tipo/paper
---

## PDF

![[eryoldas2022survey.pdf]]

## Tese Central

A configuração automática offline de algoritmos (parameter tuning) é essencial para extrair o desempenho máximo de meta-heurísticas. Esta revisão classifica e compara os principais métodos — model-free (Meta-EAs, ParamILS, REVAC, GGA++), racing (F-Race, I/F-Race, irace) e model-based (SMAC) — fornecendo um guia para pesquisadores selecionarem a abordagem adequada ao seu problema.

## Resumo

Os autores realizam uma revisão sistemática da literatura sobre configuração automática offline de parâmetros para meta-heurísticas. O artigo classifica os métodos em: (1) model-free (Meta-EAs como REVAC e GGA++, ParamILS), (2) métodos de racing (F-Race, I/F-Race, irace), (3) model-based (SMAC), e (4) baseados em design experimental. Para cada método, são discutidos o princípio de funcionamento, vantagens e limitações. O survey conclui com recomendações práticas e direções futuras, incluindo a necessidade de métodos que lidem com parâmetros condicionais e espaços de busca de alta dimensionalidade.

## Contribuições Principais

- Primeira revisão abrangente focada exclusivamente em configuração offline de algoritmos
- Taxonomia clara: model-free, racing, model-based, e experimental design
- Tabelas comparativas de vantagens/desvantagens de cada método
- Cobertura de F-Race, I/F-Race, irace, ParamILS, SMAC, REVAC, GGA++ e extensões multi-objetivo
- Recomendações para seleção de método conforme o tipo de parâmetro e objetivo de otimização

## Relevância para o TCC

Fundamenta a Seção 4.1 (Método para Avaliação) ao demonstrar que a configuração de parâmetros é um campo de pesquisa estabelecido, com métodos consolidados (irace, SMAC, ParamILS). Reforça a validade metodológica da escolha de parâmetros fixos quando acompanhada de justificativa baseada na literatura. Substitui as referências individuais a F-Race, irace e SMAC que eram inacessíveis (paywall).

## Uso no TCC

- Capítulo(s): 4 (Experimentos — Seção 4.1)
- Claim(s) apoiado(s): Parâmetros fixos são escolha metodológica válida e documentada
- Como citar na monografia: "A configuração automática de algoritmos é um campo ativo, com métodos que abrangem racing (F-Race), busca local iterada (ParamILS) e otimização baseada em modelos (SMAC), conforme revisado por ERYOLDAŞ; DURMUŞOĞLU (2022)."

## Métodos e Abordagens

- Revisão sistemática com classificação taxonômica dos métodos de configuração offline
- Análise comparativa baseada em: estratégia de busca, tipo de parâmetro, critério de parada, objetivo (single/multi), uso de surrogate models
- Cobertura de 20+ ferramentas e métodos publicados entre 1990–2022

## Evidência / Resultado Relevante

- ParamILS com adaptive capping acelera significativamente a configuração de algoritmos complexos como CPLEX
- irace (iterated F-Race) é o método de racing mais utilizado e mantido, com suporte a parâmetros categóricos e numéricos
- SMAC (model-based) utiliza random forests para modelar o desempenho do algoritmo em função dos parâmetros, sendo eficaz em espaços de alta dimensionalidade
- Métodos model-free (ParamILS, GGA++) são mais rápidos por avaliação, mas exigem mais avaliações totais

## Limitações de Uso

> [!warning] Limitação
> O survey foca em configuração offline; não cobre métodos de controle online de parâmetros. A comparação entre métodos é qualitativa — não há benchmark experimental unificado. As recomendações são baseadas na literatura revisada, não em experimentos próprios dos autores.

## Conexões

- Referencia: F-Race (Birattari et al.), irace (López-Ibáñez et al.), SMAC (Hutter et al.), ParamILS (Hutter et al.)
- Relacionado a: [[zhang2020tuning]] — tuning baseado em reinforcement learning
- Fundamenta: escolha metodológica de parâmetros fixos no Capítulo 4
- Usado em capítulo: experimentos (Seção 4.1)

## Notas e Insights

- O survey cobre exatamente os 3 métodos (F-Race, irace, SMAC) que precisávamos referenciar, unificando em uma única citação
- A classificação em model-free vs model-based ajuda a entender o trade-off entre custo por avaliação e número de avaliações
- O problema de over-tuning (viés de seleção) é discutido em profundidade — relevante para justificar por que não fizemos tuning por instância
- O artigo é CC-BY (MDPI), sem barreiras de acesso

## Citações-chave

> "Although the selection of the best performing values for free algorithm parameters [...] is a challenging and tedious task, it can lead to an effective and good performing version of these algorithms." (Eryoldaş & Durmuşoğlu, 2022, p. 1)

> "Parameter tuning became commonly utilized in industry and research and there is a significant advancement in this area." (Eryoldaş & Durmuşoğlu, 2022, p. 2)
