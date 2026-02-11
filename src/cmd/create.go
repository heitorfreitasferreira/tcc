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
	Short: "Generate points and graphs",
	Long:  `Generate random point instances and their corresponding graph files.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runMap(cmd); err != nil {
			return err
		}

		return runGraph(cmd)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String(
		"frequency",
		"3:5,4:5,5:5,6:5,7:5,8:5,9:5,10:5,20:5,30:5,40:5,50:5",
		"Frequency map in the format 'key:value,key2:value2'",
	)
}
