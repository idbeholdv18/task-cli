package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/config"
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

	err = tasks.Update(ids.Id(id), description)

	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	if err := task.WriteToJson(config.FILENAME, tasks); err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	return nil
}
