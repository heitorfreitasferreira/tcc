---
title: "Corrigir referências 'Equation' para 'Equação' no texto compilado"
tags:
  - tipo/revisao
  - status/aberto
review_id: "R03"
status: aberto
priority: "media"
source: reuniao
date_opened: "2026-06-05"
date_closed: ""
target_chapter: "todos"
correction_layers: ["latex-macro"]
evidence_layer: "monografia"
claim_ids: []
verified_by_script: ""
aliases: []
---

## Solicitação Original

> No 3.1 (Capítulo 3, Seção 1), a referência para a Equação 1 ficou em inglês ("Equation 1") em vez de português ("Equação 1"). Corrigir essa e todas as outras ocorrências semelhantes no texto.

## Análise Técnica

### Problema Identificado

Quando o comando `\autoref{eq:custo}` (usado em `cap_proposta/proposta.tex:10`) é compilado, o PDF exibe "Equation 1" em vez de "Equação 1". Isso ocorre porque a classe `ppgco.cls` define `\*autorefname` em português para Figuras, Tabelas, Capítulos, Seções etc. (linhas 528--542), mas **omite `\equationautorefname`**. Sem essa definição, o pacote `hyperref` usa o fallback padrão em inglês.

**Equações rotuladas no texto:**

| Arquivo | Linha | Label | Referenciada via `\autoref`? |
|---|---|---|---|
| `cap_fundamentacao/fundamentacao.tex` | 24 | `eq:custo` | Não no cap. 2 |
| `cap_fundamentacao/fundamentacao.tex` | 40 | `eq:objetivo` | Não localizada |
| `cap_fundamentacao/fundamentacao.tex` | 87 | `eq:reducao-ap` | Não localizada |
| `cap_proposta/proposta.tex` | 41 | `eq:pso` | Não localizada |
| `cap_proposta/proposta.tex` | 52 | `eq:aco` | Não localizada |
| `cap_proposta/proposta.tex` | 10 | `eq:custo` | **Sim — `\autoref{eq:custo}`** |

A ocorrência confirmada está em `proposta.tex:10`. Outras equações podem vir a ser referenciadas com `\autoref` futuramente, e todas herdariam o mesmo problema.

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/` (código) | Não aplicável | — |
| `src/data/` (dados) | Não aplicável | — |
| Literatura | Não aplicável | — |
| `vault/` | Não aplicável | — |
| `monografia/` (texto atual) | `proposta.tex:10` usa `\autoref{eq:custo}` → PDF mostra "Equation 1" | Sim — idioma incorreto |
| `monografia/ppgco.cls:528-542` | Define `\*autorefname` para Figura, Tabela, Capítulo, Seção, etc., mas **falta `equation`** | Sim — omissão na classe |

### Causa Raiz

A classe `ppgco.cls` (fornecida pelo PPGCO/UFU) contém traduções de `autorefname` para os tipos mais comuns de referência cruzada, mas os mantenedores da classe omitiram `\equationautorefname`. Como o `hyperref` não encontra a definição em português, usa o fallback padrão em inglês.

## Plano de Correção

> Correção determinística simples — adicionar uma linha ao `ppgco.cls` e recompilar.

### Abordagem Preferida (mais determinística)

- [ ] **Camada 1 — LaTeX-macro:** Adicionar `\renewcommand*\equationautorefname{Equação}` em `ppgco.cls` (após linha 542) ou no preâmbulo de `main_ppgco_ufu.tex`.
- [ ] **Camada 2 — Verificação:** Compilar e inspecionar o PDF para confirmar que `\autoref{eq:custo}` produz "Equação 1".

### Comandos e Passos

```bash
# 1. Adicionar a definição faltante em ppgco.cls (após linha 542)
#    \renewcommand*\equationautorefname{Equação}

# 2. Compilar a monografia
cd monografia && pdflatex main_ppgco_ufu.tex && pdflatex main_ppgco_ufu.tex

# 3. Verificar que "Equation" não aparece mais no PDF
pdftotext main_ppgco_ufu.pdf - | rg -i "equation"
# Esperado: sem matches
```

### Artefatos Afetados

- `monografia/ppgco.cls` (adicionar 1 linha)
- Alternativamente: `monografia/main_ppgco_ufu.tex` (se preferir não alterar a classe)

## Verificação

### Critério de Aceite

- O PDF compilado não contém a palavra "Equation" (em inglês) em nenhuma referência cruzada.
- `\autoref{eq:custo}` em `proposta.tex:10` produz "Equação 1" no PDF.
- Todas as demais `\autoref` produzem texto em português (já funcionavam).

### Script de Validação

```bash
# Verificar que não há "Equation" (maiúsculo) no texto do PDF compilado
pdftotext monografia/main_ppgco_ufu.pdf - | rg -c "Equation"
# Esperado: 0
```

## Notas

Registrado em sessão de orientação em 2026-06-05.

A classe `ppgco.cls` é fornecida pelo PPGCO/UFU. Se a alteração na classe for indesejada (para preservar compatibilidade com atualizações futuras), a linha `\renewcommand*\equationautorefname{Equação}` pode ser colocada no preâmbulo de `main_ppgco_ufu.tex` com o mesmo efeito.

