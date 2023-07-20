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
			// get value of flags
			doneFlgVal, err := cmd.Flags().GetBool("done")
			if err != nil {
				return err
			}
			importantFlgVal, err := cmd.Flags().GetBool("important")
			if err != nil {
				return err
			}

			switch {
			case doneFlgVal:
				todo.Todo.ListCompleted()
				return nil
			case importantFlgVal:
				todo.Todo.ListImportant()
				return nil
			}

			todo.Todo.List()

			return nil
		},
	}

	// define flags
	var (
		doneFlg      bool
		importantFlg bool
	)
	cmd.Flags().BoolVarP(&doneFlg, "done", "d", false, "List completed to-do's.")
	cmd.Flags().BoolVarP(&importantFlg, "important", "i", false, "List important to-do's.")

	return cmd
}
