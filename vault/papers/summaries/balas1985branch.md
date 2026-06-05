# Branch and Bound Methods for the Traveling Salesman Problem

**Autores:** Egon Balas e Paolo Toth  
**Ano:** 1983 (Management Science Research Report No. MSRR 488, Carnegie-Mellon University)  
**Tipo:** Survey / revisão do estado da arte  

## 1. Problema e motivação

O artigo revisa métodos enumerativos (branch and bound) para o Problema do Caixeiro Viajante (TSP), que desde o trabalho de Eastman (1958) serviu como campo de teste para o desenvolvimento de técnicas de otimização discreta. A motivação é organizar e comparar sistematicamente as principais classes de algoritmos branch and bound, classificando-as pela relaxação utilizada, e fornecer evidência computacional sobre seu desempenho empírico.

## 2. Abordagem central

- **Relaxação I — Assignment Problem (AP) com função de custo original:** resolve o AP removendo as restrições de eliminação de sub-rotas. Para o TSP assimétrico, v(AP) atinge em média 99,2% de v(TSP) em instâncias aleatórias, mas no caso simétrico cai para ~82% devido à alta frequência de sub-rotas de comprimento 2.
- **Relaxação II — 1-Tree com função Lagrangeana:** usada para o TSP simétrico. A 1-tree básica é fraca (~63% de v(TSP)), mas a versão Lagrangeana com otimização por subgradiente (Held e Karp, 1971) eleva a qualidade da cota para ~99,7% de v(TSP).
- **Relaxação III — AP com função Lagrangeana:** desenvolvida por Balas e Christofides (1981) para o TSP assimétrico. Emprega seis procedimentos sequenciais de bounding que adicionam multiplicadores Lagrangeanos a desigualdades violadas (cortes direcionados, desigualdades de eliminação de sub-rotas, desigualdades de ponto de articulação), atingindo em média 99,5% de v(TSP).
- **Regras de ramificação:** são discutidas nove regras (BR1–BR9), desde a clássica de Little et al. (1963) até regras baseadas em disjunções de cortes condicionais (BR9). As regras BR3 (partição por sub-rota mínima) e BR9 são usadas intermitentemente no algoritmo BC.
- **Seleção de subproblemas:** comparam-se as estratégias depth-first (LIFO) e breadth-first (melhor cota), com evidência de que breadth-first reduz o número de nós mas consome mais memória.

## 3. Principais resultados

- **Comparação de códigos para TSP assimétrico (três algoritmos):** SST (Smith, Srinivasan, Thompson — AP + penalidade), CT (Carpaneto e Toth — AP com pós-otimização), e BC (Balas e Christofides — AP Lagrangeano). O BC gera consistentemente menos nós da árvore de busca para n ≤ 180, mas para n ≥ 200 o CT reduz a contagem de nós devido à estratégia breadth-first.
- **Comparação para TSP simétrico:** os códigos ST (Smith e Thompson) e VJ (Volgenant e Jonker), ambos baseados em 1-tree Lagrangeana com subgradiente. O VJ mostra menor número de iterações de subgradiente e menor tempo de computação.
- **Análise estatística de crescimento:** o esforço computacional em função de n (até n=325) pode ser descrito quase igualmente bem por funções polinomiais (grau 1,4–4,4), superpolinomiais, ou exponenciais com base muito próxima de 1 (e.g., 1,012^n), sem evidência conclusiva para preferir um modelo.
- **Redução de arcos:** no algoritmo BC, a etapa de variable fixing remove em média 96–98% dos arcos antes da ramificação.

## 4. Pontos fortes e limitações

**Pontos fortes:**
- Cobertura abrangente e bem organizada das três principais famílias de relaxações.
- Inclui experimentos computacionais próprios (400 problemas aleatórios para AP, 140 para 1-tree) além da compilação de resultados da literatura.
- Análise estatística rigorosa do crescimento do esforço computacional com regressão de três modelos funcionais.
- Exemplos numéricos detalhados e exercícios ao final do capítulo.

**Limitações:**
- A relaxação AP Lagrangeana (BC) é fraca para o TSP simétrico (~96% de v(TSP)), limitando sua aplicabilidade a problemas assimétricos.
- O survey não cobre métodos heurísticos ou aproximativos, focando exclusivamente em solução exata por enumeração implícita.
- O intervalo de n testado (até 325) é modesto para os padrões atuais, e as conclusões sobre crescimento assintótico permanecem inconclusivas.
- A relaxação 2-matching Lagrangeana é discutida como promissora, mas à época carecia de implementação eficiente do algoritmo de 2-matching ponderado.

## 5. Implicações práticas para pesquisa

- Para TSP assimétrico, a relaxação AP com função Lagrangeana (procedimentos de Balas-Christofides) oferece as cotas mais fortes e deve ser a base de algoritmos exatos.
- Para TSP simétrico, a relaxação 1-tree com subgradiente continua sendo a abordagem dominante; o uso de combinação convexa de subgradientes (VJ) melhora a convergência.
- A escolha entre depth-first e breadth-first envolve um trade-off entre consumo de memória e tamanho da árvore; algoritmos híbridos que alternam entre as duas estratégias tendem a ser mais robustos.
- O fato de que funções exponenciais com base ~1,01 descrevem bem o crescimento empírico sugere que instâncias moderadas (n ≤ 300) são tratáveis na prática, apesar da NP-completude teórica.
