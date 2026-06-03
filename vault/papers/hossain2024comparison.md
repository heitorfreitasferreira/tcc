---
title: "Comparison of New and Old Optimization Algorithms for Traveling Salesman Problem on Small, Medium, and Large-Scale Benchmark Instances"
authors: [Hossain, Md Al Amin, Yılmaz Acar, Züleyha]
year: 2024
doi: "10.17798/bitlisfen.1380086"
bibtex-key: hossain2024comparison
pdf: "papers/pdfs/hossain2024comparison.pdf"
tags: [tsp, ga, aco, sa, abc, gwo, ssa, comparison, benchmark]
status: lido-parcial
rating: 4
---

## PDF

![[hossain2024comparison.pdf]]

## Resumo

Hossain e Yılmaz Acar comparam algoritmos “antigos” e “novos” para o [[tsp]] em instâncias pequenas, médias e grandes da TSPLIB. O grupo antigo inclui GA, ACO e simulated annealing. O grupo novo inclui Artificial Bee Colony, Grey Wolf Optimization e Salp Swarm Algorithm. O estudo usa instâncias `burma14`, `berlin52`, `kroA100`, `ts225`, `att532`, `rat783` e `dsj1000`, com população 100 e 1000 iterações para os métodos populacionais. As métricas incluem melhor solução, média, desvio padrão, desvio padrão percentual e tempo computacional, com testes t para comparar desempenho.

> [!warning] Leitura parcial
> O PDF é legível, mas a extração apresenta muitas palavras coladas e tabelas parcialmente desalinhadas. A nota resume o conteúdo recuperado, mas números de tabela devem ser conferidos visualmente antes de citação.

## Contribuições Principais

- Compara seis algoritmos em sete instâncias TSPLIB de diferentes escalas.
- Separa os métodos em “old” e “new” optimization algorithms.
- Inclui teste estatístico (`t-test`) para avaliar significância.
- Reporta qualidade, média, dispersão e tempo de execução.
- Mostra que resultados dependem fortemente do tamanho da instância.

## Relevância para o TCC

O artigo é diretamente relevante para o desenho experimental do TCC porque compara famílias próximas às usadas no projeto e enfatiza análise por tamanho de instância. Também reforça que uma conclusão geral como “algoritmo X é melhor” é frágil sem segmentar por escala e métrica. Para o TCC, isso apoia relatar qualidade da rota, variabilidade, tempo de execução e comportamento por instância.

## Métodos e Abordagens

- Algoritmos antigos: GA, ACO, SA.
- Algoritmos novos: ABC, SSA, GWO.
- Instâncias pequenas: `burma14`, `berlin52`, `kroA100`.
- Instâncias médias: `ts225`, `att532`.
- Instâncias grandes: `rat783`, `dsj1000`.
- Parâmetros comuns: população 100 e 1000 iterações.
- Testes t com nível de significância 0,05.

## Conexões

- [[tsp]]
- [[comparative-studies]]
- [[genetic-algorithms]]
- [[ant-colony]]
- [[almufti2025comparative]]
- [[wadi2025charting]]
- [[wu2020comparative]]
- [[chandra2022comparative]]

## Notas e Insights

- O artigo é bom como justificativa para separar análise por tamanho de instância.
- Há inconsistências textuais no PDF recuperado, incluindo provável troca de siglas em uma frase introdutória.
- Para o TCC, a lição mais segura é metodológica: múltiplas métricas, instâncias de tamanhos diferentes e teste estatístico evitam generalizações indevidas.

## Citações-chave

> “The goal of the research is to compare various algorithms' scalability, convergence, and computation times on benchmark instances of several sizes.”

> “The benchmark instances include burma14, berlin52, and kroA100 as small-sized instances, while ts225 and att532 represent medium-sized instances. Additionally, rat783 and dsj1000 are selected as large-sized instances.”
