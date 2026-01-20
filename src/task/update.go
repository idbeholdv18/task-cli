package task

import (
	"fmt"
	"task-cli/src/ids"
	"time"
)

func (task *Task) Update(description string) {
	task.Description = description
	task.UpdatedAt = time.Now()
}

func (tasks Tasks) Update(id ids.Id, description string) error {
	candidate := tasks.Find(id)

	if candidate == nil {
		return fmt.Errorf("task with id %d not found", id)
	}

	candidate.Update(description)
	return nil
}
