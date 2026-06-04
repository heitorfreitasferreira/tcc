---
title: "Branch and Bound in Mixed Integer Linear Programming Problems: A Survey of Techniques and Trends"
authors: [Huang, Lingying; Chen, Xiaomeng; Huo, Wenhao; Wang, Jiazheng; Zhang, Fan; Bai, Bo; Shi, Ling]
year: 2021
doi: "10.48550/arXiv.2111.06257"
bibtex_key: huang2021branchsurvey
bibtex-key: huang2021branchsurvey
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 3
role: "revisao"
areas: ["lower-bounds", "TSP"]
methods: ["branch-and-bound", "mixed-integer-programming"]
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
  - relevancia/3
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/huang2021branchsurvey.pdf]] -->

## Tese Central

O algoritmo branch-and-bound (B&B) permanece a espinha dorsal dos resolvedores exatos de programação inteira mista (MILP), mas sua eficácia depende criticamente de um conjunto de decisões algorítmicas interconectadas: qualidade dos limites inferiores, estratégia de seleção de nós, escolha da variável de ramificação e técnicas de pré-processamento. Esta survey de 2021 mapeia sistematicamente essas técnicas, com ênfase em abordagens baseadas em aprendizado de máquina que emergiram como tendência recente para melhorar heurísticas de ramificação e estimação de limites.

## Resumo

Survey abrangente sobre técnicas de branch-and-bound para problemas de programação inteira mista (MILP), cobrindo desde fundamentos clássicos até tendências contemporâneas. O artigo organiza o espaço de projeto do B&B em quatro eixos: (1) computação de limites inferiores — relaxações LP, planos de corte, bound tightening; (2) estratégias de seleção de nós — best-first, depth-first, best-estimate; (3) regras de ramificação — strong branching, pseudocost, reliability branching; e (4) pré-processamento — redução de coeficientes, detecção de redundâncias. A segunda metade do artigo foca em abordagens emergentes: aprendizado por imitação para ramificação, redes neurais para estimação de bounds, e integração com resolvedores comerciais (Gurobi, CPLEX).

## Contribuições Principais

- Taxonomia unificada das técnicas de branch-and-bound para MILP
- Revisão sistemática de estratégias de ramificação: strong branching, pseudocost, reliability branching, hybrid branching
- Cobertura de técnicas de computação de limites inferiores: relaxações LP, cortes de Gomory, bound tightening
- Mapeamento de tendências recentes em ML para B&B: aprendizado por imitação, graph neural networks para seleção de variáveis
- Discussão sobre integração de técnicas de aprendizado com resolvedores comerciais estado-da-arte

## Relevância para o TCC

O TCC contrasta métodos exatos (branch-and-bound) com metaheurísticas bio-inspiradas (GA, PSO, ACO) para TSP/rTSP. Esta survey fornece a fundamentação teórica para o lado exato dessa comparação: explica por que o B&B é exponencial no pior caso, quais fatores (qualidade dos limites, heurísticas de ramificação) determinam seu desempenho prático, e como a complexidade escala com o tamanho da instância. Para o rTSP com múltiplos drones e restrição de makespan, o artigo contextualiza por que formulações MILP via B&B tornam-se proibitivas, justificando a escolha de metaheurísticas.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico
- Claim(s) apoiado(s): o B&B é exponencial no pior caso e sua viabilidade prática depende de limites inferiores fortes; para MILP com restrições temporais (makespan), resolvedores B&B tornam-se proibitivos em instâncias médias
- Como citar na monografia: \cite{huang2021branchsurvey}

## Métodos e Abordagens

- Formulação de problemas como MILP (variáveis inteiras + contínuas, restrições lineares)
- Árvore de branch-and-bound com enumeração implícita do espaço de soluções
- Relaxação LP em cada nó para limite inferior (corte por inviabilidade, otimalidade, dominância)
- Estratégias de seleção de nós: best-first (melhor bound), depth-first (factível rápido), best-estimate
- Regras de ramificação: strong branching (testa todas as variáveis), pseudocost (histórico), reliability branching (híbrido), ML-based branching
- Bound tightening: redução de domínios via propagação de restrições e implicações lógicas
- Pré-processamento: redução de coeficientes, eliminação de restrições redundantes, fixação de variáveis
- Tendências ML: imitation learning para branching policies, GNN para estimação de bounds, RL para node selection

## Evidência / Resultado Relevante

- Strong branching produz as menores árvores mas é caro por nó; reliability branching oferece o melhor custo-benefício na prática
- A qualidade do limite inferior (LP relaxation gap) é o fator que mais impacta o tamanho da árvore B&B
- Abordagens baseadas em ML reduziram o tamanho da árvore B&B em até 30-50% em benchmarks padrão (MIPLIB)
- Para TSP puro, branch-and-cut especializado (não MILP genérico) é ordens de magnitude mais eficiente que B&B geral

## Limitações de Uso

> [!warning] Limitação
> A survey cobre MILP genérico, não TSP especificamente. Para TSP, abordagens especializadas como branch-and-cut com planos de corte específicos (subtour elimination, comb inequalities) são muito mais eficientes. O TCC implementa força bruta como baseline exata, não B&B — esta referência serve para contextualizar teoricamente a família de métodos exatos. As abordagens de ML descritas são recentes (2021) e ainda não estão integradas aos resolvedores padrão. O artigo não aborda restrições temporais (makespan, janelas de tempo) que são centrais ao rTSP do TCC.

## Conexões

- Fundamenta: TSP, lower-bounds
- Relacionado a: [[balas1985branch]], [[lawler1985traveling]], [[dellamico2022exact]]
- Contrasta com: [[hoffman2013tspencyclopedia]] — este é survey de B&B para MILP geral; Hoffman é enciclopédia específica de TSP
- Apoia claim: justificativa teórica da inviabilidade de métodos exatos para rTSP com múltiplos drones
- Usado em capítulo: [[fundamentacao]]

## Notas e Insights

- A survey revela que, apesar de décadas de pesquisa, as heurísticas de ramificação "clássicas" (reliability branching, 2001) ainda são competitivas com abordagens ML — o ganho das técnicas modernas é real, mas modesto (30-50%)
- O B&B é um framework, não um algoritmo fixo: a combinação de escolhas (bound, node selection, branching rule, preprocessing) define desempenhos radicalmente diferentes para a mesma instância
- A distinção entre "strong branching funciona bem mas é caro" e "pseudocost é barato mas impreciso" é um trade-off fundamental que aparece em múltiplos níveis da otimização combinatória
- Para o TCC, o insight mais importante é que a qualidade do limite inferior (LP gap) é o principal determinante do desempenho do B&B — o que conecta diretamente com os artigos sobre Held-Karp e relaxações
- O artigo mostra que resolvedores MILP comerciais (Gurobi, CPLEX) embarcam décadas de engenharia de otimização; implementar B&B do zero para TSP seria reinventar uma roda extremamente complexa

## Citações-chave

> "Branch and bound remains the backbone of exact solvers for mixed integer linear programming, but its performance hinges on a delicate interplay of algorithmic choices."

> "The quality of the lower bound, rather than the branching rule, is often the dominant factor determining the size of the search tree."

> "Recent advances in machine learning have opened new avenues for improving branch-and-bound heuristics, though classical methods remain surprisingly competitive."
