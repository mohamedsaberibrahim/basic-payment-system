// Package storage is an in-memory storage of accounts and transfers.
// This is a simple implementation that uses a map to store accounts and transfers.
// This implementation is not thread-safe.
package storage

import "github.com/mohamedsaberibrahim/basic-payment-system/models"

// A TransferStore struct represents an in-memory storage of transfers.
type TransferStore struct {
	// transfers is a map of transfer IDs to transfers.
	transfers map[string]models.Transfer
}

// TransferStoreInstance is a singleton instance of TransferStore.
var TransferStoreInstance *TransferStore

// NewTransferStorage creates a new instance of Storage.
func NewTransferStorage() *TransferStore {
	if TransferStoreInstance == nil {
		TransferStoreInstance = &TransferStore{
			transfers: make(map[string]models.Transfer),
		}
	}
	return TransferStoreInstance
}

// CreateTransfer creates a new transfer in the storage.
func (s *TransferStore) CreateTransfer(transfer models.Transfer) error {
	s.transfers[transfer.ID] = transfer
	return nil
}

// ListTransfers retrieves all transfers from the storage.
func (s *TransferStore) ListTransfers() ([]models.Transfer, error) {
	var transfers []models.Transfer
	for _, transfer := range s.transfers {
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
