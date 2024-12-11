/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"tcc/graph"
	"tcc/points"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// seed, err := cmd.Root().PersistentFlags().GetInt64("seed")
		// if err != nil {
		// 	cmd.PrintErrln("Error getting seed:", err)
		// 	return
		// }
		folder, err := cmd.Root().PersistentFlags().GetString("folder")
		if err != nil {
			cmd.PrintErrln("Error getting folder:", err)
			return
		}
		pts, err := points.Load(folder)
		if err != nil {
			cmd.PrintErrf("Error loading points from folder `%s`: %v", folder, err)
			return
		}
		grphs := graph.CreateAll(pts)

		err = graph.Save(grphs, folder)
		if err != nil {
			cmd.PrintErrf("Error saving graphs into folder `%s`: %v", folder, err)
			return
		}
		cmd.Printf("%d files saved at `%s`\n", len(pts), folder)
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
