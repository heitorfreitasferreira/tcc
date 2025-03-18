package brute

import (
	"math"
	"tcc/graph"
)

func Optimize(g graph.Graph) ([]int, float64) {
	places := make([]int, len(g)-1)
	for i := 1; i < len(g); i++ {
		places[i-1] = i
	}

	bestI := -1
	bestMksp := math.MaxFloat64
	perms := generatePermutations(places)
	for i, perm := range perms {
		mksp := g.Makespan(perm)
		if bestMksp > mksp {
			bestI = i
			bestMksp = mksp
		}
	}
	return perms[bestI], bestMksp
}

func generatePermutations(arr []int) [][]int {
	var result [][]int
	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
			return
		}
		for i := range n {
			helper(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}
	helper(arr, len(arr))
	return result
}
