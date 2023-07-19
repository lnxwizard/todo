package main

import (
	rootCmd "github.com/lnxwizard/todo/pkg/cmd/root"
	"os"
)

func main() {
	if err := rootCmd.NewCmdRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
