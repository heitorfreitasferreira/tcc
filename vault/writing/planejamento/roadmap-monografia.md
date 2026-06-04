---
title: Roadmap de Escrita da Monografia
tags:
  - writing
  - monografia
  - roadmap
  - agentes
status: pronto-com-pendencias-bibliograficas-rastreabilidade-e-validacao
created: 2026-06-02
updated: 2026-06-05
concluded: 2026-06-03
---

# Roadmap de Escrita da Monografia

Esta nota é o ponto de partida operacional para agentes que irão escrever a monografia. A escrita deve usar o `vault/` como camada organizada de conhecimento e a estrutura de diretórios da `monografia/` como destino, sem assumir conteúdo já escrito nos arquivos `.tex`.

## Hierarquia de Informação

A fonte determinística de informação do projeto é sempre o código em `src/` e os dados brutos gerados por ele em `src/data/`. O `vault/` não é fonte primária: ele é uma camada de organização mantida por agentes a comando do autor. Portanto, antes de escrever qualquer afirmação técnica, metodológica ou experimental, o agente deve conferir se ela ainda é refletida pelo código e pelos dados.

Fluxo obrigatório de informação:

```text
src/                 → define implementação, algoritmos, flags, schemas e renderizações
src/data/            → contém instâncias, grafos e resultados brutos gerados pelo código
leitura código+dados → valida o que foi realmente implementado e medido
literatura           → contextualiza, fundamenta e compara com trabalhos externos
vault/               → organiza síntese, claims, referências e plano de escrita
monografia/          → recebe o texto final, figuras, tabelas e citações
```

Ordem de autoridade quando houver conflito:

| Ordem | Fonte | Papel |
|---:|---|---|
| 1 | `src/` | Verdade sobre implementação, parâmetros disponíveis, algoritmos, schemas e endpoints |
| 2 | `src/data/` | Verdade sobre instâncias, resultados brutos, tempos, evolução e artefatos gerados |
| 3 | Literatura | Verdade externa para conceitos, métodos, lacunas e trabalhos relacionados |
| 4 | `vault/` | Organização e síntese produzida por agentes; deve ser revalidada contra código/dados |
| 5 | `monografia/` | Documento final; deve herdar informação validada das camadas anteriores |

> [!warning] Regra de validação
> Se uma nota do `vault/` contradizer o código ou os dados brutos, o código e os dados vencem. O agente deve atualizar a nota ou marcar a inconsistência antes de usar essa informação na monografia.

## Regra Central

Escrever em português acadêmico, com foco no cenário de patrulha com drones e no problema TSP-SD-ATP. Toda afirmação técnica ou experimental forte deve estar ancorada primeiro em `src/` e `src/data/`; o `vault/` serve para localizar e organizar essa evidência.

| Tipo de afirmação | Fonte obrigatória |
|---|---|
| Conceito clássico de TSP, NP-dificuldade, variantes | Literatura organizada em `vault/areas/` e `vault/papers/` |
| GA, PSO, ACO, busca exaustiva e lower bound AP como implementados | Código em `src/optimization/`, `src/graph/`, `src/shared/`; síntese em `vault/projeto/` |
| Formulação TSP-SD-ATP | Código em `src/graph/` e `src/points/`; síntese em [[problem-formulation]] |
| Arquitetura e pipeline | Código em `src/cmd/`, scripts `src/run_*.sh`, schemas em `src/shared/`; síntese em [[architecture]], [[experiment-pipeline]] |
| Resultados e conclusões empíricas | Dados brutos em `src/data/results/`; síntese em [[resultados]], [[analysis-methodology]] |
| Figuras e tabelas | Dados em `src/data/`, scripts em `scripts/`, renderização em `src/web/`, saídas em `monografia/figs/` |

> [!warning] Claims fortes
> Não escrever como fato absoluto que este é o "primeiro" estudo publicado sem formulação defensável. Preferir: "não foram identificados, na revisão realizada, estudos que comparem sistematicamente GA, PSO e ACO na variante TSP-SD-ATP".

## Ordem Recomendada de Escrita

| Ordem | Capítulo | Por que escrever nesta posição | Nota-base |
|---:|---|---|---|
| 1 | Fundamentação Teórica | Define vocabulário, citações e limites conceituais | [[fundamentacao]] |
| 2 | Proposta | Descreve a formulação e a implementação, usando vocabulário já estabilizado | [[proposta]] |
| 3 | Experimentos | Depende da proposta e precisa de números/figuras finais | [[experimentos]] |
| 4 | Introdução | Deve refletir exatamente o que a monografia entrega | [[introducao]] |
| 5 | Conclusão | Só deve fechar depois dos resultados e limitações estarem estabilizados | [[conclusao]] |

## Entregáveis Globais Antes da Escrita Final

| Entregável | Status esperado | Fonte |
|---|---|---|
| Lista de claims centrais | Criada antes de Introdução e Conclusão | [[resultados]], [[analysis-methodology]] |
| Tabela de referências por capítulo | Criada antes de Fundamentação | `vault/papers/index.md` |
| Padronização terminológica | Criada antes de qualquer capítulo | [[problem-formulation]], [[tsp-variants]] |
| Figuras/tabelas mínimas de resultados | Criadas antes de Experimentos | `scripts/`, `monografia/figs/` |
| Protocolo estatístico final | Fechado antes de Experimentos | [[analysis-methodology]] |
| Pipeline de métricas TeX auditáveis | Criado antes de Experimentos, Resumo, Abstract e Conclusão | `scripts/gerar-metricas-monografia.py`, `monografia/generated/` |
| Gate formal anti-alucinação | Executado antes da versão final e antes de qualquer entrega para banca | `scripts/check-monografia.sh`, `vault/writing/planejamento/citation-safety.md`, [[claim-evidence-matrix]] |
| Auditoria bibliográfica pós-council | Referências centrais com PDF/nota validada; periféricas classificadas como download manual, análise posterior ou descarte | `vault/papers/index.md`, seção [[#Correções Bibliográficas Pós-Auditoria]] |
| **Resumo + Abstract** | Escritos antes da introdução, revisados por último | [[validacao-modelo-facom]] |
| **Capa + Folha de Rosto** | Dados preenchidos antes da compilação final | autor, orientador, título, data |
| **Lista de Siglas** | Levantamento antes da escrita de qualquer capítulo | `vault/siglas/` (notas individuais) + `vault/bases/siglas.base` |
| **Apêndices** | Estrutura definida antes de Experimentos | `vault/writing/planejamento/apendices.md` |

## Protocolo de Validação Antes de Escrever

Todo agente deve executar esta sequência antes de produzir texto para a monografia:

1. Identificar quais afirmações o capítulo precisa fazer.
2. Conferir no `src/` como o problema, método, pipeline ou renderização está implementado.
3. Conferir em `src/data/` quais instâncias, resultados e tempos existem de fato.
4. Buscar e entender a literatura necessária para contextualizar os achados.
5. Organizar a síntese no `vault/`, corrigindo notas que estejam desatualizadas.
6. Usar o `vault/` validado como base de escrita para `monografia/`.

O agente não deve tratar uma nota do `vault/` como evidência suficiente quando a afirmação depender de implementação ou resultado experimental. Nesses casos, a nota deve apontar para o código ou para o dado bruto que sustenta a afirmação.

## Correções Bibliográficas Pós-Auditoria

Auditoria realizada em 2026-06-03 sobre `vault/papers/`: 61 notas de paper, 43 PDFs íntegros com chave correspondente, 18 notas sem PDF íntegro correspondente e 1 PDF duplicado/extra (`haroun2015.pdf`). A monografia não deve esperar que todas as 61 referências sejam completadas. O critério operacional é mais estreito: toda referência usada para sustentar claim conceitual forte, fundamentação central ou comparação metodológica precisa ter PDF íntegro, metadados corretos e nota coerente com o conteúdo do PDF.

### Política de Uso das Referências Auditadas

| Classe | Critério | Uso na monografia |
|---|---|---|
| **Usar agora** | PDF íntegro, nota coerente e metadados suficientes | Pode entrar no texto, respeitando a força do claim |
| **Corrigir antes de usar** | PDF existe, mas há erro de autores, título, DOI, status ou leitura parcial | Não citar até corrigir a nota ou o BibTeX |
| **Solicitar download manual** | Referência é relevante para argumento central, mas PDF está ausente/corrompido | Pedir download manual, analisar depois e só então citar |
| **Manter para análise posterior** | Referência é útil para trabalhos futuros ou estado da arte, mas não sustenta claim central | Não bloqueia escrita; pode ficar fora do texto final |
| **Descartar do escopo imediato** | Referência é redundante, periférica ou não altera argumento da monografia | Remover da lista de referências centrais e não bloquear escrita |

> [!warning] Regra prática
> PDF ausente não implica automaticamente corrigir agora. Se o paper não sustenta Introdução, Fundamentação, Proposta ou Experimentos, classificar como análise posterior ou descarte é preferível a atrasar a escrita.

### Correções Obrigatórias de Exatidão

| Item | Problema | Ação |
|---|---|---|
| [[gpaco2025]] | Autores na nota não batem com o PDF; PDF indica Bo-Cheng Lin, Yi Mei e Mengjie Zhang | Corrigir autores antes de usar em trabalhos futuros ou revisão de ACO moderna |
| [[neufaco2025]] | Autores na nota não batem com o PDF; PDF indica Dat Thanh Tran, Khai Quang Tran, Khoi Anh Pham, Van Khu Vu e Dong Duc Do | Corrigir autores antes de citar como estado da arte |
| [[ppaco2024]] | PDF local quase não permite extração textual; nota está corretamente limitada a `lido-parcial` | Manter apenas como fronteira de pesquisa; não usar para resultados ou metodologia |
| [[applegate2006traveling]], [[dorigo2004book]] | Livros longos marcados como `lido`, mas notas parecem leitura seletiva | Qualificar como leitura seletiva ou `lido-parcial` se a nota não for expandida |
| `haroun2015.pdf` | PDF extra/duplicado sem nota correspondente; já existe [[haroun2015performance]] | ~~Manter apenas o arquivo correspondente à chave usada ou registrar duplicata como obsoleta~~ **Resolvido em P17 (2026-06-04):** arquivo `haroun2015.pdf` não existe; apenas `haroun2015performance.pdf` está presente e é válido. Nenhuma ação necessária. |
| [[demsar2006statistical]] | Referência estatística é usada na metodologia, mas a nota do vault está vazia e fora de `vault/papers/` | Criar/preencher `vault/papers/demsar2006statistical.md` ou mover/enriquecer a nota existente antes de escrever Experimentos |

### Triagem dos PDFs Faltantes

| Decisão | Referências | Encaminhamento |
|---|---|---|
| **Download manual prioritário se forem citadas** | [[lawler1985traveling]], [[aggarwal2000angular]], [[vanhove2012route]] | Fundamentam TSP clássico, AM-TSP e turn costs. Se o texto depender delas, solicitar PDF manualmente e revisar a nota antes da citação. Se não houver tempo, usar referências já validadas como [[garey1979computers]], [[applegate2006traveling]], [[winter2002modeling]] e [[kinable2017hybrid]]. |
| **Download manual condicionado à seção de lower bounds** | [[heldkarp1971traveling]], [[balas1985branch]], [[fischetti1992additive]], [[valenzuela1997estimating]], [[leraromero2020dynamic]] | Não bloquear a monografia se o texto limitar a discussão ao lower bound AP implementado, [[heldkarp1970traveling]], [[johnson1996asymptotic]], [[righini2021efficient]], [[kinable2017hybrid]] e [[justificativa-lowerbound]]. Baixar manualmente apenas se o capítulo aprofundar bounds clássicos ou ng-path. |
| **Download manual para análise posterior** | [[dellamico2021multiple]], [[dellamico2022exact]], [[deepaco2023]], [[ahmed2024receding]], [[nagata2006eax]] | Relevantes para trabalhos relacionados ou futuros, mas não indispensáveis para demonstrar o que foi implementado. Não usar para claims fortes até recuperar PDF íntegro. |
| **Candidatas a descarte do escopo imediato** | [[hga2024hybrid]], [[sun2024hybrid]], [[huang2025matrix]], [[kappagantula2025dpso]], [[toaza2023review]] | Manter no catálogo apenas se houver intenção explícita de revisar estado da arte recente. Para a monografia atual, podem ser omitidas sem enfraquecer o argumento central. |

### Lacunas Bibliográficas que Merecem Reforço

| Lacuna | Ação recomendada | Prioridade |
|---|---|---|
| Patrulha UAV / persistent surveillance | Buscar 2 ou 3 referências diretamente sobre patrulha, vigilância persistente, monitoramento ou inspeção com UAVs; não depender apenas de TSP-D/FSTSP de entrega | Alta para Introdução |
| Metodologia estatística para metaheurísticas | Preencher [[demsar2006statistical]] e considerar Derrac et al. (2011) para testes não-paramétricos em algoritmos evolucionários | Alta para Experimentos |
| PSO discreto/permutacional | Reforçar a ponte entre PSO contínuo, random keys e permutações; [[clerc2000discretepso]] ajuda, mas a nota é parcial | Média para Fundamentação |
| Turn costs / angular routing | Usar [[winter2002modeling]] como base segura e baixar manualmente [[aggarwal2000angular]] ou [[vanhove2012route]] se a discussão angular for expandida | Média para Fundamentação |

## Rastreabilidade de Valores Quantitativos

Os valores quantitativos derivados de `src/data/results/` não devem permanecer copiados manualmente na monografia. A meta é que todo número experimental, estatístico ou tabular usado para sustentar uma conclusão seja gerado por scripts determinísticos auditáveis e incorporado ao LaTeX por macros ou fragmentos `.tex`.

O escopo não é transformar toda a prosa em saída de script. O texto interpretativo, as limitações, as justificativas metodológicas, as captions discursivas e números conceituais como $O(n^3)$ continuam manuais. O alvo de automação é o conjunto de valores que mudaria se os arquivos em `src/data/results/` fossem regenerados.

### Arquitetura Recomendada

| Artefato | Função |
|---|---|
| `scripts/gerar-metricas-monografia.py` | Carregar summaries, timing e evolution; validar cobertura esperada; calcular agregados; emitir TeX determinístico |
| `monografia/generated/metrics.tex` | Macros para valores citados no texto, resumo, abstract, experimentos e conclusão |
| `monografia/generated/tables/*.tex` | Fragmentos `tabular` para tabelas numéricas incluídas com `\input` |
| `monografia/generated/manifest.json` | Registro auditável de contagens, hashes ou timestamps dos dados usados, versão do script e arquivos gerados |

Uso esperado no LaTeX:

```tex
\input{generated/metrics.tex}

A avaliação usou \ExpTotalRuns{} execuções estruturadas.

\begin{table}[htbp]
  \centering
  \caption{Gap percentual em relação ao ótimo nas instâncias pequenas.}
  \label{tab:gap-bf}
  \input{generated/tables/tab-gap-bf.tex}
\end{table}
```

### Valores que Devem Ser Gerados

| Classe | Exemplos | Saída recomendada |
|---|---|---|
| Cobertura experimental | total de summaries/evolution/timing, runs por método, número de instâncias, sementes | Macros e tabela de cobertura |
| Busca exaustiva | ótimos das instâncias `10a`--`15c`, número de instâncias com ótimo conhecido | Tabela `.tex` e macros de contagem |
| Gap vs. ótimo | gap médio, mínimo, máximo e taxa de acerto por método | Tabela `.tex` e macros centrais |
| Instâncias grandes | melhor e média de makespan em `50a`--`50c` e `100a`--`100c` | Tabela `.tex` |
| Lower bound AP | gap médio/mínimo/máximo LB→BF, violações LB $\leq$ BF, tempo do bound | Tabela `.tex` e macros |
| Tempo computacional | médias em ms, razão ACO/GA, razão ACO/PSO, conversão para segundos no texto | Tabela `.tex` e macros |
| Estatística | Friedman/Iman-Davenport, p-valor, ranks médios, CD de Nemenyi, Wilcoxon/Holm | Macros emitidas pelo script estatístico ou por camada comum |

### Regras de Implementação

| Regra | Justificativa |
|---|---|
| Ordenar arquivos, instâncias e métodos explicitamente | Garante saída determinística e diffs estáveis |
| Validar cobertura esperada antes de gerar | Evita atualizar o texto com dataset incompleto |
| Padronizar arredondamento em um único lugar | Evita divergências entre texto, tabelas e figuras |
| Gerar vírgula decimal LaTeX como `{,}` e percentual como `\%` | Preserva tipografia correta em pt-BR |
| Gerar apenas o `tabular` quando possível | Mantém caption, label e posicionamento sob controle do capítulo |
| Oferecer modo `--check` | Falha quando os `.tex` gerados estão desatualizados em relação aos dados |
| Reusar a lógica estatística já validada em `scripts/analise-estatistica.py` | Evita duplicar Friedman, Nemenyi e Wilcoxon em outro script |

> [!warning] Anti-overengineering
> Não automatizar texto interpretativo nem cada número conceitual da monografia. Automatizar os valores que sustentam tabelas, conclusões quantitativas e testes estatísticos; manter a argumentação acadêmica sob edição humana.

## Validação Formal Anti-Alucinação

A monografia precisa de um gate explícito contra alucinações de IA. Compilar o PDF não basta: é necessário validar que citações existem e são adequadas, que claims fortes têm evidência, que números experimentais vêm de dados/scripts, que artefatos citados existem e que a conclusão não extrapola o que foi medido.

Essa etapa deve ser híbrida. Checks determinísticos bloqueiam erros objetivos; revisão humana guiada decide casos semânticos, como se uma citação realmente sustenta a frase ou se uma conclusão está forte demais.

### Comando Canônico

| Comando | Função |
|---|---|
| `scripts/check-monografia.sh` | Executa os checks automáticos e gera relatório final |
| `make check-monografia` | Atalho opcional para o script canônico |

O comando deve retornar:

| Saída | Significado |
|---|---|
| `PASS` | Sem erro bloqueante; warnings revisados ou ausentes |
| `WARN` | Texto compila, mas há pendências que exigem revisão humana documentada |
| `FAIL` | Há erro bloqueante: citação inválida, claim bloqueado, número desatualizado, artefato ausente ou referência frágil em claim forte |

### Checks Automáticos Bloqueantes

| Classe | Check |
|---|---|
| LaTeX/BibTeX | Compilar com BibTeX; falhar em citações indefinidas, referências indefinidas, arquivos ausentes, `??`, `TODO`, `FIXME` ou `lipsum` ativo |
| Citações | Extrair `\cite{...}` dos `.tex`; verificar chave no BibTeX, nota em `vault/papers/`, status de leitura/PDF e classificação em `citation-safety.md` |
| Figuras e tabelas | Verificar se todo `\includegraphics` e todo `\input` existe; quando gerado, conferir origem em `scripts/`, `src/web/` ou `monografia/generated/manifest.json` |
| Métricas | Rodar `scripts/gerar-metricas-monografia.py --check`; falhar se `monografia/generated/` estiver ausente ou desatualizado |
| Números hardcoded | Alertar ou falhar quando Resumo, Abstract, Experimentos ou Conclusão contiverem números experimentais literais que deveriam vir de macros |
| Claims bloqueados | Procurar padrões associados a claims `Bxx` da [[claim-evidence-matrix]], como ótimo em 100\%, gaps antigos, tempos antigos e superioridade universal |
| Placeholders | Falhar em conteúdo de template, texto incompleto ou marcadores de revisão não resolvidos |

### Política de Segurança de Citações

Criar `vault/writing/planejamento/citation-safety.md` para classificar cada referência citada na monografia.

| Classe | Critério | Uso permitido |
|---|---|---|
| `segura` | BibTeX correto, nota validada, PDF íntegro ou fonte bibliográfica confiável, conteúdo coerente com o uso | Pode sustentar claim forte |
| `condicional` | Referência real, mas nota parcial, PDF ausente controlado ou uso apenas contextual | Pode sustentar claim fraco ou contextual, com ressalva |
| `frágil` | PDF ausente/corrompido, metadados incompletos ou nota não validada | Não pode sustentar claim forte |
| `bloqueada` | Metadados incorretos, PDF errado, referência não conferida ou incompatível com a frase | Remover ou corrigir antes de citar |

Referências já identificadas como sensíveis devem entrar nessa classificação antes da versão final: `lawler1985traveling`, `vanhove2012route`, `dellamico2021multiple`, `dellamico2022exact`, `deepaco2023`, `nagata2006eax`, `demsar2006statistical`, `gpaco2025` e `neufaco2025`.

### Vínculo Entre Texto e Claims

Claims fortes do `.tex` devem ser conectados à matriz, no mínimo por comentários LaTeX próximos ao parágrafo:

```tex
% claim: E06
Nas 18 instâncias com busca exaustiva, o ACO apresentou o menor gap médio...
```

Alternativamente, uma macro invisível pode ser criada depois, por exemplo `\claimref{E06}`. O requisito mínimo é que um revisor consiga localizar qual claim da matriz sustenta cada trecho forte.

### Revisão Humana Guiada

O relatório do gate deve listar itens que nenhum script decide sozinho:

| Item | Pergunta de revisão |
|---|---|
| Citação semântica | A referência citada sustenta exatamente a frase? |
| Lacuna bibliográfica | A frase “não foram identificados estudos...” está limitada à revisão realizada? |
| Conclusão | A conclusão é proporcional aos dados e às limitações? |
| Claims interpretativos | O texto distingue interpretação de fato demonstrado? |
| Referências condicionais | A referência pode permanecer como contexto ou deve ser removida? |
| Estilo acadêmico | O trecho evita generalizações, causalidade indevida e linguagem promocional? |

### Critério de Aprovação

Antes da versão final, o gate anti-alucinação deve produzir um relatório arquivado no vault, por exemplo `vault/writing/planejamento/validacao-anti-alucinacao.md`, contendo:

| Seção | Conteúdo mínimo |
|---|---|
| Resultado do gate | PASS/WARN/FAIL, data, commit ou estado dos arquivos |
| Citações | Lista de chaves usadas, classificação e pendências |
| Claims | Claims fortes encontrados, IDs correspondentes e evidências |
| Números | Confirmação de que valores experimentais vêm de `monografia/generated/` |
| Figuras/tabelas | Conferência de existência e origem |
| Revisão humana | Decisões tomadas sobre warnings semânticos |

## Terminologia Oficial

Usar os termos abaixo de forma consistente:

| Conceito | Termo recomendado | Observação |
|---|---|---|
| Problema principal | TSP-SD-ATP | Definir na primeira ocorrência |
| Tradução explicativa | TSP com penalidades angulares dependentes da sequência | Usar em texto corrido |
| Apelido interno | rTSP | Evitar na monografia, exceto se explicar que é nomenclatura interna |
| Métrica objetivo | makespan | Definir como tempo/custo total da rota |
| Pontos visitados | pontos de interesse ou POIs | Manter uma forma predominante |
| Veículo | drone ou VANT | Escolher uma forma predominante; "drone" é mais direto |

## Claims Centrais Permitidos

Estes claims podem orientar a escrita, mas devem ser verificados contra os dados finais antes de entrar no texto definitivo.

| Claim | Força recomendada | Evidência esperada |
|---|---|---|
| O problema implementado difere do TSP clássico porque o custo depende do nó anterior | Forte | [[problem-formulation]] |
| O tensor 3D permite avaliar rotas em O(n), após pré-computação O(n³) | Forte | [[problem-formulation]] |
| GA, PSO e ACO são metaheurísticas canônicas para TSP e variantes | Forte | [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]] |
| ACO obteve melhor qualidade de solução nas instâncias avaliadas | Forte se os dados finais confirmarem | [[resultados]] |
| GA foi muito mais rápido que ACO, com pior qualidade em instâncias grandes | Forte se os dados finais confirmarem | [[resultados]], timing |
| PSO teve desempenho inferior nesta variante específica | Moderado; explicar pela codificação random keys com cautela | [[resultados]], [[pso]] |
| O lower bound via relaxação AP sobre matriz reduzida é um limitante inferior válido para o TSP-SD-ATP | Forte (por construção matemática) | [[lower-bounds]], código em `src/optimization/lowerbound/` |
| O bound AP é válido mas frouxo para o TSP-SD-ATP: gap ~50% (40–65%) vs brute-force nas instâncias pequenas | Forte (dados confirmam) | Auditoria em `src/data/results/` vs BF; [[auditoria-codigo-dados-vault]] |
| A variante TSP-SD-ATP não aparece como benchmark consolidado na literatura revisada | Moderado; escrever como lacuna da revisão realizada | [[comparative-studies]] |

## Roteiro por Capítulo

### Capítulo 1 — Introdução

**Objetivo:** apresentar o problema, justificar o cenário de patrulha com drones, delimitar a lacuna e declarar contribuições.

**Entradas obrigatórias:** [[introducao]], [[problem-formulation]], [[drone-routing]], [[comparative-studies]], [[resultados]].

**Estrutura sugerida:** contexto de patrulha com drones; roteamento como problema central; TSP como abstração inicial; limitação do TSP clássico para curvas; definição sucinta do TSP-SD-ATP; lacuna da revisão; objetivo geral; objetivos específicos; contribuições; organização do texto.

**Critérios de aceite:** a introdução não promete mais do que os experimentos entregam; a lacuna está escrita de forma defensável; os objetivos são mensuráveis; as contribuições correspondem a capítulos e artefatos reais.

**Riscos:** exagerar novidade; transformar a aplicação em drone delivery em vez de patrulha; antecipar resultados com números ainda não validados.

### Capítulo 2 — Fundamentação Teórica

**Objetivo:** dar suporte conceitual para o problema, os métodos e a comparação experimental.

**Entradas obrigatórias:** [[fundamentacao]], [[tsp]], [[tsp-variants]], [[bio-inspired-optimization]], [[genetic-algorithms]], [[particle-swarm]], [[ant-colony]], [[drone-routing]], [[comparative-studies]].

**Estrutura sugerida:** TSP clássico e complexidade; variantes relevantes do TSP; custos de curva e dependência de sequência; roteamento de drones; metaheurísticas bio-inspiradas; GA para TSP; PSO para TSP; ACO para TSP; estudos comparativos relacionados.

**Referências centrais seguras:** [[garey1979computers]], [[applegate2006traveling]], [[winter2002modeling]], [[holland1975adaptation]], [[goldberg1989genetic]], [[kennedy1995particle]], [[clerc2000discretepso]], [[dorigo1996ant]], [[dorigo1997ant]], [[stutzle2000mmas]], [[murray2015flying]], [[agatz2018optimization]], [[chandra2022comparative]], [[wu2020comparative]], [[halim2019combinatorial]]. Referências condicionais: [[lawler1985traveling]], [[vanhove2012route]] e [[aggarwal2000angular]] só devem ser usadas após download manual ou substituídas por fontes já validadas; [[demsar2006statistical]] precisa de nota preenchida em `vault/papers/` antes de sustentar a seção estatística.

**Critérios de aceite:** cada método tem origem, mecanismo básico e relação com TSP; a variante TSP-SD-ATP é posicionada sem inventar uma taxonomia não suportada; trabalhos relacionados são sintetizados, não listados; o capítulo prepara diretamente a proposta.

**Riscos:** excesso de revisão genérica; citar metaheurísticas modernas sem conexão com o que foi implementado; confundir TSP-D/FSTSP com o problema deste trabalho; esquecer de incluir lower bound AP na descrição dos métodos implementados.

### Capítulo 3 — Proposta

**Objetivo:** descrever o problema implementado, a representação computacional e os algoritmos usados no benchmark.

> [!note] Decisão estrutural
> O modelo FACOM prescreve 4 capítulos com "Método para Avaliação" dentro de Experimentos. Este roadmap adota **5 capítulos** porque a formulação matemática do TSP-SD-ATP (tensor 3D, função objetivo) e a descrição dos algoritmos têm densidade técnica que justifica capítulo próprio. A seção de **Método para Avaliação** (instâncias, parâmetros, métricas, baselines) permanece no capítulo de Experimentos, conforme o modelo.

**Entradas obrigatórias:** código em `src/graph/`, `src/points/`, `src/optimization/`, `src/cmd/`, `src/shared/`; depois [[proposta]], [[problem-formulation]], [[architecture]], [[ga]], [[pso]], [[aco]], [[bruteforce]], [[lower-bounds]].

**Estrutura sugerida:** formulação do TSP-SD-ATP; tensor 3D de custos; função objetivo; representação de soluções (permutação, random keys, trilha de feromônio); arquitetura da implementação; GA (codificação, operadores); PSO (codificação random keys, atualização); ACO (construção 3D, atualização de feromônio); busca exaustiva; lower bound via relaxação AP (redução 3D→2D + Hungarian); política de reprodutibilidade.

**Critérios de aceite:** a formulação matemática bate com o código real em `src/`; o tensor 3D é explicado com clareza; cada algoritmo é descrito em nível de implementação, não apenas conceitual; lower bound AP é justificado teoricamente; **não inclui instâncias, parâmetros ou métricas de avaliação** (esses vão no capítulo de Experimentos, seção Método para Avaliação).

**Riscos:** descrever GA/PSO/ACO de forma genérica sem conexão com a implementação real; não explicar a redução 3D→2D do lower bound; omitir a diferença entre a busca exaustiva (para validação, n ≤ 10) e os métodos estocásticos.

### Capítulo 4 — Experimentos e Resultados

**Objetivo:** apresentar configuração experimental, dados coletados, análise de qualidade, análise de tempo e discussão dos resultados.

**Entradas obrigatórias:** arquivos em `src/data/`, `src/data/results/`, código/scripts que geram resultados e figuras; depois [[experimentos]], [[resultados]], [[analysis-methodology]], [[experiment-pipeline]], figuras em `monografia/figs/` e fragmentos gerados em `monografia/generated/`.

**Estrutura sugerida:**
- **4.1 Método para Avaliação** (conforme modelo FACOM): instâncias e sementes; parâmetros dos métodos; métricas de avaliação (gap vs AP bound, tempo); baseline brute-force (n ≤ 10); lower bound AP como referência para instâncias grandes; plataforma e ambiente computacional
- **4.2 Experimentos**: qualidade das soluções com gap vs AP bound; tempo computacional; trade-off qualidade-tempo; curvas de convergência
- **4.3 Avaliação dos Resultados**: análise estatística (Friedman + Nemenyi + Wilcoxon); discussão por método; visualização de rotas selecionadas; ameaças à validade

**Figuras e tabelas mínimas:** tabela de cobertura experimental; tabela de ótimos brute-force; tabela comparativa AP bound vs ótimo (instâncias pequenas); tabela de gap médio (AP bound como referência); tabela de tempo nas instâncias de 100 pontos; gráfico de qualidade por método e tamanho; gráfico de tempo por método e tamanho (incluindo lower bound, ∼0ms); gráfico qualidade versus tempo; curvas de convergência; visualização de rotas selecionadas. Tabelas numéricas devem vir de `monografia/generated/tables/*.tex` quando derivadas dos resultados.

**Critérios de aceite:** todos os números derivados de resultados têm fonte rastreável em `src/data/results/` e aparecem no capítulo por macros ou `\input` gerados; gráficos têm escala, unidade e legenda; testes estatísticos são reportados apenas se executados; conclusões são proporcionais aos dados brutos; limitações experimentais aparecem antes da conclusão final.

**Riscos:** declarar significância sem teste; comparar médias sem variabilidade; usar apenas melhor caso para métodos estocásticos; esquecer que brute-force só cobre instâncias pequenas; usar AP bound como se fosse ótimo verdadeiro em vez de limitante inferior; copiar manualmente valores que deveriam vir de `monografia/generated/`.

### Capítulo 5 — Conclusão

**Objetivo:** sintetizar o que foi realizado, responder aos objetivos, delimitar limitações e propor trabalhos futuros.

**Entradas obrigatórias:** [[conclusao]], [[resultados]], [[analysis-methodology]], [[comparative-studies]], referências futuras como [[deepaco2023]], [[neufaco2025]], [[gpaco2025]], [[dellamico2021multiple]].

**Estrutura sugerida:** retomada do objetivo; resposta direta aos objetivos específicos; síntese dos principais achados; contribuições; limitações; trabalhos futuros.

**Critérios de aceite:** não apresenta resultado novo; não usa linguagem mais forte que o capítulo de Experimentos; limitações são explícitas; trabalhos futuros derivam das limitações; encerra voltando ao cenário de patrulha com drones.

**Riscos:** concluir que ACO é melhor em geral, em vez de melhor no desenho experimental realizado; minimizar limitações; propor trabalhos futuros desconectados do projeto.

> As perguntas e respostas que originaram estas tarefas estão documentadas em [[entrevista-sessao]].

## Tarefas Preparatórias Para Agentes

| ID | Tarefa | Saída esperada | Prioridade |
|---|---|---|---|
| P1 | Criar matriz claim-evidência | Nota `vault/writing/planejamento/claim-evidence-matrix.md` | Concluída |
| P2 | Atualizar índice de papers | `vault/papers/index.md` com contagem e categorias atuais | Concluída |
| P3 | Criar glossário terminológico | Nota `vault/writing/planejamento/glossario-monografia.md` | Concluída |
| P4 | Fechar protocolo estatístico | Validado em [[analysis-methodology]]; `scripts/analise-estatistica.py` corrigido; `cd-diagram.{svg,png}` regenerado | Concluída |
| P5 | Listar figuras/tabelas finais | Nota `vault/writing/planejamento/figuras-tabelas-monografia.md` | Concluída |
| P6 | Mapear cada capítulo para arquivos `.tex` | Nota `vault/writing/planejamento/mapa-capitulos-tex.md` | Concluída |
| P7 | Auditar divergências entre `src/`, `src/data/` e `vault/` | Nota `vault/writing/auditorias/auditoria-codigo-dados-vault.md` | Concluída |
| P8 | Implementar e validar lower bound para instâncias grandes (n ≥ 15) | Implementação completa: algoritmo Hungarian O(n³) em Go puro, redução 3D→2D, comando `tcc optimize lowerbound`, 30/30 instâncias executadas, figuras/overlays geradas, integração nos scripts de análise. Bound validado contra brute-force (10a: AP=4.86 ≤ BF=8.25). Documentado em [[lower-bounds]], [[justificativa-lowerbound]] | Concluída |
| P9 | Estudar efeito da codificação na comparação justa entre métodos | Concluído em [[auditoria-codificacao-metodos]]. GA usa permutação direta; PSO usa random keys; ACO usa transições 3D. Explicitar como ameaça à validade, sem nova implementação nesta monografia. | Concluída |
| P10 | Análise de sensibilidade a hiperparâmetros | Concluída em [[auditoria-hiperparametros]] como limitação/trabalho futuro. Não bloqueia a escrita; declarar que não houve tuning sistemático. | Concluída como limitação |
| P11 | Escrever Resumo (pt-BR) e Abstract (en) | Nota `vault/writing/planejamento/resumo-abstract.md` com texto final de 150–500 palavras cada, destacando objetivo, método, resultados e conclusões | Pendente |
| P12 | Preencher Capa e Folha de Rosto | Dados do autor, título definitivo, orientador, área de concentração, data em `monografia/` (via template LaTeX) | Concluída |
| P13 | Levantar e definir Lista de Siglas | Notas individuais em `vault/siglas/` (24 gerenciadas) + notas existentes em `vault/projeto/` e `vault/areas/` com tag `siglas` (GA, PSO, ACO, BF, TSP, TSP-SD-ATP). Base Obsidian em `vault/bases/siglas.base`. Script `scripts/gerar-lista-siglas.py` (scan/generate/validate). 30 siglas catalogadas (15 incluir, 5 excluir, 10 pendentes). `monografia/abrev/Abreviaturas.tex` gerado com 15 entradas | Concluída |
| P14 | Definir estrutura dos Apêndices | Nota `vault/writing/planejamento/apendices.md` listando o que vai em cada apêndice (resultados completos, pseudocódigo, instâncias exemplo) | Pendente |
| P15 | Verificar formatação ABNT no template LaTeX | Conferir citações longas (>3 linhas), alíneas, remissões internas e ambiente de siglas na classe `ppgco.cls` | Pendente |
| P16 | Corrigir exatidão bibliográfica crítica | Ajustar autores de [[gpaco2025]] (Lin, Mei, Zhangjie) e [[neufaco2025]] (Tran et al.), qualificar leituras seletivas em livros longos (applegate2006traveling, dorigo2004book → lido-parcial) e remover duplicata `haroun2015.pdf` | Concluída |
| P17 | Triar PDFs faltantes por decisão editorial | Para cada referência sem PDF íntegro: baixar manualmente se central, marcar como análise posterior se útil para futuro, ou descartar do escopo imediato se periférica | Concluída |
| P18 | Fechar bibliografia mínima por capítulo | Cada capítulo deve ter referências suficientes com PDF/nota validada; referências condicionais não podem aparecer no texto final sem validação | Pendente |
| P19 | Criar pipeline de métricas TeX auditáveis | Script `scripts/gerar-metricas-monografia.py` gerando `monografia/generated/metrics.tex`, `monografia/generated/tables/*.tex` e `manifest.json` a partir de `src/data/results/` | Concluída |
| P20 | Substituir valores quantitativos hardcoded | Experimentos, Resumo, Abstract, Introdução e Conclusão usam macros/fragmentos gerados para valores experimentais, estatísticos e tabulares | Concluída |
| P21 | Adicionar verificação de atualização dos artefatos gerados | Modo `--check` ou comando equivalente falha se `monografia/generated/` estiver desatualizado em relação aos dados/scripts | Concluída |
| P22 | Criar política de segurança de citações | Nota `vault/writing/planejamento/citation-safety.md` classificando cada referência citada como `segura`, `condicional`, `frágil` ou `bloqueada` | Concluída — 35 chaves classificadas: 17 seguras, 10 condicionais, 7 frágeis, 0 bloqueadas. Plano de ação com 16 itens por criticidade. Gate de segurança definido. |
| P23 | Implementar gate anti-alucinação | Script `scripts/check-monografia.sh` ou alvo `make check-monografia` validando LaTeX/BibTeX, citações, figuras, placeholders, métricas, claims bloqueados e números hardcoded | Concluída |
| P24 | Vincular claims fortes ao texto | Comentários `% claim: <ID>` ou mecanismo equivalente nos trechos fortes da monografia, conectando `.tex` à [[claim-evidence-matrix]] | Pendente |
| P25 | Gerar relatório final de validação anti-alucinação | Nota `vault/writing/planejamento/validacao-anti-alucinacao.md` com resultado PASS/WARN/FAIL, citações, claims, números, artefatos e decisões humanas | Pendente |
| P26 | Adicionar BibTeX Held-Karp e revisar Seção 2.8 (Lower Bounds) | BibTeX para heldkarp1970traveling (DOI: 10.1287/opre.18.6.1138), heldkarp1971traveling, johnson1996asymptotic, kinable2017hybrid, righini2021efficient, valenzuela1997estimating; revisar Seção 2.8 para citar Held-Karp como referência canônica e justificar AP vs HK para TSP-SD-ATP | Pendente |
| P27 | Corrigir documentação ACO para Ant System (não MMAS) | Código (`src/optimization/aco/ant.go`) implementa Ant System (todas as formigas depositam, sem bounds). Corrigir nota `stutzle2000mmas.md` (remover afirmação falsa) e ajustar texto da fundamentação para documentar Ant System 3D, citando apenas Dorigo (1996, 1997). OU implementar bounds MMAS no código | Concluída |
| P28 | Adicionar referências ATSP/TDTSP ao BibTeX e referencial | Entradas BibTeX para pelo menos 1 survey de ATSP (ex: Öncan et al. 2009 ou Roberti & Toth 2012) + 1 referência TDTSP (Kinable 2017 já fichado); citar na Seção 2.1 ou 2.2 para enquadrar TSP-SD-ATP como instância de classes bem estudadas | Pendente |
| P29 | Obter PDF legível de Bean (1994) e atualizar avaliação | PDF atual contém apenas metadados INFORMS; obter cópia íntegra (Sci-Hub / interlibrary loan), ler corpo do artigo (random keys), atualizar nota `bean1994genetic.md` e rating (0→≥4) | Pendente |
| P30 | Corrigir ratings inconsistentes (0→3) em artigos citados | Atualizar `bean1994genetic` ≥3 (base da codificação PSO), `winter2002modeling` ≥3 (contexto custos de curva), `vanhove2012route` ≥3, `wang2021ant` ≥3 (tuning ACO). Ratings originais são 0 para artigos citados na fundamentação | Pendente |
| P31 | Criar nota Demšar (2006) + BibTeX dos 6 lower bounds restantes | `vault/papers/demsar2006statistical.md` com resumo, avaliação de adequação ao contexto TSP e rating; BibTeX para aggarwal2000angular, balas1985branch, fischetti1992additive, karp1979patching, leraromero2020dynamic, lawler1985traveling | Pendente |
| P35 | Verificar originalidade do feromônio 3D (τ(i,j,k)) via busca intensiva na literatura | Busca sistemática em Google Scholar, Scopus, arXiv e Sci-Hub por: (a) ACO com feromônio n-dimensional (n>2) para TSP ou problemas relacionados; (b) feromônio 3D ou tensor de feromônio em scheduling/routing; (c) representações de feromônio além de arestas (hipergrafos, sequências, triplas); (d) TSP com custo dependente de sequência resolvido com ACO. Resultado: nota `vault/papers/feromonio-3d-originalidade.md` listando precedentes encontrados (ou ausência deles) com DOIs e avaliação de similaridade. Se não houver precedentes, o claim é seguro; se houver, ajustar P32 para citar e diferenciar. PRÉ-REQUISITO para P32. | Concluída — Encontrados 4 precedentes de feromônio multidimensional em ACO (Wang 2013, Wang 2015, MDACO 2025, Starzec 2026), nenhum para TSP com transições de 2ª ordem. Claim de P32 deve ser qualificado, não absoluto. |
| P32 | Documentar feromônio 3D como contribuição algorítmica original | Incluir parágrafo na Seção 2.7 explicitando que τ(i,j,k) tridimensional é adaptação inédita para ACO em TSP (literatura canônica usa τ(i,j) 2D), motivada pela dependência de triplas do TSP-SD-ATP. Buscar precedentes de feromônio n-dimensional em scheduling/routing | Concluída — Parágrafo adicionado em fundamentacao.tex com claim qualificado e citações de 3 precedentes (Wang 2013, Wang 2015, Geng 2025). BibTeX adicionado. |
| P33 | Adicionar literatura de tuning de hiperparâmetros ao referencial | 2–3 referências sobre metodologia de tuning (irace, SMAC, F-Race ou similar); citar na Seção 4.1 (Método para Avaliação) para fundamentar parâmetros fixos como escolha metodológica deliberada. Reavaliar `wang2021ant` e citar se relevante | Pendente |
| P34 | Citar shami2022pso e gad2022pso na Seção 2.6 (PSO) | Ambos estão no BibTeX mas não citados; adicionar citações na Seção 2.6 para reforçar a revisão de PSO, balanceando a cobertura ACO (13) vs PSO (5 ativos). Documentar em [[particle-swarm]] | Pendente |

## Protocolo Para Cada Agente Escritor

Antes de escrever:

1. Ler este roadmap.
2. Ler o código e/ou dados brutos relacionados ao capítulo em `src/` e `src/data/`.
3. Ler a nota-base do capítulo em `vault/writing/capitulos/`.
4. Ler as notas de projeto e área exigidas para o capítulo.
5. Separar claims em três classes: conceituais, metodológicos e experimentais.
6. Verificar se cada claim metodológico tem fonte no código.
7. Verificar se cada claim experimental tem fonte em `src/data/results/`.
8. Atualizar ou sinalizar qualquer nota do `vault/` que não reflita código ou dados atuais.
9. Conferir se cada referência citada está classificada como segura ou foi baixada/analisada manualmente após a auditoria bibliográfica.
10. Conferir se valores experimentais, estatísticos ou tabulares são macros/fragmentos gerados, não cópias manuais.
11. Conferir se claims fortes têm ID da [[claim-evidence-matrix]] ou evidência explícita indicada no trecho.

Durante a escrita:

1. Escrever em português acadêmico direto.
2. Evitar parágrafos genéricos sobre "importância" sem dado, referência ou função argumentativa.
3. Inserir citações apenas quando a referência sustenta a frase específica.
4. Preservar a narrativa de patrulha com drones.
5. Manter o problema como TSP-SD-ATP, não como benchmark genérico de otimizadores.

Após a escrita:

1. Conferir se objetivos, métodos, dados e conclusões estão alinhados.
2. Marcar claims que ainda precisam de validação.
3. Listar figuras/tabelas citadas e verificar se existem.
4. Verificar se não há promessa sem evidência.
5. Conferir formatação ABNT no capítulo: citações diretas (>3 linhas com recuo 4cm), alíneas/subalíneas, remissões internas (`\ref`/`\pageref`), uso correto de siglas (`\ac`/`\acs`/`\acl`).
6. Rodar ou atualizar o relatório de `scripts/check-monografia.sh` quando o capítulo alterar claims, citações, números, figuras ou tabelas.
7. Registrar pendências em nota separada ou no topo do capítulo.

## Protocolo de Análise Estatística (P4 — Concluído)

Com 3 métodos estocásticos × 30 instâncias × 51 sementes = 4590 execuções estocásticas, comparar apenas médias/melhores casos é insuficiente. O protocolo segue Demšar (2006), detalhado em [[analysis-methodology#Protocolo-de-Análise-Estatística]]. O script estatístico foi corrigido e validado em [[auditoria-script-analise-estatistica]].

### Resumo

| Etapa | Teste | O que responde |
|-------|-------|---------------|
| 1 | **Friedman** | Rejeita H₀: "todos os métodos são equivalentes" |
| 2 | **Nemenyi post-hoc** | Se Friedman rejeitar, mostra quais pares (GA×PSO, GA×ACO, PSO×ACO) diferem com significância |
| 3 | **Diagrama CD** | Visualização da diferença crítica entre métodos |
| 4 | **Wilcoxon signed-rank** | Alternativa pareada por instância para reforçar/confrontar Friedman+Nemenyi |

### Implementação

Script `scripts/analise-estatistica.py`:

```python
# Friedman/Iman-Davenport: 3 métodos, 30 instâncias, mediana por instância
# Nemenyi: CD com q_alpha compatível com Demšar
# Wilcoxon/Holm: pareado por instância entre cada par de métodos
```

### Saídas validadas

- Friedman/Iman-Davenport: F(2,58)=293.2222, p=4.710129e-31
- Ranks médios: ACO=1.1000, GA=1.9000, PSO=3.0000
- Nemenyi: CD=0.6050; todos os pares significativos
- Wilcoxon/Holm: todos os pares significativos
- Diagrama CD exportado para `monografia/figs/cd-diagram.{png,svg}`

## Critério de Pronto Para Escrita Definitiva

A monografia estará pronta para escrita definitiva quando estas condições forem atendidas:

| Condição | Aceite |
|---|---|
| Claims principais mapeados | Cada claim tem fonte e força definida |
| Código e dados auditados | Claims metodológicos conferidos em `src/`; claims experimentais conferidos em `src/data/` |
| Estatística resolvida | Testes executados ou análise descritiva assumida explicitamente |
| Figuras mínimas disponíveis | Cada figura tem escala, unidade e fonte |
| Métricas TeX auditáveis | Valores experimentais, estatísticos e tabelas numéricas derivadas dos dados são gerados por script determinístico |
| Artefatos gerados atualizados | `monografia/generated/` passa no modo de checagem contra `src/data/results/` e scripts de análise |
| Gate anti-alucinação aprovado | `scripts/check-monografia.sh` ou `make check-monografia` executado com resultado PASS, ou WARN com decisões humanas registradas |
| Citações classificadas | Toda referência citada consta em `citation-safety.md`; referências frágeis/bloqueadas não sustentam claims fortes |
| Claims fortes rastreados | Claims centrais do `.tex` apontam para IDs da matriz ou evidência explícita |
| Referências centrais selecionadas | Cada capítulo tem bibliografia mínima definida, sem depender de paper sem PDF íntegro |
| Pendências bibliográficas triadas | PDFs ausentes foram classificados como download manual, análise posterior ou descarte do escopo imediato |
| Terminologia estabilizada | TSP-SD-ATP, makespan, drone e POI usados de forma consistente |
| Limitações declaradas | Parâmetros fixos, instâncias sintéticas e baseline limitado aparecem no texto |
| **Pré-textuais prontos** | Capa, Folha de Rosto, Resumo, Abstract, Lista de Siglas escritos e revisados |
| **Apêndices estruturados** | Conteúdo mínimo definido (resultados completos, pseudocódigo) |
| **Formatação ABNT verificada** | Citações longas (>3 linhas), alíneas, remissões internas e siglas funcionando no template `ppgco.cls` |

## Próxima Ação Recomendada

Executar primeiro **P16–P25** para não levar referência frágil, valor quantitativo manual ou claim sem evidência ao texto final. Em paralelo, executar **P11–P15** (pré-textuais, apêndices e formatação). Depois escrever os capítulos na ordem: Proposta → Experimentos → Fundamentação → Introdução → Conclusão.
