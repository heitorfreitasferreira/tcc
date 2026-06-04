---
title: "Review of Traveling Salesman Problem Solution Methods"
authors:
  - "Yang, Lianlian"
  - "Wang, Xing"
  - "He, Zhengbing"
  - "Wang, Shuaian"
  - "Lin, Ji"
year: 2023
doi: "10.1007/978-981-97-2275-4_1"
bibtex_key: yang2023review
bibtex-key: yang2023review
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
role: "revisao"
areas:
  - tsp
  - "comparative-studies"
methods:
  - exact
  - heuristic
  - metaheuristic
  - machine-learning
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/comparative-studies
  - metodo/exact
  - metodo/heuristic
  - metodo/metaheuristic
  - metodo/machine-learning
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/4
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/yang2023review.pdf]] -->

## Tese Central

Os métodos de solução para o TSP podem ser organizados em três grandes paradigmas — exatos, heurísticos e baseados em aprendizado — cada um com um perfil distinto de *trade-off* entre qualidade de solução, tempo de computação e generalidade, sendo que a escolha do método deve considerar as características específicas da instância e os requisitos da aplicação.

## Resumo

Yang et al. (2023) apresentam uma revisão abrangente dos métodos de solução para o Problema do Caixeiro Viajante, organizando-os em três categorias principais: métodos exatos (programação inteira, *branch-and-bound*, *branch-and-cut*), métodos heurísticos e meta-heurísticos (algoritmos genéticos, *simulated annealing*, *ant colony optimization*, *particle swarm optimization*), e métodos baseados em aprendizado (redes neurais, aprendizado por reforço, *graph neural networks*). Publicado como capítulo da conferência BIC-TA 2023 (*Bio-Inspired Computing: Theories and Applications*), o artigo enfatiza a evolução histórica dos métodos e oferece uma análise comparativa das vantagens e limitações de cada abordagem. O trabalho dedica atenção especial aos métodos bio-inspirados, alinhando-se ao escopo da conferência.

## Contribuições Principais

- Propõe uma categorização tripartite dos métodos de solução do TSP (exatos, heurísticos/meta-heurísticos, aprendizado de máquina), mais atualizada que surveys anteriores que ignoravam métodos baseados em aprendizado.
- Realiza análise comparativa qualitativa entre as três famílias quanto a escalabilidade, qualidade de solução, tempo de execução e generalidade.
- Discute a emergência de abordagens híbridas que combinam meta-heurísticas com aprendizado de máquina como tendência promissora.
- Fornece uma visão histórica da evolução dos métodos, contextualizando avanços recentes em *deep reinforcement learning* para TSP.

## Relevância para o TCC

- Oferece o arcabouço conceitual para a classificação dos métodos utilizados no TCC: força bruta (exato), GA, PSO e ACO (meta-heurísticos bio-inspirados).
- A análise de *trade-offs* entre as famílias de métodos fornece subsídios para interpretar os resultados experimentais do TCC (e.g., por que GA, PSO e ACO superam força bruta em escalabilidade).
- A discussão sobre métodos bio-inspirados — foco da conferência BIC-TA — contextualiza a escolha de GA, PSO e ACO como representantes desse paradigma.
- A menção a abordagens híbridas sugere direções de trabalho futuro que podem ser mencionadas na conclusão do TCC.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico (classificação e descrição dos métodos de solução do TSP; fundamentação dos algoritmos bio-inspirados)
- Claim(s) apoiado(s): GA, PSO e ACO são métodos meta-heurísticos consolidados para o TSP; métodos exatos não escalam para instâncias médias e grandes; abordagens híbridas representam o estado da arte
- Como citar na monografia: Usar a taxonomia tripartite para classificar os métodos do TCC no referencial teórico; citar a análise de *trade-offs* para justificar a escolha de meta-heurísticas.

## Métodos e Abordagens

- Métodos exatos: *branch-and-bound*, *branch-and-cut*, programação dinâmica (Held-Karp), formulações de programação inteira — garantem otimalidade, mas com complexidade exponencial.
- Meta-heurísticas: algoritmos genéticos, *particle swarm optimization*, *ant colony optimization*, *simulated annealing* — não garantem otimalidade, mas oferecem soluções de alta qualidade em tempo polinomial.
- Métodos baseados em aprendizado: *graph neural networks*, *pointer networks*, *reinforcement learning*, *attention models* — aprendem heurísticas a partir de dados, mas dependem de generalização para instâncias não vistas.
- Abordagens híbridas: combinação de meta-heurísticas com componentes de aprendizado de máquina para melhorar exploração ou explotação.

## Evidência / Resultado Relevante

- Métodos exatos resolvem otimamente instâncias de até centenas de cidades (≈100–1000), mas tornam-se impraticáveis para instâncias maiores.
- Meta-heurísticas como ACO e GA consistentemente encontram soluções dentro de 2–5% do ótimo para instâncias com milhares de cidades em tempo razoável.
- Métodos baseados em aprendizado (e.g., *attention models*) aproximam-se do desempenho de meta-heurísticas especializadas, mas com tempo de inferência muito inferior — embora exijam treinamento caro e não generalizem bem para distribuições de instâncias diferentes das de treinamento.
- Abordagens híbridas (meta-heurística + aprendizado) têm demonstrado os melhores resultados gerais nos benchmarks mais recentes.

## Limitações de Uso

> [!warning] Limitação
> O artigo é um capítulo de conferência, não um periódico, e tem escopo mais limitado que surveys abrangentes como o de Zhang et al. (2015). A categorização em três famílias é útil, mas simplifica a diversidade interna das meta-heurísticas (não distingue adequadamente entre GA, PSO e ACO). A cobertura de métodos baseados em aprendizado é introdutória e não cobre os avanços mais recentes pós-2022 (como *diffusion models* para TSP). O artigo não realiza experimentação própria — as comparações são baseadas em resultados da literatura.

## Conexões

- Fundamenta: TSP, classificação de métodos de solução, meta-heurísticas bio-inspiradas
- Relacionado a: [[zhang2015comprehensive]], [[ilavarasi2014variants]], [[alkhalifa2025comparative]], [[alexander2020comparison]]
- Contrasta com: [[saller2025approximability]] (foco em teoria de aproximação, não em métodos práticos) e [[khoufi2019survey]] (foco em UAVs, não em classificação geral de métodos)
- Apoia claim: Meta-heurísticas (GA, PSO, ACO) são métodos estabelecidos e eficazes para o TSP; métodos exatos não escalam para instâncias realistas
- Usado em capítulo: cap_referencial_teorico

## Notas e Insights

- A estrutura tripartite (exatos, heurísticos, aprendizado) é didaticamente útil e pode ser adotada no referencial teórico do TCC como fio condutor da revisão de métodos.
- O artigo é de 2023 e cobre métodos de aprendizado — isso o torna mais atual que surveys clássicos como Zhang et al. (2015), permitindo que o TCC reconheça a existência dessa terceira via mesmo optando por meta-heurísticas tradicionais.
- A publicação na BIC-TA (conferência de computação bio-inspirada) confere viés favorável aos métodos bio-inspirados, o que deve ser considerado ao usar o artigo como fonte — ele naturalmente enfatiza as vantagens de GA, PSO e ACO.
- A menção a abordagens híbridas como tendência pode ser usada na seção de trabalhos futuros do TCC.

## Citações-chave

> "TSP solution methods can be categorized into three groups: exact methods, heuristic/metaheuristic methods, and learning-based methods." — definição da taxonomia central do artigo.

> "Metaheuristic algorithms are the most widely used approaches for solving large-scale TSP instances due to their balance between solution quality and computational efficiency." — justificativa para o foco do TCC em meta-heurísticas.

> "Hybrid approaches combining metaheuristics with machine learning techniques represent a promising direction for future TSP research." — citação útil para a seção de trabalhos futuros do TCC.
