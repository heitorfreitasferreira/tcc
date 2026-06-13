package graph

import "tcc/points"

func CreateAll(pts []points.Points2D) []*Graph {
	graphs := make([]*Graph, len(pts))
	for i, pt := range pts {
		g := New(pt)
		graphs[i] = &g
	}
	return graphs
}

func New(pts points.Points2D) Graph {
	n := len(pts)
	g := Graph{N: n, Data: make([]float64, n*n*n)}
	for k, old := range pts {
		for i, curr := range pts {
			for j, next := range pts {
				time := curr.EuclideanDistance(next) / float64(droneSpeed)
				currX, currY := old.Vector(curr)
				currVector := vector2D{currX, currY}
				nextX, nextY := curr.Vector(next)
				nextVector := vector2D{nextX, nextY}
				g.Data[k*n*n+i*n+j] = turnCost(currVector, nextVector) + time
			}
		}
	}
	return g
}
