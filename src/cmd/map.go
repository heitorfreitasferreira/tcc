/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"strings"
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
		frequency := parseFrequency(cmd)
		if frequency == nil {
			return
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

func parseFrequency(cmd *cobra.Command) map[int]int {
	frequencyStr, err := cmd.Flags().GetString("frequency") // ← Corrigido!
	if err != nil {
		cmd.PrintErrln("Erro ao obter frequency:", err)
		return nil
	}

	frequency := make(map[int]int)

	for i, pair := range strings.Split(frequencyStr, ",") {
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			cmd.PrintErrf("Entrada inválida no par #%d: '%s'. Formato esperado: 'chave:valor'\n", i+1, pair)
			return nil
		}

		key, err := strconv.Atoi(parts[0])
		if err != nil {
			cmd.PrintErrf("Chave inválida no par #%d ('%s'): %v\n", i+1, parts[0], err)
			return nil
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			cmd.PrintErrf("Valor inválido no par #%d ('%s'): %v\n", i+1, parts[1], err)
			return nil
		}

		frequency[key] = value
	}
	return frequency
}
