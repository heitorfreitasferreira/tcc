package cmd

import (
	"fmt"
	"runtime"
	"tcc/graph"
	"tcc/optimization/brute"
	"tcc/shared"
	"time"

	"github.com/spf13/cobra"
)

var bruteforceCmd = &cobra.Command{
	Use:   "bruteforce",
	Short: "Solve the instance with exhaustive search",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		instance, err := cmd.Flags().GetString("instance")
		if err != nil {
			return fmt.Errorf("get --instance: %w", err)
		}

		seed, err := cmd.Flags().GetInt64("seed")
		if err != nil {
			return fmt.Errorf("get --seed: %w", err)
		}

		parallel, err := cmd.Flags().GetBool("parallel")
		if err != nil {
			return fmt.Errorf("get --parallel: %w", err)
		}

		workers := 0
		if parallel {
			workers, err = cmd.Flags().GetInt("workers")
			if err != nil {
				return fmt.Errorf("get --workers: %w", err)
			}
		}

		params := map[string]any{"parallel": parallel}
		if parallel {
			params["workers"] = workers
		}

		run, skip, err := prepareOptimizeRun(cmd, "bruteforce", instance, seed, params)
		if err != nil {
			return err
		}
		if skip {
			return nil
		}

		loadStart := time.Now()
		g, err := graph.LoadFromFile(instance)
		if err != nil {
			return fmt.Errorf("load graph instance %q: %w", instance, err)
		}
		loadDuration := time.Since(loadStart)

		optimizeStart := time.Now()
		var result shared.OptimizationResult
		if parallel {
			if workers <= 0 {
				workers = runtime.NumCPU()
			}
			if run.progress {
				fmt.Fprintf(cmd.ErrOrStderr(), "parallel bruteforce using %d workers\n", workers)
			}
			result = brute.OptimizeParallel(g, workers, nil)
		} else {
			result = brute.Optimize(g, improvementLogger(cmd, run))
		}
		optimizeDuration := time.Since(optimizeStart)
		optimizedAt := time.Now()

		return persistOptimizeRun(cmd, run, result, loadDuration, optimizeDuration, optimizedAt)
	},
}

func init() {
	optimizeCmd.AddCommand(bruteforceCmd)
	bruteforceCmd.Flags().Bool("parallel", false, "Use parallel exhaustive search")
	bruteforceCmd.Flags().Int("workers", 0, "Number of parallel workers (default: GOMAXPROCS)")
}
