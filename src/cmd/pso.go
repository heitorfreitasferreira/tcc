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

	"github.com/spf13/cobra"
)

var psoCmd = &cobra.Command{
	Use: "pso",
	Run: func(cmd *cobra.Command, args []string) {
		seed, err := cmd.Root().PersistentFlags().GetInt64("seed")
		if err != nil {
			cmd.PrintErrln("Error getting random seed parameter:", err)
			return
		}
		instance, err := cmd.Parent().PersistentFlags().GetString("instance")
		if err != nil {
			cmd.PrintErrln("Error getting instance file path:", err)
			return
		}

		rng := rand.New(rand.NewSource(seed))

		params, err := getPsoParams(cmd)
		if err != nil {
			cmd.PrintErrln("Error getting parameters: ", err)
			return
		}
		graph, err := graph.LoadFromFile(instance)
		if err != nil {
			cmd.PrintErrln("Error getting folder:", err)
			return
		}
		sw := pso.Swarm{}
		sw.Init(&graph, params, rng)
		stats := sw.Optimize()
		fmt.Print(stats.ToCsv())
	},
}

func init() {
	optimizeCmd.AddCommand(psoCmd)
	psoCmd.Flags().Float64("c1", 2.0, "Cognitive component")
	psoCmd.Flags().Float64("c2", 2.0, "Social component")
	psoCmd.Flags().Float64("w", 0.7, "Inertia weight")
}

func getPsoParams(cmd *cobra.Command) (pso.Params, error) {
	iterations, err := cmd.Parent().PersistentFlags().GetInt("iterations")
	if err != nil {
		cmd.PrintErrln("Error getting iterations parameter:", err)
		return pso.Params{}, err
	}
	population, err := cmd.Parent().PersistentFlags().GetInt("population")
	if err != nil {
		cmd.PrintErrln("Error getting population parameter:", err)
		return pso.Params{}, err
	}
	w, err := cmd.Flags().GetFloat64("w")
	if err != nil {
		cmd.PrintErrln("Error getting inertia parameter:", err)
		return pso.Params{}, err
	}
	c1, err := cmd.Flags().GetFloat64("c1")
	if err != nil {
		cmd.PrintErrln("Error getting cognitive component (c1) parameter:", err)
		return pso.Params{}, err
	}
	c2, err := cmd.Flags().GetFloat64("c2")
	if err != nil {
		cmd.PrintErrln("Error getting social component (c2) parameter:", err)
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
