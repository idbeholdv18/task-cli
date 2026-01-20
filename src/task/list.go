package task

import (
	"fmt"
	"io"
	"text/tabwriter"
)

func (tasks Tasks) List(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 3, ' ', 0)
	fmt.Fprintln(tw, "id\tdescription\tstatus\t")
	for _, task := range tasks {
		fmt.Fprintf(tw, "%d\t%s\t%s\t\n", task.id, task.Description, task.Status)
	}
	tw.Flush()
}
