package task

import (
	"fmt"
	"task-cli/src/ids"
)

func (tasks Tasks) Delete(id ids.Id) (Tasks, error) {
	n := tasks.FindIndex(id)
	if n == -1 {
		return nil, fmt.Errorf("task with id %d not found\n", id)
	}
	return append(tasks[0:n], tasks[n+1:]...), nil
}
