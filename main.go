package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a command")
		printUsage()
		os.Exit(1)
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a description")
			os.Exit(1)
			return
		}
		description := strings.Join(os.Args[2:], " ")
		addTask(description)

	case "list":
		status := ""
		if len(os.Args) >= 3 {
			status = os.Args[2]
			if status != "todo" && status != "in-progress" && status != "done" {
				fmt.Printf("Error: Invalid status '%s'. Use: todo, in-progress, or done\n", status)
				os.Exit(1)
				return
			}
		}
		listTasks(status)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide ID and new description")
			os.Exit(1)
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("Error: Invalid ID '%s'. Please provide a number\n", os.Args[2])
			os.Exit(1)
			return
		}
		description := strings.Join(os.Args[3:], " ")
		updateTask(id, description)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			os.Exit(1)
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("Error: Invalid ID '%s'. Please provide a number\n", os.Args[2])
			os.Exit(1)
			return
		}
		deleteTask(id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			os.Exit(1)
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("Error: Invalid ID '%s'. Please provide a number\n", os.Args[2])
			os.Exit(1)
			return
		}
		markTask(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			os.Exit(1)
			return
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("Error: Invalid ID '%s'. Please provide a number\n", os.Args[2])
			os.Exit(1)
			return
		}
		markTask(id, "done")

	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Task Tracker CLI - Usage:
  add <description>              - Add a new task
  list                           - List all tasks
  list todo                      - List tasks with status 'todo'
  list in-progress               - List tasks with status 'in-progress'
  list done                      - List tasks with status 'done'
  update <id> <description>      - Update a task
  delete <id>                    - Delete a task
  mark-in-progress <id>          - Mark task as in-progress
  mark-done <id>                 - Mark task as done`)
}
