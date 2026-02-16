package repository

import (
	"context"
	"testing"
	"testing/fstest"
)

func TestListRunsParsesIterations(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{
		"results/summary/10a__aco__s1__h1.json": {
			Data: []byte(`{"run_id":"10a__aco__s1__h1","method":"aco","seed":1,"params":{"iterations":100},"result":{"best_makespan":12.5}}`),
		},
		"results/summary/10a__ga__s2__h2.json": {
			Data: []byte(`{"run_id":"10a__ga__s2__h2","method":"ga","seed":2,"params":{},"result":{"best_makespan":13.5}}`),
		},
		"results/summary/10a__pso__s3__h3.json": {
			Data: []byte(`{"run_id":"10a__pso__s3__h3","method":"pso","seed":3,"params":{"iterations":"250"},"result":{"best_makespan":11.5}}`),
		},
		"results/summary/20a__aco__s1__h4.json": {
			Data: []byte(`{"run_id":"20a__aco__s1__h4","method":"aco","seed":1,"params":{"iterations":50},"result":{"best_makespan":9.5}}`),
		},
	})

	runs, err := repo.ListRuns(context.Background(), "10a")
	if err != nil {
		t.Fatalf("ListRuns returned error: %v", err)
	}

	if len(runs) != 3 {
		t.Fatalf("expected 3 runs, got %d", len(runs))
	}

	if runs[0].RunID != "10a__aco__s1__h1" || runs[0].Iterations != 100 {
		t.Fatalf("unexpected first run: %+v", runs[0])
	}

	if runs[1].RunID != "10a__ga__s2__h2" || runs[1].Iterations != 0 {
		t.Fatalf("unexpected second run: %+v", runs[1])
	}

	if runs[2].RunID != "10a__pso__s3__h3" || runs[2].Iterations != 250 {
		t.Fatalf("unexpected third run: %+v", runs[2])
	}
}

func TestListRunsHandlesMissingSummaryDir(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{})
	runs, err := repo.ListRuns(context.Background(), "10a")
	if err != nil {
		t.Fatalf("ListRuns returned error: %v", err)
	}

	if len(runs) != 0 {
		t.Fatalf("expected no runs, got %d", len(runs))
	}
}

func TestListRunsReturnsErrorForInvalidSummary(t *testing.T) {
	t.Parallel()

	repo := NewEmbeddedDataRepository(fstest.MapFS{
		"results/summary/10a__aco__s1__h1.json": {
			Data: []byte(`{"run_id":"10a__aco__s1__h1"`),
		},
	})

	_, err := repo.ListRuns(context.Background(), "10a")
	if err == nil {
		t.Fatal("expected ListRuns to return an error")
	}
}
