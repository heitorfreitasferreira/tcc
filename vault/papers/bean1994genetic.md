---
title: "Genetic Algorithms and Random Keys for Sequencing and Optimization"
authors: [Bean, James C.]
year: 1994
doi: "10.1287/ijoc.6.2.154"
bibtex_key: bean1994genetic
bibtex-key: bean1994genetic
pdf: "papers/pdfs/bean1994genetic.pdf"
tags: [ga, metaheuristic, random-keys, encoding]
status: resumo-lido
rating: 4
classificacao: condicional
---

## Resumo

Publicado no INFORMS Journal on Computing (1994). Artigo fundacional que introduz o conceito de *random keys* para algoritmos genéticos aplicados a problemas de sequenciamento e otimização combinatória. A ideia central é representar cada item de uma sequência por um número real aleatório (a chave) e obter a permutação pela ordenação dessas chaves. Isso permite que operadores genéticos padrão (crossover e mutação contínuos) sejam aplicados sem produzir soluções infactíveis — a ordenação garante uma permutação válida. O método foi validado em problemas de sequenciamento como *single machine scheduling* e *resource-constrained project scheduling*. Com mais de 1.300 citações, é a referência canônica para random keys em computação evolutiva.

**Limitação**: o PDF completo não pôde ser obtido (paywall INFORMS, Sci-Hub não cataloga, Unpaywall confirma `oa_status: closed`). O resumo foi reconstruído a partir de metadados, citações na literatura e descrições em artigos que referenciam o método.

## Contribuições Principais

- Introdução do conceito de random keys: codificar soluções combinatórias como vetores reais, decodificar por ordenação
- Garantia de factibilidade: qualquer vetor real produz uma permutação válida
- Permite usar GA contínuo padrão sem operadores especializados para o domínio
- Validado em scheduling (single machine, resource-constrained project)

## Relevância para o TCC

Citado na Seção 2.6 (PSO) como referência para random keys. No TCC, o PSO usa codificação por random keys: cada partícula mantém um vetor real cuja ordenação gera a permutação da rota. Bean é a referência canônica para essa abordagem, que permite usar a equação contínua clássica do PSO sem modificações. A classificação `condicional` reflete que o PDF não foi lido integralmente, mas a contribuição do artigo é bem estabelecida na literatura.

## Métodos e Abordagens

- Representação: vetor de reais (chaves), um por item na sequência
- Decodificação: ordenação crescente/decrescente das chaves → permutação
- Vantagem: operadores contínuos padrão mantêm factibilidade automaticamente
- Aplicação: problemas de sequenciamento onde a ordem é a decisão principal

## Conexões

- [[genetic-algorithms]]
- [[particle-swarm]] — PSO do TCC usa random keys
- [[clerc2000discretepso]] — alternativa discreta (swap operators)
- [[holland1975adaptation]] — GA, fundação teórica
- [[goldberg1989genetic]] — textbook de referência
