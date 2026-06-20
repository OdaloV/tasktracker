package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func loadTask(filename string) ([]Task, error) {
	//check if file exists
	_, err := os.Lstat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}
	//open
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//read file
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(fileData) == 0 {
		return []Task{}, nil
	}

	//parse json data to something go can read
	var tasklist []Task
	err = json.NewDecoder(file).Decode(&tasklist)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || err.Error() == "EOF" {
			return []Task{}, nil
		}
		return nil, err
	}

	return tasklist, nil

}

func saveTask(filename string, tasklist []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(tasklist); err != nil {
		return err
	}

	return nil
}
