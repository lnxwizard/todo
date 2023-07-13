package todo

import (
	"fmt"
	"strings"
	"time"
)

// List all to-dos
func (t *TodoList) List() {
	fmt.Printf("+%s+\n", strings.Repeat("-", 112))
	fmt.Println("| # | Task                                          | Is Done? | CreatedAt              | CompletedAt            |")
	fmt.Println("|---|-----------------------------------------------|----------|------------------------|------------------------|")

	for index, item := range *t {
		index++
		task := item.Task
		done := "Yes"
		createdAt := item.CreatedAt.Format(time.RFC822)
		completedAt := item.CompletedAt.Format(time.RFC822)

		if !item.Done {
			done = "No"
			completedAt = item.CompletedAt.Format("###################")
		}

		fmt.Printf("| %d | %-45s | %-8s | %-22s | %-22s |\n", index, task, done, createdAt, completedAt)
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", 112))
	fmt.Printf("| You have %d pending to-do's%s|\n", t.CountPending(), strings.Repeat(" ", 85))
	fmt.Printf("+%s+\n", strings.Repeat("-", 112))
}
