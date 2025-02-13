package aco

import (
	"fmt"
	"math"
	"math/rand"
	"tcc/graph"
)

func Optimize(p Params, g graph.Graph, rng *rand.Rand) {
	aco := new(p, g, rng)
	var seq []int
	var mksp float64
	for range aco.Iterations {
		mksp, seq = aco.iterate()
	}
	fmt.Println(seq, mksp)
}

func (aco *ACO) iterate() (float64, []int) {
	globalBestLk := math.MaxFloat64
	globalBestid := 0
	aco.ants = []ant{}
	for i := range aco.NumberOfAnts {
		fuu := aco.walk()
		aco.ants = append(aco.ants, fuu)
		if fuu.lk < globalBestLk {
			globalBestLk = fuu.lk
			globalBestid = i
		}
	}

	return globalBestLk, aco.ants[globalBestid].seq
}

func (aco *ACO) walk() ant {
	seq := []int{}

	currNode := 0
	lastNode := 0
	visited := make([]bool, len(aco.Graph))

	visited[0] = true

	for !finished(visited) {
		cumSum := cumSum(aco.probs[lastNode][currNode])
		r := aco.rng.Float64()
		var nextNode int
		for i, value := range cumSum {
			if r < value {
				nextNode = i
				break
			}
		}
		visited[nextNode] = true
		seq = append(seq, nextNode)

		lastNode = currNode
		currNode = nextNode
	}
	return ant{
		seq: seq,
		lk:  aco.Graph.Makespan(seq),
	}
}
