#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
run_all_script="${script_dir}/run_all.sh"

seed_start="0"
seed_end="50"
methods_csv="ga,pso,aco"
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
  --seed-start=<int>     Seed inicial (padrao: 0)
  --seed-end=<int>       Seed final inclusive (padrao: 50)
  --methods=<csv>        Metodos em CSV (padrao: ga,pso,aco)
  --frequency=<spec>     Frequencia no formato n:q (ex: 10:3,11:3)
  --results-dir=<path>   Diretorio de resultados
  --folder=<path>        Pasta com .graph (padrao: src/data)
  --if-exists=<mode>     Modo para execucoes existentes (padrao: skip)
  --progress=<bool>      Exibir progresso no binario (padrao: false)
  --jobs=<int>           Paralelismo interno por chamada do run_all
  -h, --help             Mostra esta ajuda
EOF
}

build_frequency_from_graphs() {
  local folder="$1"
  local min_size="${2:-}"
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
    if [[ -n "${min_size}" ]] && ((10#${size} < 10#${min_size})); then
      continue
    fi

    counts["${size}"]=$((${counts["${size}"]:-0} + 1))
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
  --seed-start=*)
    seed_start="${1#*=}"
    ;;
  --seed-end=*)
    seed_end="${1#*=}"
    ;;
  --methods=*)
    methods_csv="${1#*=}"
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
  -h | --help)
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

if [[ ! "${seed_start}" =~ ^-?[0-9]+$ ]]; then
  echo "Invalid --seed-start: ${seed_start}. Expected an integer." >&2
  exit 1
fi

if [[ ! "${seed_end}" =~ ^-?[0-9]+$ ]]; then
  echo "Invalid --seed-end: ${seed_end}. Expected an integer." >&2
  exit 1
fi

if ((seed_start > seed_end)); then
  echo "Invalid interval: seed_start (${seed_start}) must be <= seed_end (${seed_end})." >&2
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
  if ! frequency_input="$(build_frequency_from_graphs "${data_folder}")"; then
    echo "Could not derive frequency from ${data_folder}." >&2
    exit 1
  fi
fi

IFS=',' read -r -a requested_methods <<<"${methods_csv}"
methods=()
declare -A seen_methods=()

for raw_method in "${requested_methods[@]}"; do
  method="${raw_method//[[:space:]]/}"
  if [[ -z "${method}" ]]; then
    continue
  fi

  case "${method}" in
  ga | pso | aco | lowerbound)
    if [[ -z "${seen_methods["${method}"]:-}" ]]; then
      methods+=("${method}")
      seen_methods["${method}"]=1
    fi
    ;;
  bruteforce)
    echo "Invalid method in --methods: bruteforce is not allowed in this script." >&2
    exit 1
    ;;
  *)
    echo "Invalid method in --methods: ${method}. Allowed: ga,pso,aco,lowerbound" >&2
    exit 1
    ;;
  esac
done

if [[ ${#methods[@]} -eq 0 ]]; then
  echo "No valid methods provided in --methods." >&2
  exit 1
fi

seed_count=$((seed_end - seed_start + 1))
planned_calls=$((seed_count * ${#methods[@]}))

echo "=== Experimentos multi-seed (sem bruteforce) ==="
echo "Seeds: ${seed_start}..${seed_end} (${seed_count} seeds)"
echo "Metodos: ${methods[*]}"
echo "Frequencia: ${frequency_input}"
echo "Data folder: ${data_folder}"
echo "Chamadas previstas ao run_all: ${planned_calls}"

for ((seed = seed_start; seed <= seed_end; seed++)); do
  echo "=== Seed ${seed} ==="
  for method in "${methods[@]}"; do
    echo "-> Rodando metodo: ${method}"
    cmd=(
      "${run_all_script}"
      "--method=${method}"
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
  done
done

echo "Concluido: loop de seeds para metodos heurisiticos finalizado."
