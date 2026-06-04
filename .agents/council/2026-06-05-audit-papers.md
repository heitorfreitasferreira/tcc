# Council Report — Auditoria de Resumos

**Date:** 2026-06-05
**Items:** Auditoria completa dos 87 papers do vault
**Judges:** 2 independent

## Consolidated Verdict: WARN — Classificação correta, risco bibliográfico significativo

| Critério | Judge 1 | Judge 2 | Consolidado |
|----------|---------|---------|-------------|
| OK papers spot-check | PASS 7/10 (3 status minor) | — | PASS |
| Sem-resumo verification | PASS 5/5 | — | PASS |
| Fixes (starzec, muthanna) | PASS | — | PASS |
| Incompleto verification | — | PASS 10/10 | PASS |
| PDF existence check | — | PASS (9 corrupted, 1 absent) | PASS |
| Citation risk assessment | — | **11 frágil citations** | FAIL |
| Roadmap consistency | — | **Contradição com safety policy** | FAIL |

## Issues

### 1. Status inconsistency (Judge 1: minor)
3 papers with `status: pendente` despite complete notes:
- `blum2024acobibliometric.md`
- `gad2022pso.md`
- `hassanat2019crossover.md`

**Ação:** mudar status para `resumo-lido`.

### 2. 11 frágil citations (Judge 2: critical)
Papers citados na monografia sem PDF legível ou sem resumo:

**Alto risco** (6 incompleto + 1 sem-resumo, sustentam claims metodológicos):
| Paper | Citado em | Claim | Problema |
|-------|-----------|-------|----------|
| bean1994genetic | fund.tex:69 | random keys para PSO | PDF tem só metadados INFORMS |
| clerc2000discretepso | fund.tex:69 | PSO discreto formalizado | PDF sem texto extraível |
| potvin1996ga | fund.tex:60 | revisão de crossover para TSP | PDF é scan (imagens) |
| larranaga1999ga | fund.tex:60 | sistematizou representações GA | PDF com xref corrompido |
| wu2020comparative | fund.tex:95 | comparou GA/PSO/ACO em TSPLIB | PDF corrompido — **marcado como "segura" no roadmap** |
| halim2019combinatorial | fund.tex:95 | escolha da representação | PDF muito corrompido — **marcado como "segura" no roadmap** |
| nagata2006eax | fund.tex:60 | EAX, referência de qualidade | Sem resumo, sem PDF legível |

**Médio/Baixo risco** (4 papers, claims contextuais ou futuros):
vanhove2012route, dellamico2021multiple, dellamico2022exact, deepaco2023

### 3. Roadmap contradiz própria safety policy (Judge 2: critical)
`wu2020comparative`, `halim2019combinatorial`, `clerc2000discretepso` são listados como "Referências centrais seguras" mas têm PDFs corrompidos. O P22 (citation-safety.md) classificaria como `frágil`.

## Recommended Actions (por prioridade)

| Prioridade | Ação |
|------------|------|
| 🔴 Imediata | Corrigir 3 status inconsistencies |
| 🔴 Imediata | Executar P22 (citation-safety.md) para formalizar classificação |
| 🟡 Alta | Recuperar/repor PDFs dos 6 "alto risco" |
| 🟡 Alta | Adicionar `% citation-safety: frágil` nos .tex para os 11 |
| 🟢 Média | Preencher resumos dos 18 "sem resumo" (já parcialmente cobertos por P26/P28/P31) |
