---
title: "Triagem de PDFs Faltantes — P17"
tags:
  - writing
  - auditoria
  - bibliografia
  - pdf
status: concluido
created: 2026-06-04
---

# Triagem de PDFs Faltantes — P17

Executada em 2026-06-04. Objetivo: tentar baixar referências sem PDF íntegro e classificar editorialmente as não recuperáveis.

## Metodologia

1. Sci-Hub por DOI (todas as referências com DOI)
2. Google Scholar para busca de versões abertas
3. `scripts/download-pdfs.sh --keys "..." --use-scihub` (Unpaywall + Semantic Scholar + Sci-Hub)
4. DeepACO: download direto do NeurIPS proceedings (open access)
5. **Tesble.com** (proxy Sci-Hub): encontrou 5 PDFs adicionais

## Resultado Global

| Categoria | Quantidade | Ação |
|---|---|---|
| **Recuperado** | 6 | PDF baixado e salvo |
| **Removido** | 1 | lawler1985traveling — livro real (ISBN 0471904139) com DOI incorreto (resenha). Substituído por applegate2006traveling. Removido do `.bib`, `.tex` e do escopo. |
| **Solicitar download manual** | 0 | — nenhuma referência pendente |
| **Análise posterior** | 7 | Relevantes para trabalhos futuros, mas não indispensáveis para a monografia atual |
| **Descartado do escopo imediato** | 5 | Não sustentam argumento central; podem ser omitidas sem enfraquecer o texto |

## Recuperado

| Referência | Origem | PDF | Páginas |
|---|---|---|---|
| [[deepaco2023]] | NeurIPS 2023 open access | `vault/papers/pdfs/deepaco2023.pdf` | 29 |
| [[vanhove2012route]] | Tesble / Sci-Hub | `vault/papers/pdfs/vanhove2012route.pdf` | 8 |
| [[aggarwal2000angular]] | Tesble / Sci-Hub | `vault/papers/pdfs/aggarwal2000angular.pdf` | 15 |
| [[heldkarp1971traveling]] | Tesble / Sci-Hub | `vault/papers/pdfs/heldkarp1971traveling.pdf` | 20 |
| [[dellamico2021multiple]] | Tesble / Sci-Hub | `vault/papers/pdfs/dellamico2021multiple.pdf` | 25 |
| [[dellamico2022exact]] | Tesble / Sci-Hub | `vault/papers/pdfs/dellamico2022exact.pdf` | 34 |

> deepaco2023, dellamico2021multiple, dellamico2022exact estavam corrompidos anteriormente; todos recuperados.

## Solicitar Download Manual (prioritário)

| Referência | Função | Alternativa validada |
|---|---|---|
| [[lawler1985traveling]] | TSP clássico (livro editado, Wiley 1985) | [[garey1979computers]], [[applegate2006traveling]] |

> Livro de 1985 — improvável em Sci-Hub ou CAPES. As alternativas validadas cobrem o mesmo conteúdo para a monografia.

## Análise Posterior

Referências úteis para trabalhos futuros, expansão do estado da arte ou versões estendidas do texto. Não bloqueiam a monografia atual.

### Condicionados a lower bounds (não bloqueiam se o texto usar apenas bound AP)

| Referência | Função | Motivo |
|---|---|---|
| [[balas1985branch]] | Branch and bound survey | Sem DOI; não essencial para bound AP |
| [[fischetti1992additive]] | Additive bounding ATSP | Springer paywall; não disponível no Tesble |
| [[valenzuela1997estimating]] | Estimating HK lower bound | Elsevier paywall; não disponível no Tesble |
| [[leraromero2020dynamic]] | ng-path labeling TDTSP | Springer paywall; não disponível no Tesble |

### Trabalhos relacionados / futuros

| Referência | Função | Motivo |
|---|---|---|
| [[ahmed2024receding]] | Receding horizon UAV | Elsevier paywall; não disponível no Tesble |
| [[nagata2006eax]] | EAX crossover para TSP | Sem DOI; PDF corrompido |
| [[araujo2025pso]] | PSO + TSP inteligente | PDF corrompido (Syntax Error); movido para corrupted/ |

## Descartado do Escopo Imediato

Referências que não sustentam argumento central da monografia. Mantidas no catálogo apenas como registro bibliográfico.

| Referência | Função | Motivo do descarte |
|---|---|---|
| [[hga2024hybrid]] | GA+ACO híbrido | Fora do escopo (métodos híbridos, não canônicos) |
| [[sun2024hybrid]] | PSO híbrido ternary optical | Muito específico; não é PSO canônico |
| [[huang2025matrix]] | PSO matricial multi-TSP | Multi-TSP, não TSP-SD-ATP |
| [[kappagantula2025dpso]] | DPSO-Q RL-enhanced | Fora do escopo (RL + PSO, não PSO canônico) |
| [[toaza2023review]] | Review de metaheurísticas TSP | Review genérica; [[chandra2022comparative]] e [[alexander2020comparison]] cobrem melhor |

## PDFs Corrompidos Residuais

Arquivos em `vault/papers/pdfs/corrupted/` que não puderam ser recuperados:

| Arquivo | Referência | Status |
|---|---|---|
| `araujo2025pso.pdf` | [[araujo2025pso]] | análise posterior |
| `nagata2006eax.pdf` | [[nagata2006eax]] | análise posterior |
| `vanhove2012.pdf` | [[vanhove2012route]] | obsoleto (substituído por `vanhove2012route.pdf` íntegro) |

> dellamico2021multiple e dellamico2022exact estavam em corrupted/; substituídos por versões íntegras via Tesble.

> Os arquivos corrompidos foram mantidos no diretório `corrupted/` como registro. Não devem ser usados como fonte.

## Impacto na Monografia

- **Recuperados:** 6 PDFs íntegros (deepaco2023, vanhove2012route, aggarwal2000angular, heldkarp1971traveling, dellamico2021multiple, dellamico2022exact)
- **Solicitar download manual:** 1 referência (lawler1985traveling — livro 1985 com alternativas validadas)
- **Análise posterior:** 7 referências (sem PDF acessível, para trabalhos futuros)
- **Descartados do escopo imediato:** 5 referências (não sustentam argumento central)
- **Nenhuma referência sem PDF bloqueia a escrita da monografia**
