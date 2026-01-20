package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/config"
	"task-cli/src/ids"
	"task-cli/src/task"
)

func Mark(tasks task.Tasks, argv []string) error {
	if len(argv) != 2 {
		return fmt.Errorf("update usage: task-cli mark <id> [backlog|ready|wip|done]")
	}

	id, err := strconv.Atoi(argv[0])
	if err != nil {
		return fmt.Errorf("error during parsing id: %w", err)
	}

	switch argv[1] {
	case task.StateName[task.StateBacklog]:
		{
			tasks.Mark(ids.Id(id), task.StateBacklog)
		}
	case task.StateName[task.StateReady]:
		{
			tasks.Mark(ids.Id(id), task.StateReady)
		}
	case task.StateName[task.StateInProgress]:
		{
			tasks.Mark(ids.Id(id), task.StateInProgress)
		}
	case task.StateName[task.StateDone]:
		{
			tasks.Mark(ids.Id(id), task.StateDone)
		}
	default:
		{
			return fmt.Errorf("mark usage: task-cli mark <id> [backlog|ready|wip|done]")
		}
	}

	if err := task.WriteToJson(config.FILENAME, tasks); err != nil {
		return err
	}

	return nil
}
