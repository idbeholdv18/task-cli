package main

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"task-cli/src/shared"
	"task-cli/src/task"
)

const filename = "data.json"

func main() {
	tasks := task.ReadFromJson(filename)
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: task-cli [add|update|delete|mark|list]\n")
		return
	}

	switch os.Args[1] {
	case "add":
		{
			for _, description := range os.Args[2:] {
				newTask := task.CreateTask(tasks, description)
				tasks = append(tasks, newTask)
			}
			task.WriteToJson(filename, tasks)
		}
	case "list":
		{
			task.ListTasks(tasks)
		}
	case "update":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "update usage: task-cli update <id> <new description>")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
				os.Exit(1)
			}
			candidate := task.FindTaskById(tasks, shared.CreateId(id))
			if candidate == nil {
				fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
			}
			candidate.Description = os.Args[3]
			task.WriteToJson(filename, tasks)
		}
	case "delete":
		{
			if len(os.Args) != 3 {
				fmt.Fprintf(os.Stderr, "update usage: task-cli delete <id>\n")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
				os.Exit(1)
			}
			candidate := task.FindTaskById(tasks, shared.CreateId(id))
			if candidate == nil {
				fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
			}
			filteredTasks := make([]*task.Task, 0, len(tasks)-1)
			for _, task := range tasks {
				if candidate != task {
					filteredTasks = append(filteredTasks, task)
				}
			}
			task.WriteToJson(filename, filteredTasks)
		}
	case "mark":
		{
			if len(os.Args) != 4 {
				fmt.Fprintf(os.Stderr, "update usage: task-cli mark <id> [ready|wip|done]\n")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during parsing id: %v\n", err)
				os.Exit(1)
			}
			candidate := task.FindTaskById(tasks, shared.CreateId(id))
			if candidate == nil {
				fmt.Fprintf(os.Stderr, "task with id %d not found\n", id)
				os.Exit(1)
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
				}
			}
			task.WriteToJson(filename, tasks)
		}
	default:
		{
			fmt.Println("usage: task-cli [add|update|delete|mark|list]")
			return
		}
	}
}
