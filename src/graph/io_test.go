package graph

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveSkipsExistingFiles(t *testing.T) {
	folder := t.TempDir()

	graphs := []*Graph{
		{N: 2, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8}},
		{N: 2, Data: []float64{5, 6, 7, 8, 9, 10, 11, 12}},
	}

	first, err := Save(graphs, folder)
	if err != nil {
		t.Fatalf("first save failed: %v", err)
	}
	if first.Created != 2 || first.Skipped != 0 {
		t.Fatalf("unexpected first report: %+v", first)
	}

	filePath := filepath.Join(folder, "2a.graph")
	original, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("read original file: %v", err)
	}

	updatedGraphs := []*Graph{
		{N: 2, Data: []float64{9, 10, 11, 12, 13, 14, 15, 16}},
		{N: 2, Data: []float64{13, 14, 15, 16, 17, 18, 19, 20}},
	}

	second, err := Save(updatedGraphs, folder)
	if err != nil {
		t.Fatalf("second save failed: %v", err)
	}
	if second.Created != 0 || second.Skipped != 2 {
		t.Fatalf("unexpected second report: %+v", second)
	}

	after, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("read file after second save: %v", err)
	}

	if !bytes.Equal(original, after) {
		t.Fatal("expected existing graph file to remain unchanged")
	}
}
