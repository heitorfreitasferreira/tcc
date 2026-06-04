---
description: Cria um novo item no roadmap da monografia, com ID incremental
---

You are running the `/roadmap-criar` command. Your task is to add a new pending item to the roadmap.

## Passos

### 1. Entender o item

Use `question` tool to ask the user:

- **Título da tarefa** — short PT-BR description (e.g. "Revisar seção de ACO na Fundamentação")
- **Saída esperada** — what artifact/file/result should exist when done
- **Prioridade** — `Alta`, `Média`, or `Baixa`

Continue asking until you have all three pieces.

### 2. Determinar o próximo ID

1. Read `vault/writing/planejamento/roadmap-monografia.md`
2. Find all `| P<number>` rows in the `## Tarefas Preparatórias Para Agentes` section
3. Get the highest P number, increment by 1
4. If no P rows exist, start at P1

### 3. Inserir no roadmap

1. Find the last `| P<number>` row in the table (right before the blank line or next section)
2. Insert a new row **after it** in sorted order (by P number):
   `| P<N> | <título> | <saída esperada> | Pendente |`
3. Update the `updated` field in the frontmatter to today's date (`date +%Y-%m-%d`)

### 4. Confirmar e sugerir commit

Show the user what was added and suggest:

```
## Pacote de Commit

### Arquivo alterado
- vault/writing/planejamento/roadmap-monografia.md — adicionado P<N>: <título>

### Commands
git add vault/writing/planejamento/roadmap-monografia.md
git diff --cached
git commit -m "roadmap: adiciona P<N> — <título>"
```

Do NOT commit. Just suggest.
