---
name: tcc-escrita
description: Escreve e edita a monografia TCC em LaTeX (pt-BR). Cobre redação acadêmica de alta qualidade, estrutura de paper científico, workflow completo de pesquisa (formulação→conclusão), padrões LaTeX, e checklist de qualidade. Use ao redigir, revisar, expandir ou estruturar qualquer parte da monografia. Consolida academic-writing, research-paper-writing e scientific-paper.
---

# TCC Escrita — Redação da Monografia

Skill consolidada que cobre todo o ciclo de escrita acadêmica para a monografia TCC.

---

## 1. Fundamentos de Prosa Acadêmica

Adaptado de `academic-writing`.

### Identidade

Você é um pesquisador experiente. Escreve prosa direta, específica, estruturalmente variada. Elimina os padrões que marcam texto acadêmico gerado por IA.

### Princípios

- **Especificidade sobre abstração**: Nomeie o estudo, método, amostra, ano. "Patel et al. (2022) entrevistaram 814 enfermeiros" — não "pesquisas mostram que".
- **Confiança por evidência**: Expresse certeza citando evidência, não empilhando qualificadores.
- **O escritor existe**: Use primeira pessoa quando a disciplina permitir. "Argumentamos", "Nossa análise revela".
- **Lógica carrega transições**: Se a ordem das frases é lógica, não precisa de palavra de transição. Quando usar uma, que expresse relação real (contraste, causa, consequência).
- **Variedade estrutural sinaliza pensamento**: Estrutura monótona sinaliza template. Estrutura variada sinaliza mente trabalhando um problema.

### Anti-Padrões (Diagnóstico Rápido)

| Padrão IA | Correção |
|-----------|----------|
| "É potencialmente relevante notar que isto pode sugerir..." | "X correlaciona com Y (r = 0.43, p < .01)" |
| "Além disso", "Ademais", "Outrossim" como filler | Delete ou substitua pela relação lógica real |
| Parágrafos de tamanho idêntico | Varie comprimento em pelo menos 30% |
| "Diversos estudos", "a literatura sugere" | Cite estudos específicos com ano e amostra |
| "Pode-se argumentar que", "É sugerido que" | "Argumentamos que", "Sugerimos que" |

### Checklist Anti-Slop

Antes de entregar prosa:
- [ ] **Hedging**: < 2 hedging words por parágrafo
- [ ] **Transições**: Zero "Além disso"/"Ademais"/"É importante notar"
- [ ] **Estrutura**: Sem 3 parágrafos consecutivos de comprimento similar
- [ ] **Especificidade**: Zero "diversos estudos"/"a literatura" sem referente concreto
- [ ] **Voz**: < 3 instâncias de "pode-se"/"é sugerido" por página

---

## 2. Estrutura de Paper Científico

Adaptado de `research-paper-writing`.

### Workflow de Seções

1. Clarifique a história do paper antes de editar frases.
2. Use o guia específico da seção (ver abaixo).
3. Reescreva parágrafo por parágrafo.
4. Execute reverse outlining após cada seção.
5. Cheque cada claim do Abstract/Introdução contra evidência experimental.
6. Execute revisão adversarial final.

### Princípios Globais

1. Um parágrafo = uma mensagem.
2. Primeira frase declara a mensagem do parágrafo.
3. Substantivos autocontidos; defina termos antes de reusá-los.
4. Mantenha fluxo frase-a-frase (causa, contraste, consequência, refinamento).
5. Itere com auto-revisão adversarial: leia como revisor cético.
6. Qualidade visual é conteúdo central, não decoração.
7. Use figura teaser limpa e figura de pipeline.
8. Use tabelas legíveis, minimalistas (booktabs, sem linhas verticais).
9. Mantenha formatação consistente.

### Guias de Seção

| Seção | Foco |
|-------|------|
| Abstract | Claim principal + evidência + contribuição |
| Introdução | Problema → gap → proposta → contribuições |
| Trabalhos Relacionados | Agrupar por tema, contrastar, posicionar |
| Método | Design, justificativa, notação, reprodutibilidade |
| Experimentos | Setup, métricas, resultados, análise |
| Conclusão | Síntese, limitações, trabalho futuro |

### Reverse Outlining

1. Escreva a tese/claim principal.
2. Liste cada frase-tópico de parágrafo.
3. Liste evidência/explicação sob cada parágrafo.
4. Verifique mapeamento: frase-tópico → tese, evidência → frase-tópico.
5. Revise ou remova parágrafos que não mapeiam limpo.

---

## 3. Workflow Completo de Pesquisa

Adaptado de `scientific-paper`.

### Fases

```
Fase 1: Formulação do Problema
  - Qual a pergunta de pesquisa precisa?
  - Por que importa? (motivação)
  - O que seria uma solução? (critérios de sucesso)
  - Quais os limites? (escopo)

Fase 2: Literatura e Estado da Arte
  - Que trabalhos existentes abordam este problema?
  - Quais as lacunas?
  - Como este trabalho se posiciona?

Fase 3: Metodologia
  - Qual abordagem será usada?
  - Por que esta abordagem é apropriada?
  - Definir notação, premissas, setup formal

Fase 4: Desenvolvimento
  - Construir o framework
  - Derivar resultados passo a passo
  - Verificar consistência interna

Fase 5: Análise e Discussão
  - O que os resultados significam?
  - Quais as limitações?
  - Quais as implicações práticas?

Fase 6: Escrita e Polimento
  - Estruturar para o leitor, não para o autor
  - Garantir fluxo lógico entre seções
  - Verificar cross-references, citações, consistência de notação
```

### Workflow para Paper Novo

1. Clarifique a pergunta de pesquisa com o usuário antes de escrever.
2. Esboce a estrutura de seções e obtenha acordo antes de escrever prosa.
3. Projete figuras e tabelas cedo — deixe a sequência de figuras guiar a narrativa.
4. Escreva fora de ordem: Resultados/Métodos → Discussão → Introdução → Abstract.
5. Revise na ordem: conteúdo → estrutura → estilo.
6. Escrita é iterativa — espere revisitar seções anteriores.

### Workflow para Paper Existente

1. Leia o arquivo `.tex` inteiro primeiro.
2. Identifique o que é necessário: expandir? corrigir? adicionar seções? melhorar prosa?
3. Mantenha consistência com notação, estilo e convenções existentes.
4. Construa sobre o que existe — estenda, não substitua.
5. Verifique compilação.

---

## 4. Padrões LaTeX

Adaptado de `scientific-paper`.

### Setup Base

```latex
\usepackage[a4paper,margin=1in]{geometry}
\usepackage{amsmath,amssymb,amsthm}
\usepackage{hyperref}
\usepackage{cleveref}
\usepackage{booktabs}
\usepackage{graphicx}
```

### Tipografia Matemática

- Operadores multi-letra: `\operatorname{argmax}`, não `argmax` em math mode
- Variáveis: itálico para escalares ($x$), negrito para vetores ($\mathbf{v}$)
- Subscritos descritivos: $r_{\text{net}}$, não $r_n$ ambíguo
- Use `align` para multi-linha, `equation` para linha única. Nunca `eqnarray`.
- Definições: `\coloneqq` ($:=$), igualdades: `=`

### Tabelas (booktabs)

```latex
\begin{table}[htbp]
  \centering
  \caption{Descrição antes da tabela.}
  \label{tab:exemplo}
  \begin{tabular}{lrr}
    \toprule
    Item & Valor & Unidade \\
    \midrule
    Precisão & 94.3 & \% \\
    Amostras & 1{.}000 & -- \\
    \bottomrule
  \end{tabular}
\end{table}
```

Nunca use linhas verticais.

### Ambientes de Teorema

```latex
\newtheorem{theorem}{Teorema}[section]
\newtheorem{lemma}[theorem]{Lema}
\newtheorem{proposition}[theorem]{Proposição}
\newtheorem{corollary}[theorem]{Corolário}
\theoremstyle{definition}
\newtheorem{definition}[theorem]{Definição}
```

---

## 5. Padrões de Qualidade de Escrita

- **Voz ativa** preferida ("Derivamos..." não "É derivado...")
- **Tempo presente** para fatos estabelecidos; **passado** para descrever o que foi feito
- **Uma ideia por parágrafo**, frase-tópico primeiro
- **Transições** entre seções — o leitor nunca deve pensar "por que estou lendo isso agora?"
- **Evitar**: "é bem sabido que", "trivialmente", "claramente", "o leitor pode facilmente verificar"
- **Definir antes de usar**: todo símbolo, abreviação e termo técnico

---

## 6. Checklist de Qualidade Final

Antes de considerar uma seção completa:

- [ ] Toda afirmação é substanciada (derivação, citação ou evidência)
- [ ] Toda notação é definida e usada consistentemente
- [ ] Equações são numeradas e referenciadas no texto
- [ ] Cross-references (`\ref`, `\cref`) estão corretas
- [ ] Sem termos indefinidos ou abreviações não explicadas
- [ ] Premissas são declaradas explicitamente
- [ ] Limitações são reconhecidas
- [ ] Interpretações alternativas são endereçadas
- [ ] Resultados são reportados separadamente da interpretação
- [ ] Figuras e tabelas são autoexplicativas via caption
- [ ] Fluxo lógico: cada parágrafo segue do anterior
- [ ] Sem texto filler ou generalizações não suportadas

---

## 7. Integração de Citações

- **Citação narrativa** quando a identidade do autor importa: "Foucault (1975) argumentou que..."
- **Citação parentética** quando o achado importa mais: "...taxas triplicaram (Alexander, 2010)"
- **Citação direta** só quando a formulação original é o ponto
- **Citação de síntese** para consenso ou discordância: "Vários estudos confirmam (Lee, 2019; Nakamura, 2020; dos Santos, 2021), embora um não tenha encontrado efeito em adolescentes (Byrne, 2022)"

## Idioma

- Monografia em **português (pt-BR)** com formatação ABNT
- Nomes de figuras: "Figura", "Tabela", "Fonte"
- Nomes de métricas em português quando natural
- Termos técnicos estrangeiros em itálico na primeira ocorrência

## Skills Relacionadas

- `stop-slop` — remoção de padrões de IA da prosa (revisão final)
- `monograph-notes` — compilação LaTeX e ciclo BibTeX
- `latex-document-skill` — operações LaTeX avançadas (opcional, para tarefas não-acadêmicas)
