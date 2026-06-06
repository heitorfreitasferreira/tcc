#!/usr/bin/env bash
# list-pending-roadmap.sh — Lista tarefas do sistema de fila
# Lê de vault/roadmap/tarefas/P*.md.
set -euo pipefail

TASKS_DIR="${1:-vault/roadmap/tarefas}"

if [ ! -d "$TASKS_DIR" ]; then
  echo "Erro: $TASKS_DIR não encontrado" >&2
  exit 1
fi

echo "=== Pendentes ==="
echo ""
for f in "$TASKS_DIR"/P*.md; do
  [ -f "$f" ] || continue
  id=$(grep -m1 '^task_id:' "$f" | sed 's/^task_id: *//')
  title=$(grep -m1 '^title:' "$f" | sed 's/^title: *"//;s/"$//')
  status=$(grep -m1 '^status:' "$f" | sed 's/^status: *//')
  fase=$(grep -m1 '^fase:' "$f" | sed 's/^fase: *//')
  priority=$(grep -m1 '^priority:' "$f" | sed 's/^priority: *//')
  if [[ "$status" != concluida* ]]; then
    printf "  %-5s %-55s [%s/%s]\n" "$id" "$title" "$status" "$priority"
  fi
done | sort

echo ""
echo "=== Concluídos ==="
echo ""
for f in "$TASKS_DIR"/P*.md; do
  [ -f "$f" ] || continue
  id=$(grep -m1 '^task_id:' "$f" | sed 's/^task_id: *//')
  title=$(grep -m1 '^title:' "$f" | sed 's/^title: *"//;s/"$//')
  status=$(grep -m1 '^status:' "$f" | sed 's/^status: *//')
  if [[ "$status" == concluida* ]]; then
    printf "  %-5s %s\n" "$id" "$title"
  fi
done | sort
