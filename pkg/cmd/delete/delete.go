package delete

import (
	"errors"
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
			allFlgVal, err := cmd.Flags().GetBool("all")
			if err != nil {
				return err
			}

			switch {
			case allFlgVal:
				todo.Todo.DeleteAll()

				if err := todo.Todo.Store(todo.TodoFile); err != nil {
					return err
				}

				return nil
			}

			if len(args) < 1 {
				return errors.New("index cannot be empty")
			}

			index, err := input.GetIntegerInput(2)
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

	// define flags
	var (
		allFlg bool
	)
	cmd.Flags().BoolVarP(&allFlg, "all", "a", false, "Delete all to-do's.")

	return cmd
}
