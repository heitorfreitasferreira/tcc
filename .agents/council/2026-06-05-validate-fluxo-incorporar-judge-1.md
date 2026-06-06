```json
{
  "verdict": "FAIL",
  "confidence": "HIGH",
  "key_insight": "O fluxo tem 3 falhas críticas de integridade de dados (bibtex_key duplicado no template, output arxiv: prefixado tratado como DOI, dependência circular com doiget_bibtex_export) que causariam corrupção silenciosa do vault ou falha completa da pipeline.",
  "findings": [
    {
      "severity": "critical",
      "category": "correctness",
      "description": "O template `vault/templates/paper-note.md` define ambas as propriedades `bibtex_key` e `bibtex-key` no frontmatter. O vault-semantic-schema (linha 268) explicitamente proíbe trocar `bibtex-key` por `bibtex_key`. O `/incorporar` Fase 5 diz preencher `bibtex_key`, mas o script `import-bib-to-vault.sh` e todas as 91 notas existentes usam `bibtex-key`. Isso garante que notas criadas via `/incorporar` terão valor nulo em `bibtex_key` e a propriedade `bibtex-key` ficará vazia — quebrando scripts e Bases que dependem de `bibtex-key`.",
      "fix": "Remover `bibtex_key` do template. Manter apenas `bibtex-key` e corrigir Fase 5 para usar `bibtex-key`.",
      "why": "O template foi duplicado durante alguma edição mas a divergência nunca foi reconciliada com o schema canônico.",
      "ref": "vault/templates/paper-note.md:6-7 e .opencode/command/incorporar.md:66"
    },
    {
      "severity": "critical",
      "category": "correctness",
      "description": "`scripts/extract-pdf-doi.sh` retorna `arxiv:2301.08745` como saída quando encontra arXiv ID (linha 42). A Fase 2 do `/incorporar` passa esse valor diretamente para `crossref_get_work(doi)`, `doiget_resolve_paper(doi)`, e `scholar-sidekick_checkOpenAccess(doi)`. Nenhuma dessas funções aceita o prefixo `arxiv:` — Crossref rejeita, doiget pode interpretar como inválido. O fluxo não contém nenhuma lógica de branch para tratar `arxiv:` vs `10.` prefixos.",
      "fix": "Adicionar um bloco de decisão na Fase 2: se a saída começa com `arxiv:`, usar ferramentas arXiv (arxiv_get_abstract, arxiv_download_paper) em vez de Crossref. Separar a lógica de branch explicitamente.",
      "why": "O script extract foi projetado para extração best-effort, mas a pipeline downstream não foi adaptada para consumir suas saídas heterogêneas.",
      "ref": "scripts/extract-pdf-doi.sh:40-44 e .opencode/command/incorporar.md:34-43"
    },
    {
      "severity": "critical",
      "category": "integration",
      "description": "A Fase 2 depende de `doiget_bibtex_export(doi)` para obter BibTeX, mas esta ferramenta só funciona para entradas que já existem no store local do doiget. Se o PDF foi obtido via Sci-Hub (Fase 1, fallback), ou se é um PDF local sem DOI, o doiget nunca terá essa entrada no store. A pipeline não prevê geração de BibTeX a partir dos metadados do Crossref nem via `crossref_get_work` que já foi chamado antes.",
      "fix": "Usar `scholar-sidekick_exportCitation(doi, format='bib')` como fonte primária de BibTeX, que é stateless. Ou gerar BibTeX manualmente a partir dos metadados do crossref_get_work. Mover `doiget_bibtex_export` para fallback opcional.",
      "why": "A ordem das fases implica que doiget_fetch_paper pode falhar e cair para scihub, mas a dependência circular em doiget_bibtex_export não é resolvida.",
      "ref": ".opencode/command/incorporar.md:43, Fases 1-2"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "O comando `/incorporar` Fase 5 cria notas diretamente preenchendo o template, enquanto o workflow do skill `knowledge-base` (passos 5+7) cria notas via `import-bib-to-vault.sh` e depois as enriquece. Esses dois caminhos produzem frontmatter com formatos diferentes: o script infere tags via regex no abstract (linhas 36-49 do script) e gera `tipo/paper`, `status/pendente`, `evidencia/referencia`, enquanto `/incorporar` infere tags manualmente com regras potencialmente diferentes. Notas criadas por caminhos distintos terão schemas inconsistentes.",
      "fix": "Documentar que `/incorporar` é o caminho canônico e que `import-bib-to-vault.sh` é apenas para migração em lote de entradas já existentes no .bib. Ou unificar a lógica de inferência de tags em uma fonte única.",
      "why": "O skill knowledge-base e o comando /incorporar evoluíram independentemente sem reconciliação de responsabilidades.",
      "ref": ".agents/skills/knowledge-base/SKILL.md:98-123 e .opencode/command/incorporar.md:65-77"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "A Fase 7 (Verificar claims) diz 'Para cada claim existente, verifica se o paper apoia/contradiz' mas não especifica nenhum método para realizar essa verificação. Com 60 claims em `vault/claims/`, uma leitura completa tomaria minutos e exigiria raciocínio semântico complexo. O fluxo essencialmente delega uma tarefa impossível ao agente sem heurísticas, thresholds, ou escopo reduzido.",
      "fix": "Adicionar heurísticas práticas: (1) matching por palavras-chave no título do claim vs abstract do paper; (2) limitar verificação a claims do mesmo `claim_type` que os `methods` do paper; (3) usar similaridade de cosseno entre embeddings do claim e do abstract como pré-filtro; (4) explicitar que claims com `strength: forte` e `status: validado` exigem verificação manual e não devem ser alterados automaticamente.",
      "why": "O design assume capacidade ilimitada de leitura e matching semântico que não existe nas ferramentas disponíveis.",
      "ref": ".opencode/command/incorporar.md:88-97"
    },
    {
      "severity": "significant",
      "category": "correctness",
      "description": "Fase 5 especifica `![[pdfs/<key>.pdf]]` como embed do PDF. Mas as notas residem em `vault/papers/` e os PDFs em `vault/papers/pdfs/`. O caminho `pdfs/<key>.pdf` a partir de `vault/papers/` resolveria para `vault/papers/pdfs/<key>.pdf`, que está correto. Porém o template `vault/templates/paper-note.md` não contém esse embed — ele só tem `<!-- PDF não disponível -->`. O `/incorporar` precisaria adicionar a linha de embed explicitamente, o que não está documentado como será feito (insert vs replace do placeholder).",
      "fix": "Documentar que o embed substitui o comentário `<!-- PDF não disponível -->` e usar caminho relativo `papers/pdfs/<key>.pdf` (a partir da raiz do vault). Alternativamente, usar `[[papers/pdfs/<key>.pdf]]` que é inequívoco.",
      "why": "A sintaxe de embed do Obsidian é sensível ao caminho e o placeholder no template é um comentário HTML, não um embed.",
      "ref": ".opencode/command/incorporar.md:68 e vault/templates/paper-note.md:28-29"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "O arquivo `.opencode/log/pdf-events.log` não existe (diretório vazio). O plugin `pdf-watcher.ts` (linha 49-50) escreve nele quando um PDF é detectado, mas se o watcher nunca disparou, `/incorporar fila` terá uma lista vazia ou erro de arquivo não encontrado. O fluxo não documenta o comportamento esperado quando o log está vazio ou ausente.",
      "fix": "Adicionar tratamento na Fase 1 para log ausente: exibir 'Nenhum PDF pendente na fila' e sugerir `/incorporar vault` como alternativa. O plugin também deve criar o arquivo de log na inicialização.",
      "why": "O modo `fila` depende de um side effect assíncrono (watcher) que pode nunca ter ocorrido no momento do primeiro uso.",
      "ref": ".opencode/command/incorporar.md:20, .opencode/plugins/pdf-watcher.ts:49-50"
    },
    {
      "severity": "significant",
      "category": "architecture",
      "description": "O `/incorporar` Fase 7 diz 'usa `obsidian-bases` skill se disponível' para ler claims. Mas `claims.base` é um arquivo de configuração de view do Obsidian Bases, não um banco de dados consultável por MCP. O conteúdo real dos claims está em 60 arquivos individuais `vault/claims/*.md`. A Fase 7 parece confundir o sumário (claims.base) com a fonte dos dados.",
      "fix": "Especificar que a leitura deve ser dos arquivos `vault/claims/*.md`, com `claims.base` usado apenas para referência de schema. Se houver muitos claims (>20), usar apenas claims com `status: validado` como escopo inicial.",
      "why": "claims.base é um arquivo .base do Obsidian (formato de view), não um data source. Os dados reais estão nos .md individuais.",
      "ref": ".opencode/command/incorporar.md:90 e vault/bases/claims.base"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "A lista de roles inferíveis na Fase 3 omite `trabalho-futuro`, que é um valor válido segundo `vault-semantic-schema` (linha 89). Se um paper for classificado como trabalho futuro, a inferência automática não terá essa opção.",
      "fix": "Adicionar `trabalho-futuro` à lista de roles na Fase 3.",
      "why": "Divergência entre a lista explicitada no comando e o schema completo definido na skill vault-semantic-schema.",
      "ref": ".opencode/command/incorporar.md:52 e .agents/skills/vault-semantic-schema/SKILL.md:89"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "O fluxo não especifica timeout, retry ou comportamento de falha parcial para nenhuma das 9 fases. Se o `pdf-reader` falhar na extração de texto (PDF scaneado, DRM, corrompido), as Fases 3, 7, e 8 ficam sem dados. Não há estratégia de graceful degradation.",
      "fix": "Adicionar seção de 'Tratamento de Erros' com: (a) timeout de 30s por chamada MCP; (b) se PDF ilegível, marcar `pdf_status: corrompido` e pular Fases 3, 7, 8; (c) se DOI não resolvível, marcar `validation_status: requer-validacao` e prosseguir com metadados mínimos; (d) log de cada falha em `.opencode/log/incorporar-errors.log`.",
      "why": "A pipeline assume sucesso em cascata sem pontos de recuperação.",
      "ref": ".opencode/command/incorporar.md (ausente)"
    }
  ],
  "recommendation": "Corrigir as 3 falhas críticas antes do primeiro uso real: (1) unificar `bibtex_key`/`bibtex-key` removendo `bibtex_key` do template e ajustando Fase 5; (2) adicionar branch `arxiv:` na Fase 2 com ferramentas apropriadas; (3) substituir `doiget_bibtex_export` por `scholar-sidekick_exportCitation` como fonte primária de BibTeX. As falhas significantes podem ser tratadas incrementalmente, mas a seção de tratamento de erros é pré-requisito para qualquer uso em lote.",
  "schema_version": 3
}
```

---

## Análise Completa — Judge 1

### Resumo

O `/incorporar` é um comando ambicioso que unifica 9 fases de processamento acadêmico. O design conceitual é sólido e cobre o ciclo completo: PDF → metadados → resumo → BibTeX → vault → canvas → claims → conexões → próximos passos. No entanto, a implementação especificada contém **3 falhas críticas de correção** que causariam corrupção de dados ou falha completa da pipeline, mais **6 problemas significantes** de integração e completude.

### Falhas Críticas (3)

#### 1. Duplicação de `bibtex_key` / `bibtex-key` no template

O template `vault/templates/paper-note.md` contém ambas as propriedades:

```yaml
bibtex_key: ''
bibtex-key: ''
```

O `vault-semantic-schema` (SKILL.md:268) declara explicitamente como anti-pattern: "Trocar `bibtex-key` por `bibtex_key` removendo compatibilidade com scripts." O script `import-bib-to-vault.sh` usa exclusivamente `bibtex-key`. O `/incorporar` Fase 5 (linha 66) diz preencher `bibtex_key`. Isso garante que:

- Notas criadas via `/incorporar` terão `bibtex-key: ''` (vazio)
- Scripts e Bases que dependem de `bibtex-key` falharão silenciosamente
- As duas propriedades divergirão ao longo do tempo

**Correção**: Remover `bibtex_key` do template e corrigir Fase 5 para `bibtex-key`.

#### 2. Output `arxiv:` tratado como DOI

`scripts/extract-pdf-doi.sh` (linha 42) retorna `arxiv:2301.08745` quando encontra um arXiv ID. A Fase 2 passa esse valor para:

- `crossref_get_work(doi)` — Crossref **não aceita** prefixo `arxiv:`, retornará erro
- `doiget_resolve_paper(doi)` — comportamento indefinido com prefixo não-DOI
- `scholar-sidekick_checkOpenAccess(doi)` — mesmo problema

Não há nenhuma lógica condicional para detectar o prefixo e redirecionar para ferramentas arXiv (`arxiv_get_abstract`, `arxiv_download_paper`). A pipeline inteira colapsa nesse ponto para qualquer PDF que não tenha DOI explícito mas tenha arXiv ID.

**Correção**: Adicionar branch na Fase 2: se `$EXTRACTED_ID` começa com `arxiv:`, usar ferramentas arXiv; se começa com `10.`, usar ferramentas Crossref/DOI.

#### 3. Dependência circular com `doiget_bibtex_export`

A Fase 2 (linha 43) diz: "Obtém BibTeX formatado via `doiget_bibtex_export(doi)`."

Mas `doiget_bibtex_export` **só funciona para entradas que já existem no store local do doiget**. O store é populado por `doiget_fetch_paper`. O problema:

- Se `doiget_fetch_paper` falhar (paper não-OA) e cair para scihub → doiget store **não tem** a entrada
- Se for PDF local, doiget nunca foi chamado → store **não tem** a entrada
- Se a entrada foi obtida via `doiget_metadata_only` (sem fetch) → store pode ter metadados mas `bibtex_export` depende da entrada completa

O fluxo já chama `crossref_get_work(doi)` que retorna metadados completos (título, autores, journal, ano, DOI). Com esses metadados, o BibTeX pode ser gerado estaticamente ou via `scholar-sidekick_exportCitation(doi, format='bib')`.

**Correção**: Usar `scholar-sidekick_exportCitation(doi, format='bib')` como fonte primária (stateless, sempre funciona com DOI). `doiget_bibtex_export` como fallback.

### Problemas Significantes (6)

#### 4. Duas pipelines concorrentes para criação de notas

O skill `knowledge-base` (passos 5+7) cria notas via `import-bib-to-vault.sh` e depois as enriquece. O `/incorporar` Fase 5 cria notas diretamente. Esses dois caminhos:

- Usam lógicas de inferência de tags diferentes (regex no script vs manual no `/incorporar`)
- Produzem frontmatter com estruturas potencialmente diferentes
- Não há reconciliação documentada entre os dois métodos

#### 5. Fase 7 (Verificar claims) é computacionalmente inviável

"Para cada claim existente, verifica se o paper apoia/contradiz" — com 60 claims, isso exigiria ler 60 arquivos e fazer matching semântico contra o paper. Não há:

- Heurística de pré-filtro (keywords, claim_type, área temática)
- Threshold de similaridade
- Estratégia de redução de escopo (ex: apenas claims do mesmo método)
- Mecanismo para claims com `strength: forte` que não devem ser alterados sem verificação humana

#### 6. Embed de PDF com caminho ambíguo

Fase 5 especifica `![[pdfs/<key>.pdf]]`. Da perspectiva de uma nota em `vault/papers/`, o caminho resolve para `vault/papers/pdfs/<key>.pdf` (correto no Obsidian). Mas:

- O template tem `<!-- PDF não disponível -->`, não um embed
- O comando precisa saber se deve substituir o comentário ou inserir nova linha
- Se a nota já existe, o comportamento de atualização do embed não está definido

#### 7. `.opencode/log/pdf-events.log` não existe

O diretório `.opencode/log/` está vazio. O plugin `pdf-watcher.ts` escreve no log quando um PDF é detectado. `/incorporar fila` leria um arquivo inexistente. Comportamento não definido.

#### 8. Confusão entre `claims.base` (view) e `claims/*.md` (dados)

Fase 7 (linha 90) diz para ler `vault/bases/claims.base` usando o skill `obsidian-bases`. Mas `claims.base` é um arquivo de configuração de view do Obsidian Bases — ele define filtros, fórmulas e colunas, **não contém os dados dos claims**. Os dados estão nos 60 arquivos `vault/claims/*.md`. A Fase 7 precisa ler os `.md` individuais.

#### 9. Role `trabalho-futuro` ausente

A lista de roles na Fase 3 (linha 52) omite `trabalho-futuro`, que é um valor válido no `vault-semantic-schema` (linha 89).

### Problemas Menores (2)

#### 10. Sem tratamento de erros ou graceful degradation

Nenhuma das 9 fases especifica timeouts, retry, ou comportamento de falha parcial. Se o PDF for ilegível (scaneado, DRM), as Fases 3-8 falham em cascata sem recuperação.

#### 11. Modo sem argumentos fora da tabela

A linha 22 diz "Sem argumentos, lista PDFs pendentes e pergunta qual processar." Este modo interativo não está documentado na tabela de modos de entrada e seu comportamento (como listar? qual interface?) não é especificado.

### Verdict: FAIL

As 3 falhas críticas (bibtex-key duplicado, arxiv tratado como DOI, dependência circular do doiget) são **bloqueantes** — cada uma individualmente causaria corrupção ou falha da pipeline. As falhas significantes agravam o problema mas poderiam ser tratadas incrementalmente.

A arquitetura conceitual do fluxo é correta e o escopo está bem definido. O problema está na **precisão da especificação** — os detalhes de integração entre ferramentas não foram validados contra o comportamento real dos MCPs e scripts existentes.
