# vanhove2012route

**Vanhove, S. & Fack, V.** (2012). Route planning with turn restrictions: A computational experiment. *Operations Research Letters*, 40(5), 342–348. doi:10.1016/j.orl.2012.06.001

## 1. Problema e motivação

Redes viárias reais contêm restrições de conversão (custos adicionais ou proibições de manobra) que algoritmos padrão de caminho mínimo (e.g., Dijkstra) ignoram. Embora múltiplas abordagens tenham sido propostas para incorporar tais restrições, não havia até então um estudo comparativo sistemático que determinasse qual método é mais adequado em diferentes cenários. O trabalho preenche essa lacuna por meio de experimentos computacionais com redes europeias reais.

## 2. Método ou abordagem central

- **Node splitting** (*graph-transforming*): cada nó com restrições é desmembrado em múltiplos nós (um por arco incidente); curvas legais tornam-se arcos entre os nós desmembrados, e curvas proibidas são omitidas. Nós sem restrições permanecem intactos.
- **Line graph** (*graph-transforming*): o grafo inteiro é transformado de modo que cada nó da linha representa um arco do grafo original, e cada arco da linha representa uma curva legal. O peso de um arco na linha soma o custo da curva ao peso do arco destino.
- **Direct method** (*algorithm-modifying*): modifica o algoritmo de Dijkstra para rotular arcos em vez de nós, permitindo caminhos com ciclos nodais (mas sem repetir arcos). Estendeu-se o método original (que tratava apenas proibições) para suportar também custos de conversão.
- **Lookup tables** opcionais foram avaliadas para acelerar a criação de nós virtuais e a reconversão de caminhos nas abordagens transformadoras.

## 3. Principais resultados

- **Dados sintéticos** (rede belga: 458.403 nós, 1.085.076 arcos, 2.922.504 curvas; custos/proibições aleatórias em percentuais variados):
  - *Custos de conversão*: método direto mais eficiente em memória até ~25% de curvas com custo; acima disso, *line graph* (sem lookup) é preferível. Em tempo de consulta, direto e *line graph* empatam e superam *node splitting*.
  - *Proibições de conversão*: *node splitting* é o mais rápido para <5% de curvas proibidas; a partir de 5%, o método direto assume a liderança. O método direto é sempre o mais econômico em memória para proibições.
- **Dados reais Navteq** (Luxemburgo, Bélgica, Holanda, Paris; taxas de proibição entre 0,04% e 0,57%): *node splitting* é o mais rápido e eficiente em memória — confirmando que cenários reais típicos favorecem esse método. Não houve diferença relevante entre rede urbana (Paris) e redes nacionais.
- **Lookup tables**: acrescentam até o dobro de memória com ganho de tempo insignificante (< 1 ms nas etapas auxiliares, exceto ~17 ms para criação de nós virtuais no *line graph* sem lookup), não se justificando na prática.
- **Complexidades**: *line graph* O(t_l log m); direto O(t + t_l log m); *node splitting* O((m + t_l) log(n + t_r d_max)).

## 4. Pontos fortes e limitações

- **Pontos fortes**: primeira comparação experimental abrangente das três famílias de métodos; guideline claro e acionável para aplicações reais; validação em redes reais (Navteq) com diferentes escalas (urbana e nacional); análise separada de custos e proibições.
- **Limitações**: implementação datada (Java 1.6, 2 GB RAM); apenas Dijkstra como algoritmo base (embora os autores argumentem que o comportamento se generaliza); ausência de datasets reais com custos de conversão na época; restrições aleatórias nos testes sintéticos podem não refletir padrões espaciais reais de proibições/custos.

## 5. Aplicação prática para pesquisadores

A diretriz empírica oferecida é diretamente utilizável: (a) para **custos de conversão**, use o método direto se ≤25% das curvas têm custo; caso contrário, use *line graph*; (b) para **proibições de conversão**, use *node splitting* se <5% das curvas são proibidas (caso típico real); acima disso, prefira o método direto. Lookup tables não valem o custo de memória. Para modelagem de rotas de drones com restrições de manobra (contexto TSP/rTSP), o *node splitting* tende a ser a escolha mais natural, dado que proibições de conversão geralmente são esparsas.
