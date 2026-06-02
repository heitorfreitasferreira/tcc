---
title: "Genetic Algorithms in Search, Optimization, and Machine Learning"
authors: [Goldberg, David E.]
year: 1989
doi: ""
bibtex-key: goldberg1989genetic
pdf: "papers/pdfs/goldberg1989genetic.pdf"
tags: [ga metaheuristic]
status: lido
rating: 5
---

## PDF

![[goldberg1989genetic.pdf]]

## Resumo

Livro-texto clássico que sistematizou e popularizou os Algoritmos Genéticos (GAs), apresentando-os como ferramentas práticas de busca, otimização e aprendizado de máquina. Goldberg expande o trabalho fundacional de Holland (1975) com uma abordagem tutorial que inclui fundamentos matemáticos (teorema dos esquemas, blocos construtores), implementação computacional (SGA — Simple Genetic Algorithm em Pascal) e aplicações em engenharia e ciência da computação. É a referência mais citada na literatura de GAs.

## Contribuições Principais

- Sistematização do Algoritmo Genético Simples (SGA) com implementação completa
- Formalização do Teorema dos Esquemas (*Schema Theorem*) e da Hipótese dos Blocos Construtores
- Introdução de técnicas avançadas (dominância, diploidia, inversão, nicho, compartilhamento)
- Demonstração de aplicações reais (otimização de gasodutos, estruturas, aprendizado classificador)
- Estabelecimento dos GAs como técnica de otimização reconhecida na academia e indústria

## Relevância para o TCC

O GA implementado no repositório (`src/pkg/ga/`) segue a arquitetura SGA descrita por Goldberg — população, seleção por roleta, crossover e mutação. O Teorema dos Esquemas fornece a justificativa teórica para o funcionamento do GA aplicado ao TSP. A discussão sobre nicho e compartilhamento é relevante para manter diversidade populacional em problemas de roteamento.

## Métodos e Abordagens

- Algoritmo Genético Simples (SGA): seleção por roleta, crossover de um ponto, mutação bit-flip
- Teorema dos Esquemas e Hipótese dos Blocos Construtores
- Mapeamento de código binário para parâmetros reais (fixed-point coding)
- Operadores avançados: inversão, dominância/diploidia, nicho e compartilhamento
- Sistemas Classificadores (Learning Classifier Systems) com Bucket Brigade
- Aplicações em otimização de funções, pipeline, estruturas e jogos

## Conexões

- [[holland1975adaptation]] — obra fundamental que este livro expande e torna acessível
- [[oliver1987crossover]] — operadores de cruzamento para TSP (aplicação direta dos conceitos)
- [[bean1994genetic]] — random keys (diretamente relacionado)
- [[applegate2006traveling]] — métodos exatos contrastam com a abordagem heurística dos GAs
- [[genetic-algorithms]]

## Notas e Insights

- A implementação SGA em Pascal (Apêndice) serviu de template para inúmeras implementações posteriores
- O livro foca em codificação binária; para TSP, a representação por permutação exige adaptações (PMX, OX)
- O Teorema dos Esquemas é a principal justificativa teórica para a eficácia dos GAs
- A seção sobre Sistemas Classificadores abriu caminho para GA-based machine learning
- O GA do repositório usa seleção por torneio e elitismo (não roleta), variações justificadas pela literatura posterior

## Citações-chave

> "Genetic algorithms are search algorithms based on the mechanics of natural selection and natural genetics."

> "Unlike most optimization methods, that work from a single point in the decision space [...] genetic algorithms work from an entire 'population' of points simultaneously."
