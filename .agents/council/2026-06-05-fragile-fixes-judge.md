## Judge — Frágil Fixes Validation

### Per-Paper Assessment

| Paper | Has Resumo | Substantial? | Status OK? | BibTeX matches? | Verdict |
|-------|-----------|-------------|-----------|----------------|---------|
| nagata2006eax | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ⚠️ (ver abaixo) | ⚠️ APROVADO COM RESSALVA |
| vanhove2012route | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ✅ | ✅ APROVADO |
| dellamico2021multiple | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ✅ | ✅ APROVADO |
| dellamico2022exact | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ✅ | ✅ APROVADO |
| deepaco2023 | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ✅ | ✅ APROVADO |
| demsar2006statistical | ✅ | ✅ (3 linhas, conteúdo específico) | ✅ resumo-lido | ✅ | ✅ APROVADO |

### Additional Findings

#### 1. nagata2006eax — Discrepância vault note × BibTeX (`RESOLVER`)
- **Vault note** descreve a versão de periódico: *"A Powerful Genetic Algorithm Using Edge Assembly Crossover for the Traveling Salesman Problem"*, INFORMS Journal on Computing, 2013, DOI `10.1287/ijoc.1120.0506`.
- **BibTeX** (`monografia/bib/abntex2-references.bib:460`) referencia a versão de conferência: *"Edge Assembly Crossover for the Traveling Salesman Problem"*, EvoCOP 2006, LNCS 3906, DOI `10.1007/11730095_13`.
- Ambas compartilham a mesma chave `nagata2006eax` e referem-se ao mesmo trabalho (a versão de periódico é uma extensão da de conferência), mas **título, ano, veículo e DOI divergem**. Se a monografia cita a chave `nagata2006eax`, o BibTeX produzirá a referência de 2006 (EvoCOP), mas o resumo no vault descreve a de 2013 (INFORMS).
- **Recomendação**: Ou alinhar o BibTeX com a versão de periódico (criando uma chave `nagata2013eax` ou atualizando a existente) ou ajustar o vault note para descrever a versão de conferência.

#### 2. demsar2006statistical — Resumo confere com PDF
- PDF de 30 páginas (`273KB`) presente e legível (`vault/papers/pdfs/demsar2006statistical.pdf`).
- Página 1 do PDF confirma: *"the Friedman test with the corresponding post-hoc tests"*, *"CD (critical difference) diagrams"*, *"Wilcoxon signed ranks test"*.
- Resumo do vault menciona explicitamente: Friedman, Nemenyi post-hoc, CD diagrams, Wilcoxon signed-ranks. ✅ Consistente.

#### 3. BibTeX keys — Todas as 6 chaves batem
Todas as 6 chaves existem no `monografia/bib/abntex2-references.bib` e correspondem aos campos `bibtex_key`/`bibtex-key` das notas do vault:

| Chave | BibTeX (linha) | Vault note |
|-------|---------------|------------|
| `nagata2006eax` | L460 (`@inproceedings`) | ✅ |
| `vanhove2012route` | L178 (`@article`) | ✅ |
| `dellamico2021multiple` | L145 (`@article`) | ✅ |
| `dellamico2022exact` | L156 (`@article`) | ✅ |
| `deepaco2023` | L472 (`@article`) | ✅ |
| `demsar2006statistical` | L36 (`@article`) | ✅ |

#### 4. citation-safety.md — Política atualizada, contagem de cabeçalho inconsistente (`CORRIGIR`)
- Frágeis: **0** ✅ (todos os 6 riscados com `~~` e promovidos a `resumo-lido`).
- A tabela "Seguras" lista **24 entradas** (17 originais + 6 promovidas + winter2002modeling que já estava), mas o cabeçalho diz `### Seguras (17)` — este número está desatualizado. Deveria ser `(24)` ou as promovidas deveriam estar em subseção separada.
- A linha de sumário no topo (`| **segura** (23) |`) também está defasada (deveria ser 24, ou 23 se winter2002modeling não for contada como lida). Recomendo recontar e atualizar.

### Veredito Final

| Critério | Status |
|----------|--------|
| Todos os 6 papers têm resumo? | ✅ Sim |
| Todos os resumos são substanciais? | ✅ Sim |
| Todos os status ≥ resumo-lido? | ✅ Sim |
| PDF do demsar existe e é legível? | ✅ Sim |
| Resumo do demsar menciona Friedman + Nemenyi + CD? | ✅ Sim |
| BibTeX keys batem vault ↔ .bib? | ✅ Sim (6/6) |
| Frágeis = 0 na política? | ✅ Sim |
| **Ação pendente** | Corrigir discrepância nagata2006eax (vault vs BibTeX) e contagem de cabeçalho em citation-safety.md |

**Conclusão**: As 6 classificações frágeis foram adequadamente resolvidas. Restam 2 correções cosméticas/menores (`nagata2006eax` versão divergente e contagem de cabeçalho). Nenhum bloqueio para a monografia.
