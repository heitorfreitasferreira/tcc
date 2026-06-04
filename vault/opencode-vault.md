# Instruções para Agentes — Vault de Conhecimento do TCC

## Estrutura

```
vault/
├── papers/           ← Uma nota markdown por artigo referenciado (conhecimento externo)
├── areas/            ← Notas sobre áreas de pesquisa na literatura
├── projeto/          ← Notas sobre a implementação em Go (conhecimento interno)
│   ├── index.md      ←   visão geral
│   ├── problem-formulation.md ← formulação TSP-SD-ATP
│   ├── ga.md         ←   algoritmo genético
│   ├── pso.md        ←   particle swarm
│   ├── aco.md        ←   ant colony
│   ├── bruteforce.md ←   busca exaustiva
│   ├── experiment-pipeline.md ← pipeline experimental
│   └── architecture.md ← estrutura de pacotes Go
├── templates/        ← Templates para criar novas notas
├── bases/            ← Views estruturadas do Obsidian Bases
└── canvas/           ← Grafos de conhecimento visuais (.canvas)
```

## Project Skill
- `knowledge-base` — workflow completo em `.agents/skills/knowledge-base/SKILL.md`: busca → valida DOI → adiciona BibTeX → importa → enriquece nota → atualiza canvas

## Skills Carregadas

- `obsidian-markdown` — formatação markdown compatível com Obsidian
- `obsidian-vault` — operações no vault (criar, ler, buscar notas)
- `obsidian-bases` — queries estruturadas sobre o vault
- `json-canvas` — criar/atualizar grafos de conhecimento .canvas
- `academic-researcher` — buscar e entender papers acadêmicos
- `academic-search` — estratégias de busca acadêmica
- `vault-tagger` — padronização de tags hierárquicas para grafo e buscas
- `vault-semantic-schema` — manutenção de properties, templates, relações tipadas e claims
- `vault-bases-maintainer` — manutenção de `vault/bases/*.base`

## MCPs Disponíveis

- `scihub` — baixar PDFs e metadados de papers
- `google-scholar` — buscar artigos acadêmicos

## Separação entre Conhecimento Externo e Interno

- `papers/` e `areas/` = **conhecimento da literatura** (artigos, surveys, classificações teóricas)
- `projeto/` = **conhecimento da implementação** (código Go, parâmetros, decisões de projeto, pipeline)
- `areas/` deve referenciar `projeto/` quando relevante, mas não conter detalhes de implementação
- `projeto/` deve referenciar `areas/` e `papers/` quando um conceito ou artigo fundamenta uma decisão

## Como Criar uma Nota de Artigo

1. Use Google Scholar para buscar o paper
2. (Opcional) Use SciHub para obter metadados/PDF
3. Use `vault/templates/paper-note.md` como template
4. Preencha: YAML frontmatter (título, autores, ano, DOI, propriedades e tags)
5. Escreva resumo e contribuições
6. Link para notas de área existentes: `[[tsp]]`, `[[genetic-algorithms]]`
7. Adicione tags hierárquicas: `tipo/paper`, `area/tsp`, `metodo/ga`, `papel/comparativo`

## Convenções Semânticas

- **YAML frontmatter** sempre no topo com metadados
- Wiki links para conectar conceitos/papers
- **Properties** guardam fatos consultáveis: `year`, `rating`, `reading_status`, `pdf_status`, `areas`, `methods`, `chapters`
- **Tags** guardam navegação visual e filtros de grafo: `tipo/paper`, `area/tsp`, `metodo/aco`, `papel/comparativo`
- **Status de leitura**: `pendente` | `lido-parcial` | `lido`
- **Status de escrita**: `rascunho` | `revisar` | `pronto` | `bloqueado` | `desatualizado`
- **Status de validação**: `nao-validado` | `validado` | `requer-validacao` | `bloqueado`
- **Rating**: 1-5 (relevância para o TCC)
- Nome do arquivo = bibtex-key normalizada: `kennedy1995particle.md`

### Prefixos de Tags

- `tipo/`: tipo da nota (`tipo/paper`, `tipo/area`, `tipo/projeto`, `tipo/writing`, `tipo/index`)
- `status/`: estado visível no grafo (`status/pendente`, `status/lido`, `status/revisar`)
- `area/`: área temática (`area/tsp`, `area/drone-routing`, `area/lower-bound`)
- `metodo/`: método (`metodo/ga`, `metodo/pso`, `metodo/aco`, `metodo/bruteforce`)
- `papel/`: papel na monografia (`papel/fundacional`, `papel/comparativo`, `papel/metodologico`)
- `capitulo/`: capítulo alimentado pela nota (`capitulo/fundamentacao`, `capitulo/experimentos`)
- `evidencia/`: tipo de suporte (`evidencia/referencia`, `evidencia/codigo`, `evidencia/resultado`)

### Claims e Evidências

- `vault/writing/claim-evidence-matrix.md` é a fonte canônica dos claims.
- Todo claim forte deve ter `ID`, força, evidência primária e apoio no vault.
- Claims metodológicos apontam para `src/`; claims experimentais apontam para `src/data/results/` ou scripts de análise.
- Notas de paper podem usar `claim_support` para listar IDs da matriz, mas a matriz continua sendo a fonte de verdade.

## Workflow de Busca

1. Identificar lacuna: o que precisa ser fundamentado?
2. Buscar no Google Scholar com `academic-search`
3. Baixar/ler com `academic-researcher` + SciHub
4. Criar nota no vault seguindo template
5. Atualizar canvas de conhecimento se necessário
6. Atualizar `monografia/bib/abntex2-references.bib` se for nova referência

## Canvas

O arquivo `canvas/tcc-knowledge-graph.canvas` contém o grafo de conhecimento.
Agentes podem lê-lo e modificá-lo para refletir conexões entre papers e áreas.

## Bases

- `bases/papers.base` — visão tabular de papers, prioridade de leitura, PDF e uso na monografia.
- `bases/claims.base` — visão de auditoria para a matriz de claims e eventuais notas individuais de claim.

Use Bases para localizar pendências; não duplique evidências nelas. A evidência primária continua nos arquivos indicados pela matriz de claims.
