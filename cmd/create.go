package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create new list",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// Your root command logic here
		fmt.Printf("The new list \"%s\" has been created\n", name)
	},
}

func init() {

	rootCmd.AddCommand(createCmd)
}
