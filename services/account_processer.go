// Package services is the application's business logic.
// This package is used by the controller layer to perform business logic operations.
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
)

func ProcessAccounts(url string, numWorkers int) {
	// Logging the process of accounts
	fmt.Println("Consuming accounts...")
	// Fetch JSON data from the URL
	body, err := fetchData(url)
	if err != nil {
		fmt.Println("Consuming accounts failed with error: ", err)
		os.Exit(1)
	}

	fmt.Println("Parsing accounts...")
	// Start a goroutine to parse the JSON data and send the accounts to the channel
	accounts := parseAccounts(body)

	// Start a goroutine to process the accounts concurrently
	processAccountsConcurrently(accounts, numWorkers)

	fmt.Println("Finished processing accounts successfully...")
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the JSON data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func processAccountsConcurrently(accounts []models.Account, numWorkers int) {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Create a channel to receive the parsed accounts
	accountsChannel := make(chan models.Account)

	// Send the accounts to the channel
	wg.Add(1)
	go accountsToChannel(accounts, accountsChannel, &wg)

	// Start a fixed number of goroutines to process the accounts concurrently
	wg.Add(numWorkers)
	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go processWorker(accountsChannel, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func parseAccounts(body []byte) []models.Account {
	var accounts []models.Account
	err := json.Unmarshal(body, &accounts)
	if err != nil {
		return nil
	}
	return accounts
}

func processWorker(accountsChannel chan models.Account, wg *sync.WaitGroup) {
	defer wg.Done()
	accountService := NewAccountService()

	for acc := range accountsChannel {
		err := accountService.CreateAccount(acc)
		if err != nil {
			fmt.Println("Error creating account: ", err.Error())
		}
	}
}

func accountsToChannel(accounts []models.Account, accountsChannel chan models.Account, wg *sync.WaitGroup) {
	fmt.Println("Sending accounts to the channel...")
	defer wg.Done()
	for _, acc := range accounts {
		accountsChannel <- acc
	}
	close(accountsChannel)
	fmt.Println("Closing the channel...")
}
