package aco

import (
	"fmt"
	"math/rand"
	"sync"
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

func Optimize(p Params, g graph.Graph, rng *rand.Rand) {
	aco := new(p, g, rng)
	var bestPath []int
	var bestCost float64
	var mutex sync.Mutex
	for range aco.Iterations {
		ants := make([]ant, aco.PopulationSize)
		seeds := make([]int64, aco.PopulationSize)

		// Geração segura de seeds
		for i := range aco.PopulationSize {
			seeds[i] = aco.rng.Int63()
		}

		var wg sync.WaitGroup
		for i := range aco.PopulationSize {
			wg.Add(1)
			go func(idx int, seed int64) {
				defer wg.Done()
				localRNG := rand.New(rand.NewSource(seed))
				ant := aco.walk(localRNG)

				// Atualização thread-safe do melhor caminho
				mutex.Lock()
				defer mutex.Unlock()
				if ant.lk < bestCost && len(ant.seq) == len(aco.Graph) {
					bestCost = ant.lk
					bestPath = make([]int, len(ant.seq))
					copy(bestPath, ant.seq)
				}
				ants[idx] = ant
			}(i, seeds[i])
		}
		wg.Wait()

		aco.updatePheromones(ants)
	}
	fmt.Println(bestPath, bestCost)
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
