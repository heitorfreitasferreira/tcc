---
description: Cria tarefa manual na fila ordenada do roadmap
---

# /tarefa

Cria uma tarefa manual no sistema de fila (`vault/roadmap/tarefas/P<N>.md`).

## Uso

```
/tarefa "Título da tarefa" "Descrição" [prioridade] [fase]
```

Parâmetros posicionais:

| # | Campo | Padrão | Valores |
|---|---|---|---|
| 1 | Título | obrigatório | Descrição curta |
| 2 | Descrição | vazio | Texto livre com detalhes |
| 3 | Prioridade | `media` | `alta`, `media`, `baixa` |
| 4 | Fase | `escrita` | `infra`, `literatura`, `experimentacao`, `analise`, `escrita`, `polimento`, `revisao` |

## Comportamento

1. Determina próximo número P<N> sequencial.
2. Cria `vault/roadmap/tarefas/P<N>.md` com frontmatter padronizado.
3. Se já existe tarefa com título idêntico na mesma fase, pergunta se é duplicata.
4. Registra no log de eventos.

## Exemplos

```
/tarefa "Revisar seção de Metodologia" "Ler cap3 e verificar consistência com claims" alta escrita
/tarefa "Buscar artigo sobre ACO para TSP"
/tarefa "Corrigir ortografia na introdução" media polimento
```
