---
name: vault-bases-maintainer
description: Ensina a criar e manter Obsidian Bases do vault do TCC. Use quando editar vault/bases/*.base, criar views de papers/claims, corrigir filtros, propriedades, formulas ou consultas de Obsidian Bases.
---

# Vault Bases Maintainer

Use esta skill para manter os arquivos `.base` do Obsidian em `vault/bases/`. Ela complementa a skill generica `obsidian-bases`, mas aplica as convenções especificas do vault do TCC.

## Arquivos Atuais

- `vault/bases/papers.base`: consulta papers, prioridade de leitura, PDF e uso na monografia.
- `vault/bases/claims.base`: consulta a matriz canonica de claims e eventuais notas individuais de claim.

## Regra Principal

Bases sao views, nao fonte de verdade.

- Nao escreva evidencia dentro de `.base`.
- Nao duplique tabelas de resultados experimentais nas Bases.
- Use Bases para encontrar pendencias e inconsistencias no frontmatter.
- A fonte primaria continua em `src/`, `src/data/results/`, `scripts/`, `monografia/figs/` e literatura.

## Workflow Para Editar Uma Base

1. Leia o arquivo `.base` inteiro antes de editar.
2. Confira quais propriedades existem no template correspondente.
3. Mantenha filtros simples e tolerantes a notas antigas.
4. Use formulas para compatibilidade entre schema antigo e novo.
5. Evite formulas complexas se uma propriedade explicita resolver.
6. Valide YAML depois de editar.
7. Se criar uma nova Base, documente em `vault/opencode-vault.md`.

## Padroes De Compatibilidade

O vault tem notas antigas com `status` e notas novas com `reading_status`, `writing_status` ou `validation_status`. Use formulas de fallback:

```yaml
formulas:
  leitura: 'if(reading_status, reading_status, status)'
  estado: 'if(validation_status, validation_status, status)'
```

Para propriedades novas que podem estar vazias:

```yaml
formulas:
  pdf: 'if(pdf_status, pdf_status, "")'
  secao: 'if(monografia_section, monografia_section, "")'
```

## Base De Papers

Escopo recomendado:

```yaml
filters:
  and:
    - 'file.inFolder("papers")'
    - 'file.ext == "md"'
    - 'file.basename != "index"'
```

Views uteis:

- `Papers - Visao Geral`: leitura, rating, PDF, papel, areas, metodos, capitulos.
- `Prioridade de Leitura`: notas nao lidas, agrupadas por prioridade derivada de `rating`.
- `Sem PDF Integro`: papers com `pdf_status != "integro"`.
- `Uso Na Monografia`: `role`, `chapters`, `claim_support`.

Propriedades esperadas:

```yaml
type: paper
reading_status: pendente
pdf_status: ausente
rating: 0
role: ""
areas: []
methods: []
chapters: []
claim_support: []
```

## Base De Claims

Escopo recomendado:

```yaml
filters:
  or:
    - 'type == "claim"'
    - 'type == "claim-registry"'
    - 'file.hasTag("tipo/claim")'
    - and:
        - 'file.inFolder("writing")'
        - 'file.basename == "claim-evidence-matrix"'
```

Views uteis:

- `Claims - Registro`: mostra registro canonico e notas individuais.
- `Requer Validacao`: claims/notas cujo estado nao e `validado`.
- `Claims Experimentais`: filtra `claim_type == "experimental"`.

Propriedades esperadas em notas individuais de claim:

```yaml
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
```

## Criando Nova Base

Crie uma nova `.base` somente quando houver uma pergunta recorrente que nao e bem respondida pelas Bases atuais.

Bons candidatos:

- `writing.base`: status de capitulos, auditorias e notas de escrita.
- `project-evidence.base`: notas de projeto por metodo e fonte em `src/`.
- `pdfs.base`: fila de PDFs ausentes/corrompidos se `papers.base` ficar grande demais.

Evite Bases decorativas ou que apenas repetem indices manuais.

## Nomes E Idioma

- Arquivos: kebab-case, por exemplo `papers.base`, `project-evidence.base`.
- Views: portugues claro, por exemplo `Prioridade de Leitura`.
- Display names: curtos, por exemplo `Relevancia`, `Validacao`, `Evidencia Primaria`.

## Validacao

Valide sintaxe YAML depois de editar:

```bash
ruby -e 'require "yaml"; ARGV.each { |path| YAML.load_file(path); puts "ok #{path}" }' vault/bases/papers.base vault/bases/claims.base
```

Se Obsidian mostrar erro mesmo com YAML valido, revise:

- strings com `:` sem aspas;
- formulas com aspas aninhadas;
- `formula.X` usado em `order` sem existir em `formulas`;
- filtros que referenciam propriedade com nome errado;
- expressions muito ambiciosas para o parser do Bases.

## Anti-Patterns

- Usar Base como banco de dados manual.
- Colocar texto longo em propriedade de Base.
- Criar uma Base para cada tag.
- Remover fallback para notas antigas antes da migracao completa.
- Filtrar por tag quando `file.inFolder` ou propriedade `type` e mais robusta.
