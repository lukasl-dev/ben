package main

import (
	"github.com/lukasl-dev/ben/cmd/ben/commands/run"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	cmd := cobra.Command{
		Use:   "ben",
		Short: "Ben is a tool for managing boilerplate code.",
	}
	cmd.AddCommand(run.Command())
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
