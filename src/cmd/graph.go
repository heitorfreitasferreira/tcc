/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tcc/graph"
	"tcc/points"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate graph files from points",
	Long:  `Read point files from a folder and generate graph files in the same folder.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGraph(cmd)
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}

func runGraph(cmd *cobra.Command) error {
	folder, err := cmd.Flags().GetString("folder")
	if err != nil {
		return fmt.Errorf("get --folder: %w", err)
	}

	pts, err := points.Load(folder)
	if err != nil {
		return fmt.Errorf("load points from %q: %w", folder, err)
	}

	grphs := graph.CreateAll(pts)
	report, err := graph.Save(grphs, folder)
	if err != nil {
		return fmt.Errorf("save graphs in %q: %w", folder, err)
	}

	cmd.Printf(
		"graph files: created=%d skipped=%d total=%d at `%s`\n",
		report.Created,
		report.Skipped,
		len(grphs),
		folder,
	)
	return nil
}
