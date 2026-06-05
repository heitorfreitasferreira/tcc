---
title: Sessão de Grille — Banca Simulada
tags:
- tipo/writing
- topico/monografia
- topico/revisao
created: 2026-06-02
updated: 2026-06-02
aliases:
- entrevista
type: writing
---

# Sessão de Grille — Banca Simulada

Registro completo das perguntas, respostas brutas e expansões da sessão de arguição simulada sobre a monografia de TCC. Cada entrada documenta o que foi dito e o que pode ser extraído para o texto final.

Ver também: [[roadmap-monografia]] — as ações desta sessão foram registradas como tarefas P4, P8, P9, P10.

---

## Pergunta 1 — Definição do Problema

**Banca:** Seu trabalho compara GA, PSO e ACO em um cenário de patrulha com VANTs. Qual é exatamente a função objetivo que você está minimizando? E por que ela não pode ser reduzida a um TSP clássico?

**Resposta bruta:**

> estou minimizando o tempo de patrulha, o tsp classico depende apenas de onde voce esta e onde esta indo, mas como drones se movimentam num fluido como barcos, eh ideal que se mantenha fazendo menor curva possivel, pra isso foi modelado o tensor que calcula o tempo do trajeto B->C com base onde voce estava antes de B, entao A1->B->C tem um custo, e A2->B->C tem outro, a depender do angulo (ir reto tem menor custo, dar meia volta tem custo maximo)

**Expansão para a monografia:**

A função objetivo é o **makespan** — tempo total para percorrer uma rota completa de patrulha visitando todos os pontos de interesse e retornando ao depósito. Diferentemente do TSP clássico, onde o custo de uma aresta `(u,v)` depende apenas dos nós `u` e `v`, no cenário com drones o custo de `curr→next` também depende do nó anterior (`prev`), porque o ângulo de curvatura entre os segmentos `prev→curr` e `curr→next` afeta o tempo de voo. Drones, assim como embarcações, têm inércia direcional: manter trajetória reta é mais eficiente que fazer curvas fechadas. Para capturar essa dependência, foi modelado um tensor tridimensional `G[prev][curr][next]` onde cada entrada armazena o tempo total (distância + penalidade angular) da transição `curr→next` dado que se veio de `prev`. Isso torna o problema irredutível ao TSP clássico, que opera sobre uma matriz `N×N`.

---

## Pergunta 2 — Representação do Problema

**Banca:** Você modela isso como um tensor `G[prev][curr][next]` de dimensão `N³`. Isso significa que o custo de uma transição `curr→next` depende do nó anterior. Como você garante que uma rota factível (sem nós repetidos) pode ser representada nessa estrutura? Em outras palavras, dado que o tensor tem 3 índices, como você codifica a restrição de que cada nó é visitado exatamente uma vez e a rota começa e termina no depósito?

**Resposta bruta:**

> a ideia nao eh nao ter nos repetidos, e sim passar por todos os nos no menor tempo, o tensor so modela todas as possibilidades, as metaheuristicas que geram o menor caminho, que sendo o menor caminho mesmo, tende a nao repetir nos

**Expansão para a monografia:**

O tensor `G[prev][curr][next]` não impõe restrições de factibilidade por si só — ele é apenas uma estrutura de dados que pré-computa o custo de cada triple possível. A restrição de visitar cada nó exatamente uma vez (tour Hamiltoniano) é imposta pelo algoritmo de otimização, não pelo tensor. O brute-force, GA, PSO e ACO constroem ou avaliam candidatos que são permutações dos nós (com início e fim no depósito), e o tensor é consultado para calcular o custo de cada passo da sequência. A tendência natural de um tour ótimo ou quase-ótimo é não repetir nós, pois repetições aumentam o makespan sem necessidade. O tensor unicamente fornece o custo realista de cada transição, considerando a penalidade angular; a factibilidade é responsabilidade do método de busca.

---

## Pergunta 3 — Métricas de Avaliação

**Banca:** Você usa o *makespan* como métrica única de qualidade. Três métodos estocásticos com 51 sementes cada, testados em 30 instâncias — isso gera 4590 amostras. Como você *compara* formalmente os métodos? Apenas pela melhor mediana? Isso é suficiente para afirmar "GA foi superior"?

**Resposta bruta:**

> sendo sincero nao sei, acreditava que apenas o makespan indicaria

**Expansão para a monografia (extraída do protocolo estatístico):**

Comparar apenas medianas não é suficiente para afirmar superioridade entre métodos estocásticos. A variabilidade entre sementes precisa ser considerada. O protocolo estatístico adotado segue Demšar (2006):

1. **Teste de Friedman**: rejeita a hipótese nula de que todos os métodos produzem a mesma distribuição de makespan. Aplica-se sobre as medianas de cada método em cada instância.
2. **Pós-teste de Nemenyi**: se Friedman rejeitar H₀, identifica quais pares de métodos (GA×PSO, GA×ACO, PSO×ACO) diferem com significância estatística.
3. **Diagrama de diferença crítica (CD)**: visualização gráfica dos resultados do pós-teste.
4. **Wilcoxon signed-rank**: alternativa pareada, instância por instância, para cada par de métodos, como validação adicional.

Este protocolo foi registrado no roadmap de escrita e será implementado em `scripts/analise-estatistica.py`. Conclusões sobre ordenação de métodos só serão dadas após a execução destes testes.

---

## Pergunta 4 — Instâncias e Generalização

**Banca:** Você usou 30 instâncias sintéticas com coordenadas uniformes em `[-1,1]`. Nenhuma delas reflete um cenário real de patrulha — sem clusters, sem obstáculos, sem geografia real. Como você defende que seus resultados (qualquer que seja a ordenação) se mantêm em um cenário operacional real?

**Resposta bruta:**

> a premissa eh rodar em cenarios sinteticos (tipo de tsp escolhido), fica como trabalho futuro adaptar para cenarios e mapas reais

**Expansão para a monografia:**

As 30 instâncias utilizadas seguem o formato clássico de benchmarks sintéticos para TSP: coordenadas uniformemente distribuídas em `[-1,1]`. Esta escolha é voluntária e declarada: o objetivo do trabalho é comparar metaheurísticas em uma variante pouco explorada do TSP (TSP-SD-ATP), não resolver instâncias reais de patrulha. A generalização para cenários operacionais — como mapas geográficos reais, pontos de interesse clusterizados, ou restrições de espaço aéreo — é deixada como trabalho futuro. Esta limitação está explicitada na discussão e na conclusão, de forma que o leitor não extrapole os achados além do desenho experimental executado.

---

## Pergunta 5 — Reprodutibilidade e Pipeline

**Banca:** Seu pipeline depende de um `tcc serve` para gerar figuras via `/api/render`, e os dados ficam embutidos no binário via `//go:embed`. Isso significa que um revisor precisa compilar seu código para verificar as figuras? E os dados brutos — estão acessíveis em formato padrão para análises independentes?

**Resposta bruta:**

> a ideia do trabalho eh disponibilizar o binario standalone, os dados e todo o repositorio que estamos ficara opensource tbm

**Expansão para a monografia:**

O repositório completo será disponibilizado publicamente sob licença open-source. O binário compilado (`tcc`) é standalone e pode ser baixado e executado sem necessidade de instalação de dependências — ele serve dados, renderiza figuras e executa otimizações. Os dados brutos de resultados (sumário, evolução e tempo) estão em formato JSON no diretório `src/data/results/`, acessíveis e parseáveis por qualquer linguagem. Um revisor pode tanto executar o binário pré-compilado para reproduzir figuras, quanto ler diretamente os arquivos JSON para análises independentes. O código-fonte em Go também está disponível para inspeção e compilação.

---

## Pergunta 6 — ACO e o Tensor 3D

**Banca:** Seu ACO tem um tensor de feromônio `τ_{ijk}` com `O(N³)` entradas. Para `n=100`, são ~970 mil valores de feromônio contra ~9.900 de um ACO 2D clássico. A relação sinal-ruído é ~98× pior. Isso não significa que, estruturalmente, o ACO está em desvantagem contra GA e PSO nesta variante — não por culpa do algoritmo, mas da representação? Como você separa na sua conclusão o que é limitação do ACO do que é limitação da dimensionalidade do feromônio?

**Resposta bruta:**

> nesse caso realmente aco tem uma desvantagem clara, mas se fez necessario e ocorreu por conta da natureza do trabalho, de considerar o passo anterior para ter o angulo do trajeto

**Expansão para a monografia:**

Esta observação é correta e importante. O ACO tridimensional sofre de uma diluição estrutural do feromônio: o tensor `τ_{ijk}` tem `N × (N-1) × (N-2)` entradas válidas (~970 mil para n=100), enquanto um ACO 2D clássico opera com `(N-1) × (N-2)` entradas (~9.900). Com o mesmo número de formigas depositando feromônio, cada triple recebe muito menos reforço por iteração, o que degrada a relação sinal-ruído e pode prejudicar a convergência. Esta não é uma limitação do ACO como metaheurística, mas uma consequência direta da representação tridimensional que o problema TSP-SD-ATP exige — não há como representar a penalidade angular dependente de sequência em uma matriz `τ_{ij}` sem perder informação. A discussão dos resultados deve deixar explícito que a desvantagem observada do ACO pode ser, ao menos em parte, um artefato da dimensionalidade, e não uma inferioridade algorítmica. Estratégias como MMAS (limites `[τ_min, τ_max]`), listas candidatas, ou uma formulação 2D alternativa com fatoração do turn cost na heurística são mencionadas como trabalhos futuros.

---

## Pergunta 7 — Baseline Brute-Force

**Banca:** Você executou brute-force só até `15c`. Para instâncias maiores, você não tem o ótimo global. Como você mede a qualidade absoluta das soluções encontradas por GA, PSO e ACO em instâncias de 30, 50 e 100 nós? Gap relativo entre os métodos é suficiente, ou você precisa de um lower bound para dizer "estamos a X% do ótimo"?

**Resposta bruta:**

> colocar no roadmap investigar implementacao de metodos que deem um lower bound

**Expansão para a monografia:**

Sem o ótimo global para instâncias grandes, a comparação entre métodos é apenas relativa. Para fornecer uma referência absoluta, foi investigada a viabilidade de implementar um lower bound. A análise completa está em [[lower-bounds]], com as referências fichadas no vault. As principais abordagens identificadas são:

1. **Redução 3D→2D** via `c'[j][k] = min_i cost[i][j][k]`, seguida de relaxação Assignment Problem (Hungarian, O(n³)) ou Held-Karp (1-tree + subgradiente)
2. **MDD com largura limitada** — abordagem nativa 3D, sem redução, usando Diagramas de Decisão Multivalorados ([[kinable2017hybrid]])
3. **ng-path relaxation** — relaxação de estado-espaço para labeling ([[leraromero2020dynamic]])

A abordagem via redução + AP é a mais prática para implementar em Go puro. Esta investigação foi registrada como tarefa P8 no roadmap de escrita.

---

## Pergunta 8 — Significância Prática vs. Estatística

**Banca:** Suponha que o teste de Friedman aponte diferença significativa entre GA e ACO, mas a diferença absoluta de makespan seja de 0.5% em média. Isso é estatisticamente significativo, mas é praticamente relevante para um drone em patrulha? Um operador real trocaria de algoritmo por 0.5%? Onde você traça a linha entre "diferente" e "melhor de verdade"?

**Resposta bruta:**

> considerando que o objetivo final eh diminuir tempo, essa margem muito pequena acaba sendo engolida por tempos fixos do trajeto, como tempo de checks de seguranca, abastecimetno, entao nao devem fazer tanta diferenca pratica no dia a dia, mas a diferenca obviamente chega a fazer dada a escala

**Expansão para a monografia:**

A distinção entre significância estatística e prática é crucial. Uma diferença de 0.5% no makespan pode ser estatisticamente significativa com 51 sementes, mas sua relevância operacional depende do contexto. Em um cenário de patrulha, o makespan teórico é apenas parte do tempo real de missão — tempos fixos como verificações de segurança, abastecimento, decolagem e pouso podem dominar o custo total, tornando uma diferença de 0.5% irrelevante na prática. No entanto, em operações de larga escala (frotas grandes, múltiplas missões por dia), essa fração se acumula e pode representar economia real. A discussão dos resultados deve incluir esta nuance: reportar a magnitude absoluta das diferenças junto com a significância estatística, para que o leitor possa julgar a relevância prática por si mesmo.

---

## Pergunta 9 — Representação e Codificação das Soluções

**Banca:** GA usa permutação direta, PSO usa codificação contínua por random keys, e ACO constrói rotas por transições probabilísticas. Cada método representa uma rota de forma diferente. Como você garante que está comparando algoritmos e não efeito colateral da codificação?

**Resposta bruta:**

> nao sei dizer (coloque no roadmap para estudar mais sobre, e possivelmente ajustar as implementacoes)

**Expansão para a monografia:**

Esta é uma ameaça à validade da comparação que precisa ser reconhecida. Cada método usa uma codificação diferente para representar rotas:

- **GA**: permutação direta dos nós `1..n-1`, com OX e mutação swap
- **PSO**: random keys/SPV, isto é, posição contínua convertida em permutação por ranking
- **ACO**: transições probabilísticas condicionadas a triplas `prev,curr,next`

Não é possível afirmar que a codificação não favorece ou desfavorece um método específico. A auditoria P9 foi registrada em [[auditoria-codificacao-metodos]]: os resultados devem ser apresentados como comparação entre implementações concretas, sob parâmetros e orçamento comuns, e não como prova de superioridade universal de uma metaheurística. Variantes alternativas, como PSO discreto com operadores de permutação ou GA com operadores mais orientados a arestas, ficam como trabalho futuro.

---

## Pergunta 10 — Sensibilidade a Hiperparâmetros

**Banca:** GA, PSO e ACO foram executados com um único conjunto de hiperparâmetros cada (população=100, iterações=100, e os específicos de cada método). Não houve tuning. Numa banca real, alguém vai apontar que você pode estar comparando um ACO mal configurado contra um GA razoável, e a ordenação final seria um artefato dos parâmetros escolhidos, não dos algoritmos em si. O que você responde?

**Resposta bruta:**

> isso entra em trabalhos futuros, uma analise de quais metodologias de finetuning se adequam a cada modelo nesse cenario, atualize o vault com essa informacao para trabalhos futuros

**Expansão para a monografia:**

A ausência de tuning sistemático é uma limitação reconhecida. Todos os métodos foram executados com parâmetros fixos (população=100, iterações=100, e parâmetros específicos como ρ=0.2 para ACO, w=0.7/c1=c2=2 para PSO, elitism=1/mutation_rate=0.05 para GA). A ordenação observada pode ser um artefato desta escolha única de parâmetros. Fica como trabalho futuro uma análise de sensibilidade — utilizando grid search, random search ou otimização Bayesiana — para verificar se a ordenação entre métodos se mantém sob diferentes parametrizações. Esta investigação foi registrada como tarefa P10 no roadmap.

---

## Ações Registradas no Roadmap

As seguintes tarefas foram criadas ou atualizadas como resultado desta sessão:

| ID | Tarefa | Prioridade |
|----|--------|------------|
| P4 | Fechar protocolo estatístico (Friedman, Nemenyi, Wilcoxon, CD) | Alta |
| P8 | Investigar lower bound para instâncias grandes (n ≥ 15) | Média |
| P9 | Estudar efeito da codificação na comparação justa entre métodos | Média |
| P10 | Análise de sensibilidade a hiperparâmetros | Baixa (trabalho futuro) |

> Tarefas registradas em [[roadmap-monografia#Tarefas Preparatórias Para Agentes]]
