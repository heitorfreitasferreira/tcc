package repository

import (
	"context"

	"tcc/graph"
	"tcc/points"
)

type RunRecord struct {
	RunID        string
	Method       string
	Seed         int64
	BestMakespan float64
	Iterations   int
}

type EvolutionFrame struct {
	Iter         int     `json:"iter"`
	EvalCount    int     `json:"eval_count"`
	BestMakespan float64 `json:"best_makespan"`
	BestSequence []int   `json:"best_sequence"`
}

type MapsReader interface {
	ListMaps(ctx context.Context) ([]string, error)
	LoadPoints(ctx context.Context, mapID string) (points.Points2D, error)
	LoadGraph(ctx context.Context, mapID string) (graph.Graph, error)
}

type SummaryReader interface {
	ListRuns(ctx context.Context, mapID string) ([]RunRecord, error)
}

type EvolutionReader interface {
	LoadEvolution(ctx context.Context, runID string) ([]EvolutionFrame, error)
}
