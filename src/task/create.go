package task

import (
	"time"
)

func CreateTask(tasks []Task, description string) *Task {
	return &Task{
		id:          CalculateMaxId(tasks).Next(),
		Description: description,
		Status:      StateBacklog,
		createdAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
