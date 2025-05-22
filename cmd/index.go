package cmd

import (
	"github.com/civicforge/biodata-cli/internal/indexer"
	"github.com/spf13/cobra"
)

func init() {
	indexCmd.PersistentFlags().BoolP("verbose", "v", false, "Show all indexed files and errors")
	indexCmd.PersistentFlags().BoolP("recursive", "r", true, "Recurse into subdirectories")
	indexCmd.PersistentFlags().StringArrayP("format", "", []string{}, "Only includes files of thise format(e.g geojson, csv)")
	rootCmd.AddCommand(indexCmd)
}

var indexCmd = &cobra.Command{
	Use:   "index [DIRECTORY] [FLAGS]",
	Short: "Index all supported data files in a given directory, recursively",
	Args:  cobra.MinimumNArgs(1),
	Run:   indexer.Index,
}
