package repository

import (
	"context"
	"testing"
	"testing/fstest"
)

func TestLoadEvolutionParsesJSONLFrames(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{
		"results/evolution/run-1.jsonl": {
			Data: []byte("{" +
				"\"iter\":1,\"eval_count\":10,\"best_makespan\":11.2,\"best_sequence\":[1,2,3]}\n" +
				"{\"iter\":5,\"eval_count\":32,\"best_makespan\":9.8,\"best_sequence\":[1,3,2]}\n"),
		},
	})

	frames, err := repo.LoadEvolution(context.Background(), "run-1")
	if err != nil {
		t.Fatalf("LoadEvolution returned error: %v", err)
	}

	if len(frames) != 2 {
		t.Fatalf("expected 2 frames, got %d", len(frames))
	}

	if frames[0].Iter != 1 || frames[0].EvalCount != 10 || len(frames[0].BestSequence) != 3 {
		t.Fatalf("unexpected first frame: %+v", frames[0])
	}

	if frames[1].Iter != 5 || frames[1].EvalCount != 32 {
		t.Fatalf("unexpected second frame: %+v", frames[1])
	}
}

func TestLoadEvolutionHandlesMissingFile(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{})
	frames, err := repo.LoadEvolution(context.Background(), "run-missing")
	if err != nil {
		t.Fatalf("LoadEvolution returned error: %v", err)
	}

	if len(frames) != 0 {
		t.Fatalf("expected no frames, got %d", len(frames))
	}
}

func TestLoadEvolutionReturnsErrorForInvalidLine(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{
		"results/evolution/run-1.jsonl": {
			Data: []byte("{\"iter\":1,\"eval_count\":10,\"best_makespan\":11.2,\"best_sequence\":[1,2,3]}\n{invalid json}\n"),
		},
	})

	_, err := repo.LoadEvolution(context.Background(), "run-1")
	if err == nil {
		t.Fatal("expected LoadEvolution to return an error")
	}
}
