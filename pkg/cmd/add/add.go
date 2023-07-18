package add

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/lnxwizard/todo/internal/input"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdAdd() *cobra.Command {
	// suggestion strings for add
	suggestionStrings := []string{"addd", "aad", "aadd", "ad"}

	// define `add` command
	cmd := &cobra.Command{
		Use:        "add",
		SuggestFor: suggestionStrings,
		Short:      "Add a new to-do.",
		Example: heredoc.Doc(`
		$ todo add Take a walk`),
		RunE: func(cmd *cobra.Command, args []string) error {
			task, err := input.GetInput(os.Stdin, args...)
			if err != nil {
				return err
			}

			todo.Todo.Add(task)

			if err := todo.Todo.Store(todo.TodoFile); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
