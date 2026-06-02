package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"tcc/graph"
	"tcc/points"
	"tcc/web/repository"
)

func TestRenderHandlerLoadsMultipleRuns(t *testing.T) {
	maps := fakeMapsReader{points: points.Points2D{{0, 0}, {0.1, 0.1}, {0.2, 0.2}}}
	evolution := &fakeEvolutionReader{frames: map[string][]repository.EvolutionFrame{
		"run-a": {{Iter: 1, BestMakespan: 1.0, BestSequence: []int{1, 2}}},
		"run-b": {{Iter: 1, BestMakespan: 2.0, BestSequence: []int{2, 1}}},
		"run-c": {{Iter: 1, BestMakespan: 3.0, BestSequence: []int{1, 2}}},
	}}

	req := httptest.NewRequest(http.MethodGet, "/api/render?map=10a&run=run-a&run2=run-b&run3=run-c", nil)
	rec := httptest.NewRecorder()

	NewRenderHandler(maps, evolution).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Type"); got != "image/png" {
		t.Fatalf("Content-Type = %q, want image/png", got)
	}
	if !sameStrings(evolution.loaded, []string{"run-a", "run-b", "run-c"}) {
		t.Fatalf("loaded runs = %v", evolution.loaded)
	}
}

type fakeMapsReader struct {
	points points.Points2D
}

func (f fakeMapsReader) ListMaps(context.Context) ([]string, error) { return nil, nil }
func (f fakeMapsReader) LoadPoints(context.Context, string) (points.Points2D, error) {
	return f.points, nil
}
func (f fakeMapsReader) LoadGraph(context.Context, string) (graph.Graph, error) { return nil, nil }

type fakeEvolutionReader struct {
	frames map[string][]repository.EvolutionFrame
	loaded []string
}

func (f *fakeEvolutionReader) LoadEvolution(_ context.Context, runID string) ([]repository.EvolutionFrame, error) {
	f.loaded = append(f.loaded, runID)
	return f.frames[runID], nil
}

func sameStrings(left, right []string) bool {
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
