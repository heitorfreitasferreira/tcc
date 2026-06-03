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
4. Preencha: YAML frontmatter (título, autores, ano, DOI, tags)
5. Escreva resumo e contribuições
6. Link para notas de área existentes: `[[tsp]]`, `[[genetic-algorithms]]`
7. Adicione tags: `tsp`, `ga`, `pso`, `aco`, `drone`, `metaheuristic`

## Convenções

- **YAML frontmatter** sempre no topo com metadados
- Wiki links para conectar conceitos/papers
- **Tags** em kebab-case: `genetic-algorithms`, `particle-swarm`
- **Status**: `pendente` | `lido-parcial` | `lido`
- **Rating**: 1-5 (relevância para o TCC)
- Nome do arquivo = bibtex-key normalizada: `kennedy1995particle.md`

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
