# Verdict R2 — Coverage (Revisor de Literatura)

## Position Restated (R1)

**Veredito R1:** FAIL (HIGH)

**Achados próprios (13):**
| ID | Achado | Severidade |
|----|--------|-----------|
| f-council-lit-001 | Seção 2.8 (Trabalhos Futuros) sem nenhuma citação — falha estrutural | CRITICAL |
| f-council-lit-002 | Viés ACO>PSO — referencial desbalanceado; PSO sub-representado | HIGH |
| f-council-lit-003 | Rating 0 (não citado) em múltiplos artigos do vault — inconsistência entre base de conhecimento e texto | HIGH |
| f-council-lit-004 | Demšar (2006) — artigo metodológico central para Friedmann/Nemenyi — sem nota no vault, sem resumo | HIGH |
| f-council-lit-005 | Rastreabilidade zero entre claims da monografia e evidências do vault | HIGH |
| f-council-lit-006 | Artigos ACO+deep-learning irrelevantes para o escopo TSP clássico | MEDIUM |
| f-council-lit-007 | 13 artigos no vault sem entrada BibTeX correspondente | MEDIUM |
| f-council-lit-008 | Qualidade questionável de veículos de publicação (ex.: Congresso Brasileiro de Iniciação Científica) | MEDIUM |
| f-council-lit-009 | Simulated Annealing completamente ausente do referencial | MEDIUM |
| f-council-lit-010 | 18 artigos no vault com status "pendente" — sem resumo, sem avaliação | MEDIUM |
| f-council-lit-011 | Categorização dos artigos do vault não mapeada para seções da monografia | LOW |
| f-council-lit-012 | Seção 2.4 sem citações — lacuna de cobertura localizada | LOW |
| f-council-lit-013 | wang2021 com rating 0 — inconsistência com uso do artigo no texto | LOW |

---

## Debate Notes

### STEEL-MAN — Argumento mais forte de outro juiz

**TSP: f-council-002 (CRITICAL) — ACO implementado é Ant System, NÃO MMAS.**

O Especialista TSP identificou que o código deposita feromônio de TODAS as formigas, sem bounds de feromônio — características do Ant System original (Dorigo 1992), não do MMAS (Stützle & Hoos 2000). A nota do vault afirma incorretamente que MMAS está implementado. Isso transcende minha análise de cobertura por três razões:

1. **Erro factual na monografia**: não é apenas ausência de citação — o texto afirma usar um algoritmo que não foi implementado. As referências a artigos de MMAS no referencial são portanto *misdirected*: descrevem um método que não corresponde ao código.
2. **Corrupção da base de conhecimento**: se a nota do vault está factualmente errada, toda a cadeia vault→monografia está contaminada. Meu achado f-council-lit-003 (rating 0 inconsistente) é sintoma de um problema mais profundo.
3. **Cascata de cobertura**: se o método real é Ant System e não MMAS, o referencial deveria cobrir Ant System (Dorigo 1992, Dorigo et al. 1996), Stützle & Hoos (1997, 2000) para contraste AS vs MMAS, e benchmarks de comparação AS vs MMAS. Nada disso está presente.

Este achado reformula minha preocupação com cobertura: o problema não é só *quantidade* ou *qualidade* das referências, mas *pertinência* — as referências presentes apontam para o algoritmo errado.

### CHALLENGE — Afirmação de outro juiz que considero superestimada

**Metodologista: f-council-011 — Poder estatístico do Friedmann com N=30, k=3 vs recomendação Demšar de k≥4 (severidade HIGH).**

Discordo da severidade HIGH por três razões:

1. **Demšar (2006) não estabelece k≥4 como regra rígida**: a recomendação original é que o Friedmann "requires at least 5 datasets and at least 4 algorithms", mas o próprio autor reconhece que o teste funciona com k≥3. A perda de poder com k=3 é real, mas não invalida o teste.
2. **N=30 compensa parcialmente k pequeno**: com 30 instâncias, a estatística Q de Friedmann tem poder razoável mesmo com k=3. O problema de poder é mais grave quando N é pequeno E k é pequeno — não é o caso aqui.
3. **Prática comum na literatura**: inúmeros artigos publicados em periódicos A1 usam Friedmann com k=3 e N≈30, e passam por revisão por pares. Classificar como HIGH equivale a invalidar uma fração significativa da literatura experimental em otimização.

Reclassificaria como **MEDIUM**: é uma limitação que deve ser reconhecida (e idealmente mitigada com Wilcoxon pareado como complemento), mas não é uma falha metodológica grave que invalida os resultados.

### ACKNOWLEDGE — Ponto de outro juiz que reforça minha análise

**TSP: f-council-003 (CRITICAL) — ATSP e TDTSP completamente ausentes do referencial.**

Este achado reforça diretamente meu f-council-lit-001 (Seção 2.8 sem citações) e f-council-lit-002 (viés de cobertura). O Especialista TSP demonstra que o problema estudado (TSP-SD-ATP com características assimétricas e dependentes de sequência) não encontra respaldo em nenhuma classe de literatura relevante:

- **ATSP (Asymmetric TSP)**: ausente, apesar do problema modelado ser inerentemente assimétrico (custos de voo entre pontos não são simétricos com vento, terreno, etc.)
- **TDTSP (Time-Dependent TSP)**: ausente, apesar da dependência temporal ser modelada via makespan
- **SDTSP (Sequence-Dependent TSP)**: marginalmente coberto, mas sem artigos específicos

Isso confirma que minha avaliação de FAIL não é excessivamente severa: o referencial tem lacunas *estruturais* de enquadramento, não apenas omissões pontuais. A cobertura não cobre o problema que o TCC diz resolver.

---

## Revised Verdict

**VEREDITO FINAL: FAIL (HIGH)**

**Tri-consenso**: Coverage (FAIL, HIGH), Metodologista (FAIL, HIGH), Especialista TSP (FAIL, HIGH).

**Síntese dos problemas fatais (cross-judge):**

| # | Problema | Juízes |
|---|----------|--------|
| 1 | Seção 2.8 sem citações — vácuo de referencial em seção estrutural | Coverage, TSP |
| 2 | ATSP/TDTSP ausentes — referencial não cobre a classe do problema estudado | Coverage, TSP |
| 3 | ACO implementado é Ant System, não MMAS — erro factual na monografia | TSP, Coverage (steel-man) |
| 4 | Confundimento método×codificação — validade interna comprometida | Metodologista |
| 5 | Sem literatura de tuning — parâmetros fixos sem justificativa | Metodologista |
| 6 | Demšar (2006) sem nota no vault — gap metodológico não documentado | Coverage |
| 7 | Rastreabilidade zero vault→monografia — claims sem evidência | Coverage |
| 8 | 13 artigos sem BibTeX + 5 descartados com BibTeX ativo | Coverage, Metodologista |

**Recomendações para resolução (pré-requisitos para PASS):**

1. **Emergencial**: Corrigir a nota do vault sobre MMAS → Ant System; adicionar referências de Ant System (Dorigo 1992, 1996) e MMAS (Stützle & Hoos 1997, 2000).
2. **Estrutural**: Adicionar ao menos 3-5 artigos de ATSP e 2-3 de TDTSP ao referencial; criar subseção 2.x dedicada a essas variantes.
3. **Cobertura**: Preencher Seção 2.8 com citações reais de trabalhos futuros da literatura; adicionar SA ao referencial.
4. **Base de conhecimento**: Completar os 18 artigos pendentes OU removê-los; resolver os 13 sem BibTeX; corrigir ratings inconsistentes; criar nota para Demšar (2006).
5. **Rastreabilidade**: Mapear cada claim da monografia para evidência no vault (matriz claim-evidence).

---

## JSON Verdict

```json
{
  "type": "verdict",
  "verdict": "FAIL",
  "confidence": "HIGH",
  "file": ".agents/council/2026-06-04-validate-referencial-judge-revisor-lit-R2.md",
  "debate_notes": {
    "steel_man": {
      "source": "TSP judge",
      "finding": "f-council-002",
      "summary": "ACO implementado é Ant System, não MMAS. Código mostra depósito de todas as formigas sem bounds de feromônio. Nota do vault afirma incorretamente MMAS. Erro factual que torna referências a MMAS misdirected.",
      "impact_on_my_analysis": "Reformula minha preocupação de cobertura: referências presentes apontam para algoritmo errado, não apenas ausentes."
    },
    "challenge": {
      "target": "Metodologista",
      "finding": "f-council-011",
      "claim": "Poder estatístico do Friedmann com N=30, k=3 vs recomendação Demšar de k>=4 com severidade HIGH",
      "counter": "Demšar não estabelece k>=4 como regra rígida; N=30 compensa parcialmente k=3; prática comum na literatura. Reclassificaria como MEDIUM.",
      "severity_disagreement": "HIGH -> MEDIUM"
    },
    "acknowledge": {
      "source": "TSP judge",
      "finding": "f-council-003",
      "summary": "ATSP e TDTSP completamente ausentes do referencial, apesar do problema TSP-SD-ATP ser assimétrico e com dependência de sequência.",
      "strengthens_my": ["f-council-lit-001", "f-council-lit-002"],
      "why": "Confirma que as lacunas de cobertura são estruturais (classe de problema não coberta), não pontuais."
    },
    "cross_judge_synthesis": {
      "consensus": "Tri-consenso FAIL (HIGH)",
      "all_judges": ["Coverage (FAIL, HIGH)", "Metodologista (FAIL, HIGH)", "Especialista TSP (FAIL, HIGH)"],
      "fatal_issues_count": 8,
      "key_new_insight": "Steel-man do TSP revela que o referencial não apenas omite artigos — referencia o algoritmo errado (MMAS vs Ant System)"
    }
  }
}
```
