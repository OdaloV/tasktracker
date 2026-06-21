package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func loadTask(filename string) ([]Task, error) {
	_, err := os.Lstat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("failed to access file: %w", err)
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(fileData) == 0 {
		return []Task{}, nil
	}

	var tasklist []Task
	err = json.Unmarshal(fileData, &tasklist)
	if err != nil {
		return nil, fmt.Errorf("corrupted JSON file: %w", err)
	}

	return tasklist, nil
}

func saveTask(filename string, tasklist []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(tasklist); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
