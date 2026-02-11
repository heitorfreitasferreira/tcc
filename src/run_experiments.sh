#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
run_all_script="${script_dir}/run_all.sh"

if [[ ! -x "${run_all_script}" ]]; then
	echo "run_all.sh not found or not executable at ${run_all_script}" >&2
	exit 1
fi

echo "=== Cenario 1: metodos ga, pso, aco com frequencia padrao ==="
for method in ga pso aco; do
	echo "-> Rodando metodo: ${method}"
	"${run_all_script}" "--method=${method}"
done

scenario_two_frequency="15:3,20:3,30:3,50:3,100:3"

echo "=== Cenario 2: metodos ga, pso, aco para frequencia ${scenario_two_frequency} ==="
for method in ga pso aco; do
	echo "-> Rodando metodo: ${method}"
	"${run_all_script}" "--method=${method}" "--frequency=${scenario_two_frequency}"
done

echo "=== Bruteforce por ultimo: frequencia padrao ==="
echo "-> Rodando metodo: bruteforce"
"${run_all_script}" "--method=bruteforce"

echo "Concluido: os dois cenarios foram executados, com bruteforce no final."
