package task

import (
	"task-cli/src/shared"
	"time"
)

type StatusState int

const (
	StateBacklog = iota
	StateReady
	StateInProgress
	StateDone
)

var stateName = map[StatusState]string{
	StateBacklog:    "Backlog",
	StateReady:      "Ready",
	StateInProgress: "In Progress",
	StateDone:       "Done",
}

// implementing Stringer interface
func (statusState StatusState) String() string {
	return stateName[statusState]
}

type Task struct {
	id          shared.Id
	Description string
	Status      StatusState
	createdAt   time.Time
	UpdatedAt   time.Time
}

func (task Task) Id() shared.Id {
	return task.id
}

func (task Task) CreatedAt() time.Time {
	return task.createdAt
}
