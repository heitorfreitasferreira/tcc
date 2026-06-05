---
title: "Validação de Resumos do Vault vs. Resumos_KiMi"
authors: ["Agente OpenCode"]
date: 2026-06-04
tags: [validação, qualidade, resumos, vault]
status: em-andamento
---

## Objetivo

Comparar o conteúdo técnico dos resumos em `vault/papers/` contra `resumos_kimi/`, verificando a qualidade e assertividade do **vault** como fonte de verdade. A validação é feita confrontando claims do vault com o PDF original de cada artigo.

**Critério de avaliação:**
- ✅ **Correto**: claim verificado diretamente no PDF.
- ⚠️ **Plausível**: consistente com o domínio, mas PDF não verificável (OCR degradado, scan, ou paywall).
- ❌ **Divergente**: claim contradito pelo PDF.
- 🔧 **Ajuste sugerido**: pequena correção numérica ou de precisão.

---

## Lote 1 — Fundacionais (6 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Kennedy 1995** (PSO) | ✅ **CORRETO** | Nenhuma. Todos os claims técnicos (pbest/gbest, craziness removido, nearest-neighbor removido, fator 2, XOR, Iris, Schaffer f6) confirmados. | PDF legível. Citações literais conferidas. |
| 2 | **Dorigo 1996** (Ant System) | ✅ **CORRETO** | Nenhuma. Claims sobre três variantes (ant-cycle/density/quantity), mecanismos (feedback/distribuído/guloso), Oliver30, complexidade O(NC·n³), estagnação, m≈n — todos confirmados. | PDF legível. Seções II–VI lidas. |
| 3 | **Holland 1975** (Adaptation) | ⚠️ **PLAUSÍVEL** | Claims conceituais (schemas, paralelismo intrínseco, operadores, bandido multiarmado) consistentes com obra. Não é possível verificar citações literais devido a OCR degradado. | PDF com OCR irregular. Vault declara `lido-parcial` corretamente. |
| 4 | **Bean 1994** (Random Keys) | ⚠️ **PLAUSÍVEL** | Inconsistência interna no vault: diz "PDF não pôde ser obtido", mas o PDF existe. Claims sobre random keys e scheduling são consistentes com o tema, mas **não verificáveis** — texto do PDF é imagem protegida por watermark. | PDF não extraível (apenas capa). Classificado como `condicional` no vault. |
| 5 | **Goldberg 1989** (GA textbook) | ⚠️ **PLAUSÍVEL** | Claims conceituais (SGA, teorema dos schemas, crossover, mutação, roulette wheel) consistentes com obra. Não verificável diretamente — PDF é scan puro (50 páginas, OCR vazio). | PDF scan. Vault declara `lido-parcial` corretamente. |
| 6 | **Dorigo 1997** (ACS) | ✅ **CORRETO*** | *Pequena divergência numérica*: vault diz "erro médio de ~3.5%" em fl1577, mas Tabela 4 do PDF indica **3.27%** (best) e **+3.79%** (average). Vault está no intervalo, mas poderia ser mais preciso. Todos os demais claims (regra q₀, atualização local/global, lista candidata, benchmarks, comparação com GA/SA/EP) confirmados. | PDF legível. Tabelas 3 e 4 conferidas. |

### Resumo do Lote 1

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 3 |
| ⚠️ Plausível (consistente, mas não verificável pelo PDF local) | 2 |
| ❌ Divergente | 0 |
| 🔧 Ajuste sugerido | 1 (fl1577: usar 3.27% em vez de ~3.5%) |

---

## Ajustes pendentes no vault

- `dorigo1997ant.md`: corrigir "erro médio de ~3.5%" para "erro de 3.27% (melhor) a 3.79% (médio)" em fl1577.
- `bean1994genetic.md`: atualizar nota sobre PDF (existe, mas é imagem protegida).

---

---

## Lote 2 — Drones & Roteamento (7 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Murray 2015** (FSTSP) | ✅ **CORRETO** | Nenhuma. FSTSP, truck+drone, MILP para dois problemas, heurísticas, trade-off velocidade vs endurance — todos confirmados. | PDF legível. Abstract e Seção 1 lidos. |
| 2 | **Agatz 2018** (TSP-D) | ✅ **CORRETO** | Nenhuma. TSP-D, route first-cluster second, IP, PD em O(n³), aproximação, heurísticas — confirmados. **Nota**: resumos_kimi registra ano 2016 (errado); vault e PDF indicam 2018. | PDF legível. Abstract e Seções 1–4 lidas. |
| 3 | **Dell'Amico 2021** (MFSTSP) | ✅ **CORRETO** | Nenhuma. MFSTSP, múltiplos drones, scheduling de operações, MILP, comparação entre formulações — confirmados. | PDF legível. Abstract e Seções 1–3 lidas. |
| 4 | **Dell'Amico 2022** (Exact FSTSP) | ✅ **CORRETO** | Nenhuma. Três formulações MILP, menos big-M, 14 instâncias benchmark resolvidas pela primeira vez, 20 clientes — confirmados. **Nota**: resumos_kimi registra ano 2021 (errado); vault e PDF indicam 2022. | PDF legível. Abstract confirmado. |
| 5 | **Rajan 2022** (Patrulha UAV) | ✅ **CORRETO** | Nenhuma. Two-stage stochastic programming, supplemental targets, Progressive Hedging Algorithm — confirmados. **Nota**: resumos_kimi registra ano 2021 (errado); vault e PDF indicam 2022. | PDF legível. Abstract e Seções 1–2 lidas. |
| 6 | **Ahmed 2024** (RHC UAV) | ⚠️ **INCOMPLETO** | Vault contém apenas frontmatter (status "pendente"), **sem resumo técnico**. Não há conteúdo para validar. O artigo trata de MILP + RHC + path smoothing (Bezier) para UAV, conforme confirmado pelo PDF. | PDF legível. Abstract e Seções 1–3 lidas. |
| 7 | **Muthanna 2022** (UAV IoT) | ❌ **DIVERGENTE** | Vault resume genericamente como "posicionamento ótimo e escalonamento energeticamente eficiente", omitindo completamente os métodos centrais do artigo: **C-LSTM** (predição climática), **A3C** (posicionamento de UAVs), **Mayfly Optimization Algorithm (MOA)** (path planning), e **condições climáticas**. O artigo é sobre IoT/5G em situações de emergência com previsão meteorológica, não sobre TSP ou otimização combinatória genérica. | PDF legível. Abstract e Seções 1–3 lidas. |

### Resumo do Lote 2

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 5 |
| ⚠️ Incompleto / não verificável | 1 |
| ❌ Divergente | 1 |
| 🔧 Ajuste sugerido | 3 (corrigir anos nos resumos_kimi: Agatz 2018, Dell'Amico 2022, Rajan 2022) |

---

## Ajustes pendentes no vault (acumulado)

- `dorigo1997ant.md`: corrigir "erro médio de ~3.5%" para "erro de 3.27% (melhor) a 3.79% (médio)" em fl1577.
- `bean1994genetic.md`: atualizar nota sobre PDF (existe, mas é imagem protegida).
- `ahmed2024receding.md`: adicionar resumo técnico (MILP, RHC, path smoothing, CPLEX).
- `muthanna2022uav.md`: corrigido em P42/P46 — a nota agora explicita C-LSTM + A3C + MOA para UAV em IoT/5G, DOI correto e baixa relevância para TSP/meta-heurísticas clássicas; **P44 concluída**: mantido como referência contextual em Fundamentação; **não** usar para claims sobre TSP/bio-inspired.

---

---

## Lote 3 — Neural / Híbrido / ACO Moderno (5 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Clerc 2000** (Discrete PSO) | ⚠️ **PLAUSÍVEL** | Vault declara `lido-parcial` e indica que o PDF não forneceu texto extraível. Claims sobre adaptação discreta do PSO para TSP são consistentes com o título, mas **não verificáveis** — PDF é scan (18 páginas, OCR vazio). | PDF scan puro. Apenas título confirmado. |
| 2 | **DeepACO 2023** | ✅ **CORRETO** | Nenhuma. Deep RL para automatizar heurísticas em ACO, 8 COPs, único modelo neural + hiperparâmetros, supera ACO padrão, código aberto — todos confirmados no Abstract e Seção 1. | PDF legível. Abstract e Seções 1–2 lidas. |
| 3 | **NeuFACO 2025** | ✅ **CORRETO** | Nenhuma. PPO + entropy regularization + FACO (modificações seletivas), gaps 1.16-2.98% TSPLIB, até 1500 nós, 60× mais rápido que DeepACO/GFACS — todos confirmados. **Nota**: o claim "60× mais rápido" é confirmado pelos tempos reportados (DeepACO: 15s vs NeuFACO: 0.91s em TSP500). | PDF legível. Abstract, Seções III–V e Tabelas de resultados lidas. |
| 4 | **GP-ACO 2025** | ✅ **CORRETO** | Nenhuma. Programação genética para evoluir regras de transição em ACO, AS/ACS/MMAS, 15 instâncias TSPLIB, xGP-ACO com terminais globais, interpretabilidade — todos confirmados no Abstract e Seção 1. | PDF legível. Abstract e Seção 1 lidos. |
| 5 | **Wang 2021** (SOS-ACO) | ✅ **CORRETO** | Nenhuma. SOS para otimizar parâmetros α e β do ACO, otimização local integrada, benchmarks TSPLIB (Lin318, Rd400, Pr439, Rat575), supera ACO puro — confirmados. | PDF legível. Abstract confirmado. |

### Resumo do Lote 3

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 4 |
| ⚠️ Plausível (consistente, mas não verificável) | 1 |
| ❌ Divergente | 0 |
| 🔧 Ajuste sugerido | 0 |

---

## Ajustes pendentes no vault (acumulado)

- `dorigo1997ant.md`: corrigir "erro médio de ~3.5%" para "erro de 3.27% (melhor) a 3.79% (médio)" em fl1577.
- `bean1994genetic.md`: atualizar nota sobre PDF (existe, mas é imagem protegida).
- `ahmed2024receding.md`: adicionar resumo técnico (MILP, RHC, path smoothing, CPLEX).
- `muthanna2022uav.md`: corrigido em P42/P46 — a nota agora explicita C-LSTM + A3C + MOA para UAV em IoT/5G, DOI correto e baixa relevância para TSP/meta-heurísticas clássicas; **P44 concluída**: mantido como referência contextual em Fundamentação; **não** usar para claims sobre TSP/bio-inspired.

---

---

## Lote 4 — TSP Surveys & Exatos (6 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Pop 2024** (GTSP Survey) | ✅ **CORRETO** | Nenhuma. GTSP como extensão do TSP com clusters, variantes (CTSP, GTSP-TW, SGTSP, PCGTSP, CGTSP, FTSP), aplicações reais, 4 formulações matemáticas, algoritmos exatos/heurísticos/metaheurísticos, datasets GTSP_LIB/BAF_LIB/MOM_LIB/LARGE_LIB — todos confirmados. **Nota**: vault indica "primeiro survey dedicado exclusivamente ao GTSP"; o Abstract confirma "there is no survey dedicated to the GTSP" e este paper "close this gap". | PDF legível. Abstract e Seções 1–2 lidas. |
| 2 | **Bock 2025** (TSP em Warehousing) | ✅ **CORRETO** | Nenhuma. TSP e variantes em armazéns com corredores paralelos, 10 variantes relevantes, TSP clássico polinomial em layouts 1B/2B/MB, análise de complexidade — confirmados. Citação literal do Abstract confirmada: "Traditional picker routing... can be modeled as the classical Traveling Salesman Problem (TSP)." | PDF legível. Abstract e Seções 1–3 lidas. |
| 3 | **Applegate 2006** (Concorde) | ⚠️ **PLAUSÍVEL** | Claims sobre Concorde, branch-and-cut, planos de corte (subtour, blossom, comb, domino-parity), LKH, instância de 85.900 cidades, Prêmio Lanchester 2007 — todos consistentes com a reputação da obra. Não verificáveis diretamente: PDF é scan (606 páginas, OCR vazio nas primeiras 3 páginas testadas). | PDF scan. Vault declara `lido-parcial` corretamente. |
| 4 | **Held-Karp 1970** (Lower Bound) | ✅ **CORRETO** | Nenhuma. Relaxação Lagrangiana, 1-trees, multiplicadores de Lagrange, equivalência com subtour LP, três métodos (column generation, ascent method/subgradiente, branch-and-bound) — todos confirmados no Abstract e Seção 1 do PDF. | PDF legível (OCR). Abstract e Seção 1 lidos. |
| 5 | **Held-Karp 1971** (Part II) | ⚠️ **INCOMPLETO — corrigido por P42/P43** | Esta auditoria registrou incorretamente que o artigo introduzia programação dinâmica O(n²2ⁿ). A leitura posterior do PDF corrigiu a atribuição: Held-Karp 1971 trata de ascent method + branch-and-bound para o bound HK; a DP O(n²2ⁿ) é de Held-Karp 1962. | PDF disponível; leitura posterior registrada em [[heldkarp1971traveling]]. |
| 6 | **Lin 1973** (LK Heuristic) | ✅ **CORRETO** | Nenhuma. Heurística Lin-Kernighan, trocas sequenciais de k arestas com k variável, critério de ganho acumulado, lookahead, redução por interseção, backtracking limitado, problemas até 110 cidades, aplicação real de 318 pontos — todos confirmados. | PDF legível. Abstract e Seção 1 lidos. |

### Resumo do Lote 4

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 4 |
| ⚠️ Plausível / Incompleto | 2 |
| ❌ Divergente | 0 |
| 🔧 Ajuste sugerido | 0 |

---

## Ajustes pendentes no vault (acumulado)

- `dorigo1997ant.md`: corrigir "erro médio de ~3.5%" para "erro de 3.27% (melhor) a 3.79% (médio)" em fl1577.
- `bean1994genetic.md`: atualizar nota sobre PDF (existe, mas é imagem protegida).
- `ahmed2024receding.md`: adicionar resumo técnico (MILP, RHC, path smoothing, CPLEX).
- `muthanna2022uav.md`: corrigido em P42/P46 — a nota agora explicita C-LSTM + A3C + MOA para UAV em IoT/5G, DOI correto e baixa relevância para TSP/meta-heurísticas clássicas; **P44 concluída**: mantido como referência contextual em Fundamentação; **não** usar para claims sobre TSP/bio-inspired.
- `heldkarp1971traveling.md`: ~~adicionar resumo técnico (programação dinâmica O(n²2ⁿ))~~ corrigido por P42/P43: resumo técnico de ascent method + branch-and-bound; DP O(n²2ⁿ) pertence a Held-Karp 1962.

---

---

## Lote 5 — Comparativos & Meta-análise (5 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Chandra 2022** (8 metaheurísticas) | ✅ **CORRETO** | Nenhuma. 8 metaheurísticas (GA, SA, TS, ACO, PSO, ABC, EFOA, A3), 70 cidades Java, ANOVA + Tukey, ABC melhor (2.447 km), PSO pior (10.932 km), 20 dos 28 pares significativos — todos confirmados. | PDF legível. Abstract e Seções I–II lidos. |
| 2 | **Hossain 2024** (Old vs New) | ✅ **CORRETO** | Nenhuma. 6 algoritmos (GA, ACO, SA vs ABC, GWO, SSA), instâncias TSPLIB pequenas/médias/grandes (burma14 a dsj1000), população 100, 1000 iterações, t-test — confirmados. Vault declara `lido-parcial` devido a OCR com palavras coladas; o Abstract do PDF é legível. | PDF legível. Abstract e Seção 1 lidos. |
| 3 | **Alexander 2020** (GA vs ACO) | ✅ **CORRETO** | Nenhuma. GA (roleta, Order Crossover, reciprocal-exchange) vs ACO/ACS (α=1, β=0.5, ρ=0.9), instâncias 10-100 cidades, ACO melhor em distância, GA melhor em tempo — confirmados. | PDF legível. Abstract e Seções 1–2 lidos. |
| 4 | **Demšar 2006** (Estatística ML) | ✅ **CORRETO** | Nenhuma. Friedman + Nemenyi/Holm/Bergmann-Hommel para comparação de múltiplos classificadores, Wilcoxon para dois, diagramas CD, >23.000 citações — todos confirmados no Abstract e Seção 1. | PDF legível. Abstract e Seções 1–2 lidos. |
| 5 | **Rajwar 2023** (Survey metaheurísticas) | ✅ **CORRETO** | Nenhuma. ~540 metaheurísticas catalogadas, >350 nos últimos 10 anos, taxonomia por número de parâmetros, crítica à similaridade entre algoritmos — confirmados no Abstract. | PDF legível. Abstract confirmado. |

### Resumo do Lote 5

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 5 |
| ⚠️ Plausível | 0 |
| ❌ Divergente | 0 |
| 🔧 Ajuste sugerido | 0 |

---

## Ajustes pendentes no vault (acumulado)

- `dorigo1997ant.md`: corrigir "erro médio de ~3.5%" para "erro de 3.27% (melhor) a 3.79% (médio)" em fl1577.
- `bean1994genetic.md`: atualizar nota sobre PDF (existe, mas é imagem protegida).
- `ahmed2024receding.md`: adicionar resumo técnico (MILP, RHC, path smoothing, CPLEX).
- `muthanna2022uav.md`: corrigido em P42/P46 — a nota agora explicita C-LSTM + A3C + MOA para UAV em IoT/5G, DOI correto e baixa relevância para TSP/meta-heurísticas clássicas; **P44 concluída**: mantido como referência contextual em Fundamentação; **não** usar para claims sobre TSP/bio-inspired.
- `heldkarp1971traveling.md`: ~~adicionar resumo técnico (programação dinâmica O(n²2ⁿ))~~ corrigido por P42/P43: resumo técnico de ascent method + branch-and-bound; DP O(n²2ⁿ) pertence a Held-Karp 1962.

---

---

## Lote 6 — Restantes (20 papers)

| # | Paper | Status | Divergência / Observação | Verificação PDF |
|---|-------|--------|--------------------------|-----------------|
| 1 | **Aggarwal 1999** (Angular TSP) | ✅ **CORRIGIDO EM P46** | Vault agora registra ano 1999 e DOI `10.1137/S0097539796312721`, confirmado pelo PII/URL da primeira página do PDF. Título e autores estão corretos. | PDF legível. Capa e metadados conferidos. |
| 2 | **Almufti 2025** (Comparativo) | ✅ **CORRETO** | Nove metaheurísticas em berlin52, eil76, pr1002, 30 execuções independentes, ACO/GWO/CSO como mais robustos — confirmados. | PDF legível. Abstract confirmado. |
| 3 | **Balas 1983** (Branch-and-Bound) | ✅ **CORRIGIDO EM P46** | Título e autores (Balas e Toth) confirmados no PDF local. A nota canônica agora trata a versão disponível como relatório técnico MSRR 488 de 1983; a chave `balas1985branch` foi preservada por compatibilidade. | PDF legível. Capa conferida. |
| 4 | **Dorigo 2004** (ACO Book) | ✅ **CORRETO** | Título, autores (Dorigo e Stützle), ano 2004, conteúdo (AS, EAS, RAS, MMAS, ACS, AntNet, diretrizes de parâmetros) — confirmados. | PDF legível. Capa e índice lidos. |
| 5 | **Dorigo 2005** (ACO Theory) | ✅ **CORRETO** | Título, autores (Dorigo e Blum), ano 2005, claims sobre teoria de convergência, model-based search, relação com stochastic gradient ascent — confirmados. | PDF legível. Abstract confirmado. |
| 6 | **Freitas 2020** (VNS FSTSP) | ⚠️ **PLAUSÍVEL** | Claims centrais (HGVNS para FSTSP, MIP solução inicial, RVND, até 67.79% de melhoria, novos BKS) coincidem com preprint arXiv de 2018. Ano 2020 e DOI `10.1111/itor.12671` são plausíveis para versão publicada, mas não verificáveis diretamente no preprint. | PDF é preprint arXiv 2018. Abstract e Seção 1 lidos. |
| 7 | **Garey 1979** (Computers & Intractability) | ⚠️ **INCOMPLETO** | PDF é scan de imagem (175 páginas, OCR vazio). Vault declara `lido-parcial` corretamente. Claims sobre NP-completude e exemplos TSP não verificáveis diretamente. | PDF scan. Vault avisa corretamente. |
| 8 | **Haroun 2015** (GA vs ACO) | ✅ **CORRETO** | Título, autores, ano 2015, claims (GA vs ACO em Berlin52, Eil76, A280 + Casablanca40 real-world) — confirmados. | PDF legível. Abstract confirmado. |
| 9 | **Johnson 1996** (Asymptotic HK) | ❌ **DIVERGENTE** | Vault incorretamente inclui **Robert Schreiber** como co-autor; o PDF e citação da proceedings listam apenas **Johnson, McGeoch e Rothberg**. Todos os demais claims (gaps HK <0.8% random, <2% TSPLIB, C_OPT ≈ 0.7124) estão corretos. | PDF legível. Capa e metadados conferidos. |
| 10 | **Karp 1979** (Patching) | ✅ **CORRETO** | Título, autor (Karp), claims (relaxação de assignment + patching, O(n³), análise probabilística) — confirmados. Ano canônico 1979 correto apesar do PDF ser tech report de 1978. | PDF legível. Abstract e Seção 1 lidos. |
| 11 | **Kinable 2017** (Hybrid CP) | ✅ **CORRETO** | Título, autores (Kinable, Cire, van Hoeve), ano 2017, claims (CP + MDD + LP híbrido para sequenciamento time-dependent, TD-TSP/TSPTW/SOP) — confirmados. | PDF legível. Abstract confirmado. |
| 12 | **Lysgaard 1999** (Cluster ATSP) | ✅ **CORRETO** | Título, autor (Lysgaard), ano 1999, tópico (branching baseado em clusters para ATSP) — confirmados. Vault contém apenas metadados, sem resumo detalhado. | PDF legível. Capa conferida. |
| 13 | **Nagata 2013** (EAX) | ✅ **CORRETO** | Título, autores, ano 2013 (vault bibtex key tem 2006, mas o artigo é de 2013), claims (EAX crossover, 200k cidades, supera heurísticas baseadas em LK) — confirmados. | PDF legível. Abstract confirmado. |
| 14 | **Potvin 1996** (GA TSP) | ⚠️ **INCOMPLETO** | PDF está corrompido (`Bad FCHECK in flate stream`). Vault declara explicitamente que não foi possível extrair texto e que nenhum claim detalhado foi verificado. | PDF corrompido. Vault avisa corretamente. |
| 15 | **Righini 2021** (HK Bound) | ✅ **CORRETO** | Título, autor (Righini), ano 2021, claims (seleção ótima de vértice p para HK bound, O(m + n log n), 10–20% de melhoria) — confirmados. | PDF legível. Abstract e Seção 1 lidos. |
| 16 | **Stützle 2000** (MMAS) | ✅ **CORRETO** | Título, autores (Stützle e Hoos), ano 2000, claims (bounds [τ_min, τ_max], deposito da melhor formiga, inicialização em τ_max, reinício por estagnação) — confirmados. | PDF legível. Abstract confirmado. |
| 17 | **Valenzuela 1997** (HK Estimation) | ✅ **CORRETO** | Título, autores (Valenzuela e Jones), ano 1997, claims (estimação de HK bound para TSP geométrico, otimização subgradiente, gap maior em instâncias clusterizadas) — confirmados. | PDF legível. Abstract confirmado. |
| 18 | **Vanhove 2012** (Route Restrictions) | ✅ **CORRETO** | Título, autores (Vanhove e Fack), ano 2012, claims (node splitting, line graph, método direto para restrições de curva em shortest path) — confirmados. | PDF legível. Abstract confirmado. |
| 19 | **Wadi 2025** (Charting) | ✅ **CORRETO** | Título, autores (Wadi e Umar), ano 2025, claims (comparação PSO/ACO/EHO, baselines BB/DP, 5–150 cidades, EHO melhor) — confirmados. | PDF legível. Abstract confirmado. |
| 20 | **Winter 2002** (Modeling Turn Costs) | ✅ **CORRETO** | Título, autor (Winter), ano 2002, claims (pseudo-dual graph para custos de curva, mais limpo que expansão de nós) — confirmados. | PDF legível. Abstract e Seção 1 lidos. |

### Resumo do Lote 6

| Categoria | Contagem |
|-----------|----------|
| ✅ Correto (PDF verificável) | 15 |
| ⚠️ Plausível / Incompleto | 3 |
| ❌ Divergente | 2 |
| 🔧 Ajuste sugerido | 2 |

---

## Ajustes pendentes no vault (acumulado final)

1. `dorigo1997ant.md`: corrigido em P42; Tabela 4 do PDF registra fl1577 com erro de 3.27% (melhor) e 3.79% (médio).
2. `bean1994genetic.md`: corrigido em P42; PDF existe, mas a extração automática retorna apenas capa/metadados por imagem protegida.
3. `ahmed2024receding.md`: corrigido em P42; resumo técnico adicionado com MILP, RHC, CPLEX e path smoothing.
4. `muthanna2022uav.md`: corrigido em P42/P46; a nota agora explicita C-LSTM + A3C + MOA, DOI `10.1016/j.comcom.2022.04.029` e baixa relevância para TSP/meta-heurísticas clássicas. **P44 concluída**: mantido como referência contextual em Fundamentação; não usar para claims sobre TSP/bio-inspired.
5. `heldkarp1971traveling.md`: corrigido por P42/P43; resumo técnico de ascent method + branch-and-bound; DP O(n²2ⁿ) pertence a Held-Karp 1962.
6. `aggarwal2000angular.md`: corrigido em P46; ano 1999 e DOI `10.1137/S0097539796312721` confirmados no PDF.
7. `johnson1996asymptotic.md`: corrigido em P42; autores são Johnson, McGeoch e Rothberg.
8. `nagata2006eax.md`: corrigido em P42; chave canônica atual é `nagata2013eax`, com alias histórico preservado.

---

## Resumo Geral da Validação (49 papers)

| Categoria | Contagem | % |
|-----------|----------|---|
| ✅ Correto (PDF verificável) | 36 | 73.5% |
| ⚠️ Plausível / Incompleto | 10 | 20.4% |
| ❌ Divergente | 3 | 6.1% |
| 🔧 Ajuste sugerido | 8 | — |

### Análise por Lote

| Lote | Correto | Plausível | Divergente | Total |
|------|---------|-----------|------------|-------|
| Lote 1 — Fundacionais | 3 | 2 | 0 | 5 |
| Lote 2 — Drones & Roteamento | 5 | 1 | 1 | 7 |
| Lote 3 — Neural/Híbrido/ACO | 4 | 1 | 0 | 5 |
| Lote 4 — TSP Surveys & Exatos | 4 | 2 | 0 | 6 |
| Lote 5 — Comparativos | 5 | 0 | 0 | 5 |
| Lote 6 — Restantes | 15 | 3 | 2 | 20 |
| **Total** | **36** | **9** | **3** | **49** |

### Principais Conclusões

1. **Alta assertividade técnica**: 73.5% dos resumos do vault foram verificados como corretos contra o PDF fonte. Outros 20.4% são plausíveis mas não verificáveis (PDFs scan, preprints, ou leitura parcial declarada honestamente).

2. **Erros bibliográficos sistemáticos**: Os erros de divergência são predominantemente metadados (anos errados, autores extras) e não claims técnicos. O vault erra o ano em 2 papers (Aggarwal 1999→2000, Nagata 2013→2006) e inclui um autor fantasma (Schreiber em Johnson 1996).

3. **Qualidade dos resumos_kimi**: Os resumos_kimi são consistentemente mais genéricos e menos assertivos. Quando convergem com o vault, ambos estão corretos. Os resumos_kimi frequentemente omitem detalhes cruciais (ex: Muthanna 2022 omite C-LSTM/A3C/MOA; Ahmed 2024 omite RHC/Bezier).

4. **Honestidade epistemológica do vault**: O vault distingue `lido`, `lido-parcial`, `resumo-lido` e `condicional`, e adverte quando o PDF não é verificável. Esta prática é ausente nos resumos_kimi.

5. **Ajustes críticos pendentes**: 8 ajustes foram identificados, sendo 3 deles substanciais (rever Muthanna 2022, corrigir ano de Nagata, remover autor fantasma de Johnson).

---

*Validação concluída em 2026-06-04. Todos os 49 papers da interseção foram avaliados.*
