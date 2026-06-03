#!/usr/bin/env bash
# Gera os artefatos visuais principais em padrao academico.
# Uso: bash scripts/gerar-analises.sh [addr]
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
FIGS_DIR="$REPO_DIR/monografia/figs"
ADDR="${1:-:9191}"

mkdir -p "$FIGS_DIR"

echo "=== Limpando figuras geradas antigas ==="
python3 - << 'PY'
from pathlib import Path

figs = Path('monografia/figs')
keep = {
    'facom-logo.pdf',
    'logo-ufu.png',
    'logo-ufu2.png',
    'logo-ufu3.png',
    'ppgco-logo.png',
}
for path in figs.iterdir():
    if path.name in keep:
        continue
    if path.suffix.lower() in {'.png', '.svg', '.pdf', '.tex'}:
        path.unlink()
PY

echo "=== Compilando servidor web ==="
make -C "$REPO_DIR/src" build

echo "=== Iniciando servidor em $ADDR ==="
"$REPO_DIR/src/tcc" serve --addr "$ADDR" &
SERVER_PID=$!
cleanup() { kill "$SERVER_PID" 2>/dev/null || true; }
trap cleanup EXIT
sleep 2

PORT="${ADDR##*:}"
BASE="http://localhost:$PORT"

echo "=== Diagramas TikZ de metodos ==="
python3 "$SCRIPT_DIR/gerar-diagramas-conceituais.py"

echo "=== Figuras cientificas compostas ==="
FIGURE_RENDER_BASE="$BASE" python3 "$SCRIPT_DIR/gerar-figuras-publicacao.py"

echo "=== Analises concluídas ==="
python3 - << 'PY'
from pathlib import Path
figs = Path('monografia/figs')
for ext in ['.svg', '.png', '.tex', '.pdf']:
    print(ext, len(list(figs.glob(f'*{ext}'))))
PY
