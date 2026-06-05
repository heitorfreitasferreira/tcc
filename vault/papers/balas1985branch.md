---
title: Branch and Bound Methods for the Traveling Salesman Problem
authors:
- Balas
- Egon
- Toth
- Paolo
year: 1983
doi: ""
bibtex_key: balas1985branch
bibtex-key: balas1985branch
pdf: papers/pdfs/balas1985branch.pdf
tags:
- area/atsp
- area/tsp
- evidencia/referencia
- metodo/branch-and-bound
- metodo/lower-bound
- papel/revisao
- status/resumo-lido
- tipo/paper
status: disponível
rating: 5
classificacao: recuperado
type: paper
areas:
- atsp
methods:
- branch-and-bound
- lower-bound
role: revisao
reading_status: resumo-lido
validation_status: requer-validacao
pdf_status: integro
chapters:
- fundamentacao
claim_support: []
aliases:
- Balas e Toth 1983
- Branch and Bound Methods for the TSP
note_bibkey: "O PDF local é o relatório técnico MSRR 488 de 1983 da Carnegie-Mellon University. A chave histórica balas1985branch foi preservada para compatibilidade com wikilinks e referências antigas."
---

## PDF

![[balas1985branch.pdf]]

## Tese Central

Balas e Toth revisam métodos branch-and-bound para o TSP e organizam o estado da arte por tipo de relaxação: problema de atribuição, 1-tree lagrangiana e problema de atribuição com função lagrangiana.

## Resumo

O relatório técnico revisa métodos enumerativos para resolver o TSP de forma exata. A discussão separa os algoritmos conforme a relaxação usada para produzir cotas inferiores, descreve regras de ramificação e seleção de subproblemas, e compara códigos computacionais do período.

A revisão distingue o comportamento do TSP assimétrico e simétrico. Para ATSP, a relaxação de atribuição com custo original já produz cotas fortes em instâncias aleatórias; para STSP, a relaxação 1-tree com multiplicadores lagrangianos, associada a Held e Karp, torna-se mais adequada. O texto também apresenta evidência empírica sobre crescimento computacional e sobre redução de arcos em algoritmos branch-and-bound.

## Contribuições Principais

- Organização das principais famílias de relaxações usadas em branch-and-bound para TSP.
- Comparação entre relaxação AP, 1-tree lagrangiana e AP lagrangiana.
- Discussão de regras de ramificação, seleção de subproblemas e trade-off entre memória e número de nós.
- Síntese computacional de códigos exatos para TSP simétrico e assimétrico.
- Registro de que cotas fortes reduzem drasticamente a árvore de busca, mas não eliminam o custo exponencial no pior caso.

## Relevância para o TCC

A nota sustenta a seção de fundamentação sobre métodos exatos e lower bounds. O projeto implementa um lower bound baseado em relaxação de atribuição, não um branch-and-bound completo; por isso, Balas e Toth devem ser usados para contextualizar por que relaxações de atribuição e 1-tree aparecem como componentes de métodos exatos, não para afirmar que o repositório resolve o TSP exatamente em instâncias grandes.

## Métodos e Abordagens

- Relaxação por problema de atribuição com eliminação posterior de sub-rotas.
- Relaxação 1-tree com função lagrangiana para o TSP simétrico.
- Relaxação AP lagrangiana para o TSP assimétrico.
- Regras de branching por sub-rotas, arcos e cortes condicionais.
- Estratégias de exploração depth-first e best-bound.

## Evidência / Resultado Relevante

- O PDF local identifica o documento como relatório técnico de 1983 da Carnegie-Mellon University.
- A revisão reporta que cotas AP são fortes para ATSP aleatório, enquanto a 1-tree lagrangiana é a referência para STSP.
- O texto registra que a escolha entre depth-first e best-bound troca memória por redução de nós na árvore.

## Limitações de Uso

- A versão validada localmente é relatório técnico de 1983; a chave `balas1985branch` é histórica e não deve ser interpretada como confirmação de ano 1985.
- A nota não deve ser usada para sustentar desempenho moderno de solvers exatos sem literatura posterior.
- A discussão é sobre solução exata por enumeração implícita, não sobre meta-heurísticas bioinspiradas.

## Conexões

- Fundamenta: [[lower-bounds]], [[tsp]]
- Relacionado a: [[heldkarp1971traveling]], [[lysgaard1999cluster]]
- Usado em capítulo: [[fundamentacao]]
