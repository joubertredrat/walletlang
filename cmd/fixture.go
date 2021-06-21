package cmd

import (
	"fmt"
	"os"

	"github.com/joubertredrat/walletlang/infrastructure/database"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/repository"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"github.com/spf13/cobra"
)

func NewFixtureCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "fixture",
		Short: "Popular banco com dados",
		Run: func(c *cobra.Command, args []string) {
			db := database.NewMongoDatabase(os.Getenv("MONGO_ADDR_DSN"), os.Getenv("MONGO_DATABASE"))
			repo := repository.NewUserRepositoryMongo(db)

			payers := []*entity.Payer{
				&entity.Payer{
					Amount: 100,
				},
				&entity.Payer{
					Amount: 200,
				},
				&entity.Payer{
					Amount: 150,
				},
			}

			for _, payer := range payers {
				repo.CreatePayer(payer)
			}

			fmt.Println("Done")
		},
	}
}
