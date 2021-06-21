package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Execute() {
	filename := filepath.Base(os.Args[0])

	rootCmd := &cobra.Command{
		Use:   filename,
		Short: "Walletlang: money transfer and wallet movement using golang",
	}

	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	rootCmd.AddCommand(NewApiCommand())
	rootCmd.AddCommand(NewProcessComand())
	rootCmd.AddCommand(NewFixtureCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(nil, "Error to execute cobra command", err)
	}
}
