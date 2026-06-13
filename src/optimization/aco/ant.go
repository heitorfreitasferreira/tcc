package aco

import (
	"math"
)

type ant struct {
	seq     []int
	visited []bool

	lk float64
}

func (aco *ACO) updatePheromones(ants []ant) {
	for i := range aco.pheromones {
		for j := range aco.pheromones[i] {
			for k := range aco.pheromones[i][j] {
				if i == j || j == k || i == k {
					continue
				}
				aco.pheromones[i][j][k] *= (1 - aco.Rho)
			}
		}
	}

	for _, ant := range ants {
		if ant.lk == 0 {
			continue
		}
		delta := aco.Q / ant.lk
		for idx := 2; idx < len(ant.seq); idx++ {
			prev := ant.seq[idx-2]
			curr := ant.seq[idx-1]
			next := ant.seq[idx]
			aco.pheromones[prev][curr][next] += delta
		}
	}
}

func (aco *ACO) selectNextNode(prev, curr int, visited []bool) int {
	n := aco.Graph.N
	probabilities := make([]float64, n)
	total := 0.0

	for next := 0; next < n; next++ {
		if visited[next] || next == prev || next == curr {
			continue
		}

		heuristic := 1.0 / aco.Graph.At(prev, curr, next)
		prob := math.Pow(aco.pheromones[prev][curr][next], aco.Alpha) *
			math.Pow(heuristic, aco.Beta)

		probabilities[next] = prob
		total += prob
	}

	if total == 0 {
		bestNext := -1
		bestCost := math.MaxFloat64
		for next := 0; next < n; next++ {
			if visited[next] || next == prev || next == curr {
				continue
			}
			if aco.Graph.At(prev, curr, next) < bestCost {
				bestCost = aco.Graph.At(prev, curr, next)
				bestNext = next
			}
		}
		return bestNext
	}

	cutoff := aco.rng.Float64() * total
	sum := 0.0
	for next, prob := range probabilities {
		if sum += prob; sum >= cutoff {
			return next
		}
	}
	return -1
}

func (aco *ACO) walk() ant {
	n := aco.Graph.N
	ant := ant{
		seq:     make([]int, 0, n),
		visited: make([]bool, n),
		lk:      0.0,
	}

	current := 0
	ant.seq = append(ant.seq, current)
	ant.visited[current] = true

	next := aco.rng.Intn(n)
	for next == current {
		next = aco.rng.Intn(n)
	}
	ant.seq = append(ant.seq, next)
	ant.visited[next] = true

	for len(ant.seq) < n {
		prev := ant.seq[len(ant.seq)-2]
		curr := ant.seq[len(ant.seq)-1]
		next := aco.selectNextNode(prev, curr, ant.visited)
		if next == -1 {
			break
		}

		ant.seq = append(ant.seq, next)
		ant.visited[next] = true
	}

	if len(ant.seq) == n {
		ant.lk = aco.Makespan(aco.pathWithoutOrigin(ant.seq))
	}
	return ant
}
