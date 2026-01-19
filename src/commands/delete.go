package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/config"
	"task-cli/src/shared"
	"task-cli/src/task"
)

func Delete(tasks task.Tasks, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("delete usage: task-cli delete <id>\n")
	}

	id, err := strconv.Atoi(argv[0])

	if err != nil {
		return fmt.Errorf("parse id: %w\n", err)
	}

	tasks, err = tasks.Delete(shared.Id(id))

	if err != nil {
		return err
	}

	if err := task.WriteToJson(config.FILENAME, tasks); err != nil {
		return err
	}

	return nil
}
