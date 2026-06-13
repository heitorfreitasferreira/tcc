package brute

import (
	"math"
	"os"
	"testing"

	"tcc/graph"
)

func TestOptimizeParallelMatchesSerial(t *testing.T) {
	data, err := os.ReadFile("../../data/10a.graph")
	if err != nil {
		t.Skipf("skip: cannot load 10a.graph: %v", err)
	}

	var g graph.Graph
	if err := g.UnmarshalJSON(data); err != nil {
		t.Fatalf("unmarshal graph: %v", err)
	}

	serial := Optimize(&g, nil)
	parallel := OptimizeParallel(&g, 0, nil)

	if math.Abs(serial.BestMakespan-parallel.BestMakespan) > 1e-9 {
		t.Fatalf("best makespan mismatch: serial=%.6f parallel=%.6f",
			serial.BestMakespan, parallel.BestMakespan)
	}

	if len(serial.BestSequence) != len(parallel.BestSequence) {
		t.Fatalf("sequence length mismatch: serial=%d parallel=%d",
			len(serial.BestSequence), len(parallel.BestSequence))
	}

	if serial.Evaluations != parallel.Evaluations {
		t.Fatalf("evaluation count mismatch: serial=%d parallel=%d",
			serial.Evaluations, parallel.Evaluations)
	}

	t.Logf("OK: serial=%.6f parallel=%.6f evals=%d n=%d",
		serial.BestMakespan, parallel.BestMakespan, serial.Evaluations, g.N)
}

func BenchmarkOptimizeSerial(b *testing.B) {
	g := newTestGraph(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Optimize(g, nil)
	}
}

func BenchmarkOptimizeParallel(b *testing.B) {
	g := newTestGraph(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OptimizeParallel(g, 0, nil)
	}
}

func BenchmarkOptimizeParallel4(b *testing.B) {
	g := newTestGraph(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OptimizeParallel(g, 4, nil)
	}
}

func newTestGraph(n int) *graph.Graph {
	g := &graph.Graph{N: n, Data: make([]float64, n*n*n)}
	for prev := 0; prev < n; prev++ {
		for curr := 0; curr < n; curr++ {
			for next := 0; next < n; next++ {
				if prev == curr && curr == next {
					continue
				}
				g.Data[prev*n*n+curr*n+next] = float64(1 + prev + curr + next)
			}
		}
	}
	return g
}
