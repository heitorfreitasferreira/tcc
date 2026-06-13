package brute

import (
	"math"
	"tcc/graph"
	"tcc/shared"
)

func Optimize(g *graph.Graph, onImprovement func(shared.Improvement)) shared.OptimizationResult {
	places := make([]int, g.N-1)
	for i := 1; i < g.N; i++ {
		places[i-1] = i
	}

	bestPerm := []int{}
	bestMksp := math.MaxFloat64
	evaluations := 0
	improvements := make([]shared.Improvement, 0)

	generatePermutations(places, func(perm []int) {
		evaluations++
		mksp := g.Makespan(perm)
		if bestMksp > mksp {
			delta := 0.0
			if bestMksp < math.MaxFloat64 {
				delta = mksp - bestMksp
			}

			bestPerm = append([]int(nil), perm...)
			bestMksp = mksp

			improvement := shared.Improvement{
				Iteration:    evaluations,
				Evaluation:   evaluations,
				BestMakespan: bestMksp,
				Delta:        delta,
				BestSequence: append([]int(nil), bestPerm...),
			}
			improvements = append(improvements, improvement)
			if onImprovement != nil {
				onImprovement(improvement)
			}
		}
	})

	return shared.OptimizationResult{
		BestSequence:        bestPerm,
		BestMakespan:        bestMksp,
		IterationsCompleted: evaluations,
		Evaluations:         evaluations,
		Improvements:        improvements,
	}
}

func generatePermutations(arr []int, fn func([]int)) {
	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			fn(arr)
			return
		}
		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}
	helper(arr, len(arr))
}
