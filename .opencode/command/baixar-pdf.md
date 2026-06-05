---
description: Baixa PDF de artigo por DOI/chave BibTeX e extrai resumo automaticamente
---

# /baixar-pdf

Baixa o PDF de um artigo acadêmico usando a cascata de MCPs acadêmicos (`doiget` → `scihub` → `webfetch`/`semantic-scholar`) e extrai um resumo via `pdf-reader`.

## Fluxo

1. **Identificar**: Recebe DOI ou chave BibTeX. Se for chave, busca metadados em `vault/papers/<key>.md` ou via `crossref_get_work`/`semantic-scholar`.
2. **Baixar**: Tenta `doiget_fetch_paper` (via Open Access) → `scihub_download_scihub_pdf` → `scripts/download-pdfs.sh --keys <key>`
3. **Ler**: Extrai texto com `pdf-reader` e grava nota de resumo em `vault/papers/summaries/<key>.md`
4. **Atualizar**: Atualiza campo `pdf:` no frontmatter da nota vault

## Uso

```
/baixar-pdf 10.1007/978-3-642-25566-3_40
/baixar-pdf hutter2011smac
```

## Dependências

- MCPs: `doiget`, `scihub`, `pdf-reader`, `crossref`, `semantic-scholar`
- Script: `scripts/download-pdfs.sh`
