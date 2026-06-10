---
pdf: "papers/pdfs/gad2022pso.pdf"
title: "Particle Swarm Optimization Algorithm and Its Applications: A Systematic Review"
authors: [Gad, Ahmed G.]
year: 2022
doi: "10.1007/s11831-021-09694-4"
bibtex_key: gad2022pso
bibtex-key: gad2022pso
type: paper
reading_status: resumo-lido
validation_status: nao-validado
pdf_status: disponivel
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
  - status/resumo-lido
  - evidencia/referencia
  - metodo/pso
  - area/bio-inspired-optimization
  - area/particle-swarm
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/4
---

## PDF

![[gad2022pso.pdf]]

- Status: PDF obtido e validado em 2026-06-05 via Unpaywall/Springer, versão publicada, licença CC BY.
- Validação: `pdfinfo` OK, 31 páginas, 2.344.358 bytes; metadados confirmam título, autor, periódico e DOI `10.1007/s11831-021-09694-4`.

## Tese Central

Apesar da simplicidade conceitual do PSO, a explosão de variantes publicadas desde 1995 gerou fragmentação — cada artigo propõe sua própria modificação sem validação cruzada sistemática. Gad (2022) conduz uma revisão sistemática que mapeia tendências de publicação, identifica as famílias de variantes mais influentes e analisa criticamente os domínios de aplicação onde o PSO de fato entrega vantagem competitiva, oferecendo um guia de seleção orientado a dados para pesquisadores e praticantes.

## Resumo

Gad (2022) publica no Archives of Computational Methods in Engineering uma revisão sistemática de PSO que se diferencia das surveys narrativas ao empregar metodologia PRISMA: busca estruturada em bases (Scopus, WoS, IEEE Xplore), critérios de inclusão/exclusão explícitos e análise bibliométrica quantitativa. O artigo analisa tendências temporais de publicação (1995–2021), identificando picos de crescimento associados a marcos como PSO com inércia (1998), Comprehensive Learning PSO (2006) e PSO quântico (QPSO). As variantes são organizadas em: (1) modificações nos coeficientes do algoritmo (inércia adaptativa, aceleração variável, constriction factor); (2) alterações na estrutura populacional (topologias, múltiplos enxames, niching); (3) hibridização com outros métodos (GA, DE, ACO, SA, LS); e (4) extensões multiobjetivo. A análise de aplicações cobre engenharia (design estrutural, controle, eletrônica de potência), otimização combinatória (escalonamento, roteamento), aprendizado de máquina (seleção de features, treinamento de redes neurais) e problemas de energia (despacho econômico, microgrid). O artigo conclui com uma discussão sobre reprodutibilidade e a necessidade de benchmarks padronizados.

## Contribuições Principais

- Metodologia PRISMA aplicada à revisão de PSO — critérios de busca, inclusão e exclusão transparentes e replicáveis.
- Análise bibliométrica quantitativa de tendências de publicação ao longo de 27 anos.
- Taxonomia das variantes de PSO com validação por número de citações e impacto relativo.
- Mapeamento de domínios de aplicação com frequência de uso — identifica onde PSO é mais e menos competitivo.
- Discussão sobre reprodutibilidade e qualidade metodológica dos estudos de PSO — contribuição metacientífica rara em surveys de otimização.
- Identificação de lacunas: escassez de PSO em otimização combinatória discreta, baixa padronização de benchmarks.

## Relevância para o TCC

A revisão de Gad (2022) complementa [[shami2022pso]] ao adicionar rigor metodológico (PRISMA) e análise bibliométrica — dois aspectos ausentes em surveys narrativas. A identificação de que PSO é subexplorado em otimização combinatória discreta justifica diretamente a relevância do TCC: comparar PSO com GA e ACO em TSP é responder a uma lacuna documentada na literatura. A discussão sobre reprodutibilidade reforça a importância de relatar parâmetros, sementes e configurações experimentais — prática adotada no pipeline de experimentos do projeto. A análise de domínios de aplicação contextualiza o TSP dentro do ecossistema mais amplo de uso do PSO.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): Lacuna de PSO em otimização combinatória discreta; necessidade de comparação sistemática entre métodos bio-inspirados.
- Como citar na monografia: \cite{gad2022pso} — para justificar a relevância da comparação experimental PSO vs. GA vs. ACO e para a análise de tendências de publicação.

## Métodos e Abordagens

- Metodologia: revisão sistemática PRISMA (Preferred Reporting Items for Systematic Reviews and Meta-Analyses).
- Bases de dados: Scopus, Web of Science, IEEE Xplore.
- Período: 1995–2021.
- Critérios de inclusão: artigos em inglês, revisados por pares, foco primário em PSO.
- Análise bibliométrica: contagem de publicações por ano, distribuição por periódico, análise de citações.
- Taxonomia por mecanismo: (a) parâmetros — w variável, c₁/c₂ adaptativos, fator de constrição (Clerc & Kennedy 2002); (b) população — topologia (gbest, lbest, focal, Von Neumann), múltiplos enxames cooperativos, niching (SPSO, NichePSO); (c) hibridização — GA (operadores de crossover/mutação), DE (differential evolution), ACO, SA, busca tabu; (d) multiobjetivo — MOPSO com crowding distance (Coello et al. 2004), SMPSO (Nebro et al. 2009), dMOPSO.
- Domínios cobertos: engenharia estrutural, eletrônica de potência, controle, escalonamento, roteamento, aprendizado de máquina, energia, bioinformática.

## Evidência / Resultado Relevante

- Identifica que PSO com inércia adaptativa e constriction factor domina as aplicações de engenharia contínua, mas há lacuna significativa em otimização combinatória.
- A análise bibliométrica mostra crescimento exponencial de publicações entre 2005–2015, seguido de estabilização — sugerindo maturidade do campo.
- O artigo documenta que apenas ~8% dos estudos de PSO reportam todos os parâmetros e configurações necessárias para replicação.

## Limitações de Uso

> [!warning] Limitação
> Apesar da metodologia PRISMA, a revisão não realiza meta-análise quantitativa (effect sizes, intervalos de confiança) — as conclusões sobre eficácia relativa das variantes são qualitativas. O foco em publicações indexadas em inglês pode sub-representar contribuições de comunidades não anglófonas. A cobertura de PSO para TSP especificamente é limitada — o artigo cita PSO para roteamento mas não analisa benchmarks TSPLIB.

## Conexões

- Fundamenta: [[kennedy1995particle]] — PSO original.
- Relacionado a: [[shami2022pso]] — survey abrangente com taxonomia por mecanismo (complementar, sem PRISMA).
- Relacionado a: [[zhu2025cumulative]] — survey mais recente cobrindo 2018–2025, inclui avanços posteriores.
- Relacionado a: [[zhang2015comprehensive]] — survey anterior cobrindo aplicações em engenharia e pesquisa operacional.
- Relacionado a: [[clerc2000discretepso]] — PSO discreto, lacuna identificada por Gad.
- Contrasta com: [[pop2024comprehensive]] — survey sistemática de ACO, mesmo gênero metodológico para método distinto.
- Apoia claim: Necessidade de comparar PSO com outros métodos bio-inspirados em TSP.
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A principal contribuição única de Gad (2022) sobre [[shami2022pso]] é o rigor metodológico (PRISMA) — útil para defender a validade da revisão de literatura na monografia.
- A constatação de que PSO é subexplorado em domínios discretos é um argumento forte para a relevância científica do TCC — o projeto está preenchendo uma lacuna documentada.
- O artigo fornece uma justificativa baseada em evidência para a escolha dos parâmetros default do PSO no TCC (inércia decrescente, c₁=c₂=2.0).
- A análise de reprodutibilidade (apenas 8% dos estudos reportam tudo) serve como alerta metodológico para o design experimental do TCC.
- A curva de maturidade do campo (crescimento até 2015, estabilização depois) contextualiza o TCC em um campo maduro, onde surveys e comparações sistemáticas são o próximo passo natural.

## Citações-chave

> "Despite the large number of PSO variants, there is a lack of systematic reviews that follow a rigorous methodology to analyze and compare these variants."

> "PSO has been extensively applied to continuous optimization problems, while its application to discrete and combinatorial optimization remains relatively limited."

> "The lack of standardized benchmarks and reporting practices hinders the reproducibility and fair comparison of PSO variants."
