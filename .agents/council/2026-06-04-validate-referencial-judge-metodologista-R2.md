```json
{
  "type": "verdict",
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "Confirmo FAIL com confiança reforçada. Meu R1 identificou 12 falhas (2 críticas, 3 altas, 5 médias, 2 baixas) centradas em lacunas catalográficas (Held-Karp sem BibTeX, Seção 2.8 sem citações) e fragilidades metodológicas (tuning, confundimento método×codificação, sensibilidade da função objetivo). Os achados do juiz TSP (Depth) revelam que o problema é mais grave do que eu diagnostiquei: a implementação ACO é Ant System, não MMAS como afirmado no vault (f-council-002-TSP), e o TSP-SD-ATP não está enquadrado em ATSP/TDTSP (f-council-003-TSP). Estes não são lacunas de citação — são erros factuais e omissões de enquadramento conceitual que eu não detectei. O juiz Lit (Coverage) reforça meu achado de desbalanceamento e adiciona evidência de que shami2022pso e gad2022pso existem no BibTeX mas não são citados — confirmando que o viés ACO/PSO é corrigível com material já disponível. Meu veredito FAIL se mantém, agora lastreado não apenas em falhas metodológicas/catalográficas mas também em erros factuais de implementação e lacunas de enquadramento teórico.",
  "findings": [
    {
      "id": "f-council-001",
      "title": "Held-Karp (1970) lido, rating 5, sem BibTeX — lacuna canônica em lower bounds",
      "severity": "critical",
      "category": "citation_gap",
      "description": "Mantido. Reforçado pelo juiz TSP (f-council-001-TSP) que confirma independentemente: os dois artigos (1970, 1971) estão no vault com rating 5 mas sem BibTeX. O juiz Lit (f-council-lit-001) também classifica como critical e acrescenta que lawler1985traveling (rating 5, TSP canônico) está na mesma situação.",
      "fix": "Adicionar entrada BibTeX para heldkarp1970traveling (DOI: 10.1287/opre.18.6.1138), heldkarp1971traveling, e lawler1985traveling. Citar na Seção 2.8.",
      "why": "Confirmado por 3 juízes independentes como falha crítica bloqueadora.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:78-88; vault/papers/heldkarp1970traveling.md"
    },
    {
      "id": "f-council-002",
      "title": "Seção 2.8 (Lower Bounds) sem citações — 10 artigos sem BibTeX",
      "severity": "critical",
      "category": "citation_gap",
      "description": "Mantido. O juiz Lit (f-council-lit-001) confirma: 'Numa monografia de mestrado, uma seção inteira da fundamentação sem suporte bibliográfico é inaceitável.' Ambos os juízes apontam o mesmo padrão: 10 artigos fichados, 0 citados, 0 BibTeX.",
      "fix": "Criar BibTeX para 9 artigos faltantes + citar heldkarp1970traveling, johnson1996asymptotic, valenzuela1997estimating.",
      "why": "Padrão sistêmico de não-conversão vault→BibTeX→citação.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:78-88"
    },
    {
      "id": "f-council-003",
      "title": "Ausência de literatura sobre calibração de hiperparâmetros (tuning)",
      "severity": "high",
      "category": "missing_coverage",
      "description": "Mantido. Nenhum outro juiz abordou este ponto diretamente, mas o juiz Lit (f-council-lit-013) nota que wang2021ant (tuning de parâmetros ACO) está no BibTeX com rating 0 — artigo relevante não aproveitado. Isto reforça meu achado: o material existe, mas não é usado.",
      "fix": "Adicionar 2-3 referências sobre tuning (irace, SMAC, F-Race) + reavaliar wang2021ant (rating 0→3) e citar.",
      "why": "Parâmetros fixos precisam de defesa metodológica com literatura.",
      "ref": "monografia/cap_experimentos/experimentos.tex:31"
    },
    {
      "id": "f-council-004",
      "title": "Ameaça à validade interna: confundimento método×codificação sem respaldo",
      "severity": "high",
      "category": "internal_validity",
      "description": "Mantido. O juiz TSP (f-council-013-TSP) aponta um caso concreto: Clerc (2000) PSO discreto é citado como referência mas o método implementado é random keys (Bean 1994) — duas abordagens fundamentalmente diferentes. Este é um exemplo preciso do confundimento que eu apontei em abstrato.",
      "fix": "Citar Halim e Ismail (2019) na discussão de limitações. Distinguir explicitamente PSO discreto (Clerc) de PSO com random keys (Bean) na Seção 2.6.",
      "why": "Evidência concreta do confundimento: Clerc vs Bean são métodos diferentes citados como se fossem equivalentes.",
      "ref": "monografia/cap_experimentos/experimentos.tex:164-168; vault/papers/clerc2000discretepso.md"
    },
    {
      "id": "f-council-005",
      "title": "Ausência de análise de sensibilidade da penalidade angular (Equação 4.1)",
      "severity": "high",
      "category": "missing_coverage",
      "description": "Mantido. O juiz TSP (f-council-011-TSP) adiciona contexto crucial: Aggarwal (2000) 'Angular-Metric TSP' — o artigo mais diretamente relevante para o TSP-SD-ATP — está no vault com rating 4 mas sem BibTeX. Este artigo prova resultados de aproximação para TSP com custo angular e poderia fundamentar ou calibrar a função objetivo.",
      "fix": "Adicionar BibTeX para aggarwal2000angular. Realizar análise de sensibilidade variando λ (peso do ângulo) ou justificar a normalização por π com base em Aggarwal (2000).",
      "why": "A função objetivo é a contribuição central. Aggarwal (2000) é a referência natural para validá-la e está ausente.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:22-27; vault/papers/aggarwal2000angular.md"
    },
    {
      "id": "f-council-006",
      "title": "Desbalanceamento ACO (13) vs PSO (5, sendo 3 descartados)",
      "severity": "medium",
      "category": "coverage_imbalance",
      "description": "AGRAVADO. O juiz Lit (f-council-lit-002) documenta que shami2022pso e gad2022pso — surveys abrangentes de PSO — estão no BibTeX mas não são citados. O juiz TSP (f-council-007-TSP) revela que 4 dos 13 artigos ACO são de ACO neural/RL (deepaco2023, neufaco2025, ppaco2024, gpaco2025), completamente desvinculados da implementação. O desbalanceamento real é pior: ACO efetivo = 9, PSO efetivo = 2 (apenas Kennedy 1995 e Clerc 2000, sendo que Clerc não é o método usado).",
      "fix": "Citar shami2022pso e gad2022pso na Seção 2.6. Isolar ACO+DL em categoria separada. Adicionar araujo2025pso (já no BibTeX) como referência de PSO para TSP.",
      "why": "O desbalanceamento é mais severo do que eu estimava: PSO tem efetivamente 1 artigo que descreve o método usado (Kennedy 1995).",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:64-69; vault/papers/index.md:53-68"
    },
    {
      "id": "f-council-007",
      "title": "Demšar (2006) sem nota no vault",
      "severity": "medium",
      "category": "vault_integrity",
      "description": "Mantido. Confirmado pelos juízes TSP (f-council-012-TSP) e Lit (f-council-lit-004). O juiz Lit acrescenta uma observação metodológica importante: Demšar (2006) é focado em classificadores de ML, não em meta-heurísticas de otimização — a extensão para TSP merece justificativa crítica que não existe sem a nota.",
      "fix": "Criar vault/papers/demsar2006statistical.md com avaliação crítica da adequação ao contexto TSP.",
      "why": "O teste estatístico que sustenta todas as comparações não teve seu artigo de referência avaliado criticamente.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:95"
    },
    {
      "id": "f-council-008",
      "title": "5 artigos descartados com entradas BibTeX ativas",
      "severity": "medium",
      "category": "metadata_contamination",
      "description": "Mantido. Nenhum outro juiz abordou este ponto. A inconsistência vault↔BibTeX persiste como problema de higiene catalográfica.",
      "fix": "Remover ou reclassificar entradas BibTeX de artigos descartados.",
      "ref": "monografia/bib/abntex2-references.bib"
    },
    {
      "id": "f-council-009",
      "title": "Metadados de uso (role, chapters, methods, claim_support) vazios em 100%",
      "severity": "medium",
      "category": "vault_integrity",
      "description": "Mantido. O juiz Lit (f-council-lit-005) confirma: 'NENHUM dos 84 artigos no vault possui os campos preenchidos. A view Uso Na Monografia está completamente vazia.' Classificado como HIGH pelo juiz Lit — rastreabilidade zero é um problema de auditoria acadêmica.",
      "fix": "Preencher role e chapters para ~25 artigos principais.",
      "ref": "vault/bases/papers.base:80-88"
    },
    {
      "id": "f-council-010",
      "title": "4 artigos com rating 0 citados na monografia",
      "severity": "medium",
      "category": "metadata_inconsistency",
      "description": "AGRAVADO. O juiz TSP (f-council-005-TSP) revela que bean1994genetic — rating 0 — tem PDF ilegível: 'a nota do vault reconhece que não foi possível confirmar detalhes experimentais, pseudocódigo ou conclusões.' O PSO implementado depende de um artigo cujo conteúdo o autor não conseguiu ler. O juiz Lit (f-council-lit-003) classifica como HIGH: 'Citar artigos com rating 0 mina a credibilidade da revisão.'",
      "fix": "Obter cópia legível de Bean (1994). Revisar ratings: bean1994 ≥ 3, winter2002 ≥ 3, vanhove2012 ≥ 3, wang2021ant ≥ 3.",
      "why": "Bean (1994) com rating 0 e PDF ilegível é indefensável como fundamentação do PSO.",
      "ref": "vault/papers/bean1994genetic.md"
    },
    {
      "id": "f-council-011",
      "title": "Poder estatístico do Friedman com N=30 e k=3 não discutido",
      "severity": "low",
      "category": "statistical_validity",
      "description": "Mantido. Sem novas evidências dos outros juízes.",
      "fix": "Adicionar nota sobre poder do teste com N=30, k=3, citando Demšar (2006, Seção 3.2).",
      "ref": "monografia/cap_experimentos/experimentos.tex:152-162"
    },
    {
      "id": "f-council-012",
      "title": "Validade externa: apenas instâncias sintéticas sem discussão de generalização",
      "severity": "low",
      "category": "external_validity",
      "description": "Mantido. Sem novas evidências.",
      "fix": "Adicionar referência sobre validade ecológica em otimização combinatória.",
      "ref": "monografia/cap_experimentos/experimentos.tex:164-168"
    },
    {
      "id": "f-council-013",
      "title": "NOVO: Implementação ACO é Ant System, não MMAS — erro factual no vault",
      "severity": "critical",
      "category": "implementation_error",
      "source": "TSP judge f-council-002",
      "description": "O código em src/optimization/aco/ant.go:14-39 implementa Ant System (todas as formigas depositam, sem bounds de feromônio). O vault (stutzle2000mmas.md:55) afirma que MMAS está implementado — isto é factualmente falso. A nota do vault foi escrita sem verificação do código-fonte.",
      "fix": "Opção A: implementar bounds MMAS no código. Opção B: corrigir documentação para refletir Ant System e remover afirmação falsa do vault.",
      "why": "Erro factual independente que compromete a correção da seção metodológica. Não detectado no meu R1.",
      "ref": "src/optimization/aco/ant.go:14-39; vault/papers/stutzle2000mmas.md:55"
    },
    {
      "id": "f-council-014",
      "title": "NOVO: ATSP e TDTSP completamente ausentes do referencial",
      "severity": "critical",
      "category": "framing_gap",
      "source": "TSP judge f-council-003",
      "description": "O TSP-SD-ATP gera custos assimétricos (c'(j,k) ≠ c'(k,j)) e com dependência de sequência. Isto o enquadra em ATSP e TDTSP — duas classes bem estudadas. Zero referências para qualquer uma delas no BibTeX ou vault. Kinable (2017) cobre TDTSP mas está no vault sem BibTeX.",
      "fix": "Adicionar survey de ATSP + referência de TDTSP + BibTeX para Kinable (2017) e Karp (1979).",
      "why": "Sem enquadramento em ATSP/TDTSP, o TSP-SD-ATP parece inventado ad-hoc. Lacuna de framing que eu não detectei.",
      "ref": "monografia/bib/abntex2-references.bib (zero entradas ATSP/TDTSP)"
    },
    {
      "id": "f-council-015",
      "title": "NOVO: Relaxação AP confundida com Held-Karp na nota do vault",
      "severity": "high",
      "category": "conceptual_error",
      "source": "TSP judge f-council-010",
      "description": "A nota heldkarp1970traveling.md:51-52 afirma que 'a redução c'[j][k] = min_i cost[i][j][k] permite aplicar o HK bound sobre uma matriz 2D reduzida.' Isto é conceitualmente impreciso: o bound de Held-Karp usa 1-trees com multiplicadores Lagrangianos; a redução min_i + Hungarian produz um bound AP — relacionado mas diferente e mais fraco.",
      "fix": "Corrigir a nota do vault. Na monografia, distinguir claramente relaxação AP de Held-Karp. Incluir análise do gap AP vs melhor solução.",
      "why": "Confusão conceitual que pode se propagar do vault para a monografia. Reforça meu f-council-001 e f-council-002.",
      "ref": "vault/papers/heldkarp1970traveling.md:51-52; src/optimization/lowerbound/main.go:39-57"
    },
    {
      "id": "f-council-016",
      "title": "NOVO: Feromônio 3D é inovação sem lastro — não documentado como contribuição",
      "severity": "high",
      "category": "documentation_gap",
      "source": "TSP judge f-council-004",
      "description": "O ACO usa feromônio tridimensional τ(i,j,k). Na literatura canônica de ACO para TSP, o feromônio é sempre 2D. Esta adaptação é uma contribuição algorítmica genuína, mas não está documentada como tal na monografia e não tem precedentes na literatura citada.",
      "fix": "Documentar feromônio 3D como contribuição própria. Justificar pela estrutura do TSP-SD-ATP (dependência de triplas).",
      "why": "O feromônio 3D é a característica mais distintiva da implementação. Não documentá-lo é perder a oportunidade de reivindicar novidade.",
      "ref": "src/optimization/aco/main.go:86-97"
    },
    {
      "id": "f-council-017",
      "title": "NOVO: Artigos ACO+Deep Learning (4) inflam cobertura de ACO",
      "severity": "medium",
      "category": "selection_bias",
      "source": "TSP judge f-council-007; Lit judge f-council-lit-006",
      "description": "deepaco2023, neufaco2025, ppaco2024, gpaco2025 são de ACO neural/RL, sem relação com a implementação Ant System 3D. Ocupam 4/13 posições da categoria ACO. Ambos os juízes recomendam isolamento em categoria separada.",
      "fix": "Mover para categoria 'ACO/Deep Learning — Tendências' ou remover do escopo comparativo.",
      "why": "Inflação artificial da cobertura ACO que agrava o desbalanceamento com PSO.",
      "ref": "vault/papers/index.md:31-34"
    }
  ],
  "debate_notes": {
    "steel_man": "O achado mais forte que eu não foquei no R1 é o f-council-002 do juiz TSP: a implementação ACO é Ant System, não MMAS. Meu R1 concentrou-se em lacunas catalográficas (BibTeX, citações) e fragilidades metodológicas (tuning, confundimento). O juiz TSP foi ao código-fonte e descobriu um erro factual: a nota do vault afirma incorretamente que MMAS está implementado, mas o código (ant.go:14-39) mostra depósito de todas as formigas, sem bounds de feromônio, sem reinicialização — caracterizando Ant System básico. Este não é um problema de 'falta citação' — é um problema de 'a documentação mente sobre o que o código faz'. Isto é mais grave do que qualquer lacuna catalográfica porque compromete a correção técnica da monografia. Se implementado, este achado eleva a severidade do problema ACO de 'desbalanceamento de cobertura' (meu f-council-006, severidade medium) para 'erro factual na fundamentação' (severidade critical).",
    "challenge": "Discordo parcialmente do juiz TSP quanto à severidade do f-council-008-TSP (ausência de SA e heurísticas construtivas como baselines). O juiz TSP classifica como MEDIUM e descreve como lacuna que 'enfraquece a pergunta de pesquisa'. Argumento que esta não é uma lacuna, mas uma decisão de escopo legítima. O TCC se propõe explicitamente a comparar três métodos bio-inspirados populacionais (GA, PSO, ACO). Simulated Annealing é um método de trajetória única (não populacional, não bio-inspirado no sentido de enxame/colônia). Heurísticas construtivas (nearest neighbor, Christofides) são determinísticas e não-iterativas — uma classe completamente diferente. Adicioná-las como baselines transformaria a pergunta de pesquisa de 'qual método bio-inspirado funciona melhor?' para 'métodos bio-inspirados são melhores que busca local?', que é uma pergunta diferente. Minha recomendação: reclassificar como LOW. A monografia já reconhece esta limitação na Seção 4.2.6. A adição de SA como baseline seria um experimento adicional interessante, mas sua ausência não enfraquece a validade interna da comparação GA vs PSO vs ACO.",
    "acknowledge": "O juiz TSP (f-council-003-TSP) identificou uma lacuna de enquadramento conceitual que eu não percebi: o TSP-SD-ATP, com tensor 3D e redução min_i, gera um problema fundamentalmente assimétrico (ATSP) e com dependência de sequência (TDTSP). Meu R1 focou em 'faltam citações de lower bounds' — mas o problema é mais profundo: falta o enquadramento do próprio problema nas classes de TSP que o definem. Sem referências a ATSP e TDTSP, o TSP-SD-ATP parece um problema inventado ad-hoc, quando é uma instância natural de classes bem estudadas. Este achado fortalece meu f-council-001 e f-council-002 ao revelar que não é apenas a seção de lower bounds que está descontextualizada — é a própria definição do problema que carece de enquadramento na literatura de otimização combinatória. Adicionalmente, o juiz TSP (f-council-010-TSP) descobriu um erro conceitual na nota do vault: a afirmação de equivalência entre redução-por-min + Hungarian e o bound de Held-Karp é tecnicamente incorreta. Isto é uma evidência concreta de que meu diagnóstico de 'gargalo catalográfico' estava incompleto — o problema não é apenas que as entradas BibTeX não foram criadas, mas que o conteúdo das notas do vault contém imprecisões técnicas que precisam ser corrigidas antes da citação.",
    "verdict_shift": false,
    "verdict_rationale": "Confirmo FAIL (HIGH) com confiança reforçada. Meu R1 identificou 12 falhas centradas em lacunas catalográficas e fragilidades metodológicas. O debate revelou que o problema é mais grave em três dimensões que eu subestimei: (1) erro factual — a implementação ACO não corresponde à literatura citada (Ant System ≠ MMAS); (2) lacuna de enquadramento — ATSP e TDTSP completamente ausentes, deixando o TSP-SD-ATP sem ancoragem conceitual; (3) imprecisão técnica — a nota do vault sobre Held-Karp contém um erro conceitual (confunde relaxação AP com bound HK). Estas três novas falhas são independentes das minhas 12 originais e cada uma justificaria FAIL individualmente. A combinação de falhas catalográficas (meu R1), erros factuais de implementação (juiz TSP) e falta de enquadramento teórico (juiz TSP) torna o referencial não defensável no estado atual. O lado positivo é que todas as falhas são corrigíveis com trabalho focado: as catalográficas (~3h), as de implementação/documentação (~4h), e as de enquadramento (~2h). Nenhuma requer nova coleta de dados ou reimplementação completa. O material existe — precisa ser corrigido, completado e integrado."
  },
  "recommendation": "REPROVAR com exigência de correções antes da defesa. O plano de correção do R1 permanece válido e deve ser expandido para incluir: (P0) corrigir documentação ACO (Ant System vs MMAS) e nota do vault; (P0) adicionar referências de ATSP/TDTSP; (P0) corrigir imprecisão conceitual AP vs HK na nota do vault; (P0) adicionar BibTeX para aggarwal2000angular (artigo canônico sobre TSP angular). As 12 falhas originais do R1 mais as 5 novas totalizam 17 achados, dos quais 4 são críticos e 5 são de alta severidade."
}
```

---

# R2 — Juiz Metodologista (Gaps): Debate Round 2

## 1. RESTATE: Minha Posição no R1

**Veredito R1: FAIL (HIGH)**. Identifiquei 12 falhas organizadas em três eixos:

1. **Lacunas catalográficas (2 críticas):** Held-Karp (1970) lido, rating 5, sem BibTeX (f-council-001); Seção 2.8 (Lower Bounds) inteira sem citações — 10 artigos fichados, 0 referenciados (f-council-002). Diagnóstico: gargalo no pipeline vault→BibTeX→monografia.

2. **Fragilidades metodológicas (3 altas):** Ausência de literatura de tuning (f-council-003); confundimento método×codificação sem respaldo (f-council-004); sensibilidade não analisada da função objetivo (f-council-005). Diagnóstico: o desenho experimental carece de fundamentação bibliográfica para suas escolhas de design.

3. **Problemas de integridade e rastreabilidade (5 médias, 2 baixas):** Desbalanceamento ACO×PSO, Demšar sem nota, artigos descartados com BibTeX, metadados vazios, rating 0 em citados, poder estatístico não discutido, validade externa sem respaldo.

---

## 2. STEEL-MAN: O Melhor Argumento Que Eu Não Foquei

**f-council-002-TSP (Depth): A implementação ACO é Ant System, não MMAS — erro factual no vault.**

Este é o achado mais devastador do debate. Meu R1 concentrou-se no que *falta* (citações, BibTeX, literatura de tuning). O juiz TSP foi ao código-fonte (`src/optimization/aco/ant.go:14-39`) e descobriu o que está *errado*:

- **Código:** `updatePheromones` deposita feromônio de TODAS as formigas (`for _, ant := range ants`), sem seleção de iteration-best ou global-best. Não há bounds `[τ_min, τ_max]`. Feromônio inicial = 1.0 (não τ_max). Não há reinicialização por estagnação.
- **Vault:** `stutzle2000mmas.md:55` afirma textualmente que "A fórmula τ_max = 1/(ρ·L_best) e τ_min = τ_max / (2n) é implementada diretamente no código ACO do TCC".
- **Fato:** A afirmação do vault é falsa. A implementação é um Ant System 3D básico (Dorigo 1996), não MMAS (Stützle 2000).

Este não é um problema catalográfico. É um **erro factual** que mina a credibilidade da seção metodológica. Se a banca examinadora lesse o código e comparasse com o texto, identificaria a discrepância em minutos. A gravidade está em: (a) a documentação mente sobre a implementação; (b) a nota do vault foi escrita sem verificação do código; (c) a justificativa algorítmica da monografia referencia o artigo errado para o que foi implementado.

---

## 3. CHALLENGE: Contestação de um Achado de Outro Juiz

**Discordo da classificação do f-council-008-TSP (Depth) como MEDIUM: "Ausência de heurísticas construtivas e Simulated Annealing como baselines."**

Argumento que esta severidade deve ser reduzida para **LOW**, pelas seguintes razões:

1. **Decisão de escopo, não lacuna.** O TCC se propõe explicitamente a comparar três métodos **bio-inspirados populacionais**: GA, PSO e ACO. SA é um método de **trajetória única**, não populacional e não bio-inspirado no sentido de enxame/colônia. Heurísticas construtivas (nearest neighbor, Christofides) são **determinísticas e não-iterativas** — uma classe fundamentalmente diferente. Adicioná-las transformaria a pergunta de pesquisa de "qual método bio-inspirado funciona melhor para o TSP-SD-ATP?" para "métodos bio-inspirados são melhores que busca local para o TSP-SD-ATP?" — que é uma pergunta diferente.

2. **A monografia já reconhece a limitação.** A Seção 4.2.6 discute ameaças à validade e reconhece que o escopo é limitado a métodos populacionais. A omissão de SA não é um erro — é uma boundary condition declarada.

3. **SA como baseline não é consenso metodológico.** Em surveys comparativos de meta-heurísticas para TSP, SA aparece como um entre muitos métodos, não como baseline obrigatório. Halim (2019) e Chandra (2022) incluem SA, mas também incluem Tabu Search, GRASP, VNS, etc. — não se pode exigir todos. A escolha de quais métodos incluir é do pesquisador, desde que justificada.

4. **Custo de escopo.** Adicionar SA como baseline exigiria: (a) implementação de SA para TSP-SD-ATP 3D (não-trivial — requer adaptação da função de vizinhança para custos dependentes de triplas); (b) tuning de parâmetros de SA (temperatura inicial, cooling schedule); (c) re-execução de todos os experimentos; (d) re-análise estatística com k=4. Isso expandiria o escopo em semanas.

**Recomendação:** Reclassificar f-council-008-TSP como LOW. A ausência de SA e heurísticas construtivas é uma limitação de escopo válida, não uma falha metodológica. Pode ser mencionada como "trabalho futuro" sem comprometer a validade da comparação GA vs PSO vs ACO.

---

## 4. ACKNOWLEDGE: Pontos Que Fortalecem Minha Análise

### 4.1 ATSP/TDTSP ausentes — lacuna de enquadramento (f-council-003-TSP)

O juiz TSP identificou que o TSP-SD-ATP é fundamentalmente um **ATSP** (custos assimétricos após redução min_i) e um **TDTSP** (custos dependentes de sequência), mas **nenhuma** dessas classes tem qualquer referência no BibTeX ou vault. Meu R1 focou em "faltam citações na Seção 2.8" — o juiz TSP mostrou que o problema é mais profundo: **falta enquadrar o próprio problema** nas classes de TSP que o definem.

Isto fortalece meus f-council-001 e f-council-002: não é apenas que a seção de lower bounds está sem citações; é que o TSP-SD-ATP está conceitualmente desancorado da literatura de otimização combinatória. Sem ATSP/TDTSP, o problema parece inventado ad-hoc.

### 4.2 Confusão AP vs Held-Karp na nota do vault (f-council-010-TSP)

O juiz TSP descobriu um erro conceitual preciso na nota `heldkarp1970traveling.md:51-52`: a afirmação de que "a redução c'[j][k] = min_i cost[i][j][k] permite aplicar o HK bound sobre uma matriz 2D reduzida" é tecnicamente incorreta. O bound de Held-Karp usa 1-trees com multiplicadores Lagrangianos; a redução min_i + Hungarian produz um bound de Assignment Problem — relacionado, mas diferente e mais fraco.

Isto confirma meu diagnóstico de que o gargalo não é apenas catalográfico — há **imprecisões técnicas nas notas do vault** que precisam ser corrigidas antes que o conteúdo seja transferido para a monografia. Meu R1 assumiu que as notas estavam corretas e só faltava o BibTeX; o juiz TSP mostrou que algumas notas contêm erros.

### 4.3 shami2022pso e gad2022pso no BibTeX mas não citados (f-council-lit-002)

O juiz Lit documentou que duas surveys abrangentes de PSO (IEEE Access e Archives of Computational Methods) estão no BibTeX mas **não são citadas na Seção 2.6**. Isto confirma que meu f-council-006 (desbalanceamento) é corrigível com material já disponível — não requer nova busca bibliográfica, apenas integração do que já existe.

---

## 5. VERDICT: CONFIRMADO — FAIL (HIGH)

**Meu veredito NÃO muda.** Confirmo FAIL com confiança HIGH, agora reforçada por evidências que eu não havia considerado no R1.

### 5.1 Por que mantenho FAIL

Meu R1 já justificava FAIL com 2 falhas críticas (Held-Karp sem BibTeX, Seção 2.8 sem citações) e 3 falhas altas (tuning, confundimento, sensibilidade). O debate adicionou **3 novas falhas independentes** que, sozinhas, também justificariam FAIL:

| Nova Falha | Fonte | Severidade | Natureza |
|---|---|---|---|
| ACO é Ant System, não MMAS (erro factual) | TSP f-council-002 | CRITICAL | Código não corresponde à documentação |
| ATSP/TDTSP completamente ausentes | TSP f-council-003 | CRITICAL | Problema sem enquadramento conceitual |
| Confusão AP vs Held-Karp na nota do vault | TSP f-council-010 | HIGH | Erro conceitual em nota de referência |
| Feromônio 3D não documentado como contribuição | TSP f-council-004 | HIGH | Inovação não reivindicada |
| ACO+DL infla cobertura | TSP f-council-007, Lit f-council-lit-006 | MEDIUM | Viés de seleção |

O referencial agora acumula **4 falhas críticas** (2 minhas + 2 do juiz TSP) e **5 falhas altas** (3 minhas + 2 do juiz TSP). Nenhuma é subjetiva — todas são verificáveis em arquivos específicos.

### 5.2 O que o debate revelou sobre meu R1

Meu R1 estava **correto no diagnóstico, mas incompleto no escopo**. Eu identifiquei o padrão vault→BibTeX quebrado, mas não fui ao código verificar se o que o vault afirma corresponde ao que o código implementa. O juiz TSP fez essa verificação e encontrou discrepâncias. Minhas falhas eram majoritariamente sobre *ausências* (o que falta); as falhas do juiz TSP são sobre *incorreções* (o que está errado). Ambas as classes justificam FAIL.

### 5.3 Plano de correção expandido

| Prioridade | Ação | Esforço | Fonte |
|---|---|---|---|
| P0 | Corrigir documentação ACO: Ant System ou implementar MMAS | 2-4h | TSP f-council-002 |
| P0 | Adicionar referências ATSP/TDTSP + BibTeX Kinable, Karp | 2h | TSP f-council-003 |
| P0 | Corrigir imprecisão AP vs HK na nota do vault | 30min | TSP f-council-010 |
| P0 | Adicionar BibTeX Held-Karp, Johnson, Aggarwal + citar | 1h | Meu f-council-001, TSP f-council-011 |
| P1 | Documentar feromônio 3D como contribuição própria | 1h | TSP f-council-004 |
| P1 | Adicionar literatura de tuning + citar | 1h | Meu f-council-003 |
| P1 | Criar nota Demšar (2006) | 20min | Meu f-council-007 |
| P2 | Citar shami2022pso, gad2022pso na Seção 2.6 | 30min | Lit f-council-lit-002 |
| P2 | Isolar ACO+DL em categoria separada | 15min | TSP f-council-007 |
| P2 | Obter PDF legível de Bean (1994), atualizar rating | 2h | TSP f-council-005 |

**Tempo total P0+P1 (mínimo para aprovação): ~8-10 horas.** Todas as correções são factíveis e bem delimitadas.

---

*Juiz Metodologista (Gaps) — R2 concluído. Veredito: FAIL (HIGH), confirmado com 17 achados (4 críticos, 5 altos, 6 médios, 2 baixos).*
