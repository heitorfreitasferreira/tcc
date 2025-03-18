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

	bestPerm := []int{}
	bestMksp := math.MaxFloat64

	for perm := range generatePermutations(places) {
		mksp := g.Makespan(perm)
		if bestMksp > mksp {
			bestPerm = perm
			bestMksp = mksp
		}
	}
	return bestPerm, bestMksp
}

func generatePermutations(arr []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		var helper func([]int, int)
		helper = func(arr []int, n int) {
			if n == 1 {
				tmp := make([]int, len(arr))
				copy(tmp, arr)
				ch <- tmp
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
	}()
	return ch
}
