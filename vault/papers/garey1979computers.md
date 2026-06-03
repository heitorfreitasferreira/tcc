---
title: "Computers and Intractability: A Guide to the Theory of NP-Completeness"
authors: [Garey, Michael R., Johnson, David S.]
year: 1979
doi: ""
bibtex-key: garey1979computers
pdf: "papers/pdfs/garey1979computers.pdf"
tags: [complexity foundational]
status: lido
rating: 5
---

## PDF

![[garey1979computers.pdf]]

## Resumo

Livro clássico e referência definitiva sobre NP-completude. Apresenta a teoria da NP-completude de forma acessível: o que significa um problema ser NP-completo, como provar NP-completude, e uma enciclopédia de mais de 300 problemas NP-completos organizados por categoria (grafos, redes, particionamento, programação inteira, etc.). O TSP está entre os problemas originalmente catalogados como NP-difícil. O livro fornece a base teórica que justifica porque problemas como TSP (e suas variantes como rTSP, FSTSP) exigem heurísticas e metaheurísticas — não se conhece algoritmo polinomial para resolvê-los exatamente.

## Contribuições Principais

- Organização e catalogação de centenas de problemas NP-completos em uma referência unificada
- Metodologia clara para provar NP-completude por redução polinomial
- Estabelecimento do TSP como problema NP-difícil de referência

## Relevância para o TCC

Justificativa teórica fundamental para o uso de metaheurísticas (GA, PSO, ACO) em vez de métodos exatos. O rTSP (patrulha com drones) herda a NP-dificuldade do TSP clássico, tornando heurísticas a abordagem adequada.

## Métodos e Abordagens

- Redução polinomial entre problemas
- Classes de complexidade: P, NP, NP-completo, NP-difícil
- Teorema de Cook-Levin (SAT é NP-completo) como base

## Conexões

- [[lawler1985traveling]] — survey do TSP (reafirma NP-dificuldade)
- [[applegate2006traveling]] — estudo computacional do TSP (mostra que instâncias grandes são intratáveis exatamente)
- [[lin1973effective]] — LKH heurística eficiente para TSP (necessária exatamente por ser NP-difícil)
- [[murray2015flying]] — FSTSP (NP-difícil por generalizar TSP)
- [[tsp]]

## Notas e Insights

O TSP é listado no Apêndice do livro como problema ND22 (TSP direcionado) e ND23 (TSP não-direcionado). A NP-dificuldade do TSP se reduz ao problema do Ciclo Hamiltoniano. Esta referência é a mais citada em qualquer paper que mencione "NP-difícil" ou "NP-completo".

## Citações-chave

> The traveling salesman problem is NP-hard, meaning that no polynomial-time algorithm is known for its solution.
