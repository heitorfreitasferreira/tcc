/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize an instance with a chosen algorithm",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return errors.New("missing optimization method")
	},
}

func init() {
	rootCmd.AddCommand(optimizeCmd)

	optimizeCmd.PersistentFlags().String("instance", "./data/10a.graph", "Path to instance that will be optimized")
	optimizeCmd.PersistentFlags().String("results-dir", "./data/results", "Directory to save structured optimization outputs")
	optimizeCmd.PersistentFlags().String("run-id", "", "Optional explicit run identifier")
	optimizeCmd.PersistentFlags().String("if-exists", "skip", "How to handle existing outputs: skip|overwrite|error")
	optimizeCmd.PersistentFlags().Bool("progress", true, "Print progress to stderr while running")

	// shared.HyperParams
	optimizeCmd.PersistentFlags().IntP("population", "p", 100, "Size of the population")
	optimizeCmd.PersistentFlags().IntP("iterations", "i", 100, "Number of optimization iterations")
}
