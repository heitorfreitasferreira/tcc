# Componente: Validacao e Council

## Responsabilidade

Reduzir erro, alucinacao e regressao por meio de validacao explicita.

## Council

O council usa juizes independentes para validar uma tarefa, plano ou resultado.

No OpenCode, os juizes sao subagentes `task(subagent_type="general")` que escrevem relatorios em `.agents/council/`.

## Vereditos

| Veredito | Significado | Acao padrao |
|---|---|---|
| `PASS` | Trabalho suficiente | Concluir tarefa |
| `WARN` | Diagnostico correto com pendencias ou risco | Concluir se nao houver bloqueio; criar tarefas corretivas |
| `FAIL` | Trabalho incorreto ou incompleto | Nao concluir; registrar bloqueio |

## Auditorias

Auditorias ficam em `vault/writing/auditorias/` e devem listar:

- escopo;
- fontes verificadas;
- achados;
- riscos;
- acoes realizadas;
- tarefas derivadas.

## Anti-Alucinacao

O texto final deve respeitar estas regras:

- Claim experimental precisa de dado em `src/data/results/`.
- Claim metodologico precisa de codigo ou script.
- Claim bibliografico precisa de BibTeX e nota de paper.
- Referencia condicional deve ser removida ou virar tarefa.
- Todo resultado agregado deve ser reprodutivel por script.

## Relatorios

Relatorios de validacao ficam em `.agents/council/YYYY-MM-DD-validate-<target>.md`.

Eles podem gerar novas tarefas via `/tarefa`, como ocorreu com as recomendacoes de P18.
