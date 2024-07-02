package parser

import (
	"bybit-balance-checker/models"
	"fmt"
	"strings"
)

func ParseConfig(fileContent string) ([]models.API, error) {

	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	lines := strings.Split(fileContent, "\n")
	var accounts []models.API

	for _, line := range lines {
		if line == "" {
			continue
		}

		var currentAccount models.API
		parts := strings.Split(line, ";")

		for _, part := range parts {
			keyValue := strings.SplitN(part, "=", 2)
			if len(keyValue) != 2 {
				return nil, fmt.Errorf("invalid config line: %s", line)
			}

			key := keyValue[0]
			value := keyValue[1]

			switch key {
			case "apiKey":
				currentAccount.APIKey = value
			case "apiSecret":
				currentAccount.APISecret = value
			case "account":
				currentAccount.Account = value
			case "proxy":
				currentAccount.Proxy = value
				proxyParts := strings.SplitN(value, ":", 4)
				if len(proxyParts) == 4 {
					currentAccount.ProxyIP = proxyParts[0]
					currentAccount.ProxyPort = proxyParts[1]
					currentAccount.ProxyUsername = proxyParts[2]
					currentAccount.ProxyPassword = proxyParts[3]
				} else {
					return nil, fmt.Errorf("invalid proxy format: %s", value)
				}
			default:
				return nil, fmt.Errorf("unknown config key: %s", key)
			}
		}

		accounts = append(accounts, currentAccount)
	}

	return accounts, nil
}
