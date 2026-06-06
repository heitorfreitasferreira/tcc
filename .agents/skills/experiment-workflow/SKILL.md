---
name: experiment-workflow
description: Project-specific Go CLI workflow, experiment data structures, batch scripts, and run artifacts for the TCC optimization benchmarks.
tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - bash
---

<role>
You are an agent that runs and analyzes optimization experiments for the TCC monografia. You know the exact CLI commands, data schemas, batch scripts, and coverage status of all runs. You use this knowledge to execute experiments, interpret results, and keep the monograph aligned with actual artifact data.
</role>

## Go CLI Commands

- Build: `make -C src build` → `src/tcc`; batch scripts require this binary.
- Tests: `go test ./path/to/pkg -run TestName`; full suite: `go test ./...` from `src/`.
- Lint: `make -C src lint` (requires `golangci-lint`).
- Examples: `./src/tcc create -s 42 -f ./src/data`, `./src/tcc optimize ga --instance ./src/data/10a.graph --results-dir ./src/data/results`, `./src/tcc serve --addr :8080`.

## Generation flow

1. `make -C src build`
2. `src/tcc create -s <seed> -f ./src/data` — generates `.points` + `.graph`
3. `src/tcc optimize <method> --instance <file> --seed <N> --results-dir <dir>` — single run
4. Batch scripts:
   - `src/run_all.sh` — one method/seed for instances by `--frequency=n:q`
   - `src/run_experiments.sh` — scenarios 1 and 2, single seed
   - `src/run_experiments_multi_seed.sh` — seeds 0..50 for `ga,pso,aco`
   - `src/run_experiments_bruteforce_missing.sh` — bruteforce for n >= threshold

## Instances

- `src/data/<n><label>.graph` and `.points` e.g. `10a`, `100b`, `30c`
- `n` = nodes (cities/POIs), `label` = variant (a, b, c by size)
- Scripts auto-derive frequency from present `.graph` files

## Methods and parameters

| Method | Specific params |
|--------|-----------------|
| `bruteforce` | none |
| `ga` | `--elitism`, `--mutation-rate`, `--tournament-size` |
| `pso` | `--c1`, `--c2`, `--w` |
| `aco` | `--alpha`, `--beta`, `--gama`, `--rho`, `--q` |
| `lowerbound` | none (deterministic, 1 eval) |

Heuristics accept `--population` (default 100) and `--iterations` (default 100).

## Output artifacts (per run)

Written under `<results-dir>/` (default `src/data/results`):

| Dir | File | Schema |
|-----|------|--------|
| `summary/` | `<run_id>.json` | method, instance, seed, params, status, result.best_makespan, result.best_sequence, evaluations, timing_file, evolution_file |
| `evolution/` | `<run_id>.jsonl` | 1 line per improvement: iter, eval_count, best_makespan, delta, best_sequence |
| `timing/` | `<run_id>.json` | durations_ms.load_instance, .optimize, .serialize, .total |
| `logs/` | `<run_id>.log` | stderr (not embedded in web server) |

## run_id format

`<instance>__<method>__s<seed>__h<hash8>` (current) or old format `<instance>__<method>__s<seed>__p<pop>__i<iter>__h<hash>`. Always trust `run_id` field in JSON, not filename.

## Current coverage

- Heuristics (ga, pso, aco): 51 seeds × 30 instances (10a–100c) = 1,530 runs each
- `lowerbound`: 1 seed (deterministic) for all `.graph` instances
- Bruteforce: 4 seeds (n ≤ 13) or 1 seed (n = 14), missing for n ≥ 15
- Total ~5016 files, ~4641 unique runs after dedup by `(instance, method, seed)`

## Analysis notebook

`scripts/visualizacoes.ipynb`: dedup (keep last chronologically), aggregate by `(instance, method)` with median makespan and time, generate heatmaps, convergence curves, quality×time trade-off, gap% vs bruteforce, route maps.

## Web server

`./src/tcc serve --addr :8080` — interactive viewer embedding `*.graph`, `*.points`, `results/summary`, `results/evolution`, `results/timing` (NOT `logs/`).

## Caveats

- Do not strengthen conclusions in the monograph without verifying real artifacts.
- `run_id` in filename is not stable — always read the JSON.
- Dedup by `(instance, method, seed)` required (real duplicates exist).
- Bruteforce coverage is asymmetric (baselines missing for n ≥ 15).
- Notebook, web server, and monograph text must use the same aggregated results.
