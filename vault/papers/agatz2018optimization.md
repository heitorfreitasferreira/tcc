---
title: "Optimization Approaches for the Traveling Salesman Problem with Drone"
authors: [Agatz, Niels]
year: 2018
doi: "10.1287/trsc.2017.0791"
bibtex-key: agatz2018optimization
tags: [tsp drone metaheuristic]
status: lido
rating: 5
pdf: "papers/pdfs/agatz2018optimization.pdf"
---

## PDF

![[agatz2018optimization.pdf]]

## Resumo

Propõe e formaliza o TSP-D (Travelling Salesman Problem with Drone), variante do TSP onde um veículo terrestre (caminhão) e um drone realizam entregas em cooperação. O problema difere do FSTSP de Murray e Chu (2015) ao permitir que o drone realize múltiplas entregas por lançamento e que o caminhão possa atender clientes enquanto o drone está em voo. São apresentadas formulações de programação inteira mista (MILP), heurísticas baseadas em busca local e programação dinâmica, além de experimentos computacionais em instâncias derivadas de benchmarks TSPLIB. O trabalho é referência obrigatória na literatura de otimização de entregas com drones, estabelecendo o TSP-D como problema de pesquisa ativo.

## Contribuições Principais

- Definição formal do TSP-D (TSP with Drone) com múltiplas entregas por voo do drone
- Duas formulações MILP para o TSP-D
- Heurística de roteirização e reassinalmento com busca local
- Abordagem baseada em programação dinâmica para partição de clientes entre caminhão e drone
- Experimentos em instâncias TSPLIB de até 100 clientes, demonstrando reduções de 10-40% no tempo de entrega comparado ao caminhão isolado

## Relevância para o TCC

Estabelece o TSP-D como variante central para entregas com drones, diretamente relacionada ao cenário de patrulha com drones do TCC. A modelagem de otimização cooperativa entre veículos terrestres e aéreos informa o design do problema rTSP usado no TCC. Suas heurísticas (busca local + programação dinâmica) são comparáveis às metaheurísticas bio-inspiradas (GA, PSO, ACO) implementadas no código do TCC.

## Métodos e Abordagens

- Programação Inteira Mista (MILP) com Gurobi
- Heurística de roteirização: constrói rota TSP para o caminhão e reassinala clientes para o drone via busca local
- Programação dinâmica para particionamento ótimo entre caminhão e drone
- Instâncias derivadas de TSPLIB (10–100 clientes)
- Métricas: tempo total de entrega (makespan), redução percentual vs. TSP clássico

## Conexões

- [[murray2015flying]] — FSTSP (problema predecessor que inspirou o TSP-D)
- [[dellamico2021multiple]] — extensão para múltiplos drones
- [[dellamico2022exact]] — modelos exatos para FSTSP
- [[freitas2020vns]] — VNS para FSTSP
- [[rajan2022routing]] — roteamento estocástico para patrulha UAV
- [[drone-routing]]
- [[tsp]]
- [[bio-inspired-optimization]]

## Notas e Insights

- O TSP-D difere do FSTSP na assunção de múltiplas entregas do drone por lançamento vs. uma única entrega no FSTSP
- Reduções de 10-40% no makespan vs. truck-only, dependendo da densidade de clientes elegíveis e velocidade do drone
- A heurística de reassinalmento é simples mas eficaz; metaheurísticas mais sofisticadas (GA/PSO/ACO) podem oferecer melhorias adicionais
- Tradeoff entre qualidade da solução do TSP base (LK vs. vizinho mais próximo) e desempenho global da heurística

## Citações-chave

> "Our computational study indicates that using a drone can reduce the total operational time by 10% to 40% compared to a traditional truck-only operation."
