/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/graph"
	"tcc/optimization/pso"
	"tcc/shared"
	"time"

	"github.com/spf13/cobra"
)

var psoCmd = &cobra.Command{
	Use:   "pso",
	Short: "Optimize an instance using particle swarm optimization",
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

		rng := rand.New(rand.NewSource(seed))

		params, err := getPsoParams(cmd)
		if err != nil {
			return err
		}

		run, skip, err := prepareOptimizeRun(cmd, "pso", instance, seed, psoParamsMap(params))
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

		sw := pso.Swarm{}
		sw.Init(g, params, rng)

		optimizeStart := time.Now()
		result := sw.Optimize(improvementLogger(cmd, run))
		optimizeDuration := time.Since(optimizeStart)
		optimizedAt := time.Now()

		return persistOptimizeRun(cmd, run, result, loadDuration, optimizeDuration, optimizedAt)
	},
}

func init() {
	optimizeCmd.AddCommand(psoCmd)
	psoCmd.Flags().Float64("c1", 2.0, "Cognitive component")
	psoCmd.Flags().Float64("c2", 2.0, "Social component")
	psoCmd.Flags().Float64("w", 0.7, "Inertia weight")
}

func getPsoParams(cmd *cobra.Command) (pso.Params, error) {
	iterations, err := getPositiveIntFlag(cmd, "iterations")
	if err != nil {
		return pso.Params{}, err
	}

	population, err := getPositiveIntFlag(cmd, "population")
	if err != nil {
		return pso.Params{}, err
	}

	w, err := getNonNegativeFloatFlag(cmd, "w")
	if err != nil {
		return pso.Params{}, err
	}

	c1, err := getNonNegativeFloatFlag(cmd, "c1")
	if err != nil {
		return pso.Params{}, err
	}

	c2, err := getNonNegativeFloatFlag(cmd, "c2")
	if err != nil {
		return pso.Params{}, err
	}

	return pso.Params{
		HyperParams: shared.HyperParams{
			Iterations:     iterations,
			PopulationSize: population,
		},
		C1: c1,
		C2: c2,
		W:  w,
	}, nil
}

func psoParamsMap(params pso.Params) map[string]any {
	return map[string]any{
		"population": params.HyperParams.PopulationSize,
		"iterations": params.HyperParams.Iterations,
		"w":          params.W,
		"c1":         params.C1,
		"c2":         params.C2,
	}
}
