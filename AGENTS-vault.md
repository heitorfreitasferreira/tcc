# AGENTS — Vault Obsidian

## Skill do Projeto

- `knowledge-base` (`.agents/skills/knowledge-base/SKILL.md`) — skill para expansão do vault: busca de papers, validação de DOIs, importação BibTeX, enriquecimento de notas, atualização de canvas.

## Skills Instaladas

- `obsidian-markdown` — formatação markdown compatível com Obsidian (callouts, [[links]])
- `obsidian-vault` — operações no vault (criar, ler, buscar notas)
- `obsidian-bases` — queries estruturadas sobre o vault
- `json-canvas` — criar/atualizar grafos de conhecimento `.canvas`
- `academic-researcher` — buscar e entender papers acadêmicos
- `academic-search` — estratégias de busca acadêmica

## MCPs Disponíveis (básicos)

- `scihub` — baixar PDFs e metadados de papers
- `google-scholar` — buscar artigos acadêmicos

## Estrutura do Vault

```
vault/
├── papers/           ← Uma nota markdown por artigo
├── areas/            ← Notas sobre áreas de pesquisa
├── templates/        ← Template para criar novas notas
├── canvas/           ← Grafo de conhecimento (.canvas)
└── opencode-vault.md ← Instruções completas para agentes
```

## Workflow para Novas Referências

1. Use `google-scholar` + `academic-search` skill para encontrar papers
2. Use `scihub` para baixar metadata/PDF (ou `bash scripts/download-pdfs.sh` para lote)
3. Adicione entrada em `monografia/bib/abntex2-references.bib`
4. Execute `bash scripts/import-bib-to-vault.sh <bibtex-key>` para criar nota
5. Use skill `obsidian-markdown` para escrever resumo/contribuições
6. Link para áreas existentes com `[[wiki links]]`
7. Atualize canvas em `vault/canvas/tcc-knowledge-graph.canvas`

## Importar BibTeX Existente

```bash
./scripts/import-bib-to-vault.sh                          # todas as entradas
./scripts/import-bib-to-vault.sh kennedy1995particle      # entrada única
BIB=path/to/file.bib ./scripts/import-bib-to-vault.sh     # .bib customizado
```
