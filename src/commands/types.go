package commands

import "task-cli/src/task"

var StateMap = map[string]task.StatusState{
	"backlog": task.StateBacklog,
	"ready":   task.StateReady,
	"wip":     task.StateInProgress,
	"done":    task.StateDone,
}
