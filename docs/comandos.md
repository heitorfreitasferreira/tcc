# Componente: Slash Commands

Os comandos em `.opencode/command/` sao contratos operacionais para o agente. Eles descrevem fluxos, ferramentas, validacoes e saidas esperadas.

## `/consultar`

Busca literatura em bases academicas via MCPs.

- Entrada: query textual.
- Saida: tabela de candidatos com titulo, ano, DOI, fonte e status OA.
- Nao altera arquivos por padrao.

## `/incorporar`

Ingere um paper no ecossistema.

- Entrada: DOI, arXiv ID, chave BibTeX, PDF local, `vault` ou `fila`.
- Saida: PDF em `vault/papers/pdfs/`, entrada BibTeX, nota em `vault/papers/`, canvas atualizado e log.
- Nao altera `.tex` da monografia.

## `/tarefa`

Cria tarefa manual no roadmap.

- Entrada: titulo, descricao, prioridade e fase.
- Saida: `vault/roadmap/tarefas/P<N>.md`.
- Uso recomendado para acoes pequenas e diretas.

## `/claudiney`

Registra feedback do orientador.

- Entrada: texto do feedback.
- Saida: review note em `vault/writing/review-solicitacoes/` e uma ou mais tarefas.
- Uso exclusivo para feedback externo que exige analise contextual.

## `/proximo`

Consumidor principal da fila.

- Entrada: nenhuma.
- Saida: tarefa executada, validada e marcada como concluida.
- Carrega skills conforme fase.
- Usa council por padrao quando a tarefa nao e trivial.

## `/experimento`

Executa batch experimental.

- Entrada: metodo, frequencias, seeds, jobs e politica de sobrescrita.
- Saida: resultados em `src/data/results/` e log de evento.
- Valida cobertura com `scripts/gerar-metricas-monografia.py --check`.

## `/compilar`

Valida e compila a monografia.

- Entrada: nenhuma.
- Saida: PDF em `monografia/main_ppgco_ufu.pdf` e log de compilacao.
- Executa `scripts/check-monografia.sh` antes de compilar.

## `/commitar`

Revisa e commita mudancas.

- Entrada: nenhuma.
- Saida: commit Git ou sugestao de comando.
- Deve mostrar `git status`, `git diff --stat` e commits recentes antes de confirmar.

## Comandos Depreciados

- `/roadmap-consumir`: substituido por `/proximo`.
- `/roadmap-criar`: substituido por `/tarefa` e `/claudiney`.
- `/baixar-pdf`: substituido por `/incorporar`.
