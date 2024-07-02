package internal

import (
	"bybit-balance-checker/models"
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
)

func CreateHTTPClient(api models.API) (*http.Client, error) {
	proxyURL := fmt.Sprintf("socks5://%s:%s@%s:%s", api.ProxyUsername, api.ProxyPassword, api.ProxyIP, api.ProxyPort)
	url, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL: %v", err)
	}

	dialer, err := proxy.FromURL(url, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("failed to create proxy dialer: %v", err)
	}

	httpTransport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: httpTransport,
	}

	return client, nil
}
