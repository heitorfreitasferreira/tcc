# Resumo: Applegate *et al.* (2006) — *The Traveling Salesman Problem: A Computational Study*

## 1. Problema e motivação

O livro documenta o estado da arte computacional para o Problema do Caixeiro Viajante (TSP) simétrico em meados dos anos 2000. A motivação central é expor de forma sistemática as técnicas de *branch-and-cut*, planos de corte e heurísticas de busca local que levaram à solução exata de todo o repositório TSPLIB — incluindo instâncias com até 85.900 cidades — e à obtenção de limites superiores e inferiores de altíssima precisão para instâncias com milhões de cidades. O código Concorde, disponibilizado como *software* livre, é o artefato prático que acompanha e valida cada capítulo.

## 2. Abordagem central

- **Método de planos de corte (*cutting-plane method*)** sobre a relaxação linear do TSP: parte-se de restrições de grau (*degree constraints*) e adicionam-se iterativamente desigualdades válidas violadas pela solução fracionária corrente.
- **Famílias de cortes especializadas**: cortes de sub-rotas (*subtour elimination*) com estruturas PQ-tree; *blossom inequalities*; *comb inequalities* via reconhecimento de *consecutive ones* e *domino-parity*; cortes locais (*local cuts*) baseados em redução a instâncias pequenas via mapeamento linear; e metamorfoses de cortes (*tighten, teething, gluing*).
- ***Branch-and-cut***: quando os cortes já não fecham o *gap* de integralidade, o espaço é particionado com *strong branching* e *tentative branching* (ramificação em múltiplas vias, ex. 8-way, 32-way), gerando novas rodadas de cortes em cada nó da árvore de busca.
- **Heurísticas de busca de *tours***: *Chained Lin-Kernighan*, *tour merging* via decomposição em ramos (*branch decomposition*) com programação dinâmica sobre grafos esparsos, e a variante LKH de Helsgaun com movimentos *k-opt* (até *k*=10).
- **Infraestrutura de LP**: *core LP* com precificação dinâmica de arestas, *column generation*, e uso do *solver* CPLEX (simplex primal/dual) como motor da relaxação.

## 3. Principais resultados

- **TSPLIB completo resolvido**: todas as instâncias do repositório clássico, incluindo `pla85900` (85.900 cidades, VLSI, resolvida em abr/2006 com ~136 CPU-anos em clusters Opteron/Xeon) e `Sweden TSP` (24.978 cidades). A sequência de *branch-and-cut* sobre `pla85900` reduziu o *gap* da relaxação raiz de 0,08% até a otimalidade comprovada, com *tour* de Helsgaun confirmado como ótimo.
- **World TSP** (1.904.711 cidades): limite inferior via Concorde com cortes locais progressivos (*t* = 18) atingiu *gap* de 0,058% em relação ao melhor *tour* conhecido (Helsgaun, 7.516.043.366 m). ~98% do tempo de CPU consumido pelo *solver* LP em modo sequencial.
- **Instância euclidiana aleatória de 1M de cidades** (E1M.0): cortes locais até *t* = 28 produziram *gap* de 0,029%, com 308 CPU-dias em Alpha EV6 500 MHz.
- ***Tour merging*** com 25 *tours* e limite de 1 s no Chained Lin-Kernighan: qualidade média de 1,00004 (0,004% do ótimo) para instâncias < 1.000 cidades; 37/50 *trials* produziram o *tour* ótimo. Para instâncias entre 1.000–2.000 cidades, 95/130 *trials* ótimos.
- **Heurística LKH** de Helsgaun dominou os recordes de *tours* para coleções National e VLSI; híbridos genéticos (Nguyen *et al.*) e *tour morphing* mostraram-se competitivos para instâncias muito grandes.
- **Constante de Beardwood-Halton-Hammersley** (*β*): estimativas computacionais reportadas a partir de experimentos com instâncias euclidianas aleatórias.

## 4. Pontos fortes e limitações

**Pontos fortes:**
- Tratamento exaustivo e autocontido de toda a *pipeline* computacional, da teoria poliédrica à engenharia de *software*.
- Código Concorde como artefato reproduzível e reutilizável, com documentação pública.
- Demonstração empírica de que o TSP, apesar de NP-difícil, é tratável na prática para instâncias surpreendentemente grandes com a combinação certa de cortes, ramificação e heurísticas.
- Cobertura de aplicações reais (logística, sequenciamento genético, VLSI, astronomia) que ancora a relevância prática do problema.

**Limitações:**
- O gargalo do *solver* LP (simplex dual sequencial) domina o tempo de execução (~97–98%), e o livro reconhece a ausência de paralelização eficaz para essa etapa em 2006.
- A separação de cortes ainda exige repetidas rodadas de *branch-and-cut* para instâncias muito grandes, um processo caro e pouco previsível.
- Os algoritmos de decomposição em ramos (*branch-width*) são eficazes apenas quando a largura é pequena (< 20), e a heurística para encontrar boas decomposições não oferece garantias.
- A abordagem não cobre o TSP assimétrico, variantes com restrições temporais, ou problemas dinâmicos/estocásticos.

## 5. Implicações práticas para pesquisadores

- O Concorde estabelece uma *baseline* sólida para qualquer novo método exato ou heurístico aplicado ao TSP: todo resultado deve ser comparado contra seus limites inferiores e *tours* de referência.
- A taxonomia de cortes (sub-rotas, *combs*, *domino-parity*, cortes locais) oferece um cardápio de templates poliédricos reaproveitáveis em outros problemas de otimização combinatória com estrutura de ciclos (e.g., roteamento de veículos, *pickup-and-delivery*).
- A lição de engenharia é clara: o sucesso em escala depende tanto da sofisticação dos cortes quanto da gestão cuidadosa do tamanho do LP *core*, da precificação de arestas e do descarte inteligente de cortes inativos.
- Para instâncias além do alcance exato, a combinação de LKH + *tour merging* com decomposição em ramos permanece uma estratégia heurística de alto desempenho, com *gaps* de otimalidade frequentemente abaixo de 0,01%.
