```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "key_insight": "O ecossistema cobre bem pesquisa e incorporação de literatura, mas o elo mais crítico — experimentação e validação de claims contra dados — não tem comando, skill nem MCP dedicados, criando um vácuo de rastreabilidade entre o que a monografia afirma e o que os experimentos produziram.",
  "findings": [
    {
      "severity": "critical",
      "category": "completeness",
      "description": "Nenhum comando ou skill cobre a execução de experimentos (build → create → optimize → batch scripts). O fluxo documentado pula de INCORPORAÇÃO direto para ESCRITA, sem estágio de EXPERIMENTAÇÃO. O AGENTS-experiments.md descreve pipeline com 4 métodos × 51 seeds × 30 instâncias, mas o agente não tem como dispará-lo via opencode.",
      "fix": "Criar `/experimento` (ou `/rodar`) que aceite --method, --instance, --seed e encadeie make build → tcc optimize, com verificação de cobertura pós-execução. Adicionar skill `experiment-runner` que governe o pipeline batch (run_all.sh, run_experiments_multi_seed.sh) e valide integridade dos artefatos.",
      "why": "A experimentação é o núcleo empírico da monografia. Sem automação, o agente não pode gerar novos dados para sustentar claims, ficando limitado ao que já existe em src/data/results/.",
      "ref": "AGENTS-experiments.md:4-11"
    },
    {
      "severity": "critical",
      "category": "completeness",
      "description": "Nenhum comando para compilar a monografia (pdflatex → bibtex → pdflatex × 2). O skill latex-document-skill existe, mas não há `/compilar` ou equivalente que execute o ciclo completo com verificação de erros. O agente pode editar .tex mas não pode verificar se compila.",
      "fix": "Criar `/compilar` que execute o ciclo pdflatex → bibtex → pdflatex × 2 a partir de monografia/, capture erros do .log, e reporte warnings críticos (undefined references, missing citations, overfull boxes). Alternativamente, estender o latex-document-skill com um script de compilação específico para main_ppgco_ufu.tex.",
      "why": "Sem compilação automática, edições na monografia são cegas — o agente edita .tex mas não sabe se quebrou referências cruzadas, citações ou formatação.",
      "ref": "monografia/main_ppgco_ufu.tex"
    },
    {
      "severity": "critical",
      "category": "architecture",
      "description": "Nenhum mecanismo de rastreabilidade conecta claims da monografia ↔ resultados experimentais ↔ notas do vault ↔ entradas BibTeX. O vault tem diretório claims/ e Bases, mas não há skill nem comando que valide se cada afirmação quantitativa na monografia tem respaldo em src/data/results/.",
      "fix": "Criar skill `claim-tracking` com workflow: (1) extrair claims quantitativas dos .tex, (2) cruzar com summary/*.json, (3) verificar se valores citados batem com medianas/médias dos experimentos, (4) sinalizar claims sem evidência. Adicionar comando `/consistencia` que execute essa validação e produza relatório de divergências.",
      "why": "O council anterior (2026-06-02-validate-monograph) identificou problemas de consistência (ex: ACO 71.79% vs 60.35% dependendo de tolerância) que um sistema automático de rastreabilidade poderia ter detectado antes da revisão humana.",
      "ref": ".agents/council/2026-06-02-validate-monograph.md:32"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "O MCP codegraphcontext não está listado nos 9 MCPs do ecossistema. Ele poderia indexar o código Go (src/) e estabelecer relações como 'função EvaluateSolution é chamada por GA, PSO, ACO' ou 'Claim X no vault referencia implementação Y em src/optimizer/'. Sem ele, código e claims são domínios desconectados.",
      "fix": "Adicionar codegraphcontext aos MCPs. Indexar src/ com codegraphcontext_add_code_to_graph. Criar queries cypher que liguem símbolos Go a notas do vault/projeto/. Integrar ao `/consistencia` proposto acima: para cada claim sobre implementação, verificar se o símbolo correspondente existe no grafo de código.",
      "why": "A implementação Go é o artefato central que produz os resultados. Sem capacidade de analisar o código automaticamente, o agente não pode verificar se claims sobre algoritmos (ex: 'PSO usa random keys com ordem de menor valor') correspondem ao código real.",
      "ref": "src/"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "O comando `/baixar-pdf` está marcado como [DEPRECIADO] sem substituto documentado. O ecossistema tem 4 MCPs capazes de download (doiget, scihub, arxiv, pdf-reader) mas nenhum comando unificado os expõe. O fluxo `/incorporar` depende de ter o PDF primeiro, mas não há como obtê-lo via opencode.",
      "fix": "Criar `/baixar-pdf` v2 que use o cascade de MCPs (doiget → scihub → arxiv) com fallback automático. Ou, se a intenção é que o download seja sempre manual/scripts, documentar explicitamente no help do comando e remover do fluxo `/incorporar` a dependência implícita de PDF local.",
      "why": "Comando deprecado sem substituto quebra o fluxo PESQUISA → INCORPORAÇÃO. O agente que executa `/incorporar` precisa do PDF; se `/baixar-pdf` não funciona, o fluxo para no primeiro passo.",
      "ref": "AGENTS.md:comandos"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "Os MCPs scholar-sidekick (verifyCitation, checkRetraction) e crossref (get_references) têm ferramentas de validação de bibliografia que não são usadas por nenhum comando ou skill. O arquivo monografia/bib/abntex2-references.bib contém ~19+ entradas; nenhuma foi verificada automaticamente quanto a retratação, consistência de metadados ou integridade de referências cruzadas.",
      "fix": "Adicionar etapa de validação bibliográfica ao skill knowledge-base (fase 2.5: validar entrada BibTeX com scholar-sidekick.verifyCitation + checkRetraction antes de importar para o vault). Criar `/auditar-bib` que percorra todas as entradas do .bib e produza relatório de: retratações, DOIs quebrados, títulos inconsistentes, faltantes no vault.",
      "why": "Citações são a fundação acadêmica da monografia. Uma referência retratada ou com metadados incorretos compromete a credibilidade do trabalho. O ecossistema atual confia cegamente no .bib.",
      "ref": "monografia/bib/abntex2-references.bib"
    },
    {
      "severity": "significant",
      "category": "duplication",
      "description": "Três skills cobrem escrita acadêmica com sobreposição significativa: academic-writing, research-paper-writing e scientific-paper. Similarmente, 5 MCPs cobrem busca acadêmica: crossref, semantic-scholar, scholar-sidekick, google-scholar, academic-search. A redundância não é documentada — não há guia de quando usar cada um.",
      "fix": "Documentar matriz de decisão: (a) skills de escrita — academic-writing para prosa geral em pt-BR, research-paper-writing para estrutura ML/CV/NLP (menos relevante para TCC de otimização), scientific-paper para ciclo completo LaTeX; (b) MCPs de busca — academic-search como fachada unificada (search_papers), scholar-sidekick para validação/citação, crossref para metadados estruturados, semantic-scholar para grafos de citação, google-scholar para descoberta ampla. Adicionar esta matriz ao AGENTS.md.",
      "why": "Sem documentação de escolha, o agente pode usar o skill/MCP errado para a tarefa (ex: usar research-paper-writing — otimizado para papers de ML — para editar monografia de TSP em português).",
      "ref": "AGENTS.md:skills"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "Nenhum comando para gerar figuras. scripts/gerar-figuras.sh e scripts/gerar-analises.sh existem mas não são expostos como comandos opencode. O skill scientific-figures cobre design de figuras, mas não a pipeline de geração (iniciar servidor → chamar /api/render → salvar PNG/SVG → posicionar no LaTeX).",
      "fix": "Criar `/figuras` que aceite --figura (nome) ou --todas e execute o pipeline: garante que tcc serve está rodando, chama /api/render com parâmetros corretos, salva em monografia/figs/, verifica escala/rótulos. Alternativamente, wrapper em torno de scripts/gerar-figuras.sh com validação pós-geração.",
      "why": "Figuras são o elo visual entre dados e texto. O fluxo lista FIGURAS como etapa, mas sem comando, o agente precisa de conhecimento implícito sobre scripts/ e da API do servidor para gerar cada figura.",
      "ref": "scripts/gerar-figuras.sh"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "Não há skill de domínio específico para TSP/rTSP ou otimização bioinspirada. Skills como go e golang-cli cobrem a implementação, mas nenhum skill captura conhecimento de domínio: formulação TSP-SD-ATP, propriedades das meta-heurísticas (GA, PSO, ACO), trade-offs conhecidos, pitfalls comuns (ex: encoding confound identificado pelo council anterior).",
      "fix": "Criar skill `tsp-optimization` que documente: (1) definição formal do TSP-SD-ATP, (2) propriedades de cada meta-heurística no contexto do problema, (3) armadilhas conhecidas (encoding confound, iteração vs wall-clock, tolerância de ACO), (4) convenções de nomenclatura (makespan, gap%, hit rate). Este skill seria carregado automaticamente ao editar cap_fundamentacao, cap_proposta ou cap_experimentos.",
      "why": "O council de validação da monografia identificou 4 problemas que um skill de domínio teria prevenido (encoding confound, tolerância ACO, equidade de comparação, lower bound). Conhecimento de domínio hoje está implícito no AGENTS.md e disperso em notas do vault.",
      "ref": ".agents/council/2026-06-02-validate-monograph.md:30-38"
    },
    {
      "severity": "minor",
      "category": "integration",
      "description": "O arxiv MCP tem watch_topic + check_alerts para vigilância contínua de literatura, mas nenhum tópico está configurado. O ecossistema faz pesquisa reativa (quando o usuário pede) mas não proativa (novos papers sobre TSP + drones + meta-heurísticas aparecem e o sistema alerta).",
      "fix": "Registrar watched topics via arxiv MCP: \"patrol drone\" AND (TSP OR \"traveling salesman\"), \"bio-inspired\" AND (TSP OR routing), \"multi-rotor\" AND \"path planning\" AND optimization. Adicionar verificação semanal ao fluxo de trabalho (talvez como etapa do `/roadmap-criar`).",
      "why": "A monografia está em andamento; papers publicados durante a escrita deveriam ser considerados. Sem vigilância, o referencial teórico congela no momento da busca inicial.",
      "ref": "AGENTS.md:arxiv MCP"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "O diretório vault/bases/ contém views do Obsidian Bases, skill vault-bases-maintainer existe, mas não há comando para consultar as Bases programaticamente (ex: 'quantos papers cobrem ACO?', 'quais claims não têm evidência experimental?'). O agente depende de busca textual simples.",
      "fix": "Criar `/consultar-base` que aceite --base (nome do .base) e --filtro (expressão) e retorne resultados formatados. Isso exporia as consultas estruturadas do Obsidian Bases como ferramenta de trabalho, não apenas como visualização passiva.",
      "why": "As Bases foram criadas para consultas estruturadas, mas sem comando, o agente não consegue extrair seu valor durante a escrita — precisa abrir o Obsidian manualmente.",
      "ref": "vault/bases/"
    },
    {
      "severity": "minor",
      "category": "architecture",
      "description": "O fluxo documentado omite COMPILAÇÃO e EXPERIMENTAÇÃO. O fluxo completo deveria ser: PESQUISA → INCORPORAÇÃO → ENRIQUECIMENTO → EXPERIMENTAÇÃO → ANÁLISE → ESCRITA → FIGURAS → COMPILAÇÃO → REVISÃO → TAREFAS → VALIDAÇÃO → CÓDIGO. Além disso, ANÁLISE aparece depois de FIGURAS no fluxo atual, mas a análise (notebook, estatística) normalmente precede a geração de figuras.",
      "fix": "Reordenar fluxo no AGENTS.md para refletir a dependência real: experimentos produzem dados → análise interpreta dados → figuras visualizam análises → escrita incorpora figuras → compilação produz PDF. Adicionar etapas faltantes (EXPERIMENTAÇÃO, COMPILAÇÃO).",
      "why": "Fluxo incorreto induz o agente a pular etapas ou executá-las fora de ordem. Se FIGURAS vem antes de ANÁLISE, o agente pode tentar gerar figuras antes de ter análises estatísticas que as informem.",
      "ref": "AGENTS.md:fluxo"
    }
  ],
  "recommendation": "Prioridade 1 (critica): criar `/experimento`, `/compilar` e `/consistencia` com skills correspondentes — são os três elos faltantes que impedem o agente de fechar o ciclo completa da monografia (experimentar → validar claims → compilar PDF). Prioridade 2 (significativa): integrar codegraphcontext ao ecossistema, criar `/baixar-pdf` v2, adicionar validação bibliográfica automatizada. Prioridade 3 (menor): skill de domínio TSP, vigilância de literatura com arxiv watch, `/consultar-base`, correção do fluxo documentado.",
  "schema_version": 3
}
```

---

## Análise Estendida — Judge 3 (Missing Capabilities & Logical Additions)

### Perspectiva do Juiz

Avalio o ecossistema pela lente de **completude funcional**: o que um agente precisaria fazer para levar uma monografia do zero à defesa e quais dessas capacidades estão ausentes. O ecossistema é forte em pesquisa e incorporação de literatura (4+ MCPs de busca, comando `/incorporar` com 9 fases, vault estruturado), mas fraco no que acontece **depois** que o conhecimento está no vault — transformar esse conhecimento em experimentos, análises, texto compilável e claims validadas.

### Os Três Vazios Críticos

**1. Vácuo de experimentação.** O AGENTS-experiments.md descreve um pipeline sofisticado (build → create → optimize com 4 métodos, 51 seeds, 30 instâncias), mas zero comandos opencode o expõem. Um agente que precisa rodar um novo experimento (ex: testar parâmetro diferente do ACO) teria que executar bash manualmente, sem validação de cobertura, sem verificação de integridade dos artefatos. Isso é particularmente grave porque o council anterior (2026-06-02) identificou que claims sobre ACO dependem de tolerância não documentada — se o agente pudesse re-rodar experimentos variando a tolerância, essa fragilidade teria sido detectada.

**2. Vácuo de rastreabilidade.** O vault tem estrutura sofisticada (papers/, claims/, bases/, canvas/) e a monografia tem claims quantitativas (ex: "ACO atinge 71.79% de taxa de acerto"), mas não há nenhum mecanismo automatizado que verifique se uma claim no .tex corresponde a um valor real em src/data/results/summary/*.json. O agente edita claims sem saber se elas são verdadeiras. Isso é um risco de integridade acadêmica.

**3. Vácuo de compilação.** O agente pode editar .tex (via scientific-paper skill) mas não pode verificar se o que editou compila. Sem `/compilar`, cada edição na monografia é um tiro no escuro — referências cruzadas quebradas, citações faltantes e overfull boxes só seriam descobertos por um humano rodando pdflatex manualmente.

### Oportunidades de Integração Não Exploradas

- **codegraphcontext**: Não listado. Indexar src/ permitiria queries como "quais funções implementam a penalidade angular do rTSP?" ou "a claim X na monografia referencia o código Y que não existe mais?".
- **scholar-sidekick.verifyCitation**: Não usado. Poderia validar cada entrada do .bib automaticamente.
- **arxiv.watch_topic**: Não configurado. Perde-se a oportunidade de vigilância contínua de literatura durante os meses de escrita.
- **doiget.batch_from_bibliography**: Não usado. Poderia baixar automaticamente todos os PDFs referenciados no .bib.

### Sobreposições que Precisam de Documentação

A matriz de decisão para skills de escrita e MCPs de busca é essencial. Sem ela, o agente pode:
- Usar `research-paper-writing` (otimizado para papers de ML/CV/NLP em inglês) para editar monografia de TSP em português ABNT
- Usar `google-scholar` quando `academic-search` (que agrega Crossref + Semantic Scholar) seria mais eficiente
- Usar `crossref` para busca quando `semantic-scholar` tem melhor cobertura de citações

### Fluxo Corrigido Proposto

```
PESQUISA → INCORPORAÇÃO → ENRIQUECIMENTO → EXPERIMENTAÇÃO → ANÁLISE → ESCRITA → FIGURAS → COMPILAÇÃO → CONSISTÊNCIA → REVISÃO → TAREFAS → VALIDAÇÃO → CÓDIGO
```

Mudanças em relação ao fluxo atual:
1. **+EXPERIMENTAÇÃO** após ENRIQUECIMENTO (antes de escrever, rode experimentos)
2. **+COMPILAÇÃO** após FIGURAS (produza PDF para inspeção)
3. **+CONSISTÊNCIA** após COMPILAÇÃO (valide claims contra dados experimentais)
4. **Reordenação**: ANÁLISE antes de ESCRITA e FIGURAS (não faz sentido gerar figuras antes de analisar dados)
5. **Reordenação**: CÓDIGO no final (alterações de código são consequência de validação, não etapa independente)
