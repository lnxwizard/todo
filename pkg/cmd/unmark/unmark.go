package unmark

import (
	"github.com/MakeNowJust/heredoc"
	uncompleteCmd "github.com/lnxwizard/todo/pkg/cmd/unmark/uncomplete"
	unimportantCmd "github.com/lnxwizard/todo/pkg/cmd/unmark/unimportant"
	"github.com/spf13/cobra"
)

func NewCmdUnmark() *cobra.Command {
	// suggestion strings for unmark command
	suggestionStrings := []string{"unmrka", "unmarkk", "unnmark", "uunnmark", "uunmrak", "numark"}

	cmd := &cobra.Command{
		Use:        "unmark",
		SuggestFor: suggestionStrings,
		Short:      "Unmark your to-do.",
		Example: heredoc.Doc(`
		$ todo unmark important 3
		$ todo unmark done 5`),
	}

	// add sub-commands
	cmd.AddCommand(uncompleteCmd.NewCmdDone())
	cmd.AddCommand(unimportantCmd.NewCmdImportant())

	return cmd
}
