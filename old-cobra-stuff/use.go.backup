package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Change the list to be used",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// Your root command logic here
		fmt.Printf("You are now working on \"%s\"\n", name)
	},
}

func init() {

	rootCmd.AddCommand(useCmd)
}
