/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"tcc/aco"
	"tcc/graph"
	"tcc/shared"

	"github.com/spf13/cobra"
)

// acoCmd represents the aco command
var acoCmd = &cobra.Command{
	Use:   "aco",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empoalphaers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		graph, err := graph.LoadFromFile(instance)
		if err != nil {
			cmd.PrintErrln("Error getting folder:", err)
			return
		}
		params, err := getAcoParams(cmd)
		if err != nil {
			cmd.PrintErrln("Error getting parameters: ", err)
			return
		}
		fmt.Println(aco.Optimize(params, graph, rng))
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
	iterations, err := cmd.Parent().PersistentFlags().GetInt("iterations")
	if err != nil {
		cmd.PrintErrln("Error getting iterations parameter:", err)
		return aco.Params{}, err
	}
	population, err := cmd.Parent().PersistentFlags().GetInt("population")
	if err != nil {
		cmd.PrintErrln("Error getting population parameter:", err)
		return aco.Params{}, err
	}

	alpha, err := cmd.Flags().GetFloat64("alpha")
	if err != nil {
		cmd.PrintErrln("Error getting alpha", err)
		return aco.Params{}, err
	}
	beta, err := cmd.Flags().GetFloat64("beta")
	if err != nil {
		cmd.PrintErrln("Error getting beta", err)
		return aco.Params{}, err
	}
	gama, err := cmd.Flags().GetFloat64("gama")
	if err != nil {
		cmd.PrintErrln("Error getting gama", err)
		return aco.Params{}, err
	}
	rho, err := cmd.Flags().GetFloat64("rho")
	if err != nil {
		cmd.PrintErrln("Error getting Rho", err)
		return aco.Params{}, err
	}

	q, err := cmd.Flags().GetFloat64("q")
	if err != nil {
		cmd.PrintErrln("Error getting Q", err)
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
