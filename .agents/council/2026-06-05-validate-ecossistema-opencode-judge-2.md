```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "key_insight": "The ecosystem is functionally complete but imposes high cognitive load via overlapping tool surfaces, fuzzy skill boundaries, and insufficient discoverability scaffolding — a researcher can accomplish everything but spends too much mental energy navigating the toolset rather than thinking about the monograph.",
  "findings": [
    {
      "severity": "significant",
      "category": "integration",
      "description": "Three writing-oriented skills (academic-writing, research-paper-writing, scientific-paper) share substantial overlap in trigger conditions and purpose, creating ambiguity about which to invoke for any given writing task.",
      "fix": "Merge into a single writing skill with internal modes (draft, revise, review) OR clearly document the exact trigger boundaries and non-overlapping use cases for each.",
      "why": "When a researcher wants to write a paragraph, the agent must decide among three skills with fuzzy boundaries. This wastes tokens on skill loading and forces the user to mentally model three near-identical tools.",
      "ref": "Skills: academic-writing, research-paper-writing, scientific-paper"
    },
    {
      "severity": "significant",
      "category": "architecture",
      "description": "The /incorporar command bundles 9 phases (PDF → resumo → vault → canvas → claims) into a single monolithic workflow, forcing the researcher through every stage even when only a subset is needed.",
      "fix": "Decompose into standalone sub-commands (/incorporar-resumo, /incorporar-vault, /incorporar-claims) with the option to chain, or add a --phases flag to select which stages run.",
      "why": "A researcher wanting only to update the canvas after reading a paper must re-run the full pipeline, wasting time and encouraging workarounds.",
      "ref": "Command: /incorporar"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "No discoverability mechanism exists for the 14 skills and 9 MCPs beyond AGENTS.md prose — a researcher must read and memorize the full tool surface before becoming productive.",
      "fix": "Add a /help or /skills command that lists available tools with one-line descriptions and example invocations; consider a /which-tool command that suggests the right skill/MCP for a natural-language intent.",
      "why": "The AGENTS.md file requires linear reading of ~500 lines. A new contributor or the researcher returning after weeks away will struggle to recall which MCP handles PDF downloads versus metadata.",
      "ref": "AGENTS.md, system prompt skill list"
    },
    {
      "severity": "moderate",
      "category": "duplication",
      "description": "Five MCPs (scihub, crossref, semantic-scholar, scholar-sidekick, google-scholar) provide overlapping paper search and metadata capabilities, and the /consultar-academico command queries 4 bases in parallel — but there is no guidance on when to use individual MCPs versus the composite command.",
      "fix": "Document a decision tree: /consultar-academico for broad discovery → individual MCPs for targeted follow-up (e.g., scholar-sidekick for OA checks, crossref for citation graphs).",
      "why": "Without documented boundaries, the agent may invoke multiple similar MCPs redundantly, burning tokens and increasing latency.",
      "ref": "Commands: /consultar-academico; MCPs: scihub, crossref, semantic-scholar, scholar-sidekick, google-scholar"
    },
    {
      "severity": "moderate",
      "category": "architecture",
      "description": "The /claudiney command (advisor feedback with contextual analysis) is described as the gatekeeper for monograph modifications, but its validation rules, context window, and interaction with the council validation system are unspecified.",
      "fix": "Document the exact contract: what context /claudiney receives (full chapter? diff only?), how its output is structured, and whether it triggers council validation or substitutes for it.",
      "why": "A command that gatekeeps all monograph edits must have unambiguous semantics, or researchers will bypass it when frustrated.",
      "ref": "Command: /claudiney"
    },
    {
      "severity": "minor",
      "category": "integration",
      "description": "The 10-stage pipeline (PESQUISA → INCORPORAÇÃO → ... → CÓDIGO) implies linear progression, but real research is iterative — a finding during ANÁLISE may trigger new PESQUISA, yet the ecosystem provides no explicit loop-back support or state tracking.",
      "fix": "Document the pipeline as a cycle with entry/exit points; add a /roadmap-status command that shows which stage the current task is in and what artifacts exist at each stage.",
      "why": "The linear mental model will cause frustration when the researcher discovers a gap in the literature review while writing the discussion section.",
      "ref": "Fluxo: PESQUISA → CÓDIGO"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "Three vault-oriented skills (vault-semantic-schema, vault-tagger, vault-bases-maintainer) are maintained separately but all touch the same Obsidian vault structure — a researcher editing frontmatter must mentally parse which skill governs which property.",
      "fix": "Add a vault skill dispatch table in AGENTS.md: tag changes → vault-tagger, schema changes → vault-semantic-schema, view/filter changes → vault-bases-maintainer.",
      "why": "Without clear ownership boundaries, the agent may load the wrong skill or load all three sequentially, doubling interaction cost.",
      "ref": "Skills: vault-semantic-schema, vault-tagger, vault-bases-maintainer"
    }
  ],
  "recommendation": "Consolidate the three writing skills into one with mode dispatch; decompose /incorporar into scoped sub-commands; add a /which-tool or /help mechanism for discoverability; and publish a one-page decision matrix mapping research intents to commands, skills, and MCPs.",
  "schema_version": 3
}
```

---

## Análise de Ergonomia e Experiência do Pesquisador (Judge 2)

### 1. Carga Cognitiva

O ecossistema exige que o pesquisador mantenha um modelo mental de **14 skills + 9 MCPs + ~6 comandos + pipeline de 10 estágios + schema do vault**. Isso representa aproximadamente **30 entidades distintas** para navegar antes de produzir uma única frase da monografia.

O principal gargalo cognitivo está na **sobreposição semântica entre skills de escrita** (`academic-writing`, `research-paper-writing`, `scientific-paper`). As três compartilham triggers como "escrever texto acadêmico", "revisar paper", "redigir metodologia". O pesquisador não tem como saber qual invocar sem ler o conteúdo de cada SKILL.md — e o agente, ao encontrar um gatilho ambíguo, pode carregar a skill errada ou carregar múltiplas sequencialmente. Isso é um **anti-padrão de design de skills**: skills devem ter domínios mutuamente exclusivos ou dispatch interno por modo.

O comando `/incorporar` com suas **9 fases internas** empurra complexidade para o runtime: o pesquisador inicia uma ação aparentemente simples ("incorporar este PDF") e recebe 9 etapas de processamento, muitas das quais podem ser irrelevantes para seu objetivo imediato. Isso viola o princípio de **carga cognitiva progressiva** — comece simples, ofereça profundidade sob demanda.

### 2. Troca de Contexto (Context Switching)

O fluxo de trabalho força **saltos entre domínios técnicos distintos**:

| Estágio | Ferramenta | Domínio |
|---------|-----------|---------|
| PESQUISA | MCPs (5 fontes) | Metadados acadêmicos |
| INCORPORAÇÃO | `/incorporar` | Obsidian vault |
| ENRIQUECIMENTO | skills vault-* | Frontmatter, tags, bases |
| ESCRITA | skills writing-* | LaTeX, `monografia/` |
| CÓDIGO | skills go-* | Go, `src/` |
| ANÁLISE | skills stats-* | Python, `scripts/` |
| FIGURAS | skill scientific-figures | Go render, Python, TikZ |
| VALIDAÇÃO | council | Meta-avaliação |

Em um dia típico de pesquisa, o aluno alterna entre 4-5 desses domínios. Cada transição custa **15-30 segundos de reorientação mental** mais o tempo de carregamento de skills. Com 10-20 transições por sessão, isso consome **5-10 minutos improdutivos por dia**.

O problema é agravado pela ausência de **statefulness explícito**: o ecossistema não mantém registro de "você estava escrevendo a seção 3.2 do capítulo de metodologia, e antes disso estava analisando resultados do GA na instância 20a". O pesquisador precisa reconstruir esse contexto mentalmente a cada retomada.

### 3. Descobribilidade (Discoverability)

O único mecanismo de descoberta é o arquivo `AGENTS.md` (~500 linhas de prosa densa) e a lista de skills no system prompt. Não existe:

- **Comando `/help` ou `/skills`** que liste ferramentas com descrições de uma linha
- **Comando `/which-tool`** que receba uma intenção em linguagem natural e sugira a skill/MCP correta
- **Matriz de decisão** que mapeie intenções de pesquisa a ferramentas específicas
- **Exemplos canônicos** de cada comando/skill no próprio help

A taxonomia de skills também não ajuda: `scientific-paper` vs `research-paper-writing` vs `academic-writing` são nomes que descrevem o mesmo conceito com palavras diferentes. Um nome como `writing-draft`, `writing-revise`, `writing-review` seria imediatamente mais claro.

Para MCPs, o problema é similar: `scholar-sidekick`, `academic-search`, `google-scholar`, `semantic-scholar` — qual a diferença? A resposta está na documentação de cada um, mas o custo de descobri-la é proibitivo durante o fluxo de trabalho.

### 4. Clareza da Separação de Responsabilidades

**O que funciona bem:**

- A separação entre **commands** (ações iniciadas pelo usuário) e **skills** (conhecimento especializado injetado no agente) é conceitualmente correta
- MCPs são claramente a camada de **integração externa** (APIs acadêmicas, PDFs)
- O pipeline de 10 estágios fornece um **mapa mental de alto nível** útil para orientação
- `/claudiney` como gatekeeper de alterações na monografia é um bom guardrail

**O que não funciona:**

- Skills de escrita (3) e skills de vault (3) formam **clusters com fronteiras difusas** — é um cheiro de design que indica necessidade de consolidação ou dispatch interno
- O comando `/incorporar` é um **god-command**: faz tudo (download, resumo, vault, canvas, claims) e portanto não faz nada de forma modular
- A distinção entre **pesquisa** (MCPs + `/consultar-academico`) e **incorporação** (`/incorporar`) é clara, mas a transição entre elas não é automatizada — o pesquisador precisa manualmente passar resultados de busca para incorporação
- O plugin `pdf-watcher.ts` introduz um **paradigma de interação diferente** (event-driven, file-system watcher) que não se integra conceitualmente com os comandos explícitos

### 5. Síntese

O ecossistema é **funcionalmente completo** — cobre todas as fases necessárias da pesquisa. Mas impõe **alto atrito cognitivo** por três razões estruturais:

1. **Superfície de ferramentas superdimensionada**: 30+ entidades quando ~12 bem projetadas cobririam o mesmo espaço
2. **Ausência de camada de descoberta**: o pesquisador precisa saber o que existe antes de poder usar
3. **Granularidade inconsistente**: `/incorporar` é monolítico enquanto as skills de vault são hiper-especializadas

O veredito é **WARN** (não FAIL) porque o sistema funciona — um pesquisador determinado conseguirá produzir a monografia. Mas a experiência é desnecessariamente desgastante, e o acúmulo de pequenos atritos pode levar a workarounds que comprometem a rastreabilidade monografia-evidência que o schema do vault tenta garantir.

### Recomendações Priorizadas

1. **Consolidar skills de escrita** (1-2 horas): unificar em uma skill com dispatch interno por modo
2. **Decompor `/incorporar`** (2-3 horas): extrair fases como sub-comandos independentes
3. **Criar `/which-tool`** (1 hora): comando simples que mapeia intenção → ferramenta
4. **Publicar matriz de decisão** (30 min): tabela de referência rápida no AGENTS.md
5. **Documentar contrato do `/claudiney`** (30 min): especificar entradas, saídas, e interação com council
