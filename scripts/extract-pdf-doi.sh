#!/usr/bin/env bash
set -euo pipefail
# Extrai DOI de metadata/primeiras páginas de um PDF
# Uso: scripts/extract-pdf-doi.sh <pdf-path>
# Output: DOI string ou nada se não encontrado

PDF="${1:?Uso: $0 <pdf-path>}"
[[ -f "$PDF" ]] || { echo "Arquivo não encontrado: $PDF" >&2; exit 1; }

# 1. Tenta metadados XMP/Info
if command -v pdfinfo &>/dev/null; then
    DOI=$(pdfinfo "$PDF" 2>/dev/null | grep -i 'doi' | head -1 | sed 's/.*: *//' | tr -d ' ')
    if [[ -n "$DOI" && "$DOI" =~ ^10\. ]]; then
        echo "$DOI"
        exit 0
    fi
fi

# 2. Tenta exiftool (mais completo)
if command -v exiftool &>/dev/null; then
    DOI=$(exiftool "$PDF" 2>/dev/null | grep -i 'doi' | head -1 | sed 's/.*: *//' | tr -d ' ')
    if [[ -n "$DOI" && "$DOI" =~ ^10\. ]]; then
        echo "$DOI"
        exit 0
    fi
fi

# 3. Varredura nas primeiras 3 páginas com regex DOI
if command -v pdftotext &>/dev/null; then
    TEXT=$(pdftotext -f 1 -l 3 "$PDF" - 2>/dev/null)
    DOI=$(echo "$TEXT" | grep -oP '\b10\.\d{4,9}/[-._;()/:A-Za-z0-9]+' | head -1)
    if [[ -n "$DOI" ]]; then
        # Remove trailing punctuation
        DOI=$(echo "$DOI" | sed 's/[.,;:)]$//')
        echo "$DOI"
        exit 0
    fi

    # 4. Tenta arXiv ID
    ARXIV=$(echo "$TEXT" | grep -oP 'arXiv:\s*\K\d{4}\.\d{4,5}' | head -1)
    if [[ -n "$ARXIV" ]]; then
        echo "arxiv:$ARXIV"
        exit 0
    fi
fi

exit 0
