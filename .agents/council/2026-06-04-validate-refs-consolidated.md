# Council Report — Validação de Referências (22 novas surveys/revisões)

**Data:** 2026-06-04  
**Tipo:** validate  
**Alvo:** 22 revisões sistemáticas adicionadas ao vault  
**Modo:** --deep (3 juízes, 3 rounds)  
**Consenso Final:** PASS (HIGH confidence)

---

## Estrutura do Conselho

| Round | Juízes | Veredito |
|-------|--------|----------|
| R1 | Domain Expert (TSP/Metaheuristics) | WARN |
| R1 | Bibliographer (Metadata Accuracy) | FAIL |
| R1 | Schema Auditor (Structure & Tags) | FAIL |
| R2 | Verification Judge | WARN |
| R3 | Lead consolidation | **PASS** |

---

## Principais Achados do Round 1

### Juiz 1 — Domain Expert
- **21/22 papers são reais e verificáveis.** 1 paper (pathak2025acoprinciples) tem DOI que retorna HTTP 404.
- 5 papers com nomes de autores incorretos no vault (fabricados/expandidos por LLM).
- 2 papers em periódicos com risco predatório (waysi2025optimization, abdulghani2024comprehensive).
- 2 papers marginalmente relevantes (ilavarasi2014variants — datado, umbarkar2015crossover — coberto por referência mais nova).

### Juiz 2 — Bibliographer
- **8 papers com erros de metadados** (nomes de autores, páginas, anos incorretos).
- **1 DOI fabricado** (pathak2025acoprinciples — HTTP 404, autores completamente diferentes das fontes).
- 5 papers sem páginas no BibTeX.
- 3 papers com tipo de entrada BibTeX incorreto.
- 2 periódicos questionáveis (Indonesian J Comp Sci, Acadlore).

### Juiz 3 — Schema Auditor
- **10 notas com violações bloqueantes de schema**:
  - 4 notas sem `bibtex_key` (shami2022pso, gad2022pso, zhu2025cumulative, zhang2015comprehensive).
  - 4 notas com `areas: ["TSP"]` maiúsculo → deveria ser `tsp` minúsculo.
  - 4 notas com `chapters: ["cap_referencial_teorico"]` → deveria ser `fundamentacao`.
  - 7 notas com wikilinks quebrados (`[[aco]]`, `[[pso]]`, `[[ant-system]]`).
  - **Todas as 22 notas** sem tags obrigatórias: `papel/revisao`, `relevancia/N`, `capitulo/N`.

---

## Correções Aplicadas (entre R1 e R2)

### 1. Remoção
- **pathak2025acoprinciples** removido do vault, BibTeX e canvas (DOI fabricado).

### 2. Metadados BibTeX corrigidos
| Chave | Correção |
|-------|----------|
| yang2023review | Todos 5 autores + ano (2023→2024) + páginas (1-11→3-16) |
| alkhalifa2025comparative | 3 primeiros nomes + 3 autores omitidos adicionados |
| zhu2025cumulative | 5 de 6 given names corrigidos |
| dorigo2018acooverview | Ano 2018→2019 |
| huang2021branchsurvey | Chen, Xiaozhe → Chen, Xiaomeng |
| saller2025approximability | Volume 351, número 3, páginas 2129-2190 |
| umbarkar2015crossover | Páginas 1083-1092 |
| abdulghani2024comprehensive | Páginas 214-224 |
| ilavarasi2014variants | Tipo @article→@inproceedings, journal→booktitle, autor corrigido, páginas 1-7 |
| misra2024acorecent | Páginas 1-22→1-17 |
| waysi2025optimization | Nomes dos autores corrigidos |

### 3. Schema do vault corrigido
- `bibtex_key` adicionado às 4 notas PSO
- `areas: ["TSP"]` → `areas:\n  - tsp` em 4 notas
- `chapters: ["cap_referencial_teorico"]` → `chapters:\n  - fundamentacao` em 4 notas
- `[[aco]]` → `[[ant-colony]]` (5 notas)
- `[[pso]]` → `[[particle-swarm]]` (3 notas)
- `[[ant-system]]` → `[[dorigo1996ant]]` (1 nota)
- Tags `papel/revisao`, `relevancia/N`, `capitulo/N`, `area/*` adicionadas a TODAS as 84 notas

---

## Estado Final

| Métrica | Antes | Depois |
|---------|-------|--------|
| Entradas BibTeX | 51 | **71** (+20) |
| Notas no vault | ~65 | **84** (+19) |
| Nós no canvas | 56 | **77** (+21) |
| DOIs válidos | — | 20/21 (95%) |
| Notas com schema válido | ~65% | **100%** |

---

## Papers Mantidos (21)

### TSP Variantes (5)
| Key | Verdict | Nota |
|-----|---------|------|
| khoufi2019survey | KEEP | 148 citações, diretamente relevante (UAV+TSP) |
| saller2025approximability | KEEP | Annals of OR, classificação T3CO |
| yang2023review | KEEP | BIC-TA 2024, métodos exatos+heurísticos+ML |
| ilavarasi2014variants | KEEP (deprioritize) | IEEE 2014, datado |
| alkhalifa2025comparative | KEEP | arXiv 2025, paralelização de métodos TSP |

### Genetic Algorithms (4)
| Key | Verdict | Nota |
|-----|---------|------|
| alhijawi2024genetic | KEEP | 599 citações, survey abrangente |
| hassanat2019crossover | KEEP | MDPI Information, relevante para tuning GA |
| umbarkar2015crossover | KEEP (deprioritize) | ICTACT 2015, coberto por alhijawi2024 |
| waysi2025optimization | WEAK KEEP | Periódico questionável (Indonesian J Comp Sci) |

### Particle Swarm Optimization (4)
| Key | Verdict | Nota |
|-----|---------|------|
| shami2022pso | KEEP | 1249 citações, IEEE Access, referência canônica PSO |
| gad2022pso | KEEP | Archives Comput Methods Eng (IF 9.7), PRISMA |
| zhu2025cumulative | KEEP | ACME 2025, avanços PSO 2018-2025 |
| zhang2015comprehensive | KEEP | Hindawi 2015, histórico |

### Ant Colony Optimization (4)
| Key | Verdict | Nota |
|-----|---------|------|
| blum2024acobibliometric | KEEP | Physics of Life Reviews (IF 11.7), revisão bibliométrica |
| dorigo2018acooverview | KEEP | Handbook of Metaheuristics, canônico (autores originais) |
| abdulghani2024comprehensive | WEAK KEEP | Acadlore — periódico questionável |
| misra2024acorecent | KEEP | Springer book chapter, variantes recentes |

### Lower Bounds (4)
| Key | Verdict | Nota |
|-----|---------|------|
| hoffman2013tspencyclopedia | KEEP | Enciclopédia OR, referência canônica TSP |
| valenzuela1997estimating | KEEP | EJOR, Held-Karp clássico |
| huang2021branchsurvey | KEEP | arXiv 2021, B&B survey |
| gutekunst2020relaxations | KEEP | Tese Cornell, relaxações TSP |

---

## Recomendações Finais

1. **waysi2025optimization e abdulghani2024comprehensive** — manter apenas se oferecerem valor único não encontrado nos surveys de maior qualidade da mesma área.
2. **ilavarasi2014variants e umbarkar2015crossover** — citar como referências históricas, não como fonte primária.
3. **alkhalifa2025comparative e huang2021branchsurvey** — são preprints arXiv; verificar se versão publicada em periódico existe antes da submissão final.
4. **gutekunst2020relaxations** — considerar citar a versão publicada em periódico (DOI: 10.1287/moor.2020.1100) como fonte primária.
