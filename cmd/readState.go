package cmd

import (
    "fmt"
    "os"
    "path/filepath"
)

var currentState string

func readStateFile() error {

	content, err := os.ReadFile(statePath)
    if err != nil {
        return err
    }

	currentState = filepath.Base(string(content))
	return nil
}

func GetState() {
    // Initialize the state when Cobra commands are executed
    if err := readStateFile(); err != nil {
        fmt.Println("error reading state file: owo", err)
        os.Exit(0)
    }

    
}