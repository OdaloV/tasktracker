package main

import (
	"fmt"
	"os"
	"strconv"
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
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide ID and new description")
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid ID")
			return
		}
		description := strings.Join(os.Args[3:], " ")
		updateTask(id, description)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid ID")
			return
		}
		deleteTask(id)
	case "inprogress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid ID")
			return
		}
		markTask(id, "in-progress")
	case "done":

	default:

	}

}
