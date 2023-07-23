package root

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	addCmd "github.com/lnxwizard/todo/pkg/cmd/add"
	completeCmd "github.com/lnxwizard/todo/pkg/cmd/complete"
	deleteCmd "github.com/lnxwizard/todo/pkg/cmd/delete"
	listCmd "github.com/lnxwizard/todo/pkg/cmd/list"
	markCmd "github.com/lnxwizard/todo/pkg/cmd/mark"
	unmarkCmd "github.com/lnxwizard/todo/pkg/cmd/unmark"
	"github.com/lnxwizard/todo/pkg/todo"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdRoot() *cobra.Command {
	// define root command
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple to-do list application.",
		Long:  "A simple to-do list app that you can use from the terminal.",
		Example: heredoc.Doc(`
		$ todo add Take a walk
		$ todo complete 3
		$ todo list
		$ todo delete 1`),
		Version: "1.3.0",
	}

	// add commands (a-z)
	cmd.AddCommand(addCmd.NewCmdAdd())
	cmd.AddCommand(completeCmd.NewCmdComplete())
	cmd.AddCommand(deleteCmd.NewCmdDelete())
	cmd.AddCommand(listCmd.NewCmdList())
	cmd.AddCommand(markCmd.NewCmdMark())
	cmd.AddCommand(unmarkCmd.NewCmdUnmark())

	return cmd
}

func init() {
	// load to-do file
	if err := todo.Todo.Load(todo.TodoFile); err != nil {
		fmt.Printf("Error while loading %s file: %s \n", todo.TodoFile, err)
		os.Exit(1)
	}
}
