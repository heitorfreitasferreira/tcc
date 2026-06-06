#!/usr/bin/env bash
set -euo pipefail
# scripts/migrate-roadmap.sh
# Extrai a tabela P<N> do roadmap-monografia.md e gera notas .md em vault/roadmap/tarefas/
# Uso: bash scripts/migrate-roadmap.sh

REPO="$(cd "$(dirname "$0")/.." && pwd)"
SRC="$REPO/vault/writing/planejamento/roadmap-monografia.md"
DST="$REPO/vault/roadmap/tarefas"

mkdir -p "$DST"

if [[ ! -f "$SRC" ]]; then
    echo "Erro: arquivo fonte não encontrado: $SRC" >&2
    exit 1
fi

# Extrai linhas da tabela P<N> (formato: | P<N> | titulo | saida | prioridade |)
# Pula o header da tabela
grep -P '^\| P\d+' "$SRC" | while IFS='|' read -r _ id title saida prioridade _; do
    id=$(echo "$id" | xargs)
    title=$(echo "$title" | xargs)
    saida=$(echo "$saida" | xargs)
    prioridade=$(echo "$prioridade" | xargs)

    # Normaliza prioridade
    prioridade_lower=$(echo "$prioridade" | tr '[:upper:]' '[:lower:]')
    case "$prioridade_lower" in
        concluída|concluida|concluído|concluido) status="concluida" ;;
        pendente) status="pendente" ;;
        *) status="pendente" ;;
    esac

    # Infere fase baseado no conteúdo do título
    fase=""
    title_lower=$(echo "$title" | tr '[:upper:]' '[:lower:]')
    if echo "$title_lower" | grep -qE 'claim|rastreável|rastre'; then fase="infra"
    elif echo "$title_lower" | grep -qE 'bib|paper|índice|index|bibliogr|literatura|pdf|referência|biblioteca|citação|leitura'; then fase="literatura"
    elif echo "$title_lower" | grep -qE 'experiment|execução|rodar|lower.?bound|ótimo|brute|instância|semente'; then fase="experimentacao"
    elif echo "$title_lower" | grep -qE 'estatístic|análise|metric|friedman|nemenyi|wilcoxon|cd-diagram'; then fase="analise"
    elif echo "$title_lower" | grep -qE 'capítulo|capa|folha|rosto|resumo|abstract|sigla|gloss|abnt|tex|latex|escrita|escrever|map[ae]|figura|tabela|apêndice'; then fase="escrita"
    elif echo "$title_lower" | grep -qE 'anti.?alucina|validação|validacao|compil|check|gate|consistência|formatação|revisão|revisao|limpeza|pream'; then fase="polimento"
    elif echo "$title_lower" | grep -qE 'orientador|claudiney|feedback'; then fase="revisao"
    elif echo "$title_lower" | grep -qE 'originalidade|feromônio|aco|ant.?system|mmas|codificação'; then fase="escrita"
    else fase="escrita"
    fi

    # Ajustes manuais para casos específicos
    case "$id" in
        P1|P2|P3|P5|P6) fase="infra" ;;
        P7) fase="infra" ;;
        P8) fase="experimentacao" ;;
        P9|P10) fase="analise" ;;
        P11) fase="polimento" ;;
        P12|P13) fase="escrita" ;;
        P14) fase="escrita" ;;
        P15) fase="polimento" ;;
        P16|P17|P18) fase="literatura" ;;
        P19|P20|P21) fase="experimentacao" ;;
        P22|P23) fase="polimento" ;;
        P24|P25) fase="polimento" ;;
        P26) fase="literatura" ;;
        P27) fase="escrita" ;;
        P28) fase="literatura" ;;
        P29|P30|P31) fase="literatura" ;;
        P32|P33|P34) fase="escrita" ;;
        P35) fase="literatura" ;;
        P36) fase="literatura" ;;
        P37|P38|P39) fase="polimento" ;;
        P40|P41) fase="polimento" ;;
        P42|P43|P44|P45|P46) fase="polimento" ;;
        P47) fase="polimento" ;;
    esac

    # Extrai o número da tarefa
    num=$(echo "$id" | grep -oP '\d+')
    ordem="$num"

    # Gera a nota .md
    filepath="$DST/${id}.md"
    cat > "$filepath" << NOTE
---
type: tarefa
task_id: $id
title: "$title"
status: $status
priority: $(echo "$prioridade" | tr '[:upper:]' '[:lower:]' | sed 's/ .*//')
fase: $fase
ordem: $ordem
dependencias: []
origin: manual
criado_em: 2026-06-02
saida_esperada: "$saida"
tags:
  - tipo/tarefa
  - status/$status
  - origem/manual
  - fase/$fase
---

## $title

**Saída esperada**: $saida

**Prioridade**: $prioridade

**Fase**: $fase

*Nota migrada de \`vault/writing/planejamento/roadmap-monografia.md\` em $(date -I).*
NOTE

    echo "  [$id] $title → $filepath ($status, $fase)"
done

echo ""
echo "[migrate-roadmap] Concluído. Notas em $DST/"
echo "[migrate-roadmap] O arquivo original em $SRC NÃO foi alterado."
