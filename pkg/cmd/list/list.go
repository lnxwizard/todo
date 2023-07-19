package list

import (
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	// suggest strings for list
	suggestionStrings := []string{"llist", "lsit", "li", "lis", "lsi", "ls", "listt", "llistt", "liist", "lisst"}

	// define `list` command
	cmd := &cobra.Command{
		Use:        "list",
		SuggestFor: suggestionStrings,
		Short:      "List all to-do's.",
		Example:    "$ todo list",
		RunE: func(cmd *cobra.Command, args []string) error {
			// get value of flag: completed
			completedFlgVal, err := cmd.Flags().GetBool("completed")
			if err != nil {
				return err
			}

			switch {
			case completedFlgVal:
				todo.Todo.ListCompleted()
				return nil
			}

			todo.Todo.List()

			return nil
		},
	}

	// define flags
	var (
		completedFlg bool
	)
	cmd.Flags().BoolVarP(&completedFlg, "completed", "c", false, "List completed to-do's.")

	return cmd
}
