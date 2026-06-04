#!/usr/bin/env bash
# Deterministic scaffold for advisor review requests.
# Creates Rxx note with minimal YAML + template body.
set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
REVIEWS_DIR="$REPO_DIR/vault/writing/review-solicitacoes"
TEMPLATE="$REPO_DIR/vault/templates/review-request.md"

usage() {
  cat <<EOF
Usage: $(basename "$0") <descricao>

Creates a new Rxx review request scaffold.
  <descricao>  Short description (kebab-case, e.g. "ajustar-introducao-generica")

The script:
  1. Finds the next available Rxx number
  2. Creates vault/writing/review-solicitacoes/R<num>-<descricao>.md
  3. Fills YAML with today's date, status=aberto, priority=media, source=reuniao
  4. Includes full template body for later agent analysis

Outputs the created filename on stdout.
EOF
  exit 0
}

if [[ $# -lt 1 || "$1" == "--help" ]]; then
  usage
fi

DESCRIPTION="$1"
if [[ -z "$DESCRIPTION" ]]; then
  echo "ERROR: descricao vazia" >&2
  exit 1
fi

# Sanitize: lowercase, replace spaces/special chars with hyphens, strip leading/trailing hyphens
SAFE_DESC=$(echo "$DESCRIPTION" \
  | tr '[:upper:]' '[:lower:]' \
  | sed 's/[^a-zA-Z0-9 ]/ /g' \
  | tr -s ' ' \
  | sed 's/ /-/g' \
  | sed 's/--*/-/g' \
  | sed 's/^-//;s/-$//')
# Truncate to 50 chars max for filenames
if [[ ${#SAFE_DESC} -gt 50 ]]; then
  SAFE_DESC="${SAFE_DESC:0:50}"
  SAFE_DESC="${SAFE_DESC%-}"
fi

if [[ -z "$SAFE_DESC" ]]; then
  echo "ERROR: descricao invalida apos sanitizacao" >&2
  exit 1
fi

# Find next available Rxx number
NEXT_NUM=1
while true; do
  PREFIX=$(printf "R%02d" "$NEXT_NUM")
  existing=$(find "$REVIEWS_DIR" -maxdepth 1 -name "${PREFIX}*.md" 2>/dev/null)
  if [[ -z "$existing" ]]; then
    break
  fi
  NEXT_NUM=$((NEXT_NUM + 1))
done

PREFIX=$(printf "R%02d" "$NEXT_NUM")
FILENAME="${PREFIX}-${SAFE_DESC}.md"
OUTPUT="$REVIEWS_DIR/$FILENAME"

TODAY=$(date +%Y-%m-%d)

# Generate the note
cat > "$OUTPUT" << NOTE
---
title: ""
tags:
  - tipo/revisao
  - status/aberto
review_id: "${PREFIX}"
status: aberto
priority: ""
source: reuniao
date_opened: "${TODAY}"
date_closed: ""
target_chapter: ""
correction_layers: []
evidence_layer: ""
claim_ids: []
verified_by_script: ""
aliases: []
---

<!--
ITEM REGISTRADO DURANTE SESSAO DE ORIENTACAO

Ações pendentes para o agente:
1. Preencher "title" com descricao curta.
2. Preencher "priority" (alta/media/baixa).
3. Preencher "target_chapter" (capitulo afetado).
4. Preencher "evidence_layer" (codigo/dados/literatura/vault/monografia).
5. Na seção "Análise Técnica", investigar a solicitação contra código, dados, literatura e vault.
6. Na seção "Plano de Correção", definir a abordagem mais determinística possível.
7. Remover este comentário após preencher.
-->

## Solicitação Original

> $(echo "$DESCRIPTION" | sed 's/-/ /g')

## Análise Técnica

### Problema Identificado

<!-- Qual a lacuna, inconsistência ou melhoria solicitada? -->

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| \`src/\` (código) | ... | ... |
| \`src/data/\` (dados) | ... | ... |
| Literatura | ... | ... |
| \`vault/\` | ... | ... |
| \`monografia/\` (texto atual) | ... | ... |

### Causa Raiz

<!-- O que causou o problema? Cópia manual? Desatualização? Interpretação incorreta? -->

## Plano de Correção

> Este item foi registrado em sessão de orientação e ainda não tem plano de correção.
> Revisar a análise técnica acima e definir a abordagem mais determinística.

### Abordagem Preferida (mais determinística)

- [ ] **Camada 1 — Código/Dados:** ...
- [ ] **Camada 2 — Script determinístico:** ...
- [ ] **Camada 3 — Edição agentica:** ...

### Comandos e Passos

\`\`\`bash

\`\`\`

### Artefatos Afetados

- ...

## Verificação

### Critério de Aceite

<!-- O que precisa ser verdade para considerar resolvido? -->

### Script de Validação

\`\`\`bash

\`\`\`

## Notas

Registrado em sessão de orientação em ${TODAY}.

NOTE

echo "$OUTPUT"
