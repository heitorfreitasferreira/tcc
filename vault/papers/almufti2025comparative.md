---
title: "Comparative Analysis of Metaheuristic Algorithms for Solving the Travelling Salesman Problems"
authors: [Almufti, Saman M., Shaban, Awaz Ahmed]
year: 2025
doi: "10.14419/7fk7k945"
bibtex-key: almufti2025comparative
pdf: "papers/pdfs/almufti2025comparative.pdf"
tags: [tsp, aco, gwo, abc, cso, metaheuristic, comparison]
status: lido
rating: 3
---

## PDF

![[almufti2025comparative.pdf]]

## Resumo

Almufti e Shaban comparam nove metaheurísticas no [[tsp]] clássico usando três instâncias TSPLIB: `berlin52`, `eil76` e `pr1002`. Os algoritmos avaliados são ACO, Lion Algorithm, Cuckoo Search, Grey Wolf Optimizer, Vibrating Particles System, Social Spider Optimization, Cat Swarm Optimization, Bat Algorithm e Artificial Bee Colony. O protocolo reporta 30 execuções independentes por instância e usa melhor custo, custo médio, desvio padrão e comportamento de convergência. O resultado central do artigo é que ACO, GWO e CSO apresentam o melhor equilíbrio entre qualidade de solução e robustez.

## Contribuições Principais

- Compara nove metaheurísticas em um protocolo unificado para TSP.
- Usa instâncias TSPLIB com escalas distintas: `berlin52`, `eil76` e `pr1002`.
- Avalia custo mínimo, média, desvio padrão e convergência.
- Identifica ACO, GWO e CSO como métodos mais equilibrados.
- Apresenta uma tabela sintética de forças e fraquezas por algoritmo.

## Relevância para o TCC

O artigo fortalece a justificativa de comparação empírica entre métodos bio-inspirados em TSP. Para o TCC, sua utilidade principal é contextual: ACO aparece como método robusto e estável em TSP, enquanto a discussão sobre variância e escalabilidade reforça a necessidade de múltiplas sementes, métricas de dispersão e análise por tamanho de instância. O estudo não substitui os experimentos do TCC porque usa TSP clássico.

## Métodos e Abordagens

- Benchmarks: `berlin52`, `eil76`, `pr1002`.
- Execuções: 30 rodadas independentes por algoritmo/instância.
- Métricas: melhor custo, custo médio, desvio padrão e tendência de convergência.
- ACO usa probabilidade baseada em feromônio e visibilidade heurística.
- GWO usa hierarquia alfa, beta e delta para balancear exploração e exploração local.
- CSO combina modos de busca local e rastreamento.

## Conexões

- [[tsp]]
- [[ant-colony]]
- [[comparative-studies]]
- [[hossain2024comparison]]
- [[wadi2025charting]]
- [[dorigo1997ant]]
- [[blum2005acointro]]
- [[chandra2022comparative]]

## Notas e Insights

- O artigo é útil para defender ACO como baseline forte em roteamento combinatório.
- A extração textual é boa o bastante para resumo e resultados, mas algumas tabelas vieram coladas e com quebras ruins.
- O estudo parece mais descritivo do que estatisticamente rigoroso; não foram recuperados testes estatísticos robustos na extração.
- Resultados numéricos específicos devem ser conferidos no PDF visualmente porque a extração de tabelas não preservou perfeitamente colunas.

## Citações-chave

> “The selected algorithms—Ant Colony Optimization (ACO), Lion Algorithm (LA), Cuckoo Search (CS), Grey Wolf Optimizer (GWO), Vibrating Particles System (VPS), Social Spider Optimization (SSO), Cat Swarm Optimization (CSO), Bat Algorithm (BA), and Artificial Bee Colony (ABC)—are evaluated on three standardized TSPLIB benchmark instances: berlin52, eil76, and pr1002.”

> “Notably, ACO, GWO, and CSO demonstrate superior balance between solution accuracy and robustness, making them promising candidates for large-scale combinatorial problems.”
