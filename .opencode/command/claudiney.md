---
description: Registra feedback do orientador com análise contextual e gera tarefas na fila do roadmap
---

You are running the /claudiney command to register advisor feedback about the monograph. Your task is to:

## Fase 1 — Parsear o feedback

Understand what the advisor is asking for. Extract concrete action items.

## Fase 2 — Criar scaffold da review

Run `bash scripts/registrar-review-scaffold.sh "<short-description>"` (generate a short kebab-case description from the feedback). This creates a file at `vault/writing/review-solicitacoes/R<NN>-<desc>.md`. Capture the filename from the output.

## Fase 3 — Analisar o contexto

Search the codebase for context:
- `src/` — relevant code, types, algorithms
- `src/data/results/` — relevant experiment data
- `vault/` — notes in `vault/writing/`, `vault/papers/`, `vault/areas/`
- `monografia/` — .tex files on the topic
- `vault/claims/` and `vault/bases/claims.base` — affected claims

## Fase 4 — Preencher a review note

Fill in:
- YAML `title:` with a short PT-BR description
- YAML `priority:` based on impact (alta/media/baixa)
- YAML `target_chapter:` which chapter is affected
- YAML `evidence_layer:` where the truth lives (codigo/dados/literatura/vault/monografia)
- YAML `tarefas_geradas:` (will be filled in Fase 5)
- Section "Solicitação Original" with a clear restatement
- Section "Análise Técnica" with root cause analysis
- Section "Plano de Correção" with actionable steps

## Fase 5 — ★ GERAR TAREFAS NA FILA ★

Para cada ação concreta identificada no plano de correção:
1. Executa `bash scripts/roadmap.sh tarefa criar "<título>" "<saída esperada>" <prioridade> <fase>`
2. Captura o ID gerado (ex: P50, P51, P52)
3. Adiciona `[[P50]]`, `[[P51]]` na review note

Depois, atualiza o YAML da review note:
```yaml
tarefas_geradas:
  - "[[P50]]"
  - "[[P51]]"
origin_ref nas tarefas: "[[R<NN>]]"
```

As tarefas são criadas com:
- `origin: claudiney`
- `origin_ref: "[[R<NN>]]"`
- `fase` inferida do conteúdo (literatura, experimentacao, escrita, polimento)

Para cada tarefa criada, edite o arquivo da tarefa (`vault/roadmap/tarefas/P<N>.md`) e ajuste:
- `origin: claudiney`
- Adicionar `origin_ref: "[[R<NN>]]"` no YAML

## Fase 6 — Reportar

Show summary:
```
## Feedback registrado

**Review**: R<NN> — vault/writing/review-solicitacoes/R<NN>-<desc>.md
**Tarefas criadas na fila**:
- P50: Revisar tabela X (fase: polimento, prioridade: alta)
- P51: Corrigir citação Y (fase: literatura, prioridade: alta)

**Próximo passo**: Execute /proximo para processar a fila.
```

## Notas

- `/claudiney` é exclusivo para feedback do orientador. Para criar uma tarefa sem análise contextual, use `/tarefa`.
- As tarefas entram na fila ordenada. `/proximo` as processa em ordem de `(fase, ordem)`.
- Se o feedback for vago, pergunte antes de gerar tarefas.
