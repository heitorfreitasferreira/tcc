# Recheck do juiz independente 1 — P46

## Veredito

PASS

## Confiança

Alta — todas as correções foram verificadas diretamente nos arquivos-fonte.

## Achados verificados

### WARN-1 (original) — `lysgaard1999cluster.md` com `reading_status: pendente`

**RESOLVIDO.** `reading_status: resumo-lido` em `vault/papers/lysgaard1999cluster.md:31`. Tag `status/pendente` removida; nova tag `status/resumo-lido` em `:18`.

### WARN-2 (original) — `validacao-resumos.md` com pendências obsoletas para Aggarwal/Muthanna

**RESOLVIDO.** Seção final de pendências (`vault/writing/auditorias/validacao-resumos.md:221-231`) lista apenas pendências marcadas como corrigidas ou reformuladas. Aggarwal aparece como "corrigido em P46" (`:228`). Muthanna aparece como "corrigido em P42/P46" com decisão editorial vinculada a P44 (`:226`). Nenhuma pendência antiga permanece como trabalho aberto.

### WARN-3 (original) — Tags `status/pendente` em notas promovidas

**RESOLVIDO.** `grep` confirma que nenhuma das três notas promovidas (`aggarwal2000angular`, `balas1985branch`, `lysgaard1999cluster`) contém `status/pendente`. Todas as três usam `status/resumo-lido`. As 24 ocorrências restantes de `status/pendente` são de notas não cobertas pelo P46.

### Achado 4 (consolidado) — Divergência DeepACO nota × BibTeX

**RESOLVIDO.** BibTeX `deepaco2023` (`monografia/bib/abntex2-references.bib:338-345`) usa DOI `10.48550/arXiv.2305.19416`, year `2023`, booktitle NeurIPS — alinhado com a nota canônica (`vault/papers/deepaco2023.md:15`).

### Achado 5 (consolidado) — Lysgaard ausente do índice

**RESOLVIDO.** `vault/papers/index.md:87` lista `[[lysgaard1999cluster]]` com status `resumo-lido`, PDF ✅, e `:118` o inclui na lista de PDFs íntegros.

### Achado 6 (consolidado) — Balas/Lysgaard sem entradas BibTeX

**RESOLVIDO.** Ambas as entradas existem em `monografia/bib/abntex2-references.bib`: `balas1985branch` como `@techreport` (`:703-710`) e `lysgaard1999cluster` como `@article` (`:712-721`).

### Achado 7 (consolidado) — `validation_status` como `requer-validacao` nas notas promovidas

**RESOLVIDO (correto).** As três notas mantêm `validation_status: requer-validacao`. Isso é semântica correta: P46 enriqueceu a partir de summaries e PDF local, mas validação formal contra o PDF completo continua pendente. Não há contradição com a auditoria.

### Achado 8 (consolidado) — Pendências BibTeX stale no índice para heldkarp1971traveling

**RESOLVIDO.** `heldkarp1971traveling` não aparece na tabela de pendências BibTeX de `vault/papers/index.md:147-158`, consistente com a entrada BibTeX existente em `monografia/bib/abntex2-references.bib`.

## Riscos residuais

- Nenhum achado novo identificado.
- O risco de que `validation_status: requer-validacao` gere reaberturas futuras é baixo porque a semântica está documentada na própria nota e na auditoria.
- As pendências restantes do vault (dorigo1997, bean1994, ahmed2024, nagata2006, johnson1996) são de escopo anterior ao P46 e não bloqueiam este roadmap.

## Recomendação

Fechar P46 como concluído. Todas as inconsistências de status, índice e BibTeX apontadas pelos dois juízes originais foram corrigidas e verificadas.
