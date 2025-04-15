/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/graph"
	"tcc/optimization/ga"

	"github.com/spf13/cobra"
)

// gaCmd represents the ga command
var gaCmd = &cobra.Command{
	Use:   "ga",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		instance, err := cmd.Parent().PersistentFlags().GetString("instance")
		if err != nil {
			cmd.PrintErrln("Error getting instance file path:", err)
			return
		}
		g, err := graph.LoadFromFile(instance)
		if err != nil {
			cmd.PrintErrln("Error loading graph:", err)
			return
		}

		seed, err := cmd.Root().PersistentFlags().GetInt64("seed")
		if err != nil {
			cmd.PrintErrln("Error getting random seed:", err)
			return
		}
		rng := rand.New(rand.NewSource(seed))

		params, err := getParams(cmd)
		if err != nil {
			cmd.PrintErrln("Error getting parameters:", err)
			return
		}

		bestGen, bestFen := ga.Optimize(params, g, rng)
		fmt.Println("Best solution:", bestGen)
		fmt.Println("Best fitness:", bestFen)
	},
}

func init() {
	optimizeCmd.AddCommand(gaCmd)
	gaCmd.Flags().Int("elitism", 1, "Number of elite individuals to keep each generation")
	gaCmd.Flags().Float64("mutation-rate", 0.05, "Probability of mutation (swap two genes)")
	gaCmd.Flags().Int("tournament-size", 2.0, "Tournament size for parent selection")
}

func getParams(cmd *cobra.Command) (ga.Params, error) {
	var err error
	params := ga.Params{}

	iterations, err := cmd.Parent().PersistentFlags().GetInt("iterations")
	if err != nil {
		return params, err
	}
	population, err := cmd.Parent().PersistentFlags().GetInt("population")
	if err != nil {
		return params, err
	}
	params.HyperParams.Iterations = iterations
	params.HyperParams.PopulationSize = population

	params.Elitism, err = cmd.Flags().GetInt("elitism")
	if err != nil {
		return params, err
	}
	params.MutationRate, err = cmd.Flags().GetFloat64("mutation-rate")
	if err != nil {
		return params, err
	}
	params.TournamentSize, err = cmd.Flags().GetInt("tournament-size")
	if err != nil {
		return params, err
	}
	return params, nil
}
