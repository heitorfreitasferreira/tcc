package graph

import (
	"tcc/points"
)

func CreateAll(pts []points.Points2D) []Graph {
	graphs := make([]Graph, len(pts))
	for i, pt := range pts {
		graphs[i] = New(pt)
	}
	return graphs
}

func New(pts points.Points2D) Graph {

	penalty := make(Graph, len(pts))

	for k, old := range pts {
		penalty[k] = make([][]float64, len(pts))
		for i, curr := range pts {
			penalty[k][i] = make([]float64, len(pts))
			for j, next := range pts {
				// Tempo do trajeto de curr -> next
				time := curr.EuclideanDistance(next) / float64(droneSpeed)

				currX, currY := old.Vector(curr)
				currVector := vector2D{currX, currY}
				nextX, nextY := curr.Vector(next)
				nextVector := vector2D{nextX, nextY}

				// Tempo da curva dada a direção que eu to (função da posição atual, posição anterior) e a direção que eu vou
				penalty[k][i][j] = turnCost(currVector, nextVector) + time
			}
		}
	}
	return penalty
}
