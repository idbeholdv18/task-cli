package commands

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/src/config"
	"task-cli/src/shared"
	"task-cli/src/task"
)

func Mark(tasks []*task.Task, argv []string) {
	if len(argv) != 2 {
		fmt.Fprintf(os.Stderr, "update usage: task-cli mark <id> [ready|wip|done]\n")
		os.Exit(2)
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
		os.Exit(2)
	}
	candidate := task.FindTaskById(tasks, shared.Id(id))
	if candidate == nil {
		fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
		os.Exit(2)
	}
	switch os.Args[3] {
	case task.StateName[task.StateReady]:
		{
			candidate.Status = task.StateReady
		}
	case task.StateName[task.StateInProgress]:
		{
			candidate.Status = task.StateInProgress
		}
	case task.StateName[task.StateDone]:
		{
			candidate.Status = task.StateDone
		}
	default:
		{
			fmt.Fprintf(os.Stderr, "mark usage: task-cli mark <id> [ready|wip|done]")
			os.Exit(2)
		}
	}
	task.WriteToJson(config.FILENAME, tasks)
}
