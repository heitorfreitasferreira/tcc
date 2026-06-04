---
tags: [writing, capitulo, experimentos]
created: 2026-06-02
updated: 2026-06-02
---

# Experimentos — Scaffold

## Estrutura do Capítulo

### 4.1 Configuração Experimental
- Instâncias: 30 (10-100 nós, 3 variantes cada)
- Métodos: GA, PSO, ACO, lower bound AP e brute-force (`10a..15c`)
- Parâmetros: pop=100, iter=100, 51 sementes (0-50)
- Hardware, implementação, reprodução

### 4.2 Baseline Ótima (Brute-force)
- Tabela de ótimos para `10a..15c` ([[resultados]])
- Validação dos métodos contra o ótimo

### 4.3 Qualidade da Solução
- Gap vs ótimo em instâncias pequenas
- Comparação em instâncias grandes (50, 100 nós)
- Rankings médios e significância validada por Friedman/Iman-Davenport, Nemenyi e Wilcoxon/Holm

### 4.4 Tempo Computacional
- Tempo médio por método e tamanho
- Trade-off qualidade × tempo
- Escalabilidade

### 4.5 Análise Estatística
- Friedman/Iman-Davenport sobre medianas por instância
- Pós-teste Nemenyi e Wilcoxon/Holm com os valores validados em [[analysis-methodology]]
- Não usar ANOVA/Tukey como análise principal para comparar ranks por instância

### 4.6 Discussão
- Por que ACO domina em qualidade?
- Por que PSO tem desempenho fraco?
- Implicações para o cenário de patrulha com drones

## Tabelas e Figuras Planejadas

1. Tabela de makespan ótimo (brute-force) — 18 instâncias
2. Tabela de gap médio por método (`10a..15c`)
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
- [[auditoria-codigo-dados-vault]] — cobertura auditada
- [[auditoria-script-analise-estatistica]] — bloqueio estatístico atual
- [[auditoria-codificacao-metodos]] — limitação sobre representação dos métodos
- [[chandra2022comparative]] — ANOVA+Tukey para TSP
