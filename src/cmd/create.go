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

}
