# Data Schemas

Este arquivo documenta os formatos de dados presentes em `src/data`.

## Estrutura de pastas

- `src/data/*.points`: coordenadas 2D das instancias.
- `src/data/*.graph`: matriz de custos usada pelos algoritmos.
- `src/data/results/summary/*.json`: resumo final por execucao.
- `src/data/results/timing/*.json`: tempos por execucao.
- `src/data/results/evolution/*.jsonl`: historico de melhorias (eventos).
- `src/data/results/logs/*.log`: saida textual do comando.

---

## 1) Schema: `.points` (nao versionado)

Formato JSON:

```json
[
  [x0, y0],
  [x1, y1],
  ...
]
```

Regras:

- Tipo: `array` de pontos.
- Cada ponto: `array[2]` de `float64`.
- Ordem dos pontos define o id do no (`index` no array).
- Pela geracao atual, coordenadas ficam em `[-1, 1]`.

---

## 2) Schema: `.graph` (nao versionado)

Formato JSON:

```json
[
  [
    [c000, c001, ...],
    [c010, c011, ...],
    ...
  ],
  ...
]
```

Regras:

- Tipo: `array[n][n][n]` de `float64`.
- Semantica esperada pelo `Makespan`: `g[anterior][atual][proximo]`.
- A origem fixa e o no `0`.
- A rota valida e uma permutacao dos nos `1..n-1`.

---

## 3) Schema: `summary` (`tcc.summary.v1`)

Arquivo: `src/data/results/summary/<run_id>.json`

Campos de topo:

- `schema` (`string`): `"tcc.summary.v1"`.
- `run_id` (`string`): id unico da execucao.
- `method` (`string`): `aco | ga | pso | bruteforce`.
- `instance` (`string`): caminho absoluto do `.graph`.
- `seed` (`int64`).
- `params` (`object`): hiperparametros do metodo.
- `status` (`string`): normalmente `"ok"`.
- `result` (`object`): resultado final.
- `timing_file` (`string`): caminho relativo para `timing/<run_id>.json`.
- `evolution_file` (`string`): caminho relativo para `evolution/<run_id>.jsonl`.
- `started_at` (`string`): timestamp RFC3339Nano UTC.
- `finished_at` (`string`): timestamp RFC3339Nano UTC.

`result`:

- `best_makespan` (`float64`)
- `best_sequence` (`int[]`) rota final (esperado: nos `1..n-1`)
- `iterations_completed` (`int`)
- `evaluations` (`int`)

Observacao sobre `params`:

- `aco`: `alpha`, `beta`, `gama`, `iterations`, `population`, `q`, `rho`
- `ga`: `elitism`, `iterations`, `mutation_rate`, `population`, `tournament_size`
- `pso`: `c1`, `c2`, `iterations`, `population`, `w`
- `bruteforce`: geralmente vazio (`{}`)

---

## 4) Schema: `timing` (`tcc.timing.v1`)

Arquivo: `src/data/results/timing/<run_id>.json`

Campos:

- `schema` (`string`): `"tcc.timing.v1"`
- `run_id` (`string`)
- `method` (`string`)
- `instance` (`string`)
- `durations_ms` (`object`)
- `measured` (`object`)

`durations_ms`:

- `load_instance` (`int64`)
- `optimize` (`int64`)
- `serialize` (`int64`)
- `total` (`int64`)

`measured`:

- `started_at` (`string`, RFC3339Nano UTC)
- `finished_at` (`string`, RFC3339Nano UTC)

---

## 5) Schema: `evolution` (`tcc.evolution.v1`)

Arquivo: `src/data/results/evolution/<run_id>.jsonl`

Formato: JSON Lines (1 objeto JSON por linha). Cada linha representa uma melhoria de melhor solucao global, nao cada iteracao completa.

Campos por linha:

- `schema` (`string`): `"tcc.evolution.v1"`
- `run_id` (`string`)
- `method` (`string`)
- `iter` (`int`)
- `eval_count` (`int`)
- `best_makespan` (`float64`)
- `delta` (`float64`) diferenca para o melhor anterior (tipicamente `<= 0`)
- `best_sequence` (`int[]`) rota da melhoria

Observacoes praticas:

- O ultimo `eval_count` pode ser menor que `summary.result.evaluations` (fim das melhorias antes do fim da busca).
- O numero de linhas e o numero de eventos de melhoria.

---

## 6) Logs (`.log`, nao versionado)

Arquivo: `src/data/results/logs/<...>.log`

Formato textual simples (sem schema JSON). Exemplo:

```text
run_id=<id> method=<metodo> best_makespan=<valor> summary=<caminho>
```

Uso principal: auditoria humana e depuracao.

---

## Convencoes de identificacao

`run_id` segue o formato:

```text
<instancia>__<metodo>__s<seed>__h<hash8>
```

Exemplo:

```text
13b__ga__s0__h7ed2c620
```

As tres familias de artefato (`summary`, `timing`, `evolution`) devem compartilhar o mesmo `run_id`.

