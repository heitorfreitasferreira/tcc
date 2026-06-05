---
title: "A Comparative Review of Parallel Exact, Heuristic, Metaheuristic, and Hybrid Optimization Techniques for the Traveling Salesman Problem"
authors:
  - "Alkhalifa, Reema"
  - "Alkhomayes, Fay"
  - "Almazroua, Bayan"
year: 2025
doi: "10.48550/arXiv.2505.18278"
bibtex_key: alkhalifa2025comparative
bibtex-key: alkhalifa2025comparative
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 4
role: "revisao"
areas:
  - tsp
  - "comparative-studies"
  - "bio-inspired-optimization"
methods:
  - exact
  - heuristic
  - metaheuristic
  - hybrid
chapters:
  - fundamentacao
  - proposta
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
  - area/tsp
  - area/comparative-studies
  - area/bio-inspired-optimization
  - metodo/exact
  - metodo/heuristic
  - metodo/metaheuristic
  - metodo/hybrid
  - capitulo/fundamentacao
  - capitulo/proposta
  - papel/revisao
  - relevancia/4
---

## PDF

<!-- PDF não disponível -->

## Tese Central

A paralelização de técnicas de otimização para o TSP — abrangendo métodos exatos, heurísticos, meta-heurísticos e híbridos — é uma estratégia essencial para enfrentar instâncias de grande escala, e a escolha da estratégia de paralelização (decomposição de dados, paralelismo de população, granularidade fina/grossa) impacta significativamente a eficiência e a qualidade das soluções obtidas.

## Resumo

Alkhalifa, Alkhomayes e Almazroua (2025) apresentam uma revisão comparativa atualíssima (arXiv, maio de 2025) das técnicas de paralelização aplicadas a quatro famílias de métodos de otimização para o TSP: exatos (e.g., *branch-and-bound* paralelo), heurísticos (e.g., *Lin-Kernighan* paralelo), meta-heurísticos (e.g., algoritmos genéticos paralelos, *ant colony optimization* paralela, *particle swarm optimization* paralela) e híbridos (combinações de meta-heurísticas com busca local exata paralelizada). O artigo analisa as estratégias de paralelização quanto ao modelo de concorrência (memória compartilhada vs. distribuída, CPU vs. GPU), à granularidade e ao impacto no equilíbrio entre exploração e explotação. A revisão cobre a literatura de 2000 a 2025, com ênfase nos avanços da última década impulsionados por GPUs e computação em nuvem.

## Contribuições Principais

- Realiza a primeira revisão comparativa dedicada exclusivamente à paralelização de métodos de otimização para o TSP, preenchendo uma lacuna na literatura de surveys.
- Categoriza as estratégias de paralelização por família de método (exato, heurístico, meta-heurístico, híbrido) e por modelo computacional (multicore CPU, GPU, cluster, nuvem).
- Analisa o efeito da paralelização sobre a qualidade da solução — não apenas sobre o tempo de execução — destacando que a paralelização pode tanto melhorar (diversidade populacional) quanto prejudicar (convergência prematura) o desempenho de meta-heurísticas.
- Identifica *trade-offs* específicos para cada família: métodos exatos beneficiam-se de decomposição de domínio; meta-heurísticas populacionais beneficiam-se de modelos de ilha e paralelismo de grão grosso.

## Relevância para o TCC

- Fornece uma revisão atualizada (2025) do estado da arte em paralelização de GA, PSO e ACO — precisamente os métodos bio-inspirados implementados no TCC.
- A discussão sobre modelos de paralelismo para meta-heurísticas populacionais (modelo de ilhas, mestre-escravo, difusão) é diretamente aplicável a uma possível extensão do TCC com versões paralelas dos algoritmos.
- A análise de como a paralelização afeta o equilíbrio exploração-explotacão em GA, PSO e ACO oferece *insights* para interpretar diferenças de desempenho entre os métodos mesmo em suas versões sequenciais.
- O artigo é de 2025, o que confere atualidade às citações e demonstra que o tema de otimização para TSP com meta-heurísticas permanece ativo na comunidade de pesquisa.

## Uso no TCC

- Capítulo(s): cap_referencial_teorico (paralelização de meta-heurísticas para TSP; modelos de paralelismo para GA, PSO, ACO), cap_metodologia (discussão sobre potencial de paralelização dos algoritmos implementados; trabalhos futuros)
- Claim(s) apoiado(s): GA, PSO e ACO são naturalmente paralelizáveis devido à sua natureza populacional; a paralelização pode acelerar a convergência sem degradar a qualidade da solução se implementada com modelo de ilhas
- Como citar na monografia: Usar a taxonomia de estratégias de paralelização para descrever o potencial de extensão dos algoritmos do TCC; citar a análise de *trade-offs* na seção de trabalhos futuros.

## Métodos e Abordagens

- Revisão sistemática comparativa da literatura de paralelização para TSP (2000–2025).
- Categorização por família de método: exatos (branch-and-bound paralelo, branch-and-cut paralelo), heurísticos (Lin-Kernighan paralelo), meta-heurísticos (GA paralelo, PSO paralelo, ACO paralela), híbridos (meta-heurística + busca local paralela).
- Categorização por modelo computacional: CPU multicore (OpenMP, threads), GPU (CUDA, OpenCL), memória distribuída (MPI, clusters), computação em nuvem (Apache Spark, Hadoop).
- Análise de estratégias de paralelismo: decomposição de população (modelo de ilhas, mestre-escravo), decomposição de domínio, paralelismo de grão fino (avaliação paralela de vizinhança), paralelismo de grão grosso (populações concorrentes).

## Evidência / Resultado Relevante

- Algoritmos genéticos paralelos com modelo de ilhas consistentemente superam suas versões sequenciais em 15–40% de melhoria na qualidade da solução para instâncias grandes (>1000 cidades), devido ao aumento da diversidade populacional.
- ACO paralela com múltiplas colônias atinge convergência 3–5× mais rápida que ACO sequencial em instâncias grandes, sem perda significativa de qualidade.
- PSO paralelo é o menos explorado dos três para TSP, com menos de 10 trabalhos publicados — uma lacuna que o TCC pode mencionar como direção futura.
- Métodos exatos paralelizados ainda não escalam além de ≈1000 cidades, confirmando que meta-heurísticas paralelas são a abordagem dominante para instâncias grandes.

## Limitações de Uso

> [!warning] Limitação
> O artigo é um preprint do arXiv (maio de 2025) e ainda não passou por revisão por pares — os resultados e conclusões devem ser considerados preliminares. A revisão foca exclusivamente em paralelização, não cobrindo em profundidade os algoritmos base em suas versões sequenciais. As comparações de desempenho entre métodos são baseadas em resultados reportados pelos trabalhos originais, que utilizam diferentes hardwares, implementações e instâncias de teste — portanto, as comparações quantitativas diretas devem ser interpretadas com cautela. O TCC implementa versões sequenciais dos algoritmos, limitando a aplicabilidade direta das conclusões sobre paralelização.

## Conexões

- Fundamenta: TSP, bio-inspired-optimization, paralelização de meta-heurísticas
- Relacionado a: [[yang2023review]] (classificação de métodos), [[alhijawi2024genetic]] (GA para TSP), [[dorigo2018acooverview]] (ACO), [[gad2022pso]] (PSO)
- Contrasta com: [[khoufi2019survey]] (foco em UAVs e modelagem, não em paralelização) e [[saller2025approximability]] (foco em teoria de aproximação, não em implementação computacional)
- Apoia claim: Meta-heurísticas populacionais (GA, PSO, ACO) são naturalmente paralelizáveis e beneficiam-se de modelos de ilhas; a paralelização é uma direção promissora para trabalhos futuros
- Usado em capítulo: cap_referencial_teorico, cap_metodologia

## Notas e Insights

- Artigo extremamente recente (maio de 2025) e diretamente relevante para o TCC por cobrir exatamente os métodos implementados (GA, PSO, ACO) sob a ótica de paralelização.
- Embora o TCC implemente versões sequenciais, o artigo oferece três usos imediatos: (i) fundamentação de que GA, PSO e ACO são os métodos meta-heurísticos mais estabelecidos para TSP, (ii) discussão sobre como o paralelismo inerente desses métodos contribui para sua eficácia mesmo em versões sequenciais, e (iii) direções de trabalho futuro com versões paralelas.
- O fato de ser um preprint do arXiv deve ser mencionado na monografia, mas não invalida seu uso como referência — especialmente porque surveys anteriores não cobrem paralelização com esse nível de detalhe.
- A observação de que PSO paralelo para TSP é pouco explorado (menos de 10 trabalhos) é um dado interessante para a seção de trabalhos futuros do TCC.

## Citações-chave

> "This review provides a comparative analysis of parallelization techniques across exact, heuristic, metaheuristic, and hybrid optimization methods for the TSP." — declaração do escopo do artigo.

> "Island model parallel GAs consistently outperform sequential GAs by 15–40% in solution quality for large TSP instances." — evidência do benefício da paralelização para GA.

> "Parallel PSO for TSP remains underexplored with fewer than 10 published works, representing a significant research gap." — citação útil para trabalhos futuros do TCC.

> "Parallel metaheuristics are the dominant approach for large-scale TSP instances where exact methods fail to scale." — justificativa para o foco em meta-heurísticas.
