---
name: vault-tagger
description: Padroniza tags hierarquicas no vault Obsidian do TCC. Use quando o usuario pedir para tagear notas, organizar tags, filtrar grafo do vault, revisar frontmatter ou melhorar buscas no vault/.
---

# Vault Tagger

Use esta skill para adicionar, revisar ou normalizar tags nas notas Markdown do `vault/` do TCC. O objetivo e melhorar buscas, Graph View, Bases e rastreabilidade entre literatura, projeto, escrita e evidencias.

Para propriedades YAML, relacoes tipadas, templates, aliases e claims, use tambem `vault-semantic-schema`. Para editar arquivos `.base`, use `vault-bases-maintainer`.

## Estado Atual Do Vault

- **156 notas** com frontmatter, **zero flat tags** — todas usam exclusivamente tags hierarquicas `namespace/valor`.
- Migracao completa executada por `scripts/migrate-tags.py` (idempotente, seguro reexecutar).
- Color groups do Graph View configurados em `vault/.obsidian/graph.json` por `tipo/*`.

## Principios

- Preserve o conteudo das notas; altere apenas frontmatter/tags, salvo pedido explicito.
- **Apenas tags hierarquicas** `namespace/valor`. Nunca adicione flat tags como `ga`, `pso`, `drone`, `review`.
- Cada nota deve ter pelo menos um `tipo/*` como tag primaria de classificacao.
- Tags complementam propriedades YAML (`status`, `rating`, `areas`, `methods`, `chapters`, `role`); nao as substituem.
- Nao reclassifique status academico, rating, ou evidencia sem inspecionar a nota.
- `index.md` usa `tipo/index` para exclusao de grafos; nao apague indices.
- Para migracao em lote, use `python3 scripts/migrate-tags.py` (cobre path→tipo, flat→structured, enriquecimento de propriedades).

## Workflow

1. Identifique o escopo pedido: uma nota, uma pasta (`papers/`, `areas/`, `projeto/`, `siglas/`, `writing/`) ou o vault inteiro.
2. Leia o frontmatter da nota e, se necessario, as secoes `Resumo`, `Conexoes`, `Relevancia` e links `[[...]]`.
3. Aplique tags minimas por tipo de nota (`tipo/*` obrigatorio).
4. Aplique tags tematicas por area/metodo/papel com base no conteudo.
5. Aplique tags operacionais: status, relevancia, capitulo, evidencia.
6. Preserve campos estruturados (`year`, `status`, `rating`, `doi`, `bibtex-key`, `areas`, `methods`, `role`, `reading_status`, `validation_status`, `pdf_status`, `chapters`, `claim_support`).
7. Ao editar em lote, prefira `scripts/migrate-tags.py`; faca alteracoes pequenas e verificaveis.

## Taxonomia — Namespaces Ativos

Os unicos namespaces em uso no vault (pos-migracao):

| Namespace | Uso | Exemplo |
|-----------|-----|---------|
| `tipo/` | Classificacao primaria da nota | `tipo/paper` |
| `area/` | Dominio de pesquisa | `area/tsp` |
| `metodo/` | Metodo/tecnica/algoritmo | `metodo/aco` |
| `papel/` | Funcao do paper na monografia | `papel/revisao` |
| `status/` | Estado de leitura/escrita | `status/lido` |
| `evidencia/` | Tipo de evidencia/material | `evidencia/referencia` |
| `capitulo/` | Capitulo alimentado pela nota | `capitulo/fundamentacao` |
| `relevancia/` | Espelha `rating:` (1–5) | `relevancia/5` |
| `incluir/` | Inclusao na monografia (siglas/projeto) | `incluir/sim` |
| `topico/` | Topico transversal (implementacao, monografia, ferramenta) | `topico/implementacao` |
| `forca/` | Forca da evidencia (claims) | `forca/requer-validacao` |

Namespaces obsoletos removidos: `pdf/`, `escrita/`, `artefato/` (substituidos por propriedades YAML: `pdf_status`, `writing_status`, atributos em `primary_evidence`).

### Tipo De Nota (`tipo/*`)

Uma ou mais por nota. Valores canonicos:

```yaml
tags:
  - tipo/paper        # papers/, templates/paper-note.md
  - tipo/area         # areas/
  - tipo/projeto      # projeto/
  - tipo/sigla        # siglas/
  - tipo/writing      # writing/ (capitulos, planejamento, decisoes)
  - tipo/auditoria    # writing/auditorias/
  - tipo/revisao      # writing/review-solicitacoes/
  - tipo/claim        # templates/claim-note.md, claim-evidence-matrix
  - tipo/index        # */index.md (catalogos)
  - tipo/template     # templates/
```

Mapeamento canonico por caminho:

| Caminho | `tipo/*` primario |
|---------|-------------------|
| `papers/*.md` | `tipo/paper` |
| `areas/*.md` | `tipo/area` |
| `projeto/*.md` | `tipo/projeto` |
| `siglas/*.md` | `tipo/sigla` |
| `writing/auditorias/*.md` | `tipo/auditoria` |
| `writing/review-solicitacoes/*.md` | `tipo/revisao` |
| `writing/**/*.md` (demais) | `tipo/writing` |
| `*/index.md` | `tipo/index` |
| `templates/*.md` | `tipo/template` |

### Status (`status/*`)

```yaml
tags:
  - status/pendente
  - status/resumo-lido
  - status/lido-parcial
  - status/lido
  - status/descartado
  - status/removido
  - status/analise-posterior
  - status/validado
  - status/concluido
  - status/rascunho
  - status/ativo
  - status/aberto
  - status/resolvido
  - status/atualizado
  - status/pronto-revisao
  - status/pronto-com-pendencias
  - status/auditado-bloqueios
  - status/verificado
```

Derivado da propriedade YAML `status` ou `reading_status`. Use o script de migracao para mapeamento automatico.

### Areas (`area/*`)

```yaml
tags:
  - area/tsp
  - area/atsp
  - area/tdtsp
  - area/fstsp
  - area/gtsp
  - area/tsp-variants
  - area/routing
  - area/vrp
  - area/drone-routing
  - area/bio-inspired-optimization
  - area/ant-colony
  - area/genetic-algorithms
  - area/particle-swarm
  - area/comparative-studies
  - area/lower-bound
  - area/lower-bounds
```

Regras:

- TSP classico, complexidade: `area/tsp`
- ATSP: `area/atsp`
- TDTSP: `area/tdtsp`
- FSTSP, TSP-D, drones, UAV: `area/drone-routing`
- Roteamento geral/grafos: `area/routing`
- VRP: `area/vrp`
- Otimizacao bio-inspirada generica: `area/bio-inspired-optimization`
- ACO especifico: `area/ant-colony`
- GA especifico: `area/genetic-algorithms`
- PSO especifico: `area/particle-swarm`
- Comparacoes empiricas entre metodos: `area/comparative-studies`
- Lower bounds, Held-Karp, branch-and-bound: `area/lower-bound` ou `area/lower-bounds`

### Metodos (`metodo/*`)

```yaml
tags:
  - metodo/exact
  - metodo/heuristic
  - metodo/metaheuristic
  - metodo/hybrid
  - metodo/approximation
  - metodo/aco
  - metodo/ga
  - metodo/pso
  - metodo/sa
  - metodo/ts
  - metodo/abc
  - metodo/gwo
  - metodo/cso
  - metodo/ssa
  - metodo/eho
  - metodo/gp
  - metodo/vns
  - metodo/lower-bound
  - metodo/held-karp
  - metodo/lagrangean
  - metodo/branch-and-bound
  - metodo/branch-and-cut
  - metodo/milp
  - metodo/machine-learning
  - metodo/constraint-programming
  - metodo/clustering
  - metodo/crossover
  - metodo/encoding
```

Combine tags genericas e especificas: um paper sobre ACO deve ter `metodo/metaheuristic` + `metodo/aco`.

### Papel Na Monografia (`papel/*`)

```yaml
tags:
  - papel/fundacional
  - papel/revisao
  - papel/comparativo
  - papel/benchmark
  - papel/teorico
  - papel/livro
```

Regras:

- Obras classicas (Dorigo 1996, Holland 1975, Kennedy 1995): `papel/fundacional`
- Surveys, reviews, livros: `papel/revisao`
- Estudos de comparacao empirica: `papel/comparativo`
- Resultados experimentais/benchmark: `papel/benchmark`

### Capitulos (`capitulo/*`)

```yaml
tags:
  - capitulo/introducao
  - capitulo/fundamentacao
  - capitulo/proposta
  - capitulo/experimentos
  - capitulo/conclusao
```

Use em notas `writing/`, `projeto/` e `papers/` quando a nota alimenta diretamente um capitulo.

### Evidencia (`evidencia/*`)

```yaml
tags:
  - evidencia/referencia    # papers, bibliografia
  - evidencia/codigo        # projeto/, src/
  - evidencia/dados         # src/data/results/
  - evidencia/auditoria     # writing/auditorias/
  - evidencia/estatistica   # analise estatistica
  - evidencia/metodologia   # decisoes metodologicas
  - evidencia/validacao     # verificacao, claims
```

### Inclusao Na Monografia (`incluir/*`)

Usado em `siglas/` e `projeto/` (espelha propriedade YAML `incluir`):

```yaml
tags:
  - incluir/sim
  - incluir/pendente
  - incluir/nao
```

### Topico Transversal (`topico/*`)

Topicos que cruzam tipos de nota — implementacao, escrita, ferramentas:

```yaml
tags:
  - topico/implementacao
  - topico/arquitetura
  - topico/ferramenta
  - topico/experimentos
  - topico/monografia
  - topico/formatacao
  - topico/formulacao
  - topico/visao-geral
  - topico/classificacao
  - topico/glossario
  - topico/figuras
  - topico/tabelas
  - topico/citacoes
  - topico/roadmap
  - topico/revisao
  - topico/originalidade
  - topico/multi-objetivo
```

### Relevancia (`relevancia/*`)

Espelha `rating:` para uso no Graph View:

```yaml
tags:
  - relevancia/5    # rating: 5
  - relevancia/4    # rating: 4
  - relevancia/3    # rating: 3
```

Ratings 0–2 nao geram tag de relevancia.

### Forca Da Evidencia (`forca/*`)

Uso restrito a claims e auditorias:

```yaml
tags:
  - forca/requer-validacao
```

## Exemplos Atualizados

### Paper Fundacional De ACO

```yaml
---
title: "Ant system: optimization by a colony of cooperating agents"
authors: [Dorigo, Marco and Maniezzo, Vittorio and Colorni, Alberto]
year: 1996
status: lido
rating: 5
type: paper
reading_status: lido
role: fundacional
areas: [tsp]
methods: [aco, metaheuristic]
tags:
  - tipo/paper
  - status/lido
  - area/tsp
  - area/ant-colony
  - area/bio-inspired-optimization
  - metodo/aco
  - metodo/metaheuristic
  - papel/fundacional
  - evidencia/referencia
  - capitulo/fundamentacao
  - relevancia/5
---
```

### Nota De Projeto Do GA

```yaml
---
type: projeto
areas: []
methods: [ga]
tags:
  - tipo/projeto
  - metodo/ga
  - area/genetic-algorithms
  - area/bio-inspired-optimization
  - evidencia/codigo
  - topico/implementacao
---
```

### Nota De Auditoria

```yaml
---
type: writing
status: validado
tags:
  - tipo/auditoria
  - tipo/writing
  - evidencia/auditoria
  - evidencia/estatistica
  - evidencia/metodologia
  - topico/monografia
  - status/validado
---
```

### Sigla

```yaml
---
sigla: "ACO"
definicao: "Ant Colony Optimization"
incluir: sim
tags:
  - tipo/sigla
  - incluir/sim
---
```

### Indice

```yaml
---
tags:
  - tipo/index
  - tipo/paper
  - status/atualizado
---
```

## Filtros Uteis No Obsidian Graph

Excluir indices e templates:

```text
-tag:#tipo/index -tag:#tipo/template
```

Literatura essencial:

```text
tag:#tipo/paper tag:#relevancia/5
```

Papers pendentes de leitura:

```text
tag:#tipo/paper tag:#status/pendente
```

Metodo ACO:

```text
tag:#metodo/aco
```

Area de drone-routing:

```text
tag:#area/drone-routing
```

Capitulo de experimentos:

```text
tag:#capitulo/experimentos
```

Auditorias concluidas:

```text
tag:#tipo/auditoria tag:#status/concluido
```

Projeto com evidencias de codigo:

```text
tag:#tipo/projeto tag:#evidencia/codigo
```

## Checklist Antes De Finalizar

- A nota tem pelo menos um `tipo/*`?
- Se for `index.md`, tem `tipo/index`?
- Se for paper, `status/*`, `relevancia/*`, `papel/*` e `evidencia/referencia` refletem as propriedades YAML?
- Se for sigla ou projeto com `incluir:`, tem `incluir/*` correspondente?
- **Nao ha flat tags** na lista de tags?
- A classificacao tem suporte no texto da nota, no indice ou nos links?
- O filtro `-tag:#tipo/index -tag:#tipo/template` remove ruido da visualizacao?
