---
title: Figuras e Tabelas da Monografia
tags:
- status/atualizado
- tipo/writing
- topico/figuras
- topico/monografia
- topico/tabelas
status: regenerado-pos-p4
created: 2026-06-02
updated: 2026-06-02
type: writing
---

# Figuras e Tabelas da Monografia

Catálogo operacional das figuras e tabelas candidatas para a monografia. A seleção abaixo está alinhada com [[auditoria-codigo-dados-vault]], [[claim-evidence-matrix]], [[glossario-monografia]] e [[auditoria-script-analise-estatistica]].

> [!warning] Regra de uso
> Não citar figura ou tabela no texto final sem conferir se o arquivo existe, se a escala/unidade aparece, e se o dado usado está coerente com a cobertura atual: 4638 summaries, 30 instâncias, 51 sementes para GA/PSO/ACO, 30 lowerbound e 18 brute-force (`10a..15c`).

## Inventário Atual de Artefatos

Arquivos existentes em `monografia/figs/` nesta auditoria:

| Tipo | Quantidade | Exemplos |
|---|---:|---|
| SVG | 22 | `heatmap-makespan-median.svg`, `scalability-runtime.svg`, `cd-diagram.svg` |
| PNG | 29 | previews de gráficos, route panels e logos |
| TEX | 6 | `flowchart-*.tex`, `fig-route-small-multiples.tex` |
| PDF | 6 | `flowchart-*.pdf`, logos |
| Total | 57 | inclui logos institucionais |

Artefatos principais existentes:

| Família | Arquivos existentes | Status |
|---|---|---|
| Diagramas conceituais | `diagram-angular-penalty.{png,svg}`, `diagram-tensor-3d.{png,svg}` | Prontos |
| Fluxogramas | `flowchart-ga`, `flowchart-pso`, `flowchart-aco`, `flowchart-bruteforce`, `flowchart-lowerbound` em `.tex/.pdf/.png/.svg` | Prontos |
| Heatmaps | `heatmap-makespan-median`, `heatmap-gap-vs-bf`, `heatmap-success-rate` em `.png/.svg` | Regenerados após BF `15a..15c` |
| Qualidade/tempo | `fig-quality-distribution`, `scalability-runtime`, `scatter-quality-vs-time`, `boxplot-estabilidade` em `.png/.svg` | Prontos |
| Convergência | `fig-runtime-convergence`, `convergence-last-improvement`, `convergence-overlay-{10a,30a,50a,100a}` | Prontos |
| Rotas | `route-panel-{10a,30a,100a}.png`, `fig-route-small-multiples.tex` | Regenerados |
| Estatística | `cd-diagram.{png,svg}` | Regenerado e validado após P4 |

## Seleção Recomendada Para o Corpo

Usar poucas figuras no corpo principal. Colocar o restante em apêndice ou omitir se não sustentar claim direto.

| ID | Capítulo | Claim suportado | Artefato | Status | Observação |
|---|---|---|---|---|---|
| F1 | Fundamentação/Proposta | O custo depende da sequência de três nós, não apenas da aresta atual. | `diagram-angular-penalty.svg` | Pronto | Usar com explicação da penalidade angular. |
| F2 | Proposta | O tensor 3D pré-computa custo `G[anterior][atual][próximo]`. | `diagram-tensor-3d.svg` | Pronto | Associar à avaliação O(n) após pré-computação O(n^3). |
| F3 | Proposta | Os métodos têm fluxos distintos mas compartilham a mesma função objetivo. | `flowchart-ga.tex`, `flowchart-pso.tex`, `flowchart-aco.tex`, `flowchart-lowerbound.tex` | Pronto | Preferir `.tex`/PDF em LaTeX; não usar todos se ficar repetitivo. |
| F4 | Experimentos | ACO apresenta menor makespan mediano que GA e PSO na maioria/todas as instâncias auditadas. | `heatmap-makespan-median.svg` | Regenerado | Usa cobertura atual. |
| F5 | Experimentos | Nas instâncias com ótimo, ACO tem menor gap e maior taxa de acerto que GA/PSO. | `heatmap-gap-vs-bf.svg` + `heatmap-success-rate.svg` | Regenerado | Reflete BF `10a..15c`. |
| F6 | Experimentos | ACO melhora qualidade, mas custa muito mais tempo de otimização. | `scatter-quality-vs-time.svg` | Pronto | Figura central de trade-off. |
| F7 | Experimentos | O tempo cresce por ordens de grandeza entre métodos e tamanhos. | `scalability-runtime.svg` | Pronto | Deve ter eixo/legenda em ms ou escala log explícita. |
| F8 | Experimentos | Métodos estocásticos têm variabilidade distinta entre sementes. | `boxplot-estabilidade.svg` ou `fig-quality-distribution.svg` | Pronto | Preferir uma das duas, não ambas, salvo se cada uma responder pergunta diferente. |
| F9 | Experimentos | A convergência difere entre métodos ao longo das avaliações/iterações. | `fig-runtime-convergence.svg` ou `convergence-overlay-*.svg` | Pronto | Preferir uma figura agregada; overlays por instância podem ir ao apêndice. |
| F10 | Experimentos | As rotas finais diferem qualitativamente entre métodos. | `route-panel-10a.png`, `route-panel-30a.png`, `route-panel-100a.png`, `fig-route-small-multiples.tex` | Regenerado | Painéis compostos em LaTeX. |
| F11 | Experimentos | Diferenças estatísticas entre métodos. | `cd-diagram.svg` | Regenerado e validado | Usar com os valores P4 atuais. |

## Figuras Por Capítulo

### Capítulo 1 — Introdução

| ID | Tipo | Descrição | Arquivo | Status |
|---|---|---|---|---|
| I1 | ilustração conceitual | Cenário de patrulha com drone e POIs | — | Opcional; não criar se for decorativa |

Decisão: a Introdução pode funcionar sem figura. Se uma ilustração for criada, ela deve explicar a missão de patrulha e não virar imagem decorativa de drone.

### Capítulo 2 — Fundamentação Teórica

| ID | Tipo | Descrição | Arquivo | Status |
|---|---|---|---|---|
| FT1 | diagrama | Penalidade angular entre segmentos consecutivos | `diagram-angular-penalty.{png,svg}` | Pronto |
| FT2 | diagrama | Tensor 3D de custo | `diagram-tensor-3d.{png,svg}` | Pronto |

Decisão: usar no máximo um desses no Capítulo 2 se o Capítulo 3 já for usar ambos. Evitar duplicação.

### Capítulo 3 — Proposta

| ID | Tipo | Descrição | Arquivo | Status |
|---|---|---|---|---|
| P1 | diagrama | Tensor 3D e função objetivo | `diagram-tensor-3d.svg` | Pronto |
| P2 | fluxograma | Fluxo do GA | `flowchart-ga.tex` | Pronto |
| P3 | fluxograma | Fluxo do PSO | `flowchart-pso.tex` | Pronto |
| P4 | fluxograma | Fluxo do ACO | `flowchart-aco.tex` | Pronto |
| P5 | fluxograma | Lower bound AP/Hungarian | `flowchart-lowerbound.tex` | Pronto |
| P6 | fluxograma | Busca exaustiva | `flowchart-bruteforce.tex` | Pronto |

Decisão: se o capítulo ficar visualmente pesado, agrupar fluxogramas de GA/PSO/ACO em apêndice e manter apenas o tensor + lower bound no corpo.

### Capítulo 4 — Experimentos e Resultados

| ID | Tipo | Descrição | Arquivo | Status |
|---|---|---|---|---|
| ER1 | heatmap | Makespan mediano por instância e método | `heatmap-makespan-median.{png,svg}` | Regenerado |
| ER2 | heatmap | Gap mediano ou agregado vs brute-force | `heatmap-gap-vs-bf.{png,svg}` | Regenerado com BF `10a..15c` |
| ER3 | heatmap | Taxa de acerto do ótimo | `heatmap-success-rate.{png,svg}` | Regenerado com BF `10a..15c` |
| ER4 | distribuição | Estabilidade/variabilidade entre sementes | `boxplot-estabilidade.{png,svg}` | Pronto |
| ER5 | distribuição | Qualidade agregada por método | `fig-quality-distribution.{png,svg}` | Pronto |
| ER6 | escala temporal | Tempo de otimização por tamanho | `scalability-runtime.{png,svg}` | Pronto |
| ER7 | dispersão | Trade-off qualidade-tempo | `scatter-quality-vs-time.{png,svg}` | Pronto |
| ER8 | convergência | Convergência agregada por avaliações | `fig-runtime-convergence.{png,svg}` | Pronto |
| ER9 | convergência | Última iteração com melhoria | `convergence-last-improvement.{png,svg}` | Pronto |
| ER10 | rotas | Painéis representativos por método | `route-panel-{10a,30a,100a}.png`, `fig-route-small-multiples.tex` | Regenerado |
| ER11 | estatística | Diagrama de diferença crítica | `cd-diagram.{png,svg}` | Regenerado e validado |

### Capítulo 5 — Conclusão

Nenhuma figura nova. Referenciar achados do Capítulo 4 sem introduzir resultado novo.

## Tabelas Finais Recomendadas

| ID | Capítulo | Claim suportado | Conteúdo | Fonte/Geração | Status |
|---|---|---|---|---|---|
| T1 | Proposta/Experimentos | O estudo usa 30 instâncias sintéticas com tamanhos definidos. | Instância, n, variante, arquivo `.points/.graph` | `src/data/` ou script simples | A criar/gerar |
| T2 | Proposta/Experimentos | Parâmetros foram fixos e reproduzíveis. | GA, PSO, ACO, BF, LB; flags e defaults | `src/cmd/*.go`, [[glossario-monografia]] | Manual auditada |
| T3 | Experimentos | BF fornece ótimos para 18 instâncias pequenas. | Ótimos brute-force `10a..15c` | `src/data/results/summary/*__bruteforce__*.json` | A gerar |
| T4 | Experimentos | AP bound é válido mas frouxo. | BF ótimo, LB, gap LB->BF nas 18 instâncias | [[auditoria-codigo-dados-vault]] e summaries | A gerar |
| T5 | Experimentos | ACO tem menor gap médio que GA/PSO nas instâncias com BF. | Runs, gap médio, min, max, taxa de ótimo por método | [[auditoria-codigo-dados-vault]] | A gerar |
| T6 | Experimentos | ACO domina qualidade em instâncias grandes. | Best, média, dp por método em `50a..100c` | summaries | A gerar |
| T7 | Experimentos | ACO é muito mais lento em n=100. | Tempo médio ms e razão ACO/GA, ACO/PSO | timings ligados por `timing_file` | A gerar |
| T8 | Experimentos | Teste estatístico global e pós-teste. | Friedman/Iman-Davenport, Nemenyi, Wilcoxon/Holm | `scripts/analise-estatistica.py` | Validado |

> [!note] `consolidate_results.py`
> Executado após P4 com `4638` summaries. As estatísticas foram validadas operacionalmente, mas o script ainda imprime tabelas no stdout; para uso final em LaTeX, selecionar/formatar apenas as tabelas necessárias.

## Pendências Bloqueantes

| Pendência | Afeta | Ação |
|---|---|---|
| Seleção final de tabelas ainda não foi incorporada aos `.tex` | T1–T8 | Copiar/adaptar tabelas após escrita do Capítulo 4 |
| `consolidate_results.py` imprime tabelas no stdout | T1–T7 | Transformar em arquivo `.tex` dedicado se for usado diretamente |
| Figura opcional da Introdução não existe | I1 | Só criar se apoiar argumento; caso contrário, omitir |

## Critérios de Aceite Para Figura/Tabela Final

Cada artefato aprovado para a monografia deve atender aos critérios abaixo:

| Critério | Como verificar |
|---|---|
| Claim explícito | A legenda ou seção deve dizer que comparação/achado a figura sustenta |
| Fonte rastreável | Deve apontar para `src/data/results/`, `scripts/` ou render endpoint Go |
| Escala e unidades | Eixos, ticks, unidades (`makespan`, `%`, `ms`, `n`) e legenda devem aparecer |
| Variabilidade | Gráficos de métodos estocásticos devem mostrar distribuição, dp, boxplot ou n=51 quando necessário |
| Cobertura correta | BF deve ser `10a..15c`; métodos estocásticos devem usar 51 sementes por instância |
| Texto em pt-BR | Títulos/legendas/captions finais devem estar em português acadêmico |
| Composição LaTeX | Painéis e explicações longas devem ser compostos em `.tex`, não embutidos em PNG denso |
| Estatística corrigida | Qualquer figura/tabela de significância depende do script estatístico corrigido |

## Ordem Recomendada de Finalização

1. Escolher as figuras do corpo: F1/F2, F6/F7/F8, F10 e F11.
2. Transformar T1–T8 em tabelas LaTeX finais durante a escrita do Capítulo 4.
3. Deixar painéis extras e small multiples para apêndice, se necessário.
