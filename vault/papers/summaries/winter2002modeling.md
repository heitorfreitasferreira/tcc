# Modeling Costs of Turns in Route Planning — Resumo

**Winter, S.** (2002). *GeoInformatica* 6(4), 345–361.

## 1. Problema e Motivação

Em roteamento tradicional, custos são associados apenas às arestas do grafo. Contudo, custos
de conversão (restrições de giro, ângulos de curva, tempos de espera) são relações binárias
entre arestas consecutivas e não podem ser armazenados como atributos de nós ou arestas. As
soluções vigentes — tabelas de exceção e expansão de nós — inflam o tamanho do grafo,
introduzem dois tipos de arestas com semânticas distintas e exigem modificações nos
algoritmos de caminho mínimo. O problema é particularmente relevante para navegação de
pedestres e serviços baseados em localização.

## 2. Método Central (Grafo Pseudo-Dual)

- Cada **aresta** do grafo primal torna-se um **nó** no grafo pseudo-dual.
- Cada **par de arestas consecutivas** (giro) torna-se uma **aresta** no grafo pseudo-dual.
- Custos de percurso são atribuídos aos nós pseudo-duais; custos de giro, às arestas pseudo-duais.
- Um nó virtual de origem e um nó virtual de destino conectam os gateways de partida e
  chegada, reduzindo o problema a uma origem e um destino únicos.
- Algoritmos clássicos (e.g., Dijkstra, *k*-shortest paths) executam **sem modificações**
  sobre o grafo pseudo-dual.

## 3. Principais Resultados

- **Armazenamento**: \(|N_D| = |E_G|\) e \(|E_D|\) é \(O(|E_G|)\) em redes esparsas; o grafo
  pseudo-dual é consistentemente menor que o grafo com nós expandidos em ambas as dimensões.
- **Eficiência computacional**: ganho mais que linear em relação ao grafo expandido, pois
  algoritmos de caminho mínimo são \(O(n \log n)\).
- **Aplicações demonstradas**: (i) restrições de giro com custo infinito; (ii) rota com
  número mínimo de curvas; (iii) rota de menor ângulo total; (iv) conceitos semânticos de
  giro (e.g., "siga pela mesma rua"); (v) especificação de propriedades topológicas para
  trilhas de caminhada (rotas circulares, evitar mesmo segmento no mesmo sentido).
- **Validação**: protótipo funcional em Haskell sobre rede de Viena (~2500 ruas, com
  restrições de giro e mão única); implementação Java para planejador de trilhas usando
  *k*-shortest paths.

## 4. Pontos Fortes e Limitações

**Pontos fortes:**
- Separação conceitual limpa entre custos de percurso e custos de giro.
- Menor consumo de memória que expansão de nós.
- Algoritmos padrão de caminho mínimo funcionam sem alterações.
- Flexibilidade para modelar múltiplas funções de custo (discretas, angulares, semânticas).

**Limitações:**
- O grafo pseudo-dual **não preserva a imersão espacial** do grafo original; para desenhar
  rotas é necessário referenciar o grafo primal.
- As suposições cognitivas sobre percepção de giros por pedestres **não foram validadas
  empiricamente**.
- Os protótipos implementados não são otimizados para desempenho em larga escala ou
  acesso multiusuário.
- A manutenção do grafo pseudo-dual após alterações no grafo primal permanece como
  questão em aberto.

## 5. Implicação Prática para Pesquisadores

Sempre que custos de conversão forem relevantes para o problema de roteamento (restrições
de giro, minimização de curvas, continuidade de via), o grafo pseudo-dual oferece uma
representação mais enxuta e conceitualmente mais limpa que a expansão de nós. A
construção é direta e algoritmos clássicos de caminho mínimo são aplicáveis sem
modificações. A principal contrapartida é a perda da informação espacial (coordenadas),
exigindo referência ao grafo primal para visualização.
