// cmd/greet.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task] [list]",
	Short: "Add task to list",
	Args:  cobra.ExactArgs(2), // Require exactly one argument (the name)
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		list := currentState

		if args[1] != "" {
			list = args[1]
		}

		fmt.Printf("Task %s has been added to %s!\n", task, list)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
