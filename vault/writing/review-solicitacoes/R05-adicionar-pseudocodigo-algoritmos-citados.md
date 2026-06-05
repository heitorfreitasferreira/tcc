---
title: "Adicionar pseudocódigo da busca exaustiva, lower bound e Hungarian ao apêndice"
tags:
  - tipo/revisao
  - status/aberto
review_id: "R05"
status: aberto
priority: "media"
source: reuniao
date_opened: "2026-06-05"
date_closed: ""
target_chapter: "proposta"
correction_layers: ["latex-macro", "agentico"]
evidence_layer: "codigo"
claim_ids: []
verified_by_script: ""
aliases: []
---

## Solicitação Original

> Adicione na seção de pseudocódigo os algoritmos como Hungarian e outros que você citou e não colocou.

## Análise Técnica

### Problema Identificado

O apêndice `ape_metodos/pseudocodigo.tex` contém pseudocódigo apenas para GA, PSO e ACO. Entretanto, o texto cita explicitamente outros algoritmos como parte da implementação — em particular o algoritmo de Heap (busca exaustiva), o algoritmo Hungarian e a redução 3D→2D do lower bound. Esses algoritmos são parte da proposta e devem ser documentados no mesmo apêndice.

**Algoritmos citados vs. documentados:**

| Algoritmo | Citado em | Tem pseudocódigo? | Tem fluxograma TikZ? |
|---|---|---|---|
| GA | `proposta.tex:27-31` | Sim (`ape_metodos/pseudocodigo.tex:6`) | Sim (`flowchart-ga.tex`) |
| PSO | `proposta.tex:33-44` | Sim (`ape_metodos/pseudocodigo.tex:42`) | Sim (`flowchart-pso.tex`) |
| ACO | `proposta.tex:46-57` | Sim (`ape_metodos/pseudocodigo.tex:80`) | Sim (`flowchart-aco.tex`) |
| **Busca Exaustiva (Heap)** | `proposta.tex:61` | **Não** | Sim (`flowchart-bruteforce.tex`) |
| **Lower Bound AP + Hungarian** | `proposta.tex:65`, `fundamentacao.tex:90` | **Não** | Sim (`flowchart-lowerbound.tex`) |
| **Redução 3D→2D** | `proposta.tex:65` | **Não** | Parcial (no flowchart do LB) |

**Implementação em Go (fonte da verdade para o pseudocódigo):**

| Função | Arquivo | Linha | Descrição |
|---|---|---|---|
| `hungarian(cost [][]float64) ([]int, float64)` | `src/optimization/lowerbound/hungarian.go` | 5 | Kuhn-Munkres para assignment problem |
| `reduce3Dto2D(g graph.Graph, n int) [][]float64` | `src/optimization/lowerbound/main.go` | 39 | Colapsa tensor 3D em matriz 2D por min sobre i |
| `generatePermutations(arr []int) <-chan []int` | `src/optimization/brute/main.go` | 55 | Heap's algorithm com channel-based generator |

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/optimization/lowerbound/hungarian.go:5` | Implementa Kuhn-Munkres com arrays 1-indexed (`u`, `v`, `p`, `way`) | Referência canônica para o pseudocódigo |
| `src/optimization/lowerbound/main.go:9-39` | Pipeline completo: reduce3Dto2D → hungarian → extractSequence | Referência canônica |
| `src/optimization/brute/main.go:20-79` | Otimização por busca exaustiva com channel de permutações | Referência canônica |
| `monografia/ape_metodos/pseudocodigo.tex` | Contém apenas GA, PSO, ACO (3 de 5 métodos) | Incompleto — faltam busca exaustiva e lower bound |
| `monografia/cap_proposta/proposta.tex:59-67` | Descreve busca exaustiva e lower bound em prosa | Texto correto, mas sem suporte de pseudocódigo |
| `monografia/figs/flowchart-bruteforce.tex` | Fluxograma da busca exaustiva (Heap + avaliação) | Existe, mas fluxograma ≠ pseudocódigo |
| `monografia/figs/flowchart-lowerbound.tex` | Fluxograma do lower bound (redução + Hungarian) | Existe, mas fluxograma ≠ pseudocódigo |

### Causa Raiz

O apêndice de pseudocódigo foi escrito inicialmente apenas para os três métodos meta-heurísticos (GA, PSO, ACO), provavelmente porque são os mais complexos. A busca exaustiva e o lower bound — apesar de serem métodos completos implementados e avaliados — foram omitidos por oversight. O texto em prosa os descreve, os fluxogramas TikZ os ilustram, mas falta o pseudocódigo algorítmico formal.

## Plano de Correção

> Adicionar duas novas seções ao apêndice `ape_metodos/pseudocodigo.tex`, totalizando 5 algoritmos documentados. O pseudocódigo deve refletir fielmente a implementação Go.

### Abordagem Preferida (mais determinística)

- [ ] **Camada 1 — Código:** Ler `src/optimization/brute/main.go` e `src/optimization/lowerbound/` para capturar a lógica exata.
- [ ] **Camada 2 — LaTeX-macro/agentico:** Escrever pseudocódigo da Busca Exaustiva e do Lower Bound AP + Hungarian em `ape_metodos/pseudocodigo.tex`.
- [ ] **Camada 3 — Verificação:** Compilar e inspecionar o PDF.

### Comandos e Passos

```bash
# 1. Adicionar seções ao apêndice de pseudocódigo em:
#    monografia/ape_metodos/pseudocodigo.tex
#    - Seção 4: Busca Exaustiva (Heap's algorithm + loop de avaliação)
#    - Seção 5: Lower Bound por Relaxação AP (redução 3D→2D + Hungarian)

# 2. Compilar a monografia
cd monografia && pdflatex main_ppgco_ufu.tex && pdflatex main_ppgco_ufu.tex
```

### Artefatos Afetados

- `monografia/ape_metodos/pseudocodigo.tex` — adicionar ~2 seções com `\begin{algorithm}...\end{algorithm}`
- `monografia/cap_proposta/proposta.tex` — opcional: adicionar `\autoref{alg:bruteforce}` e `\autoref{alg:lowerbound}` onde os métodos são descritos

## Verificação

### Critério de Aceite

- O apêndice de pseudocódigo contém 5 algoritmos: GA, PSO, ACO, Busca Exaustiva e Lower Bound AP.
- O pseudocódigo da busca exaustiva reflete `generatePermutations` (Heap) + loop de avaliação via tensor G.
- O pseudocódigo do lower bound reflete `reduce3Dto2D` + `hungarian` (Kuhn-Munkres).
- O Hungarian é apresentado como algoritmo auxiliar (possivelmente em subseção própria ou inline no lower bound).

### Script de Validação

```bash
# Verificar que o apêndice contém 5 ambientes algorithm
rg -c "\\\\begin\{algorithm\}" monografia/ape_metodos/pseudocodigo.tex
# Esperado: 5

# Verificar labels dos novos algoritmos
rg "\\\\label\{alg:" monografia/ape_metodos/pseudocodigo.tex
# Esperado: alg:ga, alg:pso, alg:aco, alg:bruteforce, alg:lowerbound (ou similar)
```

## Notas

Registrado em sessão de orientação em 2026-06-05.

O algoritmo Hungarian (Kuhn-Munkres) pode ser apresentado como pseudocódigo independente ou como sub-rotina dentro do lower bound. A implementação Go usa 1-indexed arrays internos com `u`, `v`, `p`, `way` — o pseudocódigo deve seguir essa estrutura. Os fluxogramas TikZ existentes (`flowchart-bruteforce.tex`, `flowchart-lowerbound.tex`) servem como referência de alto nível, mas o pseudocódigo deve ser mais detalhado e refletir a implementação real.

