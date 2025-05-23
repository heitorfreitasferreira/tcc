#!/bin/bash

while [[ $# -gt 0 ]]; do
  case "$1" in
  --frequency=*)
    frequency_input="${1#*=}"
    ;;
  --method=*)
    optimization_method="${1#*=}"
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

./tcc create --frequency "${frequency_input}"

shopt -s nullglob
files=(data/*.graph)

for file in "${files[@]##*/}"; do
  base="${file%.graph}"
  echo "{ time ./tcc optimize ${optimization_method} --instance data/${file} > \
    data/${base}.${optimization_method}; } 2> data/${base}.${optimization_method}.time"
done | parallel -j "$(nproc)"
