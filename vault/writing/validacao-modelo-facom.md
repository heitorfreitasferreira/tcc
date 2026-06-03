---
title: Validação do Roadmap vs. Modelo FACOM
tags:
  - writing
  - monografia
  - validacao
  - modelo-facom
status: draft
created: 2026-06-03
---

# Validação do Roadmap vs. Modelo FACOM

Fonte: `modelo_de_monografia_da_facom.pdf` (FACOM/UFU, abnTeX2)
Objeto: `vault/writing/roadmap-monografia.md`

## Estrutura FACOM vs. Roadmap Atual

| Elemento FACOM | No roadmap? | Situação |
|---|---|---|
| **Pré-textuais** | | |
| Capa | ❌ | Ausente. Criar folha de rosto LaTeX |
| Folha de Rosto | ❌ | Ausente. Criar com natureza do TCC, área, orientador |
| Dedicatória | ❌ | Opcional, mas sem task definida |
| Agradecimentos | ❌ | Opcional, mas sem task definida |
| Epígrafe | ❌ | Opcional, mas sem task definida |
| Resumo (pt-BR) | ❌ | Sem task de escrita. Deve conter objetivo, método, resultados, conclusões |
| Abstract (en) | ❌ | Sem task de escrita. Versão inglesa do resumo |
| Lista de Ilustrações | ❌ | Automática via LaTeX (`\listoffigures`), mas precisa ser verificada |
| Lista de Tabelas | ❌ | Automática via LaTeX (`\listoftables`), mas precisa ser verificada |
| Lista de Siglas/Abrav | ❌ | Automática via `abntex2` (`\siglas`), mas precisa ser alimentada |
| Sumário | ❌ | Automático (`\tableofcontents`), sem ação necessária |
| **Capítulos** | | |
| 1. Introdução | ✅ | Estrutura bate (Motivação, Problema, Hipótese, Objetivos, Contribuições, Organização) |
| 2. Fundamentação Teórica | ✅ | Cobre conceitos + trabalhos correlatos |
| 3. Experimentos | ⚠️ | Modelo FACOM unifica **método + experimentos + avaliação** em UM capítulo. Roadmap separa em "Proposta" (cap. 3) e "Experimentos" (cap. 4), totalizando 5 capítulos. **Divergência estrutural.** |
| 4. Conclusão | ✅ | Principais Contribuições, Trabalhos Futuros, Produção Bibliográfica |
| **Pós-textuais** | | |
| Referências | ✅ | Automático via BibTeX |
| Apêndices | ❌ | Sem task. Material produzido pelo autor (código, tabelas extras, questionários) |
| Anexos | ❌ | Sem task. Material de terceiros (artigos, documentações) |

## Resumo das Lacunas

### 1. Pré-textuais não endereçados (5 tasks faltando)
Capa, Folha de Rosto, Resumo, Abstract e Lista de Siglas precisam de tarefas explícitas de escrita/verificação. Dedicatória, Agradecimentos e Epígrafe são opcionais, mas devem ser decididos.

### 2. Estrutura de capítulos divergente
O modelo FACOM prescreve **4 capítulos**, com Método de Avaliação dentro do capítulo de Experimentos (seção 3.1). O roadmap separa em 5 capítulos, criando um capítulo "Proposta" autônomo. Isso não é proibido — a FACOM permite flexibilidade —, mas o conteúdo de "Método para Avaliação" (parâmetros, métricas, instâncias, baselines) está no capítulo de Proposta no roadmap, e não no de Experimentos como o modelo sugere.

Recomendação: ou fundir Proposta + Experimentos (4 capítulos, como o modelo), ou manter 5 capítulos mas garantir que a seção "Método para Avaliação" apareça dentro de Experimentos, não em Proposta.

### 3. Pós-textuais ausentes
Apêndices podem ser necessários para:
- Tabelas completas de resultados (por semente)
- Pseudocódigo dos algoritmos
- Resultados do lower bound por instância
- Exemplo de instância `.graph`

Anexos podem não ser necessários, mas a estrutura deve prever.

### 4. Formatação ABNT não verificada
O roadmap não menciona:
- Formatação de citações diretas (>3 linhas: recuo 4cm, fonte menor)
- Alíneas e subalíneas
- Remissões internas (`\ref`, `\pageref`)
- Uso de siglas (`\ac`, `\acs`, `\acl`)

Esses são detalhes de template LaTeX (já resolvidos pela classe `ppgco.cls`), mas precisam ser conferidos no momento da compilação final.

## Ações Recomendadas (Pré-Escrita)

- [ ] Criar/Capturar conteúdo de **Capa e Folha de Rosto** (dados: autor, título, orientador, data)
- [ ] Escrever **Resumo** (português) e **Abstract** (inglês) — 150-500 palavras cada
- [ ] Decidir sobre **Dedicatória, Agradecimentos, Epígrafe** (opcionais)
- [ ] Revisar **Lista de Siglas** — levantar siglas usadas no texto
- [ ] Decidir estrutura: **4 capítulos (modelo FACOM)** ou **5 capítulos (roadmap atual)**; se 5, mover "Método de Avaliação" para dentro de Experimentos
- [ ] Prever **Apêndices**: ao menos um apêndice com resultados completos
- [ ] Verificar template LaTeX (`ppgco.cls`) se cobre citações longas, alíneas e remissões

## Conclusão

**O roadmap está maduro para o conteúdo central**, mas **ignora completamente os elementos pré e pós-textuais**. As 10 tarefas preparatórias (P1-P10) cobrem bem a parte técnica/experimental. Faltam tarefas de escrita para:

| Prioridade | O quê |
|---|---|
| Alta | Resumo + Abstract |
| Alta | Capa + Folha de Rosto |
| Alta | Decidir estrutura 4 vs. 5 capítulos |
| Média | Lista de Siglas |
| Média | Conteúdo de Apêndices |
| Baixa | Dedicatória, Agradecimentos, Epígrafe |

A maior decisão pendente é: seguir o modelo FACOM com **4 capítulos** (Proposta fundida em Experimentos) ou manter **5 capítulos** com justificativa explícita.
