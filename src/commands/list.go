package commands

import (
	"task-cli/src/task"
)

func List(tasks []*task.Task) {
	task.ListTasks(tasks)
}
