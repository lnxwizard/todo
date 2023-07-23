package uncomplete

import (
	"errors"
	"github.com/lnxwizard/todo/internal/input"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdDone() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "done",
		SuggestFor: nil,
		Short:      "Unmark a to-do as done.",
		Example:    "$ todo unmark done 2",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 0 {
				return errors.New("index cannot be empty")
			}

			index, err := input.GetIntegerInput(3)
			if err != nil {
				return err
			}

			if err := todo.Todo.Uncomplete(index); err != nil {
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
