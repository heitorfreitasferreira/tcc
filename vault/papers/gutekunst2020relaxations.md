---
title: "Fantastic Relaxations of the TSP and How to Bound Them"
authors: [Gutekunst, Samuel C.]
year: 2020
doi: ""
bibtex_key: gutekunst2020relaxations
bibtex-key: gutekunst2020relaxations
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
role: "revisao"
areas:
  - tsp, "lower-bounds"]
methods: ["linear-programming", "held-karp", "semidefinite-programming"]
chapters:
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/lower-bound
  - metodo/exact
  - metodo/lower-bound
  - capitulo/fundamentacao
  - papel/revisao
  - relevancia/4
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/gutekunst2020relaxations.pdf]] -->

## Tese Central

A qualidade de uma relaxação para o TSP é medida pelo integrality gap — a razão entre o valor ótimo da relaxação e o valor ótimo do problema original. Gutekunst conduz um estudo abrangente das principais relaxações do TSP (Held-Karp, relaxações LP baseadas em diferentes formulações, relaxação semidefinida), caracterizando seus gaps de integralidade, relações de dominância entre elas, e como esses gaps se comportam em famílias especiais de instâncias como o TSP circulante.

## Resumo

Tese de doutorado (Cornell University, 2020) que investiga sistematicamente as relaxações do TSP e seus integrality gaps. O trabalho cobre: (1) a relaxação de Held-Karp (1-tree lagrangeana) e seu gap para o TSP circulante; (2) relaxações de programação linear baseadas em diferentes formulações (DFJ subtour, MTZ, multifluxo) e a hierarquia de força entre elas; (3) a relaxação semidefinida (SDP) do TSP e sua relação com a relaxação de Held-Karp; (4) limitantes superiores e inferiores para o integrality gap de famílias estruturadas de instâncias. A tese também aborda o "TSP circulante", uma família de instâncias com estrutura algébrica que permite análise teórica precisa dos gaps de relaxação, servindo como laboratório para entender o comportamento de diferentes relaxações.

## Contribuições Principais

- Caracterização precisa do integrality gap da relaxação de Held-Karp para o TSP circulante
- Análise comparativa da força relativa de diferentes formulações LP do TSP (DFJ, MTZ, multifluxo)
- Limites teóricos para o gap de relaxações LP em função de propriedades estruturais das instâncias
- Estudo da relaxação semidefinida (SDP) como alternativa mais forte (porém mais cara) que HK
- Identificação de instâncias onde o gap HK é grande (e.g., TSP circulante com certos parâmetros), contrariando a intuição de que HK é sempre justo
- Métodos para construir cotas superiores e inferiores do gap de integralidade sem resolver o TSP exatamente

## Relevância para o TCC

Esta tese fornece a compreensão mais profunda disponível sobre por que o limite de Held-Karp é tão bom para instâncias "típicas" e quando ele falha. Para o TCC, que usa HK como referência de qualidade nos experimentos com metaheurísticas, é crucial entender que existem famílias de instâncias (como o TSP circulante) onde o gap HK pode ser substancial (> 10-15%). Isso qualifica as conclusões experimentais: os gaps de otimalidade reportados para GA, PSO e ACO devem ser interpretados à luz da possibilidade de que o limite HK subestime o valor ótimo. A tese também estabelece a hierarquia de relaxações, permitindo posicionar corretamente HK no espectro entre relaxações fracas (MTZ) e fortes (SDP).

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): o limite HK é justo para instâncias euclidianas típicas mas pode ter gap significativo em instâncias estruturadas; a qualidade do limite inferior afeta diretamente a interpretação dos gaps de otimalidade nos experimentos
- Como citar na monografia: \cite{gutekunst2020relaxations}

## Métodos e Abordagens

- Relaxação lagrangeana de Held-Karp: 1-tree com penalidades nos graus dos vértices
- Relaxações LP: formulação DFJ (exponencial em restrições, separação por min-cut), MTZ (compacta mas fraca), multifluxo (compacta, intermediária)
- Relaxação semidefinida (SDP): relaxação do problema de atribuição quadrática com restrições de posto
- Integrality gap analysis: razão entre valor da relaxação e valor ótimo
- TSP circulante (Circulant TSP): família de instâncias definida por matriz circulante de distâncias — estrutura algébrica permite análise teórica exata dos gaps
- Métodos de bound tightening e planos de corte para fortalecer relaxações
- Limitantes superiores via heurísticas (Christofides, inserção) para estimar o gap sem o valor ótimo

## Evidência / Resultado Relevante

- O integrality gap da relaxação HK para o TSP circulante pode exceder 15% para certos parâmetros
- A relaxação SDP é estritamente mais forte que HK, mas o ganho é modesto (redução de ~1-3% no gap) e o custo computacional é ordens de magnitude maior
- A formulação DFJ é a mais forte entre as LP, mas tem número exponencial de restrições (resolvida via separação)
- Para instâncias euclidianas aleatórias (distribuição uniforme), o gap HK é consistentemente inferior a 1%, confirmando resultados empíricos anteriores
- Existe uma hierarquia estrita de relaxações: algumas formulações LP são estritamente dominadas por outras para certas famílias de instâncias

## Limitações de Uso

> [!warning] Limitação
> A tese foca em relaxações teóricas e suas propriedades matemáticas, não em implementação eficiente ou experimentação com instâncias de grande escala. O TSP circulante é uma construção teórica que não corresponde às instâncias de patrulha do TCC (pontos geográficos com distâncias euclidianas). As relaxações SDP, embora teoricamente mais fortes, são computacionalmente inviáveis para o pipeline de experimentos do TCC (que executa centenas de rodadas). A tese não aborda restrições temporais (makespan) ou múltiplos agentes (rTSP). Usar como referência conceitual para qualificar os gaps de otimalidade, não como fonte de algoritmos implementáveis.

## Conexões

- Fundamenta: [[heldkarp1970traveling]], [[heldkarp1971traveling]], TSP, lower-bounds
- Relacionado a: [[valenzuela1997estimating]], [[johnson1996asymptotic]], [[saller2025approximability]], [[applegate2006traveling]]
- Contrasta com: [[hoffman2013tspencyclopedia]] — este é análise teórica profunda dos gaps; Hoffman é visão enciclopédica ampla
- Apoia claim: o limite HK não é universalmente justo; existem instâncias com gap significativo
- Usado em capítulo: [[fundamentacao]]

## Notas e Insights

- O título "Fantastic Relaxations and How to Bound Them" é uma paródia de "Fantastic Beasts and Where to Find Them" — revela o estilo acessível da tese apesar do conteúdo matemático denso
- A descoberta de que o gap HK pode ser grande para o TSP circulante é uma contribuição teórica importante: estabelece que a "quase-otimalidade universal" do HK é uma propriedade das instâncias, não do limite em si
- A hierarquia de relaxações LP (MTZ < multifluxo < DFJ) é um resultado pedagógico valioso: formulações "equivalentes" para o problema exato podem ter forças de relaxação radicalmente diferentes
- A relaxação SDP, embora computacionalmente proibitiva para uso prático, estabelece um "teto" teórico para o que relaxações podem alcançar sem resolver o problema exato
- Para o TCC, o principal insight é metodológico: ao reportar gaps de otimalidade nos experimentos, deve-se incluir a ressalva de que o gap real (vs. valor ótimo) pode ser maior que o gap vs. HK, especialmente se as instâncias tiverem estrutura não uniforme

## Citações-chave

> "The Held-Karp relaxation is remarkably tight for typical Euclidean instances, but there exist structured families — such as the circulant TSP — where the integrality gap can exceed 15%."

> "Understanding the integrality gap of a relaxation is essential: a heuristic that achieves 99% of the Held-Karp bound may still be far from the true optimal tour if the bound itself is weak."

> "Different TSP formulations that are equivalent for the exact problem can yield relaxations of vastly different strengths, a phenomenon that underscores the art of formulation in integer programming."
