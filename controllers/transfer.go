// Package controllers is a request handler package, through using the services package and return a response.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

var (
	transferService = services.NewTransferService()
)

// TransferController is a request handler for transfer.
type TransferController struct{}

// GetAllTransfers retrieves all transfers from the storage.
func (t TransferController) GetAllTransfers(c *gin.Context) {
	transfers, err := transferService.ListAllTransfers()
	response := gin.H{"data": nil, "message": "", "error": nil}
	if err != nil {
		response["message"] = "Listing all transfers failed."
		response["error"] = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response["data"] = transfers
	response["message"] = "Listing all transfers successfully."
	c.JSON(http.StatusOK, response)
}

// CreateTransfer creates a new transfer in the storage.
func (t TransferController) CreateTransfer(c *gin.Context) {
	var transfer models.Transfer
	response := gin.H{"data": nil, "message": "", "error": nil}

	if err := c.ShouldBindJSON(&transfer); err != nil {
		response["message"] = "Creating a new transfer failed."
		response["error"] = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := transferService.CreateTransfer(transfer, *services.NewAccountService()); err != nil {
		response["message"] = "Creating a new transfer failed."
		response["error"] = err.Error()
		status := http.StatusInternalServerError
		if err == models.ErrInsufficientBalance {
			status = http.StatusBadRequest
		}
		c.JSON(status, response)
		return
	}

	response["data"] = transfer
	response["message"] = "Creating a new transfer successfully."
	c.JSON(http.StatusCreated, response)
}
