/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/graph"
	"tcc/pso"
	"tcc/shared"

	"github.com/spf13/cobra"
)

// psoCmd represents the pso command
var psoCmd = &cobra.Command{
	Use:   "pso",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		seed, err := cmd.Root().PersistentFlags().GetInt64("seed")
		if err != nil {
			cmd.PrintErrln("Error getting random seed parameter:", err)
			return
		}
		iterations, err := cmd.Parent().PersistentFlags().GetInt("iterations")
		if err != nil {
			cmd.PrintErrln("Error getting iterations parameter:", err)
			return
		}
		population, err := cmd.Parent().PersistentFlags().GetInt("population")
		if err != nil {
			cmd.PrintErrln("Error getting population parameter:", err)
			return
		}
		w, err := cmd.Flags().GetFloat64("w")
		if err != nil {
			cmd.PrintErrln("Error getting inertia parameter:", err)
			return
		}
		c1, err := cmd.Flags().GetFloat64("c1")
		if err != nil {
			cmd.PrintErrln("Error getting cognitive component (c1) parameter:", err)
			return
		}
		c2, err := cmd.Flags().GetFloat64("c2")
		if err != nil {
			cmd.PrintErrln("Error getting social component (c2) parameter:", err)
			return
		}
		instance, err := cmd.Parent().PersistentFlags().GetString("instance")
		if err != nil {
			cmd.PrintErrln("Error getting instance file path:", err)
			return
		}

		rng := rand.New(rand.NewSource(seed))

		params := pso.Params{
			HyperParams: shared.HyperParams{
				Iterations:     iterations,
				PopulationSize: population,
			},
			C1: c1,
			C2: c2,
			W:  w,
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
