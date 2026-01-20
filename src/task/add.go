package task

import (
	"task-cli/src/ids"
	"time"
)

func createTask(id ids.Id, description string) *Task {
	return &Task{
		id:          id,
		Description: description,
		Status:      StateBacklog,
		createdAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (tasks Tasks) Add(description string) Tasks {
	id, ok := tasks.FindMaxId()

	if !ok {
		id = ids.Id(0)
	} else {
		id = id.Next()
	}

	return append(tasks, createTask(id, description))
}
