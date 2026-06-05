---
title: "Computers and Intractability: A Guide to the Theory of NP-Completeness"
authors: "Michael R. Garey, David S. Johnson"
year: 1979
publisher: "W.H. Freeman and Company"
doi: null
tags:
  - np-completeness
  - computational-complexity
  - algorithm-design
  - theory
aliases:
  - Garey & Johnson (1979)
---

> **Nota:** Texto do PDF ilegível para extração automática (PDF baseado em imagem, sem camada de OCR). Resumo produzido com base no conteúdo canônico da obra.

## 1. Problema e Motivação

A obra estabelece a teoria da NP-completude como arcabouço formal para distinguir problemas computacionais tratáveis (classe P) daqueles considerados intratáveis (NP-difíceis). Diante da proliferação de problemas combinatórios para os quais não se conheciam algoritmos eficientes, os autores sistematizam as reduções polinomiais como ferramenta para provar equivalência de dificuldade entre problemas, fornecendo um guia prático para pesquisadores reconhecerem e lidarem com intratabilidade.

## 2. Método ou Abordagem Central

- Formalização das classes P, NP, NP-completo e NP-difícil, com o problema SAT (satisfatibilidade booleana) como ponto de partida via teorema de Cook (1971)
- Técnica de redução polinomial: transformação em tempo polinomial de um problema conhecido NP-completo para um problema-alvo, provando sua NP-completude
- Catálogo de mais de 300 problemas NP-completos organizados por área (teoria dos grafos, otimização combinatória, sequenciamento, autômatos, etc.)
- Diagramas de redução que mapeiam visualmente as relações de transformação entre problemas
- Estratégias de enfrentamento para problemas NP-difíceis: algoritmos de aproximação, heurísticas, casos especiais tratáveis e enumeração implícita

## 3. Principais Resultados

- Consolidação de ~300 problemas NP-completos com provas padronizadas de redução, cobrindo domínios como particionamento de grafos, cobertura de conjuntos, mochila, escalonamento e roteamento (incluindo TSP)
- Estabelecimento de uma hierarquia de dificuldade via reduções: problemas como 3-SAT, 3-Dimensional Matching, Vertex Cover, Hamiltonian Circuit e Traveling Salesman como pivôs centrais de redução
- Demonstração de que a distinção P vs. NP é a questão central em aberto da ciência da computação teórica, com implicações práticas diretas para projeto de algoritmos
- Identificação de subproblemas tratáveis (polinomiais) dentro de famílias NP-difíceis (e.g., 2-SAT, grafos bipartidos para certos problemas de cobertura)

## 4. Pontos Fortes e Limitações

**Pontos fortes:** Referência canônica e mais citada sobre NP-completude; estrutura didática que permite a pesquisadores sem formação profunda em complexidade utilizarem as reduções; catálogo abrangente que permaneceu como autoridade por décadas; notação e terminologia que se tornaram padrão na área.

**Limitações:** Não cobre avanços pós-1979 (complexidade parametrizada, algoritmos de aproximação com garantias, PCP theorem); assume familiaridade com conceitos básicos de algoritmos e teoria dos grafos; a classificação de problemas reflete o estado da arte da época (alguns problemas posteriormente resolvidos em P ou reclassificados).

## 5. Implicações Práticas para Pesquisadores

Ao enfrentar um novo problema combinatório, o fluxo de trabalho recomendado é: (i) tentar reduzi-lo a um problema conhecido em P; (ii) se isso falhar, consultar o catálogo de Garey & Johnson para tentar reduzir um problema NP-completo conhecido *para* o novo problema, provando sua NP-dificuldade; (iii) uma vez estabelecida a intratabilidade, buscar subproblemas restritos tratáveis ou projetar heurísticas e algoritmos de aproximação com fundamentação teórica. Para o contexto de TSP e otimização de rotas (relevante a esta monografia), o livro fornece as provas canônicas de NP-completude do TSP e variantes, fundamentando a justificativa teórica para o uso de meta-heurísticas bio-inspiradas.
