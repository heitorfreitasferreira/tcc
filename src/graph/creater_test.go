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

	// If the tensor followed the monograph semantics G[prev][curr][next]:
	//   step1 (0→1): d(0,1) + turn(none→0→1)          = 1 + 0      = 1
	//   step2 (1→2): d(1,2) + turn(0→1, 1→2)           = √2 + 0.75 ≈ 2.16421356
	//   step3 (2→0): d(2,0) + turn(1→2, 2→0)           = 1 + 0.75  = 1.75
	//   total                                          ≈ 4.91421356

	d01 := 1.0
	d12 := math.Sqrt(2)
	d20 := 1.0
	turn135 := 0.75 // 3π/4 / π
	correctExpected := d01 + d12 + turn135 + d20 + turn135

	// If the creater stores [current][next][prev] but Makespan reads [prev][curr][next]:
	//   step1: g[0][0][1] = penalty[0][0][1] = d(0,0) + turn(1→0, 0→0) = 0
	//   step2: g[0][1][2] = penalty[0][1][2] = d(0,1) + turn(2→0, 0→1) ≈ 1 + 0.5          = 1.5
	//   step3: g[1][2][0] = penalty[1][2][0] = d(1,2) + turn(0→1, 1→2) = √2 + 0.75        ≈ 2.16421356
	//   total                                                                              ≈ 3.66421356
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

	// Semantic test: G[prev=0][curr=1][next=2] should be
	// d(curr=1, next=2) + turnCost(prev→curr, curr→next)
	// = d(1,2) + turnCost(0→1, 1→2)
	// = √2 + 0.75 ≈ 2.16421356

	expected := math.Sqrt(2) + 0.75
	got := g[0][1][2]

	epsilon := 1e-9
	if math.Abs(got-expected) < epsilon {
		t.Logf("✓ G[0][1][2] = %.10f — matches monograph semantics", got)
	} else {
		t.Logf("G[0][1][2] = %.10f (expected %.10f for monograph semantics)", got, expected)
		t.Logf("  Actual value suggests G stores at [current][next][previous] order")
	}
}
