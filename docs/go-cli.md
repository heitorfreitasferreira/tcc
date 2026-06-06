# Componente: Codigo Experimental Go

## Objetivo

Gerar instancias, grafos e resultados experimentais para comparar metaheuristicas bioinspiradas em TSP/rTSP.

## Localizacao

```text
src/
├── cmd/                         # comandos Cobra
├── optimization/
├── graph/
├── web/                         # visualizador auxiliar
├── data/
└── run_experiments_multi_seed.sh
```

O nucleo experimental esta em `src/cmd/`, `src/optimization/`, `src/graph/` e `src/data/`. O diretorio `src/web/` contem um visualizador feito como apoio extra; ele nao define o escopo cientifico central.

## Build

```bash
make -C src build
```

## Comandos Principais

### `create`

Gera instancias e grafos.

```bash
tcc create -s 42 -f ./data
```

### `create map`

Gera somente pontos.

```bash
tcc create map -s 42 -f ./data
```

### `graph`

Gera grafos a partir de pontos existentes.

```bash
tcc graph -f ./data
```

### `optimize <method>`

Otimiza uma instancia `.graph`.

```bash
tcc optimize ga --instance ./data/10a.graph --results-dir ./data/results
```

Metodos:

- `bruteforce`
- `ga`
- `aco`
- `pso`
- `lowerbound`

### `serve`

Servidor web local para visualizar resultados.

```bash
tcc serve --addr :8080
```

## Saidas

- `summary/<run_id>.json`
- `evolution/<run_id>.jsonl`
- `timing/<run_id>.json`
- `logs/<run_id>.log`

## Relacao com o Framework

O CLI Go e a fonte primaria para claims experimentais. A monografia nao deve citar resultados que nao possam ser reconstruidos a partir dos artefatos em `src/data/results/`.

Quando o framework identifica bug, lacuna de metodo ou divergencia entre texto e codigo, ele deve criar tarefa, propor correcao com supervisao humana e reexecutar a validacao necessaria antes de atualizar a monografia.
