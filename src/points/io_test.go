package points

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveSkipsExistingFiles(t *testing.T) {
	folder := t.TempDir()

	instances := []Points2D{
		{{0, 0}, {1, 1}},
		{{2, 2}, {3, 3}},
	}

	first, err := Save(instances, folder)
	if err != nil {
		t.Fatalf("first save failed: %v", err)
	}
	if first.Created != 2 || first.Skipped != 0 {
		t.Fatalf("unexpected first report: %+v", first)
	}

	filePath := filepath.Join(folder, "2a.points")
	original, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("read original file: %v", err)
	}

	updatedInstances := []Points2D{
		{{9, 9}, {8, 8}},
		{{7, 7}, {6, 6}},
	}

	second, err := Save(updatedInstances, folder)
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
		t.Fatal("expected existing point file to remain unchanged")
	}
}
