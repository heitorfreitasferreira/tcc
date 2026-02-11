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

"${tcc_bin}" create --frequency "${frequency_input}" --folder "${data_folder}" --seed "${seed}"

shopt -s nullglob
files=("${data_folder}"/*.graph)

if [[ ${#files[@]} -eq 0 ]]; then
  echo "No graph files found in ${data_folder}" >&2
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
