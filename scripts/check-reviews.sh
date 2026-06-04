#!/usr/bin/env bash
# Deterministic validation for advisor review requests (Rxx).
set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
REVIEWS_DIR="$REPO_DIR/vault/writing/review-solicitacoes"
STATUS=0

usage() {
  cat <<EOF
Usage: $(basename "$0") [OPTION]

Validate and report on advisor review requests (Rxx).

Options:
  --help           Show this help
  --check          Validate all Rxx notes (default)
  --count          Count total review items
  --count --status <csv>  Count items with given status(es)
  --list           List all Rxx items
  --list --status <csv>   List items filtered by status(es)
  --update-index   Regenerate the index.md with current state
  --fail-on-open   Exit with code 1 if any open/analisando/planejado/em-andamento items exist
  --fail-on-high   Exit with code 1 if any high-priority open items exist

Status filter: comma-separated list (e.g. "aberto,analisando")
EOF
  exit 0
}

MODE="check"
STATUS_FILTER=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --help) usage ;;
    --check) MODE="check" ;;
    --count) MODE="count" ;;
    --list) MODE="list" ;;
    --update-index) MODE="update-index" ;;
    --fail-on-open) MODE="fail-on-open" ;;
    --fail-on-high) MODE="fail-on-high" ;;
    --status)
      shift
      STATUS_FILTER="$1"
      ;;
    *) echo "Unknown option: $1" >&2; usage ;;
  esac
  shift
done

# Collect all Rxx note files
RXX_FILES=()
while IFS= read -r -d '' f; do
  RXX_FILES+=("$f")
done < <(find "$REVIEWS_DIR" -maxdepth 1 -name 'R*.md' ! -name 'index.md' -print0 | sort -z)

# Parse YAML frontmatter helper using awk
parse_yaml() {
  local file="$1"
  local field="$2"
  awk '
    /^---$/ { count++; next }
    count == 1 {
      if ($0 ~ /^'"$field"': /) {
        sub(/^[^:]+:[[:space:]]*/, "")
        gsub(/^[[:space:]]+|[[:space:]]+$/, "")
        if (($0 ~ /^"/ && $0 ~ /"$/) || ($0 ~ /^\x27/ && $0 ~ /\x27$/)) {
          gsub(/^.|.$/, "")
        }
        print
      }
    }
    count == 2 { exit }
  ' "$file"
}

# Parse list fields (YAML arrays like [a, b] or multiline - a\n - b)
parse_yaml_list() {
  local file="$1"
  local field="$2"
  awk '
    /^---$/ { count++; next }
    count == 1 {
      if ($0 ~ /^'"$field"':[[:space:]]*\[/) {
        sub(/^[^:]+:[[:space:]]*\[/, "")
        sub(/\]/, "")
        gsub(/[[:space:]]/, "")
        split($0, arr, ",")
        for (i in arr) print arr[i]
      } else if ($0 ~ /^'"$field"':[[:space:]]*$/) {
        in_field = 1
        next
      } else if (in_field && $0 ~ /^[[:space:]]+- /) {
        sub(/^[[:space:]]*- /, "")
        print
      } else if (in_field && $0 ~ /^[[:space:]]*- \[/) {
        sub(/^[[:space:]]*- \[/, "")
        sub(/\]/, "")
        split($0, arr, ",")
        for (i in arr) print arr[i]
      } else if (in_field && $0 ~ /^[a-zA-Z]/) {
        in_field = 0
      }
    }
    count == 2 { exit }
  ' "$file"
}

# --collect all Rxx data
declare -A RXX_STATUS RXX_PRIORITY RXX_SOURCE RXX_DATE RXX_CLOSED RXX_CHAPTER RXX_EVIDENCE
declare -A RXX_LAYERS RXX_CLAIMS RXX_VERIFIED

for f in "${RXX_FILES[@]}"; do
  basename=$(basename "$f")
  id="${basename%.md}"
  RXX_STATUS["$id"]=$(parse_yaml "$f" "status")
  RXX_PRIORITY["$id"]=$(parse_yaml "$f" "priority")
  RXX_SOURCE["$id"]=$(parse_yaml "$f" "source")
  RXX_DATE["$id"]=$(parse_yaml "$f" "date_opened")
  RXX_CLOSED["$id"]=$(parse_yaml "$f" "date_closed")
  RXX_CHAPTER["$id"]=$(parse_yaml "$f" "target_chapter")
  RXX_EVIDENCE["$id"]=$(parse_yaml "$f" "evidence_layer")
  RXX_VERIFIED["$id"]=$(parse_yaml "$f" "verified_by_script")
  RXX_LAYERS["$id"]=$(parse_yaml_list "$f" "correction_layers" | tr '\n' ' ')
  RXX_CLAIMS["$id"]=$(parse_yaml_list "$f" "claim_ids" | tr '\n' ' ')
done

# --filter utility
filter_items() {
  local field="$1"
  local csv="$2"
  local -n arr=$3
  local result=()
  if [[ -z "$csv" ]]; then
    for id in "${!arr[@]}"; do
      result+=("$id")
    done
  else
    IFS=',' read -ra filters <<< "$csv"
    for id in "${!arr[@]}"; do
      val="${arr[$id]}"
      for ftr in "${filters[@]}"; do
        if [[ "$val" == "$ftr" ]]; then
          result+=("$id")
          break
        fi
      done
    done
  fi
  echo "${result[@]}"
}

# === MODE: count ===
if [[ "$MODE" == "count" ]]; then
  if [[ -n "$STATUS_FILTER" ]]; then
    ids=($(filter_items "status" "$STATUS_FILTER" RXX_STATUS))
    echo "${#ids[@]}"
  else
    echo "${#RXX_FILES[@]}"
  fi
  exit 0
fi

# === MODE: list ===
if [[ "$MODE" == "list" ]]; then
  ids=($(filter_items "status" "$STATUS_FILTER" RXX_STATUS))
  if [[ ${#ids[@]} -eq 0 ]]; then
    echo "Nenhum item de revisao encontrado."
    exit 0
  fi
  printf "%-8s %-14s %-8s %-12s %s\n" "ID" "Status" "Prior." "Capítulo" "Fonte"
  echo "----------------------------------------------------------------"
  for id in "${ids[@]}"; do
    printf "%-8s %-14s %-8s %-12s %s\n" \
      "$id" \
      "${RXX_STATUS[$id]:-}" \
      "${RXX_PRIORITY[$id]:-}" \
      "${RXX_CHAPTER[$id]:-}" \
      "${RXX_SOURCE[$id]:-}"
  done
  exit 0
fi

# === MODE: fail-on-open ===
if [[ "$MODE" == "fail-on-open" ]]; then
  ids=($(filter_items "status" "aberto,analisando,planejado,em-andamento" RXX_STATUS))
  if [[ ${#ids[@]} -gt 0 ]]; then
    echo "FAIL: ${#ids[@]} item(ns) de revisao pendente(s): ${ids[*]}" >&2
    exit 1
  fi
  echo "PASS: nenhum item de revisao pendente"
  exit 0
fi

# === MODE: fail-on-high ===
if [[ "$MODE" == "fail-on-high" ]]; then
  ids=($(filter_items "status" "aberto,analisando,planejado,em-andamento" RXX_STATUS))
  fail=0
  for id in "${ids[@]}"; do
    if [[ "${RXX_PRIORITY[$id]}" == "alta" ]]; then
      echo "FAIL: item alta prioridade pendente: $id (${RXX_STATUS[$id]})" >&2
      fail=1
    fi
  done
  if [[ $fail -eq 0 ]]; then
    echo "PASS: nenhum item de alta prioridade pendente"
  fi
  exit $fail
fi

# === MODE: update-index ===
if [[ "$MODE" == "update-index" ]]; then
  index_file="$REVIEWS_DIR/index.md"
  # Collect items by status
  aberto_ids=($(filter_items "status" "aberto" RXX_STATUS))
  andamento_ids=($(filter_items "status" "analisando,planejado,em-andamento" RXX_STATUS))
  resolvido_ids=($(filter_items "status" "resolvido" RXX_STATUS))
  cancelado_ids=($(filter_items "status" "cancelado" RXX_STATUS))

  generate_list() {
    local -n arr=$1
    local label="$2"
    echo "### $label"
    for id in "${arr[@]}"; do
      echo "- [${RXX_PRIORITY[$id]}] [[$id]] — ${RXX_STATUS[$id]} (${RXX_CHAPTER[$id]:-sem capitulo})"
    done
    echo ""
  }

  cat > "$index_file" << INDEXEOF
---
title: Índice de Solicitações de Revisão
tags:
  - tipo/index
  - revisao
  - monografia
updated: $(date +%Y-%m-%d)
---

# Índice de Solicitações de Revisão

Lista consolidada de todas as solicitações de revisão do orientador.

## Por Status

$(generate_list aberto_ids "Abertas")
$(generate_list andamento_ids "Em Andamento")
$(generate_list resolvido_ids "Resolvidas")
$(generate_list cancelado_ids "Canceladas")

## Métricas

| Métrica | Valor |
|---|---|
| Total de solicitações | ${#RXX_FILES[@]} |
| Resolvidas | ${#resolvido_ids[@]} |
| Pendentes | $((${#aberto_ids[@]} + ${#andamento_ids[@]})) |
| Canceladas | ${#cancelado_ids[@]} |

---

**Ver também:** [[review-roadmap]], [[claim-evidence-matrix]], [[roadmap-monografia]]
INDEXEOF
  echo "PASS index.md atualizado (${#RXX_FILES[@]} itens)"
  exit 0
fi

# === MODE: check (default) ===
errors=()
warnings=()

# 1. Valid IDs match filenames
for f in "${RXX_FILES[@]}"; do
  basename=$(basename "$f" .md)
  if ! [[ "$basename" =~ ^R[0-9]{2,}- ]]; then
    errors+=("$f: nome invalido. Deve ser R<numero>-<descricao>.md")
  fi
done

# 2. Check IDs (filename Rxx)
declare -A SEEN_IDS
for f in "${RXX_FILES[@]}"; do
  basename=$(basename "$f" .md)
  id="${basename%%-*}"
  if [[ -n "${SEEN_IDS[$id]:-}" ]]; then
    errors+=("ID duplicado: $id")
  fi
  SEEN_IDS["$id"]=1
done

# 3. Required YAML fields per note
REQUIRED_FIELDS=("status" "priority" "source" "date_opened" "correction_layers" "evidence_layer")
VALID_STATUSES=("aberto" "analisando" "planejado" "em-andamento" "resolvido" "cancelado")
VALID_PRIORITIES=("alta" "media" "baixa")
VALID_SOURCES=("reuniao" "email" "comentario-tex" "anotacao-manual" "banca")
VALID_LAYERS=("codigo" "dados" "pdf" "latex-macro" "vault" "agentico")
VALID_EVIDENCE=("codigo" "dados" "literatura" "vault" "monografia")

for f in "${RXX_FILES[@]}"; do
  basename=$(basename "$f" .md)
  id="$basename"

  # Required fields present (scalar vs list)
  SCALAR_FIELDS=("status" "priority" "source" "date_opened" "evidence_layer")
  LIST_FIELDS=("correction_layers")
  for field in "${SCALAR_FIELDS[@]}"; do
    val=$(parse_yaml "$f" "$field")
    if [[ -z "$val" ]]; then
      errors+=("$id: campo obrigatorio faltando: $field")
    fi
  done
  for field in "${LIST_FIELDS[@]}"; do
    val=$(parse_yaml_list "$f" "$field" | tr -d '\n')
    if [[ -z "$val" ]]; then
      errors+=("$id: campo obrigatorio faltando: $field")
    fi
  done

  # Status validity
  status=$(parse_yaml "$f" "status")
  if [[ -n "$status" ]]; then
    valid=0
    for vs in "${VALID_STATUSES[@]}"; do [[ "$status" == "$vs" ]] && valid=1 && break; done
    [[ $valid -eq 0 ]] && errors+=("$id: status invalido: $status")
  fi

  # Priority validity
  priority=$(parse_yaml "$f" "priority")
  if [[ -n "$priority" ]]; then
    valid=0
    for vp in "${VALID_PRIORITIES[@]}"; do [[ "$priority" == "$vp" ]] && valid=1 && break; done
    [[ $valid -eq 0 ]] && errors+=("$id: prioridade invalida: $priority")
  fi

  # Source validity
  source=$(parse_yaml "$f" "source")
  if [[ -n "$source" ]]; then
    valid=0
    for vs in "${VALID_SOURCES[@]}"; do [[ "$source" == "$vs" ]] && valid=1 && break; done
    [[ $valid -eq 0 ]] && errors+=("$id: fonte invalida: $source")
  fi

  # correction_layers validity
  layers="${RXX_LAYERS[$id]}"
  for layer in $layers; do
    valid=0
    for vl in "${VALID_LAYERS[@]}"; do [[ "$layer" == "$vl" ]] && valid=1 && break; done
    [[ $valid -eq 0 ]] && errors+=("$id: camada de correcao invalida: $layer")
  done

  # evidence_layer validity
  evidence=$(parse_yaml "$f" "evidence_layer")
  if [[ -n "$evidence" ]]; then
    valid=0
    for ve in "${VALID_EVIDENCE[@]}"; do [[ "$evidence" == "$ve" ]] && valid=1 && break; done
    [[ $valid -eq 0 ]] && errors+=("$id: camada de evidencia invalida: $evidence")
  fi

  # Resolved items must have date_closed
  if [[ "$status" == "resolvido" ]]; then
    closed=$(parse_yaml "$f" "date_closed")
    if [[ -z "$closed" ]]; then
      errors+=("$id: resolvido sem date_closed")
    fi
  fi

  # Agentic layer must have justification
  if [[ "$layers" == *"agentico"* ]]; then
    if ! grep -q "Camada 3.*Edição agentica" "$f" 2>/dev/null; then
      warnings+=("$id: usa camada agentica mas nao tem secao 'Camada 3 — Edicao agentica' no plano")
    fi
  fi

  # Skip agentic check for resolved items (justification already accepted)
done

# 4. Validate claim_ids against claim-evidence-matrix
CLAIM_MATRIX="$REPO_DIR/vault/writing/claim-evidence-matrix.md"
if [[ -f "$CLAIM_MATRIX" ]]; then
  declare -A VALID_CLAIMS
  while IFS= read -r line; do
    if [[ "$line" =~ \|\ +([A-Z][0-9]+)\ +\| ]]; then
      VALID_CLAIMS["${BASH_REMATCH[1]}"]=1
    fi
  done < <(grep '| [A-Z][0-9]\+ |' "$CLAIM_MATRIX")

  for f in "${RXX_FILES[@]}"; do
  basename=$(basename "$f" .md)
  id="$basename"
    claims="${RXX_CLAIMS[$id]}"
    for claim in $claims; do
      if [[ -z "${VALID_CLAIMS[$claim]:-}" ]]; then
        warnings+=("$id: claim_id '$claim' nao encontrado na claim-evidence-matrix")
      fi
    done
  done
fi

# Report
if [[ ${#errors[@]} -gt 0 ]]; then
  for e in "${errors[@]}"; do echo "ERROR $e" >&2; done
  STATUS=1
fi
if [[ ${#warnings[@]} -gt 0 ]]; then
  for w in "${warnings[@]}"; do echo "WARN $w" >&2; done
fi

if [[ $STATUS -eq 0 ]]; then
  echo "PASS check-reviews (${#RXX_FILES[@]} itens, ${#errors[@]} erros, ${#warnings[@]} warnings)"
else
  echo "FAIL check-reviews (${#RXX_FILES[@]} itens, ${#errors[@]} erros, ${#warnings[@]} warnings)" >&2
fi

exit $STATUS
