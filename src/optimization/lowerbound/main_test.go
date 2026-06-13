package lowerbound

import (
	"math"
	"testing"

	"tcc/graph"
	"tcc/shared"
)

func TestHungarianSmall(t *testing.T) {
	t.Parallel()

	cost := [][]float64{
		{math.Inf(1), 10, 15},
		{5, math.Inf(1), 20},
		{10, 25, math.Inf(1)},
	}

	assign, total := hungarian(cost)

	if len(assign) != 3 {
		t.Fatalf("assign size: got=%d want=%d", len(assign), 3)
	}

	visited := make([]bool, 3)
	for i, j := range assign {
		if visited[j] {
			t.Fatalf("coluna %d atribuída duas vezes (linha %d)", j, i)
		}
		visited[j] = true
	}

	expectedCost := 10.0 + 20.0 + 10.0
	if math.Abs(total-expectedCost) > 1e-9 {
		t.Fatalf("custo total: got=%.6f want=%.6f", total, expectedCost)
	}
}

func TestHungarianTwoNodes(t *testing.T) {
	t.Parallel()

	cost := [][]float64{
		{math.Inf(1), 42},
		{7, math.Inf(1)},
	}

	assign, total := hungarian(cost)

	if len(assign) != 2 {
		t.Fatalf("assign size: got=%d want=%d", len(assign), 2)
	}

	expectedCost := 42.0 + 7.0
	if math.Abs(total-expectedCost) > 1e-9 {
		t.Fatalf("custo total: got=%.6f want=%.6f", total, expectedCost)
	}

	if assign[0] != 1 || assign[1] != 0 {
		t.Fatalf("assign: got=%v want=[1 0]", assign)
	}
}

func TestExtractSequence(t *testing.T) {
	t.Parallel()

	assign := []int{1, 2, 3, 0}
	seq := extractSequence(assign, 4)

	expected := []int{1, 2, 3}
	if len(seq) != len(expected) {
		t.Fatalf("tamanho: got=%d want=%d", len(seq), len(expected))
	}
	for i := range expected {
		if seq[i] != expected[i] {
			t.Fatalf("seq[%d]: got=%d want=%d", i, seq[i], expected[i])
		}
	}
}

func TestExtractSequenceWithSubtours(t *testing.T) {
	t.Parallel()

	assign := []int{2, 0, 1, 4, 3}
	seq := extractSequence(assign, 5)

	visited := make([]bool, 5)
	visited[0] = true
	for _, v := range seq {
		if v <= 0 || v >= 5 {
			t.Fatalf("nó inválido: %d", v)
		}
		if visited[v] {
			t.Fatalf("nó repetido: %d", v)
		}
		visited[v] = true
	}

	if len(seq) != 4 {
		t.Fatalf("tamanho: got=%d want=%d", len(seq), 4)
	}
}

func TestReduce3Dto2D(t *testing.T) {
	t.Parallel()

	g := testGraph(4)
	cost := reduce3Dto2D(g, 4)

	for j := range 4 {
		for k := range 4 {
			if j == k {
				if !math.IsInf(cost[j][k], 1) {
					t.Fatalf("self-loop[%d][%d] deveria ser INF, got=%.2f", j, k, cost[j][k])
				}
			} else {
				if math.IsInf(cost[j][k], 1) {
					t.Fatalf("cost[%d][%d] é INF, esperado finito", j, k)
				}
			}
		}
	}
}

func TestOptimizeBasic(t *testing.T) {
	t.Parallel()

	g := testGraph(5)
	result := Optimize(g, nil)

	if result.Evaluations != 1 {
		t.Fatalf("evaluations: got=%d want=1", result.Evaluations)
	}

	if len(result.Improvements) != 1 {
		t.Fatalf("improvements: got=%d want=1", len(result.Improvements))
	}

	if math.IsInf(result.BestMakespan, 1) || result.BestMakespan <= 0 {
		t.Fatalf("best_makespan inválido: %.6f", result.BestMakespan)
	}

	if len(result.BestSequence) == 0 {
		t.Fatal("best_sequence vazia")
	}
}

func TestOptimizeBoundIsLower(t *testing.T) {
	t.Parallel()

	g := testGraph(6)
	apResult := Optimize(g, nil)
	bruteResult := bruteForce(g)

	if apResult.BestMakespan > bruteResult.BestMakespan+1e-9 {
		t.Fatalf("AP bound > ótimo: AP=%.6f ótimo=%.6f",
			apResult.BestMakespan, bruteResult.BestMakespan)
	}
}

func bruteForce(g *graph.Graph) shared.OptimizationResult {
	n := g.N
	places := make([]int, n-1)
	for i := 1; i < n; i++ {
		places[i-1] = i
	}

	bestMksp := math.MaxFloat64
	bestPerm := []int{}

	generatePermutations(places, func(perm []int) {
		mksp := g.Makespan(perm)
		if mksp < bestMksp {
			bestMksp = mksp
			bestPerm = append([]int(nil), perm...)
		}
	})

	return shared.OptimizationResult{
		BestSequence: bestPerm,
		BestMakespan: bestMksp,
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

func testGraph(nodes int) *graph.Graph {
	g := &graph.Graph{N: nodes, Data: make([]float64, nodes*nodes*nodes)}
	for prev := 0; prev < nodes; prev++ {
		for curr := 0; curr < nodes; curr++ {
			for next := 0; next < nodes; next++ {
				if prev == curr && curr == next {
					continue
				}
				g.Data[prev*nodes*nodes+curr*nodes+next] = float64(1 + prev + curr + next)
			}
		}
	}
	return g
}
