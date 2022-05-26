package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Hello() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "return Hello + [name]",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello %s\n", args[0])
		},
	}
}
