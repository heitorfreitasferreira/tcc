package points

import (
	"math/rand"
	"slices"
)

// pointsPerInstance {2:100, 30:5} : gera 100 instancias com 2 pontos, 5 instancias com 30 pontos
func CreateInstances(seed int64, pointsPerInstace map[int]int) []Points2D {
	rng := rand.New(rand.NewSource(seed))

	inst := make([]Points2D, 0)
	pointCounts := make([]int, 0, len(pointsPerInstace))
	for nPoints := range pointsPerInstace {
		pointCounts = append(pointCounts, nPoints)
	}
	slices.Sort(pointCounts)

	for _, nPoints := range pointCounts {
		nInstances := pointsPerInstace[nPoints]
		for range nInstances {
			points := make([]Coordinate2D, nPoints)

			for j := range nPoints {
				// Coordenadas no espaço R² (-1, +1)
				x := rng.Float64()*2 - 1
				y := rng.Float64()*2 - 1
				points[j][0], points[j][1] = x, y
			}

			inst = append(inst, points)
		}
	}
	return inst
}
