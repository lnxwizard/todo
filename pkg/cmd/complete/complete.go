package complete

import (
	"github.com/lnxwizard/todo/internal/input"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdComplete() *cobra.Command {
	// suggestion strings for complete
	suggestionStrings := []string{"completee", "complet", "cmoplete", "complte", "compltee", "ccomplete", "commplete"}

	// define `complete` command
	cmd := &cobra.Command{
		Use:        "complete",
		SuggestFor: suggestionStrings,
		Short:      "Mark a to-do as completed by index.",
		Example:    "$ todo complete 2",
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := input.GetIntegerInput()
			if err != nil {
				return err
			}

			if err := todo.Todo.Complete(index); err != nil {
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
