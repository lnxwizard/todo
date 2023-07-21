package mark

import (
	"github.com/MakeNowJust/heredoc"
	importantCmd "github.com/lnxwizard/todo/pkg/cmd/mark/important"
	"github.com/spf13/cobra"
)

func NewCmdMark() *cobra.Command {
	// suggestion strings for mark command
	suggestionStrings := []string{"mrak", "marrk", "markk", "mmark", "ark", "mrka", "mrakkk"}

	// define mark command
	cmd := &cobra.Command{
		Use:        "mark",
		SuggestFor: suggestionStrings,
		Short:      "Mark your to-do.",
		Example: heredoc.Doc(`
		$ todo mark important 2`),
	}

	// add sub-commands
	cmd.AddCommand(importantCmd.NewCmdImportant())

	return cmd
}
