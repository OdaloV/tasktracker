package main

import (
	"fmt"
	"strings"
)

func listTasks(status string) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	//filter tasks by status
	var filteredTasks []Task
	if status == "" {
		//show all tasks
		filteredTasks = tasks
	} else {
		for _, task := range tasks {
			if strings.EqualFold(task.Status, status) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	//if filtered results are empty
	if len(filteredTasks) == 0 {
		if status == "" {
			fmt.Println("No tasks found.")
		} else {
			fmt.Printf("No tasks with status '%s' found.\n", status)
		}
		return
	}

	//display each task
	for _, task := range filteredTasks {
		printTask(task)
	}
}

func printTask(task Task) {
	updatedStr := task.UpdatedAt.Format("2006-01-02 15:04:05")
	fmt.Printf("[%d] %s - %s (updated: %s)\n",
		task.ID,
		task.Description,
		task.Status,
		updatedStr)
}
