// main.go
package main

import (
	"fmt"
	"os"
	functions "todo/pkg"
)

func main() {

	// If default file doesnt exist yet, create and exit script
	if functions.FilesExistCheck() {
		os.Exit(0)
	}

	functions.GetState()
	for {
		// true for add, false for see
		// list > used for see, where user is asked which list they want to see; overwritten if action = true
		action, list := functions.AskAction()

		if action {
			list, task, dueDate, prio := functions.UserInputAddToList()
			fmt.Println(list, task, dueDate, prio)
			functions.AddToList(list, task, dueDate, prio)
			break

		} else if !action {
			fmt.Printf("\nHere's the content of %s:\n\n", list)
			err := functions.ReadListFile(list)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}

}
