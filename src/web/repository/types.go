// Package repository provides read-only access to experiment artifacts embedded in
// the binary. It implements three reader interfaces (MapsReader, SummaryReader,
// EvolutionReader) over an fs.FS backed by Go's embed directive.
//
// Artifacts are stored under src/data/results/ and follow the naming conventions
// described in AGENTS-experiments.md. The repository layer abstracts the filesystem
// layout so that service and handler layers need only deal with map IDs, run IDs,
// and typed records.
package repository

import (
	"context"

	"tcc/graph"
	"tcc/points"
)

// RunRecord is a summary-level descriptor for a single optimization run.
type RunRecord struct {
	RunID        string
	Method       string
	Seed         int64
	BestMakespan float64
	Iterations   int
}

// EvolutionFrame is a snapshot of optimizer state at a single iteration checkpoint.
// Fields are tagged for direct JSON serialization to the frontend.
type EvolutionFrame struct {
	Iter         int     `json:"iter"`
	EvalCount    int     `json:"eval_count"`
	BestMakespan float64 `json:"best_makespan"`
	BestSequence []int   `json:"best_sequence"`
}

// MapsReader provides access to TSP map definitions and their graph representations.
// MapsReader provides access to TSP map definitions and their graph representations.
type MapsReader interface {
	// ListMaps returns all available map IDs (sorted, natural order).
	ListMaps(ctx context.Context) ([]string, error)
	// LoadPoints returns the 2D point set for a map.
	LoadPoints(ctx context.Context, mapID string) (points.Points2D, error)
	// LoadGraph returns the adjacency/distance graph for a map.
	LoadGraph(ctx context.Context, mapID string) (graph.Graph, error)
}

// SummaryReader lists run summaries for a given map.
// SummaryReader lists run summaries for a given map.
type SummaryReader interface {
	// ListRuns returns all runs whose run_id starts with mapID (prefix match).
	ListRuns(ctx context.Context, mapID string) ([]RunRecord, error)
}

// EvolutionReader loads evolution frame data for a run.
type EvolutionReader interface {
	// LoadEvolution returns the full sequence of evolution frames for a run.
	LoadEvolution(ctx context.Context, runID string) ([]EvolutionFrame, error)
}
