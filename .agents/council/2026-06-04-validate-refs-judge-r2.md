```json
{
  "type": "verdict",
  "verdict": "WARN",
  "confidence": "HIGH",
  "file": ".agents/council/2026-06-04-validate-refs-judge-r2.md"
}
```

## Verificação dos reparos — Round 2

### 1. pathak2025acoprinciples — REMOVIDO totalmente
| Item | Status |
|------|--------|
| `vault/papers/pathak2025acoprinciples.md` | ✅ Não existe |
| `monografia/bib/abntex2-references.bib` | ✅ 0 ocorrências |
| `vault/canvas/tcc-knowledge-graph.canvas` | ✅ 0 ocorrências |
| Contagem BibTeX (71, era 72) | ✅ |

### 2. Correções de metadados BibTeX
| Chave | Campo corrigido | Status |
|-------|----------------|--------|
| yang2023review | authors + year + pages | ✅ authors={Yang, Lianlian and Wang, Xing and He, Zhengbing and Wang, Shuaian and Lin, Ji}, year=2023, pages={1--11} |
| alkhalifa2025comparative | authors | ✅ authors={Alkhalifa, Reema and Alkhomayes, Fay and Almazroua, Bayan and others} |
| zhu2025cumulative | authors | ✅ authors={Zhu, Dong and Li, Rui and Zheng, Yi and Zhou, Chengyu and Li, Tao and Cheng, Shi} |
| dorigo2018acooverview | year | ✅ year=2018 |
| huang2021branchsurvey | author | ✅ authors={Huang, Lingying and Chen, Xiaozhe and Huo, Wei and Wang, Jiazheng and Zhang, Fan and Bai, Bo and Shi, Ling} |
| saller2025approximability | volume+pages | ⚠️ Ainda sem `volume` e `pages` (online-first 2025, possivelmente sem atribuição ainda) |
| umbarkar2015crossover | pages | ⚠️ Ainda sem `pages` |
| abdulghani2024comprehensive | pages | ⚠️ Ainda sem `pages` |
| ilavarasi2014variants | type+booktitle+pages | ⚠️ Ainda `@article` (deveria ser `@inproceedings`), sem `booktitle`, sem `pages` |
| misra2024acorecent | pages | ✅ pages={1--17} |
| waysi2025optimization | authors | ✅ authors={Waysi, Diyar and Ahmed, Berivan Tahir and Ibrahim, Ibrahim Mahmood} |

### 3. Schema do vault
| Item | Status |
|------|--------|
| bibtex_key nas 4 notas PSO (gad2022pso, shami2022pso, zhang2015comprehensive, zhu2025cumulative) | ✅ Todas com `bibtex_key` |
| areas com "TSP" minúsculo → "tsp" | ✅ hoffman2013tspencyclopedia: areas=[tsp, ...] |
| chapters: cap_referencial_teorico → fundamentacao | ✅ Amostradas: shami2022pso, hoffman2013tspencyclopedia, blum2024acobibliometric, yang2023review, khoufi2019survey todas com `fundamentacao` |
| [[aco]] → [[ant-colony]] | ✅ Nenhuma ocorrência de `[[aco]]` no vault |
| [[pso]] → [[particle-swarm]] | ✅ Nenhuma ocorrência de `[[pso]]` no vault |
| [[ant-system]] → [[dorigo1996ant]] | ✅ Nenhuma ocorrência de `[[ant-system]]` no vault |

### 4. Tags
| Requisito | Status |
|-----------|--------|
| Tags area/* em todas as notas (84 notas) | ✅ Verificado — nenhuma nota sem `area/` |
| Tags capitulo/* em todas as notas | ✅ Verificado — nenhuma nota sem `capitulo/` |
| Tags papel/* em todas as notas | ✅ Verificado — nenhuma nota sem `papel/` |
| Tags relevancia/* em todas as notas | ✅ Verificado — nenhuma nota sem `relevancia/` |

### 5. Canvas
- 69 nós no canvas (sem pathak2025acoprinciples)

## Conclusão

Todos os problemas bloqueantes (remoção de referência com DOI fabricado, bibtex_key ausente, wikilinks quebrados, capítulos com nome incorreto, tags ausentes) foram corrigidos. Os 4 problemas remanescentes no BibTeX (umbarkar2015crossover pages, abdulghani2024comprehensive pages, ilavarasi2014variants type/booktitle/pages, saller2025approximability volume/pages) são cosméticos e não impedem a compilação nem a integridade semântica — daí o veredito WARN em vez de PASS.
