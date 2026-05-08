package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/osvaldsoza/cli-todo/internal/task"
)

const filePath = "taks.json"

func Losd() ([]task.Task, error) {
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
