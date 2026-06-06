# Componente: Literatura, PDFs e BibTeX

## Responsabilidade

Garantir que referencias academicas estejam rastreadas desde a busca ate o texto final.

## Fontes

| Fonte | Uso |
|---|---|
| `academic-search` | Busca ampla em Crossref e Semantic Scholar |
| `arxiv` | Preprints e PDFs arXiv |
| `google-scholar` | Busca complementar |
| `crossref` | Metadados e referencias por DOI |
| `scholar-sidekick` | BibTeX, OA, retraction e validacao de citacao |
| `doiget` | Resolucao e fetch OA |
| `pdf-reader` | Extracao de texto de PDFs |

## Pipeline de Incorporacao

1. Resolver identificador: DOI, arXiv, titulo ou chave BibTeX.
2. Obter PDF quando possivel.
3. Validar metadados.
4. Atualizar `monografia/bib/abntex2-references.bib`.
5. Criar ou atualizar `vault/papers/<key>.md`.
6. Atualizar canvas e conexoes.
7. Registrar evento no roadmap.

## Estados de Paper

| Campo | Valores comuns |
|---|---|
| `reading_status` | `pendente`, `resumo-lido`, `lido-parcial`, `lido` |
| `validation_status` | `nao-validado`, `validado`, `requer-validacao`, `bloqueado` |
| `pdf_status` | `disponivel`, `ausente`, `corrompido`, `integro` |
| `rating` | 0 a 5 |

## Referencia Condicional

Uma referencia e condicional quando aparece no texto final, mas ainda tem algum problema:

- nao tem nota no vault;
- nao tem PDF;
- PDF esta ilegivel;
- nota nao tem resumo ou validacao minima;
- claim apoiado pelo paper nao foi verificado.

Regra: referencias condicionais devem virar tarefa antes de permanecerem no texto final.

## Arquivos Canonicos

- BibTeX: `monografia/bib/abntex2-references.bib`
- Papers: `vault/papers/*.md`
- PDFs: `vault/papers/pdfs/*.pdf`
- Template: `vault/templates/paper-note.md`
- Importador: `scripts/import-bib-to-vault.sh`
- Download: `scripts/download-pdfs.sh`
