---
title: Cluster Based Branching for the Asymmetric Traveling Salesman Problem
authors:
- Lysgaard
- Jens
year: 1999
doi: 10.1016/S0377-2217(99)00133-2
bibtex_key: lysgaard1999cluster
bibtex-key: lysgaard1999cluster
pdf: papers/pdfs/lysgaard1999cluster.pdf
tags:
- area/atsp
- area/tsp
- evidencia/referencia
- metodo/branch-and-bound
- metodo/clustering
- metodo/lower-bound
- status/resumo-lido
- tipo/paper
status: disponivel
rating: 4
classificacao: nova-referencia
type: paper
areas:
- atsp
methods:
- branch-and-bound
- clustering
- lower-bound
role: revisao
reading_status: resumo-lido
validation_status: requer-validacao
pdf_status: integro
chapters:
- fundamentacao
claim_support: []
aliases:
- Cluster-based branching
- ATSP cluster branching
---

## PDF

![[lysgaard1999cluster.pdf]]

## Tese Central

Lysgaard propõe uma regra de ramificação para o ATSP que usa clusters identificados a partir de informação dual do procedimento de bounding, substituindo ramificações baseadas apenas em sub-rotas.

## Resumo

O artigo parte de algoritmos branch-and-bound para ATSP baseados na relaxação por problema de atribuição. Em vez de ramificar somente sobre sub-rotas, o método procura conjuntos de nós que podem aparecer consecutivamente em uma solução ótima. Esses conjuntos, chamados clusters, são identificados a partir de cutsets e variáveis duais produzidas pelo procedimento aditivo de bounding.

Quando um cluster é identificado, o algoritmo escolhe ramificar por nós de entrada ou saída, aplica testes de dominância para reduzir alternativas e contrai o cluster nos subproblemas filhos. A estratégia é integrada ao procedimento de bounding de Fischetti e Toth, mantendo a estrutura de branch-and-bound por melhor cota.

## Contribuições Principais

- Definição operacional de cluster para ATSP como conjunto de nós visitáveis consecutivamente em alguma solução ótima.
- Uso de informação dual de bounding para orientar decisões de ramificação.
- Procedimento com nós artificiais para testar candidatos a cluster quando a identificação direta falha.
- Testes de dominância para reduzir entradas e saídas candidatas.
- Comparação computacional contra ramificação por sub-rotas em instâncias ATSP da TSPLIB.

## Relevância para o TCC

É referência auxiliar para lower bounds e métodos exatos em variantes assimétricas do TSP. O TSP-SD-ATP do projeto induz dependência de sequência e pode ser comparado conceitualmente a problemas em que a estrutura da rota afeta o custo, mas o algoritmo de Lysgaard não é implementado no repositório e não deve ser citado como evidência experimental do TCC.

## Métodos e Abordagens

- Relaxação AP como cota inferior inicial.
- Procedimento aditivo de bounding para fortalecer cotas.
- Identificação de clusters por cutsets e variáveis duais.
- Ramificação por entrada ou saída do cluster, escolhendo a alternativa com menos subproblemas.
- Reduções por programação dinâmica para sub-rotas quando nenhum cluster é encontrado.

## Evidência / Resultado Relevante

- O PDF local confirma publicação em European Journal of Operational Research 119 (1999), páginas 314--325, DOI/PII `S0377-2217(99)00133-2`.
- O summary registra que a ramificação por clusters foi mais rápida em 21 de 22 instâncias avaliadas, mas esse dado deve ser validado no PDF antes de uso como claim quantitativo no texto final.

## Limitações de Uso

- O método é específico para ATSP e para a infraestrutura de bounding usada no artigo.
- O custo de identificar clusters cresce com o tamanho do cluster, e grandes instâncias ainda podem esbarrar em limites de tempo ou memória.
- Não fundamenta diretamente GA, PSO ou ACO; seu uso é contextual para métodos exatos/lower bounds.

## Conexões

- Relacionado a: [[balas1985branch]], [[lower-bounds]], [[tsp-variants]]
- Contrasta com: meta-heurísticas usadas no projeto
- Usado em capítulo: [[fundamentacao]]
