package main

import (
	"viacep-cli/src/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(commands.Hello())
	rootCmd.AddCommand(commands.Get())

	rootCmd.Execute()
}
