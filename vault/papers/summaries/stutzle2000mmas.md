---
title: "MAX-MIN Ant System"
authors: "Thomas Stützle, Holger H. Hoos"
year: 2000
journal: "Future Generation Computer Systems"
doi: ""
aliases: ["MMAS"]
tags: ["aco", "tsp", "qap", "metaheuristic", "optimization"]
---

## 1. Problem and Motivation

Ant System (AS), o primeiro algoritmo de Ant Colony Optimization (ACO), mostrou-se viável para problemas de otimização combinatória NP-difíceis, mas seu desempenho era insatisfatório em instâncias grandes de benchmarks como o TSP e o QAP quando comparado a algoritmos estado-da-arte. Os autores propõem o MAX-MIN Ant System (MMAS) para alcançar desempenho competitivo por meio de maior exploração das melhores soluções encontradas e mecanismos eficazes para evitar convergência prematura.

## 2. Core Method or Approach

- **Atualização elitista de feromônio**: apenas uma única formiga (a *iteration-best* ou a *global-best*) deposita feromônio a cada iteração, concentrando o reforço nas soluções de mais alta qualidade.
- **Limites explícitos de feromônio** $\tau_{\min} \leq \tau_{ij} \leq \tau_{\max}$: os limites inferior e superior evitam estagnação da busca ao impedir que diferenças relativas entre trilhas se tornem extremas. $\tau_{\max}$ é derivado assintoticamente do valor da melhor solução global, e $\tau_{\min}$ é calculado a partir da probabilidade desejada de reconstruir a melhor solução quando o algoritmo converge.
- **Inicialização das trilhas em $\tau_{\max}$**: favorece exploração nas primeiras iterações, pois a evaporação reduz lentamente as trilhas, mantendo diversidade por mais tempo em comparação com inicialização em $\tau_{\min}$.
- **Smoothing de feromônio (PTS)**: mecanismo opcional que eleva trilhas proporcionais à sua diferença para $\tau_{\max}$, facilitando escape de mínimos locais sem perder completamente a informação acumulada.
- **Estratégia mista e reinitialização**: alterna entre *iteration-best* e *global-best* para atualização com frequência crescente de *global-best* ao longo da execução; versões +ri (reinitialization) e +rs (restart) introduzem diversificação adicional ao resetar ou suavizar trilhas quando a convergência é detectada.

## 3. Main Results

- **TSP simétrico (sem busca local)**: MMAS superou consistentemente AS, AS com elitismo, AS rank-based e ACS nos benchmarks do TSPLIB (eil51, kroA100, d198, lin318, ry48p, ft70, kro124p, ftv170), com desvios médios em relação ao ótimo significativamente menores (ex.: d198: 1,09% vs 4,40% do AS).
- **TSP com busca local (3-opt e LK)**: MMAS+rs alcançou os melhores resultados médios; com LK, atingiu solução ótima em todas as execuções para lin318 e pcb442, e desvios inferiores a 0,01% para att532, rat783 e pcb1173.
- **ATSP**: MMAS+rs resolveu ft70 e ftv170 à otimalidade em todas as execuções; nos demais (ry48p, kro124p), desempenho comparável ao ACS.
- **QAP**: MMAS está entre os melhores algoritmos disponíveis. Para instâncias reais e *real-life-like* (classes iii e iv do QAPLIB), MMAS com 2-opt simples encontra as melhores soluções conhecidas em quase todas as execuções (ex.: bur26a, bur26c resolvidos em 100% das execuções, média de 3,8 s). Para instâncias aleatórias uniformes (classe i), GH e Ro-TS superam MMAS.
- **Análise FDC**: Correlação significativa entre qualidade da solução e distância ao ótimo global no TSP ($\rho \approx 0{,}45\text{--}0{,}63$) e em QAP reais ($\rho \approx 0{,}25\text{--}0{,}57$), mas correlação praticamente nula ($\rho \approx 0{,}02$) em QAP aleatórios uniformes — explicando a eficácia variável do MMAS entre classes de instâncias.

## 4. Strengths and Limitations

**Strengths**:
- Derivação matemática dos limites de feromônio com base na probabilidade de reconstrução da melhor solução, fornecendo fundamentação teórica para os parâmetros.
- Validação empírica extensa cobrindo TSP simétrico, ATSP e QAP com múltiplas variantes e busca local.
- Conexão explícita entre análise da paisagem de busca (FDC) e projeto algorítmico, motivando as escolhas de design.
- Robusto a variações razoáveis de parâmetros; os valores default funcionam bem em ampla gama de instâncias.

**Limitations**:
- Desempenho inferior em instâncias QAP com baixa correlação *fitness-distance* (classe aleatória uniforme), onde Ro-TS e GH são superiores.
- Tempos de execução elevados quando combinado com LK para TSP, ficando atrás de *genetic local search* e ILK em *runtime* competitivo.
- Atualização de feromônio requer $O(n^2)$ operações (mitigada com *candidate lists* para $O(n)$), mas ainda custosa para instâncias muito grandes.
- Mecanismos de diversificação (+ri, +rs) e PTS são heurísticos, sem ajuste fino exaustivo dos parâmetros associados.

## 5. Practical Takeaway for Researchers

- Limites de feromônio são um mecanismo simples e eficaz para prevenir convergência prematura em ACO; a derivação baseada em probabilidade de reconstrução (Equações 7–11) fornece um ponto de partida sólido para qualquer aplicação.
- Análise FDC deve preceder a aplicação de ACO a um novo problema: se não houver correlação significativa entre *fitness* e distância ao ótimo, algoritmos baseados em reforço de soluções de elite provavelmente terão desempenho insatisfatório.
- Estratégias mistas que transitam de *iteration-best* para *global-best* ao longo da execução equilibram exploração inicial e intensificação final sem necessidade de *tuning* complexo.
- Híbridos com busca local são essenciais para desempenho estado-da-arte; MMAS sem busca local é competitivo apenas contra outros ACO puros, não contra algoritmos especializados para TSP/QAP.
