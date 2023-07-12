package cli

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
)

func PrintUsage() {
	const usage = `
					A simple to-do list app that you can use from the terminal.

					Usage:
						todo [flags]

					Examples:
						$ todo --add Go Walk
						$ todo --list
						$ todo --complete 2
						$ todo --delete 5

					Flags:
						--add			Add a new to-do.
						--delete 		Delete a to-do by index.
						--complete 		Mark a to-do as completed by index.
						--list 			List all to-dos.
						--help, -h		Print help message for todo.`

	fmt.Println(heredoc.Doc(usage))
}
