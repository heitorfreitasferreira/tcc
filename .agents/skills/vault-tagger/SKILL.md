---
name: vault-tagger
description: Padroniza tags hierarquicas no vault Obsidian do TCC. Use quando o usuario pedir para tagear notas, organizar tags, filtrar grafo do vault, revisar frontmatter ou melhorar buscas no vault/.
---

# Vault Tagger

Use esta skill para adicionar, revisar ou normalizar tags nas notas Markdown do `vault/` do TCC. O objetivo e melhorar buscas, Graph View, Bases/Dataview e rastreabilidade entre literatura, projeto, escrita e evidencias.

Para propriedades YAML, relacoes tipadas, templates, aliases e claims, use tambem `vault-semantic-schema`. Para editar arquivos `.base`, use `vault-bases-maintainer`.

## Principios

- Preserve o conteudo das notas; altere apenas frontmatter/tags, salvo pedido explicito.
- Prefira tags hierarquicas com prefixo: `tipo/`, `status/`, `area/`, `metodo/`, `papel/`, `capitulo/`, `evidencia/`, `forca/`, `pdf/`, `relevancia/`, `escrita/`, `artefato/`.
- Evite tags soltas novas como `ga`, `pso`, `drone`, `review`, `paper`; normalize para `metodo/ga`, `metodo/pso`, `area/drone-routing`, `papel/revisao`, `tipo/paper`.
- Mantenha tags existentes quando ainda forem usadas por outras notas, mas acrescente a forma nova hierarquica.
- Nao reclassifique status academico, rating, PDF ou evidencia sem inspecionar a nota/indice correspondente.
- Exclua `index.md` dos grafos com `tipo/index` ou `index`; nao apague indices.

## Workflow

1. Identifique o escopo pedido: uma nota, uma pasta (`papers/`, `areas/`, `projeto/`, `writing/`) ou o vault inteiro.
2. Leia o frontmatter da nota e, se necessario, as secoes `Resumo`, `Conexoes`, `Relevancia`, `BibTeX`, `PDF`, `Status` e links `[[...]]`.
3. Aplique tags minimas por tipo de nota.
4. Aplique tags tematicas por area/metodo/papel.
5. Aplique tags operacionais: status, relevancia, PDF, capitulo, evidencia, escrita.
6. Preserve campos estruturados ja existentes (`year`, `status`, `rating`, `doi`, `bibtex`, etc.); tags complementam estes campos para busca/grafo.
7. Ao editar em lote, faca alteracoes pequenas e verificaveis; nunca invente classificacoes incertas.

## Taxonomia Recomendada

### Tipo De Nota

Use exatamente uma ou mais quando fizer sentido:

```yaml
tags:
  - tipo/paper
  - tipo/area
  - tipo/projeto
  - tipo/writing
  - tipo/index
  - tipo/template
  - tipo/auditoria
  - tipo/capitulo
```

Mapeamento por caminho:

- `vault/papers/*.md`: `tipo/paper`
- `vault/areas/*.md`: `tipo/area`
- `vault/projeto/*.md`: `tipo/projeto`
- `vault/writing/**/*.md`: `tipo/writing`
- `*/index.md`: `tipo/index`
- `vault/templates/*.md`: `tipo/template`

### Status

Use junto com o campo `status:` quando existir:

```yaml
tags:
  - status/pendente
  - status/lido-parcial
  - status/lido
  - status/revisar
  - status/atualizado
  - status/desatualizado
  - status/bloqueado
```

### Areas

```yaml
tags:
  - area/tsp
  - area/tsp-variants
  - area/routing
  - area/drone-routing
  - area/bio-inspired
  - area/comparative-studies
  - area/lower-bound
  - area/statistical-analysis
  - area/combinatorial-optimization
```

Regras rapidas:

- TSP classico, complexidade, heuristicas TSP: `area/tsp`
- FSTSP, TSP-D, UAV, drones: `area/drone-routing`
- Roteamento geral/grafos: `area/routing`
- Comparacoes GA/PSO/ACO/outros: `area/comparative-studies`
- Held-Karp, AP relaxation, branch-and-bound, bounding: `area/lower-bound`

### Metodos

```yaml
tags:
  - metodo/ga
  - metodo/pso
  - metodo/aco
  - metodo/bruteforce
  - metodo/lower-bound
  - metodo/hungarian
  - metodo/held-karp
  - metodo/branch-and-bound
  - metodo/local-search
  - metodo/lin-kernighan
  - metodo/simulated-annealing
```

Use `metodo/lower-bound` para a abordagem geral e tags especificas (`metodo/hungarian`, `metodo/held-karp`) quando o texto tratar do metodo concreto.

### Papel Na Monografia

```yaml
tags:
  - papel/fundacional
  - papel/revisao
  - papel/comparativo
  - papel/metodologico
  - papel/aplicacao
  - papel/benchmark
  - papel/limitacao
  - papel/trabalho-futuro
```

Regras rapidas:

- Obras classicas: `papel/fundacional`
- Surveys/books/reviews: `papel/revisao`
- Estudos de comparacao empirica: `papel/comparativo`
- Justifica decisao de modelagem/algoritmo: `papel/metodologico`
- Aplica em drones/patrulha/roteamento operacional: `papel/aplicacao`
- Usado para discutir ameacas, ausencia de tuning, generalizacao: `papel/limitacao`

### Capitulos

```yaml
tags:
  - capitulo/introducao
  - capitulo/fundamentacao
  - capitulo/proposta
  - capitulo/experimentos
  - capitulo/conclusao
  - capitulo/apendice
```

Use em notas `writing/`, `projeto/` e `papers/` quando a nota alimenta diretamente um capitulo.

### Evidencia

```yaml
tags:
  - evidencia/codigo
  - evidencia/dados
  - evidencia/resultado
  - evidencia/figura
  - evidencia/tabela
  - evidencia/referencia
  - evidencia/auditoria
```

Regras rapidas:

- Notas que apontam para `src/`: `evidencia/codigo`
- Notas que apontam para `src/data/results/`: `evidencia/dados` ou `evidencia/resultado`
- Notas sobre `monografia/figs/`: `evidencia/figura`
- Artigos/papers: `evidencia/referencia`
- Auditorias em `writing/`: `evidencia/auditoria`

### Forca Da Evidencia

Use especialmente em `writing/planejamento/claim-evidence-matrix.md`, auditorias e notas de resultados:

```yaml
tags:
  - forca/forte
  - forca/moderada
  - forca/fraca
  - forca/bloqueada
  - forca/requer-validacao
```

### Escrita

Use para controlar maturidade das notas de escrita:

```yaml
tags:
  - escrita/rascunho
  - escrita/revisar
  - escrita/pronto
  - escrita/bloqueado
  - escrita/desatualizado
```

### Artefatos

```yaml
tags:
  - artefato/src
  - artefato/data
  - artefato/resultados
  - artefato/figuras
  - artefato/tabelas
  - artefato/latex
  - artefato/bibtex
```

### PDF

Use somente quando houver evidencia no indice de papers ou na nota:

```yaml
tags:
  - pdf/integro
  - pdf/ausente
  - pdf/corrompido
  - pdf/capes
  - pdf/busca-manual
```

### Relevancia

Espelhe `rating:` para uso no Graph View:

```yaml
tags:
  - relevancia/5
  - relevancia/4
  - relevancia/3
  - relevancia/baixa
```

Use `relevancia/baixa` para `rating: 0`, `1` ou `2`.

## Exemplos

### Paper Fundacional De ACO

```yaml
---
title: Dorigo et al. (1996)
year: 1996
status: lido
rating: 5
tags:
  - tipo/paper
  - status/lido
  - area/tsp
  - area/bio-inspired
  - metodo/aco
  - papel/fundacional
  - evidencia/referencia
  - capitulo/fundamentacao
  - relevancia/5
  - pdf/integro
---
```

### Nota De Projeto Do GA

```yaml
---
tags:
  - tipo/projeto
  - metodo/ga
  - area/bio-inspired
  - area/tsp
  - evidencia/codigo
  - capitulo/proposta
  - artefato/src
---
```

### Nota De Experimentos

```yaml
---
tags:
  - tipo/writing
  - tipo/capitulo
  - capitulo/experimentos
  - evidencia/resultado
  - evidencia/tabela
  - evidencia/figura
  - artefato/resultados
  - artefato/figuras
  - escrita/revisar
---
```

### Indice

```yaml
---
tags:
  - tipo/index
  - tipo/paper
  - catalogo
  - tracker
  - status/atualizado
---
```

## Filtros Uteis No Obsidian Graph

Excluir indices:

```text
-tag:#tipo/index
```

Literatura essencial:

```text
tag:#tipo/paper -tag:#tipo/index tag:#relevancia/5
```

Papers pendentes:

```text
tag:#tipo/paper tag:#status/pendente -tag:#tipo/index
```

Metodo ACO:

```text
tag:#metodo/aco -tag:#tipo/index
```

Capitulo de experimentos:

```text
tag:#capitulo/experimentos -tag:#tipo/index
```

Claims com evidencia fraca:

```text
tag:#forca/fraca OR tag:#forca/requer-validacao OR tag:#forca/bloqueada
```

## Checklist Antes De Finalizar

- A nota tem `tipo/...` correto?
- Se for `index.md`, tem `tipo/index`?
- Se for paper, `status/...`, `relevancia/...`, `papel/...` e `evidencia/referencia` refletem campos existentes?
- Tags soltas antigas foram mantidas apenas se necessario, mas acompanhadas por tags hierarquicas?
- A classificacao tem suporte no texto da nota, no indice ou nos links?
- O filtro `-tag:#tipo/index` remove indices da visualizacao?
