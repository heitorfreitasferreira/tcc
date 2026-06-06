# Componente: Scripts Auxiliares

## `scripts/roadmap.sh`

Camada shell principal da fila.

Subcomandos:

- `tarefa criar`: cria `P<N>.md`.
- `tarefa listar`: lista pendentes.
- `tarefa concluir`: marca como concluida.
- `proximo`: imprime a proxima pendente.
- `log`: cria evento em `vault/roadmap/eventos/`.

## Scripts de Literatura

| Script | Papel |
|---|---|
| `scripts/import-bib-to-vault.sh` | Gera notas de papers a partir do BibTeX |
| `scripts/download-pdfs.sh` | Busca e normaliza PDFs |
| `scripts/extract-pdf-doi.sh` | Extrai DOI de PDF |
| `scripts/migrate-tags.py` | Sincroniza tags e propriedades do vault |

## Scripts de Monografia

| Script | Papel |
|---|---|
| `scripts/check-monografia.sh` | Verificacoes pre-compilacao |
| `scripts/gerar-metricas-monografia.py` | Gera ou checa metricas usadas no texto |

## Scripts de Review

| Script | Papel |
|---|---|
| `scripts/registrar-review-scaffold.sh` | Cria review note para `/claudiney` |

## Contratos

- Scripts devem ser idempotentes quando possivel.
- Saidas devem ser estruturadas e rastreaveis.
- Scripts que alteram estado devem escrever no vault ou em `src/data/results/`, nunca apenas no terminal.
- Comandos agenticos podem chamar scripts, mas devem interpretar erros e criar tarefas quando houver pendencia.
