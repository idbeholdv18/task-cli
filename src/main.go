package main

import (
	"fmt"
	"os"
	"task-cli/src/commands"
	"task-cli/src/storage"
)

const path = "data.json"
const (
	ExitCodeReadWrite = 1
	ExitCodeUsage     = 2
)

func main() {
	tasks, err := storage.ReadFromJson(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(ExitCodeReadWrite)
	}

	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: task-cli [add|update|delete|mark|list]\n")
		os.Exit(ExitCodeUsage)
	}

	args := os.Args[2:]

	switch os.Args[1] {
	case "add":
		tasks, err := commands.Add(tasks, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "add error: %s\n", err)
			os.Exit(ExitCodeUsage)
		}
		storage.WriteToJson(path, tasks)
	case "list":
		commands.List(os.Stdout, tasks)
	case "update":
		if err := commands.Update(tasks, args); err != nil {
			fmt.Fprintf(os.Stderr, "update error: %s\n", err)
			os.Exit(ExitCodeUsage)
		}
		storage.WriteToJson(path, tasks)
	case "delete":
		tasks, err := commands.Delete(tasks, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "delete error: %s\n", err)
			os.Exit(ExitCodeUsage)
		}
		storage.WriteToJson(path, tasks)
	case "mark":
		if err := commands.Mark(tasks, args); err != nil {
			fmt.Fprintf(os.Stderr, "mark error: %s\n", err)
			os.Exit(ExitCodeUsage)
		}
		storage.WriteToJson(path, tasks)
	default:
		fmt.Println("usage: task-cli [add|update|delete|mark|list]")
		os.Exit(ExitCodeUsage)
	}
}

func GetFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}
