/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "tcc",
	Short:         "RTSP path finder using bio-inspired metaheuristics",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Int64P("seed", "s", 0, "-s 42")
	rootCmd.PersistentFlags().StringP("folder", "f", "./data", "Directory to save maps and graphs")
}
