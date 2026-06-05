---
title: Exact Models for the Flying Sidekick Traveling Salesman Problem
authors:
- Dell'Amico, Mauro
- Montemanni, Roberto
- Novellani, Stefano
year: 2022
doi: 10.1111/itor.13030
bibtex_key: dellamico2022exact
bibtex-key: dellamico2022exact
pdf: papers/pdfs/dellamico2022exact.pdf
tags:
- area/drone-routing
- area/fstsp
- area/tsp
- evidencia/referencia
- metodo/branch-and-cut
- metodo/exact
- metodo/milp
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 4
type: paper
areas:
- drone-routing
- fstsp
methods:
- branch-and-cut
- exact
- milp
role: revisao
reading_status: resumo-lido
validation_status: nao-validado
pdf_status: integro
chapters:
- fundamentacao
claim_support: []
aliases:
- FSTSP exact models
---

## Resumo

Publicado no International Transactions in Operational Research (2022). Apresenta três formulações MILP aprimoradas para o FSTSP, reduzindo o número de restrições "big-M" em relação aos modelos da literatura. As novas formulações introduzem variáveis que contabilizam a presença do drone no caminhão (evitando voos infactíveis) e reduzem variáveis de temporização. Resolve pela primeira vez todas as instâncias benchmark de Murray e Chu (2015) com 10 clientes — para 14 delas, a otimalidade foi provada aqui pela primeira vez. Também resolve muitas instâncias com 20 clientes e mostra como adaptar os modelos para variantes do problema.

## Contribuições Principais

- Formulações MILP mais compactas e eficientes que as anteriores (menos big-M)
- Primeira prova de otimalidade para 14 instâncias benchmark de 10 clientes
- Resolução de instâncias com 20 clientes (antes só se resolvia até 12--13)
- Adaptação das formulações para múltiplas variantes do FSTSP da literatura

## Evidência / Resultado Relevante

- O artigo propõe três formulações MILP (`3IF`, `2IF` e `2IF-BC`) para FSTSP, reduzindo variáveis de temporização e restrições big-M.
- O summary registra que todas as 72 instâncias benchmark com 10 clientes foram resolvidas à otimalidade e que a escalabilidade degrada em instâncias com 20 clientes.
- A decomposição do objetivo em viagem do caminhão, tempos de lançamento/recolhimento e espera melhora a cota inferior inicial em relação a formulações anteriores.

## Limitações de Uso

- Contextualiza métodos exatos para FSTSP; não é evidência direta sobre TSP-SD-ATP nem sobre meta-heurísticas do projeto.
- O caso é caminhão + um drone; a formulação não cobre patrulha isolada com penalidade angular.
- Números de gap, tempo e quantidade de instâncias fechadas devem ser conferidos no PDF se forem usados no texto final.

## Relevância para o TCC

Citado na Seção 2.2 como referência de métodos exatos para drone routing. Complementa dellamico2021multiple (que trata múltiplos drones) ao focar no caso single-drone com formulações exatas aprimoradas. Relevante para contextualizar a complexidade do FSTSP e justificar o uso de meta-heurísticas.

## Conexões

- [[dellamico2021multiple]] — artigo irmão com múltiplos drones
- [[murray2015flying]] — FSTSP original
- [[agatz2018optimization]] — TSP-D
- [[tsp]]
