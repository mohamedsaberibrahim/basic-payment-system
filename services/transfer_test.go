package services

import (
	"sync"
	"testing"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
)

func TestTransferService_CreateTransfer(t *testing.T) {
	// Create TransferService and AccountService instances
	transferService := NewTransferService()
	accountService := NewAccountService()

	// Define accounts initial balances
	account1Balance := 3000.0
	account2Balance := 500.0

	// Create two accounts for testing
	account1 := models.Account{ID: "1", Name: "Test Account 1", Balance: account1Balance}
	account2 := models.Account{ID: "2", Name: "Test Account 2", Balance: account2Balance}
	// Create two accounts for testing
	accountService.CreateAccount(account1)
	accountService.CreateAccount(account2)

	// Run multiple goroutines to simulate concurrent transfers
	concurrentTransfers := 100
	var wg sync.WaitGroup
	wg.Add(concurrentTransfers)
	for i := 0; i < concurrentTransfers; i++ {
		go func() {
			defer wg.Done()

			// Create a transfer from account1 to account2
			transfer := &models.Transfer{
				From:   account1.ID,
				To:     account2.ID,
				Amount: 10,
			}
			err := transferService.CreateTransfer(transfer, *accountService)
			if err != nil {
				t.Errorf("Error creating transfer: %s", err.Error())
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify that the account balances are updated correctly
	expectedBalance1 := account1Balance - float64(10*concurrentTransfers)
	expectedBalance2 := account2Balance + float64(10*concurrentTransfers)
	updatedBalance1, _ := accountService.GetAccount(account1.ID)
	updatedBalance2, _ := accountService.GetAccount(account2.ID)
	if updatedBalance1.Balance != expectedBalance1 {
		t.Errorf("Incorrect balance for account1. Expected: %f, Actual: %f", expectedBalance1, updatedBalance1.Balance)
	}
	if updatedBalance2.Balance != expectedBalance2 {
		t.Errorf("Incorrect balance for account2. Expected: %f, Actual: %f", expectedBalance2, updatedBalance2.Balance)
	}
}
