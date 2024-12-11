/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"tcc/points"

	"github.com/spf13/cobra"
)

// mapCmd represents the map command
var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Salva no disco instâncais de pontos gerados aleatóriamente para serem usando no TSP",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		frequency := map[int]int{
			// 3: 1,
			// 6: 1,
			// 7:  1,
			// 8:  1,
			// 9:  1,
			// 10: 1,
			// 15: 1,
			// 20: 1,
			// 25: 1,
			// 30: 1,
			// 35: 1,
			// 40: 1,
			// 45: 1,
			50: 5,
		}
		seed, err := cmd.Root().PersistentFlags().GetInt64("seed")
		if err != nil {
			cmd.PrintErrln("Error getting seed:", err)
			return
		}
		folder, err := cmd.Root().PersistentFlags().GetString("folder")
		if err != nil {
			cmd.PrintErrln("Error getting folder:", err)
			return
		}

		pts := points.CreateInstances(seed, frequency)
		err = points.Save(pts, folder)
		if err != nil {
			cmd.PrintErrln("Error saving instances into folder: ", err)
		}
		cmd.Printf("%d files saved at `%s`\n", len(pts), folder)
	},
}

func init() {
	createCmd.AddCommand(mapCmd)
}
