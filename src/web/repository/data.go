package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"path"
	"sort"
	"strconv"
	"strings"

	"tcc/points"
	"tcc/shared/reporting"
)

const summaryDir = "results/summary"

type RunRecord struct {
	RunID        string
	Method       string
	Seed         int64
	BestMakespan float64
}

type Reader interface {
	ListMaps(ctx context.Context) ([]string, error)
	LoadPoints(ctx context.Context, mapID string) (points.Points2D, error)
	ListRuns(ctx context.Context, mapID string) ([]RunRecord, error)
}

type EmbeddedDataRepository struct {
	dataFS fs.FS
}

func NewEmbeddedDataRepository(dataFS fs.FS) *EmbeddedDataRepository {
	return &EmbeddedDataRepository{dataFS: dataFS}
}

func (r *EmbeddedDataRepository) ListMaps(_ context.Context) ([]string, error) {
	entries, err := fs.ReadDir(r.dataFS, ".")
	if err != nil {
		return nil, err
	}

	graphFiles := make(map[string]struct{})
	pointFiles := make(map[string]struct{})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		extension := path.Ext(name)
		baseName := strings.TrimSuffix(name, extension)

		switch extension {
		case ".graph":
			graphFiles[baseName] = struct{}{}
		case ".points":
			pointFiles[baseName] = struct{}{}
		}
	}

	ids := make([]string, 0, len(graphFiles))
	for id := range graphFiles {
		if _, ok := pointFiles[id]; ok {
			ids = append(ids, id)
		}
	}

	sort.Slice(ids, func(i, j int) bool {
		leftNum, leftSuffix, leftOK := splitMapID(ids[i])
		rightNum, rightSuffix, rightOK := splitMapID(ids[j])

		if leftOK && rightOK {
			if leftNum != rightNum {
				return leftNum < rightNum
			}

			if leftSuffix != rightSuffix {
				return leftSuffix < rightSuffix
			}

			return ids[i] < ids[j]
		}

		if leftOK {
			return true
		}

		if rightOK {
			return false
		}

		return ids[i] < ids[j]
	})

	return ids, nil
}

func (r *EmbeddedDataRepository) LoadPoints(_ context.Context, mapID string) (points.Points2D, error) {
	rawPoints, err := fs.ReadFile(r.dataFS, mapID+".points")
	if err != nil {
		return nil, err
	}

	var loaded points.Points2D
	if err := json.Unmarshal(rawPoints, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}

func (r *EmbeddedDataRepository) ListRuns(_ context.Context, mapID string) ([]RunRecord, error) {
	entries, err := fs.ReadDir(r.dataFS, summaryDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	prefix := mapID + "__"
	runs := make([]RunRecord, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		if path.Ext(fileName) != ".json" || !strings.HasPrefix(fileName, prefix) {
			continue
		}

		rawSummary, err := fs.ReadFile(r.dataFS, path.Join(summaryDir, fileName))
		if err != nil {
			return nil, err
		}

		var artifact reporting.RunSummary
		if err := json.Unmarshal(rawSummary, &artifact); err != nil {
			return nil, err
		}

		runID := strings.TrimSpace(artifact.RunID)
		if runID == "" {
			runID = strings.TrimSuffix(fileName, ".json")
		}

		runs = append(runs, RunRecord{
			RunID:        runID,
			Method:       strings.TrimSpace(artifact.Method),
			Seed:         artifact.Seed,
			BestMakespan: artifact.Result.BestMakespan,
		})
	}

	sort.Slice(runs, func(i, j int) bool {
		if runs[i].Method != runs[j].Method {
			return runs[i].Method < runs[j].Method
		}

		if runs[i].Seed != runs[j].Seed {
			return runs[i].Seed < runs[j].Seed
		}

		return runs[i].RunID < runs[j].RunID
	})

	return runs, nil
}

func splitMapID(mapID string) (int, string, bool) {
	if mapID == "" {
		return 0, "", false
	}

	split := 0
	for split < len(mapID) {
		r := mapID[split]
		if r < '0' || r > '9' {
			break
		}
		split++
	}

	if split == 0 || split == len(mapID) {
		return 0, "", false
	}

	number, err := strconv.Atoi(mapID[:split])
	if err != nil {
		return 0, "", false
	}

	return number, mapID[split:], true
}
