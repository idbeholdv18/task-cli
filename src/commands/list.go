package commands

import (
	"io"
	"task-cli/src/task"
)

func List(w io.Writer, tasks task.Tasks) {
	tasks.List(w)
}
