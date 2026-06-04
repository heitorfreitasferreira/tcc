```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "CONFIRMADO. Todas as 3 juízas independentemente chegaram a FAIL (HIGH). Meus 13 achados originais permanecem válidos, e o debate reforçou 3 dimensões adicionais que eu subestimava: (a) a lacuna de literatura de tuning de hiperparâmetros (Metod f-council-003) é metodologicamente grave porque afeta a validade de todas as conclusões experimentais; (b) o padrão de artigos rating 0 citados como autoridade (Lit f-council-lit-003) não é pontual (Bean 1994) mas sistêmico (winter2002, vanhove2012 também); (c) o diagnóstico de pipeline vault→BibTeX quebrado (Metod f-council-002, Lit f-council-lit-001) explica a causa-raiz das minhas falhas críticas f-001 e f-003.",
  "debate_notes": {
    "steel_man": {
      "source": "Gaps (Metodologista)",
      "finding_id": "f-council-003 (Metod)",
      "title": "Ausência de literatura sobre calibração de hiperparâmetros (tuning)",
      "my_R1_coverage": "NÃO abordado. Eu cataloguei 13 achados cobrindo lower bounds, implementação ACO, ATSP/TDTSP, feromônio 3D, Bean 1994, --gama, desbalanceamento, baselines, Nagata, AP vs HK, Aggarwal, Demšar, e Clerc — mas NÃO identifiquei a ausência de literatura de tuning como lacuna independente.",
      "why_compelling": "O experimento congela todos os parâmetros (pop=100, iter=100, α=1.0, β=2.0, ρ=0.2, w=0.7, c1=c2=2.0) para todas as 30 instâncias. Sem referências como irace (López-Ibáñez 2016), SMAC (Hutter 2011) ou F-Race (Birattari 2010), não há fundamentação metodológica para afirmar que os parâmetros escolhidos são razoáveis. O Metodologista corretamente observa que 'parâmetros fixos são uma escolha válida para comparação controlada, mas precisam ser defendidos com literatura'. Esta lacuna não é catalográfica (como os BibTeX faltantes) — é conceitual: a monografia não demonstra consciência de que tuning é uma subárea com metodologia própria. Eu deveria ter capturado isso como achado independente de severidade HIGH.",
      "severity_assessment": "HIGH. A ausência de literatura de tuning enfraquece a Seção 4.2 (Desenho Experimental) porque qualquer banca pode perguntar: 'por que estes parâmetros e não outros?' Sem respaldo na literatura de DoE/tuning, a resposta fica no terreno da intuição."
    },
    "challenge": {
      "target": "Gaps (Metodologista)",
      "finding_id": "f-council-008 (Metod)",
      "title": "5 artigos descartados com entradas BibTeX ativas — 'metadata contamination'",
      "my_position": "DISCORDO da severidade e da recomendação de remoção.",
      "reasoning": "O Metodologista classifica como MEDIUM e recomenda REMOVER as 5 entradas BibTeX de artigos classificados como 'descartado-escopo-imediato' (hga2024hybrid, huang2025matrix, kappagantula2025dpso, sun2024hybrid, toaza2023review). Discordo em dois pontos:\n\n1. O arquivo .bib não é o reference list final da monografia — é a base de dados bibliográfica do projeto. Manter entradas de artigos examinados e descartados documenta o processo de curadoria e evita que o autor (ou colaboradores futuros) reexaminem os mesmos artigos por engano.\n\n2. A recomendação de 'remover do BibTeX' resolve um não-problema ('risco de citação acidental') criando um problema real: perda de rastreabilidade do que foi considerado e por que foi excluído. Se o risco é citação acidental, a correção deveria ser garantir que não estão citados no texto — não apagar os registros.\n\nSeveridade real: LOW (informativo). A recomendação correta é: (a) verificar que nenhum artigo descartado é citado no texto da monografia; (b) documentar no vault POR QUE foram descartados (campo 'reason_excluded'); (c) opcionalmente adicionar um campo 'status: discarded' no BibTeX via comentário ou keyword.",
      "counter_proposal": "Manter as entradas BibTeX mas adicionar um campo `note = {Descartado: escopo imediato}` ou `keywords = {discarded}` em cada uma. Adicionar ao AGENTS.md uma política: 'artigos com keywords=discarded não devem ser citados na monografia sem reavaliação.'"
    },
    "acknowledge": {
      "source": "Gaps (Metodologista)",
      "finding_id": "f-council-002 (Metod)",
      "title": "Seção de lower bounds (2.8) sem citações — 10 artigos sem BibTeX — padrão sistêmico",
      "my_R1_coverage": "Meu f-council-001 identificou Held-Karp (1970, 1971) especificamente. Meu f-council-010 identificou a confusão conceitual AP vs HK. Mas eu NÃO cataloguei o padrão sistêmico: dos 10 artigos de lower bounds no vault, 0 têm BibTeX e 0 são citados.",
      "how_strengthens": "O Metodologista mostra que meu f-council-001 (Held-Karp sem BibTeX) não é um incidente isolado — é sintoma de um pipeline vault→BibTeX quebrado para TODA a área de lower bounds. Isso eleva a severidade do problema: não basta adicionar Held-Karp; é preciso adicionar todas as 10 entradas e revisar a Seção 2.8 inteira. O diagnóstico de 'padrão de duas velocidades' (seções centrais bem cobertas vs seções de suporte metodológico sem citações) é preciso e útil para priorização de correções."
    },
    "revised_findings_summary": {
      "confirmed_critical": [
        "f-council-001: Held-Karp sem BibTeX — CONFIRMADO e AGRAVADO pelo diagnóstico de padrão sistêmico (Metod f-002, Lit f-lit-001)",
        "f-council-002: ACO é Ant System, não MMAS — CONFIRMADO. Nenhum juiz contestou a evidência do código-fonte.",
        "f-council-003: ATSP/TDTSP ausentes — CONFIRMADO. Nenhum juiz contestou."
      ],
      "confirmed_high": [
        "f-council-004: Feromônio 3D sem documentação — CONFIRMADO. Reforçado pelo Metod f-004 (confundimento método×codificação).",
        "f-council-005: Bean 1994 ilegível — CONFIRMADO e GENERALIZADO pelo Lit f-lit-003 (padrão de rating 0 citados).",
        "f-council-006: --gama fantasma — CONFIRMADO. Nenhum juiz contestou.",
        "f-council-007: Desbalanceamento ACO 13 vs PSO 5 — CONFIRMADO e REFORÇADO pelo Lit f-lit-002 (shami2022pso e gad2022pso não citados)."
      ],
      "new_high_from_debate": [
        "f-council-R2-001 (originado do Metod f-003): Ausência de literatura de tuning de hiperparâmetros. Adiciono como HIGH porque afeta a validade de todas as conclusões experimentais. Sem irace/SMAC/F-Race no referencial, a escolha de parâmetros fixos não tem defesa metodológica.",
        "f-council-R2-002 (originado do Metod f-004): Confundimento método×codificação como ameaça à validade interna. Meu f-004 tratou do feromônio 3D como contribuição não documentada, mas o Metodologista corretamente generaliza: GA (permutação direta) vs PSO (random keys) vs ACO (feromônio 3D) significa que os efeitos de método e codificação não são isoláveis. Isso precisa ser discutido com Halim & Ismail (2019) e literatura de validade de constructo.",
        "f-council-R2-003 (originado do Lit f-lit-003): Padrão de artigos rating 0 citados como autoridade. Meu f-005 cobriu Bean (1994) isoladamente; o Lit revisor documenta que winter2002 e vanhove2012 (ambos rating 0) também são citados na Seção 2.2. Isso configura um padrão de qualidade de citação, não um incidente isolado."
      ],
      "confirmed_medium_low": [
        "f-council-008: Ausência baselines construtivas/SA — CONFIRMADO. Reforçado pelo Lit f-lit-009.",
        "f-council-009: Nagata EAX vazio — CONFIRMADO.",
        "f-council-010: Confusão AP vs HK — CONFIRMADO. Nenhum juiz contestou a imprecisão conceitual.",
        "f-council-011: Aggarwal sem BibTeX — CONFIRMADO. Reforçado pelo Lit f-lit-003 (aggarwal2000 é o artigo correto para turn costs, não winter2002/vanhove2012).",
        "f-council-012: Demšar sem nota — CONFIRMADO e AGRAVADO pelo Lit f-lit-004 (aplicabilidade a TSP não avaliada).",
        "f-council-013: Clerc citado mas não usado — CONFIRMADO."
      ]
    },
    "consensus_analysis": {
      "all_three_judges": "FAIL (HIGH)",
      "convergent_findings": [
        "Held-Karp sem BibTeX + Seção 2.8 sem citações (todos os 3 juízes)",
        "Desbalanceamento ACO vs PSO (todos os 3 juízes)",
        "Demšar sem nota no vault (todos os 3 juízes)",
        "Ausência de baselines construtivas/SA (Depth + Lit; Metod menciona indiretamente)"
      ],
      "divergent_findings": [
        "Metod f-008 (remover artigos descartados do BibTeX): DISCORDO (ver challenge acima)",
        "Metod f-005 (análise de sensibilidade angular): válido como issue experimental, mas classifico como fora do escopo de 'referencial teórico' — é lacuna de experimento, não de literatura"
      ],
      "strength_of_consensus": "FORTE. Três juízes com especialidades complementares (TSP/algoritmos, metodologia/estatística, revisão de literatura) chegaram independentemente a FAIL (HIGH). As sobreposições nos achados críticos (Held-Karp, Seção 2.8, desbalanceamento, Demšar) não são coincidência — são falhas estruturais visíveis de qualquer ângulo de análise."
    },
    "revised_recommendation": "REPROVADO — correções obrigatórias antes da qualificação/defesa, com 3 novos itens HIGH adicionados ao plano de correção original. Prioridade atualizada:\n\nBLOQUEADORAS (antes da qualificação):\n1. Adicionar Held-Karp (1970, 1971) ao BibTeX + os outros 8 artigos de lower bounds\n2. Corrigir documentação ACO: código é Ant System, não MMAS\n3. Adicionar 3-5 referências de ATSP/TDTSP\n\nALTAS (antes da versão final):\n4. [NOVO] Adicionar 2-3 referências de tuning (irace/SMAC/F-Race) e citar na Seção 4.2\n5. [NOVO] Discutir confundimento método×codificação com Halim & Ismail (2019) na Seção 4.2.6\n6. [NOVO] Corrigir ratings de winter2002, vanhove2012 (rating 0 citados como autoridade)\n7. Documentar feromônio 3D como contribuição\n8. Ler Bean (1994) — obter PDF legível\n9. Corrigir documentação do --gama\n10. Balancear referencial PSO (citar shami2022pso, gad2022pso; isolar ACO+DL)\n\nTempo adicional estimado para novos itens: +3 horas (total acumulado: 11-15 horas)."
  },
  "findings": [
    {
      "id": "f-council-001",
      "severity": "CRITICAL",
      "area": "lower-bounds",
      "title": "Held-Karp 1970 e 1971 ausentes do BibTeX — CONFIRMADO e AGRAVADO: padrão sistêmico de 10 artigos de lower bounds sem BibTeX",
      "detail": "Original: Held-Karp (1970, 1971) são os artigos canônicos sobre lower bounds Lagrangianos para TSP. Ambos fichados no vault (rating 5) mas sem BibTeX. AGRAVANTE R2: o Metodologista (f-002) e o Lit reviewer (f-lit-001) documentam que TODOS os 10 artigos de lower bounds no vault estão sem BibTeX e sem citação na Seção 2.8. Não é um incidente isolado — é um pipeline vault→BibTeX quebrado para a área inteira.",
      "fix": "Adicionar entradas BibTeX para os 10 artigos de lower bounds (heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, balas1985branch, fischetti1992additive, karp1979patching, kinable2017hybrid, righini2021efficient, aggarwal2000angular, lawler1985traveling). Revisar Seção 2.8 para citar Held-Karp como referência canônica e justificar a escolha da relaxação AP.",
      "ref": "monografia/bib/abntex2-references.bib; monografia/cap_fundamentacao/fundamentacao.tex:78-88"
    },
    {
      "id": "f-council-002",
      "severity": "CRITICAL",
      "area": "aco",
      "title": "Implementação ACO é Ant System, não MMAS — descasamento com literatura citada",
      "detail": "CONFIRMADO. Nenhum juiz contestou a evidência do código-fonte (src/optimization/aco/ant.go:14-39). A implementação usa depósito por todas as formigas (Ant System), sem bounds [τ_min, τ_max], sem reinicialização, e τ₀ = 1.0 (não τ_max). A nota vault/papers/stutzle2000mmas.md:55 afirma factualmente o contrário.",
      "fix": "Opção (a): implementar bounds MMAS no código. Opção (b): corrigir documentação para refletir Ant System + documentar 3D como contribuição sobre AS.",
      "ref": "src/optimization/aco/ant.go:14-39; vault/papers/stutzle2000mmas.md:55"
    },
    {
      "id": "f-council-003",
      "severity": "CRITICAL",
      "area": "atsp-tdtsp",
      "title": "Ausência total de referências sobre ATSP e TDTSP",
      "detail": "CONFIRMADO. O TSP-SD-ATP gera custos assimétricos e com dependência de sequência, enquadrando-se em ATSP e TDTSP. Zero entradas BibTeX para qualquer dessas classes.",
      "fix": "Adicionar survey de ATSP (Öncan et al. 2009 ou Roberti & Toth 2012), referência de TDTSP (Gouveia & Voß 1995 ou Kinable 2017), e Karp (1979).",
      "ref": "monografia/bib/abntex2-references.bib"
    },
    {
      "id": "f-council-R2-001",
      "severity": "HIGH",
      "area": "tuning",
      "title": "[NOVO — originado do Metod f-003] Ausência de literatura sobre calibração de hiperparâmetros",
      "detail": "O desenho experimental usa parâmetros fixos (pop=100, iter=100, α=1.0, β=2.0, ρ=0.2, w=0.7, c1=c2=2.0) para todas as instâncias. Não há referências a irace (López-Ibáñez 2016), SMAC (Hutter 2011), F-Race (Birattari 2010) ou qualquer metodologia de tuning/DoE. A monografia reconhece a limitação nos caps. 3-5 mas não fundamenta a escolha com literatura. Sem isso, a validade das conclusões experimentais é questionável: não se sabe se os rankings entre métodos são robustos a escolhas alternativas de parâmetros.",
      "fix": "Adicionar 2-3 referências de tuning (irace, SMAC, F-Race). Citar na Seção 4.2 para fundamentar parâmetros fixos como escolha metodológica deliberada (privilegiando reprodutibilidade sobre otimização por instância).",
      "ref": "monografia/cap_experimentos/experimentos.tex:31"
    },
    {
      "id": "f-council-R2-002",
      "severity": "HIGH",
      "area": "internal_validity",
      "title": "[NOVO — originado do Metod f-004] Confundimento método×codificação como ameaça à validade interna",
      "detail": "GA (permutação direta), PSO (random keys), ACO (feromônio 3D) usam codificações fundamentalmente diferentes. As diferenças de desempenho podem ser atribuídas tanto ao algoritmo quanto à representação — efeitos não isoláveis no desenho experimental. O Metodologista corretamente aponta que Halim & Ismail (2019), já no BibTeX, argumenta que a representação pode ser tão determinante quanto a meta-heurística e deveria ser citado nesta discussão. Meu f-004 original tratou do feromônio 3D isoladamente; o Metodologista generaliza para as três codificações.",
      "fix": "Citar Halim & Ismail (2019) na Seção 4.2.6. Adicionar 1-2 referências sobre validade de constructo em benchmarking de meta-heurísticas (Bartz-Beielstein 2006, Hooker 1995). Incluir parágrafo explicitando que a comparação mede o pacote método+codificação+parâmetros.",
      "ref": "monografia/cap_experimentos/experimentos.tex:164-168"
    },
    {
      "id": "f-council-R2-003",
      "severity": "HIGH",
      "area": "citation_quality",
      "title": "[NOVO — originado do Lit f-lit-003] Padrão de artigos com rating 0 citados como autoridade",
      "detail": "Meu f-005 original identificou Bean (1994) com rating 0 e PDF ilegível como problema isolado. O Lit reviewer documenta que winter2002 e vanhove2012 TAMBÉM têm rating 0 e são citados na Seção 2.2 como autoridade em custos de curva/mudança de direção. Três artigos com rating 0 (a pior classificação do vault) sendo usados como suporte a afirmações técnicas configura um padrão de qualidade de citação, não incidentes isolados. O artigo correto para a Seção 2.2 (Aggarwal 2000, rating 4) está no vault sem BibTeX.",
      "fix": "Substituir winter2002 e vanhove2012 por aggarwal2000angular (rating 4) como referência primária para custos angulares em TSP na Seção 2.2. Para Bean (1994): obter PDF legível ou adicionar referência alternativa de random keys. Adicionar política: artigos rating 0-1 não devem ser citados como autoridade primária.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:18,70; vault/papers/aggarwal2000angular.md"
    }
  ],
  "recommendation": "REPROVADO com 3 juízes em consenso unânime FAIL (HIGH). O debate da Rodada 2 adicionou 3 novos achados HIGH (tuning, confundimento método×codificação, padrão rating 0) e agravou 2 achados críticos existentes (Held-Karp agora reconhecido como padrão sistêmico de 10 artigos sem BibTeX; desbalanceamento ACO/PSO agora com evidência de surveys PSO não citados). As correções bloqueadoras (f-001, f-002, f-003) permanecem as mesmas. Tempo total estimado para todas as correções: 11-15 horas."
}
```
