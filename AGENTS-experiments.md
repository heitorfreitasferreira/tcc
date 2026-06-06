# AGENTS — Experiments

## Go CLI Commands

- Build: `make -C src build` → `src/tcc`; batch scripts exigem este executável.
- Testes focados: de `src/`, `go test ./path/to/pkg -run TestName`; testes completos: `go test ./...`.
- `make -C src test` roda `go mod tidy`, `go mod vendor`, depois escreve `coverage.out` e `report.json`.
- Lint: `make -C src lint` (requer `golangci-lint`).
- Exemplos: `./src/tcc create -s 42 -f ./src/data`, `./src/tcc optimize ga --instance ./src/data/10a.graph --results-dir ./src/data/results`, `./src/tcc serve --addr :8080`.

## Fluxo de geração
1. `make -C src build` → `src/tcc`
2. `src/tcc create -s <seed> -f ./src/data` — gera `.points` e `.graph`
3. `src/tcc optimize <method> --instance <file> --seed <N> --results-dir <dir>` — executa uma run
4. Scripts batch:
   - `src/run_all.sh` — roda um método/seed para instâncias definidas por `--frequency=n:q`
   - `src/run_experiments.sh` — cenários 1 e 2 para seed única
   - `src/run_experiments_multi_seed.sh` — seeds 0..50 para `ga,pso,aco`
   - `src/run_experiments_bruteforce_missing.sh` — bruteforce para instâncias >= n

## Instâncias
- `src/data/<n><label>.graph` e `.points` ex: `10a`, `100b`, `30c`
- `n` = número de nós (cidades/POIs), `label` = variante (a, b, c por tamanho)
- Scripts derivam frequência automaticamente dos arquivos `.graph` presentes

## Métodos e parâmetros
| Método | Parâmetros específicos |
|--------|------------------------|
| `bruteforce` | nenhum |
| `ga` | `--elitism`, `--mutation-rate`, `--tournament-size` |
| `pso` | `--c1`, `--c2`, `--w` |
| `aco` | `--alpha`, `--beta`, `--gama`, `--rho`, `--q` |
| `lowerbound` | nenhum (determinístico, 1 avaliação) |

Métodos heurísticos (ga, pso, aco) aceitam `--population` (padrão 100) e `--iterations` (padrão 100). `lowerbound` ignora esses flags por ser determinístico. `bruteforce` também ignora.

## Artefatos de saída (por run)
Gerados em `<results-dir>/` (padrão `src/data/results`):

| Diretório | Arquivo | Esquema |
|-----------|---------|---------|
| `summary/` | `<run_id>.json` | `method`, `instance`, `seed`, `params`, `status`, `result.best_makespan`, `result.best_sequence`, `result.evaluations`, `timing_file`, `evolution_file` |
| `evolution/` | `<run_id>.jsonl` | 1 linha por melhoria: `iter`, `eval_count`, `best_makespan`, `delta`, `best_sequence` |
| `timing/` | `<run_id>.json` | `durations_ms.load_instance`, `.optimize`, `.serialize`, `.total` |
| `logs/` | `<run_id>.log` | stderr da execução (não embutido no web server) |

## `run_id`
Formato: `<instance>__<method>__s<seed>__h<hash8>` (versão atual) ou `<instance>__<method>__s<seed>__p<pop>__i<iter>__h<hash>` (scripts antigos). Sempre confiar no campo `run_id` do JSON, não no nome do arquivo.

## Cobertura atual
- Heurísticas (ga, pso, aco): 51 seeds para 30 instâncias (10a–100c) = 1.530 runs por método
- `lowerbound`: 1 seed (determinístico) para toda instância `.graph` disponível
- Bruteforce: 4 seeds (n ≤ 13) ou 1 seed (n = 14), ausente para n ≥ 15
- Total bruto ~5016 arquivos, ~4641 runs únicas após dedup por `(instance, method, seed)`

## Análise (notebook)
`scripts/visualizacoes.ipynb`:
- Deduplica runs mantendo a última cronologicamente
- Agrega por `(instance, method)` com mediana do makespan e tempo
- Gera: heatmaps de cobertura/qualidade, curvas de convergência, trade-off qualidade×tempo, gap% vs bruteforce, mapas de rota

## Web server
`./src/tcc serve --addr :8080` — visualizador interativo embutindo os mesmos dados via `src/data/assets.go` (embed: `*.graph`, `*.points`, `results/summary`, `results/evolution`, `results/timing`; **não** embute `logs/`).

## Atenção para agentes
- Não fortalecer conclusões na monografia sem verificar artefatos reais
- Não assumir que `run_id` no nome do arquivo é estável — ler o JSON
- Dedup por `(instance, method, seed)` necessário (há duplicatas reais)
- Bruteforce tem cobertura assimétrica (faltam baselines para n ≥ 15)
- Notebook, web e texto da monografia devem usar os mesmos resultados agregados
