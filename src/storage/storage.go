package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-cli/src/task"
)

func ReadFromJson(path string) (task.Tasks, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Error when opening file: %w", err)
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

	if err := os.WriteFile(path, encoded, 0644); err != nil {
		return fmt.Errorf("error during writing file: %v", err)
	}

	return nil
}
