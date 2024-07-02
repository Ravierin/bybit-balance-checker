package main

import (
	"bybit-balance-checker/api"
	"bybit-balance-checker/parser"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

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

	for _, acc := range accounts {
		var formattedOutput string
		formattedOutput += fmt.Sprintf("Account: \"%s\":\n", acc.Account)

		fundingBalance, err := api.FetchFundingBalance(acc)
		if err != nil {
			fmt.Printf("Error fetching funding balance for account %s: %v\n", acc.Account, err)
			continue
		}
		fundingTokensOverOne := false
		formattedOutput += "  Funding:\n"
		for _, balance := range fundingBalance {
			walletBalance, _ := strconv.ParseFloat(balance.WalletBalance, 64)
			if walletBalance >= 1.0 {
				formattedOutput += fmt.Sprintf("   - %s: %s\n", balance.Coin, balance.WalletBalance)
				fundingTokensOverOne = true
			}
		}
		if !fundingTokensOverOne {
			formattedOutput += "   You homeless man(There are no coin)\n"
		}

		unifiedBalance, err := api.FetchUnifiedBalance(acc)
		if err != nil {
			fmt.Printf("Error fetching unified balance for account %s: %v\n", acc.Account, err)
			continue
		}
		unifiedTokensOverOne := false
		formattedOutput += "  Unified:\n"
		for _, balance := range unifiedBalance {
			walletBalance, _ := strconv.ParseFloat(balance.WalletBalance, 64)
			if walletBalance >= 0.0004 {
				formattedOutput += fmt.Sprintf("   - %s: %s\n", balance.Coin, balance.WalletBalance)
				unifiedTokensOverOne = true
			}
		}
		if !unifiedTokensOverOne {
			formattedOutput += "   You homeless man(There are no coin)\n"
		}

		finalOutput := fmt.Sprintf("\n%s\n", formattedOutput)

		if _, err := outputFile.WriteString(finalOutput); err != nil {
			fmt.Println("Error writing to output file:", err)
			continue
		}
	}

	fmt.Println("The results are recorded in the output.txt file.")
}
