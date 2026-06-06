---
description: Executa experimentos batch (src/run_experiments_multi_seed.sh) e registra no log do roadmap
---

# /experimento

Encapsula `src/run_experiments_multi_seed.sh` e registra o resultado no log do roadmap.

## Uso

```
/experimento --method ga --frequency 10:3 --seeds 0-2
/experimento --method aco --seeds 0-50
/experimento --method all --frequency 50:3,100:3
```

## Flags

| Flag | Default | Descrição |
|---|---|---|
| `--method` | `ga` | `ga`, `pso`, `aco`, `lowerbound`, `all` |
| `--frequency` | (auto) | `10:3,11:3,...` ou auto-detecta do diretório |
| `--seeds` | `0-50` | Range `inicio-fim` (inclusive) |
| `--jobs` | `nproc` | Paralelismo |
| `--if-exists` | `skip` | `skip`, `overwrite`, `error` |

## Fluxo

### 1. Verificar build
```bash
make -C src build
```

### 2. Executar
```bash
bash src/run_experiments_multi_seed.sh \
  --method <m> --frequency <f> --seed-start <a> --seed-end <b> \
  --if-exists skip
```

### 3. Validar cobertura
```bash
python3 scripts/gerar-metricas-monografia.py --check
```
Se FAIL → alerta sobre dados incompletos.

### 4. Logar
```bash
bash scripts/roadmap.sh log experiment \
  method=aco \
  instance=todas \
  seeds=0-50 \
  runs=1530 \
  status=ok
```

### 5. Reportar
Mostra resumo:
- Método, instâncias, seeds
- Runs executadas vs esperadas
- Status
- Evento logado em `vault/roadmap/eventos/`

## Dependências
- `src/tcc` (build)
- `src/run_experiments_multi_seed.sh`
- `parallel` (GNU parallel)
