package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/ids"
	"task-cli/src/task"
)

func Update(tasks task.Tasks, argv []string) error {
	if len(argv) != 2 {
		return fmt.Errorf("update usage: task-cli update <id> <new description>")
	}

	id, err := strconv.Atoi(argv[0])
	if err != nil {
		return fmt.Errorf("error during parsing id: %w", err)
	}

	description := argv[1]

	if err := tasks.Update(ids.Id(id), description); err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}
