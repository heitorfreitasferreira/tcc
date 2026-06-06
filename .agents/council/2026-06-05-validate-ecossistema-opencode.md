# Council Report — Validação do Ecossistema opencode do TCC

**Data**: 2026-06-05  
**Modo**: validate (3 juízes, sem perspectivas)  
**Consenso**: **WARN** (Judge-1: WARN, Judge-2: WARN, Judge-3: WARN)

---

## Sumário

O ecossistema opencode do TCC cobre bem o núcleo de pesquisa e incorporação de literatura (comandos, skills, MCPs, vault), mas tem lacunas críticas nas **pontas do pipeline** — experimentação e compilação/validação — e **sobreposições** significativas entre skills de escrita, skills de vault e MCPs de busca que aumentam a carga cognitiva sem benefício proporcional. O ecossistema é **funcional para o meio do fluxo** (pesquisa → incorporação → enriquecimento), mas **quebrado nas extremidades** (sem experimentação automatizada, sem compilação, sem rastreabilidade claim→dado).

---

## Falhas Críticas

| # | Severidade | Descrição | Juízes | Ref |
|---|---|---|---|---|
| 1 | critical | **Sem comando de experimentação**: `src/run_all.sh` e `run_experiments_multi_seed.sh` existem mas não são expostos como comandos opencode. O agente não consegue gerar dados experimentais para sustentar claims. | J1, J3 | `AGENTS-experiments.md` |
| 2 | critical | **Sem comando de compilação**: agente edita `.tex` mas não pode verificar se compila (pdflatex → bibtex → pdflatex×2). Edições são cegas. | J3 | `monografia/main_ppgco_ufu.tex` |
| 3 | critical | **Sem rastreabilidade claim→dado**: claims quantitativas na monografia (ex: "ACO atinge 71.79%") não são validadas automaticamente contra `src/data/results/summary/*.json`. | J3 | `.agents/council/2026-06-02-validate-monograph.md:32` |
| 4 | critical | **Council skill ausente do inventário**: workflow e `/roadmap-consumir` referenciam `council`, mas não está nas 14 skills listadas no ecossistema. | J1 | Workflow passo 7 |
| 5 | critical | **Loop /claudiney → /roadmap quebrado**: feedback do orientador gera análise mas não cria tarefas automaticamente. Sem rastreabilidade feedback→ação. | J1 | `/claudiney`, `/roadmap-criar` |

## Problemas Significantes (Redundâncias & Integração)

| # | Severidade | Descrição | Juízes |
|---|---|---|---|
| 6 | significant | **4 skills de escrita sobrepostas**: `academic-writing`, `research-paper-writing`, `scientific-paper`, `latex-document-skill` sem matriz de decisão. Agente pode carregar skill errada para o contexto (ex: ML-paper skill para monografia TSP em PT-BR). | J1, J2, J3 |
| 7 | significant | **2 skills de estatística sobrepostas**: `data-analysis` e `statistical-analysis` cobrem o mesmo domínio com fronteiras difusas. | J1 |
| 8 | significant | **5 MCPs de busca redundantes**: `crossref`, `semantic-scholar`, `scholar-sidekick`, `google-scholar`, `academic-search`. `/consultar-academico` dispara 4 em paralelo com resultados duplicados. | J1, J2, J3 |
| 9 | significant | **Sem comando de figuras**: `scripts/gerar-figuras.sh` e `scripts/gerar-analises.sh` não expostos como comandos. Servidor `/api/render` sem integração programática. | J1, J3 |
| 10 | significant | **codegraphcontext não integrado**: MCP capaz de indexar `src/` e ligar símbolos Go a claims do vault não está listado nem configurado. | J3 |
| 11 | significant | **Sem validação bibliográfica**: `scholar-sidekick.verifyCitation` e `checkRetraction` disponíveis mas não usados. `.bib` nunca validado quanto a retratações ou metadados incorretos. | J3 |
| 12 | significant | **/baixar-pdf depreciado sem substituto funcional**: `/incorporar` precisa de PDF; se `/baixar-pdf` não funciona e PDF não existe localmente, pipeline para no primeiro passo. | J3 |
| 13 | significant | **Alta carga cognitiva**: ~30 entidades (14 skills + 9 MCPs + 6 comandos + pipeline) para o pesquisador modelar mentalmente. Sem `/help`, `/which-tool` ou matriz de decisão. | J2 |
| 14 | significant | **/incorporar monolítico**: 9 fases em um só comando; pesquisador que só quer atualizar o canvas precisa re-executar pipeline completo. | J2 |

## Problemas Menores

| # | Severidade | Descrição | Juízes |
|---|---|---|---|
| 15 | minor | **vault-semantic-schema vs vault-tagger**: sobreposição em "tags hierárquicas" sem boundary clara. | J1, J2 |
| 16 | minor | **/baixar-pdf no inventário ativo**: comando depreciado listado entre os ativos. | J1 |
| 17 | minor | **pdf-watcher passivo**: detecta PDFs mas não dispara ação (não alimenta `/incorporar fila` automaticamente). | J1 |
| 18 | minor | **Sem skill de domínio TSP**: conhecimento sobre TSP-SD-ATP, propriedades das meta-heurísticas, armadilhas conhecidas está implícito. | J3 |
| 19 | minor | **arxiv watch não configurado**: vigilância contínua de literatura disponível mas não usada. | J3 |
| 20 | minor | **Sem /consultar-base**: `vault/bases/` existe mas sem comando para consulta programática. | J3 |
| 21 | minor | **Fluxo documentado em ordem errada**: ANÁLISE após FIGURAS (deveria ser o inverso) e EXPERIMENTAÇÃO ausente. | J3 |
| 22 | minor | **Sem estado de sessão**: ecossistema não mantém registro de "onde você parou" entre sessões. | J2 |

---

## Recomendações por Prioridade

### Imediato (antes do próximo ciclo de escrita)

1. **Criar `/experimento`** — wrapper para `make build && tcc optimize` com parâmetros (--method, --instance, --seeds)
2. **Criar `/compilar`** — ciclo pdflatex→bibtex→pdflatex×2 com captura de erros
3. **Criar `/consistencia`** — validação cruzada claims `.tex` ↔ `results/summary/*.json`
4. **Adicionar `council` ao inventário** de skills do ecossistema
5. **Estender `/claudiney`** com `--criar-tarefas` para fechar loop com roadmap

### Curto prazo (próximo sprint)

6. **Criar `/figuras`** — wrapper para `scripts/gerar-figuras.sh` com gestão de ciclo de vida do servidor
7. **Criar `/auditar-bib`** — validação de retratações e consistência de metadados no `.bib`
8. **Integrar `codegraphcontext`** — indexar `src/` e cruzar símbolos Go com claims
9. **Documentar matriz de decisão** para skills de escrita, estatística e MCPs de busca no `AGENTS.md`
10. **Criar `/which-tool`** — mapeamento intenção → ferramenta para reduzir carga cognitiva
11. **Configurar `arxiv.watch_topic`** — vigilância contínua de TSP + drones + meta-heurísticas

### Longo prazo (refinamento)

12. **Consolidar skills de escrita** em uma com dispatch interno por modo
13. **Criar skill `tsp-optimization`** com conhecimento de domínio (TSP-SD-ATP, armadilhas, convenções)
14. **Criar `/consultar-base`** para consultas programáticas às Obsidian Bases
15. **Corrigir ordem do fluxo** no `AGENTS.md`: +EXPERIMENTAÇÃO, +COMPILAÇÃO, +CONSISTÊNCIA; reordenar ANÁLISE→FIGURAS→ESCRITA
16. **Dar ação ao pdf-watcher** — escrever em fila para `/incorporar fila`
17. **Remover `/baixar-pdf`** do inventário ativo ou criar v2 funcional

---

## Fluxo Corrigido Proposto

```
PESQUISA → INCORPORAÇÃO → ENRIQUECIMENTO → EXPERIMENTAÇÃO → ANÁLISE → ESCRITA → FIGURAS → COMPILAÇÃO → CONSISTÊNCIA → REVISÃO → TAREFAS → VALIDAÇÃO → CÓDIGO
```

Mudanças: +EXPERIMENTAÇÃO, +COMPILAÇÃO, +CONSISTÊNCIA; ANÁLISE antes de ESCRITA e FIGURAS; CÓDIGO como etapa final (consequência de validação).

---

## Arquivos dos juízes

- [Judge 1](./2026-06-05-validate-ecossistema-opencode-judge-1.md) — Arquitetura geral
- [Judge 2](./2026-06-05-validate-ecossistema-opencode-judge-2.md) — Ergonomia e experiência
- [Judge 3](./2026-06-05-validate-ecossistema-opencode-judge-3.md) — Capacidades faltantes
