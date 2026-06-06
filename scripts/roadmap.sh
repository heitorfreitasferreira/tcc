#!/usr/bin/env bash
set -euo pipefail
# scripts/roadmap.sh
# Camada fina de query/log para o sistema de roadmap (fila de tarefas + eventos)
#
# Uso:
#   scripts/roadmap.sh tarefa criar "<titulo>" "<saida>" [alta|media|baixa] [fase]  → cria P<N>.md
#   scripts/roadmap.sh tarefa listar [fase]                                        → lista pendentes
#   scripts/roadmap.sh tarefa concluir P<N>                                        → marca concluída
#   scripts/roadmap.sh log incorporation key=val key=val ...                       → loga evento
#   scripts/roadmap.sh log experiment key=val key=val ...                          → loga evento
#   scripts/roadmap.sh log compilation key=val key=val ...                         → loga evento
#   scripts/roadmap.sh proximo                                                     → mostra próxima tarefa

REPO="$(cd "$(dirname "$0")/.." && pwd)"
TAREFAS_DIR="$REPO/vault/roadmap/tarefas"
EVENTOS_DIR="$REPO/vault/roadmap/eventos"
ROADMAP_README="$REPO/vault/roadmap/README.md"

mkdir -p "$TAREFAS_DIR" "$EVENTOS_DIR"

cmd="${1:-}"; shift || true

# ─── tarefa criar ──────────────────────────────────────────────
if [[ "$cmd" == "tarefa" && "${1:-}" == "criar" ]]; then
    shift
    titulo="${1:?Uso: roadmap.sh tarefa criar \"<titulo>\" \"<saida>\" [alta|media|baixa] [fase]}"; shift
    saida="${1:-}"; shift || true
    prioridade="${1:-media}"; shift || true
    fase="${1:-escrita}"; shift || true

    case "$fase" in
        infra|literatura|experimentacao|analise|escrita|polimento|revisao) ;;
        *) echo "Erro: fase inválida '$fase'. Use uma das: infra, literatura, experimentacao, analise, escrita, polimento, revisao" >&2; exit 1 ;;
    esac

    # Next ID
    last=$(ls "$TAREFAS_DIR"/P*.md 2>/dev/null | sed 's/.*P//;s/\.md//' | sort -n | tail -1)
    next=$(( ${last:-0} + 1 ))
    id="P$next"
    hoje=$(date -I)

    cat > "$TAREFAS_DIR/${id}.md" << NOTE
---
type: tarefa
task_id: $id
title: "$titulo"
status: pendente
priority: $prioridade
fase: $fase
ordem: $next
dependencias: []
origin: manual
criado_em: $hoje
saida_esperada: "$saida"
tags:
  - tipo/tarefa
  - status/pendente
  - origem/manual
  - fase/$fase
---

## $titulo

**Saída esperada**: $saida
**Prioridade**: $prioridade
**Criado em**: $hoje
NOTE

    echo "✅ Tarefa criada: $id — $titulo ($fase, $prioridade)"
    echo "   $TAREFAS_DIR/${id}.md"
    exit 0
fi

# ─── tarefa listar ─────────────────────────────────────────────
if [[ "$cmd" == "tarefa" && "${1:-}" == "listar" ]]; then
    shift || true
    filter_fase="${1:-}"
    
    if [[ -n "$filter_fase" ]]; then
        echo "## Tarefas pendentes — fase: $filter_fase"
        echo ""
    else
        echo "## Tarefas pendentes"
        echo ""
    fi

    for f in "$TAREFAS_DIR"/P*.md; do
        [[ -f "$f" ]] || continue
        status=$(grep -m1 '^status:' "$f" | sed 's/status: *//')
        [[ "$status" != "pendente" ]] && continue
        tid=$(grep -m1 '^task_id:' "$f" | sed 's/task_id: *//')
        ttl=$(grep -m1 '^title:' "$f" | sed 's/title: *"//;s/"$//')
        ffs=$(grep -m1 '^fase:' "$f" | sed 's/fase: *//')
        pri=$(grep -m1 '^priority:' "$f" | sed 's/priority: *//')
        [[ -n "$filter_fase" && "$ffs" != "$filter_fase" ]] && continue
        echo "| $tid | $ttl | $ffs | $pri |"
    done
    exit 0
fi

# ─── tarefa concluir ───────────────────────────────────────────
if [[ "$cmd" == "tarefa" && "${1:-}" == "concluir" ]]; then
    shift
    tid="${1:?Uso: roadmap.sh tarefa concluir P<N>}"
    file="$TAREFAS_DIR/${tid}.md"
    [[ -f "$file" ]] || { echo "Erro: tarefa não encontrada: $file" >&2; exit 1; }
    hoje=$(date -I)
    sed -i "s/^status: .*/status: concluida/" "$file"
    sed -i "s/^  - status\/.*/  - status\/concluida/" "$file"
    if ! grep -q '^concluido_em:' "$file"; then
        sed -i "/^criado_em:/a concluido_em: $hoje" "$file"
    else
        sed -i "s/^concluido_em: .*/concluido_em: $hoje/" "$file"
    fi
    echo "✅ $tid marcada como concluída"
    exit 0
fi

# ─── proximo ───────────────────────────────────────────────────
if [[ "$cmd" == "proximo" ]]; then
    echo "## Próxima tarefa pendente"
    echo ""

    best_file=""
    best_fase_order="zzz99999"
    for f in "$TAREFAS_DIR"/P*.md; do
        [[ -f "$f" ]] || continue
        status=$(grep -m1 '^status:' "$f" | sed 's/status: *//')
        [[ "$status" != "pendente" ]] && continue

        ffs=$(grep -m1 '^fase:' "$f" | sed 's/fase: *//')
        ord=$(grep -m1 '^ordem:' "$f" | sed 's/ordem: *//')

        fase_order_val=999
        case "$ffs" in
            infra) fase_order_val=0 ;;
            literatura) fase_order_val=100 ;;
            experimentacao) fase_order_val=200 ;;
            analise) fase_order_val=300 ;;
            escrita) fase_order_val=400 ;;
            polimento) fase_order_val=500 ;;
            revisao) fase_order_val=600 ;;
        esac
        key=$(printf "%03d%05d" "$fase_order_val" "${ord:-0}")

        if [[ "$key" < "$best_fase_order" ]]; then
            best_fase_order="$key"
            best_file="$f"
        fi
    done

    if [[ -z "$best_file" ]]; then
        echo "Nenhuma tarefa pendente."
        exit 0
    fi

    tid=$(grep -m1 '^task_id:' "$best_file" | sed 's/task_id: *//')
    ttl=$(grep -m1 '^title:' "$best_file" | sed 's/title: *"//;s/"$//')
    ffs=$(grep -m1 '^fase:' "$best_file" | sed 's/fase: *//')
    pri=$(grep -m1 '^priority:' "$best_file" | sed 's/priority: *//')
    saida=$(grep -m1 '^saida_esperada:' "$best_file" | sed 's/saida_esperada: *"//;s/"$//')
    deps=$(grep -m1 '^dependencias:' "$best_file" | sed 's/dependencias: *//')

    echo "| Campo | Valor |"
    echo "|---|---|"
    echo "| ID | $tid |"
    echo "| Tarefa | $ttl |"
    echo "| Fase | $ffs |"
    echo "| Prioridade | $pri |"
    echo "| Saída esperada | $saida |"
    echo "| Dependências | ${deps:-nenhuma} |"
    exit 0
fi

# ─── log ────────────────────────────────────────────────────────
if [[ "$cmd" == "log" ]]; then
    tipo="${1:?Uso: roadmap.sh log <incorporation|experiment|compilation> key=val ...}"; shift

    hoje=$(date -I)
    ts=$(date -Iseconds)
    slug="${tipo}-${hoje}"

    # Coleta k=v extras
    extras=""
    for kv in "$@"; do
        extras="$extras$kv\n"
    done

    # Gera nome único
    n=$(ls "$EVENTOS_DIR"/${slug}*.md 2>/dev/null | wc -l)
    n=$((n + 1))
    filepath="$EVENTOS_DIR/${slug}-${n}.md"

    case "$tipo" in
        incorporation)
            cat > "$filepath" << NOTE
---
type: log
log_type: incorporation
date: $hoje
timestamp: $ts
$extras
tags:
  - tipo/log
  - log/incorporation
---

## Incorporação — $hoje

$(for kv in "$@"; do echo "- **${kv%%=*}**: ${kv#*=}"; done)
NOTE
            ;;
        experiment)
            cat > "$filepath" << NOTE
---
type: log
log_type: experiment
date: $hoje
timestamp: $ts
$extras
tags:
  - tipo/log
  - log/experiment
---

## Experimento — $hoje

$(for kv in "$@"; do echo "- **${kv%%=*}**: ${kv#*=}"; done)
NOTE
            ;;
        compilation)
            cat > "$filepath" << NOTE
---
type: log
log_type: compilation
date: $hoje
timestamp: $ts
$extras
tags:
  - tipo/log
  - log/compilation
---

## Compilação — $hoje

$(for kv in "$@"; do echo "- **${kv%%=*}**: ${kv#*=}"; done)
NOTE
            ;;
        *)
            echo "Erro: tipo de log desconhecido: $tipo" >&2
            exit 1
            ;;
    esac

    echo "✅ Evento logado: $filepath"
    exit 0
fi

# ─── help ───────────────────────────────────────────────────────
cat << HELP
roadmap.sh — camada fina de query/log do sistema de fila

SUBCOMANDOS:
  tarefa criar "<titulo>" "<saida>" [prio] [fase]  Cria nova tarefa P<N>
  tarefa listar [fase]                              Lista tarefas pendentes
  tarefa concluir P<N>                              Marca tarefa como concluída
  proximo                                           Mostra próxima tarefa pendente
  log incorporation key=val ...                     Loga incorporação de paper
  log experiment key=val ...                        Loga experimento
  log compilation key=val ...                       Loga compilação
HELP
