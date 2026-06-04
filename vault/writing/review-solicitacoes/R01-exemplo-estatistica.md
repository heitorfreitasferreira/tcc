---
title: "Adicionar análise estatística formal (Friedman + Nemenyi)"
tags:
  - tipo/revisao
  - status/resolvido
  - exemplo
review_id: R01
status: resolvido
priority: alta
source: reuniao
date_opened: 2026-06-01
date_closed: 2026-06-03
target_chapter: experimentos
correction_layers:
  - dados
  - latex-macro
evidence_layer: dados
claim_ids:
  - C04
  - E01
  - E02
verified_by_script: scripts/check-reviews.sh --count --status resolvido
aliases:
  - R01
  - exemplo-estatistica
---

<!-- ITEM EXEMPLO — pode ser removido após entender o formato -->

## Solicitação Original

> Orientador: "Você está comparando 3 métodos estocásticos em 30 instâncias. Comparar apenas médias é insuficiente. Adicione testes estatísticos formais seguindo Demšar (2006)."

## Análise Técnica

### Problema Identificado

Capítulo de Experimentos não tinha análise estatística. Afirmações de superioridade entre métodos não eram sustentadas.

### Hierarquia de Informação — Onde Está a Verdade?

| Fonte | O que diz | Conflito? |
|---|---|---|
| `src/data/results/` | 4590 execuções estocásticas (3 métodos × 30 instâncias × 51 sementes) | OK |
| Literatura | Demšar (2006) recomenda Friedman + Nemenyi para múltiplos classificadores em datasets múltiplos | OK |
| `vault/` | [[analysis-methodology]] tem protocolo estatístico definido | Pendente de implementação |
| `monografia/` | Experimentos sem seção estatística | Ausente |

### Causa Raiz

Análise estatística não foi implementada nos scripts; capítulo foi escrito antes dos testes.

## Plano de Correção

### Abordagem Preferida (mais determinística)

- [x] **Camada 1 — Código/Dados:** Criar `scripts/analise-estatistica.py` que carrega summaries e executa Friedman, Nemenyi, Wilcoxon/Holm
- [x] **Camada 2 — Script determinístico:** Script gera `monografia/generated/tables/tab-estatistica.tex` e `monografia/figs/cd-diagram.{svg,png}`
- [ ] **Camada 3 — Edição agentica:** Desnecessário (camadas 1-2 cobrem)

### Comandos e Passos

```bash
python3 scripts/analise-estatistica.py
scripts/gerar-metricas-monografia.py
scripts/check-monografia.sh
```

### Artefatos Afetados

- `scripts/analise-estatistica.py` — criado
- `monografia/generated/tables/tab-estatistica.tex` — gerado
- `monografia/figs/cd-diagram.svg` — gerado
- `monografia/figs/cd-diagram.png` — gerado
- `monografia/cap_experimentos/experimentos.tex` — seção 4.3 adicionada

## Verificação

### Critério de Aceite

- Friedman rejeita H₀ com p < 0.05
- Nemenyi mostra pelo menos um par significativo
- Wilcoxon/Holm confirma o resultado
- Figura CD exportada
- Tabela com resultados incluída em experimentos.tex via `\input`

### Script de Validação

```bash
python3 scripts/analise-estatistica.py && echo "PASS"
```

## Evidências

- Friedman: F(2,58)=293.2222, p=4.710129e-31
- Ranks: ACO=1.1000, GA=1.9000, PSO=3.0000
- Nemenyi: CD=0.6050, todos os pares significativos
- Wilcoxon/Holm: todos os pares significativos

## Notas

Resolvido como tarefa P4 do roadmap principal. Ver [[entrevista-sessao#Pergunta 3 — Métricas de Avaliação]].
