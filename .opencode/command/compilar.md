---
description: Compila a monografia (pdflatex + check-monografia.sh) e registra no log do roadmap
---

# /compilar

Executa validação + compilação completa da monografia e registra o resultado no log do roadmap.

## Fluxo

### 1. Validar
```bash
bash scripts/check-monografia.sh
```
Se FAIL → reporta erros, pergunta se deseja continuar mesmo assim.

### 2. Compilar
```bash
cd monografia && pdflatex main_ppgco_ufu.tex && bibtex main_ppgco_ufu && pdflatex main_ppgco_ufu.tex && pdflatex main_ppgco_ufu.tex
```
Captura erros e warnings do output do LaTeX.

### 3. Logar
```bash
bash scripts/roadmap.sh log compilation \
  check_monografia=PASS \
  pdflatex_errors=0 \
  warnings=3
```

### 4. Reportar
Mostra:
- Resultado do check: PASS / WARN / FAIL
- Resultado do pdflatex: OK (0 erros) ou lista de erros
- Warnings: overfull boxes, undefined references, etc.
- PDF gerado: `monografia/main_ppgco_ufu.pdf`
- Evento logado em `vault/roadmap/eventos/`

### 5. (Opcional) Commitar mudanças
```bash
git status --short
```
Se há mudanças no working tree:
1. Gera mensagem: `docs: compilação monografia <data>`.
2. Usa toolcall `question` com opções `["SIM"]`, `["NAO"]`.
3. **SIM**: `git add -A && git commit -m "<mensagem>"`.
4. **NAO**: mostra o comando como texto e encerra sem executar.

## Dependências
- `scripts/check-monografia.sh`
- `pdflatex`, `bibtex` no PATH
