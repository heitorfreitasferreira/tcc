# Componente: Roadmap e Fila de Tarefas

## Responsabilidade

Manter uma fila ordenada de trabalho em `vault/roadmap/tarefas/`, onde cada item e uma nota Markdown com frontmatter padronizado.

## Estrutura

```text
vault/roadmap/
├── README.md
├── tarefas/P<N>.md
└── eventos/*.md
```

## Formato de Tarefa

```yaml
---
type: tarefa
task_id: P48
title: "Titulo curto"
status: pendente
priority: alta
fase: literatura
ordem: 48
dependencias: [P18]
origin: manual
criado_em: 2026-06-05
saida_esperada: "Resultado verificavel"
tags:
  - tipo/tarefa
  - status/pendente
  - origem/manual
  - fase/literatura
---
```

## Fases

| Fase | Ordem | Uso |
|---|---:|---|
| `infra` | 0 | Estrutura, scripts, migracoes |
| `literatura` | 100 | Papers, PDFs, BibTeX, vault |
| `experimentacao` | 200 | Execucao de algoritmos |
| `analise` | 300 | Estatistica e metricas |
| `escrita` | 400 | Texto dos capitulos |
| `polimento` | 500 | Validacao, formatacao, resumo |
| `revisao` | 600 | Feedback externo |

## Comandos Shell

```bash
bash scripts/roadmap.sh tarefa criar "Titulo" "Saida" alta literatura
bash scripts/roadmap.sh tarefa listar literatura
bash scripts/roadmap.sh proximo
bash scripts/roadmap.sh tarefa concluir P48
bash scripts/roadmap.sh log incorporation bibtex_key=key status=ok
```

## Regras

- `/proximo` consome a menor chave `(fase, ordem)` entre tarefas pendentes.
- Dependencias nao sao resolvidas automaticamente pelo shell; o agente deve verificar antes de executar.
- `origin` indica a entrada: `manual`, `claudiney`, `incorporar` ou `watcher`.
- Eventos sao historicos; tarefas sao estado operacional.

## Bases Obsidian

- `vault/bases/roadmap-tarefas.base`: views de pendentes, concluidas, kanban e proxima tarefa.
- `vault/bases/roadmap-log.base`: eventos de incorporacao, experimento e compilacao.
