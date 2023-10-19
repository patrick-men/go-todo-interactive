package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type Todo struct {
	Task     string `json:"task"`
	Priority string `json:"priority"`
	DueDate	 string	`json:"dueDate"`
}


var dirPath = GetPath("home")
var filePath = GetPath("file")
var statePath = GetPath("state")


func GetPath(s string) (path string) {

	// Get user 
	usr, err := user.Current()
	if err != nil {
		fmt.Println("error getting user's home directory:", err)
		return
	}

	dirPath := filepath.Join(usr.HomeDir, ".config", "todo")

	switch s {
	case "home":
		path = dirPath
	case "file":
		path = dirPath + "/default.json"
	case "state":
		path = dirPath + "/.state.json"
	}
	return path
}

// If directory exists: true
// If directory doesn't exist: false
func directoryCheck() (b bool) {
	b = true

	_, err := os.Stat(dirPath)
	if err != nil {
		b = false
	}

	return b
}


// If state file exists: true
// If state file doesn't exist: false
func stateCheck() (b bool) {
	b = true

	_, err := os.Stat(statePath)
	if err != nil {
		b = false
	}

	return b
}

// If list file exists: true
// If list file doesn't exist: false
func listCheck() (b bool) {
	b = true

	_, err := os.Stat(filePath)
	if err != nil {
		b = false
	}

	return b
}

func createDefaultPath() {
	
	_, err := os.Stat(dirPath)

	if err != nil{
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("error creating project directories:", err)
			return
		}
	}

}

func createDefaultList() {

	defaultTodoListJson := []Todo{
		{"Fix something", "M", ""},
		{"Change this to that", "H", "23.10.2023"},
		{"Answer mail", "L", "Today"},
	}

	// create default list file
	stateFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("error creating list file:", err)
		return
	}
	

	// encode default values
	jsonData, err := json.Marshal(defaultTodoListJson)
	if err != nil {
		fmt.Println("error encoding JSON:", err)
		return
	}

	// encoder := json.NewEncoder(file)
	// encoder.Encode(jsonData)

	// write to file
	err = os.WriteFile(filePath, jsonData, 0755)
	if err != nil {
		fmt.Println("error writing JSON to file:", err)
		return
	}

	fmt.Println("Since this is your first time using ToDo, we have created a default list.")

	stateFile.Close()

}

func createDefaultState() {

	// create default state file
	listFile, err := os.Create(statePath)
	if err != nil {
		fmt.Println("error creating state file:", err)
		return
	}

	// write to file
	data := []byte("default.json\n")
	err = os.WriteFile(statePath, data, 0755)
	if err != nil {
		fmt.Println("error writing to state file:", err)
		return
	}

	listFile.Close()

	
}

func FilesExistCheck() (b bool){

	b = false

	// Check if the directory exists
	if !directoryCheck() {
		createDefaultPath()
		b = true
	}

	// Check if a todo list file exists
	if !listCheck() {
		createDefaultList()
		b = true
	}

	// Check if a state file exists
	if !stateCheck() {
		createDefaultState()
		b = true
	}

	return b
}