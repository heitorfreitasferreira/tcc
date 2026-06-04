---
title: Roadmap de Revisão — Solicitações do Orientador
tags:
  - writing
  - monografia
  - revisao
  - orientador
  - roadmap
status: ativo
created: 2026-06-03
updated: 2026-06-03
aliases:
  - revisao
  - solicitacoes-revisao
---

# Roadmap de Revisão — Solicitações do Orientador

Este documento define a infraestrutura para documentar, classificar e resolver cada solicitação de revisão ou consideração feita pelo orientador sobre a monografia. Toda mudança ou melhoria solicitada deve ser rastreada em um **item de revisão** (`Rxx`) e ter um **plano de correção determinístico** antes da execução.

## Ciclo de Vida de um Item de Revisão

```text
ABERTO → ANALISANDO → PLANEJADO → EM-ANDAMENTO → RESOLVIDO
                          ↘                ↗
                       CANCELADO
```

| Estado | Significado | Quem move |
|---|---|---|
| `aberto` | Orientador solicitou; pendente de análise | Humano |
| `analisando` | Investigando o problema contra o código, dados, literatura e vault | Humano ou agente |
| `planejado` | Correção definida com método e camada de atuação; pronto para executar | Agente |
| `em-andamento` | Correção sendo implementada | Agente |
| `resolvido` | Correção aplicada e verificada | Agente |
| `cancelado` | Solicitacão descartada ou postergada (com justificativa) | Humano |

## Hierarchy de Correção (Ordem de Preferência)

Determinar a camada de atuação mais determinística disponível antes de qualquer ação agentica. Inspirado na [[roadmap-monografia#Hierarquia de Informação]].

| Ordem | Camada | Abordagem | Exemplo |
|---:|---|---|---|
| 1 | `codigo` | Alterar `src/` e regenerar `monografia/generated/` | Valor experimental incorreto → corrigir pipeline de métricas e rodar `gerar-metricas-monografia.py` |
| 2 | `dados` | Scripts Python/bash que processam `src/data/results/` e geram `.tex` | Tabela faltando instância → script que varre resultados e emite `\tabular` |
| 3 | `pdf` | Extração textual de PDFs via `pdfinfo`, `pdftotext`, `pdfplumber` | Dado de paper mal interpretado → extrair valor do PDF com script |
| 4 | `latex-macro` | Macros determinísticas em `.tex` (rg, awk, sed, Python) | Termo inconsistente em todo o texto → `rg 'termo_errado' monografia/*.tex` + `sed` |
| 5 | `vault` | Aguardar e atualizar notas do vault com info validada | Nota de fundamentação incompleta → preencher com dados de código ou literatura |
| 6 | `agentico` | Edição por IA com prompt estruturado e validação pós-correção | Reestruturação de argumento, reescrita de parágrafo, reorganização de seção |

> [!warning] Regra de ouro
> Nunca pular para a camada 6 sem antes verificar se as camadas 1-5 resolvem o problema. Um item que exige camada 6 deve ter o prompt explícito e um critério de aceite no corpo da nota.

## Schema de um Item de Revisão

Cada solicitação é uma nota individual em `review-solicitacoes/Rxx-descricao-curta.md`, com o seguinte YAML:

| Campo | Obrigatório | Descrição |
|---|---|---|
| `review_id` | Sim | `R01`, `R02`, ... |
| `status` | Sim | `aberto` / `analisando` / `planejado` / `em-andamento` / `resolvido` / `cancelado` |
| `priority` | Sim | `alta` / `media` / `baixa` |
| `source` | Sim | `reuniao` / `email` / `comentario-tex` / `anotacao-manual` / `banca` |
| `date_opened` | Sim | Data ISO |
| `date_closed` | Não | Data ISO |
| `target_chapter` | Não | Capítulo(s) afetado(s): `fundamentacao`, `proposta`, `experimentos`, `introducao`, `conclusao`, `todos` |
| `correction_layers` | Sim | Lista das camadas envolvidas na correção: `[codigo, dados, latex-macro]` |
| `evidence_layer` | Sim | Camada onde está a evidência da correção: `codigo` / `dados` / `literatura` / `vault` / `monografia` |
| `claim_ids` | Não | Lista de claims da [[claim-evidence-matrix]] afetadas: `[E06, E13]` |
| `verified_by_script` | Não | Caminho do script que valida a correção |

O corpo da nota segue a estrutura definida no template [[review-request]].

## Diretório e Nomenclatura

```
vault/writing/review-solicitacoes/
├── index.md                    ← Lista consolidada com links para todos os Rxx
├── R01-exemplo-descritivo.md
├── R02-outra-solicitacao.md
└── ...
```

Arquivos de resolução e evidência podem ser depositados em `vault/writing/review-solicitacoes/evidencias/` para organizar screenshots, logs, ou outputs de scripts.

## Scripts de Apoio

### `scripts/check-reviews.sh`

Valida o estado das revisões:

| Check | O que verifica |
|---|---|
| YAML completo | Todo Rxx tem os campos obrigatórios |
| IDs únicos | Nenhum ID duplicado entre notas |
| Links válidos | `claim_ids` referenciam IDs existentes na [[claim-evidence-matrix]] |
| Estado de resolvidos | Rxx `resolvido` tem `date_closed` preenchido |
| Hierarquia respeitada | Rxx com `correction_layers` contendo `agentico` justifica por que camadas inferiores não bastam |

### `scripts/aplicar-correcao-review.sh`

Wrapper para aplicar correções determinísticas em lote:

```bash
scripts/aplicar-correcao-review.sh R03    # Executa o plano do R03
scripts/aplicar-correcao-review.sh --list # Lista todos os Rxx planejados
scripts/aplicar-correcao-review.sh --check R03 # Verifica se R03 foi resolvido
```

Cada item `planejado` pode definir no corpo da nota um bloco `## Plano de Correção` contendo comandos shell, scripts Python ou instruções para execução. O script lê esse bloco e executa os comandos.

## Conexão com o Gate Anti-Alucinação

Itens de revisão podem afetar diretamente o resultado do [[roadmap-monografia#Gate Formal Anti-Alucinação]]:

- Rxx com `status != resolvido` e `priority == alta` devem gerar **WARN** no `scripts/check-monografia.sh`
- Rxx que alteram claims, citações, números ou figuras devem atualizar os artefatos em `monografia/generated/`
- Rxx `cancelado` devem ficar registrados com justificativa para auditoria

## View Obsidian

Uma [[reviews.base]] em `vault/bases/` permite visualizar os itens por status, prioridade, capítulo ou camada de correção, usando a estrutura de Bases do Obsidian.

---

**Ver também:** [[roadmap-monografia]], [[claim-evidence-matrix]], [[review-request]], [[check-reviews]]
