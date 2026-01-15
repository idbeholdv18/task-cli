package main

import (
	// "fmt"
	"fmt"
	"os"
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
				tasks = append(tasks, *newTask)
			}
		}
	case "list":
		{
			task.ListTasks(tasks)
		}
	default:
		{
			fmt.Println("usage: task-cli [add|update|delete|mark|list]")
		}
	}

	// fmt.Printf("%v\n", tasks)

	// fmt.Printf("%v\n", tasks)
	task.WriteToJson(filename, tasks)
}
