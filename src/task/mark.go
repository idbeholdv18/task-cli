package task

import (
	"fmt"
	"task-cli/src/ids"
)

func (task *Task) Mark(status StatusState) {
	task.Status = status
}

func (tasks Tasks) Mark(id ids.Id, status StatusState) error {
	candidate := tasks.Find(id)
	if candidate == nil {
		return fmt.Errorf("task with id %d not found", id)
	}

	candidate.Mark(status)
	return nil
}
