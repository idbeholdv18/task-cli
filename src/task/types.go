package task

import (
	"task-cli/src/ids"
	"time"
)

type StatusState int

const (
	StateBacklog = iota
	StateReady
	StateInProgress
	StateDone
)

var StateName = map[StatusState]string{
	StateBacklog:    "backlog",
	StateReady:      "ready",
	StateInProgress: "wip",
	StateDone:       "done",
}

// implementing Stringer interface
func (statusState StatusState) String() string {
	return StateName[statusState]
}

type Task struct {
	id          ids.Id
	Description string
	Status      StatusState
	createdAt   time.Time
	UpdatedAt   time.Time
}

// Tasks is a slice of non-nil *Task
type Tasks []*Task

type TaskDto struct {
	Id          ids.Id      `json:"id"`
	Description string      `json:"description"`
	Status      StatusState `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
