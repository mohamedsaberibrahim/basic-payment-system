// Package controllers is a request handler package, through using the services package and return a response.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

var (
	accountService = services.NewAccountService()
)

// AccountController is a request handler for account.
type AccountController struct{}

// GetAllAccounts retrieves all accounts from the storage.
func (a AccountController) GetAllAccounts(c *gin.Context) {
	accounts, err := accountService.GetAllAccounts()
	response := gin.H{"data": nil, "message": "", "error": nil}
	if err != nil {
		response["message"] = "Listing all accounts failed."
		response["error"] = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response["data"] = accounts
	response["message"] = "Listing all accounts successfully."
	c.JSON(http.StatusOK, response)
}
