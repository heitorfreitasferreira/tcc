---
title: 'Charting New Routes: Comparing Swarm-Based Approaches to the Traveling Salesman Problem'
authors:
- Wadi
- Ali Hassan Ahmed
- Umar
- Shahla Uthman
year: 2025
doi: 10.18280/ijcmem.130214
bibtex_key: wadi2025charting
bibtex-key: wadi2025charting
pdf: papers/pdfs/wadi2025charting.pdf
tags:
- area/tsp
- evidencia/referencia
- metodo/aco
- metodo/eho
- metodo/metaheuristic
- metodo/pso
- papel/comparativo
- status/lido-parcial
- tipo/paper
status: lido-parcial
rating: 3
type: paper
methods:
- aco
- eho
- metaheuristic
- pso
role: comparativo
reading_status: lido-parcial
validation_status: nao-validado
---

## PDF

![[wadi2025charting.pdf]]

## Resumo

Wadi e Umar comparam abordagens de enxame para resolver o [[tsp]], com foco em [[particle-swarm]], [[ant-colony]] e Elephant Herding Optimization. O artigo também usa Branch and Bound e Dynamic Programming como referências clássicas para instâncias pequenas. A avaliação considera tamanhos de 5 a 150 cidades e mede custo ótimo ou melhor custo, tempo de execução, uso de CPU e complexidade temporal. A conclusão do artigo é que EHO apresenta menores custos em grande parte dos cenários e escala melhor que métodos exatos, enquanto ACO oferece bom equilíbrio entre custo e tempo e PSO é rápido, mas sensível a parâmetros.

> [!warning] Leitura parcial
> A extração do PDF recuperou resumo, métodos e resultados, mas o texto apresenta problemas de OCR/revisão e possíveis inconsistências nas tabelas. Usar qualitativamente; conferir números antes de citar.

## Contribuições Principais

- Compara PSO, ACO e EHO no TSP com BB e DP como métodos clássicos.
- Discute escalabilidade de métodos exatos versus metaheurísticas.
- Avalia execução em diferentes números de cidades: 5, 10, 14, 19, 100 e 150.
- Inclui métricas de tempo de execução, CPU, complexidade e qualidade da solução.
- Aponta EHO como método de menor custo na maioria das instâncias do estudo.

## Relevância para o TCC

O artigo é relevante porque inclui dois dos métodos centrais do TCC: [[particle-swarm]] e [[ant-colony]]. Ele evidencia uma tensão importante para drones: métodos exatos podem servir como referência em instâncias pequenas, mas deixam de ser viáveis quando o número de pontos cresce. A comparação entre custo e tempo ajuda a enquadrar o TCC como avaliação de qualidade da rota e viabilidade computacional.

## Métodos e Abordagens

- Métodos exatos: Branch and Bound e Dynamic Programming.
- Métodos de enxame/metaheurísticos: ACO, PSO e EHO.
- ACO: construção probabilística de rotas com feromônio, visibilidade heurística, evaporação e depósito de feromônio.
- PSO: partículas representam tours; atualização orientada por `pBest` e `gBest`; adaptação discreta por trocas ou ajustes na ordem das cidades.
- EHO: população dividida em clãs; atualização guiada pelo melhor elefante/clã; substituição do pior indivíduo por rota aleatória.

## Conexões

- [[tsp]]
- [[particle-swarm]]
- [[ant-colony]]
- [[comparative-studies]]
- [[kennedy1995particle]]
- [[dorigo1996ant]]
- [[dorigo1997ant]]
- [[almufti2025comparative]]
- [[hossain2024comparison]]

## Notas e Insights

- O artigo é mais útil para discussão qualitativa sobre escalabilidade e trade-off qualidade/tempo do que para importar números diretamente.
- Há sinais de fragilidade editorial e possíveis inconsistências nos resultados; tabelas devem ser verificadas visualmente.
- A seção de trabalhos futuros é alinhada ao TCC: hibridização, busca local, ajuste adaptativo de parâmetros e aplicações em logística inteligente.

## Citações-chave

> “To solve the Traveling Salesman Problem (TSP), this research compares three swarm-based optimization algorithms: Particle Swarm Optimization (PSO), Ant Colony Optimization (ACO), and Elephant Herding Optimization (EHO).”

> “The algorithms' performance is assessed in this study based on execution time, scalability, and solution quality for a range of city sizes (5 to 150).”
