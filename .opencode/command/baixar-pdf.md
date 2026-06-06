---
description: [DEPRECIADO] Use /incorporar — incorpora artigo ao vault (PDF → resumo → vault → canvas → claims)
---

# /baixar-pdf

> **Este comando foi substituído por `/incorporar`**, que unifica download de PDF, extração de resumo, criação de nota no vault, atualização do canvas de conhecimento e verificação de claims.
>
> Use `/incorporar $ARGUMENTS` para o mesmo efeito, com integração completa ao ecossistema.

## Migração

| Antes | Agora |
|---|---|
| `/baixar-pdf 10.1007/...` | `/incorporar 10.1007/...` |
| `/baixar-pdf kennedy1995particle` | `/incorporar kennedy1995particle` |

## Dependências

- MCPs: `doiget`, `scihub`, `pdf-reader`, `crossref`, `semantic-scholar`
- Script: `scripts/download-pdfs.sh`
