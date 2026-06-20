package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please input an arguement")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		//check if description is provided
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide  description")
			return
		}
		// Join all arguments
		description := strings.Join(os.Args[2:], " ")
		addTask(description)

	case "list":
		status := ""
		if len(os.Args) >= 3 {
			status = os.Args[2]
			// Validate status
			if status != "todo" && status != "in-progress" && status != "done" {
				fmt.Printf("Error: Invalid status '%s'. Use: todo, in-progress, or done\n", status)
				return
			}
		}
		listTasks(status)

	case "update":
	case "delete":
	case "in progress":
	case "done":
	default:

	}

}
