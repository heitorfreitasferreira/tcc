---
tags: [writing, capitulo, introducao]
status: atualizado-pos-p7
created: 2026-06-02
updated: 2026-06-02
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
   - Não foram identificados, na revisão realizada, estudos que comparem sistematicamente GA, PSO e ACO na variante TSP-SD-ATP
   - Três metaheurísticas canônicas ainda precisam ser avaliadas nessa formulação com dependência angular de sequência

4. **Objetivos**
   - Implementar GA, PSO, ACO para TSP-SD-ATP
   - Comparar qualidade e tempo em 30 instâncias, 51 sementes
   - Analisar trade-off entre qualidade de solução e tempo computacional

5. **Contribuições**
   - Benchmark multi-método na variante TSP-SD-ATP dentro da revisão realizada
   - Implementação Go reproduzível com tensor 3D
   - Base experimental com 4638 summaries, 51 sementes por metaheurística e 30 instâncias
   - Comparação com busca exaustiva em 18 instâncias e lower bound AP em 30 instâncias

6. **Organização da Monografia**
   - Capítulo 2: Fundamentação teórica
   - Capítulo 3: Proposta
   - Capítulo 4: Experimentos
   - Capítulo 5: Conclusão

## Material de Apoio

- [[rajan2022routing]] — patrulha com UAV (motivação direta)
- [[murray2015flying]] — FSTSP (contexto drone routing)
- [[tsp-variants]] — classificação da variante
- [[resultados]] — dados experimentais (para mencionar na introdução)
- [[claim-evidence-matrix]] — claims liberados e bloqueados
