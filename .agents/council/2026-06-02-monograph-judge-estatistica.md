```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "findings": [
    "F1: 51 seeds per method is commendable — well above typical practice (10–30), but no power analysis was conducted to justify the choice or detect minimum effect size of interest",
    "F2: Friedman test with Iman-Davenport correction is correctly selected for 3 related groups across 30 blocks (instances); all assumptions satisfied (randomized block, ordinal data, ≥5 blocks, block independence)",
    "F3: Effect sizes (e.g. Cliff's delta, Kendall's W, or rank-biserial correlation) are absent from the report; p-values alone conflate statistical and practical significance — especially problematic with 51×30=1530 observations per method, where even trivial differences become significant",
    "F4: Different encoding representations (GA: permutation, PSO: random keys, ACO: 3D pheromone) introduce a construct-validity confound: observed performance differences conflate algorithm quality with encoding expressiveness — this is not discussed as a threat to validity",
    "F5: Lower bound gap averaging 51.36% is extremely loose (range 40.79–65.35%); it provides no practical discrimination among methods and should be de-emphasized as a quality metric",
    "F6: Only 18/30 instances have brute-force optima, all small (10–15 pts); gap analysis extrapolated to 100-pt instances is a conclusion-validity threat — small-instance gap performance does not generalize",
    "F7: No hyperparameter sensitivity analysis was conducted; conclusions about relative performance are conditional on the single chosen parameter configuration per method — external-validity threat",
    "F8: Bonferroni-Holm correction for pairwise Wilcoxon tests is correctly applied (FWER control with sequential rejection); test sequence α/3, α/2, α/1 is appropriate for 3 comparisons",
    "F9: Convergence analysis is not detailed in the provided packet — missing analysis of convergence trajectories, diversity metrics, or stagnation detection; as presented, it is inadequate",
    "F10: Timing comparison (ACO ~4262ms vs GA ~38ms vs PSO ~75ms at 100 pts) is reported without hardware specifications or variance measures — a 10× difference within a method across runs would change the conclusion; report mean±sd over seeds",
    "F11: Threats to validity (internal, external, construct, conclusion) are either absent or insufficiently discussed based on available information; at minimum, encoding confound (F4), hyperparameter dependency (F7), and small-instance generalization (F6) must be explicitly documented",
    "F12: Hit rate (ACO 71.79%, GA 26.47%, PSO 4.25%) is a useful complementary metric to gap, but its definition (exact optimum vs ε-neighborhood?) should be clearly stated",
    "F13: 30 synthetic instances in [-1,1] lack real-world topology diversity; external validity to actual patrol scenarios is limited",
    "F14: No confidence intervals reported for any point estimate (mean gap, mean time, hit rate), making sampling uncertainty invisible"
  ],
  "recommendation": "Major revisions required before defense. Address: (1) add effect sizes alongside all p-values; (2) discuss the encoding confound as a construct-validity threat; (3) de-emphasize the 51% LB gap; (4) add hyperparameter sensitivity analysis or justify choices against literature; (5) report timing with variance; (6) strengthen convergence analysis with diversity metrics and trajectory plots; (7) explicitly structure threats to validity per Shadish-Cook-Campbell framework; (8) report confidence intervals for all main estimates."
}
```

# Análise Metodológica e Estatística — Estatística

## 1. Delineamento Experimental

### 1.1 Adequação ao Problema

O delineamento emprega 3 meta-heurísticas × 30 instâncias × 51 sementes = 4.590 execuções. Trata-se de um **delineamento em blocos completos casualizados (RCBD)**, onde as instâncias são os blocos, os métodos são os tratamentos, e as sementes constituem as réplicas dentro de cada bloco. Esta estrutura é **apropriada** para o problema de comparação de algoritmos estocásticos de otimização combinatória.

Pontos fortes:
- O uso de 51 sementes é **generoso** — a literatura típica emprega 10–30 execuções (Eftimov et al., 2017; García et al., 2009).
- A cobertura de 30 instâncias (10 tamanhos × 3 variantes) permite avaliar escalabilidade.
- A inclusão de *LowerBound* via relaxação de atribuição e *BruteForce* para instâncias pequenas ancora a análise em referências objetivas.

### 1.2 Ausência de Análise de Poder

**Problema grave**: Não há menção a uma análise de poder (*power analysis*) para justificar o número de sementes. Com 51 execuções por condição, o Friedman test detectará diferenças estatisticamente significativas mesmo quando **trivialmente pequenas** do ponto de vista prático. Sem *effect size* (ver Seção 5), o leitor não consegue distinguir significância estatística de relevância prática.

*Recomendação*: Realizar *power analysis* a priori (ou *sensitivity power analysis* a posteriori) para determinar o menor *effect size* detectável com (1−β)=0.80 e α=0.05, reportando o *minimum detectable effect size* em unidades padronizadas (eg. Kendall's W para o Friedman, ou standardized rank difference).

## 2. Aplicação do Teste de Friedman

### 2.1 Adequação e Suposições

O **teste de Friedman** com correção de Iman-Davenport (estatística F aproximada) é a **escolha correta** para:
- k = 3 tratamentos relacionados (medidas repetidas nos mesmos blocos)
- n = 30 blocos (instâncias) — muito acima do mínimo de 5 exigido pela aproximação qui-quadrado
- Dados de makespan em escala de razão (ordinal ≥, que é o requisito mínimo)
- Distribuição quase certamente não-normal (típico de custos de rota)

Suposições verificadas:
| Suposição | Status | Comentário |
|-----------|--------|------------|
| Blocos aleatorizados | ✓ | Instâncias independentes, coordenadas sintéticas |
| Dados ordinais ou superiores | ✓ | Makespan é razão |
| k ≥ 3 grupos relacionados | ✓ | k=3 |
| n ≥ 5 blocos | ✓ | n=30, a aproximação F de Iman-Davenport é robusta |
| Independência entre blocos | ✓ | Instâncias geradas independentemente |

O resultado (Friedman F ≈ 293.22, p ≈ 4.71×10⁻³¹) indica rejeição forte da hipótese nula de que todos os métodos têm desempenho equivalente. Não há controvérsia aqui.

### 2.2 Post-hoc: Nemenyi e Wilcoxon-Holm

O **teste de Nemenyi** (para comparações múltiplas após Friedman) com CD = 0.6050 a α = 0.05 é corretamente aplicado. As diferenças entre as médias das ordens (ACO 1.100, GA 1.900, PSO 3.000) excedem o CD em todos os pares, confirmando três grupos distintos.

O **Wilcoxon signed-rank pareado com correção Bonferroni-Holm** é uma **escolha metodologicamente superior** ao Nemenyi cru, por ser:
1. **Menos conservador** que Bonferroni clássico: Holm sequencia os p-valores e testa em α/3, α/2, α/1, rejeitando enquanto pᵢ ≤ α/(k−i+1)
2. **Controla FWER** (Family-Wise Error Rate) rigorosamente
3. Mais poderoso que Nemenyi para detectar diferenças quando existe ordenação clara

A correção está **correta** para as 3 comparações pareadas. Se houver mais comparações implícitas (ex: ACO vs GA, ACO vs PSO, GA vs PSO, mais ACO vs BruteForce, GA vs BruteForce, PSO vs BruteForce), o número de comparações sobe para 6 e a correção precisa ser recalculada. Vale verificar se BruteForce e LowerBound foram incluídos nos testes post-hoc.

## 3. Confundimento de Representação (Ameaça à Validade de Construto)

**Achado crítico**: Cada meta-heurística utiliza uma representação de solução fundamentalmente diferente:
- **GA**: Codificação por permutação direta — espaço de busca natural para TSP (n! permutações)
- **PSO**: *Random keys* (codificação contínua [0,1]ⁿ) — mapeamento indireto, espaço de busca contínuo com mapeamento para permutações
- **ACO**: Matriz feromonal 3D τ[i][j][k] — representação construtiva probabilística

A qualidade relativa observada (ACO > GA > PSO) pode refletir não apenas a eficácia intrínseca do algoritmo, mas também a **expressividade da codificação** para o domínio. O PSO, em particular, opera em espaço contínuo e precisa de um *decoder* (random keys → permutação) que pode introduzir perda de informação ou vieses. Isso é uma **ameaça à validade de construto**: não está claro se estamos medindo "qualidade do método de busca" ou "adequação da codificação ao problema".

*Recomendação*: Discutir explicitamente esta limitação como ameaça à validade de construto. Uma análise de sensibilidade usando codificações alternativas (ex: GA com random keys, PSO com permutação via *swap* representation) fortaleceria as conclusões.

## 4. Análise do Lower Bound

### 4.1 Qualidade do LB

O gap médio de **51.36%** (mín. 40.79%, máx. 65.35%) em relação ao ótimo é **extremamente frouxo**. A relaxação de atribuição (Hungarian) sobre a matriz C'[j][k] = min_i G[i][j][k] ignora a estrutura sequencial do ângulo de patrulha, resultando em um limite que, embora válido (nunca excede o ótimo, verificado em 18 instâncias), é pouco informativo.

**Utilidade prática**: Um gap de 51% não discrimina entre soluções de qualidade. Por exemplo, se ACO tem gap de 0.43% versus o ótimo, mas gap de ~51% versus o LB, o valor informacional do LB é marginal. A inclusão do LB como referência na análise principal pode ser questionada.

*Recomendação*: Manter o LB apenas como garantia de corretude (nunca superestima o ótimo), mas não usá-lo como métrica primária de qualidade. A métrica principal deve ser gap vs ótimo (para instâncias com ótimo conhecido) e gap vs melhor solução conhecida (BKS) para instâncias maiores.

### 4.2 Validação do LB

A verificação de que o LB "nunca excede o ótimo (verificado em 18 instâncias)" é uma **validação de corretude** importante e deve ser mantida. No entanto, 18 instâncias pequenas não garantem que o LB seja válido para instâncias de 50 ou 100 pontos — isso é uma extrapolação assumida, não verificada.

## 5. Effect Sizes: Ausência Crítica

**Nenhum effect size é reportado**. Esta é uma omissão metodológica significativa. O Friedman F = 293.22 com p ≈ 10⁻³¹ estabelece que existe uma diferença, mas não quantifica **quão grande** ela é.

Métricas recomendadas que deveriam ser adicionadas:

| Effect size | Escala | Interpretação |
|-------------|--------|---------------|
| Kendall's W = F_stat / (n_blocks × (k−1)) | [0, 1] | Concordância entre blocos; W > 0.5 = large |
| Cliff's δ para pares ACO vs GA | [−1, 1] | Probabilidade de ACO < GA menos probabilidade inversa |
| η² rank (Glass) | [0, 1] | Proporção da variância entre métodos explicada |

*Impacto prático*: Com 1.530 observações por método (51 sementes × 30 instâncias), mesmo um efeito trivialmente pequeno (eg. 0.1% de diferença no gap médio) atingiria significância estatística. Sem effect size, o leitor não pode avaliar se a diferença ACO-GA (0.43% vs 5.28% de gap) é substancial — neste caso é, mas o princípio permanece.

## 6. Generalização do Gap: 18 Instâncias com Ótimo

**Ameaça à validade externa**: Apenas 18 das 30 instâncias (60%) possuem solução ótima via *BruteForce*. Todas são de pequeno porte (10–15 pontos). Conclusões sobre gap dos métodos (0.43%, 5.28%, 22.29%) são **derivadas exclusivamente dessas instâncias pequenas** e não podem ser extrapoladas para instâncias maiores (20–100 pts) sem evidência adicional.

O que se pode dizer com segurança:
- Para n ≤ 15, ACO é significativamente melhor que GA, que é melhor que PSO.
- O gap de ACO (0.43%) sugere que para instâncias pequenas, ACO encontra o ótimo na grande maioria dos casos.
- **Não se pode afirmar** que o gap de ACO se manteria abaixo de 1% para n = 100.

*Recomendação*: Separar claramente a análise do gap (apenas instâncias com ótimo, n ≤ 15) da análise de escalabilidade (todas as 30 instâncias, usando BKS como referência). Reportar gaps como "gap vs ótimo (n ≤ 15)" e "gap vs BKS (todas as instâncias)".

## 7. Ameaças à Validade

Com base nas informações disponíveis, as ameaças à validade parecem **insuficientemente discutidas** ou ausentes. Utilizando o framework Shadish, Cook & Campbell (2002):

| Tipo | Ameaça identificada | Gravidade |
|------|---------------------|-----------|
| **Validade interna** | Instâncias sintéticas sem ruído; viés de implementação (código não auditado) | Baixa |
| **Validade externa** | 30 instâncias sintéticas em [-1,1]; apenas um cenário de penalidade angular; parâmetros fixos sem sensibilidade | **Alta** |
| **Validade de construto** | Confundimento encoding/algoritmo (Seção 3); definição de *hit rate* não especificada (ótimo exato vs ε-tolerância) | **Alta** |
| **Validade de conclusão** | Ausência de effect sizes; IC não reportados; generalização do gap a partir de n ≤ 15 | **Média** |

*Recomendação*: Adicionar seção explícita "Ameaças à Validade" na metodologia, discutindo cada uma das acima e propondo mitigação ou reconhecendo limitação.

## 8. Análise de Sensibilidade de Hiperparâmetros

Não há evidência de que tenha sido realizada. Os parâmetros (GA: pop=100, mut=0.05; PSO: w=0.7, c1=c2=2.0; ACO: α=1.0, β=2.0, ρ=0.2) são fixos. Em uma comparação justa entre métodos, os parâmetros ideais de cada algoritmo devem ser calibrados (ex: grid search, random search, ou irace) para que a comparação reflita o **potencial máximo** de cada método, não o desempenho de uma configuração arbitrária.

*Impacto*: Se ACO é menos sensível a parâmetros que GA/PSO, a comparação favorece ACO não por mérito algorítmico, mas por acaso da configuração escolhida. Esta é uma **ameaça à validade externa grave**.

*Recomendação*: Realizar análise de sensibilidade ou, no mínimo, justificar cada escolha paramétrica com referência à literatura. Idealmente, usar calibração automática (irace, SMAC, ou grid search preliminar).

## 9. Análise de Convergência

O packet fornecido não detalha a análise de convergência. Para uma monografia sobre otimização, espera-se:

- **Trajetórias de convergência**: Plotar *best-so-far* vs iteração para cada método-instância-semente, com envelopes de confiança (sombreado mean±sd sobre sementes)
- **Métricas de diversidade**: Diversidade populacional ao longo das iterações (ex: diversidade de permutações para GA, diversidade de posições para PSO, entropia da matriz feromonal para ACO)
- **Detecção de estagnação**: O algoritmo converge em quantas iterações? O critério de parada (100 iterações) é suficiente?
- **Análise de *early stopping***: Todos os métodos se beneficiariam de mais iterações ou alguns já estagnaram?

*Recomendação*: Se estas análises existem na monografia mas não constam no packet, ignorar este ponto. Caso contrário, são necessárias para sustentar conclusões sobre eficiência.

## 10. Tempo de Otimização

Os tempos reportados (ACO ~4262ms, GA ~38ms, PSO ~75ms para 100 pts) precisam de:

1. **Medidas de variabilidade**: Mean ± SD sobre as 51 sementes — o tempo de ACO pode variar substancialmente dependendo da trajetória de construção de soluções.
2. **Especificações de hardware**: CPU, RAM, sistema operacional — sem isso, os tempos não são reproduzíveis.
3. **Unidade clara**: ms está ok, mas confirmar se é *wall-clock* ou *CPU time*.
4. **Análise de complexidade empírica**: Plotar tempo vs tamanho da instância em escala log-log para confirmar complexidade assintótica (O(n²) para GA, O(n³·pop) para ACO?).

A diferença de ~110× entre ACO e GA no tempo para 100 pts é enorme e tem implicações práticas: para uma aplicação real, GA pode ser preferível mesmo com gap maior, se o tempo de otimização for crítico.

## 11. Intervalos de Confiança

Nenhum intervalo de confiança é reportado. Para uma pesquisa que se pretende científica, **toda estimativa pontual deve vir acompanhada de IC 95%**:
- Gap médio de ACO: 0.43% [IC 95%: ?]
- Hit rate de ACO: 71.79% [IC 95%: ?] — especialmente importante pois é uma proporção, que tem variabilidade conhecida.
- Tempo médio de ACO para 100 pts: ~4262ms [IC 95%: ?]

*Recomendação*: Adicionar IC 95% para todas as métricas principais, preferencialmente via bootstrap (distribuição não-normal).

## 12. Interpretação de Figuras

Não posso avaliar figuras específicas sem vê-las, mas ofereço critérios mínimos que qualquer figura deve satisfazer:

- **Boxplots**: Devem mostrar a mediana (não média), IQR, *whiskers* (1.5×IQR), e *outliers*. Múltiplos boxplots lado a lado (um por método) para cada métrica.
- **CD Diagram (Demsar, 2006)**: O diagrama de diferenças críticas de Nemenyi deve ser incluído. CD = 0.6050 significa que a diferença entre ranks médios para ser significativa deve exceder este valor. ACO (1.10), GA (1.90), PSO (3.00) — diferenças > 0.605, logo todos os pares são significativos. OK.
- **Scatter plots**: Se mostram makespan vs instância ou tamanho da instância, devem usar escala apropriada (log para tempos) e cores/marcadores distintos para cada método.
- **Heatmaps**: Devem ter barra de cor com escala clara e rótulos. Cuidado com escalas divergentes que podem distorcer percepção.
- **Gráficos de convergência**: Média ± DP sobre sementes com sombreado; eixo x = iteração, eixo y = melhor custo da população (ou custo do melhor indivíduo).

## Síntese dos Problemas por Gravidade

| Gravidade | Problemas |
|-----------|-----------|
| 🔴 Crítico | Effect sizes ausentes; confundimento encoding/algoritmo; generalização do gap (n ≤ 15 → n ≤ 100); sem análise de sensibilidade paramétrica |
| 🟡 Moderado | Power analysis ausente; IC ausentes; timing sem variância; sem análise de convergência detalhada; hardware não especificado |
| 🟢 Leve | LB frouxo (51%) mas válido; 30 instâncias sintéticas (aceitável para TCC, idealmente mais diversas) |


## Checklist do Rigor Gate (First-Pass)

| Requisito | Atendido? | Observação |
|-----------|-----------|------------|
| Suposições do teste verificadas | ✓ | Friedman: bloco aleatório, ≥5 blocos, dados ordinais |
| Correção de comparações múltiplas | ✓ | Bonferroni-Holm corretamente aplicado |
| Ameaças à validade discutidas | ✗ | Ausentes ou insuficientes |
| Effect sizes reportados | ✗ | Nenhum |
| Intervalos de confiança reportados | ✗ | Nenhum |
| Sensibilidade paramétrica realizada | ✗ | Parâmetros fixos, não calibrados |
| Análise de poder | ✗ | Não realizada |
| Critério de convergência definido | ? | Depende do texto completo da monografia |
