package main

import (
	"fmt"
	"time"
)

func addTask(description string) {
	//load task
	tasks, err := loadTask(`task.json`)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}
	//while crating new task ,assign new id
	var newId int64 = 1
	maxID := tasks[0].ID
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newId = maxID + 1
	//create task
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
		fmt.Println("Error saving file:", err)
		return
	}
	fmt.Println("Task added successfully!")

}

func updateTask(id int64, description string) {
	tasks, err := loadTask("tasks.json")
	if err != nil {
		fmt.Printf("Error loading tasks", err)
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
		return
	}

	// Save changes to file
	err = saveTask("tasks.json", tasks)
	if err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		return
	}

	fmt.Printf("Task %d updated successfully\n", id)
}
