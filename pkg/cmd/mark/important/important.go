package important

import (
	"errors"
	"github.com/lnxwizard/todo/internal/input"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
)

func NewCmdImportant() *cobra.Command {
	// define suggestion strings for important command
	suggestionStrings := []string{"improtant", "import", "improtatn", "importtant", "importtantt", "importantt"}

	// define command: important
	cmd := &cobra.Command{
		Use:        "important",
		SuggestFor: suggestionStrings,
		Short:      "Mark your to-do as important.",
		Example:    "$ todo mark important 3",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("index cannot be empty")
			}

			index, err := input.GetIntegerInput(3)
			if err != nil {
				return err
			}

			if err := todo.Todo.MarkImportant(index); err != nil {
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
