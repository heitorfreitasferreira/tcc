---
description: Revisa mudanças e commita tudo com mensagem gerada
---

# /commitar

Revisa o working tree, pergunta se deseja commitar, e executa ou sugere.

## Fluxo

### 1. Mostrar estado atual

```bash
git status --short
git diff --stat
git log --oneline -5
```

Se não há nada a commitar, encerra.

### 2. Gerar mensagem de commit

Mensagem curta no formato `<tipo>: <descrição>` baseada nos arquivos alterados.

### 3. Perguntar

Use a toolcall `question` com exatamente:

- header: "Commitar?"
- question: "Mensagem: <mensagem>\n\nConfirma?"
- options: `["SIM"]`, `["NAO"]`

### 4. SIM

```bash
git add -A
git commit -m "<mensagem>"
```

### 5. NAO

Escreva em texto: "Sugestão de commit:\n\n```\ngit add -A\ngit commit -m \"<mensagem>\"\n```"

Não execute nada. Encerre.
