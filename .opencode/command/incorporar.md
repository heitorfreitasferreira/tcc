---
description: Incorpora artigo ao vault a partir de DOI, arXiv, chave BibTeX, PDF local ou processa PDFs pendentes
---

# /incorporar

Incorpora completamente um artigo acadêmico ao ecossistema do TCC: obtém PDF, extrai metadados, gera resumo, cria nota no vault, atualiza canvas de conhecimento, verifica claims e conecta com papers existentes. **Não altera o texto da monografia** — isso é feito depois via `/claudiney` ou `/tarefa`.

## Modos de entrada

`$ARGUMENTS` aceita:

| Entrada | Exemplo |
|---|---|
| DOI | `10.1007/978-3-642-25566-3_40` |
| arXiv ID | `2301.08745` ou `arXiv:2301.08745` |
| Chave BibTeX | `kennedy1995particle` |
| Caminho de PDF | `vault/papers/pdfs/meu-artigo.pdf` ou caminho absoluto |
| Palavra `vault` | Processa todos os PDFs em `vault/papers/pdfs/` sem nota correspondente |
| Palavra `fila` | Processa PDFs em `.opencode/log/pdf-events.log` ainda não incorporados |

Sem argumentos, lista PDFs pendentes e pergunta qual processar.

## Fluxo completo

### Fase 1 — Obter PDF

- **Se DOI recebido**: baixa via cascata MCP — `doiget_fetch_paper` (OA primeiro) → `scihub` → `bash scripts/download-pdfs.sh --keys <key>`. Salva em `vault/papers/pdfs/<key>.pdf`.
- **Se arXiv ID recebido** (prefixo `arxiv:` detectado): `arxiv_download_paper(id)` → salva PDF. Depois extrai metadados via `arxiv_get_abstract(id)`.
- **Se chave BibTeX recebida**: busca a chave em `monografia/bib/abntex2-references.bib`, extrai DOI. Procede com o caminho DOI. Se não encontrada, busca por título no `crossref`.
- **Se caminho de PDF recebido**: usa diretamente. Se estiver fora do vault, copia para `vault/papers/pdfs/<key>.pdf`.
- **Se "vault" ou "fila"**: itera sobre PDFs pendentes, processa um por um.
- **Se TODAS as fontes falharem** (paywall + Sci-Hub offline + script indisponível): cria nota com `pdf_status: ausente` e prossegue com metadados do Crossref. Não bloqueia a pipeline.

### Fase 2 — Extrair metadados

- Executa `bash scripts/extract-pdf-doi.sh <pdf>` para extrair DOI do PDF.
- **Branch por tipo de identificador**:
  - Se começa com `10.` (DOI): `crossref_get_work(doi)` + `doiget_resolve_paper(doi)` + `scholar-sidekick_checkOpenAccess(doi)`
  - Se começa com `arxiv:`: usa `arxiv_get_abstract(id)` + `arxiv_download_paper(id)`
  - Se nenhum: `pdf-reader` extrai texto das primeiras 3 páginas → busca título no `academic-search_search_papers` + `crossref_search_works` → confirma com usuário
- **Fonte primária de BibTeX**: `scholar-sidekick_exportCitation(doi, format='bib')` (stateless, sempre funciona com DOI).
- **Fallback BibTeX**: `doiget_bibtex_export(doi)` apenas se entrada existir no store local do doiget.

### Fase 3 — Gerar resumo

- `pdf-reader` extrai texto completo do PDF.
- Lê abstract, introduction e conclusion.
- Produz resumo estruturado em **português (PT-BR)** com 3-5 frases cobrindo: problema, método, resultado principal.
- Extrai 1-3 **citações-chave** literais (frases impactantes do paper).
- Infere automaticamente:
  - `role`: `fundacional`, `revisao`, `comparativo`, `metodologico`, `aplicacao`, `benchmark`, `limitacao`, `trabalho-futuro`
  - `areas`: ex: `tsp`, `drone-routing`, `lower-bound`, `routing`
  - `methods`: ex: `ga`, `pso`, `aco`, `exact`, `metaheuristic`, `machine-learning`
  - Tags hierárquicas correspondentes

### Fase 4 — Integrar BibTeX

- Verifica se a chave já existe em `monografia/bib/abntex2-references.bib`.
- Se não existe, adiciona entrada formatada (ABNT-style, com `doi`, autores completos, escaping LaTeX).
- Se a chave existe mas metadados diferem, pergunta se deve atualizar.

### Fase 5 — Criar/atualizar nota no vault

- Usa `bash scripts/import-bib-to-vault.sh <key>` para criar o esqueleto da nota (se ainda não existe).
- **IMPORTANTE**: usa `bibtex-key` (com hífen), **não** `bibtex_key` (com underscore). O schema canônico e todos os scripts usam `bibtex-key`.
- Preenche frontmatter: `title`, `authors`, `year`, `doi`, `bibtex-key`, `pdf`, `areas`, `methods`, `role`, `reading_status: resumo-lido`, `pdf_status: disponivel`, `rating`.
- Preenche seções:
  - **PDF**: substitui `<!-- PDF não disponível -->` por `![[papers/pdfs/<key>.pdf]]`
  - **Tese Central**: 1-2 frases
  - **Resumo**: 3-5 frases em PT-BR
  - **Contribuições Principais**: bullet list
  - **Relevância para o TCC**: conexão com drones/rTSP/metaheurísticas
  - **Métodos e Abordagens**: técnicas usadas
  - **Citações-chave**: 1-3 citações literais com `>`
  - **Conexões**: `[[wikilinks]]` para áreas e papers relacionados
- Se a nota já existe, **atualiza** (não sobrescreve) campos vazios.
- **Heurística de "campo vazio"**: null/0/[] no YAML = vazio; seção com apenas `<!-- ... -->` ou `-` = vazio.
- Executa `python3 scripts/migrate-tags.py` ao final para sincronizar tags hierárquicas com propriedades YAML.

### Fase 6 — Atualizar canvas de conhecimento

- Lê `vault/canvas/tcc-knowledge-graph.canvas`.
- Adiciona nó (`type: "file"`) para o paper se não existir.
- Posiciona abaixo do último nó existente (+70~80 em y) para evitar sobreposição.
- Adiciona arestas para áreas e métodos relevantes.
- Preserva estrutura existente do canvas.

### Fase 7 — Verificar e atualizar claims

- Lê `vault/claims/*.md` (não `claims.base` — isso é apenas view).
- **Filtro**: verifica apenas claims cujas áreas/métodos batem com as do paper (matching por keywords no título do claim vs abstract). Limitar a top-10 claims mais relevantes.
- Para cada claim relevante:
  - **Apoia** → adiciona `claim_support: [<ID>]` no frontmatter
  - **Contradiz** → registra na seção `Limitações de Uso`
  - **É fonte primária** → sugere vincular como `primary_evidence`
- Claims com `strength: forte` e `status: validado` **não** são alterados automaticamente.
- Sugere novos claims relevantes (não cria).

### Fase 8 — Conectar com papers existentes

- Obtém referências do paper via `crossref_get_references(doi)`.
- Para cada referência com DOI, busca correspondência no vault:
  - `grep -l "<doi>" vault/papers/*.md` para encontrar notas com o mesmo DOI
  - Se encontrado, adiciona `[[wikilink]]` bidirecional
- Infere similaridade por palavras-chave compartilhadas e sugere conexões adicionais.

### Fase 9 — Sugerir próximos passos

- Se o paper é altamente relevante (rating ≥ 4): sugere `/tarefa` para criar item no roadmap
- Lista claims afetados.
- Mostra resumo: `[BibTeX] [Vault] [Canvas] [Claims] [Conexões]`.

### Fase 10 — Registrar no log do roadmap

- Se paper foi incorporado com sucesso:
```bash
bash scripts/roadmap.sh log incorporation \
  bibtex_key=<key> \
  doi=<doi> \
  pdf_obtido=sim \
  nota_criada=sim \
  canvas_atualizado=sim
```
- Se paper ficou pendente de enriquecimento, também cria tarefa na fila:
```bash
bash scripts/roadmap.sh tarefa criar "Enriquecer nota: <key>" "Preencher seções pendentes em vault/papers/<key>.md" media literatura
```

## Dependências

- **MCPs**: `crossref`, `doiget`, `academic-search`, `scholar-sidekick`, `pdf-reader`, `scihub`, `arxiv`
- **Scripts**: `scripts/extract-pdf-doi.sh`, `scripts/download-pdfs.sh`, `scripts/import-bib-to-vault.sh`, `scripts/roadmap.sh`, `scripts/migrate-tags.py`
- **Skills**: `knowledge-base`, `vault-semantic-schema`, `vault-tagger`, `vault-bases-maintainer`

## Exemplos

```
/incorporar 10.1007/978-3-642-25566-3_40
/incorporar 2301.08745
/incorporar kennedy1995particle
/incorporar vault/papers/pdfs/novo-drone-survey.pdf
/incorporar vault
/incorporar fila
```

## Anti-patterns

- Não altera arquivos `.tex` da monografia (use `/tarefa` ou `/claudiney` para isso).
- Não cria claims automaticamente — apenas sugere.
- Não modifica notas de outros papers além de adicionar wikilinks bidirecionais.
- Não força preenchimento de seções que exigem leitura profunda.
- Respeita o schema do `vault-semantic-schema` — tags hierárquicas, sem flat tags.
- Usa `bibtex-key` (hífen), não `bibtex_key` (underscore).

## Relação com knowledge-base skill

O skill `knowledge-base` cobre **descoberta** de papers (busca, gap analysis, validação de DOI). O comando `/incorporar` cobre **ingestão** (PDF → vault → canvas → claims → log). Use o skill para encontrar novos papers; use o comando para processá-los.
