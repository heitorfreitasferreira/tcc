---
description: Registra solicitação de revisão do orientador com análise contextual automática
---

You are running the /claudiney command to register advisor feedback about the monograph. Your task is to:

1. **Parse the feedback** — understand what the advisor is asking for

2. **Create the scaffold** — run \`bash scripts/registrar-review-scaffold.sh "<short-description>"\` (generate a short kebab-case description from the feedback). This creates a file at \`vault/writing/review-solicitacoes/R<NN>-<desc>.md\`. Capture the filename from the output.

3. **Analyze the problem** — search the codebase for context. Be fast but thorough:
   - \`src/\` — search for relevant code files, types, algorithms the feedback mentions
   - \`src/data/results/\` — check if relevant experiment data exists
   - \`vault/\` — read notes in \`vault/writing/\`, \`vault/papers/\`, \`vault/areas/\` for context
    - \`monografia/\` — search .tex files for current text on the topic
    - \`vault/claims/\` and \`vault/bases/claims.base\` — check if existing claims are affected
    - \`vault/writing/planejamento/claim-evidence-matrix.md\` — use only for claim schema/protocol guidance

4. **Edit the note** — fill in:
   - YAML \`title:\` with a short PT-BR description
   - YAML \`priority:\` based on impact (alta/media/baixa)
   - YAML \`target_chapter:\` which chapter is affected (fundamentacao/proposta/experimentos/introducao/conclusao/todos)
    - YAML \`evidence_layer:\` where the truth lives (codigo/dados/literatura/vault/monografia) — use the hierarchy from \`vault/writing/planejamento/review-roadmap.md\`
   - YAML \`correction_layers:\` which layers the fix will likely touch
    - YAML \`claim_ids:\` if the feedback maps to specific claim IDs from \`vault/claims/\`
   - Section "Solicitação Original" with a clear restatement of the feedback
   - Section "Análise Técnica" — fill in:
     * **Problema Identificado**: what exactly needs to change
     * **Hierarquia de Informação**: what each source says (src/, src/data/, literature, vault/, monografia/)
     * **Causa Raiz**: why the problem exists (manual copy, outdated info, missing validation, etc.)
   - Section "Plano de Correção" — leave as ...
