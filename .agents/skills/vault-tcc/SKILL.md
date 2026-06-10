---
name: vault-tcc
description: Gerencia o vault Obsidian do TCC. Cobre estrutura do vault, workflow de importação de papers, enriquecimento de notas, manutenção de Bases, schema semântico, taxonomia de tags hierárquicas, e expansão da base de conhecimento. Consolida vault-reference, vault-bases-maintainer, vault-semantic-schema, vault-tagger e knowledge-base.
---

# Vault TCC

Skill unificada para manutenção do vault Obsidian do TCC.

---

## 1. Estrutura do Vault

```
vault/
├── papers/              ← Notas markdown (uma por paper)
├── areas/               ← Notas de área de pesquisa
├── claims/              ← Claims individuais
├── projeto/             ← Notas de projeto
├── siglas/              ← Notas de sigla
├── writing/             ← Notas de escrita (capítulos, auditorias, planejamento)
├── templates/           ← Templates de nota
├── bases/               ← Arquivos .base (views do Obsidian Bases)
├── canvas/              ← Knowledge graph (.canvas)
└── opencode-vault.md    ← Instruções completas para agentes
```

**Estado atual:** 156 notas, zero flat tags, 11 namespaces ativos de tags hierárquicas.

---

## 2. Workflow de Nova Referência

### Passo 1 — Identificar Lacuna
- Liste keys existentes: `ls vault/papers/*.md | sed 's/.*\///; s/\.md$//'`
- Priorize papers recentes (2020+) e surveys
- Tópicos: TSP, GA/PSO/ACO para TSP, híbridos, drone/TSP-D/rTSP, patrulha UAV

### Passo 2 — Buscar
Use `websearch` e `google-scholar` MCP em paralelo:
- "traveling salesman problem survey 2024", "genetic algorithm TSP hybrid"
- Para cada candidato: título, autores, ano, DOI/URL, abstract, journal/conference

### Passo 3 — Validar Metadata
Confirme DOIs via CrossRef: `webfetch https://api.crossref.org/works/{doi}`

### Passo 4 — Adicionar BibTeX
Append em `monografia/bib/abntex2-references.bib`:
```bibtex
@article{key2024topic,
  title   = {Título Exato},
  author  = {Sobrenome, Nome and Sobrenome2, Nome2},
  journal = {Journal},
  volume  = {X}, number = {Y}, pages = {Z--ZZ},
  year    = {2024}, doi     = {10.xxxx/xxxxx}
}
```
Regras: key = `{primeiroautor}{ano}{topico}`, nunca truncar autores, incluir DOI sempre.

### Passo 5 — Importar para o Vault
```bash
bash scripts/import-bib-to-vault.sh <bibtex-key>
# ou todos pendentes:
bash scripts/import-bib-to-vault.sh
```

### Passo 6 — Baixar PDF (quando disponível)
```bash
bash scripts/download-pdfs.sh --dry-run --min-rating 4
bash scripts/download-pdfs.sh --limit 10 --min-rating 4 --use-scihub
bash scripts/download-pdfs.sh --keys toaza2023review,larranaga1999ga
```
Fontes: Unpaywall → Semantic Scholar → OpenAlex → Sci-Hub (`--use-scihub`).
Se obtiver PDF: `pdf: "papers/pdfs/<bibtex-key>.pdf"`. Se não: `pdf: ""`.

### Passo 7 — Enriquecer Nota
Preencha cada seção da nota gerada:
- **Resumo**: 2-5 frases da contribuição central
- **Contribuições Principais**: bullet list
- **Relevância para o TCC**: conexão com patrulha de drones, rTSP, comparação de metaheurísticas
- **Métodos e Abordagens**: algoritmos, datasets, técnicas
- **Conexões**: `[[wiki links]]` para notas de área
- **Notas e Insights**: observações críticas, limitações, ideias
- **Citações-chave**: 1-2 citações impactantes com `>`
Prosa em português. Use convenções `obsidian-markdown`.

### Passo 8 — Atualizar Canvas
Leia `vault/canvas/tcc-knowledge-graph.canvas`, adicione node + edges. Posicione papers novos abaixo dos existentes (incremente y em ~70-80).

### Passo 9 — Verificar
- Frontmatter YAML sem erros
- `[[wiki links]]` apontam para arquivos reais
- `git status` mostra apenas mudanças intencionais

---

## 3. Schema Semântico

### Regra Principal

Não transforme o vault em fonte primária para claims metodológicos/experimentais.
- Claims metodológicos → `src/`
- Claims experimentais → `src/data/results/`, `scripts/`, `monografia/figs/`
- Notas do vault organizam e explicam evidências, não substituem código/dados

### Schema Mínimo por Tipo

**Papers** (`vault/papers/*.md`):
```yaml
title: ""; authors: []; year:; doi: ""; bibtex_key: ""; bibtex-key: ""
pdf: ""; type: paper
reading_status: pendente  # pendente, resumo-lido, lido-parcial, lido
validation_status: nao-validado  # nao-validado, validado, requer-validacao, bloqueado
pdf_status: ausente  # disponivel, ausente, corrompido
rating: 0
role: ""  # fundacional, revisao, comparativo, benchmark, metodologico, aplicacao
areas: []; methods: []; chapters: []; claim_support: []
aliases: []
tags: [tipo/paper, status/pendente, evidencia/referencia]
```

**Claims** (`vault/claims/*.md`):
```yaml
type: claim; claim_id: ""; claim: ""; claim_type: experimental
status: requer-validacao; strength: moderada
primary_evidence: []; vault_support: []
monografia_section: ""; usage_guidance: ""; rationale: ""
action_required: ""; last_verified: ""
tags: [tipo/claim, claim/experimental, status/requer-validacao]
```
Prefixos de claim_id: `C` conceitual, `M` metodológico, `A` algoritmo, `E` experimental, `I` interpretativo, `B` bloqueado.

**Siglas** (`vault/siglas/*.md`):
```yaml
sigla: ""; definicao: ""; incluir: pendente
ocorrencias_ac: 0; ocorrencias_texto: 0
tags: [tipo/sigla, incluir/pendente]
```

**Projeto** (`vault/projeto/*.md`):
```yaml
type: projeto; writing_status: revisar; validation_status: validado
areas: []; methods: []; chapters: []; primary_evidence: [src/]
tags: [tipo/projeto, evidencia/codigo]
```

**Writing** (`vault/writing/**/*.md`):
```yaml
type: writing; writing_status: revisar; validation_status: requer-validacao
chapters: []; primary_evidence: []
tags: [tipo/writing]
```

### Seções Padronizadas para Paper

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

### Relações Tipadas

Em `## Conexões`:
```markdown
- Fundamenta: [[ant-colony]]
- Relacionado a: [[tsp]], [[bio-inspired-optimization]]
- Contrasta com: [[pso]]
- Apoia claim: E22
- Usado em capítulo: [[experimentos]]
```

### Aliases Recomendados

```yaml
aliases:
  - TSP-SD-ATP
  - rTSP
  - Ant Colony Optimization
  - ACO
```

---

## 4. Taxonomia de Tags

**Apenas tags hierárquicas** `namespace/valor`. Nunca flat tags.

### Namespaces Ativos

| Namespace | Uso | Exemplo |
|-----------|-----|---------|
| `tipo/` | Classificação primária | `tipo/paper` |
| `area/` | Domínio de pesquisa | `area/tsp` |
| `metodo/` | Método/técnica | `metodo/aco` |
| `papel/` | Função na monografia | `papel/revisao` |
| `status/` | Estado de leitura/escrita | `status/lido` |
| `evidencia/` | Tipo de evidência | `evidencia/referencia` |
| `capitulo/` | Capítulo alimentado | `capitulo/fundamentacao` |
| `relevancia/` | Espelha `rating:` 1-5 | `relevancia/5` |
| `incluir/` | Inclusão na monografia | `incluir/sim` |
| `topico/` | Tópico transversal | `topico/implementacao` |
| `forca/` | Força da evidência | `forca/requer-validacao` |

### Mapeamento Canônico por Caminho

| Caminho | `tipo/*` |
|---------|----------|
| `papers/*.md` | `tipo/paper` |
| `areas/*.md` | `tipo/area` |
| `projeto/*.md` | `tipo/projeto` |
| `siglas/*.md` | `tipo/sigla` |
| `writing/auditorias/*.md` | `tipo/auditoria` |
| `writing/**/*.md` (demais) | `tipo/writing` |
| `*/index.md` | `tipo/index` |
| `templates/*.md` | `tipo/template` |

### Workflow de Tag

1. Identifique o escopo: uma nota, uma pasta, ou vault inteiro.
2. Leia frontmatter e seções relevantes.
3. Aplique `tipo/*` obrigatório.
4. Aplique tags temáticas (`area/`, `metodo/`, `papel/`).
5. Aplique tags operacionais (`status/`, `relevancia/`, `capitulo/`, `evidencia/`).
6. Preserve campos estruturados YAML.
7. Para lote, use `python3 scripts/migrate-tags.py`.

### Filtros Úteis no Graph

- Excluir índices: `-tag:#tipo/index -tag:#tipo/template`
- Literatura essencial: `tag:#tipo/paper tag:#relevancia/5`
- Pendentes: `tag:#tipo/paper tag:#status/pendente`
- Método ACO: `tag:#metodo/aco`
- Drone routing: `tag:#area/drone-routing`

---

## 5. Manutenção de Bases

### Regra Principal

Bases são views, não fonte de verdade.
- Não escreva evidência dentro de `.base`
- Não duplique tabelas de resultados nas Bases
- Use Bases para encontrar pendências e inconsistências no frontmatter
- Fonte primária: `src/`, `src/data/results/`, `scripts/`, `monografia/figs/`

### Workflow para Editar Base

1. Leia o arquivo `.base` inteiro.
2. Confira propriedades do template correspondente.
3. Mantenha filtros simples e tolerantes a notas antigas.
4. Use formulas para compatibilidade entre schema antigo e novo.
5. Evite formulas complexas se propriedade explícita resolver.
6. Valide YAML depois de editar.
7. Se criar nova Base, documente em `vault/opencode-vault.md`.

### Padrões de Compatibilidade

```yaml
formulas:
  leitura: 'if(reading_status, reading_status, status)'
  estado: 'if(validation_status, validation_status, status)'
  pdf: 'if(pdf_status, pdf_status, "")'
```

### Bases Existentes

- `papers.base`: papers, prioridade de leitura, PDF, uso na monografia
- `claims.base`: sumário consultável dos claims individuais
- `siglas.base`: sumário consultável das siglas

### Boas Práticas

- Nomes kebab-case: `papers.base`, `project-evidence.base`
- Views em português claro: `Prioridade de Leitura`
- Display names curtos: `Relevancia`, `Validacao`
- Validar YAML: `ruby -e 'require "yaml"; YAML.load_file(path)' vault/bases/*.base`

### Anti-Patterns de Bases

- Usar Base como banco de dados manual
- Colocar texto longo em propriedade
- Criar uma Base para cada tag
- Remover fallback para notas antigas antes da migração completa
- Filtrar por tag quando `file.inFolder` ou propriedade `type` é mais robusta

---

## 6. Atualização de Claims

1. Prefixo por tipo: `C`, `M`, `A`, `E`, `I`, `B`.
2. Claim com escopo delimitado.
3. Defina força e status.
4. Aponte para fonte primária verificável.
5. Aponte para apoio no vault.
6. Indique uso recomendado na monografia.
7. Atualize `usage_guidance`, `rationale`, `action_required` quando relevante.
8. Atualize `last_verified`.
9. Confirme que `claims.base` lista o novo registro.

Não promova claim interpretativo para forte sem evidência primária.

---

## 7. Verificação de Consistência

```bash
# Tags
python3 scripts/migrate-tags.py  # Deve reportar "Changed: 0"

# YAML
python3 -c "
import yaml, sys
with open(sys.argv[1]) as f:
    text = f.read()
_, rest = text.split('---\n', 1)
end = rest.find('\n---\n')
fm_raw = rest[:end] if end != -1 else rest
fm = yaml.safe_load(fm_raw)
flat = [t for t in fm.get('tags',[]) if '/' not in t]
if flat: print(f'FLAT TAGS: {flat}')
else: print('OK')
" vault/papers/arquivo.md
```

---

## 8. Anti-Patterns Gerais

- Usar canvas como evidência primária
- Trocar `bibtex-key` por `bibtex_key` removendo compatibilidade
- Duplicar a mesma verdade em paper, claim note, Base e canvas com valores divergentes
- **Adicionar flat tags** — o vault usa exclusivamente tags hierárquicas
- Retagear manualmente — use `scripts/migrate-tags.py`
- Paper sem DOI verificado (exceção: arXiv)
- Truncar lista de autores
- Deixar nota gerada vazia — enriqueça imediatamente
- Adicionar papers já existentes (verifique primeiro)
- Paper tangencial sem conexão com áreas do vault

## Skills Externas

- `obsidian-markdown` — callouts, `[[links]]`
- `obsidian-vault` — operações de vault
- `obsidian-bases` — API genérica de Bases
- `json-canvas` — edição de `.canvas`
- `academic-researcher` — busca de papers
- `academic-search` — estratégias de busca
