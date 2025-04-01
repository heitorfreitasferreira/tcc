package ga

import (
	"math/rand"
	"sort"
	"tcc/graph"
	"tcc/shared"
)

type Params struct {
	shared.HyperParams

	Elitism, TournamentSize int
	MutationRate            float64
}

type individual struct {
	gen []int
	fen float64
}

func Optimize(p Params, g graph.Graph, rnd *rand.Rand) ([]int, float64) {
	// Initial population
	pop := make([]individual, p.PopulationSize)
	for i := range p.PopulationSize {
		pop[i] = individual{gen: make([]int, len(g)), fen: -1.0}
		for j := range g {
			pop[i].gen[j] = j
		}
		shared.Shuffle(pop[i].gen, rnd)
		pop[i].fen = g.Makespan(pop[i].gen)
	}

	for range p.Iterations {
		newPop := make([]individual, 0, p.PopulationSize)

		if p.Elitism > 0 {
			sort.Slice(pop, func(i, j int) bool { return pop[i].fen < pop[j].fen })
			newPop = append(newPop, pop[:p.Elitism]...)
		}

		// Generate new population
		for len(newPop) < p.PopulationSize {
			// Select parents
			parent1 := selectParentTournament(pop, p, rnd)
			parent2 := selectParentTournament(pop, p, rnd)

			// Crossover
			child1, child2 := orderedCrossover(parent1.gen, parent2.gen, rnd)

			// Mutation
			mutateSwap(child1, p.MutationRate, rnd)
			mutateSwap(child2, p.MutationRate, rnd)

			// Evaluate fitness
			fen1 := g.Makespan(child1)
			fen2 := g.Makespan(child2)

			// Add to new population
			newPop = append(newPop, individual{child1, fen1}, individual{child2, fen2})
		}

		// Trim to population size if necessary
		if len(newPop) > p.PopulationSize {
			newPop = newPop[:p.PopulationSize]
		}

		pop = newPop
	}

	sort.Slice(pop, func(i, j int) bool { return pop[i].fen < pop[j].fen })
	best := pop[0]
	return best.gen, best.fen
}

func selectParentTournament(pop []individual, p Params, rnd *rand.Rand) individual {
	best := pop[rnd.Intn(len(pop))]
	for i := 1; i < int(p.TournamentSize); i++ {
		candidate := pop[rnd.Intn(len(pop))]
		if candidate.fen < best.fen {
			best = candidate
		}
	}
	return best
}

func orderedCrossover(parent1, parent2 []int, rnd *rand.Rand) ([]int, []int) {
	size := len(parent1)
	child1 := make([]int, size)
	child2 := make([]int, size)

	// Select crossover points
	pt1 := rnd.Intn(size)
	pt2 := rnd.Intn(size)
	if pt1 > pt2 {
		pt1, pt2 = pt2, pt1
	}

	// Copy segments from parents to children
	segment1 := make(map[int]bool)
	segment2 := make(map[int]bool)
	for i := pt1; i <= pt2; i++ {
		child1[i] = parent1[i]
		segment1[parent1[i]] = true
		child2[i] = parent2[i]
		segment2[parent2[i]] = true
	}

	// Fill remaining positions for child1 from parent2
	idx := 0
	for i := range size {
		if i < pt1 || i > pt2 {
			for {
				elem := parent2[idx]
				idx++
				if !segment1[elem] {
					child1[i] = elem
					break
				}
			}
		}
	}

	// Fill remaining positions for child2 from parent1
	idx = 0
	for i := range size {
		if i < pt1 || i > pt2 {
			for {
				elem := parent1[idx]
				idx++
				if !segment2[elem] {
					child2[i] = elem
					break
				}
			}
		}
	}

	return child1, child2
}

func mutateSwap(gen []int, mutationRate float64, rnd *rand.Rand) {
	if rnd.Float64() < mutationRate {
		i := rnd.Intn(len(gen))
		j := rnd.Intn(len(gen))
		gen[i], gen[j] = gen[j], gen[i]
	}
}
