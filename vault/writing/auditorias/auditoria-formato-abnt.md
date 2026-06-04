---
title: Auditoria de Formatação ABNT no Template LaTeX
tags:
  - writing
  - auditoria
  - abnt
  - latex
status: concluido-com-achados
created: 2026-06-04
related:
  - "[[roadmap-monografia]]"
  - "[[../capitulos/introducao]]"
  - "[[../capitulos/fundamentacao]]"
  - "[[../capitulos/proposta]]"
  - "[[../capitulos/experimentos]]"
  - "[[../capitulos/conclusao]]"
---

# Auditoria de Formatação ABNT — Template LaTeX

Verificação do template `ppgco.cls`, `main_ppgco_ufu.tex` e capítulos quanto à conformidade ABNT.

## Resultado Geral

O template está estruturalmente correto. A classe `ppgco.cls` herda de `abntex2`, que implementa as normas ABNT NBR 14724, NBR 10520, NBR 6023, NBR 6024 e NBR 6027. As funcionalidades ABNT existem, mas várias **não estão sendo usadas** pelos capítulos.

---

## Itens Verificados

### 1. Citações longas (>3 linhas) — recuo 4cm

- **Status:** OK (infraestrutura pronta, sem uso atual)
- O ambiente `citacao` do `abntex2` está disponível e implementa recuo de 4cm, fonte menor, sem aspas.
- Nenhum capítulo usa citação direta longa. Quando necessário, usar `\begin{citacao}...\end{citacao}`.

### 2. Alíneas

- **Status:** Infraestrutura OK, sem uso nos capítulos
- `abntex2` fornece `alineas`, `subalineas` e `incisos` (itens a), b), c)...).
- Capítulos usam apenas `itemize`. A lista de perguntas de pesquisa em `introducao.tex` e os objetivos poderiam usar `alineas` para conformidade mais estrita.

### 3. Remissões internas (`\ref`, `\pageref`)

- **Status:** Funcional, mas sem `\autoref`
- `ppgco.cls:528-542` define traduções pt-BR para `\autoref` (Figura, Tabela, Capítulo, Seção, etc.).
- Capítulos usam `\ref{}` com nome manual (ex: `Capítulo~\ref{fundamentacao}`). Funciona, mas `\autoref{}` seria mais robusto.
- `\pageref{}` nunca usado.

### 4. Ambiente de siglas

- **Status:** Pacote carregado e lista gerada, mas comandos nunca usados no texto
- Pacote `acronym` carregado em `main_ppgco_ufu.tex:44`.
- 16 siglas definidas em `abrev/Abreviaturas.tex` com `\acro{}`.
- `\listasiglas{abrev/Abreviaturas}` chamado no main.
- **Nenhum capítulo usa `\ac{}`, `\acs{}` ou `\acl{}`.** Siglas são escritas manualmente a cada ocorrência. Isso perde a funcionalidade "primeira ocorrência = nome por extenso, demais = sigla".

### 5. Espaçamento, margens e tipografia

- **Status:** OK (herdado de `abntex2`)
- Papel A4, 12pt, twoside, openright — conformes.
- `\parindent` e `\parskip` comentados em `ppgco.cls` — usa defaults do `abntex2`.

---

## Problemas Encontrados

### CRÍTICO: 4 chaves BibTeX ausentes

| Chave | Usada em | Linhas |
|---|---|---|
| `haroun2015performance` | `introducao.tex:15`, `fundamentacao.tex:95` | |
| `chandra2022comparative` | `introducao.tex:15`, `fundamentacao.tex:95` | |
| `clerc2000discretepso` | `fundamentacao.tex:69` | |
| `dorigo2004book` | `fundamentacao.tex:74` | |

Compilar com BibTeX produzirá `[?]` para essas referências.

### MODERADO: Recursos ABNT disponíveis mas não utilizados

| Recurso | Status |
|---|---|
| `\ac{}/acs{}/acl{}` (acronym) | Carregado, nunca usado |
| `\autoref{}` | Traduzido em ppgco.cls, nunca usado |
| `alineas` | Disponível, nunca usado |
| `citacao` (longa) | Disponível, nunca usado |

### MENOR: Limpeza e TODOs

- `main_ppgco_ufu.tex:75`: TODO da ficha catalográfica e folha de aprovação.
- `tabularx` usado em `proposta.tex` sem `\usepackage` explícito (herdado de `abntex2`).
- `subfig` carregado (main.tex:54) mas não usado.
- `nomencl`, `graphicx`, `hyperref` carregados duplamente (ppgco.cls + main.tex). Inofensivo.
- `makeidx` carregado (ppgco.cls:111) mas nenhum `\index{}` ou `\printindex` usado.

---

## Recomendações

1. **Adicionar as 4 chaves BibTeX ausentes** — bloqueia compilação limpa.
2. **Substituir siglas manuais por `\ac{ACO}`, `\ac{TSP}`, etc.** — consistência automática.
3. **Substituir `\ref{}` manuais por `\autoref{}`** — previne "Capítulo 4" quando deveria ser "Tabela 4".
4. **Substituir `itemize` por `alineas`** onde o conteúdo for enumeração formal dentro de seção.
5. **Resolver TODO da ficha catalográfica** antes da entrega final.
6. **Remover pacotes não usados** (`subfig`, `makeidx`) na limpeza final.
