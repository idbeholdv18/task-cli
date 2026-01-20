package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/ids"
	"task-cli/src/task"
)

func Delete(tasks task.Tasks, argv []string) (task.Tasks, error) {
	if len(argv) != 1 {
		return nil, fmt.Errorf("delete usage: task-cli delete <id>")
	}

	id, err := strconv.Atoi(argv[0])

	if err != nil {
		return nil, fmt.Errorf("parse id: %w", err)
	}

	tasks, err = tasks.Delete(ids.Id(id))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
