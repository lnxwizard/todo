package list

import (
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	// suggest strings for list
	suggestionStrings := []string{"llist", "lsit", "li", "lis", "lsi", "ls", "listt", "llist", "liist", "lisst"}

	// define `list` command
	cmd := &cobra.Command{
		Use:        "list",
		SuggestFor: suggestionStrings,
		Short:      "List all to-do's.",
		Example:    "$ todo list",
		Run: func(cmd *cobra.Command, args []string) {
			todo.Todo.List()
		},
	}

	return cmd
}
