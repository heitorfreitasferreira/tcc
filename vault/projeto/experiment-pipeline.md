---
tags: [projeto, implementacao, experimentos, pipeline, cli]
status: atualizado-pos-p7
updated: 2026-06-02
---

# Pipeline de Experimentos

## CLI — Comandos

### `tcc create`

Gera instâncias (pontos + grafos). `src/cmd/create.go`:

```bash
./tcc create -s 42 -f ./src/data
```

- `--seed` (int64): semente determinística
- `--frequency` (string): formato `n:q` (ex: `10:3,11:3,12:3`)
- Gera `.points` e `.graph` em `--folder`

### `tcc optimize <method>`

Otimiza uma instância com um método. `src/cmd/optimize.go`:

```bash
./tcc optimize ga \
  --instance ./src/data/10a.graph \
  --seed 42 \
  --population 100 \
  --iterations 100 \
  --results-dir ./src/data/results
```

Métodos: `ga`, `pso`, `aco`, `bruteforce`, `lowerbound`.

`lowerbound` executa a relaxação AP/Hungarian e retorna um limitante inferior, não uma rota factível. Ele entra como referência numérica, não como competidor estocástico.

### `tcc serve`

Servidor web para visualização de resultados.

## Scripts de Batch

### `run_all.sh`

Executa um método em múltiplas instâncias com GNU parallel:

```bash
./src/run_all.sh --method=ga --frequency=10:3,11:3 --seed=0 --jobs=4
```

### `run_experiments_multi_seed.sh`

Executa múltiplos métodos em múltiplas sementes:

```bash
./src/run_experiments_multi_seed.sh \
  --seed-start=0 --seed-end=50 \
  --methods=ga,pso,aco \
  --frequency=10:3,11:3
```

Gera 51 sementes × 3 métodos × N instâncias execuções.

### `run_experiments_bruteforce_missing.sh`

Executa brute-force para instâncias faltantes.

## Estrutura de Resultados

```
results/
├── summary/<run_id>.json       ← melhor resultado
├── evolution/<run_id>.jsonl    ← histórico de melhorias
├── timing/<run_id>.json        ← tempos de execução
└── logs/<run_id>.log           ← stderr do processo
```

### Run ID

Formato atual: `{instancia}__{metodo}__s{seed}__h{hash}`

Exemplo: `10a__ga__s0__hd186dc51`

O hash é derivado dos parâmetros canônicos da execução. Parâmetros como população e iterações ficam registrados no `summary.params`, mas não aparecem explicitamente no nome do arquivo.

## Fluxo de Execução

1. Build: `make -C src build` → `src/tcc`
2. Geração: `tcc create` → `.points` + `.graph`
3. Otimização: `run_experiments_multi_seed.sh` → resultados estruturados
4. Análise: `scripts/consolidate_results.py`, `scripts/analise-estatistica.py`, `scripts/gerar-graficos-estatisticos.py` e/ou `scripts/visualizacoes.ipynb`

Cobertura auditada atual: 4638 summaries, 4638 timings e 4638 evolutions. Distribuição: 1530 GA, 1530 PSO, 1530 ACO, 30 lowerbound e 18 brute-force. Ver [[auditoria-codigo-dados-vault]].

## Estruturas de Dados

### `shared.OptimizationResult` (`src/shared/optimization.go`)

```go
type OptimizationResult struct {
    BestSequence        []int
    BestMakespan        float64
    IterationsCompleted int
    Evaluations         int
    Improvements        []Improvement
}
```

### `shared.Improvement`

```go
type Improvement struct {
    Iteration    int
    Evaluation   int
    BestMakespan float64
    Delta        float64
    BestSequence []int
}
```

## Política de Sobrescrita

`--if-exists`: `skip` (default) | `overwrite` | `error`

## Conexões

- [[architecture]] — estrutura de pacotes
- [[ga]], [[pso]], [[aco]], [[bruteforce]], [[lower-bounds]] — métodos e referências executados
- [[problem-formulation]] — instâncias usadas
- [[comparative-studies]] — contexto experimental na literatura

## Código-fonte

- `src/cmd/optimize.go` — flags comuns
- `src/cmd/reporting.go` — persistência de resultados
- `src/shared/reporting/` — schemas e escritores
- `src/run_all.sh` — executor batch com parallel
- `src/run_experiments_multi_seed.sh` — multi-seed orchestrator
