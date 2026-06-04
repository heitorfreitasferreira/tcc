---
title: "Política de Segurança de Citações"
tags:
  - writing
  - monografia
  - citacoes
  - validacao
created: 2026-06-05
updated: 2026-06-05
---

## Classificação

Cada referência citada na monografia recebe uma de quatro classes:

| Classe | Critério | Exigência para citação |
|--------|----------|----------------------|
| **segura** (23) | PDF legível + nota com resumo substancial + status lido/resumo-lido | Pode sustentar qualquer claim |
| **condicional** (10) | PDF ilegível OU nota incompleta, mas conteúdo verificável por metadados/abstract | Pode sustentar claims contextuais/de revisão, não claims fortes |
| **frágil** (0) | PDF ausente/corrompido E nota sem resumo OU nota/vault ausente | Não pode sustentar claims; uso restrito a menções contextuais |
| **bloqueada** (0) | Chave não resolve no BibTeX | Não pode ser citada |

## Tabela de Classificação (35 chaves)

### Seguras (24)

| Key | Status | PDF |
|-----|--------|-----|
| agatz2018optimization | lido | Legível |
| applegate2006traveling | lido-parcial | Legível |
| chandra2022comparative | lido | Legível |
| dorigo1996ant | lido | Legível |
| dorigo1997ant | lido | Legível |
| dorigo2004book | lido-parcial | Legível |
| freitas2020vns | lido | Legível |
| garey1979computers | lido-parcial | Legível |
| goldberg1989genetic | lido-parcial | Legível |
| gpaco2025 | lido | Legível |
| haroun2015performance | lido-parcial | Legível |
| holland1975adaptation | lido-parcial | Legível |
| kennedy1995particle | lido | Legível |
| lin1973effective | lido | Legível |
| murray2015flying | lido | Legível |
| neufaco2025 | lido | Legível |
| rajan2022routing | lido | Legível |
| winter2002modeling | lido | Legível |
| deepaco2023 | resumo-lido | Legível |
| dellamico2021multiple | resumo-lido | Legível |
| dellamico2022exact | resumo-lido | Legível |
| demsar2006statistical | resumo-lido | Legível |
| nagata2006eax | resumo-lido | Legível |
| vanhove2012route | resumo-lido | Legível |

### Condicionais (10)

| Key | Status | Problema |
|-----|--------|----------|
| bean1994genetic | lido-parcial | PDF corpo ilegível (só metadados INFORMS) |
| clerc2000discretepso | lido-parcial | PDF sem texto extraível |
| geng2025mdaco | resumo-lido | Sem PDF, resumo de metadados |
| halim2019combinatorial | lido-parcial | PDF corrompido |
| larranaga1999ga | lido-parcial | PDF com xref corrompido |
| potvin1996ga | lido-parcial | PDF scan (imagens, sem texto) |
| starzec2026motsp | resumo-lido | Dataset, sem PDF |
| wang2013aco | resumo-lido | Sem PDF |
| wang2015aco | resumo-lido | Sem PDF |
| wu2020comparative | lido-parcial | PDF danificado (xref errors) |

### Frágeis (0 — todos resolvidos em 2026-06-05)

| Key | Status | Problema | Onde é citado |
|-----|--------|----------|---------------|
| ~~deepaco2023~~ | → resumo-lido | ~~Sem resumo~~ | Conclusão (futuro) |
| ~~dellamico2021multiple~~ | → resumo-lido | ~~Sem resumo~~ | Fund. + Intro |
| ~~dellamico2022exact~~ | → resumo-lido | ~~Sem resumo~~ | Fund. + Intro |
| ~~demsar2006statistical~~ | → resumo-lido | ~~Sem vault note, sem PDF~~ | Fund. + Exp. |
| ~~nagata2006eax~~ | → resumo-lido | ~~Sem resumo~~ | Fund. |
| ~~vanhove2012route~~ | → resumo-lido | ~~Sem resumo~~ | Fund. |

### Bloqueadas (0)

Nenhuma. Todas as 34 chaves acadêmicas resolvem no BibTeX.

---

## Análise de Risco por Seção

### Fundamentação (maior densidade de citações)

| Linha | Citações | Maior risco |
|-------|----------|-------------|
| 9 | applegate2006traveling, garey1979computers | ✅ seguro |
| 11 | lin1973effective | ✅ seguro |
| 18 | winter2002modeling, vanhove2012route | ✅ seguro |
| 48 | murray2015flying, agatz2018optimization, dellamico2021multiple, dellamico2022exact, freitas2020vns | ✅ seguro |
| 50 | rajan2022routing | ✅ seguro |
| 60 | holland1975adaptation, goldberg1989genetic, potvin1996ga, larranaga1999ga, nagata2006eax | ⚠️ 2 condicionais (potvin, larranaga) |
| 67 | kennedy1995particle | ✅ seguro |
| 69 | clerc2000discretepso, bean1994genetic | ⚠️ 2 condicionais |
| 74 | dorigo1996ant, dorigo1997ant, dorigo2004book | ✅ seguro |
| 78 | dorigo1996ant, dorigo1997ant, wang2013aco, wang2015aco, geng2025mdaco, starzec2026motsp | ⚠️ 4 condicionais |
| 95 | wu2020comparative, haroun2015performance, chandra2022comparative, halim2019combinatorial | ⚠️ 2 condicionais (wu, halim) |
| 97 | demsar2006statistical | ✅ seguro |

### Introdução

| Linha | Citações | Maior risco |
|-------|----------|-------------|
| 5 | applegate2006traveling, garey1979computers | ✅ seguro |
| 11 | murray2015flying, agatz2018optimization, dellamico2021multiple, dellamico2022exact, rajan2022routing | ⚠️ 2 frágeis |
| 15 | wu2020comparative, haroun2015performance, chandra2022comparative, halim2019combinatorial | ⚠️ 1 condicional + 1 condicional |

### Conclusão

| Linha | Citações | Maior risco |
|-------|----------|-------------|
| 28 | deepaco2023, neufaco2025, gpaco2025 | ⚠️ 1 frágil (futuro) |

---

## Plano de Ação (ordenado por criticidade)

### 🔴 Imediato — frágil em seções fundamentais

| # | Key | Ação | Bloqueia |
|---|-----|------|----------|
| 1 | demsar2006statistical | Criar vault note + obter PDF. Referência metodológica central. | Protocolo estatístico inteiro |
| 2 | nagata2006eax | Preencher vault note (PDF já existe e é legível) | Claim sobre EAX |
| 3 | vanhove2012route | Preencher vault note (PDF já existe e é legível) | Contexto turn costs |
| 4 | dellamico2021multiple | Preencher vault note (PDF já existe e é legível) | Contexto drone routing |
| 5 | dellamico2022exact | Preencher vault note (PDF já existe e é legível) | Contexto drone routing |

### 🟡 Alta — condicional que sustenta claims fortes

| # | Key | Ação |
|---|-----|------|
| 6 | bean1994genetic | Obter PDF íntegro (P29) |
| 7 | clerc2000discretepso | Obter PDF legível |
| 8 | potvin1996ga | Obter PDF legível |
| 9 | larranaga1999ga | Obter PDF legível |
| 10 | wu2020comparative | Obter PDF legível |
| 11 | halim2019combinatorial | Obter PDF legível |

### 🟢 Média — sem PDF mas baixo impacto

| # | Key | Ação |
|---|-----|------|
| 12 | wang2013aco | Obter PDF |
| 13 | wang2015aco | Obter PDF |
| 14 | geng2025mdaco | Obter PDF |
| 15 | starzec2026motsp | Obter dataset |

### ⚪ Baixa — frágil só na conclusão

| # | Key | Ação |
|---|-----|------|
| 16 | deepaco2023 | Preencher vault note (PDF existe) |

---

## Gate de Segurança

Antes da entrega final da monografia:

- [ ] Nenhuma citação `frágil` deve permanecer nas seções de Fundamentação, Introdução ou Experimentos
- [ ] Citações `condicionais` só podem sustentar claims contextuais/de revisão
- [ ] `demsar2006statistical` deve ser promovida a `segura` (referência metodológica central)
- [ ] Todos os PDFs `condicionais` devem ser substituídos ou as claims correspondentes atenuadas
- [ ] `P22` deve ser reexecutado após cada ciclo de correção de PDFs

## Conexões

- [[roadmap-monografia]] — P22, P26, P28, P29, P30, P31
- [[auditoria-codigo-dados-vault]] — auditoria de consistência código-dados-vault
- [[claim-evidence-matrix]] — matriz de claims × evidências
- [[feromonio-3d-originalidade]] — verificação de originalidade do feromônio 3D
