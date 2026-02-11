/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/graph"
	"tcc/optimization/aco"
	"tcc/shared"
	"time"

	"github.com/spf13/cobra"
)

// acoCmd represents the aco command
var acoCmd = &cobra.Command{
	Use:   "aco",
	Short: "Optimize an instance using ant colony optimization",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		seed, err := cmd.Flags().GetInt64("seed")
		if err != nil {
			return fmt.Errorf("get --seed: %w", err)
		}

		instance, err := cmd.Flags().GetString("instance")
		if err != nil {
			return fmt.Errorf("get --instance: %w", err)
		}

		params, err := getAcoParams(cmd)
		if err != nil {
			return err
		}

		run, skip, err := prepareOptimizeRun(cmd, "aco", instance, seed, acoParamsMap(params))
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

		rng := rand.New(rand.NewSource(seed))
		optimizeStart := time.Now()
		result := aco.Optimize(params, g, rng, improvementLogger(cmd, run))
		optimizeDuration := time.Since(optimizeStart)
		optimizedAt := time.Now()

		return persistOptimizeRun(cmd, run, result, loadDuration, optimizeDuration, optimizedAt)
	},
}

func init() {
	optimizeCmd.AddCommand(acoCmd)

	acoCmd.Flags().Float64("alpha", 1.0, "Alpha")
	acoCmd.Flags().Float64("beta", 2.0, "Beta")
	acoCmd.Flags().Float64("gama", 0.1, "Gama")
	acoCmd.Flags().Float64("rho", 0.5, "Rho")
	acoCmd.Flags().Float64("q", 100, "Q")
}

func getAcoParams(cmd *cobra.Command) (aco.Params, error) {
	iterations, err := cmd.Flags().GetInt("iterations")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --iterations: %w", err)
	}

	population, err := cmd.Flags().GetInt("population")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --population: %w", err)
	}

	alpha, err := cmd.Flags().GetFloat64("alpha")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --alpha: %w", err)
	}

	beta, err := cmd.Flags().GetFloat64("beta")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --beta: %w", err)
	}

	gama, err := cmd.Flags().GetFloat64("gama")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --gama: %w", err)
	}

	rho, err := cmd.Flags().GetFloat64("rho")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --rho: %w", err)
	}

	q, err := cmd.Flags().GetFloat64("q")
	if err != nil {
		return aco.Params{}, fmt.Errorf("get --q: %w", err)
	}

	return aco.Params{
		HyperParams: shared.HyperParams{
			Iterations:     iterations,
			PopulationSize: population,
		},
		Alpha: alpha,
		Beta:  beta,
		Gama:  gama,
		Rho:   rho,
		Q:     q,
	}, nil
}

func acoParamsMap(params aco.Params) map[string]any {
	return map[string]any{
		"population": params.HyperParams.PopulationSize,
		"iterations": params.HyperParams.Iterations,
		"alpha":      params.Alpha,
		"beta":       params.Beta,
		"gama":       params.Gama,
		"rho":        params.Rho,
		"q":          params.Q,
	}
}
