/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

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

func getPositiveIntFlag(cmd *cobra.Command, name string) (int, error) {
	value, err := cmd.Flags().GetInt(name)
	if err != nil {
		return 0, fmt.Errorf("get --%s: %w", name, err)
	}
	if value <= 0 {
		return 0, fmt.Errorf("invalid --%s %d: expected a positive integer", name, value)
	}
	return value, nil
}

func getNonNegativeIntFlag(cmd *cobra.Command, name string) (int, error) {
	value, err := cmd.Flags().GetInt(name)
	if err != nil {
		return 0, fmt.Errorf("get --%s: %w", name, err)
	}
	if value < 0 {
		return 0, fmt.Errorf("invalid --%s %d: expected a non-negative integer", name, value)
	}
	return value, nil
}

func getNonNegativeFloatFlag(cmd *cobra.Command, name string) (float64, error) {
	value, err := cmd.Flags().GetFloat64(name)
	if err != nil {
		return 0, fmt.Errorf("get --%s: %w", name, err)
	}
	if value < 0 {
		return 0, fmt.Errorf("invalid --%s %g: expected a non-negative number", name, value)
	}
	return value, nil
}

func getProbabilityFlag(cmd *cobra.Command, name string) (float64, error) {
	value, err := getNonNegativeFloatFlag(cmd, name)
	if err != nil {
		return 0, err
	}
	if value > 1 {
		return 0, fmt.Errorf("invalid --%s %g: expected a value between 0 and 1", name, value)
	}
	return value, nil
}
