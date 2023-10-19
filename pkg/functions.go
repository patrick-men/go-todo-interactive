package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"
)

var CurrentStateFile string
var CurrentStateName string

// Read state file, output current state filename with and without file extension
func readStateFile() error {

	content, err := os.ReadFile(statePath)
	if err != nil {
		return err
	}

	CurrentStateFile = filepath.Base(string(content))
	CurrentStateName = strings.TrimRight(CurrentStateFile, ".json\n")
	return nil
}

// Get current state + error handling
func GetState() {

	if err := readStateFile(); err != nil {
		fmt.Println("error reading state file: ", err)
		os.Exit(0)
	}

}

// Ask use whether he want to add to, or see a list. Handles wrong user input
func AskAction() (action bool) {

	message := "Do you want to (a)dd to, or (s)ee the contents of a list"
	input := strings.ToLower(readUserInput(message))

	if input == "a" || input == "ad" || input == "add" {
		action = true
		return
	} else if input == "s" || input == "se" || input == "see" {
		action = false
		return
	} else {
		fmt.Println("Error: Wrong input. Please enter a/add or s/see")
		os.Exit(0)
	}

	return
}

// Asks questions and get user input - only used if user wants to add to list
func UserInputAddToList() (list, task, dueDate, prio string) {

	message := fmt.Sprintf("Which list do you want to use [leave empty to use current: %s]", CurrentStateName)
	list = readUserInput(message)

	message = "What do you want to add to the list"
	task = readUserInput(message)

	message = "What's the due date"
	dueDate = readUserInput(message)

	message = "What's the priority"
	prio = readUserInput(message)


	return
}

func readUserInput(message string) (input string) {

	rl, err := readline.New(fmt.Sprintf("%s > ", message))
	if err != nil {
		fmt.Printf("Error reading user input: %t", err)
	}
	defer func() {
		_ = rl.Close()
	}()

	input, err = rl.Readline()
	if err != nil {
		fmt.Printf("Error reading user input: %t", err)
	}
	return

}
