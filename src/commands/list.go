package commands

import (
	"fmt"
	"task-cli/src/task"
)

func List(tasks task.Tasks, argv []string) (task.Tasks, error) {
	if len(argv) > 1 {
		return nil, fmt.Errorf("usage: task-cli list [backlog|ready|wip|done]")
	}

	if len(argv) == 0 {
		return tasks, nil
	}

	status, ok := StateMap[argv[0]]
	if !ok {
		return nil, fmt.Errorf("usage: task-cli list [backlog|ready|wip|done]")
	}

	return tasks.Filter(status), nil
}
