# TCC — Otimizacao Bioinspirada para TSP/rTSP

Este repositorio e o meu Trabalho de Conclusao de Curso. O projeto principal compara metodos de otimizacao bioinspirados aplicados a uma variante do TSP/rTSP motivada por patrulha de drones: visitar pontos de interesse com rota/tempo minimos considerando o custo da trajetoria.

O repositorio tambem contem um framework agentico de pesquisa que apoia a escrita da monografia. Esse framework nao e o produto final do TCC; ele e a infraestrutura de trabalho que organiza memoria, evidencias, backlog, literatura, validacao e escrita.

## Estrutura Geral do Repositorio

| Parte | Caminho | Papel no TCC |
|---|---|---|
| Experimentos | `src/` | Codigo Go usado para gerar instancias, grafos, algoritmos e resultados experimentais |
| CLI experimental | `src/cmd/` | Comandos Cobra para `create`, `graph`, `optimize`, metodos e reporting |
| Algoritmos | `src/optimization/` | Implementacoes de GA, PSO, ACO, brute force e lower bound |
| Dados e resultados | `src/data/` | Instancias, grafos e resultados estruturados usados como evidencia |
| Visualizador | `src/web/` | Interface extra para inspecionar resultados; util, mas fora do escopo cientifico central |
| Memoria de pesquisa | `vault/` | Vault Obsidian com papers, claims, areas, auditorias, roadmap e notas de escrita |
| Texto final | `monografia/` | Monografia LaTeX, bibliografia, figuras e PDF final |
| Framework agentico | `.opencode/`, `.agents/`, `scripts/roadmap.sh` | Camada que opera sobre codigo, vault e monografia |

## Relacao Entre as Partes

```text
src/             -> fatos de codigo e resultados experimentais
papers/PDFs      -> fatos da literatura
vault/           -> memoria organizada desses fatos
framework        -> agente que consulta, valida, corrige e escreve
monografia/      -> resultado textual final
```

O framework se baseia em fatos. Ele deve escrever ou revisar a monografia apenas a partir de evidencias verificaveis: codigo implementado, resultados experimentais, PDFs/notas de papers ou auditorias registradas.

Ele tambem se retroalimenta: quando encontra uma lacuna ou erro, cria tarefas, corrige codigo com supervisao humana, incorpora metodos da literatura, busca textos academicos com MCPs e atualiza o vault antes de alterar o texto final.

## Framework Agentico de Pesquisa

O framework transforma o processo de pesquisa e escrita em um ciclo fechado:

1. Literatura entra por `/consultar` e `/incorporar`.
2. Evidencias ficam no `vault/` como papers, claims, areas, auditorias e notas de escrita.
3. Tarefas entram na fila em `vault/roadmap/tarefas/` por `/tarefa`, `/claudiney` ou pipelines automaticos.
4. `/proximo` consome a primeira tarefa pendente por ordem de fase, executa, valida com `council` e conclui.
5. Experimentos e compilacoes geram logs em `vault/roadmap/eventos/`.
6. A monografia em `monografia/` recebe apenas texto apoiado por fatos.

## Componentes Principais

| Componente | Caminho | Funcao no framework |
|---|---|---|
| Comandos agenticos | `.opencode/command/` | Interfaces slash command para pesquisa, escrita, validacao e automacao |
| Fila de roadmap | `vault/roadmap/tarefas/` | Tarefas atomicas `P<N>.md` ordenadas por fase e ordem |
| Script de roadmap | `scripts/roadmap.sh` | Camada shell para criar, listar, concluir e logar tarefas |
| Vault semantico | `vault/` | Base de conhecimento com papers, claims, areas, auditorias e bases Obsidian |
| Bibliografia | `monografia/bib/abntex2-references.bib` | Fonte BibTeX da monografia |
| Monografia | `monografia/` | Texto LaTeX final e artefatos de compilacao |
| Codigo experimental | `src/` | Fonte primaria para claims metodologicos e experimentais |
| Skills | `.agents/skills/` e `~/.agents/skills/` | Instrucoes especializadas carregadas por fase ou dominio |
| Council | `.agents/council/` | Relatorios de validacao multi-juiz |

## Fluxo Recomendado

```text
/consultar -> /incorporar -> /tarefa ou /claudiney -> /proximo -> council -> concluir -> /compilar
```

Para trabalho experimental:

```text
/experimento -> resultados em src/data/results/ -> claims experimentais -> monografia
```

Para feedback do orientador:

```text
/claudiney -> review note -> tarefas P<N> -> /proximo
```

## Comandos Mais Usados

```bash
# Ver a proxima tarefa pendente
bash scripts/roadmap.sh proximo

# Criar tarefa manual
bash scripts/roadmap.sh tarefa criar "Titulo" "Saida esperada" alta escrita

# Marcar tarefa como concluida
bash scripts/roadmap.sh tarefa concluir P48

# Logar evento de experimento, incorporacao ou compilacao
bash scripts/roadmap.sh log experiment method=aco seeds=0-50 status=ok
```

Slash commands principais:

| Comando | Uso |
|---|---|
| `/consultar` | Busca literatura em MCPs academicos |
| `/incorporar` | DOI/PDF/arXiv/BibTeX -> vault + BibTeX + canvas + claims |
| `/tarefa` | Cria tarefa manual na fila |
| `/claudiney` | Converte feedback do orientador em review note e tarefas |
| `/proximo` | Executa a proxima tarefa pendente com validacao |
| `/experimento` | Roda experimentos batch e loga resultados |
| `/compilar` | Valida e compila a monografia |
| `/commitar` | Revisa e commita mudancas |

## Ordem da Fila

As tarefas sao ordenadas por fase e depois por `ordem` numerica:

```text
infra < literatura < experimentacao < analise < escrita < polimento < revisao
```

Essa ordem evita escrever texto final antes de fechar literatura, experimentos e evidencias.

## Documentacao

A documentacao detalhada esta em `docs/`:

| Documento | Conteudo |
|---|---|
| `docs/arquitetura.md` | Visao geral, principios e fluxo de dados |
| `docs/roadmap.md` | Fila ordenada, formato das tarefas e eventos |
| `docs/comandos.md` | Cada slash command e sua responsabilidade |
| `docs/literatura.md` | Busca, incorporacao de papers, BibTeX e PDFs |
| `docs/vault.md` | Schema semantico do vault, papers, claims e tags |
| `docs/validacao.md` | Council, auditorias, claims e anti-alucinacao |
| `docs/experimentos.md` | CLI Go, batch experimental e resultados |
| `docs/monografia.md` | Escrita, compilacao e relacao com evidencias |
| `docs/skills.md` | Skills por fase e como elas entram no fluxo |
| `docs/scripts.md` | Scripts auxiliares e contratos de entrada/saida |
| `docs/go-cli.md` | Componente Go usado para gerar dados experimentais |

## Invariantes do Repositorio

- Toda tarefa executavel deve ter `task_id`, `fase`, `ordem`, `priority`, `status`, `origin` e `saida_esperada`.
- Claims fortes na monografia precisam apontar para fonte primaria: codigo, dados, literatura validada ou auditoria.
- Papers citados no texto final devem ter entrada BibTeX e nota no vault.
- Referencias condicionais devem virar tarefa antes de permanecerem no texto final.
- Resultados experimentais devem vir de `src/data/results/`, nao de tabelas manuais.
- Mudancas relevantes devem ser validadas por council ou justificadas quando a validacao for pulada.
- O visualizador em `src/web/` pode apoiar inspecao, mas nao deve ser tratado como evidencia cientifica central.
- O vault e memoria; a monografia e o resultado; o codigo e os dados sao fontes primarias.

## Estado Operacional

Use `/proximo` para continuar o trabalho orientado pela fila. Use `bash scripts/roadmap.sh proximo` para inspecionar a proxima tarefa sem acionar o fluxo agentico completo.
