package tests

import (
	"testing"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

func createAccounts(t testing.TB, account1Balance float64, account2Balance float64, accountService *services.AccountService) (models.Account, models.Account) {
	t.Helper()

	account1 := models.Account{ID: models.NewUUID(), Name: models.NewUUID(), Balance: account1Balance}
	account2 := models.Account{ID: models.NewUUID(), Name: models.NewUUID(), Balance: account2Balance}

	accountService.CreateAccount(account1)
	accountService.CreateAccount(account2)
	return account1, account2
}

func cleanup() {
	accountService := services.NewAccountService()
	accountService.DeleteAllAccounts()
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertFloats(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertIntegers(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
