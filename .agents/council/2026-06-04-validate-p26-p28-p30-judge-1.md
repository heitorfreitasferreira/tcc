```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "findings": [
    {
      "id": "F01",
      "task": "P26",
      "fix": "5 BibTeX entries added: heldkarp1970traveling, heldkarp1971traveling, johnson1996asymptotic, kinable2017hybrid, righini2021efficient",
      "why": "All entries present at lines 662–711 of abntex2-references.bib. Authors, titles, journals, years, and DOIs match canonical references: Held & Karp (1970) Operations Research 18(6):1138–1162; Held & Karp (1971) Mathematical Programming 1(1):6–25; Johnson, McGeoch & Rothberg (1996) SODA pp.341–350; Kinable, Cire & van Hoeve (2017) EJOR 259(3):906–918; Righini (2021) OJMO 2:1–27. All cited in the new Held-Karp paragraph.",
      "ref": "abntex2-references.bib:662–711"
    },
    {
      "id": "F02",
      "task": "P26",
      "fix": "New paragraph in Seção 2.8 (Limitantes Inferiores) at fundamentacao.tex:92",
      "why": "Paragraph correctly cites Held-Karp as canonical lower bound reference, quantifies HK tightness (~1% gap for Euclidean instances via Johnson 1996), notes ongoing algorithmic improvements (Righini 2021), and justifies AP over HK: adapting HK to the 3D tensor c(i,j,k) would require three-index Lagrangian penalties, out of scope. The justification is technically sound.",
      "ref": "fundamentacao.tex:88–92"
    },
    {
      "id": "F03",
      "task": "P26",
      "fix": "LaTeX compilation blocker: `\\<1\\%` on line 92",
      "why": "The fragment `\\<1\\%` in text mode contains `\\<` which is not a valid LaTeX control sequence in standard LaTeX/abntex2. This will produce 'Undefined control sequence' error during pdflatex compilation. The backslash before `<` must be removed: `(<1\\%)` or replaced with `(\\textless 1\\%)` if the font encoding requires it.",
      "ref": "fundamentacao.tex:92"
    },
    {
      "id": "F04",
      "task": "P28",
      "fix": "BibTeX oncan2009comparative added + sentence in Seção 2.1",
      "why": "Öncan, Altınel & Laporte (2009) is a well-known ATSP survey in Computers & OR 36(3):637–654. Kinable et al. (2017) covers TDTSP (also added via P26). The sentence at line 13 correctly frames TSP-SD-ATP: the asymmetry arises because transition cost depends on the preceding node (making the effective cost matrix asymmetric); the time-dependency captures the sequence-dependence. Both citations are present and correctly placed.",
      "ref": "abntex2-references.bib:714–723, fundamentacao.tex:13"
    },
    {
      "id": "F05",
      "task": "P30",
      "fix": "Rating fixes: winter2002modeling 0→3, wang2021ant 0→3",
      "why": "Both papers are cited in the fundamentação (winter2002modeling at fundamentacao.tex:18 for turn-cost modeling; wang2021ant is in BibTeX and rated for ACO parameter tuning). Rating 3 (on apparent 0–5 scale) is appropriate for papers used as conceptual background or secondary evidence. bean1994genetic confirmed at rating 4; vanhove2012route confirmed at rating 3 — neither required changes.",
      "ref": "winter2002modeling.md:11, wang2021ant.md:11, bean1994genetic.md:11, vanhove2012route.md:11"
    },
    {
      "id": "F06",
      "task": "P38",
      "fix": "\\autoref{} used consistently in fundamentacao.tex",
      "why": "Lines 4–5 show systematic use of `\\autoref{sec:tsp}`, `\\autoref{sec:custo-dependente}`, etc. replacing manual `Capítulo~\\ref{...}` patterns. ppgco.cls defines \\autoref translations for pt-BR. No remaining raw \\ref{} found in this chapter.",
      "ref": "fundamentacao.tex:4–5"
    },
    {
      "id": "F07",
      "task": "P26/P28",
      "fix": "Roadmap tasks marked completed",
      "why": "roadmap-monografia.md lines 420, 422, 424, 432 all show 'Concluída' with accurate summaries of what was done. No false claims about completion.",
      "ref": "roadmap-monografia.md:420–432"
    }
  ],
  "recommendation": "FIX F03 before compilation: remove the spurious backslash from `\\<` on fundamentacao.tex line 92, changing it to `<`. Otherwise all P26, P28, P30, and P38 changes are correct, complete, and well-justified."
}
```

---

# Council Judge 1 — Full Analysis

## Verdict: WARN (confiança alta)

**Razão:** uma única falha bloqueante de compilação LaTeX em P26 (`\<1\%` → erro de sequência de controle indefinida). Todo o restante está correto. A correção é trivial (remover uma contrabarra).

---

## P26 — Adicionar Held-Karp BibTeX e revisar Limitantes Inferiores

### BibTeX (F01) ✓

As 5 entradas foram adicionadas corretamente ao final de `abntex2-references.bib` (linhas 662–711):

| Chave | Título | Periódico/Conferência | DOI |
|---|---|---|---|
| `heldkarp1970traveling` | The Traveling-Salesman Problem and Minimum Spanning Trees | Operations Research 18(6):1138–1162, 1970 | 10.1287/opre.18.6.1138 |
| `heldkarp1971traveling` | …Part II | Mathematical Programming 1(1):6–25, 1971 | 10.1007/BF01584070 |
| `johnson1996asymptotic` | Asymptotic Experimental Analysis for the Held-Karp Traveling Salesman Bound | SODA 1996, pp.341–350 | — |
| `kinable2017hybrid` | Hybrid Optimization Methods for Time-Dependent Sequencing Problems | EJOR 259(3):906–918, 2017 | 10.1016/j.ejor.2016.11.035 |
| `righini2021efficient` | Efficient Optimization of the Held-Karp Lower Bound | OJMO 2:1–27, 2021 | 10.5802/ojmo.11 |

Autores, títulos, volumes e DOIs conferem com as publicações canônicas. A entrada `johnson1996asymptotic` é um `@inproceedings` (SODA), sem DOI — o que é correto, pois anais de conferência da SIAM da época não tinham DOI. As demais são `@article` com DOI.

A menção a `valenzuela1997estimating` na descrição da tarefa no roadmap **não** exigiu adição, pois a entrada já existia no BibTeX (linha 560) desde antes desta tarefa. OK.

### Parágrafo na Seção 2.8 (F02) ✓

O parágrafo adicionado em `fundamentacao.tex:92`:

> A referência canônica para limitantes inferiores no TSP é a relaxação de Held--Karp \cite{heldkarp1970traveling,heldkarp1971traveling}, que utiliza árvores geradoras mínimas ponderadas por multiplicadores lagrangeanos. O limitante de Held--Karp é em média mais apertado que o AP (\<1\% do ótimo para instâncias euclidianas \cite{johnson1996asymptotic}), e algoritmos eficientes para calculá-lo continuam a ser desenvolvidos \cite{righini2021efficient}. Para o TSP-SD-ATP, entretanto, a adaptação da relaxação de Held--Karp exigiria lidar com o tensor tridimensional $c_{i,j,k}$ e com três índices na penalidade lagrangeana, o que está fora do escopo desta monografia. A relaxação AP, embora mais frouxa, é direta de implementar a partir da redução 3D$\rightarrow$2D e serve como referência inferior para as instâncias grandes.

**Análise de conteúdo:**
- Cita Held--Karp como referência canônica ✓
- Quantifica o gap do HK (~1%) citando Johnson et al. (1996) ✓
- Menciona desenvolvimento contínuo de algoritmos HK via Righini (2021) ✓
- **Justificativa AP vs HK:** a adaptação de HK exigiria lidar com o tensor 3D e penalidade lagrangeana de três índices → fora do escopo ✓
- Conclui que o AP, embora mais frouxo, é direto e serve como referência inferior ✓

A justificativa é tecnicamente correta: a relaxação de Held--Karp opera sobre 1-trees (árvores geradoras mínimas + um vértice), que exigem uma matriz de custos 2D. Adaptá-la para um tensor 3D $c_{i,j,k}$ exigiria reformular a noção de 1-tree para incluir dependência de segunda ordem, o que é um problema de pesquisa não trivial.

### Erro de compilação LaTeX (F03) ⚠️

O trecho `(\<1\%` contém `\<` (contrabarra + menor-que). Em LaTeX padrão, `\<` não é uma sequência de controle definida. O comando `pdflatex` produzirá:

```
! Undefined control sequence.
l.92 ...ais apertado que o AP (\<1
                                  \% do ótimo...
```

**Correção necessária:** remover a contrabarra espúria, alterando para `(<1\%`. O caractere `<` em modo texto é tratado corretamente com codificação T1 (padrão do `ppgco.cls` via `abntex2`). Alternativamente, usar `\textless` se a codificação for OT1.

Este é o **único bloqueador** encontrado. A correção é de 1 caractere (remover `\`).

---

## P28 — Adicionar referências ATSP/TDTSP

### BibTeX (F04) ✓

- `oncan2009comparative` adicionado (linhas 714–723): Öncan, Altınel & Laporte (2009), "A Comparative Analysis of Several Asymmetric Traveling Salesman Problem Formulations", *Computers & Operations Research* 36(3):637–654, DOI 10.1016/j.cor.2007.11.008. Este é o survey canônico de formulações ATSP — referência apropriada. ✓
- `kinable2017hybrid` já presente via P26: Kinable, Cire & van Hoeve (2017), "Hybrid Optimization Methods for Time-Dependent Sequencing Problems", *EJOR* 259(3):906–918. Cobre TDTSP com método exato baseado em *branch-and-price*. Referência apropriada para TDTSP. ✓

### Sentença na Seção 2.1 (F04) ✓

`fundamentacao.tex:13`:

> Nesse aspecto, o TSP-SD-ATP pode ser enquadrado como instância de duas classes bem estudadas: o TSP assimétrico \cite{oncan2009comparative} e o TSP dependente do tempo \cite{kinable2017hybrid}.

**Análise de enquadramento:**
- **ATSP:** O TSP-SD-ATP produz custos assimétricos efetivos porque $c_{i,j,k}$ depende do nó anterior $i$. Mesmo com distâncias euclidianas simétricas, a penalidade angular torna o custo direcional. Enquadramento correto. ✓
- **TDTSP:** A dependência da sequência temporal (nó anterior → nó atual → próximo nó) é formalmente análoga à dependência temporal do TDTSP, onde o custo de uma aresta depende do instante em que é percorrida. Enquadramento defensável. ✓

---

## P30 — Corrigir ratings inconsistentes

### Verificação (F05) ✓

| Paper | Rating antes | Rating depois | Status |
|---|---|---|---|
| `winter2002modeling` | 0 | **3** | Corrigido ✓ |
| `wang2021ant` | 0 | **3** | Corrigido ✓ |
| `bean1994genetic` | 4 | 4 | Já estava correto ✓ |
| `vanhove2012route` | 3 | 3 | Já estava correto ✓ |

Verifiquei `bean1994genetic.md:11` → `rating: 4` e `vanhove2012route.md:11` → `rating: 3`. Ambos já tinham os valores esperados antes da tarefa — a descrição no roadmap confirma: "bean1994genetic já estava 4, vanhove2012route já estava 3".

**Adequação dos ratings:**
- `winter2002modeling` (3): Citado em `fundamentacao.tex:18` como referência conceitual para custos de curva. Rating 3 é apropriado — paper relevante para contextualização, mas não é método central.
- `wang2021ant` (3): Paper sobre tuning de parâmetros ACO com SOS. Relevante como evidência de sensibilidade paramétrica do ACO, mas não é citado diretamente no texto atual da fundamentação (está no BibTeX como referência disponível). Rating 3 é apropriado.

---

## P38 — Substituir `\ref{}` por `\autoref{}`

### Verificação (F06) ✓

`fundamentacao.tex:4–5` mostra uso consistente de `\autoref{}`:

```tex
\autoref{sec:tsp} ... \autoref{sec:custo-dependente} ... \autoref{sec:drone-tsp}
```

O `ppgco.cls` (linhas 528–542) define traduções pt-BR para `\autoref` (Figura, Tabela, Capítulo, Seção etc.). Não foram encontrados `\ref{}` residuais no capítulo de fundamentação. ✓

---

## Roadmap (F07) ✓

As entradas em `roadmap-monografia.md` (linhas 420, 422, 424, 432) marcam P26, P28, P30 e P38 como "Concluída" com resumos fiéis ao trabalho realizado. ✓

---

## Conclusão

**Todas as mudanças de conteúdo estão corretas.** P26 adicionou as 5 entradas BibTeX e um parágrafo bem fundamentado justificando AP vs HK. P28 adicionou o survey ATSP canônico e enquadrou o TSP-SD-ATP corretamente. P30 corrigiu os dois ratings pendentes (os outros dois já estavam corretos). P38 aplicou `\autoref{}` consistentemente.

**O único bloqueador é F03:** `\<` em vez de `<` na linha 92 de `fundamentacao.tex`. Isso impede a compilação do LaTeX. A correção é trivial (1 caractere). Com essa correção, o veredito seria PASS.
