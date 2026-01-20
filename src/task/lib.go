package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-cli/src/ids"
)

func (task Task) ToDto() TaskDto {
	return TaskDto{
		Id:          task.id,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.createdAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func TaskFromDto(dto TaskDto) Task {
	return Task{
		id:          dto.Id,
		Description: dto.Description,
		Status:      dto.Status,
		createdAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

// FindMaxId the maximum id among all the Task elements
//
// If tasks doesn't contain any elements returns -1
//
// Although in this instance, it is possible to retrieve the ID of the last element,
// the function still iterates over the whole slice, as a change in behaviour or data
// structure could result in the maximum ID not being the last element.
func (tasks Tasks) FindMaxId() ids.Id {
	if len(tasks) == 0 {
		return -1
	}
	max := tasks[0].id
	for _, task := range tasks {
		if task.id.Compare(max) == ids.Greater {
			max = task.id
		}
	}
	return max
}

func ReadFromJson(path string) Tasks {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var dtos []TaskDto
	err = json.Unmarshal(content, &dtos)

	tasks := make([]*Task, 0, len(dtos))

	for _, dto := range dtos {
		taskFromDto := TaskFromDto(dto)
		tasks = append(tasks, &taskFromDto)
	}

	if err != nil {
		log.Fatal("Error during unmarshal: ", err)
	}

	return tasks
}

func WriteToJson(path string, tasks Tasks) error {
	dtos := make([]TaskDto, 0, len(tasks))

	for _, task := range tasks {
		dtos = append(dtos, (*task).ToDto())
	}

	encoded, err := json.Marshal(dtos)

	if err != nil {
		return fmt.Errorf("error during marshal: %v", err)
	}

	if err := os.WriteFile(path, encoded, 0644); err != nil {
		return fmt.Errorf("error during writing file: %v", err)
	}

	return nil
}
