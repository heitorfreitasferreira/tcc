# Velocidades e afins

penalizacao = penalizacao_max*(ang/180)

Soma a penalização ao custo da aresta.

# Modelagem do grafo

gerar N pontos no R², com coordenadas aleatórias, calcular o angulo e distância entre eles, e gerar o grafo. o grafo vai ser uma matriz de adjacência, onde cada aresta é uma tupla (distancia, angulo).
Gerar 1 grafos diferentes para cada qnt de vertices e salvar em um arquivo, gerar de 5,6,7,8,9,10,15,20,25,30,35,40,45,50.

# Exps

Rodar cada metaheurística N vezes (pelo menos 30), pegar a melhor solução de cada execução, cada um vai ter tempo,distancia, soma dos angulos, soma quadratica dos angulos e penalização total, solução (caminho).

# trabalhos futuros

- Implementar outras metaheurísticas
- Implementar para drone de asa fixa, considerando aceleração, velocidade máxima, perda de velocidade em curvas, etc.
- otimização multiobj (distancia e tempo tbm)
