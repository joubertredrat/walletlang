package cmd

import (
	"github.com/joubertredrat/walletlang/infrastructure/api"
	"github.com/spf13/cobra"
)

func NewApiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Iniciar o http api",
		Run: func(c *cobra.Command, args []string) {
			api.Run()
		},
	}
}
