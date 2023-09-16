package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mohamedsaberibrahim/basic-payment-system/controllers"
	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

func TestTransferController_GetAllTransfers(t *testing.T) {
	t.Cleanup(cleanup)

	transferService := services.NewTransferService()
	acc1, acc2 := createAccounts(t, 1000, 2000, services.NewAccountService())

	// Create test transfers
	transfer1 := models.Transfer{
		From: acc1.ID,
		To:   acc2.ID,
	}
	transfer2 := models.Transfer{
		From: acc2.ID,
		To:   acc1.ID,
	}

	transferService.CreateTransfer(transfer1, *services.NewAccountService())
	transferService.CreateTransfer(transfer2, *services.NewAccountService())

	// Create a test request
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Create a Gin context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	transferController := controllers.TransferController{}
	transferController.GetAllTransfers(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var response struct {
		Data    []models.Transfer `json:"data"`
		Message string            `json:"message"`
		Error   string            `json:"error"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Check the response data
	assertIntegers(t, len(response.Data), 2)
	assertStrings(t, response.Data[0].From, transfer1.From)
	assertStrings(t, response.Data[0].To, transfer1.To)
	assertFloats(t, response.Data[0].Amount, transfer1.Amount)
	assertStrings(t, response.Data[1].From, transfer2.From)
	assertStrings(t, response.Data[1].To, transfer2.To)
	assertFloats(t, response.Data[1].Amount, transfer2.Amount)
}

func TestTransferController_CreateTransfer(t *testing.T) {
	t.Cleanup(cleanup)

	accountService := services.NewAccountService()
	// Create test accounts
	acc1, acc2 := createAccounts(t, 500, 200, accountService)

	// Create a test transfer
	transfer := models.Transfer{
		From:   acc1.ID,
		To:     acc2.ID,
		Amount: 100,
	}

	requestBody := fmt.Sprintf(`{"fromAccountID": "%s", "toAccountID": "%s", "amount": "%f"}`, transfer.From, transfer.To, transfer.Amount)
	fmt.Println(requestBody)
	// Create a test request
	req, _ := http.NewRequest("POST", "/", io.Reader(bytes.NewBuffer([]byte(requestBody))))
	w := httptest.NewRecorder()

	// Create a Gin context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	transferController := controllers.TransferController{}
	transferController.CreateTransfer(c)

	// Check the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Parse the response body
	var response struct {
		Data    models.Transfer `json:"data"`
		Message string          `json:"message"`
		Error   string          `json:"error"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Check the response data
	assertStrings(t, response.Data.From, transfer.From)
	assertStrings(t, response.Data.To, transfer.To)
	assertFloats(t, response.Data.Amount, transfer.Amount)

	// Check that the account balances are updated correctly
	updatedBalance1, _ := accountService.GetAccount(acc1.ID)
	updatedBalance2, _ := accountService.GetAccount(acc2.ID)
	assertFloats(t, updatedBalance1.Balance, acc1.Balance-float64(transfer.Amount))
	assertFloats(t, updatedBalance2.Balance, acc2.Balance+float64(transfer.Amount))
}
