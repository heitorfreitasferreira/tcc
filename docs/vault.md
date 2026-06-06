# Componente: Vault Semantico

## Responsabilidade

O `vault/` e a memoria de trabalho do framework. Ele organiza literatura, claims, areas, siglas, auditorias, planejamento e evidencias para orientar a escrita da monografia.

O vault nao e o resultado final; o resultado final e `monografia/`. O vault tambem nao substitui fonte primaria; ele aponta para codigo, dados, PDFs, BibTeX e auditorias.

## Estrutura Principal

```text
vault/
├── papers/        # notas de papers
├── claims/        # claims individuais
├── areas/         # areas e metodos
├── bases/         # views Obsidian
├── canvas/        # grafo visual
├── roadmap/       # tarefas e eventos
├── siglas/        # siglas e inclusao na monografia
└── writing/       # planejamento, auditorias e reviews
```

## Tags

O vault usa somente tags hierarquicas.

| Namespace | Exemplo | Uso |
|---|---|---|
| `tipo/` | `tipo/paper` | Tipo da nota |
| `area/` | `area/tsp` | Dominio |
| `metodo/` | `metodo/aco` | Metodo ou tecnica |
| `status/` | `status/lido` | Estado operacional |
| `evidencia/` | `evidencia/referencia` | Natureza da evidencia |
| `capitulo/` | `capitulo/fundamentacao` | Capitulo afetado |
| `relevancia/` | `relevancia/5` | Rating refletido como tag |

## Claims

Claims ficam em `vault/claims/*.md` e devem apontar para fonte primaria.

Prefixos comuns:

- `C`: conceitual
- `M`: metodologico
- `A`: algoritmo
- `E`: experimental
- `I`: interpretativo
- `B`: bloqueado

## Bases

- `papers.base`: consulta de papers.
- `claims.base`: matriz consultavel de claims.
- `roadmap-tarefas.base`: fila e status.
- `roadmap-log.base`: eventos.
- `siglas.base`: controle de siglas.

## Regras

- Nao usar flat tags.
- Nao remover `bibtex-key` de papers.
- Nao usar canvas como evidencia primaria.
- Nao promover claim para forte sem fonte primaria.
- Rodar `python3 scripts/migrate-tags.py` apos alteracoes grandes.
- Usar o vault para lembrar, organizar e rastrear; usar `monografia/` para o texto final.
