package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"task-cli/src/task"
)

func ReadFromJson(f *os.File) (task.Tasks, error) {
	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	content, err := io.ReadAll(f)

	if err != nil {
		return nil, fmt.Errorf("Error during read file: %w", err)
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

func WriteToJson(f *os.File, tasks task.Tasks) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	if err := f.Truncate(0); err != nil {
		return err
	}

	dtos := make([]task.TaskDto, 0, len(tasks))

	for _, task := range tasks {
		dtos = append(dtos, task.ToDto())
	}

	encoded, err := json.MarshalIndent(dtos, "", "	")

	if err != nil {
		return fmt.Errorf("error during marshal: %v", err)
	}

	if _, err := f.Write(encoded); err != nil {
		return fmt.Errorf("error during writing file: %v", err)
	}

	return nil
}
