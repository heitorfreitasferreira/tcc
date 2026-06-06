---
title: Sistema de Roadmap
type: writing
tags:
  - tipo/writing
  - topico/roadmap
  - topico/ferramenta
aliases:
  - roadmap
  - fila de tarefas
  - task queue
---

# Sistema de Roadmap — Fila Ordenada de Tarefas

O roadmap do TCC é um sistema de **fila ordenada** com dois pontos de entrada e um consumidor.

## Estrutura

```
vault/roadmap/
├── README.md              ← este arquivo (instruções)
├── tarefas/P<N>.md        ← tarefas (uma nota por item)
└── eventos/*.md           ← log de execução (uma nota por evento)

vault/bases/
├── roadmap-tarefas.base   ← views: Próxima, Pendentes, Kanban
└── roadmap-log.base       ← views: Papers, Experimentos, Compilações
```

## Como usar

### Criar tarefa (feedback do orientador)
```
/claudiney "O orientador disse que a tabela de gap está errada..."
```
→ Cria review note em `review-solicitacoes/` + 1+ tarefas em `tarefas/`

### Criar tarefa (ideia própria)
```
/tarefa "Revisar seção de ACO na fundamentação"
```
→ Cria 1 tarefa em `tarefas/`

### Executar próxima tarefa
```
/proximo
```
→ Mostra a primeira tarefa pendente na fila, pergunta confirmação, executa, valida com council, marca concluída.

### Ver PDFs pendentes de incorporação
```
/incorporar vault
```
→ Processa PDFs em `vault/papers/pdfs/` sem nota correspondente.

### Compilar monografia
```
/compilar
```
→ `check-monografia.sh` + `pdflatex` + log.

### Rodar experimentos
```
/experimento --method ga --frequency 10:3 --seeds 0-2
```
→ Roda batch + log.

## Fases (ordem de execução)

| Fase | Ordem | Descrição |
|---|---|---|
| `infra` | 0 | Setup, migração, estrutura |
| `literatura` | 100 | Incorporar papers, fechar bibliografia |
| `experimentacao` | 200 | Rodar experimentos |
| `analise` | 300 | Análise estatística, métricas |
| `escrita` | 400 | Escrever capítulos |
| `polimento` | 500 | Validação, claims, anti-alucinação |
| `revisao` | 600 | Feedback do orientador |

`/proximo` sempre pega a tarefa com menor `(fase, ordem)`.

## Comandos do ecossistema

| Comando | Função |
|---|---|
| `/consultar` | Buscar literatura (MCPs em paralelo) |
| `/incorporar` | PDF → vault + BibTeX + canvas + claims |
| `/tarefa` | Criar tarefa manual |
| `/claudiney` | Feedback do orientador → review + tarefas |
| `/proximo` | Executar próxima tarefa da fila |
| `/compilar` | Compilar + validar monografia |
| `/experimento` | Rodar experimentos batch |

## Script auxiliar

```bash
scripts/roadmap.sh tarefa criar "titulo" "saida" [prio] [fase]
scripts/roadmap.sh tarefa listar [fase]
scripts/roadmap.sh proximo
scripts/roadmap.sh log incorporation key=val ...
```

## Referência

- Planejamento detalhado: `vault/writing/planejamento/roadmap-monografia.md`
- Claims: `vault/claims/` + `vault/bases/claims.base`
- Papers: `vault/papers/` + `vault/bases/papers.base`
- Siglas: `vault/siglas/` + `vault/bases/siglas.base`
