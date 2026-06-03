---
title: Mapa Capítulos-TeX da Monografia
tags:
  - writing
  - monografia
  - latex
  - capitulos
status: auditado-com-bloqueios
created: 2026-06-02
---

# Mapa Capítulos-TeX da Monografia

Esta nota mapeia a estrutura real de `monografia/` para os capítulos planejados no roadmap. Também registra divergências entre os arquivos `.tex` atuais e as evidências auditadas em [[auditoria-codigo-dados-vault]], [[claim-evidence-matrix]], [[glossario-monografia]] e [[figuras-tabelas-monografia]].

> [!danger] Bloqueio de escrita
> Os arquivos `.tex` atuais contêm resultados experimentais desatualizados: `5029` execuções, `55` sementes, `GA` como melhor método e `ACO` como pior. Isso contradiz os dados auditados: `4638` summaries, `51` sementes para GA/PSO/ACO, `18` brute-force, e ACO com melhor qualidade descritiva nas instâncias avaliadas. Não reaproveitar esses parágrafos sem reescrita.

## Arquivo Principal

| Elemento | Arquivo | Papel | Status |
|---|---|---|---|
| Documento principal | `monografia/main_ppgco_ufu.tex` | Classe, pacotes, capa, resumo, ordem dos capítulos, bibliografia e apêndices | Precisa revisão |
| Siglas | `monografia/abrev/Abreviaturas.tex` | Lista de abreviaturas | Precisa correção de `rTSP` e talvez inclusão de `LB` |
| Bibliografia | `monografia/bib/abntex2-references.bib` | Referências ABNT/BibTeX | Falta 11 entradas do vault, ver [[papers/index#Pendências-BibTeX]] |
| Figuras | `monografia/figs/` | Artefatos visuais | Ver [[figuras-tabelas-monografia]] |

## Ordem de Inclusão no LaTeX

`main_ppgco_ufu.tex` inclui os capítulos nesta ordem:

| Ordem | Capítulo | Include | Arquivo real |
|---:|---|---|---|
| 1 | Introdução | `\include{cap_introducao/introducao}` | `monografia/cap_introducao/introducao.tex` |
| 2 | Fundamentação Teórica | `\include{cap_fundamentacao/fundamentacao}` | `monografia/cap_fundamentacao/fundamentacao.tex` |
| 3 | Proposta | `\include{cap_proposta/proposta}` | `monografia/cap_proposta/proposta.tex` |
| 4 | Experimentos | `\include{cap_experimentos/experimentos}` | `monografia/cap_experimentos/experimentos.tex` |
| 5 | Conclusão | `\include{cap_conclusao/conclusao}` | `monografia/cap_conclusao/conclusao.tex` |

## Mapeamento Por Capítulo

### Pré-textuais e Resumo

| Campo | Arquivo | Conteúdo atual | Ação necessária |
|---|---|---|---|
| Título | `main_ppgco_ufu.tex:22`, `main_ppgco_ufu.tex:65` | Coerente com patrulha com drones e metaheurísticas bioinspiradas | Manter, salvo ajuste fino de estilo |
| Resumo | `main_ppgco_ufu.tex:161-163` | Afirma `5029` execuções e GA como melhor método | Reescrever por último, após Experimentos corrigido |
| Abstract | `main_ppgco_ufu.tex:165-167` | Repete as mesmas afirmações erradas do resumo | Reescrever após resumo em português |
| Siglas | `abrev/Abreviaturas.tex` | Corrigido para incluir TSP-SD-ATP, AP, CD e LB | Revisar apenas se novas siglas forem usadas |

### Capítulo 1 — Introdução

| Item | Estado atual | Ação para escrita final |
|---|---|---|
| Arquivo | `monografia/cap_introducao/introducao.tex` | Usar como rascunho parcial, não como texto final |
| Ponto forte | Estrutura já contém contexto, motivação, objetivos, hipótese e contribuições | Preservar macroestrutura |
| Divergência | Usa VANTs de forma predominante, enquanto o glossário prefere `drone` | Padronizar para `drone` |
| Divergência | Hipótese ainda não reflete resultado auditado: ACO melhor qualidade, GA melhor tempo | Reescrever objetivos/perguntas após Capítulo 4 estabilizado |
| Entradas do vault | [[introducao]], [[problem-formulation]], [[drone-routing]], [[comparative-studies]], [[claim-evidence-matrix]] | Conferir antes de reescrever |
| Figuras | Nenhuma obrigatória | Omitir figura decorativa; criar só se apoiar cenário de patrulha |

### Capítulo 2 — Fundamentação Teórica

| Item | Estado atual | Ação para escrita final |
|---|---|---|
| Arquivo | `monografia/cap_fundamentacao/fundamentacao.tex` | Rascunho aproveitável com ajustes |
| Ponto forte | TSP, custo dependente, drones, GA, PSO, ACO e comparativos já aparecem | Revisar citações e conexão com proposta |
| Divergência | Busca exaustiva diz “até 15 pontos”, o que é aceitável, mas deve ser amarrado a `10a..15c` | Ajustar para “instâncias executadas `10a..15c`” |
| Divergência | Pode estar faltando lower bound AP como método/relaxação na fundamentação | Incluir seção curta sobre relaxação AP, Hungarian e limites inferiores |
| Entradas do vault | [[fundamentacao]], [[TSP]], [[tsp-variants]], [[bio-inspired-optimization]], [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]], [[lower-bounds]] | Revalidar citações |
| Figuras | `diagram-angular-penalty.svg`, talvez `diagram-tensor-3d.svg` | Usar com parcimônia |

### Capítulo 3 — Proposta

| Item | Estado atual | Ação para escrita final |
|---|---|---|
| Arquivo | `monografia/cap_proposta/proposta.tex` | Precisa reescrita substancial |
| Divergência | ACO usa `rho=0,5` no texto, mas código usa `rho=0.2` | Corrigir |
| Divergência | Diz 55 sementes e 5029 summaries | Corrigir para 51 sementes e cobertura auditada |
| Divergência | Não descreve `tcc optimize lowerbound` como método implementado no pipeline | Incluir |
| Divergência | Pode simplificar demais o primeiro movimento/retorno no tensor; conferir com `src/graph/makespan.go` e `src/graph/math.go` | Reescrever com base no código |
| Entradas do vault | [[proposta]], [[problem-formulation]], [[architecture]], [[ga]], [[pso]], [[aco]], [[bruteforce]], [[lower-bounds]], [[experiment-pipeline]], [[glossario-monografia]] | Usar como base validada |
| Figuras | `diagram-tensor-3d.svg`, `flowchart-*.tex` | Integrar via LaTeX, não como PNG denso |

### Capítulo 4 — Experimentos e Resultados

| Item | Estado atual | Ação para escrita final |
|---|---|---|
| Arquivo | `monografia/cap_experimentos/experimentos.tex` | Precisa reescrita quase integral |
| Divergência crítica | Afirma `5029` execuções, `1650` por método, `79` BF | Corrigir para `4638` summaries, `1530` por método, `30` LB, `18` BF |
| Divergência crítica | Afirma GA melhor e ACO pior | Corrigir: ACO melhor qualidade, GA melhor tempo, PSO pior qualidade neste desenho |
| Divergência crítica | Tabelas embutidas têm números incompatíveis com P7 | Substituir por tabelas geradas/validadas |
| Divergência | Usa ACO `rho=0,5`; código usa `0.2` | Corrigir |
| Estatística | Seção estatística ainda não está válida | Aguardar script corrigido pelo outro agente |
| Entradas do vault | [[experimentos]], [[resultados]], [[analysis-methodology]], [[auditoria-codigo-dados-vault]], [[figuras-tabelas-monografia]], [[auditoria-script-analise-estatistica]] | Usar apenas claims liberados |
| Figuras | Heatmaps, scatter, runtime, boxplot, convergência, overlays de rota | Regerar/validar antes de citar |

### Capítulo 5 — Conclusão

| Item | Estado atual | Ação para escrita final |
|---|---|---|
| Arquivo | `monografia/cap_conclusao/conclusao.tex` | Precisa reescrita após Capítulo 4 |
| Divergência crítica | Repete `5029` execuções e GA como melhor método | Corrigir após Experimentos |
| Ponto forte | Estrutura com contribuições, limitações e trabalhos futuros é adequada | Preservar macroestrutura |
| Entradas do vault | [[conclusao]], [[resultados]], [[analysis-methodology]], [[claim-evidence-matrix]], [[comparative-studies]] | Usar depois de resultados estabilizados |
| Figuras | Nenhuma nova | Não introduzir resultado novo |

## Arquivos de Apoio e Apêndices

| Arquivo | Estado | Decisão |
|---|---|---|
| `monografia/ape_comandos/abntex2-modelo-include-comandos.tex` | Template/comandos de exemplo | Não incluir no texto final |
| `monografia/ape_sobre/sobre.tex` | Material de modelo/sobre | Verificar se ainda é necessário; provavelmente não usar |
| `monografia/figs/fig-route-small-multiples.tex` | Figura/fragmento de rotas | Candidato a apêndice, não ao corpo principal |
| `monografia/figs/flowchart-*.tex` | Diagramas TikZ | Candidatos ao Capítulo 3 ou apêndice |

## Ordem Recomendada de Reescrita dos `.tex`

Não seguir a ordem física dos arquivos. Seguir a ordem de dependência:

1. `cap_fundamentacao/fundamentacao.tex`: ajustar conceitos e lower bound.
2. `cap_proposta/proposta.tex`: reescrever implementação e pipeline contra `src/`.
3. `cap_experimentos/experimentos.tex`: reescrever somente após tabelas/figuras e estatística corrigida.
4. `cap_introducao/introducao.tex`: reescrever objetivos/contribuições para refletir exatamente o que foi entregue.
5. `cap_conclusao/conclusao.tex`: reescrever por último.
6. `main_ppgco_ufu.tex`: reescrever resumo/abstract após todos os capítulos.

## Pendências Antes de Escrever no `.tex`

| Pendência | Bloqueia | Responsável provável |
|---|---|---|
| Correção de `scripts/analise-estatistica.py` | Claims de significância em Capítulo 4 | Outro agente já delegado |
| Atualizar [[analysis-methodology]] após script corrigido | Metodologia estatística e T8/F11 | Após retorno do outro agente |
| Atualizar [[resultados]] com cobertura atual | Experimentos e Conclusão | Próximo passo deste agente, se não houver conflito |
| Auditar/ajustar `scripts/consolidate_results.py` | Tabelas T1–T7 | Próximo passo técnico |
| ~~Corrigir `abrev/Abreviaturas.tex`~~ | Siglas finais | Resolvido nesta rodada; revisar apenas novas siglas |
