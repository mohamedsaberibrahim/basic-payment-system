// Package services is the application's business logic.
// This package is used by the controller layer to perform business logic operations.
package services

import (
	"sync"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/storage"
)

// TransferService is the service layer for transfers.
type TransferService struct {
	transferStore *storage.TransferStore
	mutex         sync.Mutex
}

// NewTransferService creates a new instance of TransferService.
func NewTransferService() *TransferService {
	return &TransferService{
		transferStore: storage.NewTransferStorage(),
	}
}

// CreateTransfer creates a new transfer in the storage.
func (t *TransferService) CreateTransfer(transfer *models.Transfer, accountService AccountService) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	fromAccount, _ := accountService.GetAccount(transfer.From)
	toAccount, _ := accountService.GetAccount(transfer.From)

	if fromAccount == nil {
		return models.ErrAccountNotFound
	}
	if toAccount == nil {
		return models.ErrAccountNotFound
	}

	if fromAccount.Balance < transfer.Amount {
		return models.ErrInsufficientBalance
	}
	transfer.ID = models.NewUUID()
	t.transferStore.CreateTransfer(*transfer)
	accountService.UpdateBalance(transfer.From, -transfer.Amount)
	accountService.UpdateBalance(transfer.To, transfer.Amount)
	return nil
}

// ListAllTransfers retrieves all transfers from the storage.
func (t *TransferService) ListAllTransfers() ([]models.Transfer, error) {
	return t.transferStore.ListTransfers()
}
