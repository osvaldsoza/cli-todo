package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/osvaldsoza/cli-todo/internal/task"
)

const filePath = "tasks.json"

func Load() ([]task.Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []task.Task{}, nil
		}
		return nil, fmt.Errorf("storage.Load %w	", err)
	}
	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("storage.Load JSON inválido: %w", err)
	}
	return tasks, nil
}

func Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("storage.Save: %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("storage.Save: %w", err)
	}
	return nil
}

func NextID(tasks []task.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
