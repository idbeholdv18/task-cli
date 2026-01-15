package main

import (
	"fmt"
	"task-cli/src/task"
)

func main() {
	tasks := task.ReadFromJson("test.json")
	fmt.Printf("%v\n", tasks)

	newTask := task.CreateTask(tasks, "testing")
	tasks = append(tasks, *newTask)

	fmt.Printf("%v\n", tasks)
	task.WriteToJson("test.json", tasks)
}
