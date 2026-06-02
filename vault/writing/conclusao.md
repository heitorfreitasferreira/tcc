---
tags: [writing, capitulo, conclusao]
---

# Conclusão — Scaffold

## Estrutura do Capítulo

### 5.1 Síntese dos Resultados
- ACO domina em qualidade (gap 0% em n≤12, ~2× melhor em n=100)
- GA oferece melhor custo-benefício (ótimo em n≤12 em ~50% runs, 75× mais rápido que ACO)
- PSO apresenta desempenho consistentemente inferior na variante TSP-SD-ATP

### 5.2 Contribuições
- Primeiro benchmark multi-método na variante TSP-SD-ATP
- Implementação Go reproduzível com tensor 3D
- 4605 execuções com análise estatística
- Demonstração de que ACO é particularmente adequado para problemas com dependência de sequência

### 5.3 Limitações
- Apenas 3 métodos (GA, PSO, ACO)
- Parâmetros fixos (sem tuning sistemático)
- Instâncias sintéticas em [-1,1]²
- 51 sementes × 30 instâncias (cobertura limitada)
- Brute-force apenas até n=14

### 5.4 Trabalhos Futuros
- Métodos adicionais: ABC, GWO, SA, ILS
- Tuning de parâmetros (grid search, irace)
- Instâncias baseadas em cenários reais de patrulha
- Extensão para múltiplos drones (mTSP)
- rTSP com recompensas variáveis
- Deep learning + ACO/GA (DeepACO, NeuFACO)

## Material de Apoio

- [[resultados]] — base para síntese
- [[analysis-methodology]] — limitações estatísticas
- [[deepaco2023]], [[neufaco2025]], [[gpaco2025]] — trabalhos futuros com neural ACO
- [[dellamico2021multiple]] — múltiplos drones
- [[pop2024comprehensive]] — GTSP (extensão para clusters)
