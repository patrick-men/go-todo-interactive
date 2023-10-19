// main.go
package main

import (
	"os"
	"todo/cmd"
	"todo/internal"
)

func main() {

	// If default file is created, exit script
	if functions.FilesExistCheck() {
		os.Exit(0)
	}

	
	cmd.Execute()
}
