package commands

import (
	"fmt"
	"os"
	"task-cli/src/config"
	"task-cli/src/task"
)

func Add(tasks []*task.Task, argv []string) {
	if len(argv) < 1 {
		fmt.Fprintf(os.Stderr, "task description isn't provided")
		os.Exit(2)
	}

	for _, d := range argv {
		newTask := task.CreateTask(tasks, d)
		tasks = append(tasks, newTask)
	}

	task.WriteToJson(config.FILENAME, tasks)
}
