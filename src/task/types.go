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

type TaskDto struct {
	Id          shared.Id   `json:"id"`
	Description string      `json:"description"`
	Status      StatusState `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}


