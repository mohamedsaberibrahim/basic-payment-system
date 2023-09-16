package main

import (
	"github.com/mohamedsaberibrahim/basic-payment-system/server"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

const (
	ACCOUNTS_LINK  = "https://git.io/Jm76h"
	NUM_OF_WORKERS = 10
)

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	services.ProcessAccounts(ACCOUNTS_LINK, NUM_OF_WORKERS)

	// Start the server
	server.Init()
}
