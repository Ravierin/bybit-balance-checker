package api

import (
	"bybit-balance-checker/internal"
	"bybit-balance-checker/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func FetchUnifiedBalance(api models.API) ([]models.UnifiedCoin, error) {
	baseURL := "https://api.bybit.com"
	endpoint := "/v5/account/wallet-balance"
	queryString := "accountType=UNIFIED"
	method := "GET"

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	recvWindow := "20000"

	signature := internal.GenerateSignature(api.APISecret, api.APIKey, timestamp, recvWindow, queryString)
	requestURL := baseURL + endpoint + "?" + queryString

	req, err := http.NewRequest(method, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("X-BAPI-API-KEY", api.APIKey)
	req.Header.Add("X-BAPI-TIMESTAMP", timestamp)
	req.Header.Add("X-BAPI-RECV-WINDOW", recvWindow)
	req.Header.Add("X-BAPI-SIGN", signature)

	client, err := internal.CreateHTTPClient(api)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var response models.UnifiedResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	if response.RetCode != 0 {
		return nil, fmt.Errorf("API error: %s", response.RetMsg)
	}

	var unifiedBalances []models.UnifiedCoin
	for _, list := range response.Result.List {
		unifiedBalances = append(unifiedBalances, list.Coin...)
	}

	return unifiedBalances, nil
}
