package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewApiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Iniciar o http api",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("running api")
		},
	}
}
