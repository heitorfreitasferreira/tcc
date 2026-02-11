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

func Optimize(p Params, g graph.Graph, rng *rand.Rand, onImprovement func(shared.Improvement)) shared.OptimizationResult {
	aco := new(p, g, rng)
	result := shared.OptimizationResult{
		BestMakespan: math.MaxFloat64,
	}
	evaluationCount := 0
	bestCost := math.MaxFloat64
	var bestPath []int

	for iteration := 1; iteration <= aco.Iterations; iteration++ {
		ants := make([]ant, aco.PopulationSize)
		for i := range ants {
			a := aco.walk()
			ants[i] = a
			evaluationCount++

			if a.lk < bestCost && len(a.seq) == len(aco.Graph) {
				delta := 0.0
				if bestCost < math.MaxFloat64 {
					delta = a.lk - bestCost
				}

				bestCost = a.lk
				bestPath = make([]int, len(a.seq))
				copy(bestPath, a.seq)

				improvement := shared.Improvement{
					Iteration:    iteration,
					Evaluation:   evaluationCount,
					BestMakespan: bestCost,
					Delta:        delta,
					BestSequence: append([]int(nil), bestPath...),
				}
				result.Improvements = append(result.Improvements, improvement)
				if onImprovement != nil {
					onImprovement(improvement)
				}
			}
		}
		aco.updatePheromones(ants)
	}

	result.BestMakespan = bestCost
	result.BestSequence = append([]int(nil), bestPath...)
	result.IterationsCompleted = aco.Iterations
	result.Evaluations = evaluationCount

	return result
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
