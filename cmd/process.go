package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joubertredrat/walletlang/infrastructure/database"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/event/entity"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/repository"
	"github.com/joubertredrat/walletlang/internal/transfer/service"
	"github.com/spf13/cobra"
)

func NewProcessComand() *cobra.Command {
	return &cobra.Command{
		Use:   "process",
		Short: "Iniciar o worker para processar transações agendadas",
		Run: func(c *cobra.Command, args []string) {
			db := database.NewMongoDatabase(os.Getenv("MONGO_ADDR_DSN"), os.Getenv("MONGO_DATABASE"))
			userRepo := repository.NewUserRepositoryMongo(db)
			transactionRepo := repository.NewTransactionRepositoryMongo(db)
			processTransactionService := service.NewProcessTransaction(&userRepo, &transactionRepo)

			config := sarama.NewConfig()
			config.Consumer.Return.Errors = true
			master, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_BROKER_HOST")}, config)
			if err != nil {
				log.Panic(err)
			}
			defer func() {
				if err := master.Close(); err != nil {
					log.Panic(err)
				}
			}()
			consumer, err := master.ConsumePartition(os.Getenv("KAFKA_BROKER_TOPIC"), 0, sarama.OffsetOldest)
			if err != nil {
				log.Panic(err)
			}
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, os.Interrupt)
			doneCh := make(chan struct{})

			log.Println("Starting running proccess worker")
			go func() {
				for {
					select {
					case err := <-consumer.Errors():
						log.Println(err)
					case msg := <-consumer.Messages():
						time.Sleep(8 * time.Second)
						log.Println("Received event", string(msg.Key), string(msg.Value))
						var transactionScheduleRequest entity.TransactionScheduled
						json.Unmarshal([]byte(msg.Value), &transactionScheduleRequest)

						_, err := processTransactionService.HandleProcess(transactionScheduleRequest.ID)
						if err != nil {
							log.Println("event not processed, errors found", string(msg.Key), string(msg.Value), err.Error())
						} else {
							log.Println("event processed sucessfully", string(msg.Key), string(msg.Value))
						}

					case <-signals:
						log.Println("Interrupt is detected")
						doneCh <- struct{}{}
					}
				}
			}()
			<-doneCh
		},
	}
}
