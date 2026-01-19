package task

import "task-cli/src/shared"

// Find returns pointer to Task with specified id if exists and nil otherwise.
func (tasks Tasks) Find(id shared.Id) *Task {
	for _, task := range tasks {
		if task.id.Compare(id) == shared.Equal {
			return task
		}
	}
	return nil
}

// FindIndex returns index of the *Task with specified id if exists and -1 otherwise.
func (tasks Tasks) FindIndex(id shared.Id) int {
	for i, task := range tasks {
		if task.id.Compare(id) == shared.Equal {
			return i
		}
	}
	return -1
}
