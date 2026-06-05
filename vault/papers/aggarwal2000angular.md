---
title: The Angular-Metric Traveling Salesman Problem
authors:
- Aggarwal
- Alok
- Coppersmith
- Don
- Khanna
- Sanjeev
- Motwani
- Rajeev
- Schieber
- Baruch
year: 1999
doi: 10.1137/S0097539796312721
bibtex_key: aggarwal1999angular
bibtex-key: aggarwal1999angular
aliases:
  - aggarwal2000angular
pdf: papers/pdfs/aggarwal2000angular.pdf
tags:
- area/tsp
- area/tsp-variants
- evidencia/referencia
- metodo/approximation
- status/resumo-lido
- tipo/paper
status: disponível
rating: 4
classificacao: recuperado
type: paper
areas:
- tsp-variants
methods:
- approximation
role: revisao
reading_status: resumo-lido
validation_status: requer-validacao
pdf_status: integro
chapters:
- fundamentacao
claim_support: []
note_bibkey: "Ano de publicação: 1999 (publicação eletrônica em 7 de dezembro de 1999; volume impresso SIAM J. Comput. 2000 Vol. 29, No. 3). A chave de bibtex foi renomeada para aggarwal1999angular; o alias aggarwal2000angular é mantido para preservar wikilinks antigos do vault. DOI corrigido em P46 para 10.1137/S0097539796312721, confirmado no PII/URL da primeira página do PDF."
---

## PDF

![[aggarwal2000angular.pdf]]

## Tese Central

Aggarwal et al. formalizam o Angle-TSP, uma variante em que o custo da rota é a soma das mudanças de direção nos vértices, e mostram que a otimização angular tem dificuldade computacional própria: Angle-TSP e Angle-CCP são NP-difíceis, mas admitem aproximação polinomial de ordem `O(log n)`.

## Resumo

O artigo motiva o problema por aplicações em robótica, veículos com restrições de manobra, plataformas de alta inércia e planejamento de trajetórias em que girar pode ser mais caro que percorrer distância. A formulação considera um conjunto de pontos no plano e procura um ciclo hamiltoniano que minimize o custo angular total, definido pela soma das mudanças de direção no percurso.

Além da prova de NP-dificuldade, o trabalho apresenta algoritmos de aproximação para Angle-TSP e para a relaxação de cobertura por ciclos. A abordagem combina decomposição recursiva, programação dinâmica e remendos de ciclos, obtendo razão `O(log n)`. O artigo também analisa limites extremais para o custo angular ótimo e o compromisso entre minimizar ângulo e minimizar comprimento.

## Contribuições Principais

- Formalização do Angle-TSP como variante geométrica do TSP motivada por custo de rotação.
- Prova de NP-dificuldade para Angle-TSP e Angle-CCP por redução de one-in-three 3SAT.
- Algoritmo de aproximação `O(log n)` para Angle-TSP e Angle-CCP.
- Limites extremais essencialmente apertados para o custo angular ótimo, da ordem de `O(n / log n)`.
- Análise do trade-off entre razão de aproximação angular e razão de aproximação por comprimento.

## Relevância para o TCC

É uma referência central para justificar que custos de rotação ou mudança de direção não são detalhe cosmético em problemas de rota. O TSP-SD-ATP do projeto não é o Angle-TSP puro, mas compartilha a motivação de que o custo de visitar um ponto depende do predecessor e do sucessor, não apenas da aresta percorrida. A nota pode apoiar a fundamentação de variantes com penalidade angular, desde que o texto preserve a diferença entre custo angular acumulado e o modelo implementado no repositório.

## Métodos e Abordagens

- Redução de complexidade com gadgets de U-turn para codificar atribuições booleanas.
- Decomposição recursiva de caminhos com ângulo limitado.
- Programação dinâmica para selecionar caminhos e ciclos candidatos.
- Patching de ciclos com penalidade angular controlada.
- Análise assintótica baseada no teorema de Erdős-Szekeres para limites extremais.

## Evidência / Resultado Relevante

- O PDF registra `PII. S0097539796312721` e publicação eletrônica em 7 de dezembro de 1999.
- A aproximação `O(log n)` e a NP-dificuldade sustentam a cautela metodológica: problemas com custo angular podem exigir heurísticas ou aproximações mesmo quando a geometria parece simples.
- O resultado de trade-off mostra que otimizar suavidade angular e comprimento simultaneamente pode impor perdas em um dos objetivos.

## Limitações de Uso

- O artigo estuda Angle-TSP geométrico, não a formulação TSP-SD-ATP implementada no projeto.
- A discussão é teórica; não fornece comparação empírica com GA, PSO ou ACO.
- Claims quantitativos sobre razões de aproximação devem ser citados como resultados do modelo Angle-TSP, não como limites diretos do TSP-SD-ATP.

## Conexões

- Relacionado a: [[tsp-variants]], [[problem-formulation]]
- Complementa: [[winter2002modeling]], [[vanhove2012route]]
- Usado em capítulo: [[fundamentacao]]

## Notas e Insights

- O DOI/PII correto é `10.1137/S0097539796312721`; a nota antiga registrava `...12719`.
- A chave `aggarwal1999angular` deve ser usada em citações novas. `aggarwal2000angular` permanece apenas como alias histórico.
