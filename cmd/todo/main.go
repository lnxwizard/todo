package main

import (
	"flag"
	"fmt"
	"github.com/lnxwizard/todo/internal/cli"
	"github.com/lnxwizard/todo/pkg/todo"
	"os"
)

// change version manually
const (
	todoVersion = "1.0.0"
)

const todoFile = "todos.json"

func main() {
	// set custom help message
	flag.Usage = cli.PrintUsage

	// define flags
	var (
		addFlg      bool
		deleteFlg   int
		completeFlg int
		listFlg     bool
	)
	flag.BoolVar(&addFlg, "add", false, "Add a new to-do.")
	flag.IntVar(&deleteFlg, "delete", 0, "Delete a to-do by index.")
	flag.IntVar(&completeFlg, "complete", 0, "Mark a to-do as completed by index.")
	flag.BoolVar(&listFlg, "list", false, "List all to-dos.")

	// parse flags
	flag.Parse()

	todos := &todo.TodoList{}

	// Load all to-dos from json file
	if err := todos.Load(todoFile); err != nil {
		fmt.Printf("Error while loading to-do file: %s\n", err)
		os.Exit(1)
	}

	switch {
	case addFlg:
		task, err := cli.GetInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Printf("Error while getting to-do: %s\n", err)
			os.Exit(1)
		}

		todos.Add(task)

		if err = todos.Store(todoFile); err != nil {
			fmt.Printf("Error while loading to-do's from `%s` file: %s\n", todoFile, err)
			os.Exit(1)
		}
	case deleteFlg > 0:
		if err := todos.Delete(deleteFlg); err != nil {
			fmt.Printf("Error while deleting a to-do: %s\n", err)
			os.Exit(1)
		}

		if err := todos.Store(todoFile); err != nil {
			fmt.Printf("Error while loading to-do's from `%s` file: %s\n", todoFile, err)
			os.Exit(1)
		}
	case completeFlg > 0:
		if err := todos.Complete(completeFlg); err != nil {
			fmt.Printf("Error while marking a to-do as completed: %s\n", err)
			os.Exit(1)
		}

		if err := todos.Store(todoFile); err != nil {
			fmt.Printf("Error while loading to-do's from `%s` file: %s\n", todoFile, err)
			os.Exit(1)
		}
	case listFlg:
		todos.List()
	default:
		cli.PrintUsage()
	}
}
