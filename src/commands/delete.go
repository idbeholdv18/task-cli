package commands

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/src/config"
	"task-cli/src/shared"
	"task-cli/src/task"
)

func deleteTask(tasks []*task.Task, id int) []*task.Task {
	candidate := task.FindTaskById(tasks, shared.Id(id))
	if candidate == nil {
		fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
		return tasks
	}
	filteredTasks := make([]*task.Task, 0, len(tasks)-1)
	for _, task := range tasks {
		if candidate != task {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func Delete(tasks []*task.Task, argv []string) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "update usage: task-cli delete <id[]>\n")
		return
	}

	ids := make([]int, 0, len(argv))

	for _, candidate := range argv {
		id, err := strconv.Atoi(candidate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
			os.Exit(2)
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		tasks = deleteTask(tasks, id)
	}
	task.WriteToJson(config.FILENAME, tasks)
}
