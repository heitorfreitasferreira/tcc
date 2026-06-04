# Council Report — Validação do Referencial Teórico do TCC

**Data:** 2026-06-04  
**Tipo:** validate  
**Alvo:** Referencial teórico (61 artigos, vault + BibTeX + monografia)  
**Modo:** 3 juízes com perspectivas acadêmicas, 2 rodadas (simulando debate)  
**Fidelidade:** simulada (OpenCode — juízes re-spawned com R1 findings na R2)

---

## Veredito Final: FAIL (consenso unânime)

| Juiz | Perspectiva | R1 | R2 | Mudança |
|------|-------------|-----|-----|---------|
| **Gaps** | Metodologista | FAIL (HIGH) | FAIL (HIGH) | Confirmado, agravado |
| **Depth** | Especialista TSP | FAIL (HIGH) | FAIL (HIGH) | Confirmado, +3 achados |
| **Coverage** | Revisor de Literatura | FAIL (HIGH) | FAIL (HIGH) | Confirmado, +2 achados |

**Consenso:** 3/3 FAIL, confiança HIGH. Não houve divergência entre juízes. O debate R2 reforçou e expandiu os achados da R1; nenhum juiz revisou seu veredito. Trata-se de um **FAIL forte**: todas as perspectivas convergem independentemente para a mesma conclusão.

---

## Achados Consolidados

### CRÍTICOS (bloqueadores — 3)

#### C1. Held-Karp (1970/1971) ausente do BibTeX — lacuna canônica em lower bounds
**Juízes:** Gaps (f-001), Depth (f-001), Coverage (f-lit-001) — 3/3 concordam como CRITICAL.
- heldkarp1970 está no vault (rating 5, lido, PDF íntegro) mas sem entrada BibTeX
- heldkarp1971 está no vault (rating 5, pendente) mas sem entrada BibTeX
- Seção 2.8 da monografia descreve relaxação AP/Hungarian sem citar Held-Karp, 1-trees, ou relaxação Lagrangiana
- Gap empírico do HK bound para TSP simétrico é < 0.8% — qualquer banca de otimização identifica esta omissão em segundos
- **Fix:** Criar BibTeX para heldkarp1970/1971, citar na Seção 2.8, justificar escolha AP vs HK

#### C2. Seção 2.8 (Lower Bounds) sem citações — 10 artigos fichados, 0 referenciados
**Juízes:** Gaps (f-002), Depth (f-001, f-010), Coverage (f-lit-001) — 3/3 concordam.
- Dos 10 artigos de lower bounds no vault, 0 têm BibTeX completo e 0 são citados no texto
- Padrão sistêmico de pipeline vault→BibTeX quebrado para toda a área
- A seção menciona "subtours" sem referenciar formulação LP do TSP
- **Fix:** Criar BibTeX para 9 artigos, citar heldkarp1970 + johnson1996 + valenzuela1997 na Seção 2.8

#### C3. Implementação ACO é Ant System, NÃO MMAS — erro factual
**Juiz:** Depth (f-002) — identificado na R1, não contestado na R2.
- Código (`src/optimization/aco/ant.go:14-39`): todas as formigas depositam feromônio (estilo Ant System)
- Sem bounds [τ_min, τ_max], sem elitismo de depósito, sem reinicialização por estagnação
- Nota do vault (`stutzle2000mmas.md:55`) afirma incorretamente que "τ_max = 1/(ρ·L_best) e τ_min = τ_max/(2n) é implementada diretamente no código"
- **Fix:** Ou implementar MMAS no código, ou corrigir vault + monografia para documentar Ant System 3D

### ALTOS (devem ser resolvidos antes da versão final — 6)

#### H1. ATSP e TDTSP completamente ausentes do referencial
**Juiz:** Depth (f-003). Reforçado por Coverage na R2.
- TSP-SD-ATP gera matriz assimétrica após redução min_i — é fundamentalmente um ATSP
- Dependência de sequência nos custos caracteriza TDTSP
- Zero referências a ATSP ou TDTSP no vault ou BibTeX
- Kinable (2017) cobre TDTSP com decision diagrams e está no vault (rating 5, lido) — mas sem BibTeX

#### H2. Feromônio 3D como inovação não documentada
**Juiz:** Depth (f-004). Reforçado por Gaps (f-004) sobre confundimento método×codificação.
- τ(i,j,k) tridimensional é adaptação inédita para TSP — literatura canônica usa τ(i,j) bidimensional
- Não documentado como contribuição algorítmica original
- Risco: parecer erro de implementação em vez de adaptação deliberada

#### H3. Bean (1994) — base do PSO implementado — rating 0 e PDF ilegível
**Juiz:** Depth (f-005). Reforçado por Coverage (f-lit-003) como padrão de rating 0 em citados.
- Random keys é o mecanismo exato usado pelo PSO para decodificar vetores reais em permutações
- PDF contém apenas metadados INFORMS — corpo do artigo não foi lido
- A fundamentação do PSO repousa sobre artigo cujo conteúdo é desconhecido

#### H4. Ausência de literatura sobre tuning/calibração de hiperparâmetros
**Juiz:** Gaps (f-003). Adotado por Depth na R2 (f-R2-001).
- Parâmetros fixos (pop=100, iter=100, α=1.0, β=2.0, ρ=0.2, w=0.7, c1=c2=2.0) para todas as instâncias
- Sem referências a irace, SMAC, F-Race, ou qualquer metodologia de tuning
- wang2021ant (tuning de parâmetros ACO) está no BibTeX com rating 0 — não aproveitado

#### H5. Artigos com rating 0 usados como autoridade na fundamentação
**Juiz:** Coverage (f-lit-003). Reforçado por Depth na R2.
- winter2002 (rating 0, Seção 2.2), vanhove2012 (rating 0, Seção 2.2), bean1994 (rating 0, Seção 2.6)
- Padrão sistêmico: artigos considerados irrelevantes ou não avaliados são citados como suporte

#### H6. Demšar (2006) — base metodológica dos testes estatísticos — sem nota no vault
**Juízes:** Gaps (f-007), Depth (f-012), Coverage (f-lit-004) — 3/3 concordam.
- Único artigo do BibTeX sem nota correspondente no vault
- Protocolo Friedman + Nemenyi é a espinha dorsal das comparações experimentais
- Sem resumo, sem avaliação de adequação ao contexto TSP (artigo original foca classificadores ML)

### MÉDIOS (esperados para a defesa — 8)

| ID | Achado | Juiz principal |
|----|--------|---------------|
| M1 | Viés ACO (11) vs PSO (6) no vault — shami2022pso e gad2022pso no BibTeX mas não citados | Coverage |
| M2 | 4 artigos ACO+DL (DeepACO, NeuFACO, PPACO, GPACO) irrelevantes para escopo | Coverage + Depth |
| M3 | Confusão conceitual AP vs Held-Karp na nota do vault | Depth |
| M4 | Ausência de heurísticas construtivas e Simulated Annealing como baselines | Depth + Coverage |
| M5 | Nagata (2006) EAX com vault vazio (12 linhas, sem resumo) | Depth |
| M6 | 13 artigos no vault sem entrada BibTeX (assimetria vault↔BibTeX) | Coverage |
| M7 | 18 artigos pendentes (29.5% do vault sem resumo) — 5 já citados no texto | Coverage |
| M8 | Agravante da confusão AP vs HK: a nota do vault afirma equivalência que não existe | Depth |

### BAIXOS (desejáveis — 7)

| ID | Achado | Juiz principal |
|----|--------|---------------|
| L1 | Rastreabilidade zero (role/chapters/claim_support vazios em 100% dos artigos) | Gaps + Coverage |
| L2 | 5 artigos descartados mantendo entradas BibTeX — divergência na R2 sobre remover ou anotar | Gaps (desafiado por Depth) |
| L3 | Poder estatístico Friedmann com N=30, k=3 não discutido | Gaps (desafiado por Coverage) |
| L4 | Validade externa: apenas instâncias sintéticas sem discussão de generalização | Gaps |
| L5 | Seção 2.4 (Busca Exaustiva) sem citações | Coverage |
| L6 | 5 artigos em veículos questionáveis concentrados em estudos comparativos | Coverage |
| L7 | --gama listado na documentação mas inexistente no código | Depth |
| L8 | Clerc (2000) PSO discreto citado como referência mas método usado é random keys | Depth |

---

## Debate R2: Divergências e Convergências

### Divergências identificadas (2)

1. **Artigos descartados no BibTeX (f-council-008 Metod):** Gaps recomenda REMOVER. Depth discorda — argumenta que o `.bib` é base de dados do projeto, não reference list final, e manter entradas documenta curadoria. Propõe anotar com `keywords = {discarded}` em vez de remover. **Resolução:** Manter com anotação; adicionar política no AGENTS.md.

2. **Poder estatístico Friedmann (f-council-011 Metod):** Gaps classifica como LOW. Coverage discorda da severidade implícita, argumentando que N=30 compensa k=3 e Demšar não estabelece k≥4 como regra rígida. Reclassifica como LOW (concordante com Gaps na severidade final). **Resolução:** Manter LOW; adicionar nota na Seção 4.2.5.

### Convergências reforçadas (5)

1. **Pipeline vault→BibTeX quebrado**: Todos os 3 juízes identificaram independentemente que artigos lidos e avaliados (Held-Karp rating 5, Johnson rating 5, Kinable rating 5) não chegam ao BibTeX. Causa-raiz do C1, C2, M6, M7.

2. **Seção 2.8 como ponto de falha concentrado**: 3/3 juízes apontam esta seção como a mais frágil. Zero citações + 10 artigos disponíveis = falha de execução, não de conhecimento.

3. **Rating 0 como indicador de qualidade de revisão**: 3/3 juízes identificaram o padrão de artigos com rating 0 sendo citados. Coverage elevou de incidente isolado (Bean) para padrão sistêmico (Winter, Vanhove).

4. **Desbalanceamento ACO/PSO**: Coverage documentou viés quantitativo; Gaps mostrou implicação metodológica (confundimento); Depth apontou que shami2022pso e gad2022pso estão no BibTeX mas não citados — corrigível.

5. **Demšar como lacuna de conhecimento**: 3/3 juízes identificaram que o artigo metodológico central não foi fichado. Convergência completa.

---

## Diagnóstico de Causa-Raiz

O referencial **não sofre de falta de conhecimento** — o vault demonstra que os artigos canônicos foram identificados, lidos e avaliados. O problema é de **execução do pipeline de três etapas**:

```
vault (fichamento) → BibTeX (entrada) → monografia (citação)
     ✅                    ❌                   ❌
```

As seções "centrais" (TSP clássico, GA, ACO, drone-TSP) completaram as 3 etapas. As seções "de suporte" (lower bounds, angular, tuning, validade) pararam na etapa 1. Isso cria um **padrão de duas velocidades** no referencial.

Adicionalmente, foi identificado um **erro factual** (ACO é Ant System, não MMAS) que contamina a nota do vault e potencialmente o texto da monografia. Este erro é independente do pipeline — é de validação cruzada entre código e documentação.

---

## Plano de Correção (ordenado por prioridade)

### Bloqueadores (antes da qualificação/defesa)

| # | Ação | Esforço | Achados |
|---|------|---------|---------|
| 1 | Criar BibTeX para heldkarp1970traveling e heldkarp1971traveling | 30 min | C1 |
| 2 | Criar BibTeX para johnson1996asymptotic, kinable2017hybrid, righini2021efficient, aggarwal2000angular | 40 min | C2, H1 |
| 3 | Revisar Seção 2.8: citar Held-Karp, justificar AP vs HK, adicionar valenzuela1997 | 45 min | C1, C2 |
| 4 | Corrigir vault stutzle2000mmas.md: remover afirmação falsa sobre implementação MMAS | 10 min | C3 |
| 5 | Verificar e documentar no texto se a implementação é AS ou se MMAS será implementado | 30 min | C3 |
| 6 | Adicionar ao menos 1 referência de ATSP e 1 de TDTSP ao BibTeX e citar na Seção 2.1 ou 2.2 | 1 h | H1 |

### Alta prioridade (antes da versão final)

| # | Ação | Esforço | Achados |
|---|------|---------|---------|
| 7 | Documentar feromônio 3D como contribuição algorítmica original na Seção 2.7 | 30 min | H2 |
| 8 | Obter PDF legível de Bean (1994) via INFORMS/Sci-Hub; atualizar nota e rating | 1 h | H3 |
| 9 | Adicionar 2 referências de tuning (irace/SMAC/F-Race) + citar wang2021ant | 1 h | H4 |
| 10 | Revisar ratings: bean1994≥3, winter2002≥3, vanhove2012≥3, wang2021≥3 | 15 min | H5 |
| 11 | Criar nota vault para demsar2006statistical | 30 min | H6 |

### Média prioridade (para a defesa)

| # | Ação | Esforço | Achados |
|---|------|---------|---------|
| 12 | Citar shami2022pso e gad2022pso na Seção 2.6 | 15 min | M1 |
| 13 | Isolar 4 artigos ACO+DL em categoria separada ou remover do escopo principal | 30 min | M2 |
| 14 | Corrigir nota heldkarp1970traveling.md: remover afirmação de equivalência AP=HK | 10 min | M3 |
| 15 | Adicionar SA como baseline referenciado (Kirkpatrick 1983 ou Johnson & McGeoch 1997) | 1 h | M4 |
| 16 | Completar leitura e nota de Nagata (2006) EAX | 1 h | M5 |
| 17 | Criar BibTeX para os 6 artigos restantes de lower bounds | 1 h | M6 |

### Baixa prioridade (desejável)

| # | Ação | Esforço | Achados |
|---|------|---------|---------|
| 18 | Preencher `chapters` e `role` nos 25+ artigos citados | 2 h | L1 |
| 19 | Anotar artigos descartados no BibTeX com `keywords={discarded}` | 10 min | L2 |
| 20 | Nota sobre poder estatístico na Seção 4.2.5 | 15 min | L3 |
| 21 | Remover `--gama` da documentação (AGENTS-experiments.md) | 5 min | L7 |

**Tempo total estimado:** ~15 horas para correções completas (bloqueadores + altas + médias).  
**Mínimo para defesa:** ~5 horas (bloqueadores + altas).

---

## Conclusão do Council

O referencial teórico **não está pronto para defesa** no estado atual. As falhas são de execução, não de concepção — o vault demonstra conhecimento adequado das áreas, mas o pipeline de conversão vault→BibTeX→citação falhou para as seções de suporte metodológico (lower bounds, tuning, validade). Adicionalmente, um erro factual (ACO ≠ MMAS) precisa ser corrigido para evitar contestação durante a arguição.

Com ~5 horas de trabalho focado nos bloqueadores e nas correções de alta prioridade, o referencial atinge qualidade suficiente para uma defesa de mestrado. As correções médias e baixas (outras ~10 horas) elevariam o trabalho ao nível de excelência esperado para publicação.

**Veredito final do Council: FAIL — correções obrigatórias antes da defesa, com plano de ação bem delimitado e factível.**
