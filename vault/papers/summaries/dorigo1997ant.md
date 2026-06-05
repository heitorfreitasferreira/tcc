# dorigo1997ant — Resumo

**Título:** Ant colonies for the travelling salesman problem  
**Autores:** Marco Dorigo, Luca Maria Gambardella  
**Periódico:** *BioSystems* 43 (1997), 73–81  
**DOI:** (não informado no PDF; buscar via CrossRef)

---

## 1. Problema e motivação

O artigo propõe um algoritmo de otimização inspirado no comportamento coletivo de formigas reais para resolver o problema do caixeiro-viajante (TSP). A motivação vem da observação biológica de que colônias de formigas encontram o caminho mais curto entre o ninho e uma fonte de alimento por meio de um mecanismo de trilha de feromônio e retroalimentação positiva (autocatálise). O objetivo é transferir esse princípio emergente para um sistema artificial que gere boas soluções para instâncias simétricas e assimétricas do TSP.

## 2. Método central (Ant Colony System — ACS)

- **Agentes artificiais (formigas):** cada formiga constrói um tour completo movendo-se probabilisticamente entre cidades, escolhendo arestas com base em dois fatores: quantidade de feromônio acumulada e uma heurística gulosa (inverso da distância).
- **Regra de ação pseudoaleatória-proporcional:** com probabilidade *q₀* a formiga explora a melhor aresta conhecida; com probabilidade *1 − q₀* realiza uma exploração estocástica enviesada por feromônio e distância (Eq. 1).
- **Atualização local de feromônio:** cada vez que uma aresta é percorrida, sua trilha é reduzida pela fórmula *τ(r,s) ← (1−α)·τ(r,s) + α·τ₀*, incentivando diversificação e evitando convergência prematura.
- **Atualização global de feromônio:** ao final de cada iteração, apenas a melhor formiga deposita feromônio nas arestas de seu tour, com quantidade inversamente proporcional ao comprimento do tour (*τ ← (1−α)·τ + α / L_melhor*), reforçando as soluções de melhor qualidade.
- **Lista de candidatos:** para instâncias maiores, cada formiga restringe sua escolha às *cl* cidades mais próximas (tipicamente *cl = 20*), reduzindo o espaço de busca e acelerando a construção de tours.
- **Memória de trabalho:** cada formiga mantém uma lista tabu das cidades já visitadas para garantir tours válidos (cada cidade visitada exatamente uma vez).

## 3. Principais resultados

- **Comparação com métodos da época (Tabelas 1–3):** ACS iguala ou supera Simulated Annealing (SA), Programação Evolutiva (EP), Algoritmos Genéticos (GA), Elastic Net (EN) e Self-Organizing Map (SOM) em instâncias simétricas de 30 a 100 cidades (Oliver30, Eil50, Eil75, KroA100), tanto em qualidade de solução quanto em número de tours gerados.
- **Instâncias maiores (Tabela 4):** resultados para d198 (0,68% de erro), pcb442 (0,96%), att532 (1,67%), rat783 (2,37%) e fl1577 (3,27%), todos com *cl = 20* e obtidos em tempos de CPU viáveis (Sun Sparc-server 50 MHz).
- **TSP assimétrico (Tabela 7):** ACS encontra a solução ótima para 43X2 (43 cidades) em 220 segundos (Pentium PC), problema cuja solução exata demandava >32 horas com os melhores métodos exatos da época (branch-and-cut). Para ry48p também encontra o ótimo (14.422).
- **Efeito sinérgico da comunicação (Fig. 2):** formigas que se comunicam via feromônio encontram a solução ótima mais rapidamente do que formigas independentes, com ganho que aumenta com o número de agentes.
- **Distribuição de tempos de primeira descoberta (Fig. 3):** a comunicação altera favoravelmente a distribuição do tempo até encontrar a primeira solução ótima.
- **Lista de candidatos (Tabelas 5–6):** listas curtas (*cl = 20*) melhoram tanto o desempenho médio quanto o melhor resultado, além de reduzirem o tempo de CPU.

## 4. Pontos fortes e limitações

**Pontos fortes:**  
- Primeira demonstração convincente de que um algoritmo baseado em colônia de formigas compete com meta-heurísticas estabelecidas (SA, GA, EP).  
- Aplicável tanto a TSP simétrico quanto assimétrico, com desempenho particularmente notável no caso assimétrico.  
- Formulação clara como sistema de reinforcement learning distribuído.  
- Paralelizável e extensível a outros problemas combinatórios (ex.: QAP).

**Limitações:**  
- Para TSP simétrico, ACS não compete com heurísticas especializadas como Lin-Kernighan.  
- Os experimentos usam instâncias moderadas (≤1577 cidades); escalabilidade para dezenas de milhares de cidades não é avaliada.  
- Parâmetros (α, β, q₀, τ₀) são fixados heuristicamente; não há estudo sistemático de sensibilidade.  
- A validação estatística é limitada (15 trials por configuração, sem intervalos de confiança ou testes de hipótese).

## 5. Implicação prática para pesquisadores

O ACS fornece um *framework* geral para problemas de otimização combinatória representáveis como busca em grafo: basta definir uma representação em grafo do problema e uma heurística de distância entre nós. A comunicação indireta via feromônio (estigmergia) oferece um mecanismo simples e eficaz de cooperação entre agentes, e a técnica é diretamente paralelizável. O artigo estabelece as bases do que viria a ser a meta-heurística *Ant Colony Optimization* (ACO), aplicável a roteamento, escalonamento, atribuição quadrática e outros problemas NP-difíceis.
