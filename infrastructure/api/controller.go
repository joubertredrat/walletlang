package api

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joubertredrat/walletlang/infrastructure/database"
	"github.com/joubertredrat/walletlang/infrastructure/queue"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/event"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/repository"
	"github.com/joubertredrat/walletlang/internal/transfer/service"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (c *Controller) handleWalletsList(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "wallets list",
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleTransactionSchedule(ctx *gin.Context) {
	var request TransactionScheduleRequest
	ctx.ShouldBindBodyWith(&request, binding.JSON)

	db := database.NewMongoDatabase(os.Getenv("MONGO_ADDR_DSN"), os.Getenv("MONGO_DATABASE"))
	userRepo := repository.NewUserRepositoryMongo(db)
	transactionRepo := repository.NewTransactionRepositoryMongo(db)

	kafkaProducer := queue.NewKafkaProducer(os.Getenv("KAFKA_BROKER_HOST"), os.Getenv("KAFKA_BROKER_TOPIC"))
	eventKafka := event.NewTransactionEventKafka(kafkaProducer)

	createTransaction := service.NewCreateTransaction(&userRepo, &transactionRepo, eventKafka)
	transaction, err := createTransaction.HandleCreate(request.PayerID, request.PayeeID, request.Amount)

	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, service.CreateTransactionPayerNotFoundError) {
			statusCode = http.StatusUnprocessableEntity
		}

		if errors.Is(err, service.CreateTransactionPayeeNotFoundError) {
			statusCode = http.StatusUnprocessableEntity
		}

		response := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}

		ctx.JSON(statusCode, response)
		return
	}

	response := struct {
		ID      string `json:"id"`
		PayerID string `json:"payer_id"`
		PayeeID string `json:"payee_id"`
		Amount  uint   `json:"amount"`
		Status  string `json:"status"`
	}{
		ID:      transaction.ID,
		PayerID: transaction.Payer.ID,
		PayeeID: transaction.Payee.ID,
		Amount:  transaction.Amount,
		Status:  transaction.Status,
	}

	ctx.JSON(http.StatusCreated, response)
}
