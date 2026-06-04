#!/usr/bin/env bash
set -euo pipefail

# import-bib-to-vault.sh
# Importa referências BibTeX para notas markdown no vault.
# Uso: ./scripts/import-bib-to-vault.sh [--bib <caminho>] [--vault <caminho>]
#       ./scripts/import-bib-to-vault.sh <bibtex-key>  # importa uma específica

BIB="${BIB:-monografia/bib/abntex2-references.bib}"
VAULT="${VAULT:-vault}"

bib_file="$BIB"
vault_dir="$VAULT"

# Normaliza caminhos relativos para absolutos baseados no raiz do repo
REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
[[ "$bib_file" != /* ]] && bib_file="$REPO_ROOT/$bib_file"
[[ "$vault_dir" != /* ]] && vault_dir="$REPO_ROOT/$vault_dir"

if [[ ! -f "$bib_file" ]]; then
    echo "Erro: arquivo .bib não encontrado: $bib_file" >&2
    exit 1
fi

# Função para extrair campos do BibTeX via regex
extract_field() {
    local field="$1"
    local entry="$2"
    echo "$entry" | (grep -i "^[[:space:]]*$field[[:space:]]*=" || true) | \
        sed "s/^[[:space:]]*$field[[:space:]]*=[[:space:]]*//" | \
        sed 's/^[{"]//;s/["}],*$//' | \
        tr -d '\n'
}

# Função para extrair tags baseadas no texto do abstract/title
infer_tags() {
    local entry="$1"
    local tags=""
    local lower
    lower=$(echo "$entry" | tr '[:upper:]' '[:lower:]')
    echo "$lower" | grep -q "genetic\|mutation\|crossover\|selection" && tags="${tags}metodo/ga " || true
    echo "$lower" | grep -q "particle swarm\|pso\|swarm" && tags="${tags}metodo/pso " || true
    echo "$lower" | grep -q "ant colony\|aco\|ferom\|ant system" && tags="${tags}metodo/aco " || true
    echo "$lower" | grep -q "traveling\|tsp\|tsp-d\|fstsp\|salesman" && tags="${tags}area/tsp " || true
    echo "$lower" | grep -q "drone\|uav\|unmanned\|flying sidekick\|aerial" && tags="${tags}area/drone-routing " || true
    echo "$lower" | grep -q "heuristic\|metaheuristic\|optimization" && tags="${tags}area/bio-inspired " || true
    echo "$lower" | grep -q "np-complet\|intractability\|computacional" && tags="${tags}papel/fundacional " || true
    echo "$lower" | grep -q "route\|path\|patrol\|routing" && tags="${tags}area/routing " || true
    echo "$tags" | sed 's/ $//'
}

yaml_tag_lines() {
    local tags="$1"
    printf '  - tipo/paper\n  - status/pendente\n  - evidencia/referencia\n'
    for tag in $tags; do
        printf '  - %s\n' "$tag"
    done
}

# Normaliza nome do arquivo a partir da bibtex-key
key_to_filename() {
    echo "$1" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-zA-Z0-9_-]/-/g'
}

# Remove acentos e caracteres especiais de autor para tag
author_to_tag() {
    echo "$1" | sed 's/[[:space:]].*//' | tr '[:upper:]' '[:lower:]' | \
        sed 's/ç/c/g;s/é/e/g;s/ê/e/g;s/á/a/g;s/ã/a/g;s/à/a/g;s/ó/o/g;s/ô/o/g;s/õ/o/g;s/í/i/g;s/ú/u/g'
}

generate_note() {
    local raw_entry="$1"
    local key
    key=$(echo "$raw_entry" | head -1 | (grep -oP '@\w+\{\K[^,]+' || true) | sed 's/[[:space:]]//g')
    [[ -z "$key" ]] && return

    local title authors year doi
    title=$(extract_field "title" "$raw_entry")
    authors=$(extract_field "author" "$raw_entry")
    year=$(extract_field "year" "$raw_entry")
    doi=$(extract_field "doi" "$raw_entry")

    local filename
    filename=$(key_to_filename "$key")
    local filepath="$vault_dir/papers/$filename.md"

    if [[ -f "$filepath" ]]; then
        echo "  [skip] $key → $filepath (já existe)"
        return
    fi

    local tags
    tags=$(infer_tags "$raw_entry")

    local tag_lines
    tag_lines=$(yaml_tag_lines "$tags")

    local first_author
    first_author=$(echo "$authors" | sed 's/ and.*//' | sed 's/{//g;s/}//g')

    cat > "$filepath" << NOTE
---
title: "${title}"
authors: [${first_author}]
year: ${year}
doi: "${doi}"
bibtex_key: ${key}
bibtex-key: ${key}
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 0
role: ""
areas: []
methods: []
chapters: []
claim_support: []
aliases: []
tags:
${tag_lines}
---

## PDF

<!-- Se disponível, link para o PDF local: [[papers/pdfs/${key}.pdf]] -->

## Tese Central

<!-- Qual é a afirmação principal do artigo em 1-3 frases? -->

## Resumo

<!-- Pendente: gerar via SciHub + agente -->

## Contribuições Principais

-

## Relevância para o TCC

-

## Uso no TCC

- Capítulo(s):
- Claim(s) apoiado(s):
- Como citar na monografia:

## Métodos e Abordagens

-

## Evidência / Resultado Relevante

-

## Limitações de Uso

> [!warning] Limitação
> Registrar escopo, ameaça à validade, diferença entre o problema do paper e o TSP-SD-ATP, ou motivo para não generalizar.

## Conexões

- Fundamenta:
- Relacionado a:
- Contrasta com:
- Apoia claim:
- Usado em capítulo:

## Notas e Insights

-

## Citações-chave

>
NOTE

    echo "  [ok] $key → $filepath"
}

# Se passou bibtex-key como argumento, importa só uma
if [[ $# -ge 1 && "$1" != "--bib" && "$1" != "--vault" ]]; then
    echo "[import-bib-to-vault] Importando referência: $1"
    entry=$(awk -v RS="@" -v key="$1" 'index($0, key) && index($0, "{")' "$bib_file" | head -1)
    if [[ -z "$entry" ]]; then
        echo "Erro: chave '$1' não encontrada em $bib_file" >&2
        exit 1
    fi
    generate_note "@$entry"
    exit 0
fi

# Parse flags
while [[ $# -gt 0 ]]; do
    case "$1" in
        --bib) shift; bib_file="$1"; shift ;;
        --vault) shift; vault_dir="$1"; shift ;;
        *) shift ;;
    esac
done

echo "[import-bib-to-vault] Importando de: $bib_file"
echo "[import-bib-to-vault] Destino: $vault_dir/papers/"
echo ""

# Lê o arquivo .bib e processa cada entrada (ignora comentários)
tmpfile=$(mktemp)
awk 'BEGIN { RS="@"; count=0 } /^[a-zA-Z]/{ count++; printf "@%s\0", $0 }' "$bib_file" > "$tmpfile"
count=0
while IFS= read -r -d '' entry; do
    entry=$(echo "$entry" | sed '/^[[:space:]]*$/d')
    [[ -z "$entry" ]] && continue
    generate_note "$entry"
    count=$((count + 1))
done < "$tmpfile"
rm -f "$tmpfile"

echo ""
echo "[import-bib-to-vault] Concluído. $count entradas processadas."
