---
title: Comparacao entre notas canonicas e summaries de papers
type: writing
writing_status: concluido
validation_status: requer-validacao
created: "2026-06-05"
updated: "2026-06-05"
primary_evidence:
  - vault/papers/
  - vault/papers/summaries/
tags:
  - tipo/writing
  - tipo/auditoria
  - evidencia/auditoria
  - topico/roadmap
---

# Comparacao entre `vault/papers/` e `vault/papers/summaries/`

Auditoria criada para consumir o item P45 do [[roadmap-monografia]]. O objetivo e inventariar diferencas entre notas pareadas, sem alterar as notas canonicas em `vault/papers/`.

## Escopo e metodo

- Foram comparadas 49 notas em `vault/papers/summaries/*.md` contra notas de mesmo nome em `vault/papers/*.md`.
- Nenhum par de `summaries/` ficou sem nota canonica correspondente.
- A comparacao abaixo identifica diferencas estruturais, campos de frontmatter ausentes, trechos em `summaries/` potencialmente mais completos e casos que exigem validacao no PDF antes de incorporacao.
- Esta auditoria nao decide incorporacao de conteudo. A decisao deve ocorrer em P46 com revisao humana ou validacao no PDF quando indicada.

## Sintese executiva

- `summaries/` costuma ser mais completo em organizacao analitica: quase todas as notas trazem secoes explicitas de problema/motivacao, metodo, resultados, forcas/limitacoes e takeaway pratico.
- As notas canonicas costumam ser melhores para integracao do vault: preservam frontmatter, tags, wikilinks, relevancia para o TCC, conexoes e citacoes-chave.
- Tres notas canonicas estao praticamente vazias no corpo e devem receber prioridade em P46: [[aggarwal2000angular]], [[balas1985branch]] e [[lysgaard1999cluster]].
- Campos mais frequentemente ausentes nas notas canonicas pareadas: `chapters` e `claim_support` em 48 notas, `aliases` em 46, `pdf_status` em 41, `areas` em 31 e `methods` em 9.
- As divergencias de titulo detectadas sao majoritariamente diferencas de rotulo no resumo, inclusao de autores/ano ou capitalizacao. Casos com titulo truncado ou metadado incompleto devem ser validados em P46 antes de corrigir a nota canonica.

## Prioridades para P46

1. Prioridade alta: incorporar conteudo das notas canonicas vazias: [[aggarwal2000angular]], [[balas1985branch]], [[lysgaard1999cluster]].
2. Prioridade alta: revisar divergencias bibliograficas de DOI, ano ou placeholder: [[aggarwal2000angular]], [[balas1985branch]], [[muthanna2022uav]], [[heldkarp1971traveling]], [[chandra2022comparative]].
3. Prioridade media: incorporar secoes de forcas/limitacoes e resultados de `summaries/` nas notas muito curtas: [[deepaco2023]], [[dellamico2021multiple]], [[dellamico2022exact]], [[wang2021ant]], [[vanhove2012route]].
4. Prioridade baixa: completar campos transversais (`chapters`, `claim_support`, `aliases`, `areas`, `methods`) apenas quando houver uso concreto na monografia, em bases ou em claims.

## Decisoes P46

Decisoes registradas em 2026-06-05, a partir de resposta do autor no fluxo `/roadmap-consumir`: validar e incorporar as tres notas canonicas vazias quando possivel; validar e corrigir conflitos bibliograficos de DOI/ano/titulo; incorporar seletivamente limitacoes, resultados e takeaways das prioridades medias sem promover claims fortes.

| Paper | Decisao | Acao aplicada | Validacao/ressalva |
|---|---|---|---|
| [[aggarwal2000angular]] | Incorporar e corrigir metadado | Nota canonica enriquecida; DOI corrigido para `10.1137/S0097539796312721`; BibTeX atualizado | Confirmado na primeira pagina do PDF por `PII. S0097539796312721` e URL SIAM `/31272.html`; manter `aggarwal1999angular` como chave nova e `aggarwal2000angular` como alias historico |
| [[balas1985branch]] | Incorporar como relatorio tecnico validado localmente | Nota canonica enriquecida; ano ajustado para `1983`; DOI mantido vazio; chave historica preservada | Primeira pagina do PDF local identifica MSRR 488/Carnegie-Mellon, 1983; se uma versao publicada de 1985 for usada, criar nota/entrada bibliografica separada ou registrar edicao |
| [[lysgaard1999cluster]] | Incorporar seletivamente | Nota canonica enriquecida com tese, resumo, metodos, relevancia, evidencia e limitacoes | Primeira pagina do PDF confirma EJOR 119 (1999), pp. 314--325, PII `S0377-2217(99)00133-2`; resultados quantitativos do summary exigem conferencia antes de citacao forte |
| [[muthanna2022uav]] | Corrigir conflito de DOI e manter periferica | DOI corrigido para `10.1016/j.comcom.2022.04.029`; aliases/schema adicionados; nota reforca uso apenas contextual | Confirmado na primeira pagina do PDF; nao incorporar como referencia central de TSP/bio-inspired classico |
| [[heldkarp1971traveling]] | Preservar canonico | DOI canonico `10.1007/BF01584070` preservado; adicionados campos transversais e nota de decisao P46 | Summary tinha DOI `[to be verified]`; BibTeX e nota canonica ja estavam coerentes apos P43 |
| [[chandra2022comparative]] | Corrigir titulo | Titulo canonico e BibTeX ajustados para remover `Problem`; adicionados campos transversais e limitacao de uso | Primeira pagina do PDF exibe `A Comparative Study of Metaheuristics Methods for Solving Traveling Salesman`; DOI foi preservado do BibTeX |
| [[deepaco2023]] | Incorporar seletivo | Adicionadas evidencias, limitacoes e campos de uso como trabalho futuro | Conteudo nao sustenta claim experimental do TCC; usar apenas para fronteira de pesquisa em ACO neural |
| [[dellamico2021multiple]] | Incorporar seletivo | Adicionadas evidencias, limitacoes e aliases MFSTSP | Numeros de gap/custo do summary foram marcados para revalidacao antes de citacao formal |
| [[dellamico2022exact]] | Incorporar seletivo | Adicionadas evidencias, limitacoes e aliases | Numeros de instancias/gaps devem ser conferidos no PDF antes de uso final |
| [[wang2021ant]] | Incorporar seletivo | Adicionadas evidencias e limitacoes sobre tuning de `alpha`/`beta`; conectado a P33 | Nao comparar diretamente com Ant System puro do repositorio; resultados quantitativos exigem conferencia no PDF |
| [[vanhove2012route]] | Incorporar seletivo | Adicionadas evidencias, limitacoes e aliases de turn restrictions | Conexao com TSP-SD-ATP e conceitual; limiares percentuais devem ser validados no PDF antes de citacao formal |

## Inventario por paper

### agatz2018optimization

- `summaries/` acrescenta uma estrutura mais explicita para forcas/limitacoes, resultados e takeaway pratico.
- `vault/papers/` e mais completo em extensao geral e ja tem secoes canonicas do vault.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o titulo em `summaries/` inclui autores/ano; o titulo canonico contem apenas o titulo bibliografico.
- Validacao no PDF: nao obrigatoria para a diferenca de titulo; recomendada se os resultados forem incorporados como claim forte.
- Recomendacao preliminar: considerar incorporar apenas forcas/limitacoes e takeaway, preservando o titulo canonico.

### aggarwal2000angular

- `summaries/` contem praticamente todo o corpo analitico disponivel.
- `vault/papers/` esta sem corpo estruturado relevante.
- Campos faltantes em `vault/papers/`: `chapters`, `claim_support`.
- Divergencia inventariada em P45: antes da P46, a nota canonica registrava DOI `10.1137/S0097539796312719`, enquanto o summary registrava `10.1137/S0097539796312721`; a nota canonica tambem registrava que a chave BibTeX foi renomeada para `aggarwal1999angular` e que `aggarwal2000angular` e alias.
- Validacao no PDF: necessaria antes de transformar o resumo em nota canonica completa e antes de decidir DOI/chave/ano.
- Recomendacao preliminar: prioridade alta em P46; usar `summaries/aggarwal2000angular.md` como base inicial, mas validar metadados bibliograficos antes de incorporar.
- Decisao P46: incorporado seletivamente em `vault/papers/aggarwal2000angular.md`; DOI corrigido no vault e no BibTeX para `10.1137/S0097539796312721`, confirmado pelo PDF.

### ahmed2024receding

- `summaries/` acrescenta forcas/limitacoes, resultados em bloco proprio e takeaway pratico.
- `vault/papers/` e maior e ja contextualiza melhor a relevancia para o TCC.
- Campos faltantes em `vault/papers/`: `chapters`, `claim_support`, `aliases`.
- Divergencia: diferenca de titulo por preservacao de `{UAV}` no BibTeX/canonico e inclusao de autores/ano no summary.
- Validacao no PDF: nao obrigatoria para titulo; recomendada para resultados quantitativos.
- Recomendacao preliminar: manter a nota canonica e migrar apenas pontos de limitacao/uso pratico se forem uteis.

### alexander2020comparison

- `summaries/` e mais detalhado e separa problema, metodo, resultados, forcas/limitacoes e takeaway.
- `vault/papers/` ja tem estrutura canonica, mas e mais breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de incorporar comparacoes ou rankings.
- Recomendacao preliminar: incorporar resultados e limitacoes se o artigo continuar sendo usado como evidencia comparativa.

### almufti2025comparative

- `summaries/` e parecido em extensao, mas organiza melhor resultados e limitacoes.
- `vault/papers/` tem titulo bibliografico mais informativo.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o heading do summary usa apenas a chave `almufti2025comparative`; a nota canonica tem o titulo completo.
- Validacao no PDF: nao necessaria para manter o titulo canonico; recomendada para resultados especificos.
- Recomendacao preliminar: preservar titulo canonico e usar o summary apenas como fonte auxiliar de organizacao.

### applegate2006traveling

- `summaries/` e substancialmente mais longo e detalha resultados em secao propria.
- `vault/papers/` e mais integrado ao vault, mas mais sintetico.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o heading do summary inclui `Resumo`, autores e ano; nao contradiz o titulo canonico.
- Validacao no PDF: recomendada para numeros e afirmacoes sobre Concorde/branch-and-cut.
- Recomendacao preliminar: incorporar resultados praticos e limitacoes, com checagem no livro/PDF.

### balas1985branch

- `summaries/` contem praticamente todo o conteudo analitico disponivel.
- `vault/papers/` esta sem corpo estruturado relevante.
- Campos faltantes em `vault/papers/`: `doi`, `chapters`, `claim_support`, `aliases`.
- Divergencia: ano/versao diferem; a nota canonica usa `1985`, enquanto o summary descreve `1983` como Management Science Research Report No. MSRR 488.
- Validacao no PDF: necessaria, especialmente porque falta DOI na nota canonica e ha conflito de ano/versao.
- Recomendacao preliminar: prioridade alta em P46; completar nota e metadados somente apos validacao bibliografica da versao correta.
- Decisao P46: nota canonica enriquecida a partir do summary e do PDF local; ano ajustado para 1983 por se tratar do relatorio tecnico MSRR 488; chave `balas1985branch` mantida apenas por compatibilidade historica.

### bean1994genetic

- `summaries/` e mais detalhado em random keys, resultados, forcas/limitacoes e takeaway.
- `vault/papers/` ja contem relevancia direta para o TCC e conexoes.
- Campos faltantes em `vault/papers/`: `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o summary inclui autor/ano no heading; o titulo canonico esta correto.
- Validacao no PDF: recomendada para qualquer afirmacao sobre desempenho empirico; nao necessaria para a diferenca de titulo.
- Recomendacao preliminar: incorporar limitacoes e detalhes de random keys se o paper for citado para codificacao.

### bock2025survey

- `summaries/` e mais longo e organiza melhor forcas/limitacoes, resultados e takeaway.
- `vault/papers/` ja tem secoes canonicas e conexoes.
- Campos faltantes em `vault/papers/`: `pdf_status`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de usar taxonomias como base de claim.
- Recomendacao preliminar: incorporar apenas taxonomia/restricoes que ajudem o recorte TSP/variantes.

### chandra2022comparative

- `summaries/` e mais detalhado em resultados, forcas/limitacoes e takeaway.
- `vault/papers/` tem estrutura canonica, mas e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: titulo do summary termina em `Traveling Salesman`; titulo canonico termina em `Traveling Salesman Problem`.
- Validacao no PDF: recomendada para confirmar titulo completo e resultados comparativos.
- Recomendacao preliminar: validar metadado e incorporar detalhes somente se o artigo for mantido como referencia comparativa.
- Decisao P46: titulo corrigido no vault e no BibTeX para a forma da primeira pagina do PDF, `A Comparative Study of Metaheuristics Methods for Solving Traveling Salesman`; DOI preservado do BibTeX.

### clerc2000discretepso

- `summaries/` e mais detalhado e explicita resultados, forcas/limitacoes e takeaway.
- `vault/papers/` e curto e deve continuar marcado com cautela se o PDF estiver fragil.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de reforcar qualquer claim, pois esta referencia ja teve risco bibliografico em auditorias anteriores.
- Recomendacao preliminar: nao incorporar como evidencia forte sem PDF legivel ou fonte alternativa confiavel.

### deepaco2023

- `summaries/` e muito mais detalhado e separa metodo neural, resultados, limitacoes e takeaway.
- `vault/papers/` e muito breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de incorporar resultados, por ser referencia recente e periferica ao escopo principal.
- Recomendacao preliminar: incorporar apenas como contexto de ACO neural, se isso permanecer relevante.
- Decisao P46: incorporadas limitacoes/evidencias apenas como fronteira de pesquisa e trabalho futuro; nao sustenta resultado experimental do TCC.

### dellamico2021multiple

- `summaries/` e muito mais detalhado e inclui estudos de variantes, resultados e limitacoes.
- `vault/papers/` e breve e focada na relevancia para drones.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para variantes e resultados.
- Recomendacao preliminar: incorporar apenas o que sustenta contexto de roteamento com multiplos drones; evitar ampliar o escopo do TCC.
- Decisao P46: incorporados pontos de evidencia e limitacoes para contextualizar drone routing; numeros do summary permanecem marcados para validacao antes de citacao forte.

### dellamico2022exact

- `summaries/` e muito mais detalhado em metodo exato, resultados e limitacoes.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para claims sobre formulacoes exatas.
- Recomendacao preliminar: incorporar apenas como contraste de literatura exata para FSTSP/drone routing.
- Decisao P46: incorporados pontos de evidencia e limitacoes para contraste com metodos exatos; numeros do summary permanecem condicionados a validacao no PDF.

### demsar2006statistical

- `summaries/` e mais detalhado em protocolo estatistico, resultados e limitacoes.
- `vault/papers/` e breve, mas ja criado para sustentar a metodologia estatistica.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o heading do summary inclui autor/ano; titulo canonico esta limpo.
- Validacao no PDF: recomendada apenas para citacoes metodologicas especificas.
- Recomendacao preliminar: incorporar as limitacoes e a justificativa dos testes se ajudarem a documentar o protocolo.

### dorigo1996ant

- `summaries/` e mais detalhado em Ant System, resultados, forcas/limitacoes e takeaway.
- `vault/papers/` ja tem estrutura canonica e relevancia para o TCC.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: o summary inclui autores/ano no heading; nao contradiz o titulo canonico.
- Validacao no PDF: recomendada para numeros de desempenho e descricao precisa do Ant System.
- Recomendacao preliminar: incorporar limitacoes e resultados com cuidado, preservando a distincao Ant System versus ACS/MMAS.

### dorigo1997ant

- `summaries/` e mais detalhado em resultados, mas a nota canonica ja recebeu correcoes recentes de escopo ACS.
- `vault/papers/` tem melhor integracao com o vault.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: heading do summary usa chave + `Resumo`; titulo canonico e bibliografico.
- Validacao no PDF: recomendada antes de migrar qualquer detalhe, para nao reintroduzir confusao entre AS, ACS e ACO generico.
- Recomendacao preliminar: revisar manualmente contra as correcoes de P42 antes de incorporar.

### dorigo2004book

- `summaries/` e mais detalhado em teoria ACO, resultados/conteudos e limitacoes.
- `vault/papers/` e mais curto e integrado ao vault.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary inclui autores/ano; titulo canonico e o titulo do livro.
- Validacao no PDF: recomendada para citacoes conceituais fortes.
- Recomendacao preliminar: incorporar definicoes e limitacoes teoricas que ajudem a fundamentacao.

### dorigo2005acotheory

- `summaries/` e mais detalhado em teoria, resultados, forcas/limitacoes e takeaway.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada se usado para claims teoricos sobre convergencia.
- Recomendacao preliminar: incorporar com foco em fundamentos, nao em resultados empiricos do TCC.

### freitas2020vns

- `summaries/` e mais detalhado em metodo VNS, resultados, limitacoes e takeaway.
- `vault/papers/` tem estrutura canonica, mas menos conteudo.
- Campos faltantes em `vault/papers/`: `pdf_status`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de afirmar resultados de FSTSP.
- Recomendacao preliminar: incorporar apenas como contexto de drone routing, se citado.

### garey1979computers

- `summaries/` e parecido em extensao e destaca resultados em secao propria.
- `vault/papers/` ja contem estrutura canonica.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: nao obrigatoria para conteudo conceitual basico; recomendada para formulacoes formais.
- Recomendacao preliminar: manter nota canonica e adicionar apenas detalhes de complexidade que forem citados.

### goldberg1989genetic

- `summaries/` e mais detalhado em GA, fundamentos, limitacoes e takeaway.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary inclui autor/ano; titulo canonico esta limpo.
- Validacao no PDF: recomendada para citacoes especificas de schema theorem ou operadores.
- Recomendacao preliminar: incorporar apenas fundamentos diretamente usados na fundamentacao de GA.

### gpaco2025

- `summaries/` e mais detalhado em GP-ACO, resultados, limitacoes e takeaway.
- `vault/papers/` e mais sintetico.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary usa titulo abreviado com `GP-ACO 2025`; canonico tem titulo completo.
- Validacao no PDF: recomendada antes de incorporar claims de desempenho.
- Recomendacao preliminar: manter como referencia periferica sobre variantes modernas de ACO.

### haroun2015performance

- `summaries/` e mais detalhado em comparacao GA/ACO, resultados, limitacoes e takeaway.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada antes de usar comparacoes como evidencia.
- Recomendacao preliminar: incorporar limitacoes metodologicas para evitar claim comparativo forte demais.

### heldkarp1970traveling

- `summaries/` e mais detalhado em 1-tree, resultados, limitacoes e takeaway.
- `vault/papers/` ja contem conteudo relevante e conexoes com lower bounds.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary inclui autores/ano; titulo canonico esta limpo.
- Validacao no PDF: recomendada para detalhes matematicos do bound.
- Recomendacao preliminar: incorporar apenas se complementar a distincao Held-Karp 1970/1971 ja corrigida.

### heldkarp1971traveling

- `vault/papers/` e mais completo em extensao e ja contem alerta importante sobre Held-Karp 1962 versus 1971.
- `summaries/` ainda acrescenta forcas/limitacoes, resultados em secao propria e takeaway.
- Campos faltantes em `vault/papers/`: `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: `summaries/` marca DOI como `[to be verified]`, enquanto a nota canonica ja contem `10.1007/BF01584070`.
- Validacao no PDF: recomendada apenas para detalhes nao cobertos pela nota canonica; DOI canonico parece preferivel.
- Recomendacao preliminar: nao substituir a nota canonica; migrar apenas limitacoes se uteis.
- Decisao P46: nota canonica preservada; DOI `10.1007/BF01584070` mantido; apenas campos transversais e nota de decisao foram adicionados.

### holland1975adaptation

- `summaries/` e muito mais detalhado em fundamentos de GA, limitacoes e takeaway.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary referencia `1975/1992`; titulo canonico mantem apenas o titulo da obra.
- Validacao no PDF: recomendada para distinguir edicao original e edicao republicada.
- Recomendacao preliminar: incorporar somente fundamento conceitual necessario para GA.

### hossain2024comparison

- `summaries/` e mais detalhado em comparacao de algoritmos, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: a diferenca de titulo parece ser apenas quebra de linha YAML na nota canonica; o titulo completo continua em `Benchmark Instances`.
- Validacao no PDF: recomendada antes de usar resultados comparativos, nao por truncamento aparente do titulo.
- Recomendacao preliminar: nao corrigir titulo sem necessidade; decidir se a comparacao e relevante para o TCC e validar resultados antes de citar.

### johnson1996asymptotic

- `summaries/` e mais detalhado em resultados, limitacoes e takeaway sobre o bound Held-Karp.
- `vault/papers/` e sintetico, mas alinhado ao uso no TCC.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para qualquer percentual ou resultado assintotico citado.
- Recomendacao preliminar: incorporar resultados-chave com referencia de pagina se usados na monografia.

### karp1979patching

- `summaries/` e mais detalhado em patching, resultados e implicacoes.
- `vault/papers/` ja tem estrutura canonica.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada se a nota for usada na secao de lower bounds.
- Recomendacao preliminar: incorporar apenas se lower bounds forem expandidos alem do bound Held-Karp/AP.

### kennedy1995particle

- `summaries/` tem extensao parecida, mas organiza melhor resultados, limitacoes e takeaway.
- `vault/papers/` ja e suficiente para definicao basica de PSO.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary inclui autores/ano no heading; titulo canonico esta limpo.
- Validacao no PDF: recomendada para citacoes historicas especificas.
- Recomendacao preliminar: incorporar somente limitacoes e contexto historico, se necessario.

### kinable2017hybrid

- `summaries/` e mais detalhado em problemas dependentes do tempo, metodo, resultados e limitacoes.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: heading do summary usa chave + `Summary`; titulo canonico e bibliografico.
- Validacao no PDF: recomendada antes de usar qualquer detalhe metodologico.
- Recomendacao preliminar: incorporar apenas conexao com sequenciamento dependente de historico/tempo.

### lin1973effective

- `summaries/` tem extensao proxima e destaca forcas/limitacoes, resultados e takeaway.
- `vault/papers/` ja e relativamente completo.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary usa autores/ano no heading; titulo canonico esta limpo.
- Validacao no PDF: recomendada para detalhes de Lin-Kernighan.
- Recomendacao preliminar: incorporar apenas detalhes que ajudem a contextualizar heuristicas classicas.

### lysgaard1999cluster

- `summaries/` contem praticamente todo o corpo analitico disponivel.
- `vault/papers/` esta sem corpo estruturado relevante.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: necessaria antes de incorporacao substancial.
- Recomendacao preliminar: prioridade alta em P46 se a referencia continuar no escopo de ATSP/lower bounds.
- Decisao P46: incorporado seletivamente como referencia auxiliar de ATSP/lower bounds; resultados quantitativos do summary continuam condicionados a validacao antes de citacao formal.

### murray2015flying

- `summaries/` e mais detalhado em FSTSP, resultados, limitacoes e takeaway.
- `vault/papers/` ja contem estrutura canonica.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para resultados e formulacao FSTSP.
- Recomendacao preliminar: incorporar contexto de drone-assisted TSP se usado na introducao/fundamentacao.

### muthanna2022uav

- `vault/papers/` e mais completo e ja reflete a reclassificacao recente de baixa relevancia para TSP/bio-inspired.
- `summaries/` ainda acrescenta forcas/limitacoes e takeaway.
- Campos faltantes em `vault/papers/`: `chapters`, `claim_support`, `aliases`.
- Divergencia inventariada em P45: o summary inclui autores/ano no heading e registra DOI `10.1016/j.comcom.2022.04.029`, enquanto a nota canonica registrava `10.1016/j.comcom.2022.04.028` antes da correcao P46.
- Validacao no PDF: DOI corrigido para `10.1016/j.comcom.2022.04.029`, confirmado; incorporacao de conteudo permanece periferica.
- **P44 concluida**: [[muthanna2022uav]] mantido como referencia contextual em Fundamentacao (cenario UAV em IoT/emergencia); nao usar para claims sobre TSP/bio-inspired.

### nagata2006eax

- `vault/papers/` e ligeiramente mais completo e ja recebeu correcao recente sobre ano/titulo.
- `summaries/` acrescenta forcas/limitacoes, resultados e takeaway.
- Campos faltantes em `vault/papers/`: `areas`, `chapters`, `claim_support`.
- Divergencia: summary usa `Nagata & Kobayashi (2013)` e `TSP`; nota canonica usa titulo completo. Esta diferenca precisa ser lida junto da correcao bibliografica recente.
- Validacao no PDF: recomendada antes de qualquer alteracao, para nao reintroduzir erro de ano/chave.
- Recomendacao preliminar: preservar nota canonica; migrar apenas limitacoes/resultados apos checagem.

### neufaco2025

- `summaries/` e mais detalhado em NeuFACO, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para resultados recentes.
- Recomendacao preliminar: manter como referencia periferica de ACO neural, sem peso central na monografia.

### pop2024comprehensive

- `summaries/` e mais detalhado em survey de GTSP, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: heading do summary usa apenas a chave; titulo canonico esta completo.
- Validacao no PDF: nao necessaria para titulo; recomendada para taxonomia ou claims de survey.
- Recomendacao preliminar: incorporar apenas se GTSP for citado como variante relacionada.

### potvin1996ga

- `summaries/` e mais detalhado e separa resultados em secao propria.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para citacoes sobre operadores geneticos.
- Recomendacao preliminar: incorporar detalhes de representacao/operadores se usados em GA.

### rajan2022routing

- `summaries/` e mais detalhado em patrulhamento UAV, metodo, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: diferenca tipografica entre travessao e `--` no titulo.
- Validacao no PDF: recomendada para resultados e formulacao do problema.
- Recomendacao preliminar: incorporar com prioridade para contexto de patrulha UAV. Relacao com [[muthanna2022uav]]: P44 decidiu manter muthanna como referencia contextual (IoT/emergencia), sem substituicao. Ambos coexistem como referencias perifericas independentes.

### rajwar2023exhaustive

- `summaries/` e mais detalhado em taxonomia, aplicacoes, desafios, limitacoes e takeaway.
- `vault/papers/` tem titulo completo em YAML quebrado em duas linhas; `summaries/` e mais detalhado no corpo.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: a diferenca de titulo parece ser apenas quebra de linha YAML; nao ha evidencia de truncamento real na nota canonica.
- Validacao no PDF: recomendada para claims de survey e taxonomia, nao por truncamento aparente do titulo.
- Recomendacao preliminar: nao corrigir titulo sem necessidade; decidir se o survey e central ou apenas auxiliar.

### righini2021efficient

- `summaries/` e mais detalhado em otimizacao do lower bound Held-Karp, resultados, limitacoes e takeaway.
- `vault/papers/` e sintetico.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para resultados computacionais e detalhes do algoritmo.
- Recomendacao preliminar: incorporar se a monografia discutir implementacao/estimativa de lower bounds.

### stutzle2000mmas

- `summaries/` e mais detalhado em MMAS, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: recomendada para diferenciar MMAS de Ant System/ACO implementado.
- Recomendacao preliminar: incorporar apenas como variante relacionada, sem confundir com a implementacao atual.

### valenzuela1997estimating

- `vault/papers/` e mais completo e ja possui campos minimos, inclusive secoes de uso, evidencia e limitacoes.
- `summaries/` acrescenta principalmente um takeaway pratico.
- Campos faltantes em `vault/papers/`: nenhum campo minimo faltante detectado.
- Divergencia: nenhuma divergencia automatica relevante detectada.
- Validacao no PDF: nao obrigatoria antes de P46, exceto para citacoes numericas.
- Recomendacao preliminar: baixa prioridade; manter nota canonica e migrar apenas takeaway se util.

### vanhove2012route

- `summaries/` e muito mais detalhado e separa resultados em secao propria.
- `vault/papers/` e breve.
- Campos faltantes em `vault/papers/`: `pdf_status`, `chapters`, `claim_support`, `aliases`.
- Divergencia: heading do summary usa apenas a chave; titulo canonico esta completo.
- Validacao no PDF: recomendada se a referencia for usada para custos de curva/turn restrictions.
- Recomendacao preliminar: incorporar conteudo apenas se a monografia mantiver a discussao de penalidade angular.
- Decisao P46: incorporadas evidencias/limitacoes conceituais sobre custos de conversao; limiares percentuais permanecem condicionados a validacao no PDF antes de citacao formal.

### wadi2025charting

- `summaries/` e mais detalhado em comparacao de abordagens swarm, resultados, limitacoes e takeaway.
- `vault/papers/` e mais curto.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: heading do summary usa chave + `Summary`; titulo canonico esta completo.
- Validacao no PDF: recomendada para qualquer claim comparativo.
- Recomendacao preliminar: incorporar apenas se o artigo for usado como survey/apoio comparativo recente.

### wang2021ant

- `summaries/` e muito mais detalhado em SOS-ACO/parametrizacao, resultados, limitacoes e takeaway.
- `vault/papers/` e breve, apesar de estar citado para tuning/ACO.
- Campos faltantes em `vault/papers/`: `pdf_status`, `areas`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary usa titulo abreviado `TSP`; canonico explicita `Traveling Salesman Problem`.
- Validacao no PDF: recomendada antes de usar resultados ou parametros.
- Recomendacao preliminar: prioridade media em P46 por relacao com tuning e P33.
- Decisao P46: incorporado como apoio a tuning de hiperparametros/ACO; resultados quantitativos nao foram promovidos a claim forte.

### winter2002modeling

- `summaries/` tem extensao proxima e destaca resultados em secao propria.
- `vault/papers/` ja e relativamente completo para contexto de custos de curva.
- Campos faltantes em `vault/papers/`: `pdf_status`, `methods`, `chapters`, `claim_support`, `aliases`.
- Divergencia: summary acrescenta `Resumo` ao heading; titulo canonico esta limpo.
- Validacao no PDF: recomendada apenas para detalhes tecnicos do grafo pseudo-dual.
- Recomendacao preliminar: incorporar somente se ajudar a explicar penalidade angular/turn costs.

## Recomendacao final

Para P46, nao fazer incorporacao em massa. A abordagem mais segura e tratar cada nota por decisao explicita: incorporar, manter como auxiliar, validar no PDF, corrigir metadado ou descartar. As tres notas canonicas sem corpo devem vir primeiro; em seguida, corrigir metadados truncados e so depois enriquecer notas ja completas com limitacoes, resultados e takeaways.
