# Council Report — Validação do Fluxo `/incorporar`

**Data**: 2026-06-05  
**Modo**: validate (3 juízes, sem perspectivas)  
**Consenso**: **FAIL** (Judge-1: FAIL, Judge-2: WARN, Judge-3: WARN)

---

## Sumário

O `/incorporar` é um comando ambicioso que unifica 9 fases: PDF → metadados → resumo → BibTeX → vault → canvas → claims → conexões → sugestões. A arquitetura conceitual está correta e o escopo é bem definido. Porém, 5 falhas críticas de correção e integração bloqueiam sua execução confiável, e 8 problemas significantes precisam de ajuste antes do uso em produção.

---

## Falhas Críticas (bloqueantes)

| # | Severidade | Descrição | Juízes | Ref |
|---|---|---|---|---|
| 1 | critical | **bibtex_key / bibtex-key duplicados**: template e Fase 5 usam `bibtex_key`, schema canônico e todos os 91+ scripts usam `bibtex-key`. Notas criadas via `/incorporar` terão `bibtex-key` vazio, quebrando scripts e Bases. | J1, J2 | `vault/templates/paper-note.md:6-7`, `.opencode/command/incorporar.md:66` |
| 2 | critical | **arxiv: prefixo tratado como DOI**: `extract-pdf-doi.sh` retorna `arxiv:2301...`, mas a Fase 2 passa para Crossref/doiget (que rejeitam). Sem branch condicional, pipeline colapsa para PDFs com arXiv ID. | J1 | `scripts/extract-pdf-doi.sh:40-44`, `.opencode/command/incorporar.md:34-43` |
| 3 | critical | **doiget_bibtex_export dependência circular**: só funciona se entrada já existe no store do doiget. Se PDF veio via Sci-Hub (fallback) ou é local, retorna `null`. Crossref já tem os metadados, mas não são usados para gerar BibTeX. | J1, J2 | `.opencode/command/incorporar.md:28,43` |
| 4 | critical | **semantic-scholar MCP não disponível**: listado como dependência da Fase 2, mas não está nos MCPs ativos. O equivalente é `academic-search_search_papers` (assinatura diferente). | J2 | `.opencode/command/incorporar.md:41` |
| 5 | critical | **Sem caminho de aborto para PDF inalcançável**: se todas as fontes falharem (paywall + Sci-Hub offline + script ausente), pipeline não tem terminal state definido. | J2 | `.opencode/command/incorporar.md:26-30` |

## Problemas Significantes

| # | Severidade | Descrição | Juízes |
|---|---|---|---|
| 6 | significant | **Pipeline duplicado com knowledge-base skill**: skill cobre 7/9 fases sem que o comando declare a relação. Agente que carregue ambos fará trabalho redundante/conflitante. | J1, J3 |
| 7 | significant | **Fase 7 (claims) computacionalmente inviável**: verificar 60 claims por paper sem heurísticas de filtro consumiria >30K tokens. | J1, J2 |
| 8 | significant | **import-bib-to-vault.sh contornado**: script canônico de criação de notas é ignorado, criando dois caminhos divergentes. | J3 |
| 9 | significant | **Skills inexistentes listadas**: `obsidian-bases` e `json-canvas` não existem como skills locais. Corretos: `vault-bases-maintainer`. | J3 |
| 10 | significant | **Resolução BibTeX key → DOI ausente**: quando `$ARGUMENTS` é chave BibTeX, não há passo para extrair DOI do `.bib`. | J2 |
| 11 | significant | **Heurística de "campos vazios" ambígua**: Fase 5 diz "atualiza campos vazios" sem definir o que é vazio (placeholder HTML? bullet `-`? campo null?). | J2 |
| 12 | significant | **DOI → nome de arquivo vault sem reverse lookup**: Fase 8 cruza referências (DOIs) com vault (bibtex-key filenames) sem mecanismo de matching. | J2 |
| 13 | significant | **Embed de PDF com sintaxe ambígua**: `![[pdfs/<key>.pdf]]` vs placeholder `<!-- PDF não disponível -->`. Comportamento de insert vs replace não especificado. | J1 |

## Problemas Menores

| # | Severidade | Descrição |
|---|---|---|
| 14 | minor | **Sem tratamento de erros**: Nenhuma fase especifica timeouts, retry, ou graceful degradation. |
| 15 | minor | **Log pdf-events.log não inicializado**: diretório vazio, `/incorporar fila` leria arquivo inexistente. |
| 16 | minor | **Formato de log inconsistente**: plugin e script escrevem no mesmo arquivo com formatos diferentes. |
| 17 | minor | **Coordenadas do canvas inconsistentes**: posições na Fase 6 não batem com layout real (surveys ~x:400, não ~x:620). |
| 18 | minor | **Role `trabalho-futuro` ausente**: lista de roles inferíveis omite valor válido do schema. |
| 19 | minor | **migrate-tags.py não executado**: vault-semantic-schema exige verificação pós-criação. |
| 20 | minor | **DOIs de referências capturados como falso positivo**: regex nas primeiras 3 páginas pode capturar DOIs de citações, não do paper. |
| 21 | minor | **Sem rollback em falha parcial**: BibTeX, vault note e canvas já mutados se pipeline falhar na Fase 7. |

---

## Recomendação

**Corrigir as 5 falhas críticas antes do primeiro uso:**

1. Remover `bibtex_key` do template e corrigir Fase 5 para `bibtex-key`
2. Adicionar branch `arxiv:` na Fase 2 com ferramentas arXiv
3. Substituir `doiget_bibtex_export` por `scholar-sidekick_exportCitation(doi, format='bib')` como fonte primária
4. Substituir `semantic-scholar` por `academic-search_search_papers` na Fase 2
5. Adicionar caminho de aborto com criação de nota "apenas metadados" quando PDF inalcançável

**Tratar em seguida (problemas significantes 6-13):**
- Declarar relação `/incorporar` ↔ `knowledge-base` skill (descoberta vs ingestão)
- Adicionar heurísticas de filtro na Fase 7 (matching por keywords, limitar a top-10 claims)
- Integrar `import-bib-to-vault.sh` e `migrate-tags.py` no fluxo
- Documentar heurística de "campo vazio" e comportamento de embed do PDF

---

## Arquivos dos juízes

- [Judge 1](./2026-06-05-validate-fluxo-incorporar-judge-1.md) — Correção e integridade
- [Judge 2](./2026-06-05-validate-fluxo-incorporar-judge-2.md) — Execução prática
- [Judge 3](./2026-06-05-validate-fluxo-incorporar-judge-3.md) — Integração e coesão
