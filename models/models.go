package models

type API struct {
	APIKey        string
	APISecret     string
	Account       string
	Proxy         string
	ProxyIP       string
	ProxyPort     string
	ProxyUsername string
	ProxyPassword string
}

type FundingBalance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
	Bonus           string `json:"bonus"`
}

type UnifiedCoin struct {
	Coin            string `json:"coin"`
	WalletBalance   string `json:"walletBalance"`
	AvailableAmount string `json:"availableToWithdraw"`
}

type FundingResult struct {
	AccountType string           `json:"accountType"`
	Balance     []FundingBalance `json:"balance"`
}

type UnifiedList struct {
	Coin []UnifiedCoin `json:"coin"`
}

type UnifiedResult struct {
	List []UnifiedList `json:"list"`
}

type FundingResponse struct {
	RetCode int           `json:"retCode"`
	RetMsg  string        `json:"retMsg"`
	Result  FundingResult `json:"result"`
}

type UnifiedResponse struct {
	RetCode int           `json:"retCode"`
	RetMsg  string        `json:"retMsg"`
	Result  UnifiedResult `json:"result"`
}

type ServerTimeResponse struct {
	TimeNow string `json:"time_now"`
}
