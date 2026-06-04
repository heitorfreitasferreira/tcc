#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$REPO_ROOT"

if ! command -v pdfinfo >/dev/null 2>&1; then
  echo "Erro: pdfinfo não encontrado (instale poppler)." >&2
  exit 1
fi

exec python3 "$REPO_ROOT/scripts/download-pdfs.py" "$@"
