---
description: Consome um item do roadmap: clarifica, executa, valida com council, marca concluído, sugere commit
---

You are running the `/roadmap-consumir` command. Your task is to execute a pending roadmap item end-to-end:

**clarificar → executar → validar (council) → marcar → sugerir commit**

## Fase 1: Mostrar pendentes

Run `bash scripts/list-pending-roadmap.sh` and present the list to the user.

## Fase 2: Clarificar com o usuário

Use `question` tool to ask:

1. **Qual item?** — Offer the pending list, plus "Próximo pendente (recomendado)"
2. **O que fazer?** — Scope, requirements, design decisions
3. **Preferências?** — Format, location, style constraints

Keep asking until you have enough clarity. If the roadmap's "Saída esperada" is vague, ask the user to elaborate.

## Fase 3: Executar

1. Read relevant files (roadmap's "Saída esperada" + related vault notes + code)
2. Make the changes
3. Verify (compile, test, lint when applicable)

## Fase 4: Validar com council

1. Build a packet describing what was changed and why
2. Spawn **2 independent judges** in parallel:
   - `task(subagent_type="general")` for each
   - Each judge reads modified files and checks correctness
   - Output: `.agents/council/YYYY-MM-DD-<id>-judge-{1,2}.md`
3. Collect results, consolidate (PASS/FAIL/WARN)
4. Write report to `.agents/council/YYYY-MM-DD-<id>.md`

If verdict is FAIL or WARN with actionable issues, report to the user and stop. Do not mark as done.

## Fase 5: Marcar no roadmap

1. Edit `vault/writing/roadmap-monografia.md`
2. Change status from `Pendente` to `Concluída`
3. Update `updated` in frontmatter to today

## Fase 6: Sugerir pacote de commit

Show the user:

```
## Pacote de Commit

### Arquivos alterados
- caminho/arquivo: descrição

### Commands
git add <arquivos>
git diff --cached
git commit -m "tipo: descrição concisa"
```

Do NOT commit. Just suggest.
