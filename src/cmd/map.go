/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"tcc/points"

	"github.com/spf13/cobra"
)

// mapCmd represents the map command
var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Generate random point instances",
	Long:  `Generate random 2D point instances and save them to disk.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMap(cmd)
	},
}

func init() {
	createCmd.AddCommand(mapCmd)
}

func runMap(cmd *cobra.Command) error {
	frequency, err := parseFrequency(cmd)
	if err != nil {
		return err
	}

	seed, err := cmd.Flags().GetInt64("seed")
	if err != nil {
		return fmt.Errorf("get --seed: %w", err)
	}

	folder, err := cmd.Flags().GetString("folder")
	if err != nil {
		return fmt.Errorf("get --folder: %w", err)
	}

	pts := points.CreateInstances(seed, frequency)
	if err := points.Save(pts, folder); err != nil {
		return fmt.Errorf("save points in %q: %w", folder, err)
	}

	cmd.Printf("%d point files saved at `%s`\n", len(pts), folder)
	return nil
}

func parseFrequency(cmd *cobra.Command) (map[int]int, error) {
	frequencyStr, err := cmd.Flags().GetString("frequency")
	if err != nil {
		return nil, fmt.Errorf("get --frequency: %w", err)
	}

	frequency := make(map[int]int)

	for i, pair := range strings.Split(frequencyStr, ",") {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid pair #%d %q, expected format 'key:value'", i+1, pair)
		}

		keyText := strings.TrimSpace(parts[0])
		valueText := strings.TrimSpace(parts[1])

		key, err := strconv.Atoi(keyText)
		if err != nil {
			return nil, fmt.Errorf("invalid key in pair #%d %q: %w", i+1, keyText, err)
		}

		value, err := strconv.Atoi(valueText)
		if err != nil {
			return nil, fmt.Errorf("invalid value in pair #%d %q: %w", i+1, valueText, err)
		}

		frequency[key] = value
	}

	return frequency, nil
}
