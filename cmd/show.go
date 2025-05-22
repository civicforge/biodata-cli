package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	showCmd.PersistentFlags().StringP("pretty", "", "", "Format output nicely for terminal")
	showCmd.PersistentFlags().StringP("geometry", "", "", "Output metadata in JSON format")
	showCmd.PersistentFlags().StringP("crs", "", "", "Dump all fields unformatted")
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display full metadata for a dataset in the index",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: Implement Show")
	},
}
