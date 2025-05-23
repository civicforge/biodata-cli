package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of BioData-CLI",
	Long:  `All software has versions. This is BioData's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BioData-CLI v0.0.1 -- HEAD")
	},
}
