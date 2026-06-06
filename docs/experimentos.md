# Componente: Experimentos

## Responsabilidade

Executar e registrar experimentos comparando algoritmos de otimizacao para instancias TSP/rTSP.

## Codigo

- Modulo Go: `src/`
- CLI: `src/tcc`
- Algoritmos: `bruteforce`, `ga`, `pso`, `aco`, `lowerbound`
- Dados: `src/data/`

## Comando Agentico

`/experimento` encapsula o batch experimental.

```bash
/experimento --method ga --frequency 10:3 --seeds 0-50
/experimento --method all --frequency 50:3,100:3
```

## Fluxo

1. `make -C src build`
2. `bash src/run_experiments_multi_seed.sh ...`
3. `python3 scripts/gerar-metricas-monografia.py --check`
4. `bash scripts/roadmap.sh log experiment ...`
5. Relatar cobertura e arquivos gerados.

## Saidas Estruturadas

| Pasta | Conteudo |
|---|---|
| `src/data/results/summary/` | Melhor solucao final e metadados |
| `src/data/results/evolution/` | Eventos de melhora por iteracao |
| `src/data/results/timing/` | Tempos por etapa |
| `src/data/results/logs/` | Logs de execucao batch |

## Relacao com Claims

Claims experimentais em `vault/claims/E*.md` devem apontar para resultados ou scripts que permitam reproducao.

## Regras

- Nao preencher tabela experimental manualmente sem script.
- Registrar seeds e parametros.
- Validar cobertura antes de escrever resultados na monografia.
- Separar execucao, agregacao e interpretacao.
