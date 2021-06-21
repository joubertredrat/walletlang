package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	controller := NewController()

	router := gin.Default()
	router.POST("/api/wallets", controller.handleWalletsList)
	router.POST("/api/transactions", controller.handleTransactionSchedule)

	fmt.Println("Server running at http://0.0.0.0:8127")
	router.Run(":8127")
}
