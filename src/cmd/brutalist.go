/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tcc/brute"
	"tcc/graph"

	"github.com/spf13/cobra"
)

// brutalistCmd represents the brutalist command
var brutalistCmd = &cobra.Command{
	Use:   "brutalist",
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
		graph, err := graph.LoadFromFile(instance)
		if err != nil {
			cmd.PrintErrln("Error getting folder:", err)
			return
		}
		ind, mksp := brute.Optimize(graph)
		fmt.Println(ind, mksp)
	},
}

func init() {
	optimizeCmd.AddCommand(brutalistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// brutalistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// brutalistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
