package commands

import (
	"fmt"
	"strconv"
	"task-cli/src/ids"
	"task-cli/src/task"
)

var stateMap = map[string]task.StatusState{
	"backlog": task.StateBacklog,
	"ready":   task.StateReady,
	"wip":     task.StateInProgress,
	"done":    task.StateDone,
}

func Mark(tasks task.Tasks, argv []string) error {
	if len(argv) != 2 {
		return fmt.Errorf("update usage: task-cli mark <id> [backlog|ready|wip|done]")
	}

	id, err := strconv.Atoi(argv[0])
	if err != nil {
		return fmt.Errorf("error during parsing id: %w", err)
	}

	state, ok := stateMap[argv[1]]
	if !ok {
		return fmt.Errorf("mark usage: task-cli mark <id> [backlog|ready|wip|done]")
	}

	if err := tasks.Mark(ids.Id(id), state); err != nil {
		return fmt.Errorf("mark error: %w", err)
	}

	return nil
}
