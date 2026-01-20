package commands

import (
	"fmt"
	"task-cli/src/config"
	"task-cli/src/task"
)

func Add(tasks task.Tasks, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("add usage: task-cli add <description>")
	}

	description := argv[0]

	tasks = tasks.Add(description)

	if err := task.WriteToJson(config.FILENAME, tasks); err != nil {
		return err
	}

	return nil
}
