// Package services is the application's business logic.
// This package is used by the controller layer to perform business logic operations.
package services

import (
	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/storage"
)

// AccountService is the service layer for accounts.
type AccountService struct {
	accountStore *storage.AccountStore
}

// NewAccountService creates a new instance of AccountService.
func NewAccountService() *AccountService {
	return &AccountService{
		accountStore: storage.NewAccountStorage(),
	}
}

// GetAllAccounts retrieves all accounts from the storage.
func (a *AccountService) GetAllAccounts() ([]models.Account, error) {
	return a.accountStore.ListAccounts()
}

// GetAccount retrieves an account from the storage.
func (a *AccountService) GetAccount(id string) (*models.Account, error) {
	return a.accountStore.GetAccount(id)
}

// CreateAccount creates a new account in the storage.
func (a *AccountService) CreateAccount(account models.Account) (*models.Account, error) {
	return a.accountStore.CreateAccount(account)
}

// UpdateBalance updates an account's balance in the storage.
func (a *AccountService) UpdateBalance(id string, amount float64) error {
	account, err := a.accountStore.GetAccount(id)
	if err != nil {
		return err
	}
	account.Balance += amount
	return a.accountStore.UpdateAccount(*account)
}

// DeleteAllAccounts deletes all accounts from the storage.
func (a *AccountService) DeleteAllAccounts() error {
	return a.accountStore.DeleteAllAccounts()
}
