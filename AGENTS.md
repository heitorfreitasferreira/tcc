# AGENTS.md — Índice do Projeto

## Project Context

- TCC/monografia comparando métodos de otimização bioinspirados sobre TSP/rTSP; cenário motivador é patrulha de drones visitando pontos de interesse com rota/tempo mínimos.
- Manter código, dados de experimento, notebooks de análise e texto da monografia alinhados.

## Repository Shape

- `src/` — módulo Go (`module tcc`, Go 1.23.7), CLI Cobra; executar comandos Go de `src/` ou `make -C src ...`
- `src/data/` — `.points`, `.graph` e resultados; servidor web embarca `*.graph`, `*.points`, `results/summary`, `results/evolution`, `results/timing`
- `monografia/` — LaTeX usando `ppgco.cls`; capítulos em `cap_*`, referências em `bib/abntex2-references.bib`

## Guias por área

| Skill | Conteúdo |
|-------|----------|
| `tcc-escrita` | Redação acadêmica (prosa, estrutura, LaTeX) — use ao escrever qualquer parte da monografia |
| `stop-slop` | Remove AI tells da prosa — carregar ao revisar texto |
| `monograph-notes` | Compilação LaTeX, ciclo BibTeX, idioma pt-BR |
| `experiment-workflow` | CLI Go, workflow de experimentos, schemas de dados |
| `analise-estatistica` | Testes estatísticos, effect sizes, power analysis, reporte APA |
| `figuras-tcc` | Figuras e gráficos para a monografia (Tufte + convenções do projeto) |
| `vault-tcc` | Vault Obsidian completo: estrutura, tags, Bases, schema, importação de papers |
| `go` | Go idiomático, CLI e Cobra — sub-arquivos `idiomatic.md`, `cli.md`, `cobra.md` conforme o escopo |
| `academic-mcp` | Servidores MCP acadêmicos, scripts auxiliares, env vars |

## Escrita da Monografia

Use o subagente `@tcc-redator` para escrever, revisar ou compilar qualquer
parte da monografia. Ou carregue as skills individuais: `tcc-escrita`,
`stop-slop`, `monograph-notes`.

## Experimentos

Use o subagente `@tcc-experimentos` para orquestrar benchmarks, análises
estatísticas e geração de figuras. Ou carregue as skills individuais:
`experiment-workflow`, `analise-estatistica`, `figuras-tcc`.
