# Auditoria P10 — Sensibilidade a Hiperparâmetros

Status: `concluida-como-limitacao`

## Objetivo

Registrar como a ausência de tuning sistemático afeta a interpretação dos resultados e definir uma metodologia de sensibilidade para trabalhos futuros, sem bloquear a escrita atual da monografia.

## Parâmetros Fixos Auditados

| Método | Parâmetros usados | Evidência primária | Observação |
|---|---|---|---|
| GA | população 100, iterações 100, elitismo 1, mutação 0.05, torneio 2 | `src/cmd/optimize.go:36-37`, `src/cmd/ga.go:64-66` | Configuração simples, sem ajuste por tamanho de instância. |
| PSO | população 100, iterações 100, `c1=2.0`, `c2=2.0`, `w=0.7` | `src/cmd/optimize.go:36-37`, `src/cmd/pso.go:68-70` | Usa parâmetros canônicos, mas sem *velocity clamping* ou inércia adaptativa. |
| ACO | população 100, iterações 100, `alpha=1.0`, `beta=2.0`, `rho=0.2`, `q=100` | `src/cmd/optimize.go:36-37`, `src/cmd/aco.go:66-69` | Sensível a evaporação, peso heurístico e escala de depósito. |

## Interpretação

Os resultados atuais comparam implementações sob um orçamento comum: mesma população, mesmo número de iterações e mesmas 51 sementes para os métodos estocásticos. Essa decisão controla parte do custo experimental, mas não garante que cada método opere em sua melhor região paramétrica. Portanto, a ordenação observada deve ser apresentada como resultado da configuração avaliada, não como conclusão definitiva sobre todas as parametrizações possíveis.

O risco é particularmente relevante para ACO e PSO. No ACO, `rho`, `alpha`, `beta` e `q` alteram diretamente o equilíbrio entre exploração, reforço de trilhas e uso da heurística local. No PSO, `c1`, `c2`, `w` e eventual limitação de velocidade controlam estabilidade e diversidade; no código atual, a posição contínua é decodificada por ranking, o que pode reduzir o efeito de pequenas atualizações em `x`. No GA, taxa de mutação, tamanho do torneio e elitismo influenciam pressão seletiva e diversidade.

## Metodologia Recomendada Para Trabalho Futuro

| Estratégia | Uso recomendado | Vantagem | Custo/risco |
|---|---|---|---|
| Grid search pequeno | Primeiro diagnóstico em poucas instâncias representativas (`10a`, `30a`, `50a`, `100a`) | Fácil de explicar e reproduzir | Cresce rapidamente com o número de parâmetros. |
| Random search | Exploração inicial de faixas amplas | Mais eficiente que grid quando poucos parâmetros importam | Exige orçamento de execuções maior e controle por semente. |
| irace/F-Race | Seleção automática de configurações por desempenho empírico | Adequado para algoritmos estocásticos | Ferramenta externa e desenho experimental mais complexo. |
| Otimização Bayesiana | Ajuste fino após definir faixas plausíveis | Boa eficiência amostral | Mais difícil de justificar em monografia de escopo aplicado. |

Para esta monografia, a alternativa mais defensável seria propor como trabalho futuro um desenho em duas fases: random search ou grid reduzido para estimar faixas promissoras; depois validação cruzada por grupos de instâncias, separando instâncias usadas para tuning e instâncias usadas para avaliação final. Esse cuidado evita escolher parâmetros que se ajustam demais às mesmas instâncias reportadas como resultado.

## Texto Recomendado Para Limitações

> Os métodos foram avaliados com parâmetros fixos e orçamento computacional comum, sem tuning sistemático por método ou por tamanho de instância. Essa escolha torna o experimento reprodutível e controla a comparação operacional, mas limita a generalização dos resultados. A ordenação observada entre ACO, GA e PSO pode mudar sob outras parametrizações, especialmente porque ACO e PSO são sensíveis a parâmetros de exploração, reforço e estabilidade. Assim, os resultados devem ser interpretados como evidência para as configurações implementadas neste estudo. Uma extensão natural é executar análise de sensibilidade ou ajuste automático de parâmetros, com instâncias separadas para tuning e validação.

## Decisão P10

P10 não bloqueia a escrita definitiva. A monografia deve declarar a ausência de tuning como limitação metodológica e tratar análise de sensibilidade como trabalho futuro. Não há evidência atual para afirmar que a ordenação ACO > GA > PSO se manteria sob qualquer parametrização.

## Links

- [[claim-evidence-matrix]]
- [[auditoria-codificacao-metodos]]
- [[analysis-methodology]]
- [[resultados]]
- [[ga]]
- [[pso]]
- [[aco]]
