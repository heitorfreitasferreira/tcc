---
tags:
- area/drone-routing
- tipo/area
created: 2026-06-02
updated: 2026-06-02
type: area
---

# Roteamento de Drones

## Definição
Problemas de otimização de rotas para veículos aéreos não tripulados (VANTs/drones), incluindo patrulha, entrega e monitoramento.

## Problemas Relacionados
- **FSTSP (Flying Sidekick TSP)**: drone assiste caminhão em entregas ([[murray2015flying]])
- **TSP-D (TSP with Drone)**: variante com drone e veículo terrestre ([[agatz2018optimization]])
- **rTSP**: apelido interno para a variante TSP-SD-ATP usada neste TCC; ver [[problem-formulation]]
- **Patrulha com VANT**: roteamento para vigilância/perímetro com horizonte finito ([[rajan2022routing]])

## Papers Relacionados
- [[murray2015flying]] — FSTSP original
- [[agatz2018optimization]] — TSP-D
- [[freitas2020vns]] — VNS para FSTSP
- [[dellamico2021multiple]] — FSTSP com múltiplos drones
- [[dellamico2022exact]] — modelos exatos para FSTSP
- [[ahmed2024receding]] — planejamento de caminho UAV
- [[rajan2022routing]] — roteamento estocástico para patrulha UAV

## Implementação no Projeto

Ver [[problem-formulation]] para a formulação do problema de patrulha com drones implementada em Go.
