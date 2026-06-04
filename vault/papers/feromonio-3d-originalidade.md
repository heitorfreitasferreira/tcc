---
title: "Verificação de Originalidade — Feromônio 3D τ(i,j,k) para ACO no TSP-SD-ATP"
tags: [aco, tsp, feromonio-3d, originalidade, verificacao]
status: verificado-externo-encontrou-precedentes
created: 2026-06-05
updated: 2026-06-05
---

## Pergunta de Pesquisa

A representação de feromônio tridimensional $\tau(i,j,k)$ — em que a intensidade da trilha depende da tripla `(nó_anterior, nó_atual, próximo_nó)` — é uma contribuição algorítmica original deste TCC, ou há precedentes na literatura de ACO com feromônio n-dimensional (n > 2)?

## Metodologia de Busca

### Fontes consultadas

| Fonte | Escopo | Status |
|-------|--------|--------|
| Vault interno (84 papers) | Leitura completa de todas as notas `.md` com busca por padrões de feromônio não-2D | Concluído |
| Google Scholar | 4 consultas com termos específicos | Bloqueado por CAPTCHA |
| Semantic Scholar API | 3 consultas | Rate-limited (429) |
| DBLP API | 3 consultas | Indisponível (503) |
| arXiv search (MCP) | 3 consultas | Timeout |
| Sci-Hub | 4 consultas por keyword + 1 por título | Sem resultados relevantes |

### Termos buscados

1. `"three dimensional pheromone" ant colony`
2. `"pheromone tensor" ACO`
3. `ant colony "sequence dependent" TSP`
4. `ACO hypergraph pheromone "higher-order"`
5. `"turn cost" ant colony TSP`
6. `ant colony pheromone triple transition tsp`
7. `sequence-dependent disassembly line balancing ant colony`

## Evidência Interna (Vault — 84 papers)

### ACO canônico: apenas feromônio 2D

Todos os papers de ACO no vault descrevem exclusivamente feromônio baseado em arestas $\tau(i,j)$:

| Paper | Variante | Representação |
|-------|----------|---------------|
| [[dorigo1996ant]] | Ant System | $\tau_{ij}$ (arestas) |
| [[dorigo1997ant]] | ACS | $\tau(r,s)$ (arestas) |
| [[dorigo2004book]] | AS, EAS, RAS, MMAS, ACS | $\tau_{ij}$ (arestas), cobre todos os capítulos |
| [[dorigo2005acotheory]] | Teoria abstrata | "Componentes de solução" — instanciado como arestas no TSP |
| [[dorigo2018acooverview]] | AS, ACS, MMAS | $\tau_{ij}$ (arestas), descrição canônica |
| [[blum2005acointro]] | Introdução | Arestas |
| [[stutzle2000mmas]] | MMAS | $\tau_{ij}$ com bounds $[\tau_{min}, \tau_{max}]$ |
| [[gpaco2025]] | GP-ACO | Arestas (terminais GP: feromônio + distância) |
| [[deepaco2023]] | DeepACO | Arestas |
| [[neufaco2025]] | NeuFACO | Arestas |
| [[ppaco2024]] | Policy Gradient ACO | Arestas |
| [[misra2024acorecent]] | Survey | Arestas (todas as variantes) |

**Resultado**: zero menções a feromônio 3D, n-dimensional, ou baseado em triplas em qualquer um dos 84 papers.

### Problemas relacionados sem ACO

Problemas com dependência de sequência existem na literatura, mas são resolvidos por outros métodos:

| Problema | Paper | Método | Usa ACO? | Usa 3D? |
|----------|-------|--------|----------|---------|
| SDTSP (Sequence-Dependent TSP) | [[kinable2017hybrid]] | MDDs + LP | Não | Não |
| Turn-cost routing | [[winter2002modeling]] | Pseudo-dual graph (arestas→nós) | Não | Não |
| AM-TSP (Angular-Metric TSP) | [[aggarwal2000angular]] | Não especificado no vault | Não | Não |
| Route planning com turn restrictions | [[vanhove2012route]] | Shortest path com grafo expandido | Não | Não |

### Abstração teórica que poderia suportar 3D

[[dorigo2005acotheory]] formaliza ACO com "componentes de solução" abstratos que recebem feromônio. O texto da nota menciona: "arestas ou decisões de próximo nó podem receber feromônio". A abstração **permite** triplas como componentes, mas **nenhum artigo a instancia dessa forma**. A instanciação canônica para TSP sempre usa arestas como componentes.

## Áreas de Risco (Prior Art Potencial Não Coberto)

### 1. ACO para scheduling com 3 índices

Problemas de scheduling (job shop, flow shop) frequentemente usam formulações com 3 índices (máquina × job × posição). O dorigo2004book lista scheduling como domínio de aplicação. Se algum paper de ACO para scheduling usou feromônio $\tau(m, j, p)$ para representar "máquina m processa job j na posição p", isso seria um precedente estrutural para $\tau(i,j,k)$.

**Status**: **Não investigado.** Esta avenida foi identificada como risco real mas não foi coberta pelas buscas realizadas. As APIs de busca retornaram erros de rate-limiting e as consultas usaram termos TSP-cêntricos. Fica como limitação conhecida desta verificação.

### 2. ACO para VRP com time windows

VRPTW pode usar estados expandidos (cliente × tempo) que resultam em matrizes 3D ou tensores de feromônio.

**Ação**: buscar `"ant colony" vehicle routing "pheromone" three dimensional`.

### 3. ACO para TDTSP (Time-Dependent TSP)

O TDTSP tem custos $c(i,j,t)$ onde o tempo de partida afeta o custo da aresta. Papers de ACO para TDTSP poderiam usar feromônio indexado por tempo: $\tau(i,j,t)$.

**Ação**: buscar `ant colony time-dependent TSP pheromone`.

### 4. Hyper-heuristics com ACO

ACO aplicado como hyper-heuristic pode usar feromônio sobre sequências de heurísticas (n-gramas).

**Ação**: buscar `ant colony hyper-heuristic "pheromone" sequence`.

## Avaliação de Confiança

| Dimensão | Evidência | Confiança |
|----------|-----------|-----------|
| Literatura ACO canônica (84 papers) | Zero menções a feromônio 3D | **Alta** |
| Problemas sequence-dependent resolvidos com ACO | Nenhum encontrado no vault | **Alta** |
| Scheduling com ACO 3D | Não verificado (fora do escopo do vault) | **Baixa** — risco real |
| TDTSP com ACO | Não verificado | **Média** |
| VRP com estados expandidos | Não verificado | **Média** |

## Veredito Preliminar

**Dentro da interseção ACO ∩ TSP**: a evidência disponível (84 papers, todas as variantes canônicas AS/ACS/MMAS, surveys recentes até 2025) indica que $\tau(i,j,k)$ é **inédito**. A literatura canônica de ACO para TSP usa exclusivamente feromônio 2D baseado em arestas.

## Evidência Externa (OpenAlex — 6.993 resultados, 10 analisados + deep search)

### Precedentes de feromônio multidimensional em ACO

A busca externa encontrou **precedentes relevantes** de feromônio com dimensionalidade >2 em ACO, embora em domínios diferentes do TSP:

| Paper | Ano | Domínio | DOI | Relevância |
|-------|-----|---------|-----|------------|
| Wang et al. — CPACO | 2013 | Task allocation multi-agente | `10.1109/cc.2013.6488841` | Usa "three-dimensional path pheromone storage space" — 3D explícito |
| Wang et al. — Multi-dim pheromone | 2015 | Network coding (multicast) | `10.1109/tevc.2015.2457437` | "multi-dimensional pheromone maintenance mechanism" no IEEE TEVC |
| MDACO | 2025 | Multi-compartment VRP | `10.1109/ccdc65474.2025.11090784` | **Mais próximo**: "multi-dimensional pheromone matrix" com positional information |
| Starzec & Starzec | 2026 | Multi-objective TSP | `10.58032/agh/3gokaq` | "multi-dimensional pheromone" como modificação para MOTSP |

### Falsos positivos

Diversos papers usam "3D" no sentido de espaço físico 3D (path planning para UAVs, robôs, navios) com feromônio 2D padrão. Outros usam "tensor" no sentido de GPU (TensorACO 2024 — tensorização para aceleração, não feromônio extra-dimensional).

### Análise de similaridade

| Dimensão | CPACO 2013 | Wang 2015 | MDACO 2025 | τ(i,j,k) TCC |
|----------|------------|-----------|------------|--------------|
| Dimensionalidade | 3D (agent×task×resource) | Multi-dim (por receiver) | Multi-dim (posicional) | 3D (prev×curr×next) |
| Problema | Task allocation | Network coding | VRP | TSP |
| Semântica | Alocação de recursos | Sobreposição de rotas multicast | Ordem de visita em VRP | Transição de aresta com contexto |
| Markov? | Não | Não | Parcial (posicional) | Sim (2ª ordem) |

Nenhum dos precedentes aplica feromônio 3D a **transições de aresta com dependência de segunda ordem de Markov** (τ sobre triplas de nós consecutivos) no **TSP**. A semântica é diferente: os precedentes usam a terceira dimensão para separar contextos diferentes (agentes, receivers, posições), não para modelar uma dependência markoviana da função objetivo.

## Veredito Final (Após Evidência Externa)

O claim de originalidade **não pode ser absoluto**. Feromônio multidimensional em ACO **tem precedentes** (Wang 2013, Wang 2015, MDACO 2025, Starzec 2026). No entanto, $\tau(i,j,k)$ para TSP com transições dependentes de triplas (segunda ordem de Markov) é uma **instanciação específica sem precedente direto conhecido**. A contribuição está na **combinação** de:

1. Feromônio 3D $\tau(prev, curr, next)$ como adaptação do Ant System
2. Aplicado ao TSP (não VRP, scheduling, ou network coding)
3. Motivado pela estrutura de segunda ordem de Markov do TSP-SD-ATP
4. Com a mesma dimensionalidade do tensor de custo $G[prev][curr][next]$

## Recomendação para P32

O claim deve ser **qualificado**, não absoluto. Três camadas, da mais segura para a mais ousada:

1. **Claim seguro** (sempre verdadeiro): "No contexto de ACO para TSP, a literatura canônica usa exclusivamente feromônio $\tau(i,j)$ 2D baseado em arestas. Este trabalho adapta o Ant System para usar feromônio tridimensional $\tau(i,j,k)$, motivado pela dependência de triplas do TSP-SD-ATP."

2. **Claim qualificado** (requer qualificação): "Feromônio multidimensional em ACO tem precedentes em domínios como task allocation \cite{wang2013cpaco}, network coding \cite{wang2015multidim} e vehicle routing \cite{mdaco2025}, mas sua aplicação a transições de aresta com dependência de segunda ordem no TSP não foi encontrada na literatura consultada."

3. **Claim de diferenciação semântica** ( requer evidência forte): "Diferentemente dos usos anteriores de feromônio multidimensional (que separam contextos como agentes ou posições), o $\tau(i,j,k)$ deste trabalho modela diretamente a estrutura de segunda ordem de Markov da função objetivo do TSP-SD-ATP."

**Ação para P32**: escrever o parágrafo usando o claim #1 como base, com nota de rodapé ou menção breve aos precedentes (#2) para honestidade acadêmica. Não usar "inédito" sem qualificação. Preferir "adaptação não documentada na literatura de ACO para TSP".

## Verificação Externa

- [x] Buscar `ant colony scheduling three-index pheromone` — OpenAlex retornou CPACO 2013 (task allocation)
- [x] Buscar `ant colony time-dependent TSP pheromone` — não encontrado nos resultados
- [x] Buscar `ant colony three dimensional pheromone` no OpenAlex — 6.993 resultados, 4 precedentes identificados
- [x] Adicionar BibTeX e notas no vault para: wang2013aco, wang2015aco, geng2025mdaco, starzec2026motsp
- [x] Citar os precedentes no parágrafo do P32 (qualificação do claim)
- [x] Documentar limitação do gap de scheduling
- [x] Adicionar nota de rodapé sobre escopo e limitações da busca no LaTeX

## Conexões

- [[aco]] — implementação do ACO 3D no projeto
- [[problem-formulation]] — formulação do TSP-SD-ATP com tensor 3D
- [[dorigo1996ant]] — Ant System original (feromônio 2D)
- [[dorigo2005acotheory]] — abstração de componentes de solução
- [[kinable2017hybrid]] — SDTSP resolvido com MDDs (não ACO)
- [[claim-evidence-matrix]] — claims A07 e I01
- [[roadmap-monografia]] — P35 e P32
