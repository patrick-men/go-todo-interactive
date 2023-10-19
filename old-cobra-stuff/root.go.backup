package cmd

import (
	"fmt"
	"os"
	functions "todo/internal"

	"github.com/spf13/cobra"
)

var dirPath = functions.GetPath("home")
var filePath = functions.GetPath("file")
var statePath = functions.GetPath("state")

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "ToDo CLI tool written in Go",
	Run: func(cmd *cobra.Command, args []string) {
		// Your root command logic here
		fmt.Println(`This is a ToDo Tool`)
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
