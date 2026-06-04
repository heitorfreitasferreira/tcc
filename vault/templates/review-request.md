---
title: ""
tags:
  - tipo/revisao
  - status/aberto
review_id: ""
status: aberto
priority: ""
source: ""
date_opened: ""
date_closed: ""
target_chapter: ""
correction_layers: []
evidence_layer: ""
claim_ids: []
verified_by_script: ""
aliases: []
---

<!--
Instruções para o agente:

1. Preencher YAML acima.
2. Na seção "Solicitação Original", transcrever o feedback do orientador.
3. Na seção "Análise Técnica", investigar o problema contra código, dados, literatura e vault.
4. Na seção "Plano de Correção", definir a abordagem mais determinística possível.
5. Na seção "Verificação", definir critério de aceite e script de validação.
6. Executar o plano de correção e marcar como resolvido.
-->

## Solicitação Original

<!-- Transcrição literal do feedback do orientador -->

> ...

## Análise Técnica

### Problema Identificado

<!-- O que exatamente está errado ou precisa mudar? -->

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/` (código) | ... | ... |
| `src/data/` (dados) | ... | ... |
| Literatura | ... | ... |
| `vault/` | ... | ... |
| `monografia/` (texto atual) | ... | ... |

### Causa Raiz

<!-- O que causou o problema? Cópia manual? Desatualização? Interpretação incorreta? -->

## Plano de Correção

### Abordagem Preferida (mais determinística)

<!--
Estrutura da solução na hierarquia de confiança:
camada 1 = alterar código/scripts geradores → regenerar métricas
camada 2 = script determinístico de transformação (rg, awk, Python)
camada 3 = edição agentica com prompt estruturado
-->

- [ ] **Camada 1 — Código/Dados:** ...
- [ ] **Camada 2 — Script determinístico:** ...
- [ ] **Camada 3 — Edição agentica:** ...

### Comandos e Passos

```bash
# Comandos shell para executar a correção
```

### Artefatos Afetados

- `monografia/...` — o que muda

## Verificação

### Critério de Aceite

<!-- O que precisa ser verdade para considerar resolvido? -->

### Script de Validação

```bash
# Comando que retorna exit 0 se resolvido, exit 1 se não
```

## Evidências

<!-- Logs, screenshots, outputs de script anexados durante ou após a correção -->

## Notas

<!-- Decisões, discussões, justificativas -->
