# AGENTS.md — Índice do Projeto

## Project Context

- TCC/monografia comparando métodos de otimização bioinspirados sobre TSP/rTSP; cenário motivador é patrulha de drones visitando pontos de interesse com rota/tempo mínimos.
- Manter código, dados de experimento, notebooks de análise e texto da monografia alinhados.

## Repository Shape

- `src/` — módulo Go (`module tcc`, Go 1.23.7), CLI Cobra; executar comandos Go de `src/` ou `make -C src ...`
- `src/data/` — `.points`, `.graph` e resultados; servidor web embarca `*.graph`, `*.points`, `results/summary`, `results/evolution`, `results/timing`
- `monografia/` — LaTeX usando `ppgco.cls`; capítulos em `cap_*`, referências em `bib/abntex2-references.bib`

## Guias por área

| Arquivo | Conteúdo |
|---------|----------|
| [AGENTS-visualizacao.md](./AGENTS-visualizacao.md) | Geração de imagens, figuras, gráficos e artefatos visuais |
| [AGENTS-experiments.md](./AGENTS-experiments.md) | CLI Go, workflow de experimentos, estrutura de dados |
| [AGENTS-vault.md](./AGENTS-vault.md) | Vault Obsidian, skills instaladas, workflow de referências |
| [AGENTS-mcp.md](./AGENTS-mcp.md) | Servidores MCP acadêmicos, scripts auxiliares, env vars |
| [AGENTS-monografia.md](./AGENTS-monografia.md) | Notas de escrita e compilação |
