---
title: "Comparative Analysis of Metaheuristic Algorithms for Solving the Travelling Salesman Problems"
authors: [Almufti, Saman M., Shaban, Awaz Ahmed]
year: 2025
doi: "10.14419/7fk7k945"
bibtex-key: almufti2025comparative
pdf: "papers/pdfs/almufti2025.pdf"
tags: [tsp, aco, gwo, abc, cso, metaheuristic, comparison]
status: lido
rating: 3
---

## PDF

![[almufti2025.pdf]]

## Resumo

Compara 9 metaheurísticas — Ant Colony Optimization (ACO), Lion Algorithm (LA), Cuckoo Search (CS), Grey Wolf Optimizer (GWO), Vibrating Particles System (VPS), Social Spider Optimization (SSO), Cat Swarm Optimization (CSO), Bat Algorithm (BA) e Artificial Bee Colony (ABC) — em 3 instâncias TSPLIB (berlin52, eil76, pr1002) com 30 execuções independentes cada. ACO, GWO e CSO destacam-se com melhor equilíbrio entre acurácia e robustez. O estudo fornece um benchmark padronizado sob mesmas condições experimentais, incluindo análise de convergência e testes estatísticos.

## Contribuições Principais

- Benchmark padronizado de 9 algoritmos no mesmo framework computacional
- Análise de convergência por instância com 30 runs independentes
- Identificação de ACO, GWO e CSO como melhores desempenhos em TSP
- Comparação estatística sistemática (desvio padrão, média, melhor solução)

## Relevância para o TCC

Fornece contexto amplo de comparação entre múltiplas metaheurísticas, incluindo ACO (implementado no projeto) e GWO/CSO (não implementados, mas relevantes como baselines futuros). A metodologia de 30 runs e análise estatística serve como referência para o design experimental do TCC. O estudo reforça a importância de comparar algoritmos sob condições padronizadas.

## Métodos e Abordagens

- 9 algoritmos: ACO, LA, CS, GWO, VPS, SSO, CSO, BA, ABC
- 30 runs independentes por algoritmo-instância
- Datasets TSPLIB: berlin52 (52 cidades), eil76 (76), pr1002 (1002)
- Métricas: média, desvio padrão, melhor solução, análise de convergência
- Framework unificado de implementação para comparabilidade

## Conexões

- [[TSP]] — problema-alvo
- [[comparative-studies]] — área temática
- [[ant-colony]] — ACO incluído e bem-sucedido
- [[particle-swarm]] — PSO análogo (não incluso na comparação)
- [[genetic-algorithms]] — GA análogo (não incluso)
- [[bio-inspired-optimization]]
- [[hossain2024comparison]] — estudo comparativo similar com divisão small/medium/large

## Notas e Insights

- ACO, GWO e CSO lideram em qualidade de solução — consistente com a literatura
- A ausência de PSO e GA é notável, dado que são metaheurísticas clássicas para TSP
- A cobertura de instâncias é limitada (apenas 3); dificulta generalização
- 30 runs por instância é uma prática robusta de experimentação
- Estudo não aborda tempo de execução como métrica primária
- A inclusão de algoritmos menos comuns (VPS, SSO) agrega valor ao benchmark
- Útil como referência de quais algoritmos são competitivos em TSP

## Citações-chave

> ACO, GWO and CSO present better balance between accuracy and robustness across all tested instances.

> The comparative analysis reveals that metaheuristic performance is highly instance-dependent, with no single algorithm dominating all problem sizes.
