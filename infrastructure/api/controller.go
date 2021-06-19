package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (c *Controller) handleWalletsSeed(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "seed aplied",
	}

	ctx.JSON(http.StatusCreated, response)
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
	response := struct {
		Message string `json:"message"`
	}{
		Message: "created",
	}

	ctx.JSON(http.StatusCreated, response)
}
