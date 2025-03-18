/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Save the resource into files",
	Long:  `Save the points and the correspoding graph into disk`,
	Run: func(cmd *cobra.Command, args []string) {
		mapCmd.Run(cmd, args)
		graphCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String(
		"frequency",
		"3:5,4:5,5:5,6:5,7:5,8:5,9:5,10:5,20:5,30:5,40:5,50:5",
		"Mapa de frequência no formato 'chave:valor,chave2:valor2'",
	)
}
