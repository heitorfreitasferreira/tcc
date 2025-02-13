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
		// WIP: Sugestão: isso virar parametro num futuro distante (proavelmente nunca)
		frequency := map[int]int{
			3:  5,
			4:  5,
			5:  5,
			6:  5,
			7:  5,
			8:  5,
			9:  5,
			10: 5,
			20: 5,
			30: 5,
			40: 5,
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
