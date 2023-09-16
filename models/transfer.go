// Package models is a collection of structs and methods to represent and manipulate payment accounts.
// This package is used by the main application to store and retrieve account information.
package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
)

// A Transfer represents a payment transfer with ID, from account ID, to account ID, and amount.
type Transfer struct {
	ID        string    `json:"id,omitempty"`
	From      string    `json:"fromAccountID"`
	To        string    `json:"toAccountID"`
	Amount    float64   `json:"amount,string"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// NewUUID generates a new UUID.
func NewUUID() string {
	id := uuid.New()
	return id.String()
}
