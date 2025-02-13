package aco

import (
	"math"
	"math/rand"
	"tcc/graph"
)

type ant struct {
	seq []int
	lk  float64
}

type Params struct {
	NumberOfAnts, Iterations int

	// TODO: faz sentido usar a versão que altera Alpha e Beta ao longo das iterações?
	// https://sci-hub.se/https://www.sciencedirect.com/science/article/pii/S1876610212004377
	Alpha, Beta, Gama, Rho float64
	Q                      float64
}

type ACO struct {
	graph.Graph
	Params
	rng *rand.Rand

	pheromones [][][]float64
	probs      [][][]float64 // [anterior][atual][proximo]

	ants []ant
}

func new(p Params, g graph.Graph, rng *rand.Rand) ACO {
	pheromones := make([][][]float64, len(g))
	probs := make([][][]float64, len(g))
	for i := range len(g) {
		pheromones[i] = make([][]float64, len(g))
		probs[i] = make([][]float64, len(g))
		for j := range len(g) {
			pheromones[i][j] = make([]float64, len(g))
			probs[i][j] = make([]float64, len(g))

			for k := range len(g) {
				pheromones[i][j][k] = 1
			}
		}
	}

	ants := make([]ant, p.NumberOfAnts)
	for i := range ants {
		ants[i] = ant{
			seq: []int{},
			lk:  math.MaxFloat64,
		}
	}

	aco := ACO{
		pheromones: pheromones,
		probs:      probs,
		ants:       ants,
		Graph:      g,
		Params:     p,
		rng:        rng,
	}
	// Primeira iteração deve considerar apenas as distâncias
	aco.updateProbs()
	return aco
}

func (aco *ACO) updatePheromones() {
	// Evaporando antes de depositar
	factor := 1 - aco.Rho
	for last := range aco.probs {
		for curr := range aco.probs[last] {
			for next := range aco.probs[last][curr] {
				aco.pheromones[last][curr][next] *= factor
			}
		}
	}
	// Depositar os novos feromonios
	for _, fuu := range aco.ants {
		prevNode := 0
		currNode := 0

		for _, nextNode := range fuu.seq {
			aco.pheromones[prevNode][currNode][nextNode] += 1 / fuu.lk
			prevNode = currNode
			currNode = nextNode
		}
	}
}

func (aco *ACO) updateProbs() {
	// Suponho que o feromonio já foi att
	for last := range aco.probs {
		for curr := range aco.probs[last] {
			// aco.pheromones[last][curr] ^ aco.Alpha * aco.Graph[last][curr] ^ aco.Beta
			sum := 0.0
			numerators := make([]float64, len(aco.Graph))
			for next := range aco.probs[last][curr] {
				numerators[next] = math.Pow(aco.pheromones[last][curr][next], aco.Alpha) * math.Pow(aco.Graph[last][curr][next], aco.Beta)
				sum += numerators[next]
			}

			for next := range aco.probs[last][curr] {
				aco.probs[last][curr][next] = numerators[next] / sum
			}
		}
	}
}

// TODO: Falta fazer as formigas escolherem os caminhos + definir caso de parada
