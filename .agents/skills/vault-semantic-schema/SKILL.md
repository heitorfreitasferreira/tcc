---
name: vault-semantic-schema
description: Ensina a manter o schema semantico do vault Obsidian do TCC. Use quando editar vault/, templates, papers, claims, claim-evidence-matrix, propriedades YAML, aliases, relacoes tipadas ou rastreabilidade monografia-evidencia.
---

# Vault Semantic Schema

Use esta skill para criar, revisar ou corrigir notas do `vault/` seguindo o schema semantico incremental aprovado para o TCC. Ela complementa `vault-tagger`: use `vault-tagger` para escolher tags; use esta skill para decidir propriedades, secoes, relacoes e rastreabilidade.

## Fontes Canonicas

- Convenções gerais: `vault/opencode-vault.md`
- Template de paper: `vault/templates/paper-note.md`
- Template opcional de claim: `vault/templates/claim-note.md`
- Registro canonico de claims: `vault/writing/claim-evidence-matrix.md`
- Bases de consulta: `vault/bases/papers.base` e `vault/bases/claims.base`
- Importador de BibTeX: `scripts/import-bib-to-vault.sh`

## Regra Principal

Nao transforme o vault em fonte primaria para claims metodologicos ou experimentais.

- Claims metodologicos apontam para `src/`.
- Claims experimentais apontam para `src/data/results/`, `scripts/` ou `monografia/figs/`.
- Notas do vault organizam e explicam evidencias, mas nao substituem codigo/dados.
- `claim-evidence-matrix.md` e a fonte canonica de IDs, forca e status dos claims.

## Workflow Para Editar Nota

1. Identifique o tipo da nota: paper, area, projeto, writing, claim, index ou template.
2. Leia o frontmatter atual antes de editar.
3. Preserve propriedades existentes usadas por scripts, especialmente `bibtex-key` em papers.
4. Adicione propriedades novas sem remover compatibilidade.
5. Use tags hierarquicas apenas como camada de navegacao/grafo.
6. Padronize secoes somente quando isso melhorar busca, escrita ou rastreabilidade.
7. Se a nota sustenta claim forte, conecte ao ID da matriz em `claim_support` ou em secao textual.
8. Valide YAML/frontmatter quando alterar muitas propriedades.

## Schema Minimo Por Tipo

### Papers

Use em `vault/papers/*.md` e no template de paper:

```yaml
---
title: ""
authors: []
year:
doi: ""
bibtex_key: ""
bibtex-key: ""
pdf: ""
type: paper
reading_status: pendente
validation_status: nao-validado
pdf_status: ausente
rating: 0
role: ""
areas: []
methods: []
chapters: []
claim_support: []
aliases: []
tags:
  - tipo/paper
  - status/pendente
  - evidencia/referencia
---
```

Campos importantes:

- `reading_status`: `pendente`, `lido-parcial`, `lido`
- `validation_status`: `nao-validado`, `validado`, `requer-validacao`, `bloqueado`
- `pdf_status`: `integro`, `ausente`, `corrompido`, `capes`, `busca-manual`
- `role`: `fundacional`, `revisao`, `comparativo`, `metodologico`, `aplicacao`, `benchmark`, `limitacao`, `trabalho-futuro`
- `areas`: valores sem prefixo, por exemplo `tsp`, `drone-routing`, `lower-bound`
- `methods`: valores sem prefixo, por exemplo `ga`, `pso`, `aco`, `bruteforce`, `held-karp`
- `chapters`: `introducao`, `fundamentacao`, `proposta`, `experimentos`, `conclusao`, `apendice`
- `claim_support`: IDs da matriz, por exemplo `C04`, `A08`, `E22`

### Notas De Projeto

Use para `vault/projeto/*.md`:

```yaml
---
type: projeto
writing_status: revisar
validation_status: validado
areas: []
methods: []
chapters: []
primary_evidence:
  - src/
tags:
  - tipo/projeto
  - evidencia/codigo
---
```

Notas de projeto devem apontar para arquivos reais em `src/` quando descrevem comportamento implementado.

### Notas De Writing

Use para `vault/writing/*.md`:

```yaml
---
type: writing
writing_status: revisar
validation_status: requer-validacao
chapters: []
primary_evidence: []
tags:
  - tipo/writing
---
```

Se for auditoria, inclua `tipo/auditoria` e `evidencia/auditoria`.

### Claims

Nao crie uma pasta de claims por padrao. A matriz e suficiente para claims simples.

Use `vault/templates/claim-note.md` somente quando um claim precisar de nota propria por ser complexo, controverso ou de alto risco.

Schema minimo:

```yaml
---
type: claim
claim_id: ""
claim: ""
claim_type: experimental
status: requer-validacao
strength: moderada
primary_evidence: []
vault_support: []
monografia_section: ""
last_verified: ""
tags:
  - tipo/claim
  - evidencia/auditoria
  - forca/requer-validacao
---
```

## Secoes Padronizadas Para Paper

Use estas secoes quando criar ou enriquecer paper:

```markdown
## PDF

## Tese Central

## Resumo

## Contribuições Principais

## Relevância para o TCC

## Uso no TCC

## Métodos e Abordagens

## Evidência / Resultado Relevante

## Limitações de Uso

## Conexões

## Notas e Insights

## Citações-chave
```

Nao force preenchimento artificial. Se o paper esta pendente, deixe placeholders claros.

## Relacoes Tipadas

Em `## Conexões`, prefira relacoes semanticamente nomeadas:

```markdown
- Fundamenta: [[ant-colony]]
- Relacionado a: [[tsp]], [[bio-inspired-optimization]]
- Contrasta com: [[pso]]
- Apoia claim: E22
- Usado em capítulo: [[experimentos]]
```

Use relacoes tipadas principalmente em notas importantes. Nao precisa aplicar a todas as notas antigas de uma vez.

## Aliases Recomendados

Adicione `aliases` para termos de busca frequentes:

```yaml
aliases:
  - TSP-SD-ATP
  - rTSP
  - TSP com penalidade angular
  - TSP com custo dependente do predecessor
```

Para metodos:

```yaml
aliases:
  - Ant Colony Optimization
  - ACO
  - Otimização por Colônia de Formigas
```

## Como Atualizar Claims

Ao adicionar claim em `claim-evidence-matrix.md`:

1. Escolha prefixo por tipo: `C` conceitual, `M` metodologico, `A` algoritmo, `E` experimental, `I` interpretativo, `B` bloqueado.
2. Escreva o claim com escopo delimitado.
3. Defina força e status.
4. Aponte para fonte primaria verificavel.
5. Aponte para apoio no vault.
6. Indique uso recomendado na monografia.
7. Atualize `last_verified` se a checagem foi feita agora.

Nao promova claim interpretativo para forte sem evidencia primaria.

## Anti-Patterns

- Criar `vault/claims/` para todo claim simples.
- Usar canvas como evidencia primaria.
- Trocar `bibtex-key` por `bibtex_key` removendo compatibilidade com scripts.
- Duplicar a mesma verdade em paper, claim note, Base e canvas com valores divergentes.
- Retagear o vault inteiro sem etapa de auditoria.
- Remover tags antigas em massa sem confirmar uso em buscas ou notas.

## Verificacao

Para mudancas em `.base`, use tambem a skill `vault-bases-maintainer`.

Para frontmatter YAML, valide com uma ferramenta local quando possivel. Exemplo:

```bash
ruby -e 'require "yaml"; ARGV.each { |p| t=File.read(p); fm=t[/\A---\n(.*?)\n---/m,1]; YAML.safe_load(fm) if fm; puts "ok #{p}" }' vault/templates/paper-note.md
```
