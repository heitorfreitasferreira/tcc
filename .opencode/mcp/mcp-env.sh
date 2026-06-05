#!/usr/bin/env bash
# Carrega .env da raiz do projeto e executa o comando MCP
set -a
source "$(dirname "$0")/../../.env"
set +a
exec "$@"
