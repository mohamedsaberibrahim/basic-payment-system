// Package storage is an in-memory storage of accounts and transfers.
// This is a simple implementation that uses a map to store accounts and transfers.
// This implementation is not thread-safe.
package storage

import (
	"sync"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
)

// A AccountStore struct represents an in-memory storage of accounts.
type AccountStore struct {
	// RWMutex is a reader/writer mutual exclusion lock.
	sync.RWMutex
	// accounts is a map of account IDs to accounts.
	accounts map[string]models.Account
}

// AccountStoreInstance is a singleton instance of AccountStore.
var AccountStoreInstance *AccountStore

// NewAccountStorage creates a new instance of Storage.
func NewAccountStorage() *AccountStore {
	if AccountStoreInstance == nil {
		AccountStoreInstance = &AccountStore{
			accounts: make(map[string]models.Account),
		}
	}
	return AccountStoreInstance
}

// GetAccount retrieves an account from the storage.
func (s *AccountStore) GetAccount(id string) (*models.Account, error) {
	s.RLock()
	defer s.RUnlock()
	account, ok := s.accounts[id]
	if !ok {
		return nil, models.ErrAccountNotFound
	}
	return &account, nil
}

// CreateAccount creates a new account in the storage.
func (s *AccountStore) CreateAccount(account models.Account) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.accounts[account.ID]; ok {
		return models.ErrAccountAlreadyExists
	}
	s.accounts[account.ID] = account
	return nil
}

// ListAccounts retrieves all accounts from the storage.
func (s *AccountStore) ListAccounts() ([]models.Account, error) {
	s.RLock()
	defer s.RUnlock()

	var accounts []models.Account
	for _, account := range s.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// UpdateAccount updates an account in the storage.
func (s *AccountStore) UpdateAccount(account models.Account) error {
	s.Lock()
	defer s.Unlock()
	s.accounts[account.ID] = account
	return nil
}
