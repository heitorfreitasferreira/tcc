---
title: A Powerful Genetic Algorithm Using Edge Assembly Crossover for the Traveling Salesman Problem
authors:
- Nagata
- Yuichi
- Kobayashi
- Shigenobu
year: 2013
doi: 10.1287/ijoc.1120.0506
bibtex_key: nagata2013eax
bibtex-key: nagata2013eax
aliases:
  - nagata2006eax
pdf: papers/pdfs/nagata2006eax.pdf
tags:
- area/tsp
- evidencia/referencia
- metodo/crossover
- metodo/ga
- metodo/eax
- status/lido
- tipo/paper
status: lido
rating: 4
classificacao: recuperado
type: paper
methods:
- crossover
- ga
- eax
role: revisao
reading_status: lido
validation_status: nao-validado
pdf_status: lido
note_bibkey: "A chave histórica nagata2006eax refere-se à versão de conferência em EvoCOP 2006; este PDF é a versão de periódico (INFORMS J. Computing 2013), que estende a versão 2006. A chave atual foi renomeada para nagata2013eax para alinhar com o ano de publicação do periódico; o alias nagata2006eax é mantido para não quebrar referências antigas."
---

## Resumo

Publicado no **INFORMS Journal on Computing 25(2) (Spring 2013), pp. 346–363** (DOI 10.1287/ijoc.1120.0506; copyright © 2013 INFORMS). É a **versão de periódico estendida** do trabalho de conferência apresentado em EvoCOP 2006 (Nagata 2006a/b) e baseia-se no operador EAX (Edge Assembly Crossover) introduzido por Nagata & Kobayashi (1997). O paper propõe um GA com três melhorias substanciais ao EAX original:

1. **Localização do EAX** com implementação eficiente (Nagata 2006a): versão *localizada* que gera *offspring solutions* alterando apenas arestas dos *E-sets* (sub-grafos) e, em caso de falha, muda para versão *global*. Reduz drasticamente o custo computacional do cruzamento.
2. **Busca local integrada ao EAX**: procedimento simples que determina boas combinações dos blocos de soluções-pai; múltiplas versões *globais* do EAX (original, K-múltipla, block strategy) são discutidas.
3. **Modelo de seleção inovador** para manter diversidade populacional com custo computacional desprezível (mantém o melhor e usa regra de substituição restrita).

Resultados: o GA proposto **obtém todas as melhores soluções conhecidas** em benchmarks TSPLIB, incluindo instâncias de até 200.000 cidades, e **melhora várias BKS** (Best Known Solutions) reportadas por Helsgaun (LKH). Supera variantes do Lin-Kernighan sem usar LK diretamente. Resultados detalhados por instância são fornecidos no supplement online (DOI 10.1287/ijoc.1120.0506).

**Atenção — chave de citação**: a chave de bibtex é `nagata2013eax` (ano de publicação do periódico). A chave `nagata2006eax` (EvoCOP) é mantida como alias, mas refere-se ao mesmo trabalho — o periódico é uma extensão substancial do trabalho de conferência de 2006.

## Contribuições Principais

- Localização do EAX: gera *offspring* restrito a E-sets locais, com fallback para versão global — reduz custo sem perder qualidade
- Integração de busca local no EAX: procedimento simples em `Step 4` para gerar soluções intermediárias de boa qualidade a partir de pares de pais
- Modelo de seleção que preserva diversidade com custo desprezível: regra de seleção/restrição no `Step 5` que mantém a melhor solução e substitui baseada em critério
- Múltiplas versões globais do EAX (random, K-multiple, block) com análise empírica
- Supera LKH em instâncias grandes e melhora várias BKS da TSPLIB
- Código disponível no supplement online do paper

## Relevância para o TCC

Referência canônica para crossover em GA aplicado a TSP. O EAX é citado na Seção 2.5 (GA) como contraste ao OX (Order Crossover) implementado no TCC. Para fundamentação, é usado para justificar a importância do operador de recombinação em problemas de permutação, mostrando que operadores baseados em arestas (edge-based) superam operadores baseados em posição/ordem (position/order-based) em qualidade de solução.

## Métodos e Abordagens

- EAX (*Edge Assembly Crossover*): gera soluções-filhas combinando arestas de dois pais via ciclos AB (alternando arestas dos pais); produz *E-sets* (sub-grafos) que definem quais arestas serão trocadas
- Versão *localizada*: tenta primeiro alterar apenas arestas locais do E-set; se não gerar *offspring* válido, muda para versão *global* (modifica todo o tour)
- Versões *globais* (random, K-multiple, block): diferentes estratégias de seleção de AB-cycles dentro do E-set
- Tamanho do E-set: parâmetro que controla quantas arestas serão modificadas; ajuste automático via critério de parada
- Modelo de seleção: cada par de pais gera `N_offspring` filhos; o melhor substitui uma solução existente se passar no critério
- Busca local: procedimento simples de 2-opt/Or-opt integrado no `Step 4` para refinar soluções intermediárias
- Parâmetros: população pequena (≈30), mas com o modelo de seleção robusto

## Conexões

- [[genetic-algorithms]]
- [[tsp]]
- [[potvin1996ga]] — revisão de operadores de crossover para TSP
- [[goldberg1989genetic]] — textbook clássico de GA
- [[holland1975adaptation]] — GA, fundação teórica
- [[lin1973effective]] — Lin-Kernighan, referência para heurísticas locais

## Notas e Insights

- O EAX + busca local produz resultados **comparáveis ao LKH** com **tempo de execução competitivo** em instâncias grandes (até 200k cidades)
- O paper **melhorou várias BKS** da TSPLIB, demonstrando que o arcabouço é estado-da-arte
- O modelo de seleção é uma das chaves do sucesso: população pequena (≈30) com regra de manutenção de diversidade
- A versão *localizada* do EAX é o que torna o algoritmo computacionalmente viável para instâncias grandes
- O paper de conferência (Nagata 2006a/b) é a base; o periódico (este PDF) adiciona o modelo de seleção robusto e mais versões globais
- O código está disponível em http://dx.doi.org/10.1287/ijoc.1120.0506 (supplement)

## Citações-chave

> "To construct a powerful GA, we use edge assembly crossover (EAX) and make substantial enhancements to it: (i) localization of EAX together with its efficient implementation and (ii) the use of a local search procedure in EAX to determine good combinations of the parents' solution blocks."

> "The proposed GA improves all best-known solutions on instances with up to 200,000 cities."

> "The proposed GA is a powerful search algorithm that is competitive with state-of-the-art heuristics such as LKH."
