---
title: Guia de Claims e Evidências da Monografia
type: claim-registry
tags:
- capitulo/experimentos
- evidencia/auditoria
- tipo/auditoria
- tipo/writing
- topico/monografia
writing_status: revisar
validation_status: validado
status: migrado-para-claims-individuais
created: 2026-06-02
updated: 2026-06-04
last_verified: 2026-06-04
---

# Guia de Claims e Evidências da Monografia

Os claims da monografia não são mais mantidos como uma tabela centralizada nesta nota. Cada claim agora vive em uma nota individual em `vault/claims/`, seguindo o template de claim e expondo suas propriedades no frontmatter. O sumário operacional fica em [[claims.base]].

Esta nota permanece como guia de schema, vocabulário e processo de validação.

> [!warning] Regra de fonte primária
> Não usar uma nota do `vault/` como evidência suficiente para claims metodológicos ou experimentais. Claims metodológicos apontam para `src/`; claims experimentais apontam para `src/data/results/`, scripts de análise, figuras ou tabelas geradas. O `vault/` organiza e rastreia, mas não substitui a fonte primária.

## Onde Consultar

| Necessidade | Local |
|---|---|
| Sumário navegável de todos os claims | [[claims.base]] |
| Notas individuais dos claims | `vault/claims/` |
| Claims conceituais | `vault/claims/C*.md` |
| Claims metodológicos de formulação | `vault/claims/M*.md` |
| Claims metodológicos de algoritmos | `vault/claims/A*.md` |
| Claims experimentais | `vault/claims/E*.md` |
| Claims bloqueados | `vault/claims/B*.md` |
| Claims interpretativos | `vault/claims/I*.md` |

## Schema Mínimo de Claim

Cada nota individual de claim deve conter, no mínimo:

| Campo | Uso |
|---|---|
| `type` | Sempre `claim` |
| `claim_id` | ID estável usado no texto e nas Bases. Ex.: `E19`, `M02`, `I06` |
| `claim` | Formulação exata que pode ou não entrar na monografia |
| `claim_type` | `conceitual`, `metodologico`, `experimental`, `interpretativo` ou `bloqueado` |
| `status` | `validado`, `requer-validacao`, `bloqueado` ou `substituido` |
| `strength` | `forte`, `moderado`, `fraco`, `forte descritivo` ou `bloqueado` |
| `primary_evidence` | Código, dados, script, figura/tabela ou literatura que sustenta o claim |
| `vault_support` | Notas do vault que organizam a evidência |
| `monografia_section` | Capítulo/seção recomendada |
| `usage_guidance` | Como usar ou não usar o claim no texto |
| `rationale` | Motivo do bloqueio, ressalva ou interpretação, quando aplicável |
| `action_required` | Ação necessária para claims bloqueados ou substituídos |
| `last_verified` | Data da última checagem contra a fonte primária |

## Vocabulário Controlado

| Dimensão | Valores recomendados |
|---|---|
| Tipo | `conceitual`, `metodologico`, `experimental`, `interpretativo`, `bloqueado` |
| Status | `validado`, `requer-validacao`, `bloqueado`, `substituido` |
| Força | `forte`, `moderado`, `fraco`, `forte descritivo`, `bloqueado` |
| Evidência primária | `literatura`, `src/`, `src/data/results/`, `scripts/`, `monografia/figs/` |

## Protocolo de Uso

1. Consultar [[claims.base]] para localizar o claim pelo ID, tipo, status ou seção.
2. Abrir a nota individual do claim em `vault/claims/`.
3. Verificar `primary_evidence` antes de usar o claim em texto acadêmico.
4. Se o claim estiver `bloqueado`, não usar a formulação original; seguir `action_required`.
5. Se o claim for `interpretativo`, escrever com escopo delimitado e sem causalidade indevida.
6. Ao adicionar ou alterar claims, editar a nota individual e confirmar se [[claims.base]] continua listando o registro.

## Observação Sobre Claims Experimentais

Os claims experimentais usam a cobertura auditada em [[auditoria-codigo-dados-vault]]: GA 1530, PSO 1530, ACO 1530, lowerbound 30 e brute-force 18. Para agrupamentos por instância, normalizar o campo `instance` pelo nome-base (`10a`, `30b`, etc.), pois os summaries misturam caminhos absolutos e relativos.

## Próximas Tarefas

1. Usar [[auditoria-codificacao-metodos]] para redigir limitações sobre comparabilidade entre representações.
2. Usar [[auditoria-hiperparametros]] para redigir limitações sobre parâmetros fixos.
3. Vincular trechos fortes da monografia aos IDs em `vault/claims/` ou a comentários `% claim: <ID>`.
4. Revisar [[claims.base]] antes do gate anti-alucinação final.
