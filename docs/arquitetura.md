# Arquitetura do Repositorio e do Framework Agentico

## Objetivo

Organizar o TCC como um sistema de pesquisa rastreavel: o codigo gera evidencias experimentais, o vault guarda memoria e contexto, a monografia e o resultado final, e o framework agentico coordena a escrita e a manutencao desse ciclo.

O framework nao substitui o TCC. Ele e uma camada operacional que escreve, revisa, valida e retroalimenta o trabalho com base em fatos.

## Camadas

| Camada | Fonte canonica | Papel |
|---|---|---|
| Codigo experimental | `src/cmd/`, `src/optimization/`, `src/graph/` | Implementacao dos algoritmos, CLI e geracao de dados |
| Dados | `src/data/results/` | Saidas experimentais estruturadas |
| Visualizacao auxiliar | `src/web/` | Interface para inspecao; fora do escopo cientifico central |
| Literatura | `monografia/bib/`, `vault/papers/` | Referencias, PDFs e notas de leitura |
| Conhecimento | `vault/` | Claims, areas, auditorias, decisoes e planejamento |
| Texto final | `monografia/` | Monografia LaTeX |
| Orquestracao | `.opencode/command/`, `scripts/roadmap.sh` | Automacao agentica e fila |

## Estrutura Mental

```text
codigo + dados + PDFs -> fatos
vault                 -> memoria dos fatos
framework             -> agente que opera sobre a memoria e os fatos
monografia            -> saida textual final
```

O framework sempre deve perguntar: qual e a fonte primaria deste trecho? Se a resposta nao for codigo, resultado, PDF/nota validada ou auditoria, o trecho deve virar tarefa ou claim pendente.

## Ciclo de Trabalho

1. Descobrir lacuna, bug, demanda de escrita ou necessidade de literatura.
2. Criar tarefa ou review note.
3. Executar pelo consumidor `/proximo`.
4. Usar skills, scripts, codigo e MCPs adequados.
5. Registrar ou corrigir evidencia em codigo, dados, vault ou BibTeX.
6. Escrever ou ajustar a monografia somente depois da evidencia.
7. Validar com council.
8. Marcar tarefa como concluida e sugerir a proxima.

## Fluxo de Dados

```text
MCPs academicos -> /consultar -> /incorporar -> BibTeX + vault/papers + canvas
codigo em src/ -> experimentos -> src/data/results -> claims experimentais -> monografia
bugs ou lacunas -> tarefa -> correcao supervisionada -> novos fatos -> vault/monografia
feedback orientador -> /claudiney -> review note -> tarefas -> /proximo
```

## Decisoes de Projeto

- O vault organiza evidencia, mas nao substitui fontes primarias.
- A fila e monotona: cada tarefa tem um ID `P<N>` e so muda de `pendente` para `concluida`.
- Validacao e explicita: council gera relatorios em `.agents/council/`.
- Escrita academica so deve usar referencias existentes no BibTeX e preferencialmente com nota no vault.
- O visualizador web e ferramenta auxiliar; resultados cientificos devem vir de `src/data/results/`.

## Falhas Esperadas

- PDF paywalled: registrar `pdf_status: ausente` e criar tarefa de remediacao.
- Feedback vago: `/claudiney` deve perguntar antes de gerar tarefas.
- Council `FAIL`: nao concluir a tarefa; registrar bloqueio e gerar acao corretiva.
- Resultados incompletos: `/experimento` deve alertar pela checagem de cobertura.
