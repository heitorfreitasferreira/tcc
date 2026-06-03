#!/usr/bin/env bash
# Gera figuras de rotas para a monografia usando o endpoint /api/render.
# Uso: bash scripts/gerar-figuras.sh [addr]
# Exemplo: bash scripts/gerar-figuras.sh :9191
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
FIGS_DIR="$REPO_DIR/monografia/figs"
SUMMARY_DIR="$REPO_DIR/src/data/results/summary"
ADDR="${1:-:9191}"

mkdir -p "$FIGS_DIR"

if [ ! -x "$REPO_DIR/src/tcc" ]; then
    make -C "$REPO_DIR/src" build
fi

echo "=== Iniciando servidor em $ADDR ==="
"$REPO_DIR/src/tcc" serve --addr "$ADDR" &
SERVER_PID=$!
sleep 2

cleanup() { kill "$SERVER_PID" 2>/dev/null || true; }
trap cleanup EXIT

PORT="${ADDR##*:}"
BASE="http://localhost:$PORT"

run_id() {
    local instance="$1" method="$2" seed="${3:-0}"
    local file
    file=$(ls "$SUMMARY_DIR" 2>/dev/null | grep "^${instance}__${method}__s${seed}__" | sort | head -1 || true)
    if [ -z "$file" ]; then
        return 1
    fi
    printf '%s\n' "${file%.json}"
}

fetch() {
    local name="$1" url="$2"
    echo "  -> $name"
    local status
    status=$(curl -s -o "$FIGS_DIR/$name" -w "%{http_code}" "$BASE$url")
    local size
    size=$(wc -c < "$FIGS_DIR/$name")
    echo "     HTTP $status ${size}B"
    if [ "$status" != "200" ]; then
        rm -f "$FIGS_DIR/$name"
        return 1
    fi
}

fetch_run() {
    local instance="$1" method="$2" seed="${3:-0}"
    local rid
    if rid=$(run_id "$instance" "$method" "$seed"); then
        fetch "route-${instance}-${method}-s${seed}.png" "/api/render?map=${instance}&run=${rid}"
    else
        echo "  -> sem run: ${instance} ${method} seed ${seed}"
    fi
}

fetch_iter() {
    local instance="$1" method="$2" iteration="$3" seed="${4:-0}"
    local rid
    if rid=$(run_id "$instance" "$method" "$seed"); then
        fetch "route-${instance}-${method}-s${seed}-iter${iteration}.png" "/api/render?map=${instance}&run=${rid}&iteration=${iteration}"
    fi
}

fetch_overlay() {
    local instance="$1"
    local aco ga pso lb bf url
    aco=$(run_id "$instance" aco 0 || true)
    ga=$(run_id "$instance" ga 0 || true)
    pso=$(run_id "$instance" pso 0 || true)
    lb=$(run_id "$instance" lowerbound 0 || true)
    bf=$(run_id "$instance" bruteforce 0 || true)

    if [ -z "$aco" ] || [ -z "$ga" ] || [ -z "$pso" ]; then
        echo "  -> sem overlay: ${instance}"
        return 0
    fi

    url="/api/render?map=${instance}&run=${aco}&run2=${ga}&run3=${pso}"
    if [ -n "$lb" ]; then
        url="${url}&run4=${lb}"
    fi
    if [ -n "$bf" ]; then
        url="${url}&run5=${bf}"
    fi
    fetch "overlay-${instance}-methods.png" "$url"
}

echo "=== Gerando figuras de rotas ==="

representative_instances=(10a 10b 10c 14a 15a 20a 30a 50a 100a)
small_instances=(10a 10b 10c 11a 11b 11c 12a 12b 12c 13a 13b 13c 14a 14b 14c)

for instance in "${representative_instances[@]}"; do
    fetch "graph-${instance}-bare.png" "/api/render?map=${instance}"
    fetch_run "$instance" aco 0
    fetch_run "$instance" ga 0
    fetch_run "$instance" pso 0
    fetch_run "$instance" lowerbound 0
    fetch_overlay "$instance"
done

for instance in 10a 30a 50a 100a; do
    fetch_iter "$instance" aco 1 0
    fetch_iter "$instance" aco 25 0
    fetch_iter "$instance" aco 50 0
done

for instance in "${small_instances[@]}"; do
    fetch_run "$instance" bruteforce 0
done

echo ""
echo "=== Figuras geradas em $FIGS_DIR ==="
ls -lh "$FIGS_DIR"/*.png 2>/dev/null || true
echo ""
echo "Total: $(ls "$FIGS_DIR"/*.png 2>/dev/null | wc -l) PNGs"
