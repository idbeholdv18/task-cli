package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-cli/src/task"
)

func ReadFromJson(path string) (task.Tasks, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return task.Tasks{}, nil
		}
		return nil, err
	}

	if len(content) == 0 {
		return task.Tasks{}, nil
	}

	var dtos []task.TaskDto
	if err := json.Unmarshal(content, &dtos); err != nil {
		return nil, fmt.Errorf("Error during unmarshaling process: %w", err)
	}

	tasks := make(task.Tasks, 0, len(dtos))

	for _, dto := range dtos {
		taskFromDto := task.TaskFromDto(dto)
		tasks = append(tasks, &taskFromDto)
	}

	return tasks, nil
}

func WriteToJson(path string, tasks task.Tasks) error {
	dtos := make([]task.TaskDto, 0, len(tasks))

	for _, task := range tasks {
		dtos = append(dtos, task.ToDto())
	}

	encoded, err := json.MarshalIndent(dtos, "", "	")

	if err != nil {
		return fmt.Errorf("error during marshal: %v", err)
	}

	return os.WriteFile(path, encoded, 0644)
}
