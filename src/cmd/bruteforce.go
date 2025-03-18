/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tcc/graph"
	"tcc/optimization/brute"

	"github.com/spf13/cobra"
)

var bruteforceCmd = &cobra.Command{
	Use: "bruteforce",
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
	optimizeCmd.AddCommand(bruteforceCmd)
}
