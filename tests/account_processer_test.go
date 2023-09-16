package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

func TestProcessAccounts(t *testing.T) {
	t.Cleanup(cleanup)

	// Create a test server to mock the JSON response
	t.Run("Test processing accounts", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Return a sample JSON response
			response := `[{"id": "10", "name":"account1", "balance": "100"}, {"id": "20", "name": "account2", "balance": "200"}]`
			w.Write([]byte(response))
		}))
		defer testServer.Close()

		url := testServer.URL
		numWorkers := 2

		// Test processing accounts
		services.ProcessAccounts(url, numWorkers)

		// Call the ListAccounts function to verify that the accounts are created
		accountService := services.NewAccountService()
		accounts, _ := accountService.GetAllAccounts()
		assertIntegers(t, len(accounts), 2)
	})

	t.Run("Test processing accounts with duplicate account ID", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Return a sample JSON response
			response := `[{"id": "10", "name":"account1", "balance": "100"}, {"id": "10", "name":"account1", "balance": "100"}, {"id": "20", "name": "account2", "balance": "200"}]`
			w.Write([]byte(response))
		}))
		defer testServer.Close()

		url := testServer.URL
		numWorkers := 2

		// Test processing accounts
		services.ProcessAccounts(url, numWorkers)

		// Call the ListAccounts function to verify that the accounts are created
		accountService := services.NewAccountService()
		accounts, _ := accountService.GetAllAccounts()
		assertIntegers(t, len(accounts), 2)
	})
}

func cleanup() {
	accountService := services.NewAccountService()
	accountService.DeleteAllAccounts()
}
