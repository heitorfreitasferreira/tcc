```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "O referencial possui lacunas estruturais graves em duas frentes independentes: (a) ausência sistêmica de BibTeX para os artigos canônicos de lower bounds — Held-Karp 1970/1971, os trabalhos mais importantes sobre limites inferiores para TSP, não estão no arquivo .bib apesar de fichados no vault com rating 5 — e (b) descasamento entre a implementação ACO (Ant System básico, sem bounds de feromônio) e a literatura citada como justificativa (MMAS de Stützle 2000). Adicionalmente, o TSP-SD-ATP com tensor 3D gera um problema fundamentalmente assimétrico (ATSP) e com dependência de sequência (TDTSP), mas nenhuma dessas duas classes de problema possui qualquer referência na bibliografia.",
  "findings": [
    {
      "id": "f-council-001",
      "severity": "CRITICAL",
      "area": "lower-bounds",
      "title": "Held-Karp 1970 e 1971 ausentes do BibTeX",
      "detail": "Held-Karp (1970, 1971) são os artigos canônicos sobre lower bounds Lagrangianos para TSP — o bound de Held-Karp é equivalente à relaxação LP do TSP e tem gap empírico < 0.8%. Ambos estão fichados no vault (heldkarp1970traveling.md com rating 5, status 'lido'; heldkarp1971traveling.md com rating 5, status 'pendente'), mas NENHUM dos dois possui entrada BibTeX. O único artigo de lower bound com entrada BibTeX é Valenzuela (1997), que é uma implementação prática, não a referência canônica. Dos 10 artigos de lower bound no vault, apenas 1 tem BibTeX.",
      "fix": "Adicionar entradas BibTeX para heldkarp1970traveling e heldkarp1971traveling em abntex2-references.bib. Adicionar também johnson1996asymptotic (já lido, rating 5), kinable2017hybrid (já lido, rating 5), e righini2021efficient (já lido, rating 4).",
      "why": "O campo foi populado no vault mas a etapa de exportação para BibTeX foi omitida. A monografia ficará sem a referência canônica de lower bound para TSP, citando apenas um derivado prático (Valenzuela 1997).",
      "ref": "monografia/bib/abntex2-references.bib (sem entradas heldkarp*); vault/papers/heldkarp1970traveling.md (rating 5, lido)"
    },
    {
      "id": "f-council-002",
      "severity": "CRITICAL",
      "area": "aco",
      "title": "Implementação ACO é Ant System, não MMAS — descasamento com literatura citada",
      "detail": "O código em src/optimization/aco/ant.go:14-39 implementa atualização de feromônio no estilo Ant System (dorigo1996ant): TODAS as formigas depositam (linha 28, 'for _, ant := range ants'), sem seleção de iteration-best ou global-best. Não há bounds de feromônio [τ_min, τ_max], não há reinicialização por estagnação, e o feromônio inicial é 1.0 (não τ_max). A nota do vault para Stützle (2000) afirma incorretamente: 'A fórmula τ_max = 1/(ρ·L_best) e τ_min = τ_max / (2n) é implementada diretamente no código ACO do TCC' — isso é factualmente falso. A implementação é um Ant System 3D básico, não MMAS.",
      "fix": "Duas opções: (a) modificar a implementação para incluir os mecanismos de MMAS (bounds, elitismo de depósito, reinicialização) e manter a referência a Stützle (2000); ou (b) remover a afirmação de que MMAS está implementado e citar apenas Dorigo (1996, 1997) como base, documentando a variante 3D como contribuição própria sobre Ant System.",
      "why": "A nota do vault foi escrita assumindo conformidade com MMAS sem verificação do código-fonte. A discrepância invalida a seção de justificativa algorítmica do ACO na monografia.",
      "ref": "src/optimization/aco/ant.go:14-39 (updatePheromones sem bounds); vault/papers/stutzle2000mmas.md:55 (afirmação falsa sobre implementação)"
    },
    {
      "id": "f-council-003",
      "severity": "CRITICAL",
      "area": "atsp-tdtsp",
      "title": "Ausência total de referências sobre ATSP e TDTSP",
      "detail": "O TSP-SD-ATP com tensor 3D c(i,j,k) gera uma matriz de custos reduzida c'(j,k) = min_i c(i,j,k) que é fundamentalmente assimétrica (ATSP). Além disso, a dependência de sequência nos custos caracteriza o problema como um Time-Dependent TSP (TDTSP) ou Sequence-Dependent TSP. Nenhum artigo sobre ATSP ou TDTSP consta no BibTeX ou no vault — nem mesmo referências clássicas como o patching algorithm de Karp (1979, que está no vault mas sem BibTeX e focado em patching pós-Hungarian, não em ATSP per se) ou surveys de TDTSP. Kinable (2017) cobre TDTSP com decision diagrams mas está no vault sem BibTeX.",
      "fix": "Adicionar ao BibTeX e ao referencial: (a) uma survey de ATSP — sugestão: Öncan et al. (2009) 'A survey on the traveling salesman problem and its variants' ou Roberti & Toth (2012) 'Models and algorithms for the ATSP'; (b) uma referência sobre TDTSP — Gouveia & Voß (1995) 'A classification of formulations for the (time-dependent) traveling salesman problem' ou o próprio Kinable (2017) que já está fichado; (c) Karp (1979) 'A patching algorithm for the nonsymmetric traveling-salesman problem' que já está no vault.",
      "why": "Sem enquadramento em ATSP/TDTSP, o TSP-SD-ATP parece um problema inventado ad-hoc, quando na verdade é uma instância natural de classes bem estudadas de TSP com custos assimétricos e dependentes de sequência.",
      "ref": "monografia/bib/abntex2-references.bib (zero entradas ATSP/TDTSP); src/optimization/lowerbound/main.go:39-57 (c'(j,k) = min_i c(i,j,k) gera matriz assimétrica)"
    },
    {
      "id": "f-council-004",
      "severity": "HIGH",
      "area": "aco",
      "title": "Feromônio 3D é inovação sem lastro na literatura — precisa ser documentado como contribuição",
      "detail": "O ACO implementado usa feromônio tridimensional τ(i,j,k) (src/optimization/aco/main.go:86-97, pheromones [][][]float64). Na literatura canônica de ACO para TSP, o feromônio é sempre bidimensional τ(i,j) porque o custo depende apenas do par (cidade atual, próxima cidade). O feromônio 3D só se justifica em problemas com custos dependentes de triplas — que é exatamente o TSP-SD-ATP. Esta é uma adaptação não-trivial e inovadora, mas NÃO HÁ qualquer referência na literatura que use feromônio 3D para TSP. Os artigos canônicos (Dorigo 1996, 1997, 2004; Stützle 2000) usam todos τ(i,j).",
      "fix": "(a) Documentar explicitamente na monografia que o feromônio 3D é uma contribuição algorítmica original deste trabalho; (b) justificar a escolha citando a estrutura do TSP-SD-ATP (dependência de triplas) como motivação; (c) buscar na literatura de ACO para problemas com custos sequenciais (e.g., scheduling, routing with turn costs) se há precedentes de feromônio n-dimensional para contexto.",
      "why": "O feromônio 3D é a característica mais distintiva da implementação ACO deste TCC. Não documentá-lo como contribuição própria é perder a oportunidade de reivindicar novidade metodológica. Não referenciá-lo adequadamente na literatura existente é deixar uma lacuna de fundamentação.",
      "ref": "src/optimization/aco/main.go:86-97 (inicialização 3D); src/optimization/aco/ant.go:54 (probabilidade usa τ[prev][curr][next])"
    },
    {
      "id": "f-council-005",
      "severity": "HIGH",
      "area": "pso",
      "title": "Bean (1994) — artigo fundacional do PSO implementado — com rating 0 e PDF ilegível",
      "detail": "Bean (1994) 'Genetic Algorithms and Random Keys for Sequencing and Optimization' introduz o conceito de random keys — vetor real decodificado por ordenação para representar permutações. Este é exatamente o mecanismo usado pelo PSO implementado no TCC (random keys PSO). O artigo está no BibTeX e no vault, mas: (a) rating = 0 no vault (o mais baixo possível); (b) PDF contém apenas metadados INFORMS, corpo do artigo não foi extraído; (c) a nota do vault reconhece que 'não foi possível confirmar detalhes experimentais, pseudocódigo ou conclusões'. A fundamentação do PSO implementado repousa sobre um artigo cujo conteúdo o autor não conseguiu ler.",
      "fix": "Obter uma cópia legível do artigo (via acesso institucional à INFORMS, Sci-Hub, ou interlibrary loan). Após leitura completa, atualizar rating e nota do vault. Se o artigo permanecer inacessível, buscar referência alternativa que documente random keys para representação de permutações (ex: Snyder & Daskin 2006, 'A random-key genetic algorithm for the generalized traveling salesman problem').",
      "why": "Rating 0 para o artigo que fundamenta a estratégia de codificação do PSO é indefensável. A nota do vault é essencialmente uma inferência a partir do título — não constitui leitura acadêmica.",
      "ref": "vault/papers/bean1994genetic.md (rating: 0, status: lido-parcial, PDF ilegível); monografia/bib/abntex2-references.bib:83-92 (entrada BibTeX presente)"
    },
    {
      "id": "f-council-006",
      "severity": "HIGH",
      "area": "aco",
      "title": "Parâmetro --gama listado na documentação mas inexistente no código",
      "detail": "AGENTS-experiments.md:24 lista '--gama' como parâmetro do ACO. O código em src/optimization/aco/main.go:10-14 define Params com Alpha, Beta, Rho, Q — sem Gama. Nenhuma ocorrência de 'gama' ou 'gamma' no código-fonte Go. O parâmetro não é usado, não é inicializado, e provavelmente gera erro ou é ignorado se passado pela CLI.",
      "fix": "Remover '--gama' da documentação em AGENTS-experiments.md ou implementar o parâmetro no código (ex: controle de evaporação seletiva, peso de heurística adicional, ou outro uso documentado na literatura ACO).",
      "why": "Documentação desatualizada ou parâmetro planejado mas nunca implementado. Gera confusão sobre quais parâmetros o ACO efetivamente utiliza.",
      "ref": "AGENTS-experiments.md:24; src/optimization/aco/main.go:10-14"
    },
    {
      "id": "f-council-007",
      "severity": "HIGH",
      "area": "aco",
      "title": "Seção ACO desbalanceada: 13 artigos com 4 de ACO neural/RL irrelevantes para a comparação",
      "detail": "Dos 13 artigos ACO, 4 são de ACO neural/aprendizado por reforço (DeepACO 2023, PPACO 2024, GPACO 2025, NeuFACO 2025) — métodos que usam redes neurais para aprender políticas de construção ou parâmetros de feromônio. Estes artigos são de fronteira de pesquisa (2023-2025) e não têm relação com a implementação básica de Ant System 3D usada no TCC. Sua presença infla artificialmente a seção ACO sem contribuir para a fundamentação da implementação. Enquanto isso, GA tem 8 artigos e PSO tem apenas 5 (sendo 3 descartados).",
      "fix": "Mover DeepACO, PPACO, GPACO e NeuFACO para uma subseção de 'Tendências Recentes' ou 'Trabalhos Futuros', separando claramente da fundamentação do ACO implementado. Adicionar referências mais diretamente relevantes para balancear: (a) PSO: pelo menos +2 artigos sobre PSO para TSP com random keys ou representações contínuas; (b) GA: a nota do Nagata (2006) EAX está vazia e o PDF está corrompido — resolver ou substituir.",
      "why": "O desbalanceamento sugere viés de cobertura (ACO super-representado) e enfraquece a comparabilidade do referencial entre os três métodos. Além disso, artigos de ACO neural citados sem relação com a implementação configuram 'citation padding'.",
      "ref": "vault/papers/deepaco2023.md (pendente), ppaco2024.md (lido-parcial), gpaco2025.md (lido), neufaco2025.md (lido); contagem: 13 ACO vs 8 GA vs 5 PSO"
    },
    {
      "id": "f-council-008",
      "severity": "MEDIUM",
      "area": "baselines",
      "title": "Ausência de heurísticas construtivas e Simulated Annealing como baselines",
      "detail": "O referencial não inclui: (a) heurísticas construtivas clássicas para TSP — nearest neighbor, nearest insertion, farthest insertion, Christofides — que são baselines naturais e de baixo custo computacional; (b) Simulated Annealing (Kirkpatrick et al. 1983), que é um método de busca local estocástica amplamente usado como baseline em comparações de metaheurísticas para TSP; (c) 2-opt e 3-opt como operadores de busca local independentes (Lin 1965, além do Lin-Kernighan 1973 já citado). A única heurística não-metaheurística no referencial é o bruteforce (para n ≤ 15).",
      "fix": "Adicionar ao BibTeX e ao texto: (a) Johnson & McGeoch (1997) 'The traveling salesman problem: A case study in local optimization' como referência canônica de busca local e heurísticas construtivas para TSP; (b) Kirkpatrick, Gelatt & Vecchi (1983) 'Optimization by simulated annealing'; (c) Avaliar se a monografia deve incluir resultados experimentais com estes baselines ou apenas mencioná-los como referência contextual.",
      "why": "A ausência de baselines simples enfraquece a pergunta de pesquisa: não se sabe se GA/PSO/ACO são melhores que nearest neighbor + 2-opt, que é ordens de grandeza mais rápido. Para um TCC comparativo, baselines de complexidade inferior são essenciais para contextualizar o trade-off qualidade × tempo.",
      "ref": "monografia/bib/abntex2-references.bib (sem entradas para constructive heuristics, simulated annealing, ou 2-opt/3-opt como métodos standalone)"
    },
    {
      "id": "f-council-009",
      "severity": "MEDIUM",
      "area": "ga",
      "title": "Nagata (2006) EAX — estado da arte em crossover para TSP — com vault vazio e PDF corrompido",
      "detail": "Edge Assembly Crossover (EAX) de Nagata (2006) é reconhecido como o operador de crossover estado da arte para TSP com GA, produzindo resultados competitivos com LK e Concorde. O artigo está no BibTeX e no vault com rating 5, mas a nota do vault contém apenas metadados (12 linhas, sem resumo, sem análise). O status é 'disponível' com 'classificacao: recuperado' — indicando que o PDF foi obtido mas não processado. O packet relata 'PDF corrompido'.",
      "fix": "Obter PDF legível e completar a nota do vault com resumo, contribuições, e análise de relevância. O EAX é a principal referência para justificar por que o TCC usa OX (Order Crossover) em vez de EAX: a justificativa deve ser explícita (complexidade de implementação, adequação ao problema 3D, ou limitação de escopo).",
      "why": "Sem a nota completa, a monografia não pode discutir adequadamente a escolha de OX vs EAX, enfraquecendo a justificativa de projeto do GA.",
      "ref": "vault/papers/nagata2006eax.md (12 linhas, sem conteúdo); monografia/bib/abntex2-references.bib:463-473"
    },
    {
      "id": "f-council-010",
      "severity": "MEDIUM",
      "area": "lower-bounds",
      "title": "Relaxação AP com min_i c(i,j,k) não é validada contra Held-Karp nem contra limitantes conhecidos",
      "detail": "O lower bound implementado (src/optimization/lowerbound/main.go:39-57) reduz o tensor 3D para matriz 2D via c'(j,k) = min_i c(i,j,k) e resolve o Assignment Problem (Hungarian). Esta relaxação é correta (produz limitante inferior válido), mas: (a) não há discussão sobre quão apertado é este bound em relação ao Held-Karp ou ao ótimo; (b) o texto do vault para Held-Karp (1970) afirma 'a redução c'[j][k] = min_i cost[i][j][k] permite aplicar o HK bound sobre uma matriz 2D reduzida' — o que é conceitualmente impreciso: a redução min_i seguida de Hungarian produz um bound diferente (e mais fraco) que o Held-Karp, que usa 1-trees com multiplicadores Lagrangianos; (c) não há validação experimental do gap entre o bound AP e os resultados dos métodos heurísticos.",
      "fix": "(a) Corrigir a nota do vault de Held-Karp (1970) removendo a afirmação de equivalência entre redução-por-min + Hungarian e o bound HK; (b) na monografia, documentar a relaxação AP como um lower bound ad-hoc específico para o TSP-SD-ATP, distinguindo-o claramente do Held-Karp; (c) incluir análise experimental do gap entre o bound AP e o melhor makespan encontrado, para caracterizar a qualidade do bound.",
      "why": "A confusão conceitual entre relaxação AP (min_i + Hungarian) e relaxação Held-Karp (1-tree + Lagrange) propaga um erro técnico da nota do vault para a monografia, comprometendo a credibilidade da seção de lower bounds.",
      "ref": "vault/papers/heldkarp1970traveling.md:51-52 (afirmação imprecisa); src/optimization/lowerbound/main.go:39-57"
    },
    {
      "id": "f-council-011",
      "severity": "MEDIUM",
      "area": "angular",
      "title": "Aggarwal (2000) 'Angular-Metric TSP' — artigo mais próximo do TSP-SD-ATP — sem BibTeX",
      "detail": "Aggarwal et al. (2000) 'The Angular-Metric Traveling Salesman Problem' (SIAM Journal on Computing, doi:10.1137/S0097539796312719) é o artigo mais diretamente relevante para o TSP-SD-ATP: estuda TSP com custo angular entre arestas consecutivas e prova resultados de aproximação. Está no vault (rating 4, status 'disponível') mas NÃO está no BibTeX. Winter (2002) e Vanhove (2012), que tratam de turn costs em grafos gerais (não TSP), estão no BibTeX com rating 0.",
      "fix": "Adicionar entrada BibTeX para Aggarwal (2000). Revisar ratings: Aggarwal (2000) deve ter rating ≥ 4 por ser o artigo canônico sobre TSP com custo angular; Winter (2002) e Vanhove (2012) são sobre path planning com turn costs, não TSP — rating 0 pode ser justo se forem periféricos, mas a justificativa deve ser explícita.",
      "why": "O artigo mais relevante sobre o problema estudado está ausente da bibliografia formal. Winter (2002) e Vanhove (2012) são referências de grafos gerais, não de TSP, e têm relevância indireta.",
      "ref": "vault/papers/aggarwal2000angular.md (rating 4, sem BibTeX); monografia/bib/abntex2-references.bib (sem entrada aggarwal*)"
    },
    {
      "id": "f-council-012",
      "severity": "MEDIUM",
      "area": "estatistica",
      "title": "Demšar (2006) com BibTeX mas sem nota no vault — referência estatística órfã",
      "detail": "Demšar (2006) 'Statistical Comparisons of Classifiers over Multiple Data Sets' introduz o teste de Friedman com post-hoc de Nemenyi, que é o teste estatístico padrão para comparação de múltiplos algoritmos sobre múltiplos datasets. O artigo está no BibTeX (linhas 39-46) mas NÃO existe nota correspondente em vault/papers/. É a única referência estatística citada e não foi fichada — não há resumo, análise de aplicabilidade ao contexto de TSP, ou discussão de pressupostos (independência, normalidade).",
      "fix": "Criar vault/papers/demsar2006statistical.md com resumo, contribuições, e discussão sobre a adequação do Friedman+Nemenyi para comparação de metaheurísticas estocásticas sobre instâncias de TSP (incluindo a crítica de que Friedman assume rankings independentes, o que pode ser violado quando métodos compartilham estruturas de busca).",
      "why": "O teste estatístico que sustenta TODAS as comparações experimentais da monografia não teve seu artigo de referência lido/fichado. Isso compromete a capacidade de defender a escolha metodológica durante a defesa.",
      "ref": "monografia/bib/abntex2-references.bib:39-46; vault/papers/ (sem demsar2006statistical.md)"
    },
    {
      "id": "f-council-013",
      "severity": "LOW",
      "area": "pso",
      "title": "Clerc (2000) PSO discreto citado como referência mas não é o método usado",
      "detail": "Clerc (2000) propõe um PSO verdadeiramente discreto com operadores de troca e recombinacão sobre permutações. O PSO implementado no TCC usa random keys (Bean 1994) — representação contínua com decodificação por ordenação. São abordagens fundamentalmente diferentes. O vault reconhece que Clerc serve como 'ponte conceitual' e 'justificativa de que PSO discreto não é trivial', mas na monografia essa distinção precisa ser explícita para evitar a impressão de que o método implementado é o de Clerc.",
      "fix": "Na seção de PSO da monografia, incluir um parágrafo explícito: 'Clerc (2000) propõe um PSO discreto com operadores sobre permutações; neste trabalho, adota-se a abordagem alternativa de random keys (Bean, 1994) por [justificativa: simplicidade, preservação da dinâmica contínua do PSO, etc.]'.",
      "why": "A distinção entre PSO discreto (Clerc) e PSO com random keys (Bean) é técnica e importante. Sem ela, um leitor familiarizado com PSO para TSP pode assumir implementação incorreta.",
      "ref": "vault/papers/clerc2000discretepso.md; vault/papers/bean1994genetic.md"
    }
  ],
  "recommendation": "REPROVADO com exigência de correção antes da qualificação/defesa. As correções críticas (f-council-001, f-council-002, f-council-003) são bloqueadoras: não é aceitável defender uma monografia cujos artigos canônicos de lower bound não constam na bibliografia, cuja implementação ACO não corresponde à literatura que a justifica, e cujo problema (TSP assimétrico com dependência de sequência) não referencia as classes de problema que o enquadram. As correções de severidade HIGH (f-council-004 a f-council-007) devem ser resolvidas antes da versão final. As correções MEDIUM e LOW são esperadas para a versão de defesa mas não bloqueiam a continuidade do trabalho. Prazo estimado para correções críticas: 2-3 dias de trabalho focado (adicionar entradas BibTeX, corrigir documentação da implementação ACO, buscar e fichar 3-5 referências de ATSP/TDTSP)."
}
```

---

# Análise Técnica do Referencial — Juiz TSP e Otimização Combinatória

## 1. Sumário Executivo

**Veredito: FAIL (confiança: HIGH)**. O referencial teórico apresenta três falhas bloqueadoras independentes que, em conjunto, comprometem a correção técnica e a completude algorítmica da monografia. Nenhuma delas depende de julgamento subjetivo — são fatos verificáveis no código-fonte (`src/`), no arquivo BibTeX (`monografia/bib/abntex2-references.bib`) e nas notas do vault (`vault/papers/`).

## 2. Falhas Bloqueadoras (CRITICAL)

### 2.1 Held-Karp 1970/1971 fora do BibTeX (f-council-001)

**Evidência:** O arquivo `abntex2-references.bib` (740 linhas, 72 entradas) não contém as strings "heldkarp", "Held", ou "Karp" exceto pela entrada de Valenzuela (1997), que referencia Held-Karp no título mas é um artigo derivado. Os dois artigos canônicos estão no vault:

- `heldkarp1970traveling.md`: rating 5, status "lido", resumo completo com descrição do método de 1-tree, equivalência LP, e três algoritmos (column generation, ascent method, branch-and-bound)
- `heldkarp1971traveling.md`: rating 5, status "pendente" (Part II, formulação DP O(n²2ⁿ))

**Impacto:** A monografia discute lower bounds sem citar a referência que define o que é um lower bound para TSP. É o equivalente a discutir relatividade sem citar Einstein.

**Correção:** Adicionar entradas BibTeX. O DOI está disponível (10.1287/opre.18.6.1138 para Part I).

### 2.2 Implementação ACO não é MMAS (f-council-002)

**Evidência no código** (`src/optimization/aco/ant.go:14-39`):

```go
func (aco *ACO) updatePheromones(ants []ant) {
    // Evaporação em TODOS os τ[i][j][k]
    // Depósito: TODAS as formigas depositam (estilo Ant System)
    for _, ant := range ants {
        delta := aco.Q / ant.lk
        // deposita em τ[prev][curr][next]
    }
}
```

MMAS (Stützle 2000) requer: (a) apenas iteration-best ou global-best depositar; (b) bounds [τ_min, τ_max]; (c) inicialização em τ_max; (d) reinicialização por estagnação. **Nada disso está implementado.** A implementação é um Ant System 3D (Dorigo 1996), não MMAS.

Agravante: a nota `vault/papers/stutzle2000mmas.md:55` afirma textualmente que "A fórmula τ_max = 1/(ρ·L_best) e τ_min = τ_max / (2n) é implementada diretamente no código ACO do TCC" — **isto é factualmente falso**.

### 2.3 ATSP e TDTSP completamente ausentes (f-council-003)

O TSP-SD-ATP gera custos assimétricos (c'(j,k) ≠ c'(k,j) após redução min_i) e com dependência de sequência (c(i,j,k) depende de i, j e k). Isto o enquadra simultaneamente em duas classes bem estudadas:

- **ATSP (Asymmetric TSP)**: dezenas de surveys, algoritmos exatos (branch-and-bound, patching), e metaheurísticas especializadas
- **TDTSP (Time-Dependent TSP)**: custos que dependem da posição na sequência ou do "tempo" de visita

Zero entradas BibTeX para qualquer uma destas classes. Kinable (2017) cobre TDTSP com decision diagrams e está fichado no vault (rating 5, lido) mas também não está no BibTeX.

## 3. Problemas de Alta Severidade (HIGH)

### 3.1 Feromônio 3D sem documentação como contribuição (f-council-004)

O ACO implementado usa `pheromones [][][]float64` — um tensor tridimensional τ(i,j,k). Na literatura canônica de ACO para TSP (Dorigo 1996, 1997, 2004; Stützle 2000), o feromônio é sempre bidimensional τ(i,j). A terceira dimensão é necessária porque o TSP-SD-ATP tem custos que dependem da tripla (nó anterior, nó atual, próximo nó).

Esta é uma **contribuição algorítmica genuína** do TCC, mas não está documentada como tal. O risco é que pareça um erro de implementação (confundir TSP com problema de 3 índices) em vez de uma adaptação deliberada e justificada.

### 3.2 Bean (1994) com rating 0 e PDF ilegível (f-council-005)

Bean (1994) introduz random keys — o mecanismo exato que o PSO usa para converter vetores reais em permutações (ordenar pelos valores das chaves). O artigo tem rating 0 (o mais baixo) no vault porque o PDF contém apenas a página de metadados da INFORMS — o corpo do artigo não pôde ser lido. A nota do vault é honesta sobre isso: "não foi possível confirmar detalhes experimentais, pseudocódigo ou conclusões."

**Implicação:** A fundamentação do PSO implementado baseia-se em um artigo que o autor não leu. O rating 0 é correto dada a ilegibilidade do PDF, mas a dependência crítica do método em relação a este artigo exige que ele seja lido.

### 3.3 Parâmetro fantasma --gama (f-council-006)

`AGENTS-experiments.md:24` lista `--gama` como parâmetro do ACO. O código em `src/optimization/aco/main.go:10-14` define `Params` com `Alpha, Beta, Rho, Q` — sem Gama. O parâmetro não existe, não é usado, e não aparece em nenhum arquivo `.go`.

### 3.4 Desbalanceamento ACO (13) vs GA (8) vs PSO (5) (f-council-007)

A seção ACO tem 13 artigos, incluindo 4 de ACO neural/RL (DeepACO 2023, PPACO 2024, GPACO 2025, NeuFACO 2025) que usam redes neurais para aprender políticas de feromônio — completamente desvinculados da implementação Ant System 3D do TCC. Enquanto isso, o PSO tem apenas 5 artigos, dos quais 3 foram descartados (sun2024hybrid, huang2025matrix, kappagantula2025dpso) e 1 não é o método usado (clerc2000discretepso). Isso deixa essencialmente **1 artigo** (kennedy1995particle + araujo2025pso) para fundamentar o PSO.

## 4. Problemas de Média Severidade (MEDIUM)

### 4.1 Ausência de baselines construtivas e Simulated Annealing (f-council-008)

Sem nearest neighbor, insertion heuristics, Christofides, ou Simulated Annealing, a monografia compara GA/PSO/ACO apenas entre si e com bruteforce (para n ≤ 15). Não há como responder: "GA/PSO/ACO são melhores que nearest neighbor + 2-opt?"

### 4.2 Nagata (2006) EAX com vault vazio (f-council-009)

O operador de crossover estado da arte para TSP (EAX) está no BibTeX mas a nota do vault tem apenas 12 linhas de metadados. Sem análise do EAX, a monografia não pode justificar adequadamente por que usa OX em vez do estado da arte.

### 4.3 Confusão conceitual: relaxação AP ≠ Held-Karp (f-council-010)

A nota `heldkarp1970traveling.md:51-52` afirma que "a redução c'[j][k] = min_i cost[i][j][k] permite aplicar o HK bound sobre uma matriz 2D reduzida". Isto é conceitualmente impreciso:

- O bound de Held-Karp é o valor ótimo da relaxação Lagrangiana baseada em 1-trees com multiplicadores π_i
- A redução min_i + Hungarian produz um bound de Assignment Problem (AP) — relacionado mas diferente e geralmente mais fraco
- A equivalência é entre o bound HK e a relaxação LP (subtour elimination), não entre HK e AP

### 4.4 Demšar (2006) sem nota no vault (f-council-012)

O teste estatístico (Friedman + Nemenyi) que fundamenta todas as comparações experimentais não teve seu artigo de referência fichado. Não há discussão sobre pressupostos, adequação a metaheurísticas estocásticas, ou limitações.

## 5. Análise de Cobertura por Área

| Área | Artigos | BibTeX | Vault | Avaliação |
|------|---------|--------|-------|-----------|
| TSP Clássico | 5 | 4 | 5 | **OK** (lawler1985 sem BibTeX, aceitável como referência histórica) |
| GA | 8 | 8 | 8 | **FRÁGIL** (bean1994 ilegível, nagata2006 vazio, hga2024 descartado) |
| PSO | 5 | 5 | 5 | **INSUFICIENTE** (apenas 2 ativos, 3 descartados, 1 é o método errado) |
| ACO | 13 | 13 | 13 | **DESBALANCEADO** (4 de ACO neural irrelevantes; implementação não corresponde ao MMAS citado) |
| Lower Bounds | 10 | 1 | 10 | **CRÍTICO** (9/10 sem BibTeX, incluindo os canônicos Held-Karp) |
| Drone-TSP | 7 | 7 | 7 | **OK** (2 PDFs corrompidos: dellamico2021, dellamico2022) |
| Angular/Turn | 3 | 2 | 3 | **FRÁGIL** (aggarwal2000 canônico sem BibTeX; winter e vanhove rating 0) |
| Comparativos | 9 | 9 | 9 | **OK** (alexander2020 rating 2, baixa qualidade) |
| Estatística | 1 | 1 | 0 | **INCOMPLETO** (sem nota no vault) |
| ATSP | 0 | 0 | 0 | **AUSENTE** |
| TDTSP | 0 | 0 | 0 | **AUSENTE** (kinable2017 cobre mas está em lower-bounds sem BibTeX) |
| Heurísticas construtivas | 0 | 0 | 0 | **AUSENTE** |
| Simulated Annealing | 0 | 0 | 0 | **AUSENTE** |

## 6. Análise das Escolhas Algorítmicas

| Escolha | Fundamentação | Avaliação |
|---------|---------------|-----------|
| GA: OX crossover | Oliver (1987) + Potvin (1996) survey | **OK** — OX é bem fundamentado para TSP com permutação direta |
| GA: seleção por torneio | Goldberg (1989) | **OK** — torneio é padrão, bem documentado |
| GA: mutação por troca | Larrañaga (1999) survey | **OK** — swap mutation é padrão para representação por permutação |
| PSO: random keys | Bean (1994) | **FRÁGIL** — artigo não lido (PDF ilegível, rating 0) |
| ACO: feromônio 3D | Nenhuma referência | **INOVACÃO NÃO DOCUMENTADA** — não há precedente na literatura |
| ACO: Ant System (não MMAS) | Dorigo (1996, 1997) | **OK para AS, INCORRETO para MMAS** — código é AS, vault afirma MMAS |
| Lower bound: AP via Hungarian | Kuhn (1955) implícito | **OK como bound, IMPRECISO como justificativa** — não é HK, não deve ser apresentado como tal |
| Penalidade angular normalizada θ/π | Winter (2002) + Aggarwal (2000) | **ACEITÁVEL** — Winter modela turn costs, Aggarwal é específico para TSP angular |

## 7. Recomendações para Correção (Ordenadas por Prioridade)

### Bloqueadoras (devem ser resolvidas antes da qualificação)

1. **Adicionar Held-Karp ao BibTeX** (f-council-001): 30 minutos. DOIs disponíveis.
2. **Corrigir documentação ACO** (f-council-002): ou implementar bounds MMAS no código (2-4 horas) ou corrigir a nota do vault e o texto da monografia para refletir que a implementação é Ant System (30 minutos).
3. **Adicionar referências de ATSP/TDTSP** (f-council-003): buscar 2-3 artigos (survey de ATSP + artigo de TDTSP + Kinable 2017 para BibTeX), ~2 horas incluindo fichamento.

### Alta prioridade (antes da versão final)

4. **Ler Bean (1994)** (f-council-005): obter PDF legível e atualizar nota do vault.
5. **Documentar feromônio 3D como contribuição** (f-council-004): 1-2 parágrafos na seção de ACO.
6. **Corrigir documentação do --gama** (f-council-006): 5 minutos.
7. **Balancear referencial** (f-council-007): mover ACO neural para "Tendências Recentes", adicionar 2-3 referências de PSO.

### Média prioridade (para a defesa)

8. **Adicionar baselines** (f-council-008): avaliar se inclui experimentalmente ou apenas como discussão.
9. **Completar nota do Nagata EAX** (f-council-009).
10. **Corrigir confusão AP vs HK** (f-council-010).
11. **Criar nota do Demšar** (f-council-012).
12. **Adicionar Aggarwal ao BibTeX** (f-council-011).

## 8. Conclusão Técnica

O referencial tem méritos: os canônicos de TSP clássico (Garey & Johnson, Applegate et al., Lin-Kernighan), GA (Holland, Goldberg), PSO (Kennedy & Eberhart), e ACO (Dorigo 1996, 1997, 2004) estão presentes, lidos, e bem fichados. Os artigos de drone-TSP e comparativos cobrem adequadamente o contexto de aplicação.

No entanto, três falhas estruturais tornam o referencial **não defensável no estado atual**:

1. **Os artigos canônicos de lower bound não estão na bibliografia.** Independentemente do método de bound usado (AP relaxation), uma monografia sobre TSP que não referencia Held & Karp (1970) tem uma lacuna de fundamentação que qualquer banca de otimização combinatória identificará em 30 segundos.

2. **A implementação ACO não corresponde à literatura citada como justificativa.** Afirmar que o código implementa MMAS quando ele implementa Ant System é um erro factual que mina a credibilidade da seção metodológica.

3. **O problema estudado (TSP assimétrico com dependência de sequência) não está enquadrado nas classes de problema que o definem.** Sem referências a ATSP e TDTSP, o TSP-SD-ATP parece um problema inventado para a monografia, quando na verdade é uma instância de classes bem estabelecidas.

Com as correções acima (estimativa: 8-12 horas de trabalho), o referencial atinge qualidade técnica suficiente para uma defesa de TCC em nível de mestrado.
