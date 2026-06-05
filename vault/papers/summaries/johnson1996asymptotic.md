---
aliases: [Johnson, McGeoch, Rothberg (1996) — Asymptotic Experimental Analysis for the Held-Karp TSP Bound]
tags: [paper/summary, area/tsp, area/bounds, topic/held-karp, topic/asymptotic-analysis]
---

# Asymptotic Experimental Analysis for the Held-Karp Traveling Salesman Bound

Johnson, D. S.; McGeoch, L. A.; Rothberg, E. E. (1996). *Proceedings of the 7th Annual ACM-SIAM Symposium on Discrete Algorithms*, 341–350.

## 1. Problem and Motivation

Avaliar a qualidade de heurísticas para o TSP exige conhecer o valor da solução ótima, que é computacionalmente inviável para instâncias grandes (o maior TSP não trivial resolvido à época tinha 7.397 cidades, consumindo vários CPU-anos). O *Held-Karp lower bound* — solução da relaxação linear da formulação inteira do TSP — é candidato a substituto, mas faltava evidência empírica sistemática sobre quão próximo ele está do ótimo e qual seu comportamento assintótico.

## 2. Core Method or Approach

- Cálculo **exato** do HK bound via Simplex com oráculos de separação para restrições de subtour (viável até ~30.000 cidades).
- Aproximação **iterativa** por relaxação Lagrangeana com aceleração via MST em subgrafos esparsos e *k-d trees* para instâncias geométricas (viável até ~10⁶ cidades).
- Medição do *HK gap*: percentual pelo qual o ótimo excede o HK bound, para múltiplas classes de instâncias, com intervalos de confiança de 95%.
- Ajuste de curvas por *weighted least squares* aos dados normalizados para estimar constantes assintóticas ($C_{\text{HK}}$, $C_{\text{OPT}}$).
- Uso da topologia **toroidal** (em contraste com a planar) para eliminar o efeito de borda e acelerar a convergência assintótica.

## 3. Main Results

| Resultado | Valor |
|-----------|-------|
| HK gap médio (Euclidiano 2D planar) | ~0,76–0,77% (N ≤ 1.000) |
| HK gap para instâncias TSPLIB | < 2% (média ~0,8%) |
| HK gap para matrizes de distância aleatórias | tende a 0 com N crescente |
| $C_{\text{HK}}$ (Euclidiano 2D) | $0{,}70795 \pm 0{,}00005$ |
| $C_{\text{OPT}}$ (Euclidiano 2D) | $0{,}7124 \pm 0{,}0002$ |
| Tour ótimo esperado (matrizes aleatórias) | $\approx 2{,}042$ |
| Dimensões superiores (3D, 4D) | gaps menores que em 2D, decrescentes com $d$ |

O valor de $C_{\text{OPT}} \approx 0{,}7124$ invalida as estimativas amplamente citadas de $0{,}749$ (Beardwood et al.) e $0{,}765$ (Stein), e compromete alegações de desempenho de heurísticas baseadas nelas.

## 4. Strengths and Limitations

**Pontos fortes:** escala experimental massiva (até N=316.228 na topologia planar, ~10⁶ via aproximação iterativa); cobertura ampla de classes de instâncias (Euclidiana 2D/3D/4D, retilínea, supnorm, matrizes aleatórias, TSPLIB); metodologia estatística com intervalos de confiança; verificação cruzada entre topologias planar e toroidal; testes de sensibilidade a geradores de números aleatórios e precisão das coordenadas.

**Limitações:** para N > 1.000 os dados de gap são anedóticos (poucas instâncias); a aproximação iterativa do HK bound degrada-se em instâncias estruturadas ou com clusters (ex.: erro >3% na instância fl3795 do TSPLIB); coordenadas inteiras com precisão limitada (embora testes indiquem viés desprezível); extrapolações para o limite assintótico dependem da escolha da forma funcional da curva ajustada.

## 5. Practical Takeaway for Researchers

O HK bound é um substituto confiável para o valor ótimo do TSP: está consistentemente a menos de 0,8% do ótimo em instâncias Euclidianas aleatórias e a menos de 2% em instâncias reais do TSPLIB. Além disso, o *HK gap* tem variância muito menor que o *tour length* bruto, permitindo **redução de variância** em estudos experimentais: bons estimadores de desempenho podem ser obtidos com muito menos dados quando se normaliza pelo HK bound da instância (em vez de comparar a valores esperados). A constante $C_{\text{OPT}} \approx 0{,}7124$ deve substituir $0{,}749$ e $0{,}765$ como referência padrão. Para instâncias estruturadas, recomenda-se reportar o valor preciso do bound utilizado, devido à possível degradação das aproximações iterativas do HK bound.
