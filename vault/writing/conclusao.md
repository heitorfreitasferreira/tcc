---
tags: [writing, capitulo, conclusao]
status: atualizado-pos-p4
created: 2026-06-02
updated: 2026-06-02
---

# Conclusão — Scaffold

## Estrutura do Capítulo

### 5.1 Síntese dos Resultados
- ACO apresentou melhor qualidade descritiva nas instâncias avaliadas, tanto nas 18 instâncias com brute-force quanto nas instâncias grandes.
- GA apresentou melhor tempo computacional, especialmente em n=100, mas com makespan médio pior que ACO.
- PSO apresentou desempenho inferior nesta implementação com random keys e parâmetros fixos.
- O lower bound AP foi válido, mas frouxo: gap médio 51.36% vs brute-force nas 18 instâncias auditadas.

### 5.2 Contribuições
- Benchmark multi-método na variante TSP-SD-ATP dentro da revisão realizada
- Implementação Go reproduzível com tensor 3D
- 4638 summaries com resultados, evolução e timing estruturados
- Comparação com busca exaustiva em `10a..15c` e lower bound AP nas 30 instâncias
- Evidência estatística de diferença entre ACO, GA e PSO nas configurações avaliadas, com ACO em melhor rank médio de makespan e custo computacional elevado

### 5.3 Limitações
- Apenas 3 métodos (GA, PSO, ACO)
- Parâmetros fixos (sem tuning sistemático)
- Instâncias sintéticas em [-1,1]²
- 51 sementes × 30 instâncias (cobertura limitada)
- Brute-force apenas nas 18 instâncias `10a..15c`
- Significância estatística restrita às implementações, parâmetros e instâncias avaliadas

### 5.4 Trabalhos Futuros
- Métodos adicionais: ABC, GWO, SA, ILS
- Tuning de parâmetros e análise de sensibilidade ([[auditoria-hiperparametros]])
- Instâncias baseadas em cenários reais de patrulha
- Extensão para múltiplos drones (mTSP)
- Variantes do TSP-SD-ATP com prioridades/recompensas variáveis para POIs
- Deep learning + ACO/GA (DeepACO, NeuFACO)

## Material de Apoio

- [[resultados]] — base para síntese
- [[analysis-methodology]] — limitações estatísticas
- [[deepaco2023]], [[neufaco2025]], [[gpaco2025]] — trabalhos futuros com neural ACO
- [[dellamico2021multiple]] — múltiplos drones
- [[pop2024comprehensive]] — GTSP (extensão para clusters)
