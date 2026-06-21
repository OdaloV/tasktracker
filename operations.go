package main

import (
	"fmt"
	"os"
	"time"
)

func addTask(description string) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
		return
	}

	var newId int64 = 1
	if len(tasks) > 0 {
		maxID := tasks[0].ID
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		newId = maxID + 1
	}

	now := time.Now()
	task := Task{
		ID:          newId,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	err = saveTask("tasks.json", tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func updateTask(id int64, description string) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		os.Exit(1)
		return
	}

	err = saveTask("tasks.json", tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Task %d updated successfully\n", id)
}

func deleteTask(id int64) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
		return
	}

	found := false
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		os.Exit(1)
		return
	}

	err = saveTask("tasks.json", newTasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Task %d deleted successfully\n", id)
}

func markTask(id int64, status string) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		os.Exit(1)
		return
	}

	err = saveTask("tasks.json", tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Task %d marked as '%s'\n", id, status)
}
