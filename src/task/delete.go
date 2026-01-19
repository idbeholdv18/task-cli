package task

import (
	"fmt"
	"task-cli/src/shared"
)

func (tasks Tasks) Delete(id shared.Id) (Tasks, error) {
	n := tasks.FindIndex(id)
	if n == -1 {
		return nil, fmt.Errorf("task with id %d not found\n", id)
	}
	return append(tasks[0:n], tasks[n+1:]...), nil
}
