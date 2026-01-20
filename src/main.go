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
			if err := commands.Add(tasks, args); err != nil {
				fmt.Fprintf(os.Stderr, "add error: %s\n", err)
				os.Exit(2)
			}
		}
	case "list":
		{
			commands.List(tasks)
		}
	case "update":
		{
			if err := commands.Update(tasks, args); err != nil {
				fmt.Fprintf(os.Stderr, "update error: %s\n", err)
				os.Exit(2)
			}
		}
	case "delete":
		{
			if err := commands.Delete(tasks, args); err != nil {
				fmt.Fprintf(os.Stderr, "delete error: %s\n", err)
				os.Exit(2)
			}
		}
	case "mark":
		{
			if err := commands.Mark(tasks, args); err != nil {
				fmt.Fprintf(os.Stderr, "mark error: %s\n", err)
				os.Exit(2)
			}
		}
	default:
		{
			fmt.Println("usage: task-cli [add|update|delete|mark|list]")
			os.Exit(2)
		}
	}
}
