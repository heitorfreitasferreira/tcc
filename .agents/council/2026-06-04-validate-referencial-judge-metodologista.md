```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "O referencial tem cobertura sólida em TSP clássico, GA e ACO, mas apresenta quatro falhas metodológicas graves: (1) ausência do Held-Karp (1970) — lido, rating 5, sem BibTeX — na seção de lower bounds, impossibilitando situar a relaxação AP na literatura canônica; (2) toda a seção de lower bounds (2.8) carece de citações — 10 artigos mapeados no vault sem BibTeX; (3) ausência completa de literatura sobre tuning/sensibilidade/DoE/FLA, limitando o desenho experimental a parâmetros fixos sem fundamentação; (4) ameaça à validade interna por confundimento entre método e codificação (permutação direta vs random keys vs feromônio 3D) sem discussão respaldada por literatura.",
  "findings": [
    {
      "id": "f-council-001",
      "title": "Held-Karp (1970) lido, rating 5, sem BibTeX — lacuna canônica em lower bounds",
      "severity": "critical",
      "category": "citation_gap",
      "description": "heldkarp1970traveling tem nota no vault com rating 5, status 'lido', PDF íntegro — é o artigo fundacional do lower bound mais utilizado para TSP simétrico (gap empírico < 0.8%). No entanto, NÃO possui entrada BibTeX em abntex2-references.bib. A seção 2.8 da monografia descreve a relaxação AP/Hungarian mas nunca menciona Held-Karp, o que impede o leitor de entender como a escolha da relaxação se relaciona com o estado da arte em bounds para TSP.",
      "fix": "Adicionar entrada BibTeX para heldkarp1970traveling em monografia/bib/abntex2-references.bib (DOI: 10.1287/opre.18.6.1138). Citar na Seção 2.8 como referência canônica para lower bounds de TSP, explicando por que a relaxação AP foi escolhida em vez da relaxação Lagrangiana/Held-Karp para o TSP-SD-ATP.",
      "why": "A nota do vault já contém o resumo e a relevância para o TCC. A entrada BibTeX simplesmente não foi criada. Sem esta citação, a seção de lower bounds fica descontextualizada da literatura de referência.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:78-88; vault/papers/heldkarp1970traveling.md"
    },
    {
      "id": "f-council-002",
      "title": "Seção de lower bounds (2.8) sem citações — 10 artigos sem BibTeX",
      "severity": "critical",
      "category": "citation_gap",
      "description": "A seção 2.8 descreve a relaxação AP como lower bound mas não cita nenhum trabalho da literatura de bounds para TSP. Dos 10 artigos mapeados no vault sobre lower bounds, NENHUM possui entrada BibTeX: heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, balas1985branch, fischetti1992additive, karp1979patching, kinable2017hybrid, righini2021efficient, aggarwal2000angular. valenzuela1997estimating tem BibTeX mas não é citado. A seção menciona 'subtours' sem referenciar a formulação LP do TSP (subtour elimination constraints), que é central na literatura.",
      "fix": "Criar entradas BibTeX para os 9 artigos sem entrada. Citar pelo menos heldkarp1970traveling (canônico), johnson1996asymptotic (validação empírica do HK bound) e valenzuela1997estimating (implementação prática). Discutir por que a relaxação AP foi adotada em vez da relaxação Lagrangiana para o TSP-SD-ATP tridimensional.",
      "why": "A relaxação AP é uma escolha de design metodológico que precisa ser justificada em relação ao estado da arte. Sem citações, o leitor não tem referência para avaliar se o bound proposto é uma simplificação razoável ou uma omissão de alternativas melhores.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:78-88"
    },
    {
      "id": "f-council-003",
      "title": "Ausência de literatura sobre calibração de hiperparâmetros (tuning)",
      "severity": "high",
      "category": "missing_coverage",
      "description": "O desenho experimental usa parâmetros fixos (pop=100, iter=100, alpha=1.0, beta=2.0, rho=0.2, w=0.7, c1=c2=2.0) para todas as instâncias e métodos. A monografia reconhece essa limitação nos caps. 3, 4 e 5, mas não referencia nenhum artigo sobre metodologia de tuning de meta-heurísticas (e.g., irace, SMAC, grid search, Bayesian optimization, F-Race). A ausência dessa literatura enfraquece a justificativa de que os parâmetros escolhidos são razoáveis e torna a comparação vulnerável à crítica de que os métodos não foram otimizados igualmente.",
      "fix": "Adicionar 2-3 referências sobre tuning de meta-heurísticas (e.g., Birattari et al. 2010 — F-Race; Hutter et al. 2011 — SMAC; Lopez-Ibanez et al. 2016 — irace). Citar na Seção 4.1 ou 4.2 para fundamentar a escolha de parâmetros fixos como decisão metodológica deliberada, reconhecendo a trade-off entre reprodutibilidade e otimização por instância.",
      "why": "Parâmetros fixos são uma escolha válida para comparação controlada, mas precisam ser defendidos com literatura. Sem respaldo, a escolha parece arbitrária e abre flanco para a crítica de que os resultados poderiam ser diferentes com tuning.",
      "ref": "monografia/cap_experimentos/experimentos.tex:31; monografia/cap_proposta/proposta.tex:31,57"
    },
    {
      "id": "f-council-004",
      "title": "Ameaça à validade interna: confundimento método x codificação sem respaldo",
      "severity": "high",
      "category": "internal_validity",
      "description": "Os três métodos usam codificações distintas: GA usa permutação direta, PSO usa random keys (Bean 1994), ACO usa feromônio 3D. A monografia reconhece isso como limitação, mas não cita literatura sobre validade de constructo em comparações de meta-heurísticas (e.g., que a representação pode ser tão determinante quanto o algoritmo — Halim & Ismail 2019, já citado, poderia ser usado para respaldar essa discussão). A comparação atribui diferenças de desempenho ao 'método', mas os efeitos de codificação, operadores e hiperparâmetros não são isolados.",
      "fix": "Citar Halim e Ismail (2019) explicitamente na discussão de limitações (já está no .bib). Adicionar 1-2 referências sobre validade de constructo em benchmarking de meta-heurísticas (e.g., Bartz-Beielstein 2006; Hooker 1995). Incluir um parágrafo na Seção 4.2.6 discutindo que a comparação mede o pacote método+codificação+parâmetros, não a meta-heurística isoladamente.",
      "why": "O confundimento entre método e codificação é a ameaça mais séria à validade interna. Sem discuti-lo com respaldo bibliográfico, as conclusões sobre superioridade de ACO são metodologicamente frágeis.",
      "ref": "monografia/cap_experimentos/experimentos.tex:164-168"
    },
    {
      "id": "f-council-005",
      "title": "Ausência de análise de sensibilidade da penalidade angular",
      "severity": "high",
      "category": "missing_coverage",
      "description": "A Equação 4.1 (c_{i,j,k} = d(j,k) + theta(i,j,k)/pi) é contribuição própria do autor. A normalização por pi é uma escolha de design que coloca distância e ângulo na mesma escala. Não há análise de sensibilidade para essa ponderação implícita (distância:ângulo = 1:1/pi). Não há referências sobre como calibrar ou validar funções de custo compostas com penalidades angulares em roteamento.",
      "fix": "Adicionar uma análise de sensibilidade simples variando o peso da penalidade angular (e.g., lambda em {0, 0.5, 1.0, 2.0}) em uma instância representativa. Ou, alternativamente, citar literatura sobre modelagem de custos de manobra em roteamento para justificar a escolha da normalização.",
      "why": "A função objetivo é o núcleo da contribuição do TCC. Sem análise de sensibilidade do seu parâmetro de projeto (peso do ângulo), não se sabe se os rankings entre métodos são robustos a essa escolha.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:22-27"
    },
    {
      "id": "f-council-006",
      "title": "Desbalanceamento de cobertura: ACO (13) vs PSO (5, sendo 3 descartados)",
      "severity": "medium",
      "category": "coverage_imbalance",
      "description": "A seção de ACO referencia 13 artigos, incluindo fundamentais (Dorigo x5, Stützle, Blum) e avançados (neural/RL ACO). A seção de PSO referencia apenas 5 artigos, dos quais 3 (hga2024, huang2025, kappagantula2025) estão classificados como 'descartado-escopo-imediato' no vault. Sobram apenas Kennedy (1995) e Clerc (2000) como referências ativas. O desbalanceamento pode enviesar a discussão, fazendo o ACO parecer mais bem fundamentado que o PSO.",
      "fix": "Remover artigos descartados da lista de referenciáveis da seção PSO. Adicionar 2-3 artigos sobre PSO discreto para TSP (e.g., Shi et al. 2007; Chen et al. 2011; Zhong et al. 2012). Alternativamente, reduzir a cobertura de ACO para equilibrar ou justificar o desbalanceamento como reflexo da maturidade relativa das áreas.",
      "why": "Uma comparação justa exige que cada método tenha cobertura bibliográfica proporcional. O leitor pode questionar se o mau desempenho do PSO se deve à falta de exploração da literatura sobre variantes discretas.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:64-69"
    },
    {
      "id": "f-council-007",
      "title": "Demsar (2006) sem nota no vault — lacuna na base de conhecimento",
      "severity": "medium",
      "category": "vault_integrity",
      "description": "Demsar (2006) é a referência fundacional do protocolo estatístico da monografia, possui entrada BibTeX, mas NÃO tem nota no vault. Isso significa que o artigo não passou pelo processo de fichamento (resumo, rating, tags, conexões) que os demais 60+ artigos receberam. A base de conhecimento fica incompleta para o artigo mais citado metodologicamente.",
      "fix": "Criar vault/papers/demsar2006statistical.md usando o template de vault/templates/. Preencher com resumo do protocolo Friedman+Nemenyi, rating 5, tags [statistics, methodology, nonparametric] e conectar a artigos que também usam esse protocolo.",
      "why": "Demsar é a espinha dorsal metodológica da comparação. Sem nota no vault, não há rastreabilidade entre o protocolo estatístico da monografia e a literatura que o sustenta.",
      "ref": "monografia/cap_fundamentacao/fundamentacao.tex:95; monografia/bib/abntex2-references.bib:47-54"
    },
    {
      "id": "f-council-008",
      "title": "5 artigos descartados com entradas BibTeX ativas",
      "severity": "medium",
      "category": "metadata_contamination",
      "description": "hga2024hybrid, huang2025matrix, kappagantula2025dpso, sun2024hybrid, toaza2023review estão classificados como 'descartado-escopo-imediato' no vault mas mantêm entradas BibTeX válidas em abntex2-references.bib. Isso cria risco de citação acidental e polui o arquivo de referências. Além disso, infla artificialmente a contagem de 'artigos no referencial'.",
      "fix": "Remover as 5 entradas BibTeX de monografia/bib/abntex2-references.bib. Se alguma for mantida por cobrir área adjacente, reclassificar no vault e adicionar nota justificando.",
      "why": "Manter BibTeX para artigos explicitamente descartados contradiz a decisão editorial de escopo e cria inconsistência entre vault e monografia.",
      "ref": "monografia/bib/abntex2-references.bib:210,242,259,279,513"
    },
    {
      "id": "f-council-009",
      "title": "Metadados de uso (role, chapters, methods, claim_support) vazios em 100% dos artigos",
      "severity": "medium",
      "category": "vault_integrity",
      "description": "Nenhum dos 62 artigos no vault possui os campos role, chapters, methods ou claim_support preenchidos. A view 'Uso Na Monografia' do papers.base referencia exatamente esses campos e está, portanto, completamente inútil. Isso impede rastrear quais artigos sustentam quais claims, em quais capítulos, e com que papel (fundacional, comparativo, metodológico).",
      "fix": "Preencher pelo menos role e chapters para os ~25 artigos mais relevantes. role deve usar valores controlados como fundacional, comparativo, metodologico, contexto. chapters deve listar os números dos capítulos onde o artigo é citado (e.g., [2,4]). Atualizar a view para verificar completude.",
      "why": "Sem essa metadata, a rastreabilidade entre claims da monografia e evidências do referencial é manual e não auditável. Para uma defesa, é crucial demonstrar que cada afirmação tem respaldo bibliográfico localizável.",
      "ref": "vault/bases/papers.base:80-88"
    },
    {
      "id": "f-council-010",
      "title": "4 artigos com rating 0 citados na monografia — inconsistência de classificação",
      "severity": "medium",
      "category": "metadata_inconsistency",
      "description": "bean1994genetic (random keys), wang2021ant (parâmetros ACO), winter2002modeling (custos de curva), vanhove2012route (turn restrictions) têm rating 0 no vault mas são citados na monografia. Rating 0 deveria significar 'irrelevante' ou 'não avaliado'. Se são citados, deveriam ter rating >= 2. Essa inconsistência sugere que o sistema de rating não está sendo usado como filtro de qualidade.",
      "fix": "Revisar o rating dos 4 artigos: bean1994genetic >= 3 (base da codificação PSO), winter2002modeling >= 3 (contexto de custos angulares), vanhove2012route >= 3, wang2021ant >= 3. Atualizar as notas no vault.",
      "why": "Rating 0 para artigos citados mina a credibilidade do sistema de avaliação do vault como ferramenta de curadoria.",
      "ref": "vault/papers/bean1994genetic.md; vault/papers/wang2021ant.md; vault/papers/winter2002modeling.md; vault/papers/vanhove2012route.md"
    },
    {
      "id": "f-council-011",
      "title": "Poder estatístico do Friedman com N=30 instâncias não discutido",
      "severity": "low",
      "category": "statistical_validity",
      "description": "O teste de Friedman com 30 instâncias (blocos) e 3 métodos tem poder estatístico razoável, mas não é discutido. Demsar (2006) recomenda N >= 10 datasets e k >= 4 classifiers para poder adequado. Com k=3, o teste pode ter poder reduzido para detectar diferenças pequenas. A monografia reporta significância nos 3 pares, então o poder foi suficiente neste caso, mas a ausência de discussão é uma lacuna.",
      "fix": "Adicionar uma nota na Seção 4.2.5 sobre o poder do teste com N=30 e k=3, citando Demsar (2006, Seção 3.2) sobre a adequação do tamanho amostral. Notar que a significância observada em todos os pares mitiga a preocupação com poder insuficiente.",
      "why": "A discussão de poder estatístico é esperada em trabalhos que usam testes não-paramétricos sobre múltiplos datasets.",
      "ref": "monografia/cap_experimentos/experimentos.tex:152-162"
    },
    {
      "id": "f-council-012",
      "title": "Validade externa: apenas instâncias sintéticas sem discussão de generalização",
      "severity": "low",
      "category": "external_validity",
      "description": "As 30 instâncias são sintéticas (coordenadas em [-1,1]^2), sem correspondência com cenários reais de patrulha. A monografia reconhece isso como limitação, mas não cita literatura sobre generalização de benchmarks sintéticos para problemas reais ou sobre validação ecológica em otimização de rotas.",
      "fix": "Adicionar uma referência sobre validade ecológica em otimização combinatória (e.g., discussão em Hooker 1995 sobre 'competitive testing' vs 'scientific testing'). Incluir na Seção 4.2.6.",
      "why": "A validade externa é uma limitação reconhecida, mas sem citação de literatura sobre o tema, permanece como observação informal em vez de discussão metodológica.",
      "ref": "monografia/cap_experimentos/experimentos.tex:164-168"
    }
  ],
  "recommendation": "REPROVAR com exigência de correções antes da defesa. As falhas críticas (f-council-001, f-council-002) são de natureza catalográfica (faltam entradas BibTeX) e conceitual (a seção de lower bounds não referencia a literatura canônica). As falhas altas (f-council-003, f-council-004, f-council-005) representam riscos à validade interna e à robustez do desenho experimental. O plano de correção prioritário é: (1) adicionar as 9 entradas BibTeX faltantes, com destaque para heldkarp1970traveling; (2) revisar a Seção 2.8 para incluir Held-Karp e justificar a escolha da relaxação AP; (3) adicionar 2-3 referências sobre tuning e citar na discussão de parâmetros fixos; (4) criar nota vault para Demsar (2006). As demais falhas (médias e baixas) são importantes para a qualidade final mas não bloqueiam a defesa se as críticas forem endereçadas."
}
```

---

# Análise Metodológica do Referencial Teórico — TCC TSP-SD-ATP

**Juiz:** Gaps (Metodologista)  
**Data:** 2026-06-04  
**Veredito:** FAIL (confiança: HIGH)

## 1. Sumário Executivo

O referencial teórico possui **cobertura sólida** em TSP clássico (5 artigos), drone-TSP (7 artigos) e Algoritmos Genéticos (8 artigos). A seção de Otimização por Colônia de Formigas é a mais forte (13 artigos). O protocolo estatístico segue corretamente Demšar (2006).

No entanto, foram identificadas **4 falhas críticas/altas** que comprometem a sustentação metodológica da monografia e **8 falhas médias/baixas** que afetam a qualidade e a rastreabilidade do referencial.

## 2. Falhas Críticas

### 2.1 Held-Karp (1970) — Ausência da Referência Canônica de Lower Bounds (f-council-001)

O artigo de Held e Karp (1970), *The Traveling-Salesman Problem and Minimum Spanning Trees*, é o método canônico para obtenção de limitantes inferiores no TSP simétrico. A nota no vault (`vault/papers/heldkarp1970traveling.md`) demonstra que o autor **leu, compreendeu e atribuiu rating máximo (5)** a este artigo. A nota do vault inclusive discute como o Held-Karp bound poderia ser adaptado para o TSP-SD-ATP via redução 3D→2D.

Apesar disso:
- **Não existe entrada BibTeX** para `heldkarp1970traveling` em `abntex2-references.bib`.
- A **Seção 2.8 da monografia** descreve a relaxação AP/Hungarian sem **nenhuma menção** a Held-Karp, 1-trees, ou relaxação Lagrangiana.
- A palavra "subtour" aparece no texto descritivo, mas sem conexão com a formulação LP do TSP (subtour elimination constraints).

**Consequência:** A seção de lower bounds está metodologicamente descolada da literatura. O leitor não consegue avaliar se a relaxação AP foi uma simplificação deliberada (justificada pela tridimensionalidade do TSP-SD-ATP) ou uma omissão por desconhecimento do estado da arte.

### 2.2 Déficit de Citações na Seção de Lower Bounds (f-council-002)

Dos ~10 artigos mapeados no vault como relevantes para lower bounds, **9 não possuem entrada BibTeX** e **1 (valenzuela1997estimating) possui BibTeX mas não é citado**. Nenhum aparece na Seção 2.8. A seção inteira — que descreve um componente metodológico central do pipeline experimental — opera sem ancoragem bibliográfica.

Este é um padrão sistêmico: o vault contém informação (os artigos existem, estão fichados), mas a monografia não consegue acessá-la porque a etapa de criação de BibTeX não foi concluída.

## 3. Falhas Altas

### 3.1 Tuning de Hiperparâmetros (f-council-003)

O desenho experimental congela todos os parâmetros (pop=100, iter=100, e parâmetros específicos de cada método) para todas as 30 instâncias. A monografia reconhece a limitação, mas **não referencia literatura sobre metodologia de tuning**. Artigos como Birattari et al. (2010, F-Race), Hutter et al. (2011, SMAC) ou López-Ibáñez et al. (2016, irace) forneceriam a base para justificar parâmetros fixos como escolha metodológica (privilegiando reprodutibilidade) em vez de parecer uma omissão.

### 3.2 Validade Interna — Confundimento Método × Codificação (f-council-004)

GA usa permutação direta. PSO usa random keys. ACO usa feromônio tridimensional. As diferenças de desempenho observadas podem ser atribuídas tanto ao algoritmo quanto à codificação — e o desenho experimental **não isola esses efeitos**. A monografia menciona essa limitação na Seção 4.2.6, mas não a discute com respaldo bibliográfico. Halim e Ismail (2019), já citados no referencial, argumentam que a representação pode ser tão determinante quanto a meta-heurística — este artigo deveria ser usado para qualificar as conclusões.

### 3.3 Sensibilidade da Função Objetivo (f-council-005)

A Equação 4.1 define o custo como `c = distância + ângulo/π`. Esta é uma **contribuição original do autor**, mas **não há análise de sensibilidade** do peso relativo entre distância e ângulo. A normalização por π é uma escolha arbitrária de escala. Sem variar o peso da penalidade angular (e.g., λ ∈ {0, 0.5, 1.0, 2.0}), não é possível afirmar que os rankings entre métodos são robustos à especificação da função objetivo.

## 4. Falhas Médias

### 4.1 Desbalanceamento ACO × PSO (f-council-006)

ACO tem 13 referências (incluindo avançadas: neural ACO, RL ACO). PSO tem 5 referências, das quais 3 são artigos descartados. O referencial efetivo de PSO são apenas Kennedy (1995) e Clerc (2000). Este desbalanceamento pode enviesar a percepção do leitor.

### 4.2 Demšar sem Nota no Vault (f-council-007)

O artigo metodológico mais importante da monografia não passou pelo processo de fichamento.

### 4.3 Artigos Descartados com BibTeX (f-council-008)

5 artigos classificados como "descartado-escopo-imediato" mantêm entradas BibTeX. Inconsistência entre decisão editorial do vault e o arquivo de referências.

### 4.4 Metadados de Uso Vazios (f-council-009)

Nenhum artigo tem `role`, `chapters`, `methods` ou `claim_support` preenchidos. A view "Uso Na Monografia" do papers.base está inútil.

### 4.5 Rating 0 em Artigos Citados (f-council-010)

bean1994, wang2021, winter2002, vanhove2012 têm rating 0 mas são citados. O sistema de rating perde credibilidade como ferramenta de curadoria.

## 5. Falhas Baixas

### 5.1 Poder Estatístico (f-council-011)

O teste de Friedman com N=30 e k=3 tem poder razoável, mas a monografia não discute adequação amostral.

### 5.2 Validade Externa (f-council-012)

Apenas instâncias sintéticas, sem discussão respaldada sobre generalização.

## 6. Diagnóstico

O referencial teórico apresenta um **padrão de duas velocidades**: as seções "centrais" (TSP clássico, GA, ACO, drone-TSP) são bem cobertas e citadas, enquanto as seções "de suporte metodológico" (lower bounds, tuning, validade de constructo, sensibilidade) são frágeis ou ausentes.

A causa raiz parece ser um **gargalo catalográfico**: o vault contém artigos lidos e avaliados (heldkarp1970, johnson1996, kinable2017, etc.), mas as entradas BibTeX correspondentes não foram criadas. Isso bloqueia a citação e empobrece seções inteiras da monografia.

As falhas **NÃO são de incompetência conceitual** — a nota do vault sobre Held-Karp demonstra compreensão do método. São falhas de **execução do pipeline vault→BibTeX→monografia**.

## 7. Plano de Correção (por prioridade)

| Prioridade | Ação | Esforço |
|---|---|---|
| P0 | Criar BibTeX para `heldkarp1970traveling` e citar na Seção 2.8 | 15 min |
| P0 | Criar BibTeX para `johnson1996asymptotic` (validação empírica HK) | 10 min |
| P0 | Criar BibTeX para `kinable2017hybrid` (bounds para variantes TSP) | 10 min |
| P1 | Revisar Seção 2.8: justificar AP vs Held-Karp para TSP-SD-ATP | 30 min |
| P1 | Adicionar 2 referências de tuning (irace/F-Race/SMAC) + citar | 1 h |
| P1 | Criar nota vault para Demsar (2006) | 20 min |
| P2 | Criar BibTeX para os 6 lower-bound restantes | 1 h |
| P2 | Remover 5 entradas BibTeX de artigos descartados | 5 min |
| P2 | Atualizar rating 0→3 para 4 artigos citados | 10 min |
| P3 | Preencher `role` e `chapters` nos 25 artigos principais | 2 h |
| P3 | Análise de sensibilidade do peso angular | 4 h |

**Tempo total estimado para correções P0+P1 (mínimo para aprovação):** ~3 horas.  
**Tempo total para correções P0-P3 (qualidade de defesa):** ~10 horas.
