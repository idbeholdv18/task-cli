package commands

import (
	"fmt"
	"task-cli/src/task"
)

func Add(tasks task.Tasks, argv []string) (task.Tasks, error) {
	if len(argv) != 1 {
		return nil, fmt.Errorf("add usage: task-cli add <description>")
	}

	description := argv[0]

	tasks = tasks.Add(description)

	return tasks, nil
}
