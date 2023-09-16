package tests

import (
	"sync"
	"testing"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

func TestTransferService_CreateTransfer(t *testing.T) {
	// Create TransferService and AccountService instances
	transferService := services.NewTransferService()
	accountService := services.NewAccountService()

	t.Run("Test valid transfer", func(t *testing.T) {
		// Define accounts initial balances
		account1Balance := 3000.0
		account2Balance := 500.0

		// Create two accounts for testing
		account1, account2 := createAccounts(t, account1Balance, account2Balance, accountService)

		// Create a transfer from account1 to account2
		transfer := models.Transfer{
			From:   account1.ID,
			To:     account2.ID,
			Amount: 100,
		}
		err := transferService.CreateTransfer(transfer, *accountService)
		acc1, _ := accountService.GetAccount(transfer.From)
		acc2, _ := accountService.GetAccount(transfer.To)

		assertError(t, err, nil)
		assertFloats(t, acc1.Balance, account1Balance-transfer.Amount)
		assertFloats(t, acc2.Balance, account2Balance+transfer.Amount)
	})

	t.Run("Test invalid transfer", func(t *testing.T) {
		// Define accounts initial balances
		account1Balance := 3000.0
		account2Balance := 500.0

		// Create two accounts for testing
		account1, account2 := createAccounts(t, account1Balance, account2Balance, accountService)

		// Create a transfer from account1 to account2
		transfer := models.Transfer{
			From:   account1.ID,
			To:     account2.ID,
			Amount: -1000,
		}
		err := transferService.CreateTransfer(transfer, *accountService)
		acc1, _ := accountService.GetAccount(transfer.From)
		acc2, _ := accountService.GetAccount(transfer.To)

		assertError(t, err, models.ErrInsufficientBalance)
		assertFloats(t, acc1.Balance, account1Balance)
		assertFloats(t, acc2.Balance, account2Balance)
	})

	t.Run("Test concurrent transactions", func(t *testing.T) {
		// Define accounts initial balances
		account1Balance := 3000.0
		account2Balance := 500.0

		// Create two accounts for testing
		account1, account2 := createAccounts(t, account1Balance, account2Balance, accountService)

		// Run multiple goroutines to simulate concurrent transfers
		concurrentTransfers := 100
		var wg sync.WaitGroup
		wg.Add(concurrentTransfers)
		for i := 0; i < concurrentTransfers; i++ {
			go func() {
				defer wg.Done()

				// Create a transfer from account1 to account2
				transfer := models.Transfer{
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
		updatedBalance1, _ := accountService.GetAccount(account1.ID)
		updatedBalance2, _ := accountService.GetAccount(account2.ID)
		assertFloats(t, updatedBalance1.Balance, account1Balance-float64(10*concurrentTransfers))
		assertFloats(t, updatedBalance2.Balance, account2Balance+float64(10*concurrentTransfers))
	})
}
