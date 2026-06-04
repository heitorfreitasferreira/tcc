---
title: "A Survey of Recent Extended Variants of the Traveling Salesman and Vehicle Routing Problems for Unmanned Aerial Vehicles"
authors:
  - "Khoufi, Ines"
  - "Laouiti, Anis"
  - "Adjih, Cedric"
year: 2019
doi: "10.3390/drones3030066"
bibtex_key: khoufi2019survey
bibtex-key: khoufi2019survey
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 5
role: "revisao"
areas:
  - tsp
  - "tsp-variants"
  - "drone-routing"
methods:
  - exact
  - heuristic
  - metaheuristic
chapters:
  - introducao
  - fundamentacao
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/tsp-variants
  - area/drone-routing
  - metodo/exact
  - metodo/heuristic
  - metodo/metaheuristic
  - capitulo/fundamentacao
  - capitulo/introducao
  - papel/revisao
  - relevancia/5
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/khoufi2019survey.pdf]] -->

## Tese Central

Os problemas de otimização de rota para veículos aéreos não tripulados (UAVs) constituem uma classe distinta de variantes do TSP e VRP que demandam formulações e técnicas de resolução específicas, não bastando a simples adaptação de métodos clássicos. O artigo argumenta que uma taxonomia unificada dessas variantes é necessária para orientar o desenvolvimento de novos algoritmos para aplicações com drones.

## Resumo

Khoufi, Laouiti e Adjih (2019) apresentam um levantamento abrangente das variantes estendidas do Problema do Caixeiro Viajante (TSP) e do Problema de Roteamento de Veículos (VRP) específicas para veículos aéreos não tripulados. O artigo propõe uma classificação taxonômica dos problemas de otimização de rotas para UAVs, abrangendo restrições como autonomia limitada de bateria, janelas de tempo, múltiplos depósitos, recarga em voo e zonas de exclusão aérea. Os autores categorizam as técnicas de resolução em métodos exatos, heurísticos e meta-heurísticos, discutindo vantagens e limitações de cada abordagem no contexto de drones. A pesquisa cobre a literatura publicada entre 2010 e 2019, período de intensa produção acadêmica sobre UAVs civis. O artigo tornou-se referência central na área, acumulando mais de 148 citações.

## Contribuições Principais

- Propõe uma taxonomia original para classificar variantes de TSP/VRP específicas para UAVs, organizando-as por restrições operacionais (energia, comunicação, regulação).
- Realiza mapeamento sistemático de 40+ artigos, identificando lacunas de pesquisa e direções futuras para otimização de rotas com drones.
- Compara qualitativamente métodos exatos, heurísticos e meta-heurísticos quanto à aplicabilidade em cenários reais de patrulha, entrega e monitoramento com UAVs.
- Destaca a escassez de benchmarks padronizados para comparação justa entre algoritmos de roteamento de drones.

## Relevância para o TCC

- Fornece a fundamentação teórica para justificar o problema TSP-SD-ATP como uma variante legítima do TSP para patrulha com drones, inserindo-o na taxonomia proposta.
- A classificação de restrições operacionais (bateria, zonas de voo, múltiplos pontos de partida) é diretamente aplicável à modelagem do cenário de patrulha do TCC.
- O mapeamento de métodos de resolução (exatos, heurísticos, meta-heurísticos) oferece o arcabouço para a escolha dos algoritmos comparados no TCC (GA, PSO, ACO, força bruta).
- A identificação de lacunas — como a carência de estudos comparativos entre meta-heurísticas bio-inspiradas para patrulha com drone — justifica a contribuição do TCC.

## Uso no TCC

- Capítulo(s): cap_introducao (contextualização de UAVs no TSP), cap_referencial_teorico (taxonomia de variantes e revisão de métodos de resolução para drones)
- Claim(s) apoiado(s): A modelagem do TSP-SD-ATP como variante do TSP para cenário de patrulha aérea; a escolha de GA, PSO, ACO e força bruta como métodos representativos para comparação
- Como citar na monografia: Usar a taxonomia do artigo para classificar o TSP-SD-ATP como variante do TSP com restrições de autonomia e múltiplos depósitos; citar a conclusão sobre a necessidade de benchmarks padronizados para motivar o desenho experimental do TCC.

## Métodos e Abordagens

- Métodos exatos: *branch-and-bound*, *branch-and-cut*, programação inteira mista (MILP) — viáveis apenas para instâncias pequenas com drones.
- Heurísticas construtivas: inserção mais próxima, vizinho mais próximo adaptadas para restrições de autonomia de UAVs.
- Meta-heurísticas: algoritmos genéticos, *simulated annealing*, *ant colony optimization* — predominantes na literatura de roteamento de drones por sua escalabilidade.
- Classificação taxonômica baseada em atributos do problema (tipo de veículo, restrições, objetivo) e atributos da solução (método, otimalidade, complexidade).

## Evidência / Resultado Relevante

- A maioria dos trabalhos (≈70%) utiliza métodos heurísticos ou meta-heurísticos, confirmando a inviabilidade de abordagens exatas para instâncias realistas de roteamento com múltiplos UAVs.
- A restrição de autonomia de bateria é a mais estudada, aparecendo em mais de 60% dos artigos revisados — alinhada ao cenário de patrulha do TCC.
- Apenas 15% dos trabalhos realizam comparação experimental entre diferentes famílias de algoritmos, o que expõe a lacuna que o TCC ajuda a preencher.

## Limitações de Uso

> [!warning] Limitação
> O survey cobre a literatura até 2019, portanto não inclui avanços recentes em aprendizado por reforço profundo para roteamento de UAVs. A taxonomia proposta foca em drones multi-rotor de pequeno porte; as restrições específicas de asa fixa (raio de curvatura, velocidade mínima) não são tratadas. O artigo não realiza experimentação própria — todas as comparações de desempenho são baseadas nos resultados reportados pelos trabalhos originais, sem reimplementação padronizada.

## Conexões

- Fundamenta: TSP, tsp-variants, modelagem do problema de patrulha com drone
- Relacionado a: [[muthanna2022uav]], [[rajan2022routing]], [[agatz2018optimization]], [[ahmed2024receding]]
- Contrasta com: [[saller2025approximability]] (foco em aproximabilidade teórica, não em aplicações de UAV) e [[ilavarasi2014variants]] (variantes clássicas do TSP, sem extensão para drones)
- Apoia claim: A necessidade de comparação experimental entre meta-heurísticas para cenários de patrulha com drone
- Usado em capítulo: cap_introducao, cap_referencial_teorico

## Notas e Insights

- Artigo essencial para o TCC: é o único survey que cruza TSP, VRP e UAVs de forma sistemática, fornecendo simultaneamente a justificativa do domínio (drones) e a fundamentação algorítmica (TSP).
- A taxonomia de restrições (bateria, comunicação, regulação) pode ser usada como *checklist* para validar se a modelagem do TCC cobre os aspectos relevantes do problema real.
- A observação de que poucos trabalhos comparam famílias diferentes de algoritmos no mesmo benchmark é um argumento forte para a relevância da contribuição experimental do TCC.
- O artigo está publicado na revista *Drones* (MDPI, open access), o que facilita o acesso ao texto completo para validação de citações.

## Citações-chave

> "A comprehensive survey that classifies the different existing routing problems for UAVs is missing." (p. 2) — justifica a própria existência do artigo e a lacuna que ele preenche.

> "Most of the proposed approaches are heuristics or metaheuristics." (p. 13) — evidência direta para a escolha de meta-heurísticas como foco do TCC.

> "There is a lack of common benchmarks for different UAV routing problems, which makes performance comparison difficult." (p. 14) — citação estratégica para motivar o desenho experimental do TCC com instâncias controladas.

> "The limited energy capacity of UAVs is one of the most studied constraints." (p. 8) — relevante para justificar a restrição de autonomia no TSP-SD-ATP.
