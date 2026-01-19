package task

import "task-cli/src/shared"

func (tasks Tasks) Find(id shared.Id) (*Task, int) {
	for i, task := range tasks {
		if task.id.Compare(id) == shared.Equal {
			return task, i
		}
	}
	return nil, -1
}
