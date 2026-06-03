---
title: Auditoria do Script de Análise Estatística
tags:
  - writing
  - auditoria
  - estatistica
  - metodologia
  - monografia
status: validado
created: 2026-06-02
---

# Auditoria do Script de Análise Estatística

Esta nota audita `scripts/analise-estatistica.py` sob a ótica do protocolo descrito em [[analysis-methodology]] e dos testes recomendados por Demšar (2006) para comparação de múltiplos algoritmos em múltiplas bases/instâncias.

## Resultado Executivo

O script foi corrigido e validado após esta auditoria. A versão atual de `scripts/analise-estatistica.py` usa mediana por método×instância, calcula Friedman/Iman-Davenport com ranks médios, usa `q=2.343` para Nemenyi na escala de Demšar, implementa Wilcoxon signed-rank pareado com Bonferroni-Holm e regenera `cd-diagram.{svg,png}`.

Os valores finais validados são: ranks médios ACO `1.1000`, GA `1.9000`, PSO `3.0000`; `χ²_F=54.6000`; `F_F(2,58)=293.2222`; `p=4.710129e-31`; `CD=0.6050`. Todas as comparações par-a-par são significativas por Nemenyi e por Wilcoxon/Holm.

> [!note] Nota sobre divergência com a auditoria inicial
> A auditoria inicial registrava ranks médios ACO `1.0333`, GA `1.9667`, PSO `3.0000`. A validação final sobre a cobertura atual de `4638` summaries retornou ACO `1.1000`, GA `1.9000`, PSO `3.0000`. Usar os valores finais desta seção e de [[analysis-methodology]].

| Item | Status | Consequência |
|---|---|---|
| Unidade experimental | Parcialmente correta | Usa instância como bloco, mas a saída comunica número errado de execuções |
| Valor-resumo por método×instância | Divergente | Usa média; [[analysis-methodology]] declara mediana |
| Estatística de Friedman | Incorreta | Usa somas de ranks na fórmula para ranks médios, inflando `χ²_F` |
| Iman-Davenport | Dependente do erro anterior | O valor atual `inf` decorre da estatística inflada/perfeição artificial dos ranks por média |
| Nemenyi/CD | Conservador demais | Usa `q = 3.314`; na forma de Demšar com `CD = q_alpha sqrt(k(k+1)/(6N))`, o valor para `k=3`, `alpha=0.05` é `2.343` |
| Wilcoxon/Holm | Ausente | O protocolo declarado não está implementado |
| P-valores | Ausentes | O script não reporta p-valor de Friedman nem p-valores pareados |
| Figura | Parcialmente válida | O diagrama existe, mas usa CD inflado e deveria ser regenerado após correção |

## Protocolo Esperado Pela Literatura

Para comparar `k` algoritmos em `N` instâncias, Demšar recomenda tratar cada instância como um bloco pareado. Em cada instância, os algoritmos recebem ranks conforme o desempenho: rank 1 para o menor makespan, rank 2 para o segundo, e assim por diante. Empates devem receber rank médio.

Para o teste global de Friedman, usando ranks médios `R_j`, a estatística é:

$$
\chi_F^2 = \frac{12N}{k(k+1)}\left[\sum_{j=1}^{k} R_j^2 - \frac{k(k+1)^2}{4}\right]
$$

Forma equivalente usando somas de ranks `S_j`:

$$
\chi_F^2 = \frac{12}{Nk(k+1)}\sum_{j=1}^{k} S_j^2 - 3N(k+1)
$$

O pós-teste de Nemenyi compara diferenças entre ranks médios. A diferença crítica é:

$$
CD = q_\alpha \sqrt{\frac{k(k+1)}{6N}}
$$

Na tabela usada por Demšar para essa fórmula, `q_alpha = 2.343` para `k=3` e `alpha=0.05`. O valor `3.314` é o quantil da distribuição Studentized range em outra escala; usá-lo diretamente nessa fórmula equivale a multiplicar o CD por `sqrt(2)`.

## Problemas Encontrados no Código

### 1. Friedman Calculado com Fórmula Incompatível na Versão Antiga

O trecho crítico está em `scripts/analise-estatistica.py:123-127`:

```python
avg_ranks = {m: rank_sums[m] / N for m in METHODS}

sum_Rj_sq = sum(rank_sums[m] ** 2 for m in METHODS)
chi2 = (12 * N) / (k * (k + 1)) * (sum_Rj_sq - k * (k + 1) ** 2 / 4)
```

O script calcula `avg_ranks`, mas usa `rank_sums` em `sum_Rj_sq`. A fórmula usada exige ranks médios. Se preferir usar somas, o fator correto é `12 / (N*k*(k+1))`, com subtração `3*N*(k+1)`.

Impacto observado na versão antiga:

| Cálculo | Valor |
|---|---:|
| `χ²_F` reportado pelo script | 377640.0000 |
| `χ²_F` correto usando mediana por instância | 54.6000 |
| `F_F` Iman-Davenport correto usando mediana | 293.2222 |

O valor `377640` é impossível de defender para `N=30` e `k=3`; ele resulta da aplicação da fórmula em escala errada.

### 2. Média Usada Onde o Protocolo Declara Mediana na Versão Antiga

O protocolo em [[analysis-methodology]] define mediana do makespan por instância como valor-resumo. A versão antiga do script usava média em `scripts/analise-estatistica.py`:

```python
means = {m: statistics.mean(data[inst][m]) for m in METHODS}
```

Isso muda os ranks. Com médias, o script obtém ordenação perfeita `ACO < GA < PSO` nas 30 instâncias. Com medianas, a ordenação fica:

| Ordenação por instância | Quantidade |
|---|---:|
| ACO, GA, PSO | 29 |
| GA, ACO, PSO | 1 |

Na comparação pareada ACO×GA por mediana, há 25 instâncias com ACO melhor, 1 com GA melhor e 4 empates. Portanto, a afirmação “ACO rank 1 em todas as instâncias” é verdadeira para a média atual, mas não para a mediana declarada no protocolo.

### 3. Valor Crítico de Nemenyi em Escala Errada na Versão Antiga

A versão antiga do dicionário `NEMENYI_Q` usava `3.314` para `(3, 0.05)`. Para a fórmula de CD adotada no script, o valor compatível com Demšar é `2.343`.

Impacto nos dados atuais:

| CD | Valor |
|---|---:|
| CD correto com `q=2.343` | 0.6050 |
| CD atual com `q=3.314` | 0.8557 |

Mesmo com o CD inflado, as diferenças atuais continuam passando no Nemenyi. O problema é de reprodutibilidade e fidelidade metodológica: a figura e o texto reportam um limiar diferente do algoritmo descrito na literatura.

### 4. Wilcoxon e Holm Ausentes na Versão Antiga

[[analysis-methodology]] declara Wilcoxon signed-rank pareado com correção Bonferroni-Holm. A versão antiga não implementava essa etapa; a versão validada implementa e confirma significância nos três pares.

Valores recalculados com medianas por instância, removendo empates no Wilcoxon:

| Par | n não-zero | W | p exato bilateral | Interpretação |
|---|---:|---:|---:|---|
| ACO×GA | 26 | 1.0 | 5.96e-08 | Diferença significativa após Holm |
| ACO×PSO | 30 | 0.0 | 1.86e-09 | Diferença significativa após Holm |
| GA×PSO | 30 | 0.0 | 1.86e-09 | Diferença significativa após Holm |

Esses valores indicam que a conclusão de diferença entre métodos provavelmente permanece, mas ela precisa ser obtida por uma implementação correta, não pelo script atual.

### 5. Saída Textual Informava Número Errado de Execuções

A versão antiga imprimia `Execuções por instância: 30`. Nos dados atuais, cada método estocástico tem 51 sementes por instância. O número 30 é o número de instâncias/blocos.

Texto correto: `Instâncias/blocos: 30; sementes por método e instância: 51`.

### 6. P-valor do Teste Global Ausente na Versão Antiga

O script antigo reportava `χ²_F`, `F_F` e CD, mas não reportava p-valor. A versão validada implementa a função de sobrevivência da distribuição F por beta incompleta regularizada e reporta `p=4.710129e-31`.

```text
Friedman/Iman-Davenport: F_F(2, 58) = 293.2222, p = 4.710129e-31.
```

Sem SciPy, o p-valor do Iman-Davenport exige implementar a CDF/sobrevivência da distribuição F; a versão atual faz isso sem dependências externas.

## Valores Finais Validados

Usando 30 instâncias, três métodos (`ACO`, `GA`, `PSO`) e mediana do makespan por método×instância:

| Métrica | ACO | GA | PSO |
|---|---:|---:|---:|
| Soma de ranks | 33.0 | 57.0 | 90.0 |
| Rank médio | 1.1000 | 1.9000 | 3.0000 |

| Estatística | Valor |
|---|---:|
| `N` | 30 |
| `k` | 3 |
| `χ²_F` correto | 54.6000 |
| `F_F` Iman-Davenport | 293.2222 |
| p-valor Iman-Davenport | 4.710129e-31 |
| CD Nemenyi correto (`alpha=0.05`) | 0.6050 |

Diferenças de ranks médios:

| Par | Diferença | Comparação com CD=0.6050 |
|---|---:|---|
| ACO×GA | 0.8000 | Significativa |
| ACO×PSO | 1.9000 | Significativa |
| GA×PSO | 1.1000 | Significativa |

## Correções Aplicadas

1. `statistics.mean` substituído por `statistics.median`.
2. Friedman corrigido para usar ranks médios.
3. `NEMENYI_Q[(3, 0.05)]` corrigido para `2.343`.
4. p-valor global implementado sem SciPy.
5. Wilcoxon signed-rank pareado e correção Holm implementados.
6. `cd-diagram.svg` e `cd-diagram.png` regenerados.
7. Saída textual corrigida para distinguir instâncias/blocos de sementes por instância.
8. [[analysis-methodology]] atualizada com os valores finais.

## Decisão Para a Escrita

A monografia pode afirmar que os métodos diferem estatisticamente nas configurações avaliadas. A redação deve manter a limitação de que essa conclusão vale para as implementações, parâmetros e instâncias deste estudo, não para todas as possíveis variantes de GA, PSO e ACO.
