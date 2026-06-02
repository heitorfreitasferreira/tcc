package reporting

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const (
	SummarySchema   = "tcc.summary.v1"
	EvolutionSchema = "tcc.evolution.v1"
	TimingSchema    = "tcc.timing.v1"
)

const (
	IfExistsSkip      = "skip"
	IfExistsOverwrite = "overwrite"
	IfExistsError     = "error"
)

type Paths struct {
	BaseDir   string
	RunID     string
	Summary   string
	Evolution string
	Timing    string
	Log       string
}

type RunSummary struct {
	Schema        string         `json:"schema"`
	RunID         string         `json:"run_id"`
	Method        string         `json:"method"`
	Instance      string         `json:"instance"`
	Seed          int64          `json:"seed"`
	Params        map[string]any `json:"params"`
	Status        string         `json:"status"`
	Result        SummaryResult  `json:"result"`
	TimingFile    string         `json:"timing_file"`
	EvolutionFile string         `json:"evolution_file"`
	StartedAt     string         `json:"started_at"`
	FinishedAt    string         `json:"finished_at"`
}

type SummaryResult struct {
	BestMakespan        float64 `json:"best_makespan"`
	BestSequence        []int   `json:"best_sequence"`
	IterationsCompleted int     `json:"iterations_completed"`
	Evaluations         int     `json:"evaluations"`
}

type EvolutionRecord struct {
	Schema       string  `json:"schema"`
	RunID        string  `json:"run_id"`
	Method       string  `json:"method"`
	Iteration    int     `json:"iter"`
	EvalCount    int     `json:"eval_count"`
	BestMakespan float64 `json:"best_makespan"`
	Delta        float64 `json:"delta"`
	BestSequence []int   `json:"best_sequence"`
}

type RunTiming struct {
	Schema    string       `json:"schema"`
	RunID     string       `json:"run_id"`
	Method    string       `json:"method"`
	Instance  string       `json:"instance"`
	Durations DurationMS   `json:"durations_ms"`
	Measured  TimeMeasured `json:"measured"`
}

type DurationMS struct {
	LoadInstance int64 `json:"load_instance"`
	Optimize     int64 `json:"optimize"`
	Serialize    int64 `json:"serialize"`
	Total        int64 `json:"total"`
}

type TimeMeasured struct {
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}

type artifactState struct {
	Summary   bool
	Evolution bool
	Timing    bool
}

func (s artifactState) Any() bool {
	return s.Summary || s.Evolution || s.Timing
}

func (s artifactState) Complete() bool {
	return s.Summary && s.Evolution && s.Timing
}

func (s artifactState) Missing() []string {
	missing := make([]string, 0, 3)
	if !s.Summary {
		missing = append(missing, "summary")
	}
	if !s.Evolution {
		missing = append(missing, "evolution")
	}
	if !s.Timing {
		missing = append(missing, "timing")
	}
	return missing
}

func ParseIfExistsPolicy(policy string) (string, error) {
	switch policy {
	case IfExistsSkip, IfExistsOverwrite, IfExistsError:
		return policy, nil
	default:
		return "", fmt.Errorf("invalid --if-exists %q, expected one of: skip, overwrite, error", policy)
	}
}

func ValidateRunID(runID string) error {
	if runID == "" {
		return fmt.Errorf("run id cannot be empty")
	}

	if filepath.IsAbs(runID) {
		return fmt.Errorf("invalid --run-id %q: absolute paths are not allowed", runID)
	}

	if strings.Contains(runID, "/") || strings.Contains(runID, `\\`) {
		return fmt.Errorf("invalid --run-id %q: path separators are not allowed", runID)
	}

	if sanitizeID(runID) != runID {
		return fmt.Errorf("invalid --run-id %q: use only letters, numbers, '-' and '_'", runID)
	}

	return nil
}

func PathsForRun(resultsDir, runID string) Paths {
	return Paths{
		BaseDir:   resultsDir,
		RunID:     runID,
		Summary:   filepath.Join(resultsDir, "summary", runID+".json"),
		Evolution: filepath.Join(resultsDir, "evolution", runID+".jsonl"),
		Timing:    filepath.Join(resultsDir, "timing", runID+".json"),
		Log:       filepath.Join(resultsDir, "logs", runID+".log"),
	}
}

func (p Paths) EnsureDirs() error {
	for _, dir := range []string{
		filepath.Join(p.BaseDir, "summary"),
		filepath.Join(p.BaseDir, "evolution"),
		filepath.Join(p.BaseDir, "timing"),
		filepath.Join(p.BaseDir, "logs"),
	} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (p Paths) RelativeTiming() string {
	return filepath.ToSlash(filepath.Join("timing", p.RunID+".json"))
}

func (p Paths) RelativeEvolution() string {
	return filepath.ToSlash(filepath.Join("evolution", p.RunID+".jsonl"))
}

func CheckExisting(paths Paths, policy string) (bool, error) {
	state, err := scanArtifactState(paths)
	if err != nil {
		return false, err
	}

	if !state.Any() {
		return false, nil
	}

	if !state.Complete() {
		if policy == IfExistsError {
			return false, fmt.Errorf(
				"run has incomplete artifacts for %q: missing %s",
				paths.RunID,
				strings.Join(state.Missing(), ", "),
			)
		}

		return false, nil
	}

	switch policy {
	case IfExistsSkip:
		return true, nil
	case IfExistsOverwrite:
		return false, nil
	case IfExistsError:
		return false, fmt.Errorf("run already exists: %s", paths.Summary)
	default:
		return false, fmt.Errorf("unknown if-exists policy %q", policy)
	}
}

func scanArtifactState(paths Paths) (artifactState, error) {
	summary, err := fileExists(paths.Summary)
	if err != nil {
		return artifactState{}, fmt.Errorf("stat summary artifact %q: %w", paths.Summary, err)
	}

	evolution, err := fileExists(paths.Evolution)
	if err != nil {
		return artifactState{}, fmt.Errorf("stat evolution artifact %q: %w", paths.Evolution, err)
	}

	timing, err := fileExists(paths.Timing)
	if err != nil {
		return artifactState{}, fmt.Errorf("stat timing artifact %q: %w", paths.Timing, err)
	}

	return artifactState{
		Summary:   summary,
		Evolution: evolution,
		Timing:    timing,
	}, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func WriteJSONAtomic(path string, data any) error {
	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	payload = append(payload, '\n')
	return writeFileAtomic(path, payload)
}

func WriteJSONLinesAtomic(path string, records []EvolutionRecord) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	for _, record := range records {
		if err := encoder.Encode(record); err != nil {
			return fmt.Errorf("encode evolution record: %w", err)
		}
	}

	return writeFileAtomic(path, buf.Bytes())
}

func BuildRunID(method, instance string, seed int64, params map[string]any) string {
	base := sanitizeID(strings.TrimSuffix(filepath.Base(instance), filepath.Ext(instance)))
	method = sanitizeID(method)
	canonical := canonicalParams(params)
	sum := sha1.Sum([]byte(fmt.Sprintf("schema=v1|method=%s|instance=%s|seed=%d|params=%s", method, instance, seed, canonical)))
	hash8 := hex.EncodeToString(sum[:])[:8]
	return fmt.Sprintf("%s__%s__s%d__h%s", base, method, seed, hash8)
}

func canonicalParams(params map[string]any) string {
	if len(params) == 0 {
		return ""
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	parts := make([]string, 0, len(params))
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%v", key, params[key]))
	}
	return strings.Join(parts, ",")
}

func sanitizeID(input string) string {
	if input == "" {
		return "run"
	}

	var b strings.Builder
	b.Grow(len(input))
	for _, r := range input {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '-' || r == '_':
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}

	output := b.String()
	if output == "" {
		return "run"
	}
	return output
}

func writeFileAtomic(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("create directory %q: %w", dir, err)
	}

	tmpFile, err := os.CreateTemp(dir, filepath.Base(path)+".tmp-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()

	defer func() {
		_ = os.Remove(tmpPath)
	}()

	if _, err := tmpFile.Write(data); err != nil {
		_ = tmpFile.Close()
		return fmt.Errorf("write temp file: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		_ = tmpFile.Close()
		return fmt.Errorf("sync temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close temp file: %w", err)
	}

	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("rename temp file: %w", err)
	}

	return nil
}

func NowTimestamp(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}
