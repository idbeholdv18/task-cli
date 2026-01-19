package commands

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/src/config"
	"task-cli/src/shared"
	"task-cli/src/task"
)

func Update(tasks []*task.Task, argv []string) {
	if len(argv) != 3 {
		fmt.Fprintf(os.Stderr, "update usage: task-cli update <id> <new description>")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
		os.Exit(2)
	}
	candidate := task.FindTaskById(tasks, shared.Id(id))
	if candidate == nil {
		fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
		os.Exit(1)
	}
	candidate.Description = os.Args[3]
	task.WriteToJson(config.FILENAME, tasks)
}
