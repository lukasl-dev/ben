package main

import (
	"github.com/lukasl-dev/ben/cmd/ben/commands/run"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	cmd := cobra.Command{
		Use:   "ben",
		Short: "A command-line tool for managing sequential command procedures.",
	}
	cmd.AddCommand(run.Command())
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
