---
title: 'Computers and Intractability: A Guide to the Theory of NP-Completeness'
authors:
- Garey
- Michael R.
- Johnson
- David S.
year: 1979
doi: ''
bibtex_key: garey1979computers
bibtex-key: garey1979computers
pdf: papers/pdfs/garey1979computers.pdf
tags:
- area/tsp
- evidencia/referencia
- papel/fundacional
- status/lido-parcial
- tipo/paper
status: lido-parcial
rating: 5
type: paper
role: fundacional
reading_status: lido-parcial
validation_status: nao-validado
---

## PDF

![[garey1979computers.pdf]]

## Resumo

Garey e Johnson apresentam a teoria de NP-completude como ferramenta prática para reconhecer problemas computacionalmente intratáveis, especialmente quando não se encontra algoritmo polinomial para problemas de otimização combinatória. O PDF enfatiza a distinção entre algoritmos polinomiais e exponenciais, formaliza a passagem de problemas de otimização para problemas de decisão e usa `TRAVELING SALESMAN` como exemplo recorrente: dado um conjunto de cidades, distâncias e um limite `B`, pergunta-se se existe uma rota com comprimento total no máximo `B`. A obra mostra que provar NP-completude não encerra o estudo de um problema, mas redireciona a estratégia para casos especiais, algoritmos heurísticos, aproximações ou relaxações.

> [!warning] Leitura parcial
> O PDF local parece ser scan sem camada textual confiável. A extração exigiu OCR seletivo e apresentou erros frequentes em fórmulas, caracteres e colunas; por isso a nota fica como `lido-parcial`.

## Contribuições Principais

- Consolida a teoria de NP-completude como método para classificar problemas computacionalmente difíceis.
- Define a intratabilidade operacionalmente pela ausência esperada de algoritmo polinomial para o problema geral.
- Mostra como problemas de otimização, como TSP, podem ser analisados por versões decisórias com limite `B`.
- Inclui `TRAVELING SALESMAN`, `GEOMETRIC TRAVELING SALESMAN` e `BOTTLENECK TRAVELING SALESMAN` no catálogo de problemas difíceis.
- Explica que a prova de NP-completude orienta a busca por heurísticas, casos especiais e relaxações.

## Relevância para o TCC

A obra fundamenta a justificativa central do TCC: o TSP/rTSP usado como base para rotas de drones pertence a uma classe de problemas em que soluções exatas gerais não escalam bem. Isso legitima a comparação experimental entre força bruta e metaheurísticas bio-inspiradas, pois a dificuldade estrutural do problema torna natural estudar métodos aproximados, estocásticos e dependentes de parametrização.

## Métodos e Abordagens

- Formalização de problemas de decisão por instância e pergunta `sim/não`.
- Uso de codificações razoáveis e análise de complexidade em tempo polinomial.
- Definição de `NP` via verificabilidade polinomial e algoritmos não determinísticos.
- Reduções para demonstrar que um problema é pelo menos tão difícil quanto outro.
- Tratamento de problemas de otimização por sua versão decisória com limite de custo.

## Conexões

- [[halim2019combinatorial]]
- [[lawler1985traveling]]
- [[applegate2006traveling]]
- [[karp1979patching]]
- [[heldkarp1970traveling]]
- [[balas1985branch]]
- [[tsp]]

## Notas e Insights

- A obra não apresenta heurísticas para TSP como foco principal; sua função é justificar por que heurísticas e casos especiais são necessários.
- A passagem sobre TSP separa claramente dificuldade de decisão e dificuldade de otimização.
- O texto recomenda mudar o foco após uma prova de NP-completude: procurar casos especiais, algoritmos úteis na prática ou relaxações.
- Para o TCC, a referência deve sustentar a base de complexidade, não resultados experimentais.

## Citações-chave

Paráfrase baseada em OCR ruim: NP-completude permite provar que um problema é tão difícil quanto muitos problemas já reconhecidos como difíceis.

Paráfrase baseada em OCR ruim: se `TRAVELING SALESMAN` é NP-completo em sua forma decisória, então o problema de otimização correspondente é pelo menos tão difícil.
