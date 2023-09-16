// Package services is the application's business logic.
// This package is used by the controller layer to perform business logic operations.
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/mohamedsaberibrahim/basic-payment-system/models"
)

func ProcessAccounts(url string, numWorkers int) {
	// Fetch JSON data from the URL
	body, err := fetchData(url)
	if err != nil {
		panic(err)
	}

	// Start a goroutine to parse the JSON data and send the accounts to the channel
	accounts := parseAccounts(body)

	// Start a goroutine to process the accounts concurrently
	processAccountsConcurrently(accounts, numWorkers)

	fmt.Println("Finished processing accounts successfully.")
}

func fetchData(url string) ([]byte, error) {
	fmt.Println("Fetching JSON data from the URL...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching JSON data:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the JSON data
	fmt.Println("Reading JSON response...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return nil, err
	}
	return body, nil
}

func processAccountsConcurrently(accounts []models.Account, numWorkers int) {
	// Logging the number of accounts
	fmt.Println("Executing processAccountsConcurrently no. of accounts:", len(accounts))
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
	fmt.Println("Unmarshalling JSON data...")
	err := json.Unmarshal(body, &accounts)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
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
			fmt.Println("Error creating account: %s", err.Error())
		}
		fmt.Println("Processing account:", &acc, acc)
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
