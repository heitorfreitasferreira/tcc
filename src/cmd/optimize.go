/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize an instance based on an algorithm",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(optimizeCmd)

	optimizeCmd.PersistentFlags().String("instance", "./data/10a.graph", "Path to instance that will be optimized")

	// shared.HyperParams
	optimizeCmd.PersistentFlags().IntP("population", "p", 100, "Size of the population")
	optimizeCmd.PersistentFlags().IntP("iterations", "i", 100, "Size of the population")
}
