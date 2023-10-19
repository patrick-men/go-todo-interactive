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
		action := functions.AskAction()

		if action {
			list, task, dueDate, prio := functions.UserInputAddToList()
			fmt.Println(list, task, dueDate, prio)
			break

		} else if !action {
			fmt.Println("filler")
			break
		} 
	}

}
