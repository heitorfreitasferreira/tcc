#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
tcc_bin="${script_dir}/tcc"

seed="0"
min_size="15"
results_dir="${script_dir}/data/results"
data_folder="${script_dir}/data"
if_exists="skip"
progress="false"
jobs="$(nproc)"

usage() {
	cat <<EOF
Usage: $0 [options]

Options:
  --seed=<int>          Seed do bruteforce (padrao: 0)
  --min-size=<int>      Tamanho minimo da instancia (padrao: 15)
  --results-dir=<path>  Diretorio de resultados (padrao: src/data/results)
  --folder=<path>       Pasta com .graph (padrao: src/data)
  --if-exists=<mode>    skip|overwrite|error (padrao: skip)
  --progress=<bool>     Exibir progresso no binario (padrao: false)
  --jobs=<int>          Paralelismo das instancias (padrao: nproc)
  -h, --help            Mostra esta ajuda

Comportamento padrao:
  - roda bruteforce apenas 1 vez por instancia
  - se ja existir qualquer summary de bruteforce para a instancia, pula (if-exists=skip)
EOF
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

if [[ ! "${jobs}" =~ ^[1-9][0-9]*$ ]]; then
	echo "Invalid --jobs: ${jobs}. Expected a positive integer." >&2
	exit 1
fi

if [[ ! "${if_exists}" =~ ^(skip|overwrite|error)$ ]]; then
	echo "Invalid --if-exists: ${if_exists}. Allowed: skip|overwrite|error" >&2
	exit 1
fi

if [[ ! -x "${tcc_bin}" ]]; then
	echo "Binary not found at ${tcc_bin}. Build it first with 'make -C src build' or 'go build -o ./tcc' inside src/." >&2
	exit 1
fi

if [[ ! -d "${data_folder}" ]]; then
	echo "Data folder not found: ${data_folder}" >&2
	exit 1
fi

mkdir -p "${results_dir}/logs"

declare -a files=()

shopt -s nullglob
for path in "${data_folder}"/*.graph; do
	file="${path##*/}"
	base="${file%.graph}"

	if [[ ! "${base}" =~ ^([0-9]+)[A-Za-z]+$ ]]; then
		continue
	fi

	size="${BASH_REMATCH[1]}"
	if ((10#${size} < 10#${min_size})); then
		continue
	fi

	files+=("${path}")
done
shopt -u nullglob

if [[ ${#files[@]} -eq 0 ]]; then
	echo "Nenhum .graph encontrado com n >= ${min_size} em ${data_folder}." >&2
	exit 0
fi

mapfile -t files < <(printf '%s\n' "${files[@]}" | sort -V)

echo "=== Bruteforce para instancias >= ${min_size} ==="
echo "Seed: ${seed}"
echo "Data folder: ${data_folder}"
echo "Results dir: ${results_dir}"
echo "if_exists: ${if_exists}"
echo "Instancias candidatas: ${#files[@]}"
echo "Jobs: ${jobs}"

executed=0
skipped=0
declare -a targets=()

for path in "${files[@]}"; do
	file="${path##*/}"
	base="${file%.graph}"

	shopt -s nullglob
	existing_bf=("${results_dir}/summary/${base}__bruteforce__s"*.json)
	shopt -u nullglob

	if [[ ${#existing_bf[@]} -gt 0 ]]; then
		if [[ "${if_exists}" == "skip" ]]; then
			echo "-> Pulando ${base}: ja existe bruteforce em summary/."
			skipped=$((skipped + 1))
			continue
		fi

		if [[ "${if_exists}" == "error" ]]; then
			echo "Ja existe bruteforce em summary/ para ${base}." >&2
			exit 1
		fi
	fi

	targets+=("${path}")
done

if [[ ${#targets[@]} -eq 0 ]]; then
	echo "Concluido: nada para executar. Executadas: 0 | Puladas: ${skipped}"
	exit 0
fi

for path in "${targets[@]}"; do
	file="${path##*/}"
	base="${file%.graph}"

	path_hash="$(printf '%s' "${path}" | sha1sum | cut -c1-8)"
	run_id="${base}__bruteforce__s${seed}__h${path_hash}"
	log_file="${results_dir}/logs/${run_id}.log"

echo "-> Rodando bruteforce: ${base}" >&2
echo "\"${tcc_bin}\" optimize bruteforce --instance \"${path}\" --seed ${seed} --folder \"${data_folder}\" --results-dir \"${results_dir}\" --if-exists \"${if_exists}\" --progress=${progress} 2> \"${log_file}\""
done | parallel --halt soon,fail=1 -j "${jobs}"

executed=${#targets[@]}

echo "Concluido: bruteforce para instancias >= ${min_size}. Executadas: ${executed} | Puladas: ${skipped}"
