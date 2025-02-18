import numpy as np
import matplotlib.pyplot as plt
from matplotlib.patches import FancyArrowPatch

def plot_tsp_route(points, route):
    """
    Plota os pontos e o trajeto de um Problema do Caixeiro Viajante (TSP)
    usando setas direcionadas
    
    Parâmetros:
    points (numpy.ndarray): Matriz de pontos em R2 (coordenadas x, y)
    route (list): Sequência de índices representando a ordem da rota
    """
    plt.figure(figsize=(10, 10))
    
    # Plota os pontos
    plt.scatter(points[:, 0], points[:, 1], color='red', s=100, zorder=2, facecolors='none', edgecolors='red')
    
    # Plota a rota
    route_points = points[route]
    
    for i in range(len(route)-1):
        arrow = FancyArrowPatch(
            (route_points[i, 0], route_points[i, 1]),
            (route_points[i+1, 0], route_points[i+1, 1]),
            color='black',
            arrowstyle='->',
            mutation_scale=20,
            linewidth=2,
            zorder=1
        )
        plt.gca().add_patch(arrow)
    
    arrow = FancyArrowPatch(
        (route_points[-1, 0], route_points[-1, 1]),
        (route_points[0, 0], route_points[0, 1]),
        color='black',
        arrowstyle='->',
        mutation_scale=20,
        linewidth=2,
        zorder=1
    )
    plt.gca().add_patch(arrow)
    
    # Numera os pontos
    for i, (x, y) in enumerate(points[route]):
        plt.text(x, y, str(i), fontsize=12, 
                verticalalignment='bottom', 
                horizontalalignment='right')
    
    plt.title('Rota do Problema do Caixeiro Viajante')
    plt.xlabel('Coordenada X')
    plt.ylabel('Coordenada Y')
    plt.grid(True, linestyle='--', alpha=0.7)
    plt.axis('equal')
    plt.tight_layout()
    plt.show()

# Exemplo de uso
if __name__ == "__main__":
    # Gera pontos aleatórios em R2 entre -1 e 1
    np.random.seed(42)
    points = np.random.uniform(-1, 1, (10, 2))
    
    # Exemplo de rota (pode ser substituída por qualquer sequência)
    route = list(range(len(points))) + [0]
    
    # Plota a rota
    plot_tsp_route(points, route)
