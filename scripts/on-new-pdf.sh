#!/usr/bin/env bash
set -euo pipefail
# Script chamado pelo Obsidian Shell Commands quando um PDF novo aparece no vault.
# Configurar no plugin: Event "File created" → este script com {{event_file_path}}
#
# Uso: scripts/on-new-pdf.sh <caminho-do-pdf>

PDF="${1:?Uso: $0 <caminho-do-pdf>}"
REPO="$(cd "$(dirname "$0")/.." && pwd)"

EXT="${PDF##*.}"
[[ "${EXT,,}" != "pdf" ]] && { echo "[on-new-pdf] Ignorando não-PDF: $PDF"; exit 0; }

FILENAME="$(basename "$PDF" .pdf)"
echo "[on-new-pdf] PDF detectado: $FILENAME"

# Opção A: Se opencode estiver rodando, podemos deixar o plugin interno agir
# Apenas registra o evento para debug
echo "[on-new-pdf] $(date -Iseconds) | $PDF" >> "$REPO/.opencode/log/pdf-events.log"

# Opção B: Tenta extrair DOI e notificar
DOI=$("$REPO/scripts/extract-pdf-doi.sh" "$PDF" 2>/dev/null || true)
if [[ -n "$DOI" ]]; then
    echo "[on-new-pdf] DOI extraído: $DOI"
else
    echo "[on-new-pdf] DOI não encontrado no PDF. Será necessário processamento manual."
fi

echo "[on-new-pdf] Execute /incorporar \"$PDF\" no opencode para integrar este artigo."
