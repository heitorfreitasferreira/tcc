package repository

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strings"

	"tcc/shared/reporting"
)

// LoadEvolution reads a JSONL file (results/evolution/<runID>.jsonl) and returns
// all evolution frames. Each line is a JSON object with iter, eval_count, best_makespan,
// and best_sequence. Returns nil, nil if the file does not exist (missing evolution data
// is not an error because brute-force runs do not produce evolution files).
func (r *EmbeddedDataRepository) LoadEvolution(ctx context.Context, runID string) ([]EvolutionFrame, error) {
	runID = strings.TrimSpace(runID)
	if runID == "" {
		return nil, nil
	}

	filePath := path.Join(evolutionDir, runID+".jsonl")
	file, err := r.dataFS.Open(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}

		return nil, fmt.Errorf("open evolution for run %q: %w", runID, err)
	}
	defer file.Close()

	frames := make([]EvolutionFrame, 0)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	lineNumber := 0
	for scanner.Scan() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var record reporting.EvolutionRecord
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, fmt.Errorf("decode evolution frame line %d for run %q: %w", lineNumber, runID, err)
		}

		frames = append(frames, EvolutionFrame{
			Iter:         record.Iteration,
			EvalCount:    record.EvalCount,
			BestMakespan: record.BestMakespan,
			BestSequence: append([]int(nil), record.BestSequence...),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan evolution for run %q: %w", runID, err)
	}

	return frames, nil
}
