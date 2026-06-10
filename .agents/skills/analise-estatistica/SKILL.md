---
name: analise-estatistica
description: Toolkit completo de análise estatística para o TCC. Gera código com revisão em 4 rounds, seleciona testes apropriados, executa análises (t-test, ANOVA, chi-square, regressão, correlação, Bayesiana), calcula effect sizes, power analysis, e produz relatórios APA. Consolida data-analysis e statistical-analysis.
---

# Análise Estatística

Skill consolidada para análise estatística dos experimentos do TCC.

---

## 1. Workflow de Geração de Código

Adaptado de `data-analysis`.

### Estrutura do Código

```python
# IMPORT — pandas, numpy, scipy, statsmodels, sklearn
# LOAD DATA — carregar dos arquivos originais
# DATASET PREPARATIONS — missing values, unidades, critérios de exclusão
# DESCRIPTIVE STATISTICS — tabelas sumário se necessário
# PREPROCESSING — dummy variables, normalização
# ANALYSIS — testes estatísticos por hipótese
# SAVE ADDITIONAL RESULTS — resultados extras para pickle
```

### 4-Round Code Review

1. **Round 1 — Code Flaws**: erros matemáticos/estatísticos, cálculos errados, testes triviais
2. **Round 2 — Data Handling**: missing values, unidades, pré-processamento, escolha de teste
3. **Round 3 — Per-Table**: valores sensíveis, medidas de incerteza, dados faltantes
4. **Round 4 — Cross-Table**: completude, consistência, variáveis faltantes

### Pacotes Permitidos

`pandas`, `numpy`, `scipy`, `statsmodels`, `sklearn`, `pickle`, `pingouin`, `pymc`, `arviz`

---

## 2. Seleção de Testes

### Referência Rápida

**Comparando 2 grupos:**
- Independente, contínuo, normal → t-test independente
- Independente, contínuo, não-normal → Mann-Whitney U
- Pareado, contínuo, normal → t-test pareado
- Pareado, contínuo, não-normal → Wilcoxon signed-rank
- Binário → Chi-square ou Fisher's exact

**Comparando 3+ grupos:**
- Independente, contínuo, normal → ANOVA one-way
- Independente, contínuo, não-normal → Kruskal-Wallis
- Pareado, contínuo, normal → ANOVA medidas repetidas
- Pareado, contínuo, não-normal → Friedman

**Relações:**
- Duas variáveis contínuas → Pearson (normal) ou Spearman (não-normal)
- Outcome contínuo com preditores → Regressão linear
- Outcome binário com preditores → Regressão logística

---

## 3. Verificação de Suposições

**SEMPRE verifique suposições antes de interpretar resultados.**

### O que fazer quando violadas

**Normalidade violada:**
- Violação leve + n > 30 por grupo → prossiga com paramétrico (robusto)
- Violação moderada → use alternativa não-paramétrica
- Violação severa → transforme dados ou use não-paramétrico

**Homogeneidade de variância violada:**
- t-test → use Welch's t-test
- ANOVA → use Welch's ANOVA ou Brown-Forsythe

**Linearidade violada (regressão):**
- Adicione termos polinomiais
- Transforme variáveis
- Use modelos não-lineares ou GAM

---

## 4. Effect Sizes

| Teste | Effect Size | Pequeno | Médio | Grande |
|-------|-------------|---------|-------|--------|
| t-test | Cohen's d | 0.20 | 0.50 | 0.80 |
| ANOVA | η²_p | 0.01 | 0.06 | 0.14 |
| Correlação | r | 0.10 | 0.30 | 0.50 |
| Regressão | R² | 0.02 | 0.13 | 0.26 |
| Chi-square | Cramér's V | 0.07 | 0.21 | 0.35 |

Sempre reporte CIs para effect sizes.

---

## 5. Power Analysis

### A Priori (planejamento)

```python
from statsmodels.stats.power import tt_ind_solve_power

n = tt_ind_solve_power(effect_size=0.5, alpha=0.05, power=0.80,
                        ratio=1.0, alternative='two-sided')
```

### Sensitivity (pós-estudo)

Determine qual effect size você poderia detectar com o N disponível.

**Nota:** power analysis post-hoc geralmente não é recomendada. Use sensitivity analysis.

---

## 6. Exemplos de Código

### t-test com pingouin

```python
import pingouin as pg

result = pg.ttest(group_a, group_b, correction='auto')
t_stat = result['T'].values[0]
df = result['dof'].values[0]
p_value = result['p-val'].values[0]
cohens_d = result['cohen-d'].values[0]
ci = result['CI95%'].values[0]

print(f"t({df:.0f}) = {t_stat:.2f}, p = {p_value:.3f}, "
      f"d = {cohens_d:.2f}, 95% CI [{ci[0]:.2f}, {ci[1]:.2f}]")
```

### ANOVA com post-hoc

```python
aov = pg.anova(dv='score', between='group', data=df, detailed=True)
if aov['p-unc'].values[0] < 0.05:
    posthoc = pg.pairwise_tukey(dv='score', between='group', data=df)
eta_squared = aov['np2'].values[0]
```

---

## 7. Reporte APA

### t-test

```
Grupo A (n = 48, M = 75.2, DP = 8.5) pontuou significativamente
mais alto que Grupo B (n = 52, M = 68.3, DP = 9.2),
t(98) = 3.82, p < .001, d = 0.77, IC 95% [0.36, 1.18].
```

### ANOVA

```
ANOVA one-way revelou efeito principal significativo,
F(2, 147) = 8.45, p < .001, η²_p = .10. Post hoc Tukey HSD:
Condição A (M = 78.2, DP = 7.3) > Condição B (M = 71.5,
DP = 8.1, p = .002, d = 0.87) e Condição C (M = 70.1,
DP = 7.9, p < .001, d = 1.07).
```

---

## 8. Boas Práticas

1. Pré-registre análises quando possível
2. **Sempre** verifique suposições antes de interpretar
3. Reporte effect sizes com intervalos de confiança
4. Reporte **todas** as análises planejadas, inclusive não-significativas
5. Distinga significância estatística de significância prática
6. Visualize dados antes e depois da análise
7. Compartilhe dados e código para reprodutibilidade

## Armadilhas Comuns

1. **P-hacking**: não teste de múltiplas formas até achar significância
2. **HARKing**: não apresente achados exploratórios como confirmatórios
3. **Ignorar suposições**: verifique e reporte violações
4. **Confundir significância com importância**: p < .05 ≠ efeito relevante
5. **Não reportar effect sizes**: essencial para interpretação
6. **Múltiplas comparações**: corrija para family-wise error quando apropriado
7. **Interpretar mal p-valores**: NÃO são probabilidade da hipótese ser verdadeira
