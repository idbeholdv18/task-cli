package task

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func ListTasks(tasks []Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "id\tdescription\tstatus\t")
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t\n", task.id, task.Description, task.Status)
	}
	w.Flush()
}
