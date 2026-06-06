# Council Validation Report — Sistema de Fila

**Date**: 2026-06-05
**Mode**: quick (single-agent inline)
**Target**: Implementação do sistema de fila (roadmap migrado + 6 comandos + thin shell)

---

## Result: WARN

**Confidence**: HIGH

---

## Resumo

A implementação está funcional e internamente consistente, com 1 WARN documental e 2 WARN de completude.

---

## Findings

### F1 [WARN] — `tarefa.md` default de fase incorreto

A documentação de `/tarefa` (`.opencode/command/tarefa.md:22`) lista o default da fase como `pendente`, mas:
- `pendente` **não é uma fase válida** no sistema (é um status)
- O valor real default em `scripts/roadmap.sh:30` é `escrita`
- As fases válidas são: `infra`, `literatura`, `experimentacao`, `analise`, `escrita`, `polimento`, `revisao`

**Impacto**: Baixo — a documentação está errada, mas o script implementa o comportamento correto. Um leitor humano se confundiria.

**Correção**: Trocar `pendente` por `escrita` na tabela de parâmetros de `tarefa.md`.

### F2 [WARN] — `roadmap.sh tarefa criar` não valida `fase`

O script aceita qualquer string como fase (ex: `fase: "invalida"`). A função `proximo` no mesmo script usa um `case` statement com fallback 999 para fases desconhecidas, então uma fase inválida faria a tarefa sumir da fila (sempre ordenada após `revisao`).

**Impacto**: Médio — silencioso. Tarefas com fase inválida seriam criadas mas nunca encontradas por `proximo`.

**Correção**: Adicionar validação no `tarefa criar`:

```bash
case "$fase" in
    infra|literatura|experimentacao|analise|escrita|polimento|revisao) ;;
    *) echo "Erro: fase inválida '$fase'. Use uma das: infra, literatura, experimentacao, analise, escrita, polimento, revisao" >&2; exit 1 ;;
esac
```

### F3 [WARN] — `roadmap.sh tarefa criar` não aceita `origin` parameter

O `/claudiney` precisa pós-editar o arquivo da tarefa para setar `origin: claudiney` + `origin_ref`. Seria mais limpo aceitar `--origin claudiney` opcional.

**Impacto**: Baixo — funciona, mas é frágil (2-step em vez de 1-step).

---

## PASS Findings

### P1 — Consistência entre comandos e AGENTS.md

AGENTS.md lista 7 comandos ativos; todos os 7 têm `.opencode/command/<nome>.md`. Os 4 deprecados estão sinalizados com redirect. ✅

### P2 — Ordenação de fase

`roadmap.sh` usa infra=0, literatura=100, ..., revisao=600. Idêntico ao AGENTS.md. ✅

### P3 — `list-pending-roadmap.sh` migrado

Agora lê de `tarefas/P*.md` em vez do markdown table obsoleto. Funciona com o novo YAML frontmatter. ✅

### P4 — Deprecated commands

`roadmap-criar.md`, `roadmap-consumir.md`, `consultar-academico.md` todos apontam para o substituto correto. ✅

### P5 — `/incorporar` Fase 10 registra log

O comando chama `roadmap.sh log incorporation` que existe no script. ✅

### P6 — `proximo` ordena por (fase_order, ordem)

Usa chave composta `printf "%03d%05d"`. Sem race conditions. ✅

### P7 — `tarefa criar` ID sequencial

`ls P*.md | sort -n | tail -1` + 1. Funciona com 0 tarefas (next=1). ✅

---

## Verdict

| Aspecto | Resultado |
|---------|-----------|
| Funcionalidade core | ✅ Funcional |
| Consistência entre comandos | ✅ Consistente |
| Consistência doc ↔ script | ⚠️ WARN (F1) |
| Robustez | ⚠️ WARN (F2, F3) |
| Deprecação | ✅ Completa |
| Dados migrados (47 tarefas) | ✅ Corretos |
| Shell script | ✅ Sem bugs |

**Conclusão**: Sistema funcional e pronto para uso. Recomendo corrigir F1 (typo no doc) e F2 (validação de fase) antes de escalar o uso para terceiros. F3 é opcional.
