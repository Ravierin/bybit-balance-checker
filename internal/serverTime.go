package internal

import (
	"bybit-balance-checker/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetServerTime() (int64, error) {
	baseURL := "https://api.bybit.com/v2/public/time"
	res, err := http.Get(baseURL)
	if err != nil {
		return 0, fmt.Errorf("error getting server time: %v", err)
	}
	defer res.Body.Close()

	var response models.ServerTimeResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("error decoding server time response: %v", err)
	}

	timeNowFloat, err := strconv.ParseFloat(response.TimeNow, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing server time: %v", err)
	}

	timeNow := int64(timeNowFloat * 1000)

	return timeNow, nil
}
