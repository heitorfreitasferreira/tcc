# Componente: Skills

## Responsabilidade

Skills carregam instrucoes especializadas para cada fase ou dominio. Elas complementam os comandos agenticos.

## Skills por Fase

| Fase | Skills principais |
|---|---|
| `literatura` | `knowledge-base`, `vault-semantic-schema`, `vault-tagger` |
| `experimentacao` | `data-analysis`, `go` |
| `analise` | `data-analysis`, `statistical-analysis` |
| `escrita` | `academic-writing`, `scientific-paper`, `latex-document-skill` |
| `polimento` | `vault-semantic-schema`, `latex-document-skill` |
| `revisao` | `council`, `academic-writing`, `scientific-paper` |

## Skills de Vault

- `vault-semantic-schema`: frontmatter, relacoes, claims e propriedades.
- `vault-tagger`: taxonomia de tags hierarquicas.
- `knowledge-base`: descoberta e enriquecimento de papers.

## Skills de Escrita

- `academic-writing`: prosa academica.
- `scientific-paper`: estrutura cientifica e LaTeX.
- `research-paper-writing`: clareza, fluxo e apresentacao para revisores.
- `latex-document-skill`: compilacao, figuras, tabelas e documentos PDF.

## Skills de Codigo e Experimentos

- `go`: padroes idiomaticos de Go.
- `golang-cli`: arquitetura CLI.
- `data-analysis`: analise estatistica e relatorios.
- `statistical-analysis`: testes, intervalos, efeito e APA.

## Regra Operacional

`/proximo` deve carregar skills conforme a fase antes de executar a tarefa. Para tarefas hibridas, carregar a skill da fonte primaria da evidencia.
