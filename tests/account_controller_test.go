package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mohamedsaberibrahim/basic-payment-system/controllers"
	"github.com/mohamedsaberibrahim/basic-payment-system/models"
	"github.com/mohamedsaberibrahim/basic-payment-system/services"
)

func TestAccountController_GetAllAccounts(t *testing.T) {
	t.Cleanup(cleanup)
	accountService := services.NewAccountService()

	// Create a test account
	account := models.Account{
		ID:      "account1",
		Balance: 100,
	}

	accountService.CreateAccount(account)

	// Create a test request
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Create a Gin context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	accountController := controllers.AccountController{}
	accountController.GetAllAccounts(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var response struct {
		Data    []models.Account `json:"data"`
		Message string           `json:"message"`
		Error   string           `json:"error"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Check the response data
	if len(response.Data) != 1 {
		t.Errorf("Expected 1 account, got %d", len(response.Data))
	}
	assertStrings(t, response.Data[0].ID, account.ID)
	assertStrings(t, response.Data[0].Name, account.Name)
	assertFloats(t, response.Data[0].Balance, account.Balance)
}
