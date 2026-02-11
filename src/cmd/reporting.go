package cmd

import (
	"fmt"
	"time"

	"tcc/shared"
	"tcc/shared/reporting"

	"github.com/spf13/cobra"
)

type optimizeOutputConfig struct {
	ResultsDir string
	RunID      string
	IfExists   string
	Progress   bool
}

type optimizeRunContext struct {
	method    string
	instance  string
	seed      int64
	params    map[string]any
	runID     string
	paths     reporting.Paths
	startedAt time.Time
	progress  bool
}

func getOptimizeOutputConfig(cmd *cobra.Command) (optimizeOutputConfig, error) {
	resultsDir, err := cmd.Flags().GetString("results-dir")
	if err != nil {
		return optimizeOutputConfig{}, fmt.Errorf("get --results-dir: %w", err)
	}

	runID, err := cmd.Flags().GetString("run-id")
	if err != nil {
		return optimizeOutputConfig{}, fmt.Errorf("get --run-id: %w", err)
	}

	ifExists, err := cmd.Flags().GetString("if-exists")
	if err != nil {
		return optimizeOutputConfig{}, fmt.Errorf("get --if-exists: %w", err)
	}

	ifExists, err = reporting.ParseIfExistsPolicy(ifExists)
	if err != nil {
		return optimizeOutputConfig{}, err
	}

	progress, err := cmd.Flags().GetBool("progress")
	if err != nil {
		return optimizeOutputConfig{}, fmt.Errorf("get --progress: %w", err)
	}

	return optimizeOutputConfig{
		ResultsDir: resultsDir,
		RunID:      runID,
		IfExists:   ifExists,
		Progress:   progress,
	}, nil
}

func prepareOptimizeRun(
	cmd *cobra.Command,
	method, instance string,
	seed int64,
	params map[string]any,
) (optimizeRunContext, bool, error) {
	config, err := getOptimizeOutputConfig(cmd)
	if err != nil {
		return optimizeRunContext{}, false, err
	}

	runID := config.RunID
	if runID == "" {
		runID = reporting.BuildRunID(method, instance, seed, params)
	} else if err := reporting.ValidateRunID(runID); err != nil {
		return optimizeRunContext{}, false, err
	}

	paths := reporting.PathsForRun(config.ResultsDir, runID)
	skip, err := reporting.CheckExisting(paths, config.IfExists)
	if err != nil {
		return optimizeRunContext{}, false, fmt.Errorf("check existing run artifacts: %w", err)
	}
	if skip {
		fmt.Fprintf(cmd.ErrOrStderr(), "Skipping run %s (%s): artifacts already exist\n", runID, method)
		return optimizeRunContext{}, true, nil
	}

	if err := paths.EnsureDirs(); err != nil {
		return optimizeRunContext{}, false, fmt.Errorf("create output directories: %w", err)
	}

	return optimizeRunContext{
		method:    method,
		instance:  instance,
		seed:      seed,
		params:    params,
		runID:     runID,
		paths:     paths,
		startedAt: time.Now(),
		progress:  config.Progress,
	}, false, nil
}

func improvementLogger(cmd *cobra.Command, run optimizeRunContext) func(shared.Improvement) {
	if !run.progress {
		return nil
	}

	return func(imp shared.Improvement) {
		fmt.Fprintf(
			cmd.ErrOrStderr(),
			"[%s][%s] improvement iter=%d eval=%d best=%.6f delta=%.6f\n",
			run.method,
			run.runID,
			imp.Iteration,
			imp.Evaluation,
			imp.BestMakespan,
			imp.Delta,
		)
	}
}

func persistOptimizeRun(
	cmd *cobra.Command,
	run optimizeRunContext,
	result shared.OptimizationResult,
	loadDuration time.Duration,
	optimizeDuration time.Duration,
	optimizedAt time.Time,
) error {
	serializeStart := time.Now()

	evolutionRecords := make([]reporting.EvolutionRecord, 0, len(result.Improvements))
	for _, improvement := range result.Improvements {
		record := reporting.EvolutionRecord{
			Schema:       reporting.EvolutionSchema,
			RunID:        run.runID,
			Method:       run.method,
			Iteration:    improvement.Iteration,
			EvalCount:    improvement.Evaluation,
			BestMakespan: improvement.BestMakespan,
			Delta:        improvement.Delta,
			BestSequence: append([]int(nil), improvement.BestSequence...),
		}
		evolutionRecords = append(evolutionRecords, record)
	}

	if err := reporting.WriteJSONLinesAtomic(run.paths.Evolution, evolutionRecords); err != nil {
		return fmt.Errorf("write evolution output: %w", err)
	}

	serializeDuration := time.Since(serializeStart)
	totalDuration := time.Since(run.startedAt)

	timing := reporting.RunTiming{
		Schema:   reporting.TimingSchema,
		RunID:    run.runID,
		Method:   run.method,
		Instance: run.instance,
		Durations: reporting.DurationMS{
			LoadInstance: loadDuration.Milliseconds(),
			Optimize:     optimizeDuration.Milliseconds(),
			Serialize:    serializeDuration.Milliseconds(),
			Total:        totalDuration.Milliseconds(),
		},
		Measured: reporting.TimeMeasured{
			StartedAt:  reporting.NowTimestamp(run.startedAt),
			FinishedAt: reporting.NowTimestamp(time.Now()),
		},
	}

	if err := reporting.WriteJSONAtomic(run.paths.Timing, timing); err != nil {
		return fmt.Errorf("write timing output: %w", err)
	}

	summary := reporting.RunSummary{
		Schema:   reporting.SummarySchema,
		RunID:    run.runID,
		Method:   run.method,
		Instance: run.instance,
		Seed:     run.seed,
		Params:   run.params,
		Status:   "ok",
		Result: reporting.SummaryResult{
			BestMakespan:        result.BestMakespan,
			BestSequence:        append([]int(nil), result.BestSequence...),
			IterationsCompleted: result.IterationsCompleted,
			Evaluations:         result.Evaluations,
		},
		TimingFile:    run.paths.RelativeTiming(),
		EvolutionFile: run.paths.RelativeEvolution(),
		StartedAt:     reporting.NowTimestamp(run.startedAt),
		FinishedAt:    reporting.NowTimestamp(optimizedAt),
	}

	if err := reporting.WriteJSONAtomic(run.paths.Summary, summary); err != nil {
		return fmt.Errorf("write summary output: %w", err)
	}

	cmd.Printf(
		"run_id=%s method=%s best_makespan=%.6f summary=%s\n",
		run.runID,
		run.method,
		result.BestMakespan,
		run.paths.Summary,
	)

	return nil
}
