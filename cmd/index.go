package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Resurse bool
var Verbose bool
var Format string

func init() {
	indexCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Show all indexed files and errors")
	indexCmd.PersistentFlags().BoolVarP(&Resurse, "recursive", "r", true, "Recurse into subdirectories")
	indexCmd.PersistentFlags().StringVarP(&Format, "format", "", "", "Only includes files of thise format(e.g geojson, csv)")
	rootCmd.AddCommand(indexCmd)
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index all supported data files in a given directory, recursively",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BioData-CLI v0.0.1 -- HEAD")
	},
}
