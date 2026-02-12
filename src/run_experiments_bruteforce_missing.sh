#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
run_all_script="${script_dir}/run_all.sh"

seed="0"
min_size="15"
frequency_input=""
results_dir=""
data_folder="${script_dir}/data"
if_exists="skip"
progress="false"
jobs=""

usage() {
	cat <<EOF
Usage: $0 [options]

Options:
  --seed=<int>          Seed para o bruteforce (padrao: 0)
  --min-size=<int>      Tamanho minimo da instancia (padrao: 15)
  --frequency=<spec>    Frequencia no formato n:q (sobrescreve auto-deteccao)
  --results-dir=<path>  Diretorio de resultados
  --folder=<path>       Pasta com .graph (padrao: src/data)
  --if-exists=<mode>    Modo para execucoes existentes (padrao: skip)
  --progress=<bool>     Exibir progresso no binario (padrao: false)
  --jobs=<int>          Paralelismo interno do run_all
  -h, --help            Mostra esta ajuda
EOF
}

build_frequency_from_graphs() {
	local folder="$1"
	local min_required_size="$2"
	declare -A counts=()
	local path file base size

	shopt -s nullglob
	for path in "${folder}"/*.graph; do
		file="${path##*/}"
		base="${file%.graph}"

		if [[ ! "${base}" =~ ^([0-9]+)[A-Za-z]+$ ]]; then
			continue
		fi

		size="${BASH_REMATCH[1]}"
		if ((10#${size} < 10#${min_required_size})); then
			continue
		fi

		counts["${size}"]=$(( ${counts["${size}"]:-0} + 1 ))
	done
	shopt -u nullglob

	if [[ ${#counts[@]} -eq 0 ]]; then
		return 1
	fi

	mapfile -t sorted_sizes < <(printf '%s\n' "${!counts[@]}" | sort -n)

	local pairs=()
	local current_size
	for current_size in "${sorted_sizes[@]}"; do
		pairs+=("${current_size}:${counts["${current_size}"]}")
	done

	local IFS=,
	echo "${pairs[*]}"
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
	--min-size=*)
		min_size="${1#*=}"
		;;
	--frequency=*)
		frequency_input="${1#*=}"
		;;
	--results-dir=*)
		results_dir="${1#*=}"
		;;
	--folder=*)
		data_folder="${1#*=}"
		;;
	--if-exists=*)
		if_exists="${1#*=}"
		;;
	--progress=*)
		progress="${1#*=}"
		;;
	--jobs=*)
		jobs="${1#*=}"
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
	echo "Invalid --seed: ${seed}. Expected an integer." >&2
	exit 1
fi

if [[ ! "${min_size}" =~ ^[0-9]+$ ]]; then
	echo "Invalid --min-size: ${min_size}. Expected a positive integer." >&2
	exit 1
fi

if [[ ! -x "${run_all_script}" ]]; then
	echo "run_all.sh not found or not executable at ${run_all_script}" >&2
	exit 1
fi

if [[ ! -d "${data_folder}" ]]; then
	echo "Data folder not found: ${data_folder}" >&2
	exit 1
fi

if [[ -z "${frequency_input}" ]]; then
	if ! frequency_input="$(build_frequency_from_graphs "${data_folder}" "${min_size}")"; then
		echo "Could not derive frequency for min-size ${min_size} in ${data_folder}." >&2
		exit 1
	fi
fi

echo "=== Bruteforce para instancias sem baseline (>= ${min_size}) ==="
echo "Seed: ${seed}"
echo "Frequencia: ${frequency_input}"
echo "Data folder: ${data_folder}"
echo "if_exists: ${if_exists}"

cmd=(
	"${run_all_script}"
	"--method=bruteforce"
	"--seed=${seed}"
	"--frequency=${frequency_input}"
	"--folder=${data_folder}"
	"--if-exists=${if_exists}"
	"--progress=${progress}"
)

if [[ -n "${results_dir}" ]]; then
	cmd+=("--results-dir=${results_dir}")
fi

if [[ -n "${jobs}" ]]; then
	cmd+=("--jobs=${jobs}")
fi

"${cmd[@]}"

echo "Concluido: bruteforce para instancias >= ${min_size} finalizado."
