package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type API struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	Account   string `json:"account"`
}

type Balance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
	Bonus           string `json:"bonus"`
}

type Result struct {
	Account     string    `json:"account"`
	AccountType string    `json:"accountType"`
	Balance     []Balance `json:"balance"`
}

type Response struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  Result `json:"result"`
}

func generateSignature(secret, apiKey, timestamp, recvWindow, queryString string) string {
	message := timestamp + apiKey + recvWindow + queryString
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	baseURL := "https://api.bybit.com"
	endpoint := "/v5/asset/transfer/query-account-coins-balance"
	queryString := "accountType=FUND"
	method := "GET"

	fileContent, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var accounts []API
	err = json.Unmarshal(fileContent, &accounts)
	if err != nil {
		fmt.Println("Error parsing config JSON:", err)
		return
	}

	client := &http.Client{}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	recvWindow := "20000"

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for _, acc := range accounts {
		signature := generateSignature(acc.APISecret, acc.APIKey, timestamp, recvWindow, queryString)
		url := baseURL + endpoint + "?" + queryString

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		req.Header.Add("X-BAPI-API-KEY", acc.APIKey)
		req.Header.Add("X-BAPI-TIMESTAMP", timestamp)
		req.Header.Add("X-BAPI-RECV-WINDOW", recvWindow)
		req.Header.Add("X-BAPI-SIGN", signature)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		var formattedOutput string
		for _, balance := range response.Result.Balance {
			walletBalance, _ := strconv.ParseFloat(balance.WalletBalance, 64)
			if walletBalance > 1.0 {
				formattedOutput += fmt.Sprintf("- %s: %s\n", balance.Coin, balance.WalletBalance)
			}
		}

		finalOutput := fmt.Sprintf("Account: \"%s\":\n%s\n\n", acc.Account, formattedOutput)

		if _, err := outputFile.WriteString(finalOutput); err != nil {
			fmt.Println("Error writing to output file:", err)
			continue
		}
	}

	fmt.Println("The results are recorded in a file output.txt")
}
