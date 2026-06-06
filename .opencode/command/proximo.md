---
description: Executa a próxima tarefa pendente da fila ordenada do roadmap
---

# /proximo

Pega a primeira tarefa pendente da fila (ordenada por `fase, ordem`), mostra contexto, pergunta confirmação, executa, valida com council, e marca como concluída.

Substitui o antigo `/roadmap-consumir`.

## Fluxo

### 1. Buscar próxima tarefa
```bash
bash scripts/roadmap.sh proximo
```
Extrai: ID, título, fase, prioridade, saída esperada, dependências.

### 2. Verificar dependências
Se a tarefa tem `dependencias: [P26]` e P26 não está concluída:
- Avisa: "Esta tarefa depende de P26 (não concluída)."
- Pergunta: "[e]xecutar mesmo assim / [p]ular para a próxima / [c]ancelar"

### 3. Apresentar contexto
Mostra:
- ID, título, fase, prioridade
- Origem (manual / claudiney / incorporar / watcher)
- Se originada de `/claudiney`: link para a review note
- Saída esperada
- Artefatos existentes relacionados

### 4. Confirmar
```
Executar P18: "Fechar bibliografia mínima por capítulo"?
[s]im  [p]ular  [v]er detalhes  [c]ancelar
```

### 5. Executar
- Carrega skills relevantes para a fase:
  - `literatura` → knowledge-base, vault-semantic-schema, vault-tagger
  - `experimentacao` → data-analysis, go
  - `escrita` → academic-writing, scientific-paper, latex-document-skill
  - `polimento` → vault-semantic-schema, latex-document-skill
- Se a tarefa tem `acao: incorporar` → delega para pipeline `/incorporar`
- Realiza o trabalho conforme `saida_esperada`

### 6. Validar
- Council com 2 juízes (modo default):
  - Se tarefa é simples → pergunta se deseja pular validação
  - Se FAIL → registra bloqueio, **não** marca concluída, sugere `/claudiney`

### 7. Concluir
```bash
bash scripts/roadmap.sh tarefa concluir P<N>
```
- Se `origin == claudiney`: atualiza review note com `[[P<N>]]`
- Se `origin == watcher`: loga incorporação em `vault/roadmap/eventos/`
- Se `origin == incorporar`: atualiza nota do paper

### 8. Sugerir próximo
Mostra a próxima tarefa pendente após esta.

## Ordem da fila

Tarefas são ordenadas por `fase` (infra < literatura < experimentacao < analise < escrita < polimento < revisao) e depois por `ordem` numérica.
