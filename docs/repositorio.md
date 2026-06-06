# Estrutura Geral do Repositorio

## Projeto Principal

O repositorio e antes de tudo o TCC. O objetivo cientifico e comparar metodos de otimizacao bioinspirados sobre instancias TSP/rTSP relacionadas a patrulha de drones.

O framework agentico existe para ajudar a produzir a monografia e manter a memoria do processo, nao para ser o objeto principal da pesquisa.

## `src/` — Implementacao e Experimentos

`src/` contem o codigo Go usado como fonte primaria para claims metodologicos e experimentais.

Partes principais:

| Caminho | Papel |
|---|---|
| `src/cmd/` | Comandos Cobra usados para criar instancias, gerar grafos, otimizar e reportar resultados |
| `src/optimization/` | Implementacoes dos metodos: GA, PSO, ACO, brute force e lower bound |
| `src/graph/` | Estruturas e funcoes relacionadas aos grafos/instancias |
| `src/data/` | Instancias, grafos e resultados experimentais |
| `src/web/` | Visualizador web auxiliar para inspecionar resultados |

O visualizador em `src/web/` e um extra util para exploracao, mas nao faz parte do nucleo cientifico. A evidencia experimental deve vir dos arquivos estruturados em `src/data/results/`.

## `vault/` — Memoria do Processo

`vault/` e a memoria do framework. Ele organiza o processo de escrita e pesquisa.

Partes principais:

| Caminho | Papel |
|---|---|
| `vault/papers/` | Notas de papers e estado de leitura |
| `vault/claims/` | Claims individuais com fonte primaria e status |
| `vault/areas/` | Areas, metodos e conceitos |
| `vault/writing/` | Planejamento, auditorias, revisoes e notas de escrita |
| `vault/roadmap/` | Fila de tarefas e eventos |
| `vault/bases/` | Views Obsidian para consulta |
| `vault/canvas/` | Grafo visual de conhecimento |

O vault nao e fonte primaria por si so. Ele aponta para fontes: codigo, dados, PDFs, BibTeX, auditorias e decisoes verificadas.

## `monografia/` — Resultado Final

`monografia/` contem o texto LaTeX e o PDF final.

Partes principais:

| Caminho | Papel |
|---|---|
| `main_ppgco_ufu.tex` | Arquivo principal |
| `cap_introducao/` | Introducao |
| `cap_fundamentacao/` | Fundamentacao teorica |
| `cap_proposta/` | Proposta/metodologia |
| `cap_experimentos/` | Experimentos e resultados |
| `cap_conclusao/` | Conclusao |
| `bib/` | Bibliografia BibTeX |
| `figs/` | Figuras |

A monografia deve ser vista como saida do processo: ela recebe texto quando as evidencias correspondentes ja existem.

## Framework Agentico

O framework e formado por:

| Caminho | Papel |
|---|---|
| `.opencode/command/` | Slash commands operacionais |
| `.agents/skills/` | Skills locais do projeto |
| `.agents/council/` | Relatorios de validacao |
| `scripts/roadmap.sh` | Camada shell da fila |
| `scripts/*.sh`, `scripts/*.py` | Automacoes de suporte |

Ele atua como pesquisador assistido: busca literatura com MCPs, incorpora papers, audita referencias, cria tarefas, corrige codigo com supervisao humana, valida resultados e escreve a monografia com base em fatos.

## Retroalimentacao

O ciclo esperado e:

```text
lacuna ou bug -> tarefa -> correcao/experimento/busca -> evidencia -> vault -> monografia -> validacao -> nova tarefa se necessario
```

Exemplos:

- Se um bug em algoritmo e identificado, cria-se tarefa, corrige-se `src/`, roda-se experimento e atualiza-se claim.
- Se um metodo da literatura falta, usa-se MCPs para buscar paper, `/incorporar` para adicionar ao vault e depois adapta-se codigo ou texto.
- Se uma afirmacao da monografia nao tem evidencia, ela vira claim pendente ou tarefa de validacao.
