# Componente: Monografia

## Responsabilidade

Guardar o texto final em LaTeX e produzir PDF compilavel, com claims rastreados e referencias validadas. No desenho do repositorio, `monografia/` e a saida final do framework; `vault/` e a memoria usada para produzi-la.

## Estrutura

```text
monografia/
├── main_ppgco_ufu.tex
├── cap_introducao/
├── cap_fundamentacao/
├── cap_proposta/
├── cap_experimentos/
├── cap_conclusao/
├── bib/abntex2-references.bib
└── figs/
```

## Comando Agentico

`/compilar` executa:

1. `bash scripts/check-monografia.sh`
2. `pdflatex`, `bibtex`, `pdflatex`, `pdflatex`
3. log de compilacao via `scripts/roadmap.sh log compilation`

## Relacao com o Vault

- Papers citados devem existir no BibTeX e preferencialmente em `vault/papers/`.
- Claims fortes devem apontar para `vault/claims/` e fonte primaria.
- Auditorias em `vault/writing/auditorias/` documentam verificacoes.
- Feedback do orientador entra por `vault/writing/review-solicitacoes/`.
- Codigo e resultados em `src/` sustentam afirmacoes metodologicas e experimentais.

## Ordem Recomendada de Escrita

1. Fundamentacao e proposta com literatura validada.
2. Experimentos apos dados fechados.
3. Conclusao apos resultados e claims estabilizados.
4. Resumo e abstract por ultimo.

## Regras

- Nao citar paper que nao exista no BibTeX.
- Nao manter referencia condicional sem tarefa de remediacao.
- Nao introduzir resultado experimental sem evidencia em `src/data/results/`.
- Rodar `/compilar` antes de considerar tarefa de polimento concluida.
