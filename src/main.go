package main

import (
	"fmt"
	"os"
	"task-cli/src/commands"
	"task-cli/src/config"
	"task-cli/src/task"
)

func main() {
	tasks := task.ReadFromJson(config.FILENAME)

	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: task-cli [add|update|delete|mark|list]\n")
		return
	}

	args := os.Args[2:]

	switch os.Args[1] {
	case "add":
		{
			commands.Add(tasks, args)
		}
	case "list":
		{
			commands.List(tasks)
		}
	case "update":
		{
			commands.Update(tasks, args)
		}
	case "delete":
		{
			commands.Delete(tasks, args)
		}
	case "mark":
		{
			commands.Mark(tasks, args)
		}
	default:
		{
			fmt.Println("usage: task-cli [add|update|delete|mark|list]")
			os.Exit(2)
		}
	}
}
