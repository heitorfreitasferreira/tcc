```json
{
  "verdict": "WARN",
  "confidence": "HIGH",
  "key_insight": "O comando /incorporar solidifica um pipeline que já existia fragmentado, mas sobrepõe-se ao skill knowledge-base sem desambiguação, referencia 2 skills não encontradas localmente, e omite etapas de verificação pós-criação que o vault-semantic-schema considera obrigatórias.",
  "findings": [
    {
      "severity": "significant",
      "category": "integration",
      "description": "knowledge-base skill cobre 7 das 9 fases do /incorporar sem que o comando declare a relação — skill deve ser marcada como 'complementar para descoberta, /incorporar para ingestão' ou vice-versa",
      "fix": "Adicionar ao /incorporar.md uma seção 'Relação com knowledge-base skill': definir que o skill cobre descoberta (busca, gap analysis) enquanto o comando cobre ingestão (PDF→vault). Atualizar o skill para referenciar /incorporar como caminho preferencial para as fases 4-8.",
      "why": "Ambos descrevem o mesmo pipeline (BibTeX→vault→canvas) com granularidades diferentes; um agente que carregue knowledge-base e depois execute /incorporar fará trabalho redundante ou conflitante",
      "ref": ".agents/skills/knowledge-base/SKILL.md"
    },
    {
      "severity": "significant",
      "category": "completeness",
      "description": "Dependências listam skills obsidian-bases e json-canvas (Fase 7, linha 120) que não existem como arquivos locais — apenas vault-bases-maintainer está disponível para .base e json-canvas é built-in do sistema",
      "fix": "Substituir 'obsidian-bases' por 'vault-bases-maintainer' na lista de dependências da linha 120; verificar se json-canvas está acessível como built-in ou adicionar nota sobre como obtê-lo",
      "why": "A lista de dependências na linha 120 menciona skills que um agente tentará carregar via skill tool e falhará se não disponíveis",
      "ref": ".opencode/command/incorporar.md:120"
    },
    {
      "severity": "significant",
      "category": "integration",
      "description": "scripts/import-bib-to-vault.sh já converte BibTeX→nota vault (fase 5 parcial), mas o comando recria notas manualmente — sem definir se o script é preterido, complementar, ou ainda o caminho canônico para notas skeleton",
      "fix": "Adicionar à Fase 5: 'Se a nota não existe, execute bash scripts/import-bib-to-vault.sh <key> para criar o esqueleto, depois enriqueça as seções.' ou documentar que o comando substitui o script para novas incorporações",
      "why": "O script existe, funciona, e é referenciado por AGENTS.md e vault/opencode-vault.md — silenciosamente contorná-lo quebra a cadeia de ferramentas documentada",
      "ref": "scripts/import-bib-to-vault.sh"
    },
    {
      "severity": "minor",
      "category": "completeness",
      "description": "Nenhuma fase executa scripts/migrate-tags.py após criar/atualizar nota — vault-semantic-schema requer isso para sincronizar areas/methods/chapters/role com tags hierárquicas",
      "fix": "Adicionar ao final da Fase 5 ou como Fase 5b: 'Execute python3 scripts/migrate-tags.py para sincronizar propriedades YAML com tags hierárquicas e garantir consistência com o vault'",
      "why": "vault-semantic-schema linha 275: 'Para verificar consistência de tags no vault: python3 scripts/migrate-tags.py' — se o comando popula tags e propriedades manualmente, corre o risco de dessincronia sem essa verificação",
      "ref": ".agents/skills/vault-semantic-schema/SKILL.md:281-284"
    },
    {
      "severity": "minor",
      "category": "correctness",
      "description": "Coordenadas de posicionamento do canvas na Fase 6 (~x:620 para surveys) não batem com o layout real do canvas (area-comparative em x:400, area-neural em x:500) — risco de posicionamento inconsistente dependendo de qual fluxo posiciona o nó",
      "fix": "Atualizar coordenadas na Fase 6 para refletir o layout real: surveys em ~x:400 (coluna area-comparative), ou definir que novos papers vão abaixo do último existente (+70~80 y) como faz o knowledge-base skill Step 8",
      "why": "O canvas é mantido manualmente e por scripts; posições conflitantes entre o comando e o skill knowledge-base produzem layout desorganizado com novos papers sobrepostos ou fora da área pretendida",
      "ref": ".opencode/command/incorporar.md:82-85"
    },
    {
      "severity": "minor",
      "category": "integration",
      "description": "pdf-watcher.ts e on-new-pdf.sh escrevem ambos em .opencode/log/pdf-events.log usando formatos ligeiramente diferentes (plugin: '2026-06-05T... | path DOI: ...'; script: '2026-06-05T... | path') — /incorporar fila precisa ser tolerante a ambos",
      "fix": "Documentar no /incorporar.md que o modo fila faz parse flexível do log (split por '|', campo 2 = path, campo 3 opcional = DOI prefix), ou padronizar o formato em ambos os emissores",
      "why": "Dois emissores independentes populam o mesmo log; se o parser do modo fila for rígido, entradas de uma fonte serão ignoradas",
      "ref": ".opencode/plugins/pdf-watcher.ts:50"
    }
  ],
  "recommendation": "Aprovar com correções: (1) declarar relação /incorporar ↔ knowledge-base skill, (2) corrigir lista de skills dependentes, (3) integrar import-bib-to-vault.sh e migrate-tags.py no fluxo, (4) alinhar coordenadas do canvas com layout real. Após esses ajustes, o comando unifica coerentemente um pipeline que antes exigia 3-4 ferramentas separadas.",
  "schema_version": 3
}
```

## Análise de Integração e Coesão — Judge 3

### Escopo da avaliação

Avaliei o comando `/incorporar` contra o ecossistema existente: skills (`knowledge-base`, `vault-semantic-schema`, `vault-tagger`), scripts (`import-bib-to-vault.sh`, `extract-pdf-doi.sh`, `download-pdfs.sh`, `migrate-tags.py`), plugins (`pdf-watcher.ts`), comandos irmãos (`/baixar-pdf`, `/claudiney`, `/roadmap-criar`), estruturas de dados (`vault/canvas/`, `vault/claims/`, `vault/bases/`), e convenções (`vault/opencode-vault.md`, template `paper-note.md`).

### O que funciona bem

1. **Depreciação limpa de /baixar-pdf**: O comando antigo já aponta para `/incorporar` com tabela de migração. Sem conflito.

2. **Integração com ecossistema de comandos**: Fase 9 conecta-se a `/roadmap-criar` e `/claudiney` sem acoplamento — apenas sugere, não invoca. Correto.

3. **Respeito ao vault-semantic-schema**: Frontmatter da Fase 5 alinha-se com o schema canônico (inclui `bibtex_key` + `bibtex-key` para compatibilidade, `reading_status`, `pdf_status`, `role`, `areas`, `methods`, `claim_support`). Tags são hierárquicas. Anti-patterns explícitos.

4. **Plugin pdf-watcher → modo fila**: O plugin escreve em `.opencode/log/pdf-events.log` e `/incorporar fila` consome esse log. Pipeline de watcher→ingestão bem desenhado, sem acoplamento forte (arquivo de log como contrato).

5. **Reuso de scripts existentes**: `extract-pdf-doi.sh` (Fase 2) e `download-pdfs.sh` (Fase 1) são referenciados corretamente.

### Problemas de integração

#### 1. Sobreposição com knowledge-base skill (significativo)

O skill `knowledge-base` define um workflow de 9 passos que cobre 7 das 9 fases do `/incorporar`:

| knowledge-base Step | /incorporar Fase |
|---|---|
| Step 4: Add BibTeX | Fase 4: Integrar BibTeX |
| Step 5: Import to Vault | Fase 5: Criar nota vault |
| Step 6: Download PDF | Fase 1: Obter PDF |
| Step 7: Enrich Note | Fase 5: Preencher seções |
| Step 8: Update Canvas | Fase 6: Atualizar canvas |

Nem o comando referencia o skill, nem o skill referencia o comando. Um agente que carregue `knowledge-base` e depois execute `/incorporar` para o mesmo paper fará trabalho redundante. A divisão natural seria: skill para **descoberta** (buscar papers, identificar gaps), comando para **ingestão** (PDF→vault→canvas→claims).

#### 2. Dependências de skills inexistentes (significativo)

A linha 120 lista `obsidian-bases` e `json-canvas` como dependências. Nenhum arquivo `SKILL.md` para esses nomes existe em `.agents/skills/`. O skill correto para bases é `vault-bases-maintainer`. `json-canvas` pode ser built-in do sistema, mas não está na lista de `available_skills` do system prompt.

#### 3. Script import-bib-to-vault.sh contornado (significativo)

A Fase 5 diz "Cria `vault/papers/<bibtex-key>.md` usando o template canônico". Mas `scripts/import-bib-to-vault.sh` já faz exatamente isso — parseia o .bib e gera a nota esqueleto. O comando nem usa nem menciona o script. Isso cria dois caminhos divergentes para criar notas vault a partir de BibTeX. O script é referenciado em AGENTS.md, vault/opencode-vault.md, e no próprio knowledge-base skill.

#### 4. Ausência de migrate-tags.py no fluxo (minor)

O vault-semantic-schema (linhas 275-284) estabelece que após qualquer edição de frontmatter, deve-se executar `python3 scripts/migrate-tags.py` para verificar consistência. A Fase 5 popula `areas`, `methods`, `role` e tags hierárquicas manualmente — sem a verificação pós-criação, há risco de dessincronia entre propriedades YAML e tags.

#### 5. Coordenadas do canvas inconsistentes (minor)

A Fase 6 posiciona surveys em ~x:620. O canvas real tem `area-comparative` em x:400 e `area-neural` em x:500. O knowledge-base skill Step 8 usa estratégia diferente: "Place new papers below existing ones (increment y by ~70-80)". O comando e o skill usam estratégias de posicionamento conflitantes.

#### 6. Formato do pdf-events.log com dois emissores (minor)

`pdf-watcher.ts:50` escreve: `$(date -Iseconds) | ${p}${doiInfo}` (DOI inline após path).
`on-new-pdf.sh:19` escreve: `$(date -Iseconds) | $PDF` (apenas path, sem DOI).
Se o parser do modo `fila` assumir formato único, entradas de uma fonte serão perdidas.

### Consistência com vault-semantic-schema

O comando respeita as regras centrais:
- Tags exclusivamente hierárquicas (`namespace/valor`)
- Frontmatter com `bibtex_key` + `bibtex-key` (compatibilidade dual)
- Claims são sugeridos, não criados (anti-pattern documentado)
- Não altera `.tex` da monografia

A única violação é a omissão do passo de verificação com `migrate-tags.py`.

### Anti-patterns avoidance

O comando declara 4 anti-patterns explícitos. Todos estão alinhados com as convenções do vault e decisões prévias do council. Nenhum anti-pattern novo é introduzido.

### Recomendação

Aprovado com ressalvas (WARN). As correções necessárias são pontuais e de baixo esforço. O comando unifica coerentemente um pipeline que antes exigia coordenação manual entre 3-4 ferramentas. Após os ajustes, a integração será sólida.
