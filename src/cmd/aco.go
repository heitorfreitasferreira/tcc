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
	acoCmd.Flags().Float64("rho", 0.5, "Rho") //BUG: ρ=0.5 é muito agressivo. Literatura usa 0.1-0.3. Com ρ=0.5, τ cai para 0.001 em 10 iterações sem reforço (1.0×0.5¹⁰), agravando diluição 3D.
	acoCmd.Flags().Float64("q", 100, "Q")
}

func getAcoParams(cmd *cobra.Command) (aco.Params, error) {
	iterations, err := getPositiveIntFlag(cmd, "iterations")
	if err != nil {
		return aco.Params{}, err
	}

	population, err := getPositiveIntFlag(cmd, "population")
	if err != nil {
		return aco.Params{}, err
	}

	alpha, err := getNonNegativeFloatFlag(cmd, "alpha")
	if err != nil {
		return aco.Params{}, err
	}

	beta, err := getNonNegativeFloatFlag(cmd, "beta")
	if err != nil {
		return aco.Params{}, err
	}

	gama, err := getNonNegativeFloatFlag(cmd, "gama")
	if err != nil {
		return aco.Params{}, err
	}

	rho, err := getProbabilityFlag(cmd, "rho")
	if err != nil {
		return aco.Params{}, err
	}

	q, err := getNonNegativeFloatFlag(cmd, "q")
	if err != nil {
		return aco.Params{}, err
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
