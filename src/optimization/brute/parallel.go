package brute

import (
	"math"
	"runtime"
	"sync"

	"tcc/graph"
	"tcc/shared"
)

func OptimizeParallel(g *graph.Graph, workers int, onImprovement func(shared.Improvement)) shared.OptimizationResult {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	places := make([]int, g.N-1)
	for i := 1; i < g.N; i++ {
		places[i-1] = i
	}

	type task struct {
		first     int
		remaining []int
	}

	tasks := make(chan task, len(places))
	for _, first := range places {
		rem := make([]int, 0, len(places)-1)
		for _, v := range places {
			if v != first {
				rem = append(rem, v)
			}
		}
		tasks <- task{first, rem}
	}
	close(tasks)

	type workerResult struct {
		bestMksp float64
		bestPerm []int
		evals    int
		imps     []shared.Improvement
	}

	numWorkers := workers
	if numWorkers > len(places) {
		numWorkers = len(places)
	}

	results := make(chan workerResult, numWorkers)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			localBest := math.MaxFloat64
			var localPerm []int
			localEvals := 0
			localImps := make([]shared.Improvement, 0)

			for task := range tasks {
				combined := make([]int, 1+len(task.remaining))
				combined[0] = task.first

				generatePermutations(task.remaining, func(perm []int) {
					copy(combined[1:], perm)
					mksp := g.Makespan(combined)
					localEvals++

					if mksp < localBest {
						delta := 0.0
						if localBest < math.MaxFloat64 {
							delta = mksp - localBest
						}
						localBest = mksp
						localPerm = append([]int(nil), combined...)

						imp := shared.Improvement{
							Iteration:    localEvals,
							Evaluation:   localEvals,
							BestMakespan: localBest,
							Delta:        delta,
							BestSequence: append([]int(nil), localPerm...),
						}
						localImps = append(localImps, imp)

						if onImprovement != nil {
							mu.Lock()
							onImprovement(imp)
							mu.Unlock()
						}
					}
				})
			}

			results <- workerResult{localBest, localPerm, localEvals, localImps}
		}()
	}

	wg.Wait()
	close(results)

	globalBest := math.MaxFloat64
	var globalPerm []int
	totalEvals := 0
	allImps := make([]shared.Improvement, 0)

	for r := range results {
		totalEvals += r.evals
		if r.bestMksp < globalBest {
			globalBest = r.bestMksp
			globalPerm = r.bestPerm
		}
		allImps = append(allImps, r.imps...)
	}

	return shared.OptimizationResult{
		BestSequence:        globalPerm,
		BestMakespan:        globalBest,
		IterationsCompleted: totalEvals,
		Evaluations:         totalEvals,
		Improvements:        allImps,
	}
}
