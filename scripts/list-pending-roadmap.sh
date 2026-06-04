#!/usr/bin/env bash
# list-pending-roadmap.sh — Lista itens pendentes do roadmap-monografia.md
set -euo pipefail

ROADMAP="${1:-vault/writing/roadmap-monografia.md}"

if [ ! -f "$ROADMAP" ]; then
  echo "Erro: $ROADMAP não encontrado" >&2
  exit 1
fi

echo "=== Pendentes ==="
echo ""

grep '^| P[0-9]' "$ROADMAP" | while IFS='' read -r line; do
  id=$(echo "$line" | awk -F'|' '{print $2}' | sed 's/^ *//;s/ *$//')
  tarefa=$(echo "$line" | awk -F'|' '{print $3}' | sed 's/^ *//;s/ *$//')
  status=$(echo "$line" | awk -F'|' '{print $5}' | sed 's/^ *//;s/ *$//')

  if [[ "$status" != Concluída* ]]; then
    printf "  %-5s %-55s [%s]\n" "$id" "$tarefa" "$status"
  fi
done

echo ""
echo "=== Concluídos ==="
echo ""
grep '^| P[0-9]' "$ROADMAP" | while IFS='' read -r line; do
  id=$(echo "$line" | awk -F'|' '{print $2}' | sed 's/^ *//;s/ *$//')
  tarefa=$(echo "$line" | awk -F'|' '{print $3}' | sed 's/^ *//;s/ *$//')
  status=$(echo "$line" | awk -F'|' '{print $5}' | sed 's/^ *//;s/ *$//')

  if [[ "$status" == Concluída* ]]; then
    printf "  %-5s %s\n" "$id" "$tarefa"
  fi
done
