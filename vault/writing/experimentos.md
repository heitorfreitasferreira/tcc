---
tags: [writing, capitulo, experimentos]
---

# Experimentos — Scaffold

## Estrutura do Capítulo

### 4.1 Configuração Experimental
- Instâncias: 30 (10-100 nós, 3 variantes cada)
- Métodos: GA, PSO, ACO, Brute-force (n≤14)
- Parâmetros: pop=100, iter=100, 51 sementes (0-50)
- Hardware, implementação, reprodução

### 4.2 Baseline Ótima (Brute-force)
- Tabela de ótimos para 10 ≤ n ≤ 14 ([[resultados]])
- Validação dos métodos contra o ótimo

### 4.3 Qualidade da Solução
- Gap vs ótimo em instâncias pequenas
- Comparação em instâncias grandes (50, 100 nós)
- Rankings e diferenças significativas

### 4.4 Tempo Computacional
- Tempo médio por método e tamanho
- Trade-off qualidade × tempo
- Escalabilidade

### 4.5 Análise Estatística
- Testes de normalidade
- ANOVA / Kruskal-Wallis
- Post-hoc Tukey / Dunn

### 4.6 Discussão
- Por que ACO domina em qualidade?
- Por que PSO tem desempenho fraco?
- Implicações para o cenário de patrulha com drones

## Tabelas e Figuras Planejadas

1. Tabela de makespan ótimo (brute-force) — 15 instâncias
2. Tabela de gap médio por método (n=10..14)
3. Heatmap de makespan por instância×método (todas as 30)
4. Boxplot de gap vs ótimo (instâncias pequenas)
5. Curvas de convergência (evolution data)
6. Scatter qualidade × tempo (trade-off)
7. Tabela de tempo médio por método e tamanho
8. Mapa das melhores rotas (visualização geográfica)

## Material de Apoio

- [[resultados]] — dados consolidados
- [[analysis-methodology]] — protocolo estatístico
- [[experiment-pipeline]] — configuração experimental
- [[comparative-studies]] — referência metodológica
- [[chandra2022comparative]] — ANOVA+Tukey para TSP
