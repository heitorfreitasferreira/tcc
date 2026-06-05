---
title: Statistical Comparisons of Classifiers over Multiple Data Sets
authors:
- Demšar
- Janez
year: 2006
doi: ''
bibtex_key: demsar2006statistical
bibtex-key: demsar2006statistical
pdf: papers/pdfs/demsar2006statistical.pdf
tags:
- evidencia/estatistica
- evidencia/metodologia
- evidencia/referencia
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 5
type: paper
role: revisao
reading_status: resumo-lido
validation_status: nao-validado
---

## Resumo

Publicado no Journal of Machine Learning Research (2006). Artigo metodológico canônico para comparação estatística de múltiplos algoritmos sobre múltiplas bases de dados. Revisa a prática corrente e examina teórica e empiricamente testes adequados. Recomenda um conjunto de testes não-paramétricos simples, seguros e robustos: Wilcoxon signed-ranks para comparação de dois classificadores, e o teste de Friedman com pós-testes correspondentes (Nemenyi, Holm, Bergmann-Hommel) para comparação de mais de dois. Introduz os diagramas CD (*critical difference*) para apresentação visual dos resultados.

## Contribuições Principais

- Recomendação formal do teste de Friedman + pós-teste de Nemenyi como protocolo padrão para comparação de múltiplos algoritmos
- Introdução dos diagramas CD como forma compacta de visualizar diferenças significativas entre ranqueamentos
- Análise de poder estatístico dos testes comparados (Wilcoxon, Friedman, ANOVA com post-hoc)
- Demonstração de que Friedman é mais apropriado que ANOVA para dados de ranqueamento em ML
- Referência com >23.000 citações — padrão *de facto* em aprendizado de máquina e otimização

## Relevância para o TCC

Referência metodológica central. O protocolo estatístico da monografia (Friedman + Nemenyi + análise pareada complementar) segue exatamente a recomendação deste artigo. Citado na Seção 2.9 (Estudos Comparativos) e na Seção 4.2 (Resultados). Justifica a escolha de testes não-paramétricos que não assumem normalidade, adequados para comparar GA, PSO e ACO sobre 30 instâncias com múltiplas sementes.

## Conexões

- [[claim-evidence-matrix]] — suporta claims sobre significância estatística
- [[auditoria-hiperparametros]] — limitação de parâmetros fixos
- [[tsp]]
