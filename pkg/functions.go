package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"
)

var CurrentStateFile string
var CurrentStateName string
var UserHome = os.Getenv("HOME")

func beautifyJSON(inputJSON []byte) ([]byte, error) {
	var tasks []map[string]string
	err := json.Unmarshal(inputJSON, &tasks)
	if err != nil {
		return nil, err
	}

	var formattedTasks []string
	for i, task := range tasks {
		index := i + 1
		formattedTask := fmt.Sprintf("%d Task: %s", index, task["task"])
		if dueDate := task["dueDate"]; dueDate != "" {
			formattedTask += "\nDue Date: " + dueDate
		}
		if prio := task["priority"]; prio != "" {
			formattedTask += "\nPriority: " + prio + "\n"
		}
		formattedTasks = append(formattedTasks, formattedTask)
	}

	formattedJSON := strings.Join(formattedTasks, "\n")

	return []byte(formattedJSON), nil
}

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

func ReadListFile(list string) error {
	content, err := os.ReadFile(fmt.Sprintf("%s/.config/todo/%s.json", UserHome, list))
	if err != nil {
		fmt.Println(err)
		return err
	}

	var payload []interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	beatufiedJson, _ := beautifyJSON(content)

	fmt.Println(string(beatufiedJson))
	return err
}

// Get all files that contain todo-lists
func GetAllListFiles() (listfiles []string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var listFiles []string
	for _, file := range files {
		if file.Name() != ".state.json" {
			listFiles = append(listFiles, strings.TrimRight(file.Name(), ".json"))
			return listFiles
		}
	}
	return
}

// Get current state + error handling
func GetState() {

	if err := readStateFile(); err != nil {
		fmt.Println("error reading state file: ", err)
		os.Exit(0)
	}

}

// Ask use whether he want to add to, or see a list. Handles wrong user input
func AskAction() (action bool, list string) {

	message := "Do you want to (a)dd to, or (s)ee the contents of a list"
	input := strings.ToLower(readUserInput(message))

	if input == "a" || input == "ad" || input == "add" {
		action = true
		return
	} else if input == "s" || input == "se" || input == "see" {
		list = readUserInput(fmt.Sprintf("Which list do you want to use [leave empty to use current: %s]", CurrentStateName))
		if list == "" {
			list = CurrentStateName
		}
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

func AddToList(list, task, dueDate, prio string) {

}

// reads user input
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
