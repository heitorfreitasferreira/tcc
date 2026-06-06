---
description: Incorpora artigo ao vault a partir de DOI, arXiv, chave BibTeX, PDF local ou processa PDFs pendentes
---

# /incorporar

Incorpora completamente um artigo acadêmico ao ecossistema do TCC: obtém PDF, extrai metadados, gera resumo, cria nota no vault, atualiza canvas de conhecimento, verifica claims e conecta com papers existentes. **Não altera o texto da monografia** — isso é feito depois via `/claudiney` ou `/roadmap-criar`.

## Modos de entrada

`$ARGUMENTS` aceita:

| Entrada | Exemplo |
|---|---|
| DOI | `10.1007/978-3-642-25566-3_40` |
| arXiv ID | `2301.08745` ou `arXiv:2301.08745` |
| Chave BibTeX | `kennedy1995particle` |
| Caminho de PDF | `vault/papers/pdfs/meu-artigo.pdf` ou caminho absoluto |
| Palavra `vault` | Processa todos os PDFs em `vault/papers/pdfs/` sem nota correspondente |
| Palavra `fila` | Processa PDFs listados em `.opencode/log/pdf-events.log` que ainda não foram incorporados |

Sem argumentos, lista PDFs pendentes e pergunta qual processar.

## Fluxo completo

### Fase 1 — Obter PDF

- **Se DOI/arXiv/key recebido**: baixa via cascata MCP — `doiget_fetch_paper` (OA primeiro) → `scihub` → `bash scripts/download-pdfs.sh --keys <key>`. Salva em `vault/papers/pdfs/<key>.pdf`.
- **Se caminho de PDF recebido**: usa diretamente. Se estiver fora do vault, copia para `vault/papers/pdfs/<key>.pdf`.
- **Se "vault" ou "fila"**: itera sobre PDFs pendentes, processa um por um.

### Fase 2 — Extrair metadados

- Executa `bash scripts/extract-pdf-doi.sh <pdf>` para extrair DOI do PDF.
- Se DOI encontrado:
  - `crossref_get_work(doi)` → metadados completos (título, autores, ano, journal, abstract)
  - `doiget_resolve_paper(doi)` → fonte OA e licença
  - `scholar-sidekick_checkOpenAccess(doi)` → status OA
- Se DOI **não** encontrado:
  - `pdf-reader` extrai texto das primeiras 3 páginas
  - Busca título no `crossref_search_works` + `semantic-scholar`
  - Confirma com o usuário antes de prosseguir
- Obtém BibTeX formatado via `doiget_bibtex_export(doi)`.

### Fase 3 — Gerar resumo

- `pdf-reader` extrai texto completo do PDF.
- Lê abstract, introduction e conclusion.
- Produz resumo estruturado em **português (PT-BR)** com 3-5 frases cobrindo: problema, método, resultado principal.
- Extrai 1-3 **citações-chave** literais (frases impactantes do paper).
- Infere automaticamente:
  - `role`: `fundacional`, `revisao`, `comparativo`, `metodologico`, `aplicacao`, `benchmark`, `limitacao`
  - `areas`: ex: `tsp`, `drone-routing`, `lower-bound`, `routing`
  - `methods`: ex: `ga`, `pso`, `aco`, `exact`, `metaheuristic`, `machine-learning`
  - Tags hierárquicas correspondentes

### Fase 4 — Integrar BibTeX

- Verifica se a chave já existe em `monografia/bib/abntex2-references.bib`.
- Se não existe, adiciona entrada formatada (ABNT-style, com `doi`, autores completos, escaping LaTeX).
- Se a chave existe mas metadados diferem, pergunta se deve atualizar.

### Fase 5 — Criar/atualizar nota no vault

- Cria `vault/papers/<bibtex-key>.md` usando o template canônico (`vault/templates/paper-note.md`).
- Preenche frontmatter completo: `title`, `authors`, `year`, `doi`, `bibtex_key`, `pdf`, `areas`, `methods`, `role`, `reading_status: resumo-lido`, `pdf_status: disponivel`, `rating` (inferido por relevância).
- Preenche seções:
  - **PDF**: embed `![[pdfs/<key>.pdf]]`
  - **Tese Central**: 1-2 frases
  - **Resumo**: 3-5 frases em PT-BR
  - **Contribuições Principais**: bullet list
  - **Relevância para o TCC**: conexão com drones/rTSP/metaheurísticas
  - **Métodos e Abordagens**: técnicas usadas
  - **Citações-chave**: 1-3 citações literais com `>`
  - **Conexões**: `[[wikilinks]]` para áreas e papers relacionados
- Se a nota já existe, **atualiza** (não sobrescreve) campos vazios e enriquece seções pendentes.

### Fase 6 — Atualizar canvas de conhecimento

- Lê `vault/canvas/tcc-knowledge-graph.canvas`.
- Adiciona nó (`type: "file"`) para o paper se não existir.
- Posiciona por área: TSP (~x:400), drone (~x:840), métodos (~x:250), surveys (~x:620).
- Adiciona arestas para:
  - Áreas relevantes (ex: `area-tsp`, `area-drone`, `area-bio`)
  - Métodos relevantes (ex: `method-ga`, `method-pso`, `method-aco`)
- Preserva estrutura existente do canvas.

### Fase 7 — Verificar e atualizar claims

- Lê `vault/claims/` e `vault/bases/claims.base` (use `obsidian-bases` skill se disponível).
- Para cada claim existente, verifica se o paper:
  - **Apoia** (evidência favorável) → adiciona `claim_support: [<ID>]` no frontmatter
  - **Contradiz** (evidência contrária) → registra na seção `Limitações de Uso`
  - **É fonte primária do claim** → sugere vincular como `primary_evidence`
- Se o paper introduz um claim novo e relevante, **sugere** (não cria) um arquivo em `vault/claims/`.
- Registra claims afetados no output do comando.

### Fase 8 — Conectar com papers existentes

- Obtém lista de referências do paper via `crossref_get_references(doi)`.
- Cruza com papers existentes no vault (`ls vault/papers/*.md`).
- Para cada match:
  - Adiciona `[[wikilink]]` bidirecional nas notas (seção `Conexões`).
  - No paper existente: `- Citado por: [[novo-paper]]`
  - No novo paper: `- Fundamenta: [[paper-existente]]` ou `- Relacionado a: [[paper-existente]]`
- Infere similaridade por palavras-chave compartilhadas e sugere conexões adicionais.

### Fase 9 — Sugerir próximos passos

- Se o paper é altamente relevante (rating ≥ 4):
  - Sugere criar item no roadmap: `/roadmap-criar`
  - Sugere `/claudiney` se for evidência para capítulo em revisão
- Lista claims que precisam de atenção (atualizados, novos sugeridos).
- Mostra resumo do que foi feito: `[BibTeX] [Vault] [Canvas] [Claims] [Conexões]`.

## Dependências

- **MCPs**: `crossref`, `doiget`, `semantic-scholar`, `scholar-sidekick`, `pdf-reader`, `scihub`
- **Scripts**: `scripts/extract-pdf-doi.sh`, `scripts/download-pdfs.sh`, `scripts/import-bib-to-vault.sh`
- **Skills**: `knowledge-base`, `vault-semantic-schema`, `vault-tagger`, `json-canvas`, `obsidian-bases`

## Exemplos

```
/incorporar 10.1007/978-3-642-25566-3_40
/incorporar kennedy1995particle
/incorporar vault/papers/pdfs/novo-drone-survey.pdf
/incorporar vault
/incorporar fila
```

## Anti-patterns

- Não altera arquivos `.tex` da monografia (use `/claudiney` para isso).
- Não cria claims automaticamente — apenas sugere.
- Não modifica notas de outros papers além de adicionar wikilinks bidirecionais.
- Não força preenchimento de seções que exigem leitura profunda (ex: `Notas e Insights`).
- Respeita o schema do `vault-semantic-schema` — tags hierárquicas, sem flat tags.
