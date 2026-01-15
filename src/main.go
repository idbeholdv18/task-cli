package main

import (
	"fmt"
	"task-cli/src/task"
)

const filename = "data.json"

func main() {

	tasks := task.ReadFromJson(filename)
	fmt.Printf("%v\n", tasks)

	newTask := task.CreateTask(tasks, "testing")
	tasks = append(tasks, *newTask)

	fmt.Printf("%v\n", tasks)
	task.WriteToJson(filename, tasks)
}
