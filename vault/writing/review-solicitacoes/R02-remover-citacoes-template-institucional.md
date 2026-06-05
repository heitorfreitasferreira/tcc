---
title: "Remover citações ao template institucional no corpo da monografia"
tags:
  - tipo/revisao
  - status/aberto
review_id: "R02"
status: aberto
priority: "media"
source: reuniao
date_opened: "2026-06-05"
date_closed: ""
target_chapter: "proposta"
correction_layers: ["latex-macro", "agentico"]
evidence_layer: "monografia"
claim_ids: []
verified_by_script: ""
aliases: []
---

## Solicitação Original

> No início do Capítulo 3 (Proposta), há uma parte do texto que cita o template institucional. A monografia deve apenas seguir o template sem citá-lo. Remover essa e outras citações semelhantes.

## Análise Técnica

### Problema Identificado

O texto da monografia contém citações explícitas ao template institucional (`ppgco.cls` / modelo ABNT) no corpo do capítulo e em materiais auxiliares, o que não é apropriado para uma monografia. O template deve ser usado como suporte de formatação, não como referência citada no conteúdo acadêmico.

**Ocorrências encontradas:**

| Arquivo | Linha | Trecho | Gravidade |
|---|---|---|---|
| `monografia/cap_proposta/proposta.tex` | 4 | "...Conforme a orientação do template institucional, a revisão bibliográfica fica concentrada..." | Principal — deve ser removido |
| `monografia/ape_sobre/sobre.tex` | 3-11 | Texto explicativo sobre origem do modelo USP/ABNT e créditos a Athila Quaresma Santos | Secundário — capítulo "Sobre" do template |
| `monografia/main_ppgco_ufu.tex` | 2 | Comentário `% adaptado de modeloABNT2.tex, v1.0 athila` | Menor — comentário LaTeX, não visível no PDF |

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/` (código) | Não aplicável | — |
| `src/data/` (dados) | Não aplicável | — |
| Literatura | Não aplicável | — |
| `vault/` | Não aplicável | — |
| `monografia/` (texto atual) | O parágrafo de abertura do capítulo 3 contém referência explícita ao template institucional | Texto incorreto conforme orientação do orientador |

### Causa Raiz

O texto do Capítulo 3 foi escrito a partir do template institucional que serve como esqueleto inicial. A sentença "Conforme a orientação do template institucional..." é um texto-guia (placeholder) do template que ainda não havia sido removido. Trata-se de resquício do esqueleto de template, não de conteúdo autoral.

## Plano de Correção

> Correção textual simples — camada `latex-macro` (edição direta no `.tex`).

### Abordagem Preferida (mais determinística)

- [ ] **Camada 1 — LaTeX-macro:** Editar `monografia/cap_proposta/proposta.tex` linha 4 para remover a referência ao template, reescrevendo o parágrafo de abertura do capítulo sem perder o conteúdo informativo.
- [ ] **Camada 2 — LaTeX-macro:** Verificar e limpar `monografia/ape_sobre/sobre.tex` — decidir se o capítulo "Sobre" deve ser mantido ou removido (é parte do template original, não conteúdo da monografia).
- [ ] **Camada 3 — LaTeX-macro:** Remover comentário de atribuição em `main_ppgco_ufu.tex:2` se desejado.

### Comandos e Passos

```bash
# 1. Editar proposta.tex: remover "Conforme a orientação do template institucional,"
#    e reescrever a abertura do capítulo

# 2. Verificar se ape_sobre/sobre.tex ainda é incluído no documento principal
grep -n "sobre" monografia/main_ppgco_ufu.tex

# 3. Compilar para verificar que o PDF gerado não contém referências ao template
cd monografia && pdflatex main_ppgco_ufu.tex
```

### Artefatos Afetados

- `monografia/cap_proposta/proposta.tex` (principal)
- `monografia/ape_sobre/sobre.tex` (secundário, se mantido)
- `monografia/main_ppgco_ufu.tex` (comentário)

## Verificação

### Critério de Aceite

- Nenhuma ocorrência de "template institucional", "modelo ABNT", "modeloABNT", "Athila" ou "orientação do template" no PDF compilado (excluindo comentários LaTeX).
- O parágrafo de abertura do Capítulo 3 mantém coerência e fluidez sem a referência removida.

### Script de Validação

```bash
# Verificar que não há mais referências ao template no corpo dos .tex
rg -n "template institucional|orientação do template|conforme o template" monografia/cap_proposta/
# Esperado: sem matches
```

## Notas

Registrado em sessão de orientação em 2026-06-05.

O arquivo `ape_sobre/sobre.tex` é parte do pacote `ppgco.cls` e tipicamente serve como exemplo de uso. Verificar com o orientador se esse capítulo deve ser removido do documento final ou apenas limpo das referências ao template.

