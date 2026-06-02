#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
run_all_script="${script_dir}/run_all.sh"
seed="0"

usage() {
	echo "Usage: $0 [--seed=<int64>]"
}

while [[ $# -gt 0 ]]; do
	case "$1" in
	--seed=*)
		seed="${1#*=}"
		;;
	--seed|-s)
		if [[ $# -lt 2 ]]; then
			echo "Missing value for $1" >&2
			usage >&2
			exit 1
		fi
		seed="$2"
		shift
		;;
	-h|--help)
		usage
		exit 0
		;;
	*)
		echo "Unknown option: $1" >&2
		usage >&2
		exit 1
		;;
	esac
	shift
done

if [[ ! "${seed}" =~ ^-?[0-9]+$ ]]; then
	echo "Invalid seed: ${seed}. Expected an integer." >&2
	exit 1
fi

if [[ ! -x "${run_all_script}" ]]; then
	echo "run_all.sh not found or not executable at ${run_all_script}" >&2
	exit 1
fi

echo "=== Cenario 1: metodos ga, pso, aco, lowerbound com frequencia padrao (seed=${seed}) ==="
for method in ga pso aco lowerbound; do
	echo "-> Rodando metodo: ${method}"
	"${run_all_script}" "--method=${method}" "--seed=${seed}"
done

scenario_two_frequency="15:3,20:3,30:3,50:3,100:3"

echo "=== Cenario 2: metodos ga, pso, aco, lowerbound para frequencia ${scenario_two_frequency} (seed=${seed}) ==="
for method in ga pso aco lowerbound; do
	echo "-> Rodando metodo: ${method}"
	"${run_all_script}" "--method=${method}" "--seed=${seed}" "--frequency=${scenario_two_frequency}"
done

echo "=== Bruteforce por ultimo: frequencia padrao ==="
echo "-> Rodando metodo: bruteforce"
"${run_all_script}" "--method=bruteforce" "--seed=${seed}"

echo "Concluido: os dois cenarios foram executados, com bruteforce no final."
