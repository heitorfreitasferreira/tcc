package lowerbound

import (
	"math"
	"tcc/graph"
	"tcc/shared"
)

func Optimize(g *graph.Graph, onImprovement func(shared.Improvement)) shared.OptimizationResult {
	n := g.N

	cost := reduce3Dto2D(g, n)

	assign, bound := hungarian(cost)

	seq := extractSequence(assign, n)

	improvement := shared.Improvement{
		Iteration:    1,
		Evaluation:   1,
		BestMakespan: bound,
		Delta:        0,
		BestSequence: append([]int(nil), seq...),
	}

	if onImprovement != nil {
		onImprovement(improvement)
	}

	return shared.OptimizationResult{
		BestSequence:        seq,
		BestMakespan:        bound,
		IterationsCompleted: 1,
		Evaluations:         1,
		Improvements:        []shared.Improvement{improvement},
	}
}

func reduce3Dto2D(g *graph.Graph, n int) [][]float64 {
	cost := make([][]float64, n)
	for j := range n {
		cost[j] = make([]float64, n)
		for k := range n {
			if j == k {
				cost[j][k] = math.Inf(1)
				continue
			}
			minCost := math.Inf(1)
			for i := range n {
				if g.At(i, j, k) < minCost {
					minCost = g.At(i, j, k)
				}
			}
			cost[j][k] = minCost
		}
	}
	return cost
}

func extractSequence(assign []int, n int) []int {
	if n <= 1 {
		return nil
	}

	visited := make([]bool, n)
	seq := make([]int, 0, n-1)

	curr := assign[0]
	for curr != 0 && !visited[curr] {
		visited[curr] = true
		seq = append(seq, curr)
		curr = assign[curr]
	}
	visited[0] = true

	for i := 1; i < n; i++ {
		if !visited[i] {
			curr := i
			for !visited[curr] {
				visited[curr] = true
				seq = append(seq, curr)
				curr = assign[curr]
			}
		}
	}

	return seq
}
