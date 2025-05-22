package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	exportCmd.PersistentFlags().StringP("format", "", "", "Output format: json, csv")
	exportCmd.PersistentFlags().StringP("output", "o", "", "Path to output file")
	exportCmd.PersistentFlags().StringP("filter", "", "", "Reuse a search expression (e.g., geojson only)")
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the entire index or a filtered subset to a structured file.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: Implement export")
	},
}
