package cmd

import (
	"fmt"
	"tcc/graph"
	"tcc/optimization/lowerbound"
	"time"

	"github.com/spf13/cobra"
)

var lowerboundCmd = &cobra.Command{
	Use:   "lowerbound",
	Short: "Solve the Assignment Problem relaxation for a lower bound",
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

		run, skip, err := prepareOptimizeRun(cmd, "lowerbound", instance, seed, map[string]any{})
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
		result := lowerbound.Optimize(g, improvementLogger(cmd, run))
		optimizeDuration := time.Since(optimizeStart)
		optimizedAt := time.Now()

		return persistOptimizeRun(cmd, run, result, loadDuration, optimizeDuration, optimizedAt)
	},
}

func init() {
	optimizeCmd.AddCommand(lowerboundCmd)
}
