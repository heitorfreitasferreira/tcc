#!/bin/bash

make build

rm -rf data
./tcc create

shopt -s nullglob
files=(data/*.graph)

for file in "${files[@]##*/}"; do
    base="${file%.graph}"
    echo "./tcc optimize bruteforce --instance data/${file} > data/${base}.optimal"
done | parallel -j $(nproc)
