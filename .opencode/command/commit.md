---
description: Analisa mudanças e gera N commits atômicos com conventional commits
---

# /commit

Analisa o working tree, agrupa mudanças logicamente, e gera 1-N commits atômicos
com mensagens no formato Conventional Commits.

## Fluxo

### 1. Mostrar estado atual

```bash
git status --short
git diff --stat
git log --oneline -5
```

Se não há nada a commitar, encerra com "Nada a commitar."

### 2. Analisar e agrupar mudanças

Use `git diff` e `git diff --cached` para ler o conteúdo completo das mudanças.
Identifique grupos lógicos atômicos baseados nos arquivos alterados:

- Mudanças no mesmo pacote/diretório → mesmo commit
- Refatorações de tipos/dados compartilhados → mesmo commit (mesmo que toque vários pacotes)
- Funcionalidades independentes → commits separados
- Arquivos novos vs modificados em áreas diferentes → commits separados

Regras:
- Cada commit deve ser compilável isoladamente (`go build ./...` passa)
- Mensagem no formato `<tipo>(<escopo>): <descrição>`
- Tipos: feat, fix, refactor, docs, ci, reorg, test, data
- Escopo opcional: graph, ga, pso, aco, brute, lowerbound, web, cmd, monografia, data, workflow

Exemplo de grouping: `src/graph/` + `src/optimization/` modificados → considerar se
a mudança em optimization é consequência direta da mudança em graph (se sim,
mesmo commit; se não, commits separados).

### 3. Apresentar plano

Mostre ao usuário os commits planejados:

```
Commit 1: feat(graph): descrição
  src/graph/types.go       +10 -2
  src/graph/creater.go      +5 -3

Commit 2: fix(aco): descrição
  src/optimization/aco/ant.go  +3 -8

Confirmar? [SIM/NAO/EDITAR]
```

Use a toolcall `question` com:
- header: "Commitar?"
- question: mostrando os N commits + diff summary
- options: ["SIM", "NAO, mostrar comandos"]

### 4. SIM

Para cada commit:
```bash
git add <arquivos>
git commit -m "<tipo>(<escopo>): <descrição>"
```

### 5. NAO

Escreva os comandos git para o usuário executar manualmente. Não execute nada.

### 6. Ao final

Mostre o log resultante:
```bash
git log --oneline -5
```
