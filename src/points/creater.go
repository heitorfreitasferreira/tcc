package points

import (
	"math/rand"
)

// pointsPerInstance {2:100, 30:5} : gera 100 instancias com 2 pontos, 5 instancias com 30 pontos
func CreateInstances(seed int64, pointsPerInstace map[int]int) []Points2D {
	rng := rand.New(rand.NewSource(seed))

	inst := make([]Points2D, 0)
	for nPoints, nInstances := range pointsPerInstace {
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
