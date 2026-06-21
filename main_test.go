package main

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Test task")

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Test task" {
		t.Errorf("Expected 'Test task', got '%s'", tasks[0].Description)
	}
	if tasks[0].Status != "todo" {
		t.Errorf("Expected 'todo', got '%s'", tasks[0].Status)
	}
}

func TestAddMultipleTasks(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Task 1")
	addTask("Task 2")
	addTask("Task 3")

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}
}

func TestUpdateTask(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Original")
	updateTask(1, "Updated")

	tasks, _ := loadTask("tasks.json")
	if tasks[0].Description != "Updated" {
		t.Errorf("Expected 'Updated', got '%s'", tasks[0].Description)
	}
}

func TestUpdateTaskNotFound(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Task 1")
	updateTask(99, "New")

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Task 1")
	addTask("Task 2")
	addTask("Task 3")

	deleteTask(2)

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Task 1")
	deleteTask(99)

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestMarkTask(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Test task")

	markTask(1, "in-progress")
	tasks, _ := loadTask("tasks.json")
	if tasks[0].Status != "in-progress" {
		t.Errorf("Expected 'in-progress', got '%s'", tasks[0].Status)
	}

	markTask(1, "done")
	tasks, _ = loadTask("tasks.json")
	if tasks[0].Status != "done" {
		t.Errorf("Expected 'done', got '%s'", tasks[0].Status)
	}
}

func TestMarkTaskNotFound(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("Test task")
	markTask(99, "done")

	tasks, _ := loadTask("tasks.json")
	if tasks[0].Status != "todo" {
		t.Errorf("Expected 'todo', got '%s'", tasks[0].Status)
	}
}

func TestIDGeneration(t *testing.T) {
	cleanup()
	defer cleanup()

	addTask("First")
	addTask("Second")
	deleteTask(1)
	addTask("Third")

	tasks, _ := loadTask("tasks.json")
	if tasks[0].ID != 2 {
		t.Errorf("Expected ID 2, got %d", tasks[0].ID)
	}
	if tasks[1].ID != 3 {
		t.Errorf("Expected ID 3, got %d", tasks[1].ID)
	}
}

func TestEmptyList(t *testing.T) {
	cleanup()
	defer cleanup()

	tasks, _ := loadTask("tasks.json")
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestCorruptedJSON(t *testing.T) {
	cleanup()
	defer cleanup()

	os.WriteFile("tasks.json", []byte("{invalid"), 0644)

	_, err := loadTask("tasks.json")
	if err == nil {
		t.Errorf("Expected error for corrupted JSON")
	}
}

func cleanup() {
	os.Remove("tasks.json")
}
