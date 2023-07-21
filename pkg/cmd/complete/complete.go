package complete

import (
	"errors"
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
			allFlgVal, err := cmd.Flags().GetBool("all")
			if err != nil {
				return err
			}

			switch {
			case allFlgVal:
				todo.Todo.CompleteAll()

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

			if err := todo.Todo.Complete(index); err != nil {
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
	cmd.Flags().BoolVarP(&allFlg, "all", "a", false, "Mark all to-do's as completed.")

	return cmd
}
