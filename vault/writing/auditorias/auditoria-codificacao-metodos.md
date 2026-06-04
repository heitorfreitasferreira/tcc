---
tags: [writing, auditoria, codigo, metodologia]
created: 2026-06-02
updated: 2026-06-02
---

# Auditoria P9 — Codificação dos Métodos

Status: `concluida-com-limitacao`

## Objetivo

Avaliar se a comparação entre GA, PSO e ACO pode ser interpretada como comparação direta entre metaheurísticas ou se parte do resultado pode decorrer das codificações usadas para representar rotas no TSP-SD-ATP.

## Evidência no Código

| Método | Representação implementada | Evidência primária | Implicação |
|---|---|---|---|
| GA | Permutação explícita dos nós `1..n-1` | `src/optimization/ga/main.go:57-67` | Opera diretamente no espaço de rotas factíveis. |
| GA | Ordered Crossover (OX) e mutação por troca | `src/optimization/ga/main.go:141-201` | Preserva ordem relativa, mas não modela explicitamente triplas `prev,curr,next`. |
| PSO | Random keys: posição contínua ordenada para gerar permutação | `src/optimization/pso/particle.go:23-40` | Permite usar atualização contínua padrão, mas a ordenação não preserva adjacências de forma explícita. |
| PSO | Atualização contínua com inércia, componente cognitivo e social | `src/optimization/pso/particle.go:42-54` | Pequenas alterações em `x` podem não alterar a rota, enquanto cruzamentos de chaves podem mudar várias adjacências. |
| ACO | Construção incremental de rota com transições `prev,curr,next` | `src/optimization/aco/ant.go:42-84` | Alinha a busca ao tensor 3D usado na função objetivo. |
| ACO | Feromônio tridimensional `tau[prev][curr][next]` | `src/optimization/aco/main.go:21`, `src/optimization/aco/main.go:86-97` | Representa diretamente a dependência de sequência, com custo de memória e diluição de sinal. |

## Correção de Premissa

O roadmap antigo descrevia o GA como random keys. Essa descrição não corresponde ao código atual. O GA implementado usa permutação direta, OX e mutação swap. O PSO é o método que usa random keys/SPV por ordenação de posições contínuas. O ACO não codifica uma rota completa antes da avaliação; ele constrói a rota incrementalmente por transições probabilísticas condicionadas ao nó anterior e ao nó corrente.

## Interpretação Para a Monografia

A comparação experimental permanece válida como comparação entre as implementações avaliadas sob o mesmo conjunto de instâncias, população, iterações e sementes. Ela não deve ser apresentada como prova de superioridade intrínseca de uma família de algoritmos sobre outra. As três abordagens exploram o espaço de rotas por mecanismos distintos:

- o GA manipula permutações factíveis com recombinação e mutação;
- o PSO busca em um espaço contínuo e só depois decodifica a posição para uma permutação;
- o ACO constrói rotas por decisões locais condicionadas à sequência de três nós.

Essa diferença favorece interpretações metodológicas cautelosas. O desempenho inferior do PSO pode estar associado à codificação random keys, aos parâmetros fixos ou ao próprio mecanismo de atualização contínua aplicado a um problema de permutação com custo dependente de sequência. O bom desempenho qualitativo do ACO pode estar relacionado ao alinhamento entre feromônio tridimensional e tensor de custo, mas essa mesma representação aumenta a dimensionalidade do feromônio e torna o método mais caro.

## Texto Recomendado Para Limitações

> A comparação realizada deve ser lida como comparação entre implementações concretas dos métodos, não como ordenação universal entre GA, PSO e ACO. Embora os métodos tenham sido avaliados sob as mesmas instâncias, sementes, população e número de iterações, eles usam representações distintas: o GA manipula permutações diretamente, o PSO decodifica posições contínuas por random keys e o ACO constrói rotas por transições condicionadas a triplas de nós. Essa escolha é metodologicamente relevante porque o TSP-SD-ATP depende da sequência `prev,curr,next`; portanto, representações que preservam ou exploram adjacências podem se comportar de forma diferente de representações contínuas decodificadas por ordenação. Um estudo futuro deve comparar variantes alternativas, como PSO discreto com operadores de permutação e GA com operadores mais orientados a arestas, para isolar melhor o efeito da codificação.

## Decisão P9

P9 não exige nova implementação para a escrita atual da monografia. A ação correta é explicitar a ameaça à validade e evitar claims causais fortes sobre os algoritmos em abstrato. Uma implementação alternativa de PSO discreto ou random-key GA seria trabalho futuro, não pré-requisito para reportar os resultados atuais.

## Links

- [[claim-evidence-matrix]]
- [[glossario-monografia]]
- [[ga]]
- [[pso]]
- [[aco]]
- [[resultados]]
