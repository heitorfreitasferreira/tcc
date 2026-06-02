package render

import "testing"

func TestBuildRoutesNormalizesMultipleSequences(t *testing.T) {
	routes := BuildRoutes(5, [][]int{
		{1, 2, 3, 4},
		{0, 4, 3, 2, 1, 0},
		{9, -1},
	})

	want := [][]int{
		{0, 1, 2, 3, 4, 0},
		{0, 4, 3, 2, 1, 0},
	}

	if len(routes) != len(want) {
		t.Fatalf("BuildRoutes() length = %d, want %d", len(routes), len(want))
	}

	for i := range want {
		if routes[i].Name != defaultRouteNames[i] {
			t.Fatalf("routes[%d].Name = %q, want %q", i, routes[i].Name, defaultRouteNames[i])
		}
		if routes[i].Color != defaultRouteColors[i] {
			t.Fatalf("routes[%d].Color = %#v, want %#v", i, routes[i].Color, defaultRouteColors[i])
		}
		if !sameInts(routes[i].Indices, want[i]) {
			t.Fatalf("routes[%d].Indices = %v, want %v", i, routes[i].Indices, want[i])
		}
	}
}

func TestBuildNamedRoutesKeepsProvidedMetadata(t *testing.T) {
	routes := BuildNamedRoutes(4, []RouteSpec{
		{Name: "ACO", Sequence: []int{1, 2, 3}, Makespan: 1.25, Color: RGB{0.1, 0.2, 0.3}},
	})

	if len(routes) != 1 {
		t.Fatalf("BuildNamedRoutes() length = %d, want 1", len(routes))
	}
	if routes[0].Name != "ACO" {
		t.Fatalf("Name = %q, want ACO", routes[0].Name)
	}
	if routes[0].Makespan != 1.25 {
		t.Fatalf("Makespan = %v, want 1.25", routes[0].Makespan)
	}
	if routes[0].Color != (RGB{0.1, 0.2, 0.3}) {
		t.Fatalf("Color = %#v, want %#v", routes[0].Color, RGB{0.1, 0.2, 0.3})
	}
	if !sameInts(routes[0].Indices, []int{0, 1, 2, 3, 0}) {
		t.Fatalf("Indices = %v", routes[0].Indices)
	}
}

func TestDefaultOptionsIncludeScale(t *testing.T) {
	opts := DefaultOpts()
	if opts.ScaleStep <= 0 {
		t.Fatalf("ScaleStep = %v, want positive", opts.ScaleStep)
	}
	if opts.ScaleLabel == "" {
		t.Fatal("ScaleLabel is empty")
	}
}

func sameInts(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
