package delete

import (
	"github.com/lnxwizard/todo/internal/input"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	// suggestion strings for delete
	suggestionStrings := []string{"del", "dlete", "deelete", "deleete", "deletee", "delet", "delede", "deleetee"}

	// define `delete` command
	cmd := &cobra.Command{
		Use:        "delete",
		SuggestFor: suggestionStrings,
		Short:      "Delete a to-do by index.",
		Example:    "$ todo delete 5",
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := input.GetIntegerInput()
			if err != nil {
				return err
			}

			if err := todo.Todo.Delete(index); err != nil {
				return err
			}

			if err := todo.Todo.Store(todo.TodoFile); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
