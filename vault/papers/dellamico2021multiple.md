---
title: Modeling the Flying Sidekick Traveling Salesman Problem with Multiple Drones
authors:
- Dell'Amico, Mauro
- Montemanni, Roberto
- Novellani, Stefano
year: 2021
doi: 10.1002/net.22022
bibtex_key: dellamico2021multiple
bibtex-key: dellamico2021multiple
pdf: papers/pdfs/dellamico2021multiple.pdf
tags:
- area/drone-routing
- area/fstsp
- area/routing
- area/tsp
- evidencia/referencia
- metodo/milp
- status/resumo-lido
- tipo/paper
status: resumo-lido
rating: 4
type: paper
areas:
- drone-routing
- fstsp
- routing
methods:
- milp
role: revisao
reading_status: resumo-lido
validation_status: nao-validado
pdf_status: integro
chapters:
- fundamentacao
claim_support: []
aliases:
- MFSTSP
- Multiple Flying Sidekick Traveling Salesman Problem
---

## Resumo

Publicado na Networks (2021). Estuda o MFSTSP (Flying Sidekick TSP with Multiple Drones), no qual um caminhão e um conjunto de drones idênticos cooperam para entregar encomendas. Cada drone pode visitar um cliente por voo e deve retornar ao caminhão, com endurance limitada por bateria. Quando múltiplos drones são lançados ou recolhidos no mesmo nó, a ordem das operações torna-se relevante — ignorá-la pode gerar soluções infactíveis. O artigo propõe formulações MILP inéditas que incluem o escalonamento das operações dos drones e melhoram o tamanho das maiores instâncias resolvidas na literatura. Apresenta comparação entre formulações, entre soluções com um ou múltiplos drones, e entre variantes do modelo.

## Contribuições Principais

- Primeiras formulações MILP para FSTSP com múltiplos drones que modelam o scheduling das operações
- Comparação extensiva entre formulações e variantes do problema
- Demonstra vantagens do uso de múltiplos drones em relação ao single-drone
- Valid inequalities adaptadas para o caso multi-drone

## Evidência / Resultado Relevante

- O artigo compara quatro formulações MILP para MFSTSP e destaca a necessidade de modelar explicitamente a ordem de lançamento e recolhimento dos drones.
- O summary registra que ignorar o escalonamento de operações pode produzir soluções 5--20% mais baratas, porém infactíveis; validar no PDF antes de citar esse intervalo como número no texto final.
- O uso de múltiplos drones reduz o custo em relação ao caso com um drone, mas os experimentos são pequenos e voltados a entrega cooperativa caminhão-drone.

## Limitações de Uso

- O problema estudado é entrega com caminhão e múltiplos drones, não patrulha com um drone.
- As instâncias avaliadas são pequenas; a nota serve para contextualização de variantes, não para comparar com os experimentos do TCC.
- Claims numéricos sobre gaps e tempos devem ser revalidados no PDF antes de irem para a monografia.

## Relevância para o TCC

Citado na Seção 2.2 (TSP e variantes) como referência de drone routing com múltiplos veículos. Fornece contexto para a literatura de TSP com drones, embora o TCC trate de um único drone com penalidade angular — um problema mais restrito em número de veículos mas com complexidade adicional pela dependência de sequência.

## Conexões

- [[dellamico2022exact]] — artigo irmão com formulações exatas para FSTSP
- [[murray2015flying]] — FSTSP original (single drone)
- [[agatz2018optimization]] — TSP-D
- [[tsp]]
