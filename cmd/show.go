package cmd

import (
	"github.com/civicforge/biodata-cli/internal/show"
	"github.com/spf13/cobra"
)

func init() {
	showCmd.PersistentFlags().BoolP("pretty", "", false, "Format output nicely for terminal")
	showCmd.PersistentFlags().BoolP("json", "", false, "Output metadata in JSON format")
	showCmd.PersistentFlags().BoolP("raw", "", false, "Dump all fields unformatted")
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display full metadata for a dataset in the index",
	Args:  cobra.MinimumNArgs(1),
	Run:   show.Show,
}
