package service

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"tcc/graph"
	"tcc/points"
	"tcc/web/repository"
)

type mapsRepoStub struct {
	maps           []string
	pointsByMap    map[string]points.Points2D
	graphByMap     map[string]graph.Graph
	listErr        error
	loadPointsErr  error
	loadGraphErr   error
	loadPointsCall int
	loadGraphCall  int
}

func (s *mapsRepoStub) ListMaps(_ context.Context) ([]string, error) {
	if s.listErr != nil {
		return nil, s.listErr
	}

	return append([]string(nil), s.maps...), nil
}

func (s *mapsRepoStub) LoadPoints(_ context.Context, mapID string) (points.Points2D, error) {
	s.loadPointsCall++
	if s.loadPointsErr != nil {
		return nil, s.loadPointsErr
	}

	loaded, ok := s.pointsByMap[mapID]
	if !ok {
		return nil, errors.New("points not found")
	}

	return append(points.Points2D(nil), loaded...), nil
}

func (s *mapsRepoStub) LoadGraph(_ context.Context, mapID string) (graph.Graph, error) {
	s.loadGraphCall++
	if s.loadGraphErr != nil {
		return nil, s.loadGraphErr
	}

	loaded, ok := s.graphByMap[mapID]
	if !ok {
		return nil, errors.New("graph not found")
	}

	return loaded, nil
}

type summaryRepoStub struct {
	runsByMap map[string][]repository.RunRecord
	err       error
}

func (s *summaryRepoStub) ListRuns(_ context.Context, mapID string) ([]repository.RunRecord, error) {
	if s.err != nil {
		return nil, s.err
	}

	loaded := s.runsByMap[mapID]
	return append([]repository.RunRecord(nil), loaded...), nil
}

type evolutionRepoStub struct {
	framesByRun map[string][]repository.EvolutionFrame
	err         error
	loadCall    int
}

func (s *evolutionRepoStub) LoadEvolution(_ context.Context, runID string) ([]repository.EvolutionFrame, error) {
	s.loadCall++
	if s.err != nil {
		return nil, s.err
	}

	loaded := s.framesByRun[runID]
	return append([]repository.EvolutionFrame(nil), loaded...), nil
}

type payloadForTest struct {
	RunID      string                      `json:"run_id"`
	Iterations int                         `json:"iterations"`
	Frames     []repository.EvolutionFrame `json:"frames"`
}

func TestBuildPageLoadsEvolutionForSelectedRun(t *testing.T) {
	t.Parallel()

	runID := "10a__aco__s1__h111"
	mapsRepo := &mapsRepoStub{
		maps: []string{"10a"},
		pointsByMap: map[string]points.Points2D{
			"10a": {
				{0, 0},
				{0.3, 0.2},
				{-0.2, 0.7},
			},
		},
		graphByMap: map[string]graph.Graph{
			"10a": {
				{{1}, {2}, {3}},
				{{4}, {5}, {6}},
				{{7}, {8}, {9}},
			},
		},
	}

	summaryRepo := &summaryRepoStub{
		runsByMap: map[string][]repository.RunRecord{
			"10a": {
				{
					RunID:        runID,
					Method:       "aco",
					Seed:         1,
					BestMakespan: 12.25,
					Iterations:   100,
				},
			},
		},
	}

	evolutionRepo := &evolutionRepoStub{
		framesByRun: map[string][]repository.EvolutionFrame{
			runID: {
				{Iter: 1, EvalCount: 12, BestMakespan: 20.4, BestSequence: []int{1, 2}},
				{Iter: 8, EvalCount: 50, BestMakespan: 19.1, BestSequence: []int{2, 1}},
			},
		},
	}

	service := NewPageService(mapsRepo, summaryRepo, evolutionRepo)
	data, err := service.BuildPage(context.Background(), SelectionState{
		SelectedMap:    "10a",
		ExpandedMap:    "10a",
		SelectedMethod: "aco",
		ExpandedMethod: "aco",
		SelectedRun:    runID,
	})
	if err != nil {
		t.Fatalf("BuildPage returned error: %v", err)
	}

	if data.SelectedRun != runID {
		t.Fatalf("expected selected run %q, got %q", runID, data.SelectedRun)
	}

	if data.SelectedRunIterations != 100 {
		t.Fatalf("expected selected run iterations 100, got %d", data.SelectedRunIterations)
	}

	if data.EvolutionFramesCount != 2 {
		t.Fatalf("expected 2 evolution frames, got %d", data.EvolutionFramesCount)
	}

	if len(data.MethodGroups) != 1 || len(data.MethodGroups[0].Runs) != 1 {
		t.Fatalf("unexpected method groups: %+v", data.MethodGroups)
	}

	if data.MethodGroups[0].Runs[0].Iterations != 100 {
		t.Fatalf("expected iterations propagated to run option, got %+v", data.MethodGroups[0].Runs[0])
	}

	var payload payloadForTest
	if err := json.Unmarshal([]byte(data.EvolutionJSON), &payload); err != nil {
		t.Fatalf("failed to unmarshal evolution payload: %v", err)
	}

	if payload.RunID != runID || payload.Iterations != 100 || len(payload.Frames) != 2 {
		t.Fatalf("unexpected evolution payload: %+v", payload)
	}

	if mapsRepo.loadGraphCall != 1 {
		t.Fatalf("expected graph load once, got %d", mapsRepo.loadGraphCall)
	}

	if evolutionRepo.loadCall != 1 {
		t.Fatalf("expected evolution load once, got %d", evolutionRepo.loadCall)
	}
}

func TestBuildPageSkipsEvolutionWhenRunIsInvalid(t *testing.T) {
	t.Parallel()

	mapsRepo := &mapsRepoStub{
		maps: []string{"10a"},
		pointsByMap: map[string]points.Points2D{
			"10a": {
				{0, 0},
				{0.3, 0.2},
			},
		},
		graphByMap: map[string]graph.Graph{
			"10a": {
				{{1}, {2}},
				{{3}, {4}},
			},
		},
	}

	summaryRepo := &summaryRepoStub{
		runsByMap: map[string][]repository.RunRecord{
			"10a": {
				{RunID: "run-valid", Method: "aco", Seed: 1, BestMakespan: 10.5, Iterations: 100},
			},
		},
	}

	evolutionRepo := &evolutionRepoStub{
		framesByRun: map[string][]repository.EvolutionFrame{
			"run-valid": {
				{Iter: 1, EvalCount: 10, BestMakespan: 10, BestSequence: []int{1}},
			},
		},
	}

	service := NewPageService(mapsRepo, summaryRepo, evolutionRepo)
	data, err := service.BuildPage(context.Background(), SelectionState{
		SelectedMap:    "10a",
		ExpandedMap:    "10a",
		SelectedMethod: "aco",
		ExpandedMethod: "aco",
		SelectedRun:    "run-missing",
	})
	if err != nil {
		t.Fatalf("BuildPage returned error: %v", err)
	}

	if data.SelectedRun != "" {
		t.Fatalf("expected no selected run, got %q", data.SelectedRun)
	}

	if data.EvolutionJSON != "" || data.EvolutionFramesCount != 0 {
		t.Fatalf("expected no evolution payload, got JSON=%q frames=%d", data.EvolutionJSON, data.EvolutionFramesCount)
	}

	if mapsRepo.loadGraphCall != 0 {
		t.Fatalf("expected graph load to be skipped, got %d", mapsRepo.loadGraphCall)
	}

	if evolutionRepo.loadCall != 0 {
		t.Fatalf("expected evolution load to be skipped, got %d", evolutionRepo.loadCall)
	}
}

func TestBuildPageReturnsBaseDataForInvalidMap(t *testing.T) {
	t.Parallel()

	mapsRepo := &mapsRepoStub{
		maps: []string{"10a"},
	}

	service := NewPageService(mapsRepo, &summaryRepoStub{}, &evolutionRepoStub{})
	data, err := service.BuildPage(context.Background(), SelectionState{SelectedMap: "99z"})
	if err != nil {
		t.Fatalf("BuildPage returned error: %v", err)
	}

	if data.SelectedMap != "" {
		t.Fatalf("expected no selected map, got %q", data.SelectedMap)
	}

	if mapsRepo.loadPointsCall != 0 {
		t.Fatalf("expected points not to be loaded, got %d", mapsRepo.loadPointsCall)
	}
}
