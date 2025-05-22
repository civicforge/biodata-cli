package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	tagCmd.PersistentFlags().StringP("add", "", "", "One or more tags to add")
	tagCmd.PersistentFlags().StringP("remove", "", "", "One or more tags to remove")
	tagCmd.PersistentFlags().StringP("replace", "", "", "Replace ALL tags with provided")
	rootCmd.AddCommand(tagCmd)
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add or remove tags to a dataset in the index",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: Implement Tag")
	},
}
