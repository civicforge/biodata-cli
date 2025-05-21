package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "biodata",
	Short: "Biodata, a tool for indexing, searching, and managing large local spatial datasets",
	Long: `An easy to use CLI tool to help researchers, conservationists,
	and data professionals make sense of fragmented spatial files,
	without the need for complex GIS tools.

        Complete documentation is available at http://hugo.spf13.com`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
