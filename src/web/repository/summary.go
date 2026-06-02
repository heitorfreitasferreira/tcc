package repository

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"math"
	"path"
	"slices"
	"strconv"
	"strings"

	"tcc/shared/reporting"
)

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
			Iterations:   extractIterations(artifact.Params),
		})
	}

	slices.SortFunc(runs, func(a, b RunRecord) int {
		if order := cmp.Compare(a.Method, b.Method); order != 0 {
			return order
		}

		if order := cmp.Compare(a.Seed, b.Seed); order != 0 {
			return order
		}

		return cmp.Compare(a.RunID, b.RunID)
	})

	return runs, nil
}

func extractIterations(params map[string]any) int {
	if len(params) == 0 {
		return 0
	}

	raw, ok := params["iterations"]
	if !ok {
		return 0
	}

	iterations, ok := toInt(raw)
	if !ok || iterations < 0 {
		return 0
	}

	return iterations
}

func toInt(value any) (int, bool) {
	switch typed := value.(type) {
	case int:
		return typed, true
	case int64:
		if !fitsInt64ToInt(typed) {
			return 0, false
		}
		return int(typed), true
	case float64:
		if math.IsNaN(typed) || math.IsInf(typed, 0) {
			return 0, false
		}

		truncated := math.Trunc(typed)
		if truncated > maxInt64Float || truncated < minInt64Float {
			return 0, false
		}

		parsed := int64(truncated)
		if !fitsInt64ToInt(parsed) {
			return 0, false
		}

		return int(parsed), true
	case string:
		parsed, err := strconv.Atoi(strings.TrimSpace(typed))
		if err != nil {
			return 0, false
		}
		return parsed, true
	case json.Number:
		parsed, err := typed.Int64()
		if err == nil {
			if !fitsInt64ToInt(parsed) {
				return 0, false
			}
			return int(parsed), true
		}

		parsedFloat, err := typed.Float64()
		if err != nil || math.IsNaN(parsedFloat) || math.IsInf(parsedFloat, 0) {
			return 0, false
		}

		truncated := math.Trunc(parsedFloat)
		if truncated > maxInt64Float || truncated < minInt64Float {
			return 0, false
		}

		parsed = int64(truncated)
		if !fitsInt64ToInt(parsed) {
			return 0, false
		}

		return int(parsed), true
	default:
		return 0, false
	}
}

const (
	maxInt64Float = float64(^uint64(0) >> 1)
	minInt64Float = -maxInt64Float - 1
)

func fitsInt64ToInt(value int64) bool {
	if strconv.IntSize == 64 {
		return true
	}

	return value >= math.MinInt32 && value <= math.MaxInt32
}
