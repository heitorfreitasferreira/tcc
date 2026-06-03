---
title: Justificativa do Lower Bound — Análise e Decisão
tags: [writing, lower-bound, analise, monografia]
status: pronto-para-revisao
---

# Justificativa do Método de Lower Bound

> Pseudo-subseção para o Capítulo 4 (Experimentos). Contexto: avaliação de qualidade para instâncias grandes (n ≥ 15) onde o ótimo é desconhecido.

---

## 4.X Lower Bound para Instâncias Grandes

O problema central que esta seção endereça é: na ausência do ótimo global para instâncias com n ≥ 15, como medir a qualidade absoluta das soluções encontradas por GA, PSO e ACO?

A resposta padrão na literatura de TSP é usar um **lower bound** — um valor LB tal que LB ≤ OPT para toda instância — e expressar a qualidade como gap percentual:

$$ \text{gap} = \frac{\text{makespan} - \text{LB}}{\text{LB}} \times 100\% $$

Quanto menor o gap, mais próxima a solução está do ótimo. Para que essa métrica seja informativa, o lower bound precisa ser **apertado** (próximo do ótimo) e **computacionalmente tratável** para instâncias grandes.

### 4.X.1 Por que métodos clássicos não se aplicam diretamente

O TSP-SD-ATP é definido por um tensor de custos tridimensional `G[prev][curr][next]`, onde cada entrada incorpora a distância Euclidiana percorrida entre `curr` e `next` somada a uma penalidade angular dependente do segmento anterior `prev→curr`. Essa estrutura sequência-dependente impede a aplicação direta dos dois lower bounds mais consagrados na literatura.

O **Held-Karp bound** (Held e Karp, 1970; 1971) — baseado na relaxação Lagrangiana de 1-trees — é o padrão-ouro para o TSP simétrico, com gap empírico inferior a 0,8% para instâncias Euclideanas aleatórias (Johnson et al., 1996). Contudo, o método assume custos de aresta simétricos representados por uma matriz bidimensional `c[i][j]`. A 1-tree é construída a partir de uma árvore geradora mínima (MST), estrutura que não admite a dimensão adicional de dependência do nó anterior. O mesmo vale para a relaxação de **Assignment Problem (AP)** (Balas e Toth, 1985): embora seja a relaxação padrão para o ATSP e forneça bounds excepcionalmente apertados para custos aleatórios (Karp, 1979), ela também opera sobre uma matriz 2D.

Uma segunda linha de abordagens opera nativamente sobre estruturas com dependência de sequência. **Diagramas de Decisão Multivalorados (MDDs)** com largura limitada, propostos por Kinable, Cire e van Hoeve (2017) para o Time-Dependent TSP, permitem representar o espaço de estados do problema de forma relaxada. O estado `(S, prev, curr)` captura exatamente a semântica do tensor 3D, e o merge de estados com o mesmo par `(prev, curr)` — descartando o conjunto visitado completo — produz uma relaxação controlável pelo parâmetro de largura `W`. De forma similar, a **ng-path relaxation** (Lera-Romero, Miranda-Bront e Soulignac, 2020) reduz o espaço de estados mantendo apenas um subconjunto limitado de nós visitados por estado. Ambas as abordagens, entretanto, exigem implementações significativamente mais complexas que os métodos clássicos e, no caso do MDD, o tuning de `W` para equilibrar precisão e custo computacional não é trivial.

### 4.X.2 Abordagem adotada: redução do tensor 3D com relaxação AP

A estratégia escolhida neste trabalho combina simplicidade de implementação com respaldo teórico sólido. Dado o tensor `cost[prev][curr][next]`, define-se uma matriz de custos reduzida:

$$ c'[j][k] = \min_{i \in V} \text{cost}[i][j][k] $$

Para qualquer tour `π = (π₀, π₁, …, π_{n-1})`, vale que:

$$ \sum_{t} \text{cost}[π_{t-1}][π_t][π_{t+1}] \geq \sum_{t} c'[π_t][π_{t+1}] $$

portanto o custo ótimo do TSP-SD-ATP é superestimado (ou igualado) pelo custo ótimo do ATSP definido sobre `c'`. Consequentemente, qualquer lower bound para o ATSP sobre `c'` é também um lower bound válido para o TSP-SD-ATP. A perda de informação — o valor exato da penalidade angular para cada par `(prev, curr, next)` — é trocada pela capacidade de aplicar métodos de otimização combinatória clássicos e eficientes.

Sobre a matriz reduzida `c'`, aplicamos a **relaxação do Assignment Problem (AP)**. O AP relaxa as restrições de conectividade do ATSP (eliminação de subrotas), mantendo apenas as restrições de grau: cada vértice deve ter exatamente uma aresta de entrada e uma de saída. Formalmente:

$$ \min \sum_{j,k} c'[j][k] \, x_{jk} $$
sujeito a:
$$ \sum_{k} x_{jk} = 1 \quad \forall j $$
$$ \sum_{j} x_{jk} = 1 \quad \forall k $$
$$ x_{jk} \in \{0,1\} $$

A solução ótima deste problema — um conjunto de subrotas disjuntas que cobrem todos os vértices — pode ser encontrada pelo algoritmo Hungaro (Kuhn, 1955; Munkres, 1957) em tempo O(n³). O valor ótimo AP é um lower bound para o ATSP e, por construção, para o TSP-SD-ATP.

A escolha do AP sobre a matriz reduzida se justifica por quatro razões:

1. **Tratabilidade**: O(n³) é viável para as instâncias deste estudo (até n = 100 nós, o algoritmo executa em milissegundos em hardware moderno).
2. **Validade formal**: O bound é garantidamente um limitante inferior, independentemente da qualidade da redução.
3. **Qualidade do bound**: Para ATSP com custos não-negativos, o AP bound é tipicamente apertado — Karp (1979) demonstrou que, para matrizes de custo aleatórias U[0,1], o gap entre o AP e o ATSP ótimo é o(1). Embora os custos do TSP-SD-ATP não sejam uniformes, a redução por mínimo tende a produzir matrizes esparsas e assimétricas nas quais o AP bound se mantém informativo.
4. **Implementação**: O algoritmo Hungaro não requer bibliotecas externas — é implementável em Go puro com estruturas de dados padrão.

### 4.X.3 Limitações e refinamentos possíveis

A redução `c'[j][k] = min_i cost[i][j][k]` subestima o custo real de cada transição, e essa subestimativa acumula-se ao longo do tour. Em instâncias onde a penalidade angular é alta e varia significativamente com o nó anterior, o bound pode ser frouxo. Para mitigar essa limitação, duas linhas de refinamento podem ser exploradas em trabalhos futuros:

**Refinamento por MDD**: Substituir a relaxação AP por um MDD com largura limitada W, operando diretamente sobre o tensor 3D. O estado `(prev, curr)` captura a dependência angular, e a largura W controla o trade-off precisão-custo. Espera-se que mesmo valores moderados de W (e.g., W = 100) produzam bounds significativamente mais apertados que o AP, como demonstrado por Kinable et al. (2017) para problemas de sequenciamento posição-dependente.

**Refinamento por additive bounding**: Aplicar o procedimento aditivo de Fischetti e Toth (1992) sobre a matriz reduzida, combinando AP com arborescências. O additive bounding tipicamente adiciona 1-3% de aperto sobre o bound AP isolado.

**Refinamento por subgradiente HK**: Sobre a matriz reduzida simetrizada, aplicar o método de subgradiente de Held-Karp para obter o bound de 1-tree, que é mais apertado que o AP para TSP simétricos. Johnson et al. (1996) demonstram que esse bound fica a menos de 1% do ótimo para instâncias Euclideanas.

### 4.X.4 Resumo da decisão

| Critério | AP (reduzido) | MDD nativo | Held-Karp (reduzido) |
|----------|:-------------:|:----------:|:--------------------:|
| Opera em 3D | via redução | nativo | via redução |
| Complexidade | O(n³) | O(W·n²) | O(K·n²) |
| Implementação | Simples (Go puro) | Complexa | Moderada |
| Força do bound | Forte | Muito forte (W-dependente) | Muito forte |
| Validade formal | Garantida | Garantida | Garantida |

Para este estudo, adotamos a relaxação AP sobre a matriz reduzida como método primário de lower bound. A implementação em Go puro é viável, o bound é formalmente válido, e o algoritmo Hungaro em O(n³) escala para todas as instâncias analisadas. A análise dos gaps obtidos, confrontados com os ótimos conhecidos do brute-force para `10a..15c`, permite calibrar a interpretação dos resultados para instâncias maiores. Os refinamentos via MDD, additive bounding ou subgradiente HK são identificados como direções de trabalhos futuros, já que o bound AP auditado é válido, mas frouxo para o TSP-SD-ATP.

---

## Referências

- Balas, E., Toth, P. (1985). Branch and bound methods. In Lawler et al., *The Traveling Salesman Problem*. Wiley.
- Fischetti, M., Toth, P. (1992). An additive bounding procedure for the ATSP. *Mathematical Programming*, 53: 173–197.
- Held, M., Karp, R. M. (1970). The traveling-salesman problem and minimum spanning trees. *Operations Research*, 18: 1138–1162.
- Held, M., Karp, R. M. (1971). The traveling-salesman problem and minimum spanning trees: Part II. *Mathematical Programming*, 1: 6–25.
- Johnson, D. S., McGeoch, L. A., Rothberg, E. E., Schreiber, R. (1996). Asymptotic experimental analysis for the Held-Karp TSP bound. *Proceedings of SODA*.
- Karp, R. M. (1979). A patching algorithm for the nonsymmetric TSP. *SIAM Journal on Computing*, 8: 561–573.
- Kinable, J., Cire, A. A., van Hoeve, W.-J. (2017). Hybrid optimization methods for time-dependent sequencing problems. *European Journal of Operational Research*, 259: 887–897.
- Kuhn, H. W. (1955). The Hungarian method for the assignment problem. *Naval Research Logistics Quarterly*, 2: 83–97.
- Lera-Romero, G., Miranda-Bront, J. J., Soulignac, F. J. (2020). A dynamic programming algorithm for the TDTSP with time windows. *4OR*, 18: 399–428.
- Munkres, J. (1957). Algorithms for the assignment and transportation problems. *Journal of the Society for Industrial and Applied Mathematics*, 5: 32–38.
