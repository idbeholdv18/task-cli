package task

import "task-cli/src/ids"

// Find returns pointer to Task with specified id if exists and nil otherwise.
func (tasks Tasks) Find(id ids.Id) *Task {
	for _, task := range tasks {
		if task.id == id {
			return task
		}
	}
	return nil
}

// FindIndex returns index of the *Task with specified id if exists and -1 otherwise.
func (tasks Tasks) FindIndex(id ids.Id) int {
	for i, task := range tasks {
		if task.id == id {
			return i
		}
	}
	return -1
}
