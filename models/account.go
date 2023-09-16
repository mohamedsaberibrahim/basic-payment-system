// Package models is a collection of structs and methods to represent and manipulate payment accounts.
// This package is used by the main application to store and retrieve account information.
package models

import "errors"

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")
)

// An Account represents a payment account with ID, Name, and Balance with struct tags to map the JSON data to the struct.
type Account struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance,string"`
}
