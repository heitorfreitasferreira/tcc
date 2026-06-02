---
tags: [writing, capitulo, introducao]
---

# Introdução — Scaffold

## Estrutura do Capítulo

1. **Contexto e Motivação**
   - Patrulha com drones: vigilância, inspeção, monitoramento
   - Problema de roteamento: visitar POIs eficientemente
   - TSP como modelo matemático subjacente

2. **Problema de Pesquisa**
   - Formulação TSP-SD-ATP: TSP com penalidades angulares ([[problem-formulation]])
   - Por que o ângulo de virada importa para drones?
   - Instâncias aleatórias em [-1,1]², tensor 3D pré-computado

3. **Lacuna na Literatura**
   - Benchmarks existentes usam TSP clássico (matriz 2D simétrica)
   - Nenhum estudo comparativo sistemático na variante TSP-SD-ATP
   - Três metaheurísticas canônicas nunca comparadas nesta variante

4. **Objetivos**
   - Implementar GA, PSO, ACO para TSP-SD-ATP
   - Comparar qualidade e tempo em 30 instâncias, 51 sementes
   - Determinar qual método oferece melhor trade-off

5. **Contribuições**
   - Primeiro benchmark multi-método na variante TSP-SD-ATP
   - Implementação Go reproduzível com tensor 3D
   - Análise estatística com 4605 execuções

6. **Organização da Monografia**
   - Capítulo 2: Fundamentação teórica
   - Capítulo 3: Proposta
   - Capítulo 4: Experimentos
   - Capítulo 5: Conclusão

## Material de Apoio

- [[rajan2022routing]] — patrulha com UAV (motivação direta)
- [[murray2015flying]] — FSTSP (contexto drone routing)
- [[tsp-variants]] — classificação da variante
- [[results]] — dados experimentais (para mencionar na introdução)
