---
title: "Substituir referências a pastas do repositório nas legendas das figuras"
tags:
  - tipo/revisao
  - status/aberto
review_id: "R04"
status: aberto
priority: "media"
source: reuniao
date_opened: "2026-06-05"
date_closed: ""
target_chapter: "experimentos"
correction_layers: ["latex-macro"]
evidence_layer: "monografia"
claim_ids: []
verified_by_script: ""
aliases: []
---

## Solicitação Original

> Não gostei que nas legendas das figuras, como em "Figura 3 – Gap em relação à busca exaustiva nas instâncias com ótimo conhecido. Fonte: dados de 'src/data/results/'.", faz referência direta à pasta do repositório. Faça de uma forma mais elegante.

## Análise Técnica

### Problema Identificado

Nove legendas de figuras no capítulo de experimentos (`cap_experimentos/experimentos.tex`) referenciam caminhos do repositório como fonte dos dados. Esses caminhos são internos ao projeto e inadequados para uma monografia publicada. O padrão ABNT para dados gerados pelo próprio autor é "Fonte: elaborado pelo autor" — convenção já adotada nas figuras dos capítulos de fundamentação e proposta.

**Ocorrências afetadas:**

| Linha | Figura | Trecho atual |
|---|---|---|
| 60 | `fig:gap-vs-bf` | `Fonte: dados de \`src/data/results/'.` |
| 67 | `fig:taxa-acerto` | `Fonte: dados de \`src/data/results/'.` |
| 86 | `fig:heatmap-makespan` | `Fonte: dados de \`src/data/results/'.` |
| 120 | `fig:runtime` | `Fonte: dados de \`src/data/results/timing/'.` |
| 127 | `fig:tradeoff` | `Fonte: dados de \`src/data/results/'.` |
| 140 | `fig:estabilidade` | `Fonte: dados de \`src/data/results/'.` |
| 147 | `fig:convergencia` | `Fonte: dados de evolução em \`src/data/results/evolution/'.` |
| 160 | `fig:cd-diagram` | `Fonte: \`scripts/analise-estatistica.py'.` |

**Adicional (comentários LaTeX, invisíveis no PDF):** Arquivos em `monografia/generated/` contêm `% Source: src/data/results via scripts/gerar-metricas-monografia.py.` como comentário. Podem ser removidos ou mantidos como documentação interna.

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/` (código) | Não aplicável | — |
| `src/data/` (dados) | Os dados de fato residem em `src/data/results/` | Não — a localização está correta, o problema é expô-la na monografia |
| Literatura | Não aplicável | — |
| `vault/` | Não aplicável | — |
| `monografia/` (texto atual) | 8 legendas com caminhos de repositório + 1 referência a script Python | Sim — inadequado para publicação |
| `monografia/cap_fundamentacao/fundamentacao.tex:32` | Já usa "Fonte: elaborado pelo autor" | Padrão correto a seguir |
| `monografia/cap_proposta/proposta.tex:15` | Já usa "Fonte: elaborado pelo autor" | Padrão correto a seguir |

### Causa Raiz

As legendas foram escritas durante a fase de geração de figuras, quando a proveniência exata dos dados era informação útil para o autor. O texto jamais foi revisado para o estilo de publicação — manteve-se a anotação de trabalho como se fosse a legenda definitiva.

## Plano de Correção

> Substituição textual simples em `experimentos.tex`. Nenhum código, dado ou script precisa ser alterado.

### Abordagem Preferida (mais determinística)

- [ ] **Camada 1 — LaTeX-macro:** Substituir todas as 8 legendas com caminhos de repositório por "Fonte: elaborado pelo autor".
- [ ] **Camada 2 — LaTeX-macro (opcional):** Remover comentários `% Source: src/data/results...` dos arquivos em `generated/` se desejado.

### Comandos e Passos

```bash
# 1. Editar experimentos.tex: substituir cada caption com path por "Fonte: elaborado pelo autor."

# 2. Compilar e verificar
cd monografia && pdflatex main_ppgco_ufu.tex && pdflatex main_ppgco_ufu.tex
```

### Artefatos Afetados

- `monografia/cap_experimentos/experimentos.tex` (8 legendas)
- `monografia/generated/tables/*.tex` (comentários opcionais)

## Verificação

### Critério de Aceite

- Nenhuma legenda de figura/tabela no PDF contém caminhos de diretório (`src/`, `scripts/`, `data/`).
- Todas as figuras cujos dados foram gerados pelo autor usam "Fonte: elaborado pelo autor".
- Figuras que citam fontes externas (ex.: reproduzidas de artigos) mantêm a citação bibliográfica apropriada.

### Script de Validação

```bash
# Verificar que não há paths de repositório nas legendas do PDF
pdftotext monografia/main_ppgco_ufu.pdf - | rg "src/data|src/data/results|scripts/analise"
# Esperado: sem matches
```

## Notas

Registrado em sessão de orientação em 2026-06-05.

A opção "Fonte: elaborado pelo autor" é a mais adequada porque todos os dados e figuras foram gerados pelo próprio autor a partir dos experimentos. Se o orientador preferir uma formulação diferente (ex.: "Fonte: o autor", "Fonte: dados da pesquisa"), basta aplicar a mesma substituição com o texto desejado.

