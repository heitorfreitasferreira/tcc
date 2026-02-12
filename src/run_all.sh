#!/bin/bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
tcc_bin="${script_dir}/tcc"

while [[ $# -gt 0 ]]; do
  case "$1" in
  --frequency=*)
    frequency_input="${1#*=}"
    ;;
  --method=*)
    optimization_method="${1#*=}"
    ;;
  --seed=*)
    seed="${1#*=}"
    ;;
  --population=*)
    population="${1#*=}"
    ;;
  --iterations=*)
    iterations="${1#*=}"
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
  *)
    echo "Unknown option: $1"
    exit 1
    ;;
  esac
  shift
done

frequency_input="${frequency_input:-10:3,11:3,12:3,13:3,14:3}"
optimization_method="${optimization_method:-bruteforce}"
seed="${seed:-0}"
population="${population:-100}"
iterations="${iterations:-100}"
results_dir="${results_dir:-${script_dir}/data/results}"
data_folder="${data_folder:-${script_dir}/data}"
if_exists="${if_exists:-skip}"
progress="${progress:-false}"
jobs="${jobs:-$(nproc)}"

if [[ ! -x "${tcc_bin}" ]]; then
  echo "Binary not found at ${tcc_bin}. Build it first with 'make -C src build' or 'go build -o ./tcc' inside src/." >&2
  exit 1
fi

mkdir -p "${results_dir}/logs"

declare -A frequency_map=()
IFS=',' read -r -a frequency_pairs <<<"${frequency_input}"

for raw_pair in "${frequency_pairs[@]}"; do
  pair="${raw_pair//[[:space:]]/}"
  if [[ -z "${pair}" ]]; then
    continue
  fi

  if [[ ! "${pair}" =~ ^([0-9]+):([0-9]+)$ ]]; then
    echo "Invalid --frequency entry: ${raw_pair}" >&2
    exit 1
  fi

  size="${BASH_REMATCH[1]}"
  count="${BASH_REMATCH[2]}"
  frequency_map["${size}"]="${count}"
done

if [[ ${#frequency_map[@]} -eq 0 ]]; then
  echo "No valid frequency entries were provided." >&2
  exit 1
fi

mapfile -t sorted_sizes < <(printf '%s\n' "${!frequency_map[@]}" | sort -n)

files=()
for size in "${sorted_sizes[@]}"; do
  count="${frequency_map[${size}]}"
  for ((idx = 0; idx < 10#${count}; idx++)); do
    printf -v suffix "\\$(printf '%03o' "$((97 + idx))")"
    path="${data_folder}/${size}${suffix}.graph"
    if [[ ! -f "${path}" ]]; then
      echo "Expected graph file not found for --frequency entry: ${path}" >&2
      exit 1
    fi
    files+=("${path}")
  done
done

if [[ ${#files[@]} -eq 0 ]]; then
  echo "No graph files requested by frequency ${frequency_input}" >&2
  exit 0
fi

for path in "${files[@]}"; do
  file="${path##*/}"
  base="${file%.graph}"
  path_hash="$(printf '%s' "${path}" | sha1sum | cut -c1-8)"
  run_id="${base}__${optimization_method}__s${seed}__p${population}__i${iterations}__h${path_hash}"
  log_file="${results_dir}/logs/${run_id}.log"

  echo "\"${tcc_bin}\" optimize ${optimization_method} --instance \"${path}\" --seed ${seed} --population ${population} --iterations ${iterations} --results-dir \"${results_dir}\" --if-exists \"${if_exists}\" --progress=${progress} 2> \"${log_file}\""
done | parallel -j "${jobs}"
