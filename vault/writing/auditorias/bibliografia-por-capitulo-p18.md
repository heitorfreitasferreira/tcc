---
type: writing
writing_status: concluido
validation_status: validado
chapters:
  - introducao
  - fundamentacao
  - proposta
  - experimentos
  - conclusao
primary_evidence:
  - monografia/bib/abntex2-references.bib
  - vault/papers/
  - monografia/cap_fundamentacao/fundamentacao.tex
  - monografia/cap_introducao/introducao.tex
  - monografia/cap_experimentos/experimentos.tex
  - monografia/cap_conclusao/conclusao.tex
tags:
  - tipo/auditoria
  - tipo/writing
  - evidencia/auditoria
  - status/concluido
  - topico/monografia
  - topico/citacoes
---

# Auditoria: Bibliografia Mínima por Capítulo (P18)

## Resumo

Auditoria completa das referências citadas em cada capítulo da monografia, verificando disponibilidade de PDF, status de leitura e consistência com o vault.

## Métrica Geral

| Indicador | Valor |
|---|---|
| Entradas na bibliografia | 75 |
| Entradas citadas | 41 (54.7%) |
| Entradas não utilizadas | 34 (45.3%) |
| Papers no vault | 61 |
| Papers com PDF | 50 no disco |
| Citações totais no texto | 59 ocorrências |
| Capítulos com texto | 5 |

## Por Capítulo

### Introdução — 9 citações ✅

9 citações de 9 papers. Suficiente para apresentar contexto. Todos com PDF no disco.

### Fundamentação — 36 citações ⚠️

36 keys únicas, 38 ocorrências. Quantidade boa, mas **10 referências condicionais**:

| Chave | Rating | PDF | Leitura | Risco |
|---|---|---|---|---|
| `shami2022pso` | 5 | ausente | pendente | **ALTO** — claim substantivo (taxonomia PSO) |
| `gad2022pso` | 4 | ausente | resumo-lido | **ALTO** — claim substantivo (PRISMA, lacuna) |
| `oncan2009comparative` | 4 | ausente | pendente | **ALTO** — nota vault criada na auditoria |
| `potvin1996ga` | 4 | ilegível | lido-parcial | **MÉDIO** — claim factual (operadores crossover) |
| `clerc2000discretepso` | 5 | sem-texto | lido-parcial | **MÉDIO** — claim factual (PSO discreto) |
| `bean1994genetic` | 4 | capa-apenas | resumo-lido | **MÉDIO** — claim factual (random keys) |
| `wang2013aco` | 2 | ausente | resumo-lido | **BAIXO** — claim de suporte (feromônio 3D) |
| `wang2015aco` | 3 | ausente | resumo-lido | **BAIXO** — claim de suporte |
| `geng2025mdaco` | 2 | ausente | resumo-lido | **BAIXO** — claim de suporte |
| `starzec2026motsp` | 1 | ausente | resumo-lido | **BAIXO** — claim de suporte |

### Proposta — 0 citações ❌

Nenhuma referência bibliográfica. Capítulo puramente metodológico. Revisar se precisa de citações para a formulação do rTSP ou algoritmo.

### Experimentos — 3 citações ⚠️

Apenas `eryoldas2022survey`, `wang2021ant`, `demsar2006statistical`. Precisa de mais referências para metodologia experimental (benchmarks, configuração de parâmetros, análise estatística).

### Conclusão — 3 citações ⚠️

Apenas `deepaco2023`, `neufaco2025`, `gpaco2025` (trabalhos futuros). Precisa de mais referências para discussão e contextualização dos resultados.

## Referências Não Utilizadas — 34

### Com potencial de citação (rating ≥ 4, lido/PDF ok):

| Chave | Rating | Tema | Capítulo sugerido |
|---|---|---|---|
| `pop2024comprehensive` | 5 | Survey GTSP | introdução |
| `rajwar2023exhaustive` | 4 | Metaheurísticas (survey) | fundamentação |
| `balas1985branch` | 5 | Branch-and-bound TSP | fundamentação |
| `lysgaard1999cluster` | 4 | ATSP branching | fundamentação |
| `valenzuela1997estimating` | 4 | Held-Karp lower bound | fundamentação |
| `bock2025survey` | 3 | TSP warehousing | introdução |
| `ahmed2024receding` | 3 | UAV path planning | introdução |

### Pendentes (rating ≥ 4, sem PDF):

| Chave | Rating | Tema |
|---|---|---|
| `khoufi2019survey` | 5 | UAV TSP/VRP survey |
| `hoffman2013tspencyclopedia` | 5 | TSP encyclopedia |
| `dorigo2018acooverview` | 5 | ACO overview |
| `blum2024acobibliometric` | 5 | ACO bibliometric |
| `zhu2025cumulative` | 5 | PSO advances 2018-2025 |
| `saller2025approximability` | 5 | TSP approximability |
| `alkhalifa2025comparative` | 4 | Comparative review TSP |

### Descartados (`descartado-escopo-imediato`):

`toaza2023review`, `kappagantula2025dpso`, `sun2024hybrid`, `huang2025matrix`, `hga2024hybrid`

## Recomendações

1. **Prioridade máxima**: Obter PDF de `shami2022pso` (rating 5, usado para claim substantivo na fundamentação)
2. **Prioridade alta**: Obter PDF de `gad2022pso` (rating 4, claim substantivo PRISMA)
3. **Prioridade média**: Tentar obter PDFs de `wang2013aco`, `wang2015aco`, `geng2025mdaco`, `starzec2026motsp` para validar claim de feromônio multidimensional
4. **Proposta**: Avaliar se a formulação rTSP precisa de citação formal
5. **Experimentos**: Adicionar referências sobre metodologia experimental (e.g., benchmark TSPLIB, configuração de parâmetros)
6. **Conclusão**: Adicionar referências que contextualizem a discussão dos resultados

## Ações Realizadas

- [x] Auditoria completa de 75 entradas bib × 5 capítulos
- [x] Criação de vault note para `oncan2009comparative` (citado mas sem nota)
- [x] Identificação de 10 referências condicionais na fundamentação
- [x] Identificação de 3 capítulos com citações insuficientes
- [x] Categorização das 34 referências não utilizadas

## Data da Auditoria

2026-06-05
