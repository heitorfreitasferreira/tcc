package graph

import (
	"math"
	"testing"

	"tcc/points"
)

func TestTensorIndexing(t *testing.T) {
	pts := points.Points2D{
		{0, 0}, // node 0: base
		{1, 0}, // node 1: right
		{0, 1}, // node 2: up
	}

	g := New(pts)
	route := []int{1, 2} // 0 → 1 → 2 → 0
	got := g.Makespan(route)

	d01 := 1.0
	d12 := math.Sqrt(2)
	d20 := 1.0
	turn135 := 0.75 // 3π/4 / π
	correctExpected := d01 + d12 + turn135 + d20 + turn135

	buggyExpected := 0.0 + 1.5 + d12 + turn135

	t.Logf("=== Tensor Indexing Validation ===")
	t.Logf("Points: P0=(0,0) P1=(1,0) P2=(0,1)")
	t.Logf("Route: 0 → 1 → 2 → 0")
	t.Logf("")
	t.Logf("If CORRECT (monograph semantics G[prev][curr][next]):")
	t.Logf("  g[0][0][1] = d(0,1) = 1.000000")
	t.Logf("  g[0][1][2] = d(1,2)+turn(0→1,1→2) = √2 + 0.75 = %.10f", d12+turn135)
	t.Logf("  g[1][2][0] = d(2,0)+turn(1→2,2→0) = 1 + 0.75 = %.10f", d20+turn135)
	t.Logf("  Total = %.10f", correctExpected)
	t.Logf("")
	t.Logf("If BUGGY (creater stores [curr][next][prev], Makespan reads [prev][curr][next]):")
	t.Logf("  g[0][0][1] = 0.000000  (d(0,0)+turn(1→0,0→0)=0)")
	t.Logf("  g[0][1][2] = 1.500000  (d(0,1)+turn(2→0,0→1)=1+0.5)")
	t.Logf("  g[1][2][0] = %.10f  (d(1,2)+turn(0→1,1→2)=√2+0.75)", d12+turn135)
	t.Logf("  Total = %.10f", buggyExpected)
	t.Logf("")
	t.Logf("Actual makespan: %.10f", got)

	epsilon := 1e-9
	if math.Abs(got-correctExpected) < epsilon {
		t.Logf("✓ RESULT MATCHES CORRECT SEMANTICS — tensor indexing is correct")
	} else if math.Abs(got-buggyExpected) < epsilon {
		t.Errorf("✗ RESULT MATCHES BUGGY INDEXING — creater stores [current][next][previous]\n"+
			"  but Makespan reads [previous][current][next]. Fix required.")
	} else {
		t.Errorf("✗ Unexpected makespan %.10f — neither correct (%.10f) nor buggy (%.10f)",
			got, correctExpected, buggyExpected)
	}
}

func TestTensorCellSemantics(t *testing.T) {
	pts := points.Points2D{
		{0, 0},
		{1, 0},
		{0, 1},
	}
	g := New(pts)

	expected := math.Sqrt(2) + 0.75
	got := g.At(0, 1, 2)

	epsilon := 1e-9
	if math.Abs(got-expected) < epsilon {
		t.Logf("✓ G[0][1][2] = %.10f — matches monograph semantics", got)
	} else {
		t.Logf("G[0][1][2] = %.10f (expected %.10f for monograph semantics)", got, expected)
		t.Logf("  Actual value suggests G stores at [current][next][previous] order")
	}
}
