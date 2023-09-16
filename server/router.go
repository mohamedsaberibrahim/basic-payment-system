// Package server is the application's entry point.
// This package is used to start the application.
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedsaberibrahim/basic-payment-system/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/", health.Status)

	accountsGroup := router.Group("/accounts")
	{
		controller := new(controllers.AccountController)
		accountsGroup.GET("", controller.GetAllAccounts)
	}

	transfersGroup := router.Group("/transfers")
	{
		controller := new(controllers.TransferController)
		transfersGroup.GET("", controller.GetAllTransfers)
		transfersGroup.POST("", controller.CreateTransfer)
	}

	return router
}
