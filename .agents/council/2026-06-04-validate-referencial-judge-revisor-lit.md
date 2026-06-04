```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "A revisão de literatura apresenta viés severo de seleção (11 ACO vs 6 PSO no vault, razão ~1.8:1), referências de qualidade questionável (rating 0) usadas como autoridade na Seção 2.2, a Seção 2.8 (Lower Bounds) sem nenhuma citação apesar de 10 artigos fichados na área, e zero rastreabilidade entre claims da monografia e artigos de suporte (role/chapters/claim_support vazios em todas as 84 notas).",
  "findings": [
    {
      "id": "f-council-lit-001",
      "title": "Seção 2.8 (Lower Bounds) sem citações — 10 artigos fichados, 0 referenciados",
      "severity": "critical",
      "category": "citation_gap",
      "description": "A Seção 2.8 descreve a relaxação AP como lower bound derivando uma matriz 2D de um tensor 3D e aplicando o algoritmo Húngaro. O texto NÃO cita nenhum artigo — nem Held-Karp (1970, rating 5, lido), nem Johnson (1996, rating 5, validação empírica do HK), nem Valenzuela & Jones (1997, BibTeX existente), nem os 9 outros artigos fichados sobre bounds. A seção menciona 'subtours' sem referenciar a formulação LP do TSP (subtour elimination constraints). Numa monografia de mestrado, uma seção inteira da fundamentação sem suporte bibliográfico é inaceitável.",
      "fix": "(1) Adicionar entradas BibTeX para heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, balas1985branch, fischetti1992additive, karp1979patching, kinable2017hybrid, righini2021efficient, aggarwal2000angular. (2) Citar heldkarp1970traveling como referência canônica para lower bounds de TSP. (3) Justificar por que a relaxação AP foi adotada em vez da relaxação Lagrangiana (Held-Karp) — a nota de heldkarp1970traveling já contém resumo do artigo e gap empírico < 0.8%. (4) Referenciar valenzuela1997estimating para a técnica de estimação do HK bound.",
      "why": "A nota do vault para heldkarp1970traveling tem rating 5 e status 'lido' — o conteúdo foi lido, considerado de altíssima relevância, mas a entrada BibTeX nunca foi criada. Todos os 10 artigos da área de lower bounds foram classificados no index.md como 'Pendências BibTeX' mas nenhum foi integrado. O problema é puramente de execução: o material existe, foi avaliado como relevante, mas não foi transferido ao BibTeX nem citado no texto.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:78-88; vault/papers/index.md:146-160"
    },
    {
      "id": "f-council-lit-002",
      "title": "Viés de seleção ACO (11 artigos) vs PSO (6 artigos) — razão 1.8:1 no vault",
      "severity": "high",
      "category": "selection_bias",
      "description": "O vault contém 11 artigos categorizados como ACO (incluindo 4 irrelevantes de ACO+Deep Learning), contra apenas 6 de PSO. Na fundamentação, são 4 citações ACO vs 2 PSO. Considerando que o TCC compara GA, PSO e ACO em pé de igualdade, espera-se cobertura bibliográfica simétrica. Além disso, a seção de PSO usa apenas bean1994 (rating 0, um artigo de GA/random keys) como ponte para justificar a representação por random keys — não há revisão específica sobre PSO para TSP (shami2022pso e gad2022pso existem no BibTeX mas não são citados).",
      "fix": "(1) Citar shami2022pso (IEEE Access, survey abrangente) e gad2022pso (Archives of Comp Methods, revisão sistemática) na Seção 2.6 para dar suporte à afirmação de que 'sua aplicação a problemas de permutação exige uma etapa de codificação'. (2) Avaliar descarte de 3-4 artigos de ACO+DL (deepaco2023, neufaco2025, ppaco2024, gpaco2025) do vault se irrelevantes para o escopo comparativo, ou mover para categoria à parte. (3) Adicionar pelo menos 1 artigo adicional de PSO para TSP como referência de implementação — araujo2025pso já está no BibTeX mas não citado. (4) Considerar adicionar Simulated Annealing como baseline citada na Seção 2.9 (halim2019combinatorial já menciona SA indiretamente via tags).",
      "why": "O desbalanceamento ACO/PSO reflete o viés do pesquisador durante a fase de levantamento bibliográfico. Os artigos ACO+DL (DeepACO, NeuFACO, PPACO, GPACO) foram incluídos por interesse tangencial em deep learning, mas não contribuem para a comparação GA vs PSO vs ACO clássico que é o núcleo do TCC. Ocupam 4 das 11 posições da categoria ACO.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:64-69; vault/papers/index.md:53-68"
    },
    {
      "id": "f-council-lit-003",
      "title": "Artigos com rating 0 usados como autoridade na fundamentação",
      "severity": "high",
      "category": "citation_quality",
      "description": "Três artigos com rating 0 no vault são citados no texto da fundamentação: (a) winter2002modeling (rating 0, Sec 2.2) usado como referência para 'custos de curva', mas a nota indica que foi lido e considerado irrelevante; (b) vanhove2012route (rating 0, Sec 2.2) citado como autoridade em custos de mudança de direção, status 'pendente' (sem resumo); (c) bean1994genetic (rating 0, Sec 2.6) citado como justificativa para random keys no PSO, mas o artigo é de GA — a nota indica que a relevância é inferida 'a partir do metadado local até que o corpo do artigo esteja legível'. Isso contradiz o uso como autoridade. Citar artigos com rating 0 mina a credibilidade da revisão.",
      "fix": "(1) Para winter2002 e vanhove2012: ou reavaliar com leitura completa e atualizar rating, ou substituir por referências mais específicas sobre penalidade angular / turn costs (ex.: aggarwal2000angular, que está no vault com rating 4 mas sem BibTeX). (2) Para bean1994: manter a citação como referência histórica de random keys, mas atualizar o rating após leitura completa OU adicionar uma referência mais direta sobre random keys em PSO para TSP. (3) Adicionar política no AGENTS.md: artigos com rating 0-1 não devem ser citados como autoridade primária sem reavaliação.",
      "why": "O rating 0 significa 'irrelevante para o TCC' ou 'não avaliado adequadamente'. Citar tais artigos como suporte a afirmações técnicas é contraditório e sugere que a revisão foi feita às pressas, sem fechar o ciclo ler → avaliar → citar.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:18,70; vault/papers/winter2002modeling.md; vault/papers/vanhove2012route.md; vault/papers/bean1994genetic.md"
    },
    {
      "id": "f-council-lit-004",
      "title": "demsar2006 — base metodológica dos testes estatísticos — sem nota no vault",
      "severity": "high",
      "category": "citation_gap",
      "description": "Demšar (2006) é a referência metodológica central do protocolo estatístico adotado na monografia (Friedman + Nemenyi). A Seção 2.9 dedica um parágrafo inteiro a esta escolha. No entanto, NÃO existe nota no vault para demsar2006statistical. É o ÚNICO artigo do BibTeX sem nota correspondente. Isso significa que a base metodológica dos testes estatísticos do TCC nunca foi fichada, resumida ou avaliada quanto à adequação ao contexto TSP (o artigo original é focado em classificadores de ML).",
      "fix": "Criar nota no vault para demsar2006statistical com: (a) resumo da metodologia Friedman+Nemenyi; (b) avaliação da adequação ao contexto TSP (o artigo original compara classificadores ML, não rotas de otimização — isso precisa ser discutido); (c) rating de relevância. Adicionar referência alternativa mais específica para testes não-paramétricos em otimização combinatória se disponível.",
      "why": "Demšar (2006) é um artigo de JMLR focado em comparar classificadores sobre múltiplos datasets. Sua aplicação a meta-heurísticas de otimização é uma extensão metodológica que merece justificativa. Sem nota no vault, não há registro de que essa adequação foi avaliada criticamente.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:95; monografia/bib/abntex2-references.bib:39-46"
    },
    {
      "id": "f-council-lit-005",
      "title": "Rastreabilidade zero entre claims da monografia e artigos de suporte",
      "severity": "high",
      "category": "traceability",
      "description": "NENHUM dos 84 artigos no vault possui os campos `role`, `chapters`, `claim_support` preenchidos no frontmatter. A view 'Uso Na Monografia' no papers.base exibe exatamente esses campos e está completamente vazia. O index.md do vault tem categorização manual por área (TSP Clássico, GA, PSO, ACO, Drone, Lower Bounds, Survey/Comparativo) mas não mapeia para capítulos da monografia. Sem rastreabilidade, é impossível verificar quais claims do texto são sustentadas por quais artigos.",
      "fix": "(1) Para cada artigo citado na fundamentação, preencher no mínimo `chapters: [fundamentacao]` e `role: cornerstone|support|context|comparison`. (2) Para artigos em outras seções (metodologia, resultados), preencher chapters correspondentes. (3) Usar `claim_support` para mapear artigo → claim específico (ex.: 'ACO supera GA em qualidade com maior custo computacional' → haroun2015performance, alexander2020comparison). (4) Priorizar os 26 artigos que já têm `role`/`chapters` (via grep mostrou 26 matches) — verificar se são valores reais ou apenas cabeçalhos de template.",
      "why": "O vault tem infraestrutura de rastreabilidade (papers.base com campos role/chapters/claim_support, views 'Uso Na Monografia') mas nenhum dado foi populado. Isso inviabiliza auditoria de consistência entre claims e evidências.",
      "ref": "vault/bases/papers.base; vault/papers/index.md"
    },
    {
      "id": "f-council-lit-006",
      "title": "4 artigos ACO+Deep Learning no vault irrelevantes para o escopo comparativo",
      "severity": "medium",
      "category": "selection_bias",
      "description": "deepaco2023, neufaco2025, ppaco2024 e gpaco2025 são artigos sobre ACO aumentada por deep learning, reinforcement learning ou genetic programming. Nenhum é citado na fundamentação. Ocupam 4 das 11 posições da categoria ACO no vault e 4 das 72 entradas BibTeX. Sua presença infla artificialmente a cobertura de ACO e distorce a percepção de balanceamento.",
      "fix": "(1) Remover do BibTeX ou mover para seção 'Trabalhos Relacionados — Métodos Híbridos' se forem mencionados. (2) Se mantidos, justificar na fundamentação por que não são usados (ex.: 'extensões recentes com deep learning foram consideradas, mas fogem ao escopo desta comparação de métodos clássicos'). (3) Atualizar a categoria no index.md para 'ACO/Deep Learning' em vez de 'ACO'.",
      "why": "O TCC compara GA, PSO e ACO clássicos com operadores bio-inspirados. ACO+DL é uma família distinta que depende de treinamento supervisionado/não-supervisionado. Sua inclusão no vault sem justificativa de escopo cria ruído bibliográfico.",
      "ref": "vault/papers/index.md:31-34 (ACO); monografia/bib/abntex2-references.bib:475-503"
    },
    {
      "id": "f-council-lit-007",
      "title": "11 artigos no vault sem entrada BibTeX — assimetria vault↔BibTeX",
      "severity": "medium",
      "category": "bibliographic_integrity",
      "description": "13 artigos têm nota no vault mas NÃO têm entrada BibTeX: aggarwal2000angular, balas1985branch, fischetti1992additive, heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, karp1979patching, kinable2017hybrid, lawler1985traveling, leraromero2020dynamic, lysgaard1999cluster, muthanna2022uav, righini2021efficient. Destes, 10 são da área de Lower Bounds (a mais crítica). O index.md os lista como 'Pendências BibTeX' mas a prioridade está marcada como 'Alta' para apenas 3 deles. O artigo lawler1985traveling (rating 5, TSP Clássico) está na mesma situação e não é citado na fundamentação — é uma omissão grave, pois Lawler et al. (1985) é a referência canônica sobre TSP pré-applegate.",
      "fix": "(1) Criar entradas BibTeX para todos os 13 artigos. (2) Priorizar: lawler1985traveling (TSP canônico, rating 5) + 10 artigos de lower bounds. (3) Avaliar se lysgaard1999cluster e muthanna2022uav são realmente necessários ou podem ser removidos do vault.",
      "why": "O fluxo de trabalho vault→BibTeX→citação está quebrado. Artigos são fichados, avaliados com ratings altos (heldkarp1970 rating 5, johnson1996 rating 5, lawler1985 rating 5), mas nunca chegam ao BibTeX. O script scripts/import-bib-to-vault.sh só funciona na direção BibTeX→vault; o fluxo reverso (vault→BibTeX) não tem ferramenta.",
      "ref": "vault/papers/index.md:146-160"
    },
    {
      "id": "f-council-lit-008",
      "title": "Qualidade questionável de veículos em 5 artigos citados ou no BibTeX",
      "severity": "medium",
      "category": "publication_quality",
      "description": "Dos artigos no BibTeX: (a) alexander2020 — conferência EAI (rating 2); (b) almufti2025 — IJSciWorld, journal de baixo impacto (rating 3); (c) wadi2025 — IJCMEM, publisher questionável (rating 3); (d) araujo2025pso — arXiv preprint sem revisão por pares, sem DOI formal; (e) hossain2024 — Bitlis Eren Univ journal, regional e de baixo impacto. Para uma monografia de mestrado, a presença de 5 referências em veículos marginais não é grave por si só, mas todas estão na área de estudos comparativos — exatamente onde a qualidade das fontes é mais importante para sustentar claims de superioridade de método.",
      "fix": "(1) Para estudos comparativos (Seção 2.9), priorizar wu2020 (ACM, rating 4), chandra2022 (IJIST, rating 4), halim2019 (Archives of Comp Methods, rating 4) e haroun2015 (IJCA, rating 4). (2) Mover almufti2025, wadi2025 e alexander2020 para menção secundária ou remover. (3) Para araujo2025pso (arXiv), verificar se foi publicado em veículo com revisão por pares desde o depósito.",
      "why": "A qualidade da revisão de estudos comparativos é diretamente proporcional à qualidade dos estudos citados. Citar artigos de journals marginais como evidência de desempenho relativo enfraquece as conclusões.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:91-95; monografia/bib/abntex2-references.bib"
    },
    {
      "id": "f-council-lit-009",
      "title": "Ausência de Simulated Annealing como baseline citado na fundamentação",
      "severity": "medium",
      "category": "coverage_gap",
      "description": "SA é um baseline natural em qualquer comparação de meta-heurísticas para TSP. NÃO há nenhum artigo de SA no vault. halim2019combinatorial e chandra2022comparative têm SA nas tags, mas SA não é discutido como método na fundamentação. A ausência é notável porque a Seção 2.9 afirma que 'o desempenho relativo depende da instância, da representação e dos parâmetros' — SA seria o contraponto natural como método de trajetória única vs os três métodos populacionais comparados.",
      "fix": "(1) Adicionar ao vault pelo menos 1 artigo canônico de SA para TSP (ex.: Kirkpatrick et al. 1983, ou Černý 1985, ou Aarts & Korst 1989). (2) Citar SA na Seção 2.9 como baseline de método de busca local não-populacional, justificando sua exclusão do escopo comparativo.",
      "why": "A omissão de SA não invalida a comparação GA vs PSO vs ACO, mas enfraquece a revisão de literatura porque SA é o baseline mais comum em surveys comparativos de meta-heurísticas (presente em halim2019, chandra2022 e hossain2024, todos já no BibTeX).",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:91-95"
    },
    {
      "id": "f-council-lit-010",
      "title": "18 artigos com status 'pendente' (sem resumo) — 29.5% do vault não lido",
      "severity": "medium",
      "category": "reading_coverage",
      "description": "Dos 61 artigos no vault, 18 estão com status 'pendente' (sem seção 'Resumo' com ≥ 40 caracteres). Isso inclui lawler1985traveling (rating 5), nagata2006eax (rating 5, citado na fundamentação), vanhove2012route (rating 0, citado), dellamico2021multiple e dellamico2022exact (rating 4, citados), e 5 dos 10 artigos de lower bounds. Citar artigos não lidos é aceitável se forem referências canônicas bem conhecidas (ex.: lawler1985), mas vanhove2012 com rating 0 e status pendente sendo citado como autoridade é inaceitável.",
      "fix": "(1) Completar leitura e resumo dos 18 artigos pendentes, priorizando os já citados na fundamentação. (2) Para artigos canônicos não lidos (lawler1985, nagata2006), adicionar resumo baseado em fontes secundárias confiáveis com a nota 'resumo baseado em fonte secundária'. (3) Atualizar ratings após leitura completa.",
      "why": "A taxa de 29.5% de artigos não lidos é alta para uma monografia de mestrado. O problema é agravado pelo fato de que 5 dos artigos pendentes já são citados no texto.",
      "ref": "vault/papers/index.md:112-114"
    },
    {
      "id": "f-council-lit-011",
      "title": "Categorização do index.md não mapeia para estrutura da monografia",
      "severity": "low",
      "category": "organization",
      "description": "O index.md organiza artigos em 7 categorias: TSP Clássico, GA, PSO, ACO, Drone, Lower Bounds, Survey/Comparativo. A monografia tem 9 seções na fundamentação (2.1-2.9). O mapeamento é aproximadamente direto (ex.: GA→2.5, PSO→2.6, ACO→2.7), mas: (a) 'TSP Clássico' contém winter2002 e aggarwal2000angular que são da Seção 2.2 (custos dependentes); (b) 'Drone' contém vanhove2012 que é de curva/rota, não drone; (c) não há categoria para 'Busca Exaustiva' (Seção 2.4) — rajwar2023exhaustive está em Survey/Comparativo. Isso dificulta verificar cobertura por seção.",
      "fix": "(1) Adicionar coluna 'Seção' ao index.md mapeando cada artigo à seção da fundamentação onde é (ou deveria ser) citado. (2) Criar categoria 'Custos Dependentes de Sequência' para winter2002, vanhove2012, aggarwal2000angular. (3) Mover rajwar2023exhaustive para categoria 'Busca Exaustiva'.",
      "why": "A categorização atual foi feita por área temática, não por estrutura da monografia. Para auditoria de cobertura, o mapeamento por seção é mais útil.",
      "ref": "vault/papers/index.md"
    },
    {
      "id": "f-council-lit-012",
      "title": "Seção 2.4 (Busca Exaustiva) não cita literatura de enumeração ou branch-and-bound",
      "severity": "low",
      "category": "citation_gap",
      "description": "A Seção 2.4 descreve a busca exaustiva em 3 linhas, sem nenhuma citação. Embora a busca exaustiva seja conceitualmente simples, existem referências relevantes para métodos exatos em TSP: (a) rajwar2023exhaustive já está no vault e no BibTeX — é uma revisão de meta-heurísticas que cobre métodos exaustivos; (b) balas1985branch e fischetti1992additive no vault tratam de branch-and-bound para TSP; (c) hoffman2013tspencyclopedia (BibTeX) é uma referência enciclopédica que cobre enumeração.",
      "fix": "Citar rajwar2023exhaustive e/ou hoffman2013tspencyclopedia na Seção 2.4 para dar suporte à afirmação de que 'seu custo cresce de forma fatorial'.",
      "why": "Trata-se de omissão menor, mas consistente com o padrão de seções sem suporte bibliográfico (2.4, 2.8). Uma monografia deve referenciar até mesmo afirmações básicas.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:52-55"
    },
    {
      "id": "f-council-lit-013",
      "title": "wang2021ant — rating 0, não citado, mas no BibTeX e no vault com nota completa",
      "severity": "low",
      "category": "bibliographic_hygiene",
      "description": "wang2021ant (Applied Soft Computing, tuning de parâmetros ACO) tem rating 0 no vault ('Abordagem de tuning automático de parâmetros é relevante para a implementação de ACO no repositório. A sensibilidade paramétrica do ACO é um desafio prático no TCC'). O rating 0 contradiz o texto da nota, que afirma relevância. Não é citado na fundamentação mas está no BibTeX. Se a sensibilidade paramétrica é um desafio prático, deveria ser citado na Seção 2.7 ou 2.9.",
      "fix": "Reavaliar rating de wang2021ant: se relevante (rating ≥ 3), citar na Seção 2.7 ou 2.9. Se irrelevante, remover do BibTeX ou justificar rating 0.",
      "why": "Rating inconsistente com o conteúdo da nota. O artigo trata de tuning de parâmetros ACO, que é diretamente relevante para a implementação e para a discussão de ameaças à validade.",
      "ref": "vault/papers/wang2021ant.md; monografia/bib/abntex2-references.bib:261-269"
    }
  ],
  "recommendation": "REPROVADO — correções obrigatórias antes da defesa. Prioridade 1 (bloqueante): (a) Adicionar BibTeX e citações para os 10 artigos de lower bounds, com heldkarp1970traveling como referência canônica na Seção 2.8; (b) Criar nota vault para demsar2006 com avaliação crítica da adequação ao contexto TSP. Prioridade 2 (alta): (c) Corrigir viés ACO/PSO — citar shami2022pso e gad2022pso na Seção 2.6, remover ou isolar artigos ACO+DL; (d) Reavaliar e corrigir ratings de artigos citados com rating 0 (winter2002, vanhove2012, bean1994). Prioridade 3 (média): (e) Popular campos role/chapters/claim_support nos 26+ artigos citados para estabelecer rastreabilidade; (f) Completar leitura dos 18 artigos pendentes; (g) Adicionar SA como baseline referenciado."
}
```

# Voto do Juiz Revisor de Literatura — FAIL (HIGH confidence)

## 1. Diagnóstico Geral

A revisão de literatura da monografia sofre de **três problemas estruturais** que, combinados, comprometem a qualidade acadêmica do texto:

1. **Cobertura assimétrica**: ACO tem quase o dobro de artigos de PSO no vault (11 vs 6), com 4 dos 11 sendo de ACO+Deep Learning (irrelevantes para o escopo). A Seção 2.6 (PSO) é a mais magra de todas, apoiada em apenas 2 citações + bean1994 (rating 0, artigo de GA).

2. **Lacunas de citação**: A Seção 2.8 (Lower Bounds) tem ZERO citações apesar de 10 artigos fichados na área. A Seção 2.4 (Busca Exaustiva) também não cita ninguém. A base metodológica dos testes estatísticos (Demšar 2006) nunca foi fichada no vault.

3. **Rastreabilidade inexistente**: Nenhum artigo tem os metadados de rastreabilidade preenchidos (role, chapters, claim_support). A view "Uso Na Monografia" está vazia. É impossível auditar quais claims do texto são sustentadas por quais artigos.

## 2. Análise por Seção da Fundamentação

### 2.1 — TSP Clássico ✅
**Cobertura adequada.** lawler1985 (rating 5) está fichado mas sem BibTeX — é a única lacuna. As 4 citações (applegate2006, garey1979, lin1973 + lawler1985 não citado) cobrem bem o espectro: complexidade computacional, tratado computacional, heurística clássica.

### 2.2 — Custos Dependentes de Sequência ⚠️
**Frágil.** As duas referências citadas (winter2002, vanhove2012) têm rating 0. winter2002 foi lido e considerado irrelevante; vanhove2012 está pendente (sem resumo). O artigo mais relevante da área — aggarwal2000angular (rating 4, "Approximation algorithms for TSP with turn costs") — está no vault sem BibTeX. **Recomendação:** substituir winter2002/vanhove2012 por aggarwal2000angular como referência primária para turn costs.

### 2.3 — Drone-TSP ✅
**Cobertura adequada.** 5 citações de 7 artigos disponíveis. murray2015 e agatz2018 são referências canônicas em Transportation Research C e Transportation Science. dellamico2021/2022 estão pendentes mas são referências válidas. rajan2022 é diretamente relevante para patrulhamento.

### 2.4 — Busca Exaustiva ⚠️
**Sem citações.** A seção tem 3 linhas descrevendo enumeração fatorial. rajwar2023exhaustive está no vault e no BibTeX como "An Exhaustive Review of Metaheuristic Algorithms" mas não é citado. Recomendação: adicionar 1-2 citações.

### 2.5 — GA ✅
**Cobertura sólida.** 5 citações cobrem fundação (holland1975, goldberg1989), operadores para TSP (potvin1996, larranaga1999) e operador específico (nagata2006). oliver1987 (rating 4) está no vault/BibTeX mas não citado — é uma omissão menor, pois foi um dos primeiros estudos de crossover para TSP.

### 2.6 — PSO ⚠️
**Magro.** Apenas 3 citações: kennedy1995 (fundação), clerc2000 (PSO discreto), bean1994 (GA/random keys, rating 0). shami2022pso e gad2022pso (surveys abrangentes, IEEE Access e Archives of Comp Methods) estão no BibTeX mas não são citados. A justificativa para random keys depende exclusivamente de bean1994, que é um artigo de GA com rating 0. **Recomendação:** citar shami2022pso e gad2022pso; adicionar referência específica sobre PSO com random keys para TSP.

### 2.7 — ACO ✅
**Cobertura adequada para o texto.** 4 citações cobrem a fundação (dorigo1996, dorigo1997), o framework (dorigo2004book) e a variante MMAS (stutzle2000). A nota sobre feromônio 3D é contribuição própria e está corretamente sinalizada. O problema não está no que é citado, mas no que está no vault sem ser usado: dorigo2005acotheory, blum2005acointro, wang2021ant + 4 artigos ACO+DL.

### 2.8 — Lower Bounds 🔴
**CRÍTICO. Zero citações.** Esta é a falha mais grave da fundamentação. A seção descreve uma relaxação AP original e aplicação do algoritmo Húngaro como lower bound, mas:
- Não cita Held & Karp (1970, 1971) — o lower bound canônico para TSP (gap empírico < 0.8%)
- Não cita Johnson (1996) — validação empírica do HK bound
- Não cita Valenzuela & Jones (1997) — estimação do HK bound (BibTeX existente)
- Não referencia formulação LP do TSP (subtour elimination constraints)
- Menciona "subtours" sem contexto bibliográfico

**10 artigos fichados na área de lower bounds, 0 com BibTeX completo, 0 citados.** O index.md os lista como "Pendências BibTeX" há semanas.

### 2.9 — Estudos Comparativos ✅
**Aceitável.** 5 citações (wu2020, haroun2015, chandra2022, halim2019, demsar2006). A qualidade dos veículos é mista (ACM conference + IJCA + IJIST + Archives of Comp Methods + JMLR). A ausência de SA como baseline é notável mas não bloqueante.

## 3. Qualidade por Veículo de Publicação

| Nível | Contagem | Exemplos |
|-------|----------|----------|
| **Primeira linha** | ~18 | garey1979 (livro), applegate2006 (Princeton), dorigo1996 (IEEE Trans), kennedy1995 (IEEE), murray2015 (TRC), agatz2018 (Transportation Science), stutzle2000 (FGCS), larranaga1999 (AI Review), rajwar2023 (AI Review), pop2024 (EJOR), bock2025 (EJOR), johnson1996 (SIAM), shami2022pso (IEEE Access), dorigo2004book (MIT Press), holland1975, goldberg1989, lin1973 (Operations Research), potvin1996 (Annals of OR) |
| **Média** | ~12 | haroun2015 (IJCA), chandra2022 (IJIST), halim2019 (Archives of Comp Methods), gad2022pso (Archives of Comp Methods), clerc2000 (Springer book chapter), freitas2020 (ITOR), bean1994 (INFORMS J Computing), dellamico2021/2022 (Networks/ITOR), rajan2022 (Computers & OR) |
| **Questionável** | ~5 | alexander2020 (EAI, rating 2), almufti2025 (IJSciWorld, rating 3), wadi2025 (IJCMEM, rating 3), araujo2025pso (arXiv s/ revisão), hossain2024 (Bitlis Eren Univ) |
| **Sem DOI/Veículo obscuro** | ~5 | oliver1987 (conferência s/ DOI), nagata2006 (LNCS — OK mas s/ DOI na entrada), deepaco2023, ppaco2024 (arXiv), sun2024hybrid (Procedia Computer Science) |

**Avaliação:** A proporção de artigos de primeira linha (~40%) é adequada para mestrado. O problema é que os artigos questionáveis estão concentrados na área de estudos comparativos — exatamente onde a qualidade das fontes é mais importante.

## 4. Viés de Seleção

O viés mais evidente é **ACO > PSO**. O vault tem 11 artigos ACO vs 6 PSO (razão 1.8:1), e destes, 4 ACO são de deep learning (deepaco2023, neufaco2025, ppaco2024, gpaco2025). Removendo os de DL, a razão cai para 7:6 (balanceada). O problema é que a percepção de desbalanceamento persiste no BibTeX e no index.md.

**Recomendação:** isolar os 4 artigos ACO+DL em categoria separada ou removê-los. Adicionar 1-2 artigos de PSO para TSP (ex.: araujo2025pso já está no BibTeX).

## 5. Rastreabilidade e Metadados

**Estado atual: INEXISTENTE.** O vault tem infraestrutura completa para rastreabilidade (papers.base com campos role, chapters, claim_support, view "Uso Na Monografia") mas **nenhum dado foi populado**. 84 notas, zero claims rastreáveis.

**Recomendação:** preencher chapters e role para todos os artigos citados na fundamentação (prioridade). Preencher claim_support para artigos de estudos comparativos que sustentam claims específicas.

## 6. Sumário de Achados

| ID | Severidade | Categoria |
|----|-----------|-----------|
| f-council-lit-001 | critical | Citation gap — lower bounds |
| f-council-lit-002 | high | Selection bias — ACO vs PSO |
| f-council-lit-003 | high | Citation quality — rating 0 |
| f-council-lit-004 | high | Citation gap — demsar2006 |
| f-council-lit-005 | high | Traceability — zero |
| f-council-lit-006 | medium | Selection bias — ACO+DL |
| f-council-lit-007 | medium | BibTeX integrity |
| f-council-lit-008 | medium | Publication quality |
| f-council-lit-009 | medium | Coverage gap — SA |
| f-council-lit-010 | medium | Reading coverage |
| f-council-lit-011 | low | Organization |
| f-council-lit-012 | low | Citation gap — exaustiva |
| f-council-lit-013 | low | Rating inconsistency |

**Veredito: FAIL.** A combinação de (a) seção inteira sem citações (2.8), (b) viés de seleção documentado, (c) artigos com rating 0 citados como autoridade, e (d) zero rastreabilidade entre claims e evidências torna a revisão de literatura inadequada para defesa no estado atual. As correções são factíveis e bem delimitadas — o material existe no vault, só não foi transferido ao BibTeX e ao texto.
