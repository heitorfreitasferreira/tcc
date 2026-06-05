#!/usr/bin/env bash
# Baixa PDF de artigo acadêmico por chave BibTeX usando cascata MCP
# Uso: ./scripts/baixar-pdf-chave.sh <bibtex-key>
#   ./scripts/baixar-pdf-chave.sh hutter2011smac

set -euo pipefail
cd "$(dirname "$0")/.."

KEY="${1:?Uso: $0 <bibtex-key>}"
SUMMARY_DIR="vault/papers/summaries"
PDF_DIR="vault/papers/pdfs"
mkdir -p "$SUMMARY_DIR" "$PDF_DIR"

# 1. Extrai metadados do vault ou do BibTeX
NOTE="vault/papers/${KEY}.md"
DOI=""
TITLE=""

if [[ -f "$NOTE" ]]; then
    echo "📄 Nota vault encontrada: $NOTE"
    DOI=$(grep -E '^doi:' "$NOTE" | head -1 | sed 's/^doi:[[:space:]]*//; s/"//g')
    TITLE=$(grep -E '^title:' "$NOTE" | head -1 | sed 's/^title:[[:space:]]*//; s/"//g')
elif grep -q "$KEY" monografia/bib/abntex2-references.bib 2>/dev/null; then
    echo "📖 Chave encontrada no .bib"
    DOI=$(grep -A5 "^@.*{${KEY}," monografia/bib/abntex2-references.bib | grep -oP 'doi\s*=\s*"\K[^"]+' | head -1)
    TITLE=$(grep -A5 "^@.*{${KEY}," monografia/bib/abntex2-references.bib | grep -oP 'title\s*=\s*"\K[^"]+' | head -1)
    if [[ -n "$DOI" ]]; then
        echo "   DOI: $DOI"
    fi
fi

# 2. Tenta doiget (OA-first)
PDF_PATH="$PDF_DIR/${KEY}.pdf"
if ! command -v doiget &>/dev/null; then
    echo "⚠️  doiget não instalado — pulando fonte OA"
else
    echo "🔍 doiget: buscando $DOI..."
    doiget fetch "$DOI" 2>/dev/null && {
        # doiget salva em ~/papers/; copia pro vault
        FOUND=$(find ~/papers/ -name "*.pdf" -newer "$0" 2>/dev/null | head -1)
        if [[ -n "$FOUND" ]]; then
            cp "$FOUND" "$PDF_PATH"
            echo "✅ doiget: PDF salvo em $PDF_PATH"
        fi
    } || echo "   doiget: não disponível via OA"
fi

# 3. Fallback: script download-pdfs
if [[ ! -f "$PDF_PATH" ]]; then
    echo "🔍 download-pdfs.py: tentando fontes secundárias..."
    if [[ -n "$DOI" ]]; then
        python3 scripts/download-pdfs.py --keys "$KEY" --use-scihub --min-pages 1
    else
        python3 scripts/download-pdfs.py --keys "$KEY" --min-pages 1
    fi
fi

# 4. Gera resumo via pdfinfo + texto
if [[ -f "$PDF_PATH" ]] && command -v pdftotext &>/dev/null; then
    echo "📝 Extraindo resumo do PDF..."
    SUMMARY_FILE="$SUMMARY_DIR/${KEY}.md"
    PAGES=$(pdfinfo "$PDF_PATH" 2>/dev/null | grep -oP 'Pages:\s+\K\d+' || echo "?")
    TITLE="${TITLE:-$KEY}"
    cat > "$SUMMARY_FILE" <<- SUMMARYEOF
---
bibtex-key: "$KEY"
title: "$TITLE"
pages: $PAGES
pdf: "papers/pdfs/${KEY}.pdf"
status: "lido-parcial"
---

## Resumo automático

- **Páginas**: $PAGES
- **DOI**: $DOI

### Primeira página (amostra)

\`\`\`
$(pdftotext "$PDF_PATH" - 2>/dev/null | head -30)
\`\`\`
SUMMARYEOF
    echo "✅ Resumo salvo em $SUMMARY_FILE"
elif [[ -f "$PDF_PATH" ]]; then
    echo "ℹ️  PDF baixado em $PDF_PATH (pdftotext não disponível)"
else
    echo "❌ Falha ao baixar PDF para $KEY"
    exit 1
fi
