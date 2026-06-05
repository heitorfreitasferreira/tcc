# almufti2025comparative

**Title:** Comparative Analysis of Metaheuristic Algorithms for Solving The Travelling Salesman Problems
**Authors:** Saman M. Almufti, Awaz Ahmed Shaban
**Venue:** International Journal of Scientific World, 11(2), 26–30, 2025
**DOI:** https://doi.org/10.14419/7fk7k945

## 1. Problem and motivation

The Traveling Salesman Problem (TSP) é NP-difícil e sua explosão fatorial inviabiliza métodos exatos para instâncias grandes. Embora muitas meta-heurísticas tenham sido propostas, faltam comparações empíricas rigorosas que as avaliem sob um mesmo protocolo experimental. O artigo preenche essa lacuna com um benchmark unificado de nove algoritmos sobre instâncias padronizadas do TSPLIB.

## 2. Core method or approach

- Nove meta-heurísticas populacionais ou de inteligência de enxame: ACO, LA (Lion Algorithm), CS (Cuckoo Search), GWO, VPS, SSO, CSO (Cat Swarm Optimization), BA (Bat Algorithm) e ABC.
- Cada algoritmo é executado 30 vezes independentes sobre três instâncias do TSPLIB: *berlin52* (52 cidades), *eil76* (76 cidades) e *pr1002* (1002 cidades).
- Métricas de avaliação: melhor custo encontrado, custo médio e desvio padrão.
- Análise complementar qualitativa de forças/fraquezas e uma tabela de comparação geral (inspiração, exploração, *exploitation*, convergência, sensibilidade paramétrica, complexidade).

## 3. Main results

- **ACO** obteve os melhores resultados médios em berlin52 (7542) e eil76 (538), com desvio padrão consistentemente baixo (10 e 4).
- **CSO** produziu o menor desvio padrão em pr1002 (800) e o melhor custo médio nessa instância (260900), indicando escalabilidade.
- **GWO** mostrou desempenho competitivo e equilibrado, com baixa variância nas três instâncias.
- **LA** e **VPS** apresentaram alta variância e pior desempenho em pr1002 (std 1600 e 1500, respectivamente).
- **CS** e **BA** tiveram desempenho moderado, adequados para instâncias médias.
- Conclusão dos autores: ACO, CSO e GWO são os mais balanceados e eficazes para o TSP.

## 4. Strengths and limitations

**Strengths:**
- Protocolo experimental uniforme (30 execuções, mesmas instâncias, mesmas métricas) que permite comparação justa.
- Cobertura ampla de paradigmas (forrageamento, hierarquia social, ecolocalização, acústica etc.).
- Tabelas-resumo claras e comparação qualitativa das características algorítmicas.

**Limitations:**
- Tamanho amostral pequeno de instâncias (apenas três) e ausência de instâncias intermediárias (ex.: ~200–500 cidades).
- Não reporta *timing* nem *convergence plots* — apenas valores numéricos agregados —, o que limita a análise de custo computacional e comportamento dinâmico.
- Nenhum teste estatístico formal (e.g., Wilcoxon, Friedman) para validar diferenças entre algoritmos.
- Não descreve a configuração de parâmetros de cada algoritmo; a sensibilidade paramétrica é discutida apenas qualitativamente.
- Publicado em periódico de baixo impacto; a revisão por pares pode ser menos rigorosa.

## 5. Practical takeaway for researchers

Para problemas do tipo TSP em escalas pequenas a médias, ACO e CSO oferecem a melhor relação entre qualidade de solução e robustez. GWO é uma alternativa com baixa variância, mas requer cuidado na calibração. O *framework* comparativo serve como ponto de partida, mas precisa ser estendido com testes estatísticos, curvas de convergência e mais instâncias antes de generalizar conclusões.
