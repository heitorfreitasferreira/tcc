package ga

import (
	"cmp"
	"math"
	"math/rand"
	"slices"
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

func Optimize(p Params, g *graph.Graph, rnd *rand.Rand, onImprovement func(shared.Improvement)) shared.OptimizationResult {
	result := shared.OptimizationResult{
		BestMakespan: math.MaxFloat64,
	}

	evaluationCount := 0
	reportImprovement := func(iteration int, sequence []int, makespan float64) {
		if makespan >= result.BestMakespan {
			return
		}

		delta := 0.0
		if result.BestMakespan < math.MaxFloat64 {
			delta = makespan - result.BestMakespan
		}

		result.BestMakespan = makespan
		result.BestSequence = append([]int(nil), sequence...)

		improvement := shared.Improvement{
			Iteration:    iteration,
			Evaluation:   evaluationCount,
			BestMakespan: makespan,
			Delta:        delta,
			BestSequence: append([]int(nil), sequence...),
		}
		result.Improvements = append(result.Improvements, improvement)

		if onImprovement != nil {
			onImprovement(improvement)
		}
	}

	// Initial population
	nodes := make([]int, g.N-1)
	for node := 1; node < g.N; node++ {
		nodes[node-1] = node
	}

	pop := make([]individual, p.PopulationSize)
	for i := 0; i < p.PopulationSize; i++ {
		pop[i] = individual{gen: append([]int(nil), nodes...), fen: -1.0}
		shared.Shuffle(pop[i].gen, rnd)
		pop[i].fen = g.Makespan(pop[i].gen)
		evaluationCount++
		reportImprovement(0, pop[i].gen, pop[i].fen)
	}

	for iteration := 1; iteration <= p.Iterations; iteration++ {
		newPop := make([]individual, 0, p.PopulationSize)

		if p.Elitism > 0 {
			slices.SortFunc(pop, func(a, b individual) int { return cmp.Compare(a.fen, b.fen) })
			eliteSize := p.Elitism
			if eliteSize > len(pop) {
				eliteSize = len(pop)
			}
			newPop = append(newPop, pop[:eliteSize]...)
		}

		for len(newPop) < p.PopulationSize {
			parent1 := selectParentTournament(pop, p, rnd)
			parent2 := selectParentTournament(pop, p, rnd)

			child1, child2 := orderedCrossover(parent1.gen, parent2.gen, rnd)

			mutateSwap(child1, p.MutationRate, rnd)
			mutateSwap(child2, p.MutationRate, rnd)

			fen1 := g.Makespan(child1)
			evaluationCount++
			reportImprovement(iteration, child1, fen1)

			fen2 := g.Makespan(child2)
			evaluationCount++
			reportImprovement(iteration, child2, fen2)

			newPop = append(newPop, individual{child1, fen1}, individual{child2, fen2})
		}

		if len(newPop) > p.PopulationSize {
			newPop = newPop[:p.PopulationSize]
		}

		pop = newPop
	}

	result.IterationsCompleted = p.Iterations
	result.Evaluations = evaluationCount

	if len(result.BestSequence) == 0 && len(pop) > 0 {
		slices.SortFunc(pop, func(a, b individual) int { return cmp.Compare(a.fen, b.fen) })
		result.BestSequence = append([]int(nil), pop[0].gen...)
		result.BestMakespan = pop[0].fen
	}

	return result
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

	pt1 := rnd.Intn(size)
	pt2 := rnd.Intn(size)
	if pt1 > pt2 {
		pt1, pt2 = pt2, pt1
	}

	segment1 := make(map[int]bool)
	segment2 := make(map[int]bool)
	for i := pt1; i <= pt2; i++ {
		child1[i] = parent1[i]
		segment1[parent1[i]] = true
		child2[i] = parent2[i]
		segment2[parent2[i]] = true
	}

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
