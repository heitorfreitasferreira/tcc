# Relatório do Juiz 1 — Validação do Ecossistema opencode do TCC

```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "key_insight": "O ecossistema cobre com solidez o pipeline literatura→incorporação→escrita, mas tem três gaps arquiteturais críticos (orquestração de experimentos, council skill órfão no inventário, e ausência de loop fechado entre /claudiney e /roadmap) que quebram a rastreabilidade ponta-a-ponta do fluxo de trabalho da monografia.",
  "findings": [
    {
      "severity": "critical",
      "category": "architecture",
      "description": "Sem comando para orquestrar experimentos. Os scripts src/run_all.sh e src/run_experiments_multi_seed.sh só são acessíveis via bash cru. Nenhum comando opencode encapsula a execução de experimentos, quebrando o fluxo experimentos → dados → análise → figuras.",
      "fix": "Criar comando /experimentar que aceite parâmetros (instância, método, seeds, iterações, população) e delegue aos scripts existentes, registrando o run_id no roadmap ou em vault/writing/.",
      "why": "O pipeline da monografia depende de resultados experimentais (src/data/results/), mas a geração desses resultados está fora do ecossistema opencode. Um agente que siga o workflow descrito não consegue executar o passo 9 (ANÁLISE) sem intervenção manual externa.",
      "ref": "Workflow steps 9-10; src/run_all.sh; src/run_experiments_multi_seed.sh"
    },
    {
      "severity": "critical",
      "category": "completeness",
      "description": "Council skill é referenciado no workflow (passo 7: 'council skill → valida mudanças com multi-juízes') e em /roadmap-consumir ('validar (council)'), mas NÃO aparece no inventário de Skills do ecossistema. O inventário lista 14 skills; council é a 15ª ausente.",
      "fix": "Adicionar council ao inventário de skills no ecossistema com descrição de quando invocar e como se integra com /roadmap-consumir.",
      "why": "A skill existe no AGENTS.md (available_skills) mas não está documentada no ecossistema. Isso cria inconsistência: o workflow e comandos pressupõem sua existência, mas a documentação canônica do ecossistema a omite.",
      "ref": "ECOSSISTEMA COMPLETO > Skills section (14 listed, council missing); Workflow step 7"
    },
    {
      "severity": "critical",
      "category": "integration",
      "description": "/claudiney registra feedback do orientador em vault/writing/review-solicitacoes/ mas não gera itens no roadmap automaticamente. O feedback do orientador é um trigger natural para novas tarefas, mas o loop é quebrado: após /claudiney, o usuário precisa manualmente executar /roadmap-criar para cada ação.",
      "fix": "Estender /claudiney com flag --criar-tarefas que, ao final da análise, sugira e opcionalmente crie itens de roadmap derivados do feedback (ex.: 'corrigir tabela X', 'adicionar citação Y'). Alternativamente, criar comando /feedback-acao que leia o último review e popule o roadmap.",
      "why": "Orientador aponta problemas → não há tarefa criada → problemas persistem. Sem rastreabilidade do feedback para ação concreta, o ciclo de revisão fica dependente de memória humana.",
      "ref": "Comando /claudiney; Comando /roadmap-criar"
    },
    {
      "severity": "significant",
      "category": "duplication",
      "description": "Quatro skills com sobreposição massiva em produção de prosa acadêmica: academic-writing, research-paper-writing, scientific-paper, e latex-document-skill. Nenhuma árvore de decisão documentada diz quando usar qual.",
      "fix": "Definir matriz de decisão: academic-writing → rascunho e revisão de seções; research-paper-writing → polimento pré-submissão (ML/CV/NLP); scientific-paper → estruturação completa de paper; latex-document-skill → compilação e formatação LaTeX. Documentar no AGENTS.md ou em .agents/skills/README.md.",
      "why": "Agentes diferentes (ou o mesmo agente em sessões diferentes) podem carregar skills diferentes para a mesma tarefa, resultando em convenções inconsistentes de formatação, estilo e estrutura no .tex.",
      "ref": "Skills: academic-writing, research-paper-writing, scientific-paper, latex-document-skill"
    },
    {
      "severity": "significant",
      "category": "duplication",
      "description": "data-analysis e statistical-analysis cobrem o mesmo domínio (testes de hipótese, ANOVA, regressão) com fronteiras difusas. data-analysis promete '4-round review'; statistical-analysis promete 'APA reporting'. Na prática, um agente pode escolher qualquer uma sem critério.",
      "fix": "Unificar em uma skill com modos (--mode exploratory | confirmatory | apa-report) ou definir boundary clara: data-analysis → análise exploratória + visualização; statistical-analysis → testes confirmatórios + reporte formal.",
      "why": "Análises estatísticas inconsistentes (testes diferentes, thresholds diferentes) em seções distintas da monografia minam a credibilidade dos resultados.",
      "ref": "Skills: data-analysis, statistical-analysis"
    },
    {
      "severity": "significant",
      "category": "duplication",
      "description": "Cinco MCPs para busca de literatura (crossref, arxiv, google-scholar, semantic-scholar, academic-search) com academic-search já encapsulando crossref + semantic-scholar. /consultar-academico dispara os 4 em paralelo — cobertura máxima, mas 3x redundância de chamadas para o mesmo paper.",
      "fix": "Tornar /consultar-academico configurável (--sources crossref,arxiv) com padrão academic-search (que já consolida). Usar os MCPs individuais apenas para fallback ou busca especializada (ex.: arxiv para preprints recentes, google-scholar para cobertura cinzenta).",
      "why": "Quatro chamadas paralelas para cada consulta consomem rate limits, aumentam latência, e produzem resultados duplicados que exigem deduplicação manual. O custo operacional não se justifica para consultas de rotina.",
      "ref": "MCPs: crossref, arxiv, google-scholar, semantic-scholar, academic-search; Comando /consultar-academico"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "Plugin pdf-watcher.ts detecta PDFs novos mas não dispara ação. É um monitor passivo: loga e notifica, mas não sugere /incorporar nem popula fila de processamento. O valor agregado é zero além do que um ls ou inotify já fariam.",
      "fix": "Estender pdf-watcher para, ao detectar novo PDF, escrever o caminho em vault/.incorporar-fila (ou similar) para que /incorporar fila processe automaticamente. Alternativamente, integrar com o comando /incorporar para modo --watch.",
      "why": "Um plugin que apenas observa sem agir é peso morto na arquitetura. A detecção de PDFs é o trigger natural para o pipeline de incorporação, mas a conexão não está feita.",
      "ref": "Plugin pdf-watcher.ts; Comando /incorporar (modo 'fila')"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "Nenhum comando opencode encapsula geração de figuras. O workflow cita 'scientific-figures skill → gera figuras', mas scripts/gerar-figuras.sh e scripts/gerar-analises.sh são os artefatos canônicos e só rodam via bash.",
      "fix": "Criar comando /figuras que invoque scripts/gerar-figuras.sh e scripts/gerar-analises.sh com parâmetros (--tipo mapa|evolucao|cd|tudo, --instancia, --metodo).",
      "why": "O AGENTS.md estabelece o servidor web como fonte canônica de figuras, mas não há comando para acioná-lo. Um agente precisa saber dos scripts bash, quebrando a abstração do ecossistema opencode.",
      "ref": "Scripts: scripts/gerar-figuras.sh, scripts/gerar-analises.sh; Skill: scientific-figures"
    },
    {
      "severity": "minor",
      "category": "correctness",
      "description": "/baixar-pdf está marcado DEPRECIADO e redireciona para /incorporar, mas permanece listado no inventário de comandos. Sua presença gera confusão sobre qual comando usar para download de PDFs.",
      "fix": "Remover /baixar-pdf do inventário ativo ou movê-lo para seção 'Comandos depreciados'. Se ainda houver scripts que o referenciem, atualizar as referências.",
      "why": "Comandos depreciados no inventário ativo poluem o espaço de decisão do agente e podem ser invocados por engano.",
      "ref": "Comando /baixar-pdf"
    },
    {
      "severity": "minor",
      "category": "duplication",
      "description": "vault-semantic-schema e vault-tagger têm sobreposição de escopo em 'tags hierárquicas'. vault-semantic-schema cobre 'tags hierárquicas, frontmatter, claims'; vault-tagger cobre 'padronização de tags hierárquicas'. A padronização de tags é subconjunto do schema semântico.",
      "fix": "Consolidar vault-tagger como submódulo de vault-semantic-schema ou esclarecer que vault-semantic-schema define o schema e vault-tagger audita/enforce conformidade.",
      "why": "Duas skills que mencionam 'tags hierárquicas' criam ambiguidade sobre qual é autoridade canônica para o formato das tags.",
      "ref": "Skills: vault-semantic-schema, vault-tagger"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "Nenhum MCP ou comando expõe o endpoint /api/render do servidor Go, que é a fonte canônica de figuras segundo AGENTS.md. O servidor precisa ser iniciado manualmente (./src/tcc serve) fora do ecossistema opencode.",
      "fix": "Criar MCP render-tcc ou estender o comando /figuras (sugerido acima) para gerenciar o ciclo de vida do servidor (start → render → stop).",
      "why": "A arquitetura declara o servidor como fonte canônica, mas não oferece integração programática com ele no ecossistema opencode.",
      "ref": "AGENTS.md (web server section); src/web/handlers/render.go"
    }
  ],
  "recommendation": "Prioridade 1 (bloqueante): criar /experimentar e /figuras — sem esses comandos, o ecossistema não cobre o ciclo completo da monografia (experimentos → dados → análise → figuras). Prioridade 2 (alta): fechar o loop /claudiney → /roadmap e incluir council no inventário. Prioridade 3 (média): consolidar skills redundantes (writing, stats) e configurar /consultar-academico com fontes seletivas. Prioridade 4 (baixa): remover /baixar-pdf depreciado, dar ação ao pdf-watcher, e esclarecer boundary vault-semantic-schema vs vault-tagger.",
  "schema_version": 3
}
```

## Escopo avaliado

Ecossistema opencode completo do TCC conforme documentado no pacote: 10 comandos, 14 skills, 9 MCPs, 1 plugin, fluxo de trabalho em 10 passos, schema do vault. A avaliação concentra-se na arquitetura de integração entre camadas e na cobertura do fluxo ponta-a-ponta da monografia.

## Arquitetura geral

O ecossistema segue uma estratificação limpa em 4 camadas:

```
COMANDOS     (orquestração — entry points acionáveis pelo usuário)
SKILLS       (instrução — pacotes de conhecimento domínio-específico)
MCPs         (integração — ferramentas externas: busca, PDF, citações)
PLUGINS      (automação passiva — monitoramento de filesystem)
```

Esta estratificação é arquiteturalmente sólida. As dependências fluem na direção correta: comandos invocam skills e MCPs; skills referenciam MCPs quando necessário; plugins operam no nível de filesystem sem acoplamento com comandos.

O fluxo de trabalho em 10 passos cobre o pipeline acadêmico completo: PESQUISA → INCORPORAÇÃO → ENRIQUECIMENTO → ESCRITA → REVISÃO → TAREFAS → VALIDAÇÃO → FIGURAS → ANÁLISE → CÓDIGO. Cada passo mapeia para skills e comandos específicos, o que é correto.

## O que funciona bem

1. **Pipeline de incorporação**: `/incorporar` com 9 fases é o comando mais maduro do ecossistema. A cadeia PDF → resumo → vault → canvas → claims é bem definida e cobre o ciclo completo de ingestão de literatura.

2. **Camadas Go bem estratificadas**: `go` (fundamentos) → `golang-cli` (CLI patterns) → `golang-spf13-cobra` (framework específico) é uma hierarquia de skills exemplar, sem sobreposição.

3. **Cobertura de busca literária**: Com crossref (155M works), arxiv (preprints), google-scholar (literatura cinzenta) e semantic-scholar (citações), a cobertura é abrangente. O comando `/consultar-academico` consolida os resultados.

4. **Sistema de roadmap**: `/roadmap-criar` + `/roadmap-consumir` com IDs incrementais, fases de clarificar → executar → validar → concluir, e sugestão de commit ao final. Estrutura de task management bem pensada.

5. **Schema do vault**: Estrutura clara com papers, áreas, claims, canvas, templates, bases, e writing. A separação entre conteúdo (papers), estrutura (areas, claims) e meta (bases, templates) é correta.

## Gaps arquiteturais (detalhamento)

### Gap 1: Orquestração de experimentos (CRITICAL)

O fluxo de trabalho declara que ANÁLISE (passo 9) usa `data-analysis + statistical-analysis` e CÓDIGO (passo 10) usa `go + golang-cli + golang-spf13-cobra`. Mas entre a ESCRITA (passo 4) e a ANÁLISE (passo 9), os experimentos precisam ser executados para gerar `src/data/results/`. Esse passo está implícito mas não tem representação no ecossistema.

Os scripts `src/run_all.sh` e `src/run_experiments_multi_seed.sh` são os mecanismos canônicos, mas residem fora do ecossistema opencode — um agente não os descobre naturalmente; precisa de conhecimento prévio ou leitura do AGENTS.md.

**Consequência**: o pipeline é quebrado. Da perspectiva do ecossistema opencode, não há como gerar os dados que a análise consome.

### Gap 2: Council skill ausente do inventário (CRITICAL)

O workflow passo 7 e o comando `/roadmap-consumir` referenciam `council skill` como mecanismo de validação. O AGENTS.md lista `council` como available skill. Mas o inventário do ecossistema lista 14 skills e `council` não está entre elas.

Isso cria uma inconsistência de documentação: o ecossistema descreve um fluxo que depende de um componente não declarado. Se o inventário é a fonte canônica, `council` não existe; se o AGENTS.md é a fonte canônica, o inventário está incompleto.

### Gap 3: Loop /claudiney → /roadmap quebrado (CRITICAL)

`/claudiney` analisa o contexto (src/, data/, vault/, monografia/), preenche análise técnica e plano de correção, e escreve em `vault/writing/review-solicitacoes/`. Mas o plano de correção gerado não alimenta o roadmap.

O ciclo natural de revisão acadêmica é: orientador aponta problemas → tarefas são criadas → tarefas são executadas → nova versão é gerada. O ecossistema tem `/claudiney` para o primeiro passo e `/roadmap-*` para os dois seguintes, mas a conexão entre eles não é automatizada.

## Redundâncias (detalhamento)

### Skills de escrita (4 com sobreposição)

| Skill | Propósito declarado | Quando usar? |
|-------|-------------------|-------------|
| academic-writing | Prosa acadêmica, revisão de literatura, metodologia | ? |
| research-paper-writing | Revisão de paper ML/CV/NLP | ? |
| scientific-paper | Papers acadêmicos em LaTeX | ? |
| latex-document-skill | Documentos LaTeX completos (compilar, converter) | ? |

As 3 primeiras produzem texto acadêmico; não há critério documentado de seleção. A 4ª (latex-document-skill) é mais sobre compilação/formatação, mas também "writes and edits scientific papers in LaTeX", colidindo com scientific-paper.

### Skills de estatística (2 com sobreposição)

Ambas cobrem ANOVA, regressão, testes de hipótese. A distinção prometida (data-analysis = 4-round review; statistical-analysis = APA reporting) não é suficiente para um agente decidir qual carregar para uma tarefa concreta como "analisar os resultados do ACO vs GA na instância 10a".

### MCPs de busca (5, com 1 wrapper)

`academic-search` já encapsula crossref + semantic-scholar. Disparar os 4 individuais em paralelo no `/consultar-academico` produz resultados duplicados que precisam ser deduplicados. Para consultas de rotina, academic-search sozinho seria suficiente.

## Síntese dos modos de falha

O ecossistema é **forte no meio** (incorporação, vault, escrita) mas **fraco nas pontas**:
- **Ponta esquerda** (experimentos): sem comandos, dependente de bash manual
- **Ponta direita** (validação e figuras): council não documentado, figuras sem comando
- **Conexões internas**: /claudiney não alimenta /roadmap; pdf-watcher não alimenta /incorporar

O veredito é WARN, não FAIL, porque:
- O núcleo do ecossistema (literatura → vault → escrita) é funcional e bem estruturado
- Os gaps identificados são de integração e completude, não de correção
- Nenhum componente está quebrado; as ausências são de conectores e wrappers

## Recomendação detalhada

### Imediato (antes de qualquer novo ciclo de escrita)

1. Criar `/experimentar` — wrapper para `src/run_all.sh` e `src/run_experiments_multi_seed.sh`
2. Criar `/figuras` — wrapper para `scripts/gerar-figuras.sh` e `scripts/gerar-analises.sh`
3. Adicionar `council` ao inventário de skills do ecossistema

### Curto prazo (próximo sprint)

4. Estender `/claudiney` com `--criar-tarefas` para fechar o loop com `/roadmap`
5. Documentar matriz de decisão para skills de escrita e estatística no AGENTS.md
6. Tornar fontes do `/consultar-academico` configuráveis com default `academic-search`

### Longo prazo (refinamento)

7. Dar ação ao `pdf-watcher` (escrever em fila para `/incorporar fila`)
8. Consolidar ou esclarecer boundary `vault-semantic-schema` vs `vault-tagger`
9. Remover `/baixar-pdf` do inventário ativo
10. Avaliar criação de MCP `render-tcc` para o endpoint `/api/render`
