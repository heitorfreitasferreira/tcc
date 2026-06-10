---
description: Revisa texto da monografia aplicando regras stop-slop
---

# /revisar

Carregue o skill `stop-slop` via `skill("stop-slop")`.
Carregue o skill `academic-writing` via `skill("academic-writing")`.

Aplique as regras de remoção de AI tells no texto informado:

- Corte advérbios, voz passiva, falsos contrastes, throat-clearing openers
- Substitua agency falsa por sujeitos humanos
- Varie ritmo, elimine meta-comentário

Se um caminho de arquivo for passado como argumento, leia o arquivo e revise-o.
Se não, revise o texto informado diretamente ou peça esclarecimento.

## Uso

```
/revisar monografia/cap_introducao.tex
/revisar "O texto a ser revisado aqui"
```
