package cmd

import (
	"fmt"
	"os"
	functions "todo/internal"

	"encoding/json"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [name]",
	Short: "Shows content of current list",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if name == "content" {
			current, err := os.ReadFile(statePath)
			if err != nil {
				fmt.Println("error reading state file:", err)
				return
			}
			
			var todo functions.Todo

			err = json.Unmarshal(current, &todo)
			if err != nil {
				fmt.Println("error reading list file:", err)
				return
			}

			for i := 0; i < 3; i++ {
				fmt.Printf(`Task: %s
				Prio: %s
				DueDate: %s`,
			todo.Task, todo.Priority, todo.DueDate)
			}

			fmt.Printf(``)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}