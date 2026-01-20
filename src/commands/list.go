package commands

import "task-cli/src/task"

func List(tasks task.Tasks) {
	tasks.List()
}
