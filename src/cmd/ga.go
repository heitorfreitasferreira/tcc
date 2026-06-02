/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/graph"
	"tcc/optimization/ga"
	"time"

	"github.com/spf13/cobra"
)

// gaCmd represents the ga command
var gaCmd = &cobra.Command{
	Use:   "ga",
	Short: "Optimize an instance using genetic algorithm",
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

		params, err := getParams(cmd)
		if err != nil {
			return err
		}

		run, skip, err := prepareOptimizeRun(cmd, "ga", instance, seed, gaParamsMap(params))
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
		result := ga.Optimize(params, g, rng, improvementLogger(cmd, run))
		optimizeDuration := time.Since(optimizeStart)
		optimizedAt := time.Now()

		return persistOptimizeRun(cmd, run, result, loadDuration, optimizeDuration, optimizedAt)
	},
}

func init() {
	optimizeCmd.AddCommand(gaCmd)
	gaCmd.Flags().Int("elitism", 1, "Number of elite individuals to keep each generation")
	gaCmd.Flags().Float64("mutation-rate", 0.05, "Probability of mutation (swap two genes)")
	gaCmd.Flags().Int("tournament-size", 2, "Tournament size for parent selection")
}

func getParams(cmd *cobra.Command) (ga.Params, error) {
	var err error
	params := ga.Params{}

	iterations, err := getPositiveIntFlag(cmd, "iterations")
	if err != nil {
		return params, err
	}
	population, err := getPositiveIntFlag(cmd, "population")
	if err != nil {
		return params, err
	}
	params.HyperParams.Iterations = iterations
	params.HyperParams.PopulationSize = population

	params.Elitism, err = getNonNegativeIntFlag(cmd, "elitism")
	if err != nil {
		return params, err
	}
	params.MutationRate, err = getProbabilityFlag(cmd, "mutation-rate")
	if err != nil {
		return params, err
	}
	params.TournamentSize, err = getPositiveIntFlag(cmd, "tournament-size")
	if err != nil {
		return params, err
	}
	return params, nil
}

func gaParamsMap(params ga.Params) map[string]any {
	return map[string]any{
		"population":      params.HyperParams.PopulationSize,
		"iterations":      params.HyperParams.Iterations,
		"elitism":         params.Elitism,
		"mutation_rate":   params.MutationRate,
		"tournament_size": params.TournamentSize,
	}
}
