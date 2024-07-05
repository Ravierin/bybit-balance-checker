package main

import (
	"bybit-balance-checker/api"
	"bybit-balance-checker/models"
	"bybit-balance-checker/parser"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

func fetchAndProcessBalance(acc models.API, wg *sync.WaitGroup, outputChan chan<- string) {
	defer wg.Done()

	var formattedOutput string
	formattedOutput += fmt.Sprintf("Account: \"%s\":\n", acc.Account)

	fundingBalance, err := api.FetchFundingBalance(acc)
	if err != nil {
		fmt.Printf("Error fetching funding balance for account %s: %v\n", acc.Account, err)
		return
	}

	hasSignificantBalance := false
	fundingTokensOverOne := false
	formattedOutput += "  Funding:\n"
	for _, balance := range fundingBalance {
		walletBalance, _ := strconv.ParseFloat(balance.WalletBalance, 64)
		if walletBalance >= 1.0 {
			formattedOutput += fmt.Sprintf("   - %s: %s\n", balance.Coin, balance.WalletBalance)
			fundingTokensOverOne = true
			hasSignificantBalance = true
		}
	}
	if !fundingTokensOverOne {
		formattedOutput += "   You homeless man(There are no coin < 1)\n"
	}

	unifiedBalance, err := api.FetchUnifiedBalance(acc)
	if err != nil {
		fmt.Printf("Error fetching unified balance for account %s: %v\n", acc.Account, err)
		return
	}

	unifiedTokensOverOne := false
	formattedOutput += "  Unified:\n"
	for _, balance := range unifiedBalance {
		walletBalance, _ := strconv.ParseFloat(balance.WalletBalance, 64)
		if walletBalance >= 1 { // Отвечает за кол-ва монеток
			formattedOutput += fmt.Sprintf("   - %s: %s\n", balance.Coin, balance.WalletBalance)
			unifiedTokensOverOne = true
			hasSignificantBalance = true
		}
	}
	if !unifiedTokensOverOne {
		formattedOutput += "   You homeless man(There are no coin < 1)\n"
	}

	if hasSignificantBalance {
		finalOutput := fmt.Sprintf("\n%s\n", formattedOutput)
		outputChan <- finalOutput
	}
}

func main() {
	fileContentBytes, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	accounts, err := parser.ParseConfig(string(fileContentBytes))
	if err != nil {
		fmt.Println("Error parsing config:", err)
		return
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	var wg sync.WaitGroup
	outputChan := make(chan string, len(accounts))

	for _, acc := range accounts {
		wg.Add(1)
		go fetchAndProcessBalance(acc, &wg, outputChan)
	}

	wg.Wait()
	close(outputChan)

	for output := range outputChan {
		if _, err := outputFile.WriteString(output); err != nil {
			fmt.Println("Error writing to output file:", err)
			continue
		}
	}

	fmt.Println("The results are recorded in the output.txt file.")
}
