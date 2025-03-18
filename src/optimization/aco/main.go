package aco

import (
	"math"
	"math/rand"
	"tcc/graph"
	"tcc/shared"
)

type Params struct {
	shared.HyperParams
	Alpha, Beta, Gama, Rho float64
	Q                      float64
}

type ACO struct {
	graph.Graph
	Params
	rng *rand.Rand

	pheromones [][][]float64
}

func Optimize(p Params, g graph.Graph, rng *rand.Rand) string {
	aco := new(p, g, rng)
	sts := shared.NewStats()
	bestCost := math.MaxFloat64
	var bestPath []int

	for range aco.Iterations {
		ants := make([]ant, aco.PopulationSize)
		for i := range ants {
			a := aco.walk()
			ants[i] = a

			if a.lk < bestCost && len(a.seq) == len(aco.Graph) {
				bestCost = a.lk
				bestPath = make([]int, len(a.seq))
				copy(bestPath, a.seq)
			}
		}
		aco.updatePheromones(ants)
		sts.AddIterData(bestCost, []float64{}, bestPath)
	}
	return sts.ToCsv()
}

func new(p Params, g graph.Graph, rng *rand.Rand) *ACO {
	pheromones := make([][][]float64, len(g))
	for i := range len(g) {
		pheromones[i] = make([][]float64, len(g))
		for j := range len(g) {
			pheromones[i][j] = make([]float64, len(g))
			for k := range len(g) {
				if i != j && j != k && i != k {
					pheromones[i][j][k] = 1.0 // Valor inicial
				}
			}
		}
	}

	return &ACO{
		pheromones: pheromones,
		Graph:      g,
		Params:     p,
		rng:        rng,
	}
}
