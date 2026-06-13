package aco

import (
	"math"
	"math/rand"
	"testing"

	"tcc/graph"
	"tcc/shared"
)

func TestOptimizeUsesRoutesWithoutOrigin(t *testing.T) {
	t.Parallel()

	g := testGraph(6)
	params := Params{
		HyperParams: shared.HyperParams{
			Iterations:     8,
			PopulationSize: 16,
		},
		Alpha: 1.0,
		Beta:  2.0,
		Rho:   0.5,
		Q:     100.0,
	}

	result := Optimize(params, g, rand.New(rand.NewSource(17)), nil)

	assertRouteWithoutOrigin(t, g.N, result.BestSequence)
	for _, imp := range result.Improvements {
		assertRouteWithoutOrigin(t, g.N, imp.BestSequence)
	}

	expected := g.Makespan(result.BestSequence)
	if math.Abs(result.BestMakespan-expected) > 1e-9 {
		t.Fatalf("best_makespan inconsistente: got=%.12f expected=%.12f", result.BestMakespan, expected)
	}
}

func assertRouteWithoutOrigin(t *testing.T, nodes int, seq []int) {
	t.Helper()

	if len(seq) != nodes-1 {
		t.Fatalf("tamanho da rota inválido: got=%d expected=%d", len(seq), nodes-1)
	}

	seen := make([]bool, nodes)
	for _, node := range seq {
		if node <= 0 || node >= nodes {
			t.Fatalf("nó inválido na rota: %d", node)
		}
		if seen[node] {
			t.Fatalf("nó repetido na rota: %d", node)
		}
		seen[node] = true
	}
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
