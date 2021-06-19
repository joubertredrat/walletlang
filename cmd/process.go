package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewProcessComand() *cobra.Command {
	return &cobra.Command{
		Use:   "process",
		Short: "Iniciar o worker para processar transações agendadas",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("running proccess")
		},
	}
}
